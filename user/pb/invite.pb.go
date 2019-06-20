// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/pb/invite.proto

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

type AppendRequest struct {
	UserId               string      `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Item                 *InviteItem `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AppendRequest) Reset()         { *m = AppendRequest{} }
func (m *AppendRequest) String() string { return proto.CompactTextString(m) }
func (*AppendRequest) ProtoMessage()    {}
func (*AppendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3caf1176bf5393cc, []int{0}
}

func (m *AppendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendRequest.Unmarshal(m, b)
}
func (m *AppendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendRequest.Marshal(b, m, deterministic)
}
func (m *AppendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendRequest.Merge(m, src)
}
func (m *AppendRequest) XXX_Size() int {
	return xxx_messageInfo_AppendRequest.Size(m)
}
func (m *AppendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppendRequest proto.InternalMessageInfo

func (m *AppendRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AppendRequest) GetItem() *InviteItem {
	if m != nil {
		return m.Item
	}
	return nil
}

type SetStateRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	AppId                string   `protobuf:"bytes,2,opt,name=appId,proto3" json:"appId,omitempty"`
	State                int64    `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetStateRequest) Reset()         { *m = SetStateRequest{} }
func (m *SetStateRequest) String() string { return proto.CompactTextString(m) }
func (*SetStateRequest) ProtoMessage()    {}
func (*SetStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3caf1176bf5393cc, []int{1}
}

func (m *SetStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetStateRequest.Unmarshal(m, b)
}
func (m *SetStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetStateRequest.Marshal(b, m, deterministic)
}
func (m *SetStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetStateRequest.Merge(m, src)
}
func (m *SetStateRequest) XXX_Size() int {
	return xxx_messageInfo_SetStateRequest.Size(m)
}
func (m *SetStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetStateRequest proto.InternalMessageInfo

func (m *SetStateRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SetStateRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *SetStateRequest) GetState() int64 {
	if m != nil {
		return m.State
	}
	return 0
}

type HasInvitedRequest struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	UserId               string   `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	AppId                string   `protobuf:"bytes,4,opt,name=appId,proto3" json:"appId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HasInvitedRequest) Reset()         { *m = HasInvitedRequest{} }
func (m *HasInvitedRequest) String() string { return proto.CompactTextString(m) }
func (*HasInvitedRequest) ProtoMessage()    {}
func (*HasInvitedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3caf1176bf5393cc, []int{2}
}

func (m *HasInvitedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasInvitedRequest.Unmarshal(m, b)
}
func (m *HasInvitedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasInvitedRequest.Marshal(b, m, deterministic)
}
func (m *HasInvitedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasInvitedRequest.Merge(m, src)
}
func (m *HasInvitedRequest) XXX_Size() int {
	return xxx_messageInfo_HasInvitedRequest.Size(m)
}
func (m *HasInvitedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HasInvitedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HasInvitedRequest proto.InternalMessageInfo

func (m *HasInvitedRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *HasInvitedRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *HasInvitedRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *HasInvitedRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type HasInvitedResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	UserId               string     `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Status               int64      `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HasInvitedResponse) Reset()         { *m = HasInvitedResponse{} }
func (m *HasInvitedResponse) String() string { return proto.CompactTextString(m) }
func (*HasInvitedResponse) ProtoMessage()    {}
func (*HasInvitedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3caf1176bf5393cc, []int{3}
}

func (m *HasInvitedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasInvitedResponse.Unmarshal(m, b)
}
func (m *HasInvitedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasInvitedResponse.Marshal(b, m, deterministic)
}
func (m *HasInvitedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasInvitedResponse.Merge(m, src)
}
func (m *HasInvitedResponse) XXX_Size() int {
	return xxx_messageInfo_HasInvitedResponse.Size(m)
}
func (m *HasInvitedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HasInvitedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HasInvitedResponse proto.InternalMessageInfo

func (m *HasInvitedResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *HasInvitedResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *HasInvitedResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

type InviteSearchRequest struct {
	PageIndex            int64    `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	PageSize             int64    `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InviteSearchRequest) Reset()         { *m = InviteSearchRequest{} }
func (m *InviteSearchRequest) String() string { return proto.CompactTextString(m) }
func (*InviteSearchRequest) ProtoMessage()    {}
func (*InviteSearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3caf1176bf5393cc, []int{4}
}

func (m *InviteSearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InviteSearchRequest.Unmarshal(m, b)
}
func (m *InviteSearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InviteSearchRequest.Marshal(b, m, deterministic)
}
func (m *InviteSearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InviteSearchRequest.Merge(m, src)
}
func (m *InviteSearchRequest) XXX_Size() int {
	return xxx_messageInfo_InviteSearchRequest.Size(m)
}
func (m *InviteSearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InviteSearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InviteSearchRequest proto.InternalMessageInfo

func (m *InviteSearchRequest) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *InviteSearchRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *InviteSearchRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *InviteSearchRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type InviteSearchResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *InviteSearchResponse) Reset()         { *m = InviteSearchResponse{} }
func (m *InviteSearchResponse) String() string { return proto.CompactTextString(m) }
func (*InviteSearchResponse) ProtoMessage()    {}
func (*InviteSearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3caf1176bf5393cc, []int{5}
}

func (m *InviteSearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InviteSearchResponse.Unmarshal(m, b)
}
func (m *InviteSearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InviteSearchResponse.Marshal(b, m, deterministic)
}
func (m *InviteSearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InviteSearchResponse.Merge(m, src)
}
func (m *InviteSearchResponse) XXX_Size() int {
	return xxx_messageInfo_InviteSearchResponse.Size(m)
}
func (m *InviteSearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InviteSearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InviteSearchResponse proto.InternalMessageInfo

func (m *InviteSearchResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

type GetDetailResponse struct {
	State                *dto.State    `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Items                []*InviteItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetDetailResponse) Reset()         { *m = GetDetailResponse{} }
