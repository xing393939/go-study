```go
type Builder struct {
	addr *Builder
	buf  []byte
}
func (b *Builder) WriteString(s string) (int, error) {
	b.copyCheck() // 检查b.addr是不是自己
	b.buf = append(b.buf, s...)
	return len(s), nil
}
func (b *Builder) String() string {
	return *(*string)(unsafe.Pointer(&b.buf))
}
```