// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/thrift_proxy/v2alpha1/thrift_proxy.proto

package v2

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// Thrift transport types supported by Envoy.
type TransportType int32

const (
	// For downstream connections, the Thrift proxy will attempt to determine which transport to use.
	// For upstream connections, the Thrift proxy will use same transport as the downstream
	// connection.
	TransportType_AUTO_TRANSPORT TransportType = 0
	// The Thrift proxy will use the Thrift framed transport.
	TransportType_FRAMED TransportType = 1
	// The Thrift proxy will use the Thrift unframed transport.
	TransportType_UNFRAMED TransportType = 2
	// The Thrift proxy will assume the client is using the Thrift header transport.
	TransportType_HEADER TransportType = 3
)

var TransportType_name = map[int32]string{
	0: "AUTO_TRANSPORT",
	1: "FRAMED",
	2: "UNFRAMED",
	3: "HEADER",
}

var TransportType_value = map[string]int32{
	"AUTO_TRANSPORT": 0,
	"FRAMED":         1,
	"UNFRAMED":       2,
	"HEADER":         3,
}

func (x TransportType) String() string {
	return proto.EnumName(TransportType_name, int32(x))
}

func (TransportType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{0}
}

// Thrift Protocol types supported by Envoy.
type ProtocolType int32

const (
	// For downstream connections, the Thrift proxy will attempt to determine which protocol to use.
	// Note that the older, non-strict (or lax) binary protocol is not included in automatic protocol
	// detection. For upstream connections, the Thrift proxy will use the same protocol as the
	// downstream connection.
	ProtocolType_AUTO_PROTOCOL ProtocolType = 0
	// The Thrift proxy will use the Thrift binary protocol.
	ProtocolType_BINARY ProtocolType = 1
	// The Thrift proxy will use Thrift non-strict binary protocol.
	ProtocolType_LAX_BINARY ProtocolType = 2
	// The Thrift proxy will use the Thrift compact protocol.
	ProtocolType_COMPACT ProtocolType = 3
	// The Thrift proxy will use the Thrift "Twitter" protocol implemented by the finagle library.
	ProtocolType_TWITTER ProtocolType = 4
)

var ProtocolType_name = map[int32]string{
	0: "AUTO_PROTOCOL",
	1: "BINARY",
	2: "LAX_BINARY",
	3: "COMPACT",
	4: "TWITTER",
}

var ProtocolType_value = map[string]int32{
	"AUTO_PROTOCOL": 0,
	"BINARY":        1,
	"LAX_BINARY":    2,
	"COMPACT":       3,
	"TWITTER":       4,
}

func (x ProtocolType) String() string {
	return proto.EnumName(ProtocolType_name, int32(x))
}

func (ProtocolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{1}
}

