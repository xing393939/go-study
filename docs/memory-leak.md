### 内存泄漏

#### 参考资料
* [Go内存泄漏？不是那么简单](https://colobu.com/2019/08/28/go-memory-leak-i-dont-think-so/)
* [一些可能的内存泄漏场景](https://gfw.go101.org/article/memory-leaking.html)

#### 协程泄漏导致的内存泄漏
1. 读写channel阻塞了
1. httpServer维护了大量连接，[链接](https://mp.weixin.qq.com/s/W4eRiTw1Hbo4MkMZTgWbag)

#### 非协程泄漏导致的内存泄漏
1. 子字符串：s0 = s1\[:50]（s0是全局变量，s1是临时变量）
  * 正确做法：s0 = (strings.Builder{}).WriteString(s1\[:50])
1. 子切片：  s0 = s1\[:50]（s0是全局变量，s1是临时变量）
  * 正确做法：copy(s0, s1\[:50])
1. 在s1\[1:3]前未重置切片内的指针元素
1. time.Ticker不再使用但是没有停止它
1. time.Ticker在for循环中不断创建，见go-demos/leak-timeTicker
1. 延迟调用导致的暂时性内存泄露，见[链接](https://gfw.go101.org/article/defer-more.html#kind-of-resource-leaking)
