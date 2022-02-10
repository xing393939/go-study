### Go 汇编相关

#### 常用指令[肝了一上午golang之plan9入门](https://studygolang.com/articles/33163)

```
//数据copy
LEAQ 5(AX*2), BX // BX=AX*2+5
MOVQ $123, AX    // AX=123
MOVB $1, DI      // 1 byte
MOVW $0x10, BX   // 2 bytes
MOVD $1, DX      // 4 bytes
MOVQ $-10, AX    // 8 bytes

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

//栈扩大和缩小（没有用push、pop）
SUBQ $0x18, SP // 对SP做减法，为函数分配函数栈帧
ADDQ $0x18, SP // 对SP做加法，清除函数栈帧

//定义函数
TEXT fun·Swap(SB),NOSPLIT,$0-32 
//fun是包名
//Swap是方法名
//若不指定NOSPLIT，arguments size必须指定
//$0-32表示stack frame size + arguments size

//参考https://golang.design/under-the-hood/zh-cn/part1basic/ch01basic/asm/
//FUNCDATA 和 PCDATA 指令包含了由垃圾回收器使用的信息，他们由编译器引入。
DATA divtab<>+0x00(SB)/4, $0xf4f8fcff  // 表示的是divtab<>在0偏移处有一个4字节大小的值0xf4f8fcff
...
DATA divtab<>+0x3c(SB)/4, $0x81828384
GLOBL divtab<>(SB), RODATA, $64        // 给变量divtab<>加上RODATA只读标识，并声明占用64字节（3c+4=64）
```

#### 四个伪寄存器[plan9 汇编入门](https://github.com/cch123/golang-notes/blob/master/assembly.md#%E4%BC%AA%E5%AF%84%E5%AD%98%E5%99%A8)
* SB: Static base pointer(全局静态基指针)，一般用来声明函数或全局变量
* PC: Program counter(PC 寄存器)
* FP: Frame pointer，用来标识传参、返回值。arg0+0(FP)表示第一个传参
* SP: Stack pointer(栈指针)
  * 伪SP：指向当前栈帧的局部变量的开始位置。var0-8(SP)表示第一个局部变量(var0占8B)
  * 硬件SP：函数栈真实栈顶地址
* 伪SP和硬件SP的关系：
  * 若没有本地变量：伪SP=硬件SP+8
  * 若有本地变量：伪SP=硬件SP+16+本地变量空间大小
  * 如果是手写plan9，且如果是symbol+offset(SP)形式，则表示伪SP。如果是offset(SP)则表示硬件SP。
  * 如果是go tool objdump/go tool compile -S，看到的都是硬件SP。














