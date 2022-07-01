package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"net"
	"net/http"
	"time"
)

type serverHandler struct{}

func (sh *serverHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	w.Header().Set("server", "h2test")
	_, _ = w.Write([]byte("this is a http2 test server\n"))
}

func main() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      &serverHandler{},
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	h2c := &http2.Server{
		IdleTimeout: 1 * time.Minute,
	}
	_ = http2.ConfigureServer(server, h2c)
	lis, _ := net.Listen("tcp", ":8080")
	defer lis.Close()
	for {
		rwc, err := lis.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			continue
		}
		go h2c.ServeConn(rwc, &http2.ServeConnOpts{BaseConfig: server})
	}
}
