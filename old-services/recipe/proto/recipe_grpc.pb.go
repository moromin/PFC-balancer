// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// RecipeServiceClient is the client API for RecipeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecipeServiceClient interface {
	CreateRecipe(ctx context.Context, in *CreateRecipeRequest, opts ...grpc.CallOption) (*CreateRecipeResponse, error)
	ReadRecipe(ctx context.Context, in *ReadRecipeRequest, opts ...grpc.CallOption) (*ReadRecipeResponse, error)
}

type recipeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecipeServiceClient(cc grpc.ClientConnInterface) RecipeServiceClient {
	return &recipeServiceClient{cc}
}

func (c *recipeServiceClient) CreateRecipe(ctx context.Context, in *CreateRecipeRequest, opts ...grpc.CallOption) (*CreateRecipeResponse, error) {
	out := new(CreateRecipeResponse)
	err := c.cc.Invoke(ctx, "/food.RecipeService/CreateRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) ReadRecipe(ctx context.Context, in *ReadRecipeRequest, opts ...grpc.CallOption) (*ReadRecipeResponse, error) {
	out := new(ReadRecipeResponse)
	err := c.cc.Invoke(ctx, "/food.RecipeService/ReadRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecipeServiceServer is the server API for RecipeService service.
// All implementations should embed UnimplementedRecipeServiceServer
// for forward compatibility
type RecipeServiceServer interface {
	CreateRecipe(context.Context, *CreateRecipeRequest) (*CreateRecipeResponse, error)
	ReadRecipe(context.Context, *ReadRecipeRequest) (*ReadRecipeResponse, error)
}

// UnimplementedRecipeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRecipeServiceServer struct {
}

func (UnimplementedRecipeServiceServer) CreateRecipe(context.Context, *CreateRecipeRequest) (*CreateRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecipe not implemented")
}
func (UnimplementedRecipeServiceServer) ReadRecipe(context.Context, *ReadRecipeRequest) (*ReadRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadRecipe not implemented")
}

// UnsafeRecipeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecipeServiceServer will
// result in compilation errors.
type UnsafeRecipeServiceServer interface {
	mustEmbedUnimplementedRecipeServiceServer()
}

func RegisterRecipeServiceServer(s grpc.ServiceRegistrar, srv RecipeServiceServer) {
	s.RegisterService(&RecipeService_ServiceDesc, srv)
}

func _RecipeService_CreateRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).CreateRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/food.RecipeService/CreateRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).CreateRecipe(ctx, req.(*CreateRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_ReadRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).ReadRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/food.RecipeService/ReadRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).ReadRecipe(ctx, req.(*ReadRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecipeService_ServiceDesc is the grpc.ServiceDesc for RecipeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecipeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "food.RecipeService",
	HandlerType: (*RecipeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRecipe",
			Handler:    _RecipeService_CreateRecipe_Handler,
		},
		{
			MethodName: "ReadRecipe",
			Handler:    _RecipeService_ReadRecipe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/recipe.proto",
}
