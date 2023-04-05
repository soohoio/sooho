// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stayking/lendingpool/interestmodels/tripleslope/tripleslope.proto

package tripleslope

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// TripleSlope is the interest model used for lending pools
type TripleSlope struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *TripleSlope) Reset()      { *m = TripleSlope{} }
func (*TripleSlope) ProtoMessage() {}
func (*TripleSlope) Descriptor() ([]byte, []int) {
	return fileDescriptor_fca0f0756b6db415, []int{0}
}
func (m *TripleSlope) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TripleSlope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TripleSlope.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TripleSlope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TripleSlope.Merge(m, src)
}
func (m *TripleSlope) XXX_Size() int {
	return m.Size()
}
func (m *TripleSlope) XXX_DiscardUnknown() {
	xxx_messageInfo_TripleSlope.DiscardUnknown(m)
}

var xxx_messageInfo_TripleSlope proto.InternalMessageInfo

// Params defines the parameters for triple slope interest model
type Params struct {
	// defines the triple slope model asl follows:
	// for r_i where 0<=r_i<=1, and i in {0,1}
	// interest_rate (APR) = (m_i)x + b_i where r_{i-1} < x < r_i and (x < r_i if
	// i == 0)
	R []github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,rep,name=r,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"r"`
	M []github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,rep,name=m,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"m"`
	B []github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,rep,name=b,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"b"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_fca0f0756b6db415, []int{1}
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

func init() {
	proto.RegisterType((*TripleSlope)(nil), "stayking.lendingpool.interestmodels.tripleslope.TripleSlope")
	proto.RegisterType((*Params)(nil), "stayking.lendingpool.interestmodels.tripleslope.Params")
}

func init() {
	proto.RegisterFile("stayking/lendingpool/interestmodels/tripleslope/tripleslope.proto", fileDescriptor_fca0f0756b6db415)
}

var fileDescriptor_fca0f0756b6db415 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x2c, 0x2e, 0x49, 0xac,
	0xcc, 0xce, 0xcc, 0x4b, 0xd7, 0xcf, 0x49, 0xcd, 0x4b, 0xc9, 0xcc, 0x4b, 0x2f, 0xc8, 0xcf, 0xcf,
	0xd1, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0x4a, 0x2d, 0x2e, 0xc9, 0xcd, 0x4f, 0x49, 0xcd, 0x29, 0xd6,
	0x2f, 0x29, 0xca, 0x2c, 0xc8, 0x49, 0x2d, 0xce, 0xc9, 0x2f, 0x48, 0x45, 0x66, 0xeb, 0x15, 0x14,
	0xe5, 0x97, 0xe4, 0x0b, 0xe9, 0xc3, 0x8c, 0xd0, 0x43, 0x32, 0x42, 0x0f, 0xd5, 0x08, 0x3d, 0x24,
	0x6d, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0xbd, 0xfa, 0x20, 0x16, 0xc4, 0x18, 0x29, 0xc9,
	0xe4, 0xfc, 0xe2, 0xdc, 0xfc, 0xe2, 0x78, 0x88, 0x04, 0x84, 0x03, 0x91, 0x52, 0x6a, 0x60, 0xe4,
	0xe2, 0x0e, 0x01, 0x1b, 0x10, 0x0c, 0x32, 0x40, 0x28, 0x94, 0x8b, 0xad, 0x20, 0xb1, 0x28, 0x31,
	0xb7, 0x58, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0xc8, 0x5c, 0x8f, 0x44, 0x27, 0xe8, 0x05, 0x80,
	0xb5, 0x3b, 0xb1, 0x9c, 0xb8, 0x27, 0xcf, 0x10, 0x04, 0x35, 0xcc, 0x4a, 0xaa, 0x63, 0x81, 0x3c,
	0xc3, 0x8c, 0x05, 0xf2, 0x0c, 0xa7, 0xb6, 0xe8, 0xf2, 0x79, 0x42, 0x35, 0xfb, 0x82, 0x34, 0x7b,
	0x2a, 0x7d, 0x63, 0xe4, 0x62, 0x83, 0x68, 0x12, 0xf2, 0xe2, 0x62, 0x2c, 0x92, 0x60, 0x54, 0x60,
	0xd6, 0xe0, 0x74, 0xb2, 0x01, 0xe9, 0xbf, 0x75, 0x4f, 0x5e, 0x2d, 0x3d, 0xb3, 0x24, 0xa3, 0x34,
	0x49, 0x2f, 0x39, 0x3f, 0x17, 0xea, 0x72, 0x28, 0xa5, 0x5b, 0x9c, 0x92, 0xad, 0x5f, 0x52, 0x59,
	0x90, 0x5a, 0xac, 0xe7, 0x92, 0x9a, 0x7c, 0x69, 0x8b, 0x2e, 0x17, 0xd4, 0x63, 0x2e, 0xa9, 0xc9,
	0x41, 0x8c, 0x45, 0x20, 0xb3, 0x72, 0x25, 0x98, 0xa8, 0x61, 0x56, 0x2e, 0xc8, 0xac, 0x24, 0x09,
	0x66, 0x6a, 0x98, 0x95, 0x64, 0xc5, 0x01, 0x0b, 0x0a, 0xa7, 0xd4, 0x15, 0x8f, 0xe4, 0x18, 0x4f,
	0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18,
	0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x1d, 0xc9, 0x82, 0xe2, 0xfc, 0xfc,
	0x8c, 0xfc, 0xcc, 0x7c, 0x78, 0x72, 0xd0, 0x2f, 0x33, 0xd2, 0xaf, 0x20, 0x32, 0x59, 0x25, 0xb1,
	0x81, 0x63, 0xda, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xfc, 0x09, 0x55, 0x90, 0x02, 0x00,
	0x00,
}

