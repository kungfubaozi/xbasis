// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication/pb/inner/auth.proto

package xbasissvc_internal_authentication

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
	proto.RegisterType((*VerifyRequest)(nil), "xbasissvc.internal.authentication.VerifyRequest")
	proto.RegisterType((*VerifyResponse)(nil), "xbasissvc.internal.authentication.VerifyResponse")
}

func init() { proto.RegisterFile("authentication/pb/inner/auth.proto", fileDescriptor_8f12867eb7133256) }

var fileDescriptor_8f12867eb7133256 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6e, 0xa3, 0x30,
	0x10, 0x86, 0xc5, 0x12, 0x48, 0xe2, 0x55, 0x72, 0xb0, 0x56, 0x2b, 0x6f, 0x0e, 0x2b, 0x36, 0x5a,
	0xad, 0x90, 0x56, 0x32, 0x6d, 0xfa, 0x04, 0x3d, 0x72, 0xab, 0xdc, 0xaa, 0x77, 0x07, 0x26, 0x0a,
	0x02, 0x6c, 0xd7, 0x36, 0x55, 0xf3, 0x8c, 0x7d, 0x92, 0xbe, 0x45, 0x85, 0x0d, 0x69, 0xd2, 0x43,
	0xab, 0xde, 0xfc, 0x0d, 0x33, 0xfc, 0xff, 0x3f, 0x36, 0x5a, 0xf3, 0xce, 0xee, 0x41, 0xd8, 0xaa,
	0xe0, 0xb6, 0x92, 0x22, 0x53, 0xdb, 0xac, 0x12, 0x02, 0x74, 0xd6, 0xd7, 0xa9, 0xd2, 0xd2, 0x4a,
	0xfc, 0xe7, 0x69, 0xcb, 0x4d, 0x65, 0xcc, 0x63, 0x41, 0x2b, 0x61, 0x41, 0x0b, 0xde, 0xd0, 0xf3,
	0xb1, 0xd5, 0xff, 0x5a, 0x0a, 0xa8, 0x6b, 0x49, 0x5b, 0xc8, 0x7c, 0x77, 0x56, 0xc8, 0xb6, 0x95,
	0xc2, 0x64, 0xa5, 0x95, 0xe3, 0xd9, 0xff, 0x6f, 0xfd, 0x1c, 0xa0, 0xc5, 0x3d, 0xe8, 0x6a, 0x77,
	0x60, 0xf0, 0xd0, 0x81, 0xb1, 0xf8, 0x07, 0x8a, 0xac, 0xac, 0x41, 0x90, 0x20, 0x09, 0xd2, 0x39,
	0xf3, 0x80, 0x57, 0x68, 0x56, 0x34, 0x15, 0x08, 0x9b, 0x97, 0xe4, 0x9b, 0xfb, 0x70, 0x64, 0xfc,
	0x17, 0x2d, 0x76, 0x9d, 0x28, 0x7a, 0x71, 0x26, 0x1b, 0x30, 0x24, 0x4c, 0xc2, 0x74, 0xce, 0xce,
	0x8b, 0xf8, 0x27, 0x8a, 0x79, 0x51, 0x80, 0x31, 0x64, 0x92, 0x04, 0x69, 0xc8, 0x06, 0xea, 0xf5,
	0xcc, 0x9e, 0x6b, 0x20, 0x51, 0x12, 0xa4, 0x33, 0xe6, 0x01, 0xff, 0x46, 0x68, 0x1c, 0xcf, 0x4b,
	0x12, 0x3b, 0xc5, 0x93, 0x4a, 0x3f, 0xc5, 0x95, 0xca, 0x4b, 0x32, 0xf5, 0x2e, 0x1d, 0xac, 0x5f,
	0x02, 0xb4, 0x1c, 0xd3, 0x18, 0x25, 0x85, 0x81, 0x5e, 0xb6, 0x33, 0xa0, 0xf3, 0x72, 0xc8, 0x33,
	0xd0, 0x87, 0x81, 0xfe, 0xa1, 0xa5, 0x3f, 0xdf, 0x34, 0xdc, 0xee, 0xa4, 0x6e, 0x49, 0xe8, 0x2c,
	0xbf, 0xab, 0xe2, 0x0c, 0x45, 0xc6, 0x72, 0x0b, 0x2e, 0xd1, 0xf7, 0xcd, 0x2f, 0xea, 0xd7, 0x4d,
	0xc7, 0x15, 0x97, 0x56, 0xd2, 0xdb, 0xbe, 0x81, 0xf9, 0xbe, 0x37, 0xd7, 0xd1, 0x89, 0xeb, 0xde,
	0x8a, 0x86, 0xc6, 0x5d, 0xde, 0x90, 0xf4, 0xc8, 0x98, 0xa0, 0x29, 0x57, 0xea, 0xee, 0xa0, 0xc0,
	0x25, 0x0d, 0xd9, 0x88, 0x9b, 0x0e, 0x4d, 0xae, 0x3b, 0xbb, 0xc7, 0x2d, 0x8a, 0x7d, 0x64, 0x7c,
	0x41, 0x3f, 0x7d, 0x1c, 0xf4, 0xec, 0xae, 0x57, 0x97, 0x5f, 0x98, 0xf0, 0xfb, 0xdc, 0xc6, 0xee,
	0xdd, 0x5c, 0xbd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x32, 0x11, 0x33, 0x12, 0xad, 0x02, 0x00, 0x00,
}
