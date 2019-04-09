// Code generated by protoc-gen-go. DO NOT EDIT.
// source: permission/pb/verification.proto

package gs_service_permission

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

type HasPermissionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HasPermissionRequest) Reset()         { *m = HasPermissionRequest{} }
func (m *HasPermissionRequest) String() string { return proto.CompactTextString(m) }
func (*HasPermissionRequest) ProtoMessage()    {}
func (*HasPermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e59ca96348c63fe, []int{0}
}

func (m *HasPermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasPermissionRequest.Unmarshal(m, b)
}
func (m *HasPermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasPermissionRequest.Marshal(b, m, deterministic)
}
func (m *HasPermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasPermissionRequest.Merge(m, src)
}
func (m *HasPermissionRequest) XXX_Size() int {
	return xxx_messageInfo_HasPermissionRequest.Size(m)
}
func (m *HasPermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HasPermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HasPermissionRequest proto.InternalMessageInfo

type HasPermissionResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	User                 string     `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	AppId                string     `protobuf:"bytes,3,opt,name=appId,proto3" json:"appId,omitempty"`
	ClientId             string     `protobuf:"bytes,4,opt,name=clientId,proto3" json:"clientId,omitempty"`
	TraceId              string     `protobuf:"bytes,5,opt,name=traceId,proto3" json:"traceId,omitempty"`
	Ip                   string     `protobuf:"bytes,6,opt,name=ip,proto3" json:"ip,omitempty"`
	UserAgent            string     `protobuf:"bytes,7,opt,name=userAgent,proto3" json:"userAgent,omitempty"`
	UserDevice           string     `protobuf:"bytes,8,opt,name=userDevice,proto3" json:"userDevice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HasPermissionResponse) Reset()         { *m = HasPermissionResponse{} }
func (m *HasPermissionResponse) String() string { return proto.CompactTextString(m) }
func (*HasPermissionResponse) ProtoMessage()    {}
func (*HasPermissionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e59ca96348c63fe, []int{1}
}

func (m *HasPermissionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasPermissionResponse.Unmarshal(m, b)
}
func (m *HasPermissionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasPermissionResponse.Marshal(b, m, deterministic)
}
func (m *HasPermissionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasPermissionResponse.Merge(m, src)
}
func (m *HasPermissionResponse) XXX_Size() int {
	return xxx_messageInfo_HasPermissionResponse.Size(m)
}
func (m *HasPermissionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HasPermissionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HasPermissionResponse proto.InternalMessageInfo

func (m *HasPermissionResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *HasPermissionResponse) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *HasPermissionResponse) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *HasPermissionResponse) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *HasPermissionResponse) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *HasPermissionResponse) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *HasPermissionResponse) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *HasPermissionResponse) GetUserDevice() string {
	if m != nil {
		return m.UserDevice
	}
	return ""
}

func init() {
	proto.RegisterType((*HasPermissionRequest)(nil), "gs.service.permission.HasPermissionRequest")
	proto.RegisterType((*HasPermissionResponse)(nil), "gs.service.permission.HasPermissionResponse")
}

func init() { proto.RegisterFile("permission/pb/verification.proto", fileDescriptor_8e59ca96348c63fe) }

var fileDescriptor_8e59ca96348c63fe = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0xc7, 0xd9, 0xda, 0xed, 0xc7, 0x28, 0x1e, 0x42, 0x2b, 0xa1, 0x88, 0x94, 0x9e, 0x0a, 0x95,
	0x2c, 0xd4, 0x27, 0x10, 0x3d, 0xd8, 0x9b, 0x54, 0xf0, 0xbe, 0xcd, 0x8e, 0x6b, 0x58, 0x37, 0x13,
	0x33, 0x69, 0x1f, 0xdc, 0x27, 0x90, 0xcd, 0xd2, 0x0f, 0xa4, 0x07, 0x6f, 0xf3, 0xff, 0x98, 0x09,
	0xfc, 0x02, 0x53, 0x87, 0xbe, 0x36, 0xcc, 0x86, 0x6c, 0xe6, 0x36, 0xd9, 0x0e, 0xbd, 0xf9, 0x30,
	0x3a, 0x0f, 0x86, 0xac, 0x72, 0x9e, 0x02, 0x89, 0x71, 0xc9, 0x8a, 0xd1, 0xef, 0x8c, 0x46, 0x75,
	0x2c, 0x4f, 0x16, 0x15, 0x59, 0xac, 0x2a, 0x52, 0x35, 0x66, 0x25, 0xc5, 0x7d, 0x4d, 0x75, 0x4d,
	0x96, 0xb3, 0x22, 0xd0, 0x7e, 0x6e, 0x6f, 0xcc, 0x6e, 0x60, 0xf4, 0x92, 0xf3, 0xeb, 0x61, 0x7b,
	0x8d, 0xdf, 0x5b, 0xe4, 0x30, 0xfb, 0x49, 0x60, 0xfc, 0x27, 0x60, 0x47, 0x96, 0x51, 0x2c, 0x20,
	0xe5, 0x90, 0x07, 0x94, 0xc9, 0x34, 0x99, 0x5f, 0x2e, 0xc7, 0xaa, 0x64, 0xb5, 0xbf, 0x59, 0x04,
	0x52, 0x6f, 0x4d, 0xb8, 0x6e, 0x3b, 0x42, 0x40, 0x77, 0xcb, 0xe8, 0x65, 0x67, 0x9a, 0xcc, 0x87,
	0xeb, 0x38, 0x8b, 0x11, 0xa4, 0xb9, 0x73, 0xab, 0x42, 0x5e, 0x44, 0xb3, 0x15, 0x62, 0x02, 0x03,
	0xfd, 0x65, 0xd0, 0x86, 0x55, 0x21, 0xbb, 0x31, 0x38, 0x68, 0x21, 0xa1, 0x1f, 0x7c, 0xae, 0x71,
	0x55, 0xc8, 0x34, 0x46, 0x7b, 0x29, 0xae, 0xa1, 0x63, 0x9c, 0xec, 0x45, 0xb3, 0x63, 0x9c, 0xb8,
	0x85, 0x61, 0xf3, 0xc6, 0x63, 0x89, 0x36, 0xc8, 0x7e, 0xb4, 0x8f, 0x86, 0xb8, 0x03, 0x68, 0xc4,
	0x33, 0x36, 0xc8, 0xe4, 0x20, 0xc6, 0x27, 0xce, 0xd2, 0xc3, 0xd5, 0xfb, 0x09, 0x66, 0xb1, 0x81,
	0xf4, 0xe9, 0x13, 0x75, 0x25, 0x16, 0xea, 0x2c, 0x6a, 0x75, 0x0e, 0xdd, 0xe4, 0xfe, 0x7f, 0xe5,
	0x16, 0xe7, 0xa6, 0x17, 0xff, 0xe1, 0xe1, 0x37, 0x00, 0x00, 0xff, 0xff, 0x93, 0x10, 0x37, 0xc3,
	0xef, 0x01, 0x00, 0x00,
}
