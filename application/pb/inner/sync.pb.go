// Code generated by protoc-gen-go. DO NOT EDIT.
// source: application/pb/inner/sync.proto

package gosionsvc_internal_application

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

type CheckRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	AppId                string   `protobuf:"bytes,2,opt,name=appId,proto3" json:"appId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20c3a61a4d5ebe51, []int{0}
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

func (m *CheckRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CheckRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type UserInfo struct {
	GId                  string   `protobuf:"bytes,1,opt,name=gId,proto3" json:"gId,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Account              string   `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	Icon                 string   `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	RealName             string   `protobuf:"bytes,5,opt,name=realName,proto3" json:"realName,omitempty"`
	AppId                string   `protobuf:"bytes,6,opt,name=appId,proto3" json:"appId,omitempty"`
	Url                  string   `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_20c3a61a4d5ebe51, []int{1}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetGId() string {
	if m != nil {
		return m.GId
	}
	return ""
}

func (m *UserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserInfo) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *UserInfo) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *UserInfo) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *UserInfo) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *UserInfo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckRequest)(nil), "gosionsvc.internal.application.CheckRequest")
	proto.RegisterType((*UserInfo)(nil), "gosionsvc.internal.application.UserInfo")
}

func init() { proto.RegisterFile("application/pb/inner/sync.proto", fileDescriptor_20c3a61a4d5ebe51) }

var fileDescriptor_20c3a61a4d5ebe51 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xa9, 0x6d, 0xd3, 0x76, 0xf0, 0x20, 0x83, 0x94, 0xd0, 0x83, 0x4a, 0x4f, 0x05, 0x65,
	0x03, 0x7a, 0xf5, 0xe6, 0xa9, 0x42, 0x3d, 0x54, 0xfb, 0x00, 0xdb, 0xcd, 0x5a, 0x43, 0x92, 0x99,
	0x75, 0x77, 0x23, 0xf8, 0x4a, 0x3e, 0x9f, 0x0f, 0x20, 0x9b, 0x34, 0x35, 0x17, 0xf5, 0xe0, 0x6d,
	0xfe, 0xc9, 0xfc, 0x3f, 0x5f, 0xfe, 0x85, 0x73, 0x69, 0x4c, 0x91, 0x29, 0xe9, 0x33, 0xa6, 0xc4,
	0x6c, 0x93, 0x8c, 0x48, 0xdb, 0xc4, 0xbd, 0x93, 0x12, 0xc6, 0xb2, 0x67, 0x3c, 0xdb, 0xb1, 0xcb,
	0x98, 0xdc, 0x9b, 0x12, 0x19, 0x79, 0x6d, 0x49, 0x16, 0xa2, 0xe3, 0x99, 0x5d, 0xe6, 0x4c, 0x3a,
	0xcf, 0x59, 0x94, 0x3a, 0x69, 0x4e, 0x13, 0xc5, 0x65, 0xc9, 0xe4, 0x92, 0xd4, 0x73, 0x3b, 0x37,
	0x61, 0xf3, 0x5b, 0x38, 0xbe, 0x7b, 0xd1, 0x2a, 0x5f, 0xeb, 0xd7, 0x4a, 0x3b, 0x8f, 0x53, 0x88,
	0x2a, 0xa7, 0xed, 0x32, 0x8d, 0x7b, 0x17, 0xbd, 0xc5, 0x64, 0xbd, 0x57, 0x78, 0x0a, 0x43, 0x69,
	0xcc, 0x32, 0x8d, 0x8f, 0xea, 0x75, 0x23, 0xe6, 0x1f, 0x3d, 0x18, 0x6f, 0xc2, 0x01, 0x3d, 0x33,
	0x9e, 0x40, 0x7f, 0x77, 0xf0, 0x85, 0x11, 0x67, 0x30, 0x0e, 0x76, 0x92, 0xa5, 0xde, 0xfb, 0x0e,
	0x1a, 0x63, 0x18, 0x49, 0xa5, 0xb8, 0x22, 0x1f, 0xf7, 0xeb, 0x4f, 0xad, 0x44, 0x84, 0x41, 0xa6,
	0x98, 0xe2, 0x41, 0xbd, 0xae, 0xe7, 0x90, 0x64, 0xb5, 0x2c, 0x1e, 0x42, 0xd2, 0xb0, 0x49, 0x6a,
	0xf5, 0x37, 0x5a, 0xd4, 0x41, 0x0b, 0x34, 0x95, 0x2d, 0xe2, 0x51, 0x43, 0x53, 0xd9, 0xe2, 0xfa,
	0x73, 0x0f, 0x1b, 0xaa, 0xc4, 0x15, 0x4c, 0x9e, 0xac, 0x24, 0x67, 0xd8, 0x7a, 0x5c, 0x88, 0xdf,
	0x2b, 0x15, 0xed, 0x3f, 0xce, 0xa6, 0x62, 0xe7, 0x44, 0xdb, 0x60, 0xea, 0x59, 0x3c, 0x7a, 0xe9,
	0x2b, 0x87, 0x2b, 0x18, 0xd6, 0x35, 0xe2, 0xd5, 0x5f, 0x51, 0xdd, 0xb6, 0x7f, 0x8c, 0xbb, 0x87,
	0x68, 0x63, 0x52, 0xe9, 0xf5, 0xff, 0xd1, 0xb6, 0x51, 0xfd, 0xd0, 0x37, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x4e, 0x38, 0x9b, 0x49, 0x58, 0x02, 0x00, 0x00,
}