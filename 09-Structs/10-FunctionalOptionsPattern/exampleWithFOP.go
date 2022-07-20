package main

import "fmt"

type Person struct {
	Name      string
	Age       int
	Family    string
	Gender    string
	Height    int
	Weight    int
	HairColor string
}

type PersonOptions func(*Person)

func main() {

	person := NewPerson(SetName("Maryam"), SetFamily("Rezaei"), SetAge(23), SetGender("Female"), SetHeight(175), SetWeight(70), SetHairColor("Black"))

	fmt.Printf("Person : %+v" , person)

}

func NewPerson(options ...PersonOptions) *Person {
	person := &Person{Name: "Alireza", Age: 23}
	for _, option := range options {
		option(person)
	}
	return person
}

func SetName(name string) PersonOptions {
	return func(person *Person) {
		person.Name = name
	}
}

func SetAge(age int) PersonOptions {
	return func(person *Person) {
		person.Age = age
	}
}

func SetFamily(family string) PersonOptions {
	return func(person *Person) {
		person.Family = family
	}
}

func SetGender(gender string) PersonOptions {
	return func(person *Person) {
		person.Gender = gender
	}
}

func SetHeight(height int) PersonOptions {
	return func(person *Person) {
		person.Height = height
	}
}

func SetWeight(weight int) PersonOptions {
	return func(person *Person) {
		person.Weight = weight
	}
}

func SetHairColor(hairColor string) PersonOptions {
	return func(person *Person) {
		person.HairColor = hairColor
	}
}