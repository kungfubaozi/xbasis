// Code generated by protoc-gen-go. DO NOT EDIT.
// source: permission/pb/group.proto

package gosionsvc_external_permission

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	dto "konekko.me/gosion/commons/dto"
	math "math"
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

type GetGroupItemsRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AppId                string   `protobuf:"bytes,2,opt,name=appId,proto3" json:"appId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupItemsRequest) Reset()         { *m = GetGroupItemsRequest{} }
func (m *GetGroupItemsRequest) String() string { return proto.CompactTextString(m) }
func (*GetGroupItemsRequest) ProtoMessage()    {}
func (*GetGroupItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2256e54f4746997, []int{0}
}

func (m *GetGroupItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupItemsRequest.Unmarshal(m, b)
}
func (m *GetGroupItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupItemsRequest.Marshal(b, m, deterministic)
}
func (m *GetGroupItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupItemsRequest.Merge(m, src)
}
func (m *GetGroupItemsRequest) XXX_Size() int {
	return xxx_messageInfo_GetGroupItemsRequest.Size(m)
}
func (m *GetGroupItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupItemsRequest proto.InternalMessageInfo

func (m *GetGroupItemsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetGroupItemsRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type GetGroupItemsResponse struct {
	State                *dto.State   `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Data                 []*GroupItem `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetGroupItemsResponse) Reset()         { *m = GetGroupItemsResponse{} }
func (m *GetGroupItemsResponse) String() string { return proto.CompactTextString(m) }
func (*GetGroupItemsResponse) ProtoMessage()    {}
func (*GetGroupItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2256e54f4746997, []int{1}
}

func (m *GetGroupItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupItemsResponse.Unmarshal(m, b)
}
func (m *GetGroupItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupItemsResponse.Marshal(b, m, deterministic)
}
func (m *GetGroupItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupItemsResponse.Merge(m, src)
}
func (m *GetGroupItemsResponse) XXX_Size() int {
	return xxx_messageInfo_GetGroupItemsResponse.Size(m)
}
func (m *GetGroupItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupItemsResponse proto.InternalMessageInfo

func (m *GetGroupItemsResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *GetGroupItemsResponse) GetData() []*GroupItem {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetGroupItemDetailRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AppId                string   `protobuf:"bytes,2,opt,name=appId,proto3" json:"appId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupItemDetailRequest) Reset()         { *m = GetGroupItemDetailRequest{} }
func (m *GetGroupItemDetailRequest) String() string { return proto.CompactTextString(m) }
func (*GetGroupItemDetailRequest) ProtoMessage()    {}
func (*GetGroupItemDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2256e54f4746997, []int{2}
}

