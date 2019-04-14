// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: application/pb/application.proto

/*
Package gs_service_application is a generated protocol buffer package.

It is generated from these files:
	application/pb/application.proto

It has these top-level messages:
	EnabledRequest
	FindRequest
	ListResponse
	SimpleApplicationResponse
	CreateRequest
	RemoveRequest
	ChangeNameRequest
	AppInfo
	AppClientInfo
*/
package gs_service_application

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

// Client API for Application service

type ApplicationService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
	ChangeName(ctx context.Context, in *ChangeNameRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
	FindByAppId(ctx context.Context, in *FindRequest, opts ...client.CallOption) (*SimpleApplicationResponse, error)
	FindByClientId(ctx context.Context, in *FindRequest, opts ...client.CallOption) (*SimpleApplicationResponse, error)
	List(ctx context.Context, in *FindRequest, opts ...client.CallOption) (*ListResponse, error)
	Enabled(ctx context.Context, in *EnabledRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
}

type applicationService struct {
	c    client.Client
	name string
}

func NewApplicationService(name string, c client.Client) ApplicationService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "gs.service.application"
	}
	return &applicationService{
		c:    c,
		name: name,
	}
}

func (c *applicationService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Application.Create", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationService) Remove(ctx context.Context, in *RemoveRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Application.Remove", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationService) ChangeName(ctx context.Context, in *ChangeNameRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Application.ChangeName", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationService) FindByAppId(ctx context.Context, in *FindRequest, opts ...client.CallOption) (*SimpleApplicationResponse, error) {
	req := c.c.NewRequest(c.name, "Application.FindByAppId", in)
	out := new(SimpleApplicationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationService) FindByClientId(ctx context.Context, in *FindRequest, opts ...client.CallOption) (*SimpleApplicationResponse, error) {
	req := c.c.NewRequest(c.name, "Application.FindByClientId", in)
	out := new(SimpleApplicationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationService) List(ctx context.Context, in *FindRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Application.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationService) Enabled(ctx context.Context, in *EnabledRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Application.Enabled", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Application service

type ApplicationHandler interface {
	Create(context.Context, *CreateRequest, *gs_commons_dto.Status) error
	Remove(context.Context, *RemoveRequest, *gs_commons_dto.Status) error
	ChangeName(context.Context, *ChangeNameRequest, *gs_commons_dto.Status) error
	FindByAppId(context.Context, *FindRequest, *SimpleApplicationResponse) error
	FindByClientId(context.Context, *FindRequest, *SimpleApplicationResponse) error
	List(context.Context, *FindRequest, *ListResponse) error
	Enabled(context.Context, *EnabledRequest, *gs_commons_dto.Status) error
}

func RegisterApplicationHandler(s server.Server, hdlr ApplicationHandler, opts ...server.HandlerOption) error {
	type application interface {
		Create(ctx context.Context, in *CreateRequest, out *gs_commons_dto.Status) error
		Remove(ctx context.Context, in *RemoveRequest, out *gs_commons_dto.Status) error
		ChangeName(ctx context.Context, in *ChangeNameRequest, out *gs_commons_dto.Status) error
		FindByAppId(ctx context.Context, in *FindRequest, out *SimpleApplicationResponse) error
		FindByClientId(ctx context.Context, in *FindRequest, out *SimpleApplicationResponse) error
		List(ctx context.Context, in *FindRequest, out *ListResponse) error
		Enabled(ctx context.Context, in *EnabledRequest, out *gs_commons_dto.Status) error
	}
	type Application struct {
		application
	}
	h := &applicationHandler{hdlr}
	return s.Handle(s.NewHandler(&Application{h}, opts...))
}

type applicationHandler struct {
	ApplicationHandler
}

func (h *applicationHandler) Create(ctx context.Context, in *CreateRequest, out *gs_commons_dto.Status) error {
	return h.ApplicationHandler.Create(ctx, in, out)
}

func (h *applicationHandler) Remove(ctx context.Context, in *RemoveRequest, out *gs_commons_dto.Status) error {
	return h.ApplicationHandler.Remove(ctx, in, out)
}

func (h *applicationHandler) ChangeName(ctx context.Context, in *ChangeNameRequest, out *gs_commons_dto.Status) error {
	return h.ApplicationHandler.ChangeName(ctx, in, out)
}

func (h *applicationHandler) FindByAppId(ctx context.Context, in *FindRequest, out *SimpleApplicationResponse) error {
	return h.ApplicationHandler.FindByAppId(ctx, in, out)
}

func (h *applicationHandler) FindByClientId(ctx context.Context, in *FindRequest, out *SimpleApplicationResponse) error {
	return h.ApplicationHandler.FindByClientId(ctx, in, out)
}

func (h *applicationHandler) List(ctx context.Context, in *FindRequest, out *ListResponse) error {
	return h.ApplicationHandler.List(ctx, in, out)
}

func (h *applicationHandler) Enabled(ctx context.Context, in *EnabledRequest, out *gs_commons_dto.Status) error {
	return h.ApplicationHandler.Enabled(ctx, in, out)
}
