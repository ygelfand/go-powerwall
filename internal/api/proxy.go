package api

import (
	"crypto/tls"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	"github.com/ygelfand/go-powerwall/internal/powerwall"
)

func (api *Api) proxyRequest(c *gin.Context) {
	for _, cookie := range api.powerwall.GetAuthHeaders() {
		exists, _ := c.Cookie(cookie.Name)
		if exists == "" {
			http.SetCookie(c.Writer, cookie)
		}
	}
	api.proxy.ServeHTTP(c.Writer, c.Request)
}

func newProxy(p *powerwall.PowerwallGateway) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(p.Endpoint)
	proxy.Director = func(req *http.Request) {
		//req.Header = c.Request.Header
		req.Host = p.Endpoint.Host
		req.URL.Scheme = p.Endpoint.Scheme
		req.URL.Host = p.Endpoint.Host
		for _, cookie := range p.GetAuthHeaders() {
			req.AddCookie(cookie)
		}
	}
	proxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return proxy
}
