package main

import (
	"fmt"
)

func main() {
	chanCap := 5
	intChan1 := make(chan int, chanCap)
	intChan2 := make(chan int, chanCap)
	for i := 0; i < 2*chanCap; i++ {
		select {
		case intChan1 <- 1:
			fmt.Println(1)
			intChan2 <- 1
		default:
			fmt.Println(2)
		}
	}

	for i := 0; i < chanCap; i++ {
		<-intChan1
	}

}
