package cab

import "myproject/cabmgmt/svcparams"

type Cab struct {
	ID                int                `db:"id"`
	RegistationNumber string             `db:"reg_num"`
	Model             string             `db:"model"`
	CurState          svcparams.CabState `db:"state"`
	CurCity           int                `db:"city"`
	LastUpdateTime    int                `db:"last_update_time"`
	//driver details
	//year of passing
	//etc
}
