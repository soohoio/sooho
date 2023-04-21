// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stayking/interchainquery/v1/estimate_swap_exact_amout_out.proto

package types

import (
	fmt "fmt"
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

//=============================== EstimateSwapExactAmountOut
type EstimateSwapExactAmountOutRequest struct {
	PoolId   uint64               `protobuf:"varint,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	Routes   []SwapAmountOutRoute `protobuf:"bytes,3,rep,name=routes,proto3" json:"routes" yaml:"routes"`
	TokenOut string               `protobuf:"bytes,4,opt,name=token_out,json=tokenOut,proto3" json:"token_out,omitempty" yaml:"token_out"`
}

func (m *EstimateSwapExactAmountOutRequest) Reset()         { *m = EstimateSwapExactAmountOutRequest{} }
func (m *EstimateSwapExactAmountOutRequest) String() string { return proto.CompactTextString(m) }
func (*EstimateSwapExactAmountOutRequest) ProtoMessage()    {}
func (*EstimateSwapExactAmountOutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_721dfe0c88f1a85c, []int{0}
}
func (m *EstimateSwapExactAmountOutRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EstimateSwapExactAmountOutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EstimateSwapExactAmountOutRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EstimateSwapExactAmountOutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstimateSwapExactAmountOutRequest.Merge(m, src)
}
func (m *EstimateSwapExactAmountOutRequest) XXX_Size() int {
	return m.Size()
}
func (m *EstimateSwapExactAmountOutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EstimateSwapExactAmountOutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EstimateSwapExactAmountOutRequest proto.InternalMessageInfo

func (m *EstimateSwapExactAmountOutRequest) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *EstimateSwapExactAmountOutRequest) GetRoutes() []SwapAmountOutRoute {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *EstimateSwapExactAmountOutRequest) GetTokenOut() string {
	if m != nil {
		return m.TokenOut
	}
	return ""
}

type EstimateSwapExactAmountOutResponse struct {
	TokenInAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=token_in_amount,json=tokenInAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"token_in_amount" yaml:"token_in_amount"`
}

func (m *EstimateSwapExactAmountOutResponse) Reset()         { *m = EstimateSwapExactAmountOutResponse{} }
func (m *EstimateSwapExactAmountOutResponse) String() string { return proto.CompactTextString(m) }
func (*EstimateSwapExactAmountOutResponse) ProtoMessage()    {}
func (*EstimateSwapExactAmountOutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_721dfe0c88f1a85c, []int{1}
}
func (m *EstimateSwapExactAmountOutResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EstimateSwapExactAmountOutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EstimateSwapExactAmountOutResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EstimateSwapExactAmountOutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstimateSwapExactAmountOutResponse.Merge(m, src)
}
func (m *EstimateSwapExactAmountOutResponse) XXX_Size() int {
	return m.Size()
}
func (m *EstimateSwapExactAmountOutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EstimateSwapExactAmountOutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EstimateSwapExactAmountOutResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EstimateSwapExactAmountOutRequest)(nil), "stayking.interchainquery.v1.EstimateSwapExactAmountOutRequest")
	proto.RegisterType((*EstimateSwapExactAmountOutResponse)(nil), "stayking.interchainquery.v1.EstimateSwapExactAmountOutResponse")
}

func init() {
	proto.RegisterFile("stayking/interchainquery/v1/estimate_swap_exact_amout_out.proto", fileDescriptor_721dfe0c88f1a85c)
}

