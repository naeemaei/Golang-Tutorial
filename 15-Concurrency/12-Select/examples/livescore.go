package examples

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// https://footba11.co/json/liveFeed
// https://livescore-api.varzesh3.com/v1.0/livescore/today

// referer: https://www.varzesh3.com/

func SelectExample1() {
	resource1 := make(chan string)
	resource2 := make(chan string)
	var result = ""

	go GetResponse(resource1, "https://footba11.co/json/liveFeed")
	go GetResponse(resource2, "https://livescore-api.varzesh3.com/v1.0/livescore/today")

	select {
	case result = <-resource1:
		println("footba11 returned")
	case result = <-resource2:
		println("varzesh3 returned")
	case <-time.After(time.Second * 3):
		println("Timeout")
	default:
		println("Default")
	}
	println(result)
	PrintlnWithTime("End")
}

func SelectExample2() {
	resource1 := make(chan string)
	resource2 := make(chan string)
	var result string

	go GetResponse(resource1, "https://footba11.co/json/liveFeed")
	go GetResponse(resource2, "https://livescore-api.varzesh3.com/v1.0/livescore/today")

	for {
		select {
		case result = <-resource1:
			println("footba11 returned")
			println(result)
			return
		case result = <-resource2:
			println("varzesh3 returned")
			println(result)
			return
		default:
			println("Default")
		}
	}

}

func GetResponse(content chan<- string, url string) {
	// if url == "https://livescore-api.varzesh3.com/v1.0/livescore/today" {
	// 	time.Sleep(time.Millisecond * 50)
	// }
	// time.Sleep(time.Second * 3)
	client := http.Client{}

	request, err := http.NewRequest("GET", url, nil)

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*10000)
	defer cancel()
	request = request.WithContext(ctx)

	if err != nil {
		panic(err)
	}

	request.Header = http.Header{}
	request.Header.Add("referer", "https://www.varzesh3.com/")

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	destination := &bytes.Buffer{}

	if err = json.Indent(destination, responseBody, "", "  "); err != nil {
		panic(err)
	}
	PrintlnWithTime("Before set content")
	content <- destination.String()
	PrintlnWithTime("After set content")

}

func PrintlnWithTime(args ...any) {
	fmt.Printf("Time: %s, %v\n", time.Now().Format(time.RFC3339Nano), args)
}
