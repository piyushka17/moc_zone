// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_cloudagent_loadbalancer.proto

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

type LoadBalancerRequest struct {
	LoadBalancers        []*LoadBalancer  `protobuf:"bytes,1,rep,name=LoadBalancers,proto3" json:"LoadBalancers,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LoadBalancerRequest) Reset()         { *m = LoadBalancerRequest{} }
func (m *LoadBalancerRequest) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerRequest) ProtoMessage()    {}
func (*LoadBalancerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7464476b31ac10f8, []int{0}
}

func (m *LoadBalancerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerRequest.Unmarshal(m, b)
}
func (m *LoadBalancerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerRequest.Marshal(b, m, deterministic)
}
func (m *LoadBalancerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerRequest.Merge(m, src)
}
func (m *LoadBalancerRequest) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerRequest.Size(m)
}
func (m *LoadBalancerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerRequest proto.InternalMessageInfo

func (m *LoadBalancerRequest) GetLoadBalancers() []*LoadBalancer {
	if m != nil {
		return m.LoadBalancers
	}
	return nil
}

func (m *LoadBalancerRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type LoadBalancerResponse struct {
	LoadBalancers        []*LoadBalancer     `protobuf:"bytes,1,rep,name=LoadBalancers,proto3" json:"LoadBalancers,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LoadBalancerResponse) Reset()         { *m = LoadBalancerResponse{} }
func (m *LoadBalancerResponse) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerResponse) ProtoMessage()    {}
func (*LoadBalancerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7464476b31ac10f8, []int{1}
}

