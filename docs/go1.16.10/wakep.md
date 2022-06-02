```go
func wakep() {
	// 如果没有空闲的P
	if atomic.Load(&sched.npidle) == 0 {
		return
	}
	// 如果有spinning状态的M，或者把sched.nmspinning从0改成1失败
	if atomic.Load(&sched.nmspinning) != 0 || !atomic.Cas(&sched.nmspinning, 0, 1) {
		return
	}
	startm(nil, true)
}
```