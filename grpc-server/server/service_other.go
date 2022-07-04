package main

import (
	"go-study/grpc-server/proto"
	"io"
)

// 一个服务器端流式rpc
func (s *service) HelloWorldServerStream(req *proto.HelloRequest, srv proto.HelloService_HelloWorldServerStreamServer) error {
	res := &proto.HelloResponse{
		Response: req.Request,
	}
	_ = srv.Send(res)
	_ = srv.Send(res)
	return nil
}

// 一个客户端流式rpc
func (s *service) HelloWorldClientStream(srv proto.HelloService_HelloWorldClientStreamServer) error {
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
func (s *service) HelloWorldClientAndServerStream(srv proto.HelloService_HelloWorldClientAndServerStreamServer) error {
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
