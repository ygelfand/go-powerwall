package api

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ygelfand/go-powerwall/internal/powerwall"
)

func (app *Api) strings(c *gin.Context) {
	c.JSON(200, parseStrings(app.powerwall.Controller))
}

func parseStrings(c *powerwall.DeviceControllerResponse) map[string]interface{} {
	strings := map[string]interface{}{}
	if c == nil {
		return strings
	}
	for id, inv := range c.EsCan.Bus.Pvac {
		if !inv.PVACStatus.IsMIA {
			idx := pvacIndex(id)
			strings[fmt.Sprintf("A%s_Current", idx)] = inv.PVACLogging.PVACPVCurrentA
			strings[fmt.Sprintf("B%s_Current", idx)] = inv.PVACLogging.PVACPVCurrentB
			strings[fmt.Sprintf("C%s_Current", idx)] = inv.PVACLogging.PVACPVCurrentC
			strings[fmt.Sprintf("D%s_Current", idx)] = inv.PVACLogging.PVACPVCurrentD
			strings[fmt.Sprintf("A%s_Voltage", idx)] = inv.PVACLogging.PVACPVMeasuredVoltageA
			strings[fmt.Sprintf("B%s_Voltage", idx)] = inv.PVACLogging.PVACPVMeasuredVoltageB
			strings[fmt.Sprintf("C%s_Voltage", idx)] = inv.PVACLogging.PVACPVMeasuredVoltageC
			strings[fmt.Sprintf("D%s_Voltage", idx)] = inv.PVACLogging.PVACPVMeasuredVoltageD
			strings[fmt.Sprintf("A%s_Power", idx)] = inv.PVACLogging.PVACPVMeasuredVoltageA * inv.PVACLogging.PVACPVCurrentA
			strings[fmt.Sprintf("B%s_Power", idx)] = inv.PVACLogging.PVACPVMeasuredVoltageB * inv.PVACLogging.PVACPVCurrentB
			strings[fmt.Sprintf("C%s_Power", idx)] = inv.PVACLogging.PVACPVMeasuredVoltageC * inv.PVACLogging.PVACPVCurrentC
			strings[fmt.Sprintf("D%s_Power", idx)] = inv.PVACLogging.PVACPVMeasuredVoltageD * inv.PVACLogging.PVACPVCurrentD
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
