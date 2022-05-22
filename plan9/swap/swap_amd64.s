#include "textflag.h"

// func Swap(a,b int) (int,int)
TEXT ·Swap(SB),NOSPLIT,$0-32
    MOVQ a+0(FP), AX  // FP(Frame pointer)栈帧指针 这里指向栈帧最低位
    MOVQ b+8(FP), BX
    MOVQ BX ,ret+16(FP)
    MOVQ AX ,ret1+24(FP)
    RET
