// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: safety/pb/blacklist.proto

/*
Package xbasissvc_external_safety is a generated protocol buffer package.

It is generated from these files:
	safety/pb/blacklist.proto

It has these top-level messages:
	BlacklistSearchRequest
	BlacklistSearchResponse
	BlacklistItem
	CheckRequest
	RemoveRequest
	AddRequest
*/
package xbasissvc_external_safety

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

// Client API for Blacklist service

type BlacklistService interface {
	Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error)
	Check(ctx context.Context, in *CheckRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error)
	Search(ctx context.Context, in *BlacklistSearchRequest, opts ...client.CallOption) (*BlacklistSearchResponse, error)
}

type blacklistService struct {
	c    client.Client
	name string
}

func NewBlacklistService(name string, c client.Client) BlacklistService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "xbasissvc.external.safety"
	}
	return &blacklistService{
		c:    c,
		name: name,
	}
}

func (c *blacklistService) Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Blacklist.Add", in)
	out := new(xbasis_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blacklistService) Remove(ctx context.Context, in *RemoveRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Blacklist.Remove", in)
	out := new(xbasis_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blacklistService) Check(ctx context.Context, in *CheckRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Blacklist.Check", in)
	out := new(xbasis_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blacklistService) Search(ctx context.Context, in *BlacklistSearchRequest, opts ...client.CallOption) (*BlacklistSearchResponse, error) {
	req := c.c.NewRequest(c.name, "Blacklist.Search", in)
	out := new(BlacklistSearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Blacklist service

type BlacklistHandler interface {
	Add(context.Context, *AddRequest, *xbasis_commons_dto.Status) error
	Remove(context.Context, *RemoveRequest, *xbasis_commons_dto.Status) error
	Check(context.Context, *CheckRequest, *xbasis_commons_dto.Status) error
	Search(context.Context, *BlacklistSearchRequest, *BlacklistSearchResponse) error
}

func RegisterBlacklistHandler(s server.Server, hdlr BlacklistHandler, opts ...server.HandlerOption) error {
	type blacklist interface {
		Add(ctx context.Context, in *AddRequest, out *xbasis_commons_dto.Status) error
		Remove(ctx context.Context, in *RemoveRequest, out *xbasis_commons_dto.Status) error
		Check(ctx context.Context, in *CheckRequest, out *xbasis_commons_dto.Status) error
		Search(ctx context.Context, in *BlacklistSearchRequest, out *BlacklistSearchResponse) error
	}
	type Blacklist struct {
		blacklist
	}
	h := &blacklistHandler{hdlr}
	return s.Handle(s.NewHandler(&Blacklist{h}, opts...))
}

type blacklistHandler struct {
	BlacklistHandler
}

func (h *blacklistHandler) Add(ctx context.Context, in *AddRequest, out *xbasis_commons_dto.Status) error {
	return h.BlacklistHandler.Add(ctx, in, out)
}

func (h *blacklistHandler) Remove(ctx context.Context, in *RemoveRequest, out *xbasis_commons_dto.Status) error {
	return h.BlacklistHandler.Remove(ctx, in, out)
}

func (h *blacklistHandler) Check(ctx context.Context, in *CheckRequest, out *xbasis_commons_dto.Status) error {
	return h.BlacklistHandler.Check(ctx, in, out)
}

func (h *blacklistHandler) Search(ctx context.Context, in *BlacklistSearchRequest, out *BlacklistSearchResponse) error {
	return h.BlacklistHandler.Search(ctx, in, out)
}
