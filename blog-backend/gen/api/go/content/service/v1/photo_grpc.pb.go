// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: content/service/v1/photo.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pagination "kratos-blog/gen/api/go/pagination"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PhotoServiceClient is the client API for PhotoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PhotoServiceClient interface {
	// 获取照片列表
	ListPhoto(ctx context.Context, in *pagination.PagingRequest, opts ...grpc.CallOption) (*ListPhotoResponse, error)
	// 获取照片数据
	GetPhoto(ctx context.Context, in *GetPhotoRequest, opts ...grpc.CallOption) (*Photo, error)
	// 创建照片
	CreatePhoto(ctx context.Context, in *CreatePhotoRequest, opts ...grpc.CallOption) (*Photo, error)
	// 更新照片
	UpdatePhoto(ctx context.Context, in *UpdatePhotoRequest, opts ...grpc.CallOption) (*Photo, error)
	// 删除照片
	DeletePhoto(ctx context.Context, in *DeletePhotoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type photoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPhotoServiceClient(cc grpc.ClientConnInterface) PhotoServiceClient {
	return &photoServiceClient{cc}
}

func (c *photoServiceClient) ListPhoto(ctx context.Context, in *pagination.PagingRequest, opts ...grpc.CallOption) (*ListPhotoResponse, error) {
	out := new(ListPhotoResponse)
	err := c.cc.Invoke(ctx, "/content.service.v1.PhotoService/ListPhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoServiceClient) GetPhoto(ctx context.Context, in *GetPhotoRequest, opts ...grpc.CallOption) (*Photo, error) {
	out := new(Photo)
	err := c.cc.Invoke(ctx, "/content.service.v1.PhotoService/GetPhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoServiceClient) CreatePhoto(ctx context.Context, in *CreatePhotoRequest, opts ...grpc.CallOption) (*Photo, error) {
	out := new(Photo)
	err := c.cc.Invoke(ctx, "/content.service.v1.PhotoService/CreatePhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoServiceClient) UpdatePhoto(ctx context.Context, in *UpdatePhotoRequest, opts ...grpc.CallOption) (*Photo, error) {
	out := new(Photo)
	err := c.cc.Invoke(ctx, "/content.service.v1.PhotoService/UpdatePhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoServiceClient) DeletePhoto(ctx context.Context, in *DeletePhotoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/content.service.v1.PhotoService/DeletePhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhotoServiceServer is the server API for PhotoService service.
// All implementations must embed UnimplementedPhotoServiceServer
// for forward compatibility
type PhotoServiceServer interface {
	// 获取照片列表
	ListPhoto(context.Context, *pagination.PagingRequest) (*ListPhotoResponse, error)
	// 获取照片数据
	GetPhoto(context.Context, *GetPhotoRequest) (*Photo, error)
	// 创建照片
	CreatePhoto(context.Context, *CreatePhotoRequest) (*Photo, error)
	// 更新照片
	UpdatePhoto(context.Context, *UpdatePhotoRequest) (*Photo, error)
	// 删除照片
	DeletePhoto(context.Context, *DeletePhotoRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPhotoServiceServer()
}

// UnimplementedPhotoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPhotoServiceServer struct {
}

func (UnimplementedPhotoServiceServer) ListPhoto(context.Context, *pagination.PagingRequest) (*ListPhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPhoto not implemented")
}
func (UnimplementedPhotoServiceServer) GetPhoto(context.Context, *GetPhotoRequest) (*Photo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPhoto not implemented")
}
func (UnimplementedPhotoServiceServer) CreatePhoto(context.Context, *CreatePhotoRequest) (*Photo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePhoto not implemented")
}
func (UnimplementedPhotoServiceServer) UpdatePhoto(context.Context, *UpdatePhotoRequest) (*Photo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePhoto not implemented")
}
func (UnimplementedPhotoServiceServer) DeletePhoto(context.Context, *DeletePhotoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePhoto not implemented")
}
func (UnimplementedPhotoServiceServer) mustEmbedUnimplementedPhotoServiceServer() {}

// UnsafePhotoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PhotoServiceServer will
// result in compilation errors.
type UnsafePhotoServiceServer interface {
	mustEmbedUnimplementedPhotoServiceServer()
}

func RegisterPhotoServiceServer(s grpc.ServiceRegistrar, srv PhotoServiceServer) {
	s.RegisterService(&PhotoService_ServiceDesc, srv)
}

func _PhotoService_ListPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pagination.PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServiceServer).ListPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.service.v1.PhotoService/ListPhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServiceServer).ListPhoto(ctx, req.(*pagination.PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoService_GetPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServiceServer).GetPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.service.v1.PhotoService/GetPhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServiceServer).GetPhoto(ctx, req.(*GetPhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoService_CreatePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServiceServer).CreatePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.service.v1.PhotoService/CreatePhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServiceServer).CreatePhoto(ctx, req.(*CreatePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoService_UpdatePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServiceServer).UpdatePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.service.v1.PhotoService/UpdatePhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServiceServer).UpdatePhoto(ctx, req.(*UpdatePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoService_DeletePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServiceServer).DeletePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.service.v1.PhotoService/DeletePhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServiceServer).DeletePhoto(ctx, req.(*DeletePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PhotoService_ServiceDesc is the grpc.ServiceDesc for PhotoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PhotoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content.service.v1.PhotoService",
	HandlerType: (*PhotoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPhoto",
			Handler:    _PhotoService_ListPhoto_Handler,
		},
		{
			MethodName: "GetPhoto",
			Handler:    _PhotoService_GetPhoto_Handler,
		},
		{
			MethodName: "CreatePhoto",
			Handler:    _PhotoService_CreatePhoto_Handler,
		},
		{
			MethodName: "UpdatePhoto",
			Handler:    _PhotoService_UpdatePhoto_Handler,
		},
		{
			MethodName: "DeletePhoto",
			Handler:    _PhotoService_DeletePhoto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content/service/v1/photo.proto",
}