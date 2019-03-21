// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/pb/nops/message.proto

package gs_nops_service_message

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
	return fileDescriptor_998fee6cbe2308bd, []int{0}
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
	proto.RegisterType((*SendRequest)(nil), "gs.nops.service.message.SendRequest")
}

func init() { proto.RegisterFile("user/pb/nops/message.proto", fileDescriptor_998fee6cbe2308bd) }

var fileDescriptor_998fee6cbe2308bd = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xd9, 0x6d, 0x51, 0x4c, 0xc1, 0x43, 0x10, 0x5d, 0xf6, 0xb4, 0x88, 0x87, 0x82, 0x30,
	0x01, 0xfd, 0x09, 0x9e, 0xbd, 0x6c, 0x8b, 0x57, 0xd9, 0xee, 0x8e, 0x21, 0x2c, 0xc9, 0xc4, 0xcc,
	0xac, 0xe0, 0xbf, 0x97, 0xa4, 0x2d, 0xf4, 0xe2, 0xed, 0xe5, 0xe5, 0xe3, 0xbd, 0x79, 0xaa, 0x5d,
	0x18, 0x93, 0x89, 0x07, 0x13, 0x28, 0xb2, 0xf1, 0xc8, 0x3c, 0x58, 0x84, 0x98, 0x48, 0x48, 0x3f,
	0x58, 0x86, 0x6c, 0x03, 0x63, 0xfa, 0x71, 0x23, 0xc2, 0xe9, 0xbb, 0x7d, 0x9e, 0x29, 0xe0, 0x3c,
	0x13, 0x78, 0x34, 0x96, 0xd8, 0x51, 0x30, 0x23, 0x79, 0x4f, 0x81, 0xcd, 0x24, 0x74, 0xd6, 0xc7,
	0x94, 0x47, 0xab, 0x36, 0x3b, 0x0c, 0x53, 0x8f, 0xdf, 0x0b, 0xb2, 0xe8, 0x5b, 0x55, 0x0b, 0x35,
	0x55, 0x57, 0x6d, 0x6f, 0xfa, 0x5a, 0x48, 0x6b, 0xb5, 0x96, 0xdf, 0x88, 0x4d, 0xdd, 0x55, 0xdb,
	0x55, 0x5f, 0x74, 0xf6, 0x46, 0x9a, 0xb0, 0x59, 0x15, 0xaa, 0x68, 0xdd, 0xa9, 0xcd, 0xa9, 0x7e,
	0x9f, 0xf1, 0x75, 0xc1, 0x2f, 0xad, 0x97, 0x4f, 0x75, 0xfd, 0x7e, 0x7c, 0xea, 0xbd, 0xba, 0xcb,
	0x9d, 0x1f, 0x98, 0xdc, 0x97, 0x1b, 0x07, 0x71, 0x14, 0xde, 0x72, 0xc8, 0x13, 0xfc, 0x33, 0x09,
	0x2e, 0x4e, 0x6c, 0xef, 0x33, 0x75, 0x1e, 0x31, 0x09, 0xc1, 0x4e, 0x06, 0x59, 0xf8, 0x70, 0x55,
	0x06, 0xbd, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x04, 0x0e, 0xf4, 0x34, 0x01, 0x00, 0x00,
}
