// Code generated by protoc-gen-go. DO NOT EDIT.
// source: safety/pb/blacklist.proto

package gosionsvc_external_safety

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

type BlacklistSearchRequest struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Size                 int64    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Page                 int64    `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlacklistSearchRequest) Reset()         { *m = BlacklistSearchRequest{} }
func (m *BlacklistSearchRequest) String() string { return proto.CompactTextString(m) }
func (*BlacklistSearchRequest) ProtoMessage()    {}
func (*BlacklistSearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6f2b3a90f91407b, []int{0}
}

func (m *BlacklistSearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlacklistSearchRequest.Unmarshal(m, b)
}
func (m *BlacklistSearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlacklistSearchRequest.Marshal(b, m, deterministic)
}
func (m *BlacklistSearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlacklistSearchRequest.Merge(m, src)
}
func (m *BlacklistSearchRequest) XXX_Size() int {
	return xxx_messageInfo_BlacklistSearchRequest.Size(m)
}
func (m *BlacklistSearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlacklistSearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlacklistSearchRequest proto.InternalMessageInfo

func (m *BlacklistSearchRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *BlacklistSearchRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *BlacklistSearchRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *BlacklistSearchRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

type BlacklistSearchResponse struct {
	State                *dto.State       `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Item                 []*BlacklistItem `protobuf:"bytes,2,rep,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BlacklistSearchResponse) Reset()         { *m = BlacklistSearchResponse{} }
func (m *BlacklistSearchResponse) String() string { return proto.CompactTextString(m) }
func (*BlacklistSearchResponse) ProtoMessage()    {}
func (*BlacklistSearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6f2b3a90f91407b, []int{1}
}

func (m *BlacklistSearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlacklistSearchResponse.Unmarshal(m, b)
}
func (m *BlacklistSearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlacklistSearchResponse.Marshal(b, m, deterministic)
}
func (m *BlacklistSearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlacklistSearchResponse.Merge(m, src)
}
func (m *BlacklistSearchResponse) XXX_Size() int {
	return xxx_messageInfo_BlacklistSearchResponse.Size(m)
}
func (m *BlacklistSearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlacklistSearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlacklistSearchResponse proto.InternalMessageInfo

func (m *BlacklistSearchResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *BlacklistSearchResponse) GetItem() []*BlacklistItem {
	if m != nil {
		return m.Item
	}
	return nil
}

type BlacklistItem struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateUser           string   `protobuf:"bytes,2,opt,name=createUser,proto3" json:"createUser,omitempty"`
	CreateAt             int64    `protobuf:"varint,3,opt,name=createAt,proto3" json:"createAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlacklistItem) Reset()         { *m = BlacklistItem{} }
func (m *BlacklistItem) String() string { return proto.CompactTextString(m) }
func (*BlacklistItem) ProtoMessage()    {}
func (*BlacklistItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6f2b3a90f91407b, []int{2}
}

func (m *BlacklistItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlacklistItem.Unmarshal(m, b)
}
func (m *BlacklistItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlacklistItem.Marshal(b, m, deterministic)
}
func (m *BlacklistItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlacklistItem.Merge(m, src)
}
func (m *BlacklistItem) XXX_Size() int {
	return xxx_messageInfo_BlacklistItem.Size(m)
}
func (m *BlacklistItem) XXX_DiscardUnknown() {
	xxx_messageInfo_BlacklistItem.DiscardUnknown(m)
}

var xxx_messageInfo_BlacklistItem proto.InternalMessageInfo

func (m *BlacklistItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BlacklistItem) GetCreateUser() string {
	if m != nil {
		return m.CreateUser
	}
	return ""
}

func (m *BlacklistItem) GetCreateAt() int64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

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
	return fileDescriptor_e6f2b3a90f91407b, []int{3}
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
	return fileDescriptor_e6f2b3a90f91407b, []int{4}
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
	return fileDescriptor_e6f2b3a90f91407b, []int{5}
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
	proto.RegisterType((*BlacklistSearchRequest)(nil), "gosionsvc.external.safety.BlacklistSearchRequest")
	proto.RegisterType((*BlacklistSearchResponse)(nil), "gosionsvc.external.safety.BlacklistSearchResponse")
	proto.RegisterType((*BlacklistItem)(nil), "gosionsvc.external.safety.BlacklistItem")
	proto.RegisterType((*CheckRequest)(nil), "gosionsvc.external.safety.CheckRequest")
	proto.RegisterType((*RemoveRequest)(nil), "gosionsvc.external.safety.RemoveRequest")
	proto.RegisterType((*AddRequest)(nil), "gosionsvc.external.safety.AddRequest")
}

func init() { proto.RegisterFile("safety/pb/blacklist.proto", fileDescriptor_e6f2b3a90f91407b) }

