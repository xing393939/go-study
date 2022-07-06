package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go-study/grpc-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/structpb"
	"net"
)

type service struct {
	proto.HelloServiceServer
}

func (s *service) PostForm(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	res := &proto.HelloResponse{
		Response: req.Request,
		Data:     nil,
	}
	return res, nil
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

func ServerMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if header, ok := transport.FromServerContext(ctx); ok {
				println(header.RequestHeader().Get("user-agent"))
			}
			return handler(ctx, req)
		}
	}
}

func unaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, _ := metadata.FromIncomingContext(ctx)
		replyHeader := metadata.MD{}
		ctx = transport.NewServerContext(ctx, &Transport{
			reqHeader:   headerCarrier(md),
			replyHeader: headerCarrier(replyHeader),
		})
		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return handler(ctx, req)
		}
		h = middleware.Chain(ServerMiddleware())(h)
		reply, err := h(ctx, req)
		return reply, err
	}
}

func main() {
	go func() {
		lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", 8081))
		s := grpc.NewServer(grpc.ChainUnaryInterceptor(
			unaryServerInterceptor(),
		))
		proto.RegisterHelloServiceServer(s, &service{})
		_ = s.Serve(lis)
	}()

	go func() {
		lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
		s := http.NewServer(http.Middleware(
			ServerMiddleware(),
		))
		proto.RegisterHelloServiceHTTPServer(s, &service{})
		_ = s.Serve(lis)
	}()

	for {
	}
}
