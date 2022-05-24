package main

import (
	"errors"
	"fmt"
	"io"
)

func main() {
	a := io.EOF
	b := fmt.Errorf("%w: from b", a)

	if errors.Is(b, io.EOF) {
		fmt.Println("ok")
	}
	var c error
	if errors.As(b, &c) {
		fmt.Println("ok", c)
	}

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)

	d := fmt.Errorf("%w: from c", b)
	fmt.Println(errors.Unwrap(d))
}
