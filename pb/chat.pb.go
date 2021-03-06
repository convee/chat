// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Cmd                  *string  `protobuf:"bytes,1,req,name=cmd" json:"cmd,omitempty"`
	Data                 *string  `protobuf:"bytes,2,req,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_70a72bf71233ac49, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetCmd() string {
	if m != nil && m.Cmd != nil {
		return *m.Cmd
	}
	return ""
}

func (m *Message) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

type Login struct {
	Username             *string  `protobuf:"bytes,1,req,name=username" json:"username,omitempty"`
	Pwd                  *string  `protobuf:"bytes,2,req,name=pwd" json:"pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Login) Reset()         { *m = Login{} }
func (m *Login) String() string { return proto.CompactTextString(m) }
func (*Login) ProtoMessage()    {}
func (*Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_70a72bf71233ac49, []int{1}
}
func (m *Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Login.Unmarshal(m, b)
}
func (m *Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Login.Marshal(b, m, deterministic)
}
func (dst *Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Login.Merge(dst, src)
}
func (m *Login) XXX_Size() int {
	return xxx_messageInfo_Login.Size(m)
}
func (m *Login) XXX_DiscardUnknown() {
	xxx_messageInfo_Login.DiscardUnknown(m)
}

var xxx_messageInfo_Login proto.InternalMessageInfo

func (m *Login) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *Login) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

type Register struct {
	Username             *string  `protobuf:"bytes,1,req,name=username" json:"username,omitempty"`
	Pwd                  *string  `protobuf:"bytes,2,req,name=pwd" json:"pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Register) Reset()         { *m = Register{} }
func (m *Register) String() string { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()    {}
func (*Register) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_70a72bf71233ac49, []int{2}
}
func (m *Register) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register.Unmarshal(m, b)
}
func (m *Register) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register.Marshal(b, m, deterministic)
}
func (dst *Register) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register.Merge(dst, src)
}
func (m *Register) XXX_Size() int {
	return xxx_messageInfo_Register.Size(m)
}
func (m *Register) XXX_DiscardUnknown() {
	xxx_messageInfo_Register.DiscardUnknown(m)
}

var xxx_messageInfo_Register proto.InternalMessageInfo

func (m *Register) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *Register) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "pb.Message")
	proto.RegisterType((*Login)(nil), "pb.Login")
	proto.RegisterType((*Register)(nil), "pb.Register")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_chat_70a72bf71233ac49) }

var fileDescriptor_chat_70a72bf71233ac49 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe7, 0x62, 0xf7, 0x4d,
	0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe0, 0x62, 0x4e, 0xce, 0x4d, 0x91, 0x60, 0x54, 0x60,
	0xd2, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x84, 0xb8, 0x58, 0x52, 0x12, 0x4b, 0x12, 0x25, 0x98, 0xc0,
	0x42, 0x60, 0xb6, 0x92, 0x29, 0x17, 0xab, 0x4f, 0x7e, 0x7a, 0x66, 0x9e, 0x90, 0x14, 0x17, 0x47,
	0x69, 0x71, 0x6a, 0x51, 0x5e, 0x62, 0x6e, 0x2a, 0x54, 0x0f, 0x9c, 0x0f, 0x32, 0xaa, 0xa0, 0x3c,
	0x05, 0xaa, 0x0f, 0xc4, 0x54, 0xb2, 0xe0, 0xe2, 0x08, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x49, 0x2d,
	0x22, 0x4d, 0x27, 0x20, 0x00, 0x00, 0xff, 0xff, 0xae, 0x08, 0x6d, 0x97, 0xb2, 0x00, 0x00, 0x00,
}