func (m *GetDetailResponse) String() string { return proto.CompactTextString(m) }
func (*GetDetailResponse) ProtoMessage()    {}
func (*GetDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3caf1176bf5393cc, []int{6}
}

func (m *GetDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDetailResponse.Unmarshal(m, b)
}
func (m *GetDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDetailResponse.Marshal(b, m, deterministic)
}
func (m *GetDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDetailResponse.Merge(m, src)
}
func (m *GetDetailResponse) XXX_Size() int {
	return xxx_messageInfo_GetDetailResponse.Size(m)
}
func (m *GetDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDetailResponse proto.InternalMessageInfo

func (m *GetDetailResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *GetDetailResponse) GetItems() []*InviteItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type InviteDetail struct {
	Phone                string        `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                string        `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username             string        `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	RealName             string        `protobuf:"bytes,4,opt,name=realName,proto3" json:"realName,omitempty"`
	Items                []*InviteItem `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
	ExpiredAt            int64         `protobuf:"varint,6,opt,name=expiredAt,proto3" json:"expiredAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *InviteDetail) Reset()         { *m = InviteDetail{} }
func (m *InviteDetail) String() string { return proto.CompactTextString(m) }
func (*InviteDetail) ProtoMessage()    {}
func (*InviteDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_3caf1176bf5393cc, []int{7}
}

func (m *InviteDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InviteDetail.Unmarshal(m, b)
}
func (m *InviteDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InviteDetail.Marshal(b, m, deterministic)
}
func (m *InviteDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InviteDetail.Merge(m, src)
}
func (m *InviteDetail) XXX_Size() int {
	return xxx_messageInfo_InviteDetail.Size(m)
}
func (m *InviteDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_InviteDetail.DiscardUnknown(m)
}

var xxx_messageInfo_InviteDetail proto.InternalMessageInfo

func (m *InviteDetail) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *InviteDetail) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *InviteDetail) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *InviteDetail) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *InviteDetail) GetItems() []*InviteItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *InviteDetail) GetExpiredAt() int64 {
	if m != nil {
		return m.ExpiredAt
	}
	return 0
}

type InviteItem struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	Roles                []string `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	BindGroupId          string   `protobuf:"bytes,3,opt,name=bindGroupId,proto3" json:"bindGroupId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InviteItem) Reset()         { *m = InviteItem{} }
func (m *InviteItem) String() string { return proto.CompactTextString(m) }
func (*InviteItem) ProtoMessage()    {}
func (*InviteItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_3caf1176bf5393cc, []int{8}
}

func (m *InviteItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InviteItem.Unmarshal(m, b)
}
func (m *InviteItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InviteItem.Marshal(b, m, deterministic)
}
func (m *InviteItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InviteItem.Merge(m, src)
}
func (m *InviteItem) XXX_Size() int {
	return xxx_messageInfo_InviteItem.Size(m)
}
func (m *InviteItem) XXX_DiscardUnknown() {
	xxx_messageInfo_InviteItem.DiscardUnknown(m)
}

var xxx_messageInfo_InviteItem proto.InternalMessageInfo

func (m *InviteItem) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *InviteItem) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *InviteItem) GetBindGroupId() string {
	if m != nil {
		return m.BindGroupId
	}
	return ""
}

func init() {
	proto.RegisterType((*AppendRequest)(nil), "gosionsvc.external.user.AppendRequest")
	proto.RegisterType((*SetStateRequest)(nil), "gosionsvc.external.user.SetStateRequest")
	proto.RegisterType((*HasInvitedRequest)(nil), "gosionsvc.external.user.HasInvitedRequest")
	proto.RegisterType((*HasInvitedResponse)(nil), "gosionsvc.external.user.HasInvitedResponse")
	proto.RegisterType((*InviteSearchRequest)(nil), "gosionsvc.external.user.InviteSearchRequest")
	proto.RegisterType((*InviteSearchResponse)(nil), "gosionsvc.external.user.InviteSearchResponse")
	proto.RegisterType((*GetDetailResponse)(nil), "gosionsvc.external.user.GetDetailResponse")
	proto.RegisterType((*InviteDetail)(nil), "gosionsvc.external.user.InviteDetail")
	proto.RegisterType((*InviteItem)(nil), "gosionsvc.external.user.InviteItem")
}

func init() { proto.RegisterFile("user/pb/invite.proto", fileDescriptor_3caf1176bf5393cc) }

var fileDescriptor_3caf1176bf5393cc = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0x96, 0xe3, 0xc4, 0x4a, 0xa6, 0xff, 0x2f, 0xe8, 0x12, 0x42, 0x64, 0x71, 0x88, 0x8c, 0x40,
	0x11, 0x01, 0x47, 0x2a, 0x07, 0xc4, 0xb1, 0x02, 0xd4, 0x46, 0x42, 0x1c, 0x1c, 0xf5, 0xc2, 0x89,
	0x4d, 0x3c, 0x4a, 0xad, 0xd8, 0xbb, 0x5b, 0xef, 0x3a, 0x04, 0x78, 0x13, 0x5e, 0x89, 0x97, 0x42,
	0xeb, 0xb5, 0x9d, 0xa4, 0xad, 0x95, 0xa4, 0x37, 0x7f, 0xb3, 0x33, 0xf3, 0x7d, 0x33, 0x3b, 0x3b,
	0x86, 0x6e, 0x26, 0x31, 0x1d, 0x8b, 0xd9, 0x38, 0x62, 0xab, 0x48, 0xa1, 0x2f, 0x52, 0xae, 0x38,
	0x79, 0xb6, 0xe0, 0x32, 0xe2, 0x4c, 0xae, 0xe6, 0x3e, 0xae, 0x15, 0xa6, 0x8c, 0xc6, 0xbe, 0x76,
	0x74, 0x47, 0x4b, 0xce, 0x70, 0xb9, 0xe4, 0x7e, 0x82, 0x63, 0xe3, 0x33, 0x9e, 0xf3, 0x24, 0xe1,
	0x4c, 0x8e, 0x43, 0xc5, 0xcb, 0x6f, 0x93, 0xc5, 0xfb, 0x0e, 0xff, 0x9f, 0x0b, 0x81, 0x2c, 0x0c,
	0xf0, 0x26, 0x43, 0xa9, 0x48, 0x0f, 0x1c, 0x9d, 0x65, 0x12, 0xf6, 0xad, 0x81, 0x35, 0xec, 0x04,
	0x05, 0x22, 0xef, 0xa1, 0x19, 0x29, 0x4c, 0xfa, 0x8d, 0x81, 0x35, 0x3c, 0x39, 0x7b, 0xe1, 0xd7,
	0xb0, 0xfb, 0x93, 0x5c, 0xe3, 0x44, 0x61, 0x12, 0xe4, 0x01, 0xde, 0x15, 0x3c, 0x9a, 0xa2, 0x9a,
	0x2a, 0xaa, 0x70, 0x1f, 0x47, 0x17, 0x5a, 0x54, 0x88, 0x49, 0x98, 0x93, 0x74, 0x02, 0x03, 0xb4,
	0x55, 0xea, 0xe8, 0xbe, 0x3d, 0xb0, 0x86, 0x76, 0x60, 0x80, 0x97, 0xc0, 0xe9, 0x25, 0x95, 0x86,
	0xad, 0x12, 0xdf, 0x85, 0x96, 0xb8, 0xe6, 0x0c, 0x8b, 0xbc, 0x06, 0x68, 0x2b, 0x26, 0x34, 0x8a,
	0xcb, 0xb4, 0x39, 0xd8, 0x12, 0x61, 0xdf, 0x2f, 0xa2, 0xb9, 0x25, 0xc2, 0xbb, 0x01, 0xb2, 0x4d,
	0x27, 0x05, 0x67, 0x12, 0xc9, 0xa8, 0x94, 0x66, 0xe5, 0x5d, 0x79, 0xea, 0x2f, 0xa4, 0x5f, 0xf6,
	0x37, 0x54, 0xdc, 0x37, 0x55, 0x1b, 0x9f, 0x2d, 0xc2, 0xc6, 0x0e, 0x61, 0x0f, 0x1c, 0xed, 0x90,
	0xc9, 0xa2, 0xc0, 0x02, 0x79, 0x3f, 0xe0, 0x89, 0xe1, 0x9b, 0x22, 0x4d, 0xe7, 0xd7, 0x65, 0x8d,
	0xcf, 0xa1, 0x23, 0xe8, 0x02, 0x27, 0x2c, 0xc4, 0x75, 0xce, 0x6b, 0x07, 0x1b, 0x03, 0x71, 0xa1,
	0xad, 0xc1, 0x34, 0xfa, 0x85, 0x39, 0x8d, 0x1d, 0x54, 0x98, 0x3c, 0x06, 0x7b, 0x89, 0x3f, 0x8b,
	0x72, 0xf5, 0xa7, 0xae, 0x75, 0x45, 0xe3, 0x0c, 0xcb, 0x5a, 0x73, 0xe0, 0x7d, 0x84, 0xee, 0x2e,
	0xf1, 0x03, 0xaa, 0xf5, 0x7e, 0xc3, 0xe9, 0x05, 0xaa, 0x4f, 0xa8, 0x68, 0x14, 0x3f, 0xac, 0x5f,
	0x1f, 0xa0, 0xa5, 0x07, 0x48, 0xf6, 0x1b, 0x03, 0xfb, 0xd0, 0x91, 0x33, 0x11, 0xde, 0x5f, 0x0b,
	0xfe, 0x33, 0x56, 0x23, 0xe0, 0xa8, 0xc1, 0x70, 0xa1, 0xad, 0xd3, 0x32, 0x9a, 0x60, 0xd1, 0xab,
	0x0a, 0xeb, 0xb3, 0x14, 0x69, 0xfc, 0x55, 0x9f, 0x99, 0x9e, 0x55, 0x78, 0xa3, 0xb7, 0x75, 0xac,
	0x5e, 0x7d, 0xa7, 0xb8, 0x16, 0x51, 0x8a, 0xe1, 0xb9, 0xea, 0x3b, 0xe6, 0x4e, 0x2b, 0x83, 0xf7,
	0x0d, 0x60, 0x13, 0xb2, 0x99, 0x4f, 0xeb, 0xd6, 0x23, 0x49, 0x79, 0x8c, 0xa6, 0x59, 0x9d, 0xc0,
	0x00, 0x32, 0x80, 0x93, 0x59, 0xc4, 0xc2, 0x8b, 0x94, 0x67, 0xa2, 0x1a, 0xf4, 0x6d, 0xd3, 0xd9,
	0x9f, 0x26, 0x38, 0x26, 0x39, 0xf9, 0x0c, 0xcd, 0x2b, 0x89, 0x29, 0x79, 0xb9, 0x47, 0xb8, 0x69,
	0xa9, 0xdb, 0xbb, 0xef, 0xf2, 0x32, 0x49, 0x10, 0x1c, 0x33, 0x37, 0xe4, 0xcd, 0x9e, 0x44, 0x3b,
	0x73, 0xed, 0xbe, 0x3d, 0xd0, 0xbb, 0x18, 0x25, 0x04, 0xd8, 0x3c, 0x48, 0xf2, 0xba, 0x36, 0xf8,
	0xce, 0x92, 0x70, 0x47, 0x07, 0xf9, 0x16, 0x34, 0x73, 0xe8, 0x54, 0x63, 0x7c, 0x14, 0x4b, 0xbd,
	0xef, 0xdd, 0x67, 0xf1, 0x05, 0xda, 0xe5, 0x8a, 0x24, 0xc3, 0xda, 0xb8, 0x5b, 0x5b, 0xb4, 0xf6,
	0x02, 0x2e, 0xc1, 0x31, 0x2b, 0x9d, 0xbc, 0xaa, 0xcd, 0xb5, 0xb3, 0xf3, 0xeb, 0x32, 0xcd, 0x9c,
	0xfc, 0x1f, 0xf1, 0xee, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x9f, 0x12, 0xd3, 0x81, 0x06,
	0x00, 0x00,
}
