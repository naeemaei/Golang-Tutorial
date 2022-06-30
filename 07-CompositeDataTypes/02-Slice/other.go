package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
	// 1, 2, 3, 4, 5, 6, 7, 75, 8, 9, 10, 11, 12, 13, 14, 15
	numbers = append(numbers, 0)
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0

	_ = copy(numbers[8:], numbers[7:])
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0
	// 1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 10, 11, 12, 13, 14, 15

	numbers[7] = 75

	fmt.Printf("%v \n", numbers)

	numbers = append(numbers[:7],numbers[8:]...)

	fmt.Printf("%v \n", numbers)


	numbers = numbers[:0]
	fmt.Printf("%v \n", numbers)
	fmt.Printf("%d \n", len(numbers))
	fmt.Printf("%d \n", cap(numbers))

	

	numbers = nil
	fmt.Printf("%v \n", numbers)
	fmt.Printf("%d \n", len(numbers))
	fmt.Printf("%d \n", cap(numbers))


}