// [#comment:next free field: 6]
type ThriftProxy struct {
	// Supplies the type of transport that the Thrift proxy should use. Defaults to
	// :ref:`AUTO_TRANSPORT<envoy_api_enum_value_config.filter.network.thrift_proxy.v2alpha1.TransportType.AUTO_TRANSPORT>`.
	Transport TransportType `protobuf:"varint,2,opt,name=transport,proto3,enum=envoy.config.filter.network.thrift_proxy.v2alpha1.TransportType" json:"transport,omitempty"`
	// Supplies the type of protocol that the Thrift proxy should use. Defaults to
	// :ref:`AUTO_PROTOCOL<envoy_api_enum_value_config.filter.network.thrift_proxy.v2alpha1.ProtocolType.AUTO_PROTOCOL>`.
	Protocol ProtocolType `protobuf:"varint,3,opt,name=protocol,proto3,enum=envoy.config.filter.network.thrift_proxy.v2alpha1.ProtocolType" json:"protocol,omitempty"`
	// The human readable prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The route table for the connection manager is static and is specified in this property.
	RouteConfig *RouteConfiguration `protobuf:"bytes,4,opt,name=route_config,json=routeConfig,proto3" json:"route_config,omitempty"`
	// A list of individual Thrift filters that make up the filter chain for requests made to the
	// Thrift proxy. Order matters as the filters are processed sequentially. For backwards
	// compatibility, if no thrift_filters are specified, a default Thrift router filter
	// (`envoy.filters.thrift.router`) is used.
	ThriftFilters        []*ThriftFilter `protobuf:"bytes,5,rep,name=thrift_filters,json=thriftFilters,proto3" json:"thrift_filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ThriftProxy) Reset()         { *m = ThriftProxy{} }
func (m *ThriftProxy) String() string { return proto.CompactTextString(m) }
func (*ThriftProxy) ProtoMessage()    {}
func (*ThriftProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{0}
}

func (m *ThriftProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftProxy.Unmarshal(m, b)
}
func (m *ThriftProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftProxy.Marshal(b, m, deterministic)
}
func (m *ThriftProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftProxy.Merge(m, src)
}
func (m *ThriftProxy) XXX_Size() int {
	return xxx_messageInfo_ThriftProxy.Size(m)
}
func (m *ThriftProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftProxy.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftProxy proto.InternalMessageInfo

func (m *ThriftProxy) GetTransport() TransportType {
	if m != nil {
		return m.Transport
	}
	return TransportType_AUTO_TRANSPORT
}

func (m *ThriftProxy) GetProtocol() ProtocolType {
	if m != nil {
		return m.Protocol
	}
	return ProtocolType_AUTO_PROTOCOL
}

func (m *ThriftProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ThriftProxy) GetRouteConfig() *RouteConfiguration {
	if m != nil {
		return m.RouteConfig
	}
	return nil
}

func (m *ThriftProxy) GetThriftFilters() []*ThriftFilter {
	if m != nil {
		return m.ThriftFilters
	}
	return nil
}

// ThriftFilter configures a Thrift filter.
// [#comment:next free field: 3]
type ThriftFilter struct {
	// The name of the filter to instantiate. The name must match a supported
	// filter. The built-in filters are:
	//
	// [#comment:TODO(zuercher): Auto generate the following list]
	// * :ref:`envoy.filters.thrift.router <config_thrift_filters_router>`
	// * :ref:`envoy.filters.thrift.rate_limit <config_thrift_filters_rate_limit>`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Filter specific configuration which depends on the filter being instantiated. See the supported
	// filters for further documentation.
	//
	// Types that are valid to be assigned to ConfigType:
	//	*ThriftFilter_Config
	//	*ThriftFilter_TypedConfig
	ConfigType           isThriftFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ThriftFilter) Reset()         { *m = ThriftFilter{} }
func (m *ThriftFilter) String() string { return proto.CompactTextString(m) }
func (*ThriftFilter) ProtoMessage()    {}
func (*ThriftFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{1}
}

func (m *ThriftFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftFilter.Unmarshal(m, b)
}
func (m *ThriftFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftFilter.Marshal(b, m, deterministic)
}
func (m *ThriftFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftFilter.Merge(m, src)
}
func (m *ThriftFilter) XXX_Size() int {
	return xxx_messageInfo_ThriftFilter.Size(m)
}
func (m *ThriftFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftFilter proto.InternalMessageInfo

func (m *ThriftFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isThriftFilter_ConfigType interface {
	isThriftFilter_ConfigType()
}

type ThriftFilter_Config struct {
	Config *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type ThriftFilter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ThriftFilter_Config) isThriftFilter_ConfigType() {}

func (*ThriftFilter_TypedConfig) isThriftFilter_ConfigType() {}

func (m *ThriftFilter) GetConfigType() isThriftFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *ThriftFilter) GetConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*ThriftFilter_Config); ok {
		return x.Config
	}
	return nil
}

func (m *ThriftFilter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*ThriftFilter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ThriftFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ThriftFilter_Config)(nil),
		(*ThriftFilter_TypedConfig)(nil),
	}
}

// ThriftProtocolOptions specifies Thrift upstream protocol options. This object is used in
// in :ref:`extension_protocol_options<envoy_api_field_Cluster.extension_protocol_options>`, keyed
// by the name `envoy.filters.network.thrift_proxy`.
// [#comment:next free field: 3]
type ThriftProtocolOptions struct {
	// Supplies the type of transport that the Thrift proxy should use for upstream connections.
	// Selecting
	// :ref:`AUTO_TRANSPORT<envoy_api_enum_value_config.filter.network.thrift_proxy.v2alpha1.TransportType.AUTO_TRANSPORT>`,
	// which is the default, causes the proxy to use the same transport as the downstream connection.
	Transport TransportType `protobuf:"varint,1,opt,name=transport,proto3,enum=envoy.config.filter.network.thrift_proxy.v2alpha1.TransportType" json:"transport,omitempty"`
	// Supplies the type of protocol that the Thrift proxy should use for upstream connections.
	// Selecting
	// :ref:`AUTO_PROTOCOL<envoy_api_enum_value_config.filter.network.thrift_proxy.v2alpha1.ProtocolType.AUTO_PROTOCOL>`,
	// which is the default, causes the proxy to use the same protocol as the downstream connection.
	Protocol             ProtocolType `protobuf:"varint,2,opt,name=protocol,proto3,enum=envoy.config.filter.network.thrift_proxy.v2alpha1.ProtocolType" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ThriftProtocolOptions) Reset()         { *m = ThriftProtocolOptions{} }
func (m *ThriftProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*ThriftProtocolOptions) ProtoMessage()    {}
func (*ThriftProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{2}
}

