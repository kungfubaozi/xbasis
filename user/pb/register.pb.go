// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/pb/register.proto

package xbasissvc_external_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "konekko.me/xbasis/commons/dto"
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

type NewRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Contract             string   `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	ClientId             string   `protobuf:"bytes,3,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewRequest) Reset()         { *m = NewRequest{} }
func (m *NewRequest) String() string { return proto.CompactTextString(m) }
func (*NewRequest) ProtoMessage()    {}
func (*NewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dde35f49c07dd310, []int{0}
}

func (m *NewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewRequest.Unmarshal(m, b)
}
func (m *NewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewRequest.Marshal(b, m, deterministic)
}
func (m *NewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewRequest.Merge(m, src)
}
func (m *NewRequest) XXX_Size() int {
	return xxx_messageInfo_NewRequest.Size(m)
}
func (m *NewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewRequest proto.InternalMessageInfo

func (m *NewRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *NewRequest) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *NewRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *NewRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*NewRequest)(nil), "xbasissvc.external.user.NewRequest")
}

func init() { proto.RegisterFile("user/pb/register.proto", fileDescriptor_dde35f49c07dd310) }

var fileDescriptor_dde35f49c07dd310 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x59, 0x57, 0x64, 0xcd, 0xb1, 0x07, 0x2d, 0x3d, 0x89, 0x5e, 0x04, 0x61, 0x02, 0xfa,
	0x0e, 0x82, 0x97, 0x3d, 0xd4, 0x27, 0x48, 0x93, 0x41, 0x4a, 0x37, 0x99, 0x9a, 0x99, 0xda, 0x3d,
	0xf8, 0xf0, 0x92, 0x26, 0xd1, 0x93, 0xb7, 0x7c, 0x7c, 0x3f, 0x43, 0x3e, 0x75, 0xb3, 0x30, 0x46,
	0x3d, 0x0f, 0x3a, 0xe2, 0xc7, 0xc8, 0x82, 0x11, 0xe6, 0x48, 0x42, 0xcd, 0xed, 0x79, 0x30, 0x3c,
	0x32, 0x7f, 0x59, 0xc0, 0xb3, 0x60, 0x0c, 0xe6, 0x04, 0x69, 0xda, 0x3d, 0x4d, 0x14, 0x70, 0x9a,
	0x08, 0x3c, 0xea, 0xbc, 0xd1, 0x96, 0xbc, 0xa7, 0xc0, 0xda, 0x09, 0xd5, 0x77, 0xbe, 0x72, 0xff,
	0xad, 0xd4, 0x11, 0xd7, 0x1e, 0x3f, 0x17, 0x64, 0x69, 0x3a, 0x75, 0x48, 0x27, 0x82, 0xf1, 0xd8,
	0xee, 0xee, 0x76, 0x8f, 0xd7, 0xfd, 0x2f, 0x27, 0x67, 0x29, 0x48, 0x34, 0x56, 0xda, 0x8b, 0xec,
	0x2a, 0x6f, 0xee, 0x34, 0x62, 0x90, 0x37, 0xd7, 0xee, 0x8b, 0x2b, 0x9c, 0xdc, 0x6c, 0x98, 0x57,
	0x8a, 0xae, 0xbd, 0xcc, 0xae, 0xf2, 0x73, 0xaf, 0x0e, 0x7d, 0xa9, 0x6a, 0x5e, 0xd5, 0xfe, 0x88,
	0x6b, 0xf3, 0x00, 0xff, 0x74, 0xc1, 0xdf, 0x3f, 0xbb, 0xae, 0x8c, 0xa0, 0xc6, 0x38, 0x21, 0x78,
	0x17, 0x23, 0x0b, 0x0f, 0x57, 0x5b, 0xd8, 0xcb, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x2d,
	0x71, 0xc3, 0x38, 0x01, 0x00, 0x00,
}
