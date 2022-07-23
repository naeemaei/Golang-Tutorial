package main

import "fmt"

type IPerson interface {
	print()
}

type Person struct {
}

func main() {
	var p IPerson
	p = Person{}
	p.print()
	print2(p)
	print2(12)
	print2("ali")
}

func (person Person) print() {
	println("Hello")
}

func print2(item interface{}) {
	fmt.Printf("%v\n",item)
}