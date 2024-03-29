// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: permission/pb/inner/verification.proto

/*
Package xbasissvc_internal_permission is a generated protocol buffer package.

It is generated from these files:
	permission/pb/inner/verification.proto

It has these top-level messages:
	HasPermissionRequest
	HasPermissionResponse
	TokenInfo
*/
package xbasissvc_internal_permission

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

// Client API for Verification service

type VerificationService interface {
	// 是否有权限
	Check(ctx context.Context, in *HasPermissionRequest, opts ...client.CallOption) (*HasPermissionResponse, error)
}

type verificationService struct {
	c    client.Client
	name string
}

func NewVerificationService(name string, c client.Client) VerificationService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "xbasissvc.internal.permission"
	}
	return &verificationService{
		c:    c,
		name: name,
	}
}

func (c *verificationService) Check(ctx context.Context, in *HasPermissionRequest, opts ...client.CallOption) (*HasPermissionResponse, error) {
	req := c.c.NewRequest(c.name, "Verification.Check", in)
	out := new(HasPermissionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Verification service

type VerificationHandler interface {
	// 是否有权限
	Check(context.Context, *HasPermissionRequest, *HasPermissionResponse) error
}

func RegisterVerificationHandler(s server.Server, hdlr VerificationHandler, opts ...server.HandlerOption) error {
	type verification interface {
		Check(ctx context.Context, in *HasPermissionRequest, out *HasPermissionResponse) error
	}
	type Verification struct {
		verification
	}
	h := &verificationHandler{hdlr}
	return s.Handle(s.NewHandler(&Verification{h}, opts...))
}

type verificationHandler struct {
	VerificationHandler
}

func (h *verificationHandler) Check(ctx context.Context, in *HasPermissionRequest, out *HasPermissionResponse) error {
	return h.VerificationHandler.Check(ctx, in, out)
}
