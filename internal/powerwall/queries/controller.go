package queries

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
