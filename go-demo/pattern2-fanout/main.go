package main

import (
	"sync"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			time.Sleep(time.Second)
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// 为cs 中每一个 input channel 启动一个 output goroutine。
	// output 函数从 c 中复制值到 out 直到 c 关闭，然后调用 wg.Done。
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// 当所有的 out goroutine 结束，开启一个 goroutine 来关闭 out。
	// 这之前必须先调用 wg.Add。
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	// 建立管道.
	in := gen(2, 3)
	c1 := sq(in)
	c2 := sq(in)

	for n := range merge(c1, c2) {
		println(n)
	}
}
