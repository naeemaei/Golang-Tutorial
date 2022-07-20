package main

import "fmt"

type Test struct {
	Id   int
	name string
}

func main() {
	test := Test{Id: 1, name: "test"}
	test.Print()
	test.print2()
}

func (test Test) Print() {
	fmt.Println(test.Id, test.name)
}
func (test Test) print2() {
	fmt.Println(test.Id, test.name)
}

func AddToShoppingCart(productId int, userId int, quantity int) {
	checkUserStatus(userId)
	checkProductStatus(productId)
	checkQuantity(productId,quantity)
	checkPriceChanges(productId)
	addToCart(productId,quantity)
	updateShoppingCartTotalPrice()
}

func checkUserStatus(userId int) {
	// check user status
}

func checkProductStatus(productId int) {
	// check product status
}

func checkQuantity(productId int, quantity int) {
	// check product quantity
}

func checkPriceChanges(productId int) {
	// check price changes
}

func addToCart(productId int, quantity int) {
	// add to cart
}

func updateShoppingCartTotalPrice() {
	// update shopping cart total price
}