// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0
// source: hello_world.proto

package v1

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// The HTTPBody messages allow a response message to be specified with custom data content and a
	// custom content-type header. The values included in the HTTPBody response will be used verbatim
	// in the returned message from the gateway. Make sure you format your response carefully!
	Hello1(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	Download(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Greeter_DownloadClient, error)
	UpdateV2(ctx context.Context, in *UpdateV2Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateV2A(ctx context.Context, in *UpdateV2Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetMessage
	// It is possible to define multiple HTTP methods for one RPC by using the additional_bindings option.
	GetMessage(ctx context.Context, in *GetMessageReq, opts ...grpc.CallOption) (*GetMessageResp, error)
	UpdateMessage(ctx context.Context, in *UpdateMessageRequest, opts ...grpc.CallOption) (*UpdateMessageResp, error)
	// UpdateMessage1
	// The special name * can be used in the body mapping to define that every
	// field not bound by the path template should be mapped to the request body.
	// This enables the following alternative definition of the update method:
	UpdateMessage1(ctx context.Context, in *UpdateMessage1Request, opts ...grpc.CallOption) (*UpdateMessage1Resp, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/hello.v1.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Hello1(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/hello.v1.Greeter/Hello1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Download(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Greeter_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[0], "/hello.v1.Greeter/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_DownloadClient interface {
	Recv() (*httpbody.HttpBody, error)
	grpc.ClientStream
}

type greeterDownloadClient struct {
	grpc.ClientStream
}

func (x *greeterDownloadClient) Recv() (*httpbody.HttpBody, error) {
	m := new(httpbody.HttpBody)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) UpdateV2(ctx context.Context, in *UpdateV2Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/hello.v1.Greeter/UpdateV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) UpdateV2A(ctx context.Context, in *UpdateV2Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/hello.v1.Greeter/UpdateV2a", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetMessage(ctx context.Context, in *GetMessageReq, opts ...grpc.CallOption) (*GetMessageResp, error) {
	out := new(GetMessageResp)
	err := c.cc.Invoke(ctx, "/hello.v1.Greeter/GetMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) UpdateMessage(ctx context.Context, in *UpdateMessageRequest, opts ...grpc.CallOption) (*UpdateMessageResp, error) {
	out := new(UpdateMessageResp)
	err := c.cc.Invoke(ctx, "/hello.v1.Greeter/UpdateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) UpdateMessage1(ctx context.Context, in *UpdateMessage1Request, opts ...grpc.CallOption) (*UpdateMessage1Resp, error) {
	out := new(UpdateMessage1Resp)
	err := c.cc.Invoke(ctx, "/hello.v1.Greeter/UpdateMessage1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	// The HTTPBody messages allow a response message to be specified with custom data content and a
	// custom content-type header. The values included in the HTTPBody response will be used verbatim
	// in the returned message from the gateway. Make sure you format your response carefully!
	Hello1(context.Context, *emptypb.Empty) (*httpbody.HttpBody, error)
	Download(*emptypb.Empty, Greeter_DownloadServer) error
	UpdateV2(context.Context, *UpdateV2Request) (*emptypb.Empty, error)
	UpdateV2A(context.Context, *UpdateV2Request) (*emptypb.Empty, error)
	// GetMessage
	// It is possible to define multiple HTTP methods for one RPC by using the additional_bindings option.
	GetMessage(context.Context, *GetMessageReq) (*GetMessageResp, error)
	UpdateMessage(context.Context, *UpdateMessageRequest) (*UpdateMessageResp, error)
	// UpdateMessage1
	// The special name * can be used in the body mapping to define that every
	// field not bound by the path template should be mapped to the request body.
	// This enables the following alternative definition of the update method:
	UpdateMessage1(context.Context, *UpdateMessage1Request) (*UpdateMessage1Resp, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) Hello1(context.Context, *emptypb.Empty) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello1 not implemented")
}
func (UnimplementedGreeterServer) Download(*emptypb.Empty, Greeter_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedGreeterServer) UpdateV2(context.Context, *UpdateV2Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateV2 not implemented")
}
func (UnimplementedGreeterServer) UpdateV2A(context.Context, *UpdateV2Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateV2A not implemented")
}
func (UnimplementedGreeterServer) GetMessage(context.Context, *GetMessageReq) (*GetMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedGreeterServer) UpdateMessage(context.Context, *UpdateMessageRequest) (*UpdateMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage not implemented")
}
func (UnimplementedGreeterServer) UpdateMessage1(context.Context, *UpdateMessage1Request) (*UpdateMessage1Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage1 not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.v1.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Hello1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Hello1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.v1.Greeter/Hello1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Hello1(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).Download(m, &greeterDownloadServer{stream})
}

type Greeter_DownloadServer interface {
	Send(*httpbody.HttpBody) error
	grpc.ServerStream
}

type greeterDownloadServer struct {
	grpc.ServerStream
}

func (x *greeterDownloadServer) Send(m *httpbody.HttpBody) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_UpdateV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateV2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).UpdateV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.v1.Greeter/UpdateV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).UpdateV2(ctx, req.(*UpdateV2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_UpdateV2A_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateV2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).UpdateV2A(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.v1.Greeter/UpdateV2a",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).UpdateV2A(ctx, req.(*UpdateV2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.v1.Greeter/GetMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetMessage(ctx, req.(*GetMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_UpdateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).UpdateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.v1.Greeter/UpdateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).UpdateMessage(ctx, req.(*UpdateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_UpdateMessage1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessage1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).UpdateMessage1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.v1.Greeter/UpdateMessage1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).UpdateMessage1(ctx, req.(*UpdateMessage1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.v1.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "Hello1",
			Handler:    _Greeter_Hello1_Handler,
		},
		{
			MethodName: "UpdateV2",
			Handler:    _Greeter_UpdateV2_Handler,
		},
		{
			MethodName: "UpdateV2a",
			Handler:    _Greeter_UpdateV2A_Handler,
		},
		{
			MethodName: "GetMessage",
			Handler:    _Greeter_GetMessage_Handler,
		},
		{
			MethodName: "UpdateMessage",
			Handler:    _Greeter_UpdateMessage_Handler,
		},
		{
			MethodName: "UpdateMessage1",
			Handler:    _Greeter_UpdateMessage1_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Download",
			Handler:       _Greeter_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hello_world.proto",
}