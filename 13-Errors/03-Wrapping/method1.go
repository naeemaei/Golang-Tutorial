package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	err := CopyFile("src.txt", "dest.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func CopyFile(sourceName, destinationName string) error {

	source, err := os.Open(sourceName)

	if err != nil {
		return fmt.Errorf("during copy file could not open source file: %w", err)
	}

	/*defer*/ source.Close()

	destination, err := os.Create(destinationName)

	if err != nil {
		return fmt.Errorf("during copy file could not create destination file: %w", err)
	}

	defer destination.Close()

	_, err = io.Copy(destination, source)

	if err != nil {
		return fmt.Errorf("during copy file could not copy: %w", err)
	}
	//defer destination.Close()

	return nil

}
