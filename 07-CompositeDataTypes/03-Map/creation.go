package main

import "fmt"

func main() {

	// Map Creation:

	names := make(map[string]string)
	names1 := map[string]string{}
	var names2 map[string]string = map[string]string{}

	names["1234567890"] = "Ali Hoessini"
	names1["1234567890"] = "Ali Hoessini"
	names2["1234567890"] = "Ali Hoessini"

	fmt.Printf("%v\n", names)
	fmt.Printf("%v\n", names1)
	fmt.Printf("%v\n", names2)

}
