package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers2 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	numbers = append(numbers, 6, 7, 8, 9)

	fmt.Printf("%v \n", numbers)
	
	numbers = append(numbers, numbers2...)
	
	fmt.Printf("%v \n", numbers)
}