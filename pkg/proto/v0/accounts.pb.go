// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/proto/v0/accounts.proto

package proto

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

type Record struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Payload              *Payload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c84319968a576b, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Record) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Payload struct {
	Phoenix              *Phoenix `protobuf:"bytes,1,opt,name=phoenix,proto3" json:"phoenix,omitempty"`
	Account              *Account `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c84319968a576b, []int{1}
}

func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetPhoenix() *Phoenix {
	if m != nil {
		return m.Phoenix
	}
	return nil
}

func (m *Payload) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type Account struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DateOfBirth          string   `protobuf:"bytes,2,opt,name=dateOfBirth,proto3" json:"dateOfBirth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c84319968a576b, []int{2}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Account) GetDateOfBirth() string {
	if m != nil {
		return m.DateOfBirth
	}
	return ""
}

type Phoenix struct {
	Theme                string   `protobuf:"bytes,1,opt,name=theme,proto3" json:"theme,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Phoenix) Reset()         { *m = Phoenix{} }
func (m *Phoenix) String() string { return proto.CompactTextString(m) }
func (*Phoenix) ProtoMessage()    {}
func (*Phoenix) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c84319968a576b, []int{3}
}

func (m *Phoenix) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Phoenix.Unmarshal(m, b)
}
func (m *Phoenix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Phoenix.Marshal(b, m, deterministic)
}
func (m *Phoenix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Phoenix.Merge(m, src)
}
func (m *Phoenix) XXX_Size() int {
	return xxx_messageInfo_Phoenix.Size(m)
}
func (m *Phoenix) XXX_DiscardUnknown() {
	xxx_messageInfo_Phoenix.DiscardUnknown(m)
}

var xxx_messageInfo_Phoenix proto.InternalMessageInfo

func (m *Phoenix) GetTheme() string {
	if m != nil {
		return m.Theme
	}
	return ""
}

type Query struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c84319968a576b, []int{4}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*Record)(nil), "settings.Record")
	proto.RegisterType((*Payload)(nil), "settings.Payload")
	proto.RegisterType((*Account)(nil), "settings.Account")
	proto.RegisterType((*Phoenix)(nil), "settings.Phoenix")
	proto.RegisterType((*Query)(nil), "settings.Query")
}

func init() { proto.RegisterFile("pkg/proto/v0/accounts.proto", fileDescriptor_e3c84319968a576b) }

var fileDescriptor_e3c84319968a576b = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xcd, 0x4e, 0x83, 0x40,
	0x10, 0x0e, 0x22, 0xc5, 0x4e, 0x0f, 0xad, 0x13, 0x0f, 0x55, 0x0f, 0x12, 0x4e, 0x18, 0x13, 0x30,
	0xf5, 0x01, 0x8c, 0xbd, 0xf4, 0xa8, 0x2e, 0x37, 0x6f, 0xeb, 0x32, 0x2d, 0xa4, 0xca, 0x92, 0x65,
	0xdb, 0xc8, 0xdb, 0x9b, 0xee, 0x2e, 0xd6, 0x86, 0x13, 0xc3, 0xf7, 0x37, 0x5f, 0x66, 0xe1, 0xb6,
	0xd9, 0x6e, 0xb2, 0x46, 0x49, 0x2d, 0xb3, 0xfd, 0x63, 0xc6, 0x85, 0x90, 0xbb, 0x5a, 0xb7, 0xa9,
	0x41, 0xf0, 0xa2, 0x25, 0xad, 0xab, 0x7a, 0xd3, 0xc6, 0x2b, 0x18, 0x31, 0x12, 0x52, 0x15, 0x38,
	0x03, 0x7f, 0x4b, 0xdd, 0xdc, 0x8b, 0xbc, 0x64, 0xcc, 0x0e, 0x23, 0x3e, 0x40, 0xd8, 0xf0, 0xee,
	0x4b, 0xf2, 0x62, 0x7e, 0x16, 0x79, 0xc9, 0x64, 0x71, 0x99, 0xf6, 0xbe, 0xf4, 0xcd, 0x12, 0xac,
	0x57, 0xc4, 0x02, 0x42, 0x87, 0x19, 0x5f, 0x29, 0xa9, 0xae, 0x7e, 0x4c, 0xda, 0xa9, 0xcf, 0x12,
	0xac, 0x57, 0x1c, 0xc4, 0xae, 0xdc, 0x70, 0xc9, 0x8b, 0x25, 0x58, 0xaf, 0x88, 0x9f, 0x21, 0x74,
	0x18, 0x22, 0x9c, 0xd7, 0xfc, 0x9b, 0x5c, 0x5f, 0x33, 0x63, 0x04, 0x93, 0x82, 0x6b, 0x7a, 0x5d,
	0x2f, 0x2b, 0xa5, 0x4b, 0x93, 0x37, 0x66, 0xff, 0xa1, 0xf8, 0x0e, 0x42, 0xd7, 0x00, 0xaf, 0x20,
	0xd0, 0x25, 0xfd, 0x25, 0xd8, 0x9f, 0xf8, 0x1a, 0x82, 0xf7, 0x1d, 0xa9, 0x6e, 0x78, 0x8e, 0xc5,
	0x1a, 0xa6, 0xb9, 0x6b, 0x96, 0x93, 0xda, 0x57, 0x82, 0xf0, 0x1e, 0xfc, 0x9c, 0x34, 0xce, 0x8e,
	0x95, 0xed, 0x31, 0x6f, 0x06, 0x08, 0x26, 0xe0, 0xaf, 0x48, 0xe3, 0xf4, 0x48, 0x98, 0x3d, 0x43,
	0xe5, 0x32, 0xfc, 0x08, 0xcc, 0x2b, 0x7d, 0x8e, 0xcc, 0xe7, 0xe9, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x82, 0xf2, 0xd3, 0x55, 0xcb, 0x01, 0x00, 0x00,
}
