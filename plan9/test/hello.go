package main

func main() {
	hash := make([]int, 9, 98)
	hash[0] = 1
	hash[1] = 2
	hash[2] = 3
	say(hash)
}

//go:noinline
func say(hash2 []int) {
	hash2[3] = 4
	println(hash2)
}
