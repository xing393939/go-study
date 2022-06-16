package main

import "time"

type Conn struct{}
type Result struct{}

func (c *Conn) DoQuery(str string) Result {
	println(c, "start")
	time.Sleep(time.Millisecond)
	return Result{}
}

/*
参考：https://go.dev/blog/concurrency-timeouts
此代码ch是无缓冲的channel，若最后一行代码“return <-ch”还没有准备好接收，4个协程就都返回了
此时在select语法下，4个协程都会走default分支，那么main函数会hang住
解决方法：第一行代码改成ch := make(chan Result, 4)
*/
func Query(conns []Conn, query string) Result {
	ch := make(chan Result)
	for _, conn := range conns {
		go func(c Conn) {
			select {
			case ch <- c.DoQuery(query):
				println("passed")
			default:
				println("default")
			}
		}(conn)
	}
	return <-ch
}

func main() {
	conns := []Conn{
		{},
		{},
		{},
		{},
	}
	// 4个协程都会执行完成，但是只要有一个协程返回结果了，Query就可以走了
	Query(conns, "SELECT 1")

	time.Sleep(time.Second)
}
