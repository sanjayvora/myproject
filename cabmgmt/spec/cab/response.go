package cab

import tripspec "myproject/cabmgmt/spec/trip"

// AddCabResponse is response schema for add cab
// swagger:model AddCabResponse
type AddCabResponse struct {
	CabID int `json:"cabID"`
}

// GetCabResponse is response schema for add cab
// swagger:model GetCabResponse
type GetCabResponse struct {
	RegistationNumber string           `json:"registrationNumber"`
	Model             string           `json:"model"`
	CurCity           int              `json:"curCity"`
	Trips             []*tripspec.Trip `json:"trips"`
}

type Trips struct {
	ID        int `json:"tripID"`
	StartTime int `json:"tripStartTime"`
	EndTime   int `json:"tripEndTime"`
	StartCity int `json:"from"`
	EndCity   int `json:"to"`
}
