// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stayking/lendingpool/v1/lendingpool.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
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

// Pool is the lending pool.
type Pool struct {
	Id             uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	RemainingCoins github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=remaining_coins,json=remainingCoins,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"remaining_coins"`
	RedemptionRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=redemption_rate,json=redemptionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"redemption_rate"`
	TotalCoins     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=total_coins,json=totalCoins,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"total_coins"`
	InterestModel  *types.Any                             `protobuf:"bytes,5,opt,name=interest_model,json=interestModel,proto3" json:"interest_model,omitempty"`
	Denom          string                                 `protobuf:"bytes,6,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *Pool) Reset()      { *m = Pool{} }
func (*Pool) ProtoMessage() {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b41fc733dd0c43f, []int{0}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

// Params defines the set of params for the lending pool module.
type Params struct {
	// protocol tax rate to take from the lending pool depositors
	ProtocolTaxRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=protocol_tax_rate,json=protocolTaxRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"protocol_tax_rate"`
	// expected blocks per year
	BlocksPerYear uint64 `protobuf:"varint,6,opt,name=blocks_per_year,json=blocksPerYear,proto3" json:"blocks_per_year,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b41fc733dd0c43f, []int{1}
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

func (m *Params) GetBlocksPerYear() uint64 {
	if m != nil {
		return m.BlocksPerYear
	}
	return 0
}

// Loan defines a loan record
type Loan struct {
	Id              uint64                                      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Denom           string                                      `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Borrower        string                                      `protobuf:"bytes,3,opt,name=borrower,proto3" json:"borrower,omitempty"`
	ClientModule    string                                      `protobuf:"bytes,4,opt,name=client_module,json=clientModule,proto3" json:"client_module,omitempty"`
	InitMarkPrice   github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,5,rep,name=init_mark_price,json=initMarkPrice,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"init_mark_price"`
	TotalAssetValue github_com_cosmos_cosmos_sdk_types.Dec      `protobuf:"bytes,6,opt,name=total_asset_value,json=totalAssetValue,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"total_asset_value"`
	BorrowedValue   github_com_cosmos_cosmos_sdk_types.Dec      `protobuf:"bytes,7,opt,name=borrowed_value,json=borrowedValue,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"borrowed_value"`
}

func (m *Loan) Reset()         { *m = Loan{} }
func (m *Loan) String() string { return proto.CompactTextString(m) }
func (*Loan) ProtoMessage()    {}
func (*Loan) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b41fc733dd0c43f, []int{2}
}
func (m *Loan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Loan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Loan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Loan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Loan.Merge(m, src)
}
func (m *Loan) XXX_Size() int {
	return m.Size()
}
func (m *Loan) XXX_DiscardUnknown() {
	xxx_messageInfo_Loan.DiscardUnknown(m)
}

var xxx_messageInfo_Loan proto.InternalMessageInfo

func (m *Loan) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Loan) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Loan) GetBorrower() string {
	if m != nil {
		return m.Borrower
	}
	return ""
}

func (m *Loan) GetClientModule() string {
	if m != nil {
		return m.ClientModule
	}
	return ""
}

func (m *Loan) GetInitMarkPrice() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.InitMarkPrice
	}
	return nil
}

func init() {
	proto.RegisterType((*Pool)(nil), "stayking.lendingpool.v1.Pool")
	proto.RegisterType((*Params)(nil), "stayking.lendingpool.v1.Params")
	proto.RegisterType((*Loan)(nil), "stayking.lendingpool.v1.Loan")
}

func init() {
	proto.RegisterFile("stayking/lendingpool/v1/lendingpool.proto", fileDescriptor_7b41fc733dd0c43f)
}

