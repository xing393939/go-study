TEXT ·Test(SB),$0-8
    MOVQ $10, AX
    MOVQ $4, BX
    //LEAQ 4(BX)(AX*2), CX
    TESTQ AX, BX
    JNE ok
    MOVQ BX, ret+0(FP)
    RET
ok:
    MOVQ AX, ret+0(FP)
    RET
