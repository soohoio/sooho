// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stayking/levstakeibc/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type QueryGetHostZoneRequest struct {
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (m *QueryGetHostZoneRequest) Reset()         { *m = QueryGetHostZoneRequest{} }
func (m *QueryGetHostZoneRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetHostZoneRequest) ProtoMessage()    {}
func (*QueryGetHostZoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_98806c3a17e4b7c4, []int{0}
}
func (m *QueryGetHostZoneRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetHostZoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetHostZoneRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetHostZoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetHostZoneRequest.Merge(m, src)
}
func (m *QueryGetHostZoneRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetHostZoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetHostZoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetHostZoneRequest proto.InternalMessageInfo

func (m *QueryGetHostZoneRequest) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

type QueryGetHostZoneResponse struct {
	HostZone HostZone `protobuf:"bytes,1,opt,name=host_zone,json=hostZone,proto3" json:"host_zone"`
}

func (m *QueryGetHostZoneResponse) Reset()         { *m = QueryGetHostZoneResponse{} }
func (m *QueryGetHostZoneResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetHostZoneResponse) ProtoMessage()    {}
func (*QueryGetHostZoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98806c3a17e4b7c4, []int{1}
}
func (m *QueryGetHostZoneResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetHostZoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetHostZoneResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetHostZoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetHostZoneResponse.Merge(m, src)
}
func (m *QueryGetHostZoneResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetHostZoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetHostZoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetHostZoneResponse proto.InternalMessageInfo

func (m *QueryGetHostZoneResponse) GetHostZone() HostZone {
	if m != nil {
		return m.HostZone
	}
	return HostZone{}
}

type QueryAllHostZoneRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllHostZoneRequest) Reset()         { *m = QueryAllHostZoneRequest{} }
func (m *QueryAllHostZoneRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllHostZoneRequest) ProtoMessage()    {}
func (*QueryAllHostZoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_98806c3a17e4b7c4, []int{2}
}
func (m *QueryAllHostZoneRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllHostZoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllHostZoneRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllHostZoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllHostZoneRequest.Merge(m, src)
}
func (m *QueryAllHostZoneRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllHostZoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllHostZoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllHostZoneRequest proto.InternalMessageInfo

func (m *QueryAllHostZoneRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllHostZoneResponse struct {
	HostZone   []HostZone          `protobuf:"bytes,1,rep,name=host_zone,json=hostZone,proto3" json:"host_zone"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllHostZoneResponse) Reset()         { *m = QueryAllHostZoneResponse{} }
func (m *QueryAllHostZoneResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllHostZoneResponse) ProtoMessage()    {}
func (*QueryAllHostZoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98806c3a17e4b7c4, []int{3}
}
func (m *QueryAllHostZoneResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllHostZoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllHostZoneResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllHostZoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllHostZoneResponse.Merge(m, src)
}
func (m *QueryAllHostZoneResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllHostZoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllHostZoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllHostZoneResponse proto.InternalMessageInfo

func (m *QueryAllHostZoneResponse) GetHostZone() []HostZone {
	if m != nil {
		return m.HostZone
	}
	return nil
}

func (m *QueryAllHostZoneResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetHostZoneRequest)(nil), "stayking.levstakeibc.QueryGetHostZoneRequest")
	proto.RegisterType((*QueryGetHostZoneResponse)(nil), "stayking.levstakeibc.QueryGetHostZoneResponse")
	proto.RegisterType((*QueryAllHostZoneRequest)(nil), "stayking.levstakeibc.QueryAllHostZoneRequest")
	proto.RegisterType((*QueryAllHostZoneResponse)(nil), "stayking.levstakeibc.QueryAllHostZoneResponse")
}

func init() { proto.RegisterFile("stayking/levstakeibc/query.proto", fileDescriptor_98806c3a17e4b7c4) }

var fileDescriptor_98806c3a17e4b7c4 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xe3, 0xf2, 0xaf, 0xf3, 0x6e, 0xd6, 0x24, 0x46, 0x85, 0xb2, 0x11, 0x21, 0x86, 0x26,
	0xb0, 0xb5, 0xc0, 0x17, 0xd8, 0x0e, 0x0c, 0x24, 0x0e, 0x90, 0xe3, 0x24, 0x34, 0x39, 0x99, 0xe5,
	0x58, 0xcb, 0xfc, 0xa6, 0xb5, 0x5b, 0x51, 0x10, 0x17, 0x3e, 0x01, 0x12, 0xe2, 0xc8, 0x95, 0xaf,
	0xc1, 0xb5, 0xc7, 0x4a, 0x5c, 0x38, 0x21, 0xd4, 0xf2, 0x41, 0x50, 0x9d, 0x84, 0xa6, 0x34, 0xb4,
	0xec, 0x16, 0x2b, 0xcf, 0xfb, 0x3e, 0xbf, 0xf7, 0x79, 0x6d, 0xbc, 0x6b, 0x2c, 0x1f, 0x9e, 0x2b,
	0x2d, 0x59, 0x26, 0x06, 0xc6, 0xf2, 0x73, 0xa1, 0xe2, 0x84, 0x75, 0xfb, 0xa2, 0x37, 0xa4, 0x79,
	0x0f, 0x2c, 0x90, 0xad, 0x4a, 0x41, 0x6b, 0x8a, 0xce, 0x96, 0x04, 0x09, 0x4e, 0xc0, 0x66, 0x5f,
	0x85, 0xb6, 0x73, 0x5b, 0x02, 0xc8, 0x4c, 0x30, 0x9e, 0x2b, 0xc6, 0xb5, 0x06, 0xcb, 0xad, 0x02,
	0x6d, 0xca, 0xbf, 0xfb, 0x09, 0x98, 0x0b, 0x30, 0x2c, 0xe6, 0x46, 0x14, 0x16, 0x6c, 0x70, 0x10,
	0x0b, 0xcb, 0x0f, 0x58, 0xce, 0xa5, 0xd2, 0x4e, 0x5c, 0x6a, 0xef, 0x36, 0x72, 0xa5, 0x60, 0xec,
	0xe9, 0x1b, 0xd0, 0xa2, 0x50, 0x05, 0x8f, 0xf1, 0xcd, 0x97, 0xb3, 0x3e, 0xc7, 0xc2, 0x3e, 0x05,
	0x63, 0x4f, 0x40, 0x8b, 0x48, 0x74, 0xfb, 0xc2, 0x58, 0x72, 0x0b, 0xb7, 0x93, 0x94, 0x2b, 0x7d,
	0xaa, 0xce, 0xb6, 0xd1, 0x2e, 0xba, 0xbf, 0x11, 0xdd, 0x70, 0xe7, 0x67, 0x67, 0xc1, 0x2b, 0xbc,
	0xbd, 0x5c, 0x65, 0x72, 0xd0, 0x46, 0x90, 0x43, 0xbc, 0xf1, 0xc7, 0xc4, 0xd5, 0x6d, 0x86, 0x3e,
	0x6d, 0x4a, 0x80, 0x56, 0xa5, 0x47, 0x57, 0x47, 0x3f, 0x76, 0xbc, 0xa8, 0x9d, 0x96, 0xe7, 0x80,
	0x97, 0x50, 0x87, 0x59, 0xf6, 0x37, 0xd4, 0x13, 0x8c, 0xe7, 0x93, 0x96, 0xed, 0xef, 0xd1, 0x22,
	0x16, 0x3a, 0x8b, 0x85, 0x16, 0xc9, 0x97, 0xb1, 0xd0, 0x17, 0x5c, 0x56, 0xb5, 0x51, 0xad, 0x32,
	0xf8, 0x82, 0xca, 0x11, 0x16, 0x3c, 0x9a, 0x47, 0xb8, 0x72, 0xf9, 0x11, 0xc8, 0xf1, 0x02, 0x67,
	0xcb, 0x71, 0xee, 0xad, 0xe5, 0x2c, 0xfc, 0xeb, 0xa0, 0xe1, 0xd7, 0x16, 0xbe, 0xe6, 0x40, 0xc9,
	0x67, 0x84, 0xdb, 0x95, 0x1f, 0x79, 0xd8, 0xcc, 0xf3, 0x8f, 0x5d, 0x76, 0xe8, 0xff, 0xca, 0x0b,
	0x82, 0x20, 0x7c, 0xff, 0xed, 0xd7, 0xc7, 0xd6, 0x03, 0xb2, 0xcf, 0x56, 0xdf, 0x22, 0xf6, 0xb6,
	0xba, 0x22, 0xef, 0xc8, 0x27, 0x84, 0x37, 0x6b, 0x69, 0xae, 0x44, 0x5c, 0xde, 0xec, 0x4a, 0xc4,
	0x86, 0x25, 0x05, 0x7b, 0x0e, 0xf1, 0x0e, 0xd9, 0x59, 0x83, 0x78, 0xf4, 0x7c, 0x34, 0xf1, 0xd1,
	0x78, 0xe2, 0xa3, 0x9f, 0x13, 0x1f, 0x7d, 0x98, 0xfa, 0xde, 0x78, 0xea, 0x7b, 0xdf, 0xa7, 0xbe,
	0x77, 0x12, 0x4a, 0x65, 0xd3, 0x7e, 0x4c, 0x13, 0xb8, 0x60, 0x06, 0x20, 0x05, 0x05, 0xf3, 0x66,
	0x83, 0x90, 0xbd, 0x5e, 0xe8, 0x68, 0x87, 0xb9, 0x30, 0xf1, 0x75, 0xf7, 0x6e, 0x1e, 0xfd, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xd7, 0x40, 0x69, 0x44, 0xf7, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	HostZone(ctx context.Context, in *QueryGetHostZoneRequest, opts ...grpc.CallOption) (*QueryGetHostZoneResponse, error)
	// Queries a list of HostZone items.
	AllHostZone(ctx context.Context, in *QueryAllHostZoneRequest, opts ...grpc.CallOption) (*QueryAllHostZoneResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) HostZone(ctx context.Context, in *QueryGetHostZoneRequest, opts ...grpc.CallOption) (*QueryGetHostZoneResponse, error) {
	out := new(QueryGetHostZoneResponse)
	err := c.cc.Invoke(ctx, "/stayking.levstakeibc.Query/HostZone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllHostZone(ctx context.Context, in *QueryAllHostZoneRequest, opts ...grpc.CallOption) (*QueryAllHostZoneResponse, error) {
	out := new(QueryAllHostZoneResponse)
	err := c.cc.Invoke(ctx, "/stayking.levstakeibc.Query/AllHostZone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	HostZone(context.Context, *QueryGetHostZoneRequest) (*QueryGetHostZoneResponse, error)
	// Queries a list of HostZone items.
	AllHostZone(context.Context, *QueryAllHostZoneRequest) (*QueryAllHostZoneResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) HostZone(ctx context.Context, req *QueryGetHostZoneRequest) (*QueryGetHostZoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostZone not implemented")
}
func (*UnimplementedQueryServer) AllHostZone(ctx context.Context, req *QueryAllHostZoneRequest) (*QueryAllHostZoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllHostZone not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_HostZone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetHostZoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HostZone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stayking.levstakeibc.Query/HostZone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HostZone(ctx, req.(*QueryGetHostZoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllHostZone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllHostZoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllHostZone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stayking.levstakeibc.Query/AllHostZone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllHostZone(ctx, req.(*QueryAllHostZoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stayking.levstakeibc.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HostZone",
			Handler:    _Query_HostZone_Handler,
		},
		{
			MethodName: "AllHostZone",
			Handler:    _Query_AllHostZone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stayking/levstakeibc/query.proto",
}

func (m *QueryGetHostZoneRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetHostZoneRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetHostZoneRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetHostZoneResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetHostZoneResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetHostZoneResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.HostZone.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllHostZoneRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllHostZoneRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllHostZoneRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllHostZoneResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllHostZoneResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllHostZoneResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.HostZone) > 0 {
		for iNdEx := len(m.HostZone) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HostZone[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetHostZoneRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetHostZoneResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.HostZone.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllHostZoneRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllHostZoneResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.HostZone) > 0 {
		for _, e := range m.HostZone {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetHostZoneRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetHostZoneRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetHostZoneRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetHostZoneResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetHostZoneResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetHostZoneResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostZone", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HostZone.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllHostZoneRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllHostZoneRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllHostZoneRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
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
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllHostZoneResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllHostZoneResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllHostZoneResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostZone", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostZone = append(m.HostZone, HostZone{})
			if err := m.HostZone[len(m.HostZone)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
