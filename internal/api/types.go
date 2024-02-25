package api

import "encoding/json"

type PvacString struct {
	Connected bool    `json:"Connected"`
	Current   float64 `json:"Current"`
	Voltage   float64 `json:"Voltage"`
	State     string  `json:"State"`
}

func (s PvacString) MarshalJSON() ([]byte, error) {
	type rawPvacString PvacString
	return json.Marshal(struct {
		rawPvacString
		Power float64 `json:"Power"`
	}{rawPvacString(s), s.Power()})
}

func (s *PvacString) Power() float64 {
	return s.Current * s.Voltage
}

type pinvPower struct {
	Name                string  `json:"name,omitempty"`
	PINVFout            float64 `json:"PINV_Fout,omitempty"`
	PINVVSplit1         float64 `json:"PINV_VSplit1,omitempty"`
	PINVVSplit2         float64 `json:"PINV_VSplit2,omitempty"`
	PackagePartNumber   string  `json:"PackagePartNumber,omitempty"`
	PackageSerialNumber string  `json:"PackageSerialNumber,omitempty"`
	POut                float64 `json:"p_out,omitempty"`
	QOut                float64 `json:"q_out,omitempty"`
	VOut                float64 `json:"v_out,omitempty"`
	FOut                float64 `json:"f_out,omitempty"`
	IOut                float64 `json:"i_out,omitempty"`
}
type msaPower struct {
	PINVVSplit1 float64 `json:"PINV_VSplit1,omitempty"`
	PINVVSplit2 float64 `json:"PINV_VSplit2,omitempty"`
	PINVPSplit1 int     `json:"PINV_PSplit1,omitempty"`
	PINVPSplit2 int     `json:"PINV_PSplit2,omitempty"`
	PINVCSplit1 float64 `json:"PINV_CSplit1,omitempty"`
	PINVCSplit2 float64 `json:"PINV_CSplit2,omitempty"`
}

type islandPower struct {
	FreqL1Load float64 `json:"FreqL1_Load,omitempty"`
	FreqL1Main float64 `json:"FreqL1_Main,omitempty"`
	FreqL2Load float64 `json:"FreqL2_Load,omitempty"`
	FreqL2Main float64 `json:"FreqL2_Main,omitempty"`
	FreqL3Load float64 `json:"FreqL3_Load,omitempty"`
	FreqL3Main float64 `json:"FreqL3_Main,omitempty"`
	GridState  string  `json:"GridState,omitempty"`
	VL1NLoad   float64 `json:"VL1N_Load,omitempty"`
	VL1NMain   float64 `json:"VL1N_Main,omitempty"`
	VL2NLoad   float64 `json:"VL2N_Load,omitempty"`
	VL2NMain   float64 `json:"VL2N_Main,omitempty"`
	VL3NLoad   float64 `json:"VL3N_Load,omitempty"`
	VL3NMain   float64 `json:"VL3N_Main,omitempty"`
}

