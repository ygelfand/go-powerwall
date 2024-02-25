package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (app *Api) voltage(c *gin.Context) {
	c.JSON(200, app.parsedVoltage())
}

func (app *Api) parsedVoltage() map[string]interface{} {
	powerDevs := map[string]interface{}{}
	if app.powerwall.Controller == nil {
		return powerDevs
	}
	for id, inv := range app.powerwall.Controller.EsCan.Bus.Pinv {
		if !inv.PINVAcMeasurements.IsMIA {
			powerDevs[fmt.Sprintf("PW%v", id+1)] = pinvPower{
				PINVVSplit1: inv.PINVAcMeasurements.PINVVSplit1,
				PINVVSplit2: inv.PINVAcMeasurements.PINVVSplit2,
				PINVFout:    inv.PINVStatus.PINVFout,
				POut:        inv.PINVStatus.PINVPout,
				VOut:        inv.PINVStatus.PINVVout,
				FOut:        inv.PINVStatus.PINVFout,
				//QOut: ??
				//IOut: ??
			}
		}
	}
	msa := app.powerwall.Controller.EsCan.Bus.Msa
	if !msa.METERZAcMeasurements.IsMIA {
		powerDevs["MSA"] = msaPower{
			PINVVSplit1: msa.METERZAcMeasurements.MeterZVl1G,
			PINVVSplit2: msa.METERZAcMeasurements.MeterZVl2G,
			PINVCSplit1: msa.METERZAcMeasurements.MeterZCtaI,
			PINVCSplit2: msa.METERZAcMeasurements.MeterZCtbI,
			PINVPSplit1: msa.METERZAcMeasurements.METERZCTAInstRealPower,
			PINVPSplit2: msa.METERZAcMeasurements.METERZCTBInstRealPower,
		}
	}
	island := app.powerwall.Controller.EsCan.Bus.Islander.ISLANDAcMeasurements
	if !island.IsMIA {
		powerDevs["ISLAND"] = islandPower{
			FreqL1Load: island.ISLANDFreqL1Load,
			FreqL1Main: island.ISLANDFreqL1Main,
			FreqL2Load: island.ISLANDFreqL2Load,
			FreqL2Main: island.ISLANDFreqL2Main,
			FreqL3Load: island.ISLANDFreqL3Load,
			FreqL3Main: island.ISLANDFreqL3Main,
			GridState:  island.ISLANDGridState,
			VL1NLoad:   island.ISLANDVL1NLoad,
			VL1NMain:   island.ISLANDVL1NMain,
			VL2NLoad:   island.ISLANDVL2NLoad,
			VL2NMain:   island.ISLANDVL2NMain,
			VL3NLoad:   island.ISLANDVL3NLoad,
			VL3NMain:   island.ISLANDVL3NMain,
		}
	}
	powerDevs["grid_status"] = app.parsedGridStatus()
	return powerDevs
}
