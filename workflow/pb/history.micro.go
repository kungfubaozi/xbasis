// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: pb/history.proto

/*
Package gs_service_workflow is a generated protocol buffer package.

It is generated from these files:
	pb/history.proto

It has these top-level messages:
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

// Client API for History service

type HistoryService interface {
}

type historyService struct {
	c    client.Client
	name string
}

func NewHistoryService(name string, c client.Client) HistoryService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "gs.service.workflow"
	}
	return &historyService{
		c:    c,
		name: name,
	}
}

// Server API for History service

type HistoryHandler interface {
}

func RegisterHistoryHandler(s server.Server, hdlr HistoryHandler, opts ...server.HandlerOption) error {
	type history interface {
	}
	type History struct {
		history
	}
	h := &historyHandler{hdlr}
	return s.Handle(s.NewHandler(&History{h}, opts...))
}

type historyHandler struct {
	HistoryHandler
}
