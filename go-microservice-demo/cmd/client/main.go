package main

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"

	"github.com/win5do/go-lib/errx"
	log "github.com/win5do/go-lib/logx"

	"github.com/win5do/golang-microservice-demo/pkg/api/petpb"
)

func main() {
	_ = run("localhost:9020")
}

func run(addr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	conn, err := grpc.DialContext(ctx, "dns:///"+addr,
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
		grpc.WithBlock(),
	)
	cancel()
	if err != nil {
		return errx.WithStackOnce(err)
	}

	client := petpb.NewPetServiceClient(conn)
	resp, err := client.Ping(context.Background(), &petpb.Id{
		Id: time.Now().Format("2006-01-02"),
	})
	if err != nil {
		return errx.WithStackOnce(err)
	}

	log.Infof("resp: %s", resp.Id)
	return nil
}
