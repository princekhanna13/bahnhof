package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func RouteFinder(fromID string, toID string) string {
	url := fmt.Sprintf("https://www.mvg.de/api/fahrinfo/routing/?fromStation=%s&toStation=%s&changeLimit=0", fromID, toID)
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data = make(map[string][]ConnectionList)
	_ = json.Unmarshal(responseData, &data)
	for _, d := range data["connectionList"] {
		if len(d.ConnectionPartList) > 1 {
			continue
		}
		loc, _ := time.LoadLocation("Europe/Berlin")
		d.DepartureTime = time.Unix(0, d.Departure*int64(time.Millisecond))
		d.DepartureTime = d.DepartureTime.In(loc)
		newTime := d.DepartureTime.Add(time.Minute * time.Duration(d.ConnectionPartList[0].Delay))
		return fmt.Sprintf("Train time : %s\nTrain Number : %s\nDelay : %v minutes\nPlatform: %s", newTime.Format(time.Kitchen), d.ConnectionPartList[0].Label, d.ConnectionPartList[0].Delay, d.ConnectionPartList[0].DeparturePlatform)
	}
	return "Not Route available"
}
