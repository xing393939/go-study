package main

import "time"

type Conn struct{}
type Result struct{}

func (c *Conn) DoQuery(str string) Result {
	println(&c, "start")
	time.Sleep(time.Millisecond)
	return Result{}
}

func Query(conns []Conn, query string) Result {
	ch := make(chan Result)
	for _, conn := range conns {
		go func(c Conn) {
			select {
			case ch <- c.DoQuery(query):
				println(&c)
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
