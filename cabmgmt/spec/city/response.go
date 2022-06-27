package city

// AddCityResponse is response schema for add city
// swagger:model AddCityResponse
type AddCityResponse struct {
	CityID int `json:"cityID"`
}

// GetCityResponse is response schema for list cities
// swagger:model GetCityResponse
type GetCityResponse struct {
	CityList []*CityRecord `json:"cities"`
}

type CityRecord struct {
	Name  string `json:"cityName"`
	State string `json:"state"`
}

// DemandResponse is response schema for demands
// swagger:model DemandResponse
type DemandResponse struct {
	Records []*Demands `json:"highDemands"`
}

type Demands struct {
	CityName int   `json:"cityName"`
	PeeHour  []int `json:"peekHour"`
}
