// Code generated by protoc-gen-go.
// source: publisher.proto
// DO NOT EDIT!

/*
Package publisher is a generated protocol buffer package.

It is generated from these files:
	publisher.proto

It has these top-level messages:
	Request
	Result
	Empty
*/
package publisher

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Der              []byte  `protobuf:"bytes,1,opt,name=der" json:"der,omitempty"`
	LogURL           *string `protobuf:"bytes,2,opt,name=LogURL,json=logURL" json:"LogURL,omitempty"`
	LogPublicKey     *string `protobuf:"bytes,3,opt,name=LogPublicKey,json=logPublicKey" json:"LogPublicKey,omitempty"`
	Precert          *bool   `protobuf:"varint,4,opt,name=precert" json:"precert,omitempty"`
	StoreSCT         *bool   `protobuf:"varint,5,opt,name=storeSCT" json:"storeSCT,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetDer() []byte {
	if m != nil {
		return m.Der
	}
	return nil
}

func (m *Request) GetLogURL() string {
	if m != nil && m.LogURL != nil {
		return *m.LogURL
	}
	return ""
}

func (m *Request) GetLogPublicKey() string {
	if m != nil && m.LogPublicKey != nil {
		return *m.LogPublicKey
	}
	return ""
}

func (m *Request) GetPrecert() bool {
	if m != nil && m.Precert != nil {
		return *m.Precert
	}
	return false
}

func (m *Request) GetStoreSCT() bool {
	if m != nil && m.StoreSCT != nil {
		return *m.StoreSCT
	}
	return false
}

type Result struct {
	Sct              []byte `protobuf:"bytes,1,opt,name=sct" json:"sct,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetSct() []byte {
	if m != nil {
		return m.Sct
	}
	return nil
}

type Empty struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Result)(nil), "Result")
	proto.RegisterType((*Empty)(nil), "Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Publisher service

type PublisherClient interface {
	SubmitToCT(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Empty, error)
	SubmitToSingleCT(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Empty, error)
	SubmitToSingleCTWithResult(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type publisherClient struct {
	cc *grpc.ClientConn
}

func NewPublisherClient(cc *grpc.ClientConn) PublisherClient {
	return &publisherClient{cc}
}

func (c *publisherClient) SubmitToCT(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/Publisher/SubmitToCT", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publisherClient) SubmitToSingleCT(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/Publisher/SubmitToSingleCT", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publisherClient) SubmitToSingleCTWithResult(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/Publisher/SubmitToSingleCTWithResult", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Publisher service

type PublisherServer interface {
	SubmitToCT(context.Context, *Request) (*Empty, error)
	SubmitToSingleCT(context.Context, *Request) (*Empty, error)
	SubmitToSingleCTWithResult(context.Context, *Request) (*Result, error)
}

func RegisterPublisherServer(s *grpc.Server, srv PublisherServer) {
	s.RegisterService(&_Publisher_serviceDesc, srv)
}

func _Publisher_SubmitToCT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherServer).SubmitToCT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Publisher/SubmitToCT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherServer).SubmitToCT(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Publisher_SubmitToSingleCT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherServer).SubmitToSingleCT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Publisher/SubmitToSingleCT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherServer).SubmitToSingleCT(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Publisher_SubmitToSingleCTWithResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherServer).SubmitToSingleCTWithResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Publisher/SubmitToSingleCTWithResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherServer).SubmitToSingleCTWithResult(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Publisher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Publisher",
	HandlerType: (*PublisherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitToCT",
			Handler:    _Publisher_SubmitToCT_Handler,
		},
		{
			MethodName: "SubmitToSingleCT",
			Handler:    _Publisher_SubmitToSingleCT_Handler,
		},
		{
			MethodName: "SubmitToSingleCTWithResult",
			Handler:    _Publisher_SubmitToSingleCTWithResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "publisher.proto",
}

func init() { proto.RegisterFile("publisher.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xbb, 0xd6, 0x26, 0xe9, 0x10, 0xb0, 0xcc, 0x41, 0x96, 0x9c, 0xc2, 0x1e, 0x24, 0xa7,
	0x80, 0xfe, 0x85, 0xe2, 0xc9, 0x1c, 0xca, 0x26, 0xe2, 0xbd, 0x71, 0x48, 0x03, 0x5b, 0x37, 0xee,
	0x4e, 0x0e, 0xfd, 0x07, 0x5e, 0xfc, 0xcf, 0xd2, 0x35, 0x51, 0x11, 0x7a, 0xdb, 0xf7, 0xde, 0xb7,
	0xc3, 0x7b, 0x70, 0x33, 0x8c, 0x7b, 0xd3, 0xfb, 0x03, 0xb9, 0x72, 0x70, 0x96, 0xad, 0xfa, 0x14,
	0x10, 0x6b, 0x7a, 0x1f, 0xc9, 0x33, 0x6e, 0x60, 0xf9, 0x4a, 0x4e, 0x8a, 0x5c, 0x14, 0xa9, 0x3e,
	0x3f, 0xf1, 0x16, 0xa2, 0xca, 0x76, 0xcf, 0xba, 0x92, 0x57, 0xb9, 0x28, 0xd6, 0x3a, 0x32, 0x41,
	0xa1, 0x82, 0xb4, 0xb2, 0xdd, 0xee, 0x7c, 0xab, 0x7d, 0xa2, 0x93, 0x5c, 0x86, 0x34, 0x35, 0x7f,
	0x3c, 0x94, 0x10, 0x0f, 0x8e, 0x5a, 0x72, 0x2c, 0xaf, 0x73, 0x51, 0x24, 0x7a, 0x96, 0x98, 0x41,
	0xe2, 0xd9, 0x3a, 0xaa, 0xb7, 0x8d, 0x5c, 0x85, 0xe8, 0x47, 0xab, 0x0c, 0x22, 0x4d, 0x7e, 0x34,
	0xa1, 0x8d, 0x6f, 0x79, 0x6e, 0xe3, 0x5b, 0x56, 0x31, 0xac, 0x1e, 0x8f, 0x03, 0x9f, 0x1e, 0x3e,
	0x04, 0xac, 0x77, 0xf3, 0x10, 0xcc, 0x01, 0xea, 0x71, 0x7f, 0xec, 0xb9, 0xb1, 0xdb, 0x06, 0x93,
	0x72, 0x9a, 0x93, 0x45, 0x65, 0xa0, 0xd5, 0x02, 0xef, 0x60, 0x33, 0x13, 0x75, 0xff, 0xd6, 0x19,
	0xba, 0xc0, 0xdd, 0x43, 0xf6, 0x9f, 0x7b, 0xe9, 0xf9, 0x30, 0x15, 0xfa, 0xfd, 0x11, 0x97, 0xdf,
	0x96, 0x5a, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x91, 0x4d, 0x58, 0x51, 0x01, 0x00, 0x00,
}
