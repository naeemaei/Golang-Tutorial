package main

import "reflect"

func main() {
	Calculator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, "Ali")
}

func Calculator(numbers ...interface{}) {
	for _, item := range numbers {

		switch item.(type) {
		case int:
			println("int:", reflect.ValueOf(item).Int())
		case string:
			println("string:", reflect.ValueOf(item).String())
		}
	}
}
