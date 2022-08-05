package main

import "sync"

func SimpleGood() {
	println("Start of SimpleGood")
	println("Write 1000000 counter")
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}

	counter := 0
	wg.Add(1000000)
	for i := 0; i < 1_000_000; i++ {
		go func() {
			defer wg.Done()
			mx.Lock()
			counter++
			mx.Unlock()
		}()
	}

	wg.Wait()
	println(counter)
}

func SimpleBad() {
	println("Start of SimpleBad")
	println("Write 1000000 counter")
	wg := sync.WaitGroup{}

	counter := 0
	wg.Add(1000000)
	for i := 0; i < 1_000_000; i++ {
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	println(counter)
}
