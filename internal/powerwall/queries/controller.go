package queries

import "golang.org/x/exp/maps"

var querySet = map[string]*SignedQuery{}

func init() {
	addQuery(deviceControllerQuery)
	addQuery(deviceControllerQueryV2)
	addQuery(componentsQuery)
	addQuery(gridCodesQuery)
	addQuery(wallboxComponentsQuery)
	addQuery(protectionTripQuery)
	addQuery(selfTestQuery)
	addQuery(ie2030query)
	addQuery(gridCodesDetailsQuery)
}

func addQuery(dq *SignedQuery) {
	querySet[dq.Name] = dq
}

func GetQuery(name string) *SignedQuery {
	return querySet[name]
}

func QueryList() []string {
	return maps.Keys(querySet)
}
