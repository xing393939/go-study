package main

type sudug struct {
	name     int
	waitlink *sudug
}

func main() {
	var nextp **sudug

	first := &sudug{}
	nextp = &first
	for i := 0; i < 3; i++ {
		sp := &sudug{name: i + 1}
		*nextp = sp
		nextp = &sp.waitlink
	}

	for i := 0; i < 3; i++ {
		println(first.name)
		first = first.waitlink
	}
}
