package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       uint
}

type PersonOptions struct {
	FirstName string
	LastName  string
	Age       uint
}

func main() {

	// 1
	person1 := Person{FirstName: "Ali", LastName: "Khan", Age: 25}

	// 2
	person2 := new(Person)
	person2.FirstName = "Reza"
	person2.LastName = "Khan"
	person2.Age = 30

	// 3
	person3 := &Person{FirstName: "Sara", LastName: "Khan", Age: 35}

	// 4
	person4 := NewPerson("Maryam", "Kha", 33)

	// 5
	person5 := NewPersonWithOptions(PersonOptions{FirstName: "Milad", LastName: "Khan", Age: 30})

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
	fmt.Println(person4)
	fmt.Println(person5)

}

func NewPerson(FirstName, LastName string, Age uint) *Person {
	if len(LastName) < 4 {
		return nil
	}
	return &Person{FirstName: FirstName, LastName: LastName, Age: Age}
}

func NewPersonWithOptions(personOptions PersonOptions) *Person {
	if len(personOptions.LastName) < 4 {
		return nil
	}
	return &Person{FirstName: personOptions.FirstName, LastName: personOptions.LastName, Age: personOptions.Age}
}
