// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: permission/pb/group.proto

/*
Package xbasissvc_external_permission is a generated protocol buffer package.

It is generated from these files:
	permission/pb/group.proto

It has these top-level messages:
	SearchAppUserRequest
	SearchAppUserResponse
	MoveRequest
	GetGroupContentSizeRequest
	GetGroupContentSizeResponse
	GetGroupItemsRequest
	GetGroupItemsResponse
	GetGroupItemDetailRequest
	GroupItem
	GetGroupItemDetailResponse
	DetailItem
	DetailBindRole
	SimpleGroup
	SimpleUserNode
	AddUserRequest
*/
package xbasissvc_external_permission

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

// Client API for UserGroup service

type UserGroupService interface {
	Add(ctx context.Context, in *SimpleGroup, opts ...client.CallOption) (*xbasis_commons_dto.Status, error)
	Remove(ctx context.Context, in *SimpleGroup, opts ...client.CallOption) (*xbasis_commons_dto.Status, error)
	Rename(ctx context.Context, in *SimpleGroup, opts ...client.CallOption) (*xbasis_commons_dto.Status, error)
	AddUser(ctx context.Context, in *AddUserRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error)
	Move(ctx context.Context, in *MoveRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error)
	GetGroupItems(ctx context.Context, in *GetGroupItemsRequest, opts ...client.CallOption) (*GetGroupItemsResponse, error)
	Search(ctx context.Context, in *SearchAppUserRequest, opts ...client.CallOption) (*SearchAppUserResponse, error)
	GetGroupItemDetail(ctx context.Context, in *GetGroupItemDetailRequest, opts ...client.CallOption) (*GetGroupItemDetailResponse, error)
	GetGroupContentSize(ctx context.Context, in *GetGroupContentSizeRequest, opts ...client.CallOption) (*GetGroupContentSizeResponse, error)
}

type userGroupService struct {
	c    client.Client
	name string
}

func NewUserGroupService(name string, c client.Client) UserGroupService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "xbasissvc.external.permission"
	}
	return &userGroupService{
		c:    c,
		name: name,
	}
}

