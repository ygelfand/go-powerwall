package queries

var wallboxComponentsQuery = &SignedQuery{
	Name:      "WallboxComponentsQuery",
	SigKey:    2,
	Signature: `MIGHAkEWuEQaXGKJ7/6Y0YGjt+oyQgvz4NxVkwiYSEaNJ0hPgZXRubJbNZn1/t9pNjL4qk6//CF7gonekaFvLc9w8ccTOAJCASDV8S/929xABr3PgJ/bnKU25O4vWyUyXqqYbsNmApa9wkn0ul9ExWayQe73rmyokml8gmMZIN2oe6yeV9hLPFHS`,
	Query:     `query WallboxComponentsQuery($rootSignalNames:[String!]$sodaComponentsFilter:ComponentFilter$sodaSignalNames:[String!]$waspComponentsFilter:ComponentFilter$waspSignalNames:[String!]$stipComponentsFilter:ComponentFilter$stipSignalNames:[String!]$pchComponentsFilter:ComponentFilter$pchSignalNames:[String!]){components{root{signals(names:$rootSignalNames){name value textValue boolValue}}soda:components(filter:$sodaComponentsFilter){din signals(names:$sodaSignalNames){name value textValue boolValue timestamp}activeAlerts{name}}wasp:components(filter:$waspComponentsFilter){din signals(names:$waspSignalNames){name value textValue boolValue timestamp}activeAlerts{name}}stip:components(filter:$stipComponentsFilter){din signals(names:$stipSignalNames){name value textValue boolValue timestamp}activeAlerts{name}}pch:components(filter:$pchComponentsFilter){din signals(names:$pchSignalNames){name value textValue boolValue timestamp}activeAlerts{name}}}}`,
}
