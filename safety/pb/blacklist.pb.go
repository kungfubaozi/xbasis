// Code generated by protoc-gen-go. DO NOT EDIT.
// source: safety/pb/blacklist.proto

package gs_service_safety

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

type CheckRequest struct {
	Type                 int64    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6f2b3a90f91407b, []int{0}
}

func (m *CheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRequest.Unmarshal(m, b)
}
func (m *CheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRequest.Marshal(b, m, deterministic)
}
func (m *CheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRequest.Merge(m, src)
}
func (m *CheckRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRequest.Size(m)
}
func (m *CheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRequest proto.InternalMessageInfo

func (m *CheckRequest) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *CheckRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type RemoveRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6f2b3a90f91407b, []int{1}
}

func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AddRequest struct {
	Type                 int64    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6f2b3a90f91407b, []int{2}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *AddRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckRequest)(nil), "gs.service.safety.CheckRequest")
	proto.RegisterType((*RemoveRequest)(nil), "gs.service.safety.RemoveRequest")
	proto.RegisterType((*AddRequest)(nil), "gs.service.safety.AddRequest")
}

func init() { proto.RegisterFile("safety/pb/blacklist.proto", fileDescriptor_e6f2b3a90f91407b) }

var fileDescriptor_e6f2b3a90f91407b = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x69, 0xa7, 0x93, 0x1e, 0x54, 0x30, 0x17, 0x52, 0x07, 0xb2, 0xd2, 0xab, 0x81, 0x90,
	0x82, 0xde, 0xc9, 0x2e, 0xdc, 0x7c, 0x83, 0xf8, 0x04, 0x6d, 0x72, 0x9c, 0x21, 0x6b, 0x4e, 0x5d,
	0xce, 0x06, 0x7b, 0x4a, 0x5f, 0x49, 0x6c, 0x09, 0x2a, 0x6e, 0x37, 0xde, 0x9d, 0x90, 0xfc, 0x5f,
	0xce, 0xf7, 0xc3, 0x4d, 0xa8, 0x5f, 0x91, 0xf7, 0x55, 0xd7, 0x54, 0xcd, 0xba, 0xd6, 0x6e, 0x6d,
	0x03, 0xcb, 0x6e, 0x43, 0x4c, 0xe2, 0x6a, 0x15, 0x64, 0xc0, 0xcd, 0xce, 0x6a, 0x94, 0xc3, 0xab,
	0xc9, 0x9d, 0x23, 0x8f, 0xce, 0x91, 0x6c, 0xb1, 0x5a, 0x51, 0xb0, 0xe4, 0x2b, 0x4d, 0x6d, 0x4b,
	0x3e, 0x54, 0x86, 0x29, 0xce, 0x43, 0xbe, 0x9c, 0xc3, 0xf9, 0xf3, 0x1b, 0x6a, 0xa7, 0xf0, 0x7d,
	0x8b, 0x81, 0x85, 0x80, 0x13, 0xde, 0x77, 0x98, 0x27, 0x45, 0x32, 0x1b, 0xa9, 0x7e, 0x16, 0x39,
	0x9c, 0x69, 0xf2, 0x8c, 0x9e, 0xf3, 0xb4, 0x48, 0x66, 0x99, 0x8a, 0xc7, 0x72, 0x0a, 0x17, 0x0a,
	0x5b, 0xda, 0x61, 0x8c, 0x5f, 0x42, 0x6a, 0x4d, 0x1f, 0xce, 0x54, 0x6a, 0x4d, 0xf9, 0x08, 0xb0,
	0x30, 0xe6, 0x5f, 0xf0, 0xfb, 0x8f, 0x04, 0xb2, 0x65, 0xd4, 0x15, 0x73, 0x18, 0x2d, 0x8c, 0x11,
	0xb7, 0xf2, 0x8f, 0xb0, 0xfc, 0xfe, 0x61, 0x72, 0xfd, 0x75, 0x1d, 0x0d, 0x0d, 0x93, 0x7c, 0xe1,
	0x9a, 0xb7, 0x41, 0x2c, 0x61, 0x3c, 0x2c, 0x2a, 0x8a, 0x03, 0x80, 0x5f, 0x0e, 0x47, 0x19, 0x4f,
	0x70, 0xda, 0x57, 0x25, 0xa6, 0x07, 0x10, 0x3f, 0x4b, 0x3c, 0x46, 0x68, 0xc6, 0x7d, 0xe7, 0x0f,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x75, 0x49, 0x07, 0x02, 0xd0, 0x01, 0x00, 0x00,
}
