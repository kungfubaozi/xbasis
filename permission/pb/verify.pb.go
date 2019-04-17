// Code generated by protoc-gen-go. DO NOT EDIT.
// source: permission/pb/verify.proto

package gs_service_permission

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "konekko.me/gosion/commons/dto"
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

type AuthRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_18fc2127d2a89a53, []int{0}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AuthRequest)(nil), "gs.service.permission.AuthRequest")
}

func init() { proto.RegisterFile("permission/pb/verify.proto", fileDescriptor_18fc2127d2a89a53) }

var fileDescriptor_18fc2127d2a89a53 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x48, 0x2d, 0xca,
	0xcd, 0x2c, 0x2e, 0xce, 0xcc, 0xcf, 0xd3, 0x2f, 0x48, 0xd2, 0x2f, 0x4b, 0x2d, 0xca, 0x4c, 0xab,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4d, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0x43, 0x28, 0x93, 0xd2, 0xce, 0xce, 0xcf, 0x4b, 0xcd, 0xce, 0xce, 0xd7, 0xcb,
	0x4d, 0xd5, 0x4f, 0xcf, 0x07, 0xeb, 0x4c, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x2b, 0xd6, 0x4f, 0x29,
	0xc9, 0x87, 0xb1, 0x21, 0x66, 0x28, 0xf1, 0x72, 0x71, 0x3b, 0x96, 0x96, 0x64, 0x04, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0x18, 0x79, 0x73, 0xb1, 0x3a, 0x67, 0xa4, 0x26, 0x67, 0x0b, 0x39, 0x71,
	0xb1, 0x80, 0xc4, 0x85, 0x94, 0xf4, 0xb0, 0x5a, 0xa2, 0x87, 0xa4, 0x49, 0x4a, 0x0c, 0xa4, 0x06,
	0x66, 0x6c, 0x4a, 0x49, 0xbe, 0x5e, 0x70, 0x49, 0x62, 0x49, 0x69, 0x71, 0x12, 0x1b, 0xd8, 0x0a,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x67, 0x00, 0x2a, 0xc4, 0x00, 0x00, 0x00,
}
