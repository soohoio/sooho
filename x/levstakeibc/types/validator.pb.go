// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stayking/levstakeibc/validator.proto

package types

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

type Validator_ValidatorStatus int32

const (
	Validator_ACTIVE   Validator_ValidatorStatus = 0
	Validator_INACTIVE Validator_ValidatorStatus = 1
)

var Validator_ValidatorStatus_name = map[int32]string{
	0: "ACTIVE",
	1: "INACTIVE",
}

var Validator_ValidatorStatus_value = map[string]int32{
	"ACTIVE":   0,
	"INACTIVE": 1,
}

func (x Validator_ValidatorStatus) String() string {
	return proto.EnumName(Validator_ValidatorStatus_name, int32(x))
}

func (Validator_ValidatorStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4b22bdef2008e892, []int{1, 0}
}

type ValidatorExchangeRate struct {
	InternalTokensToSharesRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=internal_tokens_to_shares_rate,json=internalTokensToSharesRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"internal_tokens_to_shares_rate"`
	EpochNumber                uint64                                 `protobuf:"varint,2,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
}

func (m *ValidatorExchangeRate) Reset()         { *m = ValidatorExchangeRate{} }
func (m *ValidatorExchangeRate) String() string { return proto.CompactTextString(m) }
func (*ValidatorExchangeRate) ProtoMessage()    {}
func (*ValidatorExchangeRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b22bdef2008e892, []int{0}
}
func (m *ValidatorExchangeRate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorExchangeRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorExchangeRate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorExchangeRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorExchangeRate.Merge(m, src)
}
func (m *ValidatorExchangeRate) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorExchangeRate) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorExchangeRate.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorExchangeRate proto.InternalMessageInfo

func (m *ValidatorExchangeRate) GetEpochNumber() uint64 {
	if m != nil {
		return m.EpochNumber
	}
	return 0
}

type Validator struct {
	Name                 string                                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address              string                                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Status               Validator_ValidatorStatus              `protobuf:"varint,3,opt,name=status,proto3,enum=stayking.levstakeibc.Validator_ValidatorStatus" json:"status,omitempty"`
	CommissionRate       uint64                                 `protobuf:"varint,4,opt,name=commission_rate,json=commissionRate,proto3" json:"commission_rate,omitempty"`
	DelegationAmt        github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=delegation_amt,json=delegationAmt,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"delegation_amt"`
	Weight               uint64                                 `protobuf:"varint,6,opt,name=weight,proto3" json:"weight,omitempty"`
	InternalExchangeRate *ValidatorExchangeRate                 `protobuf:"bytes,7,opt,name=internal_exchange_rate,json=internalExchangeRate,proto3" json:"internal_exchange_rate,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b22bdef2008e892, []int{1}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Validator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Validator) GetStatus() Validator_ValidatorStatus {
	if m != nil {
		return m.Status
	}
	return Validator_ACTIVE
}

func (m *Validator) GetCommissionRate() uint64 {
	if m != nil {
		return m.CommissionRate
	}
	return 0
}

func (m *Validator) GetWeight() uint64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Validator) GetInternalExchangeRate() *ValidatorExchangeRate {
	if m != nil {
		return m.InternalExchangeRate
	}
	return nil
}

func init() {
	proto.RegisterEnum("stayking.levstakeibc.Validator_ValidatorStatus", Validator_ValidatorStatus_name, Validator_ValidatorStatus_value)
	proto.RegisterType((*ValidatorExchangeRate)(nil), "stayking.levstakeibc.ValidatorExchangeRate")
	proto.RegisterType((*Validator)(nil), "stayking.levstakeibc.Validator")
}

func init() {
	proto.RegisterFile("stayking/levstakeibc/validator.proto", fileDescriptor_4b22bdef2008e892)
}

