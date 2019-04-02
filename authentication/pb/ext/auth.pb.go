// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication/pb/ext/auth.proto

package gs_ext_service_authentication

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

type VerifyRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRequest) Reset()         { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()    {}
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e1e612aa43fed8d, []int{0}
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

func init() {
	proto.RegisterType((*VerifyRequest)(nil), "gs.ext.service.authentication.VerifyRequest")
}

func init() { proto.RegisterFile("authentication/pb/ext/auth.proto", fileDescriptor_2e1e612aa43fed8d) }

var fileDescriptor_2e1e612aa43fed8d = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8e, 0xcf, 0x6a, 0x84, 0x30,
	0x10, 0x87, 0xb1, 0xb4, 0xd2, 0x06, 0x7a, 0x09, 0xa5, 0x88, 0x50, 0x90, 0x9e, 0x0a, 0x2d, 0x13,
	0x68, 0x9f, 0xc0, 0x63, 0x2f, 0x3d, 0x58, 0xf0, 0xae, 0x71, 0xaa, 0x21, 0x35, 0xe3, 0x9a, 0xc9,
	0xe2, 0xbe, 0xfd, 0xa2, 0xe2, 0x82, 0x97, 0xbd, 0xcd, 0x6f, 0xfe, 0x7c, 0xf3, 0x89, 0xac, 0x0a,
	0xdc, 0xa1, 0x63, 0xa3, 0x2b, 0x36, 0xe4, 0xd4, 0x50, 0x2b, 0x9c, 0x58, 0xcd, 0x5d, 0x18, 0x46,
	0x62, 0x92, 0x2f, 0xad, 0x07, 0x9c, 0x18, 0x3c, 0x8e, 0x47, 0xa3, 0x11, 0xf6, 0x07, 0xe9, 0xbb,
	0x25, 0x87, 0xd6, 0x12, 0xf4, 0xa8, 0x5a, 0xf2, 0x33, 0x43, 0x53, 0xdf, 0x93, 0xf3, 0xaa, 0x61,
	0xda, 0xea, 0x95, 0xf5, 0x9a, 0x8b, 0xc7, 0x12, 0x47, 0xf3, 0x77, 0x2a, 0xf0, 0x10, 0xd0, 0xb3,
	0x7c, 0x12, 0x77, 0x4c, 0x16, 0x5d, 0x12, 0x65, 0xd1, 0xdb, 0x43, 0xb1, 0x06, 0x99, 0x8a, 0x7b,
	0xfd, 0x6f, 0xd0, 0xf1, 0x77, 0x93, 0xdc, 0x2c, 0x83, 0x4b, 0xfe, 0x2c, 0xc5, 0x6d, 0x1e, 0xb8,
	0x93, 0x3f, 0x22, 0x5e, 0x51, 0xf2, 0x03, 0xae, 0x1a, 0xc2, 0xee, 0x63, 0xfa, 0x3c, 0x6f, 0x6f,
	0x56, 0x0d, 0x13, 0xfc, 0x72, 0xc5, 0xc1, 0xd7, 0xf1, 0x62, 0xf8, 0x75, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0xdd, 0xfa, 0x1b, 0x62, 0x11, 0x01, 0x00, 0x00,
}