// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stayking/levstakeibc/host_zone.proto

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

type HostZone struct {
	ChainId            string                                 `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ConnectionId       string                                 `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	Bech32Prefix       string                                 `protobuf:"bytes,3,opt,name=bech32prefix,proto3" json:"bech32prefix,omitempty"`
	TransferChannelId  string                                 `protobuf:"bytes,4,opt,name=transfer_channel_id,json=transferChannelId,proto3" json:"transfer_channel_id,omitempty"`
	DelegationAccount  *ICAAccount                            `protobuf:"bytes,5,opt,name=delegation_account,json=delegationAccount,proto3" json:"delegation_account,omitempty"`
	WithdrawalAccount  *ICAAccount                            `protobuf:"bytes,6,opt,name=withdrawal_account,json=withdrawalAccount,proto3" json:"withdrawal_account,omitempty"`
	RedemptionAccount  *ICAAccount                            `protobuf:"bytes,7,opt,name=redemption_account,json=redemptionAccount,proto3" json:"redemption_account,omitempty"`
	FeeAccount         *ICAAccount                            `protobuf:"bytes,8,opt,name=fee_account,json=feeAccount,proto3" json:"fee_account,omitempty"`
	IcqAccount         *ICAAccount                            `protobuf:"bytes,9,opt,name=icq_account,json=icqAccount,proto3" json:"icq_account,omitempty"`
	HostDenom          string                                 `protobuf:"bytes,10,opt,name=host_denom,json=hostDenom,proto3" json:"host_denom,omitempty"`
	IbcDenom           string                                 `protobuf:"bytes,11,opt,name=ibc_denom,json=ibcDenom,proto3" json:"ibc_denom,omitempty"`
	LastRedemptionRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,12,opt,name=last_redemption_rate,json=lastRedemptionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"last_redemption_rate"`
	RedemptionRate     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,13,opt,name=redemption_rate,json=redemptionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"redemption_rate"`
	StakedBal          github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,14,opt,name=staked_bal,json=stakedBal,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"staked_bal"`
	Address            string                                 `protobuf:"bytes,15,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	UnbondingFrequency uint64                                 `protobuf:"varint,16,opt,name=unbonding_frequency,json=unbondingFrequency,proto3" json:"unbonding_frequency,omitempty"`
}

func (m *HostZone) Reset()         { *m = HostZone{} }
func (m *HostZone) String() string { return proto.CompactTextString(m) }
func (*HostZone) ProtoMessage()    {}
func (*HostZone) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c0f9ffee77eaaa0, []int{0}
}
func (m *HostZone) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HostZone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HostZone.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HostZone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostZone.Merge(m, src)
}
func (m *HostZone) XXX_Size() int {
	return m.Size()
}
func (m *HostZone) XXX_DiscardUnknown() {
	xxx_messageInfo_HostZone.DiscardUnknown(m)
}

var xxx_messageInfo_HostZone proto.InternalMessageInfo

func (m *HostZone) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *HostZone) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *HostZone) GetBech32Prefix() string {
	if m != nil {
		return m.Bech32Prefix
	}
	return ""
}

func (m *HostZone) GetTransferChannelId() string {
	if m != nil {
		return m.TransferChannelId
	}
	return ""
}

func (m *HostZone) GetDelegationAccount() *ICAAccount {
	if m != nil {
		return m.DelegationAccount
	}
	return nil
}

func (m *HostZone) GetWithdrawalAccount() *ICAAccount {
	if m != nil {
		return m.WithdrawalAccount
	}
	return nil
}

func (m *HostZone) GetRedemptionAccount() *ICAAccount {
	if m != nil {
		return m.RedemptionAccount
	}
	return nil
}

func (m *HostZone) GetFeeAccount() *ICAAccount {
	if m != nil {
		return m.FeeAccount
	}
	return nil
}

func (m *HostZone) GetIcqAccount() *ICAAccount {
	if m != nil {
		return m.IcqAccount
	}
	return nil
}

func (m *HostZone) GetHostDenom() string {
	if m != nil {
		return m.HostDenom
	}
	return ""
}

func (m *HostZone) GetIbcDenom() string {
	if m != nil {
		return m.IbcDenom
	}
	return ""
}

func (m *HostZone) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *HostZone) GetUnbondingFrequency() uint64 {
	if m != nil {
		return m.UnbondingFrequency
	}
	return 0
}

func init() {
	proto.RegisterType((*HostZone)(nil), "stayking.levstakeibc.HostZone")
}

