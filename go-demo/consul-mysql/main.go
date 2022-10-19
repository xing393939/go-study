package main

import (
	"context"
	"log"
	"sync"

	consulapi "github.com/hashicorp/consul/api"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

var wg sync.WaitGroup

func main() {
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	consulClient, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		h := server.Default(
			server.WithHostPorts("127.0.0.1:8888"),
		)

		h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
			kv, _, _ := consulClient.KV().Get("config/v1/local", nil)
			_, _ = ctx.Write(kv.Value)
		})
		h.Spin()
	}()

	wg.Wait()
}
