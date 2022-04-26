package main

import (
	"sync"
	"time"
)

func gen(done <-chan int, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
		println("gen exit")
	}()
	return out
}

func sq(done <-chan int, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer println("sq exit")
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func merge(done <-chan int, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer func() {
			wg.Done()
			println("output exit")
		}()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
		println("merge exit")
	}()
	return out
}

func main() {
	done := make(chan int)

	in := gen(done, 2, 3)
	c1 := sq(done, in)
	c2 := sq(done, in)

	out := merge(done, c1, c2)
	<-out

	close(done)
	time.Sleep(time.Second)
}
