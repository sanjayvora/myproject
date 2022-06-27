package cab

import "myproject/cabmgmt/svcparams"

type Cab struct {
	ID                int
	RegistationNumber string
	Model             string
	CurState          svcparams.CabState
	CurCity           int
	//driver details
	//year of passing
	//etc
}
