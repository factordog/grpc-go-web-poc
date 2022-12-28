// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: time.proto

package grpctime

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

// TimeServiceClient is the client API for TimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeServiceClient interface {
	GetCurrentTime(ctx context.Context, in *GetCurrentTimeRequest, opts ...grpc.CallOption) (*GetCurrentTimeResponse, error)
	GetGreet(ctx context.Context, in *GetGreetRequest, opts ...grpc.CallOption) (*GetGreetResponse, error)
}

type timeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeServiceClient(cc grpc.ClientConnInterface) TimeServiceClient {
	return &timeServiceClient{cc}
}

func (c *timeServiceClient) GetCurrentTime(ctx context.Context, in *GetCurrentTimeRequest, opts ...grpc.CallOption) (*GetCurrentTimeResponse, error) {
	out := new(GetCurrentTimeResponse)
	err := c.cc.Invoke(ctx, "/smpl.time.api.v1.TimeService/GetCurrentTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeServiceClient) GetGreet(ctx context.Context, in *GetGreetRequest, opts ...grpc.CallOption) (*GetGreetResponse, error) {
	out := new(GetGreetResponse)
	err := c.cc.Invoke(ctx, "/smpl.time.api.v1.TimeService/GetGreet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeServiceServer is the server API for TimeService service.
// All implementations must embed UnimplementedTimeServiceServer
// for forward compatibility
type TimeServiceServer interface {
	GetCurrentTime(context.Context, *GetCurrentTimeRequest) (*GetCurrentTimeResponse, error)
	GetGreet(context.Context, *GetGreetRequest) (*GetGreetResponse, error)
	mustEmbedUnimplementedTimeServiceServer()
}

// UnimplementedTimeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimeServiceServer struct {
}

func (UnimplementedTimeServiceServer) GetCurrentTime(context.Context, *GetCurrentTimeRequest) (*GetCurrentTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentTime not implemented")
}
func (UnimplementedTimeServiceServer) GetGreet(context.Context, *GetGreetRequest) (*GetGreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGreet not implemented")
}
func (UnimplementedTimeServiceServer) mustEmbedUnimplementedTimeServiceServer() {}

// UnsafeTimeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeServiceServer will
// result in compilation errors.
type UnsafeTimeServiceServer interface {
	mustEmbedUnimplementedTimeServiceServer()
}

func RegisterTimeServiceServer(s grpc.ServiceRegistrar, srv TimeServiceServer) {
	s.RegisterService(&TimeService_ServiceDesc, srv)
}

func _TimeService_GetCurrentTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServiceServer).GetCurrentTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smpl.time.api.v1.TimeService/GetCurrentTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServiceServer).GetCurrentTime(ctx, req.(*GetCurrentTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeService_GetGreet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServiceServer).GetGreet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smpl.time.api.v1.TimeService/GetGreet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServiceServer).GetGreet(ctx, req.(*GetGreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TimeService_ServiceDesc is the grpc.ServiceDesc for TimeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smpl.time.api.v1.TimeService",
	HandlerType: (*TimeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentTime",
			Handler:    _TimeService_GetCurrentTime_Handler,
		},
		{
			MethodName: "GetGreet",
			Handler:    _TimeService_GetGreet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "time.proto",
}
