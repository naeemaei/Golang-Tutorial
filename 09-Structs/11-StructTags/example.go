package main

import (
	"encoding/json"
	"encoding/xml"
)

type Person struct {
	Name      string	`json:"first_name" xml:"name"`
	Age       int		`json:"age,omitempty" xml:"age"`
	Family    string	`json:"last_name" xml:"lastName"`
	Gender    string	`json:"gender" xml:"gender"`
	Height    int		`json:"height" xml:"height"`
	Weight    int		`json:"weight" xml:"weight"`
	HairColor string	`json:"hair_color" xml:"hairColor"`
}

func main() {
	person := Person{Name: "Milad", Family: "Haddadi",Gender: "Male", Height: 170, Weight: 70, HairColor: "Black"}

	personJson, _ := json.MarshalIndent(person, "", "	")

	println(string(personJson))

	personJson, _ = xml.MarshalIndent(person, "", "	")

	println(string(personJson))
}


// Search Orders: OrderFilterDto 
// passengerName `json:"passengerName" xml:"passengerName" dbField:"passenger.name" `