package router

import (
	"gitee.com/autom-studio/alert-feishu/api/webhook/handler"
	"gitee.com/autom-studio/alert-feishu/internal/config"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine, config config.RootConfigStruct) {
	engine.POST(config.Main.WebhookPath, handler.NewHandler(handler.Handler, config))
}
