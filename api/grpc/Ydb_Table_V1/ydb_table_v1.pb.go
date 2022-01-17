// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_table_v1.proto

package Ydb_Table_V1

import (
	_ "github.com/yandex-cloud/ydb-go-sdk/v2/api/protos/Ydb_Table"
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

func init() { proto.RegisterFile("ydb_table_v1.proto", fileDescriptor_d67561c16f71e10d) }

var fileDescriptor_d67561c16f71e10d = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x95, 0x6f, 0x6f, 0xd3, 0x30,
	0x10, 0xc6, 0xc5, 0x1b, 0x24, 0x4c, 0xf9, 0x67, 0x10, 0x88, 0xc1, 0x18, 0x62, 0x7f, 0x90, 0x78,
	0x91, 0x52, 0xf8, 0x04, 0x6b, 0x0b, 0x12, 0x02, 0xb1, 0xd1, 0x0e, 0x24, 0x84, 0xc4, 0xe4, 0x24,
	0xa7, 0x61, 0xe6, 0x24, 0xc6, 0x76, 0xaa, 0xe5, 0xfb, 0xf0, 0x41, 0xa7, 0x26, 0x76, 0x72, 0x49,
	0x9c, 0xf6, 0x5d, 0x75, 0xcf, 0x73, 0xbf, 0x7b, 0xd2, 0x9c, 0x2e, 0x84, 0x16, 0x71, 0x78, 0x6e,
	0x58, 0x28, 0xe0, 0x7c, 0x35, 0x09, 0xa4, 0xca, 0x4c, 0x46, 0x47, 0x3f, 0xe3, 0x30, 0x38, 0x5b,
	0xd7, 0x82, 0x1f, 0x93, 0x9d, 0xc3, 0x22, 0x0e, 0xc7, 0x32, 0x0f, 0x05, 0x8f, 0xc6, 0x4c, 0xf2,
	0x71, 0xe9, 0xd1, 0xe3, 0xba, 0xaf, 0x6a, 0x7a, 0xf7, 0xff, 0x2e, 0x19, 0x95, 0x3d, 0x4b, 0x50,
	0x2b, 0x1e, 0x01, 0x5d, 0x90, 0x3b, 0x33, 0x05, 0xcc, 0xc0, 0x12, 0xb4, 0xe6, 0x59, 0x4a, 0xf7,
	0x82, 0x86, 0xdb, 0x52, 0x16, 0xf0, 0x2f, 0x07, 0x6d, 0x76, 0x5e, 0x0e, 0x1b, 0xb4, 0xcc, 0x52,
	0x5d, 0x32, 0xe7, 0x20, 0xc0, 0xcf, 0x6c, 0x29, 0x3e, 0x66, 0xc7, 0x60, 0x99, 0x1f, 0xc9, 0xad,
	0xcf, 0x00, 0xf2, 0x58, 0xf0, 0x15, 0xd0, 0x67, 0xc8, 0x5e, 0x57, 0x1d, 0xeb, 0xb9, 0x5f, 0xb4,
	0x9c, 0x2f, 0xe4, 0x76, 0x15, 0xba, 0x34, 0xd0, 0xdd, 0xde, 0xc3, 0x94, 0xbf, 0x1d, 0xeb, 0xc5,
	0x90, 0xdc, 0xa4, 0x9a, 0xab, 0x4c, 0x56, 0x2c, 0x9c, 0xaa, 0xae, 0xfa, 0x52, 0x21, 0xd1, 0x72,
	0x3e, 0x11, 0x72, 0x2c, 0x0c, 0xa8, 0x0a, 0x84, 0xbd, 0x4d, 0xd9, 0x91, 0x76, 0x07, 0xd4, 0x26,
	0xd2, 0x2c, 0x93, 0x45, 0x3f, 0x52, 0x5d, 0xf5, 0x45, 0x42, 0x62, 0x13, 0xa9, 0x2e, 0x6a, 0xea,
	0xf5, 0x6a, 0x5f, 0x24, 0xac, 0x5a, 0xd4, 0x09, 0x19, 0x2d, 0x20, 0x65, 0x09, 0x58, 0x18, 0xfe,
	0x57, 0xb1, 0xe0, 0x70, 0x7b, 0x83, 0x3a, 0x5e, 0x30, 0x1d, 0x29, 0x1e, 0xda, 0xd7, 0xd8, 0x5e,
	0x30, 0xa4, 0xf8, 0x17, 0xac, 0x65, 0xb0, 0xcc, 0x5f, 0xe4, 0xfe, 0x87, 0x2b, 0x29, 0x18, 0x4f,
	0xe7, 0xcc, 0xb0, 0x6f, 0x39, 0xa8, 0x82, 0xbe, 0x42, 0x5d, 0x5d, 0xd1, 0x91, 0xf7, 0x37, 0x7a,
	0x1a, 0xf8, 0xa9, 0x02, 0xc9, 0x14, 0xf8, 0xe1, 0x5d, 0xd1, 0x07, 0xef, 0x7b, 0x70, 0x72, 0x88,
	0x72, 0x03, 0x43, 0xc9, 0xdb, 0xa2, 0x3f, 0x79, 0xd7, 0x63, 0xe1, 0x8c, 0x50, 0xab, 0x2d, 0xa3,
	0x3f, 0x90, 0x40, 0x85, 0x3f, 0xe8, 0xb7, 0x22, 0xd9, 0x0d, 0x38, 0xdc, 0xe2, 0x6a, 0xf2, 0x4f,
	0xe1, 0x82, 0xa7, 0x67, 0x8a, 0xa5, 0x9a, 0x45, 0x66, 0x7d, 0x31, 0x70, 0xfe, 0xae, 0xe8, 0xcb,
	0xdf, 0xf7, 0x58, 0xf8, 0x6f, 0xf2, 0x60, 0x96, 0x25, 0x09, 0x37, 0x98, 0xbe, 0xdf, 0xda, 0xd7,
	0x8e, 0xea, 0xf0, 0x07, 0x9b, 0x4d, 0x96, 0x1f, 0x93, 0x87, 0x8b, 0x4c, 0x88, 0x90, 0x45, 0x97,
	0x78, 0x02, 0x7e, 0x74, 0x8f, 0xee, 0x66, 0x1c, 0x6d, 0xb3, 0xd9, 0x29, 0x17, 0xe4, 0x51, 0x6b,
	0x6b, 0x4f, 0xe4, 0x5a, 0xd6, 0xf4, 0x68, 0x68, 0xad, 0xad, 0xc1, 0xcd, 0x79, 0xbd, 0xd5, 0x67,
	0x07, 0x7d, 0x25, 0xf7, 0x96, 0x46, 0x01, 0x4b, 0x16, 0xc0, 0xe2, 0xfe, 0x0d, 0xa9, 0xab, 0xbe,
	0x1b, 0x82, 0xc4, 0x8a, 0xf6, 0xf6, 0xc6, 0xfa, 0x8a, 0x4c, 0x73, 0x71, 0xf9, 0x5d, 0x6a, 0x50,
	0xa6, 0x75, 0x45, 0x9a, 0xb2, 0xef, 0x8a, 0x60, 0xd5, 0x46, 0xfb, 0x4b, 0x1e, 0x57, 0xd1, 0xea,
	0x55, 0x62, 0xe9, 0xe0, 0xb2, 0xd7, 0xa2, 0x83, 0xbf, 0xd9, 0xe0, 0x39, 0x65, 0xca, 0x70, 0x26,
	0x9a, 0xd8, 0xd3, 0xa7, 0xe4, 0x49, 0x94, 0x25, 0x41, 0xc1, 0xd2, 0x18, 0xae, 0x82, 0x22, 0x0e,
	0x83, 0xea, 0x23, 0xba, 0x9a, 0x84, 0x37, 0xcb, 0x0f, 0xe9, 0xfb, 0xeb, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xb9, 0x2c, 0x83, 0x63, 0x93, 0x07, 0x00, 0x00,
}
