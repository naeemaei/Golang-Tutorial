package main

import (
	"fmt"
	"unsafe"
)

func main() {
	char1 := '😀'

	char2 := '😄'
	char3 := 128525

	fmt.Printf("char1: %d %T\n", char1, char1)

	fmt.Printf("char1: %d %T\n", char2, char2)
	fmt.Printf("char1: %c %T\n", char3, char3)

	myStr := "Hello World 😀 😄"

	fmt.Printf("myStr: %s %T, len: %d, size: %d \n ", myStr, myStr, len(myStr), unsafe.Sizeof(myStr))

	for i := 0; i < len(myStr); i++ {
		fmt.Printf("myStr[%d]: %c %T\n", i, myStr[i], myStr[i])
	}

	myRune := []rune("Hello World 😀 😄")

	fmt.Printf("myRune: %v %T, len: %d, size: %d \n ", myRune, myRune, len(myRune), unsafe.Sizeof(myRune))

	for i := 0; i < len(myRune); i++ {
		fmt.Printf("myStr[%d]: %c %T\n", i, myRune[i], myRune[i])
	}

	myRune = []rune("سلام دنیا")
	myStr = "سلام دنیا"

	fmt.Printf("myStr: %s %T, len: %d, size: %d \n ", myStr, myStr, len(myStr), unsafe.Sizeof(myStr))

	for i := 0; i < len(myStr); i++ {
		fmt.Printf("myStr[%d]: %c %T\n", i, myStr[i], myStr[i])
	}

	fmt.Printf("myRune: %v %T, len: %d, size: %d \n ", myRune, myRune, len(myRune), unsafe.Sizeof(myRune))

	for i := 0; i < len(myRune); i++ {
		fmt.Printf("myStr[%d]: %c %T\n", i, myRune[i], myRune[i])
	}

}
