// Code generated by protoc-gen-go. DO NOT EDIT.
// source: analysis/pb/logger.proto

package gs_service_analysis

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

type UsageFunctionRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsageFunctionRequest) Reset()         { *m = UsageFunctionRequest{} }
func (m *UsageFunctionRequest) String() string { return proto.CompactTextString(m) }
func (*UsageFunctionRequest) ProtoMessage()    {}
func (*UsageFunctionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c694e20a0de2280, []int{0}
}

func (m *UsageFunctionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsageFunctionRequest.Unmarshal(m, b)
}
func (m *UsageFunctionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsageFunctionRequest.Marshal(b, m, deterministic)
}
func (m *UsageFunctionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsageFunctionRequest.Merge(m, src)
}
func (m *UsageFunctionRequest) XXX_Size() int {
	return xxx_messageInfo_UsageFunctionRequest.Size(m)
}
func (m *UsageFunctionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UsageFunctionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UsageFunctionRequest proto.InternalMessageInfo

func (m *UsageFunctionRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type UsageFunctionStatus struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsageFunctionStatus) Reset()         { *m = UsageFunctionStatus{} }
func (m *UsageFunctionStatus) String() string { return proto.CompactTextString(m) }
func (*UsageFunctionStatus) ProtoMessage()    {}
func (*UsageFunctionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c694e20a0de2280, []int{1}
}

func (m *UsageFunctionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsageFunctionStatus.Unmarshal(m, b)
}
func (m *UsageFunctionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsageFunctionStatus.Marshal(b, m, deterministic)
}
func (m *UsageFunctionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsageFunctionStatus.Merge(m, src)
}
func (m *UsageFunctionStatus) XXX_Size() int {
	return xxx_messageInfo_UsageFunctionStatus.Size(m)
}
func (m *UsageFunctionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_UsageFunctionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_UsageFunctionStatus proto.InternalMessageInfo

func (m *UsageFunctionStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UsageFunctionStatus) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type UsageFunctionResponse struct {
	State                *dto.State             `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Status               []*UsageFunctionStatus `protobuf:"bytes,2,rep,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *UsageFunctionResponse) Reset()         { *m = UsageFunctionResponse{} }
func (m *UsageFunctionResponse) String() string { return proto.CompactTextString(m) }
func (*UsageFunctionResponse) ProtoMessage()    {}
func (*UsageFunctionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c694e20a0de2280, []int{2}
}

func (m *UsageFunctionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsageFunctionResponse.Unmarshal(m, b)
}
func (m *UsageFunctionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsageFunctionResponse.Marshal(b, m, deterministic)
}
func (m *UsageFunctionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsageFunctionResponse.Merge(m, src)
}
func (m *UsageFunctionResponse) XXX_Size() int {
	return xxx_messageInfo_UsageFunctionResponse.Size(m)
}
func (m *UsageFunctionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UsageFunctionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UsageFunctionResponse proto.InternalMessageInfo

func (m *UsageFunctionResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *UsageFunctionResponse) GetStatus() []*UsageFunctionStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type TodayVisitRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodayVisitRequest) Reset()         { *m = TodayVisitRequest{} }
func (m *TodayVisitRequest) String() string { return proto.CompactTextString(m) }
func (*TodayVisitRequest) ProtoMessage()    {}
func (*TodayVisitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c694e20a0de2280, []int{3}
}

func (m *TodayVisitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodayVisitRequest.Unmarshal(m, b)
}
func (m *TodayVisitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodayVisitRequest.Marshal(b, m, deterministic)
}
func (m *TodayVisitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodayVisitRequest.Merge(m, src)
}
func (m *TodayVisitRequest) XXX_Size() int {
	return xxx_messageInfo_TodayVisitRequest.Size(m)
}
func (m *TodayVisitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TodayVisitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TodayVisitRequest proto.InternalMessageInfo

func (m *TodayVisitRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type TodayVisitResponse struct {
	State                *dto.State            `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Data                 map[string]*VisitInfo `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TodayVisitResponse) Reset()         { *m = TodayVisitResponse{} }
func (m *TodayVisitResponse) String() string { return proto.CompactTextString(m) }
func (*TodayVisitResponse) ProtoMessage()    {}
func (*TodayVisitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c694e20a0de2280, []int{4}
}

func (m *TodayVisitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodayVisitResponse.Unmarshal(m, b)
}
func (m *TodayVisitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodayVisitResponse.Marshal(b, m, deterministic)
}
func (m *TodayVisitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodayVisitResponse.Merge(m, src)
}
func (m *TodayVisitResponse) XXX_Size() int {
	return xxx_messageInfo_TodayVisitResponse.Size(m)
}
func (m *TodayVisitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TodayVisitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TodayVisitResponse proto.InternalMessageInfo

func (m *TodayVisitResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *TodayVisitResponse) GetData() map[string]*VisitInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

type VisitInfo struct {
	Today                int64                `protobuf:"varint,1,opt,name=today,proto3" json:"today,omitempty"`
	Yesterday            int64                `protobuf:"varint,2,opt,name=yesterday,proto3" json:"yesterday,omitempty"`
	Info                 []*PlatformVisitInfo `protobuf:"bytes,3,rep,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *VisitInfo) Reset()         { *m = VisitInfo{} }
func (m *VisitInfo) String() string { return proto.CompactTextString(m) }
func (*VisitInfo) ProtoMessage()    {}
func (*VisitInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c694e20a0de2280, []int{5}
}

func (m *VisitInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VisitInfo.Unmarshal(m, b)
}
func (m *VisitInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VisitInfo.Marshal(b, m, deterministic)
}
func (m *VisitInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VisitInfo.Merge(m, src)
}
func (m *VisitInfo) XXX_Size() int {
	return xxx_messageInfo_VisitInfo.Size(m)
}
func (m *VisitInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VisitInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VisitInfo proto.InternalMessageInfo

func (m *VisitInfo) GetToday() int64 {
	if m != nil {
		return m.Today
	}
	return 0
}

func (m *VisitInfo) GetYesterday() int64 {
	if m != nil {
		return m.Yesterday
	}
	return 0
}

func (m *VisitInfo) GetInfo() []*PlatformVisitInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type PlatformVisitInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlatformVisitInfo) Reset()         { *m = PlatformVisitInfo{} }
func (m *PlatformVisitInfo) String() string { return proto.CompactTextString(m) }
func (*PlatformVisitInfo) ProtoMessage()    {}
func (*PlatformVisitInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c694e20a0de2280, []int{6}
}

func (m *PlatformVisitInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformVisitInfo.Unmarshal(m, b)
}
func (m *PlatformVisitInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformVisitInfo.Marshal(b, m, deterministic)
}
func (m *PlatformVisitInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformVisitInfo.Merge(m, src)
}
func (m *PlatformVisitInfo) XXX_Size() int {
	return xxx_messageInfo_PlatformVisitInfo.Size(m)
}
func (m *PlatformVisitInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformVisitInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformVisitInfo proto.InternalMessageInfo

func (m *PlatformVisitInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PlatformVisitInfo) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetDataResponse struct {
	State                *dto.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Data                 string     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetDataResponse) Reset()         { *m = GetDataResponse{} }
func (m *GetDataResponse) String() string { return proto.CompactTextString(m) }
func (*GetDataResponse) ProtoMessage()    {}
func (*GetDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c694e20a0de2280, []int{7}
}

func (m *GetDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataResponse.Unmarshal(m, b)
}
func (m *GetDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataResponse.Marshal(b, m, deterministic)
}
func (m *GetDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataResponse.Merge(m, src)
}
func (m *GetDataResponse) XXX_Size() int {
	return xxx_messageInfo_GetDataResponse.Size(m)
}
func (m *GetDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataResponse proto.InternalMessageInfo

func (m *GetDataResponse) GetState() *dto.State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *GetDataResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type GetDataRequest struct {
	XAxis                *XAxisRequest `protobuf:"bytes,1,opt,name=x_axis,json=xAxis,proto3" json:"x_axis,omitempty"`
	YAxis                *YAxisRequest `protobuf:"bytes,2,opt,name=y_axis,json=yAxis,proto3" json:"y_axis,omitempty"`
	Map                  string        `protobuf:"bytes,3,opt,name=map,proto3" json:"map,omitempty"`
	Sd                   int64         `protobuf:"varint,4,opt,name=sd,proto3" json:"sd,omitempty"`
	Ed                   int64         `protobuf:"varint,5,opt,name=ed,proto3" json:"ed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetDataRequest) Reset()         { *m = GetDataRequest{} }
func (m *GetDataRequest) String() string { return proto.CompactTextString(m) }
func (*GetDataRequest) ProtoMessage()    {}
func (*GetDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c694e20a0de2280, []int{8}
}

func (m *GetDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataRequest.Unmarshal(m, b)
}
func (m *GetDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataRequest.Marshal(b, m, deterministic)
}
func (m *GetDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataRequest.Merge(m, src)
}
func (m *GetDataRequest) XXX_Size() int {
	return xxx_messageInfo_GetDataRequest.Size(m)
}
func (m *GetDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataRequest proto.InternalMessageInfo

func (m *GetDataRequest) GetXAxis() *XAxisRequest {
	if m != nil {
		return m.XAxis
	}
	return nil
}

func (m *GetDataRequest) GetYAxis() *YAxisRequest {
	if m != nil {
		return m.YAxis
	}
	return nil
}

func (m *GetDataRequest) GetMap() string {
	if m != nil {
		return m.Map
	}
	return ""
}

func (m *GetDataRequest) GetSd() int64 {
	if m != nil {
		return m.Sd
	}
	return 0
}

func (m *GetDataRequest) GetEd() int64 {
	if m != nil {
		return m.Ed
	}
	return 0
}

type XAxisRequest struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Factors              []*XAxisFactor `protobuf:"bytes,2,rep,name=factors,proto3" json:"factors,omitempty"`
	Order                string         `protobuf:"bytes,3,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *XAxisRequest) Reset()         { *m = XAxisRequest{} }
func (m *XAxisRequest) String() string { return proto.CompactTextString(m) }
func (*XAxisRequest) ProtoMessage()    {}
func (*XAxisRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c694e20a0de2280, []int{9}
}

func (m *XAxisRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XAxisRequest.Unmarshal(m, b)
}
func (m *XAxisRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XAxisRequest.Marshal(b, m, deterministic)
}
func (m *XAxisRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XAxisRequest.Merge(m, src)
}
func (m *XAxisRequest) XXX_Size() int {
	return xxx_messageInfo_XAxisRequest.Size(m)
}
func (m *XAxisRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_XAxisRequest.DiscardUnknown(m)
}

var xxx_messageInfo_XAxisRequest proto.InternalMessageInfo

func (m *XAxisRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *XAxisRequest) GetFactors() []*XAxisFactor {
	if m != nil {
		return m.Factors
	}
	return nil
}

func (m *XAxisRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

type XAxisFactor struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Operation            string   `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XAxisFactor) Reset()         { *m = XAxisFactor{} }
func (m *XAxisFactor) String() string { return proto.CompactTextString(m) }
func (*XAxisFactor) ProtoMessage()    {}
func (*XAxisFactor) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c694e20a0de2280, []int{10}
}

func (m *XAxisFactor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XAxisFactor.Unmarshal(m, b)
}
func (m *XAxisFactor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XAxisFactor.Marshal(b, m, deterministic)
}
func (m *XAxisFactor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XAxisFactor.Merge(m, src)
}
func (m *XAxisFactor) XXX_Size() int {
	return xxx_messageInfo_XAxisFactor.Size(m)
}
func (m *XAxisFactor) XXX_DiscardUnknown() {
	xxx_messageInfo_XAxisFactor.DiscardUnknown(m)
}

var xxx_messageInfo_XAxisFactor proto.InternalMessageInfo

func (m *XAxisFactor) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *XAxisFactor) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *XAxisFactor) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type YAxisRequest struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size                 int64          `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Factors              []*YAxisFactor `protobuf:"bytes,3,rep,name=factors,proto3" json:"factors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *YAxisRequest) Reset()         { *m = YAxisRequest{} }
func (m *YAxisRequest) String() string { return proto.CompactTextString(m) }
func (*YAxisRequest) ProtoMessage()    {}
func (*YAxisRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c694e20a0de2280, []int{11}
}

func (m *YAxisRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_YAxisRequest.Unmarshal(m, b)
}
func (m *YAxisRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_YAxisRequest.Marshal(b, m, deterministic)
}
func (m *YAxisRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_YAxisRequest.Merge(m, src)
}
func (m *YAxisRequest) XXX_Size() int {
	return xxx_messageInfo_YAxisRequest.Size(m)
}
func (m *YAxisRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_YAxisRequest.DiscardUnknown(m)
}

var xxx_messageInfo_YAxisRequest proto.InternalMessageInfo

func (m *YAxisRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *YAxisRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *YAxisRequest) GetFactors() []*YAxisFactor {
	if m != nil {
		return m.Factors
	}
	return nil
}

type YAxisFactor struct {
	Field                string       `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Operation            string       `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	Value                string       `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Factor               *YAxisFactor `protobuf:"bytes,4,opt,name=factor,proto3" json:"factor,omitempty"`
	Name                 string       `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *YAxisFactor) Reset()         { *m = YAxisFactor{} }
func (m *YAxisFactor) String() string { return proto.CompactTextString(m) }
func (*YAxisFactor) ProtoMessage()    {}
func (*YAxisFactor) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c694e20a0de2280, []int{12}
}

func (m *YAxisFactor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_YAxisFactor.Unmarshal(m, b)
}
func (m *YAxisFactor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_YAxisFactor.Marshal(b, m, deterministic)
}
func (m *YAxisFactor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_YAxisFactor.Merge(m, src)
}
func (m *YAxisFactor) XXX_Size() int {
	return xxx_messageInfo_YAxisFactor.Size(m)
}
func (m *YAxisFactor) XXX_DiscardUnknown() {
	xxx_messageInfo_YAxisFactor.DiscardUnknown(m)
}

var xxx_messageInfo_YAxisFactor proto.InternalMessageInfo

func (m *YAxisFactor) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *YAxisFactor) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *YAxisFactor) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *YAxisFactor) GetFactor() *YAxisFactor {
	if m != nil {
		return m.Factor
	}
	return nil
}

