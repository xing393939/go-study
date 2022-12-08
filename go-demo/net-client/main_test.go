package net_client

import (
	"net/http"
	"testing"
)

var customClient = http.Client{}

func BenchmarkClient1(b *testing.B) {
	b.ReportAllocs()
	b.SetParallelism(1000)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			client := http.Client{}
			client.Get("http://127.0.0.1:8008/get?a=1")
		}
	})
}

func BenchmarkClient2(b *testing.B) {
	b.ReportAllocs()
	b.SetParallelism(1000)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			customClient.Get("http://127.0.0.1:8008/get?a=1")
		}
	})
}
