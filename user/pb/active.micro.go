// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user/pb/active.proto

/*
Package gs_service_user is a generated protocol buffer package.

It is generated from these files:
	user/pb/active.proto

It has these top-level messages:
	ActiveRequest
*/
package gs_service_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import gs_commons_dto "konekko.me/gosion/commons/dto"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = gs_commons_dto.Status{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Active service

type ActiveService interface {
	User(ctx context.Context, in *ActiveRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
}

type activeService struct {
	c    client.Client
	name string
}

func NewActiveService(name string, c client.Client) ActiveService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "gs.service.user"
	}
	return &activeService{
		c:    c,
		name: name,
	}
}

func (c *activeService) User(ctx context.Context, in *ActiveRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Active.User", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Active service

type ActiveHandler interface {
	User(context.Context, *ActiveRequest, *gs_commons_dto.Status) error
}

func RegisterActiveHandler(s server.Server, hdlr ActiveHandler, opts ...server.HandlerOption) error {
	type active interface {
		User(ctx context.Context, in *ActiveRequest, out *gs_commons_dto.Status) error
	}
	type Active struct {
		active
	}
	h := &activeHandler{hdlr}
	return s.Handle(s.NewHandler(&Active{h}, opts...))
}

type activeHandler struct {
	ActiveHandler
}

func (h *activeHandler) User(ctx context.Context, in *ActiveRequest, out *gs_commons_dto.Status) error {
	return h.ActiveHandler.User(ctx, in, out)
}