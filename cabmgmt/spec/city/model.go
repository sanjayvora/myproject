package city

type City struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	State string `db:"state"`
	//pincodes
	//country
}
