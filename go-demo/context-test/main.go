package main

import (
	"context"
	"time"
)

type MyContext struct {
	context.Context
	done chan struct{}
}

func (c *MyContext) Done() <-chan struct{} {
	if c.done == nil {
		c.done = make(chan struct{})
	}
	d := c.done
	return d
}

func main() {
	mctx := &MyContext{context.Background(), nil}

	_, childFun := context.WithCancel(mctx)
	childFun()

	time.Sleep(10 * time.Second)
}
