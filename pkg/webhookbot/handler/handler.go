package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/atompi/webhookbot/pkg/options"
	"github.com/atompi/webhookbot/pkg/util"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

type Context struct {
	GinContext *gin.Context
	Options    options.BotOptions
}

type HandlerFunc func(*Context)

func NewBotHandler(handler HandlerFunc, opts options.BotOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := new(Context)
		context.GinContext = c
		context.Options = opts
		handler(context)
	}
}

func BotHandler(c *Context) {
	bodyData := make(map[string]any)
	err := c.GinContext.BindJSON(&bodyData)
	if err != nil {
		zap.L().Sugar().Errorf("failed to read request data: %v", err)
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"error": "incorrect data format"})
		return
	}

	var alertGroupData util.AlertsGroupStruct
	bodyDataByte, err := json.Marshal(bodyData)
	if err != nil {
		bodyDataByte = nil
		zap.L().Sugar().Errorf("failed to marshal body data: %v", err)
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"error": "incorrect data format"})
		return
	}
	err = json.Unmarshal(bodyDataByte, &alertGroupData)
	if err != nil {
		alertGroupData = util.AlertsGroupStruct{}
		zap.L().Sugar().Errorf("failed to unmarshal body data: %v", err)
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"error": "incorrect data format"})
		return
	}

	var postJsonData string
	if alertGroupData.Status == "resolved" {
		postJsonData, err = util.GenPostJsonData(alertGroupData, c.Options.Template.Resolved)
	} else {
		postJsonData, err = util.GenPostJsonData(alertGroupData, c.Options.Template.Alerting)
	}
	if err != nil {
		zap.L().Sugar().Errorf("failed to generate json data: %v", err)
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"error": "incorrect data format"})
		return
	}

	postData := strings.NewReader(postJsonData)
	req, _ := http.NewRequest("POST", c.Options.Webhook, postData)
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode >= http.StatusBadRequest {
		zap.L().Sugar().Errorf("failed to sent request: %v", err)
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"error": "failed to sent request"})
		return
	}

	defer res.Body.Close()

	c.GinContext.JSON(http.StatusOK, gin.H{"status": "sent success"})
}

func RootHandler(opts options.Options) gin.HandlerFunc {
	botList := []map[string]string{}
	for _, bot := range opts.Bots {
		botMap := map[string]string{
			"name":    bot.Name,
			"path":    bot.Path,
			"webhook": bot.Webhook,
		}
		botList = append(botList, botMap)
	}
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"bots": botList})
	}
}
