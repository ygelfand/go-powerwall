package api

import "github.com/ygelfand/go-powerwall/internal/powerwall"

type Api struct {
	powerwall *powerwall.PowerwallGateway
}

func NewApi(p *powerwall.PowerwallGateway) *Api {
	return &Api{powerwall: p}
}
