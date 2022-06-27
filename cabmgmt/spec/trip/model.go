package trip

type Trip struct {
	ID        int `json:"tripID"`
	CabID     int `json:"cabID"`
	StartTime int `json:"tripStartTime"`
	EndTime   int `json:"tripEndTime"`
	StartCity int `json:"from"`
	EndCity   int `json:"to"`
}
