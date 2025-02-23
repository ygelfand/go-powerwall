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
	if app.powerwall.ControllerV2 == nil {
		return fans
	}
	for id, inv := range app.powerwall.ControllerV2.EsCan.Bus.Pvac {
		if !inv.PVACStatus.IsMIA {
			idx := pvacIndex(id)
			fans[fmt.Sprintf("A%s", idx)] = PvacFan{
				ActualRpm: inv.PVACLogging.PVAC_Fan_Speed_Actual_RPM,
				TargetRpm: inv.PVACLogging.PVAC_Fan_Speed_Target_RPM,
			}
		}
	}
	return fans
}
