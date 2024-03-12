// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_cloudagent_logicalnetwork.proto

package network

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type LogicalNetworkRequest struct {
	LogicalNetworks      []*LogicalNetwork `protobuf:"bytes,1,rep,name=LogicalNetworks,proto3" json:"LogicalNetworks,omitempty"`
	OperationType        common.Operation  `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LogicalNetworkRequest) Reset()         { *m = LogicalNetworkRequest{} }
func (m *LogicalNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*LogicalNetworkRequest) ProtoMessage()    {}
func (*LogicalNetworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_37e4212e5c247b0a, []int{0}
}

func (m *LogicalNetworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalNetworkRequest.Unmarshal(m, b)
}
func (m *LogicalNetworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalNetworkRequest.Marshal(b, m, deterministic)
}
func (m *LogicalNetworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalNetworkRequest.Merge(m, src)
}
func (m *LogicalNetworkRequest) XXX_Size() int {
	return xxx_messageInfo_LogicalNetworkRequest.Size(m)
}
func (m *LogicalNetworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalNetworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalNetworkRequest proto.InternalMessageInfo

func (m *LogicalNetworkRequest) GetLogicalNetworks() []*LogicalNetwork {
	if m != nil {
		return m.LogicalNetworks
	}
	return nil
}

func (m *LogicalNetworkRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type LogicalNetworkResponse struct {
	LogicalNetworks      []*LogicalNetwork   `protobuf:"bytes,1,rep,name=LogicalNetworks,proto3" json:"LogicalNetworks,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LogicalNetworkResponse) Reset()         { *m = LogicalNetworkResponse{} }
func (m *LogicalNetworkResponse) String() string { return proto.CompactTextString(m) }
func (*LogicalNetworkResponse) ProtoMessage()    {}
func (*LogicalNetworkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_37e4212e5c247b0a, []int{1}
}

func (m *LogicalNetworkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalNetworkResponse.Unmarshal(m, b)
}
func (m *LogicalNetworkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalNetworkResponse.Marshal(b, m, deterministic)
}
func (m *LogicalNetworkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalNetworkResponse.Merge(m, src)
}
func (m *LogicalNetworkResponse) XXX_Size() int {
	return xxx_messageInfo_LogicalNetworkResponse.Size(m)
}
func (m *LogicalNetworkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalNetworkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalNetworkResponse proto.InternalMessageInfo

func (m *LogicalNetworkResponse) GetLogicalNetworks() []*LogicalNetwork {
	if m != nil {
		return m.LogicalNetworks
	}
	return nil
}

func (m *LogicalNetworkResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *LogicalNetworkResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type LogicalNetwork struct {
	Name                         string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                           string           `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Subnets                      []*LogicalSubnet `protobuf:"bytes,3,rep,name=subnets,proto3" json:"subnets,omitempty"`
	NetworkVirtualizationEnabled bool             `protobuf:"varint,4,opt,name=networkVirtualizationEnabled,proto3" json:"networkVirtualizationEnabled,omitempty"`
	Status                       *common.Status   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	LocationName                 string           `protobuf:"bytes,6,opt,name=locationName,proto3" json:"locationName,omitempty"`
	MacPoolName                  string           `protobuf:"bytes,7,opt,name=macPoolName,proto3" json:"macPoolName,omitempty"`
	Tags                         *common.Tags     `protobuf:"bytes,8,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}         `json:"-"`
	XXX_unrecognized             []byte           `json:"-"`
	XXX_sizecache                int32            `json:"-"`
}

func (m *LogicalNetwork) Reset()         { *m = LogicalNetwork{} }
func (m *LogicalNetwork) String() string { return proto.CompactTextString(m) }
func (*LogicalNetwork) ProtoMessage()    {}
func (*LogicalNetwork) Descriptor() ([]byte, []int) {
	return fileDescriptor_37e4212e5c247b0a, []int{2}
}

func (m *LogicalNetwork) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalNetwork.Unmarshal(m, b)
}
func (m *LogicalNetwork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalNetwork.Marshal(b, m, deterministic)
}
func (m *LogicalNetwork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalNetwork.Merge(m, src)
}
func (m *LogicalNetwork) XXX_Size() int {
	return xxx_messageInfo_LogicalNetwork.Size(m)
}
func (m *LogicalNetwork) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalNetwork.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalNetwork proto.InternalMessageInfo

func (m *LogicalNetwork) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogicalNetwork) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LogicalNetwork) GetSubnets() []*LogicalSubnet {
	if m != nil {
		return m.Subnets
	}
	return nil
}

