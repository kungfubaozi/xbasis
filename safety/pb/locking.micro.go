// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: safety/pb/locking.proto

/*
Package gosionsvc_external_safety is a generated protocol buffer package.

It is generated from these files:
	safety/pb/locking.proto

It has these top-level messages:
	SearchRequest
	SearchResponse
	LockingItem
	LockRequest
	UnlockRequest
*/
package gosionsvc_external_safety

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

// Client API for Locking service

type LockingService interface {
	Lock(ctx context.Context, in *LockRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
	Unlock(ctx context.Context, in *UnlockRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
}

type lockingService struct {
	c    client.Client
	name string
}

func NewLockingService(name string, c client.Client) LockingService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "gosionsvc.external.safety"
	}
	return &lockingService{
		c:    c,
		name: name,
	}
}

func (c *lockingService) Lock(ctx context.Context, in *LockRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Locking.Lock", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lockingService) Unlock(ctx context.Context, in *UnlockRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Locking.Unlock", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lockingService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.name, "Locking.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Locking service

type LockingHandler interface {
	Lock(context.Context, *LockRequest, *gs_commons_dto.Status) error
	Unlock(context.Context, *UnlockRequest, *gs_commons_dto.Status) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
}

func RegisterLockingHandler(s server.Server, hdlr LockingHandler, opts ...server.HandlerOption) error {
	type locking interface {
		Lock(ctx context.Context, in *LockRequest, out *gs_commons_dto.Status) error
		Unlock(ctx context.Context, in *UnlockRequest, out *gs_commons_dto.Status) error
		Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error
	}
	type Locking struct {
		locking
	}
	h := &lockingHandler{hdlr}
	return s.Handle(s.NewHandler(&Locking{h}, opts...))
}

type lockingHandler struct {
	LockingHandler
}

func (h *lockingHandler) Lock(ctx context.Context, in *LockRequest, out *gs_commons_dto.Status) error {
	return h.LockingHandler.Lock(ctx, in, out)
}

func (h *lockingHandler) Unlock(ctx context.Context, in *UnlockRequest, out *gs_commons_dto.Status) error {
	return h.LockingHandler.Unlock(ctx, in, out)
}

func (h *lockingHandler) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.LockingHandler.Search(ctx, in, out)
}
