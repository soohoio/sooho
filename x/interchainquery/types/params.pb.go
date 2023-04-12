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
	// host_enabled enables or disables the host submodule.
	HostEnabled bool `protobuf:"varint,2,opt,name=host_enabled,json=hostEnabled,proto3" json:"host_enabled,omitempty" yaml:"host_enabled"`
	// allow_queries defines a list of query paths allowed to be queried on a host
	// chain.
	AllowQueries []string `protobuf:"bytes,3,rep,name=allow_queries,json=allowQueries,proto3" json:"allow_queries,omitempty" yaml:"allow_queries"`
	//define osmosis price query params
	PriceQueryChannelId    string `protobuf:"bytes,4,opt,name=price_query_channel_id,json=priceQueryChannelId,proto3" json:"price_query_channel_id,omitempty"`
	PriceQueryPoolId       string `protobuf:"bytes,5,opt,name=price_query_pool_id,json=priceQueryPoolId,proto3" json:"price_query_pool_id,omitempty"`
	PriceQueryRoutesPoolId string `protobuf:"bytes,6,opt,name=price_query_routes_pool_id,json=priceQueryRoutesPoolId,proto3" json:"price_query_routes_pool_id,omitempty"`
	PriceQueryTokenInDenom string `protobuf:"bytes,7,opt,name=price_query_token_in_denom,json=priceQueryTokenInDenom,proto3" json:"price_query_token_in_denom,omitempty"`
	PriceQueryTokenOut     string `protobuf:"bytes,8,opt,name=price_query_token_out,json=priceQueryTokenOut,proto3" json:"price_query_token_out,omitempty"`
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

func init() {
	proto.RegisterType((*Params)(nil), "stayking.interchainquery.v1.Params")
}

func init() {
	proto.RegisterFile("stayking/interchainquery/v1/params.proto", fileDescriptor_366b88762f648c8e)
}

var fileDescriptor_366b88762f648c8e = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd2, 0xb1, 0x4e, 0xea, 0x50,
	0x1c, 0x06, 0x70, 0x7a, 0xb9, 0x97, 0x0b, 0xe7, 0x72, 0x13, 0x53, 0x10, 0x1b, 0x4c, 0x0a, 0xe9,
	0xd4, 0xc5, 0x36, 0x48, 0xe2, 0x40, 0xe2, 0x82, 0x3a, 0x30, 0x89, 0x8d, 0x93, 0x4b, 0x73, 0x68,
	0x4f, 0xda, 0x13, 0xda, 0xf3, 0xaf, 0xed, 0x29, 0xda, 0xb7, 0x30, 0x3e, 0x95, 0x23, 0xa3, 0x13,
	0x31, 0xf0, 0x06, 0x3c, 0x81, 0xe9, 0x69, 0x04, 0x04, 0xb7, 0x7e, 0xf9, 0xbe, 0x5f, 0x3b, 0xf4,
	0x8f, 0xf4, 0x84, 0xe3, 0x6c, 0x4a, 0x99, 0x67, 0x52, 0xc6, 0x49, 0xec, 0xf8, 0x98, 0xb2, 0xc7,
	0x94, 0xc4, 0x99, 0x39, 0xeb, 0x99, 0x11, 0x8e, 0x71, 0x98, 0x18, 0x51, 0x0c, 0x1c, 0xe4, 0xd3,
	0xaf, 0xa5, 0xb1, 0xb7, 0x34, 0x66, 0xbd, 0x76, 0xd3, 0x03, 0x0f, 0xc4, 0xce, 0xcc, 0x9f, 0x0a,
	0xa2, 0xbd, 0x96, 0x51, 0x65, 0x2c, 0xde, 0x21, 0x0f, 0x50, 0xdd, 0x87, 0x84, 0xdb, 0x84, 0xe1,
	0x49, 0x40, 0x5c, 0xe5, 0x57, 0x57, 0xd2, 0xab, 0xc3, 0x93, 0xf5, 0xa2, 0xd3, 0xc8, 0x70, 0x18,
	0x0c, 0xb4, 0xdd, 0x56, 0xb3, 0xfe, 0xe5, 0xf1, 0xa6, 0x48, 0xf2, 0x25, 0xfa, 0x8f, 0x83, 0x00,
	0x9e, 0xec, 0xfc, 0x73, 0x94, 0x24, 0x4a, 0xb9, 0x5b, 0xd6, 0x6b, 0x43, 0x65, 0xbd, 0xe8, 0x34,
	0x0b, 0xfc, 0xad, 0xd6, 0xac, 0xba, 0xc8, 0x77, 0x45, 0x94, 0xfb, 0xa8, 0x15, 0xc5, 0xd4, 0x21,
	0xa2, 0xcf, 0x6c, 0xc7, 0xc7, 0x8c, 0x91, 0xc0, 0xa6, 0xae, 0xf2, 0xbb, 0x2b, 0xe9, 0x35, 0xab,
	0x21, 0xda, 0x7c, 0x9d, 0x5d, 0x15, 0xdd, 0xc8, 0x95, 0xcf, 0x50, 0x63, 0x17, 0x45, 0x00, 0x42,
	0xfc, 0x11, 0xe2, 0x68, 0x2b, 0xc6, 0x00, 0xf9, 0x7c, 0x80, 0xda, 0xbb, 0xf3, 0x18, 0x52, 0x4e,
	0x92, 0x8d, 0xaa, 0x08, 0xd5, 0xda, 0x2a, 0x4b, 0xf4, 0x3f, 0x5b, 0x0e, 0x53, 0xc2, 0x6c, 0xca,
	0x6c, 0x97, 0x30, 0x08, 0x95, 0xbf, 0xfb, 0xf6, 0x3e, 0xef, 0x47, 0xec, 0x3a, 0x6f, 0xe5, 0x1e,
	0x3a, 0x3e, 0xb4, 0x90, 0x72, 0xa5, 0x2a, 0x98, 0xbc, 0xc7, 0x6e, 0x53, 0x3e, 0x1c, 0xbf, 0x2d,
	0x55, 0x69, 0xbe, 0x54, 0xa5, 0x8f, 0xa5, 0x2a, 0xbd, 0xac, 0xd4, 0xd2, 0x7c, 0xa5, 0x96, 0xde,
	0x57, 0x6a, 0xe9, 0xe1, 0xc2, 0xa3, 0xdc, 0x4f, 0x27, 0x86, 0x03, 0xa1, 0x99, 0x00, 0xf8, 0x40,
	0xc1, 0xdc, 0x9c, 0xc7, 0xec, 0xdc, 0x7c, 0x3e, 0xb8, 0x11, 0x9e, 0x45, 0x24, 0x99, 0x54, 0xc4,
	0xdf, 0xee, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x27, 0x1c, 0x8b, 0x4c, 0x02, 0x00, 0x00,
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
