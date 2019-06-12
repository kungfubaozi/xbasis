// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user/pb/message.proto

/*
Package gosionsvc_external_user is a generated protocol buffer package.

It is generated from these files:
	user/pb/message.proto

It has these top-level messages:
	SendRequest
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

// Client API for Message service

type MessageService interface {
	SendMessage(ctx context.Context, in *SendRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
}

type messageService struct {
	c    client.Client
	name string
}

func NewMessageService(name string, c client.Client) MessageService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "gosionsvc.external.user"
	}
	return &messageService{
		c:    c,
		name: name,
	}
}

func (c *messageService) SendMessage(ctx context.Context, in *SendRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Message.SendMessage", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Message service

type MessageHandler interface {
	SendMessage(context.Context, *SendRequest, *gs_commons_dto.Status) error
}

func RegisterMessageHandler(s server.Server, hdlr MessageHandler, opts ...server.HandlerOption) error {
	type message interface {
		SendMessage(ctx context.Context, in *SendRequest, out *gs_commons_dto.Status) error
	}
	type Message struct {
		message
	}
	h := &messageHandler{hdlr}
	return s.Handle(s.NewHandler(&Message{h}, opts...))
}

type messageHandler struct {
	MessageHandler
}

func (h *messageHandler) SendMessage(ctx context.Context, in *SendRequest, out *gs_commons_dto.Status) error {
	return h.MessageHandler.SendMessage(ctx, in, out)
}
