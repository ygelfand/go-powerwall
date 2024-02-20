package api

import "github.com/gin-gonic/gin"

func (app *Api) Strings(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})

}
