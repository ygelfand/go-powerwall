package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (app *Api) pods(c *gin.Context) {
	c.JSON(200, app.parsedPods())
}

func (app *Api) parsedPods() map[string]podResponse {
	pods := map[string]podResponse{}
	if app.powerwall.Controller == nil {
		return pods
	}
	for id, pod := range app.powerwall.Controller.EsCan.Bus.Pod {
		if !pod.PODEnergyStatus.IsMIA {
			pod := podResponse{
				PODNomEnergyRemaining: pod.PODEnergyStatus.PODNomEnergyRemaining,
				PODNomFullPackEnergy:  pod.PODEnergyStatus.PODNomFullPackEnergy,
			}
			if len(app.powerwall.Controller.EsCan.Bus.Thc) > id {
				thc := app.powerwall.Controller.EsCan.Bus.Thc[id]
				if !thc.THCInfoMsg.IsMIA {
					pod.PackagePartNumber = thc.PackagePartNumber
					pod.PackageSerialNumber = thc.PackageSerialNumber
				}
			}
			if len(app.powerwall.Controller.EsCan.Bus.Pinv) > id {
				inv := app.powerwall.Controller.EsCan.Bus.Pinv[id]
				if !inv.PINVStatus.IsMIA {
					pod.PinvGridState = inv.PINVStatus.PINVGridState
					pod.PinvState = inv.PINVStatus.PINVState
					pod.POut = inv.PINVStatus.PINVPout
					pod.VOut = inv.PINVStatus.PINVVout
					pod.FOut = inv.PINVStatus.PINVFout
				}
			}
			pods[fmt.Sprintf("PW%v", id+1)] = pod
		}
	}
	return pods
}
