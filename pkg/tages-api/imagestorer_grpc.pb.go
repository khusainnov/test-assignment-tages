// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: imagestorer.proto

package tapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ImageServiceClient is the client API for ImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageServiceClient interface {
	UploadImage(ctx context.Context, opts ...grpc.CallOption) (ImageService_UploadImageClient, error)
	ListImages(ctx context.Context, opts ...grpc.CallOption) (ImageService_ListImagesClient, error)
}

type imageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageServiceClient(cc grpc.ClientConnInterface) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) UploadImage(ctx context.Context, opts ...grpc.CallOption) (ImageService_UploadImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ImageService_ServiceDesc.Streams[0], "/ImageService/UploadImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageServiceUploadImageClient{stream}
	return x, nil
}

type ImageService_UploadImageClient interface {
	Send(*UploadImageRequest) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type imageServiceUploadImageClient struct {
	grpc.ClientStream
}

func (x *imageServiceUploadImageClient) Send(m *UploadImageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imageServiceUploadImageClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imageServiceClient) ListImages(ctx context.Context, opts ...grpc.CallOption) (ImageService_ListImagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ImageService_ServiceDesc.Streams[1], "/ImageService/ListImages", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageServiceListImagesClient{stream}
	return x, nil
}

type ImageService_ListImagesClient interface {
	Send(*emptypb.Empty) error
	CloseAndRecv() (*ListImagesResponse, error)
	grpc.ClientStream
}

type imageServiceListImagesClient struct {
	grpc.ClientStream
}

func (x *imageServiceListImagesClient) Send(m *emptypb.Empty) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imageServiceListImagesClient) CloseAndRecv() (*ListImagesResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ListImagesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImageServiceServer is the server API for ImageService service.
// All implementations must embed UnimplementedImageServiceServer
// for forward compatibility
type ImageServiceServer interface {
	UploadImage(ImageService_UploadImageServer) error
	ListImages(ImageService_ListImagesServer) error
	mustEmbedUnimplementedImageServiceServer()
}

// UnimplementedImageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImageServiceServer struct {
}

func (UnimplementedImageServiceServer) UploadImage(ImageService_UploadImageServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (UnimplementedImageServiceServer) ListImages(ImageService_ListImagesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListImages not implemented")
}
func (UnimplementedImageServiceServer) mustEmbedUnimplementedImageServiceServer() {}

// UnsafeImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServiceServer will
// result in compilation errors.
type UnsafeImageServiceServer interface {
	mustEmbedUnimplementedImageServiceServer()
}

func RegisterImageServiceServer(s grpc.ServiceRegistrar, srv ImageServiceServer) {
	s.RegisterService(&ImageService_ServiceDesc, srv)
}

func _ImageService_UploadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImageServiceServer).UploadImage(&imageServiceUploadImageServer{stream})
}

type ImageService_UploadImageServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*UploadImageRequest, error)
	grpc.ServerStream
}

type imageServiceUploadImageServer struct {
	grpc.ServerStream
}

func (x *imageServiceUploadImageServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imageServiceUploadImageServer) Recv() (*UploadImageRequest, error) {
	m := new(UploadImageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ImageService_ListImages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImageServiceServer).ListImages(&imageServiceListImagesServer{stream})
}

type ImageService_ListImagesServer interface {
	SendAndClose(*ListImagesResponse) error
	Recv() (*emptypb.Empty, error)
	grpc.ServerStream
}

type imageServiceListImagesServer struct {
	grpc.ServerStream
}

func (x *imageServiceListImagesServer) SendAndClose(m *ListImagesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imageServiceListImagesServer) Recv() (*emptypb.Empty, error) {
	m := new(emptypb.Empty)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImageService_ServiceDesc is the grpc.ServiceDesc for ImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadImage",
			Handler:       _ImageService_UploadImage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ListImages",
			Handler:       _ImageService_ListImages_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "imagestorer.proto",
}
