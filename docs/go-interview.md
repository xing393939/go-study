### Go 程序员面试

#### 参考说明
* [Go 程序员面试](https://golang.design/go-questions/)
* 基于Go 1.16.10

#### 第1章 逃逸分析
* 堆上对象都会通过调用runtime.newobject分配，该函数会调用runtime.mallocgc
* Go语言没有哪个关键字可以指定变量一定在堆上
* [查看Go支持的操作系统和CPU架构](https://golang.google.cn/doc/install/source)
* [Linux内存管理：释放大块内存时的阻塞问题](http://hanyunqi.me/2020/05/17/Linux%E5%86%85%E5%AD%98%E7%AE%A1%E7%90%86%EF%BC%9A%E9%87%8A%E6%94%BE%E5%A4%A7%E5%9D%97%E5%86%85%E5%AD%98%E6%97%B6%E7%9A%84%E9%98%BB%E5%A1%9E%E9%97%AE%E9%A2%98/)
  * C标准库malloc，小于128KB用brk，大于128KB用mmap
  * 测试发现munmap一块20GB的内存，会阻塞其他线程的malloc（brk/sbrk/mmap）1秒左右。
  * 页框回收算法（PFRA）：大部分内核占用的内存不可回收，大部分用户的可以回收，被mlock标记的不可回收
  * 测试释放大块内存的阻塞：munmap、madvise + munmap、mlock + madvise + munmap
* [golang内存泄漏排查](https://lvbay.github.io/2020/01/20/golang%E5%86%85%E5%AD%98%E6%B3%84%E6%BC%8F%E6%8E%92%E6%9F%A5/)
  * runtime.Memstats的解读见下图
* [Go 应用内存占用太多，让排查？](https://eddycjy.gitbook.io/golang/di-1-ke-za-tan/why-vsz-large)  
  * 分析了runtime.schedinit->runtime.mallocinit
* [top 命令的虚拟内存（VIRT）、物理内存（RES）、共享内存（SHR）](http://xuwang.online/index.php/archives/253/)
  * 匿名内存：就是没有文件背景的内存，就是无法和磁盘进行交换的，比如堆栈。
  * 共享内存（SHR）包括：
    * 程序的代码段。
    * 动态库的代码段。
    * 通过 mmap 做的文件映射。
    * 通过 mmap 做的匿名映射，但指明了MAP_SHARED属性。
    * 通过 shmget 申请的共享内存。
* [A deep dive into the OS memory use of a simple Go program](https://utcc.utoronto.ca/~cks/space/blog/programming/GoProgramMemoryUse)
  * pmap -p 进程ID后，前四个内存段分别是：text、rodata(常量)、data(全局变量)、bss(未初始化)
  * cat /proc/进程ID/maps，vvar是内核和进程共享的数据，vdso是系统调用代码实现
* [Linux top 命令里的内存相关字段](https://liam.page/2020/07/17/memory-stat-in-TOP/)
  * 相关英文文章前三章节，[Memory - Part 1: Memory Types](https://techtalk.intersec.com/2013/07/memory-part-1-memory-types/)
  * VIRT：所有虚拟内存之和
  * RES：所有物理内存之和，不包含交换区的
  * SHR：所有物理内存中，共享内存部分
  * CODE：所有物理内存中，可执行代码部分
  * DATA：所有虚拟内存中，去除共享区域部分剩下的
  * ANON = RES - SHR

![Pointer](../images/interview/runtime.Memstats.png)

#### 第2章 延迟语句
* defer的执行过程
  * 返回值=xxx
  * 调用defer函数
  * 空的return
* [探究 Go 源码中 panic & recover 有哪些坑](https://www.luozhiyun.com/archives/627)
* 无法捕获的异常：
  * 内存溢出：_ = make([]int64, 1<<40)
  * map 并发读写
  * 栈内存耗尽，栈最大是1G
  * 开启一个nil协程：var f func() && go f
  * 所有协程都休眠了
* 可以捕获的异常：
  * 数组访问越界
  * 访问地址无效
  * 写nil的map
  * 写已经关闭的channel
  * 类型断言错误
  
#### 第3章 数据容器
* make和new的区别
  * make只能初始化map、slice、chan；new都能
  * make返回值，new返回值的指针
  * PS：用new初始化map、slice、chan后值是nil，所以禁止用来初始化map和chan
* make系列函数：必定在堆上分配
  * makechan：最终调用mallocgc，当make(chan int)时必定在堆上分配
  * makeslice：最终调用mallocgc
  * makemap：因为makemap函数返回值是*hmap，所以new(hmap)会执行newobject
* map只有扩容，没有缩容。如果桶过于稀疏，则迁移到新桶重新排列(内存不变)
* map在赋值的时候检查是否需要扩容，调用growWork来扩容。
  * 在赋值和删除时进行渐进式搬迁
  * 每次搬迁2个根bucket
* copy方法：copy(dst, src \[]Type)，如果dst容量小于src，容量不会变大

#### 第4章 通道
* 如何优雅的关闭通道：
  * 1个sender，1个receiver：sender端关闭
  * 1个sender，N个receiver：sender端关闭
  * N个sender，1个receiver：receiver端通过closeChan通知sender们退出，不关闭通道(GC来回收)
  * N个sender，N个receiver：receiver端通过closeChan通知sender们退出，不关闭通道(GC来回收)
* 读写一个nil通道将永久挂起，即使这个通道后续初始化了，设计缘由见[链接](https://groups.google.com/g/golang-nuts/c/QltQ0nd9HvE/m/VvDhLO07Oq4J)

#### 第5章 接口
* 默认所有类型都实现了空接口
* 具体类型转空接口：_type复制具体类型的type，复制值到一块新内存，data指向新内存
* 具体类型转非空接口：itab值在编译期间已经生成，复制值到一块新内存，tab指向itab，data指向新内存
* 接口转接口：
  * 空接口转非空接口：调用runtime.assertE2I，[见](https://go.godbolt.org/z/v1bf8co97)
  * 非空接口转空接口：eface._type = iface.tab._type，[见](https://go.godbolt.org/z/TezcajzP7)
  * 非空接口转非空接口：调用runtime.convI2I修改tab即可
* 动态派发(多态)：根据interfacetype.imethod\[i]找到tab.fun\[i]，然后执行[CALL AX](https://go.godbolt.org/z/xsv9Mj8fr)

| 用法 | 说明 |
| --- | --- |
| 接口A嵌入接口B    | A具有B的接口 |
| 结构体A嵌入结构体B| A具有B的属性和方法，强耦合 | 
| 结构体A嵌入接口B  | A可以使用B的方法，且[松耦合](https://blog.csdn.net/raoxiaoya/article/details/109998888) |

#### 第6章 unsafe包
* [聊一个string和[]byte转换问题](https://blog.huoding.com/2021/10/14/964)
* string转\[]byte：
  * `[]byte(string)`：最终调用runtime.stringtoslicebyte和copy，内存拷贝了一份
  * `*(*[]byte)(unsafe.Pointer(&s))`：内存零拷贝
* \[]byte转string：
  * `*(*string)(unsafe.Pointer(&b))`：内存零拷贝

#### 第7章 context包
* 因为valueCtx、cancelCtx结构体嵌入了Context接口，实现了松耦合：
  * 所以用WithValue()可以传入一个cancelCtx实例进而拥有cancelCtx能力，相反同理
* WithCancel(ctx) (childCtx, cancel)：
  * 如果ctx是cancelCtx类型，或者父级中有一个是cancelCtx类型，那么cancel后，childCtx.Done收到通知
  * 如果ctx是自定义ctx类型，那么内部额外维护一个协程监听ctx.Done和childCtx.Done
  
#### 第8章 error
* 包装错误：fmt.Errorf("%w", err)
* 解包错误：errors.Unwrap(err)，如果err包含Unwrap方法就执行，否则返回nil(只解一层)
* 断言错误：errors.As(err, &dst)
* 检查错误源自原始错误：errors.Is(err, io.EOF)

#### 第9章 定时器
* time.NewTimer和time.NewTicker都会生成一个timer挂在P上
* runtime包下：checkTimers检查已到时的timer，schedule会调用checkTimers
  * schedule的调度时机[见](https://xing393939.github.io/go-study/docs/go-language-design-and-implementation.html#%E7%AC%AC%E5%85%AD%E7%AB%A0-%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B)
  * checkTimers发现timer已到时，执行timer.f(timer.arg, timer.seq)
  * [go1.14基于netpoll优化timer定时器实现原理](https://xiaorui.cc/archives/6483) 
  * 每次添加/修改timer的时候，发现netpoll的pollerPollUntil>timer.when，调用netpollBreak打断netpoll

#### 第10章 反射
* [反射三大定律](https://go.dev/blog/laws-of-reflection)
  1. 可以把值转变成反射对象：reflect.ValueOf(i any) Value
  1. 可以把反射对象转变成值：func (v Value) Interface() (i any)
  1. 值是指针类型，则转变成反射对象后是CanSet的

#### 第11章 同步
* [runtime_Semacquire和runtime_Semrelease的分析](https://www.qetool.com/scripts/view/4193.html)
* sync.Mutex：
  * [go sync.Mutex 源码阅读](https://fuweid.com/post/2020-go-sync-mutex-insight/)
  * state int32：28b表示阻塞的G的个数，4b表示锁的状态
  * sema uint32：维护阻塞的G的队列
  * Lock：如果加锁失败，那么调用runtime_SemacquireMutex(sema)休眠
  * Unlock：解锁后，调用runtime_Semrelease唤醒第一个休眠的协程
* sync.WaitGroup：
  * Add：如果运行计数v==0，那么调用runtime_Semrelease唤醒休眠的协程
  * Wait：如果运行计数v>0，那么调用runtime_Semacquire(semap)休眠
* sync.Pool：
  * [golang的对象池sync.pool源码解读](https://zhuanlan.zhihu.com/p/99710992)
  * 取当前P：pool.Get()->pool.pin()：
    * 调用runtime_procPin()禁止P被抢占
    * pool没有初始化，调用pool.pinSlow()初始化
  * 取当前P本地的private或shared
  * 取其它P的shared
  * [无锁化编程](https://www.cnblogs.com/luozhiyun/p/14194872.html)
    * poolLocal.shared：本地的P可以pushHead/popHead，其他P只能popTail
    * poolDequeue：单生产者可以pushHead/popHead，多消费者只能popTail
* sync.Map：
  * [sync.Map源码分析](https://developer.aliyun.com/article/741441)
  * 1、空间换时间。通过冗余的两个数据结构(read、dirty)，减少加锁对性能的影响。
  * 2、使用只读数据(read)，避免读写冲突。
  * 3、动态调整，miss次数多了之后，将dirty数据提升为read。
  * 4、double-checking（双重检测）。
  * 5、延迟删除。删除一个键值只是打标记，只有在提升dirty的时候才清理删除的数据。
  * 6、优先从read读取、更新、删除，因为对read的读取不需要锁。


