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
	//INSERT INTO CAB (reg_num, model, state, city) values(cab.RegistationNumber, cab.Model, cab.CurState, cab.CurCity)
	return 0, nil
}
func (dl *DataLayer) DeleteCab(cabID int) error {
	//DELETE FROM CAB where id = cabID
	return nil
}
func (dl *DataLayer) ChangeCity(cabId int, city int) error {
	//UPDATE CAB set city = city where id = cabID
	return nil
}
func (dl *DataLayer) UpdateStateOfCab(cabID int, state svcparams.CabState, time int) error {
	//UPDATE CAB set state = state, last_update_time = time where id = cabID
	return nil
}
func (dl *DataLayer) GetCabForCity(cityID int) (*cabspec.Cab, error) {
	//select id, reg_num, model, state, city from cab where city = cityID AND state = CAB_IDLE order by last_update_time
	return &cabspec.Cab{}, nil
}
func (dl *DataLayer) GetCab(cabID int) (*cabspec.Cab, error) {
	//select id, reg_num, model, state, city from cab where id = cabID
	return &cabspec.Cab{}, nil
}

// City related DL functions
func (dl *DataLayer) AddCity(name, state string) (int, error) {
	//INSERT INTO city (name, state) values(name, state)
	return 0, nil
}
func (dl *DataLayer) RemoveCity(cityID int) error {
	//DELETE FROM city where id = cityID
	return nil
}

// Trip related DL functions
func (dl *DataLayer) AddTrip(cabID int, fromCityID int, toCityID int, time int) (int, error) {
	//INSERT INTO trip (cabid, fromcity, tocity, bookingtime) values(cabid, fromCityID, toCityID, time)
	return 0, nil
}
func (dl *DataLayer) GetTrip(tripID int) (*tripspec.Trip, error) {
	//select id, cabid, fromcity, tocity, starttime, endtime from trip where id = tripID
	return &tripspec.Trip{}, nil
}
func (dl *DataLayer) GetTripForCab(cabID int) ([]*tripspec.Trip, error) {
	//select id, cabid, fromcity, tocity, starttime, endtime from trip where cabid = cabID
	return []*tripspec.Trip{}, nil
}
func (dl *DataLayer) UpdateTrip(trip *tripspec.Trip) error {
	//update trip set starttime=trip.startime, endtime=trip.endtime where tripid = tripID
	return nil
}
