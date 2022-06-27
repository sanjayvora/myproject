package bl

import (
	"fmt"
	spec "myproject/cabmgmt/spec/cab"
	"myproject/cabmgmt/svcparams"
	"myproject/cabmgmt/transport"
)

type cabBL struct {
	DataLayer DL
}

func NewCabBL(dl DL) transport.CabBL {
	return &cabBL{DataLayer: dl}
}

func (c *cabBL) AddCab(r *spec.AddCabParam) (*spec.AddCabResponse, error) {
	fmt.Println("BL", "AddCab")
	cab := &spec.Cab{
		RegistationNumber: r.RegistationNumber,
		Model:             r.Model,
		CurState:          svcparams.CAB_IDLE,
	}
	cabID, err := c.DataLayer.AddCab(cab)
	if err != nil {
		return nil, err
	}
	return &spec.AddCabResponse{CabID: cabID}, nil
}

func (c *cabBL) GetCab(r *spec.GetCabRequest) (*spec.GetCabResponse, error) {
	fmt.Println("BL", "GetCab")
	cab, err := c.DataLayer.GetCab(r.CabID)
	if err != nil {
		return nil, err
	}
	resp := &spec.GetCabResponse{
		RegistationNumber: cab.RegistationNumber,
		Model:             cab.Model,
		CurCity:           cab.CurCity,
	}
	resp.Trips, err = c.DataLayer.GetTripForCab(cab.ID)
	return resp, err
}

func (c *cabBL) DeleteCab(r *spec.DeleteCabRequest) error {
	fmt.Println("BL", "DeleteCab")
	return c.DataLayer.DeleteCab(r.CabID)
}

func (c *cabBL) ChangeCity(r *spec.ChangeCityRequest) error {
	fmt.Println("BL", "ChangeCity")
	return c.DataLayer.ChangeCity(r.CabID, r.Param.CityID)
}
