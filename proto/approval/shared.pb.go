// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/approval/shared.proto

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

// 工作流模式
type WorkflowMode int32

const (
	WorkflowMode_WORKFLOW_MODE_INVALID  WorkflowMode = 0
	WorkflowMode_WORKFLOW_MODE_ALL      WorkflowMode = 1
	WorkflowMode_WORKFLOW_MODE_ANY      WorkflowMode = 2
	WorkflowMode_WORKFLOW_MODE_MAJORITY WorkflowMode = 3
)

var WorkflowMode_name = map[int32]string{
	0: "WORKFLOW_MODE_INVALID",
	1: "WORKFLOW_MODE_ALL",
	2: "WORKFLOW_MODE_ANY",
	3: "WORKFLOW_MODE_MAJORITY",
}

var WorkflowMode_value = map[string]int32{
	"WORKFLOW_MODE_INVALID":  0,
	"WORKFLOW_MODE_ALL":      1,
	"WORKFLOW_MODE_ANY":      2,
	"WORKFLOW_MODE_MAJORITY": 3,
}

func (x WorkflowMode) String() string {
	return proto.EnumName(WorkflowMode_name, int32(x))
}

func (WorkflowMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aaa2e0860ed5e576, []int{0}
}

// 任务状态
type TaskStatus int32

const (
	TaskStatus_TASK_STATUS_INVALID  TaskStatus = 0
	TaskStatus_TASK_STATUS_PENDING  TaskStatus = 1
	TaskStatus_TASK_STATUS_ACCEPTED TaskStatus = 2
	TaskStatus_TASK_STATUS_REJECTED TaskStatus = 3
)

var TaskStatus_name = map[int32]string{
	0: "TASK_STATUS_INVALID",
	1: "TASK_STATUS_PENDING",
	2: "TASK_STATUS_ACCEPTED",
	3: "TASK_STATUS_REJECTED",
}

var TaskStatus_value = map[string]int32{
	"TASK_STATUS_INVALID":  0,
	"TASK_STATUS_PENDING":  1,
	"TASK_STATUS_ACCEPTED": 2,
	"TASK_STATUS_REJECTED": 3,
}

func (x TaskStatus) String() string {
	return proto.EnumName(TaskStatus_name, int32(x))
}

func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aaa2e0860ed5e576, []int{1}
}

// 记录状态
type ActionStatus int32

const (
	ActionStatus_ACTION_STATUS_INVALID  ActionStatus = 0
	ActionStatus_ACTION_STATUS_ACCEPTED ActionStatus = 1
	ActionStatus_ACTION_STATUS_REJECTED ActionStatus = 2
)

var ActionStatus_name = map[int32]string{
	0: "ACTION_STATUS_INVALID",
	1: "ACTION_STATUS_ACCEPTED",
	2: "ACTION_STATUS_REJECTED",
}

var ActionStatus_value = map[string]int32{
	"ACTION_STATUS_INVALID":  0,
	"ACTION_STATUS_ACCEPTED": 1,
	"ACTION_STATUS_REJECTED": 2,
}

func (x ActionStatus) String() string {
	return proto.EnumName(ActionStatus_name, int32(x))
}

func (ActionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aaa2e0860ed5e576, []int{2}
}

// 状态
type Status struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa2e0860ed5e576, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// 空白请求
type BlankRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlankRequest) Reset()         { *m = BlankRequest{} }
func (m *BlankRequest) String() string { return proto.CompactTextString(m) }
func (*BlankRequest) ProtoMessage()    {}
func (*BlankRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa2e0860ed5e576, []int{1}
}

func (m *BlankRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlankRequest.Unmarshal(m, b)
}
func (m *BlankRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlankRequest.Marshal(b, m, deterministic)
}
func (m *BlankRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlankRequest.Merge(m, src)
}
func (m *BlankRequest) XXX_Size() int {
	return xxx_messageInfo_BlankRequest.Size(m)
}
func (m *BlankRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlankRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlankRequest proto.InternalMessageInfo

// 空白回复
type BlankResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlankResponse) Reset()         { *m = BlankResponse{} }
func (m *BlankResponse) String() string { return proto.CompactTextString(m) }
func (*BlankResponse) ProtoMessage()    {}
func (*BlankResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa2e0860ed5e576, []int{2}
}

