package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	GetEmployeeInfo("ali", "alavi", 1000000)

	fmt.Println("end of main")
}

func GetEmployeeInfo(firstName, lastName string, salary int) float64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
			fmt.Println("Recovering from error...")
			debug.PrintStack()
		}
	}()
	if firstName == "" {
		panic("First name is empty")
	}

	fmt.Printf("First name: %s\n", firstName)

	if lastName == "" {
		panic("Last name is empty")
	}

	fmt.Printf("Last name: %s\n", lastName)

	if salary <= 0 {
		panic("Salary is less than 0")
	}

	fmt.Printf("Salary: %d\n", salary)

	return CalculateTax(salary)
}

func CalculateTax(salary int) float64 {
	tax := float64(salary) * 0.09
	if tax > 1000 {
		panic("Tax is greater than 1000")
	}
	fmt.Printf("Tax: %f\n", tax)
	return tax
}
