// Code generated by protoc-gen-go. DO NOT EDIT.
// source: application/pb/application.proto

package gs_service_application

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

type SwitchRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	Enabled              bool     `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SwitchRequest) Reset()         { *m = SwitchRequest{} }
func (m *SwitchRequest) String() string { return proto.CompactTextString(m) }
func (*SwitchRequest) ProtoMessage()    {}
func (*SwitchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b974f2f4f838fbc1, []int{0}
}

func (m *SwitchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwitchRequest.Unmarshal(m, b)
}
func (m *SwitchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwitchRequest.Marshal(b, m, deterministic)
}
func (m *SwitchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwitchRequest.Merge(m, src)
}
func (m *SwitchRequest) XXX_Size() int {
	return xxx_messageInfo_SwitchRequest.Size(m)
}
func (m *SwitchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SwitchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SwitchRequest proto.InternalMessageInfo

func (m *SwitchRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *SwitchRequest) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type FindRequest struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRequest) Reset()         { *m = FindRequest{} }
func (m *FindRequest) String() string { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()    {}
func (*FindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b974f2f4f838fbc1, []int{1}
}

func (m *FindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRequest.Unmarshal(m, b)
}
func (m *FindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRequest.Marshal(b, m, deterministic)
}
func (m *FindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRequest.Merge(m, src)
}
func (m *FindRequest) XXX_Size() int {
	return xxx_messageInfo_FindRequest.Size(m)
}
func (m *FindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRequest proto.InternalMessageInfo

func (m *FindRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type ListResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Info                 []*AppInfo `protobuf:"bytes,2,rep,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b974f2f4f838fbc1, []int{2}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *ListResponse) GetInfo() []*AppInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type SimpleApplicationResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Info                 *AppInfo   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SimpleApplicationResponse) Reset()         { *m = SimpleApplicationResponse{} }
func (m *SimpleApplicationResponse) String() string { return proto.CompactTextString(m) }
func (*SimpleApplicationResponse) ProtoMessage()    {}
func (*SimpleApplicationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b974f2f4f838fbc1, []int{3}
}

func (m *SimpleApplicationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleApplicationResponse.Unmarshal(m, b)
}
func (m *SimpleApplicationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleApplicationResponse.Marshal(b, m, deterministic)
}
func (m *SimpleApplicationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleApplicationResponse.Merge(m, src)
}
func (m *SimpleApplicationResponse) XXX_Size() int {
	return xxx_messageInfo_SimpleApplicationResponse.Size(m)
}
func (m *SimpleApplicationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleApplicationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleApplicationResponse proto.InternalMessageInfo

func (m *SimpleApplicationResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *SimpleApplicationResponse) GetInfo() *AppInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type CreateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b974f2f4f838fbc1, []int{4}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type RemoveRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b974f2f4f838fbc1, []int{5}
}

