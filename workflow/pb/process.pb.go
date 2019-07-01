// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/process.proto

package xbasissvc_external_workflow

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	dto "konekko.me/xbasis/commons/dto"
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

type SearchProcessRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Page                 int64    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int64    `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchProcessRequest) Reset()         { *m = SearchProcessRequest{} }
func (m *SearchProcessRequest) String() string { return proto.CompactTextString(m) }
func (*SearchProcessRequest) ProtoMessage()    {}
func (*SearchProcessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{0}
}

func (m *SearchProcessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchProcessRequest.Unmarshal(m, b)
}
func (m *SearchProcessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchProcessRequest.Marshal(b, m, deterministic)
}
func (m *SearchProcessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchProcessRequest.Merge(m, src)
}
func (m *SearchProcessRequest) XXX_Size() int {
	return xxx_messageInfo_SearchProcessRequest.Size(m)
}
func (m *SearchProcessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchProcessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchProcessRequest proto.InternalMessageInfo

func (m *SearchProcessRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *SearchProcessRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SearchProcessRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchProcessRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type SearchProcessResponse struct {
	State                *dto.State           `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Data                 []*SearchProcessItem `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SearchProcessResponse) Reset()         { *m = SearchProcessResponse{} }
func (m *SearchProcessResponse) String() string { return proto.CompactTextString(m) }
func (*SearchProcessResponse) ProtoMessage()    {}
func (*SearchProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{1}
}

func (m *SearchProcessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchProcessResponse.Unmarshal(m, b)
}
func (m *SearchProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchProcessResponse.Marshal(b, m, deterministic)
}
func (m *SearchProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchProcessResponse.Merge(m, src)
}
func (m *SearchProcessResponse) XXX_Size() int {
	return xxx_messageInfo_SearchProcessResponse.Size(m)
}
func (m *SearchProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchProcessResponse proto.InternalMessageInfo

func (m *SearchProcessResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *SearchProcessResponse) GetData() []*SearchProcessItem {
	if m != nil {
		return m.Data
	}
	return nil
}

type SearchProcessItem struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AppId                string   `protobuf:"bytes,2,opt,name=appId,proto3" json:"appId,omitempty"`
	ProcessId            string   `protobuf:"bytes,3,opt,name=processId,proto3" json:"processId,omitempty"`
	UpdateAt             string   `protobuf:"bytes,4,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
	Desc                 string   `protobuf:"bytes,6,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchProcessItem) Reset()         { *m = SearchProcessItem{} }
func (m *SearchProcessItem) String() string { return proto.CompactTextString(m) }
func (*SearchProcessItem) ProtoMessage()    {}
func (*SearchProcessItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{2}
}

func (m *SearchProcessItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchProcessItem.Unmarshal(m, b)
}
func (m *SearchProcessItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchProcessItem.Marshal(b, m, deterministic)
}
func (m *SearchProcessItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchProcessItem.Merge(m, src)
}
func (m *SearchProcessItem) XXX_Size() int {
	return xxx_messageInfo_SearchProcessItem.Size(m)
}
func (m *SearchProcessItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchProcessItem.DiscardUnknown(m)
}

var xxx_messageInfo_SearchProcessItem proto.InternalMessageInfo

func (m *SearchProcessItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SearchProcessItem) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *SearchProcessItem) GetProcessId() string {
	if m != nil {
		return m.ProcessId
	}
	return ""
}

func (m *SearchProcessItem) GetUpdateAt() string {
	if m != nil {
		return m.UpdateAt
	}
	return ""
}

func (m *SearchProcessItem) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type DetailProcessRequest struct {
	ProcessId            string   `protobuf:"bytes,1,opt,name=processId,proto3" json:"processId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailProcessRequest) Reset()         { *m = DetailProcessRequest{} }
func (m *DetailProcessRequest) String() string { return proto.CompactTextString(m) }
func (*DetailProcessRequest) ProtoMessage()    {}
func (*DetailProcessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{3}
}

