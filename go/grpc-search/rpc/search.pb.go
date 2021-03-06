// Code generated by protoc-gen-go.
// source: search.proto
// DO NOT EDIT!

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	search.proto

It has these top-level messages:
	Request
	Result
*/
package rpc

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
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type Result struct {
	Title   string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Url     string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Snippet string `protobuf:"bytes,3,opt,name=snippet" json:"snippet,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Result) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Result) GetSnippet() string {
	if m != nil {
		return m.Snippet
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "rpc.Request")
	proto.RegisterType((*Result)(nil), "rpc.Result")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Google service

type GoogleClient interface {
	Search(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
	Watch(ctx context.Context, in *Request, opts ...grpc.CallOption) (Google_WatchClient, error)
}

type googleClient struct {
	cc *grpc.ClientConn
}

func NewGoogleClient(cc *grpc.ClientConn) GoogleClient {
	return &googleClient{cc}
}

func (c *googleClient) Search(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/rpc.Google/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleClient) Watch(ctx context.Context, in *Request, opts ...grpc.CallOption) (Google_WatchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Google_serviceDesc.Streams[0], c.cc, "/rpc.Google/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &googleWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Google_WatchClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type googleWatchClient struct {
	grpc.ClientStream
}

func (x *googleWatchClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Google service

type GoogleServer interface {
	Search(context.Context, *Request) (*Result, error)
	Watch(*Request, Google_WatchServer) error
}

func RegisterGoogleServer(s *grpc.Server, srv GoogleServer) {
	s.RegisterService(&_Google_serviceDesc, srv)
}

func _Google_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Google/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleServer).Search(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Google_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoogleServer).Watch(m, &googleWatchServer{stream})
}

type Google_WatchServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type googleWatchServer struct {
	grpc.ServerStream
}

func (x *googleWatchServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

var _Google_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Google",
	HandlerType: (*GoogleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Google_Search_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Google_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "search.proto",
}

func init() { proto.RegisterFile("search.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x4d, 0x2c,
	0x4a, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x56, 0x92, 0xe7,
	0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0x2c, 0x4d, 0x2d,
	0xaa, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0xbc, 0xb8, 0xd8, 0x82, 0x52,
	0x8b, 0x4b, 0x73, 0xc0, 0xf2, 0x25, 0x99, 0x25, 0x39, 0xa9, 0x30, 0x79, 0x30, 0x47, 0x48, 0x80,
	0x8b, 0xb9, 0xb4, 0x28, 0x47, 0x82, 0x09, 0x2c, 0x06, 0x62, 0x0a, 0x49, 0x70, 0xb1, 0x17, 0xe7,
	0x65, 0x16, 0x14, 0xa4, 0x96, 0x48, 0x30, 0x83, 0x45, 0x61, 0x5c, 0xa3, 0x70, 0x2e, 0x36, 0xf7,
	0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x21, 0x55, 0x2e, 0xb6, 0x60, 0xb0, 0x5b, 0x84, 0x78, 0xf4, 0x8a,
	0x0a, 0x92, 0xf5, 0xa0, 0x6e, 0x90, 0xe2, 0x86, 0xf2, 0x40, 0x16, 0x2a, 0x31, 0x08, 0xa9, 0x71,
	0xb1, 0x86, 0x27, 0x96, 0x10, 0x50, 0x65, 0xc0, 0x98, 0xc4, 0x06, 0xf6, 0x91, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x0f, 0xf3, 0x0b, 0xd7, 0xe1, 0x00, 0x00, 0x00,
}
