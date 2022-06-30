package main

import "fmt"

func main() {

	var salary float64
	var minSalary float64 = 5_600_000
	var taxPercent float64 = 0

	fmt.Print("Enter your salary: ")
	fmt.Scanln(&salary)

	switch {
	case salary <= minSalary:
		taxPercent = 0
	case salary > minSalary && salary <= minSalary*2:
		taxPercent = 0.05
	case salary > minSalary*2 && salary <= minSalary*3:
		taxPercent = 0.07
	case salary > minSalary*3 && salary <= minSalary*4:
		taxPercent = 0.10
	default:
		taxPercent = 0.15
	}

	fmt.Printf("Your tax percent is: %.2f\n", taxPercent)
	fmt.Printf("Your tax is: %.2f\n", taxPercent*salary)

	fmt.Printf("Your salary is: %.2f\n", salary-taxPercent*salary)

}
