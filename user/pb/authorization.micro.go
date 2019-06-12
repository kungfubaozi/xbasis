// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user/pb/authorization.proto

/*
Package gosionsvc_external_user is a generated protocol buffer package.

It is generated from these files:
	user/pb/authorization.proto

It has these top-level messages:
	SyncRequest
*/
package gosionsvc_external_user

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

// Client API for Authorization service

type AuthorizationService interface {
	Sync(ctx context.Context, in *SyncRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
}

type authorizationService struct {
	c    client.Client
	name string
}

func NewAuthorizationService(name string, c client.Client) AuthorizationService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "gosionsvc.external.user"
	}
	return &authorizationService{
		c:    c,
		name: name,
	}
}

func (c *authorizationService) Sync(ctx context.Context, in *SyncRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Authorization.Sync", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authorization service

type AuthorizationHandler interface {
	Sync(context.Context, *SyncRequest, *gs_commons_dto.Status) error
}

func RegisterAuthorizationHandler(s server.Server, hdlr AuthorizationHandler, opts ...server.HandlerOption) error {
	type authorization interface {
		Sync(ctx context.Context, in *SyncRequest, out *gs_commons_dto.Status) error
	}
	type Authorization struct {
		authorization
	}
	h := &authorizationHandler{hdlr}
	return s.Handle(s.NewHandler(&Authorization{h}, opts...))
}

type authorizationHandler struct {
	AuthorizationHandler
}

func (h *authorizationHandler) Sync(ctx context.Context, in *SyncRequest, out *gs_commons_dto.Status) error {
	return h.AuthorizationHandler.Sync(ctx, in, out)
}
