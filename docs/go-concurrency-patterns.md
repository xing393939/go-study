### Go 并发模式

#### 参考资料
* [Netflex开发的熔断限流包](https://learnku.com/articles/53019)
* [Go 并发模式：超时，执行下一个](https://learnku.com/docs/go-blog/go-concurrency-patterns-timing-out-and/6584)

#### 编程军规
* Never start a goroutine without knowning when it will stop

#### http请求的超时控制
* 方案1：设置http.Client的Timeout
* 方案2：传给http.Request一个超时的context，坏处是超时ctx会关闭底层连接，[导致连接频繁创建](https://mp.weixin.qq.com/s/tHBAVX3LKvqi-02cJ2jTdQ)
* 方案3：go-demo/pattern-timeout.httpRequestWithTimeout

#### 同步调用
1. 主协程+for循环子协程：
  * context.WithCancel，主协程退出前，通知所有子协程：go-demo/parttern-httpandrpc
1. 主协程+一个短任务协程，同步等待N秒，过时不候：go-demo/pattern-timeout
  * 这种情况的变种是封装成一个方法：go-demo/pattern-timeout.httpRequestWithTimeout
1. 主协程+多个短任务协程，同步等待其中一个返回即可：go-demo/pattern-get1Fromn
1. 多个协程串行执行：singleflight
1. 多个协程只有一个能执行：sync.Once
1. sync.errgroup
  * 主协程+多个短任务协程，同步等待全部完成，若有错误返回第一个错误：go-demo/parttern-errgroup.noBreakWhenHasError
  * 主协程+for循环子协程，有一个子协程失败，则其他子协程也终止：go-demo/parttern-errgroup.breakWhenHasError
  * （注意errgroup.WithContext只用来处理失败，不要把它当作父context传给下游），[bug案例](https://blog.csdn.net/EDDYCJY/article/details/119881145)

#### 管道和撤销
* [Go 并发模式：管道和撤销](https://learnku.com/docs/go-blog/pipelines/6550)
* 平方数：sq函数传入的是<-chan，传出的也是<-chan
* fan-out：N个管道来接收上游的1个管道，并发处理数据
* fan-in ：N个管道来接收上游的N个管道，汇聚到一个新的管道
* 止损：当下游提前退出时，上游可能因没有接收者而阻塞，所以需要设置缓冲区
  * 对于gen方法：因为知道要处理的数量，所以可以准确设置缓冲区
  * 对于merge方法：因为并不知道管道内有多少数据，因此不能准确设置
* 显式取消：对于merge方法，需要有一个专门的关闭管道

#### 上下文
* [Go 并发模式：上下文](https://learnku.com/docs/go-blog/context/6545)
