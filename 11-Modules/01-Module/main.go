package main

import (
	jalali "github.com/jalaali/go-jalaali"
)

func main() {
	x, y, z, _ := jalali.ToGregorian(1396, 2, 2)

	println(x, y, z)

}
