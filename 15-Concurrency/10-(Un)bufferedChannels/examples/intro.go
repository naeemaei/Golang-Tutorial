package examples

import "time"

func IntroExample() {
	numChannel := make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		numChannel <- 1
	}()

	println("Waiting for numChannel", time.Now().String())
	receivedNum := <-numChannel
	println("get data ", time.Now().String())

	println(receivedNum)

}
