```go
// 休眠M，被唤醒时绑定一个P再return
func stopm() {
	_g_ := getg()

	if _g_.m.locks != 0 {
		throw("stopm holding locks")
	}
	if _g_.m.p != 0 {
		throw("stopm holding p")
	}
	if _g_.m.spinning {
		throw("stopm spinning")
	}

	lock(&sched.lock)
	mput(_g_.m)
	unlock(&sched.lock)
	mPark()
	acquirep(_g_.m.nextp.ptr())
	_g_.m.nextp = 0
}

// 唯一休眠M的方法
//go:nosplit
func mPark() {
	g := getg()
	for {
		notesleep(&g.m.park)
		// Note, because of signal handling by this parked m,
		// a preemptive mDoFixup() may actually occur via
		// mDoFixupAndOSYield(). (See golang.org/issue/44193)
		noteclear(&g.m.park)
		if !mDoFixup() {
			return
		}
	}
}
```