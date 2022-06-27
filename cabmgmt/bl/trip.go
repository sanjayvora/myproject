package bl

import (
	"fmt"
	spec "myproject/cabmgmt/spec/trip"
	"myproject/cabmgmt/svcparams"
	"myproject/cabmgmt/transport"
)

type tripBL struct {
	DataLayer DL
}

func NewTripBL(dl DL) transport.TripBL {
	return &tripBL{DataLayer: dl}
}

func (c *tripBL) AddTrip(r *spec.AddTripParam) (*spec.AddTripResponse, error) {
	fmt.Println("BL", "AddTrip")
	cab, err := c.DataLayer.GetCabForCity(r.StartCity)
	if err != nil {
		return nil, err
	}
	tripID, err := c.DataLayer.AddTrip(cab.ID, r.StartCity, r.EndCity, r.StartTime)
	if err != nil {
		return nil, err
	}
	err = c.DataLayer.UpdateStateOfCab(cab.ID, svcparams.CAB_BOOKED)
	if err != nil {
		return nil, err
	}
	return &spec.AddTripResponse{TripID: tripID}, nil
}

func (c *tripBL) StartTrip(r *spec.TripRequest) error {
	fmt.Println("BL", "StartTrip")
	trip, err := c.DataLayer.GetTrip(r.TripID)
	if err != nil {
		return err
	}
	trip.StartTime = r.Param.Time
	c.DataLayer.UpdateTrip(trip)
	c.DataLayer.UpdateStateOfCab(trip.CabID, svcparams.CAB_ON_TRIP)
	return nil
}

func (c *tripBL) EndTrip(r *spec.TripRequest) error {
	fmt.Println("BL", "EndTrip")
	trip, err := c.DataLayer.GetTrip(r.TripID)
	if err != nil {
		return err
	}
	trip.EndTime = r.Param.Time
	c.DataLayer.UpdateTrip(trip)
	c.DataLayer.UpdateStateOfCab(trip.CabID, svcparams.CAB_IDLE)
	return nil
}
