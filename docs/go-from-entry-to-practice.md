### Go 语言从入门到实战

#### 参考资料
* [Go 语言从入门到实战](https://time.geekbang.org/course/intro/100024001)，视频课

#### 学习笔记
```
15，类方法，用类实例和类指针都可以，不同的是类实例的把值复制了一份
16，Go可以先写类，再写接口，而不需要改动原来的类

23，gmp模型，java一个线程的栈大小是1M，go是2K。java一个协程对应内核的一个kernel entry，go是多对多
24，sync.Mutex是互斥锁。sync.RWMutex是读写互斥锁。sync.WaitGroup：Add、Done和Wait。
25，channel让串行变并行
26，select + time.After变成限时chanel
27，channel被关闭时会让所有接收者立刻收到关闭信号
28，channel的关闭是一个广播机制
29，childCtx, cancle := context.withCancle(parentCtx)，把childCtx传给子协程
30，sync.Once适合单例模式，它只运行一次
31，只要有有一个协程完成任务就返回，使用channel buffer，这样其他协程不会阻塞
32，等待所有任务完成：使用waitGroup或者channel buffer
33，对象池：使用channel buffer，拿就是接收channel，放回就是写channel
34，sync.Pool，有点类似单例，get获取put释放，但是偶尔多实例化了几个也无所谓：https://mp.weixin.qq.com/s/6Nx7IGFU_FbM5AOdUzmvcw
35，单元测试
36，bentchmark：for i:=1; i<b.N; i++
  1. B/op：每次操作分配的内存字节数
  2. allocs/op：每次操作分配内存的次数
37，goconvey测试框架，还带web ui
38，反射编程
  1. reflect.TypeOf(struct).FieldByName()  元素类型
  2. reflect.ValueOf(struct).FieldByName() 元素值
  3. reflect.ValueOf(struct).methodByName().call() 类方法调用
39 reflect.DeepEqual可以比较slice、struct、map；通过反射把map赋值到struct上
40 unsafe.Pointer可以实现类型转换，但是不推荐；atomic的StorePointer和LoadPointer

47，性能分析工具
  1. wall time：指程序运行的总体时间，包含程序不耗用cpu，被阻塞时的时间
  2. cpu time：指程序占用cpu的时间
48，go test -benth=. -cpuprofile=1.prof -memprofile=2.prof
49，读多写少用sync.map；读多写多用Concurrent Map
50，GC友好的代码
  1. 复杂对象尽量传递引用
  2. 生成trace日志：go bench -trace=1；查看日志：go tool trace 1.out
  3. 尽量预置好大小，自动扩容是有代价的
51，字符串连接用string.Builder；字节连接用bytes.Buffer
52，面向错误的设计
  1. 隔离错误：微内核+N个plugin
  2. 冗余资源：standby或者热备
  3. 限流：令牌桶
  4. 限时：避免慢响应阻塞系统
  5. 断路器+降级
53，面向恢复的设计
  1. 健康检查
  2. let it crash：这样错误才能被发现和恢复
  3. 采用微服务，有针对性的恢复
  4. 有状态的迁移成无状态的
  5. 和客户端协商并降速
54，把服务治理和业务代码分开的两个思路，一个是用服务网格istio(缺点是有性能损失)，一个是service_decorators
```

#### atomic包的学习
```Go
package main

import "sync/atomic"

func main() {
	var numCreated int32
	numWorkers := 1000
	done := make(chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			atomic.AddInt32(&numCreated, 1)
			done <- 1
		}()
	}
	for i := 0; i < numWorkers; i++ {
		<-done
	}
	println(numCreated)
	for i := 0; i < numWorkers; i++ {
		go func() {
			/*
			 * atomic包对int的读写是两条指令，并不能保证原子性，那么它存在的意义是什么呢？
			 * 参考https://studygolang.com/articles/29579
			 * 1，兼容不同cpu架构，保证读的原子性
			 * 2，避免编译器优化成寄存器，要保证是对内存的读写
			 */
			tmp := atomic.LoadInt32(&numCreated)
			atomic.StoreInt32(&numCreated, tmp+1)
			done <- 1
		}()
	}
	for i := 0; i < numWorkers; i++ {
		<-done
	}
	println(numCreated)
	for i := 0; i < numWorkers; i++ {
		go func() {
			numCreated++
			done <- 1
		}()
	}
	for i := 0; i < numWorkers; i++ {
		<-done
	}
	println(numCreated)
}
```


