package main

import (
	"math/rand"
	"syscall"
	"testing"
	"time"
)

// b.N是串行执行
// go test -run=none -bench=BenchmarkSerial -count=2 .
func BenchmarkSerial(b *testing.B) {
	src := rand.NewSource(1)
	b.Log("pid:", syscall.Getpid())
	for i := 0; i < b.N; i++ {
		src.Int63()
	}
	b.ReportAllocs()
}

// b.RunParallel是并行，协程数=cpu核数
// go test -run=none -bench=BenchmarkParallel -count=2 .
func BenchmarkParallel(b *testing.B) {
	src := rand.NewSource(1)
	b.Log("pid:", syscall.Getpid())
	b.SetParallelism(0) // 设置协程数，=0无效
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			src.Int63()
		}
	})
}

// b.Run的好处是在b.Run之外的初始化不计入总测试时间
func BenchmarkRun(b *testing.B) {
	src := rand.NewSource(1)
	println("pid:", syscall.Getpid())
	time.Sleep(time.Second * 3)
	b.Run("BenchmarkRun", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src.Int63()
		}
	})
}
