package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

var m = map[[12]byte]int{}

//var m = map[string]int{}

func init() {
	for i := 0; i < 1000000; i++ {
		var arr [12]byte
		copy(arr[:], fmt.Sprint(i))
		m[arr] = i

		//m[fmt.Sprint(i)] = i
	}
}

func main() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
