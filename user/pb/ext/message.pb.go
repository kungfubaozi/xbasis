// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/pb/ext/message.proto

package gs_ext_service_user

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

type SendRequest struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Type                 int64    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	MessageType          int64    `protobuf:"varint,4,opt,name=messageType,proto3" json:"messageType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b6132722858f0ef, []int{0}
}

func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *SendRequest) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *SendRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SendRequest) GetMessageType() int64 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

func init() {
	proto.RegisterType((*SendRequest)(nil), "gs.ext.service.user.SendRequest")
}

func init() { proto.RegisterFile("user/pb/ext/message.proto", fileDescriptor_0b6132722858f0ef) }

var fileDescriptor_0b6132722858f0ef = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xd9, 0x6d, 0x51, 0x4c, 0xc1, 0x43, 0x14, 0x59, 0x7b, 0x5a, 0x3c, 0x15, 0x84, 0x09,
	0xe8, 0x4f, 0xf0, 0x2c, 0xc8, 0x56, 0xbc, 0x78, 0xda, 0x6e, 0xc6, 0x10, 0x96, 0x64, 0xd6, 0xcc,
	0xac, 0xd4, 0x7f, 0x2f, 0x49, 0x2d, 0xf4, 0xd0, 0xdb, 0x9b, 0xc7, 0xc7, 0x9b, 0xf7, 0xd4, 0xfd,
	0xcc, 0x98, 0xcc, 0xb4, 0x33, 0xb8, 0x17, 0x13, 0x90, 0xb9, 0x77, 0x08, 0x53, 0x22, 0x21, 0x7d,
	0xe3, 0x18, 0x70, 0x2f, 0xc0, 0x98, 0x7e, 0xfc, 0x80, 0x90, 0xc9, 0xf5, 0xe3, 0x48, 0x11, 0xc7,
	0x91, 0x20, 0xa0, 0x71, 0xc4, 0x9e, 0xa2, 0x19, 0x28, 0x04, 0x8a, 0x6c, 0xac, 0xd0, 0x51, 0x1f,
	0x12, 0x1e, 0x9c, 0x5a, 0x6d, 0x31, 0xda, 0x0e, 0xbf, 0x67, 0x64, 0xd1, 0xd7, 0xaa, 0x16, 0x6a,
	0xaa, 0xb6, 0xda, 0x5c, 0x75, 0xb5, 0x90, 0xd6, 0x6a, 0x29, 0xbf, 0x13, 0x36, 0x75, 0x5b, 0x6d,
	0x16, 0x5d, 0xd1, 0xd9, 0x1b, 0xc8, 0x62, 0xb3, 0x28, 0x54, 0xd1, 0xba, 0x55, 0xab, 0xff, 0x66,
	0xef, 0x19, 0x5f, 0x16, 0xfc, 0xd4, 0x7a, 0xfa, 0x54, 0x97, 0xaf, 0x87, 0x53, 0xbf, 0xa9, 0xdb,
	0xfc, 0xf3, 0x03, 0x93, 0xff, 0xf2, 0x43, 0x2f, 0x9e, 0xe2, 0x4b, 0x09, 0x81, 0x33, 0x73, 0xe0,
	0xa4, 0xde, 0xfa, 0x2e, 0x13, 0xc7, 0x01, 0x56, 0x08, 0xb6, 0xd2, 0xcb, 0xcc, 0xbb, 0x8b, 0x32,
	0xe6, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xb7, 0x73, 0xc9, 0x2b, 0x01, 0x00, 0x00,
}