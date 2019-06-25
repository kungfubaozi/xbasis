// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user/pb/invite.proto

/*
Package gosionsvc_external_user is a generated protocol buffer package.

It is generated from these files:
	user/pb/invite.proto

It has these top-level messages:
	InviteUserRequest
	AppendRequest
	SetStateRequest
	HasInvitedRequest
	HasInvitedResponse
	InviteSearchRequest
	InviteSearchResponse
	GetDetailResponse
	InviteItem
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

// Client API for Invite service

type InviteService interface {
	// 邀请用户
	// 邀请的流程并不是直接把用户放在库中
	// 需要被邀请
	// 1.如果没有注册，需要注册再进行
	User(ctx context.Context, in *InviteUserRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
	Search(ctx context.Context, in *InviteSearchRequest, opts ...client.CallOption) (*InviteSearchResponse, error)
	// 是否被邀请
	HasInvited(ctx context.Context, in *HasInvitedRequest, opts ...client.CallOption) (*HasInvitedResponse, error)
	// 获取详情
	GetDetail(ctx context.Context, in *HasInvitedRequest, opts ...client.CallOption) (*GetDetailResponse, error)
	// 完成
	SetState(ctx context.Context, in *SetStateRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
	Append(ctx context.Context, in *AppendRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error)
}

type inviteService struct {
	c    client.Client
	name string
}

func NewInviteService(name string, c client.Client) InviteService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "gosionsvc.external.user"
	}
	return &inviteService{
		c:    c,
		name: name,
	}
}

func (c *inviteService) User(ctx context.Context, in *InviteUserRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Invite.User", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteService) Search(ctx context.Context, in *InviteSearchRequest, opts ...client.CallOption) (*InviteSearchResponse, error) {
	req := c.c.NewRequest(c.name, "Invite.Search", in)
	out := new(InviteSearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteService) HasInvited(ctx context.Context, in *HasInvitedRequest, opts ...client.CallOption) (*HasInvitedResponse, error) {
	req := c.c.NewRequest(c.name, "Invite.HasInvited", in)
	out := new(HasInvitedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteService) GetDetail(ctx context.Context, in *HasInvitedRequest, opts ...client.CallOption) (*GetDetailResponse, error) {
	req := c.c.NewRequest(c.name, "Invite.GetDetail", in)
	out := new(GetDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteService) SetState(ctx context.Context, in *SetStateRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Invite.SetState", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteService) Append(ctx context.Context, in *AppendRequest, opts ...client.CallOption) (*gs_commons_dto.Status, error) {
	req := c.c.NewRequest(c.name, "Invite.Append", in)
	out := new(gs_commons_dto.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Invite service

type InviteHandler interface {
	// 邀请用户
	// 邀请的流程并不是直接把用户放在库中
	// 需要被邀请
	// 1.如果没有注册，需要注册再进行
	User(context.Context, *InviteUserRequest, *gs_commons_dto.Status) error
	Search(context.Context, *InviteSearchRequest, *InviteSearchResponse) error
	// 是否被邀请
	HasInvited(context.Context, *HasInvitedRequest, *HasInvitedResponse) error
	// 获取详情
	GetDetail(context.Context, *HasInvitedRequest, *GetDetailResponse) error
	// 完成
	SetState(context.Context, *SetStateRequest, *gs_commons_dto.Status) error
	Append(context.Context, *AppendRequest, *gs_commons_dto.Status) error
}

func RegisterInviteHandler(s server.Server, hdlr InviteHandler, opts ...server.HandlerOption) error {
	type invite interface {
		User(ctx context.Context, in *InviteUserRequest, out *gs_commons_dto.Status) error
		Search(ctx context.Context, in *InviteSearchRequest, out *InviteSearchResponse) error
		HasInvited(ctx context.Context, in *HasInvitedRequest, out *HasInvitedResponse) error
		GetDetail(ctx context.Context, in *HasInvitedRequest, out *GetDetailResponse) error
		SetState(ctx context.Context, in *SetStateRequest, out *gs_commons_dto.Status) error
		Append(ctx context.Context, in *AppendRequest, out *gs_commons_dto.Status) error
	}
	type Invite struct {
		invite
	}
	h := &inviteHandler{hdlr}
	return s.Handle(s.NewHandler(&Invite{h}, opts...))
}

type inviteHandler struct {
	InviteHandler
}

func (h *inviteHandler) User(ctx context.Context, in *InviteUserRequest, out *gs_commons_dto.Status) error {
	return h.InviteHandler.User(ctx, in, out)
}

func (h *inviteHandler) Search(ctx context.Context, in *InviteSearchRequest, out *InviteSearchResponse) error {
	return h.InviteHandler.Search(ctx, in, out)
}

func (h *inviteHandler) HasInvited(ctx context.Context, in *HasInvitedRequest, out *HasInvitedResponse) error {
	return h.InviteHandler.HasInvited(ctx, in, out)
}

func (h *inviteHandler) GetDetail(ctx context.Context, in *HasInvitedRequest, out *GetDetailResponse) error {
	return h.InviteHandler.GetDetail(ctx, in, out)
}

func (h *inviteHandler) SetState(ctx context.Context, in *SetStateRequest, out *gs_commons_dto.Status) error {
	return h.InviteHandler.SetState(ctx, in, out)
}

func (h *inviteHandler) Append(ctx context.Context, in *AppendRequest, out *gs_commons_dto.Status) error {
	return h.InviteHandler.Append(ctx, in, out)
}
