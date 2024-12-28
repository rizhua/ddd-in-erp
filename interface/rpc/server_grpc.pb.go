// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: server.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Rizhua_Dial_FullMethodName = "/rpc.Rizhua/Dial"
)

// RizhuaClient is the client API for Rizhua service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RizhuaClient interface {
	Dial(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Reply, error)
}

type rizhuaClient struct {
	cc grpc.ClientConnInterface
}

func NewRizhuaClient(cc grpc.ClientConnInterface) RizhuaClient {
	return &rizhuaClient{cc}
}

func (c *rizhuaClient) Dial(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Reply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Reply)
	err := c.cc.Invoke(ctx, Rizhua_Dial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RizhuaServer is the server API for Rizhua service.
// All implementations should embed UnimplementedRizhuaServer
// for forward compatibility.
type RizhuaServer interface {
	Dial(context.Context, *Query) (*Reply, error)
}

// UnimplementedRizhuaServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRizhuaServer struct{}

func (UnimplementedRizhuaServer) Dial(context.Context, *Query) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dial not implemented")
}
func (UnimplementedRizhuaServer) testEmbeddedByValue() {}

// UnsafeRizhuaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RizhuaServer will
// result in compilation errors.
type UnsafeRizhuaServer interface {
	mustEmbedUnimplementedRizhuaServer()
}

func RegisterRizhuaServer(s grpc.ServiceRegistrar, srv RizhuaServer) {
	// If the following call pancis, it indicates UnimplementedRizhuaServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Rizhua_ServiceDesc, srv)
}

func _Rizhua_Dial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RizhuaServer).Dial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rizhua_Dial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RizhuaServer).Dial(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

// Rizhua_ServiceDesc is the grpc.ServiceDesc for Rizhua service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rizhua_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Rizhua",
	HandlerType: (*RizhuaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Dial",
			Handler:    _Rizhua_Dial_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
