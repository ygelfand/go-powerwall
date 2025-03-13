package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (app *Api) temps(c *gin.Context) {
	temps := map[string]*float32{}
	if app.powerwall.Controller != nil {
		index := 0
		for _, msa := range app.powerwall.Controller.Components.Msa {
			for _, signal := range msa.Signals {
				if signal.Name == "THC_AmbientTemp" && signal.Value != nil {
					index++
					temps[fmt.Sprintf("PW%v_temp", index)] = signal.Value
				}
			}
		}
	}
	c.JSON(200, temps)
}
