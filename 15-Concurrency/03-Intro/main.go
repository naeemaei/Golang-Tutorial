package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	value := 0
	go task1() 
	go task2()
	go task3()

	go func() {
		// lock
		value++ //1
		//unlock
	}()
	go func() {
		// lock
		value+=2 // 3
		//unlock
	}()
	go func() {
		// lock
		value+=3 // 6
		//unlock
	}()

	println(value)

	time.Sleep(time.Second)
}

func task1() {
	println("task1")
}

func task2() {
	println("task2")
}

func task3() {
	println("task3")
}

// race condition
//await