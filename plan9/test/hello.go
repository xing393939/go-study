package main

func main() {
	chanCap := make(chan bool)

	go func() {
		for v := range chanCap {
			println(&v)
		}
	}()

	chanCap <- true
	chanCap <- true
	chanCap <- true
	chanCap <- true
}
