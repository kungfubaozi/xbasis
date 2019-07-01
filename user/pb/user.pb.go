// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/pb/user.proto

package xbasissvc_external_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	dto "konekko.me/xbasis/commons/dto"
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

type SearchRequest struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Page                 int64    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int64    `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca9694bbbb4de3af, []int{0}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *SearchRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SearchRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type SearchResponse struct {
	State                *dto.State        `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Data                 []*SimpleUserData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca9694bbbb4de3af, []int{1}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *SearchResponse) GetData() []*SimpleUserData {
	if m != nil {
		return m.Data
	}
	return nil
}

type SimpleUserData struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	RealName             string   `protobuf:"bytes,4,opt,name=realName,proto3" json:"realName,omitempty"`
	UserId               string   `protobuf:"bytes,5,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleUserData) Reset()         { *m = SimpleUserData{} }
func (m *SimpleUserData) String() string { return proto.CompactTextString(m) }
func (*SimpleUserData) ProtoMessage()    {}
func (*SimpleUserData) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca9694bbbb4de3af, []int{2}
}

func (m *SimpleUserData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleUserData.Unmarshal(m, b)
}
func (m *SimpleUserData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleUserData.Marshal(b, m, deterministic)
}
func (m *SimpleUserData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleUserData.Merge(m, src)
}
func (m *SimpleUserData) XXX_Size() int {
	return xxx_messageInfo_SimpleUserData.Size(m)
}
func (m *SimpleUserData) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleUserData.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleUserData proto.InternalMessageInfo

func (m *SimpleUserData) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SimpleUserData) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *SimpleUserData) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SimpleUserData) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *SimpleUserData) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "xbasissvc.external.user.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "xbasissvc.external.user.SearchResponse")
	proto.RegisterType((*SimpleUserData)(nil), "xbasissvc.external.user.SimpleUserData")
}

func init() { proto.RegisterFile("user/pb/user.proto", fileDescriptor_ca9694bbbb4de3af) }

var fileDescriptor_ca9694bbbb4de3af = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0x4d, 0x5a, 0xec, 0x14, 0x8b, 0x2c, 0xa2, 0xb1, 0xa7, 0xd2, 0x83, 0x2d, 0x08, 0x1b,
	0xa8, 0x47, 0xaf, 0x5e, 0xbc, 0x78, 0x48, 0xf1, 0xe0, 0x49, 0xa6, 0xcd, 0x60, 0x43, 0x3e, 0x36,
	0x66, 0x37, 0xa5, 0x7a, 0xf0, 0x17, 0xf8, 0xa3, 0x65, 0x32, 0x8d, 0xe8, 0xa1, 0x78, 0xca, 0x7b,
	0x6f, 0x5e, 0xde, 0xce, 0xdb, 0x05, 0xd5, 0x58, 0xaa, 0xa3, 0x6a, 0x1d, 0xf1, 0x57, 0x57, 0xb5,
	0x71, 0x46, 0x5d, 0xee, 0xd7, 0x68, 0x53, 0x6b, 0x77, 0x1b, 0x4d, 0x7b, 0x47, 0x75, 0x89, 0xb9,
	0xe6, 0xf1, 0xe4, 0x26, 0x33, 0x25, 0x65, 0x99, 0xd1, 0x05, 0x45, 0xe2, 0x89, 0x36, 0xa6, 0x28,
	0x4c, 0x69, 0xa3, 0xc4, 0x99, 0x0e, 0x4b, 0xca, 0xec, 0x05, 0x4e, 0x57, 0x84, 0xf5, 0x66, 0x1b,
	0xd3, 0x5b, 0x43, 0xd6, 0xa9, 0x73, 0xe8, 0xef, 0x30, 0x6f, 0x28, 0xf4, 0xa6, 0xde, 0x62, 0x18,
	0x0b, 0x51, 0x67, 0xe0, 0x67, 0xf4, 0x1e, 0xf6, 0x5a, 0x8d, 0xa1, 0x52, 0x10, 0x54, 0xf8, 0x4a,
	0xa1, 0x3f, 0xf5, 0x16, 0x7e, 0xdc, 0x62, 0xd6, 0x6c, 0xfa, 0x41, 0x61, 0x20, 0x1a, 0xe3, 0xd9,
	0x27, 0x8c, 0xbb, 0x03, 0x6c, 0x65, 0x4a, 0x4b, 0x2a, 0x82, 0xbe, 0x75, 0xe8, 0xe4, 0x84, 0xd1,
	0xf2, 0x4a, 0xcb, 0x92, 0xba, 0x5b, 0x2c, 0x71, 0x46, 0xaf, 0xd8, 0x10, 0x8b, 0x4f, 0xdd, 0x41,
	0x90, 0xa0, 0xc3, 0xb0, 0x37, 0xf5, 0x17, 0xa3, 0xe5, 0x5c, 0x1f, 0x29, 0xae, 0x57, 0x69, 0x51,
	0xe5, 0xf4, 0x64, 0xa9, 0xbe, 0x47, 0x87, 0x71, 0xfb, 0xd3, 0xec, 0xcb, 0x83, 0xf1, 0xdf, 0x81,
	0x9a, 0xc0, 0x09, 0xfb, 0x4b, 0x2c, 0xba, 0x96, 0x3f, 0x9c, 0xeb, 0x57, 0x5b, 0x53, 0xd2, 0xa1,
	0xaa, 0x10, 0x56, 0xa9, 0xc0, 0x34, 0x6f, 0xdb, 0x0e, 0x63, 0x21, 0x9c, 0x53, 0x13, 0xe6, 0x8f,
	0x9c, 0x13, 0x48, 0x4e, 0xc7, 0xd5, 0x05, 0x0c, 0x38, 0xf3, 0x21, 0x09, 0xfb, 0xed, 0xe4, 0xc0,
	0x96, 0x08, 0x01, 0xef, 0xa1, 0x9e, 0x61, 0x20, 0xd7, 0xa2, 0xae, 0x8f, 0xf7, 0xf9, 0xfd, 0x30,
	0x93, 0xf9, 0xbf, 0x3e, 0xb9, 0xdf, 0xf5, 0xa0, 0x7d, 0xd9, 0xdb, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x52, 0x40, 0x24, 0xe1, 0x35, 0x02, 0x00, 0x00,
}