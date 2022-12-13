// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: protobuf_files/systemicdb.proto

package systemicdbGrpc

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

// GetClient is the client API for Get service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetClient interface {
	ByKey(ctx context.Context, in *GetByKeyRequest, opts ...grpc.CallOption) (*SingleResponse, error)
	ByTag(ctx context.Context, in *GetByTagRequest, opts ...grpc.CallOption) (*RepeatedResponse, error)
}

type getClient struct {
	cc grpc.ClientConnInterface
}

func NewGetClient(cc grpc.ClientConnInterface) GetClient {
	return &getClient{cc}
}

func (c *getClient) ByKey(ctx context.Context, in *GetByKeyRequest, opts ...grpc.CallOption) (*SingleResponse, error) {
	out := new(SingleResponse)
	err := c.cc.Invoke(ctx, "/systemicdbGrpc.Get/ByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getClient) ByTag(ctx context.Context, in *GetByTagRequest, opts ...grpc.CallOption) (*RepeatedResponse, error) {
	out := new(RepeatedResponse)
	err := c.cc.Invoke(ctx, "/systemicdbGrpc.Get/ByTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetServer is the server API for Get service.
// All implementations must embed UnimplementedGetServer
// for forward compatibility
type GetServer interface {
	ByKey(context.Context, *GetByKeyRequest) (*SingleResponse, error)
	ByTag(context.Context, *GetByTagRequest) (*RepeatedResponse, error)
	mustEmbedUnimplementedGetServer()
}

// UnimplementedGetServer must be embedded to have forward compatible implementations.
type UnimplementedGetServer struct {
}

func (UnimplementedGetServer) ByKey(context.Context, *GetByKeyRequest) (*SingleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByKey not implemented")
}
func (UnimplementedGetServer) ByTag(context.Context, *GetByTagRequest) (*RepeatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByTag not implemented")
}
func (UnimplementedGetServer) mustEmbedUnimplementedGetServer() {}

// UnsafeGetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetServer will
// result in compilation errors.
type UnsafeGetServer interface {
	mustEmbedUnimplementedGetServer()
}

func RegisterGetServer(s grpc.ServiceRegistrar, srv GetServer) {
	s.RegisterService(&Get_ServiceDesc, srv)
}

func _Get_ByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetServer).ByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systemicdbGrpc.Get/ByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetServer).ByKey(ctx, req.(*GetByKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Get_ByTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetServer).ByTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systemicdbGrpc.Get/ByTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetServer).ByTag(ctx, req.(*GetByTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Get_ServiceDesc is the grpc.ServiceDesc for Get service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Get_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "systemicdbGrpc.Get",
	HandlerType: (*GetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ByKey",
			Handler:    _Get_ByKey_Handler,
		},
		{
			MethodName: "ByTag",
			Handler:    _Get_ByTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf_files/systemicdb.proto",
}

// SetClient is the client API for Set service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SetClient interface {
	ByKey(ctx context.Context, in *SetByKeyRequest, opts ...grpc.CallOption) (*BoolResponse, error)
	ByTag(ctx context.Context, in *SetByTagRequest, opts ...grpc.CallOption) (*BoolResponse, error)
}

type setClient struct {
	cc grpc.ClientConnInterface
}

func NewSetClient(cc grpc.ClientConnInterface) SetClient {
	return &setClient{cc}
}

func (c *setClient) ByKey(ctx context.Context, in *SetByKeyRequest, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/systemicdbGrpc.Set/ByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *setClient) ByTag(ctx context.Context, in *SetByTagRequest, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/systemicdbGrpc.Set/ByTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SetServer is the server API for Set service.
// All implementations must embed UnimplementedSetServer
// for forward compatibility
type SetServer interface {
	ByKey(context.Context, *SetByKeyRequest) (*BoolResponse, error)
	ByTag(context.Context, *SetByTagRequest) (*BoolResponse, error)
	mustEmbedUnimplementedSetServer()
}

// UnimplementedSetServer must be embedded to have forward compatible implementations.
type UnimplementedSetServer struct {
}

func (UnimplementedSetServer) ByKey(context.Context, *SetByKeyRequest) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByKey not implemented")
}
func (UnimplementedSetServer) ByTag(context.Context, *SetByTagRequest) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByTag not implemented")
}
func (UnimplementedSetServer) mustEmbedUnimplementedSetServer() {}

// UnsafeSetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SetServer will
// result in compilation errors.
type UnsafeSetServer interface {
	mustEmbedUnimplementedSetServer()
}

func RegisterSetServer(s grpc.ServiceRegistrar, srv SetServer) {
	s.RegisterService(&Set_ServiceDesc, srv)
}

func _Set_ByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetByKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetServer).ByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systemicdbGrpc.Set/ByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetServer).ByKey(ctx, req.(*SetByKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Set_ByTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetByTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetServer).ByTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systemicdbGrpc.Set/ByTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetServer).ByTag(ctx, req.(*SetByTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Set_ServiceDesc is the grpc.ServiceDesc for Set service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Set_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "systemicdbGrpc.Set",
	HandlerType: (*SetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ByKey",
			Handler:    _Set_ByKey_Handler,
		},
		{
			MethodName: "ByTag",
			Handler:    _Set_ByTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf_files/systemicdb.proto",
}

// DelClient is the client API for Del service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DelClient interface {
	ByKey(ctx context.Context, in *GetByKeyRequest, opts ...grpc.CallOption) (*BoolResponse, error)
	ByTag(ctx context.Context, in *GetByTagRequest, opts ...grpc.CallOption) (*BoolResponse, error)
}

type delClient struct {
	cc grpc.ClientConnInterface
}

func NewDelClient(cc grpc.ClientConnInterface) DelClient {
	return &delClient{cc}
}

func (c *delClient) ByKey(ctx context.Context, in *GetByKeyRequest, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/systemicdbGrpc.Del/ByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *delClient) ByTag(ctx context.Context, in *GetByTagRequest, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/systemicdbGrpc.Del/ByTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DelServer is the server API for Del service.
// All implementations must embed UnimplementedDelServer
// for forward compatibility
type DelServer interface {
	ByKey(context.Context, *GetByKeyRequest) (*BoolResponse, error)
	ByTag(context.Context, *GetByTagRequest) (*BoolResponse, error)
	mustEmbedUnimplementedDelServer()
}

// UnimplementedDelServer must be embedded to have forward compatible implementations.
type UnimplementedDelServer struct {
}

func (UnimplementedDelServer) ByKey(context.Context, *GetByKeyRequest) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByKey not implemented")
}
func (UnimplementedDelServer) ByTag(context.Context, *GetByTagRequest) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByTag not implemented")
}
func (UnimplementedDelServer) mustEmbedUnimplementedDelServer() {}

// UnsafeDelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DelServer will
// result in compilation errors.
type UnsafeDelServer interface {
	mustEmbedUnimplementedDelServer()
}

func RegisterDelServer(s grpc.ServiceRegistrar, srv DelServer) {
	s.RegisterService(&Del_ServiceDesc, srv)
}

func _Del_ByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelServer).ByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systemicdbGrpc.Del/ByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelServer).ByKey(ctx, req.(*GetByKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Del_ByTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelServer).ByTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systemicdbGrpc.Del/ByTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelServer).ByTag(ctx, req.(*GetByTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Del_ServiceDesc is the grpc.ServiceDesc for Del service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Del_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "systemicdbGrpc.Del",
	HandlerType: (*DelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ByKey",
			Handler:    _Del_ByKey_Handler,
		},
		{
			MethodName: "ByTag",
			Handler:    _Del_ByTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf_files/systemicdb.proto",
}