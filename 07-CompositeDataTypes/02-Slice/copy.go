package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers2 := []int{10, 20, 30, 40, 50}

	count := copy(numbers, numbers2)

	fmt.Printf("%v \n", numbers)
	fmt.Printf("%d \n", count)

}
