// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// ServiceRunnerClient is the client API for ServiceRunner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceRunnerClient interface {
	DescribePackage(ctx context.Context, in *DescribePackageRequest, opts ...grpc.CallOption) (*DescribePackageReply, error)
	DeletePackage(ctx context.Context, in *DeletePackageRequest, opts ...grpc.CallOption) (*DeletePackageReply, error)
	DeployPackage(ctx context.Context, in *DeployPackageRequest, opts ...grpc.CallOption) (*DeployPackageReply, error)
}

type serviceRunnerClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceRunnerClient(cc grpc.ClientConnInterface) ServiceRunnerClient {
	return &serviceRunnerClient{cc}
}

func (c *serviceRunnerClient) DescribePackage(ctx context.Context, in *DescribePackageRequest, opts ...grpc.CallOption) (*DescribePackageReply, error) {
	out := new(DescribePackageReply)
	err := c.cc.Invoke(ctx, "/ServiceRunner/DescribePackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceRunnerClient) DeletePackage(ctx context.Context, in *DeletePackageRequest, opts ...grpc.CallOption) (*DeletePackageReply, error) {
	out := new(DeletePackageReply)
	err := c.cc.Invoke(ctx, "/ServiceRunner/DeletePackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceRunnerClient) DeployPackage(ctx context.Context, in *DeployPackageRequest, opts ...grpc.CallOption) (*DeployPackageReply, error) {
	out := new(DeployPackageReply)
	err := c.cc.Invoke(ctx, "/ServiceRunner/DeployPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceRunnerServer is the server API for ServiceRunner service.
// All implementations must embed UnimplementedServiceRunnerServer
// for forward compatibility
type ServiceRunnerServer interface {
	DescribePackage(context.Context, *DescribePackageRequest) (*DescribePackageReply, error)
	DeletePackage(context.Context, *DeletePackageRequest) (*DeletePackageReply, error)
	DeployPackage(context.Context, *DeployPackageRequest) (*DeployPackageReply, error)
	mustEmbedUnimplementedServiceRunnerServer()
}

// UnimplementedServiceRunnerServer must be embedded to have forward compatible implementations.
type UnimplementedServiceRunnerServer struct {
}

func (UnimplementedServiceRunnerServer) DescribePackage(context.Context, *DescribePackageRequest) (*DescribePackageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePackage not implemented")
}
func (UnimplementedServiceRunnerServer) DeletePackage(context.Context, *DeletePackageRequest) (*DeletePackageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePackage not implemented")
}
func (UnimplementedServiceRunnerServer) DeployPackage(context.Context, *DeployPackageRequest) (*DeployPackageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployPackage not implemented")
}
func (UnimplementedServiceRunnerServer) mustEmbedUnimplementedServiceRunnerServer() {}

// UnsafeServiceRunnerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceRunnerServer will
// result in compilation errors.
type UnsafeServiceRunnerServer interface {
	mustEmbedUnimplementedServiceRunnerServer()
}

func RegisterServiceRunnerServer(s grpc.ServiceRegistrar, srv ServiceRunnerServer) {
	s.RegisterService(&ServiceRunner_ServiceDesc, srv)
}

func _ServiceRunner_DescribePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceRunnerServer).DescribePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServiceRunner/DescribePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceRunnerServer).DescribePackage(ctx, req.(*DescribePackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceRunner_DeletePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceRunnerServer).DeletePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServiceRunner/DeletePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceRunnerServer).DeletePackage(ctx, req.(*DeletePackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceRunner_DeployPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceRunnerServer).DeployPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServiceRunner/DeployPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceRunnerServer).DeployPackage(ctx, req.(*DeployPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceRunner_ServiceDesc is the grpc.ServiceDesc for ServiceRunner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceRunner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ServiceRunner",
	HandlerType: (*ServiceRunnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribePackage",
			Handler:    _ServiceRunner_DescribePackage_Handler,
		},
		{
			MethodName: "DeletePackage",
			Handler:    _ServiceRunner_DeletePackage_Handler,
		},
		{
			MethodName: "DeployPackage",
			Handler:    _ServiceRunner_DeployPackage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/sr.proto",
}
