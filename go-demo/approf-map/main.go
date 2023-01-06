package main

import (
	"fmt"
	_ "net/http/pprof"
)

type A struct {
	name string
}

func deleteSlice(a []A, i int) []A {
	return append(a[:i], a[i+1:]...)
}

func main() {
	a := []A{
		{"a"}, {"b"}, {"c"},
	}
	fmt.Println(a)
	k := 0
	for _, row := range a {
		a[k] = row

		k++
		if row.name == "c" {
			k--
		}
		fmt.Println(k, row)
	}
	a = a[:k]
	fmt.Println(a)
}
