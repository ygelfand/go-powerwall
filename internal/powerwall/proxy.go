package powerwall

import (
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

func (p *PowerwallGateway) JsonReverseProxy(method string, path string, body io.Reader) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		resp, err := p.makeApiRequest(method, path, body)
		if err != nil {
			log.Println(err)
		}
		//c.JSON(200, b64.StdEncoding.DecodeString(resp))

		c.Data(200, "application/json", resp)
	}
	return gin.HandlerFunc(fn)
}
