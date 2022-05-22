package main

type Duck2 interface {
	Quack()
}

type Cat2 struct {
	Name string
}

//go:noinline
func (c Cat2) Quack() {
	println(c.Name + " meow")
}

func draven_struct() {
	var c Duck2 = Cat2{Name: "draven"}
	c.Quack()
}
