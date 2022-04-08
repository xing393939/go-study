package main

import (
	"context"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func httpServer(ctx context.Context, ch chan<- bool) {
	srv := &http.Server{
		Addr: ":" + strconv.Itoa(5000+rand.Int()%1000),
	}

	go func() {
		<-ctx.Done()
		_ = srv.Shutdown(ctx)
	}()

	_ = srv.ListenAndServe()
	println("server退出了")
	ch <- true
}

func rpcServer(ctx context.Context, ch chan<- bool) {
	httpServer(ctx, ch)
}

func main() {
	ch := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	go httpServer(ctx, ch)
	go rpcServer(ctx, ch)

	// 主程序要退出
	time.Sleep(time.Second * 5)
	cancel()
	<-ch
	<-ch
}
