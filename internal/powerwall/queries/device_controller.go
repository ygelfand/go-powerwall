package queries

var DeviceControllerQuery = &SignedQuery{
	Name:      "DeviceControllerQuery",
	Signature: "MIGHAkIBlHJ++stgIqaruiXRFbni9KM+Gofpc/b/NFo9SkuuSgjR5cyUYEQ2A9AoZbq7Ps6kg5HCFKJZbtgRP/ph72vL1XYCQUh/5IdFZDYhmeFGs42AIE2hK+5KDhwLLWE98qD2u8z4jdRF6AYHwGpR08B9a4gO1cpOFGmBFrQ/M9BDrG2HrpG4",
	Query: ` query DeviceControllerQuery {
  control {
    systemStatus {
        nominalFullPackEnergyWh
        nominalEnergyRemainingWh
    }
    islanding {
        customerIslandMode
        contactorClosed
        microGridOK
        gridOK
    }
    meterAggregates {
      location
      realPowerW
    }
    alerts {
      active
    },
    siteShutdown {
      isShutDown
      reasons
    }
    batteryBlocks {
      din
      disableReasons
    }
    pvInverters {
      din
      disableReasons
    }
  }
  system {
    time
    supportMode {
      remoteService {
        isEnabled
        expiryTime
      }
    }
    sitemanagerStatus {
      isRunning
    }
    updateUrgencyCheck  {
      urgency
      version {
        version
        gitHash
      }
      timestamp
    }
  }
  neurio {
    isDetectingWiredMeters
    readings {
      firmwareVersion
      serial
      dataRead {
        voltageV
        realPowerW
        reactivePowerVAR
        currentA
      }
      timestamp
    }
    pairings {
      serial
      shortId
      status
      errors
      macAddress
      hostname
      isWired
      modbusPort
      modbusId
      lastUpdateTimestamp
    }
  }
  teslaRemoteMeter {
    meters {
      din
      reading {
        timestamp
        firmwareVersion
        ctReadings {
          voltageV
          realPowerW
          reactivePowerVAR
          energyExportedWs
          energyImportedWs
          currentA
        }
      }
      firmwareUpdate {
        updating
        numSteps
        currentStep
        currentStepProgress
        progress
      }
    }
  }
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
  esCan {
    bus {
      PVAC {
        packagePartNumber
        packageSerialNumber
        subPackagePartNumber
        subPackageSerialNumber
        PVAC_Status {
          isMIA
          PVAC_Pout
          PVAC_State
          PVAC_Vout
          PVAC_Fout
        }
        PVAC_InfoMsg {
          PVAC_appGitHash
        }
        PVAC_Logging {
          isMIA
          PVAC_PVCurrent_A
          PVAC_PVCurrent_B
          PVAC_PVCurrent_C
          PVAC_PVCurrent_D
          PVAC_PVMeasuredVoltage_A
          PVAC_PVMeasuredVoltage_B
          PVAC_PVMeasuredVoltage_C
          PVAC_PVMeasuredVoltage_D
          PVAC_VL1Ground
          PVAC_VL2Ground
        }
        alerts {
          isComplete
          isMIA
          active
        }
      }
      PINV {
        PINV_Status {
          isMIA
          PINV_Fout
          PINV_Pout
          PINV_Vout
          PINV_State
          PINV_GridState
        }
        PINV_AcMeasurements {
          isMIA
          PINV_VSplit1
          PINV_VSplit2
        }
        PINV_PowerCapability {
          isComplete
          isMIA
          PINV_Pnom
        }
        alerts {
          isComplete
          isMIA
          active
        }
      }
      PVS {
        PVS_Status {
          isMIA
          PVS_State
          PVS_vLL
          PVS_StringA_Connected
          PVS_StringB_Connected
          PVS_StringC_Connected
          PVS_StringD_Connected
          PVS_SelfTestState
        }
        PVS_Logging {
          PVS_numStringsLockoutBits
          PVS_sbsComplete
        }
        alerts {
          isComplete
          isMIA
          active
        }
      }
      THC {
        packagePartNumber
        packageSerialNumber
        THC_InfoMsg {
          isComplete
          isMIA
          THC_appGitHash
        }
        THC_Logging {
          THC_LOG_PW_2_0_EnableLineState
        }
      }
      POD {
        POD_EnergyStatus {
          isMIA
          POD_nom_energy_remaining
          POD_nom_full_pack_energy
        }
        POD_InfoMsg {
            POD_appGitHash
        }
      }
      MSA {
        packagePartNumber
        packageSerialNumber
        MSA_InfoMsg {
          isMIA
          MSA_pcbaId
          MSA_usageId
          MSA_appGitHash
        }
        METER_Z_AcMeasurements {
          isMIA
          lastRxTime
          METER_Z_CTA_InstRealPower
          METER_Z_CTA_InstReactivePower
          METER_Z_CTA_I
          METER_Z_VL1G
          METER_Z_CTB_InstRealPower
          METER_Z_CTB_InstReactivePower
          METER_Z_CTB_I
          METER_Z_VL2G
        }
        MSA_Status {
          lastRxTime
        }
        MSA_Debug {
          MSA_HeatingRateOccurred
        }
        alerts {
          active
        }
      }
      SYNC {
        packagePartNumber
        packageSerialNumber
        SYNC_InfoMsg {
          isMIA
          SYNC_appGitHash
        }
        METER_X_AcMeasurements {
          isMIA
          isComplete
          lastRxTime
          METER_X_CTA_InstRealPower
          METER_X_CTA_InstReactivePower
          METER_X_CTA_I
          METER_X_VL1N
          METER_X_CTB_InstRealPower
          METER_X_CTB_InstReactivePower
          METER_X_CTB_I
          METER_X_VL2N
          METER_X_CTC_InstRealPower
          METER_X_CTC_InstReactivePower
          METER_X_CTC_I
          METER_X_VL3N
        }
        METER_Y_AcMeasurements {
          isMIA
          isComplete
          lastRxTime
          METER_Y_CTA_InstRealPower
          METER_Y_CTA_InstReactivePower
          METER_Y_CTA_I
          METER_Y_VL1N
          METER_Y_CTB_InstRealPower
          METER_Y_CTB_InstReactivePower
          METER_Y_CTB_I
          METER_Y_VL2N
          METER_Y_CTC_InstRealPower
          METER_Y_CTC_InstReactivePower
          METER_Y_CTC_I
          METER_Y_VL3N
        }
        SYNC_Status {
          lastRxTime
        }
      }
      ISLANDER {
        ISLAND_GridConnection {
          ISLAND_GridConnected
          isComplete
        }
        ISLAND_AcMeasurements {
          ISLAND_VL1N_Main
          ISLAND_FreqL1_Main
          ISLAND_VL2N_Main
          ISLAND_FreqL2_Main
          ISLAND_VL3N_Main
          ISLAND_FreqL3_Main
          ISLAND_VL1N_Load
          ISLAND_FreqL1_Load
          ISLAND_VL2N_Load
          ISLAND_FreqL2_Load
          ISLAND_VL3N_Load
          ISLAND_FreqL3_Load
          ISLAND_GridState
          lastRxTime
          isComplete
          isMIA
        }
      }
    }
    enumeration {
      inProgress
      numACPW
      numPVI
    }
    firmwareUpdate {
      isUpdating
      powerwalls {
        updating
        numSteps
        currentStep
        currentStepProgress
        progress
      }
      msa {
        updating
        numSteps
        currentStep
        currentStepProgress
        progress
      }
      sync {
        updating
        numSteps
        currentStep
        currentStepProgress
        progress
      }
      pvInverters {
        updating
        numSteps
        currentStep
        currentStepProgress
        progress
      }
    }
    phaseDetection {
      inProgress
      lastUpdateTimestamp
      powerwalls {
        din
        progress
        phase
      }
    }
    inverterSelfTests {
      isRunning
      isCanceled
      pinvSelfTestsResults {
        din
        overall {
          status
          test
          summary
          setMagnitude
          setTime
          tripMagnitude
          tripTime
          accuracyMagnitude
          accuracyTime
          currentMagnitude
          timestamp
          lastError
        }
        testResults {
          status
          test
          summary
          setMagnitude
          setTime
          tripMagnitude
          tripTime
          accuracyMagnitude
          accuracyTime
          currentMagnitude
          timestamp
          lastError
        }
      }
    }
  }
}
`,
}
