package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ygelfand/go-powerwall/internal/powerwall"
)

type Api struct {
	powerwall    *powerwall.PowerwallGateway
	forceRefresh bool
}

func NewApi(p *powerwall.PowerwallGateway, forceRefresh bool) *Api {
	return &Api{powerwall: p,
		forceRefresh: forceRefresh,
	}
}

func (api *Api) Run(listen string) {

	router := gin.Default()
	base := router.Group("/api")
	{
		v1 := base.Group("/v1")
		{
			v1.GET("/strings", api.strings)
			v1.GET("/alerts", api.alerts)
			debug := v1.Group("/debug")
			{
				debug.GET("/config", api.debugConfig)
				debug.GET("/controller", api.debugController)
			}
		}
	}
	router.Run(listen)
}
