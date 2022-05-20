```go
"".test2 STEXT nosplit size=247 args=0x0 locals=0x60 funcid=0x0
	0x0000 00000 (test.go:23)	TEXT	"".test2(SB), NOSPLIT|ABIInternal, $96-0
	0x0000 00000 (test.go:23)	SUBQ	$96, SP
	0x0004 00004 (test.go:23)	MOVQ	BP, 88(SP)
	0x0009 00009 (test.go:23)	LEAQ	88(SP), BP
	// ...
	0x000e 00014 (test.go:28)	LEAQ	go.itab."".Myint,"".Myintinterface(SB), AX
	0x0015 00021 (test.go:28)	MOVQ	AX, "".mii+72(SP)         // mii.tab = go.itab."".Myint,"".Myintinterface
	0x001a 00026 (test.go:28)	LEAQ	""..stmp_0(SB), AX        // stmp_0在data段
	0x0021 00033 (test.go:28)	MOVQ	AX, "".mii+80(SP)         // mii.data = &stmp_0
	0x0026 00038 (test.go:29)	MOVQ	$0, ""..autotmp_2+40(SP)
	0x002f 00047 (test.go:29)	MOVQ	"".mii+72(SP), AX
	0x0034 00052 (test.go:29)	MOVQ	"".mii+80(SP), CX
	0x0039 00057 (test.go:29)	LEAQ	go.itab."".Myint,"".Myintinterface(SB), DX
	0x0040 00064 (test.go:29)	CMPQ	DX, AX
	0x0043 00067 (test.go:29)	JEQ	74
	0x0045 00069 (test.go:29)	JMP	213
	0x004a 00074 (test.go:29)	MOVQ	(CX), AX
	0x004d 00077 (test.go:29)	MOVQ	AX, ""..autotmp_2+40(SP)
	0x0052 00082 (test.go:29)	MOVQ	AX, (SP)                   // arg0 = stmp_0
	0x0056 00086 (test.go:29)	CALL	"".Myint.fun(SB)           // mii.fun()
	0x005b 00091 (test.go:35)	MOVQ	$0, ""..autotmp_4+32(SP)
	0x0064 00100 (test.go:35)	LEAQ	""..autotmp_4+32(SP), AX
	0x0069 00105 (test.go:35)	MOVQ	AX, ""..autotmp_3+48(SP)
	0x006e 00110 (test.go:35)	LEAQ	go.itab.*"".Myint,"".Myintinterface(SB), CX
	0x0075 00117 (test.go:35)	MOVQ	CX, "".mii2+56(SP)         // mii2.tab = go.itab.*"".Myint,"".Myintinterface
	0x007a 00122 (test.go:35)	MOVQ	AX, "".mii2+64(SP)         // mii2.data = &autotmp_4
	0x007f 00127 (test.go:36)	MOVQ	"".mii2+56(SP), AX
	0x0084 00132 (test.go:36)	MOVQ	"".mii2+64(SP), CX
	0x0089 00137 (test.go:36)	LEAQ	go.itab.*"".Myint,"".Myintinterface(SB), DX
	0x0090 00144 (test.go:36)	CMPQ	AX, DX
	0x0093 00147 (test.go:36)	JEQ	151
	0x0095 00149 (test.go:36)	JMP	180
	0x0097 00151 (test.go:36)	TESTB	AL, (CX)
	0x0099 00153 (test.go:36)	MOVQ	(CX), AX
	0x009c 00156 (test.go:36)	MOVQ	AX, ""..autotmp_5+24(SP)
	0x00a1 00161 (test.go:36)	MOVQ	AX, (SP)                   // arg0 = autotmp_5
	0x00a5 00165 (test.go:36)	CALL	"".Myint.fun(SB)           // mii2.fun()
	0x00aa 00170 (test.go:37)	MOVQ	88(SP), BP
	0x00af 00175 (test.go:37)	ADDQ	$96, SP
	0x00b3 00179 (test.go:37)	RET
	0x00b4 00180 (test.go:36)	// panic并退出
```