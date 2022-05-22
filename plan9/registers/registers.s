#include "textflag.h"

// func output(int) (int, int, int)
TEXT ·Output(SB), $8-32
    MOVQ ret+8(FP), DX      // 不带 symbol，这里的 SP 是硬件寄存器 SP
    MOVQ DX, ret2+24(FP) // 第三个返回值是基于硬件SP取到传参
    MOVQ perhapsArg1+16(SP), BX // 当前函数栈大小 > 0，所以 FP 在 SP 的上方 16 字节处
    MOVQ BX, ret2+24(FP) // 第二个返回值是基于伪SP取到传参
    MOVQ arg+0(FP), AX
    MOVQ AX, ret1+16(FP)  // 第一个返回值是基于FP取到传参
    RET

// 栈结构如下图
// ------
// ret2 (8 bytes)
// ret1 (8 bytes)
// ret0 (8 bytes)
// arg0 (8 bytes)
// ------ FP
// ret addr (8 bytes)
// caller BP (8 bytes)
// ------ pseudo SP
// frame content (8 bytes)
// ------ hardware SP