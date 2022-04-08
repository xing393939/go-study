package main

import "time"

func httpRequestWithTimeout(url string, timeout time.Duration) string {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		ch <- "get from url: " + url
	}()

	select {
	case result := <-ch:
		return result
	case <-time.After(timeout):
		return "timeout"
	}
}

func lessThan1Second(ch chan<- int) {
	time.Sleep(time.Second)
	ch <- 1
}

func main() {
	println(httpRequestWithTimeout("1", time.Second))
	println(httpRequestWithTimeout("1", time.Second*3))

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
