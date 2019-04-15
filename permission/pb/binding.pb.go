// Code generated by protoc-gen-go. DO NOT EDIT.
// source: permission/pb/binding.proto

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

type BindingRoleRequest struct {
	StructureId          string   `protobuf:"bytes,1,opt,name=structureId,proto3" json:"structureId,omitempty"`
	Id                   []string `protobuf:"bytes,2,rep,name=id,proto3" json:"id,omitempty"`
	RoleId               []string `protobuf:"bytes,3,rep,name=roleId,proto3" json:"roleId,omitempty"`
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

func (m *BindingRoleRequest) GetId() []string {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *BindingRoleRequest) GetRoleId() []string {
	if m != nil {
		return m.RoleId
	}
	return nil
}

func init() {
	proto.RegisterType((*BindingRoleRequest)(nil), "gs.service.permission.BindingRoleRequest")
}

func init() { proto.RegisterFile("permission/pb/binding.proto", fileDescriptor_5d6cdac82fda43c5) }

var fileDescriptor_5d6cdac82fda43c5 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x69, 0x0a, 0xd5, 0x8e, 0xd2, 0xc3, 0x80, 0xa5, 0xd4, 0x4b, 0xf0, 0x54, 0x11, 0x36,
	0xa0, 0x6f, 0xe0, 0x41, 0xf0, 0xe0, 0xc1, 0x94, 0xe2, 0x4d, 0x30, 0xd9, 0x21, 0x2c, 0x69, 0x76,
	0xe2, 0xce, 0xac, 0xef, 0xea, 0xdb, 0x48, 0x92, 0x16, 0x05, 0x7b, 0xcc, 0x6d, 0x77, 0xf8, 0xf7,
	0xfb, 0x3f, 0x76, 0xe0, 0xba, 0xa5, 0xd0, 0x38, 0x11, 0xc7, 0x3e, 0x6b, 0x8b, 0xac, 0x70, 0xde,
	0x3a, 0x5f, 0x99, 0x36, 0xb0, 0x32, 0x5e, 0x55, 0x62, 0x84, 0xc2, 0x97, 0x2b, 0xc9, 0xfc, 0xe6,
	0xd6, 0x77, 0x35, 0x7b, 0xaa, 0x6b, 0x36, 0x0d, 0x65, 0x15, 0xf7, 0x4f, 0x4b, 0x6e, 0x1a, 0xf6,
	0x92, 0x59, 0xe5, 0xe3, 0x79, 0x60, 0xdc, 0xbc, 0x03, 0x3e, 0x0e, 0xd0, 0x9c, 0xf7, 0x94, 0xd3,
	0x67, 0x24, 0x51, 0x4c, 0xe1, 0x42, 0x34, 0xc4, 0x52, 0x63, 0xa0, 0x67, 0xbb, 0x9a, 0xa4, 0x93,
	0xcd, 0x3c, 0xff, 0x3b, 0xc2, 0x05, 0x24, 0xce, 0xae, 0x92, 0x74, 0xba, 0x99, 0xe7, 0x89, 0xb3,
	0xb8, 0x84, 0x59, 0xe0, 0x7d, 0x17, 0x9e, 0xf6, 0xb3, 0xc3, 0xed, 0xfe, 0x3b, 0x81, 0xb3, 0x43,
	0x01, 0xbe, 0xc0, 0xf9, 0x4e, 0x28, 0x74, 0x45, 0x78, 0x6b, 0x4e, 0xca, 0x9b, 0xff, 0x32, 0xeb,
	0x65, 0x17, 0x3d, 0x5a, 0x5b, 0x65, 0xb3, 0xd5, 0x0f, 0x8d, 0x82, 0xaf, 0x70, 0xf9, 0x14, 0x7d,
	0xa9, 0x8e, 0xfd, 0x58, 0xc8, 0x2d, 0x2c, 0x76, 0xbe, 0xfb, 0xe4, 0x31, 0x3d, 0xdf, 0x00, 0x07,
	0xe8, 0xc8, 0xb6, 0xc5, 0xac, 0x5f, 0xe1, 0xc3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x69,
	0xa0, 0xc3, 0x25, 0x02, 0x00, 0x00,
}
