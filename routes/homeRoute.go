package routes

import (
	"bahnhof/destinations"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func HomeRoute() string {
	response, err := http.Get("https://www.mvg.de/api//fahrinfo/departure/de:09162:5?footway=0")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data = make(map[string][]destinations.Departure)
	_ = json.Unmarshal(responseData, &data)
	currentTime := time.Now().UnixNano() / 1000000
	var returnVal string
	for _, d := range data["departures"] {
		if d.Label == destinations.HomeDestinaion.Label && d.Destination == destinations.HomeDestinaion.Destination {
			fmt.Println(time.Unix(0, (d.DepartureTime+60000*d.Delay)*int64(time.Millisecond)))
			d.DepartureTimeinDate = time.Unix(0, (d.DepartureTime+60000*d.Delay)*int64(time.Millisecond))
			timeInMinutes := (d.DepartureTime-currentTime)/60000 + d.Delay
			d.TimeInMinutes = timeInMinutes
			returnVal = fmt.Sprintf("Time to next train: %v minutes. \nTrain will arrive at %v.\nLate by %v \nTrain is going to %v", timeInMinutes, d.DepartureTimeinDate, d.Delay, d.Destination)
			break
		}
	}
	return fmt.Sprintln(returnVal)
}
