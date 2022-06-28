### Go 进阶训练营

#### 开营直播
* [scaling memcache at Facebook](https://tech.ipalfish.com/blog/2020/04/07/fb-memcache/)
* 缓存选型：写数据既要写db也写cache，binlog异步任务补偿cache
* 缓存模式1：
  * 请求1：cacheMiss->读v1->setCache
  * 请求2：writeDB(v2)->delCache
* 缓存模式2：
  * 请求1：cacheMiss->读v1->setNX
  * 请求2：writeDB(v2)->set
  * binlog异步任务补偿cache(写写场景可能会有ABA问题)
* 缓存模式3：facebook的方案
  * 请求1：cacheMiss->读v1->setCache(v1)
  * 请求2：writeDB(v2)->delCache
  * binlog异步任务补偿delCache
* 缓存模式3：facebook的方案的规则：
  1. 第一个线程cacheMiss的时候会分配leaseId(10s过期)，后续线程没有leaseId(等待并重试)
  1. setCache的时候必须带leaseId，且是有效的
  1. delCache的时候之前发的leaseId失效，后续线程cacheMiss的时候可分配leaseId
* 缓存模式-多级缓存：
  * 写数据时，上游服务先把下游服务delCache，再自己delCache
  * 设置缓存时间，下游服务的缓存时间要大于上游
* 缓存模式-热点缓存：
  * 对于小表，起一个goroutine定时读db并更新localCache
  * 框架层面，用大顶堆维护请求频繁的key，并转为小表localCache模式
  * 框架层面，用key+后缀请求到不同的节点
* 缓存模式-缓存穿透：
  * 节点上用singleflight保证只有一个请求回源  
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

#### 第六课 评论系统
* 功能模块：对接不同场景，不同场景的评论策略不同
* 架构设计1
  * comment-bff：聚合comment-svc、account-svc，filter-svc
  * comment-svc：操作redis、mysql，通过mq与comment-job通信
  * comment-job：消峰节流
  * comment-admin：操作redis、mysql
  * （bff和svc，bff是少变的，svc是多变的，二者发版节奏不一样）
* 架构设计2
  * b站单表到了几十亿才开始分表，单表几百个G是没有问题的
  * b站开始是时候是cacheMiss读db回填redis都在svc层做，高峰时goroutine又多又慢导致oom
  * 解决办法：cacheMiss用mq异步通知job去回填redis，写操作类似
  * 内部运营体系不直接操作mysql，而是依靠es
* 存储设计-mysql：分表，冗余评论数、楼层，时间用datetime类型
* 存储设计-redis：
  * comment_subject：string，key是id，value用pb格式
  * comment_index：zset，key是id+sortType，score是楼层
  * comment_content：string，key是id，value是pb格式
* 可用性设计-缓存穿透：go的singleflight
* 可用性设计-redis热点：主动发现热点，变成localCache
  * 一个timestamp字段，一个map
  * time()=timestamp时，map\[cacheKey]++
  * time()!=timestamp时，计算map的topK，即热点的cacheKey

#### 第七课 播放历史
* 功能模块
  * 平台化：视频、文章、漫画都有历史记录
  * 变更功能：添加、删除、清空
  * 读取功能：topN、获取播放进度
  * 其他功能：暂停记录，首次观看奖励
  * 高tps写入，高tps读
* 架构设计之history-svc
  * 写kafka之前，同一个uid的数据只聚合最后一条再写
  * 为了数据的实时性，写kafka也要写redis，redis是实时的
  * uid对100取余，同时批量打包，再落到kafka不同的partition
  * 读操作先读redis，再读hbase
* 架构设计之history-job
  * 读kafka得到uid和视频id，从redis取到完整数据，再batchWrite到hbase
* 架构设计之history-bff
  * 既对外提供api，也对内提供服务
  * 依赖redis判断是否是首次播放，是否是暂停记录
* 存储设计之hbase
  * 存储时每个用户只保留最近1000条记录
  * 对于历史记录翻页，先查redis，再查hbase（一次性全捞出来做分页）
  * 对于播放页的进度，先查redis，如果cacheMiss则不查hbase（不然db压力很大）
* 存储设计之redis
  * history：zset，key是uid，score是timestamp，member是视频id
  * history_content：string，key是视频id，value是视频详情
  * 判断用户是否是首次登录：用bitmap或者boomfliter，避免热点要hash成多个key
* 可用性设计-聚合
  * history-svc对同一个用户的数据只聚合最后一条再写kafka
  * bff层先聚合100请求再请求svc层
  * 请求先在边缘节点聚合打包，再经过专网传到机房
  * 与鉴权服务保持长连接，而不是每次请求都做鉴权
* 可用性设计-判断是不是首次观看
  * 在客户端维护，header头带上最后一次观看的日期
  * 如果=当前时间，则不查redis，如果!=，查一次redis，客户端下次请求带上

#### 第八课 分布式缓存和事务
* 一致性hash：
  * 初始节点ABC，ABC各有256个虚拟节点分散在环上
  * 删除节点B，A和C的虚拟节点不变，只是范围大了点
  * 增加节点D，D生成256个虚拟节点插入在环上
* [微信抢红包的优化](https://blog.csdn.net/emprere/article/details/98857800)  
  * 根据群红包id把请求哈希到同一台机器，在进程内做合并计算，然后同一入库
* 有界负载一致性hash
  * 每个节点带有当前节点的负载信息
  * 通过常规找到某个节点时，发现负载超过一个固定值，则飘到下一个节点
* redis为什么用固定槽位来sharding？方便迁移数据
* 缓存的读写场景：
  * 方案1：写的时候updateCache，坏处是如果多个线程写cache，无法保证顺序性
  * 方案2：写的时候deleteCache，坏处1是读+填的时候，填了老数据；坏处2是缓存击穿
    * 坏处2的解决1：deleteCache的时候没有真正删除，保留10s
    * 坏处2的解决2：缓存击穿的时候，用singlefly保证只有一个线程读db
* 缓存击穿的解决：
  * singlefly
  * 分布式锁，（不建议）
  * 用mq异步回填cache
  * Facebook的方案，lease租约
* redis技巧：key名不要长，value尽量用int，大hash类型拆分key，value不能太大
* 缓存穿透（恶意用大量不存在的id来请求）
  * 空缓存设置
  * 自增id加密成字符串，这样用户模拟的id就是无效id
  * 用布隆过滤器，false就是false，true可能wrong
* 抽奖：把中奖奖池预置到list结构里面，来一个pop一个
* 分布式事务（支付宝-1w，余额宝+1w）
  * 支付宝事务更新money表和插入queue表
  * Polling publisher，余额宝定时poll queue表
  * Transaction log tailing，订阅queue表的binlog并转mq通知余额宝
  * 2PC Message Queue
* 2PC Message Queue之写入消息
  * 1，对mq发prepare请求
  * 2，支付宝-1w
  * 3，对mq发commit请求
  * 如果第三步失败了，mq会定时询问service
* 2PC Message Queue之消费消息：处理完毕后ack消息
* TCC：try、confirm、cancel

#### 第九课 网络编程
* 网络通讯协议
  * socket抽象层：建立/接收连接，读写和关闭连接，超时
* go实现网络编程
* goim长连接网关
  * 客户端httpdns，客户端pull多个ip来选择服务的ip
  * 客户端心跳保活：定时发tcp包，断线重连
  * 服务端comet：bff层，和客户端建立长连接
  * 服务端logic：通过mq和comet通信，通过grpc和具体的service交互
* id分布式生成器
  * 方案1：N个genID的server，定时向db索要一批号段
  * 方案2：snowflake：1b不用，41b是毫秒时间戳，10b是机房+机器id，12b保证随机性
  * 方案3：sonyflake：39b是时间戳(单元是10毫秒)，8b保证随机性，16b是机房+机器id
* im私信协议
  * 读扩散：N个群，读的时候去N个表拉取数据
  * 写扩散：N个群，写的时候写到用户的表下，读的时候读用户的表
  * 微信群最多1k人，适合用写扩展，而且kvdb的lsm，对写友好，读性能不佳

#### 第十课 日志指标和链路追踪
* 日志
  * Fatal会调用os.Exit，不建议使用
  * Error要么直接抛出，要么就处理它，不要既抛出又记录日志
  * Info日志不应该有关闭选项
  * Debug可以有关闭选项
  * ELK方案：客户端用filebeat/fluent->kafka->logstash->es
* 链路追踪
  * 日志里面也带上traceId和spanId，这样即使APM是采样收集，通过日志也能还原trace
* 指标
  * 监控系统的4个黄金指标分别是延迟、流量、错误和饱和度
  * 饱和度：cpu/内存/硬盘/网卡使用率
  * 长尾问题：要看95、99指标
  * B的go服务是一直开着pprof，可以定位到机器，采集30秒生成火焰图

#### 第十一课 DNS和CDN和多活架构
* DNS
  * anycast：8.8.8.8其实是一个vip，背后是一组机器
* CDN如何节约成本：
  * 不差钱：各个核心城市节点都是多线机房(电信联通移动都有)
  * 节约钱：各个核心城市节点都是单线机房，CDN边缘节点支持多线
* 动态CDN加速的优点：
  * tcp优化：例如边缘节点和源站之间可以使用bbr算法的tcp连接
  * 链路优化：通过不断的测量计算得到更快更可靠的路线
  * 连接优化：边缘节点和源站之间的连接使用长连接，而不是一个请求一个连接
  * ssl offload：在边缘节点启动https，减少源站建立https的计算压力
* 单活架构：只有一个机房
* 双活架构：两个机房，但一个机房是master数据库，一个机房是slave数据库
* 多活系统的几个术语：
  * RPO(Recovery point object)：故障时间点-已同步的数据时间点，秒级
  * RTO(Recovery target object)：系统恢复正常时间点-故障时间点，分钟级
  * WRT(Work recovery time)：异常数据修复完成时间点-系统恢复正常时间点，小时级
* 蚂蚁金服的架构：假设有3个城市
  * GlobalZone：只有1个写副本在城市1，其他2城市是CityZone(读副本)
  * RegionZone：三地五机房，城市1和城市2都是两副本，城市3一个副本，可以多读多写
* 饿了么多活：
  * 根据地区划分多个EZone，每个EZone是一个独立运行的单元
  * 各个EZone包含全量的数据，当一个EZone故障，可转移到另一个新EZone
  * EZone故障时，故障EZone和新EZone若有某些用户状态不一致，锁定他们直到恢复
* 饿了么的ShardingKey设计
  * 把地理位置、订单id、商务id、骑手id转换成核心ShardingKey
  * 再根据核心ShardingKey和路由规则，路由到相应的EZone
  * 以上逻辑部署在边缘节点上，由边缘节点直接路由到对应EZone
* 阿里的多活
  * 买家数据用RegionZone，卖家数据用GlobalZone，因为买家多
  * 买家在注册时就决定了属于哪个RegionZone
  * 多个机房+中心机房：机房1的写操作只同步给中心机房，由中心机房再广播给其他机房
  * 同城容灾：同城有2个RegionZone，一个坏了直接切到另一个，要求写数据时两个都成功才算成功
  * 异地容灾：一个城市坏了直接切到另一个，要求写数据时两个都成功才算成功
  * 阿里为了强一致性，牺牲了可用性(每次请求的延时长一些)
* 苏宁的多活
  * 多机房之间是异步复制，对比阿里，可用性好，但是存在一致性问题，灾难后需要修复数据
  * Proxy服务：例如秒杀时访问Global资源，可以先获取一批资源到本地机房，减少跨机房调用
* facebook的缓存一致性思路：A机房有masterDB，B机房是slaveDB，两机房共用cache
  * 问题：B机房writeDB，此时cache是老数据，即便deleteCache，B机房读slaveDB可能仍然是老数据
  * 解决：B机房writeDB后设置marker，读cache时发现有marker强制读masterDB，slaveDB收到数据后删除marker
* 微信朋友圈评论的因果一致性问题
  * 微信的多活类似苏宁，不同机房是异步复制
  * 问题：A机房对某动态发出评论1，B机房针对此评论发出评论2，C机房收到的顺序是评论2、评论1
  * 解决：
    * 同机房的评论id一定是递增的
    * 发出评论时一定要比当前所有可见的评论id都要大
    * 同步数据时，把此动态的所有评论打包同步(因为一个动态的所有评论不会太多)
* B站的账号服务
  * 账号是Global资源，所以注册逻辑只能路由到唯一机房
  * 但是登录/查询用户信息，可以访问读副本，登录查不到用户可以去唯一机房再查一次

#### 第十二课 消息队列
* kafka基础概念
  * broker：机器节点
  * topic：类似数据库的表
  * partition：topic可以分为多个partition，单partition内的消息是顺序的
* kafka的存储原理：
  * [Kafka文件存储机制那些事](https://tech.meituan.com/2015/01/13/kafka-fs-design-theory.html)
  * 利用顺序io和pageCache达成超高吞吐
  * 保留策略：保留xx天，保留xx文件大小
  * 每个partition对应一个文件夹，开始只有一个index和log文件，当log文件满1G了写新的index和log文件
* 分配partition的机制：
  * 指明了partition：直接分配
  * 指明了key：对key哈希得到partition
  * 都没有指定：每次生成一个随机整数，求余得到partition
  * （b站的partition机制仍然会有局部热点问题，针对热点key还是要随机一下，然后在db层兼容下）
* producer和consumer
  * 发送消息的幂等性(不建议使用)：broker会提供一个pid，producer发送的消息id是pid+partition+seq
  * 单消费组，假设4个partition，消费组有4个consumer正好一一对应，再增加consumer也收不到消息
    * 增加消费能力可以让consumer多线程消费，但是提交的offset是最小已消费的offset
  * 多消费组，增加一个消费组，即可全量镜像的消费消息了
* 消费组内的rebalance：新增或者删除consumer时，消息会重新均衡给剩下的consumer
  * broker通过心跳检测consumer是否在线
  * “避免活锁”：broker发现consumer很久没有拉取消息，或者消费速度很低就标记下线
  * “避免活锁”的坑：b站写hdfs时某节点消费慢导致被摘除，其他节点消费也更慢了，出现恶性循环
* consumer如何提交offset
  * 先提交再处理消息，“at most once”，可能丢消息
  * 先处理消息再提交，“at least once”，消息可能重复，处理时要做幂等
* 数据的可靠性级别，producer.required.acks
  * 0，producer不等待broker的ack，性能是最好的
  * 1，producer只等待leader落盘成功的ack，如果follower还没有同步完leader就挂了会丢数据
  * -1，producer等待leader和所有follower落盘后的ack，如果发送ack时leader挂了会重复数据
* 数据的可靠性，acks=-1时，默认的min.insync.replicas=1，意味着只需要一个ISR成功即可
  * 所以需要增大此值，值越大可靠性越好，性能降低
  * 若异常情况ISR的总数变得小于min.insync.replicas，发送消息会报错

#### 第十三课 语言实践
* 协程在等channel的时候gopark，再次唤醒的时候可能不在原来的M上运行，亲密性不够好
* go1.15增加runnext字段，让这种协程gopark时放在runnext，下次会优先运行它
* 每个P有256大小的运行队列，还有local freelist，存放回收的协程，以便重复利用
* 全局也有with stack的协程队列和with no stack的协程队列，以便重复利用
* GC的一些概念
  * mutator：赋值器
  * 根对象：不需要其他对象就能访问到的对象，如全局对象，栈上的对象
* GC并发有两层含义：
  * mark和sweep过程是多协程并发的，1.13是stw再mark(多协程并发)
  * mutator和collector是同时运行的
* go1.5可以在mark阶段可以并发，不过需要一小段的stw做准备工作和栈的re-scan
  * 准备工作，通知进入插入写屏障
  * re-scan，重新扫描栈
* go1.8混合写屏障：
  * 栈如果增加或者删除堆的一个引用，会把它置为灰色
  * mark阶段堆上新产生的对象直接标黑
* 协程g1写channel阻塞，g2读channel唤醒g1
  * g1把创建一个sudug(g=g1，elem=value)挂在channel的sendq上
  * g2读完channel，唤醒g1，所以称为协作式调度
* 协程g2读channel阻塞，g1写channel唤醒g2
  * 流程同上，不同在于g1直接写到g2栈的那个接收变量了  
  