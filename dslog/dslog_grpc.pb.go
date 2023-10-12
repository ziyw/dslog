// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: dslog/dslog.proto

package dslog

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DslogClient is the client API for Dslog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DslogClient interface {
	Send(ctx context.Context, in *LogMessage, opts ...grpc.CallOption) (*SendResponse, error)
	GetByTimeRange(ctx context.Context, in *TimeRange, opts ...grpc.CallOption) (Dslog_GetByTimeRangeClient, error)
	GetError(ctx context.Context, in *TimeRange, opts ...grpc.CallOption) (Dslog_GetErrorClient, error)
}

type dslogClient struct {
	cc grpc.ClientConnInterface
}

func NewDslogClient(cc grpc.ClientConnInterface) DslogClient {
	return &dslogClient{cc}
}

func (c *dslogClient) Send(ctx context.Context, in *LogMessage, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/dslog.Dslog/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dslogClient) GetByTimeRange(ctx context.Context, in *TimeRange, opts ...grpc.CallOption) (Dslog_GetByTimeRangeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Dslog_ServiceDesc.Streams[0], "/dslog.Dslog/GetByTimeRange", opts...)
	if err != nil {
		return nil, err
	}
	x := &dslogGetByTimeRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dslog_GetByTimeRangeClient interface {
	Recv() (*LogMessage, error)
	grpc.ClientStream
}

type dslogGetByTimeRangeClient struct {
	grpc.ClientStream
}

func (x *dslogGetByTimeRangeClient) Recv() (*LogMessage, error) {
	m := new(LogMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dslogClient) GetError(ctx context.Context, in *TimeRange, opts ...grpc.CallOption) (Dslog_GetErrorClient, error) {
	stream, err := c.cc.NewStream(ctx, &Dslog_ServiceDesc.Streams[1], "/dslog.Dslog/GetError", opts...)
	if err != nil {
		return nil, err
	}
	x := &dslogGetErrorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dslog_GetErrorClient interface {
	Recv() (*LogMessage, error)
	grpc.ClientStream
}

type dslogGetErrorClient struct {
	grpc.ClientStream
}

func (x *dslogGetErrorClient) Recv() (*LogMessage, error) {
	m := new(LogMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DslogServer is the server API for Dslog service.
// All implementations must embed UnimplementedDslogServer
// for forward compatibility
type DslogServer interface {
	Send(context.Context, *LogMessage) (*SendResponse, error)
	GetByTimeRange(*TimeRange, Dslog_GetByTimeRangeServer) error
	GetError(*TimeRange, Dslog_GetErrorServer) error
	mustEmbedUnimplementedDslogServer()
}

// UnimplementedDslogServer must be embedded to have forward compatible implementations.
type UnimplementedDslogServer struct {
}

func (UnimplementedDslogServer) Send(context.Context, *LogMessage) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedDslogServer) GetByTimeRange(*TimeRange, Dslog_GetByTimeRangeServer) error {
	return status.Errorf(codes.Unimplemented, "method GetByTimeRange not implemented")
}
func (UnimplementedDslogServer) GetError(*TimeRange, Dslog_GetErrorServer) error {
	return status.Errorf(codes.Unimplemented, "method GetError not implemented")
}
func (UnimplementedDslogServer) mustEmbedUnimplementedDslogServer() {}

// UnsafeDslogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DslogServer will
// result in compilation errors.
type UnsafeDslogServer interface {
	mustEmbedUnimplementedDslogServer()
}

func RegisterDslogServer(s grpc.ServiceRegistrar, srv DslogServer) {
	s.RegisterService(&Dslog_ServiceDesc, srv)
}

func _Dslog_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DslogServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dslog.Dslog/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DslogServer).Send(ctx, req.(*LogMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dslog_GetByTimeRange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TimeRange)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DslogServer).GetByTimeRange(m, &dslogGetByTimeRangeServer{stream})
}

type Dslog_GetByTimeRangeServer interface {
	Send(*LogMessage) error
	grpc.ServerStream
}

type dslogGetByTimeRangeServer struct {
	grpc.ServerStream
}

func (x *dslogGetByTimeRangeServer) Send(m *LogMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Dslog_GetError_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TimeRange)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DslogServer).GetError(m, &dslogGetErrorServer{stream})
}

type Dslog_GetErrorServer interface {
	Send(*LogMessage) error
	grpc.ServerStream
}

type dslogGetErrorServer struct {
	grpc.ServerStream
}

func (x *dslogGetErrorServer) Send(m *LogMessage) error {
	return x.ServerStream.SendMsg(m)
}

// Dslog_ServiceDesc is the grpc.ServiceDesc for Dslog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dslog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dslog.Dslog",
	HandlerType: (*DslogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Dslog_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetByTimeRange",
			Handler:       _Dslog_GetByTimeRange_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetError",
			Handler:       _Dslog_GetError_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dslog/dslog.proto",
}
