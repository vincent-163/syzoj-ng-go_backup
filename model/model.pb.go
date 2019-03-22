// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/syzoj/protoc-gen-gotype/gotype"
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

type User struct {
	Id                   *UserRef  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	UserName             *string   `protobuf:"bytes,2,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Auth                 *UserAuth `protobuf:"bytes,3,opt,name=auth" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() UserRef {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *User) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *User) GetAuth() *UserAuth {
	if m != nil {
		return m.Auth
	}
	return nil
}

type UserAuth struct {
	PasswordHash         []byte   `protobuf:"bytes,1,opt,name=password_hash,json=passwordHash" json:"password_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAuth) Reset()         { *m = UserAuth{} }
func (m *UserAuth) String() string { return proto.CompactTextString(m) }
func (*UserAuth) ProtoMessage()    {}
func (*UserAuth) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{1}
}

func (m *UserAuth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAuth.Unmarshal(m, b)
}
func (m *UserAuth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAuth.Marshal(b, m, deterministic)
}
func (m *UserAuth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAuth.Merge(m, src)
}
func (m *UserAuth) XXX_Size() int {
	return xxx_messageInfo_UserAuth.Size(m)
}
func (m *UserAuth) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAuth.DiscardUnknown(m)
}

var xxx_messageInfo_UserAuth proto.InternalMessageInfo

func (m *UserAuth) GetPasswordHash() []byte {
	if m != nil {
		return m.PasswordHash
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "syzoj.server.User")
	proto.RegisterType((*UserAuth)(nil), "syzoj.server.UserAuth")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x29, 0xae, 0xac, 0xca, 0xcf, 0xd2, 0x2b,
	0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x92, 0xe2, 0x49, 0xcf, 0x2f, 0xa9, 0x2c, 0x48, 0x85, 0xc8, 0x29,
	0xe5, 0x70, 0xb1, 0x84, 0x16, 0xa7, 0x16, 0x09, 0x49, 0x73, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x3a, 0x71, 0x77, 0xcd, 0x8c, 0x64, 0x07, 0x89, 0x06, 0xa5, 0xa6, 0x05, 0x31,
	0x65, 0xa6, 0x08, 0x49, 0x73, 0x71, 0x96, 0x16, 0xa7, 0x16, 0xc5, 0xe7, 0x25, 0xe6, 0xa6, 0x4a,
	0x30, 0x81, 0xd4, 0x04, 0x71, 0x80, 0x04, 0xfc, 0x12, 0x73, 0x53, 0x85, 0xb4, 0xb8, 0x58, 0x12,
	0x4b, 0x4b, 0x32, 0x24, 0x98, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xc4, 0xf4, 0x90, 0x2d, 0xd3, 0x03,
	0x99, 0xe2, 0x58, 0x5a, 0x92, 0x11, 0x04, 0x56, 0xa3, 0xa4, 0xcf, 0xc5, 0x01, 0x13, 0x11, 0x52,
	0xe6, 0xe2, 0x2d, 0x48, 0x2c, 0x2e, 0x2e, 0xcf, 0x2f, 0x4a, 0x89, 0xcf, 0x48, 0x2c, 0xce, 0x00,
	0x5b, 0xce, 0x13, 0xc4, 0x03, 0x13, 0xf4, 0x48, 0x2c, 0xce, 0x70, 0x52, 0x89, 0x52, 0x4a, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x07, 0x1b, 0x0d, 0x21, 0x75, 0xf3, 0xd2,
	0x75, 0xd3, 0xf3, 0xf5, 0xc1, 0xde, 0x04, 0x04, 0x00, 0x00, 0xff, 0xff, 0x50, 0x78, 0x77, 0x6c,
	0xee, 0x00, 0x00, 0x00,
}
