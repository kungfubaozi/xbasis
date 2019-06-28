// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: safety/pb/inner/security.proto

/*
Package xbasissvc_internal_safety is a generated protocol buffer package.

It is generated from these files:
	safety/pb/inner/security.proto

It has these top-level messages:
	GetRequest
	GetResponse
*/
package xbasissvc_internal_safety

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "konekko.me/xbasis/commons/dto"

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

// Client API for Security service

type SecurityService interface {
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
}

type securityService struct {
	c    client.Client
	name string
}

func NewSecurityService(name string, c client.Client) SecurityService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "xbasissvc.internal.safety"
	}
	return &securityService{
		c:    c,
		name: name,
	}
}

func (c *securityService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "Security.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Security service

type SecurityHandler interface {
	Get(context.Context, *GetRequest, *GetResponse) error
}

func RegisterSecurityHandler(s server.Server, hdlr SecurityHandler, opts ...server.HandlerOption) error {
	type security interface {
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
	}
	type Security struct {
		security
	}
	h := &securityHandler{hdlr}
	return s.Handle(s.NewHandler(&Security{h}, opts...))
}

type securityHandler struct {
	SecurityHandler
}

func (h *securityHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.SecurityHandler.Get(ctx, in, out)
}
