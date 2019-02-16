// Code generated by protoc-gen-go. DO NOT EDIT.
// source: judge.proto

package judge

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
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

type JudgeRequest struct {
	JudgerId             string   `protobuf:"bytes,1,opt,name=judger_id,json=judgerId,proto3" json:"judger_id,omitempty"`
	JudgerToken          string   `protobuf:"bytes,2,opt,name=judger_token,json=judgerToken,proto3" json:"judger_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JudgeRequest) Reset()         { *m = JudgeRequest{} }
func (m *JudgeRequest) String() string { return proto.CompactTextString(m) }
func (*JudgeRequest) ProtoMessage()    {}
func (*JudgeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b4fa635fb2baaa, []int{0}
}

func (m *JudgeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JudgeRequest.Unmarshal(m, b)
}
func (m *JudgeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JudgeRequest.Marshal(b, m, deterministic)
}
func (m *JudgeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JudgeRequest.Merge(m, src)
}
func (m *JudgeRequest) XXX_Size() int {
	return xxx_messageInfo_JudgeRequest.Size(m)
}
func (m *JudgeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JudgeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JudgeRequest proto.InternalMessageInfo

func (m *JudgeRequest) GetJudgerId() string {
	if m != nil {
		return m.JudgerId
	}
	return ""
}

func (m *JudgeRequest) GetJudgerToken() string {
	if m != nil {
		return m.JudgerToken
	}
	return ""
}

type FetchTaskResult struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Task                 *Task    `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchTaskResult) Reset()         { *m = FetchTaskResult{} }
func (m *FetchTaskResult) String() string { return proto.CompactTextString(m) }
func (*FetchTaskResult) ProtoMessage()    {}
func (*FetchTaskResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b4fa635fb2baaa, []int{1}
}

func (m *FetchTaskResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchTaskResult.Unmarshal(m, b)
}
func (m *FetchTaskResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchTaskResult.Marshal(b, m, deterministic)
}
func (m *FetchTaskResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchTaskResult.Merge(m, src)
}
func (m *FetchTaskResult) XXX_Size() int {
	return xxx_messageInfo_FetchTaskResult.Size(m)
}
func (m *FetchTaskResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchTaskResult.DiscardUnknown(m)
}

var xxx_messageInfo_FetchTaskResult proto.InternalMessageInfo

func (m *FetchTaskResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *FetchTaskResult) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type Task struct {
	TaskTag              int64    `protobuf:"varint,1,opt,name=task_tag,json=taskTag,proto3" json:"task_tag,omitempty"`
	ProblemId            string   `protobuf:"bytes,2,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	Language             string   `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`
	Code                 string   `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b4fa635fb2baaa, []int{2}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetTaskTag() int64 {
	if m != nil {
		return m.TaskTag
	}
	return 0
}

func (m *Task) GetProblemId() string {
	if m != nil {
		return m.ProblemId
	}
	return ""
}

func (m *Task) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Task) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type TaskProgress struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskProgress) Reset()         { *m = TaskProgress{} }
func (m *TaskProgress) String() string { return proto.CompactTextString(m) }
func (*TaskProgress) ProtoMessage()    {}
func (*TaskProgress) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b4fa635fb2baaa, []int{3}
}

func (m *TaskProgress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskProgress.Unmarshal(m, b)
}
func (m *TaskProgress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskProgress.Marshal(b, m, deterministic)
}
func (m *TaskProgress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskProgress.Merge(m, src)
}
func (m *TaskProgress) XXX_Size() int {
	return xxx_messageInfo_TaskProgress.Size(m)
}
func (m *TaskProgress) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskProgress.DiscardUnknown(m)
}

var xxx_messageInfo_TaskProgress proto.InternalMessageInfo

type TaskResult struct {
	Result               string   `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	Score                float32  `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskResult) Reset()         { *m = TaskResult{} }
func (m *TaskResult) String() string { return proto.CompactTextString(m) }
func (*TaskResult) ProtoMessage()    {}
func (*TaskResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b4fa635fb2baaa, []int{4}
}

