package handler

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"strings"

	"gitee.com/autom-studio/alert-feishu/internal/config"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

func genPostJsonData(structData config.AlertsGroupStruct, tmplFilePath string) (jsonData string, err error) {
	t, err := template.ParseFiles(tmplFilePath)
	if err != nil {
		zap.L().Sugar().Errorf("parse %s failed: %v", tmplFilePath, err)
		return
	}
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, structData)
	if err != nil {
		zap.L().Sugar().Errorf("generate json data failed: %v", err)
		return
	}
	jsonData = buffer.String()
	return
}

type Context struct {
	GinContext *gin.Context
	Config     config.RootConfigStruct
}

type HandlerFunc func(*Context)

func NewHandler(handler HandlerFunc, config config.RootConfigStruct) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := new(Context)
		context.GinContext = c
		context.Config = config
		handler(context)
	}
}

func Handler(c *Context) {
	bodyData := make(map[string]any)
	err := c.GinContext.BindJSON(&bodyData)
	if err != nil {
		zap.L().Sugar().Errorf("failed to read request data: %v", err)
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"error": "incorrect data format"})
		return
	}

	var alertGroupData config.AlertsGroupStruct
	bodyDataByte, err := json.Marshal(bodyData)
	if err != nil {
		bodyDataByte = nil
		zap.L().Sugar().Errorf("failed to marshal body data: %v", err)
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"error": "incorrect data format"})
		return
	}
	err = json.Unmarshal(bodyDataByte, &alertGroupData)
	if err != nil {
		alertGroupData = config.AlertsGroupStruct{}
		zap.L().Sugar().Errorf("failed to unmarshal body data: %v", err)
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"error": "incorrect data format"})
		return
	}

	var postJsonData string
	if alertGroupData.Status == "resolved" {
		postJsonData, err = genPostJsonData(alertGroupData, c.Config.Feishu.ResolvedMsgTmpl)
	} else {
		postJsonData, err = genPostJsonData(alertGroupData, c.Config.Feishu.AlertMsgTmpl)
	}
	if err != nil {
		zap.L().Sugar().Errorf("failed to generate json data: %v", err)
		c.GinContext.JSON(http.StatusBadRequest, gin.H{"error": "incorrect data format"})
		return
	}

	postData := strings.NewReader(postJsonData)
	req, _ := http.NewRequest("POST", c.Config.Feishu.WebhookUrl, postData)
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
