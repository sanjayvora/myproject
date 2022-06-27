package dl

import (
	cabspec "myproject/cabmgmt/spec/cab"
	tripspec "myproject/cabmgmt/spec/trip"
	"myproject/cabmgmt/svcparams"
)

type DataLayer struct {
}

// This is data layer implementation. To be implemented with actual interaction with RDBMS database like mysql/postgres
// Cab related DL functions
func (dl *DataLayer) AddCab(cab *cabspec.Cab) (int, error) {
	return 0, nil
}
func (dl *DataLayer) DeleteCab(cabID int) error {
	return nil
}
func (dl *DataLayer) ChangeCity(cabId int, city int) error {
	return nil
}
func (dl *DataLayer) UpdateStateOfCab(cabID int, state svcparams.CabState) error {
	return nil
}
func (dl *DataLayer) GetCabForCity(cityID int) (*cabspec.Cab, error) {
	return &cabspec.Cab{}, nil
}
func (dl *DataLayer) GetCab(cabID int) (*cabspec.Cab, error) {
	return &cabspec.Cab{}, nil
}

// City related DL functions
func (dl *DataLayer) AddCity(name, state string) (int, error) {
	return 0, nil
}
func (dl *DataLayer) RemoveCity(cityID int) error {
	return nil
}

// Trip related DL functions
func (dl *DataLayer) AddTrip(cabID int, fromCityID int, toCityID int, time int) (int, error) {
	return 0, nil
}
func (dl *DataLayer) GetTrip(tripID int) (*tripspec.Trip, error) {
	return &tripspec.Trip{}, nil
}
func (dl *DataLayer) GetTripForCab(cabID int) ([]*tripspec.Trip, error) {
	return []*tripspec.Trip{}, nil
}
func (dl *DataLayer) UpdateTrip(trip *tripspec.Trip) error {
	return nil
}
