package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var mx = sync.Mutex{}
var rwmx = sync.RWMutex{}

type Breaker struct {
	Name  string
	Count int64
}

func main() {
	var mci = Breaker{"MCI", 0}
	var irancell = Breaker{"Irancell", 0}
	var rightel = Breaker{"Rightel", 0}

	for i := 0; i < 10_000; i++ {
		go AddBreak(&mci)
		go AddBreak(&irancell)
		go AddBreak(&rightel)
	}

	time.Sleep(time.Second * 2)

	fmt.Printf("MCI: %v\n", mci)
	fmt.Printf("Irancell: %v\n", irancell)
	fmt.Printf("Rightel: %v\n", rightel)

}

func getRandomNumber(min int, max int) int {
	return (rand.Intn(max-min) + min)
}

func AddBreak(breaker *Breaker) {
	// rwmx.Lock()
	// (*breaker).Count++
	// rwmx.Unlock()
	atomic.AddInt64(&((*breaker).Count), 10)
}

func ReadBreaker(breaker *Breaker) {
	rwmx.RLock()
	fmt.Printf("%v: %v\n", breaker.Name, breaker.Count)
	rwmx.RUnlock()
}

func ResetBreaker(breaker *Breaker) {
	(*breaker).Count = 0
}