func (m *LogicalNetwork) GetNetworkVirtualizationEnabled() bool {
	if m != nil {
		return m.NetworkVirtualizationEnabled
	}
	return false
}

func (m *LogicalNetwork) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *LogicalNetwork) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *LogicalNetwork) GetMacPoolName() string {
	if m != nil {
		return m.MacPoolName
	}
	return ""
}

func (m *LogicalNetwork) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

type LogicalSubnet struct {
	Name                 string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string                    `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	AddressPrefix        string                    `protobuf:"bytes,3,opt,name=addressPrefix,proto3" json:"addressPrefix,omitempty"`
	Routes               []*common.Route           `protobuf:"bytes,4,rep,name=routes,proto3" json:"routes,omitempty"`
	Allocation           common.IPAllocationMethod `protobuf:"varint,5,opt,name=allocation,proto3,enum=moc.IPAllocationMethod" json:"allocation,omitempty"`
	Vlan                 uint32                    `protobuf:"varint,6,opt,name=vlan,proto3" json:"vlan,omitempty"`
	IpPools              []*common.IPPool          `protobuf:"bytes,7,rep,name=ipPools,proto3" json:"ipPools,omitempty"`
	Dns                  *common.Dns               `protobuf:"bytes,8,opt,name=dns,proto3" json:"dns,omitempty"`
	IsPublic             bool                      `protobuf:"varint,9,opt,name=isPublic,proto3" json:"isPublic,omitempty"`
	Tags                 *common.Tags              `protobuf:"bytes,10,opt,name=tags,proto3" json:"tags,omitempty"`
	NetworkSecurityGroup *common.NsgGroup          `protobuf:"bytes,11,opt,name=networkSecurityGroup,proto3" json:"networkSecurityGroup,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *LogicalSubnet) Reset()         { *m = LogicalSubnet{} }
func (m *LogicalSubnet) String() string { return proto.CompactTextString(m) }
func (*LogicalSubnet) ProtoMessage()    {}
func (*LogicalSubnet) Descriptor() ([]byte, []int) {
	return fileDescriptor_37e4212e5c247b0a, []int{3}
}

func (m *LogicalSubnet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalSubnet.Unmarshal(m, b)
}
func (m *LogicalSubnet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalSubnet.Marshal(b, m, deterministic)
}
func (m *LogicalSubnet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalSubnet.Merge(m, src)
}
func (m *LogicalSubnet) XXX_Size() int {
	return xxx_messageInfo_LogicalSubnet.Size(m)
}
func (m *LogicalSubnet) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalSubnet.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalSubnet proto.InternalMessageInfo

func (m *LogicalSubnet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogicalSubnet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LogicalSubnet) GetAddressPrefix() string {
	if m != nil {
		return m.AddressPrefix
	}
	return ""
}

func (m *LogicalSubnet) GetRoutes() []*common.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *LogicalSubnet) GetAllocation() common.IPAllocationMethod {
	if m != nil {
		return m.Allocation
	}
	return common.IPAllocationMethod_Invalid
}

func (m *LogicalSubnet) GetVlan() uint32 {
	if m != nil {
		return m.Vlan
	}
	return 0
}

func (m *LogicalSubnet) GetIpPools() []*common.IPPool {
	if m != nil {
		return m.IpPools
	}
	return nil
}

