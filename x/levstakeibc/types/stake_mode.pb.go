// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stayking/levstakeibc/stake_mode.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type StakingType int32

const (
	StakingType_NotLeverageType StakingType = 0
	StakingType_LeverageType    StakingType = 1
)

var StakingType_name = map[int32]string{
	0: "NotLeverageType",
	1: "LeverageType",
}

var StakingType_value = map[string]int32{
	"NotLeverageType": 0,
	"LeverageType":    1,
}

func (x StakingType) String() string {
	return proto.EnumName(StakingType_name, int32(x))
}

func (StakingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a85e3e0c98e10aa5, []int{0}
}

type StakingMode struct {
	Type StakingType `protobuf:"varint,1,opt,name=type,proto3,enum=stayking.levstakeibc.StakingType" json:"type,omitempty"`
}

func (m *StakingMode) Reset()         { *m = StakingMode{} }
func (m *StakingMode) String() string { return proto.CompactTextString(m) }
func (*StakingMode) ProtoMessage()    {}
func (*StakingMode) Descriptor() ([]byte, []int) {
	return fileDescriptor_a85e3e0c98e10aa5, []int{0}
}
func (m *StakingMode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakingMode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakingMode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakingMode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakingMode.Merge(m, src)
}
func (m *StakingMode) XXX_Size() int {
	return m.Size()
}
func (m *StakingMode) XXX_DiscardUnknown() {
	xxx_messageInfo_StakingMode.DiscardUnknown(m)
}

var xxx_messageInfo_StakingMode proto.InternalMessageInfo

func (m *StakingMode) GetType() StakingType {
	if m != nil {
		return m.Type
	}
	return StakingType_NotLeverageType
}

func init() {
	proto.RegisterEnum("stayking.levstakeibc.StakingType", StakingType_name, StakingType_value)
	proto.RegisterType((*StakingMode)(nil), "stayking.levstakeibc.StakingMode")
}

func init() {
	proto.RegisterFile("stayking/levstakeibc/stake_mode.proto", fileDescriptor_a85e3e0c98e10aa5)
}

var fileDescriptor_a85e3e0c98e10aa5 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x2e, 0x49, 0xac,
	0xcc, 0xce, 0xcc, 0x4b, 0xd7, 0xcf, 0x49, 0x2d, 0x2b, 0x2e, 0x49, 0xcc, 0x4e, 0xcd, 0x4c, 0x4a,
	0xd6, 0x07, 0x33, 0xe2, 0x73, 0xf3, 0x53, 0x52, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44,
	0x60, 0xca, 0xf4, 0x90, 0x94, 0x49, 0x49, 0x26, 0xe7, 0x17, 0xe7, 0xe6, 0x17, 0xc7, 0x83, 0xd5,
	0xe8, 0x43, 0x38, 0x10, 0x0d, 0x4a, 0x2e, 0x5c, 0xdc, 0xc1, 0x25, 0x89, 0x20, 0x1d, 0xbe, 0xf9,
	0x29, 0xa9, 0x42, 0xa6, 0x5c, 0x2c, 0x25, 0x95, 0x05, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x7c,
	0x46, 0x8a, 0x7a, 0xd8, 0x8c, 0xd3, 0x83, 0x6a, 0x08, 0xa9, 0x2c, 0x48, 0x0d, 0x02, 0x2b, 0xd7,
	0x32, 0x81, 0x9b, 0x02, 0x12, 0x14, 0x12, 0xe6, 0xe2, 0xf7, 0xcb, 0x2f, 0xf1, 0x49, 0x2d, 0x4b,
	0x2d, 0x4a, 0x4c, 0x4f, 0x05, 0x09, 0x09, 0x30, 0x08, 0x09, 0x70, 0xf1, 0xa0, 0x88, 0x30, 0x3a,
	0xf9, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x51, 0x7a, 0x66, 0x49,
	0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x71, 0x7e, 0x7e, 0x46, 0x7e, 0x66, 0xbe, 0x3e,
	0x3c, 0x00, 0xca, 0x8c, 0xf4, 0x2b, 0x50, 0x42, 0x01, 0xe4, 0x84, 0xe2, 0x24, 0x36, 0xb0, 0x87,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x4d, 0x30, 0xe1, 0x2a, 0x01, 0x00, 0x00,
}

func (m *StakingMode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakingMode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakingMode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintStakeMode(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStakeMode(dAtA []byte, offset int, v uint64) int {
	offset -= sovStakeMode(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StakingMode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovStakeMode(uint64(m.Type))
	}
	return n
}

func sovStakeMode(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStakeMode(x uint64) (n int) {
	return sovStakeMode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StakingMode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStakeMode
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
			return fmt.Errorf("proto: StakingMode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakingMode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeMode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= StakingType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStakeMode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStakeMode
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
func skipStakeMode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStakeMode
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
					return 0, ErrIntOverflowStakeMode
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
					return 0, ErrIntOverflowStakeMode
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
				return 0, ErrInvalidLengthStakeMode
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStakeMode
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStakeMode
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStakeMode        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStakeMode          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStakeMode = fmt.Errorf("proto: unexpected end of group")
)