// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	//var varName type = value
	// 1
	var number int = 12
	var firstName string = "Ali"

	println("number: ", number)
	println("firstName: ", firstName)

	// 2
	var i2 int
	var s2 string

	println("i2: ", i2)
	println("s2: ", s2)

	// 3
	var i3 = 3
	var f1 = 3.44
	var s3 = "Iran"

	println("i3: ", i3)
	println("f1: ", f1)
	println("s3: ", s3)

	// 4
	i4 := 44
	f2 := 55.2
	city := "Tehran"

	println("i4: ", i4)
	println("f2: ", f2)
	println("city: ", city)

	// 5
	var i5, f3, s5 = 55, 65.2, "Mazandaran"
	i6, f4, s6 := 55, 65.2, "Mazandaran"

	println("i5: ", i5)
	println("f3: ", f3)
	println("s5: ", s5)

	println("i6: ", i6)
	println("f4: ", f4)
	println("s6: ", s6)

	// 6

	var (
		i7 = 12
		f5 = 22.5
		s7 = "Isfahan"
	)

	println("i7: ", i7)
	println("f5: ", f5)
	println("s7: ", s7)
	fmt.Printf("Stack is 1 %s\n\n\n", string(debug.Stack()))
	fmt.Printf("Stack is 2 %s\n\n\n", string(debug.Stack()))
	debug.PrintStack()
	println("i7: ", i7)
	debug.PrintStack()

}
