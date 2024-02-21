package api

import (
	"github.com/gin-gonic/gin"
)

func (app *Api) alerts(c *gin.Context) {
	c.JSON(200, app.parsedAlerts())
}

func (app *Api) parsedAlerts() map[string]int {
	alertMap := map[string]int{}
	if app.powerwall.Controller == nil {
		return alertMap
	}
	alerts := []string{}
	alerts = append(alerts, app.powerwall.Controller.Control.Alerts.Active...)
	for _, inv := range app.powerwall.Controller.EsCan.Bus.Pinv {
		alerts = append(alerts, inv.Alerts.Active...)
	}
	for _, inv := range app.powerwall.Controller.EsCan.Bus.Pvac {
		alerts = append(alerts, inv.Alerts.Active...)
	}
	for _, inv := range app.powerwall.Controller.EsCan.Bus.Pvs {
		alerts = append(alerts, inv.Alerts.Active...)
	}
	for _, v := range alerts {
		alertMap[v] = 1
	}
	return alertMap
}
