package simple_factory

import "fmt"

type Person interface {
	Greet()
}

type person struct {
	name string
	age  int
}

func (p person) Greet() {
	fmt.Printf("Hi! My name is %s\n", p.name)
}

func NewPerson(name string, age int) Person {
	return person{
		name: name,
		age:  age,
	}
}