func (m *YAxisFactor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*UsageFunctionRequest)(nil), "gs.service.analysis.UsageFunctionRequest")
	proto.RegisterType((*UsageFunctionStatus)(nil), "gs.service.analysis.UsageFunctionStatus")
	proto.RegisterType((*UsageFunctionResponse)(nil), "gs.service.analysis.UsageFunctionResponse")
	proto.RegisterType((*TodayVisitRequest)(nil), "gs.service.analysis.TodayVisitRequest")
	proto.RegisterType((*TodayVisitResponse)(nil), "gs.service.analysis.TodayVisitResponse")
	proto.RegisterMapType((map[string]*VisitInfo)(nil), "gs.service.analysis.TodayVisitResponse.DataEntry")
	proto.RegisterType((*VisitInfo)(nil), "gs.service.analysis.VisitInfo")
	proto.RegisterType((*PlatformVisitInfo)(nil), "gs.service.analysis.PlatformVisitInfo")
	proto.RegisterType((*GetDataResponse)(nil), "gs.service.analysis.GetDataResponse")
	proto.RegisterType((*GetDataRequest)(nil), "gs.service.analysis.GetDataRequest")
	proto.RegisterType((*XAxisRequest)(nil), "gs.service.analysis.XAxisRequest")
	proto.RegisterType((*XAxisFactor)(nil), "gs.service.analysis.XAxisFactor")
	proto.RegisterType((*YAxisRequest)(nil), "gs.service.analysis.YAxisRequest")
	proto.RegisterType((*YAxisFactor)(nil), "gs.service.analysis.YAxisFactor")
}

