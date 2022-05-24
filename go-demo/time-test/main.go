package main

import "time"

func main() {
	t1 := time.After(time.Second)
	t2 := time.Tick(time.Second)

	t3 := time.AfterFunc(time.Second, func() {

	})

	t4 := time.NewTimer(time.Second)
	t5 := time.NewTicker(time.Second)

	println(&t1, &t2, &t3, &t4, &t5)
}