func (m *ThriftProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftProtocolOptions.Unmarshal(m, b)
}
func (m *ThriftProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftProtocolOptions.Marshal(b, m, deterministic)
}
func (m *ThriftProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftProtocolOptions.Merge(m, src)
}
func (m *ThriftProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_ThriftProtocolOptions.Size(m)
}
func (m *ThriftProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftProtocolOptions proto.InternalMessageInfo

func (m *ThriftProtocolOptions) GetTransport() TransportType {
	if m != nil {
		return m.Transport
	}
	return TransportType_AUTO_TRANSPORT
}

func (m *ThriftProtocolOptions) GetProtocol() ProtocolType {
	if m != nil {
		return m.Protocol
	}
	return ProtocolType_AUTO_PROTOCOL
}

func init() {
	proto.RegisterEnum("envoy.config.filter.network.thrift_proxy.v2alpha1.TransportType", TransportType_name, TransportType_value)
	proto.RegisterEnum("envoy.config.filter.network.thrift_proxy.v2alpha1.ProtocolType", ProtocolType_name, ProtocolType_value)
	proto.RegisterType((*ThriftProxy)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.ThriftProxy")
	proto.RegisterType((*ThriftFilter)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.ThriftFilter")
	proto.RegisterType((*ThriftProtocolOptions)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.ThriftProtocolOptions")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/thrift_proxy/v2alpha1/thrift_proxy.proto", fileDescriptor_e8fab7646d88fc90)
}

var fileDescriptor_e8fab7646d88fc90 = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xcd, 0xd8, 0x69, 0x69, 0xaf, 0x93, 0xc8, 0x8c, 0x8a, 0x1a, 0x22, 0x40, 0x51, 0x57, 0x51,
	0x17, 0xb6, 0x1a, 0x56, 0x2c, 0x50, 0xb1, 0xd3, 0x54, 0xad, 0xd4, 0xc6, 0xd6, 0xd4, 0x15, 0x14,
	0x09, 0x45, 0x6e, 0x6a, 0x3b, 0x86, 0xe0, 0xb1, 0xc6, 0x93, 0xd0, 0x6c, 0x59, 0xf1, 0x1f, 0xec,
	0xf8, 0x04, 0x56, 0x7c, 0x0a, 0x5b, 0x24, 0x3e, 0x02, 0x79, 0xc6, 0x69, 0x13, 0xb2, 0x4a, 0x24,
	0xd8, 0xf9, 0xbe, 0xce, 0x3d, 0xf7, 0xcc, 0x91, 0xe1, 0x28, 0x48, 0x26, 0x74, 0x6a, 0x0e, 0x68,
	0x12, 0xc6, 0x91, 0x19, 0xc6, 0x23, 0x1e, 0x30, 0x33, 0x09, 0xf8, 0x27, 0xca, 0x3e, 0x98, 0x7c,
	0xc8, 0xe2, 0x90, 0xf7, 0x53, 0x46, 0x6f, 0xa7, 0xe6, 0xa4, 0xed, 0x8f, 0xd2, 0xa1, 0x7f, 0xb0,
	0x90, 0x35, 0x52, 0x46, 0x39, 0xc5, 0x07, 0x02, 0xc5, 0x90, 0x28, 0x86, 0x44, 0x31, 0x0a, 0x14,
	0x63, 0xa1, 0x7f, 0x86, 0xd2, 0x78, 0xb9, 0xfa, 0x62, 0x46, 0xc7, 0x3c, 0x90, 0x1b, 0x1b, 0x8f,
	0x23, 0x4a, 0xa3, 0x51, 0x60, 0x8a, 0xe8, 0x7a, 0x1c, 0x9a, 0x7e, 0x52, 0x90, 0x69, 0x3c, 0xf9,
	0xbb, 0x94, 0x71, 0x36, 0x1e, 0xf0, 0xa2, 0xba, 0x3b, 0xf1, 0x47, 0xf1, 0x8d, 0xcf, 0x03, 0x73,
	0xf6, 0x51, 0x14, 0x76, 0x22, 0x1a, 0x51, 0xf1, 0x69, 0xe6, 0x5f, 0x32, 0xbb, 0xf7, 0x53, 0x05,
	0xcd, 0x13, 0x6c, 0xdc, 0x9c, 0x0c, 0x7e, 0x0f, 0xdb, 0x9c, 0xf9, 0x49, 0x96, 0x52, 0xc6, 0xeb,
	0x4a, 0x13, 0xb5, 0x6a, 0xed, 0x57, 0xc6, 0xca, 0xd7, 0x1b, 0xde, 0x0c, 0xc3, 0x9b, 0xa6, 0x81,
	0x0d, 0xdf, 0x7f, 0xfd, 0x50, 0x37, 0x3e, 0x23, 0x45, 0x47, 0xe4, 0x1e, 0x1e, 0x47, 0xb0, 0x25,
	0x48, 0x0c, 0xe8, 0xa8, 0xae, 0x8a, 0x55, 0x87, 0x6b, 0xac, 0x72, 0x0b, 0x88, 0xa5, 0x4d, 0x77,
	0xe0, 0x78, 0x1f, 0xb4, 0x8c, 0xfb, 0xf9, 0x64, 0x10, 0xc6, 0xb7, 0x75, 0xd4, 0x44, 0xad, 0x6d,
	0x7b, 0x3b, 0x6f, 0x2d, 0x33, 0xa5, 0x89, 0x08, 0xe4, 0x55, 0x57, 0x14, 0xf1, 0x10, 0x2a, 0xe2,
	0x1d, 0xfa, 0x92, 0x43, 0xbd, 0xdc, 0x44, 0x2d, 0xad, 0xdd, 0x5d, 0x83, 0x18, 0xc9, 0x61, 0x3a,
	0x62, 0x60, 0xcc, 0x7c, 0x1e, 0xd3, 0x84, 0x68, 0xec, 0x3e, 0x87, 0x43, 0xa8, 0x15, 0x83, 0x12,
	0x2e, 0xab, 0x6f, 0x34, 0xd5, 0x96, 0xb6, 0x96, 0x08, 0xf2, 0x09, 0x8f, 0x45, 0x2b, 0xa9, 0xf2,
	0xb9, 0x28, 0xdb, 0xfb, 0x86, 0xa0, 0x32, 0x5f, 0xc7, 0x4f, 0xa1, 0x9c, 0xf8, 0x1f, 0x83, 0x65,
	0x1d, 0x44, 0x1a, 0x1f, 0xc0, 0x66, 0x71, 0xbb, 0x22, 0x6e, 0xdf, 0x35, 0xa4, 0xe1, 0x8c, 0x99,
	0xe1, 0x8c, 0x0b, 0x61, 0xb8, 0x93, 0x12, 0x29, 0x1a, 0xf1, 0x0b, 0xa8, 0xf0, 0x69, 0x1a, 0xdc,
	0xcc, 0x44, 0x53, 0xc5, 0xe0, 0xce, 0xd2, 0xa0, 0x95, 0x4c, 0x4f, 0x4a, 0x44, 0x13, 0xbd, 0x52,
	0x05, 0xbb, 0x0a, 0x9a, 0x1c, 0xea, 0xe7, 0xd9, 0xbd, 0xdf, 0x08, 0x1e, 0xdd, 0xf9, 0x51, 0xbc,
	0x9e, 0x93, 0xe6, 0xd2, 0x65, 0x8b, 0xce, 0x44, 0xff, 0xcf, 0x99, 0xca, 0x3f, 0x74, 0xe6, 0xbe,
	0x03, 0xd5, 0x05, 0x42, 0x18, 0x43, 0xcd, 0xba, 0xf4, 0x9c, 0xbe, 0x47, 0xac, 0xde, 0x85, 0xeb,
	0x10, 0x4f, 0x2f, 0x61, 0x80, 0xcd, 0x63, 0x62, 0x9d, 0x77, 0x8f, 0x74, 0x84, 0x2b, 0xb0, 0x75,
	0xd9, 0x2b, 0x22, 0x25, 0xaf, 0x9c, 0x74, 0xad, 0xa3, 0x2e, 0xd1, 0xd5, 0x46, 0xf9, 0xcb, 0xd7,
	0x67, 0xa5, 0xfd, 0x77, 0x50, 0x99, 0x5f, 0x8b, 0x1f, 0x42, 0x55, 0xe0, 0xb9, 0xc4, 0xf1, 0x9c,
	0x8e, 0x73, 0x26, 0xe1, 0xec, 0xd3, 0x9e, 0x45, 0xae, 0x74, 0x84, 0x6b, 0x00, 0x67, 0xd6, 0x9b,
	0x7e, 0x11, 0x2b, 0x58, 0x83, 0x07, 0x1d, 0xe7, 0xdc, 0xb5, 0x3a, 0x9e, 0xae, 0xe6, 0x81, 0xf7,
	0xfa, 0xd4, 0xf3, 0xba, 0x44, 0x2f, 0x4b, 0x78, 0xfb, 0x0a, 0x0e, 0x63, 0x2a, 0xa5, 0x90, 0xc7,
	0xae, 0xac, 0x8a, 0xad, 0xcf, 0xfd, 0x6e, 0x04, 0x55, 0x17, 0xbd, 0x55, 0x26, 0xed, 0xeb, 0x4d,
	0x21, 0xca, 0xf3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x8e, 0x5b, 0xcb, 0xb2, 0x05, 0x00,
	0x00,
}
