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