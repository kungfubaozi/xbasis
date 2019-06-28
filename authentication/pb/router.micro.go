// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: authentication/pb/router.proto

/*
Package gosionsvc_external_authentication is a generated protocol buffer package.

It is generated from these files:
	authentication/pb/router.proto

It has these top-level messages:
	AuthorizeRequest
	LogoutRequest
	RefreshRequest
	RefreshResponse
	PushRequest
	PushResponse
*/
package gosionsvc_external_authentication

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import xbasis_commons_dto "konekko.me/xbasis/commons/dto"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = xbasis_commons_dto.Status{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Router service

type RouterService interface {
	Push(ctx context.Context, in *PushRequest, opts ...client.CallOption) (*PushResponse, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...client.CallOption) (*RefreshResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error)
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error)
}

type routerService struct {
	c    client.Client
	name string
}

func NewRouterService(name string, c client.Client) RouterService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "gosionsvc.external.authentication"
	}
	return &routerService{
		c:    c,
		name: name,
	}
}

func (c *routerService) Push(ctx context.Context, in *PushRequest, opts ...client.CallOption) (*PushResponse, error) {
	req := c.c.NewRequest(c.name, "Router.Push", in)
	out := new(PushResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerService) Refresh(ctx context.Context, in *RefreshRequest, opts ...client.CallOption) (*RefreshResponse, error) {
	req := c.c.NewRequest(c.name, "Router.Refresh", in)
	out := new(RefreshResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerService) Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Router.Logout", in)
	out := new(xbasis_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerService) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Router.Authorize", in)
	out := new(xbasis_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Router service

type RouterHandler interface {
	Push(context.Context, *PushRequest, *PushResponse) error
	Refresh(context.Context, *RefreshRequest, *RefreshResponse) error
	Logout(context.Context, *LogoutRequest, *xbasis_commons_dto.Status) error
	Authorize(context.Context, *AuthorizeRequest, *xbasis_commons_dto.Status) error
}

func RegisterRouterHandler(s server.Server, hdlr RouterHandler, opts ...server.HandlerOption) error {
	type router interface {
		Push(ctx context.Context, in *PushRequest, out *PushResponse) error
		Refresh(ctx context.Context, in *RefreshRequest, out *RefreshResponse) error
		Logout(ctx context.Context, in *LogoutRequest, out *xbasis_commons_dto.Status) error
		Authorize(ctx context.Context, in *AuthorizeRequest, out *xbasis_commons_dto.Status) error
	}
	type Router struct {
		router
	}
	h := &routerHandler{hdlr}
	return s.Handle(s.NewHandler(&Router{h}, opts...))
}

type routerHandler struct {
	RouterHandler
}

func (h *routerHandler) Push(ctx context.Context, in *PushRequest, out *PushResponse) error {
	return h.RouterHandler.Push(ctx, in, out)
}

func (h *routerHandler) Refresh(ctx context.Context, in *RefreshRequest, out *RefreshResponse) error {
	return h.RouterHandler.Refresh(ctx, in, out)
}

func (h *routerHandler) Logout(ctx context.Context, in *LogoutRequest, out *xbasis_commons_dto.Status) error {
	return h.RouterHandler.Logout(ctx, in, out)
}

func (h *routerHandler) Authorize(ctx context.Context, in *AuthorizeRequest, out *xbasis_commons_dto.Status) error {
	return h.RouterHandler.Authorize(ctx, in, out)
}
