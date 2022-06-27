package trip

// AddTripRequest is request schema to book a trip
// swagger:parameters AddTripRequest
type AddTripRequest struct {
	// in: body
	// required: true
	Param *AddTripParam `json:"param"`
}

type AddTripParam struct {
	StartCity int `json:"from"`
	EndCity   int `json:"to"`
	StartTime int `json:"startTime"`
}

// TripRequest is request schema to start/end trip
// swagger:parameters StartTripRequest EndTripRequest
type TripRequest struct {
	// in: path
	// required: true
	TripID int `json:"tripID"`
	// in: body
	// required: false
	Param *TripParam `json:"param"`
}

type TripParam struct {
	Time int `json:"time"`
}
