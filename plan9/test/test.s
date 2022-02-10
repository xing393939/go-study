TEXT ·Test(SB),$0-8
    MOVQ $10, AX
    MOVQ $6, BX
    LEAQ 4(BX)(AX*2), CX
    MOVQ CX, ret+0(FP)
    RET
