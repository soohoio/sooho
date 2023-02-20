// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stayking/stakeibc/params.proto

package types

import (
	fmt "fmt"
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

// Params defines the parameters for the module.
// next id: 18
type Params struct {
	// define epoch lengths, in stayking_epochs
	RewardsInterval                  uint64            `protobuf:"varint,1,opt,name=rewards_interval,json=rewardsInterval,proto3" json:"rewards_interval,omitempty"`
	DelegateInterval                 uint64            `protobuf:"varint,6,opt,name=delegate_interval,json=delegateInterval,proto3" json:"delegate_interval,omitempty"`
	DepositInterval                  uint64            `protobuf:"varint,2,opt,name=deposit_interval,json=depositInterval,proto3" json:"deposit_interval,omitempty"`
	RedemptionRateInterval           uint64            `protobuf:"varint,3,opt,name=redemption_rate_interval,json=redemptionRateInterval,proto3" json:"redemption_rate_interval,omitempty"`
	StaykingCommission               uint64            `protobuf:"varint,4,opt,name=stayking_commission,json=staykingCommission,proto3" json:"stayking_commission,omitempty"`
	ZoneComAddress                   map[string]string `protobuf:"bytes,5,rep,name=zone_com_address,json=zoneComAddress,proto3" json:"zone_com_address,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ReinvestInterval                 uint64            `protobuf:"varint,7,opt,name=reinvest_interval,json=reinvestInterval,proto3" json:"reinvest_interval,omitempty"`
	ValidatorRebalancingThreshold    uint64            `protobuf:"varint,8,opt,name=validator_rebalancing_threshold,json=validatorRebalancingThreshold,proto3" json:"validator_rebalancing_threshold,omitempty"`
	IcaTimeoutNanos                  uint64            `protobuf:"varint,9,opt,name=ica_timeout_nanos,json=icaTimeoutNanos,proto3" json:"ica_timeout_nanos,omitempty"`
	BufferSize                       uint64            `protobuf:"varint,10,opt,name=buffer_size,json=bufferSize,proto3" json:"buffer_size,omitempty"`
	IbcTimeoutBlocks                 uint64            `protobuf:"varint,11,opt,name=ibc_timeout_blocks,json=ibcTimeoutBlocks,proto3" json:"ibc_timeout_blocks,omitempty"`
	FeeTransferTimeoutNanos          uint64            `protobuf:"varint,12,opt,name=fee_transfer_timeout_nanos,json=feeTransferTimeoutNanos,proto3" json:"fee_transfer_timeout_nanos,omitempty"`
	MaxStakeIcaCallsPerEpoch         uint64            `protobuf:"varint,13,opt,name=max_stake_ica_calls_per_epoch,json=maxStakeIcaCallsPerEpoch,proto3" json:"max_stake_ica_calls_per_epoch,omitempty"`
	SafetyMinRedemptionRateThreshold uint64            `protobuf:"varint,14,opt,name=safety_min_redemption_rate_threshold,json=safetyMinRedemptionRateThreshold,proto3" json:"safety_min_redemption_rate_threshold,omitempty"`
	SafetyMaxRedemptionRateThreshold uint64            `protobuf:"varint,15,opt,name=safety_max_redemption_rate_threshold,json=safetyMaxRedemptionRateThreshold,proto3" json:"safety_max_redemption_rate_threshold,omitempty"`
	IbcTransferTimeoutNanos          uint64            `protobuf:"varint,16,opt,name=ibc_transfer_timeout_nanos,json=ibcTransferTimeoutNanos,proto3" json:"ibc_transfer_timeout_nanos,omitempty"`
	SafetyNumValidators              uint64            `protobuf:"varint,17,opt,name=safety_num_validators,json=safetyNumValidators,proto3" json:"safety_num_validators,omitempty"`
	SafetyMaxSlashPercent            uint64            `protobuf:"varint,18,opt,name=safety_max_slash_percent,json=safetyMaxSlashPercent,proto3" json:"safety_max_slash_percent,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_424da2190fda9dcc, []int{0}
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

func (m *Params) GetRewardsInterval() uint64 {
	if m != nil {
		return m.RewardsInterval
	}
	return 0
}

func (m *Params) GetDelegateInterval() uint64 {
	if m != nil {
		return m.DelegateInterval
	}
	return 0
}

func (m *Params) GetDepositInterval() uint64 {
	if m != nil {
		return m.DepositInterval
	}
	return 0
}

func (m *Params) GetRedemptionRateInterval() uint64 {
	if m != nil {
		return m.RedemptionRateInterval
	}
	return 0
}

func (m *Params) GetStaykingCommission() uint64 {
	if m != nil {
		return m.StaykingCommission
	}
	return 0
}

func (m *Params) GetZoneComAddress() map[string]string {
	if m != nil {
		return m.ZoneComAddress
	}
	return nil
}

func (m *Params) GetReinvestInterval() uint64 {
	if m != nil {
		return m.ReinvestInterval
	}
	return 0
}

func (m *Params) GetValidatorRebalancingThreshold() uint64 {
	if m != nil {
		return m.ValidatorRebalancingThreshold
	}
	return 0
}

func (m *Params) GetIcaTimeoutNanos() uint64 {
	if m != nil {
		return m.IcaTimeoutNanos
	}
	return 0
}

func (m *Params) GetBufferSize() uint64 {
	if m != nil {
		return m.BufferSize
	}
	return 0
}

func (m *Params) GetIbcTimeoutBlocks() uint64 {
	if m != nil {
		return m.IbcTimeoutBlocks
	}
	return 0
}

func (m *Params) GetFeeTransferTimeoutNanos() uint64 {
	if m != nil {
		return m.FeeTransferTimeoutNanos
	}
	return 0
}

func (m *Params) GetMaxStakeIcaCallsPerEpoch() uint64 {
	if m != nil {
		return m.MaxStakeIcaCallsPerEpoch
	}
	return 0
}

func (m *Params) GetSafetyMinRedemptionRateThreshold() uint64 {
	if m != nil {
		return m.SafetyMinRedemptionRateThreshold
	}
	return 0
}

func (m *Params) GetSafetyMaxRedemptionRateThreshold() uint64 {
	if m != nil {
		return m.SafetyMaxRedemptionRateThreshold
	}
	return 0
}

func (m *Params) GetIbcTransferTimeoutNanos() uint64 {
	if m != nil {
		return m.IbcTransferTimeoutNanos
	}
	return 0
}

func (m *Params) GetSafetyNumValidators() uint64 {
	if m != nil {
		return m.SafetyNumValidators
	}
	return 0
}

func (m *Params) GetSafetyMaxSlashPercent() uint64 {
	if m != nil {
		return m.SafetyMaxSlashPercent
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "stayking.stakeibc.Params")
	proto.RegisterMapType((map[string]string)(nil), "stayking.stakeibc.Params.ZoneComAddressEntry")
}

func init() { proto.RegisterFile("stayking/stakeibc/params.proto", fileDescriptor_424da2190fda9dcc) }

var fileDescriptor_424da2190fda9dcc = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xcf, 0x4e, 0x1b, 0x31,
	0x10, 0xc6, 0x13, 0xfe, 0x15, 0x4c, 0x0b, 0x89, 0x81, 0x76, 0x15, 0x89, 0x80, 0xaa, 0x1e, 0xa0,
	0xa5, 0x89, 0x44, 0x0f, 0x45, 0x70, 0xa8, 0x00, 0x81, 0xc4, 0xa1, 0x08, 0x2d, 0xa8, 0x95, 0xb8,
	0x58, 0xde, 0xdd, 0x49, 0x62, 0x65, 0xd7, 0x5e, 0xd9, 0x4e, 0x9a, 0xe4, 0x29, 0x7a, 0xec, 0xb1,
	0x8f, 0xd3, 0x23, 0xc7, 0x1e, 0x2b, 0x78, 0x91, 0xca, 0x76, 0x76, 0x37, 0x69, 0x69, 0x6f, 0xce,
	0x7c, 0xbf, 0xf9, 0x32, 0xf3, 0xc9, 0x6b, 0x54, 0x57, 0x9a, 0x0e, 0xbb, 0x8c, 0xb7, 0x9b, 0x4a,
	0xd3, 0x2e, 0xb0, 0x20, 0x6c, 0xa6, 0x54, 0xd2, 0x44, 0x35, 0x52, 0x29, 0xb4, 0xc0, 0xd5, 0x4c,
	0x6f, 0x64, 0x7a, 0x6d, 0xbd, 0x2d, 0xda, 0xc2, 0xaa, 0x4d, 0x73, 0x72, 0xe0, 0xcb, 0xbb, 0x45,
	0xb4, 0x70, 0x65, 0x3b, 0xf1, 0x2e, 0xaa, 0x48, 0xf8, 0x42, 0x65, 0xa4, 0x08, 0xe3, 0x1a, 0x64,
	0x9f, 0xc6, 0x5e, 0x79, 0xbb, 0xbc, 0x33, 0xe7, 0xaf, 0x8e, 0xeb, 0x17, 0xe3, 0x32, 0x7e, 0x83,
	0xaa, 0x11, 0xc4, 0xd0, 0xa6, 0x1a, 0x0a, 0x76, 0xc1, 0xb2, 0x95, 0x4c, 0xc8, 0xe1, 0x5d, 0x54,
	0x89, 0x20, 0x15, 0x8a, 0xe9, 0x82, 0x9d, 0x71, 0xbe, 0xe3, 0x7a, 0x8e, 0x1e, 0x20, 0x4f, 0x42,
	0x04, 0x49, 0xaa, 0x99, 0xe0, 0x44, 0x4e, 0xd9, 0xcf, 0xda, 0x96, 0xe7, 0x85, 0xee, 0x4f, 0xfe,
	0x49, 0x13, 0xad, 0x65, 0x2b, 0x93, 0x50, 0x24, 0x09, 0x53, 0x8a, 0x09, 0xee, 0xcd, 0xd9, 0x26,
	0x9c, 0x49, 0xa7, 0xb9, 0x82, 0x3f, 0xa3, 0xca, 0x48, 0x70, 0x30, 0x30, 0xa1, 0x51, 0x24, 0x41,
	0x29, 0x6f, 0x7e, 0x7b, 0x76, 0x67, 0x79, 0xff, 0x6d, 0xe3, 0xaf, 0xf0, 0x1a, 0x2e, 0xa2, 0xc6,
	0xad, 0xe0, 0x70, 0x2a, 0x92, 0x63, 0xc7, 0x9f, 0x71, 0x2d, 0x87, 0xfe, 0xca, 0x68, 0xaa, 0x68,
	0xb2, 0x91, 0xc0, 0x78, 0x1f, 0xd4, 0xc4, 0xbe, 0x4f, 0x5c, 0x36, 0x99, 0x90, 0x8f, 0x7d, 0x8e,
	0xb6, 0xfa, 0x34, 0x66, 0x11, 0xd5, 0x42, 0x12, 0x09, 0x01, 0x8d, 0x29, 0x0f, 0xcd, 0x0e, 0xba,
	0x23, 0x41, 0x75, 0x44, 0x1c, 0x79, 0x8b, 0xb6, 0x75, 0x33, 0xc7, 0xfc, 0x82, 0xba, 0xc9, 0x20,
	0xfc, 0x1a, 0x55, 0x59, 0x48, 0x89, 0x66, 0x09, 0x88, 0x9e, 0x26, 0x9c, 0x72, 0xa1, 0xbc, 0x25,
	0x17, 0x32, 0x0b, 0xe9, 0x8d, 0xab, 0x5f, 0x9a, 0x32, 0xde, 0x42, 0xcb, 0x41, 0xaf, 0xd5, 0x02,
	0x49, 0x14, 0x1b, 0x81, 0x87, 0x2c, 0x85, 0x5c, 0xe9, 0x9a, 0x8d, 0x00, 0xef, 0x21, 0xcc, 0x82,
	0x30, 0x37, 0x0b, 0x62, 0x11, 0x76, 0x95, 0xb7, 0xec, 0x56, 0x60, 0x41, 0x38, 0x76, 0x3b, 0xb1,
	0x75, 0x7c, 0x84, 0x6a, 0x2d, 0x00, 0xa2, 0x25, 0xe5, 0xca, 0x98, 0x4e, 0xcf, 0xf0, 0xd4, 0x76,
	0xbd, 0x68, 0x01, 0xdc, 0x8c, 0x81, 0xa9, 0x59, 0x3e, 0xa0, 0xcd, 0x84, 0x0e, 0x88, 0xcd, 0x99,
	0x98, 0x0d, 0x42, 0x1a, 0xc7, 0x8a, 0xa4, 0x20, 0x09, 0xa4, 0x22, 0xec, 0x78, 0xcf, 0x6c, 0xbf,
	0x97, 0xd0, 0xc1, 0xb5, 0x61, 0x2e, 0x42, 0x7a, 0x6a, 0x88, 0x2b, 0x90, 0x67, 0x46, 0xc7, 0x97,
	0xe8, 0x95, 0xa2, 0x2d, 0xd0, 0x43, 0x92, 0x30, 0x4e, 0xfe, 0xbc, 0x3c, 0x45, 0x8a, 0x2b, 0xd6,
	0x67, 0xdb, 0xb1, 0x1f, 0x19, 0xf7, 0xa7, 0xae, 0x51, 0x11, 0xe4, 0x84, 0x1f, 0x1d, 0xfc, 0xc7,
	0x6f, 0x75, 0xca, 0x8f, 0x0e, 0xfe, 0xe5, 0x77, 0x84, 0x6a, 0x36, 0xcb, 0xc7, 0xd3, 0xa9, 0xb8,
	0x74, 0x4c, 0xa6, 0x8f, 0xa5, 0xb3, 0x8f, 0x36, 0xc6, 0xc3, 0xf0, 0x5e, 0x42, 0xf2, 0x1b, 0xa0,
	0xbc, 0xaa, 0xed, 0x5b, 0x73, 0xe2, 0x65, 0x2f, 0xf9, 0x94, 0x4b, 0xf8, 0x3d, 0xf2, 0x26, 0x16,
	0x50, 0x31, 0x55, 0x1d, 0x13, 0x67, 0x08, 0x5c, 0x7b, 0xd8, 0xb6, 0x6d, 0xe4, 0x43, 0x5f, 0x1b,
	0xf5, 0xca, 0x89, 0xb5, 0x63, 0xb4, 0xf6, 0xc8, 0xf5, 0xc6, 0x15, 0x34, 0xdb, 0x85, 0xa1, 0x7d,
	0x08, 0x96, 0x7c, 0x73, 0xc4, 0xeb, 0x68, 0xbe, 0x4f, 0xe3, 0x1e, 0xd8, 0x8f, 0x78, 0xc9, 0x77,
	0x3f, 0x0e, 0x67, 0x0e, 0xca, 0x87, 0x73, 0xdf, 0xbe, 0x6f, 0x95, 0x4e, 0xce, 0x7f, 0xdc, 0xd7,
	0xcb, 0x77, 0xf7, 0xf5, 0xf2, 0xaf, 0xfb, 0x7a, 0xf9, 0xeb, 0x43, 0xbd, 0x74, 0xf7, 0x50, 0x2f,
	0xfd, 0x7c, 0xa8, 0x97, 0x6e, 0xf7, 0xda, 0x4c, 0x77, 0x7a, 0x41, 0x23, 0x14, 0x49, 0x53, 0x09,
	0xd1, 0x11, 0x4c, 0x34, 0xf3, 0x87, 0x6c, 0x50, 0x3c, 0x65, 0x7a, 0x98, 0x82, 0x0a, 0x16, 0xec,
	0x0b, 0xf5, 0xee, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0xe3, 0x6a, 0xa6, 0xec, 0x04, 0x00,
	0x00,
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
	if m.SafetyMaxSlashPercent != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SafetyMaxSlashPercent))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x90
	}
	if m.SafetyNumValidators != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SafetyNumValidators))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.IbcTransferTimeoutNanos != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.IbcTransferTimeoutNanos))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.SafetyMaxRedemptionRateThreshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SafetyMaxRedemptionRateThreshold))
		i--
		dAtA[i] = 0x78
	}
	if m.SafetyMinRedemptionRateThreshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SafetyMinRedemptionRateThreshold))
		i--
		dAtA[i] = 0x70
	}
	if m.MaxStakeIcaCallsPerEpoch != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxStakeIcaCallsPerEpoch))
		i--
		dAtA[i] = 0x68
	}
	if m.FeeTransferTimeoutNanos != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.FeeTransferTimeoutNanos))
		i--
		dAtA[i] = 0x60
	}
	if m.IbcTimeoutBlocks != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.IbcTimeoutBlocks))
		i--
		dAtA[i] = 0x58
	}
	if m.BufferSize != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BufferSize))
		i--
		dAtA[i] = 0x50
	}
	if m.IcaTimeoutNanos != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.IcaTimeoutNanos))
		i--
		dAtA[i] = 0x48
	}
	if m.ValidatorRebalancingThreshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ValidatorRebalancingThreshold))
		i--
		dAtA[i] = 0x40
	}
	if m.ReinvestInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ReinvestInterval))
		i--
		dAtA[i] = 0x38
	}
	if m.DelegateInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DelegateInterval))
		i--
		dAtA[i] = 0x30
	}
	if len(m.ZoneComAddress) > 0 {
		for k := range m.ZoneComAddress {
			v := m.ZoneComAddress[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintParams(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintParams(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintParams(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.StaykingCommission != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.StaykingCommission))
		i--
		dAtA[i] = 0x20
	}
	if m.RedemptionRateInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RedemptionRateInterval))
		i--
		dAtA[i] = 0x18
	}
	if m.DepositInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DepositInterval))
		i--
		dAtA[i] = 0x10
	}
	if m.RewardsInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RewardsInterval))
		i--
		dAtA[i] = 0x8
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
	if m.RewardsInterval != 0 {
		n += 1 + sovParams(uint64(m.RewardsInterval))
	}
	if m.DepositInterval != 0 {
		n += 1 + sovParams(uint64(m.DepositInterval))
	}
	if m.RedemptionRateInterval != 0 {
		n += 1 + sovParams(uint64(m.RedemptionRateInterval))
	}
	if m.StaykingCommission != 0 {
		n += 1 + sovParams(uint64(m.StaykingCommission))
	}
	if len(m.ZoneComAddress) > 0 {
		for k, v := range m.ZoneComAddress {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovParams(uint64(len(k))) + 1 + len(v) + sovParams(uint64(len(v)))
			n += mapEntrySize + 1 + sovParams(uint64(mapEntrySize))
		}
	}
	if m.DelegateInterval != 0 {
		n += 1 + sovParams(uint64(m.DelegateInterval))
	}
	if m.ReinvestInterval != 0 {
		n += 1 + sovParams(uint64(m.ReinvestInterval))
	}
	if m.ValidatorRebalancingThreshold != 0 {
		n += 1 + sovParams(uint64(m.ValidatorRebalancingThreshold))
	}
	if m.IcaTimeoutNanos != 0 {
		n += 1 + sovParams(uint64(m.IcaTimeoutNanos))
	}
	if m.BufferSize != 0 {
		n += 1 + sovParams(uint64(m.BufferSize))
	}
	if m.IbcTimeoutBlocks != 0 {
		n += 1 + sovParams(uint64(m.IbcTimeoutBlocks))
	}
	if m.FeeTransferTimeoutNanos != 0 {
		n += 1 + sovParams(uint64(m.FeeTransferTimeoutNanos))
	}
	if m.MaxStakeIcaCallsPerEpoch != 0 {
		n += 1 + sovParams(uint64(m.MaxStakeIcaCallsPerEpoch))
	}
	if m.SafetyMinRedemptionRateThreshold != 0 {
		n += 1 + sovParams(uint64(m.SafetyMinRedemptionRateThreshold))
	}
	if m.SafetyMaxRedemptionRateThreshold != 0 {
		n += 1 + sovParams(uint64(m.SafetyMaxRedemptionRateThreshold))
	}
	if m.IbcTransferTimeoutNanos != 0 {
		n += 2 + sovParams(uint64(m.IbcTransferTimeoutNanos))
	}
	if m.SafetyNumValidators != 0 {
		n += 2 + sovParams(uint64(m.SafetyNumValidators))
	}
	if m.SafetyMaxSlashPercent != 0 {
		n += 2 + sovParams(uint64(m.SafetyMaxSlashPercent))
	}
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardsInterval", wireType)
			}
			m.RewardsInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RewardsInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositInterval", wireType)
			}
			m.DepositInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DepositInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedemptionRateInterval", wireType)
			}
			m.RedemptionRateInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RedemptionRateInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StaykingCommission", wireType)
			}
			m.StaykingCommission = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StaykingCommission |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZoneComAddress", wireType)
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
			if m.ZoneComAddress == nil {
				m.ZoneComAddress = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthParams
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthParams
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthParams
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthParams
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipParams(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthParams
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ZoneComAddress[mapkey] = mapvalue
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegateInterval", wireType)
			}
			m.DelegateInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DelegateInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReinvestInterval", wireType)
			}
			m.ReinvestInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReinvestInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorRebalancingThreshold", wireType)
			}
			m.ValidatorRebalancingThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidatorRebalancingThreshold |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IcaTimeoutNanos", wireType)
			}
			m.IcaTimeoutNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IcaTimeoutNanos |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufferSize", wireType)
			}
			m.BufferSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BufferSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcTimeoutBlocks", wireType)
			}
			m.IbcTimeoutBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IbcTimeoutBlocks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeTransferTimeoutNanos", wireType)
			}
			m.FeeTransferTimeoutNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FeeTransferTimeoutNanos |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxStakeIcaCallsPerEpoch", wireType)
			}
			m.MaxStakeIcaCallsPerEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxStakeIcaCallsPerEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SafetyMinRedemptionRateThreshold", wireType)
			}
			m.SafetyMinRedemptionRateThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SafetyMinRedemptionRateThreshold |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SafetyMaxRedemptionRateThreshold", wireType)
			}
			m.SafetyMaxRedemptionRateThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SafetyMaxRedemptionRateThreshold |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcTransferTimeoutNanos", wireType)
			}
			m.IbcTransferTimeoutNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IbcTransferTimeoutNanos |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SafetyNumValidators", wireType)
			}
			m.SafetyNumValidators = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SafetyNumValidators |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 18:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SafetyMaxSlashPercent", wireType)
			}
			m.SafetyMaxSlashPercent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SafetyMaxSlashPercent |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
