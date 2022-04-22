package main

import (
	"fmt"
	"unsafe"
)

type myError struct{}

func (myError) Error() string { return "" }

func isNil(obj interface{}) bool {
	return obj == nil
}

type mockInterface struct {
	_type unsafe.Pointer
	value unsafe.Pointer
}

func main() {
	var s1 = make([]int, 10)
	var s2 = make([]int, 0)
	var s3 []int
	fmt.Println("s1", s1 == nil) // false 空切片分配了内存
	fmt.Println("s2", s2 == nil) // false 零切片指针指向zerobase
	fmt.Println("s3", s3 == nil) // true  nil切片指针是0

	var s []int
	fmt.Println("slice     :", s == nil) // true  nil切片
	fmt.Println("slice_func:", isNil(s)) // false 切片转换成接口类型，i._type和i.data都不为空

	var e error
	fmt.Println("error     :", e == nil) // true iface没有赋值具体类型，i._type和i.data都是nil
	fmt.Println("error_func:", isNil(e)) // true iface没有赋值具体类型，i._type和i.data都是nil
	fmt.Println((*mockInterface)(unsafe.Pointer(&e)))

	var i1 interface{} = nil
	fmt.Println("interface     :", i1 == nil) // true eface没有赋值具体类型，i._type和i.data都是nil
	fmt.Println("interface_func:", isNil(i1)) // true eface没有赋值具体类型，i._type和i.data都是nil
	fmt.Println((*mockInterface)(unsafe.Pointer(&i1)))

	var p *myError
	var i2 interface{} = p
	fmt.Println("interface     :", i1 == nil) // false i._type不为空，i.data都是nil
	fmt.Println("interface_func:", isNil(i1)) // false i._type不为空，i.data都是nil
	fmt.Println((*mockInterface)(unsafe.Pointer(&i2)))

	/*var m1 map[string]int
	var m2 = make(map[string]int)
	_, hm1 := mapTypeAndValue(m1)
	_, hm2 := mapTypeAndValue(m2)
	fmt.Println(hm1)
	fmt.Println(hm2)
	println(m1 == nil)
	println(m2 == nil)*/
}