func (c *userGroupService) Add(ctx context.Context, in *SimpleGroup, opts ...client.CallOption) (*xbasis_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "UserGroup.Add", in)
	out := new(xbasis_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupService) Remove(ctx context.Context, in *SimpleGroup, opts ...client.CallOption) (*xbasis_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "UserGroup.Remove", in)
	out := new(xbasis_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupService) Rename(ctx context.Context, in *SimpleGroup, opts ...client.CallOption) (*xbasis_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "UserGroup.Rename", in)
	out := new(xbasis_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupService) AddUser(ctx context.Context, in *AddUserRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "UserGroup.AddUser", in)
	out := new(xbasis_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupService) Move(ctx context.Context, in *MoveRequest, opts ...client.CallOption) (*xbasis_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "UserGroup.Move", in)
	out := new(xbasis_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupService) GetGroupItems(ctx context.Context, in *GetGroupItemsRequest, opts ...client.CallOption) (*GetGroupItemsResponse, error) {
	req := c.c.NewRequest(c.name, "UserGroup.GetGroupItems", in)
	out := new(GetGroupItemsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupService) Search(ctx context.Context, in *SearchAppUserRequest, opts ...client.CallOption) (*SearchAppUserResponse, error) {
	req := c.c.NewRequest(c.name, "UserGroup.Search", in)
	out := new(SearchAppUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupService) GetGroupItemDetail(ctx context.Context, in *GetGroupItemDetailRequest, opts ...client.CallOption) (*GetGroupItemDetailResponse, error) {
	req := c.c.NewRequest(c.name, "UserGroup.GetGroupItemDetail", in)
	out := new(GetGroupItemDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupService) GetGroupContentSize(ctx context.Context, in *GetGroupContentSizeRequest, opts ...client.CallOption) (*GetGroupContentSizeResponse, error) {
	req := c.c.NewRequest(c.name, "UserGroup.GetGroupContentSize", in)
	out := new(GetGroupContentSizeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserGroup service

type UserGroupHandler interface {
	Add(context.Context, *SimpleGroup, *xbasis_commons_dto.Status) error
	Remove(context.Context, *SimpleGroup, *xbasis_commons_dto.Status) error
	Rename(context.Context, *SimpleGroup, *xbasis_commons_dto.Status) error
	AddUser(context.Context, *AddUserRequest, *xbasis_commons_dto.Status) error
	Move(context.Context, *MoveRequest, *xbasis_commons_dto.Status) error
	GetGroupItems(context.Context, *GetGroupItemsRequest, *GetGroupItemsResponse) error
	Search(context.Context, *SearchAppUserRequest, *SearchAppUserResponse) error
	GetGroupItemDetail(context.Context, *GetGroupItemDetailRequest, *GetGroupItemDetailResponse) error
	GetGroupContentSize(context.Context, *GetGroupContentSizeRequest, *GetGroupContentSizeResponse) error
}

func RegisterUserGroupHandler(s server.Server, hdlr UserGroupHandler, opts ...server.HandlerOption) error {
	type userGroup interface {
		Add(ctx context.Context, in *SimpleGroup, out *xbasis_commons_dto.Status) error
		Remove(ctx context.Context, in *SimpleGroup, out *xbasis_commons_dto.Status) error
		Rename(ctx context.Context, in *SimpleGroup, out *xbasis_commons_dto.Status) error
		AddUser(ctx context.Context, in *AddUserRequest, out *xbasis_commons_dto.Status) error
		Move(ctx context.Context, in *MoveRequest, out *xbasis_commons_dto.Status) error
		GetGroupItems(ctx context.Context, in *GetGroupItemsRequest, out *GetGroupItemsResponse) error
		Search(ctx context.Context, in *SearchAppUserRequest, out *SearchAppUserResponse) error
		GetGroupItemDetail(ctx context.Context, in *GetGroupItemDetailRequest, out *GetGroupItemDetailResponse) error
		GetGroupContentSize(ctx context.Context, in *GetGroupContentSizeRequest, out *GetGroupContentSizeResponse) error
	}
	type UserGroup struct {
		userGroup
	}
	h := &userGroupHandler{hdlr}
	return s.Handle(s.NewHandler(&UserGroup{h}, opts...))
}

type userGroupHandler struct {
	UserGroupHandler
}

func (h *userGroupHandler) Add(ctx context.Context, in *SimpleGroup, out *xbasis_commons_dto.Status) error {
	return h.UserGroupHandler.Add(ctx, in, out)
}

func (h *userGroupHandler) Remove(ctx context.Context, in *SimpleGroup, out *xbasis_commons_dto.Status) error {
	return h.UserGroupHandler.Remove(ctx, in, out)
}

func (h *userGroupHandler) Rename(ctx context.Context, in *SimpleGroup, out *xbasis_commons_dto.Status) error {
	return h.UserGroupHandler.Rename(ctx, in, out)
}

func (h *userGroupHandler) AddUser(ctx context.Context, in *AddUserRequest, out *xbasis_commons_dto.Status) error {
	return h.UserGroupHandler.AddUser(ctx, in, out)
}

func (h *userGroupHandler) Move(ctx context.Context, in *MoveRequest, out *xbasis_commons_dto.Status) error {
	return h.UserGroupHandler.Move(ctx, in, out)
}

func (h *userGroupHandler) GetGroupItems(ctx context.Context, in *GetGroupItemsRequest, out *GetGroupItemsResponse) error {
	return h.UserGroupHandler.GetGroupItems(ctx, in, out)
}

func (h *userGroupHandler) Search(ctx context.Context, in *SearchAppUserRequest, out *SearchAppUserResponse) error {
	return h.UserGroupHandler.Search(ctx, in, out)
}

func (h *userGroupHandler) GetGroupItemDetail(ctx context.Context, in *GetGroupItemDetailRequest, out *GetGroupItemDetailResponse) error {
	return h.UserGroupHandler.GetGroupItemDetail(ctx, in, out)
}

func (h *userGroupHandler) GetGroupContentSize(ctx context.Context, in *GetGroupContentSizeRequest, out *GetGroupContentSizeResponse) error {
	return h.UserGroupHandler.GetGroupContentSize(ctx, in, out)
}
