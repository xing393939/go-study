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

func testCustomContext() {
	mctx := &MyContext{context.Background(), nil}

	_, childFun := context.WithCancel(mctx)
	childFun()

	time.Sleep(10 * time.Second)
}

func structExtentInterface() {
	parent := context.Background()
	currCtx, cancel := context.WithCancel(parent)
	defer cancel()
	// currCtx没有实现Deadline方法，但是parent实现了，可以兜底
	currCtx.Deadline()
}

type AA struct {
	A int
}

func (*AA) A1() {

}

type BB struct {
	AA
	B int
}

func main() {
	a := BB{}
	_ = a.A
	_ = a.B
	a.A1()
}
