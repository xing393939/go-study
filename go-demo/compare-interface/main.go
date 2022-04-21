package main

type myError struct{}

func (myError) Error() string { return "" }

func isNil(v interface{}) {
	println(v == nil)
}

func main() {
	var myErr1 myError
	var myErr2 *myError
	// println(myErr1 == nil)
	println(myErr2 == nil)
	isNil(myErr1)
	isNil(myErr2)
	var a error = myErr1
	var b error = myErr2
	println(a == nil)
	println(b == nil)
}
