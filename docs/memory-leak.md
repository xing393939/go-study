### 内存泄漏

#### 参考资料
* [Go内存泄漏？不是那么简单](https://colobu.com/2019/08/28/go-memory-leak-i-dont-think-so/)
* [一些可能的内存泄漏场景](https://gfw.go101.org/article/memory-leaking.html)

#### 协程泄漏导致的内存泄漏
1. 读写channel阻塞了
1. httpServer维护了大量连接，[链接](https://mp.weixin.qq.com/s/W4eRiTw1Hbo4MkMZTgWbag)

#### 非协程泄漏导致的内存泄漏