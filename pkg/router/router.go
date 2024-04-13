package router

import (
	metricshandler "github.com/atompi/go-kits/metrics/handler"
	"github.com/atompi/webhookbot/pkg/handler"
	"github.com/atompi/webhookbot/pkg/options"
	"github.com/gin-gonic/gin"
)

type RouterGroupFunc func(*gin.RouterGroup, options.Options)

func MetricsRouter(routerGroup *gin.RouterGroup, opts options.Options) {
	routerGroup.GET(opts.Core.Metrics.Path, metricshandler.NewPromHandler())
}

func BotRouter(routerGroup *gin.RouterGroup, opts options.Options) {
	botGroup := routerGroup.Group("bot")
	botGroup.GET("", handler.RootHandler(opts))
	for _, bot := range opts.Bots {
		botGroup.POST(bot.Path, handler.NewBotHandler(handler.BotHandler, bot))
	}
}

func Register(e *gin.Engine, opts options.Options) {
	routerGroupFuncs := []RouterGroupFunc{}

	if opts.Core.Metrics.Enable {
		e.Use(metricshandler.Handler(""))
		routerGroupFuncs = append(routerGroupFuncs, MetricsRouter)
	}

	routerGroupFuncs = append(
		routerGroupFuncs,
		BotRouter,
	)

	rootRouterGroup := e.Group("/")

	for _, routerGroup := range routerGroupFuncs {
		routerGroup(rootRouterGroup, opts)
	}
}
