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

 