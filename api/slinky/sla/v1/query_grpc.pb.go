// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: slinky/sla/v1/query.proto

package slav1

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
	Query_GetAllSLAs_FullMethodName    = "/slinky.sla.v1.Query/GetAllSLAs"
	Query_GetPriceFeeds_FullMethodName = "/slinky.sla.v1.Query/GetPriceFeeds"
	Query_Params_FullMethodName        = "/slinky.sla.v1.Query/Params"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Query is the query service for the x/sla module.
type QueryClient interface {
	// GetAllSLAs returns all SLAs that the module is currently enforcing.
	GetAllSLAs(ctx context.Context, in *GetAllSLAsRequest, opts ...grpc.CallOption) (*GetAllSLAsResponse, error)
	// GetPriceFeeds returns all price feeds that the module is currently
	// tracking. This request type inputs the SLA ID to query price feeds for.
	GetPriceFeeds(ctx context.Context, in *GetPriceFeedsRequest, opts ...grpc.CallOption) (*GetPriceFeedsResponse, error)
	// Params returns the current SLA module parameters.
	Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetAllSLAs(ctx context.Context, in *GetAllSLAsRequest, opts ...grpc.CallOption) (*GetAllSLAsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllSLAsResponse)
	err := c.cc.Invoke(ctx, Query_GetAllSLAs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetPriceFeeds(ctx context.Context, in *GetPriceFeedsRequest, opts ...grpc.CallOption) (*GetPriceFeedsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPriceFeedsResponse)
	err := c.cc.Invoke(ctx, Query_GetPriceFeeds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility.
//
// Query is the query service for the x/sla module.
type QueryServer interface {
	// GetAllSLAs returns all SLAs that the module is currently enforcing.
	GetAllSLAs(context.Context, *GetAllSLAsRequest) (*GetAllSLAsResponse, error)
	// GetPriceFeeds returns all price feeds that the module is currently
	// tracking. This request type inputs the SLA ID to query price feeds for.
	GetPriceFeeds(context.Context, *GetPriceFeedsRequest) (*GetPriceFeedsResponse, error)
	// Params returns the current SLA module parameters.
	Params(context.Context, *ParamsRequest) (*ParamsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQueryServer struct{}

func (UnimplementedQueryServer) GetAllSLAs(context.Context, *GetAllSLAsRequest) (*GetAllSLAsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSLAs not implemented")
}
func (UnimplementedQueryServer) GetPriceFeeds(context.Context, *GetPriceFeedsRequest) (*GetPriceFeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPriceFeeds not implemented")
}
func (UnimplementedQueryServer) Params(context.Context, *ParamsRequest) (*ParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}
func (UnimplementedQueryServer) testEmbeddedByValue()               {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	// If the following call pancis, it indicates UnimplementedQueryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_GetAllSLAs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllSLAsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetAllSLAs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetAllSLAs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetAllSLAs(ctx, req.(*GetAllSLAsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetPriceFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPriceFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetPriceFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetPriceFeeds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetPriceFeeds(ctx, req.(*GetPriceFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*ParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "slinky.sla.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllSLAs",
			Handler:    _Query_GetAllSLAs_Handler,
		},
		{
			MethodName: "GetPriceFeeds",
			Handler:    _Query_GetPriceFeeds_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slinky/sla/v1/query.proto",
}
