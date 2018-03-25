// Code generated by protoc-gen-go. DO NOT EDIT.
// source: risto_service.proto

/*
Package RistoService is a generated protocol buffer package.

It is generated from these files:
	risto_service.proto
	token.proto
	user.proto

It has these top-level messages:
	Token
	CreateTokenRequest
	CreateTokenResponse
	VerifyTokenRequest
	VerifyTokenResponse
	FetchUserByTokenRequest
	FetchUserByTokenResponse
	User
	CreateUserRequest
	CreateUserResponse
	HasAdminAccountRequest
	HasAdminAccountResponse
	CreateAdminAccountRequest
	CreateAdminAccountResponse
*/
package RistoService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Risto service

type RistoClient interface {
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error)
	VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	FetchUserByToken(ctx context.Context, in *FetchUserByTokenRequest, opts ...grpc.CallOption) (*FetchUserByTokenResponse, error)
	HasAdminAccount(ctx context.Context, in *HasAdminAccountRequest, opts ...grpc.CallOption) (*HasAdminAccountResponse, error)
	CreateAdminAccount(ctx context.Context, in *CreateAdminAccountRequest, opts ...grpc.CallOption) (*CreateAdminAccountResponse, error)
}

type ristoClient struct {
	cc *grpc.ClientConn
}

func NewRistoClient(cc *grpc.ClientConn) RistoClient {
	return &ristoClient{cc}
}

