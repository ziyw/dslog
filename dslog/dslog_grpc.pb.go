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
	AddLog(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
}

type dslogClient struct {
	cc grpc.ClientConnInterface
}

func NewDslogClient(cc grpc.ClientConnInterface) DslogClient {
	return &dslogClient{cc}
}

func (c *dslogClient) AddLog(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/dslog.Dslog/AddLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DslogServer is the server API for Dslog service.
// All implementations must embed UnimplementedDslogServer
// for forward compatibility
type DslogServer interface {
	AddLog(context.Context, *LogRequest) (*LogResponse, error)
	mustEmbedUnimplementedDslogServer()
}

// UnimplementedDslogServer must be embedded to have forward compatible implementations.
type UnimplementedDslogServer struct {
}

func (UnimplementedDslogServer) AddLog(context.Context, *LogRequest) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLog not implemented")
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

func _Dslog_AddLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DslogServer).AddLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dslog.Dslog/AddLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DslogServer).AddLog(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dslog_ServiceDesc is the grpc.ServiceDesc for Dslog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dslog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dslog.Dslog",
	HandlerType: (*DslogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddLog",
			Handler:    _Dslog_AddLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dslog/dslog.proto",
}