var fileDescriptor_4b22bdef2008e892 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0xf5, 0x7e, 0xcd, 0xe7, 0x92, 0x6d, 0x49, 0xd1, 0x2a, 0x54, 0x26, 0x07, 0x37, 0x44, 0x08,
	0x22, 0x55, 0xb1, 0xa5, 0xf4, 0xca, 0x25, 0xa1, 0x15, 0x8a, 0x84, 0x7a, 0x70, 0x42, 0x0f, 0x5c,
	0xac, 0x8d, 0xbd, 0xb2, 0x57, 0x89, 0x77, 0x23, 0xef, 0x24, 0xb4, 0x37, 0x7e, 0x02, 0x3f, 0xa6,
	0x57, 0x6e, 0x1c, 0x7a, 0xac, 0x7a, 0x42, 0x1c, 0x2a, 0x94, 0xfc, 0x11, 0xd4, 0xf5, 0x3a, 0x09,
	0x08, 0x09, 0x71, 0xf2, 0xcc, 0xf3, 0xcc, 0xbc, 0x37, 0x6f, 0x77, 0xf1, 0x0b, 0x05, 0xf4, 0x6a,
	0xc2, 0x45, 0xe2, 0x4f, 0xd9, 0x42, 0x01, 0x9d, 0x30, 0x3e, 0x8e, 0xfc, 0x05, 0x9d, 0xf2, 0x98,
	0x82, 0xcc, 0xbd, 0x59, 0x2e, 0x41, 0x92, 0x7a, 0x59, 0xe5, 0x6d, 0x55, 0x35, 0x9e, 0x45, 0x52,
	0x65, 0x52, 0x85, 0xba, 0xc6, 0x2f, 0x92, 0xa2, 0xa1, 0x51, 0x4f, 0x64, 0x22, 0x0b, 0xfc, 0x21,
	0x2a, 0xd0, 0xd6, 0x17, 0x84, 0x9f, 0x5e, 0x94, 0xa3, 0xcf, 0x2e, 0xa3, 0x94, 0x8a, 0x84, 0x05,
	0x14, 0x18, 0xf9, 0x84, 0xb0, 0xcb, 0x05, 0xb0, 0x5c, 0xd0, 0x69, 0x08, 0x72, 0xc2, 0x84, 0x0a,
	0x41, 0x86, 0x2a, 0xa5, 0x39, 0x53, 0x61, 0x4e, 0x81, 0x39, 0xa8, 0x89, 0xda, 0xd5, 0xfe, 0xeb,
	0x9b, 0xfb, 0x23, 0xeb, 0xfb, 0xfd, 0xd1, 0xcb, 0x84, 0x43, 0x3a, 0x1f, 0x7b, 0x91, 0xcc, 0x0c,
	0xb3, 0xf9, 0x74, 0x54, 0x3c, 0xf1, 0xe1, 0x6a, 0xc6, 0x94, 0x77, 0xca, 0xa2, 0xbb, 0xeb, 0x0e,
	0x36, 0xc2, 0x4e, 0x59, 0x14, 0x34, 0x4a, 0x8e, 0x91, 0xa6, 0x18, 0xc9, 0xa1, 0x26, 0xd0, 0x12,
	0x9e, 0xe3, 0x7d, 0x36, 0x93, 0x51, 0x1a, 0x8a, 0x79, 0x36, 0x66, 0xb9, 0xf3, 0x5f, 0x13, 0xb5,
	0x2b, 0xc1, 0x9e, 0xc6, 0xce, 0x35, 0xd4, 0xfa, 0xba, 0x83, 0xab, 0x6b, 0xfd, 0x84, 0xe0, 0x8a,
	0xa0, 0x99, 0x11, 0x16, 0xe8, 0x98, 0x74, 0xf1, 0x2e, 0x8d, 0xe3, 0x9c, 0x29, 0xa5, 0xfb, 0xab,
	0x7d, 0xe7, 0xee, 0xba, 0x53, 0x37, 0x0a, 0x7a, 0xc5, 0x9f, 0x21, 0xe4, 0x5c, 0x24, 0x41, 0x59,
	0x48, 0xde, 0x62, 0x5b, 0x01, 0x85, 0xb9, 0x72, 0x76, 0x9a, 0xa8, 0x5d, 0xeb, 0xfa, 0xde, 0x9f,
	0xdc, 0xf6, 0xd6, 0xc4, 0x9b, 0x68, 0xa8, 0xdb, 0x02, 0xd3, 0x4e, 0x5e, 0xe1, 0x83, 0x48, 0x66,
	0x19, 0x57, 0x8a, 0x4b, 0x51, 0x98, 0x56, 0xd1, 0x4b, 0xd4, 0x36, 0xb0, 0x5e, 0xf5, 0x3d, 0xae,
	0xc5, 0x6c, 0xca, 0x12, 0x0a, 0x0f, 0x85, 0x34, 0x03, 0xe7, 0x7f, 0x2d, 0xd6, 0xfb, 0x07, 0x73,
	0x07, 0x02, 0x82, 0xc7, 0x9b, 0x29, 0xbd, 0x0c, 0xc8, 0x21, 0xb6, 0x3f, 0x32, 0x9e, 0xa4, 0xe0,
	0xd8, 0x9a, 0xd6, 0x64, 0x84, 0xe2, 0xc3, 0xf5, 0xd9, 0x32, 0x73, 0xea, 0x85, 0xbc, 0xdd, 0x26,
	0x6a, 0xef, 0x75, 0x8f, 0xff, 0xb2, 0xf0, 0xf6, 0x4d, 0x09, 0xea, 0xe5, 0xa8, 0x6d, 0xb4, 0x75,
	0x8c, 0x0f, 0x7e, 0x73, 0x85, 0x60, 0x6c, 0xf7, 0xde, 0x8c, 0x06, 0x17, 0x67, 0x4f, 0x2c, 0xb2,
	0x8f, 0x1f, 0x0d, 0xce, 0x4d, 0x86, 0xfa, 0xef, 0x6e, 0x96, 0x2e, 0xba, 0x5d, 0xba, 0xe8, 0xc7,
	0xd2, 0x45, 0x9f, 0x57, 0xae, 0x75, 0xbb, 0x72, 0xad, 0x6f, 0x2b, 0xd7, 0xfa, 0xd0, 0xdd, 0x5a,
	0x5c, 0x49, 0x99, 0x4a, 0x2e, 0xfd, 0xf5, 0x03, 0x59, 0x9c, 0xf8, 0x97, 0xbf, 0xbc, 0x12, 0x6d,
	0xc4, 0xd8, 0xd6, 0x77, 0xfb, 0xe4, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x1c, 0x5c, 0x77,
	0x4a, 0x03, 0x00, 0x00,
}

