// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/ratelimit/v2/rls.proto

package v2

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Rate limit :ref:`configuration overview <config_rate_limit_service>`.
type RateLimitServiceConfig struct {
	// Specifies the gRPC service that hosts the rate limit service. The client
	// will connect to this cluster when it needs to make rate limit service
	// requests.
	GrpcService          *core.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RateLimitServiceConfig) Reset()         { *m = RateLimitServiceConfig{} }
func (m *RateLimitServiceConfig) String() string { return proto.CompactTextString(m) }
func (*RateLimitServiceConfig) ProtoMessage()    {}
func (*RateLimitServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3154ecf621be8917, []int{0}
}

func (m *RateLimitServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitServiceConfig.Unmarshal(m, b)
}
func (m *RateLimitServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitServiceConfig.Marshal(b, m, deterministic)
}
func (m *RateLimitServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitServiceConfig.Merge(m, src)
}
func (m *RateLimitServiceConfig) XXX_Size() int {
	return xxx_messageInfo_RateLimitServiceConfig.Size(m)
}
func (m *RateLimitServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitServiceConfig proto.InternalMessageInfo

func (m *RateLimitServiceConfig) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimitServiceConfig)(nil), "envoy.config.ratelimit.v2.RateLimitServiceConfig")
}

func init() {
	proto.RegisterFile("envoy/config/ratelimit/v2/rls.proto", fileDescriptor_3154ecf621be8917)
}

var fileDescriptor_3154ecf621be8917 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0x4a, 0x2c, 0x49, 0xcd, 0xc9, 0xcc,
	0xcd, 0x2c, 0xd1, 0x2f, 0x33, 0xd2, 0x2f, 0xca, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x04, 0x2b, 0xd2, 0x83, 0x28, 0xd2, 0x83, 0x2b, 0xd2, 0x2b, 0x33, 0x92, 0x52, 0x81, 0xe8,
	0x4f, 0x2c, 0xc8, 0x04, 0x69, 0x49, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0x2f, 0x2a, 0x48, 0x8e, 0x2f,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x85, 0x18, 0x20, 0x25, 0x5e, 0x96, 0x98, 0x93, 0x99, 0x92,
	0x58, 0x92, 0xaa, 0x0f, 0x63, 0x40, 0x24, 0x94, 0x8a, 0xb9, 0xc4, 0x82, 0x12, 0x4b, 0x52, 0x7d,
	0x40, 0xc6, 0x05, 0x43, 0xb4, 0x38, 0x83, 0x6d, 0x11, 0xf2, 0xe5, 0xe2, 0x41, 0x36, 0x48, 0x82,
	0x49, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x4e, 0x0f, 0xe2, 0x94, 0xc4, 0x82, 0x4c, 0xbd, 0x32, 0x23,
	0x3d, 0x90, 0x7d, 0x7a, 0xee, 0x45, 0x05, 0xc9, 0x50, 0xbd, 0x4e, 0x5c, 0xbb, 0x5e, 0x1e, 0x60,
	0x66, 0xed, 0x62, 0x64, 0x12, 0x60, 0x0c, 0xe2, 0x4e, 0x47, 0x48, 0x78, 0xb1, 0x70, 0x30, 0x0a,
	0x30, 0x79, 0xb1, 0x70, 0x30, 0x0b, 0xb0, 0x38, 0x59, 0x72, 0xa9, 0x67, 0xe6, 0x43, 0x0c, 0x2a,
	0x28, 0xca, 0xaf, 0xa8, 0xd4, 0xc3, 0xe9, 0x3d, 0x27, 0x8e, 0xa0, 0x9c, 0xe2, 0x00, 0x90, 0x4b,
	0x03, 0x18, 0xa3, 0x98, 0xca, 0x8c, 0x92, 0xd8, 0xc0, 0xce, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xa3, 0x4d, 0x98, 0xd1, 0x37, 0x01, 0x00, 0x00,
}
