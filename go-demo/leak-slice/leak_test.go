package leak_slice

import (
	"runtime"
	"testing"
)

var a []byte

//go:noinline
func getA1() []byte {
	b := make([]byte, 1024*1024*1024, 1024*1024*1024)
	a = append(a, b[0:2]...)
	return a
}

//go:noinline
func getA2() []byte {
	b := make([]byte, 1024*1024*1024, 1024*1024*1024)
	a = b[0:2]
	return a
}

func TestNoLeak(t *testing.T) {
	a := getA1()
	t.Log(a)
	runtime.GC()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	t.Log(m.Alloc)
}

func TestLeak(t *testing.T) {
	a := getA2()
	t.Log(a)
	runtime.GC()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	t.Log(m.Alloc)
}
