// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/PostReg/PostReg.proto

package go_micro_srv_PostReg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for PostReg service

type PostRegService interface {
	CallPostReg(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type postRegService struct {
	c    client.Client
	name string
}

func NewPostRegService(name string, c client.Client) PostRegService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.PostReg"
	}
	return &postRegService{
		c:    c,
		name: name,
	}
}

func (c *postRegService) CallPostReg(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PostReg.CallPostReg", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PostReg service

type PostRegHandler interface {
	CallPostReg(context.Context, *Request, *Response) error
}

func RegisterPostRegHandler(s server.Server, hdlr PostRegHandler, opts ...server.HandlerOption) error {
	type postReg interface {
		CallPostReg(ctx context.Context, in *Request, out *Response) error
	}
	type PostReg struct {
		postReg
	}
	h := &postRegHandler{hdlr}
	return s.Handle(s.NewHandler(&PostReg{h}, opts...))
}

type postRegHandler struct {
	PostRegHandler
}

func (h *postRegHandler) CallPostReg(ctx context.Context, in *Request, out *Response) error {
	return h.PostRegHandler.CallPostReg(ctx, in, out)
}