func (m *GetGroupItemDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupItemDetailRequest.Unmarshal(m, b)
}
func (m *GetGroupItemDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupItemDetailRequest.Marshal(b, m, deterministic)
}
func (m *GetGroupItemDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupItemDetailRequest.Merge(m, src)
}
func (m *GetGroupItemDetailRequest) XXX_Size() int {
	return xxx_messageInfo_GetGroupItemDetailRequest.Size(m)
}
func (m *GetGroupItemDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupItemDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupItemDetailRequest proto.InternalMessageInfo

func (m *GetGroupItemDetailRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetGroupItemDetailRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type GroupItem struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Group                bool     `protobuf:"varint,2,opt,name=group,proto3" json:"group,omitempty"`
	Icon                 string   `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Id                   string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupItem) Reset()         { *m = GroupItem{} }
func (m *GroupItem) String() string { return proto.CompactTextString(m) }
func (*GroupItem) ProtoMessage()    {}
func (*GroupItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2256e54f4746997, []int{3}
}

func (m *GroupItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupItem.Unmarshal(m, b)
}
func (m *GroupItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupItem.Marshal(b, m, deterministic)
}
func (m *GroupItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupItem.Merge(m, src)
}
func (m *GroupItem) XXX_Size() int {
	return xxx_messageInfo_GroupItem.Size(m)
}
func (m *GroupItem) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupItem.DiscardUnknown(m)
}

var xxx_messageInfo_GroupItem proto.InternalMessageInfo

func (m *GroupItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GroupItem) GetGroup() bool {
	if m != nil {
		return m.Group
	}
	return false
}

func (m *GroupItem) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *GroupItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetGroupItemDetailResponse struct {
	State                *dto.State  `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Data                 *DetailItem `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetGroupItemDetailResponse) Reset()         { *m = GetGroupItemDetailResponse{} }
func (m *GetGroupItemDetailResponse) String() string { return proto.CompactTextString(m) }
func (*GetGroupItemDetailResponse) ProtoMessage()    {}
func (*GetGroupItemDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2256e54f4746997, []int{4}
}

func (m *GetGroupItemDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupItemDetailResponse.Unmarshal(m, b)
}
func (m *GetGroupItemDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupItemDetailResponse.Marshal(b, m, deterministic)
}
func (m *GetGroupItemDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupItemDetailResponse.Merge(m, src)
}
func (m *GetGroupItemDetailResponse) XXX_Size() int {
	return xxx_messageInfo_GetGroupItemDetailResponse.Size(m)
}
func (m *GetGroupItemDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupItemDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupItemDetailResponse proto.InternalMessageInfo

func (m *GetGroupItemDetailResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *GetGroupItemDetailResponse) GetData() *DetailItem {
	if m != nil {
		return m.Data
	}
	return nil
}

type DetailItem struct {
	Username             string            `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Icon                 string            `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty"`
	Roles                []*DetailBindRole `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	RealName             string            `protobuf:"bytes,4,opt,name=realName,proto3" json:"realName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DetailItem) Reset()         { *m = DetailItem{} }
func (m *DetailItem) String() string { return proto.CompactTextString(m) }
func (*DetailItem) ProtoMessage()    {}
func (*DetailItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2256e54f4746997, []int{5}
}

func (m *DetailItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailItem.Unmarshal(m, b)
}
func (m *DetailItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailItem.Marshal(b, m, deterministic)
}
func (m *DetailItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailItem.Merge(m, src)
}
func (m *DetailItem) XXX_Size() int {
	return xxx_messageInfo_DetailItem.Size(m)
}
func (m *DetailItem) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailItem.DiscardUnknown(m)
}

var xxx_messageInfo_DetailItem proto.InternalMessageInfo

func (m *DetailItem) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *DetailItem) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *DetailItem) GetRoles() []*DetailBindRole {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *DetailItem) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

type DetailBindRole struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailBindRole) Reset()         { *m = DetailBindRole{} }
func (m *DetailBindRole) String() string { return proto.CompactTextString(m) }
func (*DetailBindRole) ProtoMessage()    {}
func (*DetailBindRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2256e54f4746997, []int{6}
}

func (m *DetailBindRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailBindRole.Unmarshal(m, b)
}
func (m *DetailBindRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailBindRole.Marshal(b, m, deterministic)
}
func (m *DetailBindRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailBindRole.Merge(m, src)
}
func (m *DetailBindRole) XXX_Size() int {
	return xxx_messageInfo_DetailBindRole.Size(m)
}
func (m *DetailBindRole) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailBindRole.DiscardUnknown(m)
}

var xxx_messageInfo_DetailBindRole proto.InternalMessageInfo