func (m *LogicalSubnet) GetDns() *common.Dns {
	if m != nil {
		return m.Dns
	}
	return nil
}

func (m *LogicalSubnet) GetIsPublic() bool {
	if m != nil {
		return m.IsPublic
	}
	return false
}

func (m *LogicalSubnet) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *LogicalSubnet) GetNetworkSecurityGroup() *common.NsgGroup {
	if m != nil {
		return m.NetworkSecurityGroup
	}
	return nil
}

func init() {
	proto.RegisterType((*LogicalNetworkRequest)(nil), "moc.cloudagent.network.LogicalNetworkRequest")
	proto.RegisterType((*LogicalNetworkResponse)(nil), "moc.cloudagent.network.LogicalNetworkResponse")
	proto.RegisterType((*LogicalNetwork)(nil), "moc.cloudagent.network.LogicalNetwork")
	proto.RegisterType((*LogicalSubnet)(nil), "moc.cloudagent.network.LogicalSubnet")
}

func init() {
	proto.RegisterFile("moc_cloudagent_logicalnetwork.proto", fileDescriptor_37e4212e5c247b0a)
}

var fileDescriptor_37e4212e5c247b0a = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xc7, 0xaf, 0x93, 0x34, 0x1f, 0x27, 0x37, 0xb9, 0xd2, 0xdc, 0xd2, 0x5a, 0x11, 0x54, 0x51,
	0x4a, 0xab, 0x6c, 0xb0, 0x45, 0x40, 0x62, 0x89, 0x1a, 0x51, 0xa1, 0x4a, 0x50, 0xa2, 0x69, 0xd5,
	0x05, 0x9b, 0x6a, 0x62, 0x4f, 0xdd, 0x51, 0xc7, 0x1e, 0x33, 0x33, 0x6e, 0x29, 0x2b, 0x96, 0x6c,
	0x79, 0x02, 0x5e, 0x02, 0x1e, 0x82, 0xb7, 0x42, 0x3e, 0x76, 0xd2, 0xb8, 0xaa, 0xa0, 0x1b, 0x56,
	0xb1, 0xcf, 0xff, 0x77, 0xbe, 0xfe, 0x9e, 0x09, 0x6c, 0xc7, 0x2a, 0x38, 0x0d, 0xa4, 0xca, 0x42,
	0x16, 0xf1, 0xc4, 0x9e, 0x4a, 0x15, 0x89, 0x80, 0xc9, 0x84, 0xdb, 0x2b, 0xa5, 0x2f, 0xbc, 0x54,
	0x2b, 0xab, 0xc8, 0x46, 0xac, 0x02, 0xef, 0x06, 0xf2, 0x4a, 0x75, 0xb0, 0x15, 0x29, 0x15, 0x49,
	0xee, 0x23, 0x35, 0xcf, 0xce, 0xfc, 0x2b, 0xcd, 0xd2, 0x94, 0x6b, 0x53, 0xe4, 0x0d, 0x36, 0xb1,
	0xb8, 0x8a, 0x63, 0x95, 0x94, 0x3f, 0xa5, 0xb0, 0xb5, 0x22, 0x94, 0xc5, 0x56, 0xf5, 0xd1, 0x37,
	0x07, 0x1e, 0xbc, 0x29, 0x26, 0x39, 0x2c, 0x64, 0xca, 0x3f, 0x64, 0xdc, 0x58, 0x32, 0x83, 0xff,
	0xaa, 0x82, 0x71, 0x9d, 0x61, 0x7d, 0xdc, 0x9d, 0xec, 0x7a, 0x77, 0x0f, 0xe9, 0xdd, 0xaa, 0x73,
	0x3b, 0x9d, 0x3c, 0x87, 0xde, 0xbb, 0x94, 0x6b, 0x66, 0x85, 0x4a, 0x8e, 0xaf, 0x53, 0xee, 0xd6,
	0x86, 0xce, 0xb8, 0x3f, 0xe9, 0x63, 0xbd, 0xa5, 0x42, 0xab, 0xd0, 0xe8, 0xbb, 0x03, 0x1b, 0xb7,
	0x27, 0x34, 0xa9, 0x4a, 0x0c, 0xff, 0x0b, 0x23, 0x4e, 0xa0, 0x49, 0xb9, 0xc9, 0xa4, 0xc5, 0xd9,
	0xba, 0x93, 0x81, 0x57, 0x18, 0xef, 0x2d, 0x8c, 0xf7, 0xa6, 0x4a, 0xc9, 0x13, 0x26, 0x33, 0x4e,
	0x4b, 0x92, 0xac, 0xc3, 0xda, 0xbe, 0xd6, 0x4a, 0xbb, 0xf5, 0xa1, 0x33, 0xee, 0xd0, 0xe2, 0x65,
	0xf4, 0xb3, 0x06, 0xfd, 0x6a, 0x75, 0x42, 0xa0, 0x91, 0xb0, 0x98, 0xbb, 0x0e, 0x72, 0xf8, 0x4c,
	0xfa, 0x50, 0x13, 0x21, 0x36, 0xeb, 0xd0, 0x9a, 0x08, 0xc9, 0x4b, 0x68, 0x99, 0x6c, 0x9e, 0x70,
	0x6b, 0xdc, 0x3a, 0xae, 0xb2, 0xf3, 0x87, 0x55, 0x8e, 0x90, 0xa6, 0x8b, 0x2c, 0x32, 0x85, 0x87,
	0x25, 0x71, 0x22, 0xb4, 0xcd, 0x98, 0x14, 0x9f, 0xd0, 0xcb, 0xfd, 0x84, 0xcd, 0x25, 0x0f, 0xdd,
	0xc6, 0xd0, 0x19, 0xb7, 0xe9, 0x6f, 0x19, 0xb2, 0x0d, 0x4d, 0x63, 0x99, 0xcd, 0x8c, 0xbb, 0x86,
	0x2e, 0x74, 0x71, 0x86, 0x23, 0x0c, 0xd1, 0x52, 0x22, 0x23, 0xf8, 0x57, 0xaa, 0x00, 0xf3, 0x0e,
	0xf3, 0xad, 0x9a, 0xb8, 0x43, 0x25, 0x46, 0x76, 0xa1, 0x1b, 0xb3, 0x60, 0xa6, 0x94, 0x44, 0xa4,
	0x95, 0x23, 0xd3, 0xc6, 0x97, 0x1f, 0xae, 0x43, 0x57, 0x05, 0xf2, 0x08, 0x1a, 0x96, 0x45, 0xc6,
	0x6d, 0x63, 0xbb, 0x0e, 0xb6, 0x3b, 0x66, 0x91, 0xa1, 0x18, 0x1e, 0x7d, 0xad, 0x43, 0xaf, 0xb2,
	0xee, 0xbd, 0xac, 0x7c, 0x0c, 0x3d, 0x16, 0x86, 0x9a, 0x1b, 0x33, 0xd3, 0xfc, 0x4c, 0x7c, 0x2c,
	0xbf, 0x4f, 0x35, 0x48, 0x46, 0xd0, 0xd4, 0x2a, 0xb3, 0xdc, 0xb8, 0x0d, 0xf4, 0x1b, 0xb0, 0x39,
	0xcd, 0x43, 0xb4, 0x54, 0xc8, 0x0b, 0x00, 0x26, 0x17, 0x8b, 0xa1, 0x27, 0xfd, 0xc9, 0x26, 0x72,
	0x07, 0xb3, 0xbd, 0xa5, 0xf0, 0x96, 0xdb, 0x73, 0x15, 0xd2, 0x15, 0x34, 0x1f, 0xf3, 0x52, 0xb2,
	0x04, 0xbd, 0xe9, 0x51, 0x7c, 0x26, 0x3b, 0xd0, 0x12, 0x69, 0xbe, 0xb9, 0x71, 0x5b, 0xd8, 0xb1,
	0x5b, 0x56, 0xca, 0x63, 0x74, 0xa1, 0x91, 0x21, 0xd4, 0xc3, 0x64, 0xe1, 0x48, 0x1b, 0x91, 0x57,
	0x89, 0x29, 0xcd, 0xcb, 0x25, 0x32, 0x80, 0xb6, 0x30, 0xb3, 0x6c, 0x2e, 0x45, 0xe0, 0x76, 0xf0,
	0xab, 0x2e, 0xdf, 0x97, 0x86, 0xc2, 0x9d, 0x86, 0x92, 0x3d, 0x58, 0x2f, 0x0f, 0xc0, 0x11, 0x0f,
	0x32, 0x2d, 0xec, 0xf5, 0x6b, 0xad, 0xb2, 0xd4, 0xed, 0x22, 0xde, 0x43, 0xfc, 0xd0, 0x44, 0x18,
	0xa4, 0x77, 0xa2, 0x93, 0xcf, 0x0e, 0xfc, 0x5f, 0x3d, 0xdf, 0x7b, 0xf9, 0xf1, 0x24, 0x02, 0x9a,
	0x07, 0xc9, 0xa5, 0xba, 0xe0, 0xe4, 0xc9, 0x3d, 0x2f, 0x61, 0xf1, 0x7f, 0x33, 0xf0, 0xee, 0x8b,
	0x17, 0x97, 0x7f, 0xf4, 0xcf, 0xf4, 0xe9, 0x7b, 0x3f, 0x12, 0xf6, 0x3c, 0x9b, 0x7b, 0x81, 0x8a,
	0xfd, 0x58, 0x04, 0x5a, 0x19, 0x75, 0x66, 0xfd, 0x58, 0x05, 0xbe, 0x4e, 0x03, 0xff, 0xa6, 0x96,
	0x5f, 0xd6, 0x9a, 0x37, 0xf1, 0x1e, 0x3f, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x5a, 0xa3,
	0x51, 0x8d, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogicalNetworkAgentClient is the client API for LogicalNetworkAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogicalNetworkAgentClient interface {
	Invoke(ctx context.Context, in *LogicalNetworkRequest, opts ...grpc.CallOption) (*LogicalNetworkResponse, error)
}

