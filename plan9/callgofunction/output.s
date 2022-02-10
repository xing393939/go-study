#include "textflag.h"

// func output(a,b int) int
TEXT ·Output(SB), NOSPLIT, $24-24
    MOVQ a+0(FP), DX // arg a
    MOVQ DX, 0(SP) // arg x，注意这里是硬件SP
    MOVQ b+8(FP), CX // arg b
    MOVQ CX, 8(SP) // arg y，注意这里是硬件SP
    CALL ·add(SB) // 在调用 add 之前，已经把参数都通过硬件SP搬到了函数的栈顶
    MOVQ 16(SP), AX // add 函数会把返回值放在这个位置
    MOVQ AX, ret+16(FP) // return result
    RET
