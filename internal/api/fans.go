package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (app *Api) fans(c *gin.Context) {
	c.JSON(200, app.parsedFans())
}

func (app *Api) parsedFans() map[string]PvacFan {
	fans := map[string]PvacFan{}
	if app.powerwall.Controller == nil {
		return fans
	}
	var indexActual, indexTarget int
	for _, msa := range app.powerwall.Controller.Components.Msa {
		for _, signal := range msa.Signals {
			if signal.Name == "PVAC_Fan_Speed_Actual_RPM" && signal.Value != nil {
				key := fmt.Sprintf("A%s", pvacIndex(indexActual))
				fan, exists := fans[key]
				if !exists {
					fan = PvacFan{} // Initialize if not present
				}
				fan.ActualRpm = int(*signal.Value)
				fans[key] = fan
				indexActual++
			} else if signal.Name == "PVAC_Fan_Speed_Target_RPM" && signal.Value != nil {
				key := fmt.Sprintf("A%s", pvacIndex(indexTarget))
				fan, exists := fans[key]
				if !exists {
					fan = PvacFan{} // Initialize if not present
				}
				fan.TargetRpm = int(*signal.Value)
				fans[key] = fan
				indexTarget++
			}
		}
	}
	return fans
}
