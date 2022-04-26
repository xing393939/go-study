package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func main() {
	fmt.Printf("%s\n", test())
}

// 参考文章：https://blog.huoding.com/2021/10/14/964
func test() []byte {
	defer runtime.GC()
	x := make([]byte, 5)
	x[0] = 'h'
	x[1] = 'e'
	x[2] = 'l'
	x[3] = 'l'
	x[4] = 'o'
	a := StringToBytesBad(string(x))
	println(cap(a))
	return a
}

// 将reflect.SliceHeader强制转成[]byte，r.Data是uintptr，它指向的内存是unused，所以被GC回收了
func StringToSliceByte(s string) []byte {
	l := len(s)
	r := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*(*reflect.StringHeader)(unsafe.Pointer(&s))).Data,
		Len:  l,
		Cap:  l,
	}))
	return r
}

// 将reflect.SliceHeader的Data指向reflect.StringHeader的Data，编译器特殊处理过，标记指向的内存是used
func StringToSliceByteOk(s string) []byte {
	var b []byte
	l := len(s)
	p := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	p.Data = (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	p.Len = l
	p.Cap = l
	return b
}

//  这种方式虽然也不会被GC回收，但是[]byte是占32B，string占16B，会产生内存越界
func StringToBytesBad(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// gin的写法规避了内存越界的问题
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

//  这种方式不会产生内存越界
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
