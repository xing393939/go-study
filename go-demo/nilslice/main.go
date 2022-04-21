package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var nilStruct struct{}
	var nilSlice []int
	emptySlice := make([]int, 0)
	zeroSlice := make([]int, 10)

	nilSliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&nilSlice))
	emptySliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&emptySlice))
	zeroSliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&zeroSlice))

	fmt.Printf("%p\n", &nilStruct)
	fmt.Printf("0x%x\n", emptySliceHeader.Data)
	fmt.Printf("%+v %+v %+v", nilSliceHeader, emptySliceHeader, zeroSliceHeader)
}
