// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package xjwt

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

// XjwtClient is the client API for Xjwt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type XjwtClient interface {
	CreateXjwt(ctx context.Context, in *CreateXjwtRequest, opts ...grpc.CallOption) (*CreateXjwtReply, error)
	UpdateXjwt(ctx context.Context, in *UpdateXjwtRequest, opts ...grpc.CallOption) (*UpdateXjwtReply, error)
	DeleteXjwt(ctx context.Context, in *DeleteXjwtRequest, opts ...grpc.CallOption) (*DeleteXjwtReply, error)
	GetXjwt(ctx context.Context, in *GetXjwtRequest, opts ...grpc.CallOption) (*GetXjwtReply, error)
	ListXjwt(ctx context.Context, in *ListXjwtRequest, opts ...grpc.CallOption) (*ListXjwtReply, error)
}

type xjwtClient struct {
	cc grpc.ClientConnInterface
}

func NewXjwtClient(cc grpc.ClientConnInterface) XjwtClient {
	return &xjwtClient{cc}
}

func (c *xjwtClient) CreateXjwt(ctx context.Context, in *CreateXjwtRequest, opts ...grpc.CallOption) (*CreateXjwtReply, error) {
	out := new(CreateXjwtReply)
	err := c.cc.Invoke(ctx, "/go_demo.xjwt.Xjwt/CreateXjwt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xjwtClient) UpdateXjwt(ctx context.Context, in *UpdateXjwtRequest, opts ...grpc.CallOption) (*UpdateXjwtReply, error) {
	out := new(UpdateXjwtReply)
	err := c.cc.Invoke(ctx, "/go_demo.xjwt.Xjwt/UpdateXjwt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xjwtClient) DeleteXjwt(ctx context.Context, in *DeleteXjwtRequest, opts ...grpc.CallOption) (*DeleteXjwtReply, error) {
	out := new(DeleteXjwtReply)
	err := c.cc.Invoke(ctx, "/go_demo.xjwt.Xjwt/DeleteXjwt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xjwtClient) GetXjwt(ctx context.Context, in *GetXjwtRequest, opts ...grpc.CallOption) (*GetXjwtReply, error) {
	out := new(GetXjwtReply)
	err := c.cc.Invoke(ctx, "/go_demo.xjwt.Xjwt/GetXjwt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xjwtClient) ListXjwt(ctx context.Context, in *ListXjwtRequest, opts ...grpc.CallOption) (*ListXjwtReply, error) {
	out := new(ListXjwtReply)
	err := c.cc.Invoke(ctx, "/go_demo.xjwt.Xjwt/ListXjwt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// XjwtServer is the server API for Xjwt service.
// All implementations must embed UnimplementedXjwtServer
// for forward compatibility
type XjwtServer interface {
	CreateXjwt(context.Context, *CreateXjwtRequest) (*CreateXjwtReply, error)
	UpdateXjwt(context.Context, *UpdateXjwtRequest) (*UpdateXjwtReply, error)
	DeleteXjwt(context.Context, *DeleteXjwtRequest) (*DeleteXjwtReply, error)
	GetXjwt(context.Context, *GetXjwtRequest) (*GetXjwtReply, error)
	ListXjwt(context.Context, *ListXjwtRequest) (*ListXjwtReply, error)
	mustEmbedUnimplementedXjwtServer()
}

// UnimplementedXjwtServer must be embedded to have forward compatible implementations.
type UnimplementedXjwtServer struct {
}

func (UnimplementedXjwtServer) CreateXjwt(context.Context, *CreateXjwtRequest) (*CreateXjwtReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateXjwt not implemented")
}
func (UnimplementedXjwtServer) UpdateXjwt(context.Context, *UpdateXjwtRequest) (*UpdateXjwtReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateXjwt not implemented")
}
func (UnimplementedXjwtServer) DeleteXjwt(context.Context, *DeleteXjwtRequest) (*DeleteXjwtReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteXjwt not implemented")
}
func (UnimplementedXjwtServer) GetXjwt(context.Context, *GetXjwtRequest) (*GetXjwtReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetXjwt not implemented")
}
func (UnimplementedXjwtServer) ListXjwt(context.Context, *ListXjwtRequest) (*ListXjwtReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListXjwt not implemented")
}
func (UnimplementedXjwtServer) mustEmbedUnimplementedXjwtServer() {}

// UnsafeXjwtServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to XjwtServer will
// result in compilation errors.
type UnsafeXjwtServer interface {
	mustEmbedUnimplementedXjwtServer()
}

func RegisterXjwtServer(s grpc.ServiceRegistrar, srv XjwtServer) {
	s.RegisterService(&Xjwt_ServiceDesc, srv)
}

func _Xjwt_CreateXjwt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateXjwtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XjwtServer).CreateXjwt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_demo.xjwt.Xjwt/CreateXjwt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XjwtServer).CreateXjwt(ctx, req.(*CreateXjwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xjwt_UpdateXjwt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateXjwtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XjwtServer).UpdateXjwt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_demo.xjwt.Xjwt/UpdateXjwt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XjwtServer).UpdateXjwt(ctx, req.(*UpdateXjwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xjwt_DeleteXjwt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteXjwtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XjwtServer).DeleteXjwt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_demo.xjwt.Xjwt/DeleteXjwt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XjwtServer).DeleteXjwt(ctx, req.(*DeleteXjwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xjwt_GetXjwt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetXjwtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XjwtServer).GetXjwt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_demo.xjwt.Xjwt/GetXjwt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XjwtServer).GetXjwt(ctx, req.(*GetXjwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xjwt_ListXjwt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListXjwtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XjwtServer).ListXjwt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_demo.xjwt.Xjwt/ListXjwt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XjwtServer).ListXjwt(ctx, req.(*ListXjwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Xjwt_ServiceDesc is the grpc.ServiceDesc for Xjwt service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Xjwt_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_demo.xjwt.Xjwt",
	HandlerType: (*XjwtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateXjwt",
			Handler:    _Xjwt_CreateXjwt_Handler,
		},
		{
			MethodName: "UpdateXjwt",
			Handler:    _Xjwt_UpdateXjwt_Handler,
		},
		{
			MethodName: "DeleteXjwt",
			Handler:    _Xjwt_DeleteXjwt_Handler,
		},
		{
			MethodName: "GetXjwt",
			Handler:    _Xjwt_GetXjwt_Handler,
		},
		{
			MethodName: "ListXjwt",
			Handler:    _Xjwt_ListXjwt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "xjwt.proto",
}
