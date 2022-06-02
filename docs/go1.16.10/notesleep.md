```go
func notesleep(n *note) {
	gp := getg()
	if gp != gp.m.g0 {
		throw("notesleep not on g0")
	}
	ns := int64(-1)
	if *cgo_yield != nil {
		// Sleep for an arbitrary-but-moderate interval to poll libc interceptors.
		ns = 10e6
	}
	for atomic.Load(key32(&n.key)) == 0 {
		gp.m.blocked = true
		futexsleep(key32(&n.key), 0, ns)
		if *cgo_yield != nil {
			asmcgocall(*cgo_yield, nil)
		}
		gp.m.blocked = false
	}
}

// Atomically, if(*addr == val) sleep Might be woken up spuriously; that's allowed.
// Don't sleep longer than ns; ns < 0 means forever.
//go:nosplit
func futexsleep(addr *uint32, val uint32, ns int64) {
	// Some Linux kernels have a bug where futex of
	// FUTEX_WAIT returns an internal error code
	// as an errno. Libpthread ignores the return value
	// here, and so can we: as it says a few lines up,
	// spurious wakeups are allowed.
	if ns < 0 {
		futex(unsafe.Pointer(addr), _FUTEX_WAIT_PRIVATE, val, nil, nil, 0)
		return
	}

	var ts timespec
	ts.setNsec(ns)
	futex(unsafe.Pointer(addr), _FUTEX_WAIT_PRIVATE, val, unsafe.Pointer(&ts), nil, 0)
}
```