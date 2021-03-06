// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.1

package proto

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type HelloServiceHTTPServer interface {
	HelloWorld(context.Context, *HelloRequest) (*HelloResponse, error)
	PostForm(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterHelloServiceHTTPServer(s *http.Server, srv HelloServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/HelloWorld", _HelloService_HelloWorld0_HTTP_Handler(srv))
	r.POST("/v1/PostForm", _HelloService_PostForm0_HTTP_Handler(srv))
}

func _HelloService_HelloWorld0_HTTP_Handler(srv HelloServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/proto.HelloService/HelloWorld")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.HelloWorld(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloResponse)
		return ctx.Result(200, reply)
	}
}

func _HelloService_PostForm0_HTTP_Handler(srv HelloServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/proto.HelloService/PostForm")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PostForm(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloResponse)
		return ctx.Result(200, reply)
	}
}

type HelloServiceHTTPClient interface {
	HelloWorld(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloResponse, err error)
	PostForm(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloResponse, err error)
}

type HelloServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewHelloServiceHTTPClient(client *http.Client) HelloServiceHTTPClient {
	return &HelloServiceHTTPClientImpl{client}
}

func (c *HelloServiceHTTPClientImpl) HelloWorld(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloResponse, error) {
	var out HelloResponse
	pattern := "/v1/HelloWorld"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/proto.HelloService/HelloWorld"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *HelloServiceHTTPClientImpl) PostForm(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloResponse, error) {
	var out HelloResponse
	pattern := "/v1/PostForm"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/proto.HelloService/PostForm"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
