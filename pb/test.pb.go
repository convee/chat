// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

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

type UserInfo struct {
	Message              *string  `protobuf:"bytes,1,req,name=message" json:"message,omitempty"`
	Length               *int32   `protobuf:"varint,2,req,name=length" json:"length,omitempty"`
	Cnt                  *int32   `protobuf:"varint,3,req,name=cnt" json:"cnt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_caa797a2fd944713, []int{0}
}
func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (dst *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(dst, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *UserInfo) GetLength() int32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *UserInfo) GetCnt() int32 {
	if m != nil && m.Cnt != nil {
		return *m.Cnt
	}
	return 0
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "pb.UserInfo")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_test_caa797a2fd944713) }

var fileDescriptor_test_caa797a2fd944713 = []byte{
	// 103 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xf2, 0xe3, 0xe2, 0x08, 0x2d,
	0x4e, 0x2d, 0xf2, 0xcc, 0x4b, 0xcb, 0x17, 0x92, 0xe0, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c,
	0x4f, 0x95, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0xc4, 0xb8, 0xd8, 0x72, 0x52,
	0xf3, 0xd2, 0x4b, 0x32, 0x24, 0x98, 0x14, 0x98, 0x34, 0x58, 0x83, 0xa0, 0x3c, 0x21, 0x01, 0x2e,
	0xe6, 0xe4, 0xbc, 0x12, 0x09, 0x66, 0xb0, 0x20, 0x88, 0x09, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdf,
	0x87, 0xd6, 0x97, 0x60, 0x00, 0x00, 0x00,
}
