// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_query_stats.proto

package Ydb_TableStats

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

// Describes select, update (insert, upsert, replace) and delete operations
type OperationStats struct {
	Rows                 uint64   `protobuf:"varint,1,opt,name=rows,proto3" json:"rows,omitempty"`
	Bytes                uint64   `protobuf:"varint,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationStats) Reset()         { *m = OperationStats{} }
func (m *OperationStats) String() string { return proto.CompactTextString(m) }
func (*OperationStats) ProtoMessage()    {}
func (*OperationStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd6647573551cb14, []int{0}
}

func (m *OperationStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationStats.Unmarshal(m, b)
}
func (m *OperationStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationStats.Marshal(b, m, deterministic)
}
func (m *OperationStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationStats.Merge(m, src)
}
func (m *OperationStats) XXX_Size() int {
	return xxx_messageInfo_OperationStats.Size(m)
}
func (m *OperationStats) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationStats.DiscardUnknown(m)
}

var xxx_messageInfo_OperationStats proto.InternalMessageInfo

func (m *OperationStats) GetRows() uint64 {
	if m != nil {
		return m.Rows
	}
	return 0
}

func (m *OperationStats) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

// Describes all operations on a table
type TableAccessStats struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Reads                *OperationStats `protobuf:"bytes,3,opt,name=reads,proto3" json:"reads,omitempty"`
	Updates              *OperationStats `protobuf:"bytes,4,opt,name=updates,proto3" json:"updates,omitempty"`
	Deletes              *OperationStats `protobuf:"bytes,5,opt,name=deletes,proto3" json:"deletes,omitempty"`
	PartitionsCount      uint64          `protobuf:"varint,6,opt,name=partitions_count,json=partitionsCount,proto3" json:"partitions_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TableAccessStats) Reset()         { *m = TableAccessStats{} }
func (m *TableAccessStats) String() string { return proto.CompactTextString(m) }
func (*TableAccessStats) ProtoMessage()    {}
func (*TableAccessStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd6647573551cb14, []int{1}
}

func (m *TableAccessStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableAccessStats.Unmarshal(m, b)
}
func (m *TableAccessStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableAccessStats.Marshal(b, m, deterministic)
}
func (m *TableAccessStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableAccessStats.Merge(m, src)
}
func (m *TableAccessStats) XXX_Size() int {
	return xxx_messageInfo_TableAccessStats.Size(m)
}
func (m *TableAccessStats) XXX_DiscardUnknown() {
	xxx_messageInfo_TableAccessStats.DiscardUnknown(m)
}

var xxx_messageInfo_TableAccessStats proto.InternalMessageInfo

func (m *TableAccessStats) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TableAccessStats) GetReads() *OperationStats {
	if m != nil {
		return m.Reads
	}
	return nil
}

func (m *TableAccessStats) GetUpdates() *OperationStats {
	if m != nil {
		return m.Updates
	}
	return nil
}

func (m *TableAccessStats) GetDeletes() *OperationStats {
	if m != nil {
		return m.Deletes
	}
	return nil
}

func (m *TableAccessStats) GetPartitionsCount() uint64 {
	if m != nil {
		return m.PartitionsCount
	}
	return 0
}

