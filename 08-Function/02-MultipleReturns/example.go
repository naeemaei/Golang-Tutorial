package main

import "fmt"

func main() {

	order1, tax1 := CalculateRoomPrice("standard", 3, 2)

	fmt.Printf("order1 price: %d, tax1: %f\n", order1, tax1)

}

func CalculateRoomPrice(roomType string, nights int, personCount int) (int, float64) {
	var price int
	var tax float64

	switch roomType {
	case "standard":
		price = nights * 140000 * personCount
	case "double":
		price = nights * 220000 * personCount
	case "suite":
		price = nights * 350000 * personCount
	}

	tax = float64(price) * 0.09
	return price, tax
}
