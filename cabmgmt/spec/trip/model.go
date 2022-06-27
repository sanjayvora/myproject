package trip

type Trip struct {
	ID          int `db:"id" json:"tripID"`
	CabID       int `db:"cabid" json:"cabID"`
	BookingTime int `db:"bookingtime" json:"bookTime"`
	StartTime   int `db:"starttime" json:"tripStartTime"`
	EndTime     int `db:"endtime" json:"tripEndTime"`
	StartCity   int `db:"fromcity" json:"from"`
	EndCity     int `db:"tocity" json:"to"`
}
