package main

import (
	"fmt"

	jalali "github.com/jalaali/go-jalaali"
	"github.com/naeemaei/moduleExample/services"
)

func main() {

	fmt.Printf("Hello, world.\n")

	year, month, day, error := jalali.ToGregorian(1401, 15, 1)
	if error == nil {
		fmt.Printf("%d/%d/%d\n", year, month, day)
	} else {
		fmt.Printf("%s\n", error)
	}

	var service services.TestService = services.TestService{}
	var service2 services.TestService2 = services.TestService2{}

	fmt.Printf("%v", service)
	fmt.Printf("%v", service2)

}
