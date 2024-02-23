package queries

import "golang.org/x/exp/maps"

var querySet = map[string]*SignedQuery{}

func init() {
	addQuery(DeviceControllerQuery)
	addQuery(ComponentsQuery)
	addQuery(GridCodesQuery)
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
