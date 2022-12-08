package main

import (
	"net"
	"os"
	"runtime/pprof"
	"testing"
)

func TestMain(m *testing.M) {
	cpuFile, _ := os.Create("./cpu.pprof")
	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()
	m.Run()
}

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		net.Dial("tcp", "baidu.com:80")
	}
}
