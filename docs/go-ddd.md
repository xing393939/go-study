### DDD 学习

#### 参考资料
* [DDD系列五讲](https://tech.taobao.org/news/wden4k)

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

#### 规范
1. 接口层的返回值是封装了错误码和DTO的Result，接口层负责捕获所有异常
1. 应用层不处理异常
1. 应用层的入参严格来说是CQE：Query、Command、Event
1. ACL接口可以在应用层和领域层，也就是他们都可以使用基础设施
  * ACL接口应该在应用层还是领域层？应用层处理的是技术问题，没有业务含义

#### DDD的分层模型

| |接口层|应用层|领域层|基础设施层|
|---|---|---|---|---|
|元素|  | Application service<br/>ACL接口 | Entity <br/>Domain Primitive<br/>Domain Service<br/>Repository接口<br/>ACL接口 | ACL具体类<br/>Repository具体类 |

