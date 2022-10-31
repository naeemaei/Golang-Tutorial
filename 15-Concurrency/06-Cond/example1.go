package main

import (
	"log"
	"sync"
	"time"
)

var myList = []int{}
var done = false

func read(name string, c *sync.Cond) {
	defer c.L.Unlock()
	c.L.Lock()
	if !done {
		c.Wait()
	}
	log.Println(name, "starts reading")
}

func write(id int, c *sync.Cond) {
	log.Println(id, "starts writing")
	time.Sleep(time.Second)
	c.L.Lock()
	myList = append(myList, id)
	if len(myList) == 100 {
		done = true
		c.Broadcast()
		log.Println(id, "wakes all")
	}
	c.L.Unlock()
}

func RunExample1() {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	for i := 0; i < 1000; i++ {
		write(i, cond)
	}

	time.Sleep(time.Second * 3)
}

//
