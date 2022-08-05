package main

import (
	"testing"
)

func BenchmarkInitializeSize(b *testing.B) {
	mySlc := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		mySlc[i] = i * 2
	}
	//log.Print(len(mySlc))
}

func BenchmarkNotInitializeSize(b *testing.B) {
	mySlc := []int{}
	for i := 0; i < b.N; i++ {
		mySlc = append(mySlc, i*2)
	}
	//log.Print(len(mySlc))
}
