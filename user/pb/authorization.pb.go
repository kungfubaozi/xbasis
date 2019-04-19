// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/pb/authorization.proto

package gs_service_user

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

type SyncRequest struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncRequest) Reset()         { *m = SyncRequest{} }
func (m *SyncRequest) String() string { return proto.CompactTextString(m) }
func (*SyncRequest) ProtoMessage()    {}
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4fbd78633bb21c2, []int{0}
}

func (m *SyncRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncRequest.Unmarshal(m, b)
}
func (m *SyncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncRequest.Marshal(b, m, deterministic)
}
func (m *SyncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncRequest.Merge(m, src)
}
func (m *SyncRequest) XXX_Size() int {
	return xxx_messageInfo_SyncRequest.Size(m)
}
func (m *SyncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncRequest proto.InternalMessageInfo

func (m *SyncRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func init() {
	proto.RegisterType((*SyncRequest)(nil), "gs.service.user.SyncRequest")
}

func init() { proto.RegisterFile("user/pb/authorization.proto", fileDescriptor_d4fbd78633bb21c2) }

var fileDescriptor_d4fbd78633bb21c2 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2f, 0x48, 0xd2, 0x4f, 0x2c, 0x2d, 0xc9, 0xc8, 0x2f, 0xca, 0xac, 0x4a, 0x2c, 0xc9, 0xcc,
	0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4f, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x03, 0xa9, 0x93, 0xd2, 0xce, 0xce, 0xcf, 0x4b, 0xcd, 0xce, 0xce, 0xd7,
	0xcb, 0x4d, 0xd5, 0x4f, 0xcf, 0x2f, 0xce, 0xcc, 0xcf, 0xd3, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf,
	0x2b, 0xd6, 0x4f, 0x29, 0xc9, 0x87, 0xb1, 0x21, 0xba, 0x95, 0x34, 0xb9, 0xb8, 0x83, 0x2b, 0xf3,
	0x92, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xa4, 0xb8, 0x38, 0x92, 0x73, 0x32, 0x53,
	0xf3, 0x4a, 0x3c, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x23, 0x5f, 0x2e,
	0x5e, 0x47, 0x64, 0xfb, 0x85, 0x6c, 0xb8, 0x58, 0x40, 0x7a, 0x85, 0x64, 0xf4, 0xd0, 0x9c, 0xa0,
	0x87, 0x64, 0xa4, 0x94, 0x18, 0x48, 0x16, 0x66, 0x69, 0x4a, 0x49, 0xbe, 0x5e, 0x70, 0x49, 0x62,
	0x49, 0x69, 0x71, 0x12, 0x1b, 0xd8, 0x01, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x50, 0x66,
	0x73, 0x85, 0xdd, 0x00, 0x00, 0x00,
}
