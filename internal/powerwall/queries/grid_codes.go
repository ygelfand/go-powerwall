package queries

var gridCodesQuery = &SignedQuery{
	Name:      "GridCodesQuery",
	SigKey:    2,
	Signature: "MIGGAkE12gu3caiB/rFRInuS5c2/YMA18295UXlixcf4E4CHMi0rqVju7reke7uuWV9W4eSfk60cSfRRTkrsQ5RpVcp2uAJBMQZMKRNATdgKn00Occ1Cc4Rj5JpQ9i+XgwKMZUkBPrWAFUTtyI3X5TdXso81T88pXKL9GbpJf3rup1SNxlw16dc=",
	Query:     `query GridCodesQuery{system{gridCodes}}`,
}

var gridCodesDetailsQuery = &SignedQuery{
	Name:      "GridCodesDetailsQuery",
	Signature: "MIGIAkIBN96jDx4iIwoglX1SEs/CRHgDJNbdlk+qVYIJHNK0I4fuQA7tOE3trH93RhlTf/EV7VEw6twafp26AfJoLAfHQHUCQgD1/Eiyl/31RAHmi0UX6lIvUVDEhBRsmtmWMcD7hUXjc2HqAI01pl56/ShqQPBlrluy5cxx9/aJMxkoVNGO6x5IbQ==",
	Query: `query GridCodeDetailsQuery($gridCode: String!, $pointNames: [String!]) {
  system {
    gridCodeSettings(gridCode: $gridCode) {
      gridCode
      gridVoltageSetting
      gridFrequencySetting
      gridPhaseSetting
      gridPerPhaseNetMeterEnabled
      gridCodePoints(names: $pointNames) {
        name
        units
        min
        max
        fileValue
      }
    }
  }
}
`,
}
