// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/interchain_accounts/host/v1/host.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the set of on-chain interchain accounts parameters.
// The following parameters may be used to disable the host submodule.
type Params struct {
	// host_enabled enables or disables the host submodule.
	HostEnabled bool `protobuf:"varint,1,opt,name=host_enabled,json=hostEnabled,proto3" json:"host_enabled,omitempty"`
	// allow_messages defines a list of sdk message typeURLs allowed to be executed on a host chain.
	AllowMessages []string `protobuf:"bytes,2,rep,name=allow_messages,json=allowMessages,proto3" json:"allow_messages,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_48e202774f13d08e, []int{0}
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

func (m *Params) GetHostEnabled() bool {
	if m != nil {
		return m.HostEnabled
	}
	return false
}

func (m *Params) GetAllowMessages() []string {
	if m != nil {
		return m.AllowMessages
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "ibc.applications.interchain_accounts.host.v1.Params")
}

func init() {
	proto.RegisterFile("ibc/applications/interchain_accounts/host/v1/host.proto", fileDescriptor_48e202774f13d08e)
}

var fileDescriptor_48e202774f13d08e = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xcf, 0xbf, 0x4a, 0xc5, 0x30,
	0x14, 0xc7, 0xf1, 0x46, 0xe1, 0xa2, 0xf5, 0xcf, 0xd0, 0xe9, 0x4e, 0xe1, 0x2a, 0x08, 0x77, 0xb0,
	0x09, 0x57, 0x87, 0xee, 0x82, 0x8b, 0x20, 0x48, 0x47, 0x97, 0x92, 0xa4, 0xa1, 0x0d, 0x24, 0x39,
	0xa5, 0x27, 0xad, 0xf8, 0x16, 0x3e, 0x96, 0x63, 0x47, 0x47, 0x69, 0x5f, 0x44, 0x5a, 0x05, 0x15,
	0xee, 0x74, 0xe0, 0x03, 0x3f, 0x38, 0xdf, 0x38, 0x33, 0x52, 0x71, 0xd1, 0x34, 0xd6, 0x28, 0x11,
	0x0c, 0x78, 0xe4, 0xc6, 0x07, 0xdd, 0xaa, 0x5a, 0x18, 0x5f, 0x08, 0xa5, 0xa0, 0xf3, 0x01, 0x79,
	0x0d, 0x18, 0x78, 0xbf, 0x5b, 0x2e, 0x6b, 0x5a, 0x08, 0x90, 0x5c, 0x1b, 0xa9, 0xd8, 0xdf, 0x21,
	0xdb, 0x33, 0x64, 0xcb, 0xa0, 0xdf, 0x5d, 0xe6, 0xf1, 0xea, 0x49, 0xb4, 0xc2, 0x61, 0x72, 0x11,
	0x9f, 0xce, 0x58, 0x68, 0x2f, 0xa4, 0xd5, 0xe5, 0x9a, 0x6c, 0xc8, 0xf6, 0x28, 0x3f, 0x99, 0xed,
	0xfe, 0x9b, 0x92, 0xab, 0xf8, 0x5c, 0x58, 0x0b, 0x2f, 0x85, 0xd3, 0x88, 0xa2, 0xd2, 0xb8, 0x3e,
	0xd8, 0x1c, 0x6e, 0x8f, 0xf3, 0xb3, 0x45, 0x1f, 0x7f, 0xf0, 0xae, 0x7c, 0x1f, 0x29, 0x19, 0x46,
	0x4a, 0x3e, 0x47, 0x4a, 0xde, 0x26, 0x1a, 0x0d, 0x13, 0x8d, 0x3e, 0x26, 0x1a, 0x3d, 0x3f, 0x54,
	0x26, 0xd4, 0x9d, 0x64, 0x0a, 0x1c, 0x57, 0x80, 0x0e, 0x90, 0x1b, 0xa9, 0xd2, 0x0a, 0x78, 0x9f,
	0x71, 0x07, 0x65, 0x67, 0x35, 0xce, 0xd1, 0xc8, 0x6f, 0xb2, 0xf4, 0xf7, 0xed, 0xf4, 0x7f, 0x6f,
	0x78, 0x6d, 0x34, 0xca, 0xd5, 0x92, 0x7b, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xb5, 0x3b,
	0x6b, 0x29, 0x01, 0x00, 0x00,
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
	if len(m.AllowMessages) > 0 {
		for iNdEx := len(m.AllowMessages) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowMessages[iNdEx])
			copy(dAtA[i:], m.AllowMessages[iNdEx])
			i = encodeVarintHost(dAtA, i, uint64(len(m.AllowMessages[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.HostEnabled {
		i--
		if m.HostEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintHost(dAtA []byte, offset int, v uint64) int {
	offset -= sovHost(v)
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
	if m.HostEnabled {
		n += 2
	}
	if len(m.AllowMessages) > 0 {
		for _, s := range m.AllowMessages {
			l = len(s)
			n += 1 + l + sovHost(uint64(l))
		}
	}
	return n
}

func sovHost(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHost(x uint64) (n int) {
	return sovHost(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
			m.HostEnabled = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowMessages", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowMessages = append(m.AllowMessages, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHost
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
func skipHost(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHost
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
					return 0, ErrIntOverflowHost
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHost
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
				return 0, ErrInvalidLengthHost
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHost
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHost
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHost        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHost          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHost = fmt.Errorf("proto: unexpected end of group")
)
