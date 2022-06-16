package main

import "unsafe"

//go:nospilt
func newMap() *map[int]unsafe.Pointer {
	m := map[int]unsafe.Pointer{}
	return &m
}

//go:nospilt
func fun() {
	a := 1
	m := *newMap()
	m[0] = unsafe.Pointer(&a)
}

func main() {

}
