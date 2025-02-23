package queries

var ie2030query = &SignedQuery{
	Name:      "IE2030Query",
	Query:     `query IEEE20305Query{ieee20305{longFormDeviceID polledResources{url name pollRateSeconds lastPolledTimestamp}controls{defaultControl{mRID setGradW opModEnergize opModMaxLimW opModImpLimW opModExpLimW opModGenLimW opModLoadLimW}activeControls{opModEnergize opModMaxLimW opModImpLimW opModExpLimW opModGenLimW opModLoadLimW}}registration{dateTimeRegistered pin}}}`,
	SigKey:    2,
	Signature: `MIGIAkIB7YgUNI6PnfOT/BSWWTWLEqow/q0EYHReDOHOfZ5KYlG2scl32/5QqGVC5nQ0tkJftgRgOM9bBl7TB2mNpvuZkNECQgDJTRVYQHAIieT1vsdlsNN8u6GPpdMzYeheEdHQll0KaCVoTf5y7nVOxkJ4ru7+JAz4q4nhGQYYLmNTBC6eFKeJ2A==`,
}
