// Code generated by protoc-gen-go. DO NOT EDIT.
// source: permission/pb/binding.proto

package xbasissvc_external_permission

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	dto "konekko.me/xbasis/commons/dto"
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

type GetTargetBindRolesRequest struct {
	User                 bool     `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	AppId                string   `protobuf:"bytes,3,opt,name=appId,proto3" json:"appId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTargetBindRolesRequest) Reset()         { *m = GetTargetBindRolesRequest{} }
func (m *GetTargetBindRolesRequest) String() string { return proto.CompactTextString(m) }
func (*GetTargetBindRolesRequest) ProtoMessage()    {}
func (*GetTargetBindRolesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d6cdac82fda43c5, []int{0}
}

func (m *GetTargetBindRolesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTargetBindRolesRequest.Unmarshal(m, b)
}
func (m *GetTargetBindRolesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTargetBindRolesRequest.Marshal(b, m, deterministic)
}
func (m *GetTargetBindRolesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTargetBindRolesRequest.Merge(m, src)
}
func (m *GetTargetBindRolesRequest) XXX_Size() int {
	return xxx_messageInfo_GetTargetBindRolesRequest.Size(m)
}
func (m *GetTargetBindRolesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTargetBindRolesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTargetBindRolesRequest proto.InternalMessageInfo

func (m *GetTargetBindRolesRequest) GetUser() bool {
	if m != nil {
		return m.User
	}
	return false
}

func (m *GetTargetBindRolesRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetTargetBindRolesRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type GetTargetBindRolesResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Data                 []string   `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetTargetBindRolesResponse) Reset()         { *m = GetTargetBindRolesResponse{} }
func (m *GetTargetBindRolesResponse) String() string { return proto.CompactTextString(m) }
func (*GetTargetBindRolesResponse) ProtoMessage()    {}
func (*GetTargetBindRolesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d6cdac82fda43c5, []int{1}
}

func (m *GetTargetBindRolesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTargetBindRolesResponse.Unmarshal(m, b)
}
func (m *GetTargetBindRolesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTargetBindRolesResponse.Marshal(b, m, deterministic)
}
func (m *GetTargetBindRolesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTargetBindRolesResponse.Merge(m, src)
}
func (m *GetTargetBindRolesResponse) XXX_Size() int {
	return xxx_messageInfo_GetTargetBindRolesResponse.Size(m)
}
func (m *GetTargetBindRolesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTargetBindRolesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTargetBindRolesResponse proto.InternalMessageInfo

func (m *GetTargetBindRolesResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *GetTargetBindRolesResponse) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

type BindingRolesRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Roles                []string `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	Override             bool     `protobuf:"varint,4,opt,name=override,proto3" json:"override,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindingRolesRequest) Reset()         { *m = BindingRolesRequest{} }
func (m *BindingRolesRequest) String() string { return proto.CompactTextString(m) }
func (*BindingRolesRequest) ProtoMessage()    {}
func (*BindingRolesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d6cdac82fda43c5, []int{2}
}

func (m *BindingRolesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindingRolesRequest.Unmarshal(m, b)
}
func (m *BindingRolesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindingRolesRequest.Marshal(b, m, deterministic)
}
func (m *BindingRolesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindingRolesRequest.Merge(m, src)
}
func (m *BindingRolesRequest) XXX_Size() int {
	return xxx_messageInfo_BindingRolesRequest.Size(m)
}
func (m *BindingRolesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BindingRolesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BindingRolesRequest proto.InternalMessageInfo

func (m *BindingRolesRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *BindingRolesRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BindingRolesRequest) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *BindingRolesRequest) GetOverride() bool {
	if m != nil {
		return m.Override
	}
	return false
}

type BindingRoleRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	RoleId               string   `protobuf:"bytes,3,opt,name=roleId,proto3" json:"roleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindingRoleRequest) Reset()         { *m = BindingRoleRequest{} }
func (m *BindingRoleRequest) String() string { return proto.CompactTextString(m) }
func (*BindingRoleRequest) ProtoMessage()    {}
func (*BindingRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d6cdac82fda43c5, []int{3}
}

func (m *BindingRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindingRoleRequest.Unmarshal(m, b)
}
func (m *BindingRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindingRoleRequest.Marshal(b, m, deterministic)
}
func (m *BindingRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindingRoleRequest.Merge(m, src)
}
func (m *BindingRoleRequest) XXX_Size() int {
	return xxx_messageInfo_BindingRoleRequest.Size(m)
}
func (m *BindingRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BindingRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BindingRoleRequest proto.InternalMessageInfo

func (m *BindingRoleRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *BindingRoleRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BindingRoleRequest) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetTargetBindRolesRequest)(nil), "xbasissvc.external.permission.GetTargetBindRolesRequest")
	proto.RegisterType((*GetTargetBindRolesResponse)(nil), "xbasissvc.external.permission.GetTargetBindRolesResponse")
	proto.RegisterType((*BindingRolesRequest)(nil), "xbasissvc.external.permission.BindingRolesRequest")
	proto.RegisterType((*BindingRoleRequest)(nil), "xbasissvc.external.permission.BindingRoleRequest")
}

func init() { proto.RegisterFile("permission/pb/binding.proto", fileDescriptor_5d6cdac82fda43c5) }

var fileDescriptor_5d6cdac82fda43c5 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0x91, 0xfc, 0xa7, 0xf6, 0xb4, 0xf8, 0xb0, 0x2d, 0x45, 0x56, 0x29, 0x18, 0x9d, 0x0c,
	0x85, 0x15, 0x75, 0x2f, 0xed, 0xd5, 0x87, 0x86, 0x5c, 0x95, 0xf8, 0x92, 0x43, 0xc2, 0x4a, 0x3b,
	0x98, 0xc5, 0xd6, 0xae, 0xb2, 0xbb, 0x32, 0x7e, 0x85, 0x3c, 0x41, 0x5e, 0x37, 0x48, 0x6b, 0x2b,
	0x36, 0xb1, 0x13, 0x27, 0x24, 0xb7, 0x19, 0x31, 0xfb, 0xfd, 0xbe, 0x19, 0xcd, 0xc0, 0x8f, 0x02,
	0x75, 0x2e, 0x8c, 0x11, 0x4a, 0xc6, 0x45, 0x1a, 0xa7, 0x42, 0x72, 0x21, 0xe7, 0xb4, 0xd0, 0xca,
	0x2a, 0xf2, 0x73, 0x9d, 0x32, 0x23, 0x8c, 0x59, 0x65, 0x14, 0xd7, 0x16, 0xb5, 0x64, 0x4b, 0xfa,
	0x58, 0x1f, 0xfe, 0x5a, 0x28, 0x89, 0x8b, 0x85, 0xa2, 0x39, 0xc6, 0xae, 0x32, 0xce, 0x54, 0x9e,
	0x2b, 0x69, 0x62, 0x6e, 0xd5, 0x36, 0x76, 0x5a, 0xd1, 0x0c, 0x86, 0x67, 0x68, 0x2f, 0x99, 0x9e,
	0xa3, 0x9d, 0x0a, 0xc9, 0x13, 0xb5, 0x44, 0x93, 0xe0, 0x6d, 0x89, 0xc6, 0x12, 0x02, 0xed, 0xd2,
	0xa0, 0x0e, 0xbc, 0x91, 0x37, 0xee, 0x25, 0x75, 0x4c, 0x06, 0xe0, 0x0b, 0x1e, 0xf8, 0x23, 0x6f,
	0xdc, 0x4f, 0x7c, 0xc1, 0xc9, 0x37, 0xe8, 0xb0, 0xa2, 0x38, 0xe7, 0x41, 0xab, 0xfe, 0xe4, 0x92,
	0x88, 0x41, 0x78, 0x48, 0xd6, 0x14, 0x4a, 0x1a, 0x24, 0x31, 0x74, 0x8c, 0x65, 0x16, 0x6b, 0xe1,
	0xcf, 0x93, 0x21, 0x75, 0x36, 0xe9, 0xd6, 0x1a, 0xb7, 0x8a, 0x5e, 0x54, 0x05, 0x89, 0xab, 0xab,
	0x8c, 0x70, 0x66, 0x59, 0xe0, 0x8f, 0x5a, 0xe3, 0x7e, 0x52, 0xc7, 0x51, 0x0e, 0x5f, 0xa7, 0x6e,
	0x2c, 0x7b, 0x9e, 0x1b, 0x3f, 0xde, 0x8e, 0x9f, 0x43, 0xae, 0x75, 0xf5, 0x2a, 0x68, 0xd5, 0x8a,
	0x2e, 0x21, 0x21, 0xf4, 0xd4, 0x0a, 0xb5, 0x16, 0x1c, 0x83, 0x76, 0xdd, 0x73, 0x93, 0x47, 0x09,
	0x90, 0x1d, 0xdc, 0xeb, 0x68, 0xdf, 0xa1, 0x5b, 0x01, 0x9a, 0x21, 0x6d, 0xb2, 0xc9, 0x7d, 0x1b,
	0x3e, 0x6d, 0x44, 0xc9, 0x15, 0xf4, 0x66, 0x06, 0x75, 0x25, 0x4e, 0x26, 0xf4, 0xd9, 0x3f, 0x4c,
	0x0f, 0xf4, 0x1d, 0x86, 0xc7, 0x86, 0x58, 0x1a, 0x72, 0x0d, 0x5f, 0xfe, 0x97, 0x32, 0xb3, 0x42,
	0xc9, 0x0f, 0xd1, 0xbf, 0x81, 0xc1, 0x4c, 0x56, 0x3b, 0xda, 0x74, 0xf0, 0xfb, 0x74, 0xc2, 0x29,
	0x80, 0x0c, 0x88, 0x03, 0xec, 0xb5, 0xf1, 0xce, 0x90, 0x3b, 0x0f, 0xc8, 0xd3, 0xa5, 0x25, 0x7f,
	0x5f, 0xa0, 0x1c, 0x3d, 0x9f, 0xf0, 0xdf, 0x1b, 0x5e, 0xba, 0x0b, 0x49, 0xbb, 0xf5, 0x75, 0xfe,
	0x79, 0x08, 0x00, 0x00, 0xff, 0xff, 0x34, 0x7b, 0x12, 0x9e, 0x08, 0x04, 0x00, 0x00,
}
