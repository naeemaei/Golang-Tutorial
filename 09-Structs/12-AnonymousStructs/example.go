package main

import "fmt"

type Person struct{
	int
	string
	float64
}
 
func main() {

	// API call and get response

	apiResponse := struct {
		ResultCode        int
		ResultMessage     string
		TransactionAmount float64
		TransactionTime   string
	}{
		ResultCode:        0,
		ResultMessage:     "Success",
		TransactionAmount: 100,
		TransactionTime:   "2020-01-01T00:00:00",
	}

	fmt.Printf("API Response: %+v \n", apiResponse)


	person := Person{1,"Alireza",19.75}

	fmt.Printf("Person: %+v \n", person)
}