package api

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

func (app *Api) justWidget(c *gin.Context) {
	if _, exists := c.GetQuery("widget"); exists {
		widget := &widgetWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}

		c.Writer = widget
		c.Next()
		widget.ResponseWriter.Write([]byte("<TODO />"))
	}
}

type widgetWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}
