package api

import (
	cache "github.com/chenyahui/gin-cache"
	"github.com/gin-gonic/gin"
)

func (app *Api) withCache() gin.HandlerFunc {
	return cache.Cache(app.cacheStore, app.expire, cache.WithCacheStrategyByRequest(func(c *gin.Context) (bool, cache.Strategy) {
		refresh := c.Query("refresh") == "true"
		return !refresh, cache.Strategy{
			CacheKey: c.Request.RequestURI,
		}
	}))
}

func (app *Api) withForcedRefresh(refresh func()) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		if app.forceRefresh || c.Query("refresh") == "true" {
			refresh()
		}
	}
	return gin.HandlerFunc(fn)
}
