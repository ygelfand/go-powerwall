package queries

var componentsQuery = &SignedQuery{
	Name:      "ComponentsQuery",
	Signature: "MIGHAkIA/m21eaB0flxf/YFl+8i1OqGoKIhYZwX8jvmgnbJ+QkMuXRHtNteCYhNtAia8XMmjQg/hf7Qib+iRfLUbxUYqp30CQWDhxwYf8dx96AtPr61tpToZuuxWhKVN7KcUHyg1kMv7WgctkVxualnLyohJKlvrhBTJadguQE3VykXq6/zvTuRK",
	DefaultParams: PointerTo(`{"hvpComponentsFilter": {"types":["PW3HVP"]},
								"pchComponentsFilter": {"types":["PCH"]},
								"bmsComponentsFilter": {"types":["PW3BMS"]},
								"pwsComponentsFilter": {"types":["PW3SAF"]}
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
