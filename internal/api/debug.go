package api

import (
	"github.com/gin-gonic/gin"
)

func (app *Api) debugConfig(c *gin.Context) {
	if app.forceRefresh || c.Query("refresh") == "true" {
		app.powerwall.UpdateConfig()
	}
	c.JSON(200, app.powerwall.Config)
}

func (app *Api) debugController(c *gin.Context) {
	if app.forceRefresh || c.Query("refresh") == "true" {
		app.powerwall.UpdateController()
	}
	c.JSON(200, app.powerwall.Controller)
}
