package main

import (
	"fmt"
)

func main() {
	var score float64

	println("Enter your score: ")
	fmt.Scanln(&score)

	switch {
	case score >= 16 && score <= 20:
		println("A")
		break
	case score >= 11 && score <= 15.99:
		println("B")
	case score >= 6 && score <= 10.99:
		println("C")
	case score >= 0 && score <= 5.99:
		println("D")
	default:
		println("Unknown")
	}
}
