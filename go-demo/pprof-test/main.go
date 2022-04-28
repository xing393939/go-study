package main

import (
	"runtime"
)

func main() {
	a := make([]byte, 512*1024)
	b := make(map[int]int, 512*1024)
	c := make([]byte, 512*1024)
	d := struct {
		A *[]byte
		B *map[int]int
		C *[]byte
	}{
		&a, &b, &c,
	}
	d.A = nil
	d.B = nil
	d.C = nil

	// pprof收集采样收集内存信息是一直开启的
	// 可以用ebpf了抓取证明：/usr/share/bcc/tools/funccount ./main:runtime.mProf_*
	// 注意build包命令加上-N -l：go build -gcflags=all="-N -l" main.go

	// 调用了runtime.GC才会sweep对象，然后才会调用runtime.mProf_Free
	runtime.GC()
}
