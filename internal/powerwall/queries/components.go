package queries

var componentsQuery = &SignedQuery{
	Name:      "ComponentsQuery",
	Signature: "MIGHAkIA/m21eaB0flxf/YFl+8i1OqGoKIhYZwX8jvmgnbJ+QkMuXRHtNteCYhNtAia8XMmjQg/hf7Qib+iRfLUbxUYqp30CQWDhxwYf8dx96AtPr61tpToZuuxWhKVN7KcUHyg1kMv7WgctkVxualnLyohJKlvrhBTJadguQE3VykXq6/zvTuRK",
	DefaultParams: PointerTo(`{"pwsComponentsFilter": {
    "types": ["PW3SAF", "PW3BMS", "PVS","PVAC", "TESYNC", "TEPINV", "TETHC", "STSTSM",  "TEMSA", "TEPINV", "BAGGR", "PW3HVP"]
  },
  "pwsSignalNames": [
		"PWS_assemblyId",
    "PWS_SelfTest",
    "PWS_PeImpTestState",
    "PWS_PvIsoTestState",
    "PWS_RelaySelfTest_State",
    "PWS_MciTestState",
    "PWS_appGitHash",
		"PWS_RSD_State",
		"PWS_RSDSelfTest_State",
		"PWS_ExtSwitch_State",
    "PWS_ProdSwitch_State"
  ],
  "pchComponentsFilter": {
    "types": ["PW3SAF", "PW3BMS", "PVS","PVAC", "TESYNC", "TEPINV", "TETHC", "STSTSM",  "TEMSA", "TEPINV", "BAGGR", "PW3HVP"]
  },
  "pchSignalNames": [
    "PCH_State",
    "PCH_PvState_A",
    "PCH_PvState_B",
    "PCH_PvState_C",
    "PCH_PvState_D",
    "PCH_PvState_E",
    "PCH_PvState_F",
    "PCH_AcFrequency",
    "PCH_AcVoltageAB",
    "PCH_AcVoltageAN",
    "PCH_AcVoltageBN",
    "PCH_packagePartNumber_1_7",
    "PCH_packagePartNumber_8_14",
    "PCH_packagePartNumber_15_20",
    "PCH_packageSerialNumber_1_7",
    "PCH_packageSerialNumber_8_14",
    "PCH_PvVoltageA",
    "PCH_PvVoltageB",
    "PCH_PvVoltageC",
    "PCH_PvVoltageD",
    "PCH_PvVoltageE",
    "PCH_PvVoltageF",
    "PCH_PvCurrentA",
    "PCH_PvCurrentB",
    "PCH_PvCurrentC",
    "PCH_PvCurrentD",
    "PCH_PvCurrentE",
    "PCH_PvCurrentF",
    "PCH_BatteryPower",
    "PCH_AcRealPowerAB",
    "PCH_SlowPvPowerSum",
    "PCH_AcMode",
    "PCH_AcFrequency",
    "PCH_DcdcState_A",
    "PCH_DcdcState_B",
    "PCH_appGitHash"
  ],
  "bmsComponentsFilter": {
    "types": ["PW3SAF", "PW3BMS", "PVS","PVAC", "TESYNC", "TEPINV", "TETHC", "STSTSM",  "TEMSA", "TEPINV", "BAGGR", "PW3HVP"]
  },
  "bmsSignalNames": [
    "BMS_nominalEnergyRemaining",
    "BMS_nominalFullPackEnergy",
    "BMS_appGitHash"
  ],
  "hvpComponentsFilter": {
    "types": ["PW3SAF", "PW3BMS", "PVS","PVAC", "TESYNC", "TEPINV", "TETHC", "STSTSM",  "TEMSA", "TEPINV", "BAGGR", "PW3HVP"]
  },
  "hvpSignalNames": [
    "HVP_State",
    "HVP_appGitHash"
  ],
  "baggrComponentsFilter": {
    "types": ["PW3SAF", "PW3BMS", "PVS","PVAC", "TESYNC", "TEPINV", "TETHC", "STSTSM",  "TEMSA", "TEPINV", "BAGGR", "PW3HVP"]
  },
  "baggrSignalNames": [
    "BAGGR_State",
    "BAGGR_OperationRequest",
    "BAGGR_NumBatteriesConnected",
    "BAGGR_NumBatteriesPresent",
    "BAGGR_NumBatteriesExpected",
    "BAGGR_LOG_BattConnectionStatus0",
    "BAGGR_LOG_BattConnectionStatus1",
    "BAGGR_LOG_BattConnectionStatus2",
    "BAGGR_LOG_BattConnectionStatus3",
		"BAGGR_ExpectedFullPackEnergy",
		"BAGGR_ExpectedEnergyRemaining"
  ]
}`),
	Query: ` query ComponentsQuery (
  $pchComponentsFilter: ComponentFilter,
  $pchSignalNames: [String!],
  $pwsComponentsFilter: ComponentFilter,
  $pwsSignalNames: [String!],
  $bmsComponentsFilter: ComponentFilter,
  $bmsSignalNames: [String!],
  $hvpComponentsFilter: ComponentFilter,
  $hvpSignalNames: [String!],
  ) {
  # TODO STST-57686: Introduce GraphQL fragments to shorten
  pw3Can {
    firmwareUpdate {
      isUpdating
      progress {
         updating
         numSteps
         currentStep
         currentStepProgress
         progress
      }
    }
  }
  components {
    pws: components(filter: $pwsComponentsFilter) {
      signals(names: $pwsSignalNames) {
        name
        value
        textValue
        boolValue
        timestamp
      }
      activeAlerts {
        name
      }
    }
    pch: components(filter: $pchComponentsFilter) {
      signals(names: $pchSignalNames) {
        name
        value
        textValue
        boolValue
        timestamp
      }
      activeAlerts {
        name
      }
    }
    bms: components(filter: $bmsComponentsFilter) {
      signals(names: $bmsSignalNames) {
        name
        value
        textValue
        boolValue
        timestamp
      }
      activeAlerts {
        name
      }
    }
    hvp: components(filter: $hvpComponentsFilter) {
      signals(names: $hvpSignalNames) {
        name
        value
        textValue
        boolValue
        timestamp
      }
      activeAlerts {
        name
      }
    }
  }
}
`,
}
