package svcparams

//This pkg maintains service level constants and environment variables.
//For now everything is defined as consts.
const (
	Port     = 12345
	BasePath = "/cabmgmt/v1"
)

type CabState int

const (
	CAB_IDLE CabState = iota
	CAB_BOOKED
	CAB_ON_TRIP
)
