package main

import "fmt"

func main() {

	var num int
	var provinceName string

	println("Please enter num: ")

	fmt.Scanln(&num)

	switch num {
	case 72, 82, 92:
		provinceName = "Mazandaran"
	case 11, 22, 33, 44, 55, 66, 77, 88, 99, 10, 20, 30, 40, 50, 60, 70, 80, 90:
		provinceName = "Tehran"
	case 13, 23, 43, 53:
		provinceName = "Isfahan"
	case 47, 57, 67:
		provinceName = "Markazi"
	default:
		provinceName = "Unknown"
	}

	println(provinceName)
}
