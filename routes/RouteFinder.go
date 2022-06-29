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

func RouteFinder(writer http.ResponseWriter, request *http.Request) {
	response, err := http.Get("https://www.mvg.de/api/fahrinfo/routing/?fromStation=de:09162:5&toStation=de:09162:700&changeLimit=0")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data = make(map[string][]destinations.ConnectionList)
	_ = json.Unmarshal(responseData, &data)
	for _, d := range data["connectionList"] {
		if len(d.ConnectionPartList) > 1 {
			continue
		}
		d.DepartureTime = time.Unix(0, d.Departure*int64(time.Millisecond))
		returnVal, _ := json.Marshal(d)
		writer.Write(returnVal)
		break
	}

	fmt.Println(data["connectionList"])
}
