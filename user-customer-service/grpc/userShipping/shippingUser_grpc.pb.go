// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: shippingUser.proto

package shipping_user

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

// ShippingClient is the client API for Shipping service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShippingClient interface {
	AddShippingAddress(ctx context.Context, in *ShippingAddressAddRequest, opts ...grpc.CallOption) (*ShippingAddressAddResponse, error)
	GetShippingAddress(ctx context.Context, in *ShippingAddressRequest, opts ...grpc.CallOption) (*ShippingAddressResponse, error)
	DeleteShippingAddress(ctx context.Context, in *ShippingAddressRequest, opts ...grpc.CallOption) (*ShippingAddressDeleteResponse, error)
	UpdateShippingAddress(ctx context.Context, in *ShippingAddressUpdateRequest, opts ...grpc.CallOption) (*ShippingAddressUpdateResponse, error)
}

type shippingClient struct {
	cc grpc.ClientConnInterface
}

func NewShippingClient(cc grpc.ClientConnInterface) ShippingClient {
	return &shippingClient{cc}
}

func (c *shippingClient) AddShippingAddress(ctx context.Context, in *ShippingAddressAddRequest, opts ...grpc.CallOption) (*ShippingAddressAddResponse, error) {
	out := new(ShippingAddressAddResponse)
	err := c.cc.Invoke(ctx, "/Shipping/AddShippingAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingClient) GetShippingAddress(ctx context.Context, in *ShippingAddressRequest, opts ...grpc.CallOption) (*ShippingAddressResponse, error) {
	out := new(ShippingAddressResponse)
	err := c.cc.Invoke(ctx, "/Shipping/GetShippingAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingClient) DeleteShippingAddress(ctx context.Context, in *ShippingAddressRequest, opts ...grpc.CallOption) (*ShippingAddressDeleteResponse, error) {
	out := new(ShippingAddressDeleteResponse)
	err := c.cc.Invoke(ctx, "/Shipping/DeleteShippingAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingClient) UpdateShippingAddress(ctx context.Context, in *ShippingAddressUpdateRequest, opts ...grpc.CallOption) (*ShippingAddressUpdateResponse, error) {
	out := new(ShippingAddressUpdateResponse)
	err := c.cc.Invoke(ctx, "/Shipping/UpdateShippingAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServer is the server API for Shipping service.
// All implementations must embed UnimplementedShippingServer
// for forward compatibility
type ShippingServer interface {
	AddShippingAddress(context.Context, *ShippingAddressAddRequest) (*ShippingAddressAddResponse, error)
	GetShippingAddress(context.Context, *ShippingAddressRequest) (*ShippingAddressResponse, error)
	DeleteShippingAddress(context.Context, *ShippingAddressRequest) (*ShippingAddressDeleteResponse, error)
	UpdateShippingAddress(context.Context, *ShippingAddressUpdateRequest) (*ShippingAddressUpdateResponse, error)
	mustEmbedUnimplementedShippingServer()
}

// UnimplementedShippingServer must be embedded to have forward compatible implementations.
type UnimplementedShippingServer struct {
}

func (UnimplementedShippingServer) AddShippingAddress(context.Context, *ShippingAddressAddRequest) (*ShippingAddressAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddShippingAddress not implemented")
}
func (UnimplementedShippingServer) GetShippingAddress(context.Context, *ShippingAddressRequest) (*ShippingAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShippingAddress not implemented")
}
func (UnimplementedShippingServer) DeleteShippingAddress(context.Context, *ShippingAddressRequest) (*ShippingAddressDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShippingAddress not implemented")
}
func (UnimplementedShippingServer) UpdateShippingAddress(context.Context, *ShippingAddressUpdateRequest) (*ShippingAddressUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShippingAddress not implemented")
}
func (UnimplementedShippingServer) mustEmbedUnimplementedShippingServer() {}

// UnsafeShippingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShippingServer will
// result in compilation errors.
type UnsafeShippingServer interface {
	mustEmbedUnimplementedShippingServer()
}

func RegisterShippingServer(s grpc.ServiceRegistrar, srv ShippingServer) {
	s.RegisterService(&Shipping_ServiceDesc, srv)
}

func _Shipping_AddShippingAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShippingAddressAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServer).AddShippingAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shipping/AddShippingAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServer).AddShippingAddress(ctx, req.(*ShippingAddressAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shipping_GetShippingAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShippingAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServer).GetShippingAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shipping/GetShippingAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServer).GetShippingAddress(ctx, req.(*ShippingAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shipping_DeleteShippingAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShippingAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServer).DeleteShippingAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shipping/DeleteShippingAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServer).DeleteShippingAddress(ctx, req.(*ShippingAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shipping_UpdateShippingAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShippingAddressUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServer).UpdateShippingAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shipping/UpdateShippingAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServer).UpdateShippingAddress(ctx, req.(*ShippingAddressUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Shipping_ServiceDesc is the grpc.ServiceDesc for Shipping service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shipping_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Shipping",
	HandlerType: (*ShippingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddShippingAddress",
			Handler:    _Shipping_AddShippingAddress_Handler,
		},
		{
			MethodName: "GetShippingAddress",
			Handler:    _Shipping_GetShippingAddress_Handler,
		},
		{
			MethodName: "DeleteShippingAddress",
			Handler:    _Shipping_DeleteShippingAddress_Handler,
		},
		{
			MethodName: "UpdateShippingAddress",
			Handler:    _Shipping_UpdateShippingAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shippingUser.proto",
}
