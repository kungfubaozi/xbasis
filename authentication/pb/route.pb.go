// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication/pb/route.proto

package gs_service_authentication

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

type LogoutRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba1f9b290d6e538f, []int{0}
}

func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (m *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(m, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

func (m *LogoutRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type RefreshRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshRequest) Reset()         { *m = RefreshRequest{} }
func (m *RefreshRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshRequest) ProtoMessage()    {}
func (*RefreshRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba1f9b290d6e538f, []int{1}
}

func (m *RefreshRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshRequest.Unmarshal(m, b)
}
func (m *RefreshRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshRequest.Marshal(b, m, deterministic)
}
func (m *RefreshRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshRequest.Merge(m, src)
}
func (m *RefreshRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshRequest.Size(m)
}
func (m *RefreshRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshRequest proto.InternalMessageInfo

func (m *RefreshRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type RefreshResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Token                string     `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RefreshResponse) Reset()         { *m = RefreshResponse{} }
func (m *RefreshResponse) String() string { return proto.CompactTextString(m) }
func (*RefreshResponse) ProtoMessage()    {}
func (*RefreshResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba1f9b290d6e538f, []int{2}
}

func (m *RefreshResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshResponse.Unmarshal(m, b)
}
func (m *RefreshResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshResponse.Marshal(b, m, deterministic)
}
func (m *RefreshResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshResponse.Merge(m, src)
}
func (m *RefreshResponse) XXX_Size() int {
	return xxx_messageInfo_RefreshResponse.Size(m)
}
func (m *RefreshResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshResponse proto.InternalMessageInfo

func (m *RefreshResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *RefreshResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type PushRequest struct {
	RouteTo              string   `protobuf:"bytes,1,opt,name=routeTo,proto3" json:"routeTo,omitempty"`
	Redirect             string   `protobuf:"bytes,2,opt,name=redirect,proto3" json:"redirect,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushRequest) Reset()         { *m = PushRequest{} }
func (m *PushRequest) String() string { return proto.CompactTextString(m) }
func (*PushRequest) ProtoMessage()    {}
func (*PushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba1f9b290d6e538f, []int{3}
}

func (m *PushRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushRequest.Unmarshal(m, b)
}
func (m *PushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushRequest.Marshal(b, m, deterministic)
}
func (m *PushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushRequest.Merge(m, src)
}
func (m *PushRequest) XXX_Size() int {
	return xxx_messageInfo_PushRequest.Size(m)
}
func (m *PushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushRequest proto.InternalMessageInfo

func (m *PushRequest) GetRouteTo() string {
	if m != nil {
		return m.RouteTo
	}
	return ""
}

func (m *PushRequest) GetRedirect() string {
	if m != nil {
		return m.Redirect
	}
	return ""
}

type PushResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	AccessToken          string     `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken         string     `protobuf:"bytes,3,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PushResponse) Reset()         { *m = PushResponse{} }
func (m *PushResponse) String() string { return proto.CompactTextString(m) }
func (*PushResponse) ProtoMessage()    {}
func (*PushResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba1f9b290d6e538f, []int{4}
}

func (m *PushResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushResponse.Unmarshal(m, b)
}
func (m *PushResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushResponse.Marshal(b, m, deterministic)
}
func (m *PushResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushResponse.Merge(m, src)
}
func (m *PushResponse) XXX_Size() int {
	return xxx_messageInfo_PushResponse.Size(m)
}
func (m *PushResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushResponse proto.InternalMessageInfo

func (m *PushResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *PushResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *PushResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func init() {
	proto.RegisterType((*LogoutRequest)(nil), "gs.service.authentication.LogoutRequest")
	proto.RegisterType((*RefreshRequest)(nil), "gs.service.authentication.RefreshRequest")
	proto.RegisterType((*RefreshResponse)(nil), "gs.service.authentication.RefreshResponse")
	proto.RegisterType((*PushRequest)(nil), "gs.service.authentication.PushRequest")
	proto.RegisterType((*PushResponse)(nil), "gs.service.authentication.PushResponse")
}

func init() { proto.RegisterFile("authentication/pb/route.proto", fileDescriptor_ba1f9b290d6e538f) }

var fileDescriptor_ba1f9b290d6e538f = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcd, 0x4e, 0xf2, 0x50,
	0x10, 0x0d, 0x7c, 0x1f, 0xa0, 0x03, 0x6a, 0x72, 0xa3, 0x06, 0x9b, 0x98, 0x90, 0x2e, 0x14, 0x25,
	0xb9, 0x8d, 0xf8, 0x08, 0xee, 0x8c, 0x0b, 0x53, 0x49, 0xdc, 0x5a, 0xca, 0x58, 0x9a, 0x4a, 0xa7,
	0xde, 0x99, 0xba, 0xf7, 0x09, 0x7c, 0x65, 0xc3, 0x6d, 0x41, 0xea, 0x0f, 0x12, 0x77, 0x9d, 0xdb,
	0x73, 0xce, 0x9c, 0x33, 0x33, 0x70, 0x1c, 0xe4, 0x32, 0xc5, 0x54, 0xe2, 0x30, 0x90, 0x98, 0x52,
	0x2f, 0x1b, 0x7b, 0x86, 0x72, 0x41, 0x9d, 0x19, 0x12, 0x52, 0x47, 0x11, 0x6b, 0x46, 0xf3, 0x12,
	0x87, 0xa8, 0xab, 0x48, 0x67, 0x90, 0x50, 0x8a, 0x49, 0x42, 0x7a, 0x86, 0x5e, 0x44, 0x3c, 0x27,
	0x87, 0x34, 0x9b, 0x51, 0xca, 0xde, 0x44, 0x68, 0xf1, 0x5d, 0xe8, 0xb8, 0x17, 0xb0, 0x73, 0x43,
	0x11, 0xe5, 0xe2, 0xe3, 0x73, 0x8e, 0x2c, 0xaa, 0x07, 0xed, 0x20, 0x0c, 0x91, 0x79, 0x44, 0x09,
	0xa6, 0xdd, 0x5a, 0xaf, 0xd6, 0xdf, 0xf6, 0x57, 0x9f, 0xdc, 0x21, 0xec, 0xfa, 0xf8, 0x68, 0x90,
	0xa7, 0x9b, 0x73, 0x46, 0xb0, 0xb7, 0xe4, 0x70, 0x46, 0x29, 0xa3, 0x1a, 0x40, 0x83, 0x25, 0x10,
	0xb4, 0xf0, 0xf6, 0xf0, 0x40, 0x47, 0xac, 0x17, 0xde, 0x26, 0x42, 0xfa, 0x6e, 0xfe, 0xd3, 0x2f,
	0x30, 0x6a, 0x1f, 0x1a, 0x62, 0xb5, 0xeb, 0x56, 0xbb, 0x28, 0xdc, 0x2b, 0x68, 0xdf, 0xe6, 0x1f,
	0x36, 0xba, 0xd0, 0xb2, 0x23, 0x1a, 0x51, 0x69, 0x61, 0x51, 0x2a, 0x07, 0xb6, 0x0c, 0x4e, 0x62,
	0x83, 0xa1, 0x94, 0x0a, 0xcb, 0xda, 0x7d, 0xad, 0x41, 0xa7, 0x50, 0xf9, 0x8b, 0xb1, 0x4f, 0xd1,
	0xeb, 0x5f, 0xa2, 0x2b, 0x17, 0x3a, 0xa6, 0x88, 0x5e, 0x40, 0xfe, 0x59, 0x48, 0xe5, 0x6d, 0xf8,
	0x56, 0x87, 0xa6, 0x3f, 0xf7, 0x6a, 0xd4, 0x3d, 0xfc, 0xcf, 0x72, 0x9e, 0xaa, 0x13, 0xfd, 0xe3,
	0x86, 0xf5, 0x4a, 0x68, 0xe7, 0xf4, 0x57, 0x5c, 0x19, 0xeb, 0x01, 0x5a, 0x65, 0x4f, 0x75, 0xb6,
	0x86, 0x53, 0x5d, 0xad, 0x73, 0xbe, 0x09, 0xb4, 0xec, 0x70, 0x0d, 0xcd, 0x27, 0x7b, 0x4b, 0xaa,
	0xbf, 0x86, 0x55, 0x39, 0x37, 0xe7, 0xf0, 0xbb, 0xe9, 0xe6, 0x3c, 0x6e, 0xda, 0xf3, 0xbc, 0x7c,
	0x0f, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x94, 0x1b, 0xa0, 0x07, 0x03, 0x00, 0x00,
}