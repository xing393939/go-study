package main

type Duck interface {
	Quack()
	Quack2()
}

type emptyDuck struct{}

func (e *emptyDuck) Quack()  {}
func (e *emptyDuck) Quack2() {}

type A struct {
	Duck
	Name string
}

func (a *A) Quack() {
	println(a.Name)
	a.Quack()
}

func (a *A) Quack2() {
	println(a.Name)
	a.Duck.Quack2()
}

func main() {
	a := &A{&emptyDuck{}, "a"}
	b := &A{a, "b"}
	b.Quack2()
}
