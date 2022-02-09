TEXT ·Test(SB),$0-8
    MOVQ $10, AX
    LEAQ 5(AX*2), BX
    MOVQ BX, ret+0(FP)
    RET
