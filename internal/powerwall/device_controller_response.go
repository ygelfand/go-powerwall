package powerwall

import "time"

type DeviceControllerResponse struct {
	Components struct {
		Msa []struct {
			ActiveAlerts []struct {
				Name string `json:"name,omitempty"`
			} `json:"activeAlerts,omitempty"`
			PartNumber   string `json:"partNumber,omitempty"`
			SerialNumber string `json:"serialNumber,omitempty"`
			Signals      []struct {
				BoolValue any      `json:"boolValue"`
				Name      string   `json:"name"`
				TextValue any      `json:"textValue"`
				Timestamp string   `json:"timestamp"`
				Value     *float32 `json:"value"`
			} `json:"signals"`
		} `json:"msa,omitempty"`
	} `json:"components,omitempty"`
	Control struct {
		Alerts struct {
			Active []string `json:"active,omitempty"`
		} `json:"alerts,omitempty"`
		BatteryBlocks []struct {
			Din            string `json:"din,omitempty"`
			DisableReasons any    `json:"disableReasons,omitempty"`
		} `json:"batteryBlocks,omitempty"`
		Islanding struct {
			ContactorClosed    bool   `json:"contactorClosed,omitempty"`
			CustomerIslandMode string `json:"customerIslandMode,omitempty"`
			GridOK             bool   `json:"gridOK,omitempty"`
			MicroGridOK        bool   `json:"microGridOK,omitempty"`
		} `json:"islanding,omitempty"`
		MeterAggregates []struct {
			Location   string  `json:"location,omitempty"`
			RealPowerW float64 `json:"realPowerW,omitempty"`
		} `json:"meterAggregates,omitempty"`
		PvInverters  []any `json:"pvInverters,omitempty"`
		SiteShutdown struct {
			IsShutDown bool  `json:"isShutDown,omitempty"`
			Reasons    []any `json:"reasons,omitempty"`
		} `json:"siteShutdown,omitempty"`
		SystemStatus struct {
			NominalEnergyRemainingWh int `json:"nominalEnergyRemainingWh,omitempty"`
			NominalFullPackEnergyWh  int `json:"nominalFullPackEnergyWh,omitempty"`
		} `json:"systemStatus,omitempty"`
	} `json:"control,omitempty"`
	EsCan struct {
		Bus struct {
			Islander struct {
				ISLANDAcMeasurements struct {
					ISLANDFreqL1Load float64 `json:"ISLAND_FreqL1_Load,omitempty"`
					ISLANDFreqL1Main float64 `json:"ISLAND_FreqL1_Main,omitempty"`
					ISLANDFreqL2Load float64 `json:"ISLAND_FreqL2_Load,omitempty"`
					ISLANDFreqL2Main float64 `json:"ISLAND_FreqL2_Main,omitempty"`
					ISLANDFreqL3Load float64 `json:"ISLAND_FreqL3_Load,omitempty"`
					ISLANDFreqL3Main float64 `json:"ISLAND_FreqL3_Main,omitempty"`
					ISLANDGridState  string  `json:"ISLAND_GridState,omitempty"`
					ISLANDVL1NLoad   float64 `json:"ISLAND_VL1N_Load,omitempty"`
					ISLANDVL1NMain   float64 `json:"ISLAND_VL1N_Main,omitempty"`
					ISLANDVL2NLoad   float64 `json:"ISLAND_VL2N_Load,omitempty"`
					ISLANDVL2NMain   float64 `json:"ISLAND_VL2N_Main,omitempty"`
					ISLANDVL3NLoad   float64 `json:"ISLAND_VL3N_Load,omitempty"`
					ISLANDVL3NMain   float64 `json:"ISLAND_VL3N_Main,omitempty"`
					IsComplete       bool    `json:"isComplete,omitempty"`
					IsMIA            bool    `json:"isMIA,omitempty"`
					LastRxTime       string  `json:"lastRxTime,omitempty"`
				} `json:"ISLAND_AcMeasurements,omitempty"`
				ISLANDGridConnection struct {
					ISLANDGridConnected string `json:"ISLAND_GridConnected,omitempty"`
					IsComplete          bool   `json:"isComplete,omitempty"`
				} `json:"ISLAND_GridConnection,omitempty"`
			} `json:"ISLANDER,omitempty"`
			Msa struct {
				METERZAcMeasurements struct {
					MeterZCtaI                 float64 `json:"METER_Z_CTA_I,omitempty"`
					METERZCTAInstReactivePower int     `json:"METER_Z_CTA_InstReactivePower,omitempty"`
					METERZCTAInstRealPower     int     `json:"METER_Z_CTA_InstRealPower,omitempty"`
					MeterZCtbI                 float64 `json:"METER_Z_CTB_I,omitempty"`
					METERZCTBInstReactivePower int     `json:"METER_Z_CTB_InstReactivePower,omitempty"`
					METERZCTBInstRealPower     int     `json:"METER_Z_CTB_InstRealPower,omitempty"`
					MeterZVl1G                 float64 `json:"METER_Z_VL1G,omitempty"`
					MeterZVl2G                 float64 `json:"METER_Z_VL2G,omitempty"`
					IsMIA                      bool    `json:"isMIA,omitempty"`
					LastRxTime                 string  `json:"lastRxTime,omitempty"`
				} `json:"METER_Z_AcMeasurements,omitempty"`
				MSAInfoMsg struct {
					MSAAppGitHash []int `json:"MSA_appGitHash,omitempty"`
					MSAAssemblyID int   `json:"MSA_assemblyId,omitempty"`
					IsMIA         bool  `json:"isMIA,omitempty"`
				} `json:"MSA_InfoMsg,omitempty"`
				MSAStatus struct {
					LastRxTime string `json:"lastRxTime,omitempty"`
				} `json:"MSA_Status,omitempty"`
				PackagePartNumber   string `json:"packagePartNumber,omitempty"`
				PackageSerialNumber string `json:"packageSerialNumber,omitempty"`
			} `json:"MSA,omitempty"`
			Pinv []struct {
				PINVAcMeasurements struct {
					PINVVSplit1 float64 `json:"PINV_VSplit1,omitempty"`
					PINVVSplit2 float64 `json:"PINV_VSplit2,omitempty"`
					IsMIA       bool    `json:"isMIA,omitempty"`
				} `json:"PINV_AcMeasurements,omitempty"`
				PINVPowerCapability struct {
					PINVPnom   int  `json:"PINV_Pnom,omitempty"`
					IsComplete bool `json:"isComplete,omitempty"`
					IsMIA      bool `json:"isMIA,omitempty"`
				} `json:"PINV_PowerCapability,omitempty"`
				PINVStatus struct {
					PINVFout      float64 `json:"PINV_Fout,omitempty"`
					PINVGridState string  `json:"PINV_GridState,omitempty"`
					PINVPout      float64 `json:"PINV_Pout,omitempty"`
					PINVState     string  `json:"PINV_State,omitempty"`
					PINVVout      float64 `json:"PINV_Vout,omitempty"`
					IsMIA         bool    `json:"isMIA,omitempty"`
				} `json:"PINV_Status,omitempty"`
				Alerts struct {
					Active     []string `json:"active,omitempty"`
					IsComplete bool     `json:"isComplete,omitempty"`
					IsMIA      bool     `json:"isMIA,omitempty"`
				} `json:"alerts,omitempty"`
			} `json:"PINV,omitempty"`
			Pod []struct {
				PODEnergyStatus struct {
					PODNomEnergyRemaining int  `json:"POD_nom_energy_remaining,omitempty"`
					PODNomFullPackEnergy  int  `json:"POD_nom_full_pack_energy,omitempty"`
					IsMIA                 bool `json:"isMIA,omitempty"`
				} `json:"POD_EnergyStatus,omitempty"`
				PODInfoMsg struct {
					PODAppGitHash []int `json:"POD_appGitHash,omitempty"`
				} `json:"POD_InfoMsg,omitempty"`
			} `json:"POD,omitempty"`
			Pvac []struct {
				PVACInfoMsg struct {
					PVACAppGitHash []int `json:"PVAC_appGitHash,omitempty"`
				} `json:"PVAC_InfoMsg,omitempty"`
				PVACLogging struct {
					PVAC_Fan_Speed_Actual_RPM int     `json:"PVAC_Fan_Speed_Actual_RPM"`
					PVAC_Fan_Speed_Target_RPM int     `json:"PVAC_Fan_Speed_Target_RPM"`
					PVACPVCurrentA            float64 `json:"PVAC_PVCurrent_A,omitempty"`
					PVACPVCurrentB            float64 `json:"PVAC_PVCurrent_B,omitempty"`
					PVACPVCurrentC            float64 `json:"PVAC_PVCurrent_C,omitempty"`
					PVACPVCurrentD            float64 `json:"PVAC_PVCurrent_D,omitempty"`
					PVACPVMeasuredVoltageA    float64 `json:"PVAC_PVMeasuredVoltage_A,omitempty"`
					PVACPVMeasuredVoltageB    float64 `json:"PVAC_PVMeasuredVoltage_B,omitempty"`
					PVACPVMeasuredVoltageC    float64 `json:"PVAC_PVMeasuredVoltage_C,omitempty"`
					PVACPVMeasuredVoltageD    float64 `json:"PVAC_PVMeasuredVoltage_D,omitempty"`
					PVACVL1Ground             float64 `json:"PVAC_VL1Ground,omitempty"`
					PVACVL2Ground             float64 `json:"PVAC_VL2Ground,omitempty"`
					IsMIA                     bool    `json:"isMIA,omitempty"`
				} `json:"PVAC_Logging,omitempty"`
				PVACStatus struct {
					PVACFout  float64 `json:"PVAC_Fout,omitempty"`
					PVACPout  float64 `json:"PVAC_Pout,omitempty"`
					PVACState string  `json:"PVAC_State,omitempty"`
					PVACVout  float64 `json:"PVAC_Vout,omitempty"`
					IsMIA     bool    `json:"isMIA,omitempty"`
				} `json:"PVAC_Status,omitempty"`
				Alerts struct {
					Active     []string `json:"active,omitempty"`
					IsComplete bool     `json:"isComplete,omitempty"`
					IsMIA      bool     `json:"isMIA,omitempty"`
				} `json:"alerts,omitempty"`
				PackagePartNumber      string `json:"packagePartNumber,omitempty"`
				PackageSerialNumber    string `json:"packageSerialNumber,omitempty"`
				SubPackagePartNumber   string `json:"subPackagePartNumber,omitempty"`
				SubPackageSerialNumber string `json:"subPackageSerialNumber,omitempty"`
			} `json:"PVAC,omitempty"`
			Pvs []struct {
				PVSStatus struct {
					PVSSelfTestState    string  `json:"PVS_SelfTestState,omitempty"`
					PVSState            string  `json:"PVS_State,omitempty"`
					PVSStringAConnected bool    `json:"PVS_StringA_Connected,omitempty"`
					PVSStringBConnected bool    `json:"PVS_StringB_Connected,omitempty"`
					PVSStringCConnected bool    `json:"PVS_StringC_Connected,omitempty"`
					PVSStringDConnected bool    `json:"PVS_StringD_Connected,omitempty"`
					PVSVLL              float64 `json:"PVS_vLL,omitempty"`
					IsMIA               bool    `json:"isMIA,omitempty"`
				} `json:"PVS_Status,omitempty"`
				Alerts struct {
					Active     []string `json:"active,omitempty"`
					IsComplete bool     `json:"isComplete,omitempty"`
					IsMIA      bool     `json:"isMIA,omitempty"`
				} `json:"alerts,omitempty"`
			} `json:"PVS,omitempty"`
			Sync struct {
				METERXAcMeasurements struct {
					MeterXCtaI                 float64   `json:"METER_X_CTA_I,omitempty"`
					METERXCTAInstReactivePower float64   `json:"METER_X_CTA_InstReactivePower,omitempty"`
					METERXCTAInstRealPower     float64   `json:"METER_X_CTA_InstRealPower,omitempty"`
					MeterXCtbI                 float64   `json:"METER_X_CTB_I,omitempty"`
					METERXCTBInstReactivePower float64   `json:"METER_X_CTB_InstReactivePower,omitempty"`
					METERXCTBInstRealPower     float64   `json:"METER_X_CTB_InstRealPower,omitempty"`
					MeterXCtcI                 float64   `json:"METER_X_CTC_I,omitempty"`
					METERXCTCInstReactivePower float64   `json:"METER_X_CTC_InstReactivePower,omitempty"`
					METERXCTCInstRealPower     float64   `json:"METER_X_CTC_InstRealPower,omitempty"`
					MeterXVl1N                 float64   `json:"METER_X_VL1N,omitempty"`
					MeterXVl2N                 float64   `json:"METER_X_VL2N,omitempty"`
					MeterXVl3N                 float64   `json:"METER_X_VL3N,omitempty"`
					IsComplete                 bool      `json:"isComplete,omitempty"`
					IsMIA                      bool      `json:"isMIA,omitempty"`
					LastRxTime                 time.Time `json:"lastRxTime,omitempty"`
				} `json:"METER_X_AcMeasurements,omitempty"`
				METERYAcMeasurements struct {
					MeterYCtaI                 float64   `json:"METER_Y_CTA_I,omitempty"`
					METERYCTAInstReactivePower float64   `json:"METER_Y_CTA_InstReactivePower,omitempty"`
					METERYCTAInstRealPower     float64   `json:"METER_Y_CTA_InstRealPower,omitempty"`
					MeterYCtbI                 float64   `json:"METER_Y_CTB_I,omitempty"`
					METERYCTBInstReactivePower float64   `json:"METER_Y_CTB_InstReactivePower,omitempty"`
					METERYCTBInstRealPower     float64   `json:"METER_Y_CTB_InstRealPower,omitempty"`
					MeterYCtcI                 float64   `json:"METER_Y_CTC_I,omitempty"`
					METERYCTCInstReactivePower float64   `json:"METER_Y_CTC_InstReactivePower,omitempty"`
					METERYCTCInstRealPower     float64   `json:"METER_Y_CTC_InstRealPower,omitempty"`
					MeterYVl1N                 float64   `json:"METER_Y_VL1N,omitempty"`
					MeterYVl2N                 float64   `json:"METER_Y_VL2N,omitempty"`
					MeterYVl3N                 float64   `json:"METER_Y_VL3N,omitempty"`
					IsComplete                 bool      `json:"isComplete,omitempty"`
					IsMIA                      bool      `json:"isMIA,omitempty"`
					LastRxTime                 time.Time `json:"lastRxTime,omitempty"`
				} `json:"METER_Y_AcMeasurements,omitempty"`
				SYNCInfoMsg struct {
					SYNCAppGitHash []int `json:"SYNC_appGitHash,omitempty"`
					IsMIA          bool  `json:"isMIA,omitempty"`
				} `json:"SYNC_InfoMsg,omitempty"`
				SYNCStatus struct {
					LastRxTime time.Time `json:"lastRxTime,omitempty"`
				} `json:"SYNC_Status,omitempty"`
				PackagePartNumber   string `json:"packagePartNumber,omitempty"`
				PackageSerialNumber string `json:"packageSerialNumber,omitempty"`
			} `json:"SYNC,omitempty"`
			Thc []struct {
				THCInfoMsg struct {
					THCAppGitHash []int `json:"THC_appGitHash,omitempty"`
					IsComplete    bool  `json:"isComplete,omitempty"`
					IsMIA         bool  `json:"isMIA,omitempty"`
				} `json:"THC_InfoMsg,omitempty"`
				THCLogging struct {
					THCLOGPW20EnableLineState string `json:"THC_LOG_PW_2_0_EnableLineState,omitempty"`
				} `json:"THC_Logging,omitempty"`
				PackagePartNumber   string `json:"packagePartNumber,omitempty"`
				PackageSerialNumber string `json:"packageSerialNumber,omitempty"`
			} `json:"THC,omitempty"`
		} `json:"bus,omitempty"`
		Enumeration    any `json:"enumeration,omitempty"`
		FirmwareUpdate struct {
			IsUpdating  bool  `json:"isUpdating,omitempty"`
			Msa         any   `json:"msa,omitempty"`
			Powerwalls  []any `json:"powerwalls,omitempty"`
			PvInverters any   `json:"pvInverters,omitempty"`
			Sync        any   `json:"sync,omitempty"`
		} `json:"firmwareUpdate,omitempty"`
		InverterSelfTests any `json:"inverterSelfTests,omitempty"`
		PhaseDetection    any `json:"phaseDetection,omitempty"`
	} `json:"esCan,omitempty"`
	Neurio struct {
		IsDetectingWiredMeters bool  `json:"isDetectingWiredMeters,omitempty"`
		Pairings               []any `json:"pairings,omitempty"`
		Readings               []struct {
			DataRead []struct {
				CurrentA         float64 `json:"currentA,omitempty"`
				ReactivePowerVAR float64 `json:"reactivePowerVAR,omitempty"`
				RealPowerW       float64 `json:"realPowerW,omitempty"`
				VoltageV         float64 `json:"voltageV,omitempty"`
			} `json:"dataRead,omitempty"`
			Serial    string `json:"serial,omitempty"`
			Timestamp string `json:"timestamp,omitempty"`
		} `json:"readings,omitempty"`
	} `json:"neurio,omitempty"`
	Pw3Can struct {
		FirmwareUpdate struct {
			IsUpdating bool `json:"isUpdating,omitempty"`
			Progress   any  `json:"progress,omitempty"`
		} `json:"firmwareUpdate,omitempty"`
	} `json:"pw3Can,omitempty"`
	System struct {
		SitemanagerStatus struct {
			IsRunning bool `json:"isRunning,omitempty"`
		} `json:"sitemanagerStatus,omitempty"`
		Time               string `json:"time,omitempty"`
		UpdateUrgencyCheck any    `json:"updateUrgencyCheck,omitempty"`
	} `json:"system,omitempty"`
}
