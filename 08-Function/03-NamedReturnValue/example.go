package main

import "fmt"

func main() {

	price, finalPrice, tax := CalculateRoomPrice("standard", 3, 2)

	fmt.Printf("order price: %d, tax: %f\n finalPrice = price + tax (%d + %f) => %d\n", price, tax, price, tax, finalPrice)

}

func CalculateRoomPrice(roomType string, nights int, personCount int) (price int, finalPrice int, tax float64) {

	switch roomType {
	case "standard":
		price = nights * 140000 * personCount
	case "double":
		price = nights * 220000 * personCount
	case "suite":
		price = nights * 350000 * personCount
	}

	tax = float64(price) * 0.09
	finalPrice = price + int(tax)

	return
}
