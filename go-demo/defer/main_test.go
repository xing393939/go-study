package main

import "testing"

func TestRecoverError(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()

	panic("custom error")
}

func TestRecoverFactor(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()

	var f func([1000]int64)
	f = func(a [1000]int64) {
		f(a)
	}
	f([1000]int64{})
}
