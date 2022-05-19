package main

import (
	"errors"
	"fmt"
	"unsafe"
)

type iface struct {
	tab  *itab
	data unsafe.Pointer
}

type eface struct {
	_type *_type
	data  unsafe.Pointer
}

type itab struct {
	inter  *interfacetype
	_type  *_type
	link   *itab
	hash   uint32 // copy of _type.hash. Used for type switches.
	bad    bool   // type does not implement interface
	inhash bool   // has this itab been added to hash?
	unused [2]byte
	fun    [1]uintptr // variable sized
}

type interfacetype struct {
	typ     _type
	pkgpath *byte
	mhdr    []int64
}

type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      uint8
	align      uint8
	fieldAlign uint8
	kind       uint8
	equal      func(unsafe.Pointer, unsafe.Pointer) bool
	gcdata     *byte
	str        int32
	ptrToThis  int32
}

func main() {
	var e1 error = errors.New("a")
	ie := (*iface)(unsafe.Pointer(&e1))
	fmt.Printf("%+v \n", ie)
	fmt.Printf("%+v \n", ie.tab)
	fmt.Printf("%+v \n", ie.tab.inter)
	fmt.Printf("%+v \n", ie.tab._type)

	var e2 interface{} = e1
	ee := (*eface)(unsafe.Pointer(&e2))
	fmt.Printf("%+v \n", ee)
}