func (c *ristoClient) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error) {
	out := new(CreateTokenResponse)
	err := grpc.Invoke(ctx, "/RistoService.Risto/createToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ristoClient) VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error) {
	out := new(VerifyTokenResponse)
	err := grpc.Invoke(ctx, "/RistoService.Risto/verifyToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ristoClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := grpc.Invoke(ctx, "/RistoService.Risto/createUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ristoClient) FetchUserByToken(ctx context.Context, in *FetchUserByTokenRequest, opts ...grpc.CallOption) (*FetchUserByTokenResponse, error) {
	out := new(FetchUserByTokenResponse)
	err := grpc.Invoke(ctx, "/RistoService.Risto/fetchUserByToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ristoClient) HasAdminAccount(ctx context.Context, in *HasAdminAccountRequest, opts ...grpc.CallOption) (*HasAdminAccountResponse, error) {
	out := new(HasAdminAccountResponse)
	err := grpc.Invoke(ctx, "/RistoService.Risto/hasAdminAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ristoClient) CreateAdminAccount(ctx context.Context, in *CreateAdminAccountRequest, opts ...grpc.CallOption) (*CreateAdminAccountResponse, error) {
	out := new(CreateAdminAccountResponse)
	err := grpc.Invoke(ctx, "/RistoService.Risto/createAdminAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Risto service

type RistoServer interface {
	CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error)
	VerifyToken(context.Context, *VerifyTokenRequest) (*VerifyTokenResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	FetchUserByToken(context.Context, *FetchUserByTokenRequest) (*FetchUserByTokenResponse, error)
	HasAdminAccount(context.Context, *HasAdminAccountRequest) (*HasAdminAccountResponse, error)
	CreateAdminAccount(context.Context, *CreateAdminAccountRequest) (*CreateAdminAccountResponse, error)
}

func RegisterRistoServer(s *grpc.Server, srv RistoServer) {
	s.RegisterService(&_Risto_serviceDesc, srv)
}

func _Risto_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RistoServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RistoService.Risto/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RistoServer).CreateToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Risto_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RistoServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RistoService.Risto/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RistoServer).VerifyToken(ctx, req.(*VerifyTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Risto_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RistoServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RistoService.Risto/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RistoServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Risto_FetchUserByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchUserByTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RistoServer).FetchUserByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RistoService.Risto/FetchUserByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RistoServer).FetchUserByToken(ctx, req.(*FetchUserByTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Risto_HasAdminAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasAdminAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RistoServer).HasAdminAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RistoService.Risto/HasAdminAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RistoServer).HasAdminAccount(ctx, req.(*HasAdminAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Risto_CreateAdminAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdminAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RistoServer).CreateAdminAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RistoService.Risto/CreateAdminAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RistoServer).CreateAdminAccount(ctx, req.(*CreateAdminAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Risto_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RistoService.Risto",
	HandlerType: (*RistoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createToken",
			Handler:    _Risto_CreateToken_Handler,
		},
		{
			MethodName: "verifyToken",
			Handler:    _Risto_VerifyToken_Handler,
		},
		{
			MethodName: "createUser",
			Handler:    _Risto_CreateUser_Handler,
		},
		{
			MethodName: "fetchUserByToken",
			Handler:    _Risto_FetchUserByToken_Handler,
		},
		{
			MethodName: "hasAdminAccount",
			Handler:    _Risto_HasAdminAccount_Handler,
		},
		{
			MethodName: "createAdminAccount",
			Handler:    _Risto_CreateAdminAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "risto_service.proto",
}

func init() { proto.RegisterFile("risto_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x3d, 0xa8, 0x87, 0x59, 0x41, 0x89, 0xb7, 0x5e, 0x5c, 0xc5, 0x55, 0x4f, 0x3d, 0xe8,
	0x13, 0xac, 0x82, 0x78, 0xb5, 0x56, 0xaf, 0x5a, 0xe3, 0x94, 0x06, 0x31, 0xa9, 0x99, 0xb4, 0xe0,
	0x9b, 0xf9, 0x78, 0x92, 0xc6, 0xa4, 0x46, 0xd2, 0xed, 0xad, 0xcc, 0xff, 0xe5, 0xfb, 0x87, 0xa1,
	0x70, 0xa8, 0x05, 0x19, 0xf5, 0x4c, 0xa8, 0x7b, 0xc1, 0x31, 0x6f, 0xb5, 0x32, 0x8a, 0xed, 0x15,
	0x76, 0xf8, 0xe0, 0x66, 0xd9, 0xc2, 0xa8, 0x77, 0x94, 0x2e, 0xca, 0xa0, 0x23, 0xd4, 0xee, 0xfb,
	0xf2, 0x7b, 0x1b, 0x76, 0x06, 0x92, 0x95, 0xb0, 0xe0, 0x1a, 0x2b, 0x83, 0xa5, 0x45, 0xd9, 0x32,
	0xff, 0x2b, 0xc8, 0x6f, 0xc6, 0xa8, 0xc0, 0xcf, 0x0e, 0xc9, 0x64, 0xc7, 0x1b, 0x08, 0x6a, 0x95,
	0x24, 0x3c, 0xd9, 0xb2, 0xd6, 0x1e, 0xb5, 0xa8, 0xbf, 0x92, 0xd6, 0xa7, 0x31, 0x9a, 0xb0, 0x46,
	0x44, 0xb0, 0xde, 0x03, 0xb8, 0x5d, 0x1f, 0x09, 0x35, 0x3b, 0x4a, 0x2d, 0x62, 0x13, 0xef, 0x5c,
	0x4e, 0x03, 0x41, 0xc9, 0xe1, 0xa0, 0x46, 0xc3, 0x1b, 0x3b, 0xbe, 0xfe, 0xdd, 0x76, 0x15, 0xbf,
	0xbb, 0xfd, 0x97, 0x7b, 0xfd, 0xd9, 0x1c, 0x16, 0x4a, 0x5e, 0x60, 0xbf, 0xa9, 0x68, 0xfd, 0xf6,
	0x21, 0xe4, 0x9a, 0x73, 0xd5, 0x49, 0xc3, 0x4e, 0xe3, 0xc7, 0x77, 0x71, 0xec, 0x2b, 0x56, 0x33,
	0x54, 0x68, 0x10, 0xc0, 0xdc, 0x65, 0xa2, 0x92, 0xf3, 0xd4, 0x01, 0x52, 0x3d, 0x17, 0xf3, 0xa0,
	0xaf, 0x7a, 0xdd, 0x1d, 0xfe, 0xa0, 0xab, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0xd5, 0x5f,
	0x64, 0x7f, 0x02, 0x00, 0x00,
}
