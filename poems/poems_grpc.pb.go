// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: poems/poems.proto

package poems

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

// PoemsClient is the client API for Poems service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PoemsClient interface {
	// Get poems list
	GetPoems(ctx context.Context, in *PoemsRequest, opts ...grpc.CallOption) (*PoemsResponse, error)
	// Get categories list
	GetCategories(ctx context.Context, in *CategoriesRequest, opts ...grpc.CallOption) (*CategoriesResponse, error)
}

type poemsClient struct {
	cc grpc.ClientConnInterface
}

func NewPoemsClient(cc grpc.ClientConnInterface) PoemsClient {
	return &poemsClient{cc}
}

func (c *poemsClient) GetPoems(ctx context.Context, in *PoemsRequest, opts ...grpc.CallOption) (*PoemsResponse, error) {
	out := new(PoemsResponse)
	err := c.cc.Invoke(ctx, "/Poems/GetPoems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poemsClient) GetCategories(ctx context.Context, in *CategoriesRequest, opts ...grpc.CallOption) (*CategoriesResponse, error) {
	out := new(CategoriesResponse)
	err := c.cc.Invoke(ctx, "/Poems/GetCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PoemsServer is the server API for Poems service.
// All implementations must embed UnimplementedPoemsServer
// for forward compatibility
type PoemsServer interface {
	// Get poems list
	GetPoems(context.Context, *PoemsRequest) (*PoemsResponse, error)
	// Get categories list
	GetCategories(context.Context, *CategoriesRequest) (*CategoriesResponse, error)
	mustEmbedUnimplementedPoemsServer()
}

// UnimplementedPoemsServer must be embedded to have forward compatible implementations.
type UnimplementedPoemsServer struct {
}

func (UnimplementedPoemsServer) GetPoems(context.Context, *PoemsRequest) (*PoemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoems not implemented")
}
func (UnimplementedPoemsServer) GetCategories(context.Context, *CategoriesRequest) (*CategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (UnimplementedPoemsServer) mustEmbedUnimplementedPoemsServer() {}

// UnsafePoemsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PoemsServer will
// result in compilation errors.
type UnsafePoemsServer interface {
	mustEmbedUnimplementedPoemsServer()
}

func RegisterPoemsServer(s grpc.ServiceRegistrar, srv PoemsServer) {
	s.RegisterService(&Poems_ServiceDesc, srv)
}

func _Poems_GetPoems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PoemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoemsServer).GetPoems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Poems/GetPoems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoemsServer).GetPoems(ctx, req.(*PoemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poems_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoemsServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Poems/GetCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoemsServer).GetCategories(ctx, req.(*CategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Poems_ServiceDesc is the grpc.ServiceDesc for Poems service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Poems_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Poems",
	HandlerType: (*PoemsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPoems",
			Handler:    _Poems_GetPoems_Handler,
		},
		{
			MethodName: "GetCategories",
			Handler:    _Poems_GetCategories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "poems/poems.proto",
}
