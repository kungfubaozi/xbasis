// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication/pb/inner/token.proto

package gosionsvc_internal_authentication

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

type GenerateRequest struct {
	Auth                 *dto.Authorize `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Route                bool           `protobuf:"varint,2,opt,name=route,proto3" json:"route,omitempty"`
	RelationId           string         `protobuf:"bytes,3,opt,name=relationId,proto3" json:"relationId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GenerateRequest) Reset()         { *m = GenerateRequest{} }
func (m *GenerateRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateRequest) ProtoMessage()    {}
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c49315c962f44d5c, []int{0}
}

func (m *GenerateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateRequest.Unmarshal(m, b)
}
func (m *GenerateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateRequest.Marshal(b, m, deterministic)
}
func (m *GenerateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateRequest.Merge(m, src)
}
func (m *GenerateRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateRequest.Size(m)
}
func (m *GenerateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateRequest proto.InternalMessageInfo

func (m *GenerateRequest) GetAuth() *dto.Authorize {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *GenerateRequest) GetRoute() bool {
	if m != nil {
		return m.Route
	}
	return false
}

func (m *GenerateRequest) GetRelationId() string {
	if m != nil {
		return m.RelationId
	}
	return ""
}

type GenerateResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	AccessToken          string     `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken         string     `protobuf:"bytes,3,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GenerateResponse) Reset()         { *m = GenerateResponse{} }
func (m *GenerateResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateResponse) ProtoMessage()    {}
func (*GenerateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c49315c962f44d5c, []int{1}
}

func (m *GenerateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateResponse.Unmarshal(m, b)
}
func (m *GenerateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateResponse.Marshal(b, m, deterministic)
}
func (m *GenerateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateResponse.Merge(m, src)
}
func (m *GenerateResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateResponse.Size(m)
}
func (m *GenerateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateResponse proto.InternalMessageInfo

func (m *GenerateResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *GenerateResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *GenerateResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func init() {
	proto.RegisterType((*GenerateRequest)(nil), "gosionsvc.internal.authentication.GenerateRequest")
	proto.RegisterType((*GenerateResponse)(nil), "gosionsvc.internal.authentication.GenerateResponse")
}

func init() {
	proto.RegisterFile("authentication/pb/inner/token.proto", fileDescriptor_c49315c962f44d5c)
}

var fileDescriptor_c49315c962f44d5c = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x41, 0x4f, 0x32, 0x31,
	0x10, 0xcd, 0x7e, 0x9f, 0x18, 0x18, 0x4c, 0x34, 0x8d, 0x26, 0xc8, 0xc1, 0xac, 0x78, 0x21, 0x21,
	0x76, 0x13, 0xf8, 0x05, 0x9e, 0x8c, 0xd7, 0xea, 0x1f, 0x28, 0xcb, 0x08, 0xcd, 0xc2, 0x0c, 0x76,
	0x66, 0x39, 0x18, 0xaf, 0xfe, 0x6f, 0xc3, 0x56, 0x22, 0x70, 0x31, 0xde, 0xda, 0xd7, 0xf7, 0xe6,
	0xbd, 0xbe, 0x81, 0x3b, 0x5f, 0xeb, 0x02, 0x49, 0x43, 0xe9, 0x35, 0x30, 0x15, 0xeb, 0x69, 0x11,
	0x88, 0x30, 0x16, 0xca, 0x15, 0x92, 0x5d, 0x47, 0x56, 0x36, 0xb7, 0x73, 0x96, 0xc0, 0x24, 0x9b,
	0xd2, 0x06, 0x52, 0x8c, 0xe4, 0x97, 0xf6, 0x50, 0xd7, 0x1f, 0x55, 0x4c, 0x58, 0x55, 0x6c, 0x57,
	0x58, 0x24, 0x76, 0x51, 0xf2, 0x6a, 0xc5, 0x24, 0xc5, 0x4c, 0x79, 0x77, 0x4e, 0xf3, 0x06, 0x1b,
	0x38, 0x7f, 0x44, 0xc2, 0xe8, 0x15, 0x1d, 0xbe, 0xd5, 0x28, 0x6a, 0xee, 0xe1, 0x64, 0x3b, 0xb1,
	0x97, 0xe5, 0xd9, 0xb0, 0x3b, 0xbe, 0xb6, 0x73, 0xb1, 0x3b, 0xcd, 0x4c, 0xd9, 0x3e, 0xd4, 0xba,
	0xe0, 0x18, 0xde, 0xd1, 0x35, 0x34, 0x73, 0x09, 0xad, 0xc8, 0xb5, 0x62, 0xef, 0x5f, 0x9e, 0x0d,
	0xdb, 0x2e, 0x5d, 0xcc, 0x0d, 0x40, 0xc4, 0x65, 0x13, 0xe8, 0x69, 0xd6, 0xfb, 0x9f, 0x67, 0xc3,
	0x8e, 0xdb, 0x43, 0x06, 0x9f, 0x19, 0x5c, 0xfc, 0x18, 0xcb, 0x9a, 0x49, 0xd0, 0x8c, 0xa0, 0x25,
	0xea, 0x15, 0xbf, 0xad, 0xaf, 0x8e, 0xad, 0x9f, 0xb7, 0x8f, 0x2e, 0x71, 0x4c, 0x0e, 0x5d, 0x5f,
	0x96, 0x28, 0xf2, 0xb2, 0xad, 0xa7, 0x71, 0xef, 0xb8, 0x7d, 0xc8, 0x0c, 0xe0, 0x2c, 0xe2, 0x6b,
	0x44, 0x59, 0x24, 0x4a, 0x4a, 0x71, 0x80, 0x8d, 0x3f, 0xa0, 0x95, 0xc8, 0x02, 0xed, 0x5d, 0x1e,
	0x33, 0xb6, 0xbf, 0xb6, 0x6c, 0x8f, 0x5a, 0xeb, 0x4f, 0xfe, 0xa4, 0x49, 0x1f, 0x9e, 0x9e, 0x36,
	0x4b, 0x98, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x14, 0xf2, 0x37, 0x80, 0xfb, 0x01, 0x00, 0x00,
}