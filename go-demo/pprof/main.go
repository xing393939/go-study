package main

import (
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"time"
)

// 一段有问题的代码
func logicCode(done chan<- bool) {
	f2 := getAFile("./mem.pprof")
	s2 := make([]byte, 1024*1024)
	_, _ = f2.Write(s2)
	_ = f2.Close()
	time.Sleep(time.Second)
	done <- true
}

func getAFile(filename string) io.WriteCloser {
	f1, err := os.Create(filename)
	if err != nil {
		fmt.Printf("create cpu pprof failed, err:%v\n", err)
		return nil
	}
	return f1
}

func main() {
	pprof.StartCPUProfile(getAFile("./cpu.pprof"))
	defer pprof.StopCPUProfile()

	done := make(chan bool)
	for i := 0; i < 100; i++ {
		go logicCode(done)
	}
	for i := 0; i < 100; i++ {
		<-done
	}

	pprof.WriteHeapProfile(getAFile("./mem.pprof"))
}