type QueryPhaseStats struct {
	DurationUs           uint64              `protobuf:"varint,1,opt,name=duration_us,json=durationUs,proto3" json:"duration_us,omitempty"`
	TableAccess          []*TableAccessStats `protobuf:"bytes,2,rep,name=table_access,json=tableAccess,proto3" json:"table_access,omitempty"`
	CpuTimeUs            uint64              `protobuf:"varint,3,opt,name=cpu_time_us,json=cpuTimeUs,proto3" json:"cpu_time_us,omitempty"`
	AffectedShards       uint64              `protobuf:"varint,4,opt,name=affected_shards,json=affectedShards,proto3" json:"affected_shards,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *QueryPhaseStats) Reset()         { *m = QueryPhaseStats{} }
func (m *QueryPhaseStats) String() string { return proto.CompactTextString(m) }
func (*QueryPhaseStats) ProtoMessage()    {}
func (*QueryPhaseStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd6647573551cb14, []int{2}
}

func (m *QueryPhaseStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryPhaseStats.Unmarshal(m, b)
}
func (m *QueryPhaseStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryPhaseStats.Marshal(b, m, deterministic)
}
func (m *QueryPhaseStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPhaseStats.Merge(m, src)
}
func (m *QueryPhaseStats) XXX_Size() int {
	return xxx_messageInfo_QueryPhaseStats.Size(m)
}
func (m *QueryPhaseStats) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPhaseStats.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPhaseStats proto.InternalMessageInfo

func (m *QueryPhaseStats) GetDurationUs() uint64 {
	if m != nil {
		return m.DurationUs
	}
	return 0
}

func (m *QueryPhaseStats) GetTableAccess() []*TableAccessStats {
	if m != nil {
		return m.TableAccess
	}
	return nil
}

func (m *QueryPhaseStats) GetCpuTimeUs() uint64 {
	if m != nil {
		return m.CpuTimeUs
	}
	return 0
}

func (m *QueryPhaseStats) GetAffectedShards() uint64 {
	if m != nil {
		return m.AffectedShards
	}
	return 0
}

type CompilationStats struct {
	FromCache            bool     `protobuf:"varint,1,opt,name=from_cache,json=fromCache,proto3" json:"from_cache,omitempty"`
	DurationUs           uint64   `protobuf:"varint,2,opt,name=duration_us,json=durationUs,proto3" json:"duration_us,omitempty"`
	CpuTimeUs            uint64   `protobuf:"varint,3,opt,name=cpu_time_us,json=cpuTimeUs,proto3" json:"cpu_time_us,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompilationStats) Reset()         { *m = CompilationStats{} }
func (m *CompilationStats) String() string { return proto.CompactTextString(m) }
func (*CompilationStats) ProtoMessage()    {}
func (*CompilationStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd6647573551cb14, []int{3}
}

func (m *CompilationStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompilationStats.Unmarshal(m, b)
}
func (m *CompilationStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompilationStats.Marshal(b, m, deterministic)
}
func (m *CompilationStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompilationStats.Merge(m, src)
}
func (m *CompilationStats) XXX_Size() int {
	return xxx_messageInfo_CompilationStats.Size(m)
}
func (m *CompilationStats) XXX_DiscardUnknown() {
	xxx_messageInfo_CompilationStats.DiscardUnknown(m)
}

var xxx_messageInfo_CompilationStats proto.InternalMessageInfo

func (m *CompilationStats) GetFromCache() bool {
	if m != nil {
		return m.FromCache
	}
	return false
}

func (m *CompilationStats) GetDurationUs() uint64 {
	if m != nil {
		return m.DurationUs
	}
	return 0
}

func (m *CompilationStats) GetCpuTimeUs() uint64 {
	if m != nil {
		return m.CpuTimeUs
	}
	return 0
}

type QueryStats struct {
	// A query might have one or more execution phases
	QueryPhases          []*QueryPhaseStats `protobuf:"bytes,1,rep,name=query_phases,json=queryPhases,proto3" json:"query_phases,omitempty"`
	Compilation          *CompilationStats  `protobuf:"bytes,2,opt,name=compilation,proto3" json:"compilation,omitempty"`
	ProcessCpuTimeUs     uint64             `protobuf:"varint,3,opt,name=process_cpu_time_us,json=processCpuTimeUs,proto3" json:"process_cpu_time_us,omitempty"`
	QueryPlan            string             `protobuf:"bytes,4,opt,name=query_plan,json=queryPlan,proto3" json:"query_plan,omitempty"`
	QueryAst             string             `protobuf:"bytes,5,opt,name=query_ast,json=queryAst,proto3" json:"query_ast,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *QueryStats) Reset()         { *m = QueryStats{} }
func (m *QueryStats) String() string { return proto.CompactTextString(m) }
func (*QueryStats) ProtoMessage()    {}
func (*QueryStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd6647573551cb14, []int{4}
}

func (m *QueryStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStats.Unmarshal(m, b)
}
func (m *QueryStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStats.Marshal(b, m, deterministic)
}
func (m *QueryStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStats.Merge(m, src)
}
func (m *QueryStats) XXX_Size() int {
	return xxx_messageInfo_QueryStats.Size(m)
}
func (m *QueryStats) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStats.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStats proto.InternalMessageInfo

func (m *QueryStats) GetQueryPhases() []*QueryPhaseStats {
	if m != nil {
		return m.QueryPhases
	}
	return nil
}

func (m *QueryStats) GetCompilation() *CompilationStats {
	if m != nil {
		return m.Compilation
	}
	return nil
}

func (m *QueryStats) GetProcessCpuTimeUs() uint64 {
	if m != nil {
		return m.ProcessCpuTimeUs
	}
	return 0
}

func (m *QueryStats) GetQueryPlan() string {
	if m != nil {
		return m.QueryPlan
	}
	return ""
}

func (m *QueryStats) GetQueryAst() string {
	if m != nil {
		return m.QueryAst
	}
	return ""
}

func init() {
	proto.RegisterType((*OperationStats)(nil), "Ydb.TableStats.OperationStats")
	proto.RegisterType((*TableAccessStats)(nil), "Ydb.TableStats.TableAccessStats")
	proto.RegisterType((*QueryPhaseStats)(nil), "Ydb.TableStats.QueryPhaseStats")
	proto.RegisterType((*CompilationStats)(nil), "Ydb.TableStats.CompilationStats")
	proto.RegisterType((*QueryStats)(nil), "Ydb.TableStats.QueryStats")
}

func init() { proto.RegisterFile("ydb_query_stats.proto", fileDescriptor_bd6647573551cb14) }

var fileDescriptor_bd6647573551cb14 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x3f, 0x6f, 0xd4, 0x30,
	0x18, 0xc6, 0x95, 0x5c, 0xae, 0x5c, 0xde, 0x54, 0x77, 0x91, 0x0b, 0x52, 0x24, 0xd4, 0xf6, 0x94,
	0x85, 0x63, 0x20, 0x43, 0x61, 0x40, 0x6c, 0x6d, 0x36, 0x16, 0xc0, 0x6d, 0x07, 0x26, 0xcb, 0xb1,
	0x7d, 0xba, 0x48, 0xf9, 0x57, 0xdb, 0x11, 0x64, 0xe4, 0x9b, 0xf1, 0xb1, 0xd8, 0x40, 0xb6, 0xef,
	0x4f, 0x9b, 0x22, 0xb5, 0x9b, 0xfd, 0xd8, 0xcf, 0xfb, 0x3e, 0xef, 0x2f, 0x0e, 0xbc, 0x1a, 0x78,
	0x41, 0xee, 0x7a, 0x21, 0x07, 0xa2, 0x34, 0xd5, 0x2a, 0xeb, 0x64, 0xab, 0x5b, 0x34, 0xff, 0xce,
	0x8b, 0xec, 0x86, 0x16, 0x95, 0xb8, 0x36, 0x6a, 0xfa, 0x09, 0xe6, 0x5f, 0x3a, 0x21, 0xa9, 0x2e,
	0xdb, 0xc6, 0x2a, 0x08, 0x41, 0x20, 0xdb, 0x1f, 0x2a, 0xf1, 0x96, 0xde, 0x2a, 0xc0, 0x76, 0x8d,
	0x5e, 0xc2, 0xb4, 0x18, 0xb4, 0x50, 0x89, 0x6f, 0x45, 0xb7, 0x49, 0x7f, 0xf9, 0x10, 0xdb, 0x52,
	0x97, 0x8c, 0x09, 0xa5, 0xf6, 0xf6, 0x86, 0xd6, 0xc2, 0xda, 0x43, 0x6c, 0xd7, 0xe8, 0x03, 0x4c,
	0xa5, 0xa0, 0x5c, 0x25, 0x93, 0xa5, 0xb7, 0x8a, 0x2e, 0xce, 0xb2, 0x87, 0x21, 0xb2, 0x87, 0x09,
	0xb0, 0xbb, 0x8c, 0x3e, 0xc2, 0x8b, 0xbe, 0xe3, 0xd4, 0xb4, 0x0d, 0x9e, 0xe5, 0xdb, 0x5d, 0x37,
	0x4e, 0x2e, 0x2a, 0x61, 0x9c, 0xd3, 0xe7, 0x39, 0xb7, 0xd7, 0xd1, 0x5b, 0x88, 0x3b, 0x2a, 0x75,
	0x69, 0x8e, 0x14, 0x61, 0x6d, 0xdf, 0xe8, 0xe4, 0xc8, 0xce, 0xbc, 0x38, 0xe8, 0xb9, 0x91, 0x3f,
	0x07, 0x33, 0x3f, 0x9e, 0xa4, 0xbf, 0x3d, 0x58, 0x7c, 0x33, 0x94, 0xbf, 0x6e, 0xa8, 0x72, 0xc5,
	0xd1, 0x39, 0x44, 0xbc, 0x77, 0xe5, 0x49, 0xbf, 0x03, 0x09, 0x3b, 0xe9, 0x56, 0xa1, 0x1c, 0x8e,
	0xb5, 0xc9, 0x42, 0xa8, 0x05, 0x97, 0xf8, 0xcb, 0xc9, 0x2a, 0xba, 0x58, 0x8e, 0x43, 0x8e, 0xd9,
	0xe2, 0x48, 0x1f, 0x14, 0x74, 0x06, 0x11, 0xeb, 0x7a, 0xa2, 0xcb, 0x5a, 0x98, 0x2e, 0x13, 0xdb,
	0x25, 0x64, 0x5d, 0x7f, 0x53, 0xd6, 0xe2, 0x56, 0xa1, 0x37, 0xb0, 0xa0, 0xeb, 0xb5, 0x60, 0x5a,
	0x70, 0xa2, 0x36, 0x54, 0x72, 0x87, 0x31, 0xc0, 0xf3, 0x9d, 0x7c, 0x6d, 0xd5, 0x54, 0x42, 0x9c,
	0xb7, 0x75, 0x57, 0x56, 0xf7, 0x1e, 0xc1, 0x29, 0xc0, 0x5a, 0xb6, 0x35, 0x61, 0x94, 0x6d, 0xdc,
	0xb7, 0x9c, 0xe1, 0xd0, 0x28, 0xb9, 0x11, 0xc6, 0x13, 0xfa, 0x8f, 0x26, 0x7c, 0x22, 0x5c, 0xfa,
	0xd7, 0x03, 0xb0, 0xd8, 0x5c, 0xbb, 0x2b, 0x38, 0x76, 0x4f, 0xb5, 0x33, 0x14, 0x0d, 0x32, 0x03,
	0xe4, 0x7c, 0x0c, 0x64, 0x04, 0x1a, 0x47, 0x77, 0x7b, 0xc1, 0xd4, 0x88, 0xd8, 0x61, 0x0c, 0x9b,
	0xe9, 0x3f, 0x4c, 0xc7, 0x93, 0xe2, 0xfb, 0x26, 0xf4, 0x0e, 0x4e, 0x3a, 0xd9, 0x1a, 0xbc, 0xe4,
	0x71, 0xfc, 0x78, 0x7b, 0x94, 0xef, 0x11, 0x9f, 0x02, 0x6c, 0x63, 0x57, 0xb4, 0xb1, 0x74, 0x43,
	0x1c, 0xba, 0x4c, 0x15, 0x6d, 0xd0, 0x6b, 0x70, 0x1b, 0x42, 0x95, 0xb6, 0x0f, 0x31, 0xc4, 0x33,
	0x2b, 0x5c, 0x2a, 0x7d, 0x75, 0x02, 0x73, 0xd6, 0xd6, 0xd9, 0x40, 0x1b, 0x2e, 0x7e, 0x66, 0x03,
	0x2f, 0xfe, 0x78, 0x5e, 0x71, 0x64, 0x7f, 0xd2, 0xf7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x32,
	0x8d, 0xc4, 0x80, 0xbd, 0x03, 0x00, 0x00,
}

const ()
