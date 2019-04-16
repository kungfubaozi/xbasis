// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: application/pb/ext/sync.proto

/*
Package gs_ext_service_application is a generated protocol buffer package.

It is generated from these files:
	application/pb/ext/sync.proto

It has these top-level messages:
	CheckRequest
	UserInfo
*/
package gs_ext_service_application

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

// Client API for Usersync service

type UsersyncService interface {
	Transport(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*gs_commons_dto.Status, error)
	Check(ctx context.Context, in *CheckRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
	Update(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*gs_commons_dto.Status, error)
}

type usersyncService struct {
	c    client.Client
	name string
}

func NewUsersyncService(name string, c client.Client) UsersyncService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "gs.ext.service.application"
	}
	return &usersyncService{
		c:    c,
		name: name,
	}
}

func (c *usersyncService) Transport(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Usersync.Transport", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersyncService) Check(ctx context.Context, in *CheckRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Usersync.Check", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersyncService) Update(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Usersync.Update", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Usersync service

type UsersyncHandler interface {
	Transport(context.Context, *UserInfo, *gs_commons_dto.Status) error
	Check(context.Context, *CheckRequest, *gs_commons_dto.Status) error
	Update(context.Context, *UserInfo, *gs_commons_dto.Status) error
}

func RegisterUsersyncHandler(s server.Server, hdlr UsersyncHandler, opts ...server.HandlerOption) error {
	type usersync interface {
		Transport(ctx context.Context, in *UserInfo, out *gs_commons_dto.Status) error
		Check(ctx context.Context, in *CheckRequest, out *gs_commons_dto.Status) error
		Update(ctx context.Context, in *UserInfo, out *gs_commons_dto.Status) error
	}
	type Usersync struct {
		usersync
	}
	h := &usersyncHandler{hdlr}
	return s.Handle(s.NewHandler(&Usersync{h}, opts...))
}

type usersyncHandler struct {
	UsersyncHandler
}

func (h *usersyncHandler) Transport(ctx context.Context, in *UserInfo, out *gs_commons_dto.Status) error {
	return h.UsersyncHandler.Transport(ctx, in, out)
}

func (h *usersyncHandler) Check(ctx context.Context, in *CheckRequest, out *gs_commons_dto.Status) error {
	return h.UsersyncHandler.Check(ctx, in, out)
}

func (h *usersyncHandler) Update(ctx context.Context, in *UserInfo, out *gs_commons_dto.Status) error {
	return h.UsersyncHandler.Update(ctx, in, out)
}
