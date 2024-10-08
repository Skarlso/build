// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// https://developers.google.com/protocol-buffers/docs/proto3

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: maintner/maintnerd/apipb/api.proto

package apipb

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
	MaintnerService_HasAncestor_FullMethodName    = "/apipb.MaintnerService/HasAncestor"
	MaintnerService_GetRef_FullMethodName         = "/apipb.MaintnerService/GetRef"
	MaintnerService_GoFindTryWork_FullMethodName  = "/apipb.MaintnerService/GoFindTryWork"
	MaintnerService_ListGoReleases_FullMethodName = "/apipb.MaintnerService/ListGoReleases"
	MaintnerService_GetDashboard_FullMethodName   = "/apipb.MaintnerService/GetDashboard"
)

// MaintnerServiceClient is the client API for MaintnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MaintnerServiceClient interface {
	// HasAncestor reports whether one commit contains another commit
	// in its git history.
	HasAncestor(ctx context.Context, in *HasAncestorRequest, opts ...grpc.CallOption) (*HasAncestorResponse, error)
	// GetRef returns information about a git ref.
	GetRef(ctx context.Context, in *GetRefRequest, opts ...grpc.CallOption) (*GetRefResponse, error)
	// GoFindTryWork finds trybot work for the coordinator to build & test.
	GoFindTryWork(ctx context.Context, in *GoFindTryWorkRequest, opts ...grpc.CallOption) (*GoFindTryWorkResponse, error)
	// ListGoReleases lists Go releases sorted by version with latest first.
	//
	// A release is considered to exist for each git tag named "goX", "goX.Y", or
	// "goX.Y.Z", as long as it has a corresponding "release-branch.goX" or
	// "release-branch.goX.Y" release branch.
	//
	// ListGoReleases returns only the latest patch versions of releases which
	// are considered supported per policy. For example, Go 1.12.6 and 1.11.11.
	// The response is guaranteed to have two versions, otherwise an error
	// is returned.
	ListGoReleases(ctx context.Context, in *ListGoReleasesRequest, opts ...grpc.CallOption) (*ListGoReleasesResponse, error)
	// GetDashboard returns the information for the build.golang.org
	// dashboard. It does not (at least currently)
	// contain any pass/fail information; it only contains information on the branches
	// and commits themselves.
	GetDashboard(ctx context.Context, in *DashboardRequest, opts ...grpc.CallOption) (*DashboardResponse, error)
}

type maintnerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMaintnerServiceClient(cc grpc.ClientConnInterface) MaintnerServiceClient {
	return &maintnerServiceClient{cc}
}

