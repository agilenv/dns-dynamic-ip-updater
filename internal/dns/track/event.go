package track

import "time"

type Event struct {
	Time      time.Time
	IP        string
	PublicAPI string
}

func NewEvent(time time.Time, ip, publicAPI string) Event {
	return Event{
		Time:      time,
		IP:        ip,
		PublicAPI: publicAPI,
	}
}