var fileDescriptor_e6f2b3a90f91407b = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x8e, 0xd3, 0x40,
	0x0c, 0x55, 0x93, 0xb6, 0xb0, 0x5e, 0x16, 0x21, 0x0b, 0x96, 0x6c, 0x0f, 0x50, 0x45, 0x42, 0x54,
	0x5a, 0x69, 0x22, 0xc2, 0x0d, 0xed, 0xa5, 0xc0, 0x01, 0x38, 0x66, 0xc5, 0x19, 0x4d, 0x33, 0xa6,
	0x8d, 0x92, 0xcc, 0x84, 0x8c, 0x5b, 0x51, 0xce, 0x7c, 0x24, 0x9f, 0x83, 0x92, 0x49, 0x4a, 0x2b,
	0x68, 0x05, 0xdc, 0x9e, 0x1d, 0xfb, 0xd9, 0xef, 0x8d, 0x03, 0x57, 0x56, 0x7e, 0x26, 0xde, 0x46,
	0xd5, 0x22, 0x5a, 0x14, 0x32, 0xcd, 0x8b, 0xcc, 0xb2, 0xa8, 0x6a, 0xc3, 0x06, 0xaf, 0x96, 0xc6,
	0x66, 0x46, 0xdb, 0x4d, 0x2a, 0xe8, 0x2b, 0x53, 0xad, 0x65, 0x21, 0x5c, 0xf5, 0xe4, 0x3a, 0x37,
	0x9a, 0xf2, 0xdc, 0x88, 0x92, 0x22, 0x57, 0x15, 0xa5, 0xa6, 0x2c, 0x8d, 0xb6, 0x91, 0x62, 0xd3,
	0x63, 0xc7, 0x13, 0xae, 0xe0, 0xf2, 0x75, 0x4f, 0x7d, 0x4b, 0xb2, 0x4e, 0x57, 0x09, 0x7d, 0x59,
	0x93, 0x65, 0x7c, 0x08, 0xa3, 0x8d, 0x2c, 0xd6, 0x14, 0x0c, 0xa6, 0x83, 0xd9, 0x59, 0xe2, 0x02,
	0x7c, 0x00, 0x7e, 0x4e, 0xdb, 0xc0, 0x6b, 0x73, 0x0d, 0x44, 0x84, 0xa1, 0xcd, 0xbe, 0x51, 0xe0,
	0x4f, 0x07, 0x33, 0x3f, 0x69, 0x71, 0x93, 0xab, 0xe4, 0x92, 0x82, 0xa1, 0xcb, 0x35, 0x38, 0xfc,
	0x3e, 0x80, 0xc7, 0xbf, 0x8d, 0xb2, 0x95, 0xd1, 0x96, 0xf0, 0x1a, 0x46, 0x96, 0x25, 0xbb, 0x59,
	0xe7, 0xf1, 0x23, 0xb1, 0xb4, 0xa2, 0xdf, 0x53, 0xb1, 0x11, 0xb7, 0xcd, 0xc7, 0xc4, 0xd5, 0xe0,
	0x0d, 0x0c, 0x33, 0xa6, 0x32, 0xf0, 0xa6, 0xfe, 0xec, 0x3c, 0x9e, 0x89, 0xa3, 0x4e, 0x88, 0xdd,
	0xb8, 0xf7, 0x4c, 0x65, 0xd2, 0x76, 0x85, 0x9f, 0xe0, 0xe2, 0x20, 0xdd, 0xec, 0xaa, 0x65, 0xd9,
	0xcb, 0x6c, 0x31, 0x3e, 0x01, 0x48, 0x6b, 0x92, 0x4c, 0x1f, 0x2d, 0xd5, 0x9d, 0xd8, 0xbd, 0x0c,
	0x4e, 0xe0, 0xae, 0x8b, 0xe6, 0xdc, 0xe9, 0xde, 0xc5, 0xe1, 0x0d, 0xdc, 0x7b, 0xb3, 0xa2, 0x34,
	0xef, 0x7d, 0x44, 0x18, 0xf2, 0xb6, 0x72, 0xfc, 0x7e, 0xd2, 0x62, 0x0c, 0xe0, 0x4e, 0x6a, 0x34,
	0x93, 0xe6, 0x8e, 0xbc, 0x0f, 0xc3, 0xa7, 0x70, 0x91, 0x50, 0x69, 0x36, 0xd4, 0xb7, 0xdf, 0x07,
	0x2f, 0x53, 0xdd, 0x72, 0x5e, 0xa6, 0xc2, 0x57, 0x00, 0x73, 0xa5, 0xfe, 0x8b, 0x3c, 0xfe, 0xe1,
	0xc1, 0xd9, 0x4e, 0x3c, 0xbe, 0x05, 0x7f, 0xae, 0x14, 0x3e, 0x3b, 0x61, 0xe0, 0xaf, 0x49, 0x93,
	0xcb, 0x3f, 0xbd, 0xc9, 0xda, 0xe2, 0x07, 0x18, 0xbb, 0x85, 0xf1, 0xd4, 0x4b, 0x1c, 0x68, 0x3a,
	0xca, 0xf5, 0x0e, 0x46, 0xad, 0x75, 0xf8, 0xfc, 0x04, 0xd5, 0xbe, 0xb9, 0x47, 0x99, 0x0c, 0x8c,
	0xdd, 0x89, 0xe1, 0x8b, 0xbf, 0xb9, 0x8f, 0x83, 0xcb, 0x9f, 0xc4, 0xff, 0xd2, 0xe2, 0x2e, 0x78,
	0x31, 0x6e, 0x7f, 0xa7, 0x97, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x65, 0xbb, 0x92, 0xae, 0xb3,
	0x03, 0x00, 0x00,
}
