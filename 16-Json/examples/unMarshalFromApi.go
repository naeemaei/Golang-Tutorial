package examples

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type BookingLocation struct {
	Country   string `json:"country"`
	State     string `json:"state"`
	StateName string `json:"stateName"`
	ZipCode   string `json:"zipcode"`
	Timezone  string `json:"timezone"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	City      string `json:"city"`
	Continent string `json:"continent"`
}

func UnmarshalExample3() {
	response := make(chan []byte)
	go GetResponse(response, "https://geolocation.onetrust.com/cookieconsentpub/v1/geo/location")

	var location = BookingLocation{}
	err := json.Unmarshal(<-response, &location)

	if err != nil {
		panic(err)
	}

	println(location.City)
	println(location.Country)
	println(location.StateName)
}

func GetResponse(content chan<- []byte, url string) {

	client := http.Client{}

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	request.Header = http.Header{}

	request.Header.Add("accept", "application/json")

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	PrintlnWithTime(string(responseBody))
	PrintlnWithTime("Before set content")
	content <- responseBody
	PrintlnWithTime("After set content")

}

func PrintlnWithTime(args ...any) {
	fmt.Printf("Time: %s, %v\n", time.Now().Format(time.RFC3339Nano), args)
}
