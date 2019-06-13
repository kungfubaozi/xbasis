// Code generated by protoc-gen-go. DO NOT EDIT.
// source: permission/pb/binding.proto

package gosionsvc_external_permission

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

type BindingRoleRequest struct {
	StructureId          string   `protobuf:"bytes,1,opt,name=structureId,proto3" json:"structureId,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	RoleId               string   `protobuf:"bytes,3,opt,name=roleId,proto3" json:"roleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindingRoleRequest) Reset()         { *m = BindingRoleRequest{} }
func (m *BindingRoleRequest) String() string { return proto.CompactTextString(m) }
func (*BindingRoleRequest) ProtoMessage()    {}
func (*BindingRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d6cdac82fda43c5, []int{0}
}

func (m *BindingRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindingRoleRequest.Unmarshal(m, b)
}
func (m *BindingRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindingRoleRequest.Marshal(b, m, deterministic)
}
func (m *BindingRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindingRoleRequest.Merge(m, src)
}
func (m *BindingRoleRequest) XXX_Size() int {
	return xxx_messageInfo_BindingRoleRequest.Size(m)
}
func (m *BindingRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BindingRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BindingRoleRequest proto.InternalMessageInfo

func (m *BindingRoleRequest) GetStructureId() string {
	if m != nil {
		return m.StructureId
	}
	return ""
}

func (m *BindingRoleRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BindingRoleRequest) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func init() {
	proto.RegisterType((*BindingRoleRequest)(nil), "gosionsvc.external.permission.BindingRoleRequest")
}

func init() { proto.RegisterFile("permission/pb/binding.proto", fileDescriptor_5d6cdac82fda43c5) }

var fileDescriptor_5d6cdac82fda43c5 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x69, 0x84, 0xaa, 0xa3, 0xf4, 0xb0, 0x87, 0x52, 0x2a, 0x42, 0xf1, 0x24, 0x08, 0xbb,
	0xa8, 0x6f, 0xe0, 0x41, 0xf0, 0x5a, 0xc9, 0x41, 0x04, 0xa5, 0xc9, 0x0e, 0x61, 0x49, 0x32, 0x13,
	0x77, 0x66, 0xc5, 0x57, 0xf7, 0x26, 0x49, 0x5a, 0x14, 0x04, 0x4f, 0xf6, 0xb6, 0x3b, 0xf3, 0xf3,
	0xcd, 0xb7, 0x3b, 0x70, 0xd6, 0x61, 0x6c, 0x83, 0x48, 0x60, 0x72, 0x5d, 0xe1, 0x8a, 0x40, 0x3e,
	0x50, 0x65, 0xbb, 0xc8, 0xca, 0xe6, 0xbc, 0xe2, 0xbe, 0x21, 0xef, 0xa5, 0xc5, 0x0f, 0xc5, 0x48,
	0x9b, 0xc6, 0x7e, 0xe7, 0x97, 0x57, 0x35, 0x13, 0xd6, 0x35, 0xdb, 0x16, 0xdd, 0x98, 0x74, 0x25,
	0xb7, 0x2d, 0x93, 0x38, 0xaf, 0xbc, 0x3b, 0x8f, 0xac, 0x8b, 0x17, 0x30, 0x77, 0x23, 0x7c, 0xcd,
	0x0d, 0xae, 0xf1, 0x2d, 0xa1, 0xa8, 0x59, 0xc1, 0x89, 0x68, 0x4c, 0xa5, 0xa6, 0x88, 0x0f, 0x7e,
	0x31, 0x59, 0x4d, 0x2e, 0x8f, 0xd7, 0x3f, 0x4b, 0x66, 0x06, 0x59, 0xf0, 0x8b, 0x6c, 0x68, 0x64,
	0xc1, 0x9b, 0x39, 0x4c, 0x23, 0x37, 0x7d, 0xf8, 0x60, 0xa8, 0x6d, 0x6f, 0x37, 0x9f, 0x19, 0x1c,
	0x6e, 0x07, 0x98, 0x1c, 0x8e, 0x72, 0xc1, 0xd8, 0x0f, 0x32, 0xd7, 0xf6, 0xcf, 0x47, 0xd8, 0xdf,
	0x52, 0xcb, 0xb9, 0xad, 0xc4, 0xee, 0xec, 0xbd, 0xb2, 0x7d, 0xd4, 0x8d, 0x26, 0x31, 0x4f, 0x70,
	0x7a, 0x9f, 0xa8, 0xd4, 0xc0, 0xf4, 0xdf, 0xe8, 0x67, 0x98, 0xe5, 0xd4, 0x7f, 0xfe, 0x3e, 0xbc,
	0x5f, 0xc1, 0x8c, 0xf0, 0x3d, 0xd9, 0x17, 0xd3, 0x61, 0xc5, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x92, 0xa5, 0x94, 0x5c, 0x4d, 0x02, 0x00, 0x00,
}
