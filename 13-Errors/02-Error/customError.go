package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpError struct {
	StatusCode int
	Message    string
}

func main() {

	response, err := GetRequest("")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)

}

func (error HttpError) Error() string {
	return fmt.Sprintf("Http error occurred: %d %s", error.StatusCode, error.Message)
}

func NewHttpError(statusCode int, message string) *HttpError {
	return &HttpError{StatusCode: statusCode, Message: message}
}

func GetRequest(url string) (string, error) {
	if len(url) == 0 {
		return "", NewHttpError(400, "Url is empty")
	}
	response, err := http.Get(url)
	if err != nil {
		return "", NewHttpError(500, "Error occurred")
	}

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", NewHttpError(200, "Body is empty")
	}
	return string(responseBody), nil
}
