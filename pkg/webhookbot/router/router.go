package router

import (
	metricshandler "github.com/atompi/go-kits/metrics/handler"
	metricsroutergroup "github.com/atompi/go-kits/metrics/router"
	"github.com/atompi/webhookbot/pkg/options"
	"github.com/atompi/webhookbot/pkg/webhookbot/handler"
	"github.com/gin-gonic/gin"
)

type RouterGroupFunc func(*gin.RouterGroup, string)

func RootRouter(routerGroup *gin.RouterGroup, opts options.Options) {
	routerGroup.GET("", handler.RootHandler(opts))
}

func Register(e *gin.Engine, opts options.Options) {
	rootRouterGroup := e.Group("/")
	routerGroups := []RouterGroupFunc{}

	if opts.Core.Metrics.Enable {
		e.Use(metricshandler.Handler(""))
		routerGroups = append(routerGroups, metricsroutergroup.MetricsRouter)
	}

	for _, routerGroup := range routerGroups {
		routerGroup(rootRouterGroup, opts.Core.Metrics.Path)
	}
}
