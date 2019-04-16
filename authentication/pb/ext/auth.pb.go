// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication/pb/ext/auth.proto

package gs_ext_service_authentication

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

type VerifyRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	FunctionRoles        []string `protobuf:"bytes,3,rep,name=functionRoles,proto3" json:"functionRoles,omitempty"`
	Funcs                string   `protobuf:"bytes,4,opt,name=funcs,proto3" json:"funcs,omitempty"`
	Access               int64    `protobuf:"varint,5,opt,name=access,proto3" json:"access,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRequest) Reset()         { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()    {}
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e1e612aa43fed8d, []int{0}
}

func (m *VerifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyRequest.Unmarshal(m, b)
}
func (m *VerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyRequest.Marshal(b, m, deterministic)
}
func (m *VerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRequest.Merge(m, src)
}
func (m *VerifyRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyRequest.Size(m)
}
func (m *VerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRequest proto.InternalMessageInfo

func (m *VerifyRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *VerifyRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *VerifyRequest) GetFunctionRoles() []string {
	if m != nil {
		return m.FunctionRoles
	}
	return nil
}

func (m *VerifyRequest) GetFuncs() string {
	if m != nil {
		return m.Funcs
	}
	return ""
}

func (m *VerifyRequest) GetAccess() int64 {
	if m != nil {
		return m.Access
	}
	return 0
}

type VerifyResponse struct {
	UserId               string     `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ClientId             string     `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	ClientPlatform       int64      `protobuf:"varint,3,opt,name=clientPlatform,proto3" json:"clientPlatform,omitempty"`
	State                *dto.State `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	AppId                string     `protobuf:"bytes,5,opt,name=appId,proto3" json:"appId,omitempty"`
	Relation             string     `protobuf:"bytes,6,opt,name=relation,proto3" json:"relation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *VerifyResponse) Reset()         { *m = VerifyResponse{} }
func (m *VerifyResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyResponse) ProtoMessage()    {}
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e1e612aa43fed8d, []int{1}
}

func (m *VerifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyResponse.Unmarshal(m, b)
}
func (m *VerifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyResponse.Marshal(b, m, deterministic)
}
func (m *VerifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyResponse.Merge(m, src)
}
func (m *VerifyResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyResponse.Size(m)
}
func (m *VerifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyResponse proto.InternalMessageInfo

func (m *VerifyResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *VerifyResponse) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *VerifyResponse) GetClientPlatform() int64 {
	if m != nil {
		return m.ClientPlatform
	}
	return 0
}

func (m *VerifyResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *VerifyResponse) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *VerifyResponse) GetRelation() string {
	if m != nil {
		return m.Relation
	}
	return ""
}

func init() {
	proto.RegisterType((*VerifyRequest)(nil), "gs.ext.service.authentication.VerifyRequest")
	proto.RegisterType((*VerifyResponse)(nil), "gs.ext.service.authentication.VerifyResponse")
}

func init() { proto.RegisterFile("authentication/pb/ext/auth.proto", fileDescriptor_2e1e612aa43fed8d) }

var fileDescriptor_2e1e612aa43fed8d = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xc1, 0x6a, 0xe3, 0x30,
	0x10, 0xc5, 0xeb, 0xc4, 0x6c, 0xb4, 0x24, 0x07, 0xb1, 0xbb, 0x98, 0x40, 0xc1, 0x84, 0x52, 0x02,
	0x69, 0x65, 0x48, 0xbf, 0xa0, 0xc7, 0xdc, 0x8a, 0x0b, 0xbd, 0x3b, 0xf2, 0x24, 0x31, 0xb6, 0x35,
	0xae, 0x67, 0x5c, 0xd2, 0x0f, 0xe9, 0x57, 0xf5, 0xa7, 0x8a, 0xa4, 0xb8, 0x90, 0x1e, 0x42, 0x6f,
	0x7a, 0x6f, 0x46, 0x6f, 0xde, 0xcc, 0x13, 0x49, 0xde, 0xf3, 0x01, 0x0c, 0x97, 0x3a, 0xe7, 0x12,
	0x4d, 0xda, 0x6e, 0x53, 0x38, 0x72, 0x6a, 0x59, 0xd5, 0x76, 0xc8, 0x28, 0xaf, 0xf6, 0xa4, 0xe0,
	0xc8, 0x8a, 0xa0, 0x7b, 0x2d, 0x35, 0xa8, 0xf3, 0x0f, 0xf3, 0x55, 0x85, 0x06, 0xaa, 0x0a, 0x55,
	0x03, 0xe9, 0x1e, 0xc9, 0x6a, 0x68, 0x6c, 0x1a, 0x34, 0x94, 0x16, 0x8c, 0xc3, 0xdb, 0x6b, 0x2d,
	0xde, 0x03, 0x31, 0x7d, 0x86, 0xae, 0xdc, 0xbd, 0x65, 0xf0, 0xd2, 0x03, 0xb1, 0xfc, 0x2b, 0xc6,
	0x8c, 0x15, 0x98, 0x38, 0x48, 0x82, 0xe5, 0x24, 0xf3, 0x40, 0xce, 0xc5, 0x6f, 0x5d, 0x97, 0x60,
	0x78, 0x53, 0xc4, 0xbf, 0x5c, 0xe1, 0x0b, 0xcb, 0x6b, 0x31, 0xdd, 0xf5, 0x46, 0xdb, 0xe1, 0x19,
	0xd6, 0x40, 0x71, 0x98, 0x84, 0xcb, 0x49, 0x76, 0x4e, 0x5a, 0x5d, 0x4b, 0x50, 0x3c, 0xf2, 0xba,
	0x0e, 0xc8, 0xff, 0x22, 0xca, 0xb5, 0x06, 0xa2, 0x78, 0x9c, 0x04, 0xcb, 0x30, 0x3b, 0xa1, 0xc5,
	0x47, 0x20, 0x66, 0x83, 0x2f, 0x6a, 0xd1, 0x10, 0xd8, 0xd6, 0x9e, 0xa0, 0xdb, 0x14, 0x27, 0x67,
	0x27, 0x74, 0xd1, 0xda, 0x8d, 0x98, 0xf9, 0xf7, 0x63, 0x9d, 0xf3, 0x0e, 0xbb, 0x26, 0x0e, 0xdd,
	0x98, 0x6f, 0xac, 0x5c, 0x89, 0x31, 0x71, 0xce, 0xe0, 0xcc, 0xfd, 0x59, 0xff, 0x53, 0x7b, 0x52,
	0xc3, 0xa1, 0x0a, 0x46, 0xf5, 0x64, 0x8b, 0x99, 0xef, 0xb1, 0x9b, 0xe4, 0x6d, 0xbb, 0x29, 0x9c,
	0xe5, 0x49, 0xe6, 0x81, 0xb5, 0xd1, 0x41, 0xed, 0x22, 0x88, 0x23, 0x6f, 0x63, 0xc0, 0xeb, 0x46,
	0x8c, 0x1e, 0x7a, 0x3e, 0x48, 0x10, 0x91, 0x5f, 0x4a, 0xde, 0xaa, 0x8b, 0x21, 0xaa, 0xb3, 0x4c,
	0xe6, 0x77, 0x3f, 0xec, 0xf6, 0x97, 0xda, 0x46, 0x2e, 0xdb, 0xfb, 0xcf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x64, 0xd9, 0xf1, 0xa6, 0x4b, 0x02, 0x00, 0x00,
}
