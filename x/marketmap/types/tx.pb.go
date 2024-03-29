// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slinky/marketmap/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// CreateMarket contains all information needed to create a new market.
type CreateMarket struct {
	// Ticker is the on-chain representation of the ticker. This is the target
	// ticker that the prices of the set of tickers will be converted to.
	Ticker Ticker `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker"`
	// Providers maps provider names to their off-chain
	// representations for the given ticker of the message.
	Providers Providers `protobuf:"bytes,2,opt,name=providers,proto3" json:"providers"`
	// Paths is the list of convertable markets that will be used to convert the
	// prices of a set of tickers to a common ticker.
	Paths Paths `protobuf:"bytes,3,opt,name=paths,proto3" json:"paths"`
}

func (m *CreateMarket) Reset()         { *m = CreateMarket{} }
func (m *CreateMarket) String() string { return proto.CompactTextString(m) }
func (*CreateMarket) ProtoMessage()    {}
func (*CreateMarket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9adadfc18297083, []int{0}
}
func (m *CreateMarket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateMarket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateMarket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateMarket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMarket.Merge(m, src)
}
func (m *CreateMarket) XXX_Size() int {
	return m.Size()
}
func (m *CreateMarket) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMarket.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMarket proto.InternalMessageInfo

func (m *CreateMarket) GetTicker() Ticker {
	if m != nil {
		return m.Ticker
	}
	return Ticker{}
}

func (m *CreateMarket) GetProviders() Providers {
	if m != nil {
		return m.Providers
	}
	return Providers{}
}

func (m *CreateMarket) GetPaths() Paths {
	if m != nil {
		return m.Paths
	}
	return Paths{}
}

// MsgUpdateMarketMap defines a message carrying a payload for updating the
// x/marketmap module.
type MsgUpdateMarketMap struct {
	// Signer is the signer of this transaction.
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	// CreateMarkets is the list of all markets to be created for the given
	// transaction.
	CreateMarkets []CreateMarket `protobuf:"bytes,2,rep,name=create_markets,json=createMarkets,proto3" json:"create_markets"`
}

func (m *MsgUpdateMarketMap) Reset()         { *m = MsgUpdateMarketMap{} }
func (m *MsgUpdateMarketMap) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateMarketMap) ProtoMessage()    {}
func (*MsgUpdateMarketMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9adadfc18297083, []int{1}
}
func (m *MsgUpdateMarketMap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateMarketMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateMarketMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateMarketMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateMarketMap.Merge(m, src)
}
func (m *MsgUpdateMarketMap) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateMarketMap) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateMarketMap.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateMarketMap proto.InternalMessageInfo

func (m *MsgUpdateMarketMap) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgUpdateMarketMap) GetCreateMarkets() []CreateMarket {
	if m != nil {
		return m.CreateMarkets
	}
	return nil
}

// MsgUpdateMarketMapResponse is the response message for MsgUpdateMarketMap.
type MsgUpdateMarketMapResponse struct {
}

