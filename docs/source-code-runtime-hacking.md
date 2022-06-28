<link rel="stylesheet" href="../images/ideal-image-slider.css">
<link rel="stylesheet" href="../images/ideal-default-theme.css">
<script src="../images/ideal-image-slider.js"></script>
<script src="../images/ideal-iis-bullet-nav.js"></script>

### runtime/HACKING

#### 参考资料
* [runtime/HACKING.md](https://github.com/golang/go/blob/master/src/runtime/HACKING.md)
* [Golang 在 runtime 中的一些骚东西](https://www.purewhite.io/2019/11/28/runtime-hacking-translate/)

#### 调度相关
* g退出后会加入到空闲的g对象池，供后续使用
* g、m和p对象都是分配在堆上且永不释放的
* 每个m都有一个g0栈，如果在unix环境还会有一个gsignal栈(一般是8K)
* m在g0系统栈上运行的代码隐含了不可抢占的含义
* 如果想要获取当前用户的g，需要使用getg().m.curg，只用getg()返回的可能是g0或者gsignal
* 判断当前是系统栈还是用户栈，可以使用 getg() == getg().m.curg

#### 同步
* runtime最底层使用mutex是安全的，阻塞M，阻止相关联的G和P被重新调度
* mutex的[lock](https://github.com/golang/go/blob/go1.16.10/src/runtime/lock_futex.go#L46)和[unlock](https://github.com/golang/go/blob/go1.16.10/src/runtime/lock_futex.go#L110)
* rwmutex也是类似的
* notesleep和notewakeup是race-free的，如果notewakeup已经发生了，那么notesleep立即返回 
* notesleep会阻止相关联的G和P被重新调度
* notetsleepg和阻塞的系统调用一样，允许释放P
* gopark挂起当前的goroutine
* goready恢复一个被挂起的goroutine
* STW期间不会有写操作，因此读操作不需要是原子的

#### Runtime-only 编译指令
* go:systemstack 表明一个函数必须在系统栈上运行
* go:nowritebarrier 告知编译器如果以下函数包含了写屏障，触发一个错误
* go:nowritebarrierrec 告知编译器如果以下函数以及它调用的函数包含了写屏障，触发一个错误
* go:nowritebarrierrec 主要用来实现写屏障自身，用来避免死循环
* go:notinheap 适用于类型声明，表明了一个类型必须不被分配在堆内内存
  1. 被用于全局变量、栈上变量，或者堆外内存上的对象
  1. 隐式的分配在runtime中是不被允许的，如new(T)、make([]T)、append()
  1. 指向堆内内存的指针，不能转换成指向go:notinheap类型对象的指针
  1. map和channel不允许有go:notinheap类型对象
  1. 指向go:notinheap类型的指针的写屏障可以被忽略
  