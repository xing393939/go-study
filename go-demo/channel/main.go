package main

func main() {
	chan1 := make(chan bool)

	go func() {
		<-chan1
		println("a")
	}()
	go func() {
		<-chan1
		println("b")
	}()

	close(chan1)
	for {

	}
}
