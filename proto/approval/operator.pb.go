// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/approval/operator.proto

package approval

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 加入的请求
type OperatorJoinRequest struct {
	Operator             string   `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	Workflow             string   `protobuf:"bytes,2,opt,name=workflow,proto3" json:"workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperatorJoinRequest) Reset()         { *m = OperatorJoinRequest{} }
func (m *OperatorJoinRequest) String() string { return proto.CompactTextString(m) }
func (*OperatorJoinRequest) ProtoMessage()    {}
func (*OperatorJoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_10682c9168bd77ed, []int{0}
}

func (m *OperatorJoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatorJoinRequest.Unmarshal(m, b)
}
func (m *OperatorJoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatorJoinRequest.Marshal(b, m, deterministic)
}
func (m *OperatorJoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatorJoinRequest.Merge(m, src)
}
func (m *OperatorJoinRequest) XXX_Size() int {
	return xxx_messageInfo_OperatorJoinRequest.Size(m)
}
func (m *OperatorJoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatorJoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OperatorJoinRequest proto.InternalMessageInfo

func (m *OperatorJoinRequest) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *OperatorJoinRequest) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

// 离开的请求
type OperatorLeaveRequest struct {
	Operator             string   `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	Workflow             string   `protobuf:"bytes,2,opt,name=workflow,proto3" json:"workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperatorLeaveRequest) Reset()         { *m = OperatorLeaveRequest{} }
func (m *OperatorLeaveRequest) String() string { return proto.CompactTextString(m) }
func (*OperatorLeaveRequest) ProtoMessage()    {}
func (*OperatorLeaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_10682c9168bd77ed, []int{1}
}

func (m *OperatorLeaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatorLeaveRequest.Unmarshal(m, b)
}
func (m *OperatorLeaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatorLeaveRequest.Marshal(b, m, deterministic)
}
func (m *OperatorLeaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatorLeaveRequest.Merge(m, src)
}
func (m *OperatorLeaveRequest) XXX_Size() int {
	return xxx_messageInfo_OperatorLeaveRequest.Size(m)
}
func (m *OperatorLeaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatorLeaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OperatorLeaveRequest proto.InternalMessageInfo

func (m *OperatorLeaveRequest) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *OperatorLeaveRequest) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

// 批量加入的请求
type OperatorBatchJoinRequest struct {
	Operator             []string `protobuf:"bytes,1,rep,name=operator,proto3" json:"operator,omitempty"`
	Workflow             string   `protobuf:"bytes,2,opt,name=workflow,proto3" json:"workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperatorBatchJoinRequest) Reset()         { *m = OperatorBatchJoinRequest{} }
func (m *OperatorBatchJoinRequest) String() string { return proto.CompactTextString(m) }
func (*OperatorBatchJoinRequest) ProtoMessage()    {}
func (*OperatorBatchJoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_10682c9168bd77ed, []int{2}
}

func (m *OperatorBatchJoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatorBatchJoinRequest.Unmarshal(m, b)
}
func (m *OperatorBatchJoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatorBatchJoinRequest.Marshal(b, m, deterministic)
}
func (m *OperatorBatchJoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatorBatchJoinRequest.Merge(m, src)
}
func (m *OperatorBatchJoinRequest) XXX_Size() int {
	return xxx_messageInfo_OperatorBatchJoinRequest.Size(m)
}
func (m *OperatorBatchJoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatorBatchJoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OperatorBatchJoinRequest proto.InternalMessageInfo

func (m *OperatorBatchJoinRequest) GetOperator() []string {
	if m != nil {
		return m.Operator
	}
	return nil
}

func (m *OperatorBatchJoinRequest) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

// 批量离开的请求
type OperatorBatchLeaveRequest struct {
	Operator             []string `protobuf:"bytes,1,rep,name=operator,proto3" json:"operator,omitempty"`
	Workflow             string   `protobuf:"bytes,2,opt,name=workflow,proto3" json:"workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperatorBatchLeaveRequest) Reset()         { *m = OperatorBatchLeaveRequest{} }
func (m *OperatorBatchLeaveRequest) String() string { return proto.CompactTextString(m) }
func (*OperatorBatchLeaveRequest) ProtoMessage()    {}
func (*OperatorBatchLeaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_10682c9168bd77ed, []int{3}
}

func (m *OperatorBatchLeaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatorBatchLeaveRequest.Unmarshal(m, b)
}
func (m *OperatorBatchLeaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatorBatchLeaveRequest.Marshal(b, m, deterministic)
}
func (m *OperatorBatchLeaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatorBatchLeaveRequest.Merge(m, src)
}
func (m *OperatorBatchLeaveRequest) XXX_Size() int {
	return xxx_messageInfo_OperatorBatchLeaveRequest.Size(m)
}
func (m *OperatorBatchLeaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatorBatchLeaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OperatorBatchLeaveRequest proto.InternalMessageInfo

func (m *OperatorBatchLeaveRequest) GetOperator() []string {
	if m != nil {
		return m.Operator
	}
	return nil
}

func (m *OperatorBatchLeaveRequest) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

// 列举的请求
type OperatorListRequest struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Workflow             string   `protobuf:"bytes,3,opt,name=workflow,proto3" json:"workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperatorListRequest) Reset()         { *m = OperatorListRequest{} }
func (m *OperatorListRequest) String() string { return proto.CompactTextString(m) }
func (*OperatorListRequest) ProtoMessage()    {}
func (*OperatorListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_10682c9168bd77ed, []int{4}
}

func (m *OperatorListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatorListRequest.Unmarshal(m, b)
}
func (m *OperatorListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatorListRequest.Marshal(b, m, deterministic)
}
func (m *OperatorListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatorListRequest.Merge(m, src)
}
func (m *OperatorListRequest) XXX_Size() int {
	return xxx_messageInfo_OperatorListRequest.Size(m)
}
func (m *OperatorListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatorListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OperatorListRequest proto.InternalMessageInfo

func (m *OperatorListRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *OperatorListRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *OperatorListRequest) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

// 列举的回复
type OperatorListResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint64   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Entity               []string `protobuf:"bytes,3,rep,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperatorListResponse) Reset()         { *m = OperatorListResponse{} }
func (m *OperatorListResponse) String() string { return proto.CompactTextString(m) }
func (*OperatorListResponse) ProtoMessage()    {}
func (*OperatorListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_10682c9168bd77ed, []int{5}
}

func (m *OperatorListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatorListResponse.Unmarshal(m, b)
}
func (m *OperatorListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatorListResponse.Marshal(b, m, deterministic)
}
func (m *OperatorListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatorListResponse.Merge(m, src)
}
func (m *OperatorListResponse) XXX_Size() int {
	return xxx_messageInfo_OperatorListResponse.Size(m)
}
func (m *OperatorListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatorListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OperatorListResponse proto.InternalMessageInfo

func (m *OperatorListResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *OperatorListResponse) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *OperatorListResponse) GetEntity() []string {
	if m != nil {
		return m.Entity
	}
	return nil
}

// 过滤的请求
type OperatorFilterRequest struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperatorFilterRequest) Reset()         { *m = OperatorFilterRequest{} }
func (m *OperatorFilterRequest) String() string { return proto.CompactTextString(m) }
func (*OperatorFilterRequest) ProtoMessage()    {}
func (*OperatorFilterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_10682c9168bd77ed, []int{6}
}

func (m *OperatorFilterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatorFilterRequest.Unmarshal(m, b)
}
func (m *OperatorFilterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatorFilterRequest.Marshal(b, m, deterministic)
}
func (m *OperatorFilterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatorFilterRequest.Merge(m, src)
}
func (m *OperatorFilterRequest) XXX_Size() int {
	return xxx_messageInfo_OperatorFilterRequest.Size(m)
}
func (m *OperatorFilterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatorFilterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OperatorFilterRequest proto.InternalMessageInfo

func (m *OperatorFilterRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *OperatorFilterRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *OperatorFilterRequest) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

// 过滤的回复
type OperatorFilterResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint64   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Entity               []string `protobuf:"bytes,3,rep,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperatorFilterResponse) Reset()         { *m = OperatorFilterResponse{} }
func (m *OperatorFilterResponse) String() string { return proto.CompactTextString(m) }
func (*OperatorFilterResponse) ProtoMessage()    {}
func (*OperatorFilterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_10682c9168bd77ed, []int{7}
}

func (m *OperatorFilterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatorFilterResponse.Unmarshal(m, b)
}
func (m *OperatorFilterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatorFilterResponse.Marshal(b, m, deterministic)
}
func (m *OperatorFilterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatorFilterResponse.Merge(m, src)
}
func (m *OperatorFilterResponse) XXX_Size() int {
	return xxx_messageInfo_OperatorFilterResponse.Size(m)
}
func (m *OperatorFilterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatorFilterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OperatorFilterResponse proto.InternalMessageInfo

func (m *OperatorFilterResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *OperatorFilterResponse) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *OperatorFilterResponse) GetEntity() []string {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterType((*OperatorJoinRequest)(nil), "approval.OperatorJoinRequest")
	proto.RegisterType((*OperatorLeaveRequest)(nil), "approval.OperatorLeaveRequest")
	proto.RegisterType((*OperatorBatchJoinRequest)(nil), "approval.OperatorBatchJoinRequest")
	proto.RegisterType((*OperatorBatchLeaveRequest)(nil), "approval.OperatorBatchLeaveRequest")
	proto.RegisterType((*OperatorListRequest)(nil), "approval.OperatorListRequest")
	proto.RegisterType((*OperatorListResponse)(nil), "approval.OperatorListResponse")
	proto.RegisterType((*OperatorFilterRequest)(nil), "approval.OperatorFilterRequest")
	proto.RegisterType((*OperatorFilterResponse)(nil), "approval.OperatorFilterResponse")
}

func init() {
	proto.RegisterFile("proto/approval/operator.proto", fileDescriptor_10682c9168bd77ed)
}

var fileDescriptor_10682c9168bd77ed = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xdf, 0x4e, 0xea, 0x40,
	0x10, 0xc6, 0x0f, 0xa7, 0xd0, 0xc0, 0x78, 0x63, 0x56, 0xc4, 0x5a, 0x03, 0x92, 0x7a, 0xc3, 0x15,
	0x24, 0xf8, 0x02, 0x86, 0x0b, 0x4d, 0x08, 0x68, 0x52, 0x1e, 0xc0, 0xac, 0xb8, 0x84, 0x86, 0xa6,
	0x5b, 0x77, 0x07, 0x88, 0xaf, 0xe7, 0x93, 0x99, 0x6e, 0xbb, 0x2d, 0x2d, 0x94, 0x10, 0x8d, 0x97,
	0xb3, 0x33, 0xfc, 0xbe, 0xf9, 0xf3, 0x51, 0x68, 0x87, 0x82, 0x23, 0x1f, 0xd0, 0x30, 0x14, 0x7c,
	0x43, 0xfd, 0x01, 0x0f, 0x99, 0xa0, 0xc8, 0x45, 0x5f, 0xbd, 0x93, 0xba, 0x4e, 0xd8, 0x37, 0x85,
	0x42, 0xb9, 0xa4, 0x82, 0xbd, 0xc7, 0x65, 0xce, 0x14, 0x2e, 0x5e, 0x92, 0x1f, 0x8e, 0xb9, 0x17,
	0xb8, 0xec, 0x63, 0xcd, 0x24, 0x12, 0x1b, 0xea, 0x9a, 0x67, 0x55, 0xba, 0x95, 0x5e, 0xc3, 0x4d,
	0xe3, 0x28, 0xb7, 0xe5, 0x62, 0xb5, 0xf0, 0xf9, 0xd6, 0xfa, 0x1f, 0xe7, 0x74, 0xec, 0x3c, 0x43,
	0x53, 0xe3, 0x26, 0x8c, 0x6e, 0xd8, 0x6f, 0x79, 0x2e, 0x58, 0x9a, 0x37, 0xa2, 0x38, 0x5f, 0x96,
	0xf7, 0x68, 0x9c, 0xcc, 0x9c, 0xc1, 0x75, 0x8e, 0x79, 0xa4, 0xd1, 0xd3, 0xa1, 0xaf, 0xd9, 0x1e,
	0x27, 0x9e, 0x44, 0x8d, 0x6b, 0x81, 0xc9, 0x17, 0x0b, 0xc9, 0x50, 0x4d, 0x6d, 0xb8, 0x49, 0x44,
	0x9a, 0x50, 0x9b, 0xf3, 0x75, 0x80, 0x8a, 0x63, 0xb8, 0x71, 0x90, 0x13, 0x30, 0x0a, 0x02, 0xc1,
	0xce, 0x66, 0x95, 0x80, 0x0c, 0x79, 0x20, 0x19, 0xe9, 0x81, 0x29, 0x91, 0xe2, 0x5a, 0x2a, 0x85,
	0xb3, 0xe1, 0x79, 0x5f, 0x1f, 0xba, 0x3f, 0x53, 0xef, 0x6e, 0x92, 0x8f, 0x34, 0x91, 0x23, 0xf5,
	0x95, 0x66, 0xd5, 0x8d, 0x83, 0xa8, 0x43, 0x16, 0xa0, 0x87, 0x9f, 0x96, 0xa1, 0xc6, 0x4d, 0x22,
	0x87, 0xc2, 0xa5, 0xd6, 0x7b, 0xf4, 0x7c, 0x64, 0xe2, 0xc7, 0x23, 0xa5, 0xfb, 0x34, 0xf2, 0x87,
	0x77, 0x42, 0x68, 0x15, 0x25, 0xfe, 0x76, 0xa8, 0xe1, 0x97, 0x01, 0x75, 0x2d, 0x49, 0x1e, 0xa0,
	0x1a, 0xd9, 0x89, 0xb4, 0x33, 0xf8, 0x81, 0xbf, 0x82, 0x7d, 0x95, 0xa5, 0x47, 0x3e, 0x0d, 0x56,
	0xba, 0x49, 0xe7, 0x1f, 0x19, 0x41, 0x4d, 0x99, 0x87, 0x74, 0xf6, 0x11, 0xbb, 0xae, 0x3a, 0xc6,
	0x18, 0x43, 0x23, 0x75, 0x36, 0x71, 0xf6, 0x39, 0x45, 0xdb, 0x1f, 0x63, 0x4d, 0x00, 0x32, 0x47,
	0x93, 0xbb, 0x12, 0xd8, 0xa9, 0x9d, 0x3d, 0x41, 0x35, 0x72, 0xda, 0xa1, 0xfd, 0xec, 0x58, 0xdc,
	0xee, 0x94, 0xa5, 0x53, 0xd0, 0x14, 0xcc, 0xf8, 0xbe, 0xe4, 0x76, 0xbf, 0x36, 0x67, 0x2e, 0xbb,
	0x5b, 0x5e, 0xa0, 0x71, 0x6f, 0xa6, 0xfa, 0x72, 0xdd, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x32,
	0xf1, 0x05, 0x7f, 0x01, 0x05, 0x00, 0x00,
}
