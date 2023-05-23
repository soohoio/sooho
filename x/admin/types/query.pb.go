// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stayking/admin/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
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

// QueryAdminsRequest is the request type for the Query/Pool RPC method.
type QueryAdminsRequest struct {
}

func (m *QueryAdminsRequest) Reset()         { *m = QueryAdminsRequest{} }
func (m *QueryAdminsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAdminsRequest) ProtoMessage()    {}
func (*QueryAdminsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9338030ca13b91aa, []int{0}
}
func (m *QueryAdminsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAdminsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAdminsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAdminsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAdminsRequest.Merge(m, src)
}
func (m *QueryAdminsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAdminsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAdminsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAdminsRequest proto.InternalMessageInfo

// QueryAdminsResponse is the response type for the Query/Pool RPC method.
type QueryAdminsResponse struct {
	// params defines the parameters of the module.
	Admins Admins `protobuf:"bytes,1,opt,name=admins,proto3" json:"admins"`
}

func (m *QueryAdminsResponse) Reset()         { *m = QueryAdminsResponse{} }
func (m *QueryAdminsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAdminsResponse) ProtoMessage()    {}
func (*QueryAdminsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9338030ca13b91aa, []int{1}
}
func (m *QueryAdminsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAdminsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAdminsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAdminsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAdminsResponse.Merge(m, src)
}
func (m *QueryAdminsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAdminsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAdminsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAdminsResponse proto.InternalMessageInfo

func (m *QueryAdminsResponse) GetAdmins() Admins {
	if m != nil {
		return m.Admins
	}
	return Admins{}
}

func init() {
	proto.RegisterType((*QueryAdminsRequest)(nil), "stayking.admin.v1.QueryAdminsRequest")
	proto.RegisterType((*QueryAdminsResponse)(nil), "stayking.admin.v1.QueryAdminsResponse")
}

func init() { proto.RegisterFile("stayking/admin/v1/query.proto", fileDescriptor_9338030ca13b91aa) }

var fileDescriptor_9338030ca13b91aa = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x2e, 0x49, 0xac,
	0xcc, 0xce, 0xcc, 0x4b, 0xd7, 0x4f, 0x4c, 0xc9, 0xcd, 0xcc, 0xd3, 0x2f, 0x33, 0xd4, 0x2f, 0x2c,
	0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0x49, 0xeb, 0x81, 0xa5,
	0xf5, 0xca, 0x0c, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xb2, 0xfa, 0x20, 0x16, 0x44, 0xa1,
	0x94, 0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e,
	0x7e, 0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31, 0x54, 0x16, 0x8b, 0x2d, 0x10, 0xf3, 0xc0, 0xd2,
	0x4a, 0x22, 0x5c, 0x42, 0x81, 0x20, 0x4b, 0x1d, 0x41, 0x62, 0xc5, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x4a, 0x7e, 0x5c, 0xc2, 0x28, 0xa2, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0xe6,
	0x5c, 0x6c, 0x60, 0xbd, 0xc5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x92, 0x7a, 0x18, 0x6e,
	0xd4, 0x83, 0x68, 0x71, 0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21, 0x08, 0xaa, 0xdc, 0xa8, 0x99, 0x91,
	0x8b, 0x15, 0x6c, 0xa0, 0x50, 0x15, 0x17, 0x1b, 0x44, 0x85, 0x90, 0x2a, 0x16, 0xcd, 0x98, 0x4e,
	0x91, 0x52, 0x23, 0xa4, 0x0c, 0xe2, 0x36, 0x25, 0xc5, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x49, 0x0b,
	0x49, 0xea, 0xe3, 0xf0, 0x70, 0xb1, 0x93, 0xdb, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31,
	0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb,
	0x31, 0x44, 0xe9, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x17, 0xe7,
	0xe7, 0x67, 0xe4, 0x67, 0xe6, 0x23, 0x8c, 0x29, 0x33, 0xd6, 0xaf, 0x80, 0x9a, 0x55, 0x52, 0x59,
	0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x0e, 0x3a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0xaa,
	0xb2, 0x73, 0xc1, 0x01, 0x00, 0x00,
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
	Admins(ctx context.Context, in *QueryAdminsRequest, opts ...grpc.CallOption) (*QueryAdminsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Admins(ctx context.Context, in *QueryAdminsRequest, opts ...grpc.CallOption) (*QueryAdminsResponse, error) {
	out := new(QueryAdminsResponse)
	err := c.cc.Invoke(ctx, "/stayking.admin.v1.Query/Admins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Admins(context.Context, *QueryAdminsRequest) (*QueryAdminsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Admins(ctx context.Context, req *QueryAdminsRequest) (*QueryAdminsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Admins not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Admins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAdminsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Admins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stayking.admin.v1.Query/Admins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Admins(ctx, req.(*QueryAdminsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stayking.admin.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Admins",
			Handler:    _Query_Admins_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stayking/admin/v1/query.proto",
}

func (m *QueryAdminsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAdminsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAdminsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAdminsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAdminsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAdminsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Admins.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryAdminsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAdminsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Admins.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAdminsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAdminsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAdminsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryAdminsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAdminsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAdminsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admins", wireType)
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
			if err := m.Admins.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
