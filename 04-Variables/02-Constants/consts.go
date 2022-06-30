// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	const pi float64 = 3.14
	const (
		name   = "iran"
		number = 25
		city   = "ahwaz"
	)

	fmt.Printf("%f  \n", pi)

	fmt.Printf("name: %s number: %d city: %s pi: %f \n", name, number, city, pi)

	const googleBaseUrl = "https://www.google.com"
	const googleMapUrl = "/maps"

	println(googleBaseUrl, googleMapUrl)
}