type logicalNetworkAgentClient struct {
	cc *grpc.ClientConn
}

func NewLogicalNetworkAgentClient(cc *grpc.ClientConn) LogicalNetworkAgentClient {
	return &logicalNetworkAgentClient{cc}
}

func (c *logicalNetworkAgentClient) Invoke(ctx context.Context, in *LogicalNetworkRequest, opts ...grpc.CallOption) (*LogicalNetworkResponse, error) {
	out := new(LogicalNetworkResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.network.LogicalNetworkAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicalNetworkAgentServer is the server API for LogicalNetworkAgent service.
type LogicalNetworkAgentServer interface {
	Invoke(context.Context, *LogicalNetworkRequest) (*LogicalNetworkResponse, error)
}

// UnimplementedLogicalNetworkAgentServer can be embedded to have forward compatible implementations.
type UnimplementedLogicalNetworkAgentServer struct {
}

func (*UnimplementedLogicalNetworkAgentServer) Invoke(ctx context.Context, req *LogicalNetworkRequest) (*LogicalNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterLogicalNetworkAgentServer(s *grpc.Server, srv LogicalNetworkAgentServer) {
	s.RegisterService(&_LogicalNetworkAgent_serviceDesc, srv)
}

func _LogicalNetworkAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogicalNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicalNetworkAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.network.LogicalNetworkAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicalNetworkAgentServer).Invoke(ctx, req.(*LogicalNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogicalNetworkAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.network.LogicalNetworkAgent",
	HandlerType: (*LogicalNetworkAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _LogicalNetworkAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_cloudagent_logicalnetwork.proto",
}
