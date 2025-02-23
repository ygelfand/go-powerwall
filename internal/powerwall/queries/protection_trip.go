package queries

var protectionTripQuery = &SignedQuery{
	Name:      "ProtectionTripQuery",
	SigKey:    2,
	Query:     `query ProtectionTripTestQuery{control{protectionTripTests{isRunning results{testType status timestamp thresholdStartValue{value unit}thresholdSetTripTime{value unit}tripThreshold{value unit}tripTime{value unit}tripTimePassCriterion{value unit}tripThresholdAccuracy{value unit}tripTimeAccuracy{value unit}measurementAtTrip{value unit}measurementAccuracy{value unit}measurementTimeAccuracy{value unit}deviation{value unit}rampStepSize{value unit}deviationPassCriterion{value unit}}}}}`,
	Signature: `MIGHAkEGsmnBC3sSrGdInJNaH9g5ErJreubbCtT+faXebk7fK+3miroCgZJennyWzsUOwh7pS7B57XyJdewCq3kS8HHteQJCATwVaH+NvdOXZKtmKk6C2mEAPnxyddU5GyYNsL8sL3z+jMceAfSoaj4XtH46ZM11l5bP578+tlvKBw/s9eRwoe+g`,
}
