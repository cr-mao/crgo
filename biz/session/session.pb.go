// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session.proto

package session

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

type Session struct {
	Version              int64    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Anonymous            bool     `protobuf:"varint,2,opt,name=anonymous,proto3" json:"anonymous,omitempty"`
	UserID               int64    `protobuf:"varint,3,opt,name=userID,proto3" json:"userID,omitempty"`
	Guid                 string   `protobuf:"bytes,4,opt,name=guid,proto3" json:"guid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{0}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Session) GetAnonymous() bool {
	if m != nil {
		return m.Anonymous
	}
	return false
}

func (m *Session) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *Session) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func init() {
	proto.RegisterType((*Session)(nil), "session.Session")
}

func init() {
	proto.RegisterFile("session.proto", fileDescriptor_3a6be1b361fa6f14)
}

var fileDescriptor_3a6be1b361fa6f14 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2e,
	0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x72, 0xb9,
	0xd8, 0x83, 0x21, 0x4c, 0x21, 0x09, 0x2e, 0xf6, 0xb2, 0xd4, 0x22, 0x10, 0x53, 0x82, 0x51, 0x81,
	0x51, 0x83, 0x39, 0x08, 0xc6, 0x15, 0x92, 0xe1, 0xe2, 0x4c, 0xcc, 0xcb, 0xcf, 0xab, 0xcc, 0xcd,
	0x2f, 0x2d, 0x96, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x08, 0x42, 0x08, 0x08, 0x89, 0x71, 0xb1, 0x95,
	0x16, 0xa7, 0x16, 0x79, 0xba, 0x48, 0x30, 0x83, 0xb5, 0x41, 0x79, 0x42, 0x42, 0x5c, 0x2c, 0xe9,
	0xa5, 0x99, 0x29, 0x12, 0x2c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x13, 0x6f, 0x14, 0x77,
	0x52, 0x66, 0x95, 0x3e, 0xd4, 0xf6, 0x24, 0x36, 0xb0, 0x6b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x32, 0x57, 0x68, 0x17, 0x9e, 0x00, 0x00, 0x00,
}