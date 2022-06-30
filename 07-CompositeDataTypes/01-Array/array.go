package main

import "fmt"

func main() {

	var myArr0 [5]int
	var myArr1 [6]int = [6]int{1, 2}
	myArr2 := [7]int{1, 2}
	myArr3 := [...]int{1, 2}

	myArr0[2] = 25
	myArr1[5] = 25
	myArr2[6] = 25
	myArr3[1] = 25

	fmt.Println(myArr0)
	fmt.Println(myArr1)
	fmt.Println(myArr2)
	fmt.Println(myArr3)

	fmt.Println(len(myArr0))
	fmt.Println(len(myArr1))
	fmt.Println(len(myArr2))
	fmt.Println(len(myArr3))

}
