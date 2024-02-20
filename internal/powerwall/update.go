package powerwall

import (
	"log"
	"time"
)

func (p *PowerwallGateway) Refresh() {
	log.Println("Refreshing state")
	res := p.RunQuery("DeviceControllerQuery", nil)
	_ = res
}

func (p *PowerwallGateway) PeriodicRefresh(interval time.Duration) {
	refreshing := false
	dataRefresh := time.NewTicker(interval)
	defer dataRefresh.Stop()
	for ; true; <-dataRefresh.C {
		if refreshing {
			return
		}
		refreshing = true
		p.Refresh()
		refreshing = false
	}
}
