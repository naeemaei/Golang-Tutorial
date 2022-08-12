package main

import (
	"bytes"
	"sync"
)

var data = make([]byte, 10000)

var pool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func RunBufferPoolExample() {

	PrintMemUsage()
	for i := 0; i < 10_000; i++ {
		UseBufferWithPool()
	}
	//runtime.GC()
	PrintMemUsage()

}

func UseBufferWithPool() {
	buf := pool.Get().(*bytes.Buffer)
	buf.Write(data)
	buf.Reset()
	pool.Put(buf)
}

func UseBuffer() {
	var buf bytes.Buffer
	buf.Write(data)
}

func RunNonBufferPoolExample() {

	PrintMemUsage()
	for i := 0; i < 10_000; i++ {
		UseBuffer()
	}
	//runtime.GC()
	PrintMemUsage()

}
