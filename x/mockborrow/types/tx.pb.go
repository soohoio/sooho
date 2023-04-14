// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stayking/mockborrow/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgBorrow struct {
	From         string                                   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Denom        string                                   `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Collateral   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=collateral,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"collateral" yaml:"collateral"`
	BorrowAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=borrow_amount,json=borrowAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"borrow_amount" yaml:"borrow_amount"`
}

func (m *MsgBorrow) Reset()         { *m = MsgBorrow{} }
func (m *MsgBorrow) String() string { return proto.CompactTextString(m) }
func (*MsgBorrow) ProtoMessage()    {}
func (*MsgBorrow) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f8368ebff528cef, []int{0}
}
func (m *MsgBorrow) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBorrow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBorrow.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBorrow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBorrow.Merge(m, src)
}
func (m *MsgBorrow) XXX_Size() int {
	return m.Size()
}
func (m *MsgBorrow) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBorrow.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBorrow proto.InternalMessageInfo

type MsgBorrowResponse struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgBorrowResponse) Reset()         { *m = MsgBorrowResponse{} }
func (m *MsgBorrowResponse) String() string { return proto.CompactTextString(m) }
func (*MsgBorrowResponse) ProtoMessage()    {}
func (*MsgBorrowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f8368ebff528cef, []int{1}
}
func (m *MsgBorrowResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBorrowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBorrowResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBorrowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBorrowResponse.Merge(m, src)
}
func (m *MsgBorrowResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgBorrowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBorrowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBorrowResponse proto.InternalMessageInfo

func (m *MsgBorrowResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MsgRepay struct {
	From   string                                   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Id     uint64                                   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount" yaml:"amount"`
}

func (m *MsgRepay) Reset()         { *m = MsgRepay{} }
func (m *MsgRepay) String() string { return proto.CompactTextString(m) }
func (*MsgRepay) ProtoMessage()    {}
func (*MsgRepay) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f8368ebff528cef, []int{2}
}
func (m *MsgRepay) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRepay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRepay.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRepay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRepay.Merge(m, src)
}
func (m *MsgRepay) XXX_Size() int {
	return m.Size()
}
func (m *MsgRepay) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRepay.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRepay proto.InternalMessageInfo

