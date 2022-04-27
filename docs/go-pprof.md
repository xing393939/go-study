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