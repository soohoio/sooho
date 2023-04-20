// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stayking/levstakeibc/params.proto

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
	DepositInterval                  uint64            `protobuf:"varint,2,opt,name=deposit_interval,json=depositInterval,proto3" json:"deposit_interval,omitempty"`
	DelegateInterval                 uint64            `protobuf:"varint,3,opt,name=delegate_interval,json=delegateInterval,proto3" json:"delegate_interval,omitempty"`
	ReinvestInterval                 uint64            `protobuf:"varint,4,opt,name=reinvest_interval,json=reinvestInterval,proto3" json:"reinvest_interval,omitempty"`
	RedemptionRateInterval           uint64            `protobuf:"varint,5,opt,name=redemption_rate_interval,json=redemptionRateInterval,proto3" json:"redemption_rate_interval,omitempty"`
	StaykingCommission               uint64            `protobuf:"varint,6,opt,name=stayking_commission,json=staykingCommission,proto3" json:"stayking_commission,omitempty"`
	ZoneComAddress                   map[string]string `protobuf:"bytes,7,rep,name=zone_com_address,json=zoneComAddress,proto3" json:"zone_com_address,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IbcTimeoutBlocks                 uint64            `protobuf:"varint,8,opt,name=ibc_timeout_blocks,json=ibcTimeoutBlocks,proto3" json:"ibc_timeout_blocks,omitempty"`
	IbcTransferTimeoutNanos          uint64            `protobuf:"varint,9,opt,name=ibc_transfer_timeout_nanos,json=ibcTransferTimeoutNanos,proto3" json:"ibc_transfer_timeout_nanos,omitempty"`
	FeeTransferTimeoutNanos          uint64            `protobuf:"varint,10,opt,name=fee_transfer_timeout_nanos,json=feeTransferTimeoutNanos,proto3" json:"fee_transfer_timeout_nanos,omitempty"`
	IcaTimeoutNanos                  uint64            `protobuf:"varint,11,opt,name=ica_timeout_nanos,json=icaTimeoutNanos,proto3" json:"ica_timeout_nanos,omitempty"`
	MaxStakeIcaCallsPerEpoch         uint64            `protobuf:"varint,12,opt,name=max_stake_ica_calls_per_epoch,json=maxStakeIcaCallsPerEpoch,proto3" json:"max_stake_ica_calls_per_epoch,omitempty"`
	ValidatorRebalancingThreshold    uint64            `protobuf:"varint,13,opt,name=validator_rebalancing_threshold,json=validatorRebalancingThreshold,proto3" json:"validator_rebalancing_threshold,omitempty"`
	BufferSize                       uint64            `protobuf:"varint,14,opt,name=buffer_size,json=bufferSize,proto3" json:"buffer_size,omitempty"`
	SafetyMinRedemptionRateThreshold uint64            `protobuf:"varint,15,opt,name=safety_min_redemption_rate_threshold,json=safetyMinRedemptionRateThreshold,proto3" json:"safety_min_redemption_rate_threshold,omitempty"`
	SafetyMaxRedemptionRateThreshold uint64            `protobuf:"varint,16,opt,name=safety_max_redemption_rate_threshold,json=safetyMaxRedemptionRateThreshold,proto3" json:"safety_max_redemption_rate_threshold,omitempty"`
	SafetyNumValidators              uint64            `protobuf:"varint,17,opt,name=safety_num_validators,json=safetyNumValidators,proto3" json:"safety_num_validators,omitempty"`
	SafetyMaxSlashPercent            uint64            `protobuf:"varint,18,opt,name=safety_max_slash_percent,json=safetyMaxSlashPercent,proto3" json:"safety_max_slash_percent,omitempty"`
	SafetyMarkPriceExpirationTime    uint64            `protobuf:"varint,19,opt,name=safety_mark_price_expiration_time,json=safetyMarkPriceExpirationTime,proto3" json:"safety_mark_price_expiration_time,omitempty"`
	LiquidationPerformanceFee        uint64            `protobuf:"varint,20,opt,name=liquidation_performance_fee,json=liquidationPerformanceFee,proto3" json:"liquidation_performance_fee,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_cad617bb0bb2e20d, []int{0}
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

