```go
// This must not split the stack because we may not even have stack bounds set up yet.
// May run during STW (because it doesn't have a P yet), so write barriers are not allowed.
//go:nosplit
//go:nowritebarrierrec
func mstart() {
	_g_ := getg()

	osStack := _g_.stack.lo == 0
	if osStack {
		// 没有设置栈就设置一下
		size := _g_.stack.hi
		if size == 0 {
			size = 8192 * sys.StackGuardMultiplier
		}
		_g_.stack.hi = uintptr(noescape(unsafe.Pointer(&size)))
		_g_.stack.lo = _g_.stack.hi - size + 1024
	}
	// 初始化栈以便执行用户的G
	_g_.stackguard0 = _g_.stack.lo + _StackGuard
	// 初始化栈以便执行g0
	_g_.stackguard1 = _g_.stackguard0
	mstart1()

	// Exit this thread.
	if mStackIsSystemAllocated() {
		osStack = true
	}
	mexit(osStack)
}

func mstart1() {
	_g_ := getg()

	if _g_ != _g_.m.g0 {
		throw("bad runtime·mstart")
	}

	// Record the caller for use as the top of stack in mcall and
	// for terminating the thread.
	// We're never coming back to mstart1 after we call schedule,
	// so other calls can reuse the current frame.
	save(getcallerpc(), getcallersp())
	asminit()
	minit()

	// 当前是m0，执行mstartm0()，主要是设置系统信号量的处理函数
	if _g_.m == &m0 {
		mstartm0()
	}

	// 如果有m的起始任务函数，则执行，比如sysmon函数
	if fn := _g_.m.mstartfn; fn != nil {
		fn()
	}

	// 当前不是m0，绑定p
	if _g_.m != &m0 {
		acquirep(_g_.m.nextp.ptr())
		_g_.m.nextp = 0
	}
	schedule()
}
```