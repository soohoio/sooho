// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stayking/records/callbacks.proto

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

// ---------------------- Transfer Callback ---------------------- //
type TransferCallback struct {
	DepositRecordId uint64 `protobuf:"varint,1,opt,name=deposit_record_id,json=depositRecordId,proto3" json:"deposit_record_id,omitempty"`
}

func (m *TransferCallback) Reset()         { *m = TransferCallback{} }
func (m *TransferCallback) String() string { return proto.CompactTextString(m) }
func (*TransferCallback) ProtoMessage()    {}
func (*TransferCallback) Descriptor() ([]byte, []int) {
	return fileDescriptor_abd4c8208668c67c, []int{0}
}
func (m *TransferCallback) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransferCallback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransferCallback.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransferCallback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferCallback.Merge(m, src)
}
func (m *TransferCallback) XXX_Size() int {
	return m.Size()
}
func (m *TransferCallback) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferCallback.DiscardUnknown(m)
}

var xxx_messageInfo_TransferCallback proto.InternalMessageInfo

func (m *TransferCallback) GetDepositRecordId() uint64 {
	if m != nil {
		return m.DepositRecordId
	}
	return 0
}

func init() {
	proto.RegisterType((*TransferCallback)(nil), "stayking.records.TransferCallback")
}

func init() { proto.RegisterFile("stayking/records/callbacks.proto", fileDescriptor_abd4c8208668c67c) }

var fileDescriptor_abd4c8208668c67c = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x2e, 0x49, 0xac,
	0xcc, 0xce, 0xcc, 0x4b, 0xd7, 0x2f, 0x4a, 0x4d, 0xce, 0x2f, 0x4a, 0x29, 0xd6, 0x4f, 0x4e, 0xcc,
	0xc9, 0x49, 0x4a, 0x4c, 0xce, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0xa9,
	0xd0, 0x83, 0xaa, 0x50, 0xb2, 0xe3, 0x12, 0x08, 0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0x4b, 0x2d, 0x72,
	0x86, 0x2a, 0x16, 0xd2, 0xe2, 0x12, 0x4c, 0x49, 0x2d, 0xc8, 0x2f, 0xce, 0x2c, 0x89, 0x87, 0x28,
	0x8b, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0xe2, 0x87, 0x4a, 0x04, 0x81, 0xc5,
	0x3d, 0x53, 0x9c, 0x3c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39,
	0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x2f,
	0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xbf, 0x38, 0x3f, 0x3f, 0x23, 0x3f,
	0x33, 0x5f, 0x1f, 0xee, 0xc0, 0x32, 0x63, 0xfd, 0x0a, 0xb8, 0x2b, 0x4b, 0x2a, 0x0b, 0x52, 0x8b,
	0x93, 0xd8, 0xc0, 0x4e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x03, 0x37, 0xd5, 0xa8, 0xc6,
	0x00, 0x00, 0x00,
}

func (m *TransferCallback) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransferCallback) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransferCallback) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DepositRecordId != 0 {
		i = encodeVarintCallbacks(dAtA, i, uint64(m.DepositRecordId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCallbacks(dAtA []byte, offset int, v uint64) int {
	offset -= sovCallbacks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TransferCallback) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DepositRecordId != 0 {
		n += 1 + sovCallbacks(uint64(m.DepositRecordId))
	}
	return n
}

func sovCallbacks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCallbacks(x uint64) (n int) {
	return sovCallbacks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransferCallback) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCallbacks
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
			return fmt.Errorf("proto: TransferCallback: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransferCallback: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositRecordId", wireType)
			}
			m.DepositRecordId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallbacks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DepositRecordId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCallbacks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCallbacks
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
func skipCallbacks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCallbacks
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
					return 0, ErrIntOverflowCallbacks
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
					return 0, ErrIntOverflowCallbacks
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
				return 0, ErrInvalidLengthCallbacks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCallbacks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCallbacks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCallbacks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCallbacks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCallbacks = fmt.Errorf("proto: unexpected end of group")
)
