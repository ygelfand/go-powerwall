package api

import "github.com/gin-gonic/gin"

func (app *Api) temps(c *gin.Context) {
	// TODO: one day
	// no way to get temps..at this time
	c.JSON(200, map[string]interface{}{})
}
