package main

type Myintinterface interface {
	fun()
}
type Myint int

func (m Myint) fun() {}

func main() {
	var mii Myintinterface = Myint(12)
	mii.fun()
}