var fileDescriptor_7b41fc733dd0c43f = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xc1, 0x6b, 0x13, 0x4f,
	0x14, 0xce, 0x26, 0x69, 0x7f, 0xed, 0xf4, 0x97, 0x84, 0x2e, 0x05, 0x63, 0xd1, 0x4d, 0x88, 0x50,
	0x22, 0xd2, 0x5d, 0x92, 0xde, 0x8a, 0x97, 0xc6, 0x5e, 0x0a, 0x16, 0xc3, 0x22, 0x82, 0x82, 0x2c,
	0x93, 0xdd, 0xe7, 0x76, 0xc8, 0xee, 0xbc, 0x30, 0x3b, 0x89, 0xdd, 0xa3, 0x17, 0xf1, 0xe8, 0xd1,
	0x63, 0x8f, 0xe2, 0xb9, 0xff, 0x81, 0x97, 0xe2, 0xa9, 0x78, 0x12, 0x0f, 0x55, 0xd2, 0xff, 0x40,
	0xf0, 0x2e, 0x3b, 0xb3, 0x89, 0xf1, 0xa4, 0x60, 0x3d, 0x25, 0x6f, 0xde, 0xdb, 0xef, 0x7b, 0xef,
	0xfb, 0xde, 0x0c, 0xb9, 0x9d, 0x48, 0x9a, 0x0e, 0x19, 0x0f, 0x9d, 0x08, 0x78, 0xc0, 0x78, 0x38,
	0x42, 0x8c, 0x9c, 0x49, 0x67, 0x31, 0xb4, 0x47, 0x02, 0x25, 0x9a, 0xd7, 0x66, 0xa5, 0xf6, 0x62,
	0x6e, 0xd2, 0xd9, 0xdc, 0x08, 0x31, 0x44, 0x55, 0xe3, 0x64, 0xff, 0x74, 0xf9, 0xe6, 0xf5, 0x10,
	0x31, 0x8c, 0xc0, 0x51, 0xd1, 0x60, 0xfc, 0xcc, 0xa1, 0x3c, 0xcd, 0x53, 0x96, 0x8f, 0x49, 0x8c,
	0x89, 0x33, 0xa0, 0x09, 0x38, 0x93, 0xce, 0x00, 0x24, 0xed, 0x38, 0x3e, 0x32, 0x3e, 0xfb, 0x54,
	0xe7, 0x3d, 0x8d, 0xa9, 0x03, 0x9d, 0x6a, 0xbd, 0x2f, 0x91, 0x72, 0x1f, 0x31, 0x32, 0x6f, 0x92,
	0x22, 0x0b, 0xea, 0x46, 0xd3, 0x68, 0x97, 0x7b, 0x95, 0x6f, 0x17, 0x8d, 0xd5, 0x94, 0xc6, 0xd1,
	0x6e, 0x8b, 0x05, 0x2d, 0xb7, 0xc8, 0x02, 0x13, 0x48, 0x4d, 0x40, 0x4c, 0x19, 0x67, 0x3c, 0xf4,
	0x32, 0xe8, 0xa4, 0x5e, 0x6c, 0x1a, 0xed, 0xd5, 0xde, 0xdd, 0xb3, 0x8b, 0x46, 0xe1, 0xf3, 0x45,
	0x63, 0x2b, 0x64, 0xf2, 0x68, 0x3c, 0xb0, 0x7d, 0x8c, 0x73, 0x86, 0xfc, 0x67, 0x3b, 0x09, 0x86,
	0x8e, 0x4c, 0x47, 0x90, 0xd8, 0xfb, 0xe0, 0x7f, 0x3c, 0xdd, 0x26, 0x79, 0x03, 0xfb, 0xe0, 0xbb,
	0xd5, 0x39, 0xe8, 0xbd, 0x0c, 0x53, 0xd3, 0x04, 0x10, 0x8f, 0x24, 0x43, 0xee, 0x09, 0x2a, 0xa1,
	0x5e, 0xba, 0x1a, 0x9a, 0x19, 0xa8, 0x4b, 0x25, 0x98, 0x4f, 0xc9, 0x9a, 0x44, 0x49, 0xa3, 0x7c,
	0x92, 0xf2, 0x15, 0x50, 0x10, 0x05, 0xa8, 0xa7, 0x78, 0x40, 0xaa, 0x8c, 0x4b, 0x10, 0x90, 0x48,
	0x2f, 0xc6, 0x00, 0xa2, 0xfa, 0x52, 0xd3, 0x68, 0xaf, 0x75, 0x37, 0x6c, 0xed, 0xa1, 0x3d, 0xf3,
	0xd0, 0xde, 0xe3, 0x69, 0xcf, 0xfc, 0x70, 0xba, 0x5d, 0x3d, 0xc8, 0xeb, 0x0f, 0xb3, 0xf2, 0x03,
	0xb7, 0xc2, 0x16, 0x63, 0x73, 0x83, 0x2c, 0x05, 0xc0, 0x31, 0xae, 0x2f, 0x67, 0x9d, 0xba, 0x3a,
	0xd8, 0x5d, 0x79, 0x75, 0xd2, 0x28, 0xbc, 0x39, 0x69, 0x14, 0x5a, 0x27, 0x06, 0x59, 0xee, 0x53,
	0x41, 0xe3, 0xc4, 0x3c, 0x22, 0xeb, 0x0a, 0xdd, 0xc7, 0xc8, 0x93, 0xf4, 0x58, 0x6b, 0x68, 0x5c,
	0xc1, 0x80, 0xb5, 0x19, 0xec, 0x43, 0x7a, 0xac, 0x44, 0xdc, 0x22, 0xb5, 0x41, 0x84, 0xfe, 0x30,
	0xf1, 0x46, 0x20, 0xbc, 0x14, 0xa8, 0x50, 0xed, 0x95, 0xdd, 0x8a, 0x3e, 0xee, 0x83, 0x78, 0x0c,
	0x54, 0xec, 0x96, 0x55, 0x8b, 0xdf, 0x4b, 0xa4, 0x7c, 0x1f, 0x29, 0x37, 0xab, 0x3f, 0x17, 0x4d,
	0x6d, 0xd6, 0x7c, 0xb6, 0xe2, 0xc2, 0x6c, 0xe6, 0x26, 0x59, 0x19, 0xa0, 0x10, 0xf8, 0x1c, 0x84,
	0xde, 0x00, 0x77, 0x1e, 0x9b, 0xb7, 0x48, 0xc5, 0x8f, 0x18, 0x70, 0x25, 0xee, 0x38, 0x02, 0xed,
	0x9f, 0xfb, 0xbf, 0x3e, 0x3c, 0x54, 0x67, 0x66, 0x4a, 0x6a, 0x8c, 0x33, 0xe9, 0xc5, 0x54, 0x0c,
	0xbd, 0x91, 0x60, 0x3e, 0xd4, 0x97, 0x9a, 0xa5, 0xf6, 0x5a, 0xf7, 0x86, 0x9d, 0x0f, 0x95, 0xdd,
	0x16, 0x3b, 0xbf, 0x2d, 0xd9, 0x84, 0x99, 0x77, 0xbd, 0x9d, 0x4c, 0xa3, 0x77, 0x5f, 0x1a, 0x77,
	0xfe, 0x4c, 0x23, 0xe5, 0x77, 0xe6, 0x16, 0x93, 0x87, 0x54, 0x0c, 0xfb, 0x19, 0x8f, 0xf9, 0xd2,
	0x20, 0xeb, 0x7a, 0xbd, 0x68, 0x92, 0x80, 0xf4, 0x26, 0x34, 0x1a, 0x83, 0xd2, 0xe6, 0x77, 0xec,
	0x7f, 0xe9, 0x90, 0x22, 0xdd, 0xcb, 0x38, 0x1f, 0x65, 0x94, 0xe6, 0x0b, 0x83, 0x54, 0x73, 0xd5,
	0x82, 0xbc, 0x8b, 0xff, 0xfe, 0x79, 0x17, 0x95, 0x19, 0xa3, 0xea, 0xa1, 0xd7, 0x7f, 0x3b, 0xb5,
	0x8c, 0xb3, 0xa9, 0x65, 0x9c, 0x4f, 0x2d, 0xe3, 0xeb, 0xd4, 0x32, 0x5e, 0x5f, 0x5a, 0x85, 0xf3,
	0x4b, 0xab, 0xf0, 0xe9, 0xd2, 0x2a, 0x3c, 0xe9, 0x2e, 0x50, 0x24, 0x88, 0x47, 0xc8, 0xd0, 0x99,
	0xbf, 0xa0, 0x93, 0xae, 0x73, 0xfc, 0xcb, 0x33, 0xaa, 0x28, 0x07, 0xcb, 0x6a, 0x11, 0x77, 0x7e,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x03, 0xf1, 0xfb, 0xe6, 0x6b, 0x05, 0x00, 0x00,
}

