package main

type name int

func (n name) print() {
	println("print", n)
}

func (n *name) pprint() {
	println("pprint", *n)
}

func main() {
	var n name
	defer n.print()
	defer n.pprint()
	n = 3
}
