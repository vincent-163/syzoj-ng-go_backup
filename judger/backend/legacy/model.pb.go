// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package legacy

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

type ProblemData struct {
	TimeLimit            *int64   `protobuf:"varint,1,opt,name=time_limit,json=timeLimit" json:"time_limit,omitempty"`
	MemoryLimit          *int64   `protobuf:"varint,2,opt,name=memory_limit,json=memoryLimit" json:"memory_limit,omitempty"`
	FileIo               *bool    `protobuf:"varint,3,opt,name=file_io,json=fileIo" json:"file_io,omitempty"`
	FileIoInputName      *string  `protobuf:"bytes,4,opt,name=file_io_input_name,json=fileIoInputName" json:"file_io_input_name,omitempty"`
	FileIoOutputName     *string  `protobuf:"bytes,5,opt,name=file_io_output_name,json=fileIoOutputName" json:"file_io_output_name,omitempty"`
	Type                 *int32   `protobuf:"varint,6,opt,name=type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProblemData) Reset()         { *m = ProblemData{} }
func (m *ProblemData) String() string { return proto.CompactTextString(m) }
func (*ProblemData) ProtoMessage()    {}
func (*ProblemData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

func (m *ProblemData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProblemData.Unmarshal(m, b)
}
func (m *ProblemData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProblemData.Marshal(b, m, deterministic)
}
func (m *ProblemData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProblemData.Merge(m, src)
}
func (m *ProblemData) XXX_Size() int {
	return xxx_messageInfo_ProblemData.Size(m)
}
func (m *ProblemData) XXX_DiscardUnknown() {
	xxx_messageInfo_ProblemData.DiscardUnknown(m)
}

var xxx_messageInfo_ProblemData proto.InternalMessageInfo

func (m *ProblemData) GetTimeLimit() int64 {
	if m != nil && m.TimeLimit != nil {
		return *m.TimeLimit
	}
	return 0
}

func (m *ProblemData) GetMemoryLimit() int64 {
	if m != nil && m.MemoryLimit != nil {
		return *m.MemoryLimit
	}
	return 0
}

func (m *ProblemData) GetFileIo() bool {
	if m != nil && m.FileIo != nil {
		return *m.FileIo
	}
	return false
}

func (m *ProblemData) GetFileIoInputName() string {
	if m != nil && m.FileIoInputName != nil {
		return *m.FileIoInputName
	}
	return ""
}

func (m *ProblemData) GetFileIoOutputName() string {
	if m != nil && m.FileIoOutputName != nil {
		return *m.FileIoOutputName
	}
	return ""
}

func (m *ProblemData) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

type SubmissionContent struct {
	Language             *string  `protobuf:"bytes,1,opt,name=language" json:"language,omitempty"`
	Code                 *string  `protobuf:"bytes,2,opt,name=code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmissionContent) Reset()         { *m = SubmissionContent{} }
func (m *SubmissionContent) String() string { return proto.CompactTextString(m) }
func (*SubmissionContent) ProtoMessage()    {}
func (*SubmissionContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{1}
}

func (m *SubmissionContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmissionContent.Unmarshal(m, b)
}
func (m *SubmissionContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmissionContent.Marshal(b, m, deterministic)
}
func (m *SubmissionContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmissionContent.Merge(m, src)
}
func (m *SubmissionContent) XXX_Size() int {
	return xxx_messageInfo_SubmissionContent.Size(m)
}
func (m *SubmissionContent) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmissionContent.DiscardUnknown(m)
}

var xxx_messageInfo_SubmissionContent proto.InternalMessageInfo

func (m *SubmissionContent) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

func (m *SubmissionContent) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func init() {
	proto.RegisterType((*ProblemData)(nil), "syzoj.judge.legacy.ProblemData")
	proto.RegisterType((*SubmissionContent)(nil), "syzoj.judge.legacy.SubmissionContent")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xd0, 0xcf, 0x4a, 0xfc, 0x30,
	0x10, 0x07, 0x70, 0xf2, 0xdb, 0x3f, 0xbf, 0x6d, 0x56, 0x50, 0xe3, 0xc1, 0x22, 0x08, 0x75, 0x4f,
	0x05, 0xd9, 0x16, 0xc4, 0x27, 0x70, 0xbd, 0x2c, 0x88, 0x4a, 0xbc, 0x79, 0x29, 0x69, 0x3b, 0xc6,
	0xac, 0x49, 0x66, 0x69, 0x93, 0x43, 0x7d, 0x52, 0x1f, 0x47, 0x9a, 0xac, 0x7b, 0x09, 0x33, 0xdf,
	0xf9, 0x10, 0x98, 0xa1, 0x4b, 0x83, 0x2d, 0xe8, 0x62, 0xdf, 0xa1, 0x43, 0xc6, 0xfa, 0xe1, 0x1b,
	0x77, 0xc5, 0xce, 0xb7, 0x12, 0x0a, 0x0d, 0x52, 0x34, 0xc3, 0xea, 0x87, 0xd0, 0xe5, 0x6b, 0x87,
	0xb5, 0x06, 0xf3, 0x28, 0x9c, 0x60, 0xd7, 0x94, 0x3a, 0x65, 0xa0, 0xd2, 0xca, 0x28, 0x97, 0x92,
	0x8c, 0xe4, 0x13, 0x9e, 0x8c, 0xc9, 0xd3, 0x18, 0xb0, 0x1b, 0x7a, 0x62, 0xc0, 0x60, 0x37, 0x1c,
	0xc0, 0xbf, 0x00, 0x96, 0x31, 0x8b, 0xe4, 0x92, 0xfe, 0xff, 0x50, 0x1a, 0x2a, 0x85, 0xe9, 0x24,
	0x23, 0xf9, 0x82, 0xcf, 0xc7, 0x76, 0x8b, 0xec, 0x96, 0xb2, 0xc3, 0xa0, 0x52, 0x76, 0xef, 0x5d,
	0x65, 0x85, 0x81, 0x74, 0x9a, 0x91, 0x3c, 0xe1, 0xa7, 0xd1, 0x6c, 0xc7, 0xfc, 0x59, 0x18, 0x60,
	0x6b, 0x7a, 0xf1, 0x87, 0xd1, 0xbb, 0xa3, 0x9e, 0x05, 0x7d, 0x16, 0xf5, 0x4b, 0x18, 0x04, 0xce,
	0xe8, 0xd4, 0x0d, 0x7b, 0x48, 0xe7, 0x19, 0xc9, 0x67, 0x3c, 0xd4, 0xab, 0x0d, 0x3d, 0x7f, 0xf3,
	0xb5, 0x51, 0x7d, 0xaf, 0xd0, 0x6e, 0xd0, 0x3a, 0xb0, 0x8e, 0x5d, 0xd1, 0x85, 0x16, 0x56, 0x7a,
	0x21, 0x21, 0x6c, 0x97, 0xf0, 0x63, 0x3f, 0x7e, 0xd2, 0x60, 0x0b, 0x61, 0xa9, 0x84, 0x87, 0xfa,
	0xe1, 0xfe, 0xfd, 0x4e, 0x2a, 0xf7, 0xe9, 0xeb, 0xa2, 0x41, 0x53, 0x86, 0x03, 0xc6, 0x77, 0x6d,
	0xe5, 0x5a, 0x62, 0x19, 0x8e, 0xd9, 0x95, 0xb5, 0x68, 0xbe, 0xc0, 0xb6, 0x65, 0xbc, 0xea, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x25, 0x75, 0xcb, 0x77, 0x01, 0x00, 0x00,
}
