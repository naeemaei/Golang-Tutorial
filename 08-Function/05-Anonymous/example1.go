package main

import "fmt"

func main() {

	myFunc := func() {
		fmt.Println("Hello World")
	}

	x := 1200

	myFunc()
	myFunc()

	println(func(numbers ...int) int {
		var total int
		for _, number := range numbers {
			total += number
			x++
			println(&x)
		}
		return total
	}(1, 2, 3, 4, 5))

	println(x)

	sum := func(numbers ...int) int {
		var total int
		for _, number := range numbers {
			total += number
			x++
			println(&x)
		}
		return total
	}

	println(sum(1, 2, 3, 4, 5))
	println(x)
	println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	println(x)

}
