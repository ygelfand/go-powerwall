package api

import (
	"time"

	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"github.com/ygelfand/go-powerwall/internal/powerwall"
)

type Api struct {
	powerwall    *powerwall.PowerwallGateway
	forceRefresh bool
	cacheStore   persist.CacheStore
	expire       time.Duration
}

func NewApi(p *powerwall.PowerwallGateway, forceRefresh bool) *Api {
	return &Api{powerwall: p,
		forceRefresh: forceRefresh,
		cacheStore:   persist.NewMemoryStore(10 * time.Minute),
		expire:       60 * time.Second,
	}
}

func (api *Api) Run(listen string) {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	base := router.Group("/api")
	{
		v1 := base.Group("/v1")
		{
			v1.GET("/strings", api.withForcedRefresh(api.powerwall.UpdateController), api.strings)
			v1.GET("/alerts", api.withForcedRefresh(api.powerwall.UpdateController), api.alerts)
			v1.GET("/aggregates", api.withCache(), api.powerwall.JSONReverseProxy("GET", "meters/aggregates", nil))
			v1.GET("/soe", api.withCache(), api.powerwall.JSONReverseProxy("GET", "system_status/soe", nil))
			debug := v1.Group("/debug")
			{
				debug.GET("/config", api.withForcedRefresh(api.powerwall.UpdateConfig), api.debugConfig)
				debug.GET("/controller", api.withForcedRefresh(api.powerwall.UpdateController), api.debugController)
			}
		}
	}
	router.Run(listen)
}
