// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stayking/interchainquery/v1/messages.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	crypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// MsgSubmitQueryResponse represents a message type to fulfil a query request.
type MsgSubmitQueryResponse struct {
	ChainId     string           `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty" yaml:"chain_id"`
	QueryId     string           `protobuf:"bytes,2,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty" yaml:"query_id"`
	Result      []byte           `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty" yaml:"result"`
	ProofOps    *crypto.ProofOps `protobuf:"bytes,4,opt,name=proof_ops,json=proofOps,proto3" json:"proof_ops,omitempty" yaml:"proof_ops"`
	Height      int64            `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty" yaml:"height"`
	FromAddress string           `protobuf:"bytes,6,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
}

func (m *MsgSubmitQueryResponse) Reset()         { *m = MsgSubmitQueryResponse{} }
func (m *MsgSubmitQueryResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitQueryResponse) ProtoMessage()    {}
func (*MsgSubmitQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1ebc5c6d8a15b8, []int{0}
}
func (m *MsgSubmitQueryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitQueryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitQueryResponse.Merge(m, src)
}
func (m *MsgSubmitQueryResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitQueryResponse proto.InternalMessageInfo

// MsgSubmitQueryResponseResponse defines the MsgSubmitQueryResponse response
// type.
type MsgSubmitQueryResponseResponse struct {
}

func (m *MsgSubmitQueryResponseResponse) Reset()         { *m = MsgSubmitQueryResponseResponse{} }
func (m *MsgSubmitQueryResponseResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitQueryResponseResponse) ProtoMessage()    {}
func (*MsgSubmitQueryResponseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1ebc5c6d8a15b8, []int{1}
}
func (m *MsgSubmitQueryResponseResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitQueryResponseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitQueryResponseResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitQueryResponseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitQueryResponseResponse.Merge(m, src)
}
func (m *MsgSubmitQueryResponseResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitQueryResponseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitQueryResponseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitQueryResponseResponse proto.InternalMessageInfo

type MsgSendQueryAllBalances struct {
	Creator    string             `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ChannelId  string             `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Address    string             `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Pagination *query.PageRequest `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *MsgSendQueryAllBalances) Reset()         { *m = MsgSendQueryAllBalances{} }
func (m *MsgSendQueryAllBalances) String() string { return proto.CompactTextString(m) }
func (*MsgSendQueryAllBalances) ProtoMessage()    {}
func (*MsgSendQueryAllBalances) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1ebc5c6d8a15b8, []int{2}
}
func (m *MsgSendQueryAllBalances) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendQueryAllBalances) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendQueryAllBalances.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendQueryAllBalances) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendQueryAllBalances.Merge(m, src)
}
func (m *MsgSendQueryAllBalances) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendQueryAllBalances) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendQueryAllBalances.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendQueryAllBalances proto.InternalMessageInfo

func (m *MsgSendQueryAllBalances) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSendQueryAllBalances) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *MsgSendQueryAllBalances) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgSendQueryAllBalances) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type MsgSendQueryAllBalancesResponse struct {
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *MsgSendQueryAllBalancesResponse) Reset()         { *m = MsgSendQueryAllBalancesResponse{} }
func (m *MsgSendQueryAllBalancesResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendQueryAllBalancesResponse) ProtoMessage()    {}
func (*MsgSendQueryAllBalancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1ebc5c6d8a15b8, []int{3}
}
func (m *MsgSendQueryAllBalancesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendQueryAllBalancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendQueryAllBalancesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendQueryAllBalancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendQueryAllBalancesResponse.Merge(m, src)
}
func (m *MsgSendQueryAllBalancesResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendQueryAllBalancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendQueryAllBalancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendQueryAllBalancesResponse proto.InternalMessageInfo

func (m *MsgSendQueryAllBalancesResponse) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgSubmitQueryResponse)(nil), "stayking.interchainquery.v1.MsgSubmitQueryResponse")
	proto.RegisterType((*MsgSubmitQueryResponseResponse)(nil), "stayking.interchainquery.v1.MsgSubmitQueryResponseResponse")
	proto.RegisterType((*MsgSendQueryAllBalances)(nil), "stayking.interchainquery.v1.MsgSendQueryAllBalances")
	proto.RegisterType((*MsgSendQueryAllBalancesResponse)(nil), "stayking.interchainquery.v1.MsgSendQueryAllBalancesResponse")
}

func init() {
	proto.RegisterFile("stayking/interchainquery/v1/messages.proto", fileDescriptor_ea1ebc5c6d8a15b8)
}

var fileDescriptor_ea1ebc5c6d8a15b8 = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x3f, 0x6f, 0xd4, 0x3e,
	0x18, 0xc7, 0xcf, 0xbd, 0xfe, 0xfa, 0xc7, 0xed, 0x4f, 0x40, 0x5a, 0xc1, 0x71, 0xa5, 0xc9, 0x29,
	0x03, 0x1c, 0x15, 0xd8, 0xea, 0x55, 0x62, 0x68, 0xc5, 0xd0, 0x1b, 0x90, 0x3a, 0x14, 0x4a, 0xba,
	0xb1, 0x9c, 0x7c, 0x89, 0xeb, 0xb3, 0x48, 0xec, 0x34, 0xf6, 0x55, 0xbd, 0x95, 0x89, 0x11, 0x89,
	0x85, 0xb1, 0x2f, 0xa2, 0x1b, 0x23, 0x0b, 0x63, 0x05, 0x0b, 0x53, 0x85, 0x5a, 0x06, 0x58, 0xfb,
	0x0a, 0x50, 0xec, 0x24, 0xad, 0xda, 0xc2, 0xc0, 0x14, 0xfb, 0xf9, 0x7e, 0x1e, 0xfb, 0x79, 0x1e,
	0x3f, 0x4f, 0xe0, 0x92, 0xd2, 0x64, 0xf4, 0x9a, 0x0b, 0x86, 0xb9, 0xd0, 0x34, 0x0b, 0x07, 0x84,
	0x8b, 0xdd, 0x21, 0xcd, 0x46, 0x78, 0x6f, 0x19, 0x27, 0x54, 0x29, 0xc2, 0xa8, 0x42, 0x69, 0x26,
	0xb5, 0x74, 0x16, 0x4a, 0x16, 0x5d, 0x62, 0xd1, 0xde, 0x72, 0x73, 0x9e, 0x49, 0x26, 0x0d, 0x87,
	0xf3, 0x95, 0x75, 0x69, 0xde, 0x0d, 0xa5, 0x4a, 0xa4, 0xea, 0x59, 0xc1, 0x6e, 0x0a, 0xe9, 0x1e,
	0x93, 0x92, 0xc5, 0x14, 0x93, 0x94, 0x63, 0x22, 0x84, 0xd4, 0x44, 0x73, 0x29, 0x4a, 0x75, 0x51,
	0x53, 0x11, 0xd1, 0x2c, 0xe1, 0x42, 0xe3, 0x30, 0x1b, 0xa5, 0x5a, 0xe2, 0x34, 0x93, 0x72, 0xa7,
	0x90, 0x97, 0xec, 0x51, 0xb8, 0x4f, 0x14, 0xc5, 0x65, 0xbc, 0x7d, 0xaa, 0xc9, 0x32, 0x4e, 0x09,
	0xe3, 0xc2, 0x9c, 0x65, 0x59, 0xff, 0xd7, 0x18, 0xbc, 0xbd, 0xa9, 0xd8, 0xf6, 0xb0, 0x9f, 0x70,
	0xfd, 0x32, 0x67, 0x03, 0xaa, 0x52, 0x29, 0x14, 0x75, 0x10, 0x9c, 0x32, 0x59, 0xf4, 0x78, 0xd4,
	0x00, 0x2d, 0xd0, 0x9e, 0xee, 0xce, 0x9d, 0x1d, 0x7b, 0x37, 0x46, 0x24, 0x89, 0x57, 0xfd, 0x52,
	0xf1, 0x83, 0x49, 0xb3, 0xdc, 0x88, 0x72, 0xde, 0x5c, 0x96, 0xf3, 0x63, 0x97, 0xf9, 0x52, 0xf1,
	0x83, 0x49, 0xb3, 0xdc, 0x88, 0x9c, 0x87, 0x70, 0x22, 0xa3, 0x6a, 0x18, 0xeb, 0x46, 0xbd, 0x05,
	0xda, 0xb3, 0xdd, 0x5b, 0x67, 0xc7, 0xde, 0xff, 0x96, 0xb6, 0x76, 0x3f, 0x28, 0x00, 0xe7, 0x39,
	0x9c, 0x36, 0x09, 0xf6, 0x64, 0xaa, 0x1a, 0xe3, 0x2d, 0xd0, 0x9e, 0xe9, 0x2c, 0xa0, 0xf3, 0x22,
	0x20, 0x5b, 0x04, 0xb4, 0x95, 0x33, 0x2f, 0x52, 0xd5, 0x9d, 0x3f, 0x3b, 0xf6, 0x6e, 0xda, 0xa3,
	0x2a, 0x3f, 0x3f, 0x98, 0x4a, 0x0b, 0x3d, 0xbf, 0x7a, 0x40, 0x39, 0x1b, 0xe8, 0xc6, 0x7f, 0x2d,
	0xd0, 0xae, 0x5f, 0xbc, 0xda, 0xda, 0xfd, 0xa0, 0x00, 0x9c, 0x35, 0x38, 0xbb, 0x93, 0xc9, 0xa4,
	0x47, 0xa2, 0x28, 0xa3, 0x4a, 0x35, 0x26, 0x4c, 0x66, 0x8d, 0x2f, 0x87, 0x8f, 0xe7, 0x8b, 0x17,
	0x5b, 0xb7, 0xca, 0xb6, 0xce, 0xb8, 0x60, 0xc1, 0x4c, 0x4e, 0x17, 0xa6, 0xd5, 0xd9, 0xb7, 0x07,
	0x5e, 0xed, 0xc3, 0x81, 0x07, 0x7e, 0x1e, 0x78, 0x35, 0xbf, 0x05, 0xdd, 0xeb, 0x4b, 0x5d, 0x7e,
	0xfd, 0x43, 0x00, 0xef, 0xe4, 0x08, 0x15, 0x91, 0x01, 0xd6, 0xe3, 0xb8, 0x4b, 0x62, 0x22, 0x42,
	0xaa, 0x9c, 0x06, 0x9c, 0x0c, 0x33, 0x4a, 0xb4, 0xcc, 0xec, 0x6b, 0x04, 0xe5, 0xd6, 0x59, 0x84,
	0x30, 0x1c, 0x10, 0x21, 0x68, 0x5c, 0x95, 0x3e, 0x98, 0x2e, 0x2c, 0x1b, 0x51, 0xee, 0x58, 0x06,
	0x5f, 0xb7, 0x8e, 0xc5, 0xd6, 0x79, 0x06, 0xe1, 0x79, 0x43, 0x14, 0x75, 0xbd, 0x8f, 0x8a, 0xb4,
	0xf2, 0xee, 0x41, 0x65, 0x07, 0x9b, 0xee, 0x41, 0x5b, 0x84, 0xd1, 0x80, 0xee, 0x0e, 0xa9, 0xd2,
	0xc1, 0x05, 0x4f, 0xff, 0x29, 0xf4, 0xfe, 0x10, 0x75, 0xd5, 0x4c, 0x4d, 0x38, 0xa5, 0x72, 0x4f,
	0x11, 0x52, 0x13, 0xfe, 0x78, 0x50, 0xed, 0x3b, 0x9f, 0x00, 0xac, 0x6f, 0x2a, 0xe6, 0x7c, 0x04,
	0x70, 0xee, 0xba, 0x46, 0x5c, 0x41, 0x7f, 0x99, 0x2d, 0x74, 0x7d, 0x49, 0x9b, 0x6b, 0xff, 0xe0,
	0x54, 0xbd, 0x43, 0xe7, 0xcd, 0xd7, 0x1f, 0xef, 0xc7, 0x1e, 0xf9, 0x0f, 0xae, 0x0c, 0xbe, 0xde,
	0xaf, 0x66, 0x49, 0x99, 0x03, 0x8c, 0x79, 0x15, 0x2c, 0x75, 0xb7, 0x3e, 0x9f, 0xb8, 0xe0, 0xe8,
	0xc4, 0x05, 0xdf, 0x4f, 0x5c, 0xf0, 0xee, 0xd4, 0xad, 0x1d, 0x9d, 0xba, 0xb5, 0x6f, 0xa7, 0x6e,
	0xed, 0xd5, 0x13, 0xc6, 0xf5, 0x60, 0xd8, 0x47, 0xa1, 0x4c, 0xb0, 0x92, 0x72, 0x20, 0xb9, 0xc4,
	0xd5, 0x9f, 0x65, 0xaf, 0x83, 0xf7, 0xaf, 0xde, 0x32, 0x4a, 0xa9, 0xea, 0x4f, 0x98, 0x11, 0x5d,
	0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x46, 0xa1, 0x61, 0x87, 0x04, 0x00, 0x00,
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
	// SubmitQueryResponse defines a method for submit query responses.
	SubmitQueryResponse(ctx context.Context, in *MsgSubmitQueryResponse, opts ...grpc.CallOption) (*MsgSubmitQueryResponseResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitQueryResponse(ctx context.Context, in *MsgSubmitQueryResponse, opts ...grpc.CallOption) (*MsgSubmitQueryResponseResponse, error) {
	out := new(MsgSubmitQueryResponseResponse)
	err := c.cc.Invoke(ctx, "/stayking.interchainquery.v1.Msg/SubmitQueryResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SubmitQueryResponse defines a method for submit query responses.
	SubmitQueryResponse(context.Context, *MsgSubmitQueryResponse) (*MsgSubmitQueryResponseResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitQueryResponse(ctx context.Context, req *MsgSubmitQueryResponse) (*MsgSubmitQueryResponseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitQueryResponse not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitQueryResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitQueryResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitQueryResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stayking.interchainquery.v1.Msg/SubmitQueryResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitQueryResponse(ctx, req.(*MsgSubmitQueryResponse))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stayking.interchainquery.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitQueryResponse",
			Handler:    _Msg_SubmitQueryResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stayking/interchainquery/v1/messages.proto",
}

func (m *MsgSubmitQueryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitQueryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitQueryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0x32
	}
	if m.Height != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x28
	}
	if m.ProofOps != nil {
		{
			size, err := m.ProofOps.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessages(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Result) > 0 {
		i -= len(m.Result)
		copy(dAtA[i:], m.Result)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.QueryId) > 0 {
		i -= len(m.QueryId)
		copy(dAtA[i:], m.QueryId)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.QueryId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitQueryResponseResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitQueryResponseResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitQueryResponseResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSendQueryAllBalances) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendQueryAllBalances) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendQueryAllBalances) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessages(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendQueryAllBalancesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendQueryAllBalancesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendQueryAllBalancesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessages(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessages(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSubmitQueryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.QueryId)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.ProofOps != nil {
		l = m.ProofOps.Size()
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovMessages(uint64(m.Height))
	}
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *MsgSubmitQueryResponseResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSendQueryAllBalances) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *MsgSendQueryAllBalancesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovMessages(uint64(m.Sequence))
	}
	return n
}

func sovMessages(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSubmitQueryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgSubmitQueryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitQueryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = append(m.Result[:0], dAtA[iNdEx:postIndex]...)
			if m.Result == nil {
				m.Result = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofOps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ProofOps == nil {
				m.ProofOps = &crypto.ProofOps{}
			}
			if err := m.ProofOps.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func (m *MsgSubmitQueryResponseResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgSubmitQueryResponseResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitQueryResponseResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func (m *MsgSendQueryAllBalances) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgSendQueryAllBalances: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendQueryAllBalances: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func (m *MsgSendQueryAllBalancesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgSendQueryAllBalancesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendQueryAllBalancesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func skipMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
				return 0, ErrInvalidLengthMessages
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessages
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessages
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessages        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessages          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessages = fmt.Errorf("proto: unexpected end of group")
)
