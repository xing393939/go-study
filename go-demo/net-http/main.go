package main

import (
	"io"
	"net"
	"net/http"
	"time"
)

func server() {
	addr := ":803"
	server := http.Server{
		Addr: addr,
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello"))
	})
	// net.Listen最终会执行系统调用socket、bind、listen
	listener, _ := net.Listen("tcp", addr)
	// 每accept一个连接就会起一个协程
	_ = server.Serve(listener)
}

func client() {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	// 最终执行net.http.(*Transport).roundTrip(&Request)
	resp, _ := client.Get("http://httpbin.org/get?a=1")
	body, _ := io.ReadAll(resp.Body)
	println(string(body))
}

func main() {
	client()
}
