```go
// mainPC is a function value for runtime.main, to be passed to newproc.
DATA	runtime·mainPC+0(SB)/8,$runtime·main<ABIInternal>(SB)

func main() {
	g := getg()
	g.m.g0.racectx = 0

	// Max stack size is 1 GB on 64-bit, 250 MB on 32-bit.
	if sys.PtrSize == 8 {
		maxstacksize = 1000000000
	} else {
		maxstacksize = 250000000
	}
	maxstackceiling = 2 * maxstacksize

	// Allow newproc to start new Ms.
	mainStarted = true

	if GOARCH != "wasm" {
		atomic.Store(&sched.sysmonStarting, 1)
		systemstack(func() {
			// 启动系统监控线程，不需要绑定p
			newm(sysmon, nil, -1)
		})
	}

	lockOSThread()

	if g.m != &m0 {
		throw("runtime.main not on m0")
	}
	m0.doesPark = true

	runtimeInitTime = nanotime()
	if runtimeInitTime == 0 {
		throw("nanotime returning zero")
	}

	if debug.inittrace != 0 {
		inittrace.id = getg().goid
		inittrace.active = true
	}

	// 执行运行时的init方法
	doInit(&runtime_inittask)

	needUnlock := true
	defer func() {
		if needUnlock {
			unlockOSThread()
		}
	}()

	gcenable()

	main_init_done = make(chan bool)
	if iscgo {
		if _cgo_thread_start == nil {
			throw("_cgo_thread_start missing")
		}
		if GOOS != "windows" {
			if _cgo_setenv == nil {
				throw("_cgo_setenv missing")
			}
			if _cgo_unsetenv == nil {
				throw("_cgo_unsetenv missing")
			}
		}
		if _cgo_notify_runtime_init_done == nil {
			throw("_cgo_notify_runtime_init_done missing")
		}
		// Start the template thread in case we enter Go from
		// a C-created thread and need to create a new thread.
		startTemplateThread()
		cgocall(_cgo_notify_runtime_init_done, nil)
	}

	// 执行用户程序包的init方法
	doInit(&main_inittask)
	inittrace.active = false

	close(main_init_done)

	needUnlock = false
	unlockOSThread()

	// 如果编译参数-buildmode=c-archive或者c-shared，就不往下走了
	if isarchive || islibrary {
		return
	}

	// 执行main包的main方法
	fn := main_main
	fn()
	if raceenabled {
		racefini()
	}

	if atomic.Load(&runningPanicDefers) != 0 {
		// Running deferred functions should not take long.
		for c := 0; c < 1000; c++ {
			if atomic.Load(&runningPanicDefers) == 0 {
				break
			}
			Gosched()
		}
	}
	if atomic.Load(&panicking) != 0 {
		gopark(nil, nil, waitReasonPanicWait, traceEvGoStop, 1)
	}

	// main协程执行系统调用的exit，会导致整个进程都退出
	exit(0)
	for {
		var x *int32
		*x = 0
	}
}
```