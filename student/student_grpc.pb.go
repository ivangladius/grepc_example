// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: student.proto

package __

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

// NetworkClient is the client API for Network service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkClient interface {
	WelcomeStudent(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Welcome, error)
}

type networkClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkClient(cc grpc.ClientConnInterface) NetworkClient {
	return &networkClient{cc}
}

func (c *networkClient) WelcomeStudent(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Welcome, error) {
	out := new(Welcome)
	err := c.cc.Invoke(ctx, "/student.Network/WelcomeStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServer is the server API for Network service.
// All implementations must embed UnimplementedNetworkServer
// for forward compatibility
type NetworkServer interface {
	WelcomeStudent(context.Context, *Student) (*Welcome, error)
	mustEmbedUnimplementedNetworkServer()
}

// UnimplementedNetworkServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkServer struct {
}

func (UnimplementedNetworkServer) WelcomeStudent(context.Context, *Student) (*Welcome, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WelcomeStudent not implemented")
}
func (UnimplementedNetworkServer) mustEmbedUnimplementedNetworkServer() {}

// UnsafeNetworkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkServer will
// result in compilation errors.
type UnsafeNetworkServer interface {
	mustEmbedUnimplementedNetworkServer()
}

func RegisterNetworkServer(s grpc.ServiceRegistrar, srv NetworkServer) {
	s.RegisterService(&Network_ServiceDesc, srv)
}

func _Network_WelcomeStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Student)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).WelcomeStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.Network/WelcomeStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).WelcomeStudent(ctx, req.(*Student))
	}
	return interceptor(ctx, in, info, handler)
}

// Network_ServiceDesc is the grpc.ServiceDesc for Network service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Network_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "student.Network",
	HandlerType: (*NetworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WelcomeStudent",
			Handler:    _Network_WelcomeStudent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "student.proto",
}