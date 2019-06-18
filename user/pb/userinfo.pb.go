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
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
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

func (m *GetInfoByIdRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetInfoResponse)(nil), "gosionsvc.external.user.GetInfoResponse")
	proto.RegisterType((*GetInfoByIdRequest)(nil), "gosionsvc.external.user.GetInfoByIdRequest")
}

func init() { proto.RegisterFile("user/pb/userinfo.proto", fileDescriptor_76493317093c0d65) }

var fileDescriptor_76493317093c0d65 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x25, 0xda, 0x96, 0x76, 0x14, 0x84, 0x05, 0x6b, 0x08, 0x1e, 0x4a, 0x4f, 0x81, 0xc8, 0x2e,
	0xd4, 0x3f, 0xf0, 0x52, 0x0a, 0xe2, 0x61, 0xc5, 0x0f, 0x48, 0x93, 0x69, 0x08, 0x69, 0x76, 0xd2,
	0xec, 0x44, 0xf4, 0x8f, 0xfc, 0x4c, 0xd9, 0x4d, 0x62, 0x0f, 0x22, 0x9e, 0x76, 0xde, 0xbe, 0x37,
	0x6f, 0x1e, 0x0f, 0x96, 0x9d, 0xc5, 0x56, 0x35, 0x7b, 0xe5, 0xde, 0xd2, 0x1c, 0x48, 0x36, 0x2d,
	0x31, 0x89, 0xbb, 0x82, 0x6c, 0x49, 0xc6, 0xbe, 0x67, 0x12, 0x3f, 0x18, 0x5b, 0x93, 0x1e, 0xa5,
	0x93, 0x44, 0x49, 0x45, 0x06, 0xab, 0x8a, 0x64, 0x8d, 0xaa, 0xd7, 0xa8, 0x8c, 0xea, 0x9a, 0x8c,
	0x55, 0x39, 0xd3, 0x38, 0xf7, 0x2e, 0xeb, 0xaf, 0x00, 0x6e, 0xb6, 0xc8, 0x3b, 0x73, 0x20, 0x8d,
	0xb6, 0x21, 0x63, 0x51, 0x24, 0x30, 0xb5, 0x9c, 0x32, 0x86, 0xc1, 0x2a, 0x88, 0xaf, 0x36, 0xb7,
	0xb2, 0xb0, 0x72, 0xdc, 0xca, 0x99, 0xe4, 0xab, 0x23, 0x75, 0xaf, 0x11, 0x11, 0xcc, 0xdd, 0x55,
	0x93, 0xd6, 0x18, 0x5e, 0xac, 0x82, 0x78, 0xa1, 0x7f, 0xb0, 0xb8, 0x87, 0x45, 0xd1, 0x52, 0xd7,
	0xbc, 0x38, 0xf2, 0xd2, 0x93, 0xe7, 0x0f, 0x21, 0x60, 0x52, 0x66, 0x64, 0xc2, 0x89, 0x27, 0xfc,
	0x3c, 0xba, 0xb9, 0x38, 0xe1, 0xf4, 0xec, 0xe6, 0xf0, 0xfa, 0x01, 0xc4, 0x90, 0xf4, 0xe9, 0x73,
	0x97, 0x6b, 0x3c, 0x75, 0x68, 0x59, 0x2c, 0x61, 0xe6, 0x15, 0xb9, 0x4f, 0xbb, 0xd0, 0x03, 0xda,
	0x9c, 0x60, 0xfe, 0x36, 0x6c, 0x0a, 0x84, 0xeb, 0x2d, 0xf2, 0x33, 0x65, 0xe9, 0xd1, 0xe3, 0x44,
	0xfe, 0xd1, 0x9d, 0xfc, 0x7d, 0x20, 0x8a, 0xff, 0x13, 0x8f, 0xbd, 0xed, 0x67, 0xbe, 0xd2, 0xc7,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0xd8, 0x27, 0x6e, 0xb2, 0x01, 0x00, 0x00,
}
