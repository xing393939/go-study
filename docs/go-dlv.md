### Go dlv调试

#### dlv用法
```
// 编译并进入调试模式
go build -gcflags=all="-N -l" hello.go
dlv exec ./hello
b main.main
c

// 常用命令
n                          执行go代码的下一行代码
regs                       打印所有寄存器的值
help regs                  获取regs命令的帮助说明
p var0                     打印var0
x -count 24 -size 8 0x10a0 打印0x10a0地址的值，8字节一组，共24组
```

#### dlv常用命令
![dlv](../images/dlv.jpg)

#### slice的内存结构
```
func main() {
	hash := make([]int, 9, 98)
	hash[0] = 1
	hash[1] = 2
	hash[2] = 3
	say(hash)
}
//go:noinline
func say(hash2 []int) {
	hash2[3] = 4
	_ = hash2
}
// 编译并进入调试模式后执行
b main.say
c
n
x -count 16 -size 8 %rsp
// 结果：第0~3的8B不管，第4~6的8B是slice的结构，第7~9是array数据
0xc00003c410:   0x0000000000410eed   0x000000c00005e000   0x000000c00003c770   0x0000000000459aba   0x000000c00003c448   0x0000000000000009   0x0000000000000062   0x0000000000000001   
0xc00003c450:   0x0000000000000002   0x0000000000000003   0x0000000000000000   0x0000000000000000   0x0000000000000000   0x0000000000000000   0x0000000000000000   0x0000000000000000
```

#### map的内存结构
```
func main() {
	hash := make(map[string]int, 24)
	say(hash)
}

//go:noinline
func say(hash2 map[string]int) {
	fmt.Printf("%v\n", getHMap(hash2))
}

func getHMap(m interface{}) *hmap {
	ei := (*emptyInterface)(unsafe.Pointer(&m))
	return (*hmap)(ei.value)
}

type emptyInterface struct {
	_type unsafe.Pointer
	value unsafe.Pointer
}

type hmap struct {
	count      int
	flags      uint8
	B          uint8
	noverflow  uint16
	hash0      uint32
	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr
	extra      unsafe.Pointer
}

// 编译并进入调试模式后执行
b main.say
c
// 此时rax存储的即是hmap的地址
x -count 6 -size 8 0x000000c000074150
0xc000074150:   0x0000000000000000   0x69587b7700000200   0x000000c000078a80   0x0000000000000000   0x0000000000000000   0x0000000000000000
```