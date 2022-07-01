package main

import (
	"context"
	"fmt"
	"go-study/grpc-server/proto"
	"google.golang.org/grpc"
	"io"
	"net"
)

type server struct {
	proto.HelloServiceServer
}

// 正常模式
func (s *server) HelloWorld(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	res := &proto.HelloResponse{
		Response: req.Request,
	}
	return res, nil
}

// 一个服务器端流式rpc
func (s *server) HelloWorldServerStream(req *proto.HelloRequest, srv proto.HelloService_HelloWorldServerStreamServer) error {
	res := &proto.HelloResponse{
		Response: req.Request,
	}
	_ = srv.Send(res)
	_ = srv.Send(res)
	return nil
}

// 一个客户端流式rpc
func (s *server) HelloWorldClientStream(srv proto.HelloService_HelloWorldClientStreamServer) error {
	str := ""
	for {
		req, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic("unknown error")
		}
		str += req.Request
	}
	res := &proto.HelloResponse{
		Response: str,
	}
	return srv.SendAndClose(res)
}

// 一个客户端和服务器端双向流式rpc
func (s *server) HelloWorldClientAndServerStream(srv proto.HelloService_HelloWorldClientAndServerStreamServer) error {
	for {
		req, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic("unknown error")
		}
		res := &proto.HelloResponse{
			Response: req.Request,
		}
		_ = srv.Send(res)
	}
	return nil
}

func main() {
	lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	s := grpc.NewServer()
	proto.RegisterHelloServiceServer(s, &server{})
	_ = s.Serve(lis)
}
