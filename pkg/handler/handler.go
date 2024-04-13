package handler

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"net/http"

	"github.com/atompi/webhookbot/pkg/options"
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
	bodyData, err := io.ReadAll(c.GinContext.Request.Body)
	if err != nil {
		zap.L().Sugar().Errorf("failed to read request data: %v", err)
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"error": "read body failed"})
		return
	}

	m := map[string]interface{}{}
	err = json.Unmarshal(bodyData, &m)
	if err != nil {
		zap.L().Sugar().Errorf("failed to unmarshal body data: %v", err)
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"error": "incorrect data format"})
		return
	}

	var tmplFilePath string
	if m["status"] == "resolved" {
		tmplFilePath = c.Options.Template.Resolved
	} else {
		tmplFilePath = c.Options.Template.Alerting
	}
	t, err := template.ParseFiles(tmplFilePath)
	if err != nil {
		zap.L().Sugar().Errorf("failed to parse template: %v", err)
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"error": "parse template failed"})
		return
	}

	postData := new(bytes.Buffer)
	err = t.Execute(postData, m)
	if err != nil {
		zap.L().Sugar().Errorf("failed to generate post body: %v", err)
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"error": "generate post body failed"})
		return
	}

	req, _ := http.NewRequest("POST", c.Options.Webhook, postData)
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode >= http.StatusBadRequest {
		zap.L().Sugar().Errorf("failed to send request: %v", err)
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"error": "send request failed"})
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
