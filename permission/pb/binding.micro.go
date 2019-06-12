// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: permission/pb/binding.proto

/*
Package gosionsvc_external_permission is a generated protocol buffer package.

It is generated from these files:
	permission/pb/binding.proto

It has these top-level messages:
	BindingRoleRequest
*/
package gosionsvc_external_permission

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

// Client API for Binding service

type BindingService interface {
	UserRole(ctx context.Context, in *BindingRoleRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
	FunctionRole(ctx context.Context, in *BindingRoleRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
	UnbindUserRole(ctx context.Context, in *BindingRoleRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
	UnbindFunctionRole(ctx context.Context, in *BindingRoleRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
}

type bindingService struct {
	c    client.Client
	name string
}

func NewBindingService(name string, c client.Client) BindingService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "gosionsvc.external.permission"
	}
	return &bindingService{
		c:    c,
		name: name,
	}
}

func (c *bindingService) UserRole(ctx context.Context, in *BindingRoleRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Binding.UserRole", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bindingService) FunctionRole(ctx context.Context, in *BindingRoleRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Binding.FunctionRole", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bindingService) UnbindUserRole(ctx context.Context, in *BindingRoleRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Binding.UnbindUserRole", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bindingService) UnbindFunctionRole(ctx context.Context, in *BindingRoleRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Binding.UnbindFunctionRole", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Binding service

type BindingHandler interface {
	UserRole(context.Context, *BindingRoleRequest, *gs_commons_dto.Status) error
	FunctionRole(context.Context, *BindingRoleRequest, *gs_commons_dto.Status) error
	UnbindUserRole(context.Context, *BindingRoleRequest, *gs_commons_dto.Status) error
	UnbindFunctionRole(context.Context, *BindingRoleRequest, *gs_commons_dto.Status) error
}

func RegisterBindingHandler(s server.Server, hdlr BindingHandler, opts ...server.HandlerOption) error {
	type binding interface {
		UserRole(ctx context.Context, in *BindingRoleRequest, out *gs_commons_dto.Status) error
		FunctionRole(ctx context.Context, in *BindingRoleRequest, out *gs_commons_dto.Status) error
		UnbindUserRole(ctx context.Context, in *BindingRoleRequest, out *gs_commons_dto.Status) error
		UnbindFunctionRole(ctx context.Context, in *BindingRoleRequest, out *gs_commons_dto.Status) error
	}
	type Binding struct {
		binding
	}
	h := &bindingHandler{hdlr}
	return s.Handle(s.NewHandler(&Binding{h}, opts...))
}

type bindingHandler struct {
	BindingHandler
}

func (h *bindingHandler) UserRole(ctx context.Context, in *BindingRoleRequest, out *gs_commons_dto.Status) error {
	return h.BindingHandler.UserRole(ctx, in, out)
}

func (h *bindingHandler) FunctionRole(ctx context.Context, in *BindingRoleRequest, out *gs_commons_dto.Status) error {
	return h.BindingHandler.FunctionRole(ctx, in, out)
}

func (h *bindingHandler) UnbindUserRole(ctx context.Context, in *BindingRoleRequest, out *gs_commons_dto.Status) error {
	return h.BindingHandler.UnbindUserRole(ctx, in, out)
}

func (h *bindingHandler) UnbindFunctionRole(ctx context.Context, in *BindingRoleRequest, out *gs_commons_dto.Status) error {
	return h.BindingHandler.UnbindFunctionRole(ctx, in, out)
}
