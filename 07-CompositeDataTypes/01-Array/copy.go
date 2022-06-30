package main

import "fmt"

func main() {
	numbers := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	numbers2 := &numbers

	numbers3 := numbers[:2]

	numbers4 := numbers3

	numbers[0] = 100

	println(&numbers)
	println(&numbers2)
	println(&numbers3)
	println(&numbers4)

	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("numbers2: %v\n", numbers2)

	changeValue(&numbers)
	changeValue(numbers2)

	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("numbers2: %v\n", numbers2)

}

func changeValue(mayArray *[8]int) {
	mayArray[0] = 555
	mayArray[1] = 666
}
