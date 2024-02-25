package powerwall

import (
	"net/http"
	"net/url"

	"golang.org/x/sync/semaphore"
)

type PowerwallGateway struct {
	endpoint   *url.URL
	password   string
	authToken  string
	httpClient *http.Client
	Din        string
	refreshSem *semaphore.Weighted
	Config     *ConfigResponse
	Controller *DeviceControllerResponse
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

type podResponse struct {
	Name                    string  `json:"name"`
	PODActiveHeating        any     `json:"POD_ActiveHeating"`
	PODChargeComplete       any     `json:"POD_ChargeComplete"`
	PODChargeRequest        any     `json:"POD_ChargeRequest"`
	PODDischargeComplete    any     `json:"POD_DischargeComplete"`
	PODPermanentlyFaulted   any     `json:"POD_PermanentlyFaulted"`
	PODPersistentlyFaulted  any     `json:"POD_PersistentlyFaulted"`
	PODEnableLine           any     `json:"POD_enable_line"`
	PODAvailableChargePower any     `json:"POD_available_charge_power"`
	PODAvailableDischgPower any     `json:"POD_available_dischg_power"`
	PODNomEnergyRemaining   int     `json:"POD_nom_energy_remaining"`
	PODNomEnergyToBeCharged any     `json:"POD_nom_energy_to_be_charged"`
	PODNomFullPackEnergy    int     `json:"POD_nom_full_pack_energy"`
	PackagePartNumber       string  `json:"PackagePartNumber"`
	PackageSerialNumber     string  `json:"PackageSerialNumber"`
	PinvState               string  `json:"pinv_state"`
	PinvGridState           string  `json:"pinv_grid_state"`
	POut                    int     `json:"p_out"`
	QOut                    int     `json:"q_out"`
	VOut                    float64 `json:"v_out"`
	FOut                    float64 `json:"f_out"`
	IOut                    float64 `json:"i_out"`
	EnergyCharged           int     `json:"energy_charged"`
	EnergyDischarged        int     `json:"energy_discharged"`
	OffGrid                 int     `json:"off_grid"`
	VfMode                  int     `json:"vf_mode"`
	WobbleDetected          int     `json:"wobble_detected"`
	ChargePowerClamped      int     `json:"charge_power_clamped"`
	BackupReady             int     `json:"backup_ready"`
	OpSeqState              string  `json:"OpSeqState"`
	Version                 string  `json:"version"`
}
