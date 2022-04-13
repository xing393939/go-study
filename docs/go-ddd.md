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

