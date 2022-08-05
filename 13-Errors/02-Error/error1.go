package main

import (
	"io/ioutil"
	"net/http"
)

func main() {

	response, err := http.Get("https://dummyjson3123.com/products/categories")

	if err != nil {
		println("an error has occurred on get request")
		return
	}

	println(response.Status)

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		println("an error has occurred on reading the response body")
	}

	println(string(responseBody))

}
