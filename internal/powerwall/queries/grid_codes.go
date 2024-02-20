package queries

var GridCodesQuery = &SignedQuery{
	Name:      "GridCodesQuery",
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
