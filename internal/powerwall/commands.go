package powerwall

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/sync/semaphore"
)

func NewPowerwallGateway(endpoint string, password string) *PowerwallGateway {
	url, err := url.Parse(endpoint)
	if err != nil {
		log.Fatalf("Invalid endpoint: %s, %s", url, err)
	}
	pwr := &PowerwallGateway{
		password: password,
		endpoint: url,
	}

	pwr.httpClient = pwr.getClient()
	pwr.Din = *pwr.getDin()
	pwr.refreshSem = semaphore.NewWeighted(1)
	pwr.refreshAuthToken()
	return pwr
}

func (p *PowerwallGateway) getClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
