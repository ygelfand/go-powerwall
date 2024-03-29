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
		Endpoint: url,
	}

	pwr.httpClient = pwr.getClient()
	pwr.Din = *pwr.getDin()
	pwr.refreshSem = semaphore.NewWeighted(1)
	pwr.authSem = semaphore.NewWeighted(1)
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
