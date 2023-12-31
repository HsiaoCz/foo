// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: pb/pv1/foo.proto

package pv1

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

const (
	Foo_UserSignup_FullMethodName = "/pv1.Foo/UserSignup"
	Foo_UserLogin_FullMethodName  = "/pv1.Foo/UserLogin"
)

// FooClient is the client API for Foo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FooClient interface {
	// user api
	UserSignup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	UserLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type fooClient struct {
	cc grpc.ClientConnInterface
}

func NewFooClient(cc grpc.ClientConnInterface) FooClient {
	return &fooClient{cc}
}

func (c *fooClient) UserSignup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, Foo_UserSignup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fooClient) UserLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, Foo_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FooServer is the server API for Foo service.
// All implementations must embed UnimplementedFooServer
// for forward compatibility
type FooServer interface {
	// user api
	UserSignup(context.Context, *SignupRequest) (*SignupResponse, error)
	UserLogin(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedFooServer()
}

// UnimplementedFooServer must be embedded to have forward compatible implementations.
type UnimplementedFooServer struct {
}

func (UnimplementedFooServer) UserSignup(context.Context, *SignupRequest) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignup not implemented")
}
func (UnimplementedFooServer) UserLogin(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedFooServer) mustEmbedUnimplementedFooServer() {}

// UnsafeFooServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FooServer will
// result in compilation errors.
type UnsafeFooServer interface {
	mustEmbedUnimplementedFooServer()
}

func RegisterFooServer(s grpc.ServiceRegistrar, srv FooServer) {
	s.RegisterService(&Foo_ServiceDesc, srv)
}

func _Foo_UserSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FooServer).UserSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Foo_UserSignup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FooServer).UserSignup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Foo_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FooServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Foo_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FooServer).UserLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Foo_ServiceDesc is the grpc.ServiceDesc for Foo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Foo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pv1.Foo",
	HandlerType: (*FooServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserSignup",
			Handler:    _Foo_UserSignup_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _Foo_UserLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/pv1/foo.proto",
}
