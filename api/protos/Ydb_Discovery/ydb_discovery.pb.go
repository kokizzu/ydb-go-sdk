// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_discovery.proto

package Ydb_Discovery

import (
	Ydb_Operations "github.com/yandex-cloud/ydb-go-sdk/v2/api/protos/Ydb_Operations"
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

type ListEndpointsRequest struct {
	Database             string   `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Service              []string `protobuf:"bytes,2,rep,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListEndpointsRequest) Reset()         { *m = ListEndpointsRequest{} }
func (m *ListEndpointsRequest) String() string { return proto.CompactTextString(m) }
func (*ListEndpointsRequest) ProtoMessage()    {}
func (*ListEndpointsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b240df5737142d, []int{0}
}

func (m *ListEndpointsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEndpointsRequest.Unmarshal(m, b)
}
func (m *ListEndpointsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEndpointsRequest.Marshal(b, m, deterministic)
}
func (m *ListEndpointsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEndpointsRequest.Merge(m, src)
}
func (m *ListEndpointsRequest) XXX_Size() int {
	return xxx_messageInfo_ListEndpointsRequest.Size(m)
}
func (m *ListEndpointsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEndpointsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListEndpointsRequest proto.InternalMessageInfo

func (m *ListEndpointsRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *ListEndpointsRequest) GetService() []string {
	if m != nil {
		return m.Service
	}
	return nil
}

type EndpointInfo struct {
	// This is an address (usually fqdn) and port of this node's grpc endpoint
	Address    string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port       uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	LoadFactor float32  `protobuf:"fixed32,3,opt,name=load_factor,json=loadFactor,proto3" json:"load_factor,omitempty"`
	Ssl        bool     `protobuf:"varint,4,opt,name=ssl,proto3" json:"ssl,omitempty"`
	Service    []string `protobuf:"bytes,5,rep,name=service,proto3" json:"service,omitempty"`
	Location   string   `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	NodeId     uint32   `protobuf:"varint,7,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Optional ipv4 and/or ipv6 addresses of the endpoint, which clients may
	// use instead of a dns name in the address field.
	IpV4 []string `protobuf:"bytes,8,rep,name=ip_v4,json=ipV4,proto3" json:"ip_v4,omitempty"`
	IpV6 []string `protobuf:"bytes,9,rep,name=ip_v6,json=ipV6,proto3" json:"ip_v6,omitempty"`
	// Optional value for grpc.ssl_target_name_override option that must be
	// used when connecting to this endpoint. This may be specified when an ssl
	// endpoint is using certificate chain valid for a balancer hostname, and
	// not this specific node hostname.
	SslTargetNameOverride string   `protobuf:"bytes,10,opt,name=ssl_target_name_override,json=sslTargetNameOverride,proto3" json:"ssl_target_name_override,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *EndpointInfo) Reset()         { *m = EndpointInfo{} }
func (m *EndpointInfo) String() string { return proto.CompactTextString(m) }
func (*EndpointInfo) ProtoMessage()    {}
func (*EndpointInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b240df5737142d, []int{1}
}

func (m *EndpointInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointInfo.Unmarshal(m, b)
}
func (m *EndpointInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointInfo.Marshal(b, m, deterministic)
}
func (m *EndpointInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointInfo.Merge(m, src)
}
func (m *EndpointInfo) XXX_Size() int {
	return xxx_messageInfo_EndpointInfo.Size(m)
}
func (m *EndpointInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointInfo proto.InternalMessageInfo

func (m *EndpointInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EndpointInfo) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *EndpointInfo) GetLoadFactor() float32 {
	if m != nil {
		return m.LoadFactor
	}
	return 0
}

func (m *EndpointInfo) GetSsl() bool {
	if m != nil {
		return m.Ssl
	}
	return false
}

func (m *EndpointInfo) GetService() []string {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *EndpointInfo) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *EndpointInfo) GetNodeId() uint32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *EndpointInfo) GetIpV4() []string {
	if m != nil {
		return m.IpV4
	}
	return nil
}

func (m *EndpointInfo) GetIpV6() []string {
	if m != nil {
		return m.IpV6
	}
	return nil
}

func (m *EndpointInfo) GetSslTargetNameOverride() string {
	if m != nil {
		return m.SslTargetNameOverride
	}
	return ""
}

type ListEndpointsResult struct {
	Endpoints            []*EndpointInfo `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	SelfLocation         string          `protobuf:"bytes,2,opt,name=self_location,json=selfLocation,proto3" json:"self_location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListEndpointsResult) Reset()         { *m = ListEndpointsResult{} }
func (m *ListEndpointsResult) String() string { return proto.CompactTextString(m) }
func (*ListEndpointsResult) ProtoMessage()    {}
func (*ListEndpointsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b240df5737142d, []int{2}
}

func (m *ListEndpointsResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEndpointsResult.Unmarshal(m, b)
}
func (m *ListEndpointsResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEndpointsResult.Marshal(b, m, deterministic)
}
func (m *ListEndpointsResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEndpointsResult.Merge(m, src)
}
func (m *ListEndpointsResult) XXX_Size() int {
	return xxx_messageInfo_ListEndpointsResult.Size(m)
}
func (m *ListEndpointsResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEndpointsResult.DiscardUnknown(m)
}

var xxx_messageInfo_ListEndpointsResult proto.InternalMessageInfo

func (m *ListEndpointsResult) GetEndpoints() []*EndpointInfo {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ListEndpointsResult) GetSelfLocation() string {
	if m != nil {
		return m.SelfLocation
	}
	return ""
}

type ListEndpointsResponse struct {
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ListEndpointsResponse) Reset()         { *m = ListEndpointsResponse{} }
func (m *ListEndpointsResponse) String() string { return proto.CompactTextString(m) }
func (*ListEndpointsResponse) ProtoMessage()    {}
func (*ListEndpointsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b240df5737142d, []int{3}
}

func (m *ListEndpointsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEndpointsResponse.Unmarshal(m, b)
}
func (m *ListEndpointsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEndpointsResponse.Marshal(b, m, deterministic)
}
func (m *ListEndpointsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEndpointsResponse.Merge(m, src)
}
func (m *ListEndpointsResponse) XXX_Size() int {
	return xxx_messageInfo_ListEndpointsResponse.Size(m)
}
func (m *ListEndpointsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEndpointsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListEndpointsResponse proto.InternalMessageInfo

func (m *ListEndpointsResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type WhoAmIRequest struct {
	// Include user groups in response
	IncludeGroups        bool     `protobuf:"varint,1,opt,name=include_groups,json=includeGroups,proto3" json:"include_groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WhoAmIRequest) Reset()         { *m = WhoAmIRequest{} }
func (m *WhoAmIRequest) String() string { return proto.CompactTextString(m) }
func (*WhoAmIRequest) ProtoMessage()    {}
func (*WhoAmIRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b240df5737142d, []int{4}
}

func (m *WhoAmIRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhoAmIRequest.Unmarshal(m, b)
}
func (m *WhoAmIRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhoAmIRequest.Marshal(b, m, deterministic)
}
func (m *WhoAmIRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhoAmIRequest.Merge(m, src)
}
func (m *WhoAmIRequest) XXX_Size() int {
	return xxx_messageInfo_WhoAmIRequest.Size(m)
}
func (m *WhoAmIRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WhoAmIRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WhoAmIRequest proto.InternalMessageInfo

func (m *WhoAmIRequest) GetIncludeGroups() bool {
	if m != nil {
		return m.IncludeGroups
	}
	return false
}

type WhoAmIResult struct {
	// User SID (Security ID)
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// List of group SIDs (Security IDs) for the user
	Groups               []string `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WhoAmIResult) Reset()         { *m = WhoAmIResult{} }
func (m *WhoAmIResult) String() string { return proto.CompactTextString(m) }
func (*WhoAmIResult) ProtoMessage()    {}
func (*WhoAmIResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b240df5737142d, []int{5}
}

func (m *WhoAmIResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhoAmIResult.Unmarshal(m, b)
}
func (m *WhoAmIResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhoAmIResult.Marshal(b, m, deterministic)
}
func (m *WhoAmIResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhoAmIResult.Merge(m, src)
}
func (m *WhoAmIResult) XXX_Size() int {
	return xxx_messageInfo_WhoAmIResult.Size(m)
}
func (m *WhoAmIResult) XXX_DiscardUnknown() {
	xxx_messageInfo_WhoAmIResult.DiscardUnknown(m)
}

var xxx_messageInfo_WhoAmIResult proto.InternalMessageInfo

func (m *WhoAmIResult) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *WhoAmIResult) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

type WhoAmIResponse struct {
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *WhoAmIResponse) Reset()         { *m = WhoAmIResponse{} }
func (m *WhoAmIResponse) String() string { return proto.CompactTextString(m) }
func (*WhoAmIResponse) ProtoMessage()    {}
func (*WhoAmIResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b240df5737142d, []int{6}
}

func (m *WhoAmIResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhoAmIResponse.Unmarshal(m, b)
}
func (m *WhoAmIResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhoAmIResponse.Marshal(b, m, deterministic)
}
func (m *WhoAmIResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhoAmIResponse.Merge(m, src)
}
func (m *WhoAmIResponse) XXX_Size() int {
	return xxx_messageInfo_WhoAmIResponse.Size(m)
}
func (m *WhoAmIResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WhoAmIResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WhoAmIResponse proto.InternalMessageInfo

func (m *WhoAmIResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func init() {
	proto.RegisterType((*ListEndpointsRequest)(nil), "Ydb.Discovery.ListEndpointsRequest")
	proto.RegisterType((*EndpointInfo)(nil), "Ydb.Discovery.EndpointInfo")
	proto.RegisterType((*ListEndpointsResult)(nil), "Ydb.Discovery.ListEndpointsResult")
	proto.RegisterType((*ListEndpointsResponse)(nil), "Ydb.Discovery.ListEndpointsResponse")
	proto.RegisterType((*WhoAmIRequest)(nil), "Ydb.Discovery.WhoAmIRequest")
	proto.RegisterType((*WhoAmIResult)(nil), "Ydb.Discovery.WhoAmIResult")
	proto.RegisterType((*WhoAmIResponse)(nil), "Ydb.Discovery.WhoAmIResponse")
}

func init() { proto.RegisterFile("ydb_discovery.proto", fileDescriptor_a2b240df5737142d) }

var fileDescriptor_a2b240df5737142d = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x5b, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0xb6, 0xeb, 0xe5, 0xb4, 0x1d, 0xc8, 0x65, 0x60, 0xc6, 0x03, 0x51, 0x10, 0x52,
	0x78, 0x49, 0xa5, 0x31, 0x75, 0x82, 0x37, 0x26, 0x2e, 0xaa, 0x54, 0xb1, 0x29, 0x42, 0x20, 0x9e,
	0x22, 0x27, 0x76, 0x87, 0xa5, 0x34, 0x0e, 0x3e, 0x4e, 0x45, 0x3f, 0x0e, 0xdf, 0x92, 0x47, 0x14,
	0x37, 0x4e, 0xb7, 0xbd, 0xee, 0xed, 0xfc, 0x7f, 0x3e, 0x17, 0x9f, 0xe3, 0x63, 0x98, 0xed, 0x78,
	0x9a, 0x70, 0x89, 0x99, 0xda, 0x0a, 0xbd, 0x8b, 0x4a, 0xad, 0x8c, 0x22, 0xd3, 0x9f, 0x3c, 0x8d,
	0x3e, 0x3a, 0x78, 0xfa, 0x66, 0xc7, 0xd3, 0x79, 0x59, 0xa5, 0xb9, 0xcc, 0xe6, 0xac, 0x94, 0x73,
	0xeb, 0x84, 0xf3, 0x3a, 0x52, 0x95, 0x42, 0x33, 0x23, 0x55, 0xb1, 0x8f, 0x0c, 0x56, 0xf0, 0x64,
	0x25, 0xd1, 0x7c, 0x2a, 0x78, 0xa9, 0x64, 0x61, 0x30, 0x16, 0xbf, 0x2b, 0x81, 0x86, 0x9c, 0xc2,
	0x90, 0x33, 0xc3, 0x52, 0x86, 0x82, 0x7a, 0xbe, 0x17, 0x8e, 0xe2, 0x56, 0x13, 0x0a, 0x03, 0x14,
	0x7a, 0x2b, 0x33, 0x41, 0x3b, 0x7e, 0x37, 0x1c, 0xc5, 0x4e, 0x06, 0x7f, 0x3b, 0x30, 0x71, 0xa9,
	0x96, 0xc5, 0x5a, 0xd5, 0xae, 0x8c, 0x73, 0x2d, 0x10, 0x9b, 0x2c, 0x4e, 0x12, 0x02, 0xbd, 0x52,
	0x69, 0x43, 0x3b, 0xbe, 0x17, 0x4e, 0x63, 0x6b, 0x93, 0x97, 0x30, 0xce, 0x15, 0xe3, 0xc9, 0x9a,
	0x65, 0x46, 0x69, 0xda, 0xf5, 0xbd, 0xb0, 0x13, 0x43, 0x8d, 0x3e, 0x5b, 0x42, 0x1e, 0x43, 0x17,
	0x31, 0xa7, 0x3d, 0xdf, 0x0b, 0x87, 0x71, 0x6d, 0xde, 0xbe, 0xcb, 0xd1, 0x9d, 0xbb, 0xd4, 0x1d,
	0xe4, 0x2a, 0xb3, 0xbd, 0xd2, 0xfe, 0xbe, 0x03, 0xa7, 0xc9, 0x33, 0x18, 0x14, 0x8a, 0x8b, 0x44,
	0x72, 0x3a, 0xb0, 0xf5, 0xfb, 0xb5, 0x5c, 0x72, 0x32, 0x83, 0x23, 0x59, 0x26, 0xdb, 0x73, 0x3a,
	0xb4, 0xc9, 0x7a, 0xb2, 0xfc, 0x7e, 0xee, 0xe0, 0x82, 0x8e, 0x5a, 0xb8, 0x20, 0x17, 0x40, 0x11,
	0xf3, 0xc4, 0x30, 0x7d, 0x23, 0x4c, 0x52, 0xb0, 0x8d, 0x48, 0xea, 0xe1, 0x6b, 0xc9, 0x05, 0x05,
	0x5b, 0xee, 0x04, 0x31, 0xff, 0x66, 0x8f, 0xbf, 0xb2, 0x8d, 0xb8, 0x6a, 0x0e, 0x83, 0x0a, 0x66,
	0xf7, 0x26, 0x8e, 0x55, 0x6e, 0xc8, 0x3b, 0x18, 0x09, 0x87, 0xa8, 0xe7, 0x77, 0xc3, 0xf1, 0xd9,
	0x8b, 0xe8, 0xce, 0xb3, 0x46, 0xb7, 0x27, 0x1b, 0x1f, 0xbc, 0xc9, 0x2b, 0x98, 0xa2, 0xc8, 0xd7,
	0x49, 0xdb, 0x6e, 0xc7, 0xd6, 0x9f, 0xd4, 0x70, 0xd5, 0xb0, 0xe0, 0x1a, 0x4e, 0xee, 0x97, 0x2d,
	0x55, 0x81, 0x82, 0x5c, 0xc0, 0xa8, 0x5d, 0x0a, 0xfb, 0x48, 0xe3, 0xb3, 0xe7, 0xb6, 0xf0, 0x95,
	0xa3, 0x78, 0x30, 0xe3, 0x83, 0x6f, 0xb0, 0x80, 0xe9, 0x8f, 0x5f, 0xea, 0xc3, 0x66, 0xe9, 0x76,
	0xe6, 0x35, 0x1c, 0xcb, 0x22, 0xcb, 0x2b, 0x2e, 0x92, 0x1b, 0xad, 0xaa, 0x72, 0xff, 0xe6, 0xc3,
	0x78, 0xda, 0xd0, 0x2f, 0x16, 0x06, 0xef, 0x61, 0xe2, 0xe2, 0x6c, 0xe7, 0x04, 0x7a, 0x15, 0x0a,
	0xdd, 0x2c, 0x88, 0xb5, 0xc9, 0x53, 0xe8, 0x37, 0x29, 0xf6, 0x1b, 0xd6, 0xa8, 0x60, 0x09, 0xc7,
	0x6d, 0xec, 0xc3, 0xae, 0x7f, 0x19, 0x01, 0xcd, 0xd4, 0x26, 0xda, 0xb1, 0x82, 0x8b, 0x3f, 0xd1,
	0x8e, 0xa7, 0x51, 0xfb, 0xab, 0x2e, 0x1f, 0xb5, 0x43, 0xbf, 0xb6, 0x5f, 0xe7, 0x9f, 0xe7, 0xa5,
	0x7d, 0xfb, 0x61, 0xde, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xcf, 0xd4, 0x1d, 0x81, 0x03,
	0x00, 0x00,
}
