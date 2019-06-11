// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: pb/form.proto

/*
Package gs_service_workflow is a generated protocol buffer package.

It is generated from these files:
	pb/form.proto

It has these top-level messages:
	CheckFiledValueRequest
	CheckFieldValueResponse
	GetAllTypeFieldsRequest
	GetAllTypeFieldsResponse
	CreatePlaceholderRequest
	CreatePlaceholderResponse
	DeletePlaceholderRequest
	DeletePlaceholderResponse
	UpdatePlaceholderRequest
	UpdatePlaceholderResponse
	AddFieldRequest
	AddFieldResponse
	RemoveFieldRequest
	RemoveFieldResponse
	UpdateFieldPropsRequest
	UpdateFieldPropsResponse
*/
package gs_service_workflow

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

// Client API for Form service

type FormService interface {
	CreatePlaceholder(ctx context.Context, in *CreatePlaceholderRequest, opts ...client.CallOption) (*CreatePlaceholderResponse, error)
	DeletePlaceholder(ctx context.Context, in *DeletePlaceholderRequest, opts ...client.CallOption) (*DeletePlaceholderResponse, error)
	UpdatePlaceholder(ctx context.Context, in *UpdatePlaceholderRequest, opts ...client.CallOption) (*UpdatePlaceholderResponse, error)
	AddField(ctx context.Context, in *AddFieldRequest, opts ...client.CallOption) (*AddFieldResponse, error)
	RemoveField(ctx context.Context, in *RemoveFieldRequest, opts ...client.CallOption) (*RemoveFieldResponse, error)
	UpdateFieldProps(ctx context.Context, in *UpdateFieldPropsRequest, opts ...client.CallOption) (*UpdateFieldPropsResponse, error)
	GetAllTypeFields(ctx context.Context, in *GetAllTypeFieldsRequest, opts ...client.CallOption) (*GetAllTypeFieldsResponse, error)
	// 检查form的filed value是否符合
	CheckFiledValue(ctx context.Context, in *CheckFiledValueRequest, opts ...client.CallOption) (*CheckFieldValueResponse, error)
}

type formService struct {
	c    client.Client
	name string
}

func NewFormService(name string, c client.Client) FormService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "gs.service.workflow"
	}
	return &formService{
		c:    c,
		name: name,
	}
}

