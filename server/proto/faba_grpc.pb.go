// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// FabaServiceClient is the client API for FabaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FabaServiceClient interface {
	RegisterEmployee(ctx context.Context, in *RegisterEmployeeRequest, opts ...grpc.CallOption) (*RegisterEmployeeResponse, error)
}

type fabaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFabaServiceClient(cc grpc.ClientConnInterface) FabaServiceClient {
	return &fabaServiceClient{cc}
}

func (c *fabaServiceClient) RegisterEmployee(ctx context.Context, in *RegisterEmployeeRequest, opts ...grpc.CallOption) (*RegisterEmployeeResponse, error) {
	out := new(RegisterEmployeeResponse)
	err := c.cc.Invoke(ctx, "/proto.FabaService/RegisterEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FabaServiceServer is the server API for FabaService service.
// All implementations must embed UnimplementedFabaServiceServer
// for forward compatibility
type FabaServiceServer interface {
	RegisterEmployee(context.Context, *RegisterEmployeeRequest) (*RegisterEmployeeResponse, error)
	mustEmbedUnimplementedFabaServiceServer()
}

// UnimplementedFabaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFabaServiceServer struct {
}

func (UnimplementedFabaServiceServer) RegisterEmployee(context.Context, *RegisterEmployeeRequest) (*RegisterEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterEmployee not implemented")
}
func (UnimplementedFabaServiceServer) mustEmbedUnimplementedFabaServiceServer() {}

// UnsafeFabaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FabaServiceServer will
// result in compilation errors.
type UnsafeFabaServiceServer interface {
	mustEmbedUnimplementedFabaServiceServer()
}

func RegisterFabaServiceServer(s grpc.ServiceRegistrar, srv FabaServiceServer) {
	s.RegisterService(&FabaService_ServiceDesc, srv)
}

func _FabaService_RegisterEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FabaServiceServer).RegisterEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FabaService/RegisterEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FabaServiceServer).RegisterEmployee(ctx, req.(*RegisterEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FabaService_ServiceDesc is the grpc.ServiceDesc for FabaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FabaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FabaService",
	HandlerType: (*FabaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterEmployee",
			Handler:    _FabaService_RegisterEmployee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "faba.proto",
}