func (this *TripleSlope) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TripleSlope)
	if !ok {
		that2, ok := that.(TripleSlope)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Params.Equal(&that1.Params) {
		return false
	}
	return true
}
func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.R) != len(that1.R) {
		return false
	}
	for i := range this.R {
		if !this.R[i].Equal(that1.R[i]) {
			return false
		}
	}
	if len(this.M) != len(that1.M) {
		return false
	}
	for i := range this.M {
		if !this.M[i].Equal(that1.M[i]) {
			return false
		}
	}
	if len(this.B) != len(that1.B) {
		return false
	}
	for i := range this.B {
		if !this.B[i].Equal(that1.B[i]) {
			return false
		}
	}
	return true
}
func (m *TripleSlope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TripleSlope) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TripleSlope) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTripleslope(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
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
	if len(m.B) > 0 {
		for iNdEx := len(m.B) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.B[iNdEx].Size()
				i -= size
				if _, err := m.B[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintTripleslope(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.M) > 0 {
		for iNdEx := len(m.M) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.M[iNdEx].Size()
				i -= size
				if _, err := m.M[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintTripleslope(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.R) > 0 {
		for iNdEx := len(m.R) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.R[iNdEx].Size()
				i -= size
				if _, err := m.R[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintTripleslope(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTripleslope(dAtA []byte, offset int, v uint64) int {
	offset -= sovTripleslope(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TripleSlope) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovTripleslope(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.R) > 0 {
		for _, e := range m.R {
			l = e.Size()
			n += 1 + l + sovTripleslope(uint64(l))
		}
	}
	if len(m.M) > 0 {
		for _, e := range m.M {
			l = e.Size()
			n += 1 + l + sovTripleslope(uint64(l))
		}
	}
	if len(m.B) > 0 {
		for _, e := range m.B {
			l = e.Size()
			n += 1 + l + sovTripleslope(uint64(l))
		}
	}
	return n
}

func sovTripleslope(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTripleslope(x uint64) (n int) {
	return sovTripleslope(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TripleSlope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTripleslope
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
			return fmt.Errorf("proto: TripleSlope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TripleSlope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTripleslope
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
				return ErrInvalidLengthTripleslope
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTripleslope
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTripleslope(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTripleslope
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTripleslope
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
				return fmt.Errorf("proto: wrong wireType = %d for field R", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTripleslope
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
				return ErrInvalidLengthTripleslope
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTripleslope
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.R = append(m.R, v)
			if err := m.R[len(m.R)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field M", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTripleslope
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
				return ErrInvalidLengthTripleslope
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTripleslope
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.M = append(m.M, v)
			if err := m.M[len(m.M)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTripleslope
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
				return ErrInvalidLengthTripleslope
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTripleslope
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.B = append(m.B, v)
			if err := m.B[len(m.B)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTripleslope(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTripleslope
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
func skipTripleslope(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTripleslope
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
					return 0, ErrIntOverflowTripleslope
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
					return 0, ErrIntOverflowTripleslope
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
				return 0, ErrInvalidLengthTripleslope
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTripleslope
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTripleslope
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTripleslope        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTripleslope          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTripleslope = fmt.Errorf("proto: unexpected end of group")
)