// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: userApi.proto

package go_micro_api_userApi

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserApi service

func NewUserApiEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserApi service

type UserApiService interface {
	Register(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Login(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	SendMessage(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type userApiService struct {
	c    client.Client
	name string
}

func NewUserApiService(name string, c client.Client) UserApiService {
	return &userApiService{
		c:    c,
		name: name,
	}
}

func (c *userApiService) Register(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserApi.Register", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiService) Login(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserApi.Login", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiService) SendMessage(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserApi.SendMessage", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserApi service

type UserApiHandler interface {
	Register(context.Context, *Request, *Response) error
	Login(context.Context, *Request, *Response) error
	SendMessage(context.Context, *Request, *Response) error
}

func RegisterUserApiHandler(s server.Server, hdlr UserApiHandler, opts ...server.HandlerOption) error {
	type userApi interface {
		Register(ctx context.Context, in *Request, out *Response) error
		Login(ctx context.Context, in *Request, out *Response) error
		SendMessage(ctx context.Context, in *Request, out *Response) error
	}
	type UserApi struct {
		userApi
	}
	h := &userApiHandler{hdlr}
	return s.Handle(s.NewHandler(&UserApi{h}, opts...))
}

type userApiHandler struct {
	UserApiHandler
}

func (h *userApiHandler) Register(ctx context.Context, in *Request, out *Response) error {
	return h.UserApiHandler.Register(ctx, in, out)
}

func (h *userApiHandler) Login(ctx context.Context, in *Request, out *Response) error {
	return h.UserApiHandler.Login(ctx, in, out)
}

func (h *userApiHandler) SendMessage(ctx context.Context, in *Request, out *Response) error {
	return h.UserApiHandler.SendMessage(ctx, in, out)
}
