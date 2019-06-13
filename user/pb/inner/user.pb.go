// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/pb/inner/user.proto

package gosionsvc_internal_user

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

type ExistsRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExistsRequest) Reset()         { *m = ExistsRequest{} }
func (m *ExistsRequest) String() string { return proto.CompactTextString(m) }
func (*ExistsRequest) ProtoMessage()    {}
func (*ExistsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2adf0be4aa740dac, []int{0}
}

func (m *ExistsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExistsRequest.Unmarshal(m, b)
}
func (m *ExistsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExistsRequest.Marshal(b, m, deterministic)
}
func (m *ExistsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExistsRequest.Merge(m, src)
}
func (m *ExistsRequest) XXX_Size() int {
	return xxx_messageInfo_ExistsRequest.Size(m)
}
func (m *ExistsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExistsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExistsRequest proto.InternalMessageInfo

func (m *ExistsRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type SimpleUserInfo struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	UserId               string     `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Username             string     `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Icon                 string     `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	RealName             string     `protobuf:"bytes,5,opt,name=realName,proto3" json:"realName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SimpleUserInfo) Reset()         { *m = SimpleUserInfo{} }
func (m *SimpleUserInfo) String() string { return proto.CompactTextString(m) }
func (*SimpleUserInfo) ProtoMessage()    {}
func (*SimpleUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2adf0be4aa740dac, []int{1}
}

func (m *SimpleUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleUserInfo.Unmarshal(m, b)
}
func (m *SimpleUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleUserInfo.Marshal(b, m, deterministic)
}
func (m *SimpleUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleUserInfo.Merge(m, src)
}
func (m *SimpleUserInfo) XXX_Size() int {
	return xxx_messageInfo_SimpleUserInfo.Size(m)
}
func (m *SimpleUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleUserInfo proto.InternalMessageInfo

func (m *SimpleUserInfo) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *SimpleUserInfo) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SimpleUserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SimpleUserInfo) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *SimpleUserInfo) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

type GetUserInfoByIdRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInfoByIdRequest) Reset()         { *m = GetUserInfoByIdRequest{} }
func (m *GetUserInfoByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoByIdRequest) ProtoMessage()    {}
func (*GetUserInfoByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2adf0be4aa740dac, []int{2}
}

func (m *GetUserInfoByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoByIdRequest.Unmarshal(m, b)
}
func (m *GetUserInfoByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoByIdRequest.Marshal(b, m, deterministic)
}
func (m *GetUserInfoByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoByIdRequest.Merge(m, src)
}
func (m *GetUserInfoByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoByIdRequest.Size(m)
}
func (m *GetUserInfoByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoByIdRequest proto.InternalMessageInfo

func (m *GetUserInfoByIdRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*ExistsRequest)(nil), "gosionsvc.internal.user.ExistsRequest")
	proto.RegisterType((*SimpleUserInfo)(nil), "gosionsvc.internal.user.SimpleUserInfo")
	proto.RegisterType((*GetUserInfoByIdRequest)(nil), "gosionsvc.internal.user.GetUserInfoByIdRequest")
}

func init() { proto.RegisterFile("user/pb/inner/user.proto", fileDescriptor_2adf0be4aa740dac) }

var fileDescriptor_2adf0be4aa740dac = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x51, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0xb6, 0xa5, 0x8e, 0xa8, 0xb0, 0x60, 0x0c, 0x79, 0x92, 0x3c, 0x58, 0xa1, 0xb0,
	0x2b, 0xf5, 0x06, 0x82, 0x48, 0x7d, 0xf0, 0x21, 0xc5, 0x03, 0xa4, 0xc9, 0x58, 0x96, 0x24, 0x3b,
	0x31, 0xb3, 0x11, 0x3d, 0x8e, 0x27, 0xf1, 0x6a, 0xb2, 0x89, 0x91, 0x56, 0x1a, 0x7c, 0x9b, 0x7f,
	0xe7, 0xdb, 0x9f, 0x99, 0x7f, 0x20, 0x68, 0x18, 0x6b, 0x55, 0xad, 0x95, 0x36, 0x06, 0x6b, 0xe5,
	0x94, 0xac, 0x6a, 0xb2, 0x24, 0x2e, 0x36, 0xc4, 0x9a, 0x0c, 0xbf, 0xa5, 0x52, 0x1b, 0x8b, 0xb5,
	0x49, 0x0a, 0xe9, 0xda, 0xe1, 0x3c, 0x27, 0x83, 0x79, 0x4e, 0xb2, 0x44, 0xd5, 0x31, 0x2a, 0xa5,
	0xb2, 0x24, 0xc3, 0x2a, 0xb3, 0xd4, 0xd7, 0x9d, 0x4b, 0x34, 0x83, 0x93, 0xfb, 0x77, 0xcd, 0x96,
	0x63, 0x7c, 0x6d, 0x90, 0xad, 0xf0, 0x61, 0xe2, 0x5c, 0x96, 0x59, 0xe0, 0x5d, 0x7a, 0xd7, 0x47,
	0xf1, 0x8f, 0x8a, 0x3e, 0x3d, 0x38, 0x5d, 0xe9, 0xb2, 0x2a, 0xf0, 0xd9, 0x3d, 0x98, 0x17, 0x12,
	0x73, 0x18, 0xb3, 0x4d, 0x2c, 0xb6, 0xe4, 0xf1, 0xe2, 0x5c, 0x6e, 0x58, 0xf6, 0xee, 0x99, 0x25,
	0xb9, 0x72, 0xcd, 0xb8, 0x63, 0xb6, 0x7c, 0x0f, 0xb6, 0x7d, 0x45, 0x08, 0x53, 0x57, 0x99, 0xa4,
	0xc4, 0xe0, 0xb0, 0xed, 0xfc, 0x6a, 0x21, 0x60, 0xa4, 0x53, 0x32, 0xc1, 0xa8, 0x7d, 0x6f, 0x6b,
	0xc7, 0xd7, 0x98, 0x14, 0x4f, 0x8e, 0x1f, 0x77, 0x7c, 0xaf, 0xa3, 0x1b, 0xf0, 0x1f, 0xd0, 0xf6,
	0xf3, 0xdd, 0x7d, 0x2c, 0xb3, 0x7f, 0xb6, 0x5a, 0x7c, 0x79, 0x30, 0x72, 0xbc, 0x78, 0x84, 0xa9,
	0xe6, 0x2e, 0x09, 0x71, 0x25, 0x07, 0xa2, 0x95, 0x3b, 0x51, 0x85, 0xfe, 0xbe, 0x85, 0x1b, 0x16,
	0x39, 0x9c, 0xfd, 0x19, 0x43, 0xa8, 0x41, 0xcb, 0xfd, 0x03, 0x87, 0xb3, 0xc1, 0x0f, 0xbb, 0x47,
	0x58, 0x4f, 0xda, 0x3b, 0xde, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x48, 0xaa, 0xab, 0x29,
	0x02, 0x00, 0x00,
}
