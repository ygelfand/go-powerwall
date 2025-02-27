package powerwall

import (
	"net/http"
	"net/url"
	"sync/atomic"

	"golang.org/x/sync/semaphore"
)

type PowerwallGateway struct {
	Endpoint     *url.URL
	password     string
	authToken    string
	userRecord   string
	httpClient   *http.Client
	Din          string
	refreshSem   *semaphore.Weighted
	authSem      *semaphore.Weighted
	Config       *ConfigResponse
	Controller   *DeviceControllerResponse
	ControllerV2 *DeviceControllerResponse
	lastRefresh  atomic.Int64
}

type State struct {
	PvStrings     map[string]PvString
	Inverters     []Inverter
	Pods          []Pod
	Temperature   float32
	InverterPower float32
}

type Inverter struct {
	State           string
	GridState       string
	FrequencyOut    float32
	PowerOut        float32
	VoltageOut      float32
	ActiveAlerts    []string
	VoltageLine1    float32
	VoltageLine2    float32
	PowerCapability float32
	Mia             bool
	Battery         *Pod
}

type Pod struct {
	EnergyRemaining float32
	EnergyFull      float32
	Mia             bool
}

type PvString struct {
	Current   float32
	Voltage   float32
	Power     float32
	Connected bool
	Mia       bool
}

type loginResponse struct {
	Email     string   `json:"email"`
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	Roles     []string `json:"roles"`
	Token     string   `json:"token"`
	Provider  string   `json:"provider"`
	LoginTime string   `json:"loginTime"`
}
