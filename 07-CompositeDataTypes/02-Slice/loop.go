package main

import (
	"fmt"
	"strings"
)

func main() {

	names := []string{"Hamed", "Farzaneh", "Reza", "Sara"}

	for _, item := range names {
		item = strings.ToUpper(item)
	}

	fmt.Printf("%v \n", names)

	for i, item := range names {
		names[i] = strings.ToUpper(item)
	}

	fmt.Printf("%v \n", names)
}
