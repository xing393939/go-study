package main

import (
	"fmt"
	"unsafe"
)

func main() {
	hash := make(map[string]int, 24)
	say(hash)
}

//go:noinline
func say(hash2 map[string]int) {
	hash2["3"] = 4
	fmt.Printf("%v", getHMap(hash2))
}

func getHMap(m interface{}) *hmap {
	ei := (*emptyInterface)(unsafe.Pointer(&m))
	return (*hmap)(ei.value)
}

type emptyInterface struct {
	_type unsafe.Pointer
	value unsafe.Pointer
}

type hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/gc/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra unsafe.Pointer // optional fields
}
