package destinations

import "time"

type ConnectionList struct {
	Departure          int64 `json:"departure"`
	Arrival            int64 `json:"arrival"`
	DepartureTime      time.Time
	ConnectionPartList []ConnectionPart `json:"connectionPartList"`
}

type ConnectionPart struct {
	Label string `json:"label"`
	Delay int64  `json:"delay"`
}