func (m *DetailProcessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailProcessRequest.Unmarshal(m, b)
}
func (m *DetailProcessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailProcessRequest.Marshal(b, m, deterministic)
}
func (m *DetailProcessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailProcessRequest.Merge(m, src)
}
func (m *DetailProcessRequest) XXX_Size() int {
	return xxx_messageInfo_DetailProcessRequest.Size(m)
}
func (m *DetailProcessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailProcessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetailProcessRequest proto.InternalMessageInfo

func (m *DetailProcessRequest) GetProcessId() string {
	if m != nil {
		return m.ProcessId
	}
	return ""
}

type DetailProcessResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	NodeDataArray        string     `protobuf:"bytes,2,opt,name=nodeDataArray,proto3" json:"nodeDataArray,omitempty"`
	LinkDataArray        string     `protobuf:"bytes,3,opt,name=linkDataArray,proto3" json:"linkDataArray,omitempty"`
	Desc                 string     `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	Name                 string     `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DetailProcessResponse) Reset()         { *m = DetailProcessResponse{} }
func (m *DetailProcessResponse) String() string { return proto.CompactTextString(m) }
func (*DetailProcessResponse) ProtoMessage()    {}
func (*DetailProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{4}
}

func (m *DetailProcessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailProcessResponse.Unmarshal(m, b)
}
func (m *DetailProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailProcessResponse.Marshal(b, m, deterministic)
}
func (m *DetailProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailProcessResponse.Merge(m, src)
}
func (m *DetailProcessResponse) XXX_Size() int {
	return xxx_messageInfo_DetailProcessResponse.Size(m)
}
func (m *DetailProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DetailProcessResponse proto.InternalMessageInfo

func (m *DetailProcessResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *DetailProcessResponse) GetNodeDataArray() string {
	if m != nil {
		return m.NodeDataArray
	}
	return ""
}

func (m *DetailProcessResponse) GetLinkDataArray() string {
	if m != nil {
		return m.LinkDataArray
	}
	return ""
}

func (m *DetailProcessResponse) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *DetailProcessResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetImageRequest struct {
	ProcessId            string   `protobuf:"bytes,1,opt,name=processId,proto3" json:"processId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetImageRequest) Reset()         { *m = GetImageRequest{} }
func (m *GetImageRequest) String() string { return proto.CompactTextString(m) }
func (*GetImageRequest) ProtoMessage()    {}
func (*GetImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{5}
}

func (m *GetImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetImageRequest.Unmarshal(m, b)
}
func (m *GetImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetImageRequest.Marshal(b, m, deterministic)
}
func (m *GetImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetImageRequest.Merge(m, src)
}
func (m *GetImageRequest) XXX_Size() int {
	return xxx_messageInfo_GetImageRequest.Size(m)
}
func (m *GetImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetImageRequest proto.InternalMessageInfo

func (m *GetImageRequest) GetProcessId() string {
	if m != nil {
		return m.ProcessId
	}
	return ""
}

type GetImageResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Image                string     `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetImageResponse) Reset()         { *m = GetImageResponse{} }
func (m *GetImageResponse) String() string { return proto.CompactTextString(m) }
func (*GetImageResponse) ProtoMessage()    {}
func (*GetImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{6}
}

func (m *GetImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetImageResponse.Unmarshal(m, b)
}
func (m *GetImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetImageResponse.Marshal(b, m, deterministic)
}
func (m *GetImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetImageResponse.Merge(m, src)
}
func (m *GetImageResponse) XXX_Size() int {
	return xxx_messageInfo_GetImageResponse.Size(m)
}
func (m *GetImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetImageResponse proto.InternalMessageInfo

func (m *GetImageResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *GetImageResponse) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type CreateRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc                 string   `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	AppId                string   `protobuf:"bytes,4,opt,name=appId,proto3" json:"appId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{7}
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

func (m *CreateRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

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

func (m *CreateRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type CreateResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	ProcessId            string     `protobuf:"bytes,2,opt,name=processId,proto3" json:"processId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{8}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *CreateResponse) GetProcessId() string {
	if m != nil {
		return m.ProcessId
	}
	return ""
}

type BuildRequest struct {
	ProcessId            string   `protobuf:"bytes,1,opt,name=processId,proto3" json:"processId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildRequest) Reset()         { *m = BuildRequest{} }
func (m *BuildRequest) String() string { return proto.CompactTextString(m) }
func (*BuildRequest) ProtoMessage()    {}
func (*BuildRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{9}
}

func (m *BuildRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildRequest.Unmarshal(m, b)
}
func (m *BuildRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildRequest.Marshal(b, m, deterministic)
}
func (m *BuildRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildRequest.Merge(m, src)
}
func (m *BuildRequest) XXX_Size() int {
	return xxx_messageInfo_BuildRequest.Size(m)
}
func (m *BuildRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuildRequest proto.InternalMessageInfo

func (m *BuildRequest) GetProcessId() string {
	if m != nil {
		return m.ProcessId
	}
	return ""
}

type BuildResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BuildResponse) Reset()         { *m = BuildResponse{} }
func (m *BuildResponse) String() string { return proto.CompactTextString(m) }
func (*BuildResponse) ProtoMessage()    {}
func (*BuildResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{10}
}

func (m *BuildResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildResponse.Unmarshal(m, b)
}
func (m *BuildResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildResponse.Marshal(b, m, deterministic)
}
func (m *BuildResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildResponse.Merge(m, src)
}
func (m *BuildResponse) XXX_Size() int {
	return xxx_messageInfo_BuildResponse.Size(m)
}
func (m *BuildResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BuildResponse proto.InternalMessageInfo

func (m *BuildResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

type DeleteProcessRequest struct {
	ProcessId            string   `protobuf:"bytes,1,opt,name=processId,proto3" json:"processId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProcessRequest) Reset()         { *m = DeleteProcessRequest{} }
func (m *DeleteProcessRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteProcessRequest) ProtoMessage()    {}
func (*DeleteProcessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{11}
}

func (m *DeleteProcessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProcessRequest.Unmarshal(m, b)
}
func (m *DeleteProcessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProcessRequest.Marshal(b, m, deterministic)
}
func (m *DeleteProcessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProcessRequest.Merge(m, src)
}
func (m *DeleteProcessRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteProcessRequest.Size(m)
}
func (m *DeleteProcessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProcessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProcessRequest proto.InternalMessageInfo

func (m *DeleteProcessRequest) GetProcessId() string {
	if m != nil {
		return m.ProcessId
	}
	return ""
}

type DeleteProcessResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DeleteProcessResponse) Reset()         { *m = DeleteProcessResponse{} }
func (m *DeleteProcessResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteProcessResponse) ProtoMessage()    {}
func (*DeleteProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{12}
}

func (m *DeleteProcessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProcessResponse.Unmarshal(m, b)
}
func (m *DeleteProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProcessResponse.Marshal(b, m, deterministic)
}
func (m *DeleteProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProcessResponse.Merge(m, src)
}
func (m *DeleteProcessResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteProcessResponse.Size(m)
}
func (m *DeleteProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProcessResponse proto.InternalMessageInfo

func (m *DeleteProcessResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

type UpdateProcessRequest struct {
	ProcessId            string   `protobuf:"bytes,1,opt,name=processId,proto3" json:"processId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc                 string   `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	NodeDataArray        string   `protobuf:"bytes,4,opt,name=nodeDataArray,proto3" json:"nodeDataArray,omitempty"`
	LinkDataArray        string   `protobuf:"bytes,5,opt,name=linkDataArray,proto3" json:"linkDataArray,omitempty"`
	AppId                string   `protobuf:"bytes,6,opt,name=appId,proto3" json:"appId,omitempty"`
	Image                string   `protobuf:"bytes,7,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateProcessRequest) Reset()         { *m = UpdateProcessRequest{} }
func (m *UpdateProcessRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateProcessRequest) ProtoMessage()    {}
func (*UpdateProcessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{13}
}

func (m *UpdateProcessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProcessRequest.Unmarshal(m, b)
}
func (m *UpdateProcessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProcessRequest.Marshal(b, m, deterministic)
}
func (m *UpdateProcessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProcessRequest.Merge(m, src)
}
func (m *UpdateProcessRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateProcessRequest.Size(m)
}
func (m *UpdateProcessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProcessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProcessRequest proto.InternalMessageInfo

func (m *UpdateProcessRequest) GetProcessId() string {
	if m != nil {
		return m.ProcessId
	}
	return ""
}

func (m *UpdateProcessRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateProcessRequest) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *UpdateProcessRequest) GetNodeDataArray() string {
	if m != nil {
		return m.NodeDataArray
	}
	return ""
}

func (m *UpdateProcessRequest) GetLinkDataArray() string {
	if m != nil {
		return m.LinkDataArray
	}
	return ""
}

func (m *UpdateProcessRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *UpdateProcessRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type UpdateProcessResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	ProcessId            string     `protobuf:"bytes,2,opt,name=processId,proto3" json:"processId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateProcessResponse) Reset()         { *m = UpdateProcessResponse{} }
func (m *UpdateProcessResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateProcessResponse) ProtoMessage()    {}
func (*UpdateProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{14}
}

func (m *UpdateProcessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProcessResponse.Unmarshal(m, b)
}
func (m *UpdateProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProcessResponse.Marshal(b, m, deterministic)
}
func (m *UpdateProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProcessResponse.Merge(m, src)
}
func (m *UpdateProcessResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateProcessResponse.Size(m)
}
func (m *UpdateProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProcessResponse proto.InternalMessageInfo

func (m *UpdateProcessResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *UpdateProcessResponse) GetProcessId() string {
	if m != nil {
		return m.ProcessId
	}
	return ""
}

type OpenRequest struct {
	ProcessId            string   `protobuf:"bytes,1,opt,name=processId,proto3" json:"processId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenRequest) Reset()         { *m = OpenRequest{} }
func (m *OpenRequest) String() string { return proto.CompactTextString(m) }
func (*OpenRequest) ProtoMessage()    {}
func (*OpenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{15}
}

func (m *OpenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenRequest.Unmarshal(m, b)
}
func (m *OpenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenRequest.Marshal(b, m, deterministic)
}
func (m *OpenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenRequest.Merge(m, src)
}
func (m *OpenRequest) XXX_Size() int {
	return xxx_messageInfo_OpenRequest.Size(m)
}
func (m *OpenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OpenRequest proto.InternalMessageInfo

func (m *OpenRequest) GetProcessId() string {
	if m != nil {
		return m.ProcessId
	}
	return ""
}

type OpenResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *OpenResponse) Reset()         { *m = OpenResponse{} }
func (m *OpenResponse) String() string { return proto.CompactTextString(m) }
func (*OpenResponse) ProtoMessage()    {}
func (*OpenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{16}
}

func (m *OpenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenResponse.Unmarshal(m, b)
}
func (m *OpenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenResponse.Marshal(b, m, deterministic)
}
func (m *OpenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenResponse.Merge(m, src)
}
func (m *OpenResponse) XXX_Size() int {
	return xxx_messageInfo_OpenResponse.Size(m)
}
func (m *OpenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OpenResponse proto.InternalMessageInfo

func (m *OpenResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

type LaunchRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LaunchRequest) Reset()         { *m = LaunchRequest{} }
func (m *LaunchRequest) String() string { return proto.CompactTextString(m) }
func (*LaunchRequest) ProtoMessage()    {}
func (*LaunchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{17}
}

func (m *LaunchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchRequest.Unmarshal(m, b)
}
func (m *LaunchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchRequest.Marshal(b, m, deterministic)
}
func (m *LaunchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchRequest.Merge(m, src)
}
func (m *LaunchRequest) XXX_Size() int {
	return xxx_messageInfo_LaunchRequest.Size(m)
}
func (m *LaunchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchRequest proto.InternalMessageInfo

func (m *LaunchRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type LaunchResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LaunchResponse) Reset()         { *m = LaunchResponse{} }
func (m *LaunchResponse) String() string { return proto.CompactTextString(m) }
func (*LaunchResponse) ProtoMessage()    {}
func (*LaunchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7499dfaf74a17214, []int{18}
}

func (m *LaunchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchResponse.Unmarshal(m, b)
}
func (m *LaunchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchResponse.Marshal(b, m, deterministic)
}
func (m *LaunchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchResponse.Merge(m, src)
}
func (m *LaunchResponse) XXX_Size() int {
	return xxx_messageInfo_LaunchResponse.Size(m)
}
func (m *LaunchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchResponse proto.InternalMessageInfo

func (m *LaunchResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func init() {
	proto.RegisterType((*SearchProcessRequest)(nil), "xbasissvc.external.workflow.SearchProcessRequest")
	proto.RegisterType((*SearchProcessResponse)(nil), "xbasissvc.external.workflow.SearchProcessResponse")
	proto.RegisterType((*SearchProcessItem)(nil), "xbasissvc.external.workflow.SearchProcessItem")
	proto.RegisterType((*DetailProcessRequest)(nil), "xbasissvc.external.workflow.DetailProcessRequest")
	proto.RegisterType((*DetailProcessResponse)(nil), "xbasissvc.external.workflow.DetailProcessResponse")
	proto.RegisterType((*GetImageRequest)(nil), "xbasissvc.external.workflow.GetImageRequest")
	proto.RegisterType((*GetImageResponse)(nil), "xbasissvc.external.workflow.GetImageResponse")
	proto.RegisterType((*CreateRequest)(nil), "xbasissvc.external.workflow.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "xbasissvc.external.workflow.CreateResponse")
	proto.RegisterType((*BuildRequest)(nil), "xbasissvc.external.workflow.BuildRequest")
	proto.RegisterType((*BuildResponse)(nil), "xbasissvc.external.workflow.BuildResponse")
	proto.RegisterType((*DeleteProcessRequest)(nil), "xbasissvc.external.workflow.DeleteProcessRequest")
	proto.RegisterType((*DeleteProcessResponse)(nil), "xbasissvc.external.workflow.DeleteProcessResponse")
	proto.RegisterType((*UpdateProcessRequest)(nil), "xbasissvc.external.workflow.UpdateProcessRequest")
	proto.RegisterType((*UpdateProcessResponse)(nil), "xbasissvc.external.workflow.UpdateProcessResponse")
	proto.RegisterType((*OpenRequest)(nil), "xbasissvc.external.workflow.OpenRequest")
	proto.RegisterType((*OpenResponse)(nil), "xbasissvc.external.workflow.OpenResponse")
	proto.RegisterType((*LaunchRequest)(nil), "xbasissvc.external.workflow.LaunchRequest")
	proto.RegisterType((*LaunchResponse)(nil), "xbasissvc.external.workflow.LaunchResponse")
}

func init() { proto.RegisterFile("pb/process.proto", fileDescriptor_7499dfaf74a17214) }

var fileDescriptor_7499dfaf74a17214 = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x95, 0x13, 0x27, 0x6d, 0xa6, 0x4d, 0xbf, 0x7e, 0xab, 0x54, 0x0a, 0x86, 0x8b, 0x62, 0x71,
	0x91, 0xfe, 0xe0, 0x88, 0xc0, 0x3d, 0xb4, 0x54, 0x82, 0x48, 0x48, 0x20, 0x57, 0x5c, 0x20, 0x24,
	0xaa, 0x6d, 0x3c, 0x6d, 0xa3, 0x24, 0x5e, 0x63, 0x6f, 0x68, 0xcb, 0x2b, 0xc0, 0x23, 0xf1, 0x18,
	0x88, 0xe7, 0x41, 0xbb, 0xeb, 0xf8, 0xaf, 0x6e, 0xec, 0x58, 0xe2, 0x6e, 0x3d, 0x99, 0xd9, 0x33,
	0xe7, 0xec, 0x99, 0xec, 0xc2, 0xb6, 0x77, 0xde, 0xf7, 0x7c, 0x36, 0xc2, 0x20, 0xb0, 0x3c, 0x9f,
	0x71, 0x46, 0x1e, 0xde, 0x9c, 0xd3, 0x60, 0x1c, 0x04, 0xdf, 0x46, 0x16, 0xde, 0x70, 0xf4, 0x5d,
	0x3a, 0xb5, 0xae, 0x99, 0x3f, 0xb9, 0x98, 0xb2, 0x6b, 0xe3, 0x60, 0xc2, 0x5c, 0x9c, 0x4c, 0x98,
	0x35, 0xc3, 0xbe, 0xca, 0xeb, 0x8f, 0xd8, 0x6c, 0xc6, 0xdc, 0xa0, 0xef, 0x70, 0xb6, 0x58, 0xab,
	0x9d, 0xcc, 0x2b, 0xe8, 0x9c, 0x22, 0xf5, 0x47, 0x57, 0x1f, 0x14, 0x80, 0x8d, 0x5f, 0xe7, 0x18,
	0x70, 0xd2, 0x81, 0x06, 0xf5, 0xbc, 0xa1, 0xd3, 0xd5, 0x76, 0xb5, 0x5e, 0xcb, 0x56, 0x1f, 0x84,
	0x80, 0xee, 0xd2, 0x19, 0x76, 0x6b, 0x32, 0x28, 0xd7, 0x22, 0xe6, 0xd1, 0x4b, 0xec, 0xd6, 0x77,
	0xb5, 0x5e, 0xdd, 0x96, 0x6b, 0x11, 0x0b, 0xc6, 0xdf, 0xb1, 0xab, 0xab, 0x98, 0x58, 0x9b, 0x3f,
	0x35, 0xd8, 0xc9, 0x40, 0x05, 0x1e, 0x73, 0x03, 0x24, 0x7d, 0x68, 0x04, 0x9c, 0x72, 0x94, 0x58,
	0x1b, 0x83, 0x07, 0x96, 0xea, 0xda, 0x5a, 0x74, 0xea, 0x70, 0x66, 0x9d, 0x8a, 0x04, 0x5b, 0xe5,
	0x91, 0x63, 0xd0, 0x1d, 0xca, 0x69, 0xb7, 0xb6, 0x5b, 0xef, 0x6d, 0x0c, 0x2c, 0x6b, 0x89, 0x1a,
	0x56, 0x0a, 0x72, 0xc8, 0x71, 0x66, 0xcb, 0x5a, 0xf3, 0x87, 0x06, 0xff, 0xdf, 0xf9, 0x2d, 0x22,
	0xa8, 0x25, 0x08, 0x46, 0x52, 0xd4, 0x92, 0x52, 0x3c, 0x82, 0x56, 0x78, 0x26, 0x43, 0x47, 0x72,
	0x6f, 0xd9, 0x71, 0x80, 0x18, 0xb0, 0x3e, 0xf7, 0x1c, 0xca, 0xf1, 0x88, 0x4b, 0x11, 0x5a, 0x76,
	0xf4, 0x2d, 0x30, 0x1c, 0x0c, 0x46, 0xdd, 0xa6, 0xc2, 0x10, 0x6b, 0xf3, 0x05, 0x74, 0x4e, 0x90,
	0xd3, 0xf1, 0x34, 0x73, 0x0c, 0x29, 0x14, 0x2d, 0x83, 0x62, 0xfe, 0xd2, 0x60, 0x27, 0x53, 0x56,
	0x55, 0xd2, 0x27, 0xd0, 0x76, 0x99, 0x83, 0x27, 0x94, 0xd3, 0x23, 0xdf, 0xa7, 0xb7, 0x21, 0xd9,
	0x74, 0x50, 0x64, 0x4d, 0xc7, 0xee, 0x24, 0xce, 0x52, 0xc4, 0xd3, 0xc1, 0x88, 0xa0, 0x1e, 0x13,
	0x8c, 0x84, 0x6d, 0xc4, 0xc2, 0x9a, 0x7d, 0xf8, 0xef, 0x0d, 0xf2, 0xe1, 0x8c, 0x5e, 0x62, 0x39,
	0xbe, 0x9f, 0x60, 0x3b, 0x2e, 0xa8, 0xca, 0xb4, 0x03, 0x8d, 0xb1, 0xd8, 0x61, 0x71, 0x9c, 0xf2,
	0xc3, 0x3c, 0x83, 0xf6, 0x6b, 0x1f, 0x45, 0x5a, 0xd8, 0xc9, 0x36, 0xd4, 0x27, 0x78, 0x1b, 0xf6,
	0x20, 0x96, 0xf7, 0x99, 0x5f, 0x52, 0xad, 0x27, 0xa8, 0x46, 0x7e, 0xd1, 0x13, 0x7e, 0x31, 0xcf,
	0x60, 0x6b, 0x01, 0x50, 0xb5, 0xf3, 0x94, 0x38, 0xb5, 0xac, 0x38, 0x87, 0xb0, 0x79, 0x3c, 0x1f,
	0x4f, 0x9d, 0x72, 0x52, 0xbe, 0x82, 0x76, 0x98, 0x5d, 0xb1, 0x1b, 0x65, 0xd9, 0x29, 0x72, 0x5c,
	0xc9, 0xb2, 0x6f, 0x85, 0x63, 0x53, 0x55, 0x55, 0xf1, 0x7f, 0x6b, 0xd0, 0xf9, 0x28, 0x67, 0x6a,
	0x95, 0x06, 0x4a, 0x9f, 0xe2, 0x9d, 0x81, 0xd0, 0x4b, 0x0d, 0x44, 0x23, 0x6f, 0x20, 0x22, 0x47,
	0x34, 0x93, 0xff, 0x20, 0x91, 0x11, 0xd7, 0x92, 0x46, 0xbc, 0x80, 0x9d, 0x0c, 0xab, 0x7f, 0x63,
	0x97, 0x03, 0xd8, 0x78, 0xef, 0xa1, 0x5b, 0xee, 0xd4, 0x5e, 0xc2, 0xa6, 0x4a, 0xae, 0x7a, 0x58,
	0x8f, 0xa1, 0xfd, 0x8e, 0xce, 0xdd, 0xd1, 0xd5, 0xbd, 0xe3, 0x65, 0x1e, 0xc1, 0xd6, 0x22, 0xa5,
	0x22, 0xca, 0xe0, 0x4f, 0x13, 0xd6, 0x42, 0xd9, 0x08, 0x85, 0xa6, 0xda, 0x8e, 0xec, 0x2f, 0xbd,
	0x1f, 0x52, 0x6d, 0x19, 0x07, 0xa5, 0x72, 0xc3, 0xfe, 0xbe, 0x40, 0x43, 0xce, 0x10, 0xd9, 0x5b,
	0x5a, 0x95, 0x9c, 0x4a, 0x63, 0xbf, 0x4c, 0x6a, 0xb8, 0x3f, 0x83, 0xa6, 0x9a, 0x15, 0xf2, 0x6c,
	0x69, 0x55, 0xde, 0x18, 0x1a, 0x83, 0x55, 0x4a, 0x62, 0x40, 0xe5, 0xbd, 0x02, 0xc0, 0xbc, 0xb1,
	0x2b, 0x00, 0xcc, 0xf7, 0xf4, 0x67, 0xd0, 0x85, 0xaf, 0x48, 0x6f, 0x69, 0x6d, 0xc2, 0xa7, 0xc6,
	0x5e, 0x89, 0xcc, 0xa4, 0x7c, 0xe2, 0x72, 0x2c, 0x94, 0xef, 0xee, 0xc5, 0x5b, 0x28, 0x5f, 0xde,
	0xa5, 0x7b, 0x09, 0xeb, 0x8b, 0xeb, 0x89, 0x1c, 0x2e, 0xad, 0xcf, 0x5c, 0x7b, 0xc6, 0xd3, 0x92,
	0xd9, 0x31, 0x33, 0xf5, 0x74, 0x29, 0x60, 0x96, 0xf7, 0xb2, 0x2b, 0x60, 0x96, 0xfb, 0x42, 0x3b,
	0x6f, 0xca, 0xc7, 0xe2, 0xf3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x1e, 0x7b, 0x20, 0x8a,
	0x0a, 0x00, 0x00,
}
