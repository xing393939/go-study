package main

import "time"

func main()  {
	c3 := time.After(time.Microsecond)
	c4 := c3



	go func() {
		<-c4
		println(4)
	}()

	go func() {
		<-c3
		println(3)
	}()

	time.Sleep(time.Second)
}
