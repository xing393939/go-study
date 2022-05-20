```go
go.itab.*"".Myint,"".Myintinterface SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 5a 95 d9 8b 00 00 00 00 00 00 00 00 00 00 00 00  Z...............
	rel 0+8 t=1 `type."".Myintinterface`+0
	rel 8+8 t=1 type.*"".Myint+0
	rel 24+8 t=1 "".(*Myint).fun+0

type itab struct {
    inter *interfacetype `type."".Myintinterface`
    _type *_type         `type.*"".Myint`
    hash  uint32
    _     [4]byte
    fun   [1]uintptr     `"".(*Myint).fun`
}
```