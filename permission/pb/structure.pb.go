// Code generated by protoc-gen-go. DO NOT EDIT.
// source: permission/pb/structure.proto

package gs_service_permission

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	dto "konekko.me/gosion/commons/dto"
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

type CreateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AppId                string   `protobuf:"bytes,2,opt,name=appId,proto3" json:"appId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f325e6597bcc547a, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type EnabledRequest struct {
	StructureId          string   `protobuf:"bytes,1,opt,name=structureId,proto3" json:"structureId,omitempty"`
	Opening              bool     `protobuf:"varint,2,opt,name=opening,proto3" json:"opening,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnabledRequest) Reset()         { *m = EnabledRequest{} }
func (m *EnabledRequest) String() string { return proto.CompactTextString(m) }
func (*EnabledRequest) ProtoMessage()    {}
func (*EnabledRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f325e6597bcc547a, []int{1}
}

func (m *EnabledRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnabledRequest.Unmarshal(m, b)
}
func (m *EnabledRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnabledRequest.Marshal(b, m, deterministic)
}
func (m *EnabledRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnabledRequest.Merge(m, src)
}
func (m *EnabledRequest) XXX_Size() int {
	return xxx_messageInfo_EnabledRequest.Size(m)
}
func (m *EnabledRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EnabledRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EnabledRequest proto.InternalMessageInfo

func (m *EnabledRequest) GetStructureId() string {
	if m != nil {
		return m.StructureId
	}
	return ""
}

func (m *EnabledRequest) GetOpening() bool {
	if m != nil {
		return m.Opening
	}
	return false
}

type GetListRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListRequest) Reset()         { *m = GetListRequest{} }
func (m *GetListRequest) String() string { return proto.CompactTextString(m) }
func (*GetListRequest) ProtoMessage()    {}
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f325e6597bcc547a, []int{2}
}

func (m *GetListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListRequest.Unmarshal(m, b)
}
func (m *GetListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListRequest.Marshal(b, m, deterministic)
}
func (m *GetListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListRequest.Merge(m, src)
}
func (m *GetListRequest) XXX_Size() int {
	return xxx_messageInfo_GetListRequest.Size(m)
}
func (m *GetListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetListRequest proto.InternalMessageInfo

func (m *GetListRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type GetListResponse struct {
	State                *dto.State         `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Data                 []*SimpleStructure `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetListResponse) Reset()         { *m = GetListResponse{} }
func (m *GetListResponse) String() string { return proto.CompactTextString(m) }
func (*GetListResponse) ProtoMessage()    {}
func (*GetListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f325e6597bcc547a, []int{3}
}

func (m *GetListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListResponse.Unmarshal(m, b)
}
func (m *GetListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListResponse.Marshal(b, m, deterministic)
}
func (m *GetListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListResponse.Merge(m, src)
}
func (m *GetListResponse) XXX_Size() int {
	return xxx_messageInfo_GetListResponse.Size(m)
}
func (m *GetListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetListResponse proto.InternalMessageInfo

func (m *GetListResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *GetListResponse) GetData() []*SimpleStructure {
	if m != nil {
		return m.Data
	}
	return nil
}

type SimpleStructure struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateAt             int64    `protobuf:"varint,2,opt,name=createAt,proto3" json:"createAt,omitempty"`
	Opening              bool     `protobuf:"varint,3,opt,name=opening,proto3" json:"opening,omitempty"`
	Id                   string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleStructure) Reset()         { *m = SimpleStructure{} }
func (m *SimpleStructure) String() string { return proto.CompactTextString(m) }
func (*SimpleStructure) ProtoMessage()    {}
func (*SimpleStructure) Descriptor() ([]byte, []int) {
	return fileDescriptor_f325e6597bcc547a, []int{4}
}

func (m *SimpleStructure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleStructure.Unmarshal(m, b)
}
func (m *SimpleStructure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleStructure.Marshal(b, m, deterministic)
}
func (m *SimpleStructure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleStructure.Merge(m, src)
}
func (m *SimpleStructure) XXX_Size() int {
	return xxx_messageInfo_SimpleStructure.Size(m)
}
func (m *SimpleStructure) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleStructure.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleStructure proto.InternalMessageInfo

func (m *SimpleStructure) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SimpleStructure) GetCreateAt() int64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

func (m *SimpleStructure) GetOpening() bool {
	if m != nil {
		return m.Opening
	}
	return false
}

func (m *SimpleStructure) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "gs.service.permission.CreateRequest")
	proto.RegisterType((*EnabledRequest)(nil), "gs.service.permission.EnabledRequest")
	proto.RegisterType((*GetListRequest)(nil), "gs.service.permission.GetListRequest")
	proto.RegisterType((*GetListResponse)(nil), "gs.service.permission.GetListResponse")
	proto.RegisterType((*SimpleStructure)(nil), "gs.service.permission.SimpleStructure")
}

func init() { proto.RegisterFile("permission/pb/structure.proto", fileDescriptor_f325e6597bcc547a) }

