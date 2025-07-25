// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: order.proto

package order

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
	Order_FindOrderHistory_FullMethodName = "/order.Order/FindOrderHistory"
	Order_FindOrderCurrent_FullMethodName = "/order.Order/FindOrderCurrent"
	Order_Add_FullMethodName              = "/order.Order/Add"
	Order_FindByOrderId_FullMethodName    = "/order.Order/FindByOrderId"
	Order_CancelOrder_FullMethodName      = "/order.Order/CancelOrder"
)

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	FindOrderHistory(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error)
	FindOrderCurrent(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error)
	Add(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*AddOrderRes, error)
	FindByOrderId(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*ExchangeOrderOrigin, error)
	CancelOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*CancelOrderRes, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) FindOrderHistory(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderRes)
	err := c.cc.Invoke(ctx, Order_FindOrderHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) FindOrderCurrent(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderRes)
	err := c.cc.Invoke(ctx, Order_FindOrderCurrent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) Add(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*AddOrderRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddOrderRes)
	err := c.cc.Invoke(ctx, Order_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) FindByOrderId(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*ExchangeOrderOrigin, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExchangeOrderOrigin)
	err := c.cc.Invoke(ctx, Order_FindByOrderId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CancelOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*CancelOrderRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelOrderRes)
	err := c.cc.Invoke(ctx, Order_CancelOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility.
type OrderServer interface {
	FindOrderHistory(context.Context, *OrderReq) (*OrderRes, error)
	FindOrderCurrent(context.Context, *OrderReq) (*OrderRes, error)
	Add(context.Context, *OrderReq) (*AddOrderRes, error)
	FindByOrderId(context.Context, *OrderReq) (*ExchangeOrderOrigin, error)
	CancelOrder(context.Context, *OrderReq) (*CancelOrderRes, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrderServer struct{}

func (UnimplementedOrderServer) FindOrderHistory(context.Context, *OrderReq) (*OrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOrderHistory not implemented")
}
func (UnimplementedOrderServer) FindOrderCurrent(context.Context, *OrderReq) (*OrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOrderCurrent not implemented")
}
func (UnimplementedOrderServer) Add(context.Context, *OrderReq) (*AddOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedOrderServer) FindByOrderId(context.Context, *OrderReq) (*ExchangeOrderOrigin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByOrderId not implemented")
}
func (UnimplementedOrderServer) CancelOrder(context.Context, *OrderReq) (*CancelOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}
func (UnimplementedOrderServer) testEmbeddedByValue()               {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	// If the following call pancis, it indicates UnimplementedOrderServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_FindOrderHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).FindOrderHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_FindOrderHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).FindOrderHistory(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_FindOrderCurrent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).FindOrderCurrent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_FindOrderCurrent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).FindOrderCurrent(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).Add(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_FindByOrderId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).FindByOrderId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_FindByOrderId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).FindByOrderId(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CancelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CancelOrder(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindOrderHistory",
			Handler:    _Order_FindOrderHistory_Handler,
		},
		{
			MethodName: "FindOrderCurrent",
			Handler:    _Order_FindOrderCurrent_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Order_Add_Handler,
		},
		{
			MethodName: "FindByOrderId",
			Handler:    _Order_FindByOrderId_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _Order_CancelOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
