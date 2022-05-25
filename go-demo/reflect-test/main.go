package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := int(2)
	s := []int{}
	m := make(map[int]int)
	c := make(chan int)
	p := new(int)

	fmt.Printf("%v \n", reflect.ValueOf(&i).Elem().CanSet())
	fmt.Printf("%v \n", reflect.ValueOf(&s).Elem().CanSet())
	fmt.Printf("%v \n", reflect.ValueOf(&m).Elem().CanSet())
	fmt.Printf("%v \n", reflect.ValueOf(&c).Elem().CanSet())
	fmt.Printf("%v \n", reflect.ValueOf(p).Elem().CanSet())
}
