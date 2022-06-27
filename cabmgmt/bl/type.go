package bl

import (
	cabspec "myproject/cabmgmt/spec/cab"
	tripspec "myproject/cabmgmt/spec/trip"
	"myproject/cabmgmt/svcparams"
)

type DL interface {
	AddCab(cab *cabspec.Cab) (int, error)
	DeleteCab(cabID int) error
	ChangeCity(cabID int, cityID int) error
	UpdateStateOfCab(cabID int, state svcparams.CabState) error
	GetCabForCity(cityID int) (*cabspec.Cab, error)
	GetCab(cabID int) (*cabspec.Cab, error)

	AddCity(name string, state string) (int, error)
	RemoveCity(cityID int) error

	AddTrip(cabID int, from int, to int, time int) (int, error)
	GetTrip(tripID int) (*tripspec.Trip, error)
	GetTripForCab(cabID int) ([]*tripspec.Trip, error)
	UpdateTrip(trip *tripspec.Trip) error
}
