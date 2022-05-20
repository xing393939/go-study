```go
"".test1 STEXT nosplit size=79 args=0x0 locals=0x28 funcid=0x0
	0x0000 00000 (test.go:12)	TEXT	"".test1(SB), NOSPLIT|ABIInternal, $40-0
	0x0000 00000 (test.go:12)	SUBQ	$40, SP
	0x0004 00004 (test.go:12)	MOVQ	BP, 32(SP)
	0x0009 00009 (test.go:12)	LEAQ	32(SP), BP
	// ...
	0x000e 00014 (test.go:13)	MOVQ	$1, "".m+8(SP)           // m := 1
	0x0017 00023 (test.go:14)	MOVQ	$1, (SP)                 // copy m's value
	0x0020 00032 (test.go:14)	CALL	`"".Myint.fun`(SB)       // m.fun()
	0x0025 00037 (test.go:16)	LEAQ	"".m+8(SP), AX
	0x002a 00042 (test.go:16)	MOVQ	AX, "".m2+24(SP)         // m2 := &m 
	0x002f 00047 (test.go:17)	TESTB	AL, (AX)
	0x0031 00049 (test.go:17)	MOVQ	"".m+8(SP), AX
	0x0036 00054 (test.go:17)	MOVQ	AX, ""..autotmp_2+16(SP)
	0x003b 00059 (test.go:17)	MOVQ	AX, (SP)                 // copy m's value
	0x0040 00064 (test.go:17)	CALL	`"".Myint.fun`(SB)       // m2.fun()
	0x0045 00069 (test.go:18)	MOVQ	32(SP), BP
	0x004a 00074 (test.go:18)	ADDQ	$40, SP
	0x004e 00078 (test.go:18)	RET
```