package main

import "encoding/json"

type T string

func (t T) Error() string { return "" }

func (t *T) UnmarshalJSON([]byte) error { return nil }

func main() {
	var t1 T
	var t2 = &t1
	var myError error
	var myUnmarshaler json.Unmarshaler
	var t1Slice = []T{t1}
	var t2Slice = []*T{t2}
	var t1Map = map[int]T{0: t1}
	var t2Map = map[int]*T{0: t2}
	const t1Const T = ""

	// 第一行-局部变量
	_ = t1.Error()
	_ = t1.UnmarshalJSON(nil)
	// 第一行-slice
	_ = t1Slice[0].Error()
	_ = t1Slice[0].UnmarshalJSON(nil)
	// 第一行-map
	_ = t1Map[0].Error()
	// _ = t1Map[0].UnmarshalJSON(nil) // 报错
	// 第一行-const常量
	_ = t1Const.Error()
	// _ = t1Const.UnmarshalJSON(nil)  // 报错

	// 第二行-局部变量
	_ = t2.Error()
	_ = t2.UnmarshalJSON(nil)
	// 第二行-slice
	_ = t2Slice[0].Error()
	_ = t2Slice[0].UnmarshalJSON(nil)
	// 第二行-map
	_ = t2Map[0].Error()
	_ = t2Map[0].UnmarshalJSON(nil)

	// 第三行
	myError = t1
	// myUnmarshaler = t1 // T初始化的变量t1不能实现接口

	// 第四行
	myError = t2
	myUnmarshaler = t2

	println(myError, myUnmarshaler)
}
