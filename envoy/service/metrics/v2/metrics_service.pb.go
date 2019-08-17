// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/metrics/v2/metrics_service.proto

package v2

import (
	context "context"
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_go "github.com/prometheus/client_model/go"
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

type StreamMetricsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamMetricsResponse) Reset()         { *m = StreamMetricsResponse{} }
func (m *StreamMetricsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamMetricsResponse) ProtoMessage()    {}
func (*StreamMetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_587c4e7585395bd5, []int{0}
}

func (m *StreamMetricsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMetricsResponse.Unmarshal(m, b)
}
func (m *StreamMetricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMetricsResponse.Marshal(b, m, deterministic)
}
func (m *StreamMetricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMetricsResponse.Merge(m, src)
}
func (m *StreamMetricsResponse) XXX_Size() int {
	return xxx_messageInfo_StreamMetricsResponse.Size(m)
}
func (m *StreamMetricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMetricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMetricsResponse proto.InternalMessageInfo

type StreamMetricsMessage struct {
	// Identifier data effectively is a structured metadata. As a performance optimization this will
	// only be sent in the first message on the stream.
	Identifier *StreamMetricsMessage_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// A list of metric entries
	EnvoyMetrics         []*_go.MetricFamily `protobuf:"bytes,2,rep,name=envoy_metrics,json=envoyMetrics,proto3" json:"envoy_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *StreamMetricsMessage) Reset()         { *m = StreamMetricsMessage{} }
func (m *StreamMetricsMessage) String() string { return proto.CompactTextString(m) }
func (*StreamMetricsMessage) ProtoMessage()    {}
func (*StreamMetricsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_587c4e7585395bd5, []int{1}
}

func (m *StreamMetricsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMetricsMessage.Unmarshal(m, b)
}
func (m *StreamMetricsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMetricsMessage.Marshal(b, m, deterministic)
}
func (m *StreamMetricsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMetricsMessage.Merge(m, src)
}
func (m *StreamMetricsMessage) XXX_Size() int {
	return xxx_messageInfo_StreamMetricsMessage.Size(m)
}
func (m *StreamMetricsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMetricsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMetricsMessage proto.InternalMessageInfo

func (m *StreamMetricsMessage) GetIdentifier() *StreamMetricsMessage_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *StreamMetricsMessage) GetEnvoyMetrics() []*_go.MetricFamily {
	if m != nil {
		return m.EnvoyMetrics
	}
	return nil
}

type StreamMetricsMessage_Identifier struct {
	// The node sending metrics over the stream.
	Node                 *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StreamMetricsMessage_Identifier) Reset()         { *m = StreamMetricsMessage_Identifier{} }
func (m *StreamMetricsMessage_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamMetricsMessage_Identifier) ProtoMessage()    {}
func (*StreamMetricsMessage_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_587c4e7585395bd5, []int{1, 0}
}

func (m *StreamMetricsMessage_Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMetricsMessage_Identifier.Unmarshal(m, b)
}
func (m *StreamMetricsMessage_Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMetricsMessage_Identifier.Marshal(b, m, deterministic)
}
func (m *StreamMetricsMessage_Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMetricsMessage_Identifier.Merge(m, src)
}
func (m *StreamMetricsMessage_Identifier) XXX_Size() int {
	return xxx_messageInfo_StreamMetricsMessage_Identifier.Size(m)
}
func (m *StreamMetricsMessage_Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMetricsMessage_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMetricsMessage_Identifier proto.InternalMessageInfo

func (m *StreamMetricsMessage_Identifier) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamMetricsResponse)(nil), "envoy.service.metrics.v2.StreamMetricsResponse")
	proto.RegisterType((*StreamMetricsMessage)(nil), "envoy.service.metrics.v2.StreamMetricsMessage")
	proto.RegisterType((*StreamMetricsMessage_Identifier)(nil), "envoy.service.metrics.v2.StreamMetricsMessage.Identifier")
}

func init() {
	proto.RegisterFile("envoy/service/metrics/v2/metrics_service.proto", fileDescriptor_587c4e7585395bd5)
}

var fileDescriptor_587c4e7585395bd5 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xff, 0x9b, 0xfe, 0xf5, 0x30, 0xb5, 0x22, 0x51, 0x69, 0x09, 0x1e, 0x4a, 0x0f, 0xd2,
	0xd3, 0x06, 0xe2, 0x41, 0xbc, 0x16, 0xac, 0x78, 0xa8, 0x94, 0xf4, 0xa4, 0x97, 0xb2, 0x4d, 0x46,
	0x5d, 0x68, 0xb2, 0x61, 0x77, 0x5d, 0xec, 0xd1, 0x8b, 0x88, 0x8f, 0xe4, 0xc9, 0xd7, 0xf1, 0x2d,
	0x24, 0xd9, 0x4d, 0x35, 0x60, 0x41, 0x6f, 0x93, 0xcc, 0xf7, 0xfd, 0xe6, 0x9b, 0x1d, 0xa0, 0x98,
	0x1b, 0xb1, 0x0a, 0x15, 0x4a, 0xc3, 0x13, 0x0c, 0x33, 0xd4, 0x92, 0x27, 0x2a, 0x34, 0x51, 0x5d,
	0xce, 0x5d, 0x8b, 0x16, 0x52, 0x68, 0xe1, 0xf7, 0x2a, 0x3d, 0xad, 0x7f, 0x3a, 0x11, 0x35, 0x51,
	0x70, 0x64, 0x49, 0xac, 0xe0, 0xa5, 0x3b, 0x11, 0x12, 0xc3, 0x05, 0x53, 0xce, 0x17, 0x74, 0x6a,
	0xa5, 0xfd, 0xec, 0x1a, 0xb6, 0xe4, 0x29, 0xd3, 0x18, 0xd6, 0x85, 0x6d, 0x0c, 0xba, 0x70, 0x38,
	0xd3, 0x12, 0x59, 0x36, 0xb1, 0xfa, 0x18, 0x55, 0x21, 0x72, 0x85, 0x83, 0x27, 0x0f, 0x0e, 0x1a,
	0x9d, 0x09, 0x2a, 0xc5, 0xee, 0xd0, 0xbf, 0x06, 0xe0, 0x29, 0xe6, 0x9a, 0xdf, 0x72, 0x94, 0x3d,
	0xd2, 0x27, 0xc3, 0x76, 0x74, 0x46, 0x37, 0xc5, 0xa4, 0x3f, 0x31, 0xe8, 0xe5, 0x1a, 0x10, 0x7f,
	0x83, 0xf9, 0x17, 0xd0, 0xa9, 0x38, 0x73, 0xe7, 0xef, 0x79, 0xfd, 0xd6, 0xb0, 0x1d, 0x0d, 0x28,
	0x17, 0x65, 0xdc, 0x0c, 0xf5, 0x3d, 0x3e, 0x28, 0x9a, 0x2c, 0x39, 0xe6, 0x9a, 0x5a, 0xe6, 0x98,
	0x65, 0x7c, 0xb9, 0x8a, 0x77, 0x2a, 0xa3, 0x1b, 0x13, 0x9c, 0x03, 0x7c, 0x8d, 0xf0, 0x4f, 0xe1,
	0x7f, 0x2e, 0x52, 0x74, 0x59, 0xbb, 0x2e, 0x2b, 0x2b, 0x78, 0x99, 0xaf, 0x7c, 0x38, 0x7a, 0x25,
	0x52, 0x1c, 0xc1, 0xdb, 0xc7, 0x7b, 0x6b, 0xeb, 0x95, 0x78, 0x7b, 0x24, 0xae, 0x0c, 0xd1, 0x33,
	0x81, 0x5d, 0x87, 0x9c, 0xd9, 0xcd, 0x7c, 0x0d, 0x9d, 0xc6, 0x46, 0x3e, 0xfd, 0xdb, 0xea, 0x41,
	0xf8, 0x4b, 0xfd, 0xfa, 0x10, 0xff, 0x86, 0x64, 0x34, 0x86, 0x63, 0x2e, 0xac, 0xb1, 0x90, 0xe2,
	0x71, 0xb5, 0x91, 0x31, 0xda, 0x6f, 0xe6, 0x9d, 0x96, 0x47, 0x9e, 0x92, 0x1b, 0xcf, 0x44, 0x2f,
	0x84, 0x2c, 0xb6, 0xab, 0xa3, 0x9f, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xa2, 0x17, 0x60,
	0x86, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsServiceClient interface {
	// Envoy will connect and send StreamMetricsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure.
	StreamMetrics(ctx context.Context, opts ...grpc.CallOption) (MetricsService_StreamMetricsClient, error)
}

type metricsServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetricsServiceClient(cc *grpc.ClientConn) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) StreamMetrics(ctx context.Context, opts ...grpc.CallOption) (MetricsService_StreamMetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MetricsService_serviceDesc.Streams[0], "/envoy.service.metrics.v2.MetricsService/StreamMetrics", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsServiceStreamMetricsClient{stream}
	return x, nil
}

type MetricsService_StreamMetricsClient interface {
	Send(*StreamMetricsMessage) error
	CloseAndRecv() (*StreamMetricsResponse, error)
	grpc.ClientStream
}

type metricsServiceStreamMetricsClient struct {
	grpc.ClientStream
}

func (x *metricsServiceStreamMetricsClient) Send(m *StreamMetricsMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsServiceStreamMetricsClient) CloseAndRecv() (*StreamMetricsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamMetricsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetricsServiceServer is the server API for MetricsService service.
type MetricsServiceServer interface {
	// Envoy will connect and send StreamMetricsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure.
	StreamMetrics(MetricsService_StreamMetricsServer) error
}

// UnimplementedMetricsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMetricsServiceServer struct {
}

func (*UnimplementedMetricsServiceServer) StreamMetrics(srv MetricsService_StreamMetricsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMetrics not implemented")
}

func RegisterMetricsServiceServer(s *grpc.Server, srv MetricsServiceServer) {
	s.RegisterService(&_MetricsService_serviceDesc, srv)
}

func _MetricsService_StreamMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServiceServer).StreamMetrics(&metricsServiceStreamMetricsServer{stream})
}

type MetricsService_StreamMetricsServer interface {
	SendAndClose(*StreamMetricsResponse) error
	Recv() (*StreamMetricsMessage, error)
	grpc.ServerStream
}

type metricsServiceStreamMetricsServer struct {
	grpc.ServerStream
}

func (x *metricsServiceStreamMetricsServer) SendAndClose(m *StreamMetricsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsServiceStreamMetricsServer) Recv() (*StreamMetricsMessage, error) {
	m := new(StreamMetricsMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.metrics.v2.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMetrics",
			Handler:       _MetricsService_StreamMetrics_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/metrics/v2/metrics_service.proto",
}
