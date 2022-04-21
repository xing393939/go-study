package main

import (
	"go-study/plan9/backtrace"
	"go-study/plan9/callgofunction"
	"go-study/plan9/math"
	"go-study/plan9/nilcompare"
	"go-study/plan9/registers"
	"go-study/plan9/swap"
	"go-study/plan9/test"
)

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

func (c Cat) Quack() {
	c.Name = "bbbb"
}

func a() {
	backtrace.ZzzPrintTrace('\n')
}

func b() {
	a()
}

func main() {

	println("====")

	nilcompare.Test()

	println(registers.Output(987654321))

	println(callgofunction.Output(2, 3))

	println(swap.Swap(1, 2))

	println(math.Add(10, 11))
	println(math.Sub(99, 15))
	println(math.Mul(11, 12))

	println(test.Test())
}
