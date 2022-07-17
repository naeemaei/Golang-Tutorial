package main

import "fmt"

func main() {

	sum, multiply := Calculator("Ali", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Printf("sum: %d, multiply: %d\n", sum, multiply)
}

func Calculator(name string, numbers ...int) (sum int, mul int) {
	mul = 1
	for _, number := range numbers {
		sum += number
		mul *= number
	}
	return
}
