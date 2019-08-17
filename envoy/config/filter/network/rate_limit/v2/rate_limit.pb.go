// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/rate_limit/v2/rate_limit.proto

package v2

import (
	fmt "fmt"
	ratelimit "github.com/envoyproxy/go-control-plane/envoy/api/v2/ratelimit"
	v2 "github.com/envoyproxy/go-control-plane/envoy/config/ratelimit/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type RateLimit struct {
	// The prefix to use when emitting :ref:`statistics <config_network_filters_rate_limit_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The rate limit domain to use in the rate limit service request.
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// The rate limit descriptor list to use in the rate limit service request.
	Descriptors []*ratelimit.RateLimitDescriptor `protobuf:"bytes,3,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *duration.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// The filter's behaviour in case the rate limiting service does
	// not respond back. When it is set to true, Envoy will not allow traffic in case of
	// communication failure between rate limiting service and the proxy.
	// Defaults to false.
	FailureModeDeny bool `protobuf:"varint,5,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	// Configuration for an external rate limit service provider. If not
	// specified, any calls to the rate limit service will immediately return
	// success.
	RateLimitService     *v2.RateLimitServiceConfig `protobuf:"bytes,6,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_34e9a222968daa71, []int{0}
}

func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimit.Unmarshal(m, b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimit.Size(m)
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetDescriptors() []*ratelimit.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimit) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func (m *RateLimit) GetRateLimitService() *v2.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.config.filter.network.rate_limit.v2.RateLimit")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/rate_limit/v2/rate_limit.proto", fileDescriptor_34e9a222968daa71)
}

var fileDescriptor_34e9a222968daa71 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xbf, 0x8e, 0xd3, 0x30,
	0x18, 0x97, 0xd3, 0x5e, 0xb9, 0xba, 0x12, 0x1c, 0x16, 0x12, 0xe1, 0x06, 0x08, 0x20, 0xa1, 0x72,
	0x48, 0xb6, 0x08, 0x03, 0x82, 0x31, 0x74, 0x04, 0x51, 0x85, 0x09, 0x96, 0xc8, 0xd7, 0x7c, 0x89,
	0xac, 0x4b, 0xf3, 0x45, 0x8e, 0x6b, 0xae, 0x8f, 0x00, 0x23, 0x13, 0xcf, 0xc2, 0xc4, 0x9b, 0x30,
	0xf3, 0x16, 0x28, 0xb1, 0xdb, 0xf4, 0x3a, 0xb1, 0x7d, 0xce, 0xef, 0xcf, 0xe7, 0xdf, 0x2f, 0xa6,
	0x6f, 0xa1, 0xb6, 0xb8, 0x15, 0x2b, 0xac, 0x0b, 0x55, 0x8a, 0x42, 0x55, 0x06, 0xb4, 0xa8, 0xc1,
	0x7c, 0x45, 0x7d, 0x25, 0xb4, 0x34, 0x90, 0x55, 0x6a, 0xad, 0x8c, 0xb0, 0xf1, 0xc1, 0x89, 0x37,
	0x1a, 0x0d, 0xb2, 0xe7, 0xbd, 0x96, 0x3b, 0x2d, 0x77, 0x5a, 0xee, 0xb5, 0xfc, 0x80, 0x6d, 0xe3,
	0xf3, 0x67, 0x6e, 0x8d, 0x6c, 0xd4, 0xce, 0xc9, 0xd9, 0xee, 0x27, 0x67, 0x79, 0xfe, 0xf4, 0xc6,
	0x75, 0x06, 0x5e, 0x27, 0xaa, 0x5a, 0x4f, 0x7a, 0x58, 0x22, 0x96, 0x15, 0x88, 0xfe, 0x74, 0xb9,
	0x29, 0x44, 0xbe, 0xd1, 0xd2, 0x28, 0xac, 0x3d, 0x7e, 0xdf, 0xca, 0x4a, 0xe5, 0xd2, 0x80, 0xd8,
	0x0d, 0x1e, 0xb8, 0x57, 0x62, 0x89, 0xfd, 0x28, 0xba, 0xc9, 0x7d, 0x7d, 0xf2, 0x6d, 0x44, 0xa7,
	0xa9, 0x34, 0xf0, 0xbe, 0xdb, 0xc4, 0x2e, 0xe8, 0xac, 0x35, 0xd2, 0x64, 0x8d, 0x86, 0x42, 0x5d,
	0x87, 0x24, 0x22, 0xf3, 0x69, 0x32, 0xfd, 0xf5, 0xf7, 0xf7, 0x68, 0xac, 0x83, 0x88, 0xa4, 0xb4,
	0x43, 0x97, 0x3d, 0xc8, 0x1e, 0xd3, 0x49, 0x8e, 0x6b, 0xa9, 0xea, 0x30, 0x38, 0xa6, 0x79, 0x80,
	0x7d, 0xa6, 0xb3, 0x1c, 0xda, 0x95, 0x56, 0x8d, 0x41, 0xdd, 0x86, 0xa3, 0x68, 0x34, 0x9f, 0xc5,
	0x2f, 0xb8, 0x6b, 0x4e, 0x36, 0x8a, 0xdb, 0x98, 0x0f, 0x25, 0xec, 0xaf, 0xb1, 0xd8, 0x6b, 0x12,
	0xda, 0x99, 0x9e, 0xfc, 0x20, 0xc1, 0x29, 0x49, 0x0f, 0xbd, 0xd8, 0x1b, 0x7a, 0xcb, 0xa8, 0x35,
	0xe0, 0xc6, 0x84, 0xe3, 0x88, 0xcc, 0x67, 0xf1, 0x03, 0xee, 0x8a, 0xe1, 0xbb, 0x62, 0xf8, 0xc2,
	0x17, 0x93, 0x8c, 0x7f, 0xfe, 0x79, 0x44, 0xd2, 0x1d, 0x9f, 0x5d, 0xd0, 0xbb, 0x85, 0x54, 0xd5,
	0x46, 0x43, 0xb6, 0xc6, 0x1c, 0xb2, 0x1c, 0xea, 0x6d, 0x78, 0x12, 0x91, 0xf9, 0x69, 0x7a, 0xc7,
	0x03, 0x1f, 0x30, 0x87, 0x05, 0xd4, 0x5b, 0x76, 0x45, 0xd9, 0xf0, 0x2f, 0xb3, 0x16, 0xb4, 0x55,
	0x2b, 0x08, 0x27, 0xfd, 0xc6, 0x97, 0xfc, 0xc6, 0x13, 0x18, 0x82, 0xd8, 0x78, 0xc8, 0xf2, 0xc9,
	0x49, 0xde, 0xf5, 0x1c, 0x1f, 0xe7, 0x3b, 0x09, 0xce, 0x48, 0x7a, 0xa6, 0x8f, 0x38, 0xc9, 0x47,
	0xfa, 0x5a, 0xa1, 0x33, 0x6d, 0x34, 0x5e, 0x6f, 0xf9, 0x7f, 0x3f, 0xb1, 0xe4, 0xf6, 0x7e, 0xe1,
	0xb2, 0x8b, 0xbf, 0x24, 0x5f, 0x02, 0x1b, 0x5f, 0x4e, 0xfa, 0x2e, 0x5e, 0xfd, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0x7d, 0x95, 0x59, 0xea, 0xe8, 0x02, 0x00, 0x00,
}