func (m *LoadBalancerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerResponse.Unmarshal(m, b)
}
func (m *LoadBalancerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerResponse.Marshal(b, m, deterministic)
}
func (m *LoadBalancerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerResponse.Merge(m, src)
}
func (m *LoadBalancerResponse) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerResponse.Size(m)
}
func (m *LoadBalancerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerResponse proto.InternalMessageInfo

func (m *LoadBalancerResponse) GetLoadBalancers() []*LoadBalancer {
	if m != nil {
		return m.LoadBalancers
	}
	return nil
}

func (m *LoadBalancerResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *LoadBalancerResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type LoadBalancerPrecheckRequest struct {
	LoadBalancers        []*LoadBalancer `protobuf:"bytes,1,rep,name=LoadBalancers,proto3" json:"LoadBalancers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LoadBalancerPrecheckRequest) Reset()         { *m = LoadBalancerPrecheckRequest{} }
func (m *LoadBalancerPrecheckRequest) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerPrecheckRequest) ProtoMessage()    {}
func (*LoadBalancerPrecheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7464476b31ac10f8, []int{2}
}

func (m *LoadBalancerPrecheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerPrecheckRequest.Unmarshal(m, b)
}
func (m *LoadBalancerPrecheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerPrecheckRequest.Marshal(b, m, deterministic)
}
func (m *LoadBalancerPrecheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerPrecheckRequest.Merge(m, src)
}
func (m *LoadBalancerPrecheckRequest) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerPrecheckRequest.Size(m)
}
func (m *LoadBalancerPrecheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerPrecheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerPrecheckRequest proto.InternalMessageInfo

func (m *LoadBalancerPrecheckRequest) GetLoadBalancers() []*LoadBalancer {
	if m != nil {
		return m.LoadBalancers
	}
	return nil
}

type LoadBalancerPrecheckResponse struct {
	// The precheck result: true if the precheck criteria is passed; otherwise, false
	Result *wrappers.BoolValue `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	// The error message if the precheck is not passed; otherwise, empty string
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadBalancerPrecheckResponse) Reset()         { *m = LoadBalancerPrecheckResponse{} }
func (m *LoadBalancerPrecheckResponse) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerPrecheckResponse) ProtoMessage()    {}
func (*LoadBalancerPrecheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7464476b31ac10f8, []int{3}
}

func (m *LoadBalancerPrecheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerPrecheckResponse.Unmarshal(m, b)
}
func (m *LoadBalancerPrecheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerPrecheckResponse.Marshal(b, m, deterministic)
}
func (m *LoadBalancerPrecheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerPrecheckResponse.Merge(m, src)
}
func (m *LoadBalancerPrecheckResponse) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerPrecheckResponse.Size(m)
}
func (m *LoadBalancerPrecheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerPrecheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerPrecheckResponse proto.InternalMessageInfo

func (m *LoadBalancerPrecheckResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *LoadBalancerPrecheckResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type LoadbalancerInboundNatRule struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	FrontendPort         uint32          `protobuf:"varint,2,opt,name=frontendPort,proto3" json:"frontendPort,omitempty"`
	BackendPort          uint32          `protobuf:"varint,3,opt,name=backendPort,proto3" json:"backendPort,omitempty"`
	Protocol             common.Protocol `protobuf:"varint,4,opt,name=protocol,proto3,enum=moc.Protocol" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LoadbalancerInboundNatRule) Reset()         { *m = LoadbalancerInboundNatRule{} }
func (m *LoadbalancerInboundNatRule) String() string { return proto.CompactTextString(m) }
func (*LoadbalancerInboundNatRule) ProtoMessage()    {}
func (*LoadbalancerInboundNatRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_7464476b31ac10f8, []int{4}
}

func (m *LoadbalancerInboundNatRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadbalancerInboundNatRule.Unmarshal(m, b)
}
func (m *LoadbalancerInboundNatRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadbalancerInboundNatRule.Marshal(b, m, deterministic)
}
func (m *LoadbalancerInboundNatRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadbalancerInboundNatRule.Merge(m, src)
}
func (m *LoadbalancerInboundNatRule) XXX_Size() int {
	return xxx_messageInfo_LoadbalancerInboundNatRule.Size(m)
}
func (m *LoadbalancerInboundNatRule) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadbalancerInboundNatRule.DiscardUnknown(m)
}

var xxx_messageInfo_LoadbalancerInboundNatRule proto.InternalMessageInfo

func (m *LoadbalancerInboundNatRule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoadbalancerInboundNatRule) GetFrontendPort() uint32 {
	if m != nil {
		return m.FrontendPort
	}
	return 0
}

func (m *LoadbalancerInboundNatRule) GetBackendPort() uint32 {
	if m != nil {
		return m.BackendPort
	}
	return 0
}

func (m *LoadbalancerInboundNatRule) GetProtocol() common.Protocol {
	if m != nil {
		return m.Protocol
	}
	return common.Protocol_All
}

type LoadBalancingRule struct {
	FrontendPort         uint32          `protobuf:"varint,1,opt,name=frontendPort,proto3" json:"frontendPort,omitempty"`
	BackendPort          uint32          `protobuf:"varint,2,opt,name=backendPort,proto3" json:"backendPort,omitempty"`
	Protocol             common.Protocol `protobuf:"varint,3,opt,name=protocol,proto3,enum=moc.Protocol" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LoadBalancingRule) Reset()         { *m = LoadBalancingRule{} }
func (m *LoadBalancingRule) String() string { return proto.CompactTextString(m) }
func (*LoadBalancingRule) ProtoMessage()    {}
func (*LoadBalancingRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_7464476b31ac10f8, []int{5}
}

func (m *LoadBalancingRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancingRule.Unmarshal(m, b)
}
func (m *LoadBalancingRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancingRule.Marshal(b, m, deterministic)
}
func (m *LoadBalancingRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancingRule.Merge(m, src)
}
func (m *LoadBalancingRule) XXX_Size() int {
	return xxx_messageInfo_LoadBalancingRule.Size(m)
}
func (m *LoadBalancingRule) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancingRule.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancingRule proto.InternalMessageInfo

func (m *LoadBalancingRule) GetFrontendPort() uint32 {
	if m != nil {
		return m.FrontendPort
	}
	return 0
}

func (m *LoadBalancingRule) GetBackendPort() uint32 {
	if m != nil {
		return m.BackendPort
	}
	return 0
}

func (m *LoadBalancingRule) GetProtocol() common.Protocol {
	if m != nil {
		return m.Protocol
	}
	return common.Protocol_All
}

type LoadBalancer struct {
	Name                 string                        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string                        `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	FrontendIP           string                        `protobuf:"bytes,3,opt,name=frontendIP,proto3" json:"frontendIP,omitempty"`
	Backendpoolnames     []string                      `protobuf:"bytes,4,rep,name=backendpoolnames,proto3" json:"backendpoolnames,omitempty"`
	Networkid            string                        `protobuf:"bytes,5,opt,name=networkid,proto3" json:"networkid,omitempty"`
	Loadbalancingrules   []*LoadBalancingRule          `protobuf:"bytes,6,rep,name=loadbalancingrules,proto3" json:"loadbalancingrules,omitempty"`
	Nodefqdn             string                        `protobuf:"bytes,7,opt,name=nodefqdn,proto3" json:"nodefqdn,omitempty"`
	GroupName            string                        `protobuf:"bytes,8,opt,name=groupName,proto3" json:"groupName,omitempty"`
	LocationName         string                        `protobuf:"bytes,9,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Status               *common.Status                `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	Tags                 *common.Tags                  `protobuf:"bytes,11,opt,name=tags,proto3" json:"tags,omitempty"`
	ReplicationCount     uint32                        `protobuf:"varint,12,opt,name=replicationCount,proto3" json:"replicationCount,omitempty"`
	InboundNatRules      []*LoadbalancerInboundNatRule `protobuf:"bytes,13,rep,name=inboundNatRules,proto3" json:"inboundNatRules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *LoadBalancer) Reset()         { *m = LoadBalancer{} }
func (m *LoadBalancer) String() string { return proto.CompactTextString(m) }
func (*LoadBalancer) ProtoMessage()    {}
func (*LoadBalancer) Descriptor() ([]byte, []int) {
	return fileDescriptor_7464476b31ac10f8, []int{6}
}

func (m *LoadBalancer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancer.Unmarshal(m, b)
}
func (m *LoadBalancer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancer.Marshal(b, m, deterministic)
}
func (m *LoadBalancer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancer.Merge(m, src)
}
func (m *LoadBalancer) XXX_Size() int {
	return xxx_messageInfo_LoadBalancer.Size(m)
}
func (m *LoadBalancer) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancer.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancer proto.InternalMessageInfo

func (m *LoadBalancer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoadBalancer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LoadBalancer) GetFrontendIP() string {
	if m != nil {
		return m.FrontendIP
	}
	return ""
}

func (m *LoadBalancer) GetBackendpoolnames() []string {
	if m != nil {
		return m.Backendpoolnames
	}
	return nil
}

func (m *LoadBalancer) GetNetworkid() string {
	if m != nil {
		return m.Networkid
	}
	return ""
}

func (m *LoadBalancer) GetLoadbalancingrules() []*LoadBalancingRule {
	if m != nil {
		return m.Loadbalancingrules
	}
	return nil
}

func (m *LoadBalancer) GetNodefqdn() string {
	if m != nil {
		return m.Nodefqdn
	}
	return ""
}

func (m *LoadBalancer) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *LoadBalancer) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *LoadBalancer) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *LoadBalancer) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *LoadBalancer) GetReplicationCount() uint32 {
	if m != nil {
		return m.ReplicationCount
	}
	return 0
}

func (m *LoadBalancer) GetInboundNatRules() []*LoadbalancerInboundNatRule {
	if m != nil {
		return m.InboundNatRules
	}
	return nil
}

func init() {
	proto.RegisterType((*LoadBalancerRequest)(nil), "moc.cloudagent.network.LoadBalancerRequest")
	proto.RegisterType((*LoadBalancerResponse)(nil), "moc.cloudagent.network.LoadBalancerResponse")
	proto.RegisterType((*LoadBalancerPrecheckRequest)(nil), "moc.cloudagent.network.LoadBalancerPrecheckRequest")
	proto.RegisterType((*LoadBalancerPrecheckResponse)(nil), "moc.cloudagent.network.LoadBalancerPrecheckResponse")
	proto.RegisterType((*LoadbalancerInboundNatRule)(nil), "moc.cloudagent.network.LoadbalancerInboundNatRule")
	proto.RegisterType((*LoadBalancingRule)(nil), "moc.cloudagent.network.LoadBalancingRule")
	proto.RegisterType((*LoadBalancer)(nil), "moc.cloudagent.network.LoadBalancer")
}

func init() { proto.RegisterFile("moc_cloudagent_loadbalancer.proto", fileDescriptor_7464476b31ac10f8) }

var fileDescriptor_7464476b31ac10f8 = []byte{
	// 694 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0xd3, 0x4a,
	0x14, 0x7d, 0x4e, 0xd2, 0xbc, 0xe4, 0xa6, 0xe9, 0x7b, 0x6f, 0x5e, 0x05, 0x56, 0x28, 0x55, 0x08,
	0x5d, 0xa4, 0x80, 0x6c, 0x91, 0xf6, 0x0f, 0x10, 0xc4, 0xa2, 0xa8, 0x2a, 0xd1, 0x50, 0x21, 0x81,
	0x90, 0xaa, 0x89, 0x3d, 0x71, 0xad, 0xd8, 0x73, 0xdd, 0x99, 0x31, 0x15, 0x7b, 0x16, 0xfc, 0x03,
	0xd6, 0xec, 0x91, 0x58, 0xf1, 0xcb, 0xf8, 0x03, 0x28, 0x63, 0x27, 0xb1, 0xdb, 0xa8, 0x1f, 0x12,
	0xac, 0x12, 0x9f, 0x7b, 0xe6, 0xcc, 0x99, 0x33, 0x77, 0x2e, 0x3c, 0x88, 0xd1, 0x3b, 0xf1, 0x22,
	0x4c, 0x7d, 0x16, 0x70, 0xa1, 0x4f, 0x22, 0x64, 0xfe, 0x98, 0x45, 0x4c, 0x78, 0x5c, 0x3a, 0x89,
	0x44, 0x8d, 0xe4, 0x4e, 0x8c, 0x9e, 0xb3, 0xa4, 0x38, 0x82, 0xeb, 0x73, 0x94, 0xd3, 0xce, 0x76,
	0x80, 0x18, 0x44, 0xdc, 0x35, 0xac, 0x71, 0x3a, 0x71, 0xcf, 0x25, 0x4b, 0x12, 0x2e, 0x55, 0xb6,
	0xae, 0x73, 0xd7, 0x48, 0x63, 0x1c, 0xa3, 0xc8, 0x7f, 0xf2, 0xc2, 0x76, 0xa1, 0x90, 0x8b, 0x15,
	0xeb, 0xbd, 0x2f, 0x16, 0xfc, 0x7f, 0x88, 0xcc, 0x1f, 0xe6, 0x3e, 0x28, 0x3f, 0x4b, 0xb9, 0xd2,
	0xe4, 0x25, 0xb4, 0x8b, 0xb0, 0xb2, 0xad, 0x6e, 0xb5, 0xdf, 0x1a, 0xec, 0x38, 0xab, 0x0d, 0x3a,
	0x25, 0x8d, 0xf2, 0x52, 0xb2, 0x0f, 0xed, 0x57, 0x09, 0x97, 0x4c, 0x87, 0x28, 0x8e, 0x3f, 0x26,
	0xdc, 0xae, 0x74, 0xad, 0xfe, 0xc6, 0x60, 0xc3, 0x68, 0x2d, 0x2a, 0xb4, 0x4c, 0xea, 0x7d, 0xb3,
	0x60, 0xb3, 0xec, 0x4c, 0x25, 0x28, 0x14, 0xff, 0xad, 0xd6, 0x06, 0x50, 0xa7, 0x5c, 0xa5, 0x91,
	0x36, 0x9e, 0x5a, 0x83, 0x8e, 0x93, 0x05, 0xed, 0xcc, 0x83, 0x76, 0x86, 0x88, 0xd1, 0x1b, 0x16,
	0xa5, 0x9c, 0xe6, 0x4c, 0xb2, 0x09, 0x6b, 0x2f, 0xa4, 0x44, 0x69, 0x57, 0xbb, 0x56, 0xbf, 0x49,
	0xb3, 0x8f, 0x5e, 0x08, 0xf7, 0x8a, 0xd2, 0x23, 0xc9, 0xbd, 0x53, 0xee, 0x4d, 0xff, 0x40, 0x9e,
	0xbd, 0x53, 0xd8, 0x5a, 0xbd, 0x55, 0x1e, 0xd0, 0xf2, 0x50, 0xd6, 0xed, 0x0f, 0x55, 0x29, 0x1e,
	0xea, 0xab, 0x05, 0x9d, 0xc3, 0x42, 0x97, 0x1e, 0x88, 0x31, 0xa6, 0xc2, 0x3f, 0x62, 0x9a, 0xa6,
	0x11, 0x27, 0x04, 0x6a, 0x82, 0xc5, 0xdc, 0x6c, 0xd3, 0xa4, 0xe6, 0x3f, 0xe9, 0xc1, 0xfa, 0x44,
	0xa2, 0xd0, 0x5c, 0xf8, 0x23, 0x94, 0x59, 0xae, 0x6d, 0x5a, 0xc2, 0x48, 0x17, 0x5a, 0x63, 0xe6,
	0x4d, 0xe7, 0x94, 0xaa, 0xa1, 0x14, 0x21, 0xb2, 0x0b, 0x0d, 0x63, 0xd6, 0xc3, 0xc8, 0xae, 0x99,
	0x6e, 0x69, 0x9b, 0xa4, 0x46, 0x39, 0x48, 0x17, 0xe5, 0xde, 0x27, 0x0b, 0xfe, 0x5b, 0xc6, 0x11,
	0x8a, 0xc0, 0x58, 0xbb, 0x68, 0xc3, 0xba, 0xde, 0x46, 0xe5, 0x6a, 0x1b, 0xd5, 0xab, 0x6d, 0x7c,
	0xaf, 0xc1, 0x7a, 0xf1, 0x56, 0x56, 0x86, 0xb3, 0x01, 0x95, 0xd0, 0xcf, 0x23, 0xae, 0x84, 0x3e,
	0xd9, 0x01, 0x98, 0x3b, 0x3a, 0x18, 0x65, 0xfd, 0x34, 0xac, 0x7d, 0xfe, 0x61, 0x5b, 0xb4, 0x80,
	0x93, 0x47, 0xf0, 0x6f, 0x6e, 0x2a, 0x41, 0x8c, 0x66, 0x42, 0xca, 0xae, 0x75, 0xab, 0xfd, 0x26,
	0xbd, 0x84, 0x93, 0x2d, 0x68, 0xe6, 0x2d, 0x14, 0xfa, 0xf6, 0x9a, 0xd9, 0x68, 0x09, 0x90, 0xb7,
	0x40, 0x96, 0x43, 0x27, 0x14, 0x81, 0x4c, 0x23, 0xae, 0xec, 0xba, 0x69, 0xc5, 0xdd, 0xeb, 0x5b,
	0x31, 0x0f, 0x97, 0xae, 0x10, 0x21, 0x5d, 0x68, 0x08, 0xf4, 0xf9, 0xe4, 0xcc, 0x17, 0xf6, 0xdf,
	0x85, 0x83, 0x2c, 0xd0, 0x99, 0xb5, 0x40, 0x62, 0x9a, 0x1c, 0xcd, 0x52, 0x69, 0x64, 0xd6, 0x16,
	0xc0, 0xec, 0xc2, 0x22, 0xf4, 0xcc, 0xf3, 0x37, 0x84, 0xa6, 0x21, 0x94, 0x30, 0xf2, 0x10, 0xea,
	0x4a, 0x33, 0x9d, 0x2a, 0x1b, 0x4c, 0x63, 0xb7, 0x8c, 0xe5, 0xd7, 0x06, 0xa2, 0x79, 0x89, 0xdc,
	0x87, 0x9a, 0x66, 0x81, 0xb2, 0x5b, 0x86, 0xd2, 0x34, 0x94, 0x63, 0x16, 0x28, 0x6a, 0xe0, 0x59,
	0x98, 0x92, 0x27, 0x51, 0x98, 0xc9, 0x3e, 0xc7, 0x54, 0x68, 0x7b, 0xdd, 0xdc, 0xfc, 0x25, 0x9c,
	0xbc, 0x87, 0x7f, 0xc2, 0x52, 0xc7, 0x2b, 0xbb, 0x6d, 0xb2, 0x1a, 0x5c, 0x95, 0xd5, 0xea, 0xc7,
	0x42, 0x2f, 0x4a, 0x0d, 0x7e, 0x96, 0x1a, 0x97, 0xcb, 0x67, 0x33, 0x25, 0xc2, 0xa1, 0x7e, 0x20,
	0x3e, 0xe0, 0x94, 0x93, 0xc7, 0x37, 0x9a, 0x0d, 0xd9, 0x7c, 0xe9, 0x3c, 0xb9, 0x19, 0x39, 0x9b,
	0x10, 0xbd, 0xbf, 0xc8, 0x39, 0x34, 0xe6, 0x73, 0x83, 0xec, 0xdd, 0x64, 0xed, 0x85, 0x81, 0xd6,
	0xd9, 0xbf, 0xdd, 0xa2, 0xf9, 0xc6, 0xc3, 0xa7, 0xef, 0xdc, 0x20, 0xd4, 0xa7, 0xe9, 0xd8, 0xf1,
	0x30, 0x76, 0xe3, 0xd0, 0x93, 0xa8, 0x70, 0xa2, 0xdd, 0x18, 0x3d, 0x57, 0x26, 0x9e, 0xbb, 0x54,
	0x74, 0x73, 0xc5, 0x71, 0xdd, 0x3c, 0xb2, 0xbd, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x88, 0x59,
	0x9d, 0xe5, 0x40, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoadBalancerAgentClient is the client API for LoadBalancerAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadBalancerAgentClient interface {
	Invoke(ctx context.Context, in *LoadBalancerRequest, opts ...grpc.CallOption) (*LoadBalancerResponse, error)
	// Prechecks whether the system is able to create specified load balancers (but does not actually create them).
	Precheck(ctx context.Context, in *LoadBalancerPrecheckRequest, opts ...grpc.CallOption) (*LoadBalancerPrecheckResponse, error)
}

type loadBalancerAgentClient struct {
	cc *grpc.ClientConn
}

func NewLoadBalancerAgentClient(cc *grpc.ClientConn) LoadBalancerAgentClient {
	return &loadBalancerAgentClient{cc}
}

func (c *loadBalancerAgentClient) Invoke(ctx context.Context, in *LoadBalancerRequest, opts ...grpc.CallOption) (*LoadBalancerResponse, error) {
	out := new(LoadBalancerResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.network.LoadBalancerAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancerAgentClient) Precheck(ctx context.Context, in *LoadBalancerPrecheckRequest, opts ...grpc.CallOption) (*LoadBalancerPrecheckResponse, error) {
	out := new(LoadBalancerPrecheckResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.network.LoadBalancerAgent/Precheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoadBalancerAgentServer is the server API for LoadBalancerAgent service.
type LoadBalancerAgentServer interface {
	Invoke(context.Context, *LoadBalancerRequest) (*LoadBalancerResponse, error)
	// Prechecks whether the system is able to create specified load balancers (but does not actually create them).
	Precheck(context.Context, *LoadBalancerPrecheckRequest) (*LoadBalancerPrecheckResponse, error)
}

// UnimplementedLoadBalancerAgentServer can be embedded to have forward compatible implementations.
type UnimplementedLoadBalancerAgentServer struct {
}

func (*UnimplementedLoadBalancerAgentServer) Invoke(ctx context.Context, req *LoadBalancerRequest) (*LoadBalancerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (*UnimplementedLoadBalancerAgentServer) Precheck(ctx context.Context, req *LoadBalancerPrecheckRequest) (*LoadBalancerPrecheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Precheck not implemented")
}

func RegisterLoadBalancerAgentServer(s *grpc.Server, srv LoadBalancerAgentServer) {
	s.RegisterService(&_LoadBalancerAgent_serviceDesc, srv)
}

func _LoadBalancerAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.network.LoadBalancerAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerAgentServer).Invoke(ctx, req.(*LoadBalancerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancerAgent_Precheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadBalancerPrecheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerAgentServer).Precheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.network.LoadBalancerAgent/Precheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerAgentServer).Precheck(ctx, req.(*LoadBalancerPrecheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoadBalancerAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.network.LoadBalancerAgent",
	HandlerType: (*LoadBalancerAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _LoadBalancerAgent_Invoke_Handler,
		},
		{
			MethodName: "Precheck",
			Handler:    _LoadBalancerAgent_Precheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_cloudagent_loadbalancer.proto",
}