func (this *Pool) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Pool)
	if !ok {
		that2, ok := that.(Pool)
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
	if this.Id != that1.Id {
		return false
	}
	if !this.RemainingCoins.Equal(that1.RemainingCoins) {
		return false
	}
	if !this.RedemptionRate.Equal(that1.RedemptionRate) {
		return false
	}
	if !this.TotalCoins.Equal(that1.TotalCoins) {
		return false
	}
	if !this.InterestModel.Equal(that1.InterestModel) {
		return false
	}
	if this.Denom != that1.Denom {
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
	if !this.ProtocolTaxRate.Equal(that1.ProtocolTaxRate) {
		return false
	}
	if this.BlocksPerYear != that1.BlocksPerYear {
		return false
	}
	return true
}
func (this *Loan) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Loan)
	if !ok {
		that2, ok := that.(Loan)
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
	if this.Id != that1.Id {
		return false
	}
	if this.Denom != that1.Denom {
		return false
	}
	if this.Borrower != that1.Borrower {
		return false
	}
	if this.ClientModule != that1.ClientModule {
		return false
	}
	if len(this.InitMarkPrice) != len(that1.InitMarkPrice) {
		return false
	}
	for i := range this.InitMarkPrice {
		if !this.InitMarkPrice[i].Equal(&that1.InitMarkPrice[i]) {
			return false
		}
	}
	if !this.TotalAssetValue.Equal(that1.TotalAssetValue) {
		return false
	}
	if !this.BorrowedValue.Equal(that1.BorrowedValue) {
		return false
	}
	return true
}
func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintLendingpool(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x32
	}
	if m.InterestModel != nil {
		{
			size, err := m.InterestModel.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLendingpool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	{
		size := m.TotalCoins.Size()
		i -= size
		if _, err := m.TotalCoins.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLendingpool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.RedemptionRate.Size()
		i -= size
		if _, err := m.RedemptionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLendingpool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.RemainingCoins.Size()
		i -= size
		if _, err := m.RemainingCoins.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLendingpool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Id != 0 {
		i = encodeVarintLendingpool(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
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
	if m.BlocksPerYear != 0 {
		i = encodeVarintLendingpool(dAtA, i, uint64(m.BlocksPerYear))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.ProtocolTaxRate.Size()
		i -= size
		if _, err := m.ProtocolTaxRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLendingpool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Loan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Loan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Loan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BorrowedValue.Size()
		i -= size
		if _, err := m.BorrowedValue.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLendingpool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.TotalAssetValue.Size()
		i -= size
		if _, err := m.TotalAssetValue.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLendingpool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.InitMarkPrice) > 0 {
		for iNdEx := len(m.InitMarkPrice) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitMarkPrice[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLendingpool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ClientModule) > 0 {
		i -= len(m.ClientModule)
		copy(dAtA[i:], m.ClientModule)
		i = encodeVarintLendingpool(dAtA, i, uint64(len(m.ClientModule)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Borrower) > 0 {
		i -= len(m.Borrower)
		copy(dAtA[i:], m.Borrower)
		i = encodeVarintLendingpool(dAtA, i, uint64(len(m.Borrower)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintLendingpool(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintLendingpool(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLendingpool(dAtA []byte, offset int, v uint64) int {
	offset -= sovLendingpool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLendingpool(uint64(m.Id))
	}
	l = m.RemainingCoins.Size()
	n += 1 + l + sovLendingpool(uint64(l))
	l = m.RedemptionRate.Size()
	n += 1 + l + sovLendingpool(uint64(l))
	l = m.TotalCoins.Size()
	n += 1 + l + sovLendingpool(uint64(l))
	if m.InterestModel != nil {
		l = m.InterestModel.Size()
		n += 1 + l + sovLendingpool(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovLendingpool(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ProtocolTaxRate.Size()
	n += 1 + l + sovLendingpool(uint64(l))
	if m.BlocksPerYear != 0 {
		n += 1 + sovLendingpool(uint64(m.BlocksPerYear))
	}
	return n
}

func (m *Loan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLendingpool(uint64(m.Id))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovLendingpool(uint64(l))
	}
	l = len(m.Borrower)
	if l > 0 {
		n += 1 + l + sovLendingpool(uint64(l))
	}
	l = len(m.ClientModule)
	if l > 0 {
		n += 1 + l + sovLendingpool(uint64(l))
	}
	if len(m.InitMarkPrice) > 0 {
		for _, e := range m.InitMarkPrice {
			l = e.Size()
			n += 1 + l + sovLendingpool(uint64(l))
		}
	}
	l = m.TotalAssetValue.Size()
	n += 1 + l + sovLendingpool(uint64(l))
	l = m.BorrowedValue.Size()
	n += 1 + l + sovLendingpool(uint64(l))
	return n
}

func sovLendingpool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLendingpool(x uint64) (n int) {
	return sovLendingpool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLendingpool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingpool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainingCoins", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingpool
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
				return ErrInvalidLengthLendingpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLendingpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RemainingCoins.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedemptionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingpool
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
				return ErrInvalidLengthLendingpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLendingpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RedemptionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalCoins", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingpool
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
				return ErrInvalidLengthLendingpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLendingpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalCoins.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterestModel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingpool
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
				return ErrInvalidLengthLendingpool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLendingpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InterestModel == nil {
				m.InterestModel = &types.Any{}
			}
			if err := m.InterestModel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingpool
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
				return ErrInvalidLengthLendingpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLendingpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLendingpool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLendingpool
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
				return ErrIntOverflowLendingpool
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
				return fmt.Errorf("proto: wrong wireType = %d for field ProtocolTaxRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingpool
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
				return ErrInvalidLengthLendingpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLendingpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ProtocolTaxRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksPerYear", wireType)
			}
			m.BlocksPerYear = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingpool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksPerYear |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLendingpool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLendingpool
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
func (m *Loan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLendingpool
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
			return fmt.Errorf("proto: Loan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Loan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingpool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingpool
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
				return ErrInvalidLengthLendingpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLendingpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Borrower", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingpool
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
				return ErrInvalidLengthLendingpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLendingpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Borrower = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientModule", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingpool
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
				return ErrInvalidLengthLendingpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLendingpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientModule = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitMarkPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingpool
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
				return ErrInvalidLengthLendingpool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLendingpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitMarkPrice = append(m.InitMarkPrice, types1.DecCoin{})
			if err := m.InitMarkPrice[len(m.InitMarkPrice)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalAssetValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingpool
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
				return ErrInvalidLengthLendingpool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLendingpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalAssetValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowedValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingpool
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
				return ErrInvalidLengthLendingpool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLendingpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BorrowedValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLendingpool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLendingpool
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
func skipLendingpool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLendingpool
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
					return 0, ErrIntOverflowLendingpool
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
					return 0, ErrIntOverflowLendingpool
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
				return 0, ErrInvalidLengthLendingpool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLendingpool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLendingpool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLendingpool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLendingpool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLendingpool = fmt.Errorf("proto: unexpected end of group")
)
