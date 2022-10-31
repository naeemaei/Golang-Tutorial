package examples

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func RunHttpUExample1() {
	now := time.Now()

	response := make(chan string)
	for i := 0; i < total; i++ {
		go Get1(response, "https://jsonplaceholder.typicode.com/todos/"+strconv.Itoa(i+1), i)
	}

	for i := 0; i < total; i++ {
		fmt.Println("time: ", time.Now(), ":", <-response)
	}

	println("time: ", time.Since(now).Milliseconds(), "milliseconds")
}

func RunHttpUExample2() {
	now := time.Now()

	response := make(chan RequestResponse)
	for i := 0; i < total; i++ {
		go GetU2(response)
		response <- RequestResponse{Url: "https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(i+1), Index: i + 1}
	}

	for i := 0; i < total; i++ {
		fmt.Println("time: ", time.Now(), ":", (<-response).Index)
	}

	println("time: ", time.Since(now).Milliseconds(), "milliseconds")
}

func RunHttpUExample3() {
	now := time.Now()

	response := make(chan RequestResponse)
	for i := 0; i < total; i++ {
		go Get3(response, "https://jsonplaceholder.typicode.com/todos/"+strconv.Itoa(i+1), i)
	}

	for i := 0; i < total; i++ {
		fmt.Println("time: ", time.Now(), ":", (<-response).Index)
	}

	println("time: ", time.Since(now).Milliseconds(), "milliseconds")
}

func GetU1(content chan<- string, url string, index int) {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content <- string(responseBody)
	println("added to channel buffer: ", index)
}
func GetU2(content chan RequestResponse) {
	requestResponse := <-content
	response, err := http.Get(requestResponse.Url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content <- RequestResponse{
		Url:   requestResponse.Url,
		Body:  string(responseBody),
		Index: requestResponse.Index,
	}
	println("added to channel buffer: ", requestResponse.Index)
}

func GetU3(content chan RequestResponse, url string, index int) {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	time.Sleep(time.Millisecond * 1100)
	content <- RequestResponse{
		Url:   url,
		Body:  string(responseBody),
		Index: index,
	}
	println("added to channel buffer: ", index)
}
