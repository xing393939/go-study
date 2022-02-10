package main

import (
	"go-study/plan9/math"
	"go-study/plan9/registers"
	"go-study/plan9/swap"
	"go-study/plan9/test"
)

func main() {
	println(test.Test())

	println(registers.Output(987654321))

	a, b := swap.Swap(1, 2)
	println(a, b)

	println(math.Add(10, 11))
	println(math.Sub(99, 15))
	println(math.Mul(11, 12))
}
