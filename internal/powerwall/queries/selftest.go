package queries

var selfTestQuery = &SignedQuery{
	Name:      "SelfTestQuery",
	SigKey:    2,
	Query:     `query PinvSelfTestQuery{esCan{inverterSelfTests{isRunning isCanceled pinvSelfTestsResults{din overall{status test summary setMagnitude setTime tripMagnitude tripTime accuracyMagnitude accuracyTime currentMagnitude timestamp lastError}testResults{status test summary setMagnitude setTime tripMagnitude tripTime accuracyMagnitude accuracyTime currentMagnitude timestamp lastError}}}}}`,
	Signature: "MIGIAkIAqdGpzuLPflDSPGuVyrqsOU4DnyhhMUHRwMlH1QYTxBep0IU0smVWNJ6L3F34k2LBbugQjm2kwGabgS0e2kjyXs8CQgFDy0GMLWZb1MoaWPAr4NQxwKzrElwWZO19a1NiwiujR/la5OuBQ4lg8wOq04qzWw7Ti/+MNsD3Eacp6uL4OTwi5g==",
}
