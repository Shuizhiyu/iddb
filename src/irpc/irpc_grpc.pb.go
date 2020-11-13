// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package irpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// PushTableClient is the client API for PushTable service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushTableClient interface {
	PushTable(ctx context.Context, in *Table, opts ...grpc.CallOption) (*IrpcStatus, error)
}

type pushTableClient struct {
	cc grpc.ClientConnInterface
}

func NewPushTableClient(cc grpc.ClientConnInterface) PushTableClient {
	return &pushTableClient{cc}
}

func (c *pushTableClient) PushTable(ctx context.Context, in *Table, opts ...grpc.CallOption) (*IrpcStatus, error) {
	out := new(IrpcStatus)
	err := c.cc.Invoke(ctx, "/irpc.PushTable/PushTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushTableServer is the server API for PushTable service.
// All implementations must embed UnimplementedPushTableServer
// for forward compatibility
type PushTableServer interface {
	PushTable(context.Context, *Table) (*IrpcStatus, error)
	mustEmbedUnimplementedPushTableServer()
}

// UnimplementedPushTableServer must be embedded to have forward compatible implementations.
type UnimplementedPushTableServer struct {
}

func (UnimplementedPushTableServer) PushTable(context.Context, *Table) (*IrpcStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushTable not implemented")
}
func (UnimplementedPushTableServer) mustEmbedUnimplementedPushTableServer() {}

// UnsafePushTableServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushTableServer will
// result in compilation errors.
type UnsafePushTableServer interface {
	mustEmbedUnimplementedPushTableServer()
}

func RegisterPushTableServer(s *grpc.Server, srv PushTableServer) {
	s.RegisterService(&_PushTable_serviceDesc, srv)
}

func _PushTable_PushTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Table)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushTableServer).PushTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/irpc.PushTable/PushTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushTableServer).PushTable(ctx, req.(*Table))
	}
	return interceptor(ctx, in, info, handler)
}

var _PushTable_serviceDesc = grpc.ServiceDesc{
	ServiceName: "irpc.PushTable",
	HandlerType: (*PushTableServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushTable",
			Handler:    _PushTable_PushTable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "irpc.proto",
}

// ExecuterCallerClient is the client API for ExecuterCaller service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecuterCallerClient interface {
	ExecuterCall(ctx context.Context, in *IrpcCallReq, opts ...grpc.CallOption) (*IrpcStatus, error)
}

type executerCallerClient struct {
	cc grpc.ClientConnInterface
}

func NewExecuterCallerClient(cc grpc.ClientConnInterface) ExecuterCallerClient {
	return &executerCallerClient{cc}
}

func (c *executerCallerClient) ExecuterCall(ctx context.Context, in *IrpcCallReq, opts ...grpc.CallOption) (*IrpcStatus, error) {
	out := new(IrpcStatus)
	err := c.cc.Invoke(ctx, "/irpc.ExecuterCaller/ExecuterCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecuterCallerServer is the server API for ExecuterCaller service.
// All implementations must embed UnimplementedExecuterCallerServer
// for forward compatibility
type ExecuterCallerServer interface {
	ExecuterCall(context.Context, *IrpcCallReq) (*IrpcStatus, error)
	mustEmbedUnimplementedExecuterCallerServer()
}

// UnimplementedExecuterCallerServer must be embedded to have forward compatible implementations.
type UnimplementedExecuterCallerServer struct {
}

func (UnimplementedExecuterCallerServer) ExecuterCall(context.Context, *IrpcCallReq) (*IrpcStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuterCall not implemented")
}
func (UnimplementedExecuterCallerServer) mustEmbedUnimplementedExecuterCallerServer() {}

// UnsafeExecuterCallerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExecuterCallerServer will
// result in compilation errors.
type UnsafeExecuterCallerServer interface {
	mustEmbedUnimplementedExecuterCallerServer()
}

func RegisterExecuterCallerServer(s *grpc.Server, srv ExecuterCallerServer) {
	s.RegisterService(&_ExecuterCaller_serviceDesc, srv)
}

func _ExecuterCaller_ExecuterCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IrpcCallReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecuterCallerServer).ExecuterCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/irpc.ExecuterCaller/ExecuterCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecuterCallerServer).ExecuterCall(ctx, req.(*IrpcCallReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExecuterCaller_serviceDesc = grpc.ServiceDesc{
	ServiceName: "irpc.ExecuterCaller",
	HandlerType: (*ExecuterCallerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuterCall",
			Handler:    _ExecuterCaller_ExecuterCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "irpc.proto",
}
