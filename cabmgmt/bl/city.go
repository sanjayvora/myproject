package bl

import (
	"fmt"
	spec "myproject/cabmgmt/spec/city"
	"myproject/cabmgmt/transport"
)

type cityBL struct {
	DataLayer DL
}

func NewCityBL(dl DL) transport.CityBL {
	return &cityBL{DataLayer: dl}
}

func (c *cityBL) AddCity(r *spec.AddCityParam) (*spec.AddCityResponse, error) {
	fmt.Println("BL", "AddCity")
	cityID, err := c.DataLayer.AddCity(r.Name, r.State)
	if err != nil {
		return nil, err
	}
	return &spec.AddCityResponse{CityID: cityID}, nil
}

func (c *cityBL) GetCities(r *spec.GetCityRequest) (*spec.GetCityResponse, error) {
	fmt.Println("BL", "GetCities")
	return nil, nil
}

func (c *cityBL) DeleteCity(r *spec.DeleteCityRequest) error {
	fmt.Println("BL", "DeleteCity")
	return c.DataLayer.RemoveCity(r.CityID)
}

func (c *cityBL) GetDemand(r *spec.GetDemandRequest) (*spec.DemandResponse, error) {
	return nil, nil
}
