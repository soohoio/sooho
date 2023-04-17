// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stayking/interchainquery/v1/params.proto

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
type Params struct {
	HostEnabled bool `protobuf:"varint,2,opt,name=host_enabled,json=hostEnabled,proto3" json:"host_enabled,omitempty" yaml:"host_enabled"`
	// allow_queries defines a list of query paths allowed to be queried on a host chain.
	AllowQueries []string `protobuf:"bytes,3,rep,name=allow_queries,json=allowQueries,proto3" json:"allow_queries,omitempty" yaml:"allow_queries"`
	//define osmosis price query params
	PriceQueryChannelId    string `protobuf:"bytes,4,opt,name=price_query_channel_id,json=priceQueryChannelId,proto3" json:"price_query_channel_id,omitempty"`
	PriceQueryPoolId       string `protobuf:"bytes,5,opt,name=price_query_pool_id,json=priceQueryPoolId,proto3" json:"price_query_pool_id,omitempty"`
	PriceQueryRoutesPoolId string `protobuf:"bytes,6,opt,name=price_query_routes_pool_id,json=priceQueryRoutesPoolId,proto3" json:"price_query_routes_pool_id,omitempty"`
	PriceQueryTokenInDenom string `protobuf:"bytes,7,opt,name=price_query_token_in_denom,json=priceQueryTokenInDenom,proto3" json:"price_query_token_in_denom,omitempty"`
	PriceQueryTokenOut     string `protobuf:"bytes,8,opt,name=price_query_token_out,json=priceQueryTokenOut,proto3" json:"price_query_token_out,omitempty"`
	PriceQueryPath         string `protobuf:"bytes,9,opt,name=price_query_path,json=priceQueryPath,proto3" json:"price_query_path,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_366b88762f648c8e, []int{0}
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

func (m *Params) GetHostEnabled() bool {
	if m != nil {
		return m.HostEnabled
	}
	return false
}

func (m *Params) GetAllowQueries() []string {
	if m != nil {
		return m.AllowQueries
	}
	return nil
}

func (m *Params) GetPriceQueryChannelId() string {
	if m != nil {
		return m.PriceQueryChannelId
	}
	return ""
}

func (m *Params) GetPriceQueryPoolId() string {
	if m != nil {
		return m.PriceQueryPoolId
	}
	return ""
}

func (m *Params) GetPriceQueryRoutesPoolId() string {
	if m != nil {
		return m.PriceQueryRoutesPoolId
	}
	return ""
}

func (m *Params) GetPriceQueryTokenInDenom() string {
	if m != nil {
		return m.PriceQueryTokenInDenom
	}
	return ""
}

func (m *Params) GetPriceQueryTokenOut() string {
	if m != nil {
		return m.PriceQueryTokenOut
	}
	return ""
}

func (m *Params) GetPriceQueryPath() string {
	if m != nil {
		return m.PriceQueryPath
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "stayking.interchainquery.v1.Params")
}

func init() {
	proto.RegisterFile("stayking/interchainquery/v1/params.proto", fileDescriptor_366b88762f648c8e)
}

var fileDescriptor_366b88762f648c8e = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x8e, 0x9a, 0x40,
	0x18, 0x80, 0xa5, 0xb6, 0x56, 0xa7, 0xb6, 0x31, 0xa3, 0xb5, 0xc4, 0x26, 0x68, 0x38, 0x71, 0x29,
	0xc4, 0x9a, 0xf4, 0x60, 0xd2, 0x8b, 0x6d, 0x0f, 0x9e, 0x6a, 0x49, 0x4f, 0xbd, 0x90, 0x11, 0x26,
	0x30, 0x11, 0xe6, 0xa7, 0x30, 0xd8, 0xf2, 0x16, 0xfb, 0x40, 0xfb, 0x00, 0x7b, 0xf4, 0xb8, 0x27,
	0xb3, 0xd1, 0x37, 0xf0, 0x09, 0x36, 0x0c, 0x59, 0x45, 0xdd, 0x1b, 0x7f, 0xbe, 0xef, 0x63, 0xf2,
	0x27, 0x3f, 0x32, 0x52, 0x41, 0xf2, 0x15, 0xe3, 0xbe, 0xc5, 0xb8, 0xa0, 0x89, 0x1b, 0x10, 0xc6,
	0xff, 0x66, 0x34, 0xc9, 0xad, 0xf5, 0xd8, 0x8a, 0x49, 0x42, 0xa2, 0xd4, 0x8c, 0x13, 0x10, 0x80,
	0x3f, 0x3e, 0x99, 0xe6, 0x85, 0x69, 0xae, 0xc7, 0x83, 0x9e, 0x0f, 0x3e, 0x48, 0xcf, 0x2a, 0xbe,
	0xca, 0x44, 0xbf, 0xad, 0xa3, 0xc6, 0x42, 0xfe, 0x03, 0x4f, 0x51, 0x3b, 0x80, 0x54, 0x38, 0x94,
	0x93, 0x65, 0x48, 0x3d, 0xf5, 0xc5, 0x48, 0x31, 0x9a, 0xb3, 0x0f, 0x87, 0xed, 0xb0, 0x9b, 0x93,
	0x28, 0x9c, 0xea, 0x55, 0xaa, 0xdb, 0x6f, 0x8a, 0xf1, 0x47, 0x39, 0xe1, 0xaf, 0xe8, 0x2d, 0x09,
	0x43, 0xf8, 0xe7, 0x14, 0xcf, 0x31, 0x9a, 0xaa, 0xf5, 0x51, 0xdd, 0x68, 0xcd, 0xd4, 0xc3, 0x76,
	0xd8, 0x2b, 0xe3, 0x33, 0xac, 0xdb, 0x6d, 0x39, 0xff, 0x2a, 0x47, 0x3c, 0x41, 0xfd, 0x38, 0x61,
	0x2e, 0x95, 0x3c, 0x77, 0xdc, 0x80, 0x70, 0x4e, 0x43, 0x87, 0x79, 0xea, 0xcb, 0x91, 0x62, 0xb4,
	0xec, 0xae, 0xa4, 0x85, 0x9d, 0x7f, 0x2b, 0xd9, 0xdc, 0xc3, 0x9f, 0x50, 0xb7, 0x1a, 0xc5, 0x00,
	0xb2, 0x78, 0x25, 0x8b, 0xce, 0xa9, 0x58, 0x00, 0x14, 0xfa, 0x14, 0x0d, 0xaa, 0x7a, 0x02, 0x99,
	0xa0, 0xe9, 0xb1, 0x6a, 0xc8, 0xaa, 0x7f, 0xaa, 0x6c, 0xc9, 0x9f, 0x6f, 0x05, 0xac, 0x28, 0x77,
	0x18, 0x77, 0x3c, 0xca, 0x21, 0x52, 0x5f, 0x5f, 0xb6, 0xbf, 0x0b, 0x3e, 0xe7, 0xdf, 0x0b, 0x8a,
	0xc7, 0xe8, 0xfd, 0x75, 0x0b, 0x99, 0x50, 0x9b, 0x32, 0xc3, 0x17, 0xd9, 0xcf, 0x4c, 0x60, 0x03,
	0x75, 0xce, 0x36, 0x23, 0x22, 0x50, 0x5b, 0xd2, 0x7e, 0x57, 0x59, 0x8b, 0x88, 0x60, 0xb6, 0xb8,
	0xdb, 0x69, 0xca, 0x66, 0xa7, 0x29, 0x0f, 0x3b, 0x4d, 0xb9, 0xd9, 0x6b, 0xb5, 0xcd, 0x5e, 0xab,
	0xdd, 0xef, 0xb5, 0xda, 0x9f, 0x2f, 0x3e, 0x13, 0x41, 0xb6, 0x34, 0x5d, 0x88, 0xac, 0x14, 0x20,
	0x00, 0x06, 0xd6, 0xf1, 0x90, 0xd6, 0x9f, 0xad, 0xff, 0x57, 0xd7, 0x24, 0xf2, 0x98, 0xa6, 0xcb,
	0x86, 0xbc, 0x8b, 0xc9, 0x63, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x70, 0xcf, 0x5e, 0x76, 0x02,
	0x00, 0x00,
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
	if len(m.PriceQueryPath) > 0 {
		i -= len(m.PriceQueryPath)
		copy(dAtA[i:], m.PriceQueryPath)
		i = encodeVarintParams(dAtA, i, uint64(len(m.PriceQueryPath)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.PriceQueryTokenOut) > 0 {
		i -= len(m.PriceQueryTokenOut)
		copy(dAtA[i:], m.PriceQueryTokenOut)
		i = encodeVarintParams(dAtA, i, uint64(len(m.PriceQueryTokenOut)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.PriceQueryTokenInDenom) > 0 {
		i -= len(m.PriceQueryTokenInDenom)
		copy(dAtA[i:], m.PriceQueryTokenInDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.PriceQueryTokenInDenom)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.PriceQueryRoutesPoolId) > 0 {
		i -= len(m.PriceQueryRoutesPoolId)
		copy(dAtA[i:], m.PriceQueryRoutesPoolId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.PriceQueryRoutesPoolId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.PriceQueryPoolId) > 0 {
		i -= len(m.PriceQueryPoolId)
		copy(dAtA[i:], m.PriceQueryPoolId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.PriceQueryPoolId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PriceQueryChannelId) > 0 {
		i -= len(m.PriceQueryChannelId)
		copy(dAtA[i:], m.PriceQueryChannelId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.PriceQueryChannelId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AllowQueries) > 0 {
		for iNdEx := len(m.AllowQueries) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowQueries[iNdEx])
			copy(dAtA[i:], m.AllowQueries[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.AllowQueries[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.HostEnabled {
		i--
		if m.HostEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
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
	if m.HostEnabled {
		n += 2
	}
	if len(m.AllowQueries) > 0 {
		for _, s := range m.AllowQueries {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = len(m.PriceQueryChannelId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.PriceQueryPoolId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.PriceQueryRoutesPoolId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.PriceQueryTokenInDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.PriceQueryTokenOut)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.PriceQueryPath)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.HostEnabled = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowQueries", wireType)
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
			m.AllowQueries = append(m.AllowQueries, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceQueryChannelId", wireType)
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
			m.PriceQueryChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceQueryPoolId", wireType)
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
			m.PriceQueryPoolId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceQueryRoutesPoolId", wireType)
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
			m.PriceQueryRoutesPoolId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceQueryTokenInDenom", wireType)
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
			m.PriceQueryTokenInDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceQueryTokenOut", wireType)
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
			m.PriceQueryTokenOut = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceQueryPath", wireType)
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
			m.PriceQueryPath = string(dAtA[iNdEx:postIndex])
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