var fileDescriptor_721dfe0c88f1a85c = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x77, 0xda, 0x65, 0x6d, 0x23, 0xd5, 0x12, 0xaa, 0x2c, 0x15, 0x92, 0x75, 0x0e, 0xb2,
	0xa0, 0x66, 0xd8, 0x0a, 0x1e, 0xbc, 0x88, 0x0b, 0x05, 0xd7, 0x4b, 0x25, 0xde, 0x3c, 0x18, 0xa6,
	0xc9, 0x90, 0x1d, 0xb6, 0x99, 0x97, 0x66, 0xde, 0x6c, 0x9b, 0x6f, 0xe1, 0xc9, 0xcf, 0xd4, 0x63,
	0x8f, 0xe2, 0x21, 0xc8, 0xee, 0x27, 0x30, 0x9f, 0x40, 0x32, 0x89, 0x6b, 0x51, 0xd8, 0x53, 0x1e,
	0x79, 0xef, 0xfd, 0x7f, 0xff, 0x37, 0x7f, 0xe7, 0xad, 0x46, 0x5e, 0x2e, 0xa4, 0x4a, 0x99, 0x54,
	0x28, 0x8a, 0x78, 0xce, 0xa5, 0xba, 0x34, 0xa2, 0x28, 0xd9, 0x72, 0xc2, 0x84, 0x46, 0x99, 0x71,
	0x14, 0x91, 0xbe, 0xe2, 0x79, 0x24, 0xae, 0x79, 0x8c, 0x11, 0xcf, 0xc0, 0x60, 0x04, 0x06, 0x83,
	0xbc, 0x00, 0x04, 0xf7, 0xc9, 0x1f, 0x81, 0xe0, 0x1f, 0x81, 0x60, 0x39, 0x39, 0x3e, 0x4a, 0x21,
	0x05, 0x3b, 0xc7, 0x9a, 0xaa, 0x5d, 0x39, 0x7e, 0xb1, 0x8d, 0x69, 0x51, 0x05, 0x18, 0x14, 0xed,
	0x34, 0xfd, 0x45, 0x9c, 0xa7, 0xa7, 0x9d, 0x91, 0x4f, 0x57, 0x3c, 0x3f, 0x6d, 0x6c, 0xbc, 0xcb,
	0xc0, 0x28, 0x3c, 0x33, 0x18, 0x8a, 0x4b, 0x23, 0x34, 0xba, 0xcf, 0x9d, 0x7b, 0x39, 0xc0, 0x45,
	0x24, 0x93, 0xe1, 0xce, 0x88, 0x8c, 0xfb, 0x53, 0xb7, 0xae, 0xfc, 0x07, 0x25, 0xcf, 0x2e, 0xde,
	0xd0, 0xae, 0x41, 0xc3, 0x41, 0x53, 0xcd, 0x12, 0xf7, 0x8b, 0x33, 0xb0, 0x04, 0x3d, 0xdc, 0x1d,
	0xed, 0x8e, 0xef, 0x9f, 0xb0, 0x60, 0xcb, 0x11, 0x41, 0x03, 0xfd, 0xcb, 0x6b, 0xf6, 0xa6, 0x8f,
	0x6e, 0x2a, 0xbf, 0x57, 0x57, 0xfe, 0x41, 0x0b, 0x68, 0xc5, 0x68, 0xd8, 0xa9, 0xba, 0x13, 0x67,
	0x1f, 0x61, 0x21, 0x54, 0xf3, 0x4c, 0xc3, 0xfe, 0x88, 0x8c, 0xf7, 0xa7, 0x47, 0x75, 0xe5, 0x1f,
	0xb6, 0xd3, 0x9b, 0x16, 0x0d, 0xf7, 0x6c, 0x7d, 0x66, 0xf0, 0x43, 0x7f, 0x8f, 0x1c, 0xee, 0x84,
	0x03, 0x2d, 0x54, 0x22, 0x0a, 0xfa, 0x8d, 0x38, 0x74, 0xdb, 0xcd, 0x3a, 0x07, 0xa5, 0x85, 0x9b,
	0x3b, 0x0f, 0x5b, 0x31, 0xa9, 0x6c, 0x2e, 0x0a, 0x87, 0xc4, 0xd2, 0xde, 0x37, 0xfe, 0x7e, 0x54,
	0xfe, 0xb3, 0x54, 0xe2, 0xdc, 0x9c, 0x07, 0x31, 0x64, 0x2c, 0x06, 0x9d, 0x81, 0xee, 0x3e, 0x2f,
	0x75, 0xb2, 0x60, 0x58, 0xe6, 0x42, 0x07, 0x33, 0x85, 0x75, 0xe5, 0x3f, 0xbe, 0xeb, 0x6d, 0x23,
	0x47, 0xc3, 0x03, 0xfb, 0x67, 0xa6, 0x5a, 0xf8, 0xf4, 0xe3, 0xcd, 0xca, 0x23, 0xb7, 0x2b, 0x8f,
	0xfc, 0x5c, 0x79, 0xe4, 0xeb, 0xda, 0xeb, 0xdd, 0xae, 0xbd, 0xde, 0xf7, 0xb5, 0xd7, 0xfb, 0xfc,
	0xfa, 0x0e, 0x4a, 0x03, 0xcc, 0x41, 0x02, 0xdb, 0xe4, 0xbc, 0x3c, 0x61, 0xd7, 0xff, 0x85, 0x6d,
	0xf1, 0xe7, 0x03, 0x9b, 0xf2, 0xab, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x41, 0x68, 0x91, 0x24,
	0x89, 0x02, 0x00, 0x00,
}

func (m *EstimateSwapExactAmountOutRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EstimateSwapExactAmountOutRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EstimateSwapExactAmountOutRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenOut) > 0 {
		i -= len(m.TokenOut)
		copy(dAtA[i:], m.TokenOut)
		i = encodeVarintEstimateSwapExactAmoutOut(dAtA, i, uint64(len(m.TokenOut)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Routes) > 0 {
		for iNdEx := len(m.Routes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Routes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEstimateSwapExactAmoutOut(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.PoolId != 0 {
		i = encodeVarintEstimateSwapExactAmoutOut(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x10
	}
	return len(dAtA) - i, nil
}

func (m *EstimateSwapExactAmountOutResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EstimateSwapExactAmountOutResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EstimateSwapExactAmountOutResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TokenInAmount.Size()
		i -= size
		if _, err := m.TokenInAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEstimateSwapExactAmoutOut(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintEstimateSwapExactAmoutOut(dAtA []byte, offset int, v uint64) int {
	offset -= sovEstimateSwapExactAmoutOut(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EstimateSwapExactAmountOutRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovEstimateSwapExactAmoutOut(uint64(m.PoolId))
	}
	if len(m.Routes) > 0 {
		for _, e := range m.Routes {
			l = e.Size()
			n += 1 + l + sovEstimateSwapExactAmoutOut(uint64(l))
		}
	}
	l = len(m.TokenOut)
	if l > 0 {
		n += 1 + l + sovEstimateSwapExactAmoutOut(uint64(l))
	}
	return n
}

func (m *EstimateSwapExactAmountOutResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TokenInAmount.Size()
	n += 1 + l + sovEstimateSwapExactAmoutOut(uint64(l))
	return n
}

func sovEstimateSwapExactAmoutOut(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEstimateSwapExactAmoutOut(x uint64) (n int) {
	return sovEstimateSwapExactAmoutOut(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EstimateSwapExactAmountOutRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEstimateSwapExactAmoutOut
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
			return fmt.Errorf("proto: EstimateSwapExactAmountOutRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EstimateSwapExactAmountOutRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEstimateSwapExactAmoutOut
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Routes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEstimateSwapExactAmoutOut
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
				return ErrInvalidLengthEstimateSwapExactAmoutOut
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEstimateSwapExactAmoutOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Routes = append(m.Routes, SwapAmountOutRoute{})
			if err := m.Routes[len(m.Routes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEstimateSwapExactAmoutOut
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
				return ErrInvalidLengthEstimateSwapExactAmoutOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEstimateSwapExactAmoutOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenOut = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEstimateSwapExactAmoutOut(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEstimateSwapExactAmoutOut
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
func (m *EstimateSwapExactAmountOutResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEstimateSwapExactAmoutOut
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
			return fmt.Errorf("proto: EstimateSwapExactAmountOutResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EstimateSwapExactAmountOutResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenInAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEstimateSwapExactAmoutOut
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
				return ErrInvalidLengthEstimateSwapExactAmoutOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEstimateSwapExactAmoutOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenInAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEstimateSwapExactAmoutOut(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEstimateSwapExactAmoutOut
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
func skipEstimateSwapExactAmoutOut(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEstimateSwapExactAmoutOut
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
					return 0, ErrIntOverflowEstimateSwapExactAmoutOut
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
					return 0, ErrIntOverflowEstimateSwapExactAmoutOut
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
				return 0, ErrInvalidLengthEstimateSwapExactAmoutOut
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEstimateSwapExactAmoutOut
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEstimateSwapExactAmoutOut
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEstimateSwapExactAmoutOut        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEstimateSwapExactAmoutOut          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEstimateSwapExactAmoutOut = fmt.Errorf("proto: unexpected end of group")
)