func (m *TaskResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResult.Unmarshal(m, b)
}
func (m *TaskResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResult.Marshal(b, m, deterministic)
}
func (m *TaskResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResult.Merge(m, src)
}
func (m *TaskResult) XXX_Size() int {
	return xxx_messageInfo_TaskResult.Size(m)
}
func (m *TaskResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResult.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResult proto.InternalMessageInfo

func (m *TaskResult) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *TaskResult) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

type SetTaskResultMessage struct {
	JudgerId             string               `protobuf:"bytes,1,opt,name=judger_id,json=judgerId,proto3" json:"judger_id,omitempty"`
	JudgerToken          string               `protobuf:"bytes,2,opt,name=judger_token,json=judgerToken,proto3" json:"judger_token,omitempty"`
	TaskTag              int64                `protobuf:"varint,3,opt,name=task_tag,json=taskTag,proto3" json:"task_tag,omitempty"`
	Result               *TaskResult          `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	Myint                *wrappers.Int64Value `protobuf:"bytes,5,opt,name=myint,proto3" json:"myint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SetTaskResultMessage) Reset()         { *m = SetTaskResultMessage{} }
func (m *SetTaskResultMessage) String() string { return proto.CompactTextString(m) }
func (*SetTaskResultMessage) ProtoMessage()    {}
func (*SetTaskResultMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b4fa635fb2baaa, []int{5}
}

func (m *SetTaskResultMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTaskResultMessage.Unmarshal(m, b)
}
func (m *SetTaskResultMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTaskResultMessage.Marshal(b, m, deterministic)
}
func (m *SetTaskResultMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTaskResultMessage.Merge(m, src)
}
func (m *SetTaskResultMessage) XXX_Size() int {
	return xxx_messageInfo_SetTaskResultMessage.Size(m)
}
func (m *SetTaskResultMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTaskResultMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SetTaskResultMessage proto.InternalMessageInfo

func (m *SetTaskResultMessage) GetJudgerId() string {
	if m != nil {
		return m.JudgerId
	}
	return ""
}

func (m *SetTaskResultMessage) GetJudgerToken() string {
	if m != nil {
		return m.JudgerToken
	}
	return ""
}

func (m *SetTaskResultMessage) GetTaskTag() int64 {
	if m != nil {
		return m.TaskTag
	}
	return 0
}

func (m *SetTaskResultMessage) GetResult() *TaskResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *SetTaskResultMessage) GetMyint() *wrappers.Int64Value {
	if m != nil {
		return m.Myint
	}
	return nil
}

func init() {
	proto.RegisterType((*JudgeRequest)(nil), "JudgeRequest")
	proto.RegisterType((*FetchTaskResult)(nil), "FetchTaskResult")
	proto.RegisterType((*Task)(nil), "Task")
	proto.RegisterType((*TaskProgress)(nil), "TaskProgress")
	proto.RegisterType((*TaskResult)(nil), "TaskResult")
	proto.RegisterType((*SetTaskResultMessage)(nil), "SetTaskResultMessage")
}

func init() { proto.RegisterFile("judge.proto", fileDescriptor_a2b4fa635fb2baaa) }

var fileDescriptor_a2b4fa635fb2baaa = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0xed, 0x36, 0xbb, 0x6d, 0x32, 0x49, 0x28, 0x1a, 0x95, 0x6a, 0xbb, 0x11, 0x08, 0xcc, 0xa5,
	0x27, 0x57, 0x14, 0xc4, 0x01, 0x71, 0xe0, 0x42, 0xa5, 0x20, 0x81, 0xd0, 0x12, 0x71, 0xad, 0x9c,
	0xdd, 0xc1, 0x94, 0x6c, 0xd6, 0x8b, 0xed, 0x15, 0xea, 0x97, 0xf1, 0x0d, 0xfc, 0x15, 0xb2, 0xbd,
	0x9b, 0x26, 0x05, 0x4e, 0xdc, 0xfc, 0xde, 0x78, 0xe6, 0x8d, 0x9f, 0x1f, 0x8c, 0xbf, 0xb5, 0xa5,
	0x24, 0xde, 0x68, 0x65, 0x55, 0x36, 0x93, 0x4a, 0xc9, 0x8a, 0xce, 0x3d, 0x5a, 0xb6, 0x5f, 0xce,
	0x69, 0xdd, 0xd8, 0x9b, 0xae, 0xf8, 0xe8, 0x6e, 0xf1, 0x87, 0x16, 0x4d, 0x43, 0xda, 0x84, 0x3a,
	0xfb, 0x00, 0x93, 0x77, 0x6e, 0x56, 0x4e, 0xdf, 0x5b, 0x32, 0x16, 0x67, 0x30, 0xf2, 0xb3, 0xf5,
	0xd5, 0x75, 0x99, 0x46, 0x8f, 0xa3, 0xb3, 0x51, 0x3e, 0x0c, 0xc4, 0xbc, 0xc4, 0x27, 0x30, 0xe9,
	0x8a, 0x56, 0xad, 0xa8, 0x4e, 0xf7, 0x7d, 0x3d, 0x2c, 0xa3, 0x17, 0x8e, 0x62, 0x97, 0x70, 0x74,
	0x49, 0xb6, 0xf8, 0xba, 0x10, 0x66, 0x95, 0x93, 0x69, 0x2b, 0x8b, 0x29, 0x1c, 0x9a, 0xb6, 0x28,
	0xc8, 0x18, 0x3f, 0x70, 0x98, 0xf7, 0x10, 0x4f, 0x21, 0xb6, 0xc2, 0xac, 0xfc, 0x9c, 0xf1, 0x45,
	0xc2, 0x7d, 0x93, 0xa7, 0x58, 0x03, 0xb1, 0x43, 0x78, 0x0a, 0x43, 0x87, 0xaf, 0xac, 0x90, 0xbe,
	0x7b, 0x90, 0x1f, 0x3a, 0xbc, 0x10, 0x12, 0x1f, 0x02, 0x34, 0x5a, 0x2d, 0x2b, 0x5a, 0xbb, 0x5d,
	0xc3, 0x2e, 0xa3, 0x8e, 0x99, 0x97, 0x98, 0xc1, 0xb0, 0x12, 0xb5, 0x6c, 0x85, 0xa4, 0x74, 0x10,
	0x1e, 0xd2, 0x63, 0x44, 0x88, 0x0b, 0x55, 0x52, 0x1a, 0x7b, 0xde, 0x9f, 0xd9, 0x3d, 0x98, 0x38,
	0xc5, 0x8f, 0x5a, 0x49, 0x4d, 0xc6, 0xb0, 0x57, 0x00, 0x5b, 0x8f, 0x38, 0x81, 0x03, 0xed, 0x4f,
	0x9d, 0x50, 0x87, 0xf0, 0x18, 0x12, 0x53, 0x28, 0x1d, 0x24, 0xf6, 0xf3, 0x00, 0xd8, 0xaf, 0x08,
	0x8e, 0x3f, 0x91, 0xbd, 0xed, 0x7f, 0x4f, 0xc6, 0x38, 0xe1, 0xff, 0xb4, 0x77, 0xc7, 0x8e, 0xc1,
	0xae, 0x1d, 0x4f, 0x37, 0x1b, 0xc6, 0xde, 0xce, 0x31, 0xbf, 0x95, 0xdf, 0xac, 0xfb, 0x0c, 0x92,
	0xf5, 0xcd, 0x75, 0x6d, 0xd3, 0xc4, 0xdf, 0x99, 0xf1, 0x10, 0x0f, 0xde, 0xc7, 0x83, 0xcf, 0x6b,
	0xfb, 0xf2, 0xc5, 0x67, 0x51, 0xb5, 0x94, 0x87, 0x9b, 0x17, 0x3f, 0x23, 0x48, 0x7c, 0x44, 0x90,
	0xc3, 0x68, 0xf3, 0xb7, 0x38, 0xe5, 0xdb, 0xb9, 0xc9, 0xee, 0xf3, 0x3b, 0xdf, 0xce, 0xf6, 0xf0,
	0x35, 0x1c, 0x75, 0x26, 0xf4, 0xa6, 0xe2, 0x94, 0x6f, 0xc3, 0xec, 0xe4, 0x0f, 0xfd, 0xb7, 0x2e,
	0xbb, 0x6c, 0xef, 0x2c, 0xc2, 0x37, 0x30, 0xdd, 0xb1, 0x10, 0x1f, 0xf0, 0xbf, 0x59, 0xfa, 0xef,
	0x19, 0xcb, 0x03, 0xcf, 0x3c, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x8f, 0x86, 0xda, 0x2e,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JudgeClient is the client API for Judge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JudgeClient interface {
	FetchTask(ctx context.Context, in *JudgeRequest, opts ...grpc.CallOption) (*FetchTaskResult, error)
	SetTaskProgress(ctx context.Context, opts ...grpc.CallOption) (Judge_SetTaskProgressClient, error)
	SetTaskResult(ctx context.Context, in *SetTaskResultMessage, opts ...grpc.CallOption) (*empty.Empty, error)
}

type judgeClient struct {
	cc *grpc.ClientConn
}

func NewJudgeClient(cc *grpc.ClientConn) JudgeClient {
	return &judgeClient{cc}
}

func (c *judgeClient) FetchTask(ctx context.Context, in *JudgeRequest, opts ...grpc.CallOption) (*FetchTaskResult, error) {
	out := new(FetchTaskResult)
	err := c.cc.Invoke(ctx, "/Judge/FetchTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeClient) SetTaskProgress(ctx context.Context, opts ...grpc.CallOption) (Judge_SetTaskProgressClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Judge_serviceDesc.Streams[0], "/Judge/SetTaskProgress", opts...)
	if err != nil {
		return nil, err
	}
	x := &judgeSetTaskProgressClient{stream}
	return x, nil
}

type Judge_SetTaskProgressClient interface {
	Send(*TaskProgress) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type judgeSetTaskProgressClient struct {
	grpc.ClientStream
}

func (x *judgeSetTaskProgressClient) Send(m *TaskProgress) error {
	return x.ClientStream.SendMsg(m)
}

func (x *judgeSetTaskProgressClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *judgeClient) SetTaskResult(ctx context.Context, in *SetTaskResultMessage, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Judge/SetTaskResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JudgeServer is the server API for Judge service.
type JudgeServer interface {
	FetchTask(context.Context, *JudgeRequest) (*FetchTaskResult, error)
	SetTaskProgress(Judge_SetTaskProgressServer) error
	SetTaskResult(context.Context, *SetTaskResultMessage) (*empty.Empty, error)
}

func RegisterJudgeServer(s *grpc.Server, srv JudgeServer) {
	s.RegisterService(&_Judge_serviceDesc, srv)
}

func _Judge_FetchTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JudgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServer).FetchTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Judge/FetchTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServer).FetchTask(ctx, req.(*JudgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Judge_SetTaskProgress_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(JudgeServer).SetTaskProgress(&judgeSetTaskProgressServer{stream})
}

type Judge_SetTaskProgressServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*TaskProgress, error)
	grpc.ServerStream
}

type judgeSetTaskProgressServer struct {
	grpc.ServerStream
}

func (x *judgeSetTaskProgressServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *judgeSetTaskProgressServer) Recv() (*TaskProgress, error) {
	m := new(TaskProgress)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Judge_SetTaskResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTaskResultMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServer).SetTaskResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Judge/SetTaskResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServer).SetTaskResult(ctx, req.(*SetTaskResultMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Judge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Judge",
	HandlerType: (*JudgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchTask",
			Handler:    _Judge_FetchTask_Handler,
		},
		{
			MethodName: "SetTaskResult",
			Handler:    _Judge_SetTaskResult_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SetTaskProgress",
			Handler:       _Judge_SetTaskProgress_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "judge.proto",
}
