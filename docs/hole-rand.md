### Go 坑之math/rand

<link rel="stylesheet" type="text/css" href="../images/jquery.dialog.css">
<script type=text/javascript src="../images/jquery.dialog-code.js"></script>

#### 案例
* 在使用mq的时候使用math/rand包的随机数来作为mq的唯一id，在请求高峰时，唯一id重复导致业务不正确

#### 重现代码
```go
// 并发1000可能出现两种情况：
// 1.panic: runtime error: index out of range 
// 2.panic: duplicate
func TestMathRand(t *testing.T) {
	m := sync.Map{}
	eg := errgroup.Group{}
	ns := rand.NewSource(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		eg.Go(func() error {
			s := int64(0)
			for i := 0; i < 100; i++ {
				s = ns.Int63()
			}
			if _,ok := m.Load(s); ok {
				panic("duplicate")
			}
			m.Store(s, struct {}{})
			return nil
		})
	}
	eg.Wait()
}
```

#### 原因
* 使用math/rand有两种方式：
  1. 直接使用rand.Int63()，随机源是全局变量globalRand = &lockedSource{src: NewSource(1)}，可以看到随机种子是1
  2. 使用rand.NewSource(mySeed) && ns.Int63()，随机源是rngSource{}
* 两者的区别是lockedSource对数据的读写加锁了，而rngSource没有，并发情况下会导致数据错乱，也就是上面代码出现的两种情况。