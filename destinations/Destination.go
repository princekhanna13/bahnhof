package destinations

import "time"

type Departure struct {
	Label               string `json:"label"`
	DepartureTime       int64  `json:"departureTime"`
	Destination         string `json:"destination"`
	Delay               int64  `json:"delay"`
	DepartureTimeinDate time.Time
	TimeInMinutes       int64
}

var HomeDestinaion = Departure{
	Label:       "S8",
	Destination: "Flughafen München",
}

var WorkDestinaion = Departure{
	Label:       "S8",
	Destination: "Ostbahnhof",
}
