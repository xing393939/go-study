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
func test1() {
	m := Myint(1) // CALL "".Myint.fun(SB)
	m.fun()

	m2 := &m
	m2.fun()      // 还是CALL "".Myint.fun(SB)
}

//go:nosplit
func test2() {
	var mii Myintinterface = Myint(12)
	mii.fun()     // CALL "".Myint.fun(SB)

	var mii2 Myintinterface = new(Myint)
	mii2.fun()    // 还是CALL "".Myint.fun(SB)
}
```

####  test1函数对应汇编

<div class="DialogCode" data-code="demo/interface_test1"></div>

####  test2函数对应汇编

<div class="DialogCode" data-code="demo/interface_test2"></div>