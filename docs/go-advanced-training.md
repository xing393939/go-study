### Go 进阶训练营

#### 开营直播
* [scaling memcache at Facebook](https://tech.ipalfish.com/blog/2020/04/07/fb-memcache/)
* 缓存选型：写数据既要写db也写cache，binlog异步任务补偿cache
* 缓存模式1：
  * 请求1：cacheMiss->读v1->setCache
  * 请求2：writeDB(v2)->delCache
* 缓存模式2：
  * 请求1：cacheMiss->读v1->setNX
  * 请求2：writeDB(v2)->setEX  
  * binlog异步任务补偿cache(假设不会早于请求2的setEX)
* 缓存模式3：facebook的方案
  * 请求1：cacheMiss->读v1->setCache(v1)，若存在v2版本set失败
  * 请求2：writeDB(v2)->delCache  
  * binlog异步任务补偿setCache(v2)
* 缓存模式-多级缓存：
  * 写数据时，上游服务先把下游服务delCache，再自己delCache
  * 设置缓存时间，下游服务的缓存时间要大于上游
* 缓存模式-热点缓存：
  * 对于小表，起一个goroutine定时读db并更新localCache
  * 框架层面，用大顶堆维护请求频繁的key，并转为小表localCache模式
  * 框架层面，用key+后缀请求到不同的节点
* 缓存模式-缓存穿透：
  * 通过哈希key保证一定访问service的某节点，节点上用singlefly保证只有一个请求回源  
* 缓存技巧：
  * 大的hash key拆成多个key
  * 避免超大value

#### 第一课 微服务
* 微服务的优点
  * 拆分成小的服务，容易测试，容易维护
  * 单一职责，专注才能做好
  * 促使我们尽早的设计api和建立契约
  * 兼容性和可移植性比效率更重要
* 微服务的定义：原子服务，独立进程，隔离部署，去中心化的服务治理
* 微服务概览
  * 库组件服务化：传统项目库组件的改变需要全局重新部署，现在只需要更新对应服务
  * 按业务组织服务：you build it, you run it。开发者利用运维能力，对业务全生命周期负责
  * 去中心化：数据库、治理、语言去中心化
  * 基础设施自动化
  * 可用性、兼容性设计：隔离、容错、限流、降级
  * 接口的兼容性：发送要保守，接收要开放
* 微服务设计
  * 增加BFF层，BFF层也要拆分：按业务、按重要性
  * 按部门职能划分微服务：视频、稿件
  * 按DDD领域划分微服务：除了视频、稿件，还有创作(既需要视频也需要稿件)
* gRPC和服务发现
  * gRPC：多语言、多种消息类型(pb/json/xml)、轻量级、IDL(文件定义服务)、支持http2
  * 客户端发现模式：从注册中心获取服务的IPs，然后请求服务
  * 服务端发现模式：客户端->LB->服务端
  * Envoy属于客户端发现，但是对用户代码无侵入
  * zookeeper做服务发现的问题：要求强一致性，3个挂断1个就不可用，重新选举时间长
  * B站实现了类似Eureka的服务发现，保证最终一致性
* 多集群和多租户：
  * a->b->c，要测试b的新版本b'，设置入站header头env=testb
  * 这样就变成了a->b'->c

#### 第二课 异常处理
* error和exception
  * 不要用panic来参与业务逻辑，panic是表示factor error
* error type
  * 哨兵error：预定义的error实例
  * 避免使用哨兵error：比如io包有个io.EOF，file包依赖io包，用户使用file包可能需要依赖io包
* handing error
  * you should handing error once：要么打印日志，要么返回错误
  * error.Wrap(err, "/file/path") 包装错误
  * error.Cause() 获取原始错误和哨兵error比较
* error inspection
  * fmt.Errorf("%w /file/path", err) 包装错误
  * error.Is(err, io.EOF) 和原始错误比较
* 为什么要使用"github.com/pkg/errors"
  * 打印的时候使用%+v可以打印堆栈信息

#### 第三课 并行编程
* goroutine-让调用方决定要不要继续运行
  * ListDirectory(dir) ([]string, error) 缺点是必须扫描全部文件后再返回
  * ListDirectory(dir) chan string 缺点是要靠关闭chan来通知调用结束，但是不知道是报错还是扫描完成
  * ListDirectory(dir, func(string) error) 参考filepath.WalkDir()
* goroutine泄露的例子
  * go func(接收或者写入一个阻塞的chan)
  * select条件1是超时1秒，条件2是等待某协程的chan数据，若条件1执行，条件2的goroutine因为没人接收而阻塞
* 如果不知道协程什么时候结束或者不能控制它结束，就不要用协程
  * go 协程1和go协程2时传入一个chan，并在协程1和协程2中判断这个chan关闭时协程退出
* tracker上报逻辑
  * 方案1：go tracker.send(结束时waitGroup.Add)，在main函数有waitGroup.Done保证程序退出时都send完。缺点是起了大量协程
  * 方案2：启动一个专门的协程接收chan并send数据，同时判断ctx.Done
* 数据的发送者才能去关闭channel
* memory model的可见性
  * 线程1修改了a，线程2不一定会读到a的新值，要满足可见性则必须把a的修改同步到每个cpu
* memory model的原子性
  * 64位的机器只能保证大小在64位以内数据的原子性
* data race：当两个线程竞争同一资源时，如果对资源的访问顺序敏感，就称存在竞争条件。
* go run -race 既检查原子性，也检查可见性
* sync.Atomic：可以保证原子性和可见性
* sync.Mutex 
  * Barging模式：锁释放后，等待者或者runner都可以去抢锁
  * Handoff模式：锁释放后，把锁给等待者
* sync.errorGroup：并行执行协程ABC，ABC有一个失败则3个都退出
* sync.Pool：协程ABC都需要申请一块零时空间，不使用Pool则需要申请3份内存   
* concurrency patterns
  * https://go.dev/blog/pipelines
  * https://go.dev/blog/concurrency-timeouts
* design philosophy
* context的WithValue 里面的值必须是请求级别的，和请求是同生命周期的，比如traceId

#### 第四课 工程化实践
* 工程项目结构：
  * api放pb文件属于协议定义，configs是配置，pkg是基础库包，cmd是命令入口
  * app是业务代码
* api设计
  * [Google API Design Guide](https://google-cloud.gitbook.io/api-design-guide/)
  * 利用http的状态码
  * rpc通讯使用gRPC
  * api换版本号（破坏性变更）：修改/删除/重命名字段/服务等
  * 错误码：每个服务都有自己的错误码，调用其他服务的错误码在内部消耗不要暴露
* 配置管理
  * 方案1：Conf{A, B, C} init(Conf) 缺点是不能区分零值和默认值
  * 方案2：type Option func(*Conf) init(...Option) 可以区分零值和默认值
  * 线上配置变更：配置尽量简单，合理性检查，配置和应用版本同步上线和回滚
* 模块单元测试
  * 不用程序mock外部资源，用docker生成mysql、redis、mc等等实例

#### 第五课 微服务实战
* 隔离
  * 动静分离：接口和静态资源分离
  * 服务隔离：把表频繁更新的字段放到一个单独的表；读写分离；CQRS，服务和查询分离
  * 轻重隔离：业务按重要性分离；mq按重要性使用不同的topic
  * 热点隔离：cache热点可以把remotecache变成localcache
  * 线程隔离：使用线程池来控制线程总数，或者用信号量了控制线程总数
  * 物理隔离：使用docker来进程隔离；集群隔离；多机房
* 超时传递：服务A调服务B
  * A的剩余时间额度是100ms，调用B允许的最长时间是200ms，两者取其小100ms
  * B接收请求，得到header头grpc-timeout=100ms
  * B的rootCtx=WithTimeout(ctx, 100ms)
* 过载保护
  * 令牌桶：流入的速度是固定的，流入10个/s，桶大20，qps则是10~20
  * 漏桶：流出的速率是固定的，桶干则qps是0
  * 令牌桶和漏桶的缺点是不好确定阈值
  * 过载保护：cpu达到90%，队列长度>N
  * 过载保护的效果：一个服务的极限qps是100，当qps高于100，只服务100个请求
  * 过载保护的缺点：如果qps是1000甚至更多，过载保护依然扛不住
* 分布式限流(用在bff层)：
  * 方案1：使用redis计数，缺点是热点问题
  * 方案2：各个node设定初始qps额度，根据最近的qps去center申请增加额度
* 上游限流：
  * 发现下游错误率超过50%，就减少请求，错误率低了又加大请求
  * gutter：A集群超过阈值，增加一个小集群(便宜)并设置一个小阈值
* 客户端限流：
  * 接口第一次挂sleep 1s，第二次sleep 3s...
  * sleep时间由接口的返回字段控制
* 降级
  * 降级一般在bff层，最好是自动降级
  * 参考谷歌sre，团队设置mttr(事故持续时间)指标和sop事故处理手册
* 重试退避
  * 避免重试风暴：1.1倍重试，重试请求超过了正常请求的10%就要打住
  * http错误码：A->B->C，B调C超时，C给B是504超时，B给A是503错误，B会重试，但是A不会
* 负载均衡
  * 方案1：轮询请求，这样请求均匀，但是不同node的处理能力不同，会导致不同node的cpu负载不一
  * 方案2：p2c算法，收集客户端的cpu负载（靠response header或者health check），随机取2个node，取最优的


