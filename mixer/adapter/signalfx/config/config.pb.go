// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/signalfx/config/config.proto

// The `signalfx` adapter collects Istio metrics and trace spans and sends them
// to [SignalFx](https://signalfx.com).
//
// This adapter supports the [metric template](https://istio.io/docs/reference/config/policy-and-telemetry/templates/metric/)
// and the [tracespan template](https://istio.io/docs/reference/config/policy-and-telemetry/templates/tracespan/).
//
// If sending trace spans, this adapter can make use of certain conventions in
// the tracespan format that is configured to send to this adapter.  Here is an
// example tracespan spec that will work well:
//
// ```yaml
// apiVersion: config.istio.io/v1alpha2
// kind: instance
// metadata:
//   name: signalfx
// spec:
//   compiledTemplate: tracespan
//   params:
//     traceId: request.headers["x-b3-traceid"] | ""
//     spanId: request.headers["x-b3-spanid"] | ""
//     parentSpanId: request.headers["x-b3-parentspanid"] | ""
//     # If the path contains query parameters, they will be split off and put into
//     # tags such that the span name sent to SignalFx will consist only of the path
//     # itself.
//     spanName: request.path | "/"
//     startTime: request.time
//     endTime: response.time
//     # If this is >=500, the span will get an 'error' tag
//     httpStatusCode: response.code | 0
//     clientSpan: context.reporter.kind == "outbound"
//     # Span tags below that do not have comments are useful but optional and will
//     # be passed to SignalFx unmodified. The tags that have comments are interpreted
//     # in a special manner, but are still optional.
//     spanTags:
//       # This is used to determine whether the span pertains to the client or
//       # server side of the request.
//       context.reporter.local: context.reporter.local
//       # This gets put into the remoteEndpoint.ipv4 field
//       destination.ip: destination.ip | ip("0.0.0.0")
//       # This gets flattened out to individual tags of the form
//       # 'destination.labels.<key>: <value>'.
//       destination.labels: destination.labels
//       # This gets put into the remoteEndpoint.name field
//       destination.name: destination.name | "unknown"
//       destination.namespace: destination.namespace | "unknown"
//       request.host: request.host | ""
//       request.method: request.method | ""
//       request.path: request.path | ""
//       request.size: request.size | 0
//       request.useragent: request.useragent | ""
//       response.size: response.size | 0
//       # This gets put into the localEndpoint.name field
//       source.name: source.name | "unknown"
//       # This gets put into the localEndpoint.ipv4 field
//       source.ip: source.ip | ip("0.0.0.0")
//       source.namespace: source.namespace | "unknown"
//       # This gets flattened out to individual tags of the form
//       # 'source.labels.<key>: <value>'.
//       source.labels: source.labels
//       source.version: source.labels["version"] | "unknown"
//  ```

package config

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Describes what kind of metric this is.
type Params_MetricConfig_Type int32

const (
	// None is the default and is invalid
	NONE Params_MetricConfig_Type = 0
	// Values with the same set of dimensions will be added together
	// as a continuously incrementing value.
	COUNTER Params_MetricConfig_Type = 1
	// A histogram distribution.  This will result in several metrics
	// emitted for each unique set of dimensions.
	HISTOGRAM Params_MetricConfig_Type = 2
)

var Params_MetricConfig_Type_name = map[int32]string{
	0: "NONE",
	1: "COUNTER",
	2: "HISTOGRAM",
}

var Params_MetricConfig_Type_value = map[string]int32{
	"NONE":      0,
	"COUNTER":   1,
	"HISTOGRAM": 2,
}

func (Params_MetricConfig_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_41883b00dc025265, []int{0, 0, 0}
}

