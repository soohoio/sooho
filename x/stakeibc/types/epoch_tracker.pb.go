// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/stakeibc/epoch_tracker.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type EpochTracker struct {
	EpochIdentifier    string `protobuf:"bytes,1,opt,name=epoch_identifier,json=epochIdentifier,proto3" json:"epoch_identifier,omitempty"`
	EpochNumber        uint64 `protobuf:"varint,2,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
	NextEpochStartTime uint64 `protobuf:"varint,3,opt,name=next_epoch_start_time,json=nextEpochStartTime,proto3" json:"next_epoch_start_time,omitempty"`
	Duration           uint64 `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (m *EpochTracker) Reset()         { *m = EpochTracker{} }
func (m *EpochTracker) String() string { return proto.CompactTextString(m) }
func (*EpochTracker) ProtoMessage()    {}
func (*EpochTracker) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7c48143f24adf66, []int{0}
}
func (m *EpochTracker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochTracker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochTracker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochTracker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochTracker.Merge(m, src)
}
func (m *EpochTracker) XXX_Size() int {
	return m.Size()
}
func (m *EpochTracker) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochTracker.DiscardUnknown(m)
}

var xxx_messageInfo_EpochTracker proto.InternalMessageInfo

func (m *EpochTracker) GetEpochIdentifier() string {
	if m != nil {
		return m.EpochIdentifier
	}
	return ""
}

func (m *EpochTracker) GetEpochNumber() uint64 {
	if m != nil {
		return m.EpochNumber
	}
	return 0
}

func (m *EpochTracker) GetNextEpochStartTime() uint64 {
	if m != nil {
		return m.NextEpochStartTime
	}
	return 0
}

func (m *EpochTracker) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func init() {
	proto.RegisterType((*EpochTracker)(nil), "stride.stakeibc.EpochTracker")
}

func init() {
	proto.RegisterFile("stride/stakeibc/epoch_tracker.proto", fileDescriptor_e7c48143f24adf66)
}

var fileDescriptor_e7c48143f24adf66 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xd0, 0x4f, 0x4a, 0xc3, 0x40,
	0x14, 0x06, 0xf0, 0x8c, 0x16, 0xd1, 0xb1, 0x50, 0x19, 0x10, 0x82, 0x8b, 0xa1, 0xea, 0xa6, 0x82,
	0x24, 0x88, 0x37, 0x10, 0x14, 0xdc, 0xb8, 0xa8, 0x5d, 0xb9, 0x09, 0xf9, 0xf3, 0x6c, 0x1e, 0x21,
	0x79, 0x61, 0xe6, 0x05, 0xda, 0x5b, 0x78, 0x0f, 0x2f, 0xe2, 0xb2, 0x4b, 0x97, 0x92, 0x5c, 0x44,
	0x32, 0x83, 0x75, 0x39, 0xdf, 0xf7, 0x1b, 0x1e, 0x7c, 0xf2, 0xda, 0xb2, 0xc1, 0x02, 0x62, 0xcb,
	0x69, 0x05, 0x98, 0xe5, 0x31, 0xb4, 0x94, 0x97, 0x09, 0x9b, 0x34, 0xaf, 0xc0, 0x44, 0xad, 0x21,
	0x26, 0x35, 0xf3, 0x28, 0xfa, 0x43, 0x57, 0x9f, 0x42, 0x4e, 0x1f, 0x47, 0xb8, 0xf2, 0x4e, 0xdd,
	0xc8, 0x33, 0xff, 0x11, 0x0b, 0x68, 0x18, 0xdf, 0x11, 0x4c, 0x28, 0xe6, 0x62, 0x71, 0xb2, 0x9c,
	0xb9, 0xfc, 0x79, 0x1f, 0xab, 0x4b, 0x39, 0xf5, 0xb4, 0xe9, 0xea, 0x0c, 0x4c, 0x78, 0x30, 0x17,
	0x8b, 0xc9, 0xf2, 0xd4, 0x65, 0x2f, 0x2e, 0x52, 0x77, 0xf2, 0xbc, 0x81, 0x0d, 0x27, 0xde, 0x59,
	0x4e, 0x0d, 0x27, 0x8c, 0x35, 0x84, 0x87, 0xce, 0xaa, 0xb1, 0x74, 0xe7, 0x5f, 0xc7, 0x6a, 0x85,
	0x35, 0xa8, 0x0b, 0x79, 0x5c, 0x74, 0x26, 0x65, 0xa4, 0x26, 0x9c, 0x38, 0xb5, 0x7f, 0x3f, 0x3c,
	0x7d, 0xf5, 0x5a, 0xec, 0x7a, 0x2d, 0x7e, 0x7a, 0x2d, 0x3e, 0x06, 0x1d, 0xec, 0x06, 0x1d, 0x7c,
	0x0f, 0x3a, 0x78, 0xbb, 0x5d, 0x23, 0x97, 0x5d, 0x16, 0xe5, 0x54, 0xc7, 0x96, 0xa8, 0x24, 0xa4,
	0x71, 0x89, 0x6d, 0x85, 0xcd, 0x3a, 0xde, 0xfc, 0x8f, 0xc2, 0xdb, 0x16, 0x6c, 0x76, 0xe4, 0xd6,
	0xb8, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x66, 0xc8, 0x49, 0xf1, 0x34, 0x01, 0x00, 0x00,
}

func (m *EpochTracker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochTracker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochTracker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Duration != 0 {
		i = encodeVarintEpochTracker(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x20
	}
	if m.NextEpochStartTime != 0 {
		i = encodeVarintEpochTracker(dAtA, i, uint64(m.NextEpochStartTime))
		i--
		dAtA[i] = 0x18
	}
	if m.EpochNumber != 0 {
		i = encodeVarintEpochTracker(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x10
	}
	if len(m.EpochIdentifier) > 0 {
		i -= len(m.EpochIdentifier)
		copy(dAtA[i:], m.EpochIdentifier)
		i = encodeVarintEpochTracker(dAtA, i, uint64(len(m.EpochIdentifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEpochTracker(dAtA []byte, offset int, v uint64) int {
	offset -= sovEpochTracker(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EpochTracker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EpochIdentifier)
	if l > 0 {
		n += 1 + l + sovEpochTracker(uint64(l))
	}
	if m.EpochNumber != 0 {
		n += 1 + sovEpochTracker(uint64(m.EpochNumber))
	}
	if m.NextEpochStartTime != 0 {
		n += 1 + sovEpochTracker(uint64(m.NextEpochStartTime))
	}
	if m.Duration != 0 {
		n += 1 + sovEpochTracker(uint64(m.Duration))
	}
	return n
}

func sovEpochTracker(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEpochTracker(x uint64) (n int) {
	return sovEpochTracker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EpochTracker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpochTracker
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
			return fmt.Errorf("proto: EpochTracker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochTracker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochTracker
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
				return ErrInvalidLengthEpochTracker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEpochTracker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextEpochStartTime", wireType)
			}
			m.NextEpochStartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextEpochStartTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEpochTracker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEpochTracker
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
func skipEpochTracker(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEpochTracker
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
					return 0, ErrIntOverflowEpochTracker
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
					return 0, ErrIntOverflowEpochTracker
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
				return 0, ErrInvalidLengthEpochTracker
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEpochTracker
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEpochTracker
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEpochTracker        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEpochTracker          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEpochTracker = fmt.Errorf("proto: unexpected end of group")
)
