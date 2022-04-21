### 细节原理

<link rel="stylesheet" type="text/css" href="../images/jquery.dialog.css">
<script type=text/javascript src="../images/jquery.dialog-code.js"></script>

#### 如何高效地拼接字符串
* var b strings.Builder && b.WriteString("asong") && b.string()

<div class="DialogCode" data-code="strings/WriteString"></div>

#### defer 的执行顺序
* [Go 语言笔试面试题](https://geektutu.com/post/qa-golang-1.html)
* 需要注意有名返回值在defer中可以被修改，如func get() (i int)

<div class="DialogCode" data-code="demo/testDefer"></div>

#### 空结构体
* 空结构体不占内存，指针指向runtime.zerobase，三个常见用途
  * 定义集合set，value可以是空结构体
  * 使用channel的信号，但是不需要值
  * 只包含方法的结构体：type Lamp struct{}
  
#### 比较两个interface
* 两种情况下interface相等
  * 两个interface均等于nil，此时类型T和值V都是unset
  * 类型T相同，且对应的值V相等
  
#### nil和interface
1. 一个接口等于nil：有且仅当类型T和值V都是unset
1. 两个接口比较时，会先比较类型T，再比较值V
1. 接口与非接口比较时，会先将非接口转换成接口

#### 值接收者和指针接收者
* [非接口的任意类型T都能够调用*T的方法吗？反过来呢](https://geektutu.com/post/qa-golang-2.html#Q7-%E9%9D%9E%E6%8E%A5%E5%8F%A3%E9%9D%9E%E6%8E%A5%E5%8F%A3%E7%9A%84%E4%BB%BB%E6%84%8F%E7%B1%BB%E5%9E%8B-T-%E9%83%BD%E8%83%BD%E5%A4%9F%E8%B0%83%E7%94%A8-T-%E7%9A%84%E6%96%B9%E6%B3%95%E5%90%97%EF%BC%9F%E5%8F%8D%E8%BF%87%E6%9D%A5%E5%91%A2%EF%BC%9F)
* T初始化的变量t1，如果是“不可寻址的”，则不能调用指针接收者定义的方法，哪些是“不可寻址的”：
  * map中的元素(slice的元素可寻址)
  * const常量
  * 字符串中的字节
  * 包级别的函数
* 测试代码：go-demo/receiver

|                | func (T) Error() | func (*T) UnmarshalJSON() |
| ---            | ---              | ---                       |
| T初始化的变量t1  | 可以调用         | 要求变量是可寻址的        | 
| *T初始化的变量t2 | 可以调用         | 可以调用                  |
| T初始化的变量t1能否实现接口  | 可以 | 不能                      |
| *T初始化的变t2量能否实现接口 | 可以 | 可以                      |
