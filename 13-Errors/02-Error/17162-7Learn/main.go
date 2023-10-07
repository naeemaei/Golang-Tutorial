package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type IOError struct {
	FileName string
	Message  string
	Err      error
}

func (error *IOError) Unwrap() error {
	return error.Err
}

func (error IOError) Error() string {
	return fmt.Sprintf("IO error occurred: FileName: %s Message: %s Detail: %s", error.FileName, error.Message, error.Err.Error())
}

func CopyFile(srcName, dstName string) error {
	src, err := os.Open(srcName)
	if err != nil {
		return &IOError{FileName: srcName, Message: "during copy file could not open source file", Err: err}
	}
	//defer src.Close()
	dst, err := os.Create(dstName)
	if err != nil {
		return &IOError{FileName: srcName, Message: "during copy file could not create destination file", Err: err}
	}
	dst.Close()
	_, err = io.Copy(dst, src)
	fmt.Printf("io error: %s\n", err)
	var ioErr IOError = IOError{}
	if errors.Is(err,ioErr){
		fmt.Printf("cast error: %s\n", ioErr)
	}
	if err != nil {
		return &IOError{FileName: srcName, Message: "during copy file could not copy", Err: err}
	}
	return nil
}
func main() {
	err := CopyFile("src.txt", "dst.txt")
	fmt.Printf("%s\n", err)
	fmt.Printf("Unwrap Error: %s", errors.Unwrap(err))

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
