package main

func main() {
	a := new([]int)
	println((*a)[0])
	*a = append(*a, 1)
}
