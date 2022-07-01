package main

import "fmt"

type Person struct {
	Name   string
	Family string
	Age    int
}

func main() {

	persons := make(map[string]Person)

	persons["1234567890"] = Person{Name: "Ali", Family: "Hoessini", Age: 30}
	persons["4832425354"] = Person{Name: "Reza", Family: "Rezaei", Age: 31}
	persons["8579456213"] = Person{Name: "Milad", Family: "Mohammadi", Age: 31}

	fmt.Printf("%v\n", persons)

	persons["1234567890"] = Person{Name: "Farhad", Family: "Davoodi", Age: 41}

	fmt.Printf("%v\n", persons)

	delete(persons, "8579456213")

	fmt.Printf("%v\n", persons)

	farhad, ok := persons["1234567840"]

	if ok {
		fmt.Printf("%v\n", farhad)
	} else {
		fmt.Printf("Not Found\n")
	}

}
