package main

import "time"

func main() {
	c1 := make(chan bool)
	c2 := make(chan bool)
	c3 := time.After(time.Microsecond)
	c4 := time.After(time.Microsecond)
	select {
	case <-c1:
		println(1)
	case <-c2:
		println(2)
	case <-c3:
		println(3)
	case <-c4:
		println(4)
	}
}
