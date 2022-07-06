// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: account.proto

package pb

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

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountClient interface {
	CreateAccountSecret(ctx context.Context, in *CreateAccountSecretRequest, opts ...grpc.CallOption) (*CreateAccountSecretResponse, error)
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	VersionOfAccountMicroService(ctx context.Context, in *VersionOfAccountMicroServiceRequest, opts ...grpc.CallOption) (*VersionOfAccountMicroServiceResponse, error)
	GenerateSetPasswordToken(ctx context.Context, in *GenerateSetPasswordTokenRequest, opts ...grpc.CallOption) (*GenerateSetPasswordTokenResponse, error)
	SetPassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*SetPasswordResponse, error)
	GenerateRecoveryPasswordToken(ctx context.Context, in *GenerateRecoveryPasswordTokenRequest, opts ...grpc.CallOption) (*GenerateRecoveryPasswordTokenResponse, error)
	RecoveryPassword(ctx context.Context, in *RecoveryPasswordRequest, opts ...grpc.CallOption) (*RecoveryPasswordResponse, error)
	GetAccountById(ctx context.Context, in *GetAccountByIdRequest, opts ...grpc.CallOption) (*GetAccountByIdResponse, error)
	GetAccountByEmail(ctx context.Context, in *GetAccountByEmailRequest, opts ...grpc.CallOption) (*GetAccountByEmailResponse, error)
	GetAccountByPhone(ctx context.Context, in *GetAccountByPhoneRequest, opts ...grpc.CallOption) (*GetAccountByPhoneResponse, error)
	GetAccountByDocument(ctx context.Context, in *GetAccountByDocumentRequest, opts ...grpc.CallOption) (*GetAccountByDocumentResponse, error)
	GetAllAccounts(ctx context.Context, in *GetAllAccountsRequest, opts ...grpc.CallOption) (*GetAllAccountsResponse, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error)
}

type accountClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountClient(cc grpc.ClientConnInterface) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) CreateAccountSecret(ctx context.Context, in *CreateAccountSecretRequest, opts ...grpc.CallOption) (*CreateAccountSecretResponse, error) {
	out := new(CreateAccountSecretResponse)
	err := c.cc.Invoke(ctx, "/iuris.account.v1.Account/CreateAccountSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/iuris.account.v1.Account/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) VersionOfAccountMicroService(ctx context.Context, in *VersionOfAccountMicroServiceRequest, opts ...grpc.CallOption) (*VersionOfAccountMicroServiceResponse, error) {
	out := new(VersionOfAccountMicroServiceResponse)
	err := c.cc.Invoke(ctx, "/iuris.account.v1.Account/VersionOfAccountMicroService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GenerateSetPasswordToken(ctx context.Context, in *GenerateSetPasswordTokenRequest, opts ...grpc.CallOption) (*GenerateSetPasswordTokenResponse, error) {
	out := new(GenerateSetPasswordTokenResponse)
	err := c.cc.Invoke(ctx, "/iuris.account.v1.Account/GenerateSetPasswordToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) SetPassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*SetPasswordResponse, error) {
	out := new(SetPasswordResponse)
	err := c.cc.Invoke(ctx, "/iuris.account.v1.Account/SetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GenerateRecoveryPasswordToken(ctx context.Context, in *GenerateRecoveryPasswordTokenRequest, opts ...grpc.CallOption) (*GenerateRecoveryPasswordTokenResponse, error) {
	out := new(GenerateRecoveryPasswordTokenResponse)
	err := c.cc.Invoke(ctx, "/iuris.account.v1.Account/GenerateRecoveryPasswordToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) RecoveryPassword(ctx context.Context, in *RecoveryPasswordRequest, opts ...grpc.CallOption) (*RecoveryPasswordResponse, error) {
	out := new(RecoveryPasswordResponse)
	err := c.cc.Invoke(ctx, "/iuris.account.v1.Account/RecoveryPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetAccountById(ctx context.Context, in *GetAccountByIdRequest, opts ...grpc.CallOption) (*GetAccountByIdResponse, error) {
	out := new(GetAccountByIdResponse)
	err := c.cc.Invoke(ctx, "/iuris.account.v1.Account/GetAccountById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetAccountByEmail(ctx context.Context, in *GetAccountByEmailRequest, opts ...grpc.CallOption) (*GetAccountByEmailResponse, error) {
	out := new(GetAccountByEmailResponse)
	err := c.cc.Invoke(ctx, "/iuris.account.v1.Account/GetAccountByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetAccountByPhone(ctx context.Context, in *GetAccountByPhoneRequest, opts ...grpc.CallOption) (*GetAccountByPhoneResponse, error) {
	out := new(GetAccountByPhoneResponse)
	err := c.cc.Invoke(ctx, "/iuris.account.v1.Account/GetAccountByPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetAccountByDocument(ctx context.Context, in *GetAccountByDocumentRequest, opts ...grpc.CallOption) (*GetAccountByDocumentResponse, error) {
	out := new(GetAccountByDocumentResponse)
	err := c.cc.Invoke(ctx, "/iuris.account.v1.Account/GetAccountByDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetAllAccounts(ctx context.Context, in *GetAllAccountsRequest, opts ...grpc.CallOption) (*GetAllAccountsResponse, error) {
	out := new(GetAllAccountsResponse)
	err := c.cc.Invoke(ctx, "/iuris.account.v1.Account/GetAllAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error) {
	out := new(UpdateAccountResponse)
	err := c.cc.Invoke(ctx, "/iuris.account.v1.Account/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error) {
	out := new(DeleteAccountResponse)
	err := c.cc.Invoke(ctx, "/iuris.account.v1.Account/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
// All implementations should embed UnimplementedAccountServer
// for forward compatibility
type AccountServer interface {
	CreateAccountSecret(context.Context, *CreateAccountSecretRequest) (*CreateAccountSecretResponse, error)
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	VersionOfAccountMicroService(context.Context, *VersionOfAccountMicroServiceRequest) (*VersionOfAccountMicroServiceResponse, error)
	GenerateSetPasswordToken(context.Context, *GenerateSetPasswordTokenRequest) (*GenerateSetPasswordTokenResponse, error)
	SetPassword(context.Context, *SetPasswordRequest) (*SetPasswordResponse, error)
	GenerateRecoveryPasswordToken(context.Context, *GenerateRecoveryPasswordTokenRequest) (*GenerateRecoveryPasswordTokenResponse, error)
	RecoveryPassword(context.Context, *RecoveryPasswordRequest) (*RecoveryPasswordResponse, error)
	GetAccountById(context.Context, *GetAccountByIdRequest) (*GetAccountByIdResponse, error)
	GetAccountByEmail(context.Context, *GetAccountByEmailRequest) (*GetAccountByEmailResponse, error)
	GetAccountByPhone(context.Context, *GetAccountByPhoneRequest) (*GetAccountByPhoneResponse, error)
	GetAccountByDocument(context.Context, *GetAccountByDocumentRequest) (*GetAccountByDocumentResponse, error)
	GetAllAccounts(context.Context, *GetAllAccountsRequest) (*GetAllAccountsResponse, error)
	UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error)
}

// UnimplementedAccountServer should be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (UnimplementedAccountServer) CreateAccountSecret(context.Context, *CreateAccountSecretRequest) (*CreateAccountSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccountSecret not implemented")
}
func (UnimplementedAccountServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAccountServer) VersionOfAccountMicroService(context.Context, *VersionOfAccountMicroServiceRequest) (*VersionOfAccountMicroServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VersionOfAccountMicroService not implemented")
}
func (UnimplementedAccountServer) GenerateSetPasswordToken(context.Context, *GenerateSetPasswordTokenRequest) (*GenerateSetPasswordTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateSetPasswordToken not implemented")
}
func (UnimplementedAccountServer) SetPassword(context.Context, *SetPasswordRequest) (*SetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPassword not implemented")
}
func (UnimplementedAccountServer) GenerateRecoveryPasswordToken(context.Context, *GenerateRecoveryPasswordTokenRequest) (*GenerateRecoveryPasswordTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateRecoveryPasswordToken not implemented")
}
func (UnimplementedAccountServer) RecoveryPassword(context.Context, *RecoveryPasswordRequest) (*RecoveryPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoveryPassword not implemented")
}
func (UnimplementedAccountServer) GetAccountById(context.Context, *GetAccountByIdRequest) (*GetAccountByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountById not implemented")
}
func (UnimplementedAccountServer) GetAccountByEmail(context.Context, *GetAccountByEmailRequest) (*GetAccountByEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByEmail not implemented")
}
func (UnimplementedAccountServer) GetAccountByPhone(context.Context, *GetAccountByPhoneRequest) (*GetAccountByPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByPhone not implemented")
}
func (UnimplementedAccountServer) GetAccountByDocument(context.Context, *GetAccountByDocumentRequest) (*GetAccountByDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByDocument not implemented")
}
func (UnimplementedAccountServer) GetAllAccounts(context.Context, *GetAllAccountsRequest) (*GetAllAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAccounts not implemented")
}
func (UnimplementedAccountServer) UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedAccountServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}

// UnsafeAccountServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServer will
// result in compilation errors.
type UnsafeAccountServer interface {
	mustEmbedUnimplementedAccountServer()
}

func RegisterAccountServer(s grpc.ServiceRegistrar, srv AccountServer) {
	s.RegisterService(&Account_ServiceDesc, srv)
}

func _Account_CreateAccountSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).CreateAccountSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iuris.account.v1.Account/CreateAccountSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).CreateAccountSecret(ctx, req.(*CreateAccountSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iuris.account.v1.Account/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_VersionOfAccountMicroService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionOfAccountMicroServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).VersionOfAccountMicroService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iuris.account.v1.Account/VersionOfAccountMicroService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).VersionOfAccountMicroService(ctx, req.(*VersionOfAccountMicroServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GenerateSetPasswordToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateSetPasswordTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GenerateSetPasswordToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iuris.account.v1.Account/GenerateSetPasswordToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GenerateSetPasswordToken(ctx, req.(*GenerateSetPasswordTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_SetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).SetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iuris.account.v1.Account/SetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).SetPassword(ctx, req.(*SetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GenerateRecoveryPasswordToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRecoveryPasswordTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GenerateRecoveryPasswordToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iuris.account.v1.Account/GenerateRecoveryPasswordToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GenerateRecoveryPasswordToken(ctx, req.(*GenerateRecoveryPasswordTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_RecoveryPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoveryPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).RecoveryPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iuris.account.v1.Account/RecoveryPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).RecoveryPassword(ctx, req.(*RecoveryPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetAccountById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetAccountById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iuris.account.v1.Account/GetAccountById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetAccountById(ctx, req.(*GetAccountByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetAccountByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetAccountByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iuris.account.v1.Account/GetAccountByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetAccountByEmail(ctx, req.(*GetAccountByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetAccountByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountByPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetAccountByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iuris.account.v1.Account/GetAccountByPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetAccountByPhone(ctx, req.(*GetAccountByPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetAccountByDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountByDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetAccountByDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iuris.account.v1.Account/GetAccountByDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetAccountByDocument(ctx, req.(*GetAccountByDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetAllAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetAllAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iuris.account.v1.Account/GetAllAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetAllAccounts(ctx, req.(*GetAllAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iuris.account.v1.Account/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).UpdateAccount(ctx, req.(*UpdateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iuris.account.v1.Account/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Account_ServiceDesc is the grpc.ServiceDesc for Account service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Account_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iuris.account.v1.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccountSecret",
			Handler:    _Account_CreateAccountSecret_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _Account_CreateAccount_Handler,
		},
		{
			MethodName: "VersionOfAccountMicroService",
			Handler:    _Account_VersionOfAccountMicroService_Handler,
		},
		{
			MethodName: "GenerateSetPasswordToken",
			Handler:    _Account_GenerateSetPasswordToken_Handler,
		},
		{
			MethodName: "SetPassword",
			Handler:    _Account_SetPassword_Handler,
		},
		{
			MethodName: "GenerateRecoveryPasswordToken",
			Handler:    _Account_GenerateRecoveryPasswordToken_Handler,
		},
		{
			MethodName: "RecoveryPassword",
			Handler:    _Account_RecoveryPassword_Handler,
		},
		{
			MethodName: "GetAccountById",
			Handler:    _Account_GetAccountById_Handler,
		},
		{
			MethodName: "GetAccountByEmail",
			Handler:    _Account_GetAccountByEmail_Handler,
		},
		{
			MethodName: "GetAccountByPhone",
			Handler:    _Account_GetAccountByPhone_Handler,
		},
		{
			MethodName: "GetAccountByDocument",
			Handler:    _Account_GetAccountByDocument_Handler,
		},
		{
			MethodName: "GetAllAccounts",
			Handler:    _Account_GetAllAccounts_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _Account_UpdateAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _Account_DeleteAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}
