### Go 汇编相关

#### 常用指令
* [肝了一上午golang之plan9入门](https://studygolang.com/articles/33163)

```
//数据copy
MOVQ $123, AX  //AX=123
MOVB $1, DI    // 1 byte
MOVW $0x10, BX // 2bytes
MOVD $1, DX    // 4 bytes
MOVQ $-10, AX  // 8 bytes

//计算指令
ADDQ AX, BX  // BX += AX
SUBQ AX, BX  // BX -= AX
IMULQ AX, BX // BX *= AX

//跳转指令
JMP addr   // 跳转到地址，地址可为代码中的地址 不过实际上手写不会出现这种东西
JMP label  // 跳转到标签 可以跳转到同一函数内的标签位置
JMP 2(PC)  // 以当前置顶为基础，PC+2
JMP -2(PC) // 以当前置顶为基础，PC-2
JNZ target // 如果zero flag被set过，则跳转
CMPQ SI CX 
JLS 0x0185 // 如果SI<CX，则跳转到0x0185

//数据传输和

//栈扩大和缩小（没有用push、pop）
SUBQ $0x18, SP // 对SP做减法，为函数分配函数栈帧
ADDQ $0x18, SP // 对SP做加法，清除函数栈帧


```



