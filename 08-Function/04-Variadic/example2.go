package main

import "fmt"

func main() {
	PrintLog("Ali", 1, "Reza", 3, "Error", 5, 6, "Info", 8, 9, false)

}

func PrintLog(logs ...interface{}) {
	for i, item := range logs {
		fmt.Printf("%d: %v\n", i, item)
	}
}
