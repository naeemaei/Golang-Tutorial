package main

import (
	"encoding/json"
	"fmt"
)

const apiUrl = "https://api.github.com/users/"

type Order struct {
	Id     int
	Price  int
	Status OrderStatus
}

type OrderStatus int

const (
	Created        OrderStatus = 0
	Processing                 = 1
	WaitForPayment             = 2
	PaymentDone                = 3
	Issue                      = 4
	Refund                     = 5
)

func main() {

	order1 := Order{Id: 1, Price: 100, Status: Created}
	order2 := Order{Id: 2, Price: 300, Status: PaymentDone}
	order3 := Order{Id: 3, Price: 200, Status: Issue}

	order1Json, _ := json.Marshal(order1)
	order2Json, _ := json.Marshal(order2)
	order3Json, _ := json.Marshal(order3)
	println(string(order1Json))
	println(string(order2Json))
	println(string(order3Json))
	fmt.Printf("%+v", order3)
}
