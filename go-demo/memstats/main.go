package main

import (
	"C"
	"os"
	"runtime/pprof"
)

type canceler struct {
	code int
	data uintptr
}

func sliceMemoryUse(num int) {
	max := make([]canceler, num, num)
	_ = &max
}

func mapMemoryUse(num int) {
	max := make(map[canceler]struct{}, num)
	_ = &max
}

func main() {
	println(C.int(0))

	num := 1024 * 1024
	mapMemoryUse(num)

	f, _ := os.Create("mem.out")
	pprof.WriteHeapProfile(f)
	f.Close()
}
