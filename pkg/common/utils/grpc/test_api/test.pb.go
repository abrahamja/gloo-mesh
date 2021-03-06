// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package test_api

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Request)(nil), "grpc.test.Request")
	proto.RegisterType((*Response)(nil), "grpc.test.Response")
}

func init() {
	proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e)
}

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0x2f, 0x2a, 0x48, 0xd6, 0x03, 0x09, 0x28,
	0x71, 0x72, 0xb1, 0x07, 0xa5, 0x16, 0x96, 0x82, 0x98, 0x5c, 0x5c, 0x1c, 0x41, 0xa9, 0xc5, 0x05,
	0xf9, 0x79, 0xc5, 0xa9, 0x46, 0x4e, 0x5c, 0xdc, 0x21, 0xa9, 0xc5, 0x25, 0xc1, 0xa9, 0x45, 0x65,
	0x99, 0xc9, 0xa9, 0x42, 0xc6, 0x5c, 0x6c, 0x9e, 0x79, 0x65, 0xf9, 0xd9, 0xa9, 0x42, 0x42, 0x7a,
	0x70, 0xbd, 0x7a, 0x50, 0x8d, 0x52, 0xc2, 0x28, 0x62, 0x10, 0x13, 0x94, 0x18, 0x9c, 0x78, 0xa2,
	0xb8, 0xf4, 0xf4, 0x41, 0x82, 0xf1, 0x89, 0x05, 0x99, 0x49, 0x6c, 0x60, 0xab, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xb3, 0x1a, 0xde, 0xc4, 0x88, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	Invoke(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) Invoke(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.test.TestService/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	Invoke(context.Context, *Request) (*Response, error)
}

// UnimplementedTestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (*UnimplementedTestServiceServer) Invoke(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.test.TestService/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Invoke(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.test.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _TestService_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
