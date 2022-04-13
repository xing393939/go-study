```go
func wakep() {
	if atomic.Load(&sched.npidle) == 0 {
		return
	}
	// be conservative about spinning threads
	if atomic.Load(&sched.nmspinning) != 0 || !atomic.Cas(&sched.nmspinning, 0, 1) {
		return
	}
	startm(nil, true)
}
```