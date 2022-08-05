package main

import "fmt"

type Number interface{
	int | int16 | int32 | int64 | float32 | float64
}

func main() {
	x := Sum(2, 7)
	fmt.Printf("%d\n", x)

	y := Sum(2.5, 7.5)
	fmt.Printf("%f\n", y)
}

func SumInts(a, b int) int {
	return a + b
}

func SumFloats(a, b float64) float64 {
	return a + b
}

func Sum[T Number](a, b T) T {
	return a + b
}


// func Call[Input any, Output any](headers string, uri string, input Input) Output {
// 	// http request() 
// 	//
	
// }
