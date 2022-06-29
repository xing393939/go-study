package string_slice

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
