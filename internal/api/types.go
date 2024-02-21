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
