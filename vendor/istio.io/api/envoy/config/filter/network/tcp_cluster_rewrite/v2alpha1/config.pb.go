// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/network/tcp_cluster_rewrite/v2alpha1/config.proto

package v2alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// TcpClusterRewrite is the config for the TCP cluster rewrite filter.
type TcpClusterRewrite struct {
	// Specifies the regex pattern to be matched in the cluster name.
	ClusterPattern string `protobuf:"bytes,1,opt,name=cluster_pattern,json=clusterPattern,proto3" json:"cluster_pattern,omitempty"`
	// Specifies the replacement for the matched cluster pattern.
	ClusterReplacement string `protobuf:"bytes,2,opt,name=cluster_replacement,json=clusterReplacement,proto3" json:"cluster_replacement,omitempty"`
}

func (m *TcpClusterRewrite) Reset()                    { *m = TcpClusterRewrite{} }
func (m *TcpClusterRewrite) String() string            { return proto.CompactTextString(m) }
func (*TcpClusterRewrite) ProtoMessage()               {}
func (*TcpClusterRewrite) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *TcpClusterRewrite) GetClusterPattern() string {
	if m != nil {
		return m.ClusterPattern
	}
	return ""
}

func (m *TcpClusterRewrite) GetClusterReplacement() string {
	if m != nil {
		return m.ClusterReplacement
	}
	return ""
}

func init() {
	proto.RegisterType((*TcpClusterRewrite)(nil), "istio.envoy.config.filter.network.tcp_cluster_rewrite.v2alpha1.TcpClusterRewrite")
}
func (m *TcpClusterRewrite) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TcpClusterRewrite) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ClusterPattern) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ClusterPattern)))
		i += copy(dAtA[i:], m.ClusterPattern)
	}
	if len(m.ClusterReplacement) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ClusterReplacement)))
		i += copy(dAtA[i:], m.ClusterReplacement)
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TcpClusterRewrite) Size() (n int) {
	var l int
	_ = l
	l = len(m.ClusterPattern)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.ClusterReplacement)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TcpClusterRewrite) Unmarshal(dAtA []byte) error {
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
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TcpClusterRewrite: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TcpClusterRewrite: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterPattern", wireType)
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterPattern = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterReplacement", wireType)
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterReplacement = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
			iNdEx += length
			if length < 0 {
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

func init() {
	proto.RegisterFile("envoy/config/filter/network/tcp_cluster_rewrite/v2alpha1/config.proto", fileDescriptorConfig)
}

var fileDescriptorConfig = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4d, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2f, 0x49, 0x2e, 0x88, 0x4f, 0xce, 0x29, 0x2d,
	0x2e, 0x49, 0x2d, 0x8a, 0x2f, 0x4a, 0x2d, 0x2f, 0xca, 0x2c, 0x49, 0xd5, 0x2f, 0x33, 0x4a, 0xcc,
	0x29, 0xc8, 0x48, 0x34, 0x84, 0x6a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xb2, 0xcb, 0x2c,
	0x2e, 0xc9, 0xcc, 0xd7, 0x03, 0x1b, 0xa6, 0x07, 0x95, 0x81, 0x18, 0xa6, 0x07, 0x35, 0x4c, 0x0f,
	0x8b, 0x61, 0x7a, 0x30, 0xc3, 0x94, 0x72, 0xb9, 0x04, 0x43, 0x92, 0x0b, 0x9c, 0x21, 0xd2, 0x41,
	0x10, 0x59, 0x21, 0x75, 0x2e, 0x7e, 0x98, 0x86, 0x82, 0xc4, 0x92, 0x92, 0xd4, 0xa2, 0x3c, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x3e, 0xa8, 0x70, 0x00, 0x44, 0x54, 0x48, 0x9f, 0x4b, 0x18,
	0x61, 0x72, 0x41, 0x4e, 0x62, 0x72, 0x6a, 0x6e, 0x6a, 0x5e, 0x89, 0x04, 0x13, 0x58, 0xb1, 0x50,
	0x32, 0xcc, 0x54, 0xb8, 0x8c, 0x53, 0xf0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e,
	0x78, 0x24, 0xc7, 0x18, 0xe5, 0x0a, 0x71, 0x7c, 0x66, 0xbe, 0x7e, 0x62, 0x41, 0xa6, 0x3e, 0xb9,
	0x01, 0x92, 0xc4, 0x06, 0x0e, 0x0a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0x0d, 0xa4,
	0x2b, 0x53, 0x01, 0x00, 0x00,
}
