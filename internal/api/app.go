package api

import (
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/ygelfand/go-powerwall/internal/powerwall"
)

type Api struct {
	powerwall    *powerwall.PowerwallGateway
	forceRefresh bool
	cacheStore   persist.CacheStore
	expire       time.Duration
	proxy        *httputil.ReverseProxy
}

func NewApi(p *powerwall.PowerwallGateway, forceRefresh bool) *Api {
	return &Api{
		powerwall:    p,
		forceRefresh: forceRefresh,
		cacheStore:   persist.NewMemoryStore(10 * time.Minute),
		expire:       60 * time.Second,
		proxy:        newProxy(p),
	}
}

func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(5*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(func(c *gin.Context) {
			c.String(http.StatusRequestTimeout, "timeout")
		}),
	)
}

func defaultMiddleware() gin.HandlersChain {
	return gin.HandlersChain{
		timeoutMiddleware(),
	}
}

func (api *Api) Run(listen string) {
	router := gin.Default()
	router.Use(defaultMiddleware()...)
	router.SetTrustedProxies(nil)
	base := router.Group("/api")
	{
		v1 := base.Group("/v1", api.withCache(), api.withForcedRefresh(api.powerwall.TryRefresh))
		{
			v1.GET("/strings", api.strings)
			v1.GET("/fans", api.fans)
			v1.GET("/alerts", api.alerts)
			v1.GET("/freq", api.voltage)
			v1.GET("/temps", api.temps)
			v1.GET("/pod", api.pods)
		}
		noRefresh := base.Group("/v1", api.withCache())
		{
			noRefresh.GET("/aggregates", api.powerwall.JSONReverseProxy("GET", "meters/aggregates", nil))
			noRefresh.GET("/soe", api.powerwall.JSONReverseProxy("GET", "system_status/soe", nil))
			noRefresh.GET("/status", api.powerwall.JSONReverseProxy("GET", "status", nil))
		}
		debug := router.Group("/debug")
		{
			debug.GET("/config", api.withForcedRefresh(api.powerwall.UpdateConfig), api.debugConfig)
			debug.GET("/controller", api.withForcedRefresh(api.powerwall.UpdateController), api.debugController)
			debug.GET("/controller2", api.withForcedRefresh(api.powerwall.UpdateControllerV2), api.debugControllerV2)
		}
	}

	router.GET("/", api.withCache(), api.proxyRequest, api.justWidget)

	router.NoRoute(api.proxyRequest)
	router.NoMethod(api.proxyRequest)
	router.Run(listen)
}