var fileDescriptor_f325e6597bcc547a = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcd, 0x8a, 0xdb, 0x30,
	0x10, 0xc7, 0x89, 0x93, 0x34, 0xc9, 0x84, 0x26, 0xa0, 0x36, 0xad, 0x31, 0x14, 0x82, 0x69, 0x43,
	0x20, 0x20, 0x43, 0x7a, 0x6a, 0x6f, 0xa5, 0xf4, 0x23, 0x90, 0x93, 0x4d, 0xe9, 0x5e, 0x1d, 0x6b,
	0x30, 0xc2, 0xb1, 0xa4, 0xb5, 0xe4, 0x3d, 0xec, 0xbb, 0xee, 0xbb, 0x2c, 0x2b, 0xc7, 0x4e, 0xbc,
	0xc4, 0xec, 0x65, 0x6f, 0x9a, 0xd1, 0xcc, 0x6f, 0x46, 0xff, 0x19, 0xc1, 0x27, 0x85, 0x45, 0xce,
	0xb5, 0xe6, 0x52, 0x04, 0xea, 0x10, 0x68, 0x53, 0x94, 0x89, 0x29, 0x0b, 0xa4, 0xaa, 0x90, 0x46,
	0x92, 0x45, 0xaa, 0xa9, 0xc6, 0xe2, 0x8e, 0x27, 0x48, 0xcf, 0x91, 0xde, 0x26, 0x93, 0x02, 0xb3,
	0x4c, 0xd2, 0x1c, 0x83, 0x54, 0xda, 0xe4, 0x44, 0xe6, 0xb9, 0x14, 0x3a, 0x60, 0x46, 0xd6, 0xe7,
	0x8a, 0xe1, 0x7f, 0x83, 0xb7, 0x3f, 0x0b, 0x8c, 0x0d, 0x86, 0x78, 0x5b, 0xa2, 0x36, 0x84, 0xc0,
	0x40, 0xc4, 0x39, 0xba, 0xbd, 0x65, 0x6f, 0x3d, 0x09, 0xed, 0x99, 0xbc, 0x87, 0x61, 0xac, 0xd4,
	0x8e, 0xb9, 0x8e, 0x75, 0x56, 0x86, 0xbf, 0x87, 0xd9, 0x2f, 0x11, 0x1f, 0x8e, 0xc8, 0xea, 0xdc,
	0x25, 0x4c, 0x9b, 0x1e, 0x77, 0xec, 0x84, 0xb8, 0x74, 0x11, 0x17, 0x46, 0x52, 0xa1, 0xe0, 0x22,
	0xb5, 0xac, 0x71, 0x58, 0x9b, 0xfe, 0x0a, 0x66, 0x7f, 0xd0, 0xec, 0xb9, 0x36, 0x35, 0xad, 0xa9,
	0xda, 0xbb, 0xac, 0x7a, 0x0f, 0xf3, 0x26, 0x4e, 0x2b, 0x29, 0x34, 0x92, 0x0d, 0x0c, 0xb5, 0x89,
	0x4d, 0xd5, 0xf3, 0x74, 0xbb, 0xa0, 0xa9, 0xa6, 0xf5, 0x2b, 0x99, 0x91, 0x34, 0x7a, 0xba, 0x0c,
	0xab, 0x18, 0xf2, 0x1d, 0x06, 0x2c, 0x36, 0xb1, 0xeb, 0x2c, 0xfb, 0xeb, 0xe9, 0x76, 0x45, 0xaf,
	0x6a, 0x48, 0x23, 0x9e, 0xab, 0x23, 0x46, 0x75, 0xe7, 0xa1, 0xcd, 0xf1, 0x33, 0x98, 0x3f, 0xbb,
	0xb8, 0x2a, 0x97, 0x07, 0xe3, 0xc4, 0x6a, 0xfa, 0xc3, 0xd8, 0x57, 0xf6, 0xc3, 0xc6, 0xbe, 0x14,
	0xa0, 0xdf, 0x12, 0x80, 0xcc, 0xc0, 0xe1, 0xcc, 0x1d, 0x58, 0x8e, 0xc3, 0xd9, 0xf6, 0xc1, 0x81,
	0xc9, 0xb9, 0x4e, 0x04, 0xef, 0xaa, 0x39, 0xfd, 0xd3, 0x58, 0x9c, 0xdd, 0x9f, 0x3b, 0xfa, 0x6f,
	0xcd, 0xd4, 0xfb, 0x70, 0x4d, 0x91, 0x52, 0x93, 0xff, 0xf0, 0xb1, 0x0a, 0xfc, 0x5d, 0x8a, 0xc4,
	0x70, 0x29, 0x5e, 0x0b, 0xfc, 0x17, 0x46, 0xa7, 0xd5, 0x20, 0x5f, 0x3a, 0x40, 0xed, 0xd5, 0xe9,
	0x24, 0xdd, 0xc0, 0xe8, 0x34, 0xee, 0x4e, 0x52, 0x7b, 0x6d, 0xbc, 0xd5, 0x4b, 0x61, 0xd5, 0xd6,
	0x1c, 0xde, 0xd8, 0x0f, 0xf0, 0xf5, 0x31, 0x00, 0x00, 0xff, 0xff, 0xee, 0x2b, 0xd9, 0xd0, 0x65,
	0x03, 0x00, 0x00,
}