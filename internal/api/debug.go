package api

import (
	"github.com/gin-gonic/gin"
)

func (app *Api) debugConfig(c *gin.Context) {
	c.JSON(200, app.powerwall.Config)
}

func (app *Api) debugController(c *gin.Context) {
	c.JSON(200, app.powerwall.Controller)
}
