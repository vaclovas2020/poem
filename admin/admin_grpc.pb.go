// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: admin/admin.proto

package admin

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

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	// Add new poem
	AddPoem(ctx context.Context, in *AdminPoem, opts ...grpc.CallOption) (*PoemResponse, error)
	// Add new category
	AddCategory(ctx context.Context, in *AdminCategory, opts ...grpc.CallOption) (*CategoryResponse, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) AddPoem(ctx context.Context, in *AdminPoem, opts ...grpc.CallOption) (*PoemResponse, error) {
	out := new(PoemResponse)
	err := c.cc.Invoke(ctx, "/Admin/AddPoem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AddCategory(ctx context.Context, in *AdminCategory, opts ...grpc.CallOption) (*CategoryResponse, error) {
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, "/Admin/AddCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	// Add new poem
	AddPoem(context.Context, *AdminPoem) (*PoemResponse, error)
	// Add new category
	AddCategory(context.Context, *AdminCategory) (*CategoryResponse, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) AddPoem(context.Context, *AdminPoem) (*PoemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPoem not implemented")
}
func (UnimplementedAdminServer) AddCategory(context.Context, *AdminCategory) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategory not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_AddPoem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminPoem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AddPoem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Admin/AddPoem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AddPoem(ctx, req.(*AdminPoem))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AddCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AddCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Admin/AddCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AddCategory(ctx, req.(*AdminCategory))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPoem",
			Handler:    _Admin_AddPoem_Handler,
		},
		{
			MethodName: "AddCategory",
			Handler:    _Admin_AddCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/admin.proto",
}
