// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: careerhub/posting_service/provider_grpc/provider_grpc/grpc.proto

package provider_grpc

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

// ProviderGrpcClient is the client API for ProviderGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderGrpcClient interface {
	IsCompanyRegistered(ctx context.Context, in *CompanyId, opts ...grpc.CallOption) (*BoolResponse, error)
	GetAllHiring(ctx context.Context, in *Site, opts ...grpc.CallOption) (*JobPostings, error)
	CloseJobPostings(ctx context.Context, in *JobPostings, opts ...grpc.CallOption) (*BoolResponse, error)
	RegisterJobPostingInfo(ctx context.Context, in *JobPostingInfo, opts ...grpc.CallOption) (*BoolResponse, error)
	RegisterCompany(ctx context.Context, in *Company, opts ...grpc.CallOption) (*BoolResponse, error)
}

type providerGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderGrpcClient(cc grpc.ClientConnInterface) ProviderGrpcClient {
	return &providerGrpcClient{cc}
}

func (c *providerGrpcClient) IsCompanyRegistered(ctx context.Context, in *CompanyId, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/careerhub.posting_service.provider_grpc.ProviderGrpc/IsCompanyRegistered", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerGrpcClient) GetAllHiring(ctx context.Context, in *Site, opts ...grpc.CallOption) (*JobPostings, error) {
	out := new(JobPostings)
	err := c.cc.Invoke(ctx, "/careerhub.posting_service.provider_grpc.ProviderGrpc/GetAllHiring", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerGrpcClient) CloseJobPostings(ctx context.Context, in *JobPostings, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/careerhub.posting_service.provider_grpc.ProviderGrpc/CloseJobPostings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerGrpcClient) RegisterJobPostingInfo(ctx context.Context, in *JobPostingInfo, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/careerhub.posting_service.provider_grpc.ProviderGrpc/RegisterJobPostingInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerGrpcClient) RegisterCompany(ctx context.Context, in *Company, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/careerhub.posting_service.provider_grpc.ProviderGrpc/RegisterCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderGrpcServer is the server API for ProviderGrpc service.
// All implementations must embed UnimplementedProviderGrpcServer
// for forward compatibility
type ProviderGrpcServer interface {
	IsCompanyRegistered(context.Context, *CompanyId) (*BoolResponse, error)
	GetAllHiring(context.Context, *Site) (*JobPostings, error)
	CloseJobPostings(context.Context, *JobPostings) (*BoolResponse, error)
	RegisterJobPostingInfo(context.Context, *JobPostingInfo) (*BoolResponse, error)
	RegisterCompany(context.Context, *Company) (*BoolResponse, error)
	mustEmbedUnimplementedProviderGrpcServer()
}

// UnimplementedProviderGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedProviderGrpcServer struct {
}

func (UnimplementedProviderGrpcServer) IsCompanyRegistered(context.Context, *CompanyId) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsCompanyRegistered not implemented")
}
func (UnimplementedProviderGrpcServer) GetAllHiring(context.Context, *Site) (*JobPostings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllHiring not implemented")
}
func (UnimplementedProviderGrpcServer) CloseJobPostings(context.Context, *JobPostings) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseJobPostings not implemented")
}
func (UnimplementedProviderGrpcServer) RegisterJobPostingInfo(context.Context, *JobPostingInfo) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterJobPostingInfo not implemented")
}
func (UnimplementedProviderGrpcServer) RegisterCompany(context.Context, *Company) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCompany not implemented")
}
func (UnimplementedProviderGrpcServer) mustEmbedUnimplementedProviderGrpcServer() {}

// UnsafeProviderGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderGrpcServer will
// result in compilation errors.
type UnsafeProviderGrpcServer interface {
	mustEmbedUnimplementedProviderGrpcServer()
}

func RegisterProviderGrpcServer(s grpc.ServiceRegistrar, srv ProviderGrpcServer) {
	s.RegisterService(&ProviderGrpc_ServiceDesc, srv)
}

func _ProviderGrpc_IsCompanyRegistered_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderGrpcServer).IsCompanyRegistered(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.posting_service.provider_grpc.ProviderGrpc/IsCompanyRegistered",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderGrpcServer).IsCompanyRegistered(ctx, req.(*CompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderGrpc_GetAllHiring_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Site)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderGrpcServer).GetAllHiring(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.posting_service.provider_grpc.ProviderGrpc/GetAllHiring",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderGrpcServer).GetAllHiring(ctx, req.(*Site))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderGrpc_CloseJobPostings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobPostings)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderGrpcServer).CloseJobPostings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.posting_service.provider_grpc.ProviderGrpc/CloseJobPostings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderGrpcServer).CloseJobPostings(ctx, req.(*JobPostings))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderGrpc_RegisterJobPostingInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobPostingInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderGrpcServer).RegisterJobPostingInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.posting_service.provider_grpc.ProviderGrpc/RegisterJobPostingInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderGrpcServer).RegisterJobPostingInfo(ctx, req.(*JobPostingInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderGrpc_RegisterCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Company)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderGrpcServer).RegisterCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.posting_service.provider_grpc.ProviderGrpc/RegisterCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderGrpcServer).RegisterCompany(ctx, req.(*Company))
	}
	return interceptor(ctx, in, info, handler)
}

// ProviderGrpc_ServiceDesc is the grpc.ServiceDesc for ProviderGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProviderGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "careerhub.posting_service.provider_grpc.ProviderGrpc",
	HandlerType: (*ProviderGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsCompanyRegistered",
			Handler:    _ProviderGrpc_IsCompanyRegistered_Handler,
		},
		{
			MethodName: "GetAllHiring",
			Handler:    _ProviderGrpc_GetAllHiring_Handler,
		},
		{
			MethodName: "CloseJobPostings",
			Handler:    _ProviderGrpc_CloseJobPostings_Handler,
		},
		{
			MethodName: "RegisterJobPostingInfo",
			Handler:    _ProviderGrpc_RegisterJobPostingInfo_Handler,
		},
		{
			MethodName: "RegisterCompany",
			Handler:    _ProviderGrpc_RegisterCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "careerhub/posting_service/provider_grpc/provider_grpc/grpc.proto",
}