func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type ChangeNameRequest struct {
	AppId                string   `protobuf:"bytes,2,opt,name=appId,proto3" json:"appId,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeNameRequest) Reset()         { *m = ChangeNameRequest{} }
func (m *ChangeNameRequest) String() string { return proto.CompactTextString(m) }
func (*ChangeNameRequest) ProtoMessage()    {}
func (*ChangeNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b974f2f4f838fbc1, []int{6}
}

func (m *ChangeNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeNameRequest.Unmarshal(m, b)
}
func (m *ChangeNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeNameRequest.Marshal(b, m, deterministic)
}
func (m *ChangeNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeNameRequest.Merge(m, src)
}
func (m *ChangeNameRequest) XXX_Size() int {
	return xxx_messageInfo_ChangeNameRequest.Size(m)
}
func (m *ChangeNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeNameRequest proto.InternalMessageInfo

func (m *ChangeNameRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ChangeNameRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AppInfo struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc                 string           `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	AppId                string           `protobuf:"bytes,3,opt,name=appId,proto3" json:"appId,omitempty"`
	Enabled              int64            `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	CreateAt             int64            `protobuf:"varint,5,opt,name=createAt,proto3" json:"createAt,omitempty"`
	Clients              []*AppClientInfo `protobuf:"bytes,6,rep,name=clients,proto3" json:"clients,omitempty"`
	Settings             *dto.AppSettings `protobuf:"bytes,7,opt,name=settings,proto3" json:"settings,omitempty"`
	CreateUserId         string           `protobuf:"bytes,8,opt,name=createUserId,proto3" json:"createUserId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AppInfo) Reset()         { *m = AppInfo{} }
func (m *AppInfo) String() string { return proto.CompactTextString(m) }
func (*AppInfo) ProtoMessage()    {}
func (*AppInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b974f2f4f838fbc1, []int{7}
}

func (m *AppInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppInfo.Unmarshal(m, b)
}
func (m *AppInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppInfo.Marshal(b, m, deterministic)
}
func (m *AppInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppInfo.Merge(m, src)
}
func (m *AppInfo) XXX_Size() int {
	return xxx_messageInfo_AppInfo.Size(m)
}
func (m *AppInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AppInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AppInfo proto.InternalMessageInfo

func (m *AppInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AppInfo) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *AppInfo) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *AppInfo) GetEnabled() int64 {
	if m != nil {
		return m.Enabled
	}
	return 0
}

func (m *AppInfo) GetCreateAt() int64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

func (m *AppInfo) GetClients() []*AppClientInfo {
	if m != nil {
		return m.Clients
	}
	return nil
}

func (m *AppInfo) GetSettings() *dto.AppSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *AppInfo) GetCreateUserId() string {
	if m != nil {
		return m.CreateUserId
	}
	return ""
}

type AppClientInfo struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Enabled              int64    `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Platform             int64    `protobuf:"varint,3,opt,name=platform,proto3" json:"platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppClientInfo) Reset()         { *m = AppClientInfo{} }
func (m *AppClientInfo) String() string { return proto.CompactTextString(m) }
func (*AppClientInfo) ProtoMessage()    {}
func (*AppClientInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b974f2f4f838fbc1, []int{8}
}

func (m *AppClientInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppClientInfo.Unmarshal(m, b)
}
func (m *AppClientInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppClientInfo.Marshal(b, m, deterministic)
}
func (m *AppClientInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppClientInfo.Merge(m, src)
}
func (m *AppClientInfo) XXX_Size() int {
	return xxx_messageInfo_AppClientInfo.Size(m)
}
func (m *AppClientInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AppClientInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AppClientInfo proto.InternalMessageInfo

func (m *AppClientInfo) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *AppClientInfo) GetEnabled() int64 {
	if m != nil {
		return m.Enabled
	}
	return 0
}

func (m *AppClientInfo) GetPlatform() int64 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func init() {
	proto.RegisterType((*SwitchRequest)(nil), "gs.service.application.SwitchRequest")
	proto.RegisterType((*FindRequest)(nil), "gs.service.application.FindRequest")
	proto.RegisterType((*ListResponse)(nil), "gs.service.application.ListResponse")
	proto.RegisterType((*SimpleApplicationResponse)(nil), "gs.service.application.SimpleApplicationResponse")
	proto.RegisterType((*CreateRequest)(nil), "gs.service.application.CreateRequest")
	proto.RegisterType((*RemoveRequest)(nil), "gs.service.application.RemoveRequest")
	proto.RegisterType((*ChangeNameRequest)(nil), "gs.service.application.ChangeNameRequest")
	proto.RegisterType((*AppInfo)(nil), "gs.service.application.AppInfo")
	proto.RegisterType((*AppClientInfo)(nil), "gs.service.application.AppClientInfo")
}

func init() { proto.RegisterFile("application/pb/application.proto", fileDescriptor_b974f2f4f838fbc1) }

