package main

type Duck interface {
	Say()
}

type Dog struct{}

func (c *Dog) Say() {}

type Cat struct{}

func (c Cat) Say() {}

func main() {
	var a Duck = &Cat{}
	var b Duck = &Cat{}
	println(a == b)
}
