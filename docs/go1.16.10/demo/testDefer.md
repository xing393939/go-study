```go
//go:nosplit
func testUnnamed() int {
	i := 1
	defer func() {
		i = 2
	}()
	return i
}
// go tool compile -S -N -l main.go |grep -A 42 '.testUnnamed '
"".testUnnamed STEXT nosplit size=127 args=0x8 locals=0x68 funcid=0x0
	0x0000 00000 (main.go:26)	SUBQ	$104, SP
	0x0004 00004 (main.go:26)	MOVQ	BP, 96(SP)
	0x0009 00009 (main.go:26)	LEAQ	96(SP), BP                       // 设置BP
	0x000e 00014 (main.go:26)	MOVQ	$0, "".~r0+112(SP)               // r0 = 0
	0x0017 00023 (main.go:27)	MOVQ	$1, "".i+8(SP)                   // i = 1
	0x0020 00032 (main.go:28)	MOVL	$8, ""..autotmp_2+16(SP)         // _defer.siz = 8
	0x0028 00040 (main.go:28)	LEAQ	"".testUnnamed.func1·f(SB), AX
	0x002f 00047 (main.go:28)	MOVQ	AX, ""..autotmp_2+40(SP)         // _defer.fn = func1
	0x0034 00052 (main.go:28)	LEAQ	"".i+8(SP), AX
	0x0039 00057 (main.go:28)	MOVQ	AX, ""..autotmp_2+88(SP)         // _defer.varp = &i
	0x003e 00062 (main.go:28)	LEAQ	""..autotmp_2+16(SP), AX
	0x0043 00067 (main.go:28)	MOVQ	AX, (SP)                         // arg0 = &_defer
	0x0047 00071 (main.go:28)	CALL	runtime.deferprocStack(SB)       // 加入defer链 
	0x004c 00076 (main.go:28)	TESTL	AX, AX
	0x004e 00078 (main.go:28)	JNE	111
	0x0050 00080 (main.go:28)	JMP	82
	0x0052 00082 (main.go:31)	MOVQ	"".i+8(SP), AX
	0x0057 00087 (main.go:31)	MOVQ	AX, "".~r0+112(SP)               // r0 = i
	0x005c 00092 (main.go:31)	XCHGL	AX, AX
	0x005d 00093 (main.go:31)	NOP
	0x0060 00096 (main.go:31)	CALL	runtime.deferreturn(SB)          // 执行defer链 
	0x0065 00101 (main.go:31)	MOVQ	96(SP), BP                       // 恢复BP
	0x006a 00106 (main.go:31)	ADDQ	$104, SP                         // 恢复栈
	0x006e 00110 (main.go:31)	RET                                      // 返回
	0x006f 00111 (main.go:28)	XCHGL	AX, AX
	0x0070 00112 (main.go:28)	CALL	runtime.deferreturn(SB)
	0x0075 00117 (main.go:28)	MOVQ	96(SP), BP
	0x007a 00122 (main.go:28)	ADDQ	$104, SP
	0x007e 00126 (main.go:28)	RET

//go:nosplit
func testNamed() (i int) {
	i = 1
	defer func() {
		i = 2
	}()
	return i
}
// go tool compile -S -N -l main.go |grep -A 42 '.testNamed '
"".testNamed STEXT nosplit size=114 args=0x8 locals=0x60 funcid=0x0
	0x0000 00000 (main.go:17)	SUBQ	$96, SP
	0x0004 00004 (main.go:17)	MOVQ	BP, 88(SP)
	0x0009 00009 (main.go:17)	LEAQ	88(SP), BP                     // 设置BP
	0x000e 00014 (main.go:17)	MOVQ	$0, "".i+104(SP)               // i = 0
	0x0017 00023 (main.go:18)	MOVQ	$1, "".i+104(SP)               // i = 1
	0x0020 00032 (main.go:19)	MOVL	$8, ""..autotmp_1+8(SP)        // _defer.siz = 8 
	0x0028 00040 (main.go:19)	LEAQ	"".testNamed.func1·f(SB), AX  
	0x002f 00047 (main.go:19)	MOVQ	AX, ""..autotmp_1+32(SP)       // _defer.fn = func1  
	0x0034 00052 (main.go:19)	LEAQ	"".i+104(SP), AX
	0x0039 00057 (main.go:19)	MOVQ	AX, ""..autotmp_1+80(SP)       // _defer.varp = &i
	0x003e 00062 (main.go:19)	LEAQ	""..autotmp_1+8(SP), AX
	0x0043 00067 (main.go:19)	MOVQ	AX, (SP)                       // arg0 = &_defer   
	0x0047 00071 (main.go:19)	CALL	runtime.deferprocStack(SB)     // 加入defer链  
	0x004c 00076 (main.go:19)	TESTL	AX, AX
	0x004e 00078 (main.go:19)	JNE	98
	0x0050 00080 (main.go:19)	JMP	82
	0x0052 00082 (main.go:22)	XCHGL	AX, AX
	0x0053 00083 (main.go:22)	CALL	runtime.deferreturn(SB)
	0x0058 00088 (main.go:22)	MOVQ	88(SP), BP
	0x005d 00093 (main.go:22)	ADDQ	$96, SP
	0x0061 00097 (main.go:22)	RET
	0x0062 00098 (main.go:19)	XCHGL	AX, AX
	0x0063 00099 (main.go:19)	CALL	runtime.deferreturn(SB)
	0x0068 00104 (main.go:19)	MOVQ	88(SP), BP
	0x006d 00109 (main.go:19)	ADDQ	$96, SP
	0x0071 00113 (main.go:19)	RET
```