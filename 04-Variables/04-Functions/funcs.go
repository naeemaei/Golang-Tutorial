package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "this is golang tutorial go go"

	fmt.Println(strings.Contains(myString, "go1"))
	fmt.Println(strings.ContainsAny(myString, "go1"))
	fmt.Println(strings.Count(myString, "go"))
	fmt.Println(strings.Cut(myString, "go1"))

	myStringArray := strings.Split(myString, " ")

	println("Word Count: ", len(myStringArray))

	for _, item := range myStringArray {
		println(item)
	}

	myStringArray2 := strings.Join(myStringArray, ",")

	println(myStringArray2)

	println(strings.Repeat("iran ", 10))

	println(strings.Replace(myString, "golang", "go", 3))
	println(strings.ReplaceAll(myString, "go", "golang"))

	println(strings.Compare("golang", "golang"))
	println(strings.Compare("Golang", "golang"))
	println(strings.Compare("Golang", "GOLANG"))

	println(strings.EqualFold("Golang", "GOLANG"))
	println(strings.EqualFold("Golang", "golang"))


	println(strings.HasPrefix("Iran", "Ir"))
	println(strings.HasPrefix("Iran", "IR"))

	println(strings.HasSuffix("Iran", "an"))
	println(strings.HasSuffix("Iran", "n"))

	println(strings.Index("Iran", "an"))
	println(strings.Index("Iran", "r"))


	println(strings.ToLower("IrAn"))
	println(strings.ToUpper("IrAn"))
	println(strings.Title("Iran is a country"))

	println(strings.Trim("  Iran is a country   ", " "))
	println(strings.TrimLeft("!!Iran is a country!!", "!"))
	println(strings.TrimRight("!!Iran is a country!!", "!"))
	println(strings.Trim("!!Iran is a country!!", "!"))





}
