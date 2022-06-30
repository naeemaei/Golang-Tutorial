package main

import "fmt"

func main() {

	var salary float64
	var minSalary float64 = 5_600_000
	var taxPercent float64 = 0
	var knowledgeBase bool = true

	fmt.Print("Enter your salary: ")
	fmt.Scanln(&salary)

	if salary <= minSalary {
		taxPercent = 0
	} else if knowledgeBase || salary > minSalary && salary <= minSalary*2 {
		taxPercent = 0.05
	} else if salary > minSalary*2 && salary <= minSalary*3 {
		taxPercent = 0.07
	} else if salary > minSalary*3 && salary <= minSalary*4 {
		taxPercent = 0.10
	} else {
		taxPercent = 0.15
	}

	fmt.Printf("Your tax percent is: %.2f\n", taxPercent)
	fmt.Printf("Your tax is: %.2f\n", taxPercent*salary)

	fmt.Printf("Your salary is: %.2f\n", salary-taxPercent*salary)

}
