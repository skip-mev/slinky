// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: slinky/alerts/v1/tx.proto

package alertsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Msg_Alert_FullMethodName        = "/slinky.alerts.v1.Msg/Alert"
	Msg_Conclusion_FullMethodName   = "/slinky.alerts.v1.Msg/Conclusion"
	Msg_UpdateParams_FullMethodName = "/slinky.alerts.v1.Msg/UpdateParams"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Msg is the message service for the x/alerts module.
type MsgClient interface {
	// Alert creates a new alert. On alert creation (if valid), the alert will be
	// saved to state, and its bond will be escrowed until a corresponding
	// Conclusion is filed to close the alert.
	Alert(ctx context.Context, in *MsgAlert, opts ...grpc.CallOption) (*MsgAlertResponse, error)
	// Conclusion closes an alert. On alert conclusion (if valid), the alert will
	// be marked as Concluded, the bond for the alert will either be burned or
	// returned, and a set of incentives will be issued to the validators deemed
	// malicious by the conclusion.
	Conclusion(ctx context.Context, in *MsgConclusion, opts ...grpc.CallOption) (*MsgConclusionResponse, error)
	// UpdateParams updates the parameters of the alerts module. Specifically, the
	// only address that is capable of submitting this Msg is the
	// module-authority, in general, the x/gov module-account. The process for
	// executing this message will be via governance proposal
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Alert(ctx context.Context, in *MsgAlert, opts ...grpc.CallOption) (*MsgAlertResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgAlertResponse)
	err := c.cc.Invoke(ctx, Msg_Alert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Conclusion(ctx context.Context, in *MsgConclusion, opts ...grpc.CallOption) (*MsgConclusionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgConclusionResponse)
	err := c.cc.Invoke(ctx, Msg_Conclusion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
//
// Msg is the message service for the x/alerts module.
type MsgServer interface {
	// Alert creates a new alert. On alert creation (if valid), the alert will be
	// saved to state, and its bond will be escrowed until a corresponding
	// Conclusion is filed to close the alert.
	Alert(context.Context, *MsgAlert) (*MsgAlertResponse, error)
	// Conclusion closes an alert. On alert conclusion (if valid), the alert will
	// be marked as Concluded, the bond for the alert will either be burned or
	// returned, and a set of incentives will be issued to the validators deemed
	// malicious by the conclusion.
	Conclusion(context.Context, *MsgConclusion) (*MsgConclusionResponse, error)
	// UpdateParams updates the parameters of the alerts module. Specifically, the
	// only address that is capable of submitting this Msg is the
	// module-authority, in general, the x/gov module-account. The process for
	// executing this message will be via governance proposal
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) Alert(context.Context, *MsgAlert) (*MsgAlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Alert not implemented")
}
func (UnimplementedMsgServer) Conclusion(context.Context, *MsgConclusion) (*MsgConclusionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Conclusion not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
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

func _Msg_Alert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAlert)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Alert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Alert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Alert(ctx, req.(*MsgAlert))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Conclusion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConclusion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Conclusion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Conclusion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Conclusion(ctx, req.(*MsgConclusion))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "slinky.alerts.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Alert",
			Handler:    _Msg_Alert_Handler,
		},
		{
			MethodName: "Conclusion",
			Handler:    _Msg_Conclusion_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slinky/alerts/v1/tx.proto",
}
