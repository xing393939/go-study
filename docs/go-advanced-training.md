### Go 进阶训练营

#### 开营直播
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

#### 第一课-微服务
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
  * 增加BFF层
  * 
* grpc和服务发现
* 多集群和多用户

 