package main

import (
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

type imethod struct {
	name int32
	ityp int32
}
type name struct {
	bytes *byte
}
type interfacetype struct {
	typ     _type
	pkgpath name
	mhdr    []imethod
}
type itab struct {
	inter *interfacetype
	_type *_type
	hash  uint32
	_     [4]byte
	fun   [1]uintptr
}
type _type struct {
	size       uintptr
	ptrdata    uintptr
	hash       uint32
	tflag      uint8
	align      uint8
	fieldalign uint8
	kind       uint8
	alg        uintptr
	gcdata     *byte
	str        int32
	ptrToThis  int32
}

type MsgI interface {
	MsgFun()
}

type Msg struct {
	Name string
	Age  int
}

func (m Msg) MsgFun() {
	fmt.Println(m)
}

func main() {
	var a MsgI = Msg{}
	ea := (*iface)(unsafe.Pointer(&a))
	fmt.Println(ea.tab._type)
	fmt.Println(&ea.tab.inter.typ)
}
