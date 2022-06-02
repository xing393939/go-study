```go
//go:nosplit
func acquirem() *m {
	_g_ := getg()
	_g_.m.locks++
	return _g_.m
}
```