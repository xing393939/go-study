```go
//go:nosplit
func releasem(mp *m) {
	_g_ := getg()
	mp.locks--
	if mp.locks == 0 && _g_.preempt {
		// restore the preemption request in case we've cleared it in newstack
		_g_.stackguard0 = stackPreempt
	}
}
```