func (c *maintnerServiceClient) HasAncestor(ctx context.Context, in *HasAncestorRequest, opts ...grpc.CallOption) (*HasAncestorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HasAncestorResponse)
	err := c.cc.Invoke(ctx, MaintnerService_HasAncestor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *maintnerServiceClient) GetRef(ctx context.Context, in *GetRefRequest, opts ...grpc.CallOption) (*GetRefResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRefResponse)
	err := c.cc.Invoke(ctx, MaintnerService_GetRef_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *maintnerServiceClient) GoFindTryWork(ctx context.Context, in *GoFindTryWorkRequest, opts ...grpc.CallOption) (*GoFindTryWorkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GoFindTryWorkResponse)
	err := c.cc.Invoke(ctx, MaintnerService_GoFindTryWork_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *maintnerServiceClient) ListGoReleases(ctx context.Context, in *ListGoReleasesRequest, opts ...grpc.CallOption) (*ListGoReleasesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListGoReleasesResponse)
	err := c.cc.Invoke(ctx, MaintnerService_ListGoReleases_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *maintnerServiceClient) GetDashboard(ctx context.Context, in *DashboardRequest, opts ...grpc.CallOption) (*DashboardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DashboardResponse)
	err := c.cc.Invoke(ctx, MaintnerService_GetDashboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MaintnerServiceServer is the server API for MaintnerService service.
// All implementations must embed UnimplementedMaintnerServiceServer
// for forward compatibility.
type MaintnerServiceServer interface {
	// HasAncestor reports whether one commit contains another commit
	// in its git history.
	HasAncestor(context.Context, *HasAncestorRequest) (*HasAncestorResponse, error)
	// GetRef returns information about a git ref.
	GetRef(context.Context, *GetRefRequest) (*GetRefResponse, error)
	// GoFindTryWork finds trybot work for the coordinator to build & test.
	GoFindTryWork(context.Context, *GoFindTryWorkRequest) (*GoFindTryWorkResponse, error)
	// ListGoReleases lists Go releases sorted by version with latest first.
	//
	// A release is considered to exist for each git tag named "goX", "goX.Y", or
	// "goX.Y.Z", as long as it has a corresponding "release-branch.goX" or
	// "release-branch.goX.Y" release branch.
	//
	// ListGoReleases returns only the latest patch versions of releases which
	// are considered supported per policy. For example, Go 1.12.6 and 1.11.11.
	// The response is guaranteed to have two versions, otherwise an error
	// is returned.
	ListGoReleases(context.Context, *ListGoReleasesRequest) (*ListGoReleasesResponse, error)
	// GetDashboard returns the information for the build.golang.org
	// dashboard. It does not (at least currently)
	// contain any pass/fail information; it only contains information on the branches
	// and commits themselves.
	GetDashboard(context.Context, *DashboardRequest) (*DashboardResponse, error)
	mustEmbedUnimplementedMaintnerServiceServer()
}

// UnimplementedMaintnerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMaintnerServiceServer struct{}

func (UnimplementedMaintnerServiceServer) HasAncestor(context.Context, *HasAncestorRequest) (*HasAncestorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasAncestor not implemented")
}
func (UnimplementedMaintnerServiceServer) GetRef(context.Context, *GetRefRequest) (*GetRefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRef not implemented")
}
func (UnimplementedMaintnerServiceServer) GoFindTryWork(context.Context, *GoFindTryWorkRequest) (*GoFindTryWorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoFindTryWork not implemented")
}
func (UnimplementedMaintnerServiceServer) ListGoReleases(context.Context, *ListGoReleasesRequest) (*ListGoReleasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGoReleases not implemented")
}
func (UnimplementedMaintnerServiceServer) GetDashboard(context.Context, *DashboardRequest) (*DashboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDashboard not implemented")
}
func (UnimplementedMaintnerServiceServer) mustEmbedUnimplementedMaintnerServiceServer() {}
func (UnimplementedMaintnerServiceServer) testEmbeddedByValue()                         {}

// UnsafeMaintnerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MaintnerServiceServer will
// result in compilation errors.
type UnsafeMaintnerServiceServer interface {
	mustEmbedUnimplementedMaintnerServiceServer()
}

func RegisterMaintnerServiceServer(s grpc.ServiceRegistrar, srv MaintnerServiceServer) {
	// If the following call pancis, it indicates UnimplementedMaintnerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MaintnerService_ServiceDesc, srv)
}

func _MaintnerService_HasAncestor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasAncestorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintnerServiceServer).HasAncestor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaintnerService_HasAncestor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintnerServiceServer).HasAncestor(ctx, req.(*HasAncestorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaintnerService_GetRef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintnerServiceServer).GetRef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaintnerService_GetRef_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintnerServiceServer).GetRef(ctx, req.(*GetRefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaintnerService_GoFindTryWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoFindTryWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintnerServiceServer).GoFindTryWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaintnerService_GoFindTryWork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintnerServiceServer).GoFindTryWork(ctx, req.(*GoFindTryWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaintnerService_ListGoReleases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGoReleasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintnerServiceServer).ListGoReleases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaintnerService_ListGoReleases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintnerServiceServer).ListGoReleases(ctx, req.(*ListGoReleasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaintnerService_GetDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintnerServiceServer).GetDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaintnerService_GetDashboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintnerServiceServer).GetDashboard(ctx, req.(*DashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MaintnerService_ServiceDesc is the grpc.ServiceDesc for MaintnerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MaintnerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apipb.MaintnerService",
	HandlerType: (*MaintnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HasAncestor",
			Handler:    _MaintnerService_HasAncestor_Handler,
		},
		{
			MethodName: "GetRef",
			Handler:    _MaintnerService_GetRef_Handler,
		},
		{
			MethodName: "GoFindTryWork",
			Handler:    _MaintnerService_GoFindTryWork_Handler,
		},
		{
			MethodName: "ListGoReleases",
			Handler:    _MaintnerService_ListGoReleases_Handler,
		},
		{
			MethodName: "GetDashboard",
			Handler:    _MaintnerService_GetDashboard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "maintner/maintnerd/apipb/api.proto",
}
