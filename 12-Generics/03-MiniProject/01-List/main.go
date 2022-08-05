package main

import (
	"fmt"
	"genericList/generics"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	genericInt()
	genericString()
	genericPerson()
}

func genericInt() {
	list1 := generics.List[int]{Items: []int{}}

	list1.Add(10)
	list1.Add(20)
	list1.Add(30)
	fmt.Printf("%v\n", list1.Items)

	list1.InsertAt(1, 15)
	fmt.Printf("%v\n", list1.Items)

	list1.RemoveAt(2)
	fmt.Printf("%v\n", list1.Items)

	list1.Remove(30)
	fmt.Printf("%v\n", list1.Items)
}

func genericString() {
	list1 := generics.List[string]{Items: []string{}}

	list1.Add("Ali")
	list1.Add("Reza")
	list1.Add("Milad")
	fmt.Printf("%v\n", list1.Items)

	list1.Remove("Reza")
	fmt.Printf("%v\n", list1.Items)

	list1.InsertAt(1, "Reza1")
	fmt.Printf("%v\n", list1.Items)

}
func genericPerson() {
	list1 := generics.List[Person]{Items: []Person{}}
	list1.Add(Person{Name: "Ali", Age: 30})
	list1.Add(Person{Name: "Reza", Age: 20})
	list1.Add(Person{Name: "Milad", Age: 10})

	fmt.Printf("%v\n", list1.Items)

	list1.Remove(Person{Name: "Reza", Age: 20})
	fmt.Printf("%v\n", list1.Items)

	list1.InsertAt(1, Person{Name: "Reza1", Age: 20})
	fmt.Printf("%v\n", list1.Items)

}
