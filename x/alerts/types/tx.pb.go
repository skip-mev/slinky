// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slinky/alerts/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgAlert defines a message to create an alert.
type MsgAlert struct {
	// alert is the alert to be filed
	Alert Alert `protobuf:"bytes,1,opt,name=alert,proto3" json:"alert"`
}

func (m *MsgAlert) Reset()         { *m = MsgAlert{} }
func (m *MsgAlert) String() string { return proto.CompactTextString(m) }
func (*MsgAlert) ProtoMessage()    {}
func (*MsgAlert) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddd37b6330b105e2, []int{0}
}
func (m *MsgAlert) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAlert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAlert.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAlert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAlert.Merge(m, src)
}
func (m *MsgAlert) XXX_Size() int {
	return m.Size()
}
func (m *MsgAlert) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAlert.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAlert proto.InternalMessageInfo

func (m *MsgAlert) GetAlert() Alert {
	if m != nil {
		return m.Alert
	}
	return Alert{}
}

type MsgAlertResponse struct {
}

func (m *MsgAlertResponse) Reset()         { *m = MsgAlertResponse{} }
func (m *MsgAlertResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAlertResponse) ProtoMessage()    {}
func (*MsgAlertResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddd37b6330b105e2, []int{1}
}
func (m *MsgAlertResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAlertResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAlertResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAlertResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAlertResponse.Merge(m, src)
}
func (m *MsgAlertResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAlertResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAlertResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAlertResponse proto.InternalMessageInfo

// MsgConclusion defines a message carrying a Conclusion made by the SecondTier,
// which will be used to close an alert. And trigger any ramifications of the
// conclusion.
type MsgConclusion struct {
	// signer is the signer of this transaction (notice, this may not always be a
	// node from the SecondTier)
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	// conclusion is the conclusion to be filed
	Conclusion *types.Any `protobuf:"bytes,2,opt,name=conclusion,proto3" json:"conclusion,omitempty"`
}

func (m *MsgConclusion) Reset()         { *m = MsgConclusion{} }
func (m *MsgConclusion) String() string { return proto.CompactTextString(m) }
func (*MsgConclusion) ProtoMessage()    {}
func (*MsgConclusion) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddd37b6330b105e2, []int{2}
}
func (m *MsgConclusion) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConclusion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConclusion.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConclusion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConclusion.Merge(m, src)
}
func (m *MsgConclusion) XXX_Size() int {
	return m.Size()
}
func (m *MsgConclusion) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConclusion.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConclusion proto.InternalMessageInfo

func (m *MsgConclusion) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgConclusion) GetConclusion() *types.Any {
	if m != nil {
		return m.Conclusion
	}
	return nil
}

type MsgConclusionResponse struct {
}

func (m *MsgConclusionResponse) Reset()         { *m = MsgConclusionResponse{} }
func (m *MsgConclusionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgConclusionResponse) ProtoMessage()    {}
func (*MsgConclusionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddd37b6330b105e2, []int{3}
}
func (m *MsgConclusionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConclusionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConclusionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConclusionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConclusionResponse.Merge(m, src)
}
func (m *MsgConclusionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgConclusionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConclusionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConclusionResponse proto.InternalMessageInfo

// MsgUpdateParams defines the message type expected by the UpdateParams rpc. It
// contains an authority address, and the new Params for the x/alerts module.
type MsgUpdateParams struct {
	// authority is the address of the authority that is submitting the update
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// params is the new set of parameters for the x/alerts module
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddd37b6330b105e2, []int{4}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddd37b6330b105e2, []int{5}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAlert)(nil), "slinky.alerts.v1.MsgAlert")
	proto.RegisterType((*MsgAlertResponse)(nil), "slinky.alerts.v1.MsgAlertResponse")
	proto.RegisterType((*MsgConclusion)(nil), "slinky.alerts.v1.MsgConclusion")
	proto.RegisterType((*MsgConclusionResponse)(nil), "slinky.alerts.v1.MsgConclusionResponse")
	proto.RegisterType((*MsgUpdateParams)(nil), "slinky.alerts.v1.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "slinky.alerts.v1.MsgUpdateParamsResponse")
}

func init() { proto.RegisterFile("slinky/alerts/v1/tx.proto", fileDescriptor_ddd37b6330b105e2) }

