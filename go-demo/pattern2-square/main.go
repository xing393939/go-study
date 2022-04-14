package main

import "time"

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

func main() {
	// 建立管道.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	println(<-out) // 4
	println(<-out) // 9
}
