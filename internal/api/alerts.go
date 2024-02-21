package api

import (
	"slices"

	"github.com/gin-gonic/gin"
)

func (app *Api) alerts(c *gin.Context) {
	c.JSON(200, app.parsedAlerts())
}

func (app *Api) parsedAlerts() []string {
	alerts := []string{}
	if app.powerwall.Controller == nil {
		return alerts
	}
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
	slices.Sort(alerts)
	return slices.Compact(alerts)
}
