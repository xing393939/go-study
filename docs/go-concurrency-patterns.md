### Go 并发模式

#### 参考资料
* []

#### 编程军规
* Never start a goroutine without knowning when it will stop

#### 同步调用
1. 主协程+for循环子协程：
  * context.WithCancel，主协程退出前，通知所有子协程：go-demo/parttern-httpandrpc
1. 主协程+一个短任务协程，同步等待N秒，过时不候：go-demo/pattern-timeout
1. 主协程+多个短任务协程，同步等待其中一个返回即可：go-demo/pattern-get1Fromn
1. 主协程+多个短任务协程，同步等待全部完成，有一个失败立即返回：