var fileDescriptor_ddd37b6330b105e2 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0x8e, 0x0b, 0x89, 0xe8, 0x01, 0xa2, 0x58, 0x41, 0x49, 0x8c, 0xea, 0x50, 0x2f, 0x40, 0xa4,
	0xf8, 0x48, 0x2a, 0x65, 0xc8, 0x96, 0x30, 0x74, 0x8a, 0x84, 0x82, 0xca, 0x80, 0x90, 0x90, 0xe3,
	0x1c, 0x57, 0xab, 0xf1, 0x9d, 0xe5, 0xe7, 0x44, 0xcd, 0x86, 0x18, 0x61, 0xe1, 0x27, 0xf0, 0x13,
	0x32, 0x74, 0xe1, 0x1f, 0x54, 0xb0, 0x54, 0x4c, 0x4c, 0x08, 0x25, 0x43, 0xf8, 0x19, 0xc8, 0xe7,
	0xbb, 0x38, 0xad, 0x53, 0x65, 0xb1, 0xfc, 0xee, 0xfb, 0xde, 0xf7, 0xbd, 0xef, 0xdd, 0xa1, 0x0a,
	0x8c, 0x3c, 0x76, 0x3a, 0xc5, 0xce, 0x88, 0x84, 0x11, 0xe0, 0x49, 0x03, 0x47, 0x67, 0x76, 0x10,
	0xf2, 0x88, 0xeb, 0x7b, 0x09, 0x64, 0x27, 0x90, 0x3d, 0x69, 0x18, 0xfb, 0x19, 0xb2, 0xc4, 0x44,
	0x83, 0x51, 0x71, 0x39, 0xf8, 0x1c, 0xde, 0x8b, 0x0a, 0x27, 0x85, 0x84, 0x4a, 0x49, 0x85, 0x7d,
	0xa0, 0x71, 0x9b, 0x0f, 0x54, 0x02, 0x0f, 0x1d, 0xdf, 0x63, 0x1c, 0x8b, 0xaf, 0x3c, 0x2a, 0x52,
	0x4e, 0x79, 0xa2, 0x11, 0xff, 0x29, 0x71, 0xca, 0x39, 0x1d, 0x11, 0x2c, 0xaa, 0xc1, 0xf8, 0x03,
	0x76, 0xd8, 0x54, 0x42, 0x66, 0x66, 0x2c, 0x4a, 0x18, 0x01, 0x4f, 0x9a, 0x5b, 0x03, 0x74, 0xa7,
	0x07, 0xb4, 0x13, 0xa3, 0xfa, 0x21, 0xca, 0x0b, 0x5a, 0x59, 0x7b, 0xa2, 0x3d, 0xbb, 0xdb, 0x2c,
	0xd9, 0xd7, 0x43, 0xda, 0x82, 0xd7, 0xbd, 0x7d, 0xf1, 0xa7, 0x9a, 0xeb, 0x27, 0xdc, 0xf6, 0xc1,
	0xbf, 0x6f, 0xd5, 0xdc, 0xe7, 0xe5, 0xac, 0x56, 0x96, 0x4e, 0x67, 0xca, 0x4b, 0xe9, 0x5a, 0x3a,
	0xda, 0x53, 0xff, 0x7d, 0x02, 0x01, 0x67, 0x40, 0xac, 0x9f, 0x1a, 0xba, 0xdf, 0x03, 0xfa, 0x92,
	0x33, 0x77, 0x34, 0x06, 0x8f, 0x33, 0xfd, 0x05, 0x2a, 0x80, 0x47, 0x19, 0x09, 0x85, 0xfd, 0x6e,
	0xb7, 0xfc, 0xeb, 0xbc, 0x5e, 0x94, 0x8b, 0xea, 0x0c, 0x87, 0x21, 0x01, 0x78, 0x1d, 0x85, 0x1e,
	0xa3, 0x7d, 0xc9, 0xd3, 0x8f, 0x11, 0x72, 0x57, 0xfd, 0xe5, 0x1d, 0x31, 0x74, 0xd1, 0x4e, 0x76,
	0x61, 0xab, 0x5d, 0xd8, 0x1d, 0x36, 0xed, 0x56, 0x7f, 0x9c, 0xd7, 0x1f, 0x67, 0xd2, 0xa4, 0xe6,
	0xfd, 0x35, 0xa1, 0x76, 0x23, 0x4e, 0xf4, 0x69, 0x39, 0xab, 0x49, 0x9f, 0x38, 0xdc, 0xfe, 0x86,
	0x70, 0x69, 0xbb, 0x55, 0x42, 0x8f, 0xae, 0x1c, 0xac, 0x62, 0x7e, 0xd7, 0xd0, 0x83, 0x1e, 0xd0,
	0xe3, 0x60, 0xe8, 0x44, 0xe4, 0x95, 0x13, 0x3a, 0x3e, 0xe8, 0x2d, 0xb4, 0xeb, 0x8c, 0xa3, 0x13,
	0x1e, 0x7a, 0xd1, 0x74, 0x6b, 0xd6, 0x94, 0xaa, 0xb7, 0x50, 0x21, 0x10, 0x0a, 0x32, 0x6a, 0x39,
	0x7b, 0x3f, 0x89, 0x83, 0xbc, 0x20, 0xc9, 0x6e, 0xb7, 0x54, 0x9e, 0x54, 0x2b, 0x8e, 0x54, 0xdd,
	0x10, 0x69, 0x7d, 0x4e, 0xab, 0x82, 0x4a, 0xd7, 0x8e, 0x54, 0xac, 0xe6, 0x97, 0x1d, 0x74, 0xab,
	0x07, 0x54, 0x3f, 0x42, 0xf9, 0xe4, 0xe9, 0x18, 0xd9, 0x59, 0xd4, 0x95, 0x1b, 0xd6, 0xcd, 0x98,
	0x12, 0xd4, 0xdf, 0x20, 0xb4, 0xf6, 0x14, 0xaa, 0x1b, 0x3b, 0x52, 0x82, 0xf1, 0x74, 0x0b, 0x61,
	0xa5, 0xfb, 0x0e, 0xdd, 0xbb, 0xb2, 0xfb, 0x83, 0x8d, 0x8d, 0xeb, 0x14, 0xe3, 0xf9, 0x56, 0x8a,
	0x52, 0x37, 0xf2, 0x1f, 0x97, 0xb3, 0x9a, 0xd6, 0x3d, 0xba, 0x98, 0x9b, 0xda, 0xe5, 0xdc, 0xd4,
	0xfe, 0xce, 0x4d, 0xed, 0xeb, 0xc2, 0xcc, 0x5d, 0x2e, 0xcc, 0xdc, 0xef, 0x85, 0x99, 0x7b, 0x5b,
	0xa7, 0x5e, 0x74, 0x32, 0x1e, 0xd8, 0x2e, 0xf7, 0x31, 0x9c, 0x7a, 0x41, 0xdd, 0x27, 0x13, 0xec,
	0x72, 0xc6, 0x88, 0x1b, 0xe1, 0x49, 0x33, 0xdd, 0x7d, 0x34, 0x0d, 0x08, 0x0c, 0x0a, 0xe2, 0xd1,
	0x1e, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x29, 0xd2, 0xcc, 0xec, 0x79, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
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
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Alert(ctx context.Context, in *MsgAlert, opts ...grpc.CallOption) (*MsgAlertResponse, error) {
	out := new(MsgAlertResponse)
	err := c.cc.Invoke(ctx, "/slinky.alerts.v1.Msg/Alert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Conclusion(ctx context.Context, in *MsgConclusion, opts ...grpc.CallOption) (*MsgConclusionResponse, error) {
	out := new(MsgConclusionResponse)
	err := c.cc.Invoke(ctx, "/slinky.alerts.v1.Msg/Conclusion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/slinky.alerts.v1.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
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
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Alert(ctx context.Context, req *MsgAlert) (*MsgAlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Alert not implemented")
}
func (*UnimplementedMsgServer) Conclusion(ctx context.Context, req *MsgConclusion) (*MsgConclusionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Conclusion not implemented")
}
func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
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
		FullMethod: "/slinky.alerts.v1.Msg/Alert",
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
		FullMethod: "/slinky.alerts.v1.Msg/Conclusion",
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
		FullMethod: "/slinky.alerts.v1.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
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

func (m *MsgAlert) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAlert) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAlert) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Alert.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgAlertResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAlertResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAlertResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgConclusion) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConclusion) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConclusion) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Conclusion != nil {
		{
			size, err := m.Conclusion.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgConclusionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConclusionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConclusionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAlert) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Alert.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgAlertResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgConclusion) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Conclusion != nil {
		l = m.Conclusion.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgConclusionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAlert) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAlert: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAlert: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alert", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Alert.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgAlertResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAlertResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAlertResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgConclusion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgConclusion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConclusion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Conclusion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Conclusion == nil {
				m.Conclusion = &types.Any{}
			}
			if err := m.Conclusion.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgConclusionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgConclusionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConclusionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
