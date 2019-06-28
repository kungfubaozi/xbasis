// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user/pb/oauth.proto

/*
Package xbasissvc_external_user is a generated protocol buffer package.

It is generated from these files:
	user/pb/oauth.proto

It has these top-level messages:
	BindOAuthRequest
	UnbindOAuthRequest
	OAuthLoginRequest
	OAuthLoginResponse
*/
package xbasissvc_external_user

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

// Client API for OAuth service

type OAuthService interface {
	Login(ctx context.Context, in *OAuthLoginRequest, opts ...client.CallOption) (*OAuthLoginResponse, error)
	BindOAuth(ctx context.Context, in *BindOAuthRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error)
	UnbindOAuth(ctx context.Context, in *UnbindOAuthRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error)
}

type oAuthService struct {
	c    client.Client
	name string
}

func NewOAuthService(name string, c client.Client) OAuthService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "xbasissvc.external.user"
	}
	return &oAuthService{
		c:    c,
		name: name,
	}
}

func (c *oAuthService) Login(ctx context.Context, in *OAuthLoginRequest, opts ...client.CallOption) (*OAuthLoginResponse, error) {
	req := c.c.NewRequest(c.name, "OAuth.Login", in)
	out := new(OAuthLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthService) BindOAuth(ctx context.Context, in *BindOAuthRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "OAuth.BindOAuth", in)
	out := new(xbasis_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthService) UnbindOAuth(ctx context.Context, in *UnbindOAuthRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "OAuth.UnbindOAuth", in)
	out := new(xbasis_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OAuth service

type OAuthHandler interface {
	Login(context.Context, *OAuthLoginRequest, *OAuthLoginResponse) error
	BindOAuth(context.Context, *BindOAuthRequest, *xbasis_commons_dto.Status) error
	UnbindOAuth(context.Context, *UnbindOAuthRequest, *xbasis_commons_dto.Status) error
}

func RegisterOAuthHandler(s server.Server, hdlr OAuthHandler, opts ...server.HandlerOption) error {
	type oAuth interface {
		Login(ctx context.Context, in *OAuthLoginRequest, out *OAuthLoginResponse) error
		BindOAuth(ctx context.Context, in *BindOAuthRequest, out *xbasis_commons_dto.Status) error
		UnbindOAuth(ctx context.Context, in *UnbindOAuthRequest, out *xbasis_commons_dto.Status) error
	}
	type OAuth struct {
		oAuth
	}
	h := &oAuthHandler{hdlr}
	return s.Handle(s.NewHandler(&OAuth{h}, opts...))
}

type oAuthHandler struct {
	OAuthHandler
}

func (h *oAuthHandler) Login(ctx context.Context, in *OAuthLoginRequest, out *OAuthLoginResponse) error {
	return h.OAuthHandler.Login(ctx, in, out)
}

func (h *oAuthHandler) BindOAuth(ctx context.Context, in *BindOAuthRequest, out *xbasis_commons_dto.Status) error {
	return h.OAuthHandler.BindOAuth(ctx, in, out)
}

func (h *oAuthHandler) UnbindOAuth(ctx context.Context, in *UnbindOAuthRequest, out *xbasis_commons_dto.Status) error {
	return h.OAuthHandler.UnbindOAuth(ctx, in, out)
}
