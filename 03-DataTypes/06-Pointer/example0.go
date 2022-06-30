// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	i, j := 20, 40

	var ip *int = &i
	var jp *int = &j

	fmt.Println("Address of i:", &i)
	fmt.Println("Address of ip  pointer:", ip)

	fmt.Println("Address of j:", &j)
	fmt.Println("Address of jp pointer:", jp)

	i2 := i
	i2 = i2 + 2
	fmt.Println("value of i after increase i2", i)
	fmt.Println("value of i2 after increase i2", i2)

	fmt.Println("Address of i2:", &i2)

	var ip2 *int = &i

	*ip2 = *ip2 + 2
	fmt.Println("value of ip2", *ip2)
	fmt.Println("value of i", i)
}
