// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: tag.proto

package tag

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

// TagClient is the client API for Tag service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagClient interface {
	AddTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (Tag_AddTagClient, error)
	ListTag(ctx context.Context, in *ListTagRequest, opts ...grpc.CallOption) (Tag_ListTagClient, error)
	DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (Tag_DeleteTagClient, error)
}

type tagClient struct {
	cc grpc.ClientConnInterface
}

func NewTagClient(cc grpc.ClientConnInterface) TagClient {
	return &tagClient{cc}
}

func (c *tagClient) AddTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (Tag_AddTagClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tag_ServiceDesc.Streams[0], "/datasource.Tag/AddTag", opts...)
	if err != nil {
		return nil, err
	}
	x := &tagAddTagClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tag_AddTagClient interface {
	Recv() (*TagMessageResponse, error)
	grpc.ClientStream
}

type tagAddTagClient struct {
	grpc.ClientStream
}

func (x *tagAddTagClient) Recv() (*TagMessageResponse, error) {
	m := new(TagMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tagClient) ListTag(ctx context.Context, in *ListTagRequest, opts ...grpc.CallOption) (Tag_ListTagClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tag_ServiceDesc.Streams[1], "/datasource.Tag/ListTag", opts...)
	if err != nil {
		return nil, err
	}
	x := &tagListTagClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tag_ListTagClient interface {
	Recv() (*ListTagResponse, error)
	grpc.ClientStream
}

type tagListTagClient struct {
	grpc.ClientStream
}

func (x *tagListTagClient) Recv() (*ListTagResponse, error) {
	m := new(ListTagResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tagClient) DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (Tag_DeleteTagClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tag_ServiceDesc.Streams[2], "/datasource.Tag/DeleteTag", opts...)
	if err != nil {
		return nil, err
	}
	x := &tagDeleteTagClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tag_DeleteTagClient interface {
	Recv() (*TagMessageResponse, error)
	grpc.ClientStream
}

type tagDeleteTagClient struct {
	grpc.ClientStream
}

func (x *tagDeleteTagClient) Recv() (*TagMessageResponse, error) {
	m := new(TagMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TagServer is the server API for Tag service.
// All implementations must embed UnimplementedTagServer
// for forward compatibility
type TagServer interface {
	AddTag(*CreateTagRequest, Tag_AddTagServer) error
	ListTag(*ListTagRequest, Tag_ListTagServer) error
	DeleteTag(*DeleteTagRequest, Tag_DeleteTagServer) error
	mustEmbedUnimplementedTagServer()
}

// UnimplementedTagServer must be embedded to have forward compatible implementations.
type UnimplementedTagServer struct {
}

func (UnimplementedTagServer) AddTag(*CreateTagRequest, Tag_AddTagServer) error {
	return status.Errorf(codes.Unimplemented, "method AddTag not implemented")
}
func (UnimplementedTagServer) ListTag(*ListTagRequest, Tag_ListTagServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTag not implemented")
}
func (UnimplementedTagServer) DeleteTag(*DeleteTagRequest, Tag_DeleteTagServer) error {
	return status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}
func (UnimplementedTagServer) mustEmbedUnimplementedTagServer() {}

// UnsafeTagServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagServer will
// result in compilation errors.
type UnsafeTagServer interface {
	mustEmbedUnimplementedTagServer()
}

func RegisterTagServer(s grpc.ServiceRegistrar, srv TagServer) {
	s.RegisterService(&Tag_ServiceDesc, srv)
}

func _Tag_AddTag_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateTagRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TagServer).AddTag(m, &tagAddTagServer{stream})
}

type Tag_AddTagServer interface {
	Send(*TagMessageResponse) error
	grpc.ServerStream
}

type tagAddTagServer struct {
	grpc.ServerStream
}

func (x *tagAddTagServer) Send(m *TagMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Tag_ListTag_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTagRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TagServer).ListTag(m, &tagListTagServer{stream})
}

type Tag_ListTagServer interface {
	Send(*ListTagResponse) error
	grpc.ServerStream
}

type tagListTagServer struct {
	grpc.ServerStream
}

func (x *tagListTagServer) Send(m *ListTagResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Tag_DeleteTag_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DeleteTagRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TagServer).DeleteTag(m, &tagDeleteTagServer{stream})
}

type Tag_DeleteTagServer interface {
	Send(*TagMessageResponse) error
	grpc.ServerStream
}

type tagDeleteTagServer struct {
	grpc.ServerStream
}

func (x *tagDeleteTagServer) Send(m *TagMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Tag_ServiceDesc is the grpc.ServiceDesc for Tag service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tag_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "datasource.Tag",
	HandlerType: (*TagServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddTag",
			Handler:       _Tag_AddTag_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListTag",
			Handler:       _Tag_ListTag_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DeleteTag",
			Handler:       _Tag_DeleteTag_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tag.proto",
}
