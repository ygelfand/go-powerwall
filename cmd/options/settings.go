package options

type PowerwallOptions struct {
	Endpoint  string
	Password  string
	DebugMode bool
}

type ProxyOptions struct {
	*PowerwallOptions
	RefreshInterval uint32
	OnDemand        bool
	ListenOn        string
}