func (m *ValidatorExchangeRate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorExchangeRate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorExchangeRate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EpochNumber != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.InternalTokensToSharesRate.Size()
		i -= size
		if _, err := m.InternalTokensToSharesRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InternalExchangeRate != nil {
		{
			size, err := m.InternalExchangeRate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintValidator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Weight != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.Weight))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.DelegationAmt.Size()
		i -= size
		if _, err := m.DelegationAmt.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.CommissionRate != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.CommissionRate))
		i--
		dAtA[i] = 0x20
	}
	if m.Status != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorExchangeRate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.InternalTokensToSharesRate.Size()
	n += 1 + l + sovValidator(uint64(l))
	if m.EpochNumber != 0 {
		n += 1 + sovValidator(uint64(m.EpochNumber))
	}
	return n
}

func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovValidator(uint64(m.Status))
	}
	if m.CommissionRate != 0 {
		n += 1 + sovValidator(uint64(m.CommissionRate))
	}
	l = m.DelegationAmt.Size()
	n += 1 + l + sovValidator(uint64(l))
	if m.Weight != 0 {
		n += 1 + sovValidator(uint64(m.Weight))
	}
	if m.InternalExchangeRate != nil {
		l = m.InternalExchangeRate.Size()
		n += 1 + l + sovValidator(uint64(l))
	}
	return n
}

func sovValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidator(x uint64) (n int) {
	return sovValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorExchangeRate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: ValidatorExchangeRate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorExchangeRate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalTokensToSharesRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InternalTokensToSharesRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Validator_ValidatorStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissionRate", wireType)
			}
			m.CommissionRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommissionRate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationAmt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DelegationAmt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			m.Weight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Weight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalExchangeRate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InternalExchangeRate == nil {
				m.InternalExchangeRate = &ValidatorExchangeRate{}
			}
			if err := m.InternalExchangeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func skipValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
				return 0, ErrInvalidLengthValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValidator = fmt.Errorf("proto: unexpected end of group")
)
