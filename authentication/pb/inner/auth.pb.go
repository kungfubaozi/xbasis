// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication/pb/inner/auth.proto

package gosionsvc_internal_authentication

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

type VerifyRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	FunctionRoles        []string `protobuf:"bytes,3,rep,name=functionRoles,proto3" json:"functionRoles,omitempty"`
	Funcs                string   `protobuf:"bytes,4,opt,name=funcs,proto3" json:"funcs,omitempty"`
	Access               int64    `protobuf:"varint,5,opt,name=access,proto3" json:"access,omitempty"`
	Share                bool     `protobuf:"varint,6,opt,name=share,proto3" json:"share,omitempty"`
	FunctionId           string   `protobuf:"bytes,7,opt,name=functionId,proto3" json:"functionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRequest) Reset()         { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()    {}
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f12867eb7133256, []int{0}
}

func (m *VerifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyRequest.Unmarshal(m, b)
}
func (m *VerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyRequest.Marshal(b, m, deterministic)
}
func (m *VerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRequest.Merge(m, src)
}
func (m *VerifyRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyRequest.Size(m)
}
func (m *VerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRequest proto.InternalMessageInfo

func (m *VerifyRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *VerifyRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *VerifyRequest) GetFunctionRoles() []string {
	if m != nil {
		return m.FunctionRoles
	}
	return nil
}

func (m *VerifyRequest) GetFuncs() string {
	if m != nil {
		return m.Funcs
	}
	return ""
}

func (m *VerifyRequest) GetAccess() int64 {
	if m != nil {
		return m.Access
	}
	return 0
}

func (m *VerifyRequest) GetShare() bool {
	if m != nil {
		return m.Share
	}
	return false
}

func (m *VerifyRequest) GetFunctionId() string {
	if m != nil {
		return m.FunctionId
	}
	return ""
}

type VerifyResponse struct {
	UserId               string     `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ClientId             string     `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	ClientPlatform       int64      `protobuf:"varint,3,opt,name=clientPlatform,proto3" json:"clientPlatform,omitempty"`
	State                *dto.State `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	AppId                string     `protobuf:"bytes,5,opt,name=appId,proto3" json:"appId,omitempty"`
	Relation             string     `protobuf:"bytes,6,opt,name=relation,proto3" json:"relation,omitempty"`
	AppType              int64      `protobuf:"varint,7,opt,name=appType,proto3" json:"appType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *VerifyResponse) Reset()         { *m = VerifyResponse{} }
func (m *VerifyResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyResponse) ProtoMessage()    {}
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f12867eb7133256, []int{1}
}

func (m *VerifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyResponse.Unmarshal(m, b)
}
func (m *VerifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyResponse.Marshal(b, m, deterministic)
}
func (m *VerifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyResponse.Merge(m, src)
}
func (m *VerifyResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyResponse.Size(m)
}
func (m *VerifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyResponse proto.InternalMessageInfo

func (m *VerifyResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *VerifyResponse) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *VerifyResponse) GetClientPlatform() int64 {
	if m != nil {
		return m.ClientPlatform
	}
	return 0
}

func (m *VerifyResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *VerifyResponse) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *VerifyResponse) GetRelation() string {
	if m != nil {
		return m.Relation
	}
	return ""
}

func (m *VerifyResponse) GetAppType() int64 {
	if m != nil {
		return m.AppType
	}
	return 0
}

func init() {
	proto.RegisterType((*VerifyRequest)(nil), "gosionsvc.internal.authentication.VerifyRequest")
	proto.RegisterType((*VerifyResponse)(nil), "gosionsvc.internal.authentication.VerifyResponse")
}

func init() { proto.RegisterFile("authentication/pb/inner/auth.proto", fileDescriptor_8f12867eb7133256) }

var fileDescriptor_8f12867eb7133256 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6e, 0xdb, 0x20,
	0x18, 0xc7, 0xe5, 0x39, 0x71, 0x62, 0xa6, 0xe4, 0x80, 0xb6, 0x09, 0xe5, 0x30, 0x79, 0xd1, 0x34,
	0x59, 0x8a, 0x84, 0xb7, 0xec, 0x09, 0x76, 0xf4, 0x6d, 0xa2, 0x55, 0xef, 0xc4, 0xfe, 0x92, 0x58,
	0xb6, 0x81, 0x02, 0xae, 0x94, 0x67, 0xec, 0x73, 0xf4, 0x3d, 0x2a, 0x20, 0x8e, 0x9a, 0x1e, 0x5a,
	0xf5, 0xc6, 0xef, 0xe3, 0xe3, 0x0f, 0x3f, 0x00, 0xad, 0xf9, 0x60, 0x8f, 0x20, 0x6c, 0x53, 0x71,
	0xdb, 0x48, 0x51, 0xa8, 0x5d, 0xd1, 0x08, 0x01, 0xba, 0x70, 0x75, 0xaa, 0xb4, 0xb4, 0x12, 0xff,
	0x38, 0x48, 0xd3, 0x48, 0x61, 0x1e, 0x2a, 0xda, 0x08, 0x0b, 0x5a, 0xf0, 0x8e, 0x5e, 0x2f, 0x5b,
	0x6d, 0x5a, 0x29, 0xa0, 0x6d, 0x25, 0xed, 0xa1, 0x08, 0xdd, 0x45, 0x25, 0xfb, 0x5e, 0x0a, 0x53,
	0xd4, 0x56, 0x8e, 0xe3, 0x90, 0xb7, 0x7e, 0x8c, 0xd0, 0xe2, 0x0e, 0x74, 0xb3, 0x3f, 0x31, 0xb8,
	0x1f, 0xc0, 0x58, 0xfc, 0x05, 0x4d, 0xad, 0x6c, 0x41, 0x90, 0x28, 0x8b, 0xf2, 0x94, 0x05, 0xc0,
	0x2b, 0x34, 0xaf, 0xba, 0x06, 0x84, 0x2d, 0x6b, 0xf2, 0xc9, 0x4f, 0x5c, 0x18, 0xff, 0x44, 0x8b,
	0xfd, 0x20, 0x2a, 0xb7, 0x39, 0x93, 0x1d, 0x18, 0x12, 0x67, 0x71, 0x9e, 0xb2, 0xeb, 0xa2, 0xcb,
	0x75, 0x05, 0x43, 0x26, 0x21, 0xd7, 0x03, 0xfe, 0x86, 0x12, 0x5e, 0x55, 0x60, 0x0c, 0x99, 0x66,
	0x51, 0x1e, 0xb3, 0x33, 0xb9, 0x6e, 0x73, 0xe4, 0x1a, 0x48, 0x92, 0x45, 0xf9, 0x9c, 0x05, 0xc0,
	0xdf, 0x11, 0x1a, 0x43, 0xcb, 0x9a, 0xcc, 0x7c, 0xd0, 0x8b, 0xca, 0xfa, 0x29, 0x42, 0xcb, 0xd1,
	0xc6, 0x28, 0x29, 0x0c, 0xb8, 0x0d, 0x06, 0x03, 0xba, 0xac, 0xcf, 0x3e, 0x67, 0x7a, 0x53, 0xe8,
	0x17, 0x5a, 0x86, 0xf1, 0xff, 0x8e, 0xdb, 0xbd, 0xd4, 0x3d, 0x89, 0xfd, 0xe1, 0x5e, 0x55, 0xf1,
	0x06, 0x4d, 0x8d, 0xe5, 0x16, 0xbc, 0xd2, 0xe7, 0xed, 0x57, 0x7a, 0x30, 0x74, 0xbc, 0xde, 0xda,
	0x4a, 0x7a, 0xe3, 0x26, 0x59, 0xe8, 0x71, 0x46, 0x5c, 0xa9, 0xb2, 0xf6, 0xa2, 0x29, 0x0b, 0xe0,
	0x8e, 0xa1, 0xa1, 0xf3, 0x0f, 0xe7, 0x55, 0x53, 0x76, 0x61, 0x4c, 0xd0, 0x8c, 0x2b, 0x75, 0x7b,
	0x52, 0xe0, 0x55, 0x63, 0x36, 0xe2, 0x76, 0x40, 0x93, 0x7f, 0x83, 0x3d, 0xe2, 0x1e, 0x25, 0x41,
	0x17, 0xff, 0xa6, 0xef, 0x7e, 0x0c, 0x7a, 0xf5, 0xce, 0xab, 0x3f, 0x1f, 0x58, 0x11, 0xee, 0x72,
	0x97, 0xf8, 0x3f, 0xf3, 0xf7, 0x39, 0x00, 0x00, 0xff, 0xff, 0x62, 0xcc, 0x37, 0x90, 0xa9, 0x02,
	0x00, 0x00,
}
