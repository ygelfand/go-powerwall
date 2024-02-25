package api

import (
	"encoding/json"
	"log"
)

var gridmap = map[string]int{
	"SystemGridConnected":      1,
	"SystemIslandedActive":     0,
	"SystemTransitionToGrid":   -1,
	"SystemTransitionToIsland": -1,
	"SystemIslandedReady":      -1,
	"SystemMicroGridFaulted":   0,
	"SystemWaitForUser":        0,
}

func (app *Api) parsedGridStatus() int {
	status, err := app.powerwall.MakeAPIRequest("GET", "system_status/grid_status", nil)
	if err != nil {
		log.Printf("Failed to get grid status: %s", err)
		return 0
	}
	statusResp := &gridResponse{}
	err = json.Unmarshal(status, statusResp)
	if err != nil {
		log.Printf("Failed to parse grid status: %s", err)
		return 0
	}
	return gridmap[statusResp.GridStatus]
}

type gridResponse struct {
	GridStatus        string `json:"grid_status"`
	GridServiceActoun bool   `json:"grid_services_active"`
}
