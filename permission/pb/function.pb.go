// Code generated by protoc-gen-go. DO NOT EDIT.
// source: permission/pb/function.proto

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

type FunctionRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Api                  string   `protobuf:"bytes,3,opt,name=api,proto3" json:"api,omitempty"`
	Type                 int64    `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	BindGroupId          string   `protobuf:"bytes,5,opt,name=bindGroupId,proto3" json:"bindGroupId,omitempty"`
	AuthTypes            []int64  `protobuf:"varint,6,rep,packed,name=authTypes,proto3" json:"authTypes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FunctionRequest) Reset()         { *m = FunctionRequest{} }
func (m *FunctionRequest) String() string { return proto.CompactTextString(m) }
func (*FunctionRequest) ProtoMessage()    {}
func (*FunctionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ec5ed375fcb9fb2, []int{0}
}

func (m *FunctionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionRequest.Unmarshal(m, b)
}
func (m *FunctionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionRequest.Marshal(b, m, deterministic)
}
func (m *FunctionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionRequest.Merge(m, src)
}
func (m *FunctionRequest) XXX_Size() int {
	return xxx_messageInfo_FunctionRequest.Size(m)
}
func (m *FunctionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionRequest proto.InternalMessageInfo

func (m *FunctionRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *FunctionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FunctionRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *FunctionRequest) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *FunctionRequest) GetBindGroupId() string {
	if m != nil {
		return m.BindGroupId
	}
	return ""
}

func (m *FunctionRequest) GetAuthTypes() []int64 {
	if m != nil {
		return m.AuthTypes
	}
	return nil
}

type FunctionGroupRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	BindGroupId          string   `protobuf:"bytes,3,opt,name=bindGroupId,proto3" json:"bindGroupId,omitempty"`
	Type                 int64    `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FunctionGroupRequest) Reset()         { *m = FunctionGroupRequest{} }
func (m *FunctionGroupRequest) String() string { return proto.CompactTextString(m) }
func (*FunctionGroupRequest) ProtoMessage()    {}
func (*FunctionGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ec5ed375fcb9fb2, []int{1}
}

func (m *FunctionGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionGroupRequest.Unmarshal(m, b)
}
func (m *FunctionGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionGroupRequest.Marshal(b, m, deterministic)
}
func (m *FunctionGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionGroupRequest.Merge(m, src)
}
func (m *FunctionGroupRequest) XXX_Size() int {
	return xxx_messageInfo_FunctionGroupRequest.Size(m)
}
func (m *FunctionGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionGroupRequest proto.InternalMessageInfo

func (m *FunctionGroupRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *FunctionGroupRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FunctionGroupRequest) GetBindGroupId() string {
	if m != nil {
		return m.BindGroupId
	}
	return ""
}

func (m *FunctionGroupRequest) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func init() {
	proto.RegisterType((*FunctionRequest)(nil), "gs.service.permission.FunctionRequest")
	proto.RegisterType((*FunctionGroupRequest)(nil), "gs.service.permission.FunctionGroupRequest")
}

func init() { proto.RegisterFile("permission/pb/function.proto", fileDescriptor_7ec5ed375fcb9fb2) }

var fileDescriptor_7ec5ed375fcb9fb2 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xd1, 0x4a, 0xfb, 0x30,
	0x18, 0xc5, 0xe9, 0x3f, 0xdb, 0xd8, 0xbe, 0x5d, 0xfc, 0x25, 0x4c, 0x29, 0x63, 0x17, 0x65, 0x17,
	0x32, 0x18, 0xa4, 0xa0, 0xcf, 0xe0, 0xd4, 0x2b, 0xa5, 0xfa, 0x02, 0xd9, 0x12, 0x67, 0x28, 0xc9,
	0x17, 0x9b, 0xb4, 0xb0, 0xc7, 0xf1, 0xb5, 0x7c, 0x1a, 0x69, 0x6a, 0xd9, 0x98, 0x13, 0x91, 0xed,
	0xee, 0xf4, 0x70, 0xbe, 0x5f, 0x0f, 0x87, 0xc0, 0xc4, 0xca, 0x42, 0x2b, 0xe7, 0x14, 0x9a, 0xd4,
	0x2e, 0xd3, 0x97, 0xd2, 0xac, 0xbc, 0x42, 0xc3, 0x6c, 0x81, 0x1e, 0xe9, 0xf9, 0xda, 0x31, 0x27,
	0x8b, 0x4a, 0xad, 0x24, 0xdb, 0x06, 0xc7, 0xf3, 0x1c, 0x8d, 0xcc, 0x73, 0x64, 0x5a, 0xa6, 0x6b,
	0x0c, 0xb7, 0x2b, 0xd4, 0x1a, 0x8d, 0x4b, 0x85, 0xc7, 0x56, 0x37, 0x8c, 0xe9, 0x7b, 0x04, 0xff,
	0x17, 0x5f, 0xd8, 0x4c, 0xbe, 0x95, 0xd2, 0x79, 0x3a, 0x82, 0x2e, 0xb7, 0xf6, 0x5e, 0xc4, 0x51,
	0x12, 0xcd, 0x06, 0x59, 0xf3, 0x41, 0x29, 0x74, 0x0c, 0xd7, 0x32, 0xfe, 0x17, 0xcc, 0xa0, 0xe9,
	0x19, 0x10, 0x6e, 0x55, 0x4c, 0x82, 0x55, 0xcb, 0x3a, 0xe5, 0x37, 0x56, 0xc6, 0x9d, 0x24, 0x9a,
	0x91, 0x2c, 0x68, 0x9a, 0xc0, 0x70, 0xa9, 0x8c, 0xb8, 0x2d, 0xb0, 0xac, 0xa9, 0xdd, 0x90, 0xde,
	0xb5, 0xe8, 0x04, 0x06, 0xbc, 0xf4, 0xaf, 0xcf, 0x1b, 0x2b, 0x5d, 0xdc, 0x4b, 0xc8, 0x8c, 0x64,
	0x5b, 0x63, 0x5a, 0xc1, 0xa8, 0xad, 0x18, 0x0e, 0xfe, 0xde, 0x73, 0xaf, 0x01, 0xf9, 0xde, 0xe0,
	0x40, 0xef, 0xab, 0x0f, 0x02, 0xfd, 0xf6, 0xc7, 0xf4, 0x06, 0x08, 0x17, 0x82, 0x5e, 0xb2, 0x83,
	0xa3, 0xb3, 0xbd, 0x0d, 0xc7, 0x17, 0x75, 0xae, 0x9d, 0x5a, 0x78, 0x64, 0x4f, 0x9e, 0xfb, 0xd2,
	0xd1, 0x3b, 0xe8, 0x15, 0x32, 0x74, 0x3a, 0x96, 0xb4, 0x80, 0x8e, 0xc6, 0xea, 0x78, 0xce, 0x03,
	0xf4, 0xb9, 0x68, 0x76, 0xa0, 0xf3, 0x5f, 0x58, 0xbb, 0xf3, 0xff, 0x08, 0x7c, 0x84, 0x41, 0x5d,
	0xec, 0x84, 0xc4, 0x0c, 0x86, 0xcd, 0x68, 0xa7, 0x63, 0x2e, 0x7b, 0xe1, 0xfd, 0x5f, 0x7f, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xfa, 0x16, 0x43, 0xcc, 0x63, 0x03, 0x00, 0x00,
}
