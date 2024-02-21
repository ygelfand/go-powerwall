package api

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *Api) strings(c *gin.Context) {
	c.JSON(200, app.parsedStrings())
}

func (app *Api) parsedStrings() map[string]PvacString {
	strings := map[string]PvacString{}
	if app.powerwall.Controller == nil {
		return strings
	}
	for id, inv := range app.powerwall.Controller.EsCan.Bus.Pvac {
		if !inv.PVACStatus.IsMIA {
			idx := pvacIndex(id)
			strings[fmt.Sprintf("A%s", idx)] = PvacString{
				Current:   inv.PVACLogging.PVACPVCurrentA,
				Voltage:   inv.PVACLogging.PVACPVMeasuredVoltageA,
				State:     inv.PVACStatus.PVACState,
				Connected: app.powerwall.Controller.EsCan.Bus.Pvs[id].PVSStatus.PVSStringAConnected,
			}
			strings[fmt.Sprintf("B%s", idx)] = PvacString{
				Current:   inv.PVACLogging.PVACPVCurrentB,
				Voltage:   inv.PVACLogging.PVACPVMeasuredVoltageB,
				State:     inv.PVACStatus.PVACState,
				Connected: app.powerwall.Controller.EsCan.Bus.Pvs[id].PVSStatus.PVSStringBConnected,
			}
			strings[fmt.Sprintf("C%s", idx)] = PvacString{
				Current:   inv.PVACLogging.PVACPVCurrentC,
				Voltage:   inv.PVACLogging.PVACPVMeasuredVoltageC,
				State:     inv.PVACStatus.PVACState,
				Connected: app.powerwall.Controller.EsCan.Bus.Pvs[id].PVSStatus.PVSStringCConnected,
			}
			strings[fmt.Sprintf("D%s", idx)] = PvacString{
				Current:   inv.PVACLogging.PVACPVCurrentD,
				Voltage:   inv.PVACLogging.PVACPVMeasuredVoltageD,
				State:     inv.PVACStatus.PVACState,
				Connected: app.powerwall.Controller.EsCan.Bus.Pvs[id].PVSStatus.PVSStringDConnected,
			}
		}
	}
	return strings
}

func pvacIndex(num int) string {
	if num == 0 {
		return ""
	}
	return strconv.Itoa(num)
}
