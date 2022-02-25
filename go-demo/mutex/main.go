package main

import (
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex

	go func() {
		println(1)
		mu.Lock()

		println(1)
		time.Sleep(time.Second * 30)
	}()

	go func() {
		println(2)
		mu.Lock()

		println(2)
		time.Sleep(time.Second * 30)
	}()

	mu.Lock()

	println(3)
	time.Sleep(time.Second * 1)
	println(3)
}
