```go
type."".Myintinterface SRODATA size=104
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 08 22 ad e1 07 08 08 14 00 00 00 00 00 00 00 00  ."..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 18 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.interequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*main.Myintinterface.+0
	rel 44+4 t=5 type.*"".Myintinterface+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".Myintinterface+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+4 t=5 type..namedata.fun-+0
	rel 100+4 t=5 type.func()+0
	
type interfacetype struct {
	typ     _type      // 第0~47字节
	pkgpath name       // 第48~55字节
	mhdr    []imethod  // 第56~103字节
}

type _type struct {
	size       uintptr
	ptrdata    uintptr
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
```