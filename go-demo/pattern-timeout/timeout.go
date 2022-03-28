package main

import "time"

func lessThan1Second(ch chan<- int) {
	time.Sleep(time.Second)
	ch <- 1
}

func main() {
	ch := make(chan int, 1)
	go lessThan1Second(ch)

	// 最多等1秒，如果ch还没有值就走
	select {
	case <-ch:
		println("get it")
	case <-time.After(time.Second):
		println("timeout")
	}
}
