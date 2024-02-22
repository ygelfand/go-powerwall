package powerwall

import (
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

func (p *PowerwallGateway) JSONReverseProxy(method string, path string, body io.Reader) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		resp, err := p.makeAPIRequest(method, path, body)
		if err != nil {
			log.Println(err)
		}
		c.Data(200, "application/json", resp)
	}
	return gin.HandlerFunc(fn)
}
