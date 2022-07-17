package main

import (
	"fmt"
	"sort"
)

func main() {

	numbers := []int{12, 5, 48, 32, 1, 2, 74, 36, 90, 100, 95, 42}

	fmt.Printf("%v\n", numbers)

	sortingFunc := func(i, j int) bool {
		return numbers[i] < numbers[j]
	}
	sort.Slice(numbers, sortingFunc)

	fmt.Printf("%v\n", numbers)
}
