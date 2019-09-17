// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/GetImageCd/GetImageCd.proto

package go_micro_srv_GetImageCd

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

// Client API for GetImageCd service

type GetImageCdService interface {
	CallGetImageCd(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type getImageCdService struct {
	c    client.Client
	name string
}

func NewGetImageCdService(name string, c client.Client) GetImageCdService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.GetImageCd"
	}
	return &getImageCdService{
		c:    c,
		name: name,
	}
}

func (c *getImageCdService) CallGetImageCd(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "GetImageCd.CallGetImageCd", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GetImageCd service

type GetImageCdHandler interface {
	CallGetImageCd(context.Context, *Request, *Response) error
}

func RegisterGetImageCdHandler(s server.Server, hdlr GetImageCdHandler, opts ...server.HandlerOption) error {
	type getImageCd interface {
		CallGetImageCd(ctx context.Context, in *Request, out *Response) error
	}
	type GetImageCd struct {
		getImageCd
	}
	h := &getImageCdHandler{hdlr}
	return s.Handle(s.NewHandler(&GetImageCd{h}, opts...))
}

type getImageCdHandler struct {
	GetImageCdHandler
}

func (h *getImageCdHandler) CallGetImageCd(ctx context.Context, in *Request, out *Response) error {
	return h.GetImageCdHandler.CallGetImageCd(ctx, in, out)
}
