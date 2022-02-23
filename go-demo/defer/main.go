package _defer

func say() {
	defer hello(13)
	a := 11
	_ = a
}

//go:noinline
func hello(c int) {
	b := 12
	c = b
}
