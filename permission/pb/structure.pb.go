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
	User                 bool     `protobuf:"varint,2,opt,name=user,proto3" json:"user,omitempty"`
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

func (m *GetListRequest) GetUser() bool {
	if m != nil {
		return m.User
	}
	return false
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
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x41, 0xab, 0xda, 0x40,
	0x10, 0xc6, 0xa8, 0x55, 0x47, 0xaa, 0xb0, 0xad, 0x6d, 0x08, 0x14, 0x24, 0xb4, 0x22, 0x08, 0x1b,
	0xb0, 0xa7, 0x7a, 0x2b, 0xa5, 0x2d, 0x82, 0xa7, 0x84, 0xd2, 0x5e, 0x63, 0x76, 0x08, 0x4b, 0xcc,
	0xee, 0x76, 0x77, 0xd3, 0x43, 0xff, 0xc6, 0xfb, 0xc3, 0x0f, 0x37, 0x26, 0xea, 0x43, 0x79, 0x97,
	0x77, 0x9b, 0x99, 0x9d, 0xef, 0x9b, 0x99, 0xef, 0x5b, 0xf8, 0xa0, 0x50, 0x97, 0xdc, 0x18, 0x2e,
	0x45, 0xa4, 0xf6, 0x91, 0xb1, 0xba, 0xca, 0x6c, 0xa5, 0x91, 0x2a, 0x2d, 0xad, 0x24, 0xb3, 0xdc,
	0x50, 0x83, 0xfa, 0x1f, 0xcf, 0x90, 0x9e, 0x3b, 0x83, 0x55, 0x21, 0x05, 0x16, 0x85, 0xa4, 0x25,
	0x46, 0xb9, 0x74, 0xe0, 0x4c, 0x96, 0xa5, 0x14, 0x26, 0x62, 0x56, 0x36, 0x71, 0xcd, 0x11, 0x7e,
	0x81, 0xd7, 0xdf, 0x34, 0xa6, 0x16, 0x63, 0xfc, 0x5b, 0xa1, 0xb1, 0x84, 0x40, 0x4f, 0xa4, 0x25,
	0xfa, 0x9d, 0x79, 0x67, 0x39, 0x8a, 0x5d, 0x4c, 0xde, 0x42, 0x3f, 0x55, 0x6a, 0xcb, 0x7c, 0xcf,
	0x15, 0xeb, 0x24, 0xdc, 0xc1, 0xe4, 0xbb, 0x48, 0xf7, 0x07, 0x64, 0x0d, 0x76, 0x0e, 0xe3, 0x76,
	0xc7, 0x2d, 0x3b, 0x51, 0x5c, 0x96, 0x88, 0x0f, 0x03, 0xa9, 0x50, 0x70, 0x91, 0x3b, 0xae, 0x61,
	0xdc, 0xa4, 0xe1, 0x06, 0x26, 0x3f, 0xd1, 0xee, 0xb8, 0xb1, 0x0d, 0x5b, 0x3b, 0xb5, 0x73, 0x31,
	0xf5, 0xb8, 0x5f, 0x65, 0x50, 0x9f, 0xe0, 0x2e, 0x0e, 0xff, 0xc3, 0xb4, 0xc5, 0x1a, 0x25, 0x85,
	0x41, 0xb2, 0x82, 0xbe, 0xb1, 0xa9, 0xad, 0xef, 0x18, 0xaf, 0x67, 0x34, 0x37, 0xb4, 0xb9, 0x9c,
	0x59, 0x49, 0x93, 0xe3, 0x63, 0x5c, 0xf7, 0x90, 0x0d, 0xf4, 0x58, 0x6a, 0x53, 0xdf, 0x9b, 0x77,
	0x97, 0xe3, 0xf5, 0x82, 0xde, 0xd4, 0x95, 0x26, 0xbc, 0x54, 0x07, 0x4c, 0x9a, 0x6b, 0x62, 0x87,
	0x09, 0x0b, 0x98, 0x3e, 0x79, 0xb8, 0x29, 0x61, 0x00, 0xc3, 0xcc, 0xe9, 0xfc, 0xd5, 0xba, 0xd5,
	0xbb, 0x71, 0x9b, 0x5f, 0x8a, 0xd2, 0xbd, 0x12, 0x85, 0x4c, 0xc0, 0xe3, 0xcc, 0xef, 0x39, 0x1e,
	0x8f, 0xb3, 0xf5, 0x83, 0x07, 0xa3, 0xf3, 0x9c, 0x04, 0xde, 0xd4, 0xde, 0xfd, 0x32, 0xa8, 0xcf,
	0xe5, 0x8f, 0x77, 0xf6, 0xbf, 0xf2, 0x39, 0x78, 0x77, 0x4b, 0x91, 0xca, 0x90, 0xdf, 0xf0, 0xbe,
	0x6e, 0xfc, 0x51, 0x89, 0xcc, 0x72, 0x29, 0x5e, 0x8a, 0xf8, 0x0f, 0x0c, 0x4e, 0x26, 0x91, 0x4f,
	0x77, 0x88, 0xae, 0x3f, 0x40, 0xb0, 0x78, 0xae, 0xad, 0xf6, 0x7a, 0xff, 0xca, 0x7d, 0xe5, 0xcf,
	0x8f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x18, 0x4d, 0x13, 0xa7, 0x2f, 0x03, 0x00, 0x00,
}
