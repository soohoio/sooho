// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stayking/stakeibc/epoch_tracker.proto

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
	return fileDescriptor_f83fb06cc54ac384, []int{0}
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
	proto.RegisterType((*EpochTracker)(nil), "stayking.stakeibc.EpochTracker")
}

func init() {
	proto.RegisterFile("stayking/stakeibc/epoch_tracker.proto", fileDescriptor_f83fb06cc54ac384)
}

var fileDescriptor_f83fb06cc54ac384 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xd0, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0x07, 0xf0, 0x45, 0x87, 0x68, 0x1c, 0xa8, 0x01, 0xa1, 0x78, 0x08, 0x53, 0x10, 0xe6, 0xa5,
	0x41, 0xf6, 0x06, 0x82, 0x87, 0x5d, 0x3c, 0xcc, 0x9d, 0xbc, 0x94, 0xb4, 0xfb, 0x5c, 0x3f, 0x4a,
	0x93, 0x92, 0x7c, 0x95, 0xed, 0x2d, 0x7c, 0x0f, 0x5f, 0xc4, 0xe3, 0x8e, 0x1e, 0xa5, 0x7d, 0x11,
	0x69, 0xca, 0xea, 0x31, 0xff, 0xff, 0x2f, 0x7c, 0xf0, 0xe7, 0xf7, 0x9e, 0xf4, 0xae, 0x40, 0xb3,
	0x51, 0x9e, 0x74, 0x01, 0x98, 0x66, 0x0a, 0x2a, 0x9b, 0xe5, 0x09, 0x39, 0x9d, 0x15, 0xe0, 0xe2,
	0xca, 0x59, 0xb2, 0xe2, 0xea, 0xc0, 0xe2, 0x03, 0xbb, 0xfb, 0x62, 0x7c, 0xf2, 0xdc, 0xd1, 0x55,
	0x2f, 0xc5, 0x03, 0xbf, 0xec, 0xbf, 0xe2, 0x1a, 0x0c, 0xe1, 0x3b, 0x82, 0x8b, 0xd8, 0x94, 0xcd,
	0xce, 0x96, 0x17, 0x21, 0x5f, 0x0c, 0xb1, 0xb8, 0xe5, 0x93, 0x9e, 0x9a, 0xba, 0x4c, 0xc1, 0x45,
	0x47, 0x53, 0x36, 0x1b, 0x2f, 0xcf, 0x43, 0xf6, 0x12, 0x22, 0xf1, 0xc8, 0xaf, 0x0d, 0x6c, 0x29,
	0xe9, 0x9d, 0x27, 0xed, 0x28, 0x21, 0x2c, 0x21, 0x3a, 0x0e, 0x56, 0x74, 0x65, 0x38, 0xff, 0xda,
	0x55, 0x2b, 0x2c, 0x41, 0xdc, 0xf0, 0xd3, 0x75, 0xed, 0x34, 0xa1, 0x35, 0xd1, 0x38, 0xa8, 0xe1,
	0xfd, 0xb4, 0xf8, 0x6e, 0x24, 0xdb, 0x37, 0x92, 0xfd, 0x36, 0x92, 0x7d, 0xb6, 0x72, 0xb4, 0x6f,
	0xe5, 0xe8, 0xa7, 0x95, 0xa3, 0x37, 0xb5, 0x41, 0xca, 0xeb, 0x34, 0xce, 0x6c, 0xa9, 0xbc, 0xb5,
	0xb9, 0x45, 0xab, 0x86, 0x51, 0x3e, 0xe6, 0x6a, 0xfb, 0xbf, 0x0c, 0xed, 0x2a, 0xf0, 0xe9, 0x49,
	0x98, 0x64, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x84, 0x1b, 0x65, 0x1c, 0x3b, 0x01, 0x00, 0x00,
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
