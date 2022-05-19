### 接口

<link rel="stylesheet" type="text/css" href="../images/jquery.dialog.css">
<script type=text/javascript src="../images/jquery.dialog-code.js"></script>

#### 测试代码
```go
type Myintinterface interface {
	fun()
}

type Myint int

func (m Myint) fun() {}

//go:nosplit
func test2() {
	var mii Myintinterface = Myint(12)
	mii.fun()

	var mii2 Myintinterface = new(Myint)
	mii2.fun()
}
```

####  对应分析

<div class="DialogCode" data-code="demo/value_recevier"></div>