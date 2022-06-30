package main

import "fmt"

func main() {

	name := "Mahtab"
	age := 28
	nationalCode := 4832145478
	score := 7.5
	print("My name is ", name, " and my age is ", age, " and my national code is ", nationalCode, " and my score is ", score, "\n")
	println("My name is", name, "and my age is", age, "and my national code is", nationalCode, "and my score is", score)
	// println()
	fmt.Printf("My name is %s and my age is %d and my national code is %d and my score is %f\n", name, age, nationalCode, score)

	fmt.Printf("name: My type is %T\n", name)
	fmt.Printf("nationalCode binary is %b", nationalCode)

}
