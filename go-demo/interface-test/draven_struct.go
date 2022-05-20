package main

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

//go:noinline
func (c Cat) Quack() {
	println(c.Name + " meow")
}

func draven_struct() {
	var c Duck = Cat{Name: "draven"}
	c.Quack()
}