func init() {
	proto.RegisterFile("stayking/levstakeibc/host_zone.proto", fileDescriptor_4c0f9ffee77eaaa0)
}

var fileDescriptor_4c0f9ffee77eaaa0 = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xbb, 0x6e, 0xdb, 0x3e,
	0x14, 0xc6, 0xad, 0xff, 0x3f, 0x8d, 0x6d, 0xe6, 0xd6, 0x30, 0x19, 0x94, 0x14, 0x55, 0x0c, 0xb7,
	0x08, 0x32, 0x34, 0x12, 0xe0, 0x6c, 0x45, 0x97, 0x5c, 0x50, 0xd4, 0x40, 0x8b, 0x02, 0x1a, 0xb3,
	0x08, 0x14, 0x79, 0x2c, 0x11, 0x91, 0x49, 0x47, 0xa4, 0x93, 0xb8, 0x4f, 0xd1, 0x87, 0xe9, 0x43,
	0x64, 0x0c, 0x3a, 0x05, 0x1d, 0x82, 0x22, 0x7e, 0x83, 0x3e, 0x41, 0x21, 0xea, 0x62, 0xb9, 0xc8,
	0xe0, 0x02, 0x9d, 0x4c, 0x9e, 0xef, 0x3b, 0xbf, 0x8f, 0xa6, 0x88, 0x83, 0x5e, 0x2b, 0x4d, 0x26,
	0x17, 0x5c, 0x44, 0x5e, 0x02, 0x57, 0x4a, 0x93, 0x0b, 0xe0, 0x21, 0xf5, 0x62, 0xa9, 0x74, 0xf0,
	0x45, 0x0a, 0x70, 0x47, 0xa9, 0xd4, 0x12, 0x6f, 0x97, 0x2e, 0xb7, 0xe6, 0xda, 0xdd, 0x7f, 0xb2,
	0x97, 0x53, 0x12, 0x10, 0x4a, 0xe5, 0x58, 0xe8, 0xbc, 0x7b, 0x77, 0x3b, 0x92, 0x91, 0x34, 0x4b,
	0x2f, 0x5b, 0x15, 0xd5, 0x1d, 0x2a, 0xd5, 0x50, 0xaa, 0x20, 0x17, 0xf2, 0x4d, 0x2e, 0x75, 0xef,
	0x9b, 0xa8, 0xf5, 0x41, 0x2a, 0x7d, 0x2e, 0x05, 0xe0, 0x1d, 0xd4, 0xa2, 0x31, 0xe1, 0x22, 0xe0,
	0xcc, 0xb6, 0x3a, 0xd6, 0x41, 0xdb, 0x6f, 0x9a, 0x7d, 0x9f, 0xe1, 0x57, 0x68, 0x8d, 0x4a, 0x21,
	0x80, 0x6a, 0x2e, 0x8d, 0xfe, 0x9f, 0xd1, 0x57, 0x67, 0xc5, 0x3e, 0xc3, 0x5d, 0xb4, 0x1a, 0x02,
	0x8d, 0x8f, 0x7a, 0xa3, 0x14, 0x06, 0xfc, 0xc6, 0xfe, 0x3f, 0xf7, 0xd4, 0x6b, 0xd8, 0x45, 0x5b,
	0x3a, 0x25, 0x42, 0x0d, 0x20, 0x0d, 0x68, 0x4c, 0x84, 0x80, 0x24, 0xc3, 0x2d, 0x19, 0xeb, 0x66,
	0x29, 0x9d, 0xe6, 0x4a, 0x9f, 0xe1, 0xcf, 0x08, 0x33, 0x48, 0x20, 0x22, 0x26, 0xb8, 0xf8, 0xb7,
	0xf6, 0xb3, 0x8e, 0x75, 0xb0, 0xd2, 0xeb, 0xb8, 0x4f, 0x5d, 0x96, 0xdb, 0x3f, 0x3d, 0x3e, 0xce,
	0x7d, 0xfe, 0xe6, 0xac, 0xb7, 0x28, 0x65, 0xc0, 0x6b, 0xae, 0x63, 0x96, 0x92, 0x6b, 0x92, 0x54,
	0xc0, 0xe5, 0x45, 0x81, 0xb3, 0xde, 0x1a, 0x30, 0x05, 0x06, 0xc3, 0xd1, 0xdc, 0x09, 0x9b, 0x8b,
	0x02, 0x67, 0xbd, 0x25, 0xf0, 0x18, 0xad, 0x0c, 0x00, 0x2a, 0x52, 0x6b, 0x41, 0x12, 0x1a, 0x00,
	0xd4, 0x10, 0x9c, 0x5e, 0x56, 0x88, 0xf6, 0xa2, 0x08, 0x4e, 0x2f, 0x4b, 0xc4, 0x4b, 0x84, 0xcc,
	0xdb, 0x64, 0x20, 0xe4, 0xd0, 0x46, 0xe6, 0xfb, 0xb4, 0xb3, 0xca, 0x59, 0x56, 0xc0, 0x2f, 0x50,
	0x9b, 0x87, 0xb4, 0x50, 0x57, 0x8c, 0xda, 0xe2, 0x21, 0xcd, 0x45, 0x81, 0xb6, 0x13, 0xa2, 0x74,
	0x50, 0xbb, 0x97, 0x94, 0x68, 0xb0, 0x57, 0x33, 0xdf, 0xc9, 0xbb, 0xdb, 0x87, 0xbd, 0xc6, 0x8f,
	0x87, 0xbd, 0xfd, 0x88, 0xeb, 0x78, 0x1c, 0xba, 0x54, 0x0e, 0x8b, 0x47, 0x59, 0xfc, 0x1c, 0x2a,
	0x76, 0xe1, 0xe9, 0xc9, 0x08, 0x94, 0x7b, 0x06, 0xf4, 0xfb, 0xb7, 0x43, 0x54, 0xbc, 0xd9, 0x33,
	0xa0, 0x3e, 0xce, 0xc8, 0x7e, 0x05, 0xf6, 0x89, 0x06, 0x0c, 0x68, 0xe3, 0xcf, 0xa8, 0xb5, 0x7f,
	0x10, 0xb5, 0x9e, 0xce, 0xc7, 0x7c, 0x42, 0xc8, 0xdc, 0x1a, 0x0b, 0x42, 0x92, 0xd8, 0xeb, 0x26,
	0xc1, 0xfd, 0x8b, 0x84, 0xbe, 0xd0, 0x7e, 0x3b, 0x27, 0x9c, 0x90, 0x04, 0xbf, 0x41, 0x4d, 0xc2,
	0x58, 0x0a, 0x4a, 0xd9, 0x1b, 0x86, 0x85, 0x7f, 0x3d, 0xec, 0xad, 0x4f, 0xc8, 0x30, 0x79, 0xdb,
	0x2d, 0x84, 0xae, 0x5f, 0x5a, 0xb0, 0x87, 0xb6, 0xc6, 0x22, 0x94, 0x82, 0x71, 0x11, 0x05, 0x83,
	0x14, 0x2e, 0xc7, 0x20, 0xe8, 0xc4, 0x7e, 0xde, 0xb1, 0x0e, 0x96, 0x7c, 0x5c, 0x49, 0xef, 0x4b,
	0xe5, 0xe4, 0xe3, 0xed, 0xa3, 0x63, 0xdd, 0x3d, 0x3a, 0xd6, 0xcf, 0x47, 0xc7, 0xfa, 0x3a, 0x75,
	0x1a, 0x77, 0x53, 0xa7, 0x71, 0x3f, 0x75, 0x1a, 0xe7, 0xbd, 0xda, 0x59, 0x95, 0x94, 0xb1, 0xe4,
	0xd2, 0xab, 0x06, 0xcc, 0x55, 0xcf, 0xbb, 0x99, 0x9b, 0x32, 0xe6, 0xec, 0xe1, 0xb2, 0x99, 0x17,
	0x47, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x25, 0xea, 0xeb, 0xc6, 0x04, 0x00, 0x00,
}

