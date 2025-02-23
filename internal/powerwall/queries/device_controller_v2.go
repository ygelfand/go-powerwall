package queries

var deviceControllerQueryV2 = &SignedQuery{
	Name:      "DeviceControllerQueryV2",
	SigKey:    2,
	Signature: "MIGHAkFt638Fpr6LmaBsuw+voQMcsRxqDQWVx3FLuBhbM43p8DpzuVK2e8IsrTa4V5of91pdaqt1NcAeCv/boHO5/N1GiQJCAL4ZCVs8R1qerAvQ6Ota6xrfOB7FfyO5CoIVeAlyFlw06VJqInTnryBIRXU8TL02BpStex3FVWbA6M6zFMYOK5AY",
	DefaultParams: PointerTo(`{"msaComp":{"types" :["PVS","PVAC", "TESYNC", "TEPINV", "TETHC", "STSTSM",  "TEMSA", "TEPINV" ]},
	"msaSignals":[
	"MSA_pcbaId",
	"MSA_usageId",
	"MSA_appGitHash",
	"MSA_HeatingRateOccurred",
	"THC_AmbientTemp",
  "MSA_HeatingRateLockoutEnabled",
  "MSA_ExtPower_Status",
  "PINV_ACStageFault",
  "THC_State",
  "THC_Logging",
	"METER_Z_CTA_InstRealPower",
	"METER_Z_CTA_InstReactivePower",
	"METER_Z_CTA_I",
	"METER_Z_VL1G",
	"METER_Z_CTB_InstRealPower",
	"METER_Z_CTB_InstReactivePower",
	"METER_Z_CTB_I",
	"METER_Z_VL2G"]}
	`),
	Query: `query DeviceControllerQuery{control{systemStatus{nominalFullPackEnergyWh nominalEnergyRemainingWh}islanding{customerIslandMode contactorClosed microGridOK gridOK disableReasons}meterAggregates{location realPowerW}alerts{active}siteShutdown{isShutDown reasons}batteryBlocks{din disableReasons}pvInverters{din disableReasons}}system{time supportMode{remoteService{isEnabled expiryTime sessionId}}sitemanagerStatus{isRunning}}neurio{isDetectingWiredMeters readings{firmwareVersion serial dataRead{voltageV realPowerW reactivePowerVAR currentA}timestamp}pairings{serial shortId status errors macAddress hostname isWired modbusPort modbusId lastUpdateTimestamp}}teslaRemoteMeter{meters{din reading{timestamp firmwareVersion ctReadings{voltageV realPowerW reactivePowerVAR energyExportedWs energyImportedWs currentA}}firmwareUpdate{updating numSteps currentStep currentStepProgress progress}}detectedWired{din serialPort}}pw3Can{firmwareUpdate{isUpdating progress{updating numSteps currentStep currentStepProgress progress}}enumeration{inProgress}}esCan{bus{PVAC{packagePartNumber packageSerialNumber subPackagePartNumber subPackageSerialNumber PVAC_Status{isMIA PVAC_Pout PVAC_State PVAC_Vout PVAC_Fout}PVAC_ControlMeasurements{PVAC_FanSelfTestState}PVAC_InfoMsg{PVAC_appGitHash}PVAC_Logging{isMIA PVAC_PVCurrent_A PVAC_PVCurrent_B PVAC_PVCurrent_C PVAC_PVCurrent_D PVAC_PVMeasuredVoltage_A PVAC_PVMeasuredVoltage_B PVAC_PVMeasuredVoltage_C PVAC_PVMeasuredVoltage_D PVAC_VL1Ground PVAC_VL2Ground PVAC_Fan_Speed_Actual_RPM PVAC_Fan_Speed_Target_RPM}alerts{isComplete isMIA active}}PINV{PINV_Status{isMIA PINV_Fout PINV_Pout PINV_Vout PINV_State PINV_GridState}PINV_AcMeasurements{isMIA PINV_VSplit1 PINV_VSplit2}PINV_PowerCapability{isComplete isMIA PINV_Pnom}alerts{isComplete isMIA active}}PVS{PVS_Status{isMIA PVS_State PVS_vLL PVS_StringA_Connected PVS_StringB_Connected PVS_StringC_Connected PVS_StringD_Connected PVS_SelfTestState}PVS_Logging{PVS_numStringsLockoutBits PVS_sbsComplete}alerts{isComplete isMIA active}}THC{packagePartNumber packageSerialNumber THC_InfoMsg{isComplete isMIA THC_appGitHash}THC_Logging{THC_LOG_PW_2_0_EnableLineState}}POD{POD_EnergyStatus{isMIA POD_nom_energy_remaining POD_nom_full_pack_energy}POD_InfoMsg{POD_appGitHash}}SYNC{packagePartNumber packageSerialNumber SYNC_InfoMsg{isMIA SYNC_appGitHash SYNC_assemblyId}METER_X_AcMeasurements{isMIA isComplete METER_X_CTA_InstRealPower METER_X_CTA_InstReactivePower METER_X_CTA_I METER_X_VL1N METER_X_CTB_InstRealPower METER_X_CTB_InstReactivePower METER_X_CTB_I METER_X_VL2N METER_X_CTC_InstRealPower METER_X_CTC_InstReactivePower METER_X_CTC_I METER_X_VL3N}METER_Y_AcMeasurements{isMIA isComplete METER_Y_CTA_InstRealPower METER_Y_CTA_InstReactivePower METER_Y_CTA_I METER_Y_VL1N METER_Y_CTB_InstRealPower METER_Y_CTB_InstReactivePower METER_Y_CTB_I METER_Y_VL2N METER_Y_CTC_InstRealPower METER_Y_CTC_InstReactivePower METER_Y_CTC_I METER_Y_VL3N}}ISLANDER{ISLAND_GridConnection{ISLAND_GridConnected isComplete}ISLAND_AcMeasurements{ISLAND_VL1N_Main ISLAND_FreqL1_Main ISLAND_VL2N_Main ISLAND_FreqL2_Main ISLAND_VL3N_Main ISLAND_FreqL3_Main ISLAND_VL1N_Load ISLAND_FreqL1_Load ISLAND_VL2N_Load ISLAND_FreqL2_Load ISLAND_VL3N_Load ISLAND_FreqL3_Load ISLAND_GridState isComplete isMIA}}}enumeration{inProgress numACPW numPVI}firmwareUpdate{isUpdating powerwalls{updating numSteps currentStep currentStepProgress progress}msa{updating numSteps currentStep currentStepProgress progress}msa1{updating numSteps currentStep currentStepProgress progress}sync{updating numSteps currentStep currentStepProgress progress}pvInverters{updating numSteps currentStep currentStepProgress progress}}phaseDetection{inProgress lastUpdateTimestamp powerwalls{din progress phase}}}components{msa:components(filter:{types:[TEMSA]}){partNumber serialNumber signals(names:["MSA_pcbaId" "MSA_usageId" "MSA_appGitHash" "MSA_HeatingRateOccurred" "METER_Z_CTA_InstRealPower" "METER_Z_CTA_InstReactivePower" "METER_Z_CTA_I" "METER_Z_VL1G" "METER_Z_CTB_InstRealPower" "METER_Z_CTB_InstReactivePower" "METER_Z_CTB_I" "METER_Z_VL2G"]){name value textValue boolValue timestamp}activeAlerts{name}}}}`,
}
