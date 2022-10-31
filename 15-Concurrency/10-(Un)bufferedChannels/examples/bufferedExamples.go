package examples

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type RequestResponse struct {
	Url          string
	ResponseBody string
	Index        int
}

func BufferedExample() {
	total := 100
	wg := sync.WaitGroup{}
	now := time.Now()
	// buffered channel unbuffered channel
	// send only, receive only, send and receive
	// sender := make(chan<- string)
	// receiver := make(<-chan string)
	wg.Add(total)
	content := make(chan RequestResponse, total)
	for i := 0; i < total; i++ {
		go GetHttpRequestBuffered(content, &wg)
		content <- RequestResponse{Url: "https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(i+1), Index: i + 1}
	}

	wg.Wait()
	close(content)

	for item := range content {
		PrintlnWithTime(item)
	}

	PrintlnWithTime("Time: ", time.Since(now).Milliseconds())
}

func GetHttpRequestBuffered(content chan RequestResponse, wg *sync.WaitGroup) {
	defer wg.Done()
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
	PrintlnWithTime("Before set content")
	content <- RequestResponse{Url: requestResponse.Url, ResponseBody: string(responseBody), Index: requestResponse.Index}

	PrintlnWithTime("After set content")
}

func PrintlnWithTime(args ...any) {
	fmt.Printf("Time: %s, %v\n", time.Now().Format(time.RFC3339Nano), args)
}
