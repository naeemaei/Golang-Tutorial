package main

import "math/rand"

type CreditCard struct {
	CardNumber string
	ExpireDate string
	Cvv2       string
	BankName   string
}

func main() {

	cards := []CreditCard{
		{CardNumber: "6037991725253535", ExpireDate: "01/01", Cvv2: "123", BankName: "Melli"},
		{CardNumber: "5892101245457878", ExpireDate: "03/03", Cvv2: "245", BankName: "Sepah"},
		{CardNumber: "6104981247653214", ExpireDate: "02/04", Cvv2: "741", BankName: "Mellat"},
		{CardNumber: "6219861047653214", ExpireDate: "01/02", Cvv2: "1023", BankName: "Saman"},
	}

	for _, card := range cards {

		if card.ExpireDate < "01/03" {
			println("Your card", card.CardNumber, " is expired")
			continue
		}
		var remainAmount = getBankAccountRemainAmount(card.CardNumber, card.ExpireDate)
		println("Your card", card.CardNumber, " is valid, remain amount is ", remainAmount)
	}

}

func getBankAccountRemainAmount(cardNumber string, expireDate string) int {
	min := 1000000
	max := 30000000
	if expireDate < "01/03" {
		return 0
	}
	return (rand.Intn(max-min) + min)
}
