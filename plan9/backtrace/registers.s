#include "textflag.h"

TEXT ·ZzzPrintTrace(SB), NOSPLIT, $16-16
    MOVQ	(TLS), CX
	MOVQ	BP, DX
ok:
	MOVQ	0(DX), R8
	CMPQ    (R8), CX
	JLS	exit0
	MOVQ	R8, bp-8(SP)
	MOVQ	8(DX), SI
	MOVQ	SI, i-16(SP)
	CALL	·zzzPrintP(SB)
	MOVQ	bp-8(SP), DX
	JMP	ok
exit0:
    CALL	·zzzPrintLn(SB)
	RET
