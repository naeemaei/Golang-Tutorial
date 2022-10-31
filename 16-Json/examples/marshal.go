package examples

import "encoding/json"

type Person struct {
	Name   string `json:"name"`
	Family string `json:"family"`
	Age    int    `json:"age,omitempty"`
}

func MarshalExample1() {
	person1 := Person{Name: "Ali", Family: "Rezaee", Age: 17}
	person2 := Person{Name: "Milad", Family: "Mohammadi", Age: 20}
	person3 := Person{Name: "Peyman", Family: "Mohammadi", Age: 0}

	person1Json, err := json.Marshal(person1)
	if err != nil {
		panic(err)
	}
	person2Json, err := json.Marshal(person2)
	if err != nil {
		panic(err)
	}
	person3Json, err := json.Marshal(person3)
	if err != nil {
		panic(err)
	}

	println(string(person1Json))
	println(string(person2Json))
	println(string(person3Json))
}

func MarshalExample2() {
	person1 := Person{Name: "Ali", Family: "Rezaee", Age: 17}
	person2 := Person{Name: "Milad", Family: "Mohammadi", Age: 20}
	person3 := Person{Name: "Peyman", Family: "Mohammadi", Age: 0}

	persons := []Person{person1, person2, person3}
	personsJson, err := json.Marshal(persons)
	if err != nil {
		panic(err)
	}

	println(string(personsJson))
}