func (m *BlankResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlankResponse.Unmarshal(m, b)
}
func (m *BlankResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlankResponse.Marshal(b, m, deterministic)
}
func (m *BlankResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlankResponse.Merge(m, src)
}
func (m *BlankResponse) XXX_Size() int {
	return xxx_messageInfo_BlankResponse.Size(m)
}
func (m *BlankResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlankResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlankResponse proto.InternalMessageInfo

func (m *BlankResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// 工作流实体
type WorkflowEntity struct {
	Uuid                 string       `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Mode                 WorkflowMode `protobuf:"varint,3,opt,name=mode,proto3,enum=approval.WorkflowMode" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *WorkflowEntity) Reset()         { *m = WorkflowEntity{} }
func (m *WorkflowEntity) String() string { return proto.CompactTextString(m) }
func (*WorkflowEntity) ProtoMessage()    {}
func (*WorkflowEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa2e0860ed5e576, []int{3}
}

func (m *WorkflowEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowEntity.Unmarshal(m, b)
}
func (m *WorkflowEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowEntity.Marshal(b, m, deterministic)
}
func (m *WorkflowEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowEntity.Merge(m, src)
}
func (m *WorkflowEntity) XXX_Size() int {
	return xxx_messageInfo_WorkflowEntity.Size(m)
}
func (m *WorkflowEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowEntity.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowEntity proto.InternalMessageInfo

func (m *WorkflowEntity) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *WorkflowEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkflowEntity) GetMode() WorkflowMode {
	if m != nil {
		return m.Mode
	}
	return WorkflowMode_WORKFLOW_MODE_INVALID
}

// 任务实体
type TaskEntity struct {
	Uuid                 string     `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Subject              string     `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Body                 string     `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Meta                 string     `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
	State                TaskStatus `protobuf:"varint,5,opt,name=state,proto3,enum=approval.TaskStatus" json:"state,omitempty"`
	Workflow             string     `protobuf:"bytes,6,opt,name=workflow,proto3" json:"workflow,omitempty"`
	UpdatedAt            int64      `protobuf:"varint,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TaskEntity) Reset()         { *m = TaskEntity{} }
func (m *TaskEntity) String() string { return proto.CompactTextString(m) }
func (*TaskEntity) ProtoMessage()    {}
func (*TaskEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa2e0860ed5e576, []int{4}
}

func (m *TaskEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskEntity.Unmarshal(m, b)
}
func (m *TaskEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskEntity.Marshal(b, m, deterministic)
}
func (m *TaskEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskEntity.Merge(m, src)
}
func (m *TaskEntity) XXX_Size() int {
	return xxx_messageInfo_TaskEntity.Size(m)
}
func (m *TaskEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskEntity.DiscardUnknown(m)
}

var xxx_messageInfo_TaskEntity proto.InternalMessageInfo

func (m *TaskEntity) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *TaskEntity) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *TaskEntity) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *TaskEntity) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *TaskEntity) GetState() TaskStatus {
	if m != nil {
		return m.State
	}
	return TaskStatus_TASK_STATUS_INVALID
}

func (m *TaskEntity) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

func (m *TaskEntity) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

// 记录实体
type ActionEntity struct {
	Uuid                 string       `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Task                 string       `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	Operator             string       `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	State                ActionStatus `protobuf:"varint,4,opt,name=state,proto3,enum=approval.ActionStatus" json:"state,omitempty"`
	UpdatedAt            int64        `protobuf:"varint,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ActionEntity) Reset()         { *m = ActionEntity{} }
func (m *ActionEntity) String() string { return proto.CompactTextString(m) }
func (*ActionEntity) ProtoMessage()    {}
func (*ActionEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa2e0860ed5e576, []int{5}
}

func (m *ActionEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionEntity.Unmarshal(m, b)
}
func (m *ActionEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionEntity.Marshal(b, m, deterministic)
}
func (m *ActionEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionEntity.Merge(m, src)
}
func (m *ActionEntity) XXX_Size() int {
	return xxx_messageInfo_ActionEntity.Size(m)
}
func (m *ActionEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionEntity.DiscardUnknown(m)
}

var xxx_messageInfo_ActionEntity proto.InternalMessageInfo

func (m *ActionEntity) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ActionEntity) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

func (m *ActionEntity) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ActionEntity) GetState() ActionStatus {
	if m != nil {
		return m.State
	}
	return ActionStatus_ACTION_STATUS_INVALID
}

func (m *ActionEntity) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func init() {
	proto.RegisterEnum("approval.WorkflowMode", WorkflowMode_name, WorkflowMode_value)
	proto.RegisterEnum("approval.TaskStatus", TaskStatus_name, TaskStatus_value)
	proto.RegisterEnum("approval.ActionStatus", ActionStatus_name, ActionStatus_value)
	proto.RegisterType((*Status)(nil), "approval.Status")
	proto.RegisterType((*BlankRequest)(nil), "approval.BlankRequest")
	proto.RegisterType((*BlankResponse)(nil), "approval.BlankResponse")
	proto.RegisterType((*WorkflowEntity)(nil), "approval.WorkflowEntity")
	proto.RegisterType((*TaskEntity)(nil), "approval.TaskEntity")
	proto.RegisterType((*ActionEntity)(nil), "approval.ActionEntity")
}

func init() {
	proto.RegisterFile("proto/approval/shared.proto", fileDescriptor_aaa2e0860ed5e576)
}

var fileDescriptor_aaa2e0860ed5e576 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x5f, 0x6f, 0xda, 0x3c,
	0x14, 0xc6, 0x5f, 0xf3, 0xb7, 0x9c, 0x97, 0xa1, 0xcc, 0x2b, 0x2c, 0x65, 0xbb, 0x40, 0xb9, 0x42,
	0x68, 0xa2, 0x52, 0x27, 0x4d, 0xda, 0x65, 0x06, 0xd9, 0x44, 0x0b, 0xa4, 0x32, 0xd9, 0x50, 0xaf,
	0x90, 0x21, 0xde, 0xc6, 0x80, 0x38, 0x8b, 0x9d, 0x55, 0xfd, 0x32, 0xfb, 0x3c, 0xfb, 0x58, 0x93,
	0x9d, 0x04, 0x08, 0x9d, 0x7a, 0x77, 0xce, 0xf3, 0x98, 0xf3, 0x7b, 0x8e, 0x4d, 0xe0, 0x55, 0x18,
	0x71, 0xc9, 0x2f, 0x69, 0x18, 0x46, 0xfc, 0x17, 0xdd, 0x5e, 0x8a, 0xef, 0x34, 0x62, 0x7e, 0x5f,
	0xab, 0xf8, 0x2c, 0x93, 0xad, 0x77, 0x50, 0x99, 0x49, 0x2a, 0x63, 0x81, 0x31, 0x94, 0x56, 0xdc,
	0x67, 0x26, 0xea, 0xa0, 0x6e, 0x99, 0xe8, 0x1a, 0x9b, 0x50, 0xdd, 0x31, 0x21, 0xe8, 0x37, 0x66,
	0x16, 0x3a, 0xa8, 0x5b, 0x23, 0x59, 0x6b, 0x35, 0xa0, 0xfe, 0x61, 0x4b, 0x83, 0x0d, 0x61, 0x3f,
	0x63, 0x26, 0xa4, 0xf5, 0x1e, 0x9e, 0xa5, 0xbd, 0x08, 0x79, 0x20, 0x18, 0xee, 0x42, 0x45, 0xe8,
	0xc1, 0x7a, 0xe0, 0xff, 0x57, 0x46, 0x3f, 0x63, 0xf6, 0x13, 0x20, 0x49, 0x7d, 0xcb, 0x87, 0xc6,
	0x9c, 0x47, 0x9b, 0xaf, 0x5b, 0x7e, 0xef, 0x04, 0x72, 0x2d, 0x1f, 0x54, 0x94, 0x38, 0x5e, 0xfb,
	0xfa, 0x97, 0x35, 0xa2, 0x6b, 0xa5, 0x05, 0x74, 0x97, 0xe5, 0xd0, 0x35, 0xee, 0x41, 0x69, 0xa7,
	0x22, 0x17, 0x3b, 0xa8, 0xdb, 0xb8, 0x6a, 0x1d, 0x08, 0xd9, 0xbc, 0x09, 0xf7, 0x19, 0xd1, 0x67,
	0xac, 0x3f, 0x08, 0xc0, 0xa3, 0x62, 0xf3, 0x04, 0xc2, 0x84, 0xaa, 0x88, 0x97, 0x3f, 0xd8, 0x4a,
	0x66, 0xdb, 0xa6, 0xad, 0x3a, 0xbd, 0xe4, 0xfe, 0x83, 0x06, 0xd5, 0x88, 0xae, 0x95, 0xb6, 0x63,
	0x92, 0x9a, 0xa5, 0x44, 0x53, 0x35, 0xee, 0x41, 0x59, 0x2d, 0xc5, 0xcc, 0xb2, 0x4e, 0x74, 0x7e,
	0x48, 0xa4, 0xd0, 0xe9, 0xde, 0xc9, 0x11, 0xdc, 0x86, 0xb3, 0xfb, 0x34, 0xa6, 0x59, 0xd1, 0x33,
	0xf6, 0x3d, 0x7e, 0x0d, 0xb5, 0x38, 0xf4, 0xa9, 0x64, 0xbe, 0x2d, 0xcd, 0x6a, 0x07, 0x75, 0x8b,
	0xe4, 0x20, 0x58, 0xbf, 0x11, 0xd4, 0xed, 0x95, 0x5c, 0xf3, 0xe0, 0xe9, 0xfb, 0x92, 0x54, 0x6c,
	0xb2, 0xfb, 0x52, 0xb5, 0x42, 0xf2, 0x90, 0x45, 0x54, 0xf2, 0x28, 0x5d, 0x65, 0xdf, 0xe3, 0x37,
	0x59, 0xf4, 0xd2, 0xe9, 0x65, 0x26, 0xa8, 0x7c, 0xf8, 0x5c, 0xc0, 0xf2, 0x49, 0xc0, 0x9e, 0x80,
	0xfa, 0xf1, 0x0b, 0xe0, 0x0b, 0x68, 0xce, 0x5d, 0x72, 0xf3, 0x71, 0xec, 0xce, 0x17, 0x13, 0x77,
	0xe8, 0x2c, 0x46, 0xd3, 0x2f, 0xf6, 0x78, 0x34, 0x34, 0xfe, 0xc3, 0x4d, 0x78, 0x9e, 0xb7, 0xec,
	0xf1, 0xd8, 0x40, 0xff, 0x90, 0xa7, 0x77, 0x46, 0x01, 0xb7, 0xa1, 0x95, 0x97, 0x27, 0xf6, 0xb5,
	0x4b, 0x46, 0xde, 0x9d, 0x51, 0xec, 0x45, 0xc9, 0xfb, 0xa6, 0xff, 0xe6, 0x97, 0xf0, 0xc2, 0xb3,
	0x67, 0x37, 0x8b, 0x99, 0x67, 0x7b, 0x9f, 0x67, 0x47, 0xc0, 0x13, 0xe3, 0xd6, 0x99, 0x0e, 0x47,
	0xd3, 0x4f, 0x06, 0xc2, 0x26, 0x9c, 0x1f, 0x1b, 0xf6, 0x60, 0xe0, 0xdc, 0x7a, 0xce, 0xd0, 0x28,
	0x9c, 0x3a, 0xc4, 0xb9, 0x76, 0x06, 0xca, 0x29, 0xf6, 0x68, 0xf6, 0x10, 0x29, 0xf5, 0x02, 0x9a,
	0xf6, 0xc0, 0x1b, 0xb9, 0xd3, 0xc7, 0xdc, 0x36, 0xb4, 0xf2, 0xd6, 0x1e, 0x80, 0x1e, 0x7b, 0x7b,
	0x44, 0x61, 0x59, 0xd1, 0x5f, 0xec, 0xdb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xde, 0x56, 0xcb,
	0xbc, 0xd0, 0x03, 0x00, 0x00,
}
