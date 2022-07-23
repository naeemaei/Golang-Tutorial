package main

import "fmt"

type Person struct {
	Name string
}
func main() {
	Print("Hello")
	Print(123)
	Print(true)
	Print([]int{1, 2, 3})
	Print(map[string]int{"a": 1, "b": 2})
	Print(Person{"Ali"})
}

func AddOrder(order interface{}) interface{}{
	
	if 1 == 2{
		return 1
	}

	if 1 == 3{
		return "Cancel"
	}

	if 1 == 4{
		return true
	}

	if 1 == 4{
		return false
	}

	return "OK"
}

func Print(input interface{}) {
	switch input.(type) {
	case string:
		fmt.Println("String:", input)
	case int:
		fmt.Println("Integer:", input)
		case bool:
			fmt.Println("Boolean:", input)
		case []int:
			fmt.Println("Slice:", input)
		case map[string]int:
			fmt.Println("Map:", input)
		case Person:
			fmt.Println("Person:", input)
	}
}


// object