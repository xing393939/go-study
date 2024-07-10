### 细节原理

<link rel="stylesheet" type="text/css" href="../images/jquery.dialog.css">
<script type=text/javascript src="../images/jquery.dialog-code.js"></script>

#### 对齐规则
* [Go 夜读-内存对齐](https://www.bilibili.com/video/BV1iZ4y1j7TT)
* [在 Go 中恰到好处的内存对齐](https://eddycjy.gitbook.io/golang/di-1-ke-za-tan/go-memory-align)
* 编译器默认对齐长度：#pragma pack(n)，一般是8
* 对齐规则：
  * 结构体的成员变量，第一个成员的偏移量为0。之后的成员的偏移量需要是min(编译器默认, 当前成员的类型长度)的倍数
  * 结构体本身的类型长度=max(编译器默认, 所有成员的类型长度的最大值)
* [在线测试代码](https://go.dev/play/p/25Pr9TmqW-C)

#### 如何高效地拼接字符串
* var b strings.Builder && b.WriteString("asong") && b.string()

<div class="DialogCode" data-code="strings/WriteString"></div>

#### defer 的执行顺序
* [Go 语言笔试面试题](https://geektutu.com/post/qa-golang-1.html)
* 需要注意有名返回值在defer中可以被修改，如func testNamed() (i int)

<div class="DialogCode" data-code="demo/testDefer"></div>

#### 空结构体
* 空结构体不占内存，指针指向runtime.zerobase，三个常见用途
  * 定义集合set，value可以是空结构体
  * 使用channel的信号，但是不需要值
  * 只包含方法的结构体：type Lamp struct{}
  
#### 比较两个interface
* 两种情况下interface相等
  * 两个interface均等于nil，此时类型T和值V都是unset
  * 类型T相同，且对应的值V相等

```go
//go:nosplit
func testEq(a interface{}, b interface{}) bool {
	return a == b
}

TEXT    "".testEq(SB), NOSPLIT|ABIInternal, $40-40
        SUBQ    $40, SP
        MOVQ    BP, 32(SP)
        LEAQ    32(SP), BP          // 设置栈和保存BP
        MOVB    $0, "".~r2+80(SP)   // r2 = 0
        MOVQ    "".b+64(SP), AX
        CMPQ    "".a+48(SP), AX     // 比较a._type和b._type  
        SETEQ   AL
        JEQ     testNamed_pc36      // 相等跳转 
        JMP     testNamed_pc91      // 不相等跳转
testNamed_pc36:
        MOVQ    "".a+48(SP), AX
        MOVQ    "".a+56(SP), CX
        MOVQ    "".b+72(SP), DX
        MOVQ    AX, (SP)
        MOVQ    CX, 8(SP)
        MOVQ    DX, 16(SP)
        CALL    runtime.efaceeq(SB) // efaceeq(a._type, a.value, b.value)
        MOVBLZX 24(SP), AX
        JMP     testNamed_pc77
testNamed_pc77:
        MOVB    AL, "".~r2+80(SP)
        MOVQ    32(SP), BP
        ADDQ    $40, SP
        RET
testNamed_pc91:
        JMP     testNamed_pc77
        RET
```

#### 和nil比较
```go
//go:nosplit
func testCompareNil(a []int) bool {
	return a == nil
}
TEXT    "".testCompareNil(SB), NOSPLIT|ABIInternal, $0-32
        MOVB    $0, "".~r1+32(SP)  // r1 = 0
        CMPQ    "".a+8(SP), $0     // 比较a.Data和0
        SETEQ   "".~r1+32(SP)      // 设置r1
        RET
//go:nosplit
func testCompareNil(a map[int]int) bool {
	return a == nil
}
TEXT    "".testCompareNil(SB), NOSPLIT|ABIInternal, $0-16
        MOVB    $0, "".~r1+16(SP)  // r1 = 0
        CMPQ    "".a+8(SP), $0     // 比较a(*hmap)和0
        SETEQ   "".~r1+16(SP)      // 设置r1
        RET
//go:nosplit
func testCompareNil(a interface{}) bool {
	return a == nil
}
TEXT    "".testCompareNil(SB), NOSPLIT|ABIInternal, $0-24
        MOVB    $0, "".~r1+24(SP) // r1 = 0
        CMPQ    "".a+8(SP), $0    // 比较a._type和0
        SETEQ   "".~r1+24(SP)     // 设置r1
        RET
```

#### nil和interface
1. 一个接口等于nil：类型T是nil(eface和iface都适用)
  * _type和data都是nil：var a interface{} = nil或 var b error = nil
  * _type不为0data为0：var a *int && var b interface{} = a
1. 两个接口比较时，会先比较类型T，再比较值V
1. 接口与非接口比较时，会先将非接口转换成接口

#### 值接收者和指针接收者
* [非接口的任意类型T都能够调用*T的方法吗？反过来呢](https://geektutu.com/post/qa-golang-2.html#Q7-%E9%9D%9E%E6%8E%A5%E5%8F%A3%E9%9D%9E%E6%8E%A5%E5%8F%A3%E7%9A%84%E4%BB%BB%E6%84%8F%E7%B1%BB%E5%9E%8B-T-%E9%83%BD%E8%83%BD%E5%A4%9F%E8%B0%83%E7%94%A8-T-%E7%9A%84%E6%96%B9%E6%B3%95%E5%90%97%EF%BC%9F%E5%8F%8D%E8%BF%87%E6%9D%A5%E5%91%A2%EF%BC%9F)
* T初始化的变量t1，如果是“不可寻址的”，则不能调用指针接收者定义的方法，哪些是“不可寻址的”：
  * map中的元素(slice的元素可寻址)
  * const常量
  * 字符串中的字节
  * 包级别的函数
* 测试代码：go-demo/receiver

|                | func (T) Error() | func (*T) UnmarshalJSON() |
| ---            | ---              | ---                       |
| T初始化的变量t1  | 可以调用         | 要求变量是可寻址的        | 
| *T初始化的变量t2 | 可以调用         | 可以调用                  |
| T初始化的变量t1能否实现接口  | 可以 | 不能                      |
| *T初始化的变量t2能否实现接口 | 可以 | 可以                      |
