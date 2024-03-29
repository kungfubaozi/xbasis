// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: pb/instance.proto

/*
Package xbasissvc_external_workflow is a generated protocol buffer package.

It is generated from these files:
	pb/instance.proto

It has these top-level messages:
	SearchInstanceRequest
	SearchInstanceResponse
	SubmitRequest
	SubmitResponse
	ContinueRequest
	ContinueResponse
	RestartRequest
	RestartResponse
	StopRequest
	StopResponse
	GetAllInstancesRequest
	GetAllInstancesResponse
	GetMyLaunchInstancesRequest
	GetMyLaunchInstancesResponse
*/
package xbasissvc_external_workflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
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

// Client API for Instance service

type InstanceService interface {
	// 我发起的所有实例（流程）
	GetMyLaunchInstances(ctx context.Context, in *GetMyLaunchInstancesRequest, opts ...client.CallOption) (*GetMyLaunchInstancesResponse, error)
	// 所有实例
	GetAllInstances(ctx context.Context, in *GetAllInstancesRequest, opts ...client.CallOption) (*GetAllInstancesResponse, error)
	// 停止
	Stop(ctx context.Context, in *StopRequest, opts ...client.CallOption) (*StopResponse, error)
	// 重新开始
	Restart(ctx context.Context, in *RestartRequest, opts ...client.CallOption) (*RestartResponse, error)
	// 继续执行
	Continue(ctx context.Context, in *ContinueRequest, opts ...client.CallOption) (*ContinueResponse, error)
	Submit(ctx context.Context, in *SubmitRequest, opts ...client.CallOption) (*SubmitResponse, error)
	Search(ctx context.Context, in *SearchInstanceRequest, opts ...client.CallOption) (*SearchInstanceResponse, error)
}

type instanceService struct {
	c    client.Client
	name string
}

func NewInstanceService(name string, c client.Client) InstanceService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "xbasissvc.external.workflow"
	}
	return &instanceService{
		c:    c,
		name: name,
	}
}

func (c *instanceService) GetMyLaunchInstances(ctx context.Context, in *GetMyLaunchInstancesRequest, opts ...client.CallOption) (*GetMyLaunchInstancesResponse, error) {
	req := c.c.NewRequest(c.name, "Instance.GetMyLaunchInstances", in)
	out := new(GetMyLaunchInstancesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceService) GetAllInstances(ctx context.Context, in *GetAllInstancesRequest, opts ...client.CallOption) (*GetAllInstancesResponse, error) {
	req := c.c.NewRequest(c.name, "Instance.GetAllInstances", in)
	out := new(GetAllInstancesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceService) Stop(ctx context.Context, in *StopRequest, opts ...client.CallOption) (*StopResponse, error) {
	req := c.c.NewRequest(c.name, "Instance.Stop", in)
	out := new(StopResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceService) Restart(ctx context.Context, in *RestartRequest, opts ...client.CallOption) (*RestartResponse, error) {
	req := c.c.NewRequest(c.name, "Instance.Restart", in)
	out := new(RestartResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceService) Continue(ctx context.Context, in *ContinueRequest, opts ...client.CallOption) (*ContinueResponse, error) {
	req := c.c.NewRequest(c.name, "Instance.Continue", in)
	out := new(ContinueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceService) Submit(ctx context.Context, in *SubmitRequest, opts ...client.CallOption) (*SubmitResponse, error) {
	req := c.c.NewRequest(c.name, "Instance.Submit", in)
	out := new(SubmitResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceService) Search(ctx context.Context, in *SearchInstanceRequest, opts ...client.CallOption) (*SearchInstanceResponse, error) {
	req := c.c.NewRequest(c.name, "Instance.Search", in)
	out := new(SearchInstanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Instance service

type InstanceHandler interface {
	// 我发起的所有实例（流程）
	GetMyLaunchInstances(context.Context, *GetMyLaunchInstancesRequest, *GetMyLaunchInstancesResponse) error
	// 所有实例
	GetAllInstances(context.Context, *GetAllInstancesRequest, *GetAllInstancesResponse) error
	// 停止
	Stop(context.Context, *StopRequest, *StopResponse) error
	// 重新开始
	Restart(context.Context, *RestartRequest, *RestartResponse) error
	// 继续执行
	Continue(context.Context, *ContinueRequest, *ContinueResponse) error
	Submit(context.Context, *SubmitRequest, *SubmitResponse) error
	Search(context.Context, *SearchInstanceRequest, *SearchInstanceResponse) error
}

func RegisterInstanceHandler(s server.Server, hdlr InstanceHandler, opts ...server.HandlerOption) error {
	type instance interface {
		GetMyLaunchInstances(ctx context.Context, in *GetMyLaunchInstancesRequest, out *GetMyLaunchInstancesResponse) error
		GetAllInstances(ctx context.Context, in *GetAllInstancesRequest, out *GetAllInstancesResponse) error
		Stop(ctx context.Context, in *StopRequest, out *StopResponse) error
		Restart(ctx context.Context, in *RestartRequest, out *RestartResponse) error
		Continue(ctx context.Context, in *ContinueRequest, out *ContinueResponse) error
		Submit(ctx context.Context, in *SubmitRequest, out *SubmitResponse) error
		Search(ctx context.Context, in *SearchInstanceRequest, out *SearchInstanceResponse) error
	}
	type Instance struct {
		instance
	}
	h := &instanceHandler{hdlr}
	return s.Handle(s.NewHandler(&Instance{h}, opts...))
}

type instanceHandler struct {
	InstanceHandler
}

func (h *instanceHandler) GetMyLaunchInstances(ctx context.Context, in *GetMyLaunchInstancesRequest, out *GetMyLaunchInstancesResponse) error {
	return h.InstanceHandler.GetMyLaunchInstances(ctx, in, out)
}

func (h *instanceHandler) GetAllInstances(ctx context.Context, in *GetAllInstancesRequest, out *GetAllInstancesResponse) error {
	return h.InstanceHandler.GetAllInstances(ctx, in, out)
}

func (h *instanceHandler) Stop(ctx context.Context, in *StopRequest, out *StopResponse) error {
	return h.InstanceHandler.Stop(ctx, in, out)
}

func (h *instanceHandler) Restart(ctx context.Context, in *RestartRequest, out *RestartResponse) error {
	return h.InstanceHandler.Restart(ctx, in, out)
}

func (h *instanceHandler) Continue(ctx context.Context, in *ContinueRequest, out *ContinueResponse) error {
	return h.InstanceHandler.Continue(ctx, in, out)
}

func (h *instanceHandler) Submit(ctx context.Context, in *SubmitRequest, out *SubmitResponse) error {
	return h.InstanceHandler.Submit(ctx, in, out)
}

func (h *instanceHandler) Search(ctx context.Context, in *SearchInstanceRequest, out *SearchInstanceResponse) error {
	return h.InstanceHandler.Search(ctx, in, out)
}
