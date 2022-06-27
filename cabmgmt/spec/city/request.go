package city

// AddCityRequest is request schema to add city record
// swagger:parameters AddCityRequest
type AddCityRequest struct {
	// in: body
	// required: true
	Param *AddCityParam `json:"param"`
}

type AddCityParam struct {
	Name  string `json:"cityName"`
	State string `json:"state"`
	//pincodes
	//country
}

// DeleteCityRequest is request schema to delete a city
// swagger:parameters DeleteCityRequest
type DeleteCityRequest struct {
	// in: path
	// required: true
	CityID int `json:"cityID"`
}

// GetCityRequest is request schema to get list of cities
// swagger:parameters GetCityRequest
type GetCityRequest struct {
}

// GetDemandRequest is request schema to get high demand city and peek time
// swagger:parameters GetDemandRequest
type GetDemandRequest struct {
	// in: query
	// required: false
	Count int `json:"count"`
}
