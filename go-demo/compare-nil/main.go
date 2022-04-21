package main

import (
	"fmt"
)

func main() {
	temp := 1
	pointer1 := &temp
	println(pointer1 == nil)
	pointer1 = nil

	var s1 = make([]int, 10)
	var s2 = make([]int, 0)
	var s3 []int
	println("s1", s1 == nil) // 空切片分配了内存
	println("s2", s2 == nil) // 零切片指针指向zerobase
	println("s3", s3 == nil) // nil切片指针是0

	var m1 map[string]int
	var m2 = make(map[string]int)
	_, hm1 := mapTypeAndValue(m1)
	_, hm2 := mapTypeAndValue(m2)
	fmt.Println(hm1)
	fmt.Println(hm2)
	println(m1 == nil)
	println(m2 == nil)
}
