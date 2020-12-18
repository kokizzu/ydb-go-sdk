// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_scripting.proto

package Ydb_Scripting

import (
	Ydb "github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb"
	Ydb_Operations "github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb_Operations"
	Ydb_Table "github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb_Table"
	Ydb_TableStats "github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb_TableStats"
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

type ExplainYqlRequest_Mode int32

const (
	ExplainYqlRequest_PARSE    ExplainYqlRequest_Mode = 0
	ExplainYqlRequest_VALIDATE ExplainYqlRequest_Mode = 1
	ExplainYqlRequest_PLAN     ExplainYqlRequest_Mode = 2
)

var ExplainYqlRequest_Mode_name = map[int32]string{
	0: "PARSE",
	1: "VALIDATE",
	2: "PLAN",
}

var ExplainYqlRequest_Mode_value = map[string]int32{
	"PARSE":    0,
	"VALIDATE": 1,
	"PLAN":     2,
}

func (x ExplainYqlRequest_Mode) String() string {
	return proto.EnumName(ExplainYqlRequest_Mode_name, int32(x))
}

func (ExplainYqlRequest_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49d60dbac494bf1d, []int{3, 0}
}

type ExecuteYqlRequest struct {
	OperationParams      *Ydb_Operations.OperationParams     `protobuf:"bytes,1,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	Script               string                              `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
	Parameters           map[string]*Ydb.TypedValue          `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CollectStats         Ydb_Table.QueryStatsCollection_Mode `protobuf:"varint,4,opt,name=collect_stats,json=collectStats,proto3,enum=Ydb.Table.QueryStatsCollection_Mode" json:"collect_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *ExecuteYqlRequest) Reset()         { *m = ExecuteYqlRequest{} }
func (m *ExecuteYqlRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteYqlRequest) ProtoMessage()    {}
func (*ExecuteYqlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49d60dbac494bf1d, []int{0}
}

func (m *ExecuteYqlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteYqlRequest.Unmarshal(m, b)
}
func (m *ExecuteYqlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteYqlRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteYqlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteYqlRequest.Merge(m, src)
}
func (m *ExecuteYqlRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteYqlRequest.Size(m)
}
func (m *ExecuteYqlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteYqlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteYqlRequest proto.InternalMessageInfo

func (m *ExecuteYqlRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if m != nil {
		return m.OperationParams
	}
	return nil
}

func (m *ExecuteYqlRequest) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

func (m *ExecuteYqlRequest) GetParameters() map[string]*Ydb.TypedValue {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *ExecuteYqlRequest) GetCollectStats() Ydb_Table.QueryStatsCollection_Mode {
	if m != nil {
		return m.CollectStats
	}
	return Ydb_Table.QueryStatsCollection_STATS_COLLECTION_UNSPECIFIED
}

type ExecuteYqlResponse struct {
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ExecuteYqlResponse) Reset()         { *m = ExecuteYqlResponse{} }
func (m *ExecuteYqlResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteYqlResponse) ProtoMessage()    {}
func (*ExecuteYqlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_49d60dbac494bf1d, []int{1}
}

func (m *ExecuteYqlResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteYqlResponse.Unmarshal(m, b)
}
func (m *ExecuteYqlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteYqlResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteYqlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteYqlResponse.Merge(m, src)
}
func (m *ExecuteYqlResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteYqlResponse.Size(m)
}
func (m *ExecuteYqlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteYqlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteYqlResponse proto.InternalMessageInfo

func (m *ExecuteYqlResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type ExecuteYqlResult struct {
	ResultSets           []*Ydb.ResultSet           `protobuf:"bytes,1,rep,name=result_sets,json=resultSets,proto3" json:"result_sets,omitempty"`
	QueryStats           *Ydb_TableStats.QueryStats `protobuf:"bytes,2,opt,name=query_stats,json=queryStats,proto3" json:"query_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ExecuteYqlResult) Reset()         { *m = ExecuteYqlResult{} }
func (m *ExecuteYqlResult) String() string { return proto.CompactTextString(m) }
func (*ExecuteYqlResult) ProtoMessage()    {}
func (*ExecuteYqlResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_49d60dbac494bf1d, []int{2}
}

func (m *ExecuteYqlResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteYqlResult.Unmarshal(m, b)
}
func (m *ExecuteYqlResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteYqlResult.Marshal(b, m, deterministic)
}
func (m *ExecuteYqlResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteYqlResult.Merge(m, src)
}
func (m *ExecuteYqlResult) XXX_Size() int {
	return xxx_messageInfo_ExecuteYqlResult.Size(m)
}
func (m *ExecuteYqlResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteYqlResult.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteYqlResult proto.InternalMessageInfo

func (m *ExecuteYqlResult) GetResultSets() []*Ydb.ResultSet {
	if m != nil {
		return m.ResultSets
	}
	return nil
}

func (m *ExecuteYqlResult) GetQueryStats() *Ydb_TableStats.QueryStats {
	if m != nil {
		return m.QueryStats
	}
	return nil
}

type ExplainYqlRequest struct {
	OperationParams      *Ydb_Operations.OperationParams `protobuf:"bytes,1,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	Script               string                          `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
	Mode                 ExplainYqlRequest_Mode          `protobuf:"varint,3,opt,name=mode,proto3,enum=Ydb.Scripting.ExplainYqlRequest_Mode" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ExplainYqlRequest) Reset()         { *m = ExplainYqlRequest{} }
func (m *ExplainYqlRequest) String() string { return proto.CompactTextString(m) }
func (*ExplainYqlRequest) ProtoMessage()    {}
func (*ExplainYqlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49d60dbac494bf1d, []int{3}
}

func (m *ExplainYqlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExplainYqlRequest.Unmarshal(m, b)
}
func (m *ExplainYqlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExplainYqlRequest.Marshal(b, m, deterministic)
}
func (m *ExplainYqlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExplainYqlRequest.Merge(m, src)
}
func (m *ExplainYqlRequest) XXX_Size() int {
	return xxx_messageInfo_ExplainYqlRequest.Size(m)
}
func (m *ExplainYqlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExplainYqlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExplainYqlRequest proto.InternalMessageInfo

func (m *ExplainYqlRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if m != nil {
		return m.OperationParams
	}
	return nil
}

func (m *ExplainYqlRequest) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

func (m *ExplainYqlRequest) GetMode() ExplainYqlRequest_Mode {
	if m != nil {
		return m.Mode
	}
	return ExplainYqlRequest_PARSE
}

type ExplainYqlResponse struct {
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ExplainYqlResponse) Reset()         { *m = ExplainYqlResponse{} }
func (m *ExplainYqlResponse) String() string { return proto.CompactTextString(m) }
func (*ExplainYqlResponse) ProtoMessage()    {}
func (*ExplainYqlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_49d60dbac494bf1d, []int{4}
}

func (m *ExplainYqlResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExplainYqlResponse.Unmarshal(m, b)
}
func (m *ExplainYqlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExplainYqlResponse.Marshal(b, m, deterministic)
}
func (m *ExplainYqlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExplainYqlResponse.Merge(m, src)
}
func (m *ExplainYqlResponse) XXX_Size() int {
	return xxx_messageInfo_ExplainYqlResponse.Size(m)
}
func (m *ExplainYqlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExplainYqlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExplainYqlResponse proto.InternalMessageInfo

func (m *ExplainYqlResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type ExplainYqlResult struct {
	ParametersTypes      map[string]*Ydb.Type `protobuf:"bytes,1,rep,name=parameters_types,json=parametersTypes,proto3" json:"parameters_types,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Plan                 string               `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ExplainYqlResult) Reset()         { *m = ExplainYqlResult{} }
func (m *ExplainYqlResult) String() string { return proto.CompactTextString(m) }
func (*ExplainYqlResult) ProtoMessage()    {}
func (*ExplainYqlResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_49d60dbac494bf1d, []int{5}
}

func (m *ExplainYqlResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExplainYqlResult.Unmarshal(m, b)
}
func (m *ExplainYqlResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExplainYqlResult.Marshal(b, m, deterministic)
}
func (m *ExplainYqlResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExplainYqlResult.Merge(m, src)
}
func (m *ExplainYqlResult) XXX_Size() int {
	return xxx_messageInfo_ExplainYqlResult.Size(m)
}
func (m *ExplainYqlResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ExplainYqlResult.DiscardUnknown(m)
}

var xxx_messageInfo_ExplainYqlResult proto.InternalMessageInfo

func (m *ExplainYqlResult) GetParametersTypes() map[string]*Ydb.Type {
	if m != nil {
		return m.ParametersTypes
	}
	return nil
}

func (m *ExplainYqlResult) GetPlan() string {
	if m != nil {
		return m.Plan
	}
	return ""
}

func init() {
	proto.RegisterEnum("Ydb.Scripting.ExplainYqlRequest_Mode", ExplainYqlRequest_Mode_name, ExplainYqlRequest_Mode_value)
	proto.RegisterType((*ExecuteYqlRequest)(nil), "Ydb.Scripting.ExecuteYqlRequest")
	proto.RegisterMapType((map[string]*Ydb.TypedValue)(nil), "Ydb.Scripting.ExecuteYqlRequest.ParametersEntry")
	proto.RegisterType((*ExecuteYqlResponse)(nil), "Ydb.Scripting.ExecuteYqlResponse")
	proto.RegisterType((*ExecuteYqlResult)(nil), "Ydb.Scripting.ExecuteYqlResult")
	proto.RegisterType((*ExplainYqlRequest)(nil), "Ydb.Scripting.ExplainYqlRequest")
	proto.RegisterType((*ExplainYqlResponse)(nil), "Ydb.Scripting.ExplainYqlResponse")
	proto.RegisterType((*ExplainYqlResult)(nil), "Ydb.Scripting.ExplainYqlResult")
	proto.RegisterMapType((map[string]*Ydb.Type)(nil), "Ydb.Scripting.ExplainYqlResult.ParametersTypesEntry")
}

func init() { proto.RegisterFile("ydb_scripting.proto", fileDescriptor_49d60dbac494bf1d) }

var fileDescriptor_49d60dbac494bf1d = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x49, 0x5a, 0x35, 0x93, 0xb6, 0x31, 0x0b, 0x42, 0x26, 0x97, 0x46, 0x11, 0x95, 0x82,
	0x84, 0x36, 0x28, 0x20, 0xf1, 0x75, 0x4a, 0x21, 0x87, 0xa2, 0xa6, 0x98, 0x4d, 0x55, 0xa9, 0x27,
	0xcb, 0x1f, 0x2b, 0x64, 0xc5, 0xb1, 0x37, 0xde, 0x35, 0xaa, 0x6f, 0xfc, 0x4a, 0x6e, 0xdc, 0xf8,
	0x11, 0x1c, 0xd1, 0xee, 0x3a, 0x8e, 0xdd, 0x96, 0x08, 0x09, 0x89, 0xdb, 0x78, 0xe6, 0xcd, 0x78,
	0xe7, 0xbd, 0xb7, 0x0b, 0x0f, 0xf2, 0xc0, 0x73, 0xb8, 0x9f, 0x86, 0x4c, 0x84, 0xf1, 0x17, 0xcc,
	0xd2, 0x44, 0x24, 0xe8, 0xe0, 0x2a, 0xf0, 0xf0, 0x7c, 0x9d, 0xec, 0x3d, 0x5b, 0x84, 0x8b, 0x70,
	0x99, 0x8e, 0x58, 0xe6, 0x45, 0xa1, 0x3f, 0x72, 0x59, 0x38, 0x52, 0x38, 0x3e, 0x92, 0xcd, 0x09,
	0xa3, 0xa9, 0x2b, 0xc2, 0x24, 0xd6, 0xcd, 0xbd, 0xe1, 0x56, 0xf4, 0x57, 0x37, 0xca, 0xe8, 0x5f,
	0x21, 0x85, 0xeb, 0x45, 0x6b, 0x24, 0xde, 0x8a, 0x5c, 0x65, 0x34, 0xcd, 0x1d, 0x2e, 0x5c, 0xc1,
	0x35, 0x7e, 0xf0, 0xb3, 0x01, 0xf7, 0xa7, 0xd7, 0xd4, 0xcf, 0x04, 0xbd, 0x5a, 0x45, 0x84, 0xae,
	0x32, 0xca, 0x05, 0xfa, 0x08, 0x66, 0x79, 0x58, 0x87, 0xb9, 0xa9, 0xbb, 0xe4, 0x96, 0xd1, 0x37,
	0x86, 0x9d, 0xf1, 0x11, 0x96, 0x1b, 0x7f, 0x5a, 0x17, 0xf9, 0x26, 0xb4, 0x15, 0x8c, 0x74, 0x93,
	0x7a, 0x02, 0x3d, 0x82, 0x5d, 0xcd, 0x9a, 0xd5, 0xe8, 0x1b, 0xc3, 0x36, 0x29, 0xbe, 0x90, 0x0d,
	0xa0, 0x26, 0x53, 0x41, 0x53, 0x6e, 0x35, 0xfb, 0xcd, 0x61, 0x67, 0xfc, 0x1c, 0xd7, 0xf8, 0xc4,
	0xb7, 0x4e, 0x86, 0xed, 0xb2, 0x65, 0x1a, 0x8b, 0x34, 0x27, 0x95, 0x19, 0xe8, 0x14, 0x0e, 0xfc,
	0x24, 0x8a, 0xa8, 0x2f, 0xf4, 0x8a, 0x56, 0xab, 0x6f, 0x0c, 0x0f, 0xc7, 0x4f, 0xd4, 0xd0, 0x0b,
	0x45, 0xd2, 0x67, 0x49, 0xc0, 0x5c, 0x16, 0xdf, 0x6b, 0xa4, 0x54, 0x63, 0x96, 0x04, 0x94, 0xec,
	0x17, 0xad, 0xaa, 0xd8, 0x3b, 0x87, 0xee, 0x8d, 0x3f, 0x21, 0x13, 0x9a, 0x0b, 0x9a, 0x2b, 0x1a,
	0xda, 0x44, 0x86, 0xe8, 0x18, 0x76, 0x94, 0x48, 0x6a, 0xb1, 0xce, 0xb8, 0xab, 0xff, 0x93, 0x33,
	0x1a, 0x5c, 0xca, 0x34, 0xd1, 0xd5, 0xb7, 0x8d, 0xd7, 0xc6, 0x60, 0x06, 0xa8, 0xba, 0x0b, 0x67,
	0x49, 0xcc, 0x29, 0x7a, 0x05, 0xed, 0x92, 0xad, 0x82, 0xdf, 0xc7, 0x7f, 0xe4, 0x97, 0x6c, 0xb0,
	0x83, 0x6f, 0x06, 0x98, 0xb5, 0x79, 0x59, 0x24, 0xd0, 0x08, 0x3a, 0xa9, 0x8a, 0x1c, 0x4e, 0x85,
	0xd4, 0x4b, 0x32, 0x7a, 0xa8, 0xe6, 0x69, 0xc4, 0x9c, 0x0a, 0x02, 0xe9, 0x3a, 0xe4, 0xe8, 0x1d,
	0x74, 0x2a, 0x86, 0x28, 0xb6, 0xe8, 0x6d, 0xd8, 0x52, 0x54, 0x54, 0x28, 0x23, 0xb0, 0x2a, 0xe3,
	0xc1, 0x0f, 0x43, 0x1a, 0x87, 0x45, 0x6e, 0x18, 0xff, 0x67, 0xe3, 0xbc, 0x81, 0xd6, 0x32, 0x09,
	0xa8, 0xd5, 0x54, 0xea, 0x1e, 0xdf, 0xb2, 0xcc, 0x8d, 0x33, 0x69, 0x79, 0x55, 0xcb, 0xe0, 0x29,
	0xb4, 0xe4, 0x17, 0x6a, 0xc3, 0x8e, 0x3d, 0x21, 0xf3, 0xa9, 0x79, 0x0f, 0xed, 0xc3, 0xde, 0xe5,
	0xe4, 0xec, 0xf4, 0xc3, 0xe4, 0x62, 0x6a, 0x1a, 0x68, 0x0f, 0x5a, 0xf6, 0xd9, 0xe4, 0xdc, 0x6c,
	0x68, 0xc5, 0x36, 0xa3, 0xfe, 0x55, 0xb1, 0xef, 0x4a, 0xb1, 0xca, 0x3c, 0xa9, 0x98, 0x03, 0xe6,
	0xc6, 0xbe, 0x8e, 0xc8, 0x19, 0x5d, 0xcb, 0xf6, 0x72, 0xcb, 0x56, 0xb2, 0xb5, 0x72, 0x0f, 0xa4,
	0xe1, 0x8a, 0xcb, 0xd0, 0x65, 0xf5, 0x2c, 0x42, 0xd0, 0x62, 0x91, 0x1b, 0x17, 0x04, 0xaa, 0xb8,
	0x37, 0x83, 0x87, 0x77, 0x35, 0xdf, 0xe1, 0xef, 0xa3, 0xba, 0xbf, 0xdb, 0xa5, 0xbf, 0x2b, 0xce,
	0x3e, 0xc1, 0x60, 0xf9, 0xc9, 0x12, 0xe7, 0x6e, 0x1c, 0xd0, 0x6b, 0x9c, 0x07, 0x1e, 0x2e, 0xdf,
	0xc8, 0x93, 0x6e, 0xb9, 0x80, 0xad, 0xde, 0xa0, 0x5f, 0x86, 0xe1, 0xed, 0xaa, 0x77, 0xe7, 0xc5,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x1f, 0x4e, 0x2c, 0x4f, 0x05, 0x00, 0x00,
}

const ()

// SetOperationParams implements ydb generic interface for setting
// operation parameters inside driver implementation.
func (m *ExecuteYqlRequest) SetOperationParams(v *Ydb_Operations.OperationParams) {
	m.OperationParams = v
}

// SetOperationParams implements ydb generic interface for setting
// operation parameters inside driver implementation.
func (m *ExplainYqlRequest) SetOperationParams(v *Ydb_Operations.OperationParams) {
	m.OperationParams = v
}