func (m *MsgUpdateMarketMapResponse) Reset()         { *m = MsgUpdateMarketMapResponse{} }
func (m *MsgUpdateMarketMapResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateMarketMapResponse) ProtoMessage()    {}
func (*MsgUpdateMarketMapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9adadfc18297083, []int{2}
}
func (m *MsgUpdateMarketMapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateMarketMapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateMarketMapResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateMarketMapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateMarketMapResponse.Merge(m, src)
}
func (m *MsgUpdateMarketMapResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateMarketMapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateMarketMapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateMarketMapResponse proto.InternalMessageInfo

// MsgParams defines the Msg/Params request type. It contains the
// new parameters for the x/marketmap module.
type MsgParams struct {
	// Params defines the new parameters for the x/marketmap module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// Authority defines the authority that is updating the x/marketmap module
	// parameters.
	Authority string `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
}

func (m *MsgParams) Reset()         { *m = MsgParams{} }
func (m *MsgParams) String() string { return proto.CompactTextString(m) }
func (*MsgParams) ProtoMessage()    {}
func (*MsgParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9adadfc18297083, []int{3}
}
func (m *MsgParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgParams.Merge(m, src)
}
func (m *MsgParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgParams proto.InternalMessageInfo

func (m *MsgParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *MsgParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

// MsgParamsResponse defines the Msg/Params response type.
type MsgParamsResponse struct {
}

func (m *MsgParamsResponse) Reset()         { *m = MsgParamsResponse{} }
func (m *MsgParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgParamsResponse) ProtoMessage()    {}
func (*MsgParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9adadfc18297083, []int{4}
}
func (m *MsgParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgParamsResponse.Merge(m, src)
}
func (m *MsgParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateMarket)(nil), "slinky.marketmap.v1.CreateMarket")
	proto.RegisterType((*MsgUpdateMarketMap)(nil), "slinky.marketmap.v1.MsgUpdateMarketMap")
	proto.RegisterType((*MsgUpdateMarketMapResponse)(nil), "slinky.marketmap.v1.MsgUpdateMarketMapResponse")
	proto.RegisterType((*MsgParams)(nil), "slinky.marketmap.v1.MsgParams")
	proto.RegisterType((*MsgParamsResponse)(nil), "slinky.marketmap.v1.MsgParamsResponse")
}

func init() { proto.RegisterFile("slinky/marketmap/v1/tx.proto", fileDescriptor_e9adadfc18297083) }

var fileDescriptor_e9adadfc18297083 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6a, 0x13, 0x41,
	0x18, 0xc7, 0x77, 0x8c, 0x0d, 0x64, 0xaa, 0x95, 0x4e, 0x0b, 0xc6, 0xb5, 0x6c, 0xeb, 0x82, 0x5a,
	0x82, 0xdd, 0xb5, 0x11, 0x0a, 0xed, 0xcd, 0x08, 0xde, 0x56, 0x42, 0xd4, 0x8b, 0x97, 0xb2, 0xdd,
	0x0c, 0x93, 0x65, 0xdd, 0x9d, 0x61, 0x66, 0x1a, 0x9a, 0x9b, 0x78, 0xf4, 0x20, 0x3e, 0x82, 0x8f,
	0x90, 0x83, 0xaf, 0x20, 0x14, 0xbc, 0x44, 0x4f, 0x9e, 0x44, 0x92, 0x43, 0x7c, 0x0c, 0xd9, 0x99,
	0xd9, 0x24, 0x98, 0x35, 0xf4, 0x12, 0x76, 0xbe, 0xef, 0xf7, 0x9f, 0xf9, 0x7f, 0xdf, 0x9f, 0xc0,
	0x1d, 0xf1, 0x36, 0xce, 0x92, 0x81, 0x9f, 0x86, 0x3c, 0xc1, 0x32, 0x0d, 0x99, 0xdf, 0x3f, 0xf4,
	0xe5, 0x85, 0xc7, 0x38, 0x95, 0x14, 0x6d, 0xe9, 0xae, 0x37, 0xeb, 0x7a, 0xfd, 0x43, 0xfb, 0x76,
	0x44, 0x45, 0x4a, 0x85, 0x9f, 0x0a, 0x92, 0xc3, 0xa9, 0x20, 0x9a, 0xb6, 0xb7, 0x09, 0x25, 0x54,
	0x7d, 0xfa, 0xf9, 0x97, 0xa9, 0xde, 0xd1, 0xf8, 0xa9, 0x6e, 0xe8, 0x83, 0x69, 0x6d, 0x86, 0x69,
	0x9c, 0x51, 0x5f, 0xfd, 0x9a, 0xd2, 0x5e, 0x99, 0x1f, 0x7d, 0x58, 0x45, 0xb0, 0x90, 0x87, 0xa9,
	0xb9, 0xd6, 0xfd, 0x0a, 0xe0, 0x8d, 0x67, 0x1c, 0x87, 0x12, 0x07, 0x8a, 0x41, 0xc7, 0xb0, 0x2a,
	0xe3, 0x28, 0xc1, 0xbc, 0x0e, 0xf6, 0xc0, 0xfe, 0x7a, 0xf3, 0xae, 0x57, 0x32, 0x97, 0xf7, 0x4a,
	0x21, 0xad, 0xeb, 0x97, 0xbf, 0x76, 0xad, 0x8e, 0x11, 0xa0, 0x16, 0xac, 0x31, 0x4e, 0xfb, 0x71,
	0x17, 0x73, 0x51, 0xbf, 0xa6, 0xd4, 0x4e, 0xa9, 0xba, 0x5d, 0x50, 0xe6, 0x82, 0xb9, 0x0c, 0x1d,
	0xc1, 0x35, 0x16, 0xca, 0x9e, 0xa8, 0x57, 0x94, 0xde, 0x2e, 0xd7, 0xe7, 0x84, 0xd1, 0x6a, 0xdc,
	0xfd, 0x0e, 0x20, 0x0a, 0x04, 0x79, 0xcd, 0xba, 0xb3, 0x51, 0x82, 0x90, 0xa1, 0xc7, 0xb0, 0x2a,
	0x62, 0x92, 0x99, 0x69, 0x6a, 0xad, 0xfa, 0x8f, 0x2f, 0x07, 0xdb, 0x66, 0xaf, 0x4f, 0xbb, 0x5d,
	0x8e, 0x85, 0x78, 0x29, 0x79, 0x9c, 0x91, 0x8e, 0xe1, 0xd0, 0x0b, 0xb8, 0x11, 0xa9, 0x7d, 0x9c,
	0xea, 0x27, 0xf3, 0x49, 0x2a, 0xfb, 0xeb, 0xcd, 0x7b, 0xa5, 0x4e, 0x16, 0x57, 0x67, 0x0c, 0xdd,
	0x8c, 0x16, 0x6a, 0xe2, 0xe4, 0xf8, 0xcf, 0xe7, 0x5d, 0xeb, 0xfd, 0x74, 0xd8, 0x30, 0x0f, 0x7c,
	0x98, 0x0e, 0x1b, 0xf7, 0x4d, 0x2c, 0x17, 0x0b, 0xc1, 0x2c, 0x9b, 0x77, 0x77, 0xa0, 0xbd, 0x5c,
	0xed, 0x60, 0xc1, 0x68, 0x26, 0xb0, 0xfb, 0x11, 0xc0, 0x5a, 0x20, 0x48, 0x5b, 0xa5, 0x99, 0xc7,
	0xa6, 0x73, 0x5d, 0x19, 0x9b, 0x86, 0x8b, 0xd8, 0xb4, 0x00, 0x1d, 0xc1, 0x5a, 0x78, 0x2e, 0x7b,
	0x94, 0xc7, 0x72, 0xa0, 0x62, 0x5b, 0xb5, 0xa6, 0x39, 0x7a, 0xb2, 0x91, 0x4f, 0x35, 0x3f, 0xbb,
	0x5b, 0x70, 0x73, 0xe6, 0xa7, 0x70, 0xd9, 0xfc, 0x06, 0x60, 0x25, 0x10, 0x04, 0x25, 0xf0, 0xd6,
	0xbf, 0xd9, 0x3c, 0x2c, 0xb5, 0xb8, 0x3c, 0xb1, 0xed, 0x5f, 0x11, 0x2c, 0x1e, 0x45, 0x6d, 0x58,
	0x35, 0x6b, 0x71, 0xfe, 0x27, 0xd5, 0x7d, 0xfb, 0xc1, 0xea, 0x7e, 0x71, 0xa3, 0xbd, 0xf6, 0x6e,
	0x3a, 0x6c, 0x80, 0xd6, 0xf3, 0xcb, 0xb1, 0x03, 0x46, 0x63, 0x07, 0xfc, 0x1e, 0x3b, 0xe0, 0xd3,
	0xc4, 0xb1, 0x46, 0x13, 0xc7, 0xfa, 0x39, 0x71, 0xac, 0x37, 0x8f, 0x48, 0x2c, 0x7b, 0xe7, 0x67,
	0x5e, 0x44, 0x53, 0x5f, 0x24, 0x31, 0x3b, 0x48, 0x71, 0xdf, 0x2f, 0x89, 0x59, 0x0e, 0x18, 0x16,
	0x67, 0x55, 0xf5, 0xe7, 0x7b, 0xf2, 0x37, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x70, 0x1d, 0x7a, 0x52,
	0x04, 0x00, 0x00,
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
	// UpdateMarketMap creates markets from the given message.
	UpdateMarketMap(ctx context.Context, in *MsgUpdateMarketMap, opts ...grpc.CallOption) (*MsgUpdateMarketMapResponse, error)
	// Params defines a method for updating the x/marketmap module parameters.
	Params(ctx context.Context, in *MsgParams, opts ...grpc.CallOption) (*MsgParamsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateMarketMap(ctx context.Context, in *MsgUpdateMarketMap, opts ...grpc.CallOption) (*MsgUpdateMarketMapResponse, error) {
	out := new(MsgUpdateMarketMapResponse)
	err := c.cc.Invoke(ctx, "/slinky.marketmap.v1.Msg/UpdateMarketMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Params(ctx context.Context, in *MsgParams, opts ...grpc.CallOption) (*MsgParamsResponse, error) {
	out := new(MsgParamsResponse)
	err := c.cc.Invoke(ctx, "/slinky.marketmap.v1.Msg/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateMarketMap creates markets from the given message.
	UpdateMarketMap(context.Context, *MsgUpdateMarketMap) (*MsgUpdateMarketMapResponse, error)
	// Params defines a method for updating the x/marketmap module parameters.
	Params(context.Context, *MsgParams) (*MsgParamsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateMarketMap(ctx context.Context, req *MsgUpdateMarketMap) (*MsgUpdateMarketMapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMarketMap not implemented")
}
func (*UnimplementedMsgServer) Params(ctx context.Context, req *MsgParams) (*MsgParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateMarketMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateMarketMap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateMarketMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slinky.marketmap.v1.Msg/UpdateMarketMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateMarketMap(ctx, req.(*MsgUpdateMarketMap))
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
		FullMethod: "/slinky.marketmap.v1.Msg/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Params(ctx, req.(*MsgParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "slinky.marketmap.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateMarketMap",
			Handler:    _Msg_UpdateMarketMap_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Msg_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slinky/marketmap/v1/tx.proto",
}

func (m *CreateMarket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateMarket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateMarket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Paths.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Providers.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Ticker.MarshalToSizedBuffer(dAtA[:i])
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

func (m *MsgUpdateMarketMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateMarketMap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateMarketMap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CreateMarkets) > 0 {
		for iNdEx := len(m.CreateMarkets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CreateMarkets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
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

func (m *MsgUpdateMarketMapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateMarketMapResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateMarketMapResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
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

func (m *MsgParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *CreateMarket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Ticker.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.Providers.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.Paths.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateMarketMap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.CreateMarkets) > 0 {
		for _, e := range m.CreateMarkets {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgUpdateMarketMapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgParamsResponse) Size() (n int) {
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
func (m *CreateMarket) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CreateMarket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateMarket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticker", wireType)
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
			if err := m.Ticker.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Providers", wireType)
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
			if err := m.Providers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paths", wireType)
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
			if err := m.Paths.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgUpdateMarketMap) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateMarketMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateMarketMap: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field CreateMarkets", wireType)
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
			m.CreateMarkets = append(m.CreateMarkets, CreateMarket{})
			if err := m.CreateMarkets[len(m.CreateMarkets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgUpdateMarketMapResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateMarketMapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateMarketMapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 2:
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
func (m *MsgParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
