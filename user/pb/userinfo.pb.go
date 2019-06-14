// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/pb/userinfo.proto

package gosionsvc_external_user

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

type GetInfoResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Username             string     `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	GroupName            string     `protobuf:"bytes,3,opt,name=groupName,proto3" json:"groupName,omitempty"`
	Icon                 string     `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	UserInfo             string     `protobuf:"bytes,5,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetInfoResponse) Reset()         { *m = GetInfoResponse{} }
func (m *GetInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetInfoResponse) ProtoMessage()    {}
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_76493317093c0d65, []int{0}
}

func (m *GetInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoResponse.Unmarshal(m, b)
}
func (m *GetInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoResponse.Merge(m, src)
}
func (m *GetInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetInfoResponse.Size(m)
}
func (m *GetInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoResponse proto.InternalMessageInfo

func (m *GetInfoResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *GetInfoResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetInfoResponse) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *GetInfoResponse) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *GetInfoResponse) GetUserInfo() string {
	if m != nil {
		return m.UserInfo
	}
	return ""
}

type GetInfoByIdRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoByIdRequest) Reset()         { *m = GetInfoByIdRequest{} }
func (m *GetInfoByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetInfoByIdRequest) ProtoMessage()    {}
func (*GetInfoByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_76493317093c0d65, []int{1}
}

func (m *GetInfoByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoByIdRequest.Unmarshal(m, b)
}
func (m *GetInfoByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoByIdRequest.Marshal(b, m, deterministic)
}
func (m *GetInfoByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoByIdRequest.Merge(m, src)
}
func (m *GetInfoByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetInfoByIdRequest.Size(m)
}
func (m *GetInfoByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoByIdRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetInfoResponse)(nil), "gosionsvc.external.user.GetInfoResponse")
	proto.RegisterType((*GetInfoByIdRequest)(nil), "gosionsvc.external.user.GetInfoByIdRequest")
}

func init() { proto.RegisterFile("user/pb/userinfo.proto", fileDescriptor_76493317093c0d65) }

var fileDescriptor_76493317093c0d65 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x89, 0xb6, 0xd2, 0xae, 0x82, 0xb0, 0xf8, 0x27, 0x04, 0x0f, 0xa5, 0xa7, 0x40, 0x60,
	0x17, 0xea, 0x1b, 0x78, 0x29, 0x05, 0xf1, 0x10, 0xf1, 0x01, 0xd2, 0x64, 0x1a, 0x42, 0x9a, 0x99,
	0x34, 0x33, 0x11, 0x7d, 0x23, 0x1f, 0x53, 0x76, 0xd3, 0xd8, 0x83, 0x88, 0xa7, 0x9d, 0x6f, 0xbf,
	0xdf, 0x7e, 0x3b, 0x7c, 0xea, 0xae, 0x67, 0xe8, 0x6c, 0xbb, 0xb5, 0xee, 0xac, 0x70, 0x47, 0xa6,
	0xed, 0x48, 0x48, 0xdf, 0x97, 0xc4, 0x15, 0x21, 0xbf, 0xe7, 0x06, 0x3e, 0x04, 0x3a, 0xcc, 0xf6,
	0xc6, 0x21, 0x51, 0x52, 0x13, 0x42, 0x5d, 0x93, 0x69, 0xc0, 0x0e, 0x8c, 0xcd, 0xa9, 0x69, 0x08,
	0xd9, 0x16, 0x42, 0xe3, 0x3c, 0xa4, 0x2c, 0xbf, 0x02, 0x75, 0xbd, 0x06, 0xd9, 0xe0, 0x8e, 0x52,
	0xe0, 0x96, 0x90, 0x41, 0x27, 0x6a, 0xca, 0x92, 0x09, 0x84, 0xc1, 0x22, 0x88, 0x2f, 0x57, 0xb7,
	0xa6, 0x64, 0x33, 0xbe, 0x2a, 0x84, 0xcc, 0xab, 0x33, 0xd3, 0x81, 0xd1, 0x91, 0x9a, 0xb9, 0x5f,
	0x31, 0x6b, 0x20, 0x3c, 0x5b, 0x04, 0xf1, 0x3c, 0xfd, 0xd1, 0xfa, 0x41, 0xcd, 0xcb, 0x8e, 0xfa,
	0xf6, 0xc5, 0x99, 0xe7, 0xde, 0x3c, 0x5d, 0x68, 0xad, 0x26, 0x55, 0x4e, 0x18, 0x4e, 0xbc, 0xe1,
	0xe7, 0x31, 0xcd, 0xad, 0x13, 0x4e, 0x4f, 0x69, 0x4e, 0x2f, 0x6f, 0x94, 0x3e, 0x6e, 0xfa, 0xf4,
	0xb9, 0x29, 0x52, 0x38, 0xf4, 0xc0, 0xb2, 0x3a, 0xa8, 0xd9, 0xdb, 0x91, 0xd0, 0xa0, 0xae, 0xd6,
	0x20, 0xcf, 0x94, 0x67, 0x7b, 0xaf, 0x13, 0xf3, 0x47, 0x47, 0xe6, 0x77, 0x50, 0x14, 0xff, 0x07,
	0x8f, 0xfd, 0x6c, 0x2f, 0x7c, 0x75, 0x8f, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x67, 0x22, 0x9e,
	0x70, 0x9a, 0x01, 0x00, 0x00,
}