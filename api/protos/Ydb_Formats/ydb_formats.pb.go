// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_formats.proto

package Ydb_Formats

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

type ArrowBatchSettings struct {
	Schema               []byte   `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArrowBatchSettings) Reset()         { *m = ArrowBatchSettings{} }
func (m *ArrowBatchSettings) String() string { return proto.CompactTextString(m) }
func (*ArrowBatchSettings) ProtoMessage()    {}
func (*ArrowBatchSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec3de800d6b71469, []int{0}
}

func (m *ArrowBatchSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrowBatchSettings.Unmarshal(m, b)
}
func (m *ArrowBatchSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrowBatchSettings.Marshal(b, m, deterministic)
}
func (m *ArrowBatchSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrowBatchSettings.Merge(m, src)
}
func (m *ArrowBatchSettings) XXX_Size() int {
	return xxx_messageInfo_ArrowBatchSettings.Size(m)
}
func (m *ArrowBatchSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrowBatchSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ArrowBatchSettings proto.InternalMessageInfo

func (m *ArrowBatchSettings) GetSchema() []byte {
	if m != nil {
		return m.Schema
	}
	return nil
}

type CsvSettings struct {
	// Number of rows to skip before CSV data. It should be present only in the first upsert of CSV file.
	SkipRows uint32 `protobuf:"varint,1,opt,name=skip_rows,json=skipRows,proto3" json:"skip_rows,omitempty"`
	// Fields delimiter in CSV file. It's "," if not set.
	Delimiter []byte `protobuf:"bytes,2,opt,name=delimiter,proto3" json:"delimiter,omitempty"`
	// String value that would be interpreted as NULL.
	NullValue            []byte   `protobuf:"bytes,3,opt,name=null_value,json=nullValue,proto3" json:"null_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CsvSettings) Reset()         { *m = CsvSettings{} }
func (m *CsvSettings) String() string { return proto.CompactTextString(m) }
func (*CsvSettings) ProtoMessage()    {}
func (*CsvSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec3de800d6b71469, []int{1}
}

func (m *CsvSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CsvSettings.Unmarshal(m, b)
}
func (m *CsvSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CsvSettings.Marshal(b, m, deterministic)
}
func (m *CsvSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsvSettings.Merge(m, src)
}
func (m *CsvSettings) XXX_Size() int {
	return xxx_messageInfo_CsvSettings.Size(m)
}
func (m *CsvSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_CsvSettings.DiscardUnknown(m)
}

var xxx_messageInfo_CsvSettings proto.InternalMessageInfo

func (m *CsvSettings) GetSkipRows() uint32 {
	if m != nil {
		return m.SkipRows
	}
	return 0
}

func (m *CsvSettings) GetDelimiter() []byte {
	if m != nil {
		return m.Delimiter
	}
	return nil
}

func (m *CsvSettings) GetNullValue() []byte {
	if m != nil {
		return m.NullValue
	}
	return nil
}

func init() {
	proto.RegisterType((*ArrowBatchSettings)(nil), "Ydb.Formats.ArrowBatchSettings")
	proto.RegisterType((*CsvSettings)(nil), "Ydb.Formats.CsvSettings")
}

func init() { proto.RegisterFile("ydb_formats.proto", fileDescriptor_ec3de800d6b71469) }

var fileDescriptor_ec3de800d6b71469 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xac, 0x4c, 0x49, 0x8a,
	0x4f, 0xcb, 0x2f, 0xca, 0x4d, 0x2c, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x8e,
	0x4c, 0x49, 0xd2, 0x73, 0x83, 0x08, 0x29, 0xe9, 0x70, 0x09, 0x39, 0x16, 0x15, 0xe5, 0x97, 0x3b,
	0x25, 0x96, 0x24, 0x67, 0x04, 0xa7, 0x96, 0x94, 0x64, 0xe6, 0xa5, 0x17, 0x0b, 0x89, 0x71, 0xb1,
	0x15, 0x27, 0x67, 0xa4, 0xe6, 0x26, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x41, 0x79, 0x4a,
	0xe9, 0x5c, 0xdc, 0xce, 0xc5, 0x65, 0x70, 0x65, 0xd2, 0x5c, 0x9c, 0xc5, 0xd9, 0x99, 0x05, 0xf1,
	0x45, 0xf9, 0xe5, 0xc5, 0x60, 0x95, 0xbc, 0x41, 0x1c, 0x20, 0x81, 0xa0, 0xfc, 0xf2, 0x62, 0x21,
	0x19, 0x2e, 0xce, 0x94, 0xd4, 0x9c, 0xcc, 0xdc, 0xcc, 0x92, 0xd4, 0x22, 0x09, 0x26, 0xb0, 0x31,
	0x08, 0x01, 0x21, 0x59, 0x2e, 0xae, 0xbc, 0xd2, 0x9c, 0x9c, 0xf8, 0xb2, 0xc4, 0x9c, 0xd2, 0x54,
	0x09, 0x66, 0x88, 0x34, 0x48, 0x24, 0x0c, 0x24, 0xe0, 0x24, 0xcd, 0x25, 0x96, 0x9c, 0x9f, 0xab,
	0x57, 0x99, 0x98, 0x97, 0x92, 0x5a, 0xa1, 0x57, 0x99, 0x92, 0xa4, 0x07, 0xf5, 0xc3, 0x0f, 0x46,
	0xc6, 0x24, 0x36, 0xb0, 0x3f, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x06, 0x8e, 0x0c,
	0xdc, 0x00, 0x00, 0x00,
}
