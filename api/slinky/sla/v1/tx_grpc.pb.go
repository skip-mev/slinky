// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: slinky/sla/v1/tx.proto

package slav1

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
	Msg_AddSLAs_FullMethodName    = "/slinky.sla.v1.Msg/AddSLAs"
	Msg_RemoveSLAs_FullMethodName = "/slinky.sla.v1.Msg/RemoveSLAs"
	Msg_Params_FullMethodName     = "/slinky.sla.v1.Msg/Params"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// AddSLA defines a method for adding a new SLAs to the store. Note, this will
	// overwrite any existing SLA with the same Ticker.
	AddSLAs(ctx context.Context, in *MsgAddSLAs, opts ...grpc.CallOption) (*MsgAddSLAsResponse, error)
	// RemoveSLA defines a method for removing existing SLAs from the store. Note,
	// this will not panic if the SLA does not exist.
	RemoveSLAs(ctx context.Context, in *MsgRemoveSLAs, opts ...grpc.CallOption) (*MsgRemoveSLAsResponse, error)
	// Params defines a method for updating the SLA module parameters.
	Params(ctx context.Context, in *MsgParams, opts ...grpc.CallOption) (*MsgParamsResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddSLAs(ctx context.Context, in *MsgAddSLAs, opts ...grpc.CallOption) (*MsgAddSLAsResponse, error) {
	out := new(MsgAddSLAsResponse)
	err := c.cc.Invoke(ctx, Msg_AddSLAs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveSLAs(ctx context.Context, in *MsgRemoveSLAs, opts ...grpc.CallOption) (*MsgRemoveSLAsResponse, error) {
	out := new(MsgRemoveSLAsResponse)
	err := c.cc.Invoke(ctx, Msg_RemoveSLAs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Params(ctx context.Context, in *MsgParams, opts ...grpc.CallOption) (*MsgParamsResponse, error) {
	out := new(MsgParamsResponse)
	err := c.cc.Invoke(ctx, Msg_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// AddSLA defines a method for adding a new SLAs to the store. Note, this will
	// overwrite any existing SLA with the same Ticker.
	AddSLAs(context.Context, *MsgAddSLAs) (*MsgAddSLAsResponse, error)
	// RemoveSLA defines a method for removing existing SLAs from the store. Note,
	// this will not panic if the SLA does not exist.
	RemoveSLAs(context.Context, *MsgRemoveSLAs) (*MsgRemoveSLAsResponse, error)
	// Params defines a method for updating the SLA module parameters.
	Params(context.Context, *MsgParams) (*MsgParamsResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) AddSLAs(context.Context, *MsgAddSLAs) (*MsgAddSLAsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSLAs not implemented")
}
func (UnimplementedMsgServer) RemoveSLAs(context.Context, *MsgRemoveSLAs) (*MsgRemoveSLAsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSLAs not implemented")
}
func (UnimplementedMsgServer) Params(context.Context, *MsgParams) (*MsgParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_AddSLAs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddSLAs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddSLAs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddSLAs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddSLAs(ctx, req.(*MsgAddSLAs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveSLAs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveSLAs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveSLAs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RemoveSLAs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveSLAs(ctx, req.(*MsgRemoveSLAs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Params(ctx, req.(*MsgParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "slinky.sla.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSLAs",
			Handler:    _Msg_AddSLAs_Handler,
		},
		{
			MethodName: "RemoveSLAs",
			Handler:    _Msg_RemoveSLAs_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Msg_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slinky/sla/v1/tx.proto",
}
