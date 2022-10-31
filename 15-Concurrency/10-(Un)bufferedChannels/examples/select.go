package examples

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func SelectExample() {

	now := time.Now()
	resource1 := make(chan string)
	resource2 := make(chan string)
	go GetHttpRequestSelect(resource1, "https://livescore-api.varzesh3.com/v1.0/livescore/today")
	go GetHttpRequestSelect(resource2, "https://footba11.co/json/liveFeed")

	select {
	case result1 := <-resource1:
		fmt.Println(result1)
		break
	case result2 := <-resource2:
		fmt.Println(result2)
		break
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout")
		// default:
		// 	fmt.Println("Default")
	}

	println("Time: ", time.Since(now).Milliseconds())
}

func GetHttpRequestSelect(content chan<- string, url string) {
	if url == "https://livescore-api.varzesh3.com/v1.0/livescore/today" {
		time.Sleep(3 * time.Second)
	}
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{}

	req.Header.Add("referer", "https://www.varzesh3.com/")

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	dst := &bytes.Buffer{}
	if err := json.Indent(dst, responseBody, "", "  "); err != nil {
		panic(err)
	}

	PrintlnWithTime("Before set content")
	content <- dst.String()

	PrintlnWithTime("After set content")
}
