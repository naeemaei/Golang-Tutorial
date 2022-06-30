package main

import "fmt"

func main() {

	myArray := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	// mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// mySlice1 := make([]int, 8)

	// mySlice2 := make([]int, 8, 16)

	slc1 := myArray[2:6]

	myArray[2] = 20

	// fmt.Printf("%v\n", myArray)
	// fmt.Printf("%v\n", slc1)

	// fmt.Println("slc1 length: ", len(slc1))
	// fmt.Println("slc1 cap: ", cap(slc1))

	// fmt.Println("myArray length: ", len(myArray))
	// fmt.Println("myArray cap: ", cap(myArray))

	slc1 = append(slc1, 99)
	slc1 = append(slc1, 98)
	slc1 = append(slc1, 97)
	slc1 = append(slc1, 96, 95, 94, 93, 92, 91, 90)

	myArray[7] = 101

	fmt.Printf("%v\n", myArray)
	fmt.Printf("%v\n", slc1)

	fmt.Println("slc1 length: ", len(slc1))
	fmt.Println("slc1 cap: ", cap(slc1))

	fmt.Println("myArray length: ", len(myArray))
	fmt.Println("myArray cap: ", cap(myArray))
}
