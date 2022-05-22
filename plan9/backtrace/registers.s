#include "textflag.h"

TEXT ·ZzzPrintTrace(SB), NOSPLIT, $32-1
    MOVQ	(TLS), AX   // AX=g
    MOVQ	(AX), R9    // R9=g.stack.lo
    MOVQ	8(AX), R10  // R10=g.stack.hi
	MOVQ	BP, DX
ok111:
	MOVQ	0(DX), R8
	CMPQ    (R8), R9
	JLS	exit000
	CMPQ    R10, (R8)
	JLS	exit000
	MOVQ	R8, r8-8(SP)
	MOVQ	R9, lo-16(SP)
	MOVQ	R10, hi-24(SP)
	MOVQ	8(DX), SI
	MOVQ	SI, i-32(SP)
	CALL	·zzzPrintP(SB)
	MOVQ	r8-8(SP), DX
	MOVQ	lo-16(SP), R9
	MOVQ	hi-24(SP), R10
	JMP	ok111
exit000:
    MOVB	b+0(FP), SI
    MOVQ	SI, i-32(SP)
    CALL	·zzzPrintB(SB)
	RET
