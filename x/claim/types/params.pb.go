// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/claim/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the claim module's parameters.
type Params struct {
	Airdrops []*Airdrop `protobuf:"bytes,1,rep,name=airdrops,proto3" json:"airdrops,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7ac871d3875dc3, []int{0}
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

func (m *Params) GetAirdrops() []*Airdrop {
	if m != nil {
		return m.Airdrops
	}
	return nil
}

type Airdrop struct {
	AirdropIdentifier string `protobuf:"bytes,1,opt,name=airdrop_identifier,json=airdropIdentifier,proto3" json:"airdrop_identifier,omitempty" yaml:"airdrop_identifier"`
	// seconds
	AirdropStartTime time.Time `protobuf:"bytes,2,opt,name=airdrop_start_time,json=airdropStartTime,proto3,stdtime" json:"airdrop_start_time" yaml:"airdrop_start_time"`
	// seconds
	AirdropDuration time.Duration `protobuf:"bytes,3,opt,name=airdrop_duration,json=airdropDuration,proto3,stdduration" json:"airdrop_duration,omitempty" yaml:"airdrop_duration"`
	// denom of claimable asset
	ClaimDenom string `protobuf:"bytes,4,opt,name=claim_denom,json=claimDenom,proto3" json:"claim_denom,omitempty"`
	// airdrop distribution account
	DistributorAddress string `protobuf:"bytes,5,opt,name=distributor_address,json=distributorAddress,proto3" json:"distributor_address,omitempty"`
	// ustrd tokens claimed so far in the current period
	ClaimedSoFar github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=claimed_so_far,json=claimedSoFar,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"claimed_so_far"`
}

func (m *Airdrop) Reset()         { *m = Airdrop{} }
func (m *Airdrop) String() string { return proto.CompactTextString(m) }
func (*Airdrop) ProtoMessage()    {}
func (*Airdrop) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7ac871d3875dc3, []int{1}
}
func (m *Airdrop) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Airdrop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Airdrop.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Airdrop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Airdrop.Merge(m, src)
}
func (m *Airdrop) XXX_Size() int {
	return m.Size()
}
func (m *Airdrop) XXX_DiscardUnknown() {
	xxx_messageInfo_Airdrop.DiscardUnknown(m)
}

var xxx_messageInfo_Airdrop proto.InternalMessageInfo

func (m *Airdrop) GetAirdropIdentifier() string {
	if m != nil {
		return m.AirdropIdentifier
	}
	return ""
}

func (m *Airdrop) GetAirdropStartTime() time.Time {
	if m != nil {
		return m.AirdropStartTime
	}
	return time.Time{}
}

func (m *Airdrop) GetAirdropDuration() time.Duration {
	if m != nil {
		return m.AirdropDuration
	}
	return 0
}

func (m *Airdrop) GetClaimDenom() string {
	if m != nil {
		return m.ClaimDenom
	}
	return ""
}

func (m *Airdrop) GetDistributorAddress() string {
	if m != nil {
		return m.DistributorAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "stride.claim.Params")
	proto.RegisterType((*Airdrop)(nil), "stride.claim.Airdrop")
}

func init() { proto.RegisterFile("stride/claim/params.proto", fileDescriptor_dd7ac871d3875dc3) }