func (m *Params) GetDepositInterval() uint64 {
	if m != nil {
		return m.DepositInterval
	}
	return 0
}

func (m *Params) GetDelegateInterval() uint64 {
	if m != nil {
		return m.DelegateInterval
	}
	return 0
}

func (m *Params) GetReinvestInterval() uint64 {
	if m != nil {
		return m.ReinvestInterval
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

func (m *Params) GetIbcTimeoutBlocks() uint64 {
	if m != nil {
		return m.IbcTimeoutBlocks
	}
	return 0
}

func (m *Params) GetIbcTransferTimeoutNanos() uint64 {
	if m != nil {
		return m.IbcTransferTimeoutNanos
	}
	return 0
}

func (m *Params) GetFeeTransferTimeoutNanos() uint64 {
	if m != nil {
		return m.FeeTransferTimeoutNanos
	}
	return 0
}

func (m *Params) GetIcaTimeoutNanos() uint64 {
	if m != nil {
		return m.IcaTimeoutNanos
	}
	return 0
}

func (m *Params) GetMaxStakeIcaCallsPerEpoch() uint64 {
	if m != nil {
		return m.MaxStakeIcaCallsPerEpoch
	}
	return 0
}

func (m *Params) GetValidatorRebalancingThreshold() uint64 {
	if m != nil {
		return m.ValidatorRebalancingThreshold
	}
	return 0
}

func (m *Params) GetBufferSize() uint64 {
	if m != nil {
		return m.BufferSize
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

func (m *Params) GetSafetyMarkPriceExpirationTime() uint64 {
	if m != nil {
		return m.SafetyMarkPriceExpirationTime
	}
	return 0
}

func (m *Params) GetLiquidationPerformanceFee() uint64 {
	if m != nil {
		return m.LiquidationPerformanceFee
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "stayking.levstakeibc.Params")
	proto.RegisterMapType((map[string]string)(nil), "stayking.levstakeibc.Params.ZoneComAddressEntry")
}

func init() { proto.RegisterFile("stayking/levstakeibc/params.proto", fileDescriptor_cad617bb0bb2e20d) }

var fileDescriptor_cad617bb0bb2e20d = []byte{
	// 719 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x4d, 0x4f, 0x1b, 0x39,
	0x18, 0xc7, 0x13, 0xde, 0x76, 0x63, 0x76, 0x21, 0x38, 0x61, 0x77, 0x96, 0x15, 0x01, 0x56, 0x7b,
	0x60, 0x5f, 0x94, 0x54, 0xe9, 0xa1, 0x88, 0x4a, 0xad, 0x00, 0x81, 0x8a, 0xd4, 0xa2, 0x28, 0xa0,
	0x1e, 0xb8, 0x58, 0x1e, 0xcf, 0x93, 0xc4, 0xca, 0x8c, 0x3d, 0xb5, 0x9d, 0x34, 0xc9, 0xa7, 0xe8,
	0xb1, 0xc7, 0x7e, 0x9c, 0x1e, 0x39, 0xf6, 0x52, 0xa9, 0x82, 0x2f, 0x52, 0xd9, 0x4e, 0x26, 0x09,
	0x82, 0xde, 0x26, 0xcf, 0xff, 0xf7, 0x7b, 0x3c, 0x7e, 0xc6, 0x31, 0xda, 0xd3, 0x86, 0x0e, 0xbb,
	0x5c, 0xb4, 0x6b, 0x31, 0xf4, 0xb5, 0xa1, 0x5d, 0xe0, 0x21, 0xab, 0xa5, 0x54, 0xd1, 0x44, 0x57,
	0x53, 0x25, 0x8d, 0xc4, 0xe5, 0x09, 0x52, 0x9d, 0x41, 0xb6, 0xca, 0x6d, 0xd9, 0x96, 0x0e, 0xa8,
	0xd9, 0x27, 0xcf, 0xfe, 0xf5, 0xb5, 0x80, 0x56, 0x1a, 0x4e, 0xc6, 0xff, 0xa0, 0xa2, 0x82, 0xf7,
	0x54, 0x45, 0x9a, 0x70, 0x61, 0x40, 0xf5, 0x69, 0x1c, 0xe4, 0x77, 0xf3, 0xfb, 0x4b, 0xcd, 0xf5,
	0x71, 0xfd, 0x7c, 0x5c, 0xb6, 0x68, 0x04, 0xa9, 0xd4, 0xdc, 0x4c, 0xd1, 0x05, 0x8f, 0x8e, 0xeb,
	0x19, 0xfa, 0x1f, 0xda, 0x88, 0x20, 0x86, 0x36, 0x35, 0x30, 0x65, 0x17, 0x1d, 0x5b, 0x9c, 0x04,
	0xb3, 0xb0, 0x02, 0x2e, 0xfa, 0xa0, 0x67, 0x1a, 0x2f, 0x79, 0x78, 0x12, 0x64, 0xf0, 0x01, 0x0a,
	0x14, 0x44, 0x90, 0xa4, 0x86, 0x4b, 0x41, 0xd4, 0xdc, 0x02, 0xcb, 0xce, 0xf9, 0x6d, 0x9a, 0x37,
	0x67, 0x97, 0xa9, 0xa1, 0xd2, 0x64, 0x44, 0x84, 0xc9, 0x24, 0xe1, 0x5a, 0x73, 0x29, 0x82, 0x15,
	0x27, 0xe1, 0x49, 0x74, 0x92, 0x25, 0xf8, 0x1a, 0x15, 0x47, 0x52, 0x80, 0x85, 0x09, 0x8d, 0x22,
	0x05, 0x5a, 0x07, 0x3f, 0xed, 0x2e, 0xee, 0xaf, 0xd6, 0x9f, 0x54, 0x1f, 0x1a, 0x76, 0xd5, 0x8f,
	0xb4, 0x7a, 0x2d, 0x05, 0x9c, 0xc8, 0xe4, 0xc8, 0x2b, 0xa7, 0xc2, 0xa8, 0x61, 0x73, 0x6d, 0x34,
	0x57, 0xc4, 0xff, 0x23, 0xcc, 0x43, 0x46, 0x0c, 0x4f, 0x40, 0xf6, 0x0c, 0x09, 0x63, 0xc9, 0xba,
	0x3a, 0xf8, 0xd9, 0x6f, 0x9a, 0x87, 0xec, 0xca, 0x07, 0xc7, 0xae, 0x8e, 0x9f, 0xa3, 0x2d, 0x47,
	0x2b, 0x2a, 0x74, 0x0b, 0x54, 0xa6, 0x09, 0x2a, 0xa4, 0x0e, 0x0a, 0xce, 0xfa, 0xdd, 0x5a, 0x63,
	0x60, 0x6c, 0x5f, 0xd8, 0xd8, 0xca, 0x2d, 0x80, 0xc7, 0x64, 0xe4, 0xe5, 0x16, 0xc0, 0x83, 0xf2,
	0xbf, 0x68, 0x83, 0x33, 0x7a, 0xcf, 0x59, 0xf5, 0x1f, 0x9d, 0x33, 0x3a, 0xc7, 0xbe, 0x44, 0xdb,
	0x09, 0x1d, 0x10, 0x37, 0x0e, 0x62, 0x2d, 0x46, 0xe3, 0x58, 0x93, 0x14, 0x14, 0x81, 0x54, 0xb2,
	0x4e, 0xf0, 0x8b, 0xf3, 0x82, 0x84, 0x0e, 0x2e, 0x2d, 0x73, 0xce, 0xe8, 0x89, 0x25, 0x1a, 0xa0,
	0x4e, 0x6d, 0x8e, 0xcf, 0xd0, 0x4e, 0x9f, 0xc6, 0x3c, 0xa2, 0x46, 0x2a, 0xa2, 0x20, 0xa4, 0x31,
	0x15, 0xcc, 0x7e, 0x2e, 0xd3, 0x51, 0xa0, 0x3b, 0x32, 0x8e, 0x82, 0x5f, 0x5d, 0x8b, 0xed, 0x0c,
	0x6b, 0x4e, 0xa9, 0xab, 0x09, 0x84, 0x77, 0xd0, 0x6a, 0xd8, 0x6b, 0xd9, 0xbd, 0x6a, 0x3e, 0x82,
	0x60, 0xcd, 0x39, 0xc8, 0x97, 0x2e, 0xf9, 0x08, 0xf0, 0x05, 0xfa, 0x5b, 0xd3, 0x16, 0x98, 0x21,
	0x49, 0xb8, 0x20, 0xf7, 0xcf, 0xd3, 0x74, 0xb5, 0x75, 0x67, 0xee, 0x7a, 0xf6, 0x0d, 0x17, 0xcd,
	0xb9, 0x93, 0x35, 0x5d, 0x70, 0xa6, 0x1f, 0x1d, 0xfc, 0xa0, 0x5f, 0x71, 0xae, 0x1f, 0x1d, 0x3c,
	0xd6, 0xaf, 0x8e, 0x36, 0xc7, 0xfd, 0x44, 0x2f, 0x21, 0xd9, 0x66, 0x75, 0xb0, 0xe1, 0x1a, 0x94,
	0x7c, 0x78, 0xd1, 0x4b, 0xde, 0x66, 0x11, 0x7e, 0x86, 0x82, 0x99, 0x77, 0xd0, 0x31, 0xd5, 0x1d,
	0x3b, 0x7a, 0x06, 0xc2, 0x04, 0xd8, 0x69, 0x9b, 0xd9, 0xba, 0x97, 0x36, 0x6d, 0xf8, 0x10, 0xbf,
	0x42, 0x7b, 0x99, 0xa8, 0xba, 0x24, 0x55, 0x9c, 0x01, 0x81, 0x41, 0xca, 0x15, 0x75, 0x7b, 0xb0,
	0xdf, 0x3e, 0x28, 0xf9, 0xb9, 0x4f, 0x3a, 0xa8, 0x6e, 0xc3, 0x62, 0xa7, 0x19, 0x65, 0x0f, 0x02,
	0x7e, 0x81, 0xfe, 0x8c, 0xf9, 0xbb, 0x9e, 0x7d, 0x23, 0x2b, 0xa6, 0xa0, 0x5a, 0x52, 0x25, 0x54,
	0x30, 0x20, 0x2d, 0x80, 0xa0, 0xec, 0x7a, 0xfc, 0x31, 0x83, 0x34, 0xa6, 0xc4, 0x19, 0xc0, 0xd6,
	0x11, 0x2a, 0x3d, 0xf0, 0xdf, 0xc1, 0x45, 0xb4, 0xd8, 0x85, 0xa1, 0xbb, 0x95, 0x0a, 0x4d, 0xfb,
	0x88, 0xcb, 0x68, 0xb9, 0x4f, 0xe3, 0x1e, 0xb8, 0xeb, 0xa7, 0xd0, 0xf4, 0x3f, 0x0e, 0x17, 0x0e,
	0xf2, 0x87, 0x4b, 0x1f, 0x3f, 0xed, 0xe4, 0x8e, 0x5f, 0x7f, 0xbe, 0xad, 0xe4, 0x6f, 0x6e, 0x2b,
	0xf9, 0x6f, 0xb7, 0x95, 0xfc, 0x87, 0xbb, 0x4a, 0xee, 0xe6, 0xae, 0x92, 0xfb, 0x72, 0x57, 0xc9,
	0x5d, 0xd7, 0xdb, 0xdc, 0x74, 0x7a, 0x61, 0x95, 0xc9, 0xa4, 0xa6, 0xa5, 0xec, 0x48, 0x2e, 0x6b,
	0xd9, 0xdd, 0xda, 0xaf, 0xd7, 0x06, 0x73, 0x17, 0xac, 0x19, 0xa6, 0xa0, 0xc3, 0x15, 0x77, 0x69,
	0x3e, 0xfd, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xcd, 0x19, 0xc6, 0x85, 0x05, 0x00, 0x00,
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
	if m.LiquidationPerformanceFee != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.LiquidationPerformanceFee))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa0
	}
	if m.SafetyMarkPriceExpirationTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SafetyMarkPriceExpirationTime))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x98
	}
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
	if m.SafetyMaxRedemptionRateThreshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SafetyMaxRedemptionRateThreshold))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.SafetyMinRedemptionRateThreshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SafetyMinRedemptionRateThreshold))
		i--
		dAtA[i] = 0x78
	}
	if m.BufferSize != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BufferSize))
		i--
		dAtA[i] = 0x70
	}
	if m.ValidatorRebalancingThreshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ValidatorRebalancingThreshold))
		i--
		dAtA[i] = 0x68
	}
	if m.MaxStakeIcaCallsPerEpoch != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxStakeIcaCallsPerEpoch))
		i--
		dAtA[i] = 0x60
	}
	if m.IcaTimeoutNanos != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.IcaTimeoutNanos))
		i--
		dAtA[i] = 0x58
	}
	if m.FeeTransferTimeoutNanos != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.FeeTransferTimeoutNanos))
		i--
		dAtA[i] = 0x50
	}
	if m.IbcTransferTimeoutNanos != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.IbcTransferTimeoutNanos))
		i--
		dAtA[i] = 0x48
	}
	if m.IbcTimeoutBlocks != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.IbcTimeoutBlocks))
		i--
		dAtA[i] = 0x40
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
			dAtA[i] = 0x3a
		}
	}
	if m.StaykingCommission != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.StaykingCommission))
		i--
		dAtA[i] = 0x30
	}
	if m.RedemptionRateInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RedemptionRateInterval))
		i--
		dAtA[i] = 0x28
	}
	if m.ReinvestInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ReinvestInterval))
		i--
		dAtA[i] = 0x20
	}
	if m.DelegateInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DelegateInterval))
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
	if m.DelegateInterval != 0 {
		n += 1 + sovParams(uint64(m.DelegateInterval))
	}
	if m.ReinvestInterval != 0 {
		n += 1 + sovParams(uint64(m.ReinvestInterval))
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
	if m.IbcTimeoutBlocks != 0 {
		n += 1 + sovParams(uint64(m.IbcTimeoutBlocks))
	}
	if m.IbcTransferTimeoutNanos != 0 {
		n += 1 + sovParams(uint64(m.IbcTransferTimeoutNanos))
	}
	if m.FeeTransferTimeoutNanos != 0 {
		n += 1 + sovParams(uint64(m.FeeTransferTimeoutNanos))
	}
	if m.IcaTimeoutNanos != 0 {
		n += 1 + sovParams(uint64(m.IcaTimeoutNanos))
	}
	if m.MaxStakeIcaCallsPerEpoch != 0 {
		n += 1 + sovParams(uint64(m.MaxStakeIcaCallsPerEpoch))
	}
	if m.ValidatorRebalancingThreshold != 0 {
		n += 1 + sovParams(uint64(m.ValidatorRebalancingThreshold))
	}
	if m.BufferSize != 0 {
		n += 1 + sovParams(uint64(m.BufferSize))
	}
	if m.SafetyMinRedemptionRateThreshold != 0 {
		n += 1 + sovParams(uint64(m.SafetyMinRedemptionRateThreshold))
	}
	if m.SafetyMaxRedemptionRateThreshold != 0 {
		n += 2 + sovParams(uint64(m.SafetyMaxRedemptionRateThreshold))
	}
	if m.SafetyNumValidators != 0 {
		n += 2 + sovParams(uint64(m.SafetyNumValidators))
	}
	if m.SafetyMaxSlashPercent != 0 {
		n += 2 + sovParams(uint64(m.SafetyMaxSlashPercent))
	}
	if m.SafetyMarkPriceExpirationTime != 0 {
		n += 2 + sovParams(uint64(m.SafetyMarkPriceExpirationTime))
	}
	if m.LiquidationPerformanceFee != 0 {
		n += 2 + sovParams(uint64(m.LiquidationPerformanceFee))
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
		case 4:
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
		case 5:
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
		case 6:
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
		case 7:
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
		case 8:
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
		case 9:
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
		case 10:
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
		case 11:
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
		case 12:
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
		case 13:
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
		case 14:
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
		case 15:
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
		case 16:
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
		case 19:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SafetyMarkPriceExpirationTime", wireType)
			}
			m.SafetyMarkPriceExpirationTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SafetyMarkPriceExpirationTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationPerformanceFee", wireType)
			}
			m.LiquidationPerformanceFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LiquidationPerformanceFee |= uint64(b&0x7F) << shift
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
