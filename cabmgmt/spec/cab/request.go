package cab

// AddCabRequest is request schema to register cab
// swagger:parameters AddCabRequest
type AddCabRequest struct {
	// in: body
	// required: true
	Param *AddCabParam `json:"param"`
}

type AddCabParam struct {
	RegistationNumber string `json:"registrationNumber"`
	Model             string `json:"cabModel"`
	//driver details
	//year of passing
	//etc
}

// GetCabRequest is request schema to get cab details
// swagger:parameters GetCabRequest
type GetCabRequest struct {
	// in: path
	// required: true
	CabID int `json:"cabID"`
	// Fetch cab history
	// in: query
	// required: false
	GetCabHistory bool `json:"getHistory"`
	// Last n trip records for the cab.
	// Default brings last *5* trips
	// Only applicable if getHistory is set to true
	// in: query
	// required: false
	RecordCount int `json:"count"`
}

// DeleteCabRequest is request schema to delete a cab
// swagger:parameters DeleteCabRequest
type DeleteCabRequest struct {
	// in: path
	// required: true
	CabID int `json:"cabID"`
}

// ChangeCityRequest is request schema to delete a cab
// swagger:parameters ChangeCityRequest
type ChangeCityRequest struct {
	// in: path
	// required: true
	CabID int `json:"cabID"`
	// in: body
	// required: true
	Param *ChangeCityParam `json:"param"`
}

type ChangeCityParam struct {
	CityID int `json:"cityID"`
}
