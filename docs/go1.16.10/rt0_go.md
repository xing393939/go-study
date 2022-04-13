```go
TEXT runtime·rt0_go<ABIInternal>(SB),NOSPLIT,$0
	// 初始化SP栈
	MOVQ	DI, AX		// argc
	MOVQ	SI, BX		// argv
	SUBQ	$(4*8+7), SP
	ANDQ	$~15, SP
	MOVQ	AX, 16(SP)
	MOVQ	BX, 24(SP)

	// 初始化g0
	MOVQ	$runtime·g0(SB), DI
	LEAQ	(-64*1024+104)(SP), BX
	MOVQ	BX, g_stackguard0(DI)
	MOVQ	BX, g_stackguard1(DI)
	MOVQ	BX, (g_stack+stack_lo)(DI)
	MOVQ	SP, (g_stack+stack_hi)(DI)

	// 省略检查CPU的相关代码
	...

	// 初始化m0和m0.tls
	LEAQ	runtime·m0+m_tls(SB), DI
	CALL	runtime·settls(SB)

	// 检查m0.tls是否生效
	get_tls(BX)
	MOVQ	$0x123, g(BX)
	MOVQ	runtime·m0+m_tls(SB), AX
	CMPQ	AX, $0x123
	JEQ 2(PC)
	CALL	runtime·abort(SB)
ok:
	get_tls(BX)
	LEAQ	runtime·g0(SB), CX
	MOVQ	CX, g(BX)		    // 设置m0.tls[0] = &g0
	LEAQ	runtime·m0(SB), AX

	// save m->g0 = g0
	MOVQ	CX, m_g0(AX)
	// save m0 to g0->m
	MOVQ	AX, g_m(CX)

	MOVL	16(SP), AX		    // copy argc
	MOVL	AX, 0(SP)
	MOVQ	24(SP), AX		    // copy argv
	MOVQ	AX, 8(SP)
	CALL	runtime·args(SB)
	CALL	runtime·osinit(SB)
	CALL	runtime·schedinit(SB)

	// 创建main协程
	MOVQ	$runtime·mainPC(SB), AX
	PUSHQ	AX
	PUSHQ	$0
	CALL	runtime·newproc(SB)
	POPQ	AX
	POPQ	AX

	// 启动m0
	CALL	runtime·mstart(SB)

	CALL	runtime·abort(SB)	// mstart()不返回，应该不会到这里
	RET
```