func (m *HostZone) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostZone) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HostZone) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UnbondingFrequency != 0 {
		i = encodeVarintHostZone(dAtA, i, uint64(m.UnbondingFrequency))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintHostZone(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x7a
	}
	{
		size := m.StakedBal.Size()
		i -= size
		if _, err := m.StakedBal.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintHostZone(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x72
	{
		size := m.RedemptionRate.Size()
		i -= size
		if _, err := m.RedemptionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintHostZone(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size := m.LastRedemptionRate.Size()
		i -= size
		if _, err := m.LastRedemptionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintHostZone(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if len(m.IbcDenom) > 0 {
		i -= len(m.IbcDenom)
		copy(dAtA[i:], m.IbcDenom)
		i = encodeVarintHostZone(dAtA, i, uint64(len(m.IbcDenom)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.HostDenom) > 0 {
		i -= len(m.HostDenom)
		copy(dAtA[i:], m.HostDenom)
		i = encodeVarintHostZone(dAtA, i, uint64(len(m.HostDenom)))
		i--
		dAtA[i] = 0x52
	}
	if m.IcqAccount != nil {
		{
			size, err := m.IcqAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHostZone(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.FeeAccount != nil {
		{
			size, err := m.FeeAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHostZone(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.RedemptionAccount != nil {
		{
			size, err := m.RedemptionAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHostZone(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.WithdrawalAccount != nil {
		{
			size, err := m.WithdrawalAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHostZone(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.DelegationAccount != nil {
		{
			size, err := m.DelegationAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHostZone(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TransferChannelId) > 0 {
		i -= len(m.TransferChannelId)
		copy(dAtA[i:], m.TransferChannelId)
		i = encodeVarintHostZone(dAtA, i, uint64(len(m.TransferChannelId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Bech32Prefix) > 0 {
		i -= len(m.Bech32Prefix)
		copy(dAtA[i:], m.Bech32Prefix)
		i = encodeVarintHostZone(dAtA, i, uint64(len(m.Bech32Prefix)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintHostZone(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintHostZone(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHostZone(dAtA []byte, offset int, v uint64) int {
	offset -= sovHostZone(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HostZone) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovHostZone(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovHostZone(uint64(l))
	}
	l = len(m.Bech32Prefix)
	if l > 0 {
		n += 1 + l + sovHostZone(uint64(l))
	}
	l = len(m.TransferChannelId)
	if l > 0 {
		n += 1 + l + sovHostZone(uint64(l))
	}
	if m.DelegationAccount != nil {
		l = m.DelegationAccount.Size()
		n += 1 + l + sovHostZone(uint64(l))
	}
	if m.WithdrawalAccount != nil {
		l = m.WithdrawalAccount.Size()
		n += 1 + l + sovHostZone(uint64(l))
	}
	if m.RedemptionAccount != nil {
		l = m.RedemptionAccount.Size()
		n += 1 + l + sovHostZone(uint64(l))
	}
	if m.FeeAccount != nil {
		l = m.FeeAccount.Size()
		n += 1 + l + sovHostZone(uint64(l))
	}
	if m.IcqAccount != nil {
		l = m.IcqAccount.Size()
		n += 1 + l + sovHostZone(uint64(l))
	}
	l = len(m.HostDenom)
	if l > 0 {
		n += 1 + l + sovHostZone(uint64(l))
	}
	l = len(m.IbcDenom)
	if l > 0 {
		n += 1 + l + sovHostZone(uint64(l))
	}
	l = m.LastRedemptionRate.Size()
	n += 1 + l + sovHostZone(uint64(l))
	l = m.RedemptionRate.Size()
	n += 1 + l + sovHostZone(uint64(l))
	l = m.StakedBal.Size()
	n += 1 + l + sovHostZone(uint64(l))
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovHostZone(uint64(l))
	}
	if m.UnbondingFrequency != 0 {
		n += 2 + sovHostZone(uint64(m.UnbondingFrequency))
	}
	return n
}

func sovHostZone(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHostZone(x uint64) (n int) {
	return sovHostZone(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HostZone) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHostZone
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
			return fmt.Errorf("proto: HostZone: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostZone: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bech32Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bech32Prefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransferChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DelegationAccount == nil {
				m.DelegationAccount = &ICAAccount{}
			}
			if err := m.DelegationAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawalAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.WithdrawalAccount == nil {
				m.WithdrawalAccount = &ICAAccount{}
			}
			if err := m.WithdrawalAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedemptionAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RedemptionAccount == nil {
				m.RedemptionAccount = &ICAAccount{}
			}
			if err := m.RedemptionAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FeeAccount == nil {
				m.FeeAccount = &ICAAccount{}
			}
			if err := m.FeeAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IcqAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IcqAccount == nil {
				m.IcqAccount = &ICAAccount{}
			}
			if err := m.IcqAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastRedemptionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastRedemptionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedemptionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RedemptionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakedBal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StakedBal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingFrequency", wireType)
			}
			m.UnbondingFrequency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondingFrequency |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHostZone(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHostZone
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
func skipHostZone(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHostZone
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
					return 0, ErrIntOverflowHostZone
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
					return 0, ErrIntOverflowHostZone
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
				return 0, ErrInvalidLengthHostZone
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHostZone
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHostZone
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHostZone        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHostZone          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHostZone = fmt.Errorf("proto: unexpected end of group")
)