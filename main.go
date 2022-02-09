package main

import (
	"go-study/plan9/math"
	"go-study/plan9/swap"
)

func main() {
	a, b := swap.Swap(1, 2)
	println(a, b)

	println(math.Add(10, 11))
	println(math.Sub(99, 15))
	println(math.Mul(11, 12))
}
