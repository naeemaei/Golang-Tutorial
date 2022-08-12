package main

import (
	"sync"
	"testing"
)

type DbConnection struct {
	DbName   string
	UserId   string
	Password string
	Timeout  int
}

var connectionPool = sync.Pool{
	New: func() any {
		return &DbConnection{
			DbName:   "DbTest",
			UserId:   "UserTest",
			Password: "PasswordTest",
			Timeout:  20,
		}
	},
}

func BenchmarkWithoutPool(b *testing.B) {
	var p *DbConnection
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			p = &DbConnection{}
			p.Timeout = 45
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var p *DbConnection
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			p = connectionPool.Get().(*DbConnection)

			connectionPool.Put(p)
		}
	}
}
