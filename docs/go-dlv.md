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

#### slice和map的内存结构
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



```
