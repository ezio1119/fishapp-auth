// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package auth_grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TokenPair struct {
	IdToken              string   `protobuf:"bytes,1,opt,name=idToken,proto3" json:"idToken,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenPair) Reset()         { *m = TokenPair{} }
func (m *TokenPair) String() string { return proto.CompactTextString(m) }
func (*TokenPair) ProtoMessage()    {}
func (*TokenPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *TokenPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenPair.Unmarshal(m, b)
}
func (m *TokenPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenPair.Marshal(b, m, deterministic)
}
func (m *TokenPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenPair.Merge(m, src)
}
func (m *TokenPair) XXX_Size() int {
	return xxx_messageInfo_TokenPair.Size(m)
}
func (m *TokenPair) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenPair.DiscardUnknown(m)
}

var xxx_messageInfo_TokenPair proto.InternalMessageInfo

func (m *TokenPair) GetIdToken() string {
	if m != nil {
		return m.IdToken
	}
	return ""
}

func (m *TokenPair) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func init() {
	proto.RegisterType((*TokenPair)(nil), "auth_grpc.TokenPair")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0xb1, 0xe3, 0xd3, 0x8b, 0x0a, 0x92, 0xa5,
	0xa4, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0x12, 0x49, 0xa5, 0x69, 0xfa, 0xa9, 0xb9,
	0x05, 0x25, 0x95, 0x10, 0x75, 0x52, 0x72, 0xe8, 0x92, 0xe5, 0x45, 0x89, 0x05, 0x05, 0xa9, 0x45,
	0xc5, 0x10, 0x79, 0x25, 0x4f, 0x2e, 0xce, 0x90, 0xfc, 0xec, 0xd4, 0xbc, 0x80, 0xc4, 0xcc, 0x22,
	0x21, 0x09, 0x2e, 0xf6, 0xcc, 0x14, 0x30, 0x57, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6,
	0x15, 0x52, 0xe2, 0xe2, 0x29, 0x4a, 0x4d, 0x2b, 0x4a, 0x2d, 0xce, 0x80, 0x48, 0x33, 0x81, 0xa5,
	0x51, 0xc4, 0x8c, 0xae, 0x33, 0x72, 0x71, 0x3b, 0x96, 0x96, 0x64, 0x04, 0xa7, 0x16, 0x95, 0x65,
	0x26, 0xa7, 0x0a, 0xd9, 0x71, 0xf1, 0x05, 0x41, 0xe4, 0x3d, 0x5d, 0x20, 0xa6, 0x88, 0xe9, 0x41,
	0x5c, 0xa3, 0x07, 0x73, 0x8d, 0x9e, 0x2b, 0xc8, 0xa9, 0x52, 0x22, 0x7a, 0x70, 0xdf, 0xe8, 0x21,
	0x5c, 0xe3, 0xc2, 0xc5, 0xe7, 0x9c, 0x91, 0x9a, 0x9c, 0xed, 0x94, 0x93, 0x98, 0x9c, 0xed, 0x93,
	0x59, 0x5c, 0x82, 0x53, 0xbf, 0x14, 0x86, 0xb8, 0x53, 0x7e, 0x7e, 0x4e, 0x58, 0x62, 0x4e, 0x69,
	0xaa, 0x90, 0x13, 0x17, 0x8f, 0x63, 0x4a, 0x0a, 0x45, 0x66, 0x24, 0xb1, 0x81, 0xc5, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x03, 0xec, 0x56, 0xd9, 0x81, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	RefreshIDToken(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TokenPair, error)
	CheckBlackList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
	AddBlackList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) RefreshIDToken(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TokenPair, error) {
	out := new(TokenPair)
	err := c.cc.Invoke(ctx, "/auth_grpc.AuthService/RefreshIDToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CheckBlackList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/auth_grpc.AuthService/CheckBlackList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AddBlackList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/auth_grpc.AuthService/AddBlackList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	RefreshIDToken(context.Context, *empty.Empty) (*TokenPair, error)
	CheckBlackList(context.Context, *empty.Empty) (*wrappers.BoolValue, error)
	AddBlackList(context.Context, *empty.Empty) (*wrappers.BoolValue, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) RefreshIDToken(ctx context.Context, req *empty.Empty) (*TokenPair, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshIDToken not implemented")
}
func (*UnimplementedAuthServiceServer) CheckBlackList(ctx context.Context, req *empty.Empty) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckBlackList not implemented")
}
func (*UnimplementedAuthServiceServer) AddBlackList(ctx context.Context, req *empty.Empty) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlackList not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_RefreshIDToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RefreshIDToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_grpc.AuthService/RefreshIDToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RefreshIDToken(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CheckBlackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CheckBlackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_grpc.AuthService/CheckBlackList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CheckBlackList(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AddBlackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AddBlackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_grpc.AuthService/AddBlackList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AddBlackList(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth_grpc.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RefreshIDToken",
			Handler:    _AuthService_RefreshIDToken_Handler,
		},
		{
			MethodName: "CheckBlackList",
			Handler:    _AuthService_CheckBlackList_Handler,
		},
		{
			MethodName: "AddBlackList",
			Handler:    _AuthService_AddBlackList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
