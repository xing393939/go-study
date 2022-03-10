### 全局变量

#### 基本介绍
* 源码基于Go 1.16.10
* linux系统，amd64架构
* cpu：物理核4个、逻辑核8个(cat /proc/cpuinfo)

#### 全局变量
```
// 示例代码
func main() {
    fmt.Println("Hello World")
}
// dlv断点位置：b main.main。此时查看线程一共有5个(threads命令)
runtime.allp  []*runtime.p // 共8个p
runtime.allgs []*runtime.g // 共5个g
runtime.allm  *runtime.m   // 以m.alllink为指针的单链表，共6个m。
                           // 头节点runtime.allm.id=5，尾节点是runtime.m0.id=0
runtime.sched runtime.schedt {
	nmidle: 3,             // 空闲的m有3个
	mnext: 5,              // m的id自增号
	maxmcount: 10000,      // m最多能有多少个
	midle: muintptr,       // m的空闲链表
	pidle: puintptr,       // p的空闲链表
	npidle: 7,             // 空闲的p有7个
	runq: runtime.gQueue   // 全局运行队列
    runqsize: 0,           // 全局运行队列长度
}
```

#### GMP的各自重要属性
```
g.atomicstatus 
	_Gidle = 0      // 空闲
	_Grunnable = 1  // 可运行
	_Grunning = 2   // 运行中
	_Gsyscall = 3   // 系统调用中，此时拥有m，但是和p分离，也不在运行队列中
	_Gwaiting = 4   // 阻塞中
	_Gdead = 6      // 停止运行
	_Gcopystack = 8 // 栈扩缩，不在运行队列上
	_Gpreempted = 9 // 被抢占，不在运行队列上
	_Gscan = 0x1000 // GC时间
m {	
    p     puintptr // 当前绑定的p
	nextp puintptr // 暂存的p
	oldp  puintptr // 执行系统调用之前的p
}
p.status
	_Pidle = 0  // 空闲
	_Prunning = 1  // 被m持有，正在运行
	_Psyscall = 2  // 系统调用中
	_Pgcstop = 3   // 被m持有，gc时间
	_Pdead = 4     // p不再被使用
```
