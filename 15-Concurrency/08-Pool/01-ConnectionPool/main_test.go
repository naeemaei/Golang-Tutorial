package main

import "testing"

func BenchmarkWithoutPool(b *testing.B) {
	var connection *DbConnection
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10000; i++ {
			connection = &DbConnection{
				Host:     "localhost",
				DbName:   "test",
				User:     "root",
				Password: "root",
			}

			connection.DbName = "test"
		}
	}

}

func BenchmarkWithPool(b *testing.B) {
	var connection *DbConnection
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10000; i++ {
			connection = connectionPool.Get().(*DbConnection)
			connection.DbName = "test"
			connectionPool.Put(connection)
		}
	}
}
