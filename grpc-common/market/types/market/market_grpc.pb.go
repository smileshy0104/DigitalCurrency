// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: market.proto

package market

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
	Market_FindSymbolThumbTrend_FullMethodName    = "/market.Market/FindSymbolThumbTrend"
	Market_FindSymbolInfo_FullMethodName          = "/market.Market/FindSymbolInfo"
	Market_FindCoinInfo_FullMethodName            = "/market.Market/FindCoinInfo"
	Market_FindAllCoin_FullMethodName             = "/market.Market/FindAllCoin"
	Market_HistoryKline_FullMethodName            = "/market.Market/HistoryKline"
	Market_FindExchangeCoinVisible_FullMethodName = "/market.Market/FindExchangeCoinVisible"
	Market_FindCoinById_FullMethodName            = "/market.Market/FindCoinById"
)

// MarketClient is the client API for Market service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketClient interface {
	FindSymbolThumbTrend(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*SymbolThumbRes, error)
	FindSymbolInfo(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*ExchangeCoin, error)
	FindCoinInfo(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*Coin, error)
	FindAllCoin(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*CoinList, error)
	HistoryKline(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*HistoryRes, error)
	FindExchangeCoinVisible(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*ExchangeCoinRes, error)
	FindCoinById(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*Coin, error)
}

type marketClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketClient(cc grpc.ClientConnInterface) MarketClient {
	return &marketClient{cc}
}

func (c *marketClient) FindSymbolThumbTrend(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*SymbolThumbRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SymbolThumbRes)
	err := c.cc.Invoke(ctx, Market_FindSymbolThumbTrend_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) FindSymbolInfo(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*ExchangeCoin, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExchangeCoin)
	err := c.cc.Invoke(ctx, Market_FindSymbolInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) FindCoinInfo(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*Coin, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Coin)
	err := c.cc.Invoke(ctx, Market_FindCoinInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) FindAllCoin(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*CoinList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CoinList)
	err := c.cc.Invoke(ctx, Market_FindAllCoin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) HistoryKline(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*HistoryRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HistoryRes)
	err := c.cc.Invoke(ctx, Market_HistoryKline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) FindExchangeCoinVisible(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*ExchangeCoinRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExchangeCoinRes)
	err := c.cc.Invoke(ctx, Market_FindExchangeCoinVisible_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) FindCoinById(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*Coin, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Coin)
	err := c.cc.Invoke(ctx, Market_FindCoinById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketServer is the server API for Market service.
// All implementations must embed UnimplementedMarketServer
// for forward compatibility.
type MarketServer interface {
	FindSymbolThumbTrend(context.Context, *MarketReq) (*SymbolThumbRes, error)
	FindSymbolInfo(context.Context, *MarketReq) (*ExchangeCoin, error)
	FindCoinInfo(context.Context, *MarketReq) (*Coin, error)
	FindAllCoin(context.Context, *MarketReq) (*CoinList, error)
	HistoryKline(context.Context, *MarketReq) (*HistoryRes, error)
	FindExchangeCoinVisible(context.Context, *MarketReq) (*ExchangeCoinRes, error)
	FindCoinById(context.Context, *MarketReq) (*Coin, error)
	mustEmbedUnimplementedMarketServer()
}

// UnimplementedMarketServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMarketServer struct{}

func (UnimplementedMarketServer) FindSymbolThumbTrend(context.Context, *MarketReq) (*SymbolThumbRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSymbolThumbTrend not implemented")
}
func (UnimplementedMarketServer) FindSymbolInfo(context.Context, *MarketReq) (*ExchangeCoin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSymbolInfo not implemented")
}
func (UnimplementedMarketServer) FindCoinInfo(context.Context, *MarketReq) (*Coin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCoinInfo not implemented")
}
func (UnimplementedMarketServer) FindAllCoin(context.Context, *MarketReq) (*CoinList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllCoin not implemented")
}
func (UnimplementedMarketServer) HistoryKline(context.Context, *MarketReq) (*HistoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HistoryKline not implemented")
}
func (UnimplementedMarketServer) FindExchangeCoinVisible(context.Context, *MarketReq) (*ExchangeCoinRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindExchangeCoinVisible not implemented")
}
func (UnimplementedMarketServer) FindCoinById(context.Context, *MarketReq) (*Coin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCoinById not implemented")
}
func (UnimplementedMarketServer) mustEmbedUnimplementedMarketServer() {}
func (UnimplementedMarketServer) testEmbeddedByValue()                {}

// UnsafeMarketServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketServer will
// result in compilation errors.
type UnsafeMarketServer interface {
	mustEmbedUnimplementedMarketServer()
}

func RegisterMarketServer(s grpc.ServiceRegistrar, srv MarketServer) {
	// If the following call pancis, it indicates UnimplementedMarketServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Market_ServiceDesc, srv)
}

func _Market_FindSymbolThumbTrend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).FindSymbolThumbTrend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_FindSymbolThumbTrend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).FindSymbolThumbTrend(ctx, req.(*MarketReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_FindSymbolInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).FindSymbolInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_FindSymbolInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).FindSymbolInfo(ctx, req.(*MarketReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_FindCoinInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).FindCoinInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_FindCoinInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).FindCoinInfo(ctx, req.(*MarketReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_FindAllCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).FindAllCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_FindAllCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).FindAllCoin(ctx, req.(*MarketReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_HistoryKline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).HistoryKline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_HistoryKline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).HistoryKline(ctx, req.(*MarketReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_FindExchangeCoinVisible_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).FindExchangeCoinVisible(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_FindExchangeCoinVisible_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).FindExchangeCoinVisible(ctx, req.(*MarketReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_FindCoinById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).FindCoinById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_FindCoinById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).FindCoinById(ctx, req.(*MarketReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Market_ServiceDesc is the grpc.ServiceDesc for Market service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Market_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "market.Market",
	HandlerType: (*MarketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindSymbolThumbTrend",
			Handler:    _Market_FindSymbolThumbTrend_Handler,
		},
		{
			MethodName: "FindSymbolInfo",
			Handler:    _Market_FindSymbolInfo_Handler,
		},
		{
			MethodName: "FindCoinInfo",
			Handler:    _Market_FindCoinInfo_Handler,
		},
		{
			MethodName: "FindAllCoin",
			Handler:    _Market_FindAllCoin_Handler,
		},
		{
			MethodName: "HistoryKline",
			Handler:    _Market_HistoryKline_Handler,
		},
		{
			MethodName: "FindExchangeCoinVisible",
			Handler:    _Market_FindExchangeCoinVisible_Handler,
		},
		{
			MethodName: "FindCoinById",
			Handler:    _Market_FindCoinById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "market.proto",
}
