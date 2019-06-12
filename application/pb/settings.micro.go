// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: application/pb/settings.proto

/*
Package gosionsvc_external_application is a generated protocol buffer package.

It is generated from these files:
	application/pb/settings.proto

It has these top-level messages:
	UpdateRequest
	EnabledRequest
	GetRequest
*/
package gosionsvc_external_application

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

// Client API for Settings service

type SettingsService interface {
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
	EnabledClient(ctx context.Context, in *EnabledRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
}

type settingsService struct {
	c    client.Client
	name string
}

func NewSettingsService(name string, c client.Client) SettingsService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "gosionsvc.external.application"
	}
	return &settingsService{
		c:    c,
		name: name,
	}
}

func (c *settingsService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Settings.Update", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsService) EnabledClient(ctx context.Context, in *EnabledRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Settings.EnabledClient", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Settings service

type SettingsHandler interface {
	Update(context.Context, *UpdateRequest, *gs_commons_dto.Status) error
	EnabledClient(context.Context, *EnabledRequest, *gs_commons_dto.Status) error
}

func RegisterSettingsHandler(s server.Server, hdlr SettingsHandler, opts ...server.HandlerOption) error {
	type settings interface {
		Update(ctx context.Context, in *UpdateRequest, out *gs_commons_dto.Status) error
		EnabledClient(ctx context.Context, in *EnabledRequest, out *gs_commons_dto.Status) error
	}
	type Settings struct {
		settings
	}
	h := &settingsHandler{hdlr}
	return s.Handle(s.NewHandler(&Settings{h}, opts...))
}

type settingsHandler struct {
	SettingsHandler
}

func (h *settingsHandler) Update(ctx context.Context, in *UpdateRequest, out *gs_commons_dto.Status) error {
	return h.SettingsHandler.Update(ctx, in, out)
}

func (h *settingsHandler) EnabledClient(ctx context.Context, in *EnabledRequest, out *gs_commons_dto.Status) error {
	return h.SettingsHandler.EnabledClient(ctx, in, out)
}
