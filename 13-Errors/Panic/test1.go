package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	defer func() {
		fmt.Println("Start of Defer main")
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
			fmt.Println("Recovering from error...")
			debug.PrintStack()
		}
	}()
	nums := []int{1, 2, 3, 4, 5}
	function1(nums, 6)

	Divide(10, 2)

	Divide(10, 0)
}

func function1(numbers []int, index int) {
	defer func() {
		fmt.Println("Start of Defer function1")
	}()
	// 1 - index out of range
	//fmt.Println(numbers[index])
	// 2 - panic
	//panic("panic")
	// 3 - log
	//log.Fatal("Fatal")
	//fmt.Println(12)
}

func Divide(a, b int) {
	defer func() {
		fmt.Println("Defer of divide")
	}()
	fmt.Printf("Start of divide: %d , %d\n", a, b)
	fmt.Printf("Result: %d\n", a/b)
	fmt.Printf("End of divide: %d , %d\n", a, b)

}