// Configuration format for the `signalfx` adapter.
type Params struct {
	// Required. The set of metrics to send to SignalFx. If an Istio metric is
	// configured to be sent to this adapter, it must have a corresponding
	// description here.
	Metrics []*Params_MetricConfig `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	// Optional. The URL of the SignalFx ingest server to use.  Will default to
	// the global ingest server if not specified.
	IngestUrl string `protobuf:"bytes,2,opt,name=ingest_url,json=ingestUrl,proto3" json:"ingest_url,omitempty"`
	// Required. The access token for the SignalFx organization that should
	// receive the metrics.
	AccessToken string `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// Optional. Specifies how frequently to send metrics to SignalFx.  Metrics
	// reported to this adapter are collected and reported as a timeseries.
	// This will be rounded to the nearest second and rounded values less than
	// one second are not valid. Defaults to 10 seconds if not specified.
	DatapointInterval time.Duration `protobuf:"bytes,4,opt,name=datapoint_interval,json=datapointInterval,proto3,stdduration" json:"datapoint_interval"`
	// Optional.  If set to false, metrics won't be sent (but trace spans will
	// be sent, unless otherwise disabled).
	EnableMetrics bool `protobuf:"varint,5,opt,name=enable_metrics,json=enableMetrics,proto3" json:"enable_metrics,omitempty"`
	// Optional.  If set to false, trace spans won't be sent (but metrics will
	// be sent, unless otherwise disabled).
	EnableTracing bool `protobuf:"varint,6,opt,name=enable_tracing,json=enableTracing,proto3" json:"enable_tracing,omitempty"`
	// Optional.  The number of trace spans that the adapter will buffer before
	// dropping them.  This defaults to 1000 spans but can be configured higher
	// if needed.  An error message will be logged if spans are dropped.
	TracingBufferSize uint32 `protobuf:"varint,7,opt,name=tracing_buffer_size,json=tracingBufferSize,proto3" json:"tracing_buffer_size,omitempty"`
	// Optional. The uniform probability ([0.0, 1.0]) that a given span gets
	// sampled if its parent was not already sampled.  Child spans will always
	// be sampled if their parent is.  If not provided, defaults to sending all
	// spans.
	TracingSampleProbability float64 `protobuf:"fixed64,8,opt,name=tracing_sample_probability,json=tracingSampleProbability,proto3" json:"tracing_sample_probability,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_41883b00dc025265, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// Describes what metrics should be sent to SignalFx and in what form.
type Params_MetricConfig struct {
	// Required.  The name of the metric as it is sent to the adapter.  In
	// Kubernetes this is of the form "<name>.metric.<namespace>" where
	// "<name>" is the name field of the metric resource, and "<namespace>"
	// is the namespace of the metric resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The metric type of the metric
	Type Params_MetricConfig_Type `protobuf:"varint,4,opt,name=type,proto3,enum=adapter.signalfx.config.Params_MetricConfig_Type" json:"type,omitempty"`
}

func (m *Params_MetricConfig) Reset()      { *m = Params_MetricConfig{} }
func (*Params_MetricConfig) ProtoMessage() {}
func (*Params_MetricConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_41883b00dc025265, []int{0, 0}
}
func (m *Params_MetricConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params_MetricConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params_MetricConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params_MetricConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params_MetricConfig.Merge(m, src)
}
func (m *Params_MetricConfig) XXX_Size() int {
	return m.Size()
}
func (m *Params_MetricConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Params_MetricConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Params_MetricConfig proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("adapter.signalfx.config.Params_MetricConfig_Type", Params_MetricConfig_Type_name, Params_MetricConfig_Type_value)
	proto.RegisterType((*Params)(nil), "adapter.signalfx.config.Params")
	proto.RegisterType((*Params_MetricConfig)(nil), "adapter.signalfx.config.Params.MetricConfig")
}

func init() {
	proto.RegisterFile("mixer/adapter/signalfx/config/config.proto", fileDescriptor_41883b00dc025265)
}

var fileDescriptor_41883b00dc025265 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x7d, 0xad, 0xc9, 0x9f, 0x4b, 0x53, 0xa5, 0x07, 0x12, 0x26, 0x12, 0x57, 0x53, 0x09,
	0xc9, 0x42, 0xd5, 0x59, 0x84, 0x95, 0x01, 0x52, 0x02, 0x74, 0x48, 0x52, 0x39, 0xe9, 0xc2, 0x62,
	0x9d, 0x93, 0x8b, 0x75, 0xc2, 0xf1, 0x59, 0xe7, 0x0b, 0x6a, 0x3a, 0xf1, 0x11, 0xd8, 0xe0, 0x23,
	0xf0, 0x21, 0xf8, 0x00, 0x19, 0x33, 0x76, 0x02, 0xe2, 0x2c, 0x8c, 0xfd, 0x08, 0x28, 0x67, 0x87,
	0x76, 0x61, 0xe8, 0xe4, 0xbb, 0xe7, 0xf9, 0x3d, 0xf2, 0x3d, 0xef, 0x0b, 0x9f, 0x4d, 0xf9, 0x05,
	0x93, 0x2e, 0x1d, 0xd3, 0x44, 0x31, 0xe9, 0xa6, 0x3c, 0x8c, 0x69, 0x34, 0xb9, 0x70, 0x47, 0x22,
	0x9e, 0xf0, 0xb0, 0xf8, 0x90, 0x44, 0x0a, 0x25, 0xd0, 0xc3, 0x82, 0x22, 0x5b, 0x8a, 0xe4, 0x76,
	0xf3, 0x41, 0x28, 0x42, 0xa1, 0x19, 0x77, 0x73, 0xca, 0xf1, 0x26, 0x0e, 0x85, 0x08, 0x23, 0xe6,
	0xea, 0x5b, 0x30, 0x9b, 0xb8, 0xe3, 0x99, 0xa4, 0x8a, 0x8b, 0x38, 0xf7, 0x8f, 0x7e, 0x98, 0xb0,
	0x74, 0x46, 0x25, 0x9d, 0xa6, 0xe8, 0x2d, 0x2c, 0x4f, 0x99, 0x92, 0x7c, 0x94, 0x5a, 0xc0, 0xde,
	0x75, 0x6a, 0xad, 0x63, 0xf2, 0x9f, 0x7f, 0x91, 0x3c, 0x41, 0xba, 0x1a, 0x3f, 0xd1, 0x9a, 0xb7,
	0x0d, 0xa3, 0xc7, 0x10, 0xf2, 0x38, 0x64, 0xa9, 0xf2, 0x67, 0x32, 0xb2, 0x76, 0x6c, 0xe0, 0x54,
	0xbd, 0x6a, 0xae, 0x9c, 0xcb, 0x08, 0x3d, 0x81, 0x7b, 0x74, 0x34, 0x62, 0x69, 0xea, 0x2b, 0xf1,
	0x91, 0xc5, 0xd6, 0xae, 0x06, 0x6a, 0xb9, 0x36, 0xdc, 0x48, 0xc8, 0x83, 0x68, 0x4c, 0x15, 0x4d,
	0x04, 0x8f, 0x95, 0xcf, 0x63, 0xc5, 0xe4, 0x27, 0x1a, 0x59, 0xa6, 0x0d, 0x9c, 0x5a, 0xeb, 0x11,
	0xc9, 0x1b, 0x91, 0x6d, 0x23, 0xf2, 0xa6, 0x68, 0xd4, 0xae, 0x2c, 0x7e, 0x1e, 0x1a, 0xdf, 0x7e,
	0x1d, 0x02, 0xef, 0xe0, 0x5f, 0xfc, 0xb4, 0x48, 0xa3, 0xa7, 0x70, 0x9f, 0xc5, 0x34, 0x88, 0x98,
	0xbf, 0x2d, 0x79, 0xcf, 0x06, 0x4e, 0xc5, 0xab, 0xe7, 0x6a, 0xb7, 0x78, 0xfc, 0x0d, 0xa6, 0x24,
	0x1d, 0xf1, 0x38, 0xb4, 0x4a, 0xb7, 0xb1, 0x61, 0x2e, 0x22, 0x02, 0xef, 0x17, 0xbe, 0x1f, 0xcc,
	0x26, 0x13, 0x26, 0xfd, 0x94, 0x5f, 0x32, 0xab, 0x6c, 0x03, 0xa7, 0xee, 0x1d, 0x14, 0x56, 0x5b,
	0x3b, 0x03, 0x7e, 0xc9, 0xd0, 0x4b, 0xd8, 0xdc, 0xf2, 0x29, 0x9d, 0x26, 0x11, 0xf3, 0x13, 0x29,
	0x02, 0x1a, 0xf0, 0x88, 0xab, 0xb9, 0x55, 0xb1, 0x81, 0x03, 0x3c, 0xab, 0x20, 0x06, 0x1a, 0x38,
	0xbb, 0xf1, 0x9b, 0x5f, 0x01, 0xdc, 0xbb, 0x3d, 0x6b, 0x84, 0xa0, 0x19, 0xd3, 0x29, 0xb3, 0x80,
	0x9e, 0x9d, 0x3e, 0xa3, 0x0e, 0x34, 0xd5, 0x3c, 0x61, 0x7a, 0x4c, 0xfb, 0xad, 0xe7, 0x77, 0xd9,
	0x1d, 0x19, 0xce, 0x13, 0xe6, 0xe9, 0xf8, 0xd1, 0x31, 0x34, 0x37, 0x37, 0x54, 0x81, 0x66, 0xaf,
	0xdf, 0xeb, 0x34, 0x0c, 0x54, 0x83, 0xe5, 0x93, 0xfe, 0x79, 0x6f, 0xd8, 0xf1, 0x1a, 0x00, 0xd5,
	0x61, 0xf5, 0xfd, 0xe9, 0x60, 0xd8, 0x7f, 0xe7, 0xbd, 0xee, 0x36, 0x76, 0xda, 0xaf, 0x16, 0x2b,
	0x6c, 0x2c, 0x57, 0xd8, 0xb8, 0x5a, 0x61, 0xe3, 0x7a, 0x85, 0x8d, 0xcf, 0x19, 0x06, 0xdf, 0x33,
	0x6c, 0x2c, 0x32, 0x0c, 0x96, 0x19, 0x06, 0xbf, 0x33, 0x0c, 0xfe, 0x64, 0xd8, 0xb8, 0xce, 0x30,
	0xf8, 0xb2, 0xc6, 0xc6, 0x72, 0x8d, 0x8d, 0xab, 0x35, 0x36, 0x3e, 0x94, 0xf2, 0xe7, 0x04, 0x25,
	0xbd, 0xc7, 0x17, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xa3, 0x64, 0x48, 0x04, 0x03, 0x00,
	0x00,
}

func (x Params_MetricConfig_Type) String() string {
	s, ok := Params_MetricConfig_Type_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TracingSampleProbability != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.TracingSampleProbability))))
		i--
		dAtA[i] = 0x41
	}
	if m.TracingBufferSize != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.TracingBufferSize))
		i--
		dAtA[i] = 0x38
	}
	if m.EnableTracing {
		i--
		if m.EnableTracing {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.EnableMetrics {
		i--
		if m.EnableMetrics {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DatapointInterval, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DatapointInterval):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintConfig(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if len(m.AccessToken) > 0 {
		i -= len(m.AccessToken)
		copy(dAtA[i:], m.AccessToken)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.AccessToken)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.IngestUrl) > 0 {
		i -= len(m.IngestUrl)
		copy(dAtA[i:], m.IngestUrl)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.IngestUrl)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Metrics) > 0 {
		for iNdEx := len(m.Metrics) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Metrics[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Params_MetricConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params_MetricConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params_MetricConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Metrics) > 0 {
		for _, e := range m.Metrics {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	l = len(m.IngestUrl)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.AccessToken)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DatapointInterval)
	n += 1 + l + sovConfig(uint64(l))
	if m.EnableMetrics {
		n += 2
	}
	if m.EnableTracing {
		n += 2
	}
	if m.TracingBufferSize != 0 {
		n += 1 + sovConfig(uint64(m.TracingBufferSize))
	}
	if m.TracingSampleProbability != 0 {
		n += 9
	}
	return n
}

func (m *Params_MetricConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovConfig(uint64(m.Type))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForMetrics := "[]*Params_MetricConfig{"
	for _, f := range this.Metrics {
		repeatedStringForMetrics += strings.Replace(fmt.Sprintf("%v", f), "Params_MetricConfig", "Params_MetricConfig", 1) + ","
	}
	repeatedStringForMetrics += "}"
	s := strings.Join([]string{`&Params{`,
		`Metrics:` + repeatedStringForMetrics + `,`,
		`IngestUrl:` + fmt.Sprintf("%v", this.IngestUrl) + `,`,
		`AccessToken:` + fmt.Sprintf("%v", this.AccessToken) + `,`,
		`DatapointInterval:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.DatapointInterval), "Duration", "types.Duration", 1), `&`, ``, 1) + `,`,
		`EnableMetrics:` + fmt.Sprintf("%v", this.EnableMetrics) + `,`,
		`EnableTracing:` + fmt.Sprintf("%v", this.EnableTracing) + `,`,
		`TracingBufferSize:` + fmt.Sprintf("%v", this.TracingBufferSize) + `,`,
		`TracingSampleProbability:` + fmt.Sprintf("%v", this.TracingSampleProbability) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Params_MetricConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params_MetricConfig{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metrics = append(m.Metrics, &Params_MetricConfig{})
			if err := m.Metrics[len(m.Metrics)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IngestUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IngestUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DatapointInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DatapointInterval, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableMetrics", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableMetrics = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableTracing", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableTracing = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TracingBufferSize", wireType)
			}
			m.TracingBufferSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TracingBufferSize |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field TracingSampleProbability", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.TracingSampleProbability = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params_MetricConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MetricConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Params_MetricConfig_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthConfig
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)
