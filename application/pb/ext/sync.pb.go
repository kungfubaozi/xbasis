// Code generated by protoc-gen-go. DO NOT EDIT.
// source: application/pb/ext/sync.proto

package gs_ext_service_application

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
	return fileDescriptor_e164f260a2b33b3a, []int{0}
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
	return fileDescriptor_e164f260a2b33b3a, []int{1}
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
	proto.RegisterType((*CheckRequest)(nil), "gs.ext.service.application.CheckRequest")
	proto.RegisterType((*UserInfo)(nil), "gs.ext.service.application.UserInfo")
}

func init() { proto.RegisterFile("application/pb/ext/sync.proto", fileDescriptor_e164f260a2b33b3a) }

var fileDescriptor_e164f260a2b33b3a = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xe9, 0xbf, 0x6d, 0xda, 0x2e, 0xff, 0x83, 0x0c, 0x52, 0x42, 0x40, 0x90, 0xe2, 0xa1,
	0x20, 0x6c, 0x40, 0xaf, 0xde, 0x04, 0x21, 0x17, 0x0f, 0xd5, 0x3e, 0xc0, 0x76, 0x33, 0xd6, 0x90,
	0x64, 0x67, 0xdd, 0x9d, 0x48, 0x7d, 0x25, 0x9f, 0xcc, 0xc7, 0x90, 0x4d, 0x9a, 0xda, 0x4b, 0xbd,
	0x78, 0x9b, 0x6f, 0x32, 0xdf, 0x8f, 0x2f, 0xdf, 0x8a, 0x0b, 0x65, 0x6d, 0x55, 0x68, 0xc5, 0x05,
	0x99, 0xd4, 0x6e, 0x52, 0xdc, 0x71, 0xea, 0x3f, 0x8c, 0x96, 0xd6, 0x11, 0x13, 0x24, 0x5b, 0x2f,
	0x71, 0xc7, 0xd2, 0xa3, 0x7b, 0x2f, 0x34, 0xca, 0xa3, 0xeb, 0xe4, 0xba, 0x24, 0x83, 0x65, 0x49,
	0xb2, 0xc6, 0x74, 0x4b, 0x3e, 0x00, 0x34, 0xd5, 0x35, 0x19, 0x9f, 0xe6, 0x4c, 0xfd, 0xdc, 0x81,
	0x16, 0x77, 0xe2, 0xff, 0xfd, 0x2b, 0xea, 0x72, 0x85, 0x6f, 0x0d, 0x7a, 0x86, 0xb9, 0x88, 0x1a,
	0x8f, 0x2e, 0xcb, 0xe3, 0xc1, 0xe5, 0x60, 0x39, 0x5b, 0xed, 0x15, 0x9c, 0x8b, 0xb1, 0xb2, 0x36,
	0xcb, 0xe3, 0x7f, 0xed, 0xba, 0x13, 0x8b, 0xcf, 0x81, 0x98, 0xae, 0xc3, 0x81, 0x79, 0x21, 0x38,
	0x13, 0xc3, 0xed, 0xc1, 0x17, 0x46, 0x48, 0xc4, 0x34, 0xd8, 0x8d, 0xaa, 0x71, 0xef, 0x3b, 0x68,
	0x88, 0xc5, 0x44, 0x69, 0x4d, 0x8d, 0xe1, 0x78, 0xd8, 0x7e, 0xea, 0x25, 0x80, 0x18, 0x15, 0x9a,
	0x4c, 0x3c, 0x6a, 0xd7, 0xed, 0x1c, 0x48, 0x0e, 0x55, 0xf5, 0x18, 0x48, 0xe3, 0x8e, 0xd4, 0xeb,
	0x9f, 0x68, 0xd1, 0x51, 0xb4, 0x90, 0xa6, 0x71, 0x55, 0x3c, 0xe9, 0xd2, 0x34, 0xae, 0xba, 0xf9,
	0xda, 0x87, 0x0d, 0x35, 0x42, 0x26, 0x66, 0xcf, 0x4e, 0x19, 0x6f, 0xc9, 0x31, 0x5c, 0xc9, 0xd3,
	0x75, 0xca, 0xfe, 0xff, 0x92, 0x79, 0xb8, 0xea, 0xdb, 0xcb, 0x99, 0xe4, 0x13, 0x2b, 0x6e, 0x3c,
	0x64, 0x62, 0xdc, 0x56, 0x08, 0xcb, 0xdf, 0x30, 0xc7, 0x2d, 0x9f, 0x44, 0x3d, 0x88, 0x68, 0x6d,
	0x73, 0xc5, 0xf8, 0xb7, 0x48, 0x9b, 0xa8, 0x7d, 0xdc, 0xdb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x7d, 0xe9, 0xff, 0x09, 0x46, 0x02, 0x00, 0x00,
}