// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: authentication/pb/inner/token.proto

/*
Package xbasissvc_internal_authentication is a generated protocol buffer package.

It is generated from these files:
	authentication/pb/inner/token.proto

It has these top-level messages:
	GenerateRequest
	GenerateResponse
*/
package xbasissvc_internal_authentication

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

// Client API for Token service

type TokenService interface {
	Generate(ctx context.Context, in *GenerateRequest, opts ...client.CallOption) (*GenerateResponse, error)
}

type tokenService struct {
	c    client.Client
	name string
}

func NewTokenService(name string, c client.Client) TokenService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "xbasissvc.internal.authentication"
	}
	return &tokenService{
		c:    c,
		name: name,
	}
}

func (c *tokenService) Generate(ctx context.Context, in *GenerateRequest, opts ...client.CallOption) (*GenerateResponse, error) {
	req := c.c.NewRequest(c.name, "Token.Generate", in)
	out := new(GenerateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Token service

type TokenHandler interface {
	Generate(context.Context, *GenerateRequest, *GenerateResponse) error
}

func RegisterTokenHandler(s server.Server, hdlr TokenHandler, opts ...server.HandlerOption) error {
	type token interface {
		Generate(ctx context.Context, in *GenerateRequest, out *GenerateResponse) error
	}
	type Token struct {
		token
	}
	h := &tokenHandler{hdlr}
	return s.Handle(s.NewHandler(&Token{h}, opts...))
}

type tokenHandler struct {
	TokenHandler
}

func (h *tokenHandler) Generate(ctx context.Context, in *GenerateRequest, out *GenerateResponse) error {
	return h.TokenHandler.Generate(ctx, in, out)
}
