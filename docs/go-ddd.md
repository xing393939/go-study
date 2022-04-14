### DDD 学习

#### 参考资料
* [DDD系列五讲](https://tech.taobao.org/news/wden4k)
* [领域驱动设计在互联网业务开发中的实践](https://tech.meituan.com/2017/12/22/ddd-in-practice.html)
* [从理论到实践，一文带你掌握DDD](https://mp.weixin.qq.com/s/x4HjK8t6mPAg1vQWa3PrSg)
* [领域驱动设计](https://book.douban.com/subject/35728308/)

#### Domain Primitive
* 定义：DP是一个在特定领域里，拥有准确定义的、可自我验证的、拥有行为的Value Object
* 三个原则：
  1. 让隐性的概念显性化：电话号原本只是一个传参，起始可以显性化，它也有自己的逻辑：验证手机号、获取区号。
  1. 将隐性的上下文显性化：money可以只代表金额，也可以是包含货币类型的金额。
  1. 封装多对象行为：转账业务涉及多个对象：money、货币类型、汇率，封装后对调用方友好。
* DP和DTO(Data Transfer Object)
  * 功能：DTO用于数据传输，DP代表业务域的对象
  * 关联性：DTO没有关联性，DP数据有关联性
  * 行为：DTO无行为，DP有丰富的行为和业务逻辑
* 常见的DP使用场景：
  * 有格式限制的string：如Name、PhoneNumber、OrderNumber
  * 有限制的Integer：如OrderId要大于0，Percentage的范围是0~100
  * Double或BigDecimal：如Money、Temperature
  * 复杂的数据结构  
* PS：Domain Primitive可以说是Value Object的进阶版。

#### 三种对象类型
1. DO(Data Object)：作为数据库表的映射
1. Entity：作为领域层的入参和出参。正常业务使用的对象模型，和持久化的方式无关
1. DTO(Data Transfer Object)：作为应用层的入参和出参。CQRS里面的Query、Command、Event和Request、Response都属于DTO

#### 贫血模型
* 什么是贫血模型：模型是对数据库表的映射，仅包含实体的数据和属性，而业务逻辑都分散在service、controller、utils中
* 贫血模型的缺陷：
  * 无法保护模型对象的完整性和一致性，因为对象的所有属性是公开的
  * 很难从对象属性上看出有哪些业务逻辑，比如long类型的值是否可以为负数
  * 代码逻辑重复：比如校验逻辑很容易出现在多个地方
  * 代码的健壮性差，模型的变化可能导致从上到下的代码都要变更
  * 强依赖底层实现：比如绑定了底层数据库、网络协议、第三方api

#### 规范
1. 接口层的返回值是封装了错误码和DTO的Result，接口层负责捕获所有异常
1. 接口层的功能：
  * 网络协议的转化：比如http、grpc
  * 前置缓存：针对只读场景的优化
  * 异常处理：统一的异常捕获
  * 鉴权、session、限流、日志：一些微服务框架可能已经包含这些，或者放在独立的网关层
1. 应用层不处理异常
1. 应用层的入参严格来说是CQE：Query、Command、Event
1. ACL接口可以在应用层和领域层，也就是他们都可以使用基础设施
  * ACL接口应该在应用层还是领域层？应用层处理的是技术问题，没有业务含义

#### DDD的分层模型

| |接口层|应用层|领域层|基础设施层|
|---|---|---|---|---|
|元素|  | Application service<br/>ACL接口 | Entity <br/>Domain Primitive<br/>Domain Service<br/>Repository接口<br/>ACL接口 | ACL具体类<br/>Repository具体类 |

