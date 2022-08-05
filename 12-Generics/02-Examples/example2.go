package main

import "fmt"

type Number interface {
	int | float64
}

func main() {

	mySlice := []int{1, 2, 3, 4, 5}

	fmt.Printf("%d\n", Sum(mySlice))

	mySlice2 := []float64{1.5, 2.5, 3.5, 4.5, 5.5}

	fmt.Printf("%f\n", Sum(mySlice2))
}

func Sum[T Number](slc []T) (sum T) {
	for _, v := range slc {
		sum += v
	}
	return
}