func (c *formService) CreatePlaceholder(ctx context.Context, in *CreatePlaceholderRequest, opts ...client.CallOption) (*CreatePlaceholderResponse, error) {
	req := c.c.NewRequest(c.name, "Form.CreatePlaceholder", in)
	out := new(CreatePlaceholderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formService) DeletePlaceholder(ctx context.Context, in *DeletePlaceholderRequest, opts ...client.CallOption) (*DeletePlaceholderResponse, error) {
	req := c.c.NewRequest(c.name, "Form.DeletePlaceholder", in)
	out := new(DeletePlaceholderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formService) UpdatePlaceholder(ctx context.Context, in *UpdatePlaceholderRequest, opts ...client.CallOption) (*UpdatePlaceholderResponse, error) {
	req := c.c.NewRequest(c.name, "Form.UpdatePlaceholder", in)
	out := new(UpdatePlaceholderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formService) AddField(ctx context.Context, in *AddFieldRequest, opts ...client.CallOption) (*AddFieldResponse, error) {
	req := c.c.NewRequest(c.name, "Form.AddField", in)
	out := new(AddFieldResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formService) RemoveField(ctx context.Context, in *RemoveFieldRequest, opts ...client.CallOption) (*RemoveFieldResponse, error) {
	req := c.c.NewRequest(c.name, "Form.RemoveField", in)
	out := new(RemoveFieldResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formService) UpdateFieldProps(ctx context.Context, in *UpdateFieldPropsRequest, opts ...client.CallOption) (*UpdateFieldPropsResponse, error) {
	req := c.c.NewRequest(c.name, "Form.UpdateFieldProps", in)
	out := new(UpdateFieldPropsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formService) GetAllTypeFields(ctx context.Context, in *GetAllTypeFieldsRequest, opts ...client.CallOption) (*GetAllTypeFieldsResponse, error) {
	req := c.c.NewRequest(c.name, "Form.GetAllTypeFields", in)
	out := new(GetAllTypeFieldsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formService) CheckFiledValue(ctx context.Context, in *CheckFiledValueRequest, opts ...client.CallOption) (*CheckFieldValueResponse, error) {
	req := c.c.NewRequest(c.name, "Form.CheckFiledValue", in)
	out := new(CheckFieldValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Form service

type FormHandler interface {
	CreatePlaceholder(context.Context, *CreatePlaceholderRequest, *CreatePlaceholderResponse) error
	DeletePlaceholder(context.Context, *DeletePlaceholderRequest, *DeletePlaceholderResponse) error
	UpdatePlaceholder(context.Context, *UpdatePlaceholderRequest, *UpdatePlaceholderResponse) error
	AddField(context.Context, *AddFieldRequest, *AddFieldResponse) error
	RemoveField(context.Context, *RemoveFieldRequest, *RemoveFieldResponse) error
	UpdateFieldProps(context.Context, *UpdateFieldPropsRequest, *UpdateFieldPropsResponse) error
	GetAllTypeFields(context.Context, *GetAllTypeFieldsRequest, *GetAllTypeFieldsResponse) error
	// 检查form的filed value是否符合
	CheckFiledValue(context.Context, *CheckFiledValueRequest, *CheckFieldValueResponse) error
}

func RegisterFormHandler(s server.Server, hdlr FormHandler, opts ...server.HandlerOption) error {
	type form interface {
		CreatePlaceholder(ctx context.Context, in *CreatePlaceholderRequest, out *CreatePlaceholderResponse) error
		DeletePlaceholder(ctx context.Context, in *DeletePlaceholderRequest, out *DeletePlaceholderResponse) error
		UpdatePlaceholder(ctx context.Context, in *UpdatePlaceholderRequest, out *UpdatePlaceholderResponse) error
		AddField(ctx context.Context, in *AddFieldRequest, out *AddFieldResponse) error
		RemoveField(ctx context.Context, in *RemoveFieldRequest, out *RemoveFieldResponse) error
		UpdateFieldProps(ctx context.Context, in *UpdateFieldPropsRequest, out *UpdateFieldPropsResponse) error
		GetAllTypeFields(ctx context.Context, in *GetAllTypeFieldsRequest, out *GetAllTypeFieldsResponse) error
		CheckFiledValue(ctx context.Context, in *CheckFiledValueRequest, out *CheckFieldValueResponse) error
	}
	type Form struct {
		form
	}
	h := &formHandler{hdlr}
	return s.Handle(s.NewHandler(&Form{h}, opts...))
}

type formHandler struct {
	FormHandler
}

func (h *formHandler) CreatePlaceholder(ctx context.Context, in *CreatePlaceholderRequest, out *CreatePlaceholderResponse) error {
	return h.FormHandler.CreatePlaceholder(ctx, in, out)
}

func (h *formHandler) DeletePlaceholder(ctx context.Context, in *DeletePlaceholderRequest, out *DeletePlaceholderResponse) error {
	return h.FormHandler.DeletePlaceholder(ctx, in, out)
}

func (h *formHandler) UpdatePlaceholder(ctx context.Context, in *UpdatePlaceholderRequest, out *UpdatePlaceholderResponse) error {
	return h.FormHandler.UpdatePlaceholder(ctx, in, out)
}

func (h *formHandler) AddField(ctx context.Context, in *AddFieldRequest, out *AddFieldResponse) error {
	return h.FormHandler.AddField(ctx, in, out)
}

func (h *formHandler) RemoveField(ctx context.Context, in *RemoveFieldRequest, out *RemoveFieldResponse) error {
	return h.FormHandler.RemoveField(ctx, in, out)
}

func (h *formHandler) UpdateFieldProps(ctx context.Context, in *UpdateFieldPropsRequest, out *UpdateFieldPropsResponse) error {
	return h.FormHandler.UpdateFieldProps(ctx, in, out)
}

func (h *formHandler) GetAllTypeFields(ctx context.Context, in *GetAllTypeFieldsRequest, out *GetAllTypeFieldsResponse) error {
	return h.FormHandler.GetAllTypeFields(ctx, in, out)
}

func (h *formHandler) CheckFiledValue(ctx context.Context, in *CheckFiledValueRequest, out *CheckFieldValueResponse) error {
	return h.FormHandler.CheckFiledValue(ctx, in, out)
}
