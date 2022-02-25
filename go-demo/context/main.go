package main

import (
	"context"
	"time"
)

type A struct {
}

func main() {
	ctx1, cancel := context.WithCancel(context.Background())
	go func() {
		ctx2, _ := context.WithCancel(ctx1)
		go func() {
			<-ctx2.Done()
			println("ctx3")
		}()

		<-ctx1.Done()
		println("ctx2")
	}()

	cancel()
	time.Sleep(time.Second)
}
