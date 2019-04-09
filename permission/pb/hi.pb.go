// Code generated by protoc-gen-go. DO NOT EDIT.
// source: permission/pb/hi.proto

package gs_service_permission

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type HiRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HiRequest) Reset()         { *m = HiRequest{} }
func (m *HiRequest) String() string { return proto.CompactTextString(m) }
func (*HiRequest) ProtoMessage()    {}
func (*HiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b15dd198008fde79, []int{0}
}

func (m *HiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiRequest.Unmarshal(m, b)
}
func (m *HiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiRequest.Marshal(b, m, deterministic)
}
func (m *HiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiRequest.Merge(m, src)
}
func (m *HiRequest) XXX_Size() int {
	return xxx_messageInfo_HiRequest.Size(m)
}
func (m *HiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HiRequest proto.InternalMessageInfo

func (m *HiRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HiResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HiResponse) Reset()         { *m = HiResponse{} }
func (m *HiResponse) String() string { return proto.CompactTextString(m) }
func (*HiResponse) ProtoMessage()    {}
func (*HiResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b15dd198008fde79, []int{1}
}

func (m *HiResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiResponse.Unmarshal(m, b)
}
func (m *HiResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiResponse.Marshal(b, m, deterministic)
}
func (m *HiResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiResponse.Merge(m, src)
}
func (m *HiResponse) XXX_Size() int {
	return xxx_messageInfo_HiResponse.Size(m)
}
func (m *HiResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HiResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HiResponse proto.InternalMessageInfo

func (m *HiResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*HiRequest)(nil), "gs.service.permission.HiRequest")
	proto.RegisterType((*HiResponse)(nil), "gs.service.permission.HiResponse")
}

func init() { proto.RegisterFile("permission/pb/hi.proto", fileDescriptor_b15dd198008fde79) }

var fileDescriptor_b15dd198008fde79 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x48, 0x2d, 0xca,
	0xcd, 0x2c, 0x2e, 0xce, 0xcc, 0xcf, 0xd3, 0x2f, 0x48, 0xd2, 0xcf, 0xc8, 0xd4, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x4d, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x43,
	0x28, 0x51, 0x92, 0xe7, 0xe2, 0xf4, 0xc8, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12,
	0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95,
	0xe4, 0xb8, 0xb8, 0x40, 0x0a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x04, 0xb8, 0x98, 0x73,
	0x8b, 0xd3, 0xa1, 0x0a, 0x40, 0x4c, 0xa3, 0x50, 0x2e, 0x26, 0x8f, 0x4c, 0x21, 0x7f, 0x2e, 0x8e,
	0xe0, 0xc4, 0x4a, 0x8f, 0xd4, 0x9c, 0x9c, 0x7c, 0x21, 0x05, 0x3d, 0xac, 0x56, 0xe9, 0xc1, 0xed,
	0x91, 0x52, 0xc4, 0xa3, 0x02, 0x62, 0x51, 0x12, 0x1b, 0xd8, 0xd5, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xef, 0x8c, 0xb8, 0xf8, 0xcf, 0x00, 0x00, 0x00,
}
