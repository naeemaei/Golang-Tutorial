package examples

import (
	"io"
	"net/http"
	"strconv"
	"time"
)

func UnbufferedExample() {

	now := time.Now()
	// buffered channel unbuffered channel
	// send only, receive only, send and receive
	// sender := make(chan<- string)
	// receiver := make(<-chan string)
	content := make(chan string)
	for i := 0; i < 100; i++ {
		go GetHttpRequestUnbuffered(content, i+1)
	}

	for i := 0; i < 100; i++ {
		response := <-content
		println(response)
	}

	println("Time: ", time.Since(now).Milliseconds())
}

func GetHttpRequestUnbuffered(content chan<- string, index int) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(index))

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	PrintlnWithTime("Before set content")
	content <- string(responseBody)

	PrintlnWithTime("After set content")
}


type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	City      string `json:"city"`
}