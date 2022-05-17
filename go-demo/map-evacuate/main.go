package main

import (
	"fmt"
	"unsafe"
)

type hmap struct {
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    uintptr // array of 2^B Buckets. may be nil if count==0.
	oldbuckets uintptr // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr // progress counter for evacuation (buckets less than this have been evacuated)

	extra uintptr // optional fields
}

type bmap struct {
	topbits  [8]uint8
	keys     [8]uint8
	values   [8]uint8
	overflow uintptr
}

type emptyInterface struct {
	_type unsafe.Pointer
	value unsafe.Pointer
}

func getPointer(m interface{}) unsafe.Pointer {
	ei := (*emptyInterface)(unsafe.Pointer(&m))
	return ei.value
}

func main() {
	m := make(map[uint8]uint8, 9)
	m[1] = 1
	m[2] = 2
	m[3] = 3
	m[4] = 4
	hm := (*hmap)(getPointer(m))
	b0 := (*bmap)(unsafe.Pointer(hm.buckets))
	b1 := (*bmap)(unsafe.Pointer(hm.buckets + unsafe.Sizeof(bmap{})))
	fmt.Printf("%#v\n", hm)
	fmt.Printf("%#v\n", b0)
	fmt.Printf("%#v\n", b1)

	// 模拟触发搬迁
	hm.count = 14
	m[5] = 5
	fmt.Printf("%#v\n", hm)
}
