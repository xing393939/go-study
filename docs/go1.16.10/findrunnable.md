```go
// 找到一个可以运行的G，不找到就让M休眠，然后等待唤醒，直到找到一个G返回
func findrunnable() (gp *g, inheritTime bool) {
	_g_ := getg()

	// 此处和handoffp中的条件必须一致：如果findrunnable将返回G运行，则handoffp必须启动M.
top:
	// 当前m绑定的p
	_p_ := _g_.m.p.ptr()
	
	// 省略...

	// 再尝试从本地队列中获取G
	if gp, inheritTime := runqget(_p_); gp != nil {
		return gp, inheritTime
	}

	// 尝试从全局队列中获取G
	if sched.runqsize != 0 {
		lock(&sched.lock)
		gp := globrunqget(_p_, 0)
		unlock(&sched.lock)
		if gp != nil {
			return gp, false
		}
	}

	// 从网络IO轮询器中找到就绪的G，把这个G变为可运行的G
	if netpollinited() && atomic.Load(&netpollWaiters) > 0 && atomic.Load64(&sched.lastpoll) != 0 {
		if gp := netpoll(false); gp != nil { // non-blocking
			// netpoll returns list of goroutines linked by schedlink.
			// 如果找到的可运行的网络IO的G列表，则把相关的G插入全局队列
			injectglist(gp.schedlink.ptr())
			// 更改G的状态为_Grunnable，以便下次M能找到这些G来执行
			casgstatus(gp, _Gwaiting, _Grunnable)
			// goroutine trace事件记录-unpark
			if trace.enabled {
				traceGoUnpark(gp, 0)
			}
			return gp, false
		}
	}

	// Steal work from other P's.
	procs := uint32(gomaxprocs)
	// 如果其他P都是空闲的，就不从其他P哪里偷取G了
	if atomic.Load(&sched.npidle) == procs-1 {
		// Either GOMAXPROCS=1 or everybody, except for us, is idle already.
		// New work can appear from returning syscall/cgocall, network or timers.
		// Neither of that submits to local run queues, so no point in stealing.
		goto stop
	}
	
	// 如果当前的M没在自旋且空闲P的数目小于正在自旋的M个数的2倍，那么让该M进入自旋状态
	if !_g_.m.spinning && 2*atomic.Load(&sched.nmspinning) >= procs-atomic.Load(&sched.npidle) {
		goto stop
	}

	// 如果M为非自旋，那么设置为自旋状态
	if !_g_.m.spinning {
		_g_.m.spinning = true
		atomic.Xadd(&sched.nmspinning, 1)
	}
	// 随机选一个P，尝试从这P中偷取一些G
	for i := 0; i < 4; i++ { // 尝试四次
		for enum := stealOrder.start(fastrand()); !enum.done(); enum.next() {
			if sched.gcwaiting != 0 {
				goto top
			}
			stealRunNextG := i > 2 // first look for ready queues with more than 1 g
			// 从allp[enum.position()]偷去一半的G，并返回其中的一个
			if gp := runqsteal(_p_, allp[enum.position()], stealRunNextG); gp != nil {
				return gp, false
			}
		}
	}

stop:
	// 当前的M找不到G来运行。如果此时P处于 GC mark 阶段
	// 那么此时可以安全的扫描和黑化对象，和返回 gcBgMarkWorker 来运行
	if gcBlackenEnabled != 0 && _p_.gcBgMarkWorker != 0 && gcMarkWorkAvailable(_p_) {
		// 设置gcMarkWorkerMode 为 gcMarkWorkerIdleMode
		_p_.gcMarkWorkerMode = gcMarkWorkerIdleMode
		// 获取gcBgMarkWorker goroutine
		gp := _p_.gcBgMarkWorker.ptr()
		casgstatus(gp, _Gwaiting, _Grunnable)
		if trace.enabled {
			traceGoUnpark(gp, 0)
		}
		return gp, false
	}

	// Before we drop our P, make a snapshot of the allp slice,
	// which can change underfoot once we no longer block
	// safe-points. We don't need to snapshot the contents because
	// everything up to cap(allp) is immutable.
	allpSnapshot := allp

	// return P and block
	lock(&sched.lock)
	if sched.gcwaiting != 0 || _p_.runSafePointFn != 0 {
		unlock(&sched.lock)
		goto top
	}
	// 再次从全局队列中获取G
	if sched.runqsize != 0 {
		gp := globrunqget(_p_, 0)
		unlock(&sched.lock)
		return gp, false
	}

	// 将当前对M和P解绑
	if releasep() != _p_ {
		throw("findrunnable: wrong p")
	}
	// 将p放入p空闲链表
	pidleput(_p_)
	unlock(&sched.lock)

	wasSpinning := _g_.m.spinning
	// M取消自旋状态
	if _g_.m.spinning {
		_g_.m.spinning = false
		if int32(atomic.Xadd(&sched.nmspinning, -1)) < 0 {
			throw("findrunnable: negative nmspinning")
		}
	}

	// 再次检查所有的P，有没有可以运行的G
	for _, _p_ := range allpSnapshot {
		// 如果p的本地队列有G
		if !runqempty(_p_) {
			lock(&sched.lock)
			// 获取另外一个空闲P
			_p_ = pidleget()
			unlock(&sched.lock)
			if _p_ != nil {
				// 如果P不是nil，将M绑定P
				acquirep(_p_)
				// 如果是自旋，设置M为自旋
				if wasSpinning {
					_g_.m.spinning = true
					atomic.Xadd(&sched.nmspinning, 1)
				}
				// 返回到函数开头，从本地p获取G
				goto top
			}
			break
		}
	}

	// Check for idle-priority GC work again.
	if gcBlackenEnabled != 0 && gcMarkWorkAvailable(nil) {
		lock(&sched.lock)
		_p_ = pidleget()
		if _p_ != nil && _p_.gcBgMarkWorker == 0 {
			pidleput(_p_)
			_p_ = nil
		}
		unlock(&sched.lock)
		if _p_ != nil {
			acquirep(_p_)
			if wasSpinning {
				_g_.m.spinning = true
				atomic.Xadd(&sched.nmspinning, 1)
			}
			// Go back to idle GC check.
			goto stop
		}
	}

	// 再次检查netpoll
	if netpollinited() && atomic.Load(&netpollWaiters) > 0 && atomic.Xchg64(&sched.lastpoll, 0) != 0 {
		if _g_.m.p != 0 {
			throw("findrunnable: netpoll with p")
		}
		if _g_.m.spinning {
			throw("findrunnable: netpoll with spinning")
		}
		gp := netpoll(true) // block until new work is available
		atomic.Store64(&sched.lastpoll, uint64(nanotime()))
		if gp != nil {
			lock(&sched.lock)
			_p_ = pidleget()
			unlock(&sched.lock)
			if _p_ != nil {
				acquirep(_p_)
				injectglist(gp.schedlink.ptr())
				casgstatus(gp, _Gwaiting, _Grunnable)
				if trace.enabled {
					traceGoUnpark(gp, 0)
				}
				return gp, false
			}
			injectglist(gp)
		}
	}
	// 实在找不到G，那就休眠吧
	// 且此时的M一定不是自旋状态
	stopm()
	goto top
}
```