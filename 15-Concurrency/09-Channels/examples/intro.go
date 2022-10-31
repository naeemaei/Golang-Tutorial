package examples

import (
	"fmt"
	"time"
)

func RunIntro() {
	numChannel := make(chan int)

	receivedNum := <-numChannel
	go SendDataToChannel(numChannel)

	PrintlnWithTime("Received number:", receivedNum)

	receivedNum = <-numChannel
	PrintlnWithTime("Received number:", receivedNum)

	time.Sleep(time.Second * 2)
}

func SendDataToChannel(numChannel chan int) {
	PrintlnWithTime("Before sending 1 to channel")
	numChannel <- 1 // ChannelName <- data
	PrintlnWithTime("After sending 1 to channel")

	PrintlnWithTime("Before sending 2 to channel")
	numChannel <- 2
	PrintlnWithTime("After sending 2 to channel")

	PrintlnWithTime("Before sending 3 to channel")
	numChannel <- 3
	PrintlnWithTime("After sending 3 to channel")

}

func PrintlnWithTime(args ...any) {
	fmt.Printf("Time: %s, %v\n", time.Now().Format(time.RFC3339Nano), args)
}
