package main

import "fmt"

type Animal interface {
	Eat()
	Sleep()
	Walk()
}

type Human interface {
	Animal
	Speak()
	Think()
}

type Employee struct {
	Human
	Name string
	Age  int
}

type Cat struct {
	Name string
}

func main() {
	employee := &Employee{Name: "Ali", Age: 30}
	cat := &Cat{Name: "Cat"}

	var human Human = employee
	var animal Animal = cat

	human.Eat()
	human.Sleep()
	human.Walk()
	human.Speak()
	human.Think()

	animal.Eat()
	animal.Sleep()
	animal.Walk()
}

func (cat *Cat) Eat() {
	fmt.Println("Cat is eating")
}

func (cat *Cat) Sleep() {
	fmt.Println("Cat is sleeping")
}

func (cat *Cat) Walk() {
	fmt.Println("Cat is walking")
}

func (employee *Employee) Speak() {
	fmt.Println("Employee is speaking")
}

func (employee *Employee) Think() {
	fmt.Println("Employee is thinking")
}

func (employee *Employee) Eat() {
	fmt.Println("Employee is eating")
}

func (employee *Employee) Sleep() {
	fmt.Println("Employee is sleeping")
}

func (employee *Employee) Walk() {
	fmt.Println("Employee is walking")
}


