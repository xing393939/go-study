package main

import (
	"testing"
	"unsafe"
)

var data = []byte{64, 65, 66, 67}
var str string

func BenchmarkA(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			str = *(*string)(unsafe.Pointer(&data))
			_ = str
		}
	})
}

func BenchmarkB(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			str = string(data)
			_ = str
		}
	})
}

func BenchmarkC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str = *(*string)(unsafe.Pointer(&data))
		_ = str
	}
}

func BenchmarkD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str = string(data)
		_ = str
	}
}

func fib1(n int) int {
	if n < 2 {
		return n
	}
	if n == 2 {
		return 1
	}
	fnSub1 := 1
	fnSub2 := 1
	result := 0
	for i := 3; i <= n; i++ {
		result = fnSub1 + fnSub2
		fnSub2 = fnSub1
		fnSub1 = result
	}
	return result
}

func fib2(n int) int {
	if n < 2 {
		return n
	}
	return fib2(n-2) + fib2(n-1)

}

func TestFib(t *testing.T) {
	for i := 0; i < 15; i++ {
		if fib1(i) != fib2(i) {
			t.Fatal("err", i)
		}
	}
}

func BenchmarkFib1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib1(15)
	}
}

func BenchmarkFib2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib2(15)
	}
}
