package leak_timeTicker

import (
	"testing"
	"time"
)

func BenchmarkTicker1(b *testing.B) {
	temp := 1
	for i := 1; i < b.N; i++ {
		select {
		case <-time.NewTicker(time.Microsecond * 10).C:
			temp++
		case <-time.NewTicker(time.Microsecond * 20).C:
			temp++
		case <-time.NewTicker(time.Microsecond * 10).C:
			temp++
		case <-time.NewTicker(time.Microsecond * 20).C:
			temp++
		}
	}
	b.Log(temp)
}

func BenchmarkTicker2(b *testing.B) {
	temp := 1
	t1 := time.NewTicker(time.Microsecond * 10)
	t2 := time.NewTicker(time.Microsecond * 20)
	for i := 1; i < b.N; i++ {
		select {
		case <-t1.C:
			temp++
		case <-t2.C:
			temp++
		}
	}
	b.Log(temp)
}

func BenchmarkAfter(b *testing.B) {
	temp := 1
	for i := 1; i < b.N; i++ {
		select {
		case <-time.After(time.Microsecond * 10):
			temp++
		case <-time.After(time.Microsecond * 20):
			temp++
		}
	}
	b.Log(temp)
}
