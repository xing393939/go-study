package main

import "sync"

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Done()
	wg.Wait()
	println("done")

	pool := sync.Pool{
		New: func() interface{} {
			return new(int)
		},
	}
	fish := pool.Get()
	println(fish.(*int))
	pool.Put(fish)
}
