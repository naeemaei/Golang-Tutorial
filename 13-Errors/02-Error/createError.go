package main

import (
	"errors"
	"fmt"
)

func main() {
	output, err := createErrorMethod2(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)

}

func createErrorMethod1(number int) (int, error) {
	if number == 0 {
		return 0, errors.New("Number is not valid")
	}
	return number * 5, nil
}

func createErrorMethod2(number int) (int, error) {
	if number == 0 {
		return 0, fmt.Errorf("Number is not valid: %d", number)
	}
	return number * 5, nil
}
