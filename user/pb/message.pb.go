// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/pb/message.proto

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
	return fileDescriptor_14fa1320b0eb76da, []int{0}
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
	proto.RegisterType((*SendRequest)(nil), "gs.service.user.SendRequest")
}

func init() { proto.RegisterFile("user/pb/message.proto", fileDescriptor_14fa1320b0eb76da) }

var fileDescriptor_14fa1320b0eb76da = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x85, 0xe9, 0x0f, 0x8a, 0x29, 0x28, 0x04, 0x94, 0x52, 0x5c, 0x14, 0x57, 0x05, 0x21, 0x01,
	0x7d, 0x04, 0xd7, 0xba, 0x68, 0x7d, 0x81, 0xb6, 0x19, 0x42, 0x29, 0xc9, 0xd4, 0xce, 0x54, 0xf0,
	0xed, 0x25, 0xe9, 0x2d, 0x94, 0xbb, 0x3b, 0xe7, 0xe4, 0x0b, 0xf9, 0x22, 0x1e, 0x37, 0x82, 0x55,
	0x2f, 0x83, 0x76, 0x40, 0xd4, 0x5b, 0x50, 0xcb, 0x8a, 0x8c, 0xf2, 0xc1, 0x92, 0x22, 0x58, 0x7f,
	0xa7, 0x11, 0x54, 0x20, 0xaa, 0xd7, 0x19, 0x3d, 0xcc, 0x33, 0x2a, 0x07, 0xda, 0x22, 0x4d, 0xe8,
	0xf5, 0x88, 0xce, 0xa1, 0x27, 0x6d, 0x18, 0x8f, 0xbc, 0xdf, 0x7e, 0xb1, 0xa2, 0xe8, 0xc0, 0x9b,
	0x16, 0x7e, 0x36, 0x20, 0x96, 0xf7, 0x22, 0x65, 0x2c, 0x93, 0x3a, 0x69, 0xee, 0xda, 0x94, 0x51,
	0x4a, 0x91, 0xf3, 0xdf, 0x02, 0x65, 0x5a, 0x27, 0x4d, 0xd6, 0xc6, 0x1c, 0xb6, 0x11, 0x0d, 0x94,
	0x59, 0xa4, 0x62, 0x96, 0xb5, 0x28, 0x2e, 0x56, 0xdf, 0x01, 0xcf, 0x23, 0x7e, 0x9e, 0xde, 0xbe,
	0xc4, 0xed, 0xe7, 0x5e, 0xe5, 0xc7, 0xfe, 0xe6, 0x51, 0x9f, 0xd5, 0xd5, 0x0f, 0xd4, 0xc9, 0xa8,
	0x7a, 0x0a, 0xa7, 0x87, 0xb3, 0x61, 0x54, 0x1d, 0xf7, 0xbc, 0xd1, 0x70, 0x13, 0xfd, 0xdf, 0xff,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x63, 0xed, 0x52, 0xc6, 0x16, 0x01, 0x00, 0x00,
}