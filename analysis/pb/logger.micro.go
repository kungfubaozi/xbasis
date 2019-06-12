// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: analysis/pb/logger.proto

/*
Package gosionsvc_external_analysis is a generated protocol buffer package.

It is generated from these files:
	analysis/pb/logger.proto

It has these top-level messages:
	UsageFunctionRequest
	UsageFunctionStatus
	UsageFunctionResponse
	TodayVisitRequest
	TodayVisitResponse
	VisitInfo
	PlatformVisitInfo
	GetDataResponse
	GetDataRequest
	XAxisRequest
	XAxisFactor
	YAxisRequest
	YAxisFactor
*/
package gosionsvc_external_analysis

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "konekko.me/gosion/commons/dto"

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

// Client API for Logger service

type LoggerService interface {
	GetAxisData(ctx context.Context, in *GetDataRequest, opts ...client.CallOption) (*GetDataResponse, error)
	// 今日使用情况
	TodayVisit(ctx context.Context, in *TodayVisitRequest, opts ...client.CallOption) (*TodayVisitResponse, error)
	// 使用功能统计
	UsageFunction(ctx context.Context, in *UsageFunctionRequest, opts ...client.CallOption) (*UsageFunctionResponse, error)
}

type loggerService struct {
	c    client.Client
	name string
}

func NewLoggerService(name string, c client.Client) LoggerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "gosionsvc.external.analysis"
	}
	return &loggerService{
		c:    c,
		name: name,
	}
}

func (c *loggerService) GetAxisData(ctx context.Context, in *GetDataRequest, opts ...client.CallOption) (*GetDataResponse, error) {
	req := c.c.NewRequest(c.name, "Logger.GetAxisData", in)
	out := new(GetDataResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggerService) TodayVisit(ctx context.Context, in *TodayVisitRequest, opts ...client.CallOption) (*TodayVisitResponse, error) {
	req := c.c.NewRequest(c.name, "Logger.TodayVisit", in)
	out := new(TodayVisitResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggerService) UsageFunction(ctx context.Context, in *UsageFunctionRequest, opts ...client.CallOption) (*UsageFunctionResponse, error) {
	req := c.c.NewRequest(c.name, "Logger.UsageFunction", in)
	out := new(UsageFunctionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Logger service

type LoggerHandler interface {
	GetAxisData(context.Context, *GetDataRequest, *GetDataResponse) error
	// 今日使用情况
	TodayVisit(context.Context, *TodayVisitRequest, *TodayVisitResponse) error
	// 使用功能统计
	UsageFunction(context.Context, *UsageFunctionRequest, *UsageFunctionResponse) error
}

func RegisterLoggerHandler(s server.Server, hdlr LoggerHandler, opts ...server.HandlerOption) error {
	type logger interface {
		GetAxisData(ctx context.Context, in *GetDataRequest, out *GetDataResponse) error
		TodayVisit(ctx context.Context, in *TodayVisitRequest, out *TodayVisitResponse) error
		UsageFunction(ctx context.Context, in *UsageFunctionRequest, out *UsageFunctionResponse) error
	}
	type Logger struct {
		logger
	}
	h := &loggerHandler{hdlr}
	return s.Handle(s.NewHandler(&Logger{h}, opts...))
}

type loggerHandler struct {
	LoggerHandler
}

func (h *loggerHandler) GetAxisData(ctx context.Context, in *GetDataRequest, out *GetDataResponse) error {
	return h.LoggerHandler.GetAxisData(ctx, in, out)
}

func (h *loggerHandler) TodayVisit(ctx context.Context, in *TodayVisitRequest, out *TodayVisitResponse) error {
	return h.LoggerHandler.TodayVisit(ctx, in, out)
}

func (h *loggerHandler) UsageFunction(ctx context.Context, in *UsageFunctionRequest, out *UsageFunctionResponse) error {
	return h.LoggerHandler.UsageFunction(ctx, in, out)
}
