package main

import "fmt"

const monthDays1 int = 31
const monthDays2 int = 30
const monthDays3 int = 29

func main() {

	var month int
	var totalDays int = 0

	println("Please enter a month number: ")

	fmt.Scanln(&month)

	switch month {
	case 12:
		totalDays += monthDays3
		fallthrough
	case 11:
		totalDays += monthDays2
		fallthrough
	case 10:
		totalDays += monthDays2
		fallthrough
	case 9:
		totalDays += monthDays2
		fallthrough
	case 8:
		totalDays += monthDays2
		fallthrough
	case 7:
		totalDays += monthDays2
		fallthrough
	case 6:
		totalDays += monthDays1
		fallthrough
	case 5:
		totalDays += monthDays1
		fallthrough
	case 4:
		totalDays += monthDays1
		fallthrough
	case 3:
		totalDays += monthDays1
		fallthrough
	case 2:
		totalDays += monthDays1
		fallthrough
	case 1:
		totalDays += monthDays1

	}

	println("Total days: ", totalDays)
}
