package main

func main() {
	defer println("1 in main")
	defer func() {
		defer func() {
			println("2")
			panic("panic again and again")
		}()
		println("3")
		panic("panic again")
	}()

	panic("panic once")
}
