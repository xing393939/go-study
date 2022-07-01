package main

import (
	"context"
	"fmt"
	"go-study/grpc-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c := proto.NewHelloServiceClient(conn)
	res, _ := c.HelloWorld(context.Background(), &proto.HelloRequest{Request: "John"})
	fmt.Println(res)
}