func (m *DetailBindRole) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DetailBindRole) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SimpleGroup struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	BindGroupId          string   `protobuf:"bytes,4,opt,name=bindGroupId,proto3" json:"bindGroupId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleGroup) Reset()         { *m = SimpleGroup{} }
func (m *SimpleGroup) String() string { return proto.CompactTextString(m) }
func (*SimpleGroup) ProtoMessage()    {}
func (*SimpleGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2256e54f4746997, []int{7}
}

func (m *SimpleGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleGroup.Unmarshal(m, b)
}
func (m *SimpleGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleGroup.Marshal(b, m, deterministic)
}
func (m *SimpleGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleGroup.Merge(m, src)
}
func (m *SimpleGroup) XXX_Size() int {
	return xxx_messageInfo_SimpleGroup.Size(m)
}
func (m *SimpleGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleGroup.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleGroup proto.InternalMessageInfo

func (m *SimpleGroup) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *SimpleGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SimpleGroup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SimpleGroup) GetBindGroupId() string {
	if m != nil {
		return m.BindGroupId
	}
	return ""
}

type SimpleUserNode struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleUserNode) Reset()         { *m = SimpleUserNode{} }
func (m *SimpleUserNode) String() string { return proto.CompactTextString(m) }
func (*SimpleUserNode) ProtoMessage()    {}
func (*SimpleUserNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2256e54f4746997, []int{8}
}

func (m *SimpleUserNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleUserNode.Unmarshal(m, b)
}
func (m *SimpleUserNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleUserNode.Marshal(b, m, deterministic)
}
func (m *SimpleUserNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleUserNode.Merge(m, src)
}
func (m *SimpleUserNode) XXX_Size() int {
	return xxx_messageInfo_SimpleUserNode.Size(m)
}
func (m *SimpleUserNode) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleUserNode.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleUserNode proto.InternalMessageInfo

func (m *SimpleUserNode) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetGroupItemsRequest)(nil), "gosionsvc.external.permission.GetGroupItemsRequest")
	proto.RegisterType((*GetGroupItemsResponse)(nil), "gosionsvc.external.permission.GetGroupItemsResponse")
	proto.RegisterType((*GetGroupItemDetailRequest)(nil), "gosionsvc.external.permission.GetGroupItemDetailRequest")
	proto.RegisterType((*GroupItem)(nil), "gosionsvc.external.permission.GroupItem")
	proto.RegisterType((*GetGroupItemDetailResponse)(nil), "gosionsvc.external.permission.GetGroupItemDetailResponse")
	proto.RegisterType((*DetailItem)(nil), "gosionsvc.external.permission.DetailItem")
	proto.RegisterType((*DetailBindRole)(nil), "gosionsvc.external.permission.DetailBindRole")
	proto.RegisterType((*SimpleGroup)(nil), "gosionsvc.external.permission.SimpleGroup")
	proto.RegisterType((*SimpleUserNode)(nil), "gosionsvc.external.permission.SimpleUserNode")
}

func init() { proto.RegisterFile("permission/pb/group.proto", fileDescriptor_b2256e54f4746997) }

var fileDescriptor_b2256e54f4746997 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x41, 0x6f, 0x12, 0x41,
	0x14, 0xce, 0x2e, 0x05, 0xe1, 0x11, 0x39, 0x4c, 0xda, 0x86, 0x6e, 0x62, 0x42, 0xf6, 0x84, 0x36,
	0x5d, 0x12, 0xda, 0x83, 0x26, 0xf5, 0x80, 0x9a, 0x34, 0x24, 0xb5, 0xd1, 0xad, 0x3d, 0x78, 0x5c,
	0x98, 0x17, 0x32, 0x61, 0x77, 0x66, 0xdd, 0x19, 0x88, 0xf1, 0xe8, 0x49, 0x7f, 0x84, 0xbf, 0xc4,
	0x3f, 0x67, 0x66, 0x06, 0x86, 0xc5, 0xa2, 0x94, 0x86, 0xdb, 0xcc, 0xf0, 0xbe, 0xef, 0x7b, 0xef,
	0xfb, 0x36, 0x0f, 0x38, 0xc9, 0xb1, 0xc8, 0x98, 0x94, 0x4c, 0xf0, 0x5e, 0x3e, 0xea, 0x4d, 0x0a,
	0x31, 0xcb, 0xa3, 0xbc, 0x10, 0x4a, 0x90, 0x67, 0x13, 0xa1, 0x9f, 0xe5, 0x7c, 0x1c, 0xe1, 0x57,
	0x85, 0x05, 0x4f, 0xd2, 0x68, 0x55, 0x1d, 0x9c, 0x4e, 0x05, 0xc7, 0xe9, 0x54, 0x44, 0x19, 0xf6,
	0x6c, 0x65, 0x6f, 0x2c, 0xb2, 0x4c, 0x70, 0xd9, 0xa3, 0x4a, 0x2c, 0xcf, 0x96, 0x2b, 0xbc, 0x84,
	0xc3, 0x2b, 0x54, 0x57, 0x9a, 0x7d, 0xa8, 0x30, 0x93, 0x31, 0x7e, 0x99, 0xa1, 0x54, 0xa4, 0x05,
	0x3e, 0xa3, 0x6d, 0xaf, 0xe3, 0x75, 0x1b, 0xb1, 0xcf, 0x28, 0x39, 0x84, 0x6a, 0x92, 0xe7, 0x43,
	0xda, 0xf6, 0xcd, 0x93, 0xbd, 0x84, 0xdf, 0x3d, 0x38, 0xfa, 0x0b, 0x2e, 0x73, 0xc1, 0x25, 0x92,
	0x53, 0xa8, 0x4a, 0x95, 0x28, 0x34, 0x14, 0xcd, 0xfe, 0x51, 0x34, 0x91, 0xd1, 0x52, 0x99, 0x2a,
	0x11, 0xdd, 0xea, 0x1f, 0x63, 0x5b, 0x43, 0x2e, 0xe1, 0x80, 0x26, 0x2a, 0x69, 0xfb, 0x9d, 0x4a,
	0xb7, 0xd9, 0xef, 0x46, 0xff, 0x9d, 0x2f, 0x72, 0x6a, 0xb1, 0x41, 0x85, 0x03, 0x38, 0x29, 0xf7,
	0xf0, 0x0e, 0x55, 0xc2, 0xd2, 0xdd, 0xe6, 0xf8, 0x0c, 0x0d, 0x87, 0x27, 0x04, 0x0e, 0x78, 0x92,
	0xe1, 0x02, 0x64, 0xce, 0x1a, 0x66, 0x12, 0x30, 0xb0, 0x7a, 0x6c, 0x2f, 0xba, 0x92, 0x8d, 0x05,
	0x6f, 0x57, 0x6c, 0xa5, 0x3e, 0x2f, 0x04, 0x0f, 0x96, 0x82, 0xe1, 0x0f, 0x0f, 0x82, 0x4d, 0xed,
	0x3d, 0xc6, 0xa7, 0xd7, 0xce, 0x27, 0x5d, 0xfb, 0x7c, 0x8b, 0x4f, 0x56, 0xa9, 0x64, 0xd4, 0x2f,
	0x0f, 0x60, 0xf5, 0x48, 0x02, 0xa8, 0xcf, 0xa4, 0x46, 0xb9, 0x59, 0xdd, 0xdd, 0x4d, 0xe6, 0x97,
	0x26, 0x7b, 0x0b, 0xd5, 0x42, 0xa4, 0x28, 0xdb, 0x15, 0x13, 0xd3, 0xd9, 0x83, 0xe4, 0xdf, 0x30,
	0x4e, 0x63, 0x91, 0x62, 0x6c, 0xb1, 0x5a, 0xb4, 0xc0, 0x24, 0xbd, 0xd1, 0xa2, 0xd6, 0x24, 0x77,
	0x0f, 0x2f, 0xa0, 0xb5, 0x0e, 0xda, 0x18, 0x85, 0x35, 0xd8, 0x77, 0x06, 0x33, 0x68, 0xde, 0xb2,
	0x2c, 0x4f, 0xd1, 0x58, 0xbc, 0x0a, 0xd8, 0x2b, 0x05, 0xec, 0x88, 0xfc, 0x7b, 0x44, 0x15, 0xf7,
	0x69, 0x74, 0xa0, 0x39, 0x62, 0x9c, 0xda, 0xa4, 0x96, 0x11, 0x96, 0x9f, 0xc2, 0x2e, 0xb4, 0xac,
	0xd4, 0x9d, 0xc4, 0xe2, 0x46, 0x50, 0x24, 0xc7, 0x50, 0xd3, 0x9e, 0x39, 0xb9, 0xc5, 0xad, 0xff,
	0xbb, 0x06, 0x0d, 0x5d, 0x64, 0x7b, 0xba, 0x86, 0xda, 0x35, 0xe3, 0xd3, 0x4f, 0x82, 0xbc, 0xd8,
	0x62, 0x5a, 0x69, 0x92, 0xe0, 0x78, 0xd3, 0xb7, 0x30, 0x93, 0x9a, 0xed, 0x8e, 0xa7, 0x8c, 0x4f,
	0xf7, 0xc2, 0x36, 0x84, 0xca, 0x80, 0xd2, 0x7d, 0x35, 0x16, 0x63, 0x26, 0xe6, 0xb8, 0x3f, 0x36,
	0x13, 0xd4, 0x3e, 0xd8, 0x3e, 0xc0, 0x93, 0x01, 0xa5, 0x3a, 0x12, 0x72, 0xf6, 0x20, 0xba, 0x65,
	0xc4, 0xff, 0x64, 0xfc, 0x08, 0xf5, 0xf7, 0x62, 0x8e, 0xfb, 0xa4, 0xfc, 0x06, 0x4f, 0xd7, 0xb6,
	0x29, 0x39, 0xdf, 0xb6, 0x0a, 0x37, 0xac, 0xee, 0xe0, 0x62, 0x37, 0xd0, 0x62, 0x11, 0xfd, 0xf4,
	0x80, 0xdc, 0xdf, 0x53, 0xe4, 0xe5, 0x0e, 0x64, 0x6b, 0x9b, 0x37, 0x78, 0xf5, 0x08, 0xa4, 0xed,
	0x65, 0x54, 0x33, 0xff, 0x4d, 0xe7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xde, 0x2a, 0x2e,
	0x04, 0x07, 0x00, 0x00,
}
