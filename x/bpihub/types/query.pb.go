// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bpihub/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
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

// this line is used by starport scaffolding # 3
type QueryGetProviderRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetProviderRequest) Reset()         { *m = QueryGetProviderRequest{} }
func (m *QueryGetProviderRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetProviderRequest) ProtoMessage()    {}
func (*QueryGetProviderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4fa37fddb4f26ee, []int{0}
}
func (m *QueryGetProviderRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetProviderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetProviderRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetProviderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetProviderRequest.Merge(m, src)
}
func (m *QueryGetProviderRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetProviderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetProviderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetProviderRequest proto.InternalMessageInfo

func (m *QueryGetProviderRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryGetProviderResponse struct {
	Provider *Provider `protobuf:"bytes,1,opt,name=Provider,json=provider,proto3" json:"Provider,omitempty"`
}

func (m *QueryGetProviderResponse) Reset()         { *m = QueryGetProviderResponse{} }
func (m *QueryGetProviderResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetProviderResponse) ProtoMessage()    {}
func (*QueryGetProviderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4fa37fddb4f26ee, []int{1}
}
func (m *QueryGetProviderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetProviderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetProviderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetProviderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetProviderResponse.Merge(m, src)
}
func (m *QueryGetProviderResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetProviderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetProviderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetProviderResponse proto.InternalMessageInfo

func (m *QueryGetProviderResponse) GetProvider() *Provider {
	if m != nil {
		return m.Provider
	}
	return nil
}

type QueryAllProviderRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllProviderRequest) Reset()         { *m = QueryAllProviderRequest{} }
func (m *QueryAllProviderRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllProviderRequest) ProtoMessage()    {}
func (*QueryAllProviderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4fa37fddb4f26ee, []int{2}
}
func (m *QueryAllProviderRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllProviderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllProviderRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllProviderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllProviderRequest.Merge(m, src)
}
func (m *QueryAllProviderRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllProviderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllProviderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllProviderRequest proto.InternalMessageInfo

func (m *QueryAllProviderRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllProviderResponse struct {
	Provider   []*Provider         `protobuf:"bytes,1,rep,name=Provider,json=provider,proto3" json:"Provider,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllProviderResponse) Reset()         { *m = QueryAllProviderResponse{} }
func (m *QueryAllProviderResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllProviderResponse) ProtoMessage()    {}
func (*QueryAllProviderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4fa37fddb4f26ee, []int{3}
}
func (m *QueryAllProviderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllProviderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllProviderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllProviderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllProviderResponse.Merge(m, src)
}
func (m *QueryAllProviderResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllProviderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllProviderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllProviderResponse proto.InternalMessageInfo

func (m *QueryAllProviderResponse) GetProvider() []*Provider {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *QueryAllProviderResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetProviderRequest)(nil), "madtechworks.bpihub.bpihub.QueryGetProviderRequest")
	proto.RegisterType((*QueryGetProviderResponse)(nil), "madtechworks.bpihub.bpihub.QueryGetProviderResponse")
	proto.RegisterType((*QueryAllProviderRequest)(nil), "madtechworks.bpihub.bpihub.QueryAllProviderRequest")
	proto.RegisterType((*QueryAllProviderResponse)(nil), "madtechworks.bpihub.bpihub.QueryAllProviderResponse")
}

func init() { proto.RegisterFile("bpihub/query.proto", fileDescriptor_f4fa37fddb4f26ee) }

var fileDescriptor_f4fa37fddb4f26ee = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x2a, 0xc8, 0xcc,
	0x28, 0x4d, 0xd2, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92,
	0xca, 0x4d, 0x4c, 0x29, 0x49, 0x4d, 0xce, 0x28, 0xcf, 0x2f, 0xca, 0x2e, 0xd6, 0x83, 0x28, 0x80,
	0x52, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89, 0x79,
	0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0x9d, 0x52, 0x5a, 0xc9, 0xf9, 0xc5,
	0xb9, 0xf9, 0xc5, 0xfa, 0x49, 0x89, 0xc5, 0xa9, 0x10, 0x23, 0xf5, 0xcb, 0x0c, 0x93, 0x52, 0x4b,
	0x12, 0x0d, 0xf5, 0x0b, 0x12, 0xd3, 0x33, 0xf3, 0xc0, 0x8a, 0xa1, 0x6a, 0x45, 0xa1, 0x36, 0x17,
	0x14, 0xe5, 0x97, 0x65, 0xa6, 0xa4, 0x16, 0x41, 0x84, 0x95, 0x34, 0xb9, 0xc4, 0x03, 0x41, 0x1a,
	0xdd, 0x53, 0x4b, 0x02, 0xa0, 0x32, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x7c, 0x5c,
	0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x4c, 0x99, 0x29, 0x4a, 0x31, 0x5c,
	0x12, 0x98, 0x4a, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x1c, 0xb8, 0x38, 0x60, 0x62, 0x60,
	0x1d, 0xdc, 0x46, 0x2a, 0x7a, 0xb8, 0xbd, 0xa5, 0x07, 0xd7, 0xcf, 0x01, 0x73, 0x8e, 0x52, 0x22,
	0xd4, 0x21, 0x8e, 0x39, 0x39, 0xe8, 0x0e, 0x71, 0xe3, 0xe2, 0x42, 0x78, 0x07, 0x6a, 0xbc, 0x9a,
	0x1e, 0xc4, 0xef, 0x7a, 0x20, 0xbf, 0xeb, 0x41, 0x82, 0x13, 0xea, 0x77, 0xbd, 0x80, 0xc4, 0xf4,
	0x54, 0xa8, 0xde, 0x20, 0x24, 0x9d, 0x4a, 0x4b, 0x19, 0xa1, 0x3e, 0x40, 0xb1, 0x03, 0xab, 0x0f,
	0x98, 0x49, 0xf7, 0x81, 0x90, 0x3b, 0x8a, 0x33, 0x99, 0xc0, 0xce, 0x54, 0x27, 0xe8, 0x4c, 0x88,
	0xf5, 0xc8, 0xee, 0x34, 0xba, 0xcd, 0xc4, 0xc5, 0x0a, 0x76, 0xa7, 0xd0, 0x0a, 0x46, 0x84, 0xab,
	0x84, 0x8c, 0xf1, 0xb9, 0x07, 0x47, 0x24, 0x4a, 0x99, 0x90, 0xa6, 0x09, 0xe2, 0x1a, 0x25, 0xc3,
	0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x69, 0x0b, 0x69, 0xea, 0x23, 0xeb, 0xd6, 0x87, 0x26, 0x21, 0xb4,
	0x94, 0xa4, 0x5f, 0x9d, 0x99, 0x52, 0x2b, 0xb4, 0x8c, 0x91, 0x8b, 0x1b, 0x66, 0x8e, 0x63, 0x4e,
	0x0e, 0x11, 0xae, 0xc5, 0x8c, 0x69, 0x22, 0x5c, 0x8b, 0x25, 0xea, 0x94, 0x74, 0xc0, 0xae, 0x55,
	0x13, 0x52, 0x21, 0xc6, 0xb5, 0x4e, 0xee, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8,
	0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7,
	0x10, 0xa5, 0x9b, 0x9e, 0x59, 0x02, 0xb2, 0x28, 0x39, 0x3f, 0x17, 0xab, 0x49, 0x15, 0x30, 0x46,
	0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0x07, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x89, 0x8e, 0xbc, 0x54, 0xd4, 0x03, 0x00, 0x00,
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
	// this line is used by starport scaffolding # 2
	Provider(ctx context.Context, in *QueryGetProviderRequest, opts ...grpc.CallOption) (*QueryGetProviderResponse, error)
	ProviderAll(ctx context.Context, in *QueryAllProviderRequest, opts ...grpc.CallOption) (*QueryAllProviderResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Provider(ctx context.Context, in *QueryGetProviderRequest, opts ...grpc.CallOption) (*QueryGetProviderResponse, error) {
	out := new(QueryGetProviderResponse)
	err := c.cc.Invoke(ctx, "/madtechworks.bpihub.bpihub.Query/Provider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ProviderAll(ctx context.Context, in *QueryAllProviderRequest, opts ...grpc.CallOption) (*QueryAllProviderResponse, error) {
	out := new(QueryAllProviderResponse)
	err := c.cc.Invoke(ctx, "/madtechworks.bpihub.bpihub.Query/ProviderAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// this line is used by starport scaffolding # 2
	Provider(context.Context, *QueryGetProviderRequest) (*QueryGetProviderResponse, error)
	ProviderAll(context.Context, *QueryAllProviderRequest) (*QueryAllProviderResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Provider(ctx context.Context, req *QueryGetProviderRequest) (*QueryGetProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Provider not implemented")
}
func (*UnimplementedQueryServer) ProviderAll(ctx context.Context, req *QueryAllProviderRequest) (*QueryAllProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProviderAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Provider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Provider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/madtechworks.bpihub.bpihub.Query/Provider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Provider(ctx, req.(*QueryGetProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ProviderAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ProviderAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/madtechworks.bpihub.bpihub.Query/ProviderAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ProviderAll(ctx, req.(*QueryAllProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "madtechworks.bpihub.bpihub.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Provider",
			Handler:    _Query_Provider_Handler,
		},
		{
			MethodName: "ProviderAll",
			Handler:    _Query_ProviderAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bpihub/query.proto",
}

func (m *QueryGetProviderRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetProviderRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetProviderRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetProviderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetProviderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetProviderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Provider != nil {
		{
			size, err := m.Provider.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllProviderRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllProviderRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllProviderRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryAllProviderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllProviderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllProviderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Provider) > 0 {
		for iNdEx := len(m.Provider) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Provider[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryGetProviderRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryGetProviderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Provider != nil {
		l = m.Provider.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllProviderRequest) Size() (n int) {
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

func (m *QueryAllProviderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Provider) > 0 {
		for _, e := range m.Provider {
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
func (m *QueryGetProviderRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetProviderRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetProviderRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
func (m *QueryGetProviderResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetProviderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetProviderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
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
			if m.Provider == nil {
				m.Provider = &Provider{}
			}
			if err := m.Provider.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllProviderRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllProviderRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllProviderRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryAllProviderResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllProviderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllProviderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
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
			m.Provider = append(m.Provider, &Provider{})
			if err := m.Provider[len(m.Provider)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
