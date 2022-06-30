package main

import (
	"fmt"
	"math/bits"
	"unsafe"
)

func main() {

	var num int = 100000000000000000
	var num8 int8 = 100
	var num16 int16 = 10000
	var num32 int32 = 10000000
	var num64 int64 = 100000000000000000

	fmt.Printf("num %d bytes \n", unsafe.Sizeof(num))
	fmt.Printf("num8 %d bytes \n", unsafe.Sizeof(num8))
	fmt.Printf("num16 %d bytes \n", unsafe.Sizeof(num16))
	fmt.Printf("num32 %d bytes \n", unsafe.Sizeof(num32))
	fmt.Printf("num64 %d bytes \n", unsafe.Sizeof(num64))

	var a = bits.UintSize

	fmt.Println(a)

	var fnum float32 = 10.2 // 3.4 * 10^ -38
	var fnum8 float64       // 1.7 * 10 ^ -308

	fmt.Printf("fnum %d bytes \n", unsafe.Sizeof(fnum))
	fmt.Printf("fnum8 %d bytes \n", unsafe.Sizeof(fnum8))

}
