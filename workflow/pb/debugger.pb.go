// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/debugger.proto

package xbasissvc_external_workflow

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

type NextRequest struct {
	InstanceId           string   `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NextRequest) Reset()         { *m = NextRequest{} }
func (m *NextRequest) String() string { return proto.CompactTextString(m) }
func (*NextRequest) ProtoMessage()    {}
func (*NextRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d226d22c86bd491, []int{0}
}

func (m *NextRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NextRequest.Unmarshal(m, b)
}
func (m *NextRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NextRequest.Marshal(b, m, deterministic)
}
func (m *NextRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NextRequest.Merge(m, src)
}
func (m *NextRequest) XXX_Size() int {
	return xxx_messageInfo_NextRequest.Size(m)
}
func (m *NextRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NextRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NextRequest proto.InternalMessageInfo

func (m *NextRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

type NextResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NextResponse) Reset()         { *m = NextResponse{} }
func (m *NextResponse) String() string { return proto.CompactTextString(m) }
func (*NextResponse) ProtoMessage()    {}
func (*NextResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d226d22c86bd491, []int{1}
}

func (m *NextResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NextResponse.Unmarshal(m, b)
}
func (m *NextResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NextResponse.Marshal(b, m, deterministic)
}
func (m *NextResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NextResponse.Merge(m, src)
}
func (m *NextResponse) XXX_Size() int {
	return xxx_messageInfo_NextResponse.Size(m)
}
func (m *NextResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NextResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NextResponse proto.InternalMessageInfo

type RunRequest struct {
	Json                 string   `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunRequest) Reset()         { *m = RunRequest{} }
func (m *RunRequest) String() string { return proto.CompactTextString(m) }
func (*RunRequest) ProtoMessage()    {}
func (*RunRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d226d22c86bd491, []int{2}
}

func (m *RunRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunRequest.Unmarshal(m, b)
}
func (m *RunRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunRequest.Marshal(b, m, deterministic)
}
func (m *RunRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunRequest.Merge(m, src)
}
func (m *RunRequest) XXX_Size() int {
	return xxx_messageInfo_RunRequest.Size(m)
}
func (m *RunRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunRequest proto.InternalMessageInfo

func (m *RunRequest) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

type RunResponse struct {
	InstanceId           string   `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunResponse) Reset()         { *m = RunResponse{} }
func (m *RunResponse) String() string { return proto.CompactTextString(m) }
func (*RunResponse) ProtoMessage()    {}
func (*RunResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d226d22c86bd491, []int{3}
}

func (m *RunResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunResponse.Unmarshal(m, b)
}
func (m *RunResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunResponse.Marshal(b, m, deterministic)
}
func (m *RunResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunResponse.Merge(m, src)
}
func (m *RunResponse) XXX_Size() int {
	return xxx_messageInfo_RunResponse.Size(m)
}
func (m *RunResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RunResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RunResponse proto.InternalMessageInfo

func (m *RunResponse) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func init() {
	proto.RegisterType((*NextRequest)(nil), "xbasissvc.external.workflow.NextRequest")
	proto.RegisterType((*NextResponse)(nil), "xbasissvc.external.workflow.NextResponse")
	proto.RegisterType((*RunRequest)(nil), "xbasissvc.external.workflow.RunRequest")
	proto.RegisterType((*RunResponse)(nil), "xbasissvc.external.workflow.RunResponse")
}

func init() { proto.RegisterFile("pb/debugger.proto", fileDescriptor_1d226d22c86bd491) }

var fileDescriptor_1d226d22c86bd491 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x48, 0xd2, 0x4f,
	0x49, 0x4d, 0x2a, 0x4d, 0x4f, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xae,
	0x48, 0x4a, 0x2c, 0xce, 0x2c, 0x2e, 0x2e, 0x4b, 0xd6, 0x4b, 0xad, 0x28, 0x49, 0x2d, 0xca, 0x4b,
	0xcc, 0xd1, 0x2b, 0xcf, 0x2f, 0xca, 0x4e, 0xcb, 0xc9, 0x2f, 0x57, 0xd2, 0xe5, 0xe2, 0xf6, 0x4b,
	0xad, 0x28, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe3, 0xe2, 0xca, 0xcc, 0x2b,
	0x2e, 0x49, 0xcc, 0x4b, 0x4e, 0xf5, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x12,
	0x51, 0xe2, 0xe3, 0xe2, 0x81, 0x28, 0x2f, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x55, 0x52, 0xe0, 0xe2,
	0x0a, 0x2a, 0xcd, 0x83, 0xe9, 0x16, 0xe2, 0x62, 0xc9, 0x2a, 0xce, 0xcf, 0x83, 0xea, 0x03, 0xb3,
	0x41, 0x16, 0x80, 0x55, 0x40, 0x34, 0x10, 0xb2, 0xc0, 0xe8, 0x20, 0x23, 0x17, 0x87, 0x0b, 0xd4,
	0xfd, 0x42, 0x11, 0x5c, 0xcc, 0x41, 0xa5, 0x79, 0x42, 0xea, 0x7a, 0x78, 0x7c, 0xa0, 0x87, 0xb0,
	0x5f, 0x4a, 0x83, 0xb0, 0x42, 0xa8, 0x33, 0xa2, 0xb9, 0x58, 0x40, 0xfe, 0x10, 0xc2, 0xaf, 0x03,
	0x29, 0x64, 0xa4, 0x34, 0x89, 0x50, 0x09, 0x31, 0x3c, 0x89, 0x0d, 0x1c, 0xee, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xbc, 0xce, 0x7c, 0x77, 0x8c, 0x01, 0x00, 0x00,
}
