// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stayking/levstakeibc/callbacks.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type SplitDelegation struct {
	Validator string                                 `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	Amount    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
}

func (m *SplitDelegation) Reset()         { *m = SplitDelegation{} }
func (m *SplitDelegation) String() string { return proto.CompactTextString(m) }
func (*SplitDelegation) ProtoMessage()    {}
func (*SplitDelegation) Descriptor() ([]byte, []int) {
	return fileDescriptor_aba33a0433787ff2, []int{0}
}
func (m *SplitDelegation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SplitDelegation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SplitDelegation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SplitDelegation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SplitDelegation.Merge(m, src)
}
func (m *SplitDelegation) XXX_Size() int {
	return m.Size()
}
func (m *SplitDelegation) XXX_DiscardUnknown() {
	xxx_messageInfo_SplitDelegation.DiscardUnknown(m)
}

var xxx_messageInfo_SplitDelegation proto.InternalMessageInfo

func (m *SplitDelegation) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

type DelegateCallback struct {
	HostZoneId       string             `protobuf:"bytes,1,opt,name=host_zone_id,json=hostZoneId,proto3" json:"host_zone_id,omitempty"`
	DepositRecordId  uint64             `protobuf:"varint,2,opt,name=deposit_record_id,json=depositRecordId,proto3" json:"deposit_record_id,omitempty"`
	SplitDelegations []*SplitDelegation `protobuf:"bytes,3,rep,name=split_delegations,json=splitDelegations,proto3" json:"split_delegations,omitempty"`
}

func (m *DelegateCallback) Reset()         { *m = DelegateCallback{} }
func (m *DelegateCallback) String() string { return proto.CompactTextString(m) }
func (*DelegateCallback) ProtoMessage()    {}
func (*DelegateCallback) Descriptor() ([]byte, []int) {
	return fileDescriptor_aba33a0433787ff2, []int{1}
}
func (m *DelegateCallback) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegateCallback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegateCallback.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegateCallback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegateCallback.Merge(m, src)
}
func (m *DelegateCallback) XXX_Size() int {
	return m.Size()
}
func (m *DelegateCallback) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegateCallback.DiscardUnknown(m)
}

var xxx_messageInfo_DelegateCallback proto.InternalMessageInfo

func (m *DelegateCallback) GetHostZoneId() string {
	if m != nil {
		return m.HostZoneId
	}
	return ""
}

func (m *DelegateCallback) GetDepositRecordId() uint64 {
	if m != nil {
		return m.DepositRecordId
	}
	return 0
}

func (m *DelegateCallback) GetSplitDelegations() []*SplitDelegation {
	if m != nil {
		return m.SplitDelegations
	}
	return nil
}

func init() {
	proto.RegisterType((*SplitDelegation)(nil), "stayking.levstakeibc.SplitDelegation")
	proto.RegisterType((*DelegateCallback)(nil), "stayking.levstakeibc.DelegateCallback")
}

func init() {
	proto.RegisterFile("stayking/levstakeibc/callbacks.proto", fileDescriptor_aba33a0433787ff2)
}

var fileDescriptor_aba33a0433787ff2 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x41, 0x6b, 0xea, 0x40,
	0x18, 0x4c, 0x9e, 0x0f, 0xc1, 0x7d, 0x0f, 0xd4, 0xe0, 0x41, 0xe4, 0x11, 0x45, 0x5e, 0x8b, 0x14,
	0x9a, 0x45, 0xfb, 0x0f, 0x6c, 0x29, 0x08, 0x3d, 0xa5, 0x37, 0x2f, 0x61, 0x93, 0x5d, 0xe2, 0x92,
	0xb8, 0x5f, 0xc8, 0xb7, 0xa6, 0xb5, 0xbf, 0xa2, 0x3f, 0xa6, 0x3f, 0xc2, 0xa3, 0xc7, 0xd2, 0x83,
	0x14, 0xfd, 0x23, 0x25, 0x31, 0xb5, 0x5a, 0x7a, 0xda, 0x65, 0x66, 0x96, 0xd9, 0x99, 0x21, 0xff,
	0x51, 0xb3, 0x65, 0x24, 0x55, 0x48, 0x63, 0x91, 0xa1, 0x66, 0x91, 0x90, 0x7e, 0x40, 0x03, 0x16,
	0xc7, 0x3e, 0x0b, 0x22, 0x74, 0x92, 0x14, 0x34, 0x58, 0xad, 0x4f, 0x95, 0x73, 0xa4, 0xea, 0xb4,
	0x42, 0x08, 0xa1, 0x10, 0xd0, 0xfc, 0xb6, 0xd7, 0x76, 0xec, 0x00, 0x70, 0x0e, 0x48, 0x7d, 0x86,
	0x82, 0x66, 0x43, 0x5f, 0x68, 0x36, 0xa4, 0x01, 0x48, 0xb5, 0xe7, 0xfb, 0x0f, 0xa4, 0x7e, 0x9f,
	0xc4, 0x52, 0xdf, 0x88, 0x58, 0x84, 0x4c, 0x4b, 0x50, 0xd6, 0x3f, 0x52, 0xcb, 0x58, 0x2c, 0x39,
	0xd3, 0x90, 0xb6, 0xcd, 0x9e, 0x39, 0xa8, 0xb9, 0x5f, 0x80, 0x75, 0x4b, 0xaa, 0x6c, 0x0e, 0x0b,
	0xa5, 0xdb, 0xbf, 0x72, 0x6a, 0xec, 0xac, 0x36, 0x5d, 0xe3, 0x6d, 0xd3, 0x3d, 0x0f, 0xa5, 0x9e,
	0x2d, 0x7c, 0x27, 0x80, 0x39, 0x2d, 0x3d, 0xf7, 0xc7, 0x25, 0xf2, 0x88, 0xea, 0x65, 0x22, 0xd0,
	0x99, 0x28, 0xed, 0x96, 0xaf, 0xfb, 0x2f, 0x26, 0x69, 0x94, 0xa6, 0xe2, 0xba, 0x0c, 0x68, 0xf5,
	0xc8, 0xdf, 0x19, 0xa0, 0xf6, 0x9e, 0x40, 0x09, 0x4f, 0xf2, 0xd2, 0x9d, 0xe4, 0xd8, 0x14, 0x94,
	0x98, 0x70, 0xeb, 0x82, 0x34, 0xb9, 0x48, 0x00, 0xa5, 0xf6, 0x52, 0x11, 0x40, 0xca, 0x73, 0x59,
	0xfe, 0x93, 0xdf, 0x6e, 0xbd, 0x24, 0xdc, 0x02, 0x9f, 0x70, 0xcb, 0x25, 0x4d, 0xcc, 0xb3, 0x79,
	0xfc, 0x10, 0x0e, 0xdb, 0x95, 0x5e, 0x65, 0xf0, 0x67, 0x74, 0xe6, 0xfc, 0xd4, 0xa1, 0xf3, 0xad,
	0x0a, 0xb7, 0x81, 0xa7, 0x00, 0x8e, 0xef, 0x56, 0x5b, 0xdb, 0x5c, 0x6f, 0x6d, 0xf3, 0x7d, 0x6b,
	0x9b, 0xcf, 0x3b, 0xdb, 0x58, 0xef, 0x6c, 0xe3, 0x75, 0x67, 0x1b, 0xd3, 0xd1, 0x51, 0x01, 0x08,
	0x30, 0x03, 0x09, 0xf4, 0x30, 0x67, 0x36, 0xa2, 0x8f, 0x27, 0x9b, 0x16, 0x85, 0xf8, 0xd5, 0x62,
	0x84, 0xab, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0xe3, 0xe2, 0xda, 0xf8, 0x01, 0x00, 0x00,
}

func (m *SplitDelegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SplitDelegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SplitDelegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCallbacks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintCallbacks(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DelegateCallback) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegateCallback) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegateCallback) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SplitDelegations) > 0 {
		for iNdEx := len(m.SplitDelegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SplitDelegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCallbacks(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.DepositRecordId != 0 {
		i = encodeVarintCallbacks(dAtA, i, uint64(m.DepositRecordId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.HostZoneId) > 0 {
		i -= len(m.HostZoneId)
		copy(dAtA[i:], m.HostZoneId)
		i = encodeVarintCallbacks(dAtA, i, uint64(len(m.HostZoneId)))
		i--
		dAtA[i] = 0xa
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
func (m *SplitDelegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovCallbacks(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovCallbacks(uint64(l))
	return n
}

func (m *DelegateCallback) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HostZoneId)
	if l > 0 {
		n += 1 + l + sovCallbacks(uint64(l))
	}
	if m.DepositRecordId != 0 {
		n += 1 + sovCallbacks(uint64(m.DepositRecordId))
	}
	if len(m.SplitDelegations) > 0 {
		for _, e := range m.SplitDelegations {
			l = e.Size()
			n += 1 + l + sovCallbacks(uint64(l))
		}
	}
	return n
}

func sovCallbacks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCallbacks(x uint64) (n int) {
	return sovCallbacks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SplitDelegation) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SplitDelegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SplitDelegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallbacks
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
				return ErrInvalidLengthCallbacks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCallbacks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallbacks
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
				return ErrInvalidLengthCallbacks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCallbacks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *DelegateCallback) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DelegateCallback: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegateCallback: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostZoneId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallbacks
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
				return ErrInvalidLengthCallbacks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCallbacks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostZoneId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SplitDelegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallbacks
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
				return ErrInvalidLengthCallbacks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCallbacks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SplitDelegations = append(m.SplitDelegations, &SplitDelegation{})
			if err := m.SplitDelegations[len(m.SplitDelegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
