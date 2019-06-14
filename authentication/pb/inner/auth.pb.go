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
	Access               int64    `protobuf:"varint,4,opt,name=access,proto3" json:"access,omitempty"`
	Share                bool     `protobuf:"varint,5,opt,name=share,proto3" json:"share,omitempty"`
	FunctionId           string   `protobuf:"bytes,6,opt,name=functionId,proto3" json:"functionId,omitempty"`
	AppId                string   `protobuf:"bytes,7,opt,name=appId,proto3" json:"appId,omitempty"`
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

func (m *VerifyRequest) GetAppId() string {
	if m != nil {
		return m.AppId
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
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xd1, 0x6e, 0xdb, 0x20,
	0x14, 0x86, 0xe5, 0x39, 0x76, 0x12, 0xa6, 0xe4, 0x02, 0x6d, 0x13, 0xca, 0xc5, 0xe4, 0x45, 0xd3,
	0x64, 0x29, 0x12, 0xde, 0xb2, 0x27, 0xd8, 0xa5, 0xef, 0x26, 0x5a, 0xf5, 0x9e, 0xd8, 0x27, 0x89,
	0x65, 0x1b, 0x28, 0xe0, 0x4a, 0x79, 0xc6, 0x3e, 0x47, 0xdf, 0xa3, 0x32, 0xd8, 0x69, 0xd2, 0x8b,
	0x56, 0xbd, 0xe3, 0x3b, 0x70, 0xf8, 0xff, 0x1f, 0x0e, 0x5a, 0xf3, 0xce, 0x1e, 0x41, 0xd8, 0xaa,
	0xe0, 0xb6, 0x92, 0x22, 0x53, 0xbb, 0xac, 0x12, 0x02, 0x74, 0xd6, 0xd7, 0xa9, 0xd2, 0xd2, 0x4a,
	0xfc, 0xe3, 0x20, 0x4d, 0x25, 0x85, 0x79, 0x28, 0x68, 0x25, 0x2c, 0x68, 0xc1, 0x1b, 0x7a, 0xdd,
	0xb6, 0xda, 0xd4, 0x52, 0x40, 0x5d, 0x4b, 0xda, 0x42, 0xe6, 0x4f, 0x67, 0x85, 0x6c, 0x5b, 0x29,
	0x4c, 0x56, 0x5a, 0x39, 0xae, 0xfd, 0x7d, 0xeb, 0xc7, 0x00, 0x2d, 0xee, 0x40, 0x57, 0xfb, 0x13,
	0x83, 0xfb, 0x0e, 0x8c, 0xc5, 0x5f, 0x50, 0x64, 0x65, 0x0d, 0x82, 0x04, 0x49, 0x90, 0xce, 0x99,
	0x07, 0xbc, 0x42, 0xb3, 0xa2, 0xa9, 0x40, 0xd8, 0xbc, 0x24, 0x9f, 0xdc, 0xc6, 0x99, 0xf1, 0x4f,
	0xb4, 0xd8, 0x77, 0xa2, 0xe8, 0xc5, 0x99, 0x6c, 0xc0, 0x90, 0x30, 0x09, 0xd3, 0x39, 0xbb, 0x2e,
	0xe2, 0x6f, 0x28, 0xe6, 0x45, 0x01, 0xc6, 0x90, 0x49, 0x12, 0xa4, 0x21, 0x1b, 0xa8, 0xd7, 0x33,
	0x47, 0xae, 0x81, 0x44, 0x49, 0x90, 0xce, 0x98, 0x07, 0xfc, 0x1d, 0xa1, 0xb1, 0x3d, 0x2f, 0x49,
	0xec, 0x14, 0x2f, 0x2a, 0x7d, 0x17, 0x57, 0x2a, 0x2f, 0xc9, 0xd4, 0xbb, 0x74, 0xb0, 0x7e, 0x0a,
	0xd0, 0x72, 0x4c, 0x63, 0x94, 0x14, 0x06, 0x7a, 0xd9, 0xce, 0x80, 0xce, 0xcb, 0x21, 0xcf, 0x40,
	0x6f, 0x06, 0xfa, 0x85, 0x96, 0x7e, 0xfd, 0xbf, 0xe1, 0x76, 0x2f, 0x75, 0x4b, 0x42, 0x67, 0xf9,
	0x55, 0x15, 0x6f, 0x50, 0x64, 0x2c, 0xb7, 0xe0, 0x12, 0x7d, 0xde, 0x7e, 0xa5, 0x07, 0x43, 0xc7,
	0xe7, 0x2d, 0xad, 0xa4, 0x37, 0xfd, 0x26, 0xf3, 0x67, 0x5e, 0x1c, 0x47, 0x17, 0x8e, 0x7b, 0x1b,
	0x1a, 0x1a, 0xf7, 0x71, 0x43, 0xca, 0x33, 0x63, 0x82, 0xa6, 0x5c, 0xa9, 0xdb, 0x93, 0x02, 0x97,
	0x32, 0x64, 0x23, 0x6e, 0x3b, 0x34, 0xf9, 0xd7, 0xd9, 0x23, 0x6e, 0x51, 0xec, 0xe3, 0xe2, 0xdf,
	0xf4, 0xdd, 0xc1, 0xa0, 0x57, 0xff, 0xbc, 0xfa, 0xf3, 0x81, 0x0e, 0xff, 0x96, 0xbb, 0xd8, 0xcd,
	0xcc, 0xdf, 0xe7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xdd, 0x56, 0x92, 0xa9, 0x02, 0x00, 0x00,
}