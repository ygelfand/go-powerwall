package powerwall

import (
	"encoding/json"
	"log"
	"time"
)

func (p *PowerwallGateway) Refresh(force bool) {
	if !force {
		if !p.refreshSem.TryAcquire(1) {
			return
		}
		defer p.refreshSem.Release(1)
	}
	p.UpdateController()
	p.UpdateConfig()
}

func (p *PowerwallGateway) UpdateController() {
	log.Println("Refreshing Controller")
	res := p.RunQuery("DeviceControllerQuery", nil)
	if res != nil {
		err := json.Unmarshal([]byte(*res), &p.Controller)
		if err != nil {
			log.Println(err)
		}
	}
	res = p.RunQuery("DeviceControllerQueryV2", nil)
	if res != nil {
		err := json.Unmarshal([]byte(*res), &p.ControllerV2)
		if err != nil {
			log.Println(err)
		}
	}
}

func (p *PowerwallGateway) UpdateConfig() {
	log.Println("Refreshing Config")
	res := p.GetConfig()
	if res != nil {
		err := json.Unmarshal([]byte(*res), &p.Config)
		if err != nil {
			log.Println(err)
		}
	}
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
		p.Refresh(false)
		refreshing = false
	}
}
