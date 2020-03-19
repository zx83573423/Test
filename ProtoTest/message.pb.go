// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package ProtoTest

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

type A struct {
	MNA                  int32    `protobuf:"varint,1,opt,name=m_nA,json=mNA,proto3" json:"m_nA,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *A) Reset()         { *m = A{} }
func (m *A) String() string { return proto.CompactTextString(m) }
func (*A) ProtoMessage()    {}
func (*A) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *A) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_A.Unmarshal(m, b)
}
func (m *A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_A.Marshal(b, m, deterministic)
}
func (m *A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_A.Merge(m, src)
}
func (m *A) XXX_Size() int {
	return xxx_messageInfo_A.Size(m)
}
func (m *A) XXX_DiscardUnknown() {
	xxx_messageInfo_A.DiscardUnknown(m)
}

var xxx_messageInfo_A proto.InternalMessageInfo

func (m *A) GetMNA() int32 {
	if m != nil {
		return m.MNA
	}
	return 0
}

type T struct {
	MMapTest             map[int32]*A `protobuf:"bytes,1,rep,name=m_mapTest,json=mMapTest,proto3" json:"m_mapTest,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *T) Reset()         { *m = T{} }
func (m *T) String() string { return proto.CompactTextString(m) }
func (*T) ProtoMessage()    {}
func (*T) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *T) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_T.Unmarshal(m, b)
}
func (m *T) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_T.Marshal(b, m, deterministic)
}
func (m *T) XXX_Merge(src proto.Message) {
	xxx_messageInfo_T.Merge(m, src)
}
func (m *T) XXX_Size() int {
	return xxx_messageInfo_T.Size(m)
}
func (m *T) XXX_DiscardUnknown() {
	xxx_messageInfo_T.DiscardUnknown(m)
}

var xxx_messageInfo_T proto.InternalMessageInfo

func (m *T) GetMMapTest() map[int32]*A {
	if m != nil {
		return m.MMapTest
	}
	return nil
}

func init() {
	proto.RegisterType((*A)(nil), "ProtoTest.A")
	proto.RegisterType((*T)(nil), "ProtoTest.T")
	proto.RegisterMapType((map[int32]*A)(nil), "ProtoTest.T.MMapTestEntry")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x0c, 0x00, 0x51, 0x21, 0xa9,
	0xc5, 0x25, 0x4a, 0x62, 0x5c, 0x8c, 0x8e, 0x42, 0x82, 0x5c, 0x2c, 0xb9, 0xf1, 0x79, 0x8e, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0xcc, 0xb9, 0x7e, 0x8e, 0x4a, 0xed, 0x8c, 0x5c, 0x8c, 0x21,
	0x42, 0xe6, 0x5c, 0x9c, 0xb9, 0xf1, 0xb9, 0x89, 0x05, 0x20, 0xa5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a,
	0xdc, 0x46, 0x52, 0x7a, 0x70, 0xcd, 0x7a, 0x21, 0x7a, 0xbe, 0xbe, 0x10, 0x49, 0xd7, 0xbc, 0x92,
	0xa2, 0xca, 0x20, 0x8e, 0x5c, 0x28, 0x57, 0xca, 0x93, 0x8b, 0x17, 0x45, 0x4a, 0x48, 0x80, 0x8b,
	0x39, 0x3b, 0xb5, 0x12, 0x66, 0x43, 0x76, 0x6a, 0xa5, 0x90, 0x12, 0x17, 0x6b, 0x59, 0x62, 0x4e,
	0x69, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0xb7, 0x11, 0x0f, 0x92, 0xb9, 0x8e, 0x41, 0x10, 0x29,
	0x2b, 0x26, 0x0b, 0xc6, 0x24, 0x36, 0xb0, 0x9b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x79,
	0xd4, 0xe4, 0x99, 0xc4, 0x00, 0x00, 0x00,
}