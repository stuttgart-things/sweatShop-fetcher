// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: repository.proto

package apiclient

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

// FetcherServiceClient is the client API for FetcherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FetcherServiceClient interface {
	// ListFiles returns a list of files in the repo
	ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*FileList, error)
	// Returns a bool val if the repository is valid and has proper access
	TestRepository(ctx context.Context, in *TestRepositoryRequest, opts ...grpc.CallOption) (*TestRepositoryResponse, error)
}

type fetcherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFetcherServiceClient(cc grpc.ClientConnInterface) FetcherServiceClient {
	return &fetcherServiceClient{cc}
}

func (c *fetcherServiceClient) ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*FileList, error) {
	out := new(FileList)
	err := c.cc.Invoke(ctx, "/fetcher.FetcherService/ListFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fetcherServiceClient) TestRepository(ctx context.Context, in *TestRepositoryRequest, opts ...grpc.CallOption) (*TestRepositoryResponse, error) {
	out := new(TestRepositoryResponse)
	err := c.cc.Invoke(ctx, "/fetcher.FetcherService/TestRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FetcherServiceServer is the server API for FetcherService service.
// All implementations should embed UnimplementedFetcherServiceServer
// for forward compatibility
type FetcherServiceServer interface {
	// ListFiles returns a list of files in the repo
	ListFiles(context.Context, *ListFilesRequest) (*FileList, error)
	// Returns a bool val if the repository is valid and has proper access
	TestRepository(context.Context, *TestRepositoryRequest) (*TestRepositoryResponse, error)
}

// UnimplementedFetcherServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFetcherServiceServer struct {
}

func (UnimplementedFetcherServiceServer) ListFiles(context.Context, *ListFilesRequest) (*FileList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedFetcherServiceServer) TestRepository(context.Context, *TestRepositoryRequest) (*TestRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestRepository not implemented")
}

// UnsafeFetcherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FetcherServiceServer will
// result in compilation errors.
type UnsafeFetcherServiceServer interface {
	mustEmbedUnimplementedFetcherServiceServer()
}

func RegisterFetcherServiceServer(s grpc.ServiceRegistrar, srv FetcherServiceServer) {
	s.RegisterService(&FetcherService_ServiceDesc, srv)
}

func _FetcherService_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetcherServiceServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fetcher.FetcherService/ListFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetcherServiceServer).ListFiles(ctx, req.(*ListFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FetcherService_TestRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetcherServiceServer).TestRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fetcher.FetcherService/TestRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetcherServiceServer).TestRepository(ctx, req.(*TestRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FetcherService_ServiceDesc is the grpc.ServiceDesc for FetcherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FetcherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fetcher.FetcherService",
	HandlerType: (*FetcherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFiles",
			Handler:    _FetcherService_ListFiles_Handler,
		},
		{
			MethodName: "TestRepository",
			Handler:    _FetcherService_TestRepository_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "repository.proto",
}
