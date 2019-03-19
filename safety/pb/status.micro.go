// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: safety/pb/status.proto

/*
Package gs_service_safety is a generated protocol buffer package.

It is generated from these files:
	safety/pb/status.proto

It has these top-level messages:
	TestRequest
*/
package gs_service_safety

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

// Client API for Status service

type StatusService interface {
	Test(ctx context.Context, in *TestRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
}

type statusService struct {
	c    client.Client
	name string
}

func NewStatusService(name string, c client.Client) StatusService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "gs.service.safety"
	}
	return &statusService{
		c:    c,
		name: name,
	}
}

func (c *statusService) Test(ctx context.Context, in *TestRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Status.test", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Status service

type StatusHandler interface {
	Test(context.Context, *TestRequest, *gs_commons_dto.Status) error
}

func RegisterStatusHandler(s server.Server, hdlr StatusHandler, opts ...server.HandlerOption) error {
	type status interface {
		Test(ctx context.Context, in *TestRequest, out *gs_commons_dto.Status) error
	}
	type Status struct {
		status
	}
	h := &statusHandler{hdlr}
	return s.Handle(s.NewHandler(&Status{h}, opts...))
}

type statusHandler struct {
	StatusHandler
}

func (h *statusHandler) Test(ctx context.Context, in *TestRequest, out *gs_commons_dto.Status) error {
	return h.StatusHandler.Test(ctx, in, out)
}
