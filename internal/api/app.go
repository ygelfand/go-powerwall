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
		timeout.WithTimeout(500*time.Millisecond),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(func(c *gin.Context) {
			c.String(http.StatusRequestTimeout, "timeout")
		}),
	)
}

func (api *Api) Run(listen string) {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(timeoutMiddleware())
	router.SetTrustedProxies(nil)
	base := router.Group("/api")
	{
		v1 := base.Group("/v1")
		{
			v1.GET("/strings", api.withCache(), api.withForcedRefresh(api.powerwall.UpdateController), api.strings)
			v1.GET("/fans", api.withCache(), api.withForcedRefresh(api.powerwall.UpdateController), api.fans)
			v1.GET("/alerts", api.withCache(), api.withForcedRefresh(api.powerwall.UpdateController), api.alerts)
			v1.GET("/freq", api.withCache(), api.withForcedRefresh(api.powerwall.UpdateController), api.voltage)
			v1.GET("/temps", api.withCache(), api.withForcedRefresh(api.powerwall.UpdateController), api.temps)
			v1.GET("/pod", api.withCache(), api.withForcedRefresh(api.powerwall.UpdateController), api.pods)
			v1.GET("/aggregates", api.withCache(), api.powerwall.JSONReverseProxy("GET", "meters/aggregates", nil))
			v1.GET("/soe", api.withCache(), api.powerwall.JSONReverseProxy("GET", "system_status/soe", nil))
			v1.GET("/status", api.withCache(), api.powerwall.JSONReverseProxy("GET", "status", nil))
			debug := v1.Group("/debug")
			{
				debug.GET("/config", api.withForcedRefresh(api.powerwall.UpdateConfig), api.debugConfig)
				debug.GET("/controller", api.withForcedRefresh(api.powerwall.UpdateController), api.debugController)
			}
		}
	}

	router.GET("/", api.withCache(), api.proxyRequest, api.justWidget)

	router.NoRoute(api.proxyRequest)
	router.NoMethod(api.proxyRequest)
	router.Run(listen)
}
