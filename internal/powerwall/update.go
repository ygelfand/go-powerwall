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
	now := time.Now().UnixNano()
	last := p.lastRefresh.Load()
	if now-last < 30*time.Second.Nanoseconds() {
		log.Println("Skipping refresh, last refresh was too recent.")
		return
	}

	if !p.lastRefresh.CompareAndSwap(last, now) {
		log.Println("Another thread updated the timestamp, skipping refresh.")
		return
	}
	p.UpdateController()
}

func (p *PowerwallGateway) TryRefresh() {
	p.Refresh(false)
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
}

func (p *PowerwallGateway) UpdateControllerV2() {
	log.Println("Refreshing Controller V2")
	res := p.RunQuery("DeviceControllerQueryV2", nil)
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
