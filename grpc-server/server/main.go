package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go-study/grpc-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
	"net"
)

type service struct {
	proto.HelloServiceServer
}

type MyStruct struct {
	A int
	B string
}

// 正常模式
func (s *service) HelloWorld(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	m := map[string]interface{}{
		"a": 1,
		"b": []interface{}{
			map[string]interface{}{
				"type":   "home",
				"number": "212 555-1234",
			},
			map[string]interface{}{
				"type":   "office",
				"number": "646 555-4567",
			},
		},
		"c": map[string]interface{}{
			"c": "ccc",
		},
		// "d": MyStruct{A: 1, B: "b"}, 不能用自定义的结构体
	}
	mys, _ := structpb.NewStruct(m)
	res := &proto.HelloResponse{
		Response: req.Request,
		Data:     mys,
	}
	return res, nil
}

func main() {
	lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", 8081))
	s := grpc.NewServer()
	proto.RegisterHelloServiceServer(s, &service{})
	go func() {
		lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
		s := http.NewServer()
		proto.RegisterHelloServiceHTTPServer(s, &service{})
		_ = s.Serve(lis)
	}()
	_ = s.Serve(lis)
}
