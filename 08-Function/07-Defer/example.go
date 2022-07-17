package main

import (
	"io"
	"os"
)

func main() {

	defer println("Bye")
	defer println("Good")

	for i := 0; i < 10; i++ {
		defer println(i)
	}

	println("Hello")

	CopyFile("F:\\GolangTutorial\\08-Function\\07-Defer\\destination.txt", "F:\\GolangTutorial\\08-Function\\07-Defer\\source.txt")

}

func CopyFile(destinationName, sourceName string) (written int64, err error) {

	source, err := os.Open(sourceName)

	if err != nil {
		return
	}

	defer source.Close()

	destination, err := os.Create(destinationName)

	if err != nil {
		return
	}

	defer destination.Close()

	return io.Copy(destination, source)

}
