// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: pb/debugger.proto

/*
Package gs_service_workflow is a generated protocol buffer package.

It is generated from these files:
	pb/debugger.proto

It has these top-level messages:
	NextRequest
	NextResponse
	RunRequest
	RunResponse
*/
package gs_service_workflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Debugger service

type DebuggerService interface {
	Run(ctx context.Context, in *RunRequest, opts ...client.CallOption) (*RunResponse, error)
	Next(ctx context.Context, in *NextRequest, opts ...client.CallOption) (*NextResponse, error)
}

type debuggerService struct {
	c    client.Client
	name string
}

func NewDebuggerService(name string, c client.Client) DebuggerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "gs.service.workflow"
	}
	return &debuggerService{
		c:    c,
		name: name,
	}
}

func (c *debuggerService) Run(ctx context.Context, in *RunRequest, opts ...client.CallOption) (*RunResponse, error) {
	req := c.c.NewRequest(c.name, "Debugger.Run", in)
	out := new(RunResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debuggerService) Next(ctx context.Context, in *NextRequest, opts ...client.CallOption) (*NextResponse, error) {
	req := c.c.NewRequest(c.name, "Debugger.Next", in)
	out := new(NextResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Debugger service

type DebuggerHandler interface {
	Run(context.Context, *RunRequest, *RunResponse) error
	Next(context.Context, *NextRequest, *NextResponse) error
}

func RegisterDebuggerHandler(s server.Server, hdlr DebuggerHandler, opts ...server.HandlerOption) error {
	type debugger interface {
		Run(ctx context.Context, in *RunRequest, out *RunResponse) error
		Next(ctx context.Context, in *NextRequest, out *NextResponse) error
	}
	type Debugger struct {
		debugger
	}
	h := &debuggerHandler{hdlr}
	return s.Handle(s.NewHandler(&Debugger{h}, opts...))
}

type debuggerHandler struct {
	DebuggerHandler
}

func (h *debuggerHandler) Run(ctx context.Context, in *RunRequest, out *RunResponse) error {
	return h.DebuggerHandler.Run(ctx, in, out)
}

func (h *debuggerHandler) Next(ctx context.Context, in *NextRequest, out *NextResponse) error {
	return h.DebuggerHandler.Next(ctx, in, out)
}