func init() { proto.RegisterFile("analysis/pb/logger.proto", fileDescriptor_7c694e20a0de2280) }

var fileDescriptor_7c694e20a0de2280 = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0x92, 0xb6, 0xa8, 0xa7, 0x63, 0x6c, 0xde, 0x26, 0x45, 0x13, 0x42, 0x25, 0xa0, 0xd1,
	0x31, 0x94, 0x8a, 0xc2, 0xc5, 0x34, 0x09, 0x01, 0x12, 0xdb, 0x34, 0x89, 0x0b, 0x14, 0x7e, 0xb6,
	0x5e, 0x20, 0xe4, 0x35, 0x6e, 0x15, 0xb5, 0xb1, 0x43, 0xec, 0x4e, 0x0d, 0xdc, 0xf3, 0x26, 0x5c,
	0xf0, 0x02, 0xbc, 0x0f, 0x6f, 0x82, 0xfc, 0x93, 0xb4, 0x65, 0xa1, 0x2b, 0x93, 0xb8, 0xf3, 0x71,
	0xbe, 0xef, 0x9c, 0xef, 0xf3, 0x39, 0x76, 0xc0, 0xc5, 0x14, 0x8f, 0x32, 0x1e, 0xf1, 0x76, 0x72,
	0xde, 0x1e, 0xb1, 0xc1, 0x80, 0xa4, 0x7e, 0x92, 0x32, 0xc1, 0xd0, 0xc6, 0x80, 0xfb, 0x9c, 0xa4,
	0x17, 0x51, 0x8f, 0xf8, 0x39, 0x68, 0x7b, 0x6f, 0xc8, 0x28, 0x19, 0x0e, 0x99, 0x1f, 0x93, 0xf6,
	0x80, 0xf1, 0x88, 0xd1, 0x76, 0x8f, 0xc5, 0x31, 0xa3, 0xbc, 0x1d, 0x0a, 0x96, 0xaf, 0x75, 0x06,
	0xef, 0x11, 0x6c, 0xbe, 0xe7, 0x78, 0x40, 0x8e, 0xc6, 0xb4, 0x27, 0x22, 0x46, 0x03, 0xf2, 0x79,
	0x4c, 0xb8, 0x40, 0x9b, 0x50, 0xc5, 0x49, 0x72, 0x12, 0xba, 0x56, 0xd3, 0x6a, 0xd5, 0x03, 0x1d,
	0x78, 0xcf, 0x61, 0x63, 0x0e, 0xfd, 0x56, 0x60, 0x31, 0xe6, 0x08, 0x41, 0x85, 0xe2, 0x98, 0x18,
	0xac, 0x5a, 0xcb, 0x04, 0x3d, 0x36, 0xa6, 0xc2, 0xb5, 0x9b, 0x56, 0xcb, 0x09, 0x74, 0xe0, 0x7d,
	0xb3, 0x60, 0xeb, 0x8f, 0x7a, 0x3c, 0x61, 0x94, 0x13, 0xb4, 0x07, 0x55, 0x2e, 0xb0, 0xd0, 0x49,
	0x1a, 0x9d, 0x2d, 0x7f, 0xc0, 0xfd, 0x5c, 0x6a, 0x28, 0x98, 0x2f, 0x4b, 0x91, 0x40, 0x63, 0xd0,
	0x0b, 0xa8, 0x71, 0x55, 0xda, 0xb5, 0x9b, 0x4e, 0xab, 0xd1, 0x69, 0xf9, 0x25, 0x07, 0xe1, 0x97,
	0x48, 0x0d, 0x0c, 0xcf, 0xdb, 0x85, 0xf5, 0x77, 0x2c, 0xc4, 0xd9, 0x87, 0x88, 0x47, 0x62, 0xb1,
	0xe9, 0x5f, 0x16, 0xa0, 0x59, 0xec, 0x75, 0x04, 0x1f, 0x42, 0x25, 0xc4, 0x02, 0x1b, 0xb9, 0x8f,
	0x4b, 0xe5, 0x5e, 0xae, 0xe1, 0xbf, 0xc2, 0x02, 0x1f, 0x52, 0x91, 0x66, 0x81, 0xa2, 0x6f, 0x9f,
	0x42, 0xbd, 0xd8, 0x42, 0x6b, 0xe0, 0x0c, 0x49, 0x66, 0xb4, 0xca, 0x25, 0x7a, 0x0a, 0xd5, 0x0b,
	0x3c, 0x1a, 0x13, 0x75, 0xe6, 0x8d, 0xce, 0x9d, 0xd2, 0x32, 0xaa, 0xc2, 0x09, 0xed, 0xb3, 0x40,
	0x83, 0x0f, 0xec, 0x7d, 0xcb, 0xfb, 0x0a, 0xf5, 0x62, 0x5f, 0x1e, 0x83, 0x90, 0x5a, 0x54, 0x6a,
	0x27, 0xd0, 0x01, 0xba, 0x0d, 0xf5, 0x8c, 0x70, 0x41, 0x52, 0xf9, 0x45, 0x37, 0x75, 0xba, 0x81,
	0x0e, 0xa0, 0x12, 0xd1, 0x3e, 0x73, 0x1d, 0x65, 0x70, 0xa7, 0xb4, 0xf2, 0x9b, 0x11, 0x16, 0x7d,
	0x96, 0xc6, 0x53, 0x05, 0x8a, 0xe3, 0x3d, 0x83, 0xf5, 0x4b, 0x9f, 0xfe, 0x61, 0xa6, 0x02, 0xb8,
	0x75, 0x4c, 0x84, 0x3c, 0x97, 0xeb, 0xf5, 0x06, 0x15, 0xbd, 0x51, 0x95, 0xe4, 0xda, 0xfb, 0x69,
	0xc1, 0x6a, 0x91, 0x54, 0x0f, 0xc7, 0x3e, 0xd4, 0x26, 0x9f, 0xf0, 0x24, 0xe2, 0x26, 0xe9, 0xdd,
	0x52, 0x8f, 0x67, 0x2f, 0x27, 0x11, 0x37, 0x94, 0xa0, 0x3a, 0x91, 0x91, 0x64, 0x66, 0x9a, 0x69,
	0x2f, 0x60, 0x76, 0xe7, 0x98, 0x99, 0x62, 0xae, 0x81, 0x13, 0xe3, 0xc4, 0x75, 0x74, 0x8b, 0x63,
	0x9c, 0xa0, 0x55, 0xb0, 0x79, 0xe8, 0x56, 0x94, 0x7f, 0x9b, 0x87, 0x32, 0x26, 0xa1, 0x5b, 0xd5,
	0x31, 0x09, 0x3d, 0x01, 0x2b, 0xb3, 0x12, 0x4a, 0x8f, 0xf1, 0x00, 0x6e, 0xf4, 0x71, 0x4f, 0xb0,
	0x34, 0xbf, 0x3e, 0xcd, 0xbf, 0x5b, 0x39, 0x52, 0xc0, 0x20, 0x27, 0xc8, 0x16, 0xb0, 0x34, 0x24,
	0xa9, 0xd1, 0xa4, 0x03, 0xef, 0x14, 0x1a, 0x33, 0x68, 0x09, 0xea, 0x47, 0x64, 0x54, 0xdc, 0x23,
	0x15, 0xc8, 0x01, 0x62, 0x09, 0x49, 0xb1, 0xbc, 0x8d, 0xe6, 0xb0, 0xa7, 0x1b, 0x92, 0xa3, 0x67,
	0xd7, 0x24, 0x56, 0x81, 0x97, 0xc2, 0x4a, 0xf7, 0x2a, 0x3b, 0x08, 0x2a, 0x3c, 0xfa, 0x42, 0xcc,
	0x50, 0xa8, 0xf5, 0xac, 0x45, 0x67, 0x81, 0xc5, 0x6e, 0x89, 0x45, 0xef, 0xbb, 0x05, 0x8d, 0xee,
	0xff, 0x70, 0x23, 0x07, 0x41, 0x17, 0x51, 0x0d, 0x5c, 0x46, 0x94, 0xc1, 0x17, 0xbe, 0xab, 0x53,
	0xdf, 0x9d, 0x1f, 0x36, 0xd4, 0x5e, 0xab, 0xbf, 0x01, 0x3a, 0x83, 0xc6, 0x31, 0x11, 0x92, 0x27,
	0x27, 0x16, 0xdd, 0x2b, 0xcd, 0x3b, 0x3f, 0xcf, 0xdb, 0xf7, 0x17, 0x83, 0xcc, 0x4d, 0xfa, 0x08,
	0x30, 0x7d, 0x97, 0xd0, 0xce, 0x95, 0x0f, 0x97, 0xce, 0xfd, 0x60, 0xc9, 0x07, 0x0e, 0xf5, 0xe1,
	0xe6, 0xdc, 0x2b, 0x8d, 0x76, 0xaf, 0x7e, 0xc9, 0xf3, 0x22, 0x0f, 0x97, 0x81, 0xea, 0x3a, 0xe7,
	0x35, 0xf5, 0xb7, 0x7b, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x67, 0xd7, 0x4b, 0x3b, 0x4b, 0x07,
	0x00, 0x00,
}
