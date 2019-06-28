// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/pb/update.proto

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

type UpdateRequest struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8d02e1088640a13, []int{0}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type UpdatePasswordRequest struct {
	Original             string   `protobuf:"bytes,1,opt,name=original,proto3" json:"original,omitempty"`
	New                  string   `protobuf:"bytes,2,opt,name=new,proto3" json:"new,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePasswordRequest) Reset()         { *m = UpdatePasswordRequest{} }
func (m *UpdatePasswordRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePasswordRequest) ProtoMessage()    {}
func (*UpdatePasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8d02e1088640a13, []int{1}
}

func (m *UpdatePasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePasswordRequest.Unmarshal(m, b)
}
func (m *UpdatePasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePasswordRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePasswordRequest.Merge(m, src)
}
func (m *UpdatePasswordRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePasswordRequest.Size(m)
}
func (m *UpdatePasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePasswordRequest proto.InternalMessageInfo

func (m *UpdatePasswordRequest) GetOriginal() string {
	if m != nil {
		return m.Original
	}
	return ""
}

func (m *UpdatePasswordRequest) GetNew() string {
	if m != nil {
		return m.New
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateRequest)(nil), "xbasissvc.external.user.UpdateRequest")
	proto.RegisterType((*UpdatePasswordRequest)(nil), "xbasissvc.external.user.UpdatePasswordRequest")
}

func init() { proto.RegisterFile("user/pb/update.proto", fileDescriptor_d8d02e1088640a13) }

var fileDescriptor_d8d02e1088640a13 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xa9, 0xc5, 0x5a, 0x07, 0x04, 0x59, 0x14, 0x43, 0x4e, 0xd2, 0x83, 0x28, 0xc2, 0x06,
	0xf4, 0x19, 0x7a, 0x10, 0xa1, 0x94, 0x48, 0xbd, 0x6f, 0x92, 0x41, 0x43, 0x92, 0x9d, 0xb8, 0x33,
	0xb1, 0x7d, 0x66, 0x9f, 0x42, 0x36, 0xe9, 0x0a, 0x0a, 0xea, 0x25, 0xb7, 0x99, 0xe5, 0x9f, 0x8f,
	0xf9, 0xf7, 0x1f, 0x38, 0xeb, 0x18, 0x5d, 0xd2, 0x66, 0x49, 0xd7, 0x16, 0x46, 0x50, 0xb7, 0x8e,
	0x84, 0xd4, 0xc5, 0x2e, 0x33, 0x5c, 0x32, 0xbf, 0xe7, 0x1a, 0x77, 0x82, 0xce, 0x9a, 0x5a, 0x7b,
	0x61, 0x7c, 0x5b, 0x91, 0xc5, 0xaa, 0x22, 0xdd, 0x60, 0x32, 0x68, 0x92, 0x9c, 0x9a, 0x86, 0x2c,
	0x27, 0x85, 0x50, 0xa8, 0x07, 0xca, 0xe2, 0x06, 0x4e, 0x36, 0x3d, 0x35, 0xc5, 0xb7, 0x0e, 0x59,
	0x54, 0x04, 0x47, 0x39, 0x59, 0x41, 0x2b, 0xd1, 0xe4, 0x72, 0x72, 0x7d, 0x9c, 0x86, 0x76, 0xb1,
	0x84, 0xf3, 0x41, 0xba, 0x36, 0xcc, 0x5b, 0x72, 0x45, 0x18, 0x89, 0x61, 0x4e, 0xae, 0x7c, 0x29,
	0xad, 0xa9, 0xf7, 0x33, 0x5f, 0xbd, 0x3a, 0x85, 0xa9, 0xc5, 0x6d, 0x74, 0xd0, 0x3f, 0xfb, 0xf2,
	0xee, 0x63, 0x0a, 0xb3, 0x81, 0xa3, 0x56, 0x30, 0xdf, 0xb0, 0xdf, 0xbc, 0x41, 0x75, 0xa5, 0x7f,
	0xf1, 0xa3, 0xbf, 0xed, 0x17, 0xc7, 0x7b, 0x9d, 0x0e, 0x3e, 0x0a, 0x21, 0xfd, 0x24, 0x46, 0x3a,
	0xf6, 0xbc, 0x14, 0x4d, 0xbd, 0x1a, 0x8b, 0xf7, 0x08, 0x87, 0xeb, 0x57, 0xb2, 0xa3, 0xc1, 0x96,
	0x8d, 0x29, 0xeb, 0xb1, 0x9c, 0xfa, 0x9f, 0x7b, 0xc8, 0xc9, 0x8e, 0xc2, 0x7b, 0x86, 0x79, 0x48,
	0x55, 0xe9, 0x7f, 0x78, 0x3f, 0xe2, 0xff, 0x8b, 0x9b, 0xcd, 0xfa, 0x2b, 0xbb, 0xff, 0x0c, 0x00,
	0x00, 0xff, 0xff, 0xca, 0x21, 0xd8, 0x57, 0xc3, 0x02, 0x00, 0x00,
}
