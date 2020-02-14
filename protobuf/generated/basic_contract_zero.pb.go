// Code generated by protoc-gen-go. DO NOT EDIT.
// source: basic_contract_zero.proto

package client

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

//basic_contract_zero
type AddressList struct {
	Value                []*Address `protobuf:"bytes,1,rep,name=value,proto3" json:"value" form:"value"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AddressList) Reset()         { *m = AddressList{} }
func (m *AddressList) String() string { return proto.CompactTextString(m) }
func (*AddressList) ProtoMessage()    {}
func (*AddressList) Descriptor() ([]byte, []int) {
	return fileDescriptor_feca40e9e32ac99b, []int{0}
}

func (m *AddressList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressList.Unmarshal(m, b)
}
func (m *AddressList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressList.Marshal(b, m, deterministic)
}
func (m *AddressList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressList.Merge(m, src)
}
func (m *AddressList) XXX_Size() int {
	return xxx_messageInfo_AddressList.Size(m)
}
func (m *AddressList) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressList.DiscardUnknown(m)
}

var xxx_messageInfo_AddressList proto.InternalMessageInfo

func (m *AddressList) GetValue() []*Address {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*AddressList)(nil), "client.AddressList")
}

func init() { proto.RegisterFile("basic_contract_zero.proto", fileDescriptor_feca40e9e32ac99b) }

var fileDescriptor_feca40e9e32ac99b = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x4a, 0x2c, 0xce,
	0x4c, 0x8e, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0x89, 0xaf, 0x4a, 0x2d, 0xca, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x91, 0xe2, 0x81,
	0xd0, 0x10, 0x51, 0x25, 0x13, 0x2e, 0x6e, 0xc7, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0x62, 0x9f, 0xcc,
	0xe2, 0x12, 0x21, 0x55, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05, 0x66, 0x0d,
	0x6e, 0x23, 0x7e, 0x3d, 0xa8, 0x62, 0xa8, 0x9a, 0x20, 0x88, 0x6c, 0x12, 0x1b, 0x58, 0xb3, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x1d, 0xdf, 0x58, 0x6f, 0x00, 0x00, 0x00,
}