var fileDescriptor_dd7ac871d3875dc3 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xbf, 0x6f, 0x13, 0x31,
	0x18, 0xcd, 0x51, 0x08, 0xe0, 0x54, 0xfc, 0x30, 0x20, 0x2e, 0x91, 0xb8, 0x8b, 0x4e, 0x02, 0x05,
	0x09, 0x6c, 0x51, 0x36, 0x98, 0x12, 0x2a, 0xa4, 0x4a, 0x0c, 0x28, 0xed, 0xc4, 0x72, 0x72, 0x62,
	0xe7, 0x6a, 0x35, 0xce, 0x77, 0xb2, 0x1d, 0x89, 0xfc, 0x05, 0xac, 0x1d, 0xf9, 0x93, 0x3a, 0x76,
	0x44, 0x0c, 0x01, 0x25, 0x1b, 0x63, 0x47, 0x26, 0x64, 0x9f, 0xaf, 0x0d, 0xcd, 0x74, 0x77, 0xef,
	0x3d, 0xbf, 0xf7, 0xbe, 0xcf, 0x87, 0xda, 0xc6, 0x6a, 0xc9, 0x05, 0x1d, 0x4f, 0x99, 0x54, 0xb4,
	0x64, 0x9a, 0x29, 0x43, 0x4a, 0x0d, 0x16, 0xf0, 0x6e, 0x45, 0x11, 0x4f, 0x75, 0x1e, 0x17, 0x50,
	0x80, 0x27, 0xa8, 0x7b, 0xab, 0x34, 0x9d, 0xa4, 0x00, 0x28, 0xa6, 0x82, 0xfa, 0xaf, 0xd1, 0x7c,
	0x42, 0xf9, 0x5c, 0x33, 0x2b, 0x61, 0x16, 0xf8, 0xf4, 0x3a, 0x6f, 0xa5, 0x12, 0xc6, 0x32, 0x55,
	0x56, 0x82, 0xec, 0x3d, 0x6a, 0x7e, 0xf6, 0xa1, 0xf8, 0x0d, 0xba, 0xc3, 0xa4, 0xe6, 0x1a, 0x4a,
	0x13, 0x47, 0xdd, 0x9d, 0x5e, 0x6b, 0xef, 0x09, 0xd9, 0x6c, 0x40, 0xfa, 0x15, 0x3b, 0xbc, 0x94,
	0x65, 0x7f, 0x77, 0xd0, 0xed, 0x80, 0xe2, 0x4f, 0x08, 0x07, 0x3c, 0x97, 0x5c, 0xcc, 0xac, 0x9c,
	0x48, 0xa1, 0xe3, 0xa8, 0x1b, 0xf5, 0xee, 0x0e, 0x9e, 0x5d, 0x2c, 0xd3, 0xf6, 0x82, 0xa9, 0xe9,
	0xbb, 0x6c, 0x5b, 0x93, 0x0d, 0x1f, 0x06, 0xf0, 0xe0, 0x12, 0xc3, 0x70, 0xe5, 0x66, 0x2c, 0xd3,
	0x36, 0x77, 0xbd, 0xe3, 0x1b, 0xdd, 0xa8, 0xd7, 0xda, 0xeb, 0x90, 0x6a, 0x28, 0x52, 0x0f, 0x45,
	0x8e, 0xea, 0xa1, 0x06, 0xcf, 0xcf, 0x96, 0x69, 0x63, 0x3b, 0xed, 0xca, 0x23, 0x3b, 0xfd, 0x95,
	0x46, 0xc3, 0x07, 0x81, 0x38, 0x74, 0xb8, 0x3b, 0x8d, 0xbf, 0x45, 0xa8, 0x06, 0xf3, 0x7a, 0x87,
	0xf1, 0x8e, 0xcf, 0x6b, 0x6f, 0xe5, 0xed, 0x07, 0xc1, 0xa0, 0xef, 0xe2, 0xfe, 0x2c, 0xd3, 0xce,
	0xf5, 0xa3, 0xaf, 0x40, 0x49, 0x2b, 0x54, 0x69, 0x17, 0x17, 0xcb, 0xf4, 0xe9, 0xff, 0x65, 0x6a,
	0x4d, 0xf6, 0xdd, 0x55, 0xb9, 0x1f, 0xe0, 0xda, 0x13, 0xa7, 0xa8, 0xe5, 0xf7, 0x9d, 0x73, 0x31,
	0x03, 0x15, 0xdf, 0x74, 0x1b, 0x1c, 0x22, 0x0f, 0xed, 0x3b, 0x04, 0x53, 0xf4, 0x88, 0x4b, 0x77,
	0x33, 0xa3, 0xb9, 0x05, 0x9d, 0x33, 0xce, 0xb5, 0x30, 0x26, 0xbe, 0xe5, 0x85, 0x78, 0x83, 0xea,
	0x57, 0x0c, 0x3e, 0x42, 0xf7, 0xfc, 0x71, 0xc1, 0x73, 0x03, 0xf9, 0x84, 0xe9, 0xb8, 0xe9, 0xaf,
	0x85, 0xb8, 0xf6, 0x3f, 0x97, 0xe9, 0x8b, 0x42, 0xda, 0xe3, 0xf9, 0x88, 0x8c, 0x41, 0xd1, 0x31,
	0x18, 0x05, 0x26, 0x3c, 0x5e, 0x1b, 0x7e, 0x42, 0xed, 0xa2, 0x14, 0x86, 0x1c, 0xcc, 0xec, 0x70,
	0x37, 0xb8, 0x1c, 0xc2, 0x47, 0xa6, 0x07, 0x1f, 0xce, 0x56, 0x49, 0x74, 0xbe, 0x4a, 0xa2, 0xdf,
	0xab, 0x24, 0x3a, 0x5d, 0x27, 0x8d, 0xf3, 0x75, 0xd2, 0xf8, 0xb1, 0x4e, 0x1a, 0x5f, 0x5e, 0x6e,
	0xf8, 0x19, 0x80, 0x63, 0x90, 0x40, 0x8d, 0x65, 0x8b, 0x13, 0x39, 0x2b, 0xe8, 0xd7, 0xf0, 0xab,
	0x7b, 0xdb, 0x51, 0xd3, 0xef, 0xf4, 0xed, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x25, 0x20, 0x48,
	0x55, 0x07, 0x03, 0x00, 0x00,
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
	if len(m.Airdrops) > 0 {
		for iNdEx := len(m.Airdrops) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Airdrops[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Airdrop) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Airdrop) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Airdrop) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ClaimedSoFar.Size()
		i -= size
		if _, err := m.ClaimedSoFar.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.DistributorAddress) > 0 {
		i -= len(m.DistributorAddress)
		copy(dAtA[i:], m.DistributorAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DistributorAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ClaimDenom) > 0 {
		i -= len(m.ClaimDenom)
		copy(dAtA[i:], m.ClaimDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ClaimDenom)))
		i--
		dAtA[i] = 0x22
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.AirdropDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.AirdropDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.AirdropStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintParams(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.AirdropIdentifier) > 0 {
		i -= len(m.AirdropIdentifier)
		copy(dAtA[i:], m.AirdropIdentifier)
		i = encodeVarintParams(dAtA, i, uint64(len(m.AirdropIdentifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
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
	if len(m.Airdrops) > 0 {
		for _, e := range m.Airdrops {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func (m *Airdrop) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AirdropIdentifier)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.AirdropDuration)
	n += 1 + l + sovParams(uint64(l))
	l = len(m.ClaimDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DistributorAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.ClaimedSoFar.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
				return fmt.Errorf("proto: wrong wireType = %d for field Airdrops", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Airdrops = append(m.Airdrops, &Airdrop{})
			if err := m.Airdrops[len(m.Airdrops)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Airdrop) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Airdrop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Airdrop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AirdropIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.AirdropStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.AirdropDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistributorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedSoFar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClaimedSoFar.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