type MsgRepayResponse struct {
	Change github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=change,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"change" yaml:"borrow_amount"`
}

func (m *MsgRepayResponse) Reset()         { *m = MsgRepayResponse{} }
func (m *MsgRepayResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRepayResponse) ProtoMessage()    {}
func (*MsgRepayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f8368ebff528cef, []int{3}
}
func (m *MsgRepayResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRepayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRepayResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRepayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRepayResponse.Merge(m, src)
}
func (m *MsgRepayResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRepayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRepayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRepayResponse proto.InternalMessageInfo

func (m *MsgRepayResponse) GetChange() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Change
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgBorrow)(nil), "stayking.mockborrow.v1.MsgBorrow")
	proto.RegisterType((*MsgBorrowResponse)(nil), "stayking.mockborrow.v1.MsgBorrowResponse")
	proto.RegisterType((*MsgRepay)(nil), "stayking.mockborrow.v1.MsgRepay")
	proto.RegisterType((*MsgRepayResponse)(nil), "stayking.mockborrow.v1.MsgRepayResponse")
}

func init() { proto.RegisterFile("stayking/mockborrow/v1/tx.proto", fileDescriptor_2f8368ebff528cef) }

var fileDescriptor_2f8368ebff528cef = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xbf, 0x6e, 0xd3, 0x40,
	0x1c, 0xf6, 0x25, 0x69, 0xd4, 0x1e, 0x14, 0xd1, 0x53, 0x84, 0x4c, 0x06, 0x3b, 0x98, 0x25, 0x0c,
	0xdc, 0xc9, 0x61, 0xeb, 0xd6, 0x20, 0x24, 0x24, 0x94, 0xc5, 0x48, 0x0c, 0x2c, 0xe8, 0xec, 0x1c,
	0x17, 0x2b, 0xb1, 0x7f, 0x91, 0xef, 0x1a, 0x1a, 0xb1, 0x30, 0x76, 0x83, 0x47, 0xe8, 0xcc, 0xc0,
	0xc0, 0x53, 0x74, 0x41, 0xea, 0xc8, 0x54, 0x50, 0xb2, 0x30, 0xf3, 0x04, 0xa8, 0x77, 0x8e, 0x1b,
	0xa4, 0x8a, 0x00, 0x43, 0xa7, 0x5c, 0xec, 0xef, 0xfb, 0x7d, 0x7f, 0xce, 0x77, 0xd8, 0x57, 0x9a,
	0xcf, 0xc7, 0x69, 0x2e, 0x59, 0x06, 0xc9, 0x38, 0x86, 0xa2, 0x80, 0x37, 0x6c, 0x16, 0x32, 0x7d,
	0x44, 0xa7, 0x05, 0x68, 0x20, 0x77, 0x56, 0x00, 0x7a, 0x09, 0xa0, 0xb3, 0xb0, 0xdd, 0x92, 0x20,
	0xc1, 0x40, 0xd8, 0xc5, 0xca, 0xa2, 0xdb, 0x5e, 0x02, 0x2a, 0x03, 0xc5, 0x62, 0xae, 0x04, 0x9b,
	0x85, 0xb1, 0xd0, 0x3c, 0x64, 0x09, 0xa4, 0xb9, 0x7d, 0x1f, 0x7c, 0xa9, 0xe1, 0x9d, 0x81, 0x92,
	0x7d, 0x33, 0x86, 0x10, 0xdc, 0x78, 0x5d, 0x40, 0xe6, 0xa2, 0x0e, 0xea, 0xee, 0x44, 0x66, 0x4d,
	0x5a, 0x78, 0x6b, 0x28, 0x72, 0xc8, 0xdc, 0x9a, 0x79, 0x68, 0xff, 0x90, 0x77, 0x08, 0xe3, 0x04,
	0x26, 0x13, 0xae, 0x45, 0xc1, 0x27, 0x6e, 0xbd, 0x53, 0xef, 0xde, 0xe8, 0xdd, 0xa5, 0x56, 0x8d,
	0x5e, 0xa8, 0xd1, 0x52, 0x8d, 0x3e, 0x86, 0x34, 0xef, 0x3f, 0x39, 0x3d, 0xf7, 0x9d, 0x9f, 0xe7,
	0xfe, 0xde, 0x9c, 0x67, 0x93, 0xfd, 0xe0, 0x92, 0x1a, 0x7c, 0xfc, 0xe6, 0x77, 0x65, 0xaa, 0x47,
	0x87, 0x31, 0x4d, 0x20, 0x63, 0xa5, 0x5f, 0xfb, 0xf3, 0x50, 0x0d, 0xc7, 0x4c, 0xcf, 0xa7, 0x42,
	0x99, 0x29, 0x2a, 0x5a, 0xd3, 0x24, 0xc7, 0x08, 0xef, 0xda, 0xf8, 0xaf, 0x78, 0x06, 0x87, 0xb9,
	0x76, 0x1b, 0x9b, 0x5c, 0x3c, 0x2d, 0x5d, 0xb4, 0xac, 0x8b, 0xdf, 0xd8, 0xff, 0x66, 0xe4, 0xa6,
	0xe5, 0x1e, 0x18, 0xea, 0xfe, 0xf6, 0xf1, 0x89, 0xef, 0xfc, 0x38, 0xf1, 0x9d, 0xe0, 0x3e, 0xde,
	0xab, 0xea, 0x8c, 0x84, 0x9a, 0x42, 0xae, 0x04, 0xb9, 0x85, 0x6b, 0xe9, 0xd0, 0x94, 0xda, 0x88,
	0x6a, 0xe9, 0x30, 0xf8, 0x84, 0xf0, 0xf6, 0x40, 0xc9, 0x48, 0x4c, 0xf9, 0xfc, 0xca, 0xce, 0x2d,
	0xa1, 0xbe, 0x22, 0x10, 0x8d, 0x9b, 0x7f, 0x1b, 0xf1, 0xa0, 0x8c, 0xb8, 0x6b, 0x23, 0xfe, 0x4f,
	0xb6, 0x52, 0x6b, 0x2d, 0xd5, 0x7b, 0x84, 0x6f, 0xaf, 0x0c, 0x57, 0xa9, 0xde, 0xe2, 0x66, 0x32,
	0xe2, 0xb9, 0x14, 0x2e, 0xba, 0xbe, 0xde, 0x4b, 0xc9, 0xde, 0x67, 0x84, 0xeb, 0x03, 0x25, 0xc9,
	0x0b, 0xdc, 0x2c, 0xbf, 0xdd, 0x7b, 0xf4, 0xea, 0x83, 0x41, 0xab, 0xfd, 0x68, 0x3f, 0xd8, 0x08,
	0xa9, 0xc2, 0x3d, 0xc7, 0x5b, 0x76, 0x7b, 0x3a, 0x7f, 0xe0, 0x18, 0x44, 0xbb, 0xbb, 0x09, 0xb1,
	0x1a, 0xda, 0x7f, 0x76, 0xba, 0xf0, 0xd0, 0xd9, 0xc2, 0x43, 0xdf, 0x17, 0x1e, 0xfa, 0xb0, 0xf4,
	0x9c, 0xb3, 0xa5, 0xe7, 0x7c, 0x5d, 0x7a, 0xce, 0xcb, 0x70, 0xad, 0x00, 0x05, 0x30, 0x82, 0x14,
	0x58, 0x75, 0x11, 0xcc, 0x7a, 0xec, 0x68, 0xfd, 0x36, 0x30, 0x7d, 0xc4, 0x4d, 0x73, 0x80, 0x1f,
	0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xb4, 0x6a, 0x33, 0x31, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Borrow(ctx context.Context, in *MsgBorrow, opts ...grpc.CallOption) (*MsgBorrowResponse, error)
	Repay(ctx context.Context, in *MsgRepay, opts ...grpc.CallOption) (*MsgRepayResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Borrow(ctx context.Context, in *MsgBorrow, opts ...grpc.CallOption) (*MsgBorrowResponse, error) {
	out := new(MsgBorrowResponse)
	err := c.cc.Invoke(ctx, "/stayking.mockborrow.v1.Msg/Borrow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Repay(ctx context.Context, in *MsgRepay, opts ...grpc.CallOption) (*MsgRepayResponse, error) {
	out := new(MsgRepayResponse)
	err := c.cc.Invoke(ctx, "/stayking.mockborrow.v1.Msg/Repay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Borrow(context.Context, *MsgBorrow) (*MsgBorrowResponse, error)
	Repay(context.Context, *MsgRepay) (*MsgRepayResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Borrow(ctx context.Context, req *MsgBorrow) (*MsgBorrowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Borrow not implemented")
}
func (*UnimplementedMsgServer) Repay(ctx context.Context, req *MsgRepay) (*MsgRepayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Repay not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Borrow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBorrow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Borrow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stayking.mockborrow.v1.Msg/Borrow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Borrow(ctx, req.(*MsgBorrow))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Repay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRepay)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Repay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stayking.mockborrow.v1.Msg/Repay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Repay(ctx, req.(*MsgRepay))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stayking.mockborrow.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Borrow",
			Handler:    _Msg_Borrow_Handler,
		},
		{
			MethodName: "Repay",
			Handler:    _Msg_Repay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stayking/mockborrow/v1/tx.proto",
}

func (m *MsgBorrow) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBorrow) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBorrow) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BorrowAmount) > 0 {
		for iNdEx := len(m.BorrowAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BorrowAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Collateral) > 0 {
		for iNdEx := len(m.Collateral) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Collateral[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBorrowResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBorrowResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBorrowResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgRepay) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRepay) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRepay) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x18
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRepayResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRepayResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRepayResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Change) > 0 {
		for iNdEx := len(m.Change) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Change[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgBorrow) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Collateral) > 0 {
		for _, e := range m.Collateral {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.BorrowAmount) > 0 {
		for _, e := range m.BorrowAmount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgBorrowResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgRepay) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgRepayResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Change) > 0 {
		for _, e := range m.Change {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgBorrow) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBorrow: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBorrow: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collateral = append(m.Collateral, types.Coin{})
			if err := m.Collateral[len(m.Collateral)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BorrowAmount = append(m.BorrowAmount, types.Coin{})
			if err := m.BorrowAmount[len(m.BorrowAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgBorrowResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBorrowResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBorrowResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRepay) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRepay: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRepay: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRepayResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRepayResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRepayResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Change", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Change = append(m.Change, types.Coin{})
			if err := m.Change[len(m.Change)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
