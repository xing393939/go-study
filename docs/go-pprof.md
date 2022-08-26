### pprof 工具

#### Go性能和问题诊断
* [Go性能和问题诊断](https://live.geekbang.org/room/1423)
* 常见OOM原因
  * 推荐看煎鱼的公众号
  * 协程泄露：channel、lock、IO、busy loop
  * 对象泄露：指针泄露、资源关闭、定时器
* pprof是通过采样数据和采样率来还原出真实数据
* pprof线上采样会带来5%的性能开销
* pprof工具痛点：
  * 第一时间抓到现场比较难
  * 需要手工敲命令
  * 报错时需要和正常时间点的数据做diff
* pprof平台化：
  * 在ui上点一下就收集
  * 定时收集
  * 当出现异常时收集
* 红蓝差分火焰图：红色表示上升，蓝色表示衰减
* 除了pprof之外的工具：
  * atop命令
  * sysstat包：iostat、mpstat、sar、strace
  * gdb
  * perf命令
  * viewcore

#### 性能优化究竟应该怎么做
* [性能优化究竟应该怎么做](https://talkgo.org/t/topic/2127)
* pprof可以收集的：
  * on-cpu profile
  * memory
  * goroutine
* fgprof：弥补pprof不能抓休眠协程的短板
* trace：诊断运行时的bug
* perf
* OOM常见情况：
  * 锁问题：锁粒度大了，会有大量协程卡在这里
  * 高cpu：高并发时GC次数变多
  * 高内存：使用sync.Pool、少用map(销毁时要注意销毁子元素)

#### pprof指标
* [pprof性能调优](https://www.topgoer.com/%E5%85%B6%E4%BB%96/pprof%E6%80%A7%E8%83%BD%E8%B0%83%E4%BC%98.html)
* [你不知道的 Go 之 pprof](https://darjun.github.io/2021/06/09/youdontknowgo/pprof/)
* cpu指标：每隔10ms会中断一次，记录每个协程当前执行的堆栈
  * 采样频率是写死的，不能更改，官方通过测试认为是合理的
  * 采样范围是采样开始~采样结束
  * 先注册定时器，10ms发送一个SIGPROF信号给进程，开启一个写缓冲协程来收集堆栈数据
  * flat：真正消耗cpu的函数，比如函数a、函数b
  * cum：如果函数c调用了a和b，那么c的时间就是a+b
* Goroutine和ThreadCreate
  * Groutine是用户创建的协程和main协程
  * ThreadCreate是所有线程信息
  * 两者都是先STW再收集
* 内存指标：每分配512KB内存则打点当前对象，free打点对象时更新统计
  * [万字长文图解 Go 内存管理分析](https://mp.weixin.qq.com/s/rydO2JK-r8JjG9v_Uy7gXg)
  * [Go语言核心36讲51](https://itcn.blog/p/1648835511945839.html)
  * 默认程序是开启内存采样的，证明代码见：go-demo/pprof-test
  * runtime.MemProfileRate = 512*1024，每分配512KB采样一次，=1表示每次分配都采样
  * 采样数据是程序启动以来到采样点时的内存情况
  * pprof.WriteHeapProfile收集的是内存快照，要实时数据用runtime.ReadMemStats(会STW)
  * 局限性：只能记录在堆上分配且参与GC的内存，栈内存和CGO无法被记录
  * alloc_objects : 历史总分配的累计
  * alloc_space ：历史总分配累计
  * inuse_objects：当前正在使用的对象数，包括业务没有使用但未GC的
  * inuse_space：当前正在使用的内存  
* Block和Mutex
  * Block：runtime.SetBlockProfileRate(10)，阻塞耗时超过10纳秒则记录，=1则每次都记录
  * Mutex：runtime.SetMutexProfileFraction(10)，每次加锁有1/10概率被记录，=为1则每次都记录
* 三种收集方式：
  * 直接使用runtime包
  * 使用net/http/pprof暴露http服务
  * 使用go test -bench . -cpuprofile=cpu.pprof
   