```go
func newproc(siz int32, fn *funcval) {
	argp := add(unsafe.Pointer(&fn), sys.PtrSize) //argp指针指向第一个参数的位置
	gp := getg()           //caller的g
	pc := getcallerpc()    //caller的pc
	systemstack(func() {
		newg := newproc1(fn, argp, siz, gp, pc)

		_p_ := getg().m.p.ptr()
		runqput(_p_, newg, true) //放入运行队列

		if mainStarted {
			wakep()
		}
	})
}
```