var fileDescriptor_b974f2f4f838fbc1 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdd, 0x8a, 0x13, 0x31,
	0x14, 0xa6, 0xff, 0xdd, 0xd3, 0xad, 0x60, 0xd0, 0x65, 0xac, 0x17, 0x96, 0xd1, 0xc5, 0xca, 0xc2,
	0x14, 0xbb, 0x17, 0x7b, 0x25, 0x4b, 0x2d, 0xb8, 0x14, 0x44, 0x31, 0xc5, 0x07, 0x48, 0x67, 0x4e,
	0xdb, 0xd0, 0x4e, 0x12, 0x9b, 0xec, 0x8a, 0x17, 0x3e, 0x89, 0xcf, 0xe6, 0xbb, 0xc8, 0x24, 0x9d,
	0xb6, 0x53, 0x37, 0x4b, 0x41, 0xbc, 0xcb, 0x49, 0xbe, 0x7c, 0xe7, 0xe4, 0x3b, 0x5f, 0x0e, 0x74,
	0x99, 0x52, 0x2b, 0x1e, 0x33, 0xc3, 0xa5, 0xe8, 0xab, 0x69, 0x7f, 0x2f, 0x8c, 0xd4, 0x5a, 0x1a,
	0x49, 0xce, 0xe6, 0x3a, 0xd2, 0xb8, 0xbe, 0xe3, 0x31, 0x46, 0x7b, 0xa7, 0x9d, 0x8b, 0xa5, 0x14,
	0xb8, 0x5c, 0xca, 0x28, 0xc5, 0xfe, 0x5c, 0xea, 0xec, 0x7e, 0x2c, 0xd3, 0x54, 0x0a, 0xdd, 0x4f,
	0x8c, 0xcc, 0xd7, 0x8e, 0x24, 0xbc, 0x86, 0xf6, 0xe4, 0x3b, 0x37, 0xf1, 0x82, 0xe2, 0xb7, 0x5b,
	0xd4, 0x86, 0x3c, 0x81, 0x1a, 0x53, 0x6a, 0x9c, 0x04, 0xa5, 0x6e, 0xa9, 0x77, 0x42, 0x5d, 0x40,
	0x02, 0x68, 0xa0, 0x60, 0xd3, 0x15, 0x26, 0x41, 0xb9, 0x5b, 0xea, 0x35, 0x69, 0x1e, 0x86, 0xaf,
	0xa1, 0xf5, 0x81, 0x8b, 0x24, 0xbf, 0x1e, 0x40, 0x23, 0x96, 0xc2, 0xa0, 0x30, 0x1b, 0x82, 0x3c,
	0x0c, 0x15, 0x9c, 0x7e, 0xe4, 0xda, 0x50, 0xd4, 0x4a, 0x0a, 0x8d, 0xe4, 0x02, 0x6a, 0xda, 0x30,
	0x83, 0x16, 0xd7, 0x1a, 0x3c, 0x8d, 0xe6, 0x3a, 0xca, 0x6b, 0x4b, 0x8c, 0x8c, 0x26, 0xd9, 0x21,
	0x75, 0x18, 0x72, 0x09, 0x55, 0x2e, 0x66, 0x32, 0x28, 0x77, 0x2b, 0xbd, 0xd6, 0xe0, 0x45, 0x74,
	0xff, 0xd3, 0xa3, 0xa1, 0x52, 0x63, 0x31, 0x93, 0xd4, 0x82, 0xc3, 0x9f, 0xf0, 0x6c, 0xc2, 0x53,
	0xb5, 0xc2, 0xe1, 0x0e, 0xf2, 0xaf, 0xe9, 0x4b, 0xc7, 0xa7, 0xbf, 0x82, 0xf6, 0x68, 0x8d, 0x19,
	0xcb, 0x46, 0x1b, 0x02, 0x55, 0xc1, 0x52, 0xdc, 0x08, 0x63, 0xd7, 0xd9, 0x5e, 0x82, 0x3a, 0xb6,
	0xcc, 0x27, 0xd4, 0xae, 0xc3, 0x73, 0x68, 0x53, 0x4c, 0xe5, 0x1d, 0x3e, 0xd8, 0x93, 0xf0, 0x1d,
	0x3c, 0x1e, 0x2d, 0x98, 0x98, 0xe3, 0x27, 0x96, 0xfe, 0x0d, 0x2d, 0xef, 0xb7, 0x2f, 0xcf, 0x5c,
	0xd9, 0x65, 0x0e, 0x7f, 0x95, 0xa1, 0xb1, 0x29, 0xf8, 0xd8, 0xca, 0x76, 0xec, 0x15, 0x8f, 0x39,
	0xaa, 0xdd, 0x52, 0xaf, 0xb2, 0x35, 0x07, 0xe9, 0x40, 0x33, 0xb6, 0x12, 0x0c, 0x4d, 0x50, 0xb3,
	0x47, 0xdb, 0x98, 0x5c, 0x43, 0x23, 0x5e, 0x71, 0x14, 0x46, 0x07, 0x75, 0xdb, 0xd5, 0xf3, 0x07,
	0x64, 0x1d, 0x59, 0xa4, 0x15, 0x37, 0xbf, 0x45, 0xae, 0xa0, 0xa9, 0xd1, 0x18, 0x2e, 0xe6, 0x3a,
	0x68, 0xd8, 0xc6, 0x3c, 0x3f, 0x6c, 0xe2, 0x50, 0xa9, 0xc9, 0x06, 0x42, 0xb7, 0x60, 0x12, 0xc2,
	0xa9, 0xab, 0xe2, 0xab, 0xc6, 0xf5, 0x38, 0x09, 0x9a, 0xf6, 0x31, 0x85, 0xbd, 0x90, 0x41, 0xbb,
	0x90, 0xd6, 0x3e, 0xc5, 0x45, 0x79, 0x1b, 0xb6, 0xf1, 0xe1, 0xef, 0x28, 0x0a, 0xa0, 0x56, 0xcc,
	0xcc, 0xe4, 0x3a, 0xb5, 0x9a, 0x55, 0xe8, 0x36, 0x1e, 0xfc, 0xae, 0x42, 0x6b, 0xcf, 0x99, 0xe4,
	0x06, 0xea, 0xce, 0x2f, 0xc4, 0xab, 0x44, 0xc1, 0x4f, 0x9d, 0xb3, 0xfb, 0x3c, 0x7b, 0xab, 0x33,
	0x22, 0xe7, 0x1f, 0x3f, 0x51, 0xc1, 0x5f, 0x5e, 0xa2, 0xcf, 0x00, 0x3b, 0x87, 0x91, 0x37, 0xde,
	0xaa, 0x0e, 0x5d, 0xe8, 0x25, 0x44, 0x37, 0x2c, 0xde, 0xff, 0x18, 0x5a, 0xe3, 0xbc, 0xf4, 0x31,
	0xee, 0x4d, 0x94, 0xce, 0x5b, 0x1f, 0xc8, 0xff, 0xb7, 0x17, 0xf0, 0xc8, 0xa5, 0x19, 0xe5, 0x1d,
	0xfa, 0x5f, 0x99, 0xbe, 0x40, 0x35, 0x1b, 0x6a, 0xc7, 0xf1, 0xbf, 0xf2, 0x81, 0x0a, 0x73, 0xf1,
	0x06, 0xea, 0x6e, 0x22, 0xfb, 0xbb, 0x57, 0x98, 0xd8, 0x3e, 0xb1, 0xa7, 0x75, 0x3b, 0xe1, 0x2f,
	0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xfe, 0x1b, 0xf6, 0x4a, 0x06, 0x00, 0x00,
}
