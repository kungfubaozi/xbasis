// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/pb/grant.proto

package gs_service_user

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

type StatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_875be15474041634, []int{0}
}

func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (m *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(m, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

type GrantRequest struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrantRequest) Reset()         { *m = GrantRequest{} }
func (m *GrantRequest) String() string { return proto.CompactTextString(m) }
func (*GrantRequest) ProtoMessage()    {}
func (*GrantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_875be15474041634, []int{1}
}

func (m *GrantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrantRequest.Unmarshal(m, b)
}
func (m *GrantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrantRequest.Marshal(b, m, deterministic)
}
func (m *GrantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrantRequest.Merge(m, src)
}
func (m *GrantRequest) XXX_Size() int {
	return xxx_messageInfo_GrantRequest.Size(m)
}
func (m *GrantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GrantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GrantRequest proto.InternalMessageInfo

func (m *GrantRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func init() {
	proto.RegisterType((*StatusRequest)(nil), "gs.service.user.StatusRequest")
	proto.RegisterType((*GrantRequest)(nil), "gs.service.user.GrantRequest")
}

func init() { proto.RegisterFile("user/pb/grant.proto", fileDescriptor_875be15474041634) }

var fileDescriptor_875be15474041634 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2f, 0x48, 0xd2, 0x4f, 0x2f, 0x4a, 0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x4f, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x03, 0xc9, 0x4b, 0x69,
	0x67, 0xe7, 0xe7, 0xa5, 0x66, 0x67, 0xe7, 0xeb, 0xe5, 0xa6, 0xea, 0xa7, 0xe7, 0x17, 0x67, 0xe6,
	0xe7, 0xe9, 0x27, 0xe7, 0xe7, 0xe6, 0xe6, 0xe7, 0x15, 0xeb, 0xa7, 0x94, 0xe4, 0xc3, 0xd8, 0x10,
	0xdd, 0x4a, 0xfc, 0x5c, 0xbc, 0xc1, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x4a, 0x5a, 0x5c, 0x3c, 0xee, 0x20, 0xd3, 0xa1, 0x7c, 0x21, 0x29, 0x2e, 0x8e, 0xe4,
	0x9c, 0xcc, 0xd4, 0xbc, 0x12, 0xcf, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x38, 0xdf,
	0x68, 0x0a, 0x23, 0x97, 0x80, 0x63, 0x41, 0x41, 0x4e, 0x66, 0x72, 0x62, 0x49, 0x66, 0x7e, 0x1e,
	0x58, 0x9f, 0x90, 0x1d, 0x17, 0x2b, 0xd8, 0x79, 0x42, 0xb2, 0x7a, 0x68, 0x2e, 0xd3, 0x43, 0x36,
	0x58, 0x4a, 0x0c, 0x24, 0x0d, 0x73, 0x4c, 0x4a, 0x49, 0xbe, 0x1e, 0xc4, 0x1d, 0x42, 0x0e, 0x5c,
	0x6c, 0xc5, 0x10, 0x96, 0x1c, 0x86, 0x01, 0x28, 0x4e, 0xc5, 0x65, 0x42, 0x12, 0x1b, 0xd8, 0x6b,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xaa, 0x6a, 0x79, 0x2f, 0x01, 0x00, 0x00,
}
