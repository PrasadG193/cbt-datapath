// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: cbt.proto

package grpc

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

// VolumeSnapshotDeltaServiceClient is the client API for VolumeSnapshotDeltaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VolumeSnapshotDeltaServiceClient interface {
	ListVolumeSnapshotDeltas(ctx context.Context, in *VolumeSnapshotDeltaRequest, opts ...grpc.CallOption) (*VolumeSnapshotDeltaResponse, error)
}

type volumeSnapshotDeltaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVolumeSnapshotDeltaServiceClient(cc grpc.ClientConnInterface) VolumeSnapshotDeltaServiceClient {
	return &volumeSnapshotDeltaServiceClient{cc}
}

func (c *volumeSnapshotDeltaServiceClient) ListVolumeSnapshotDeltas(ctx context.Context, in *VolumeSnapshotDeltaRequest, opts ...grpc.CallOption) (*VolumeSnapshotDeltaResponse, error) {
	out := new(VolumeSnapshotDeltaResponse)
	err := c.cc.Invoke(ctx, "/cbt.VolumeSnapshotDeltaService/ListVolumeSnapshotDeltas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VolumeSnapshotDeltaServiceServer is the server API for VolumeSnapshotDeltaService service.
// All implementations must embed UnimplementedVolumeSnapshotDeltaServiceServer
// for forward compatibility
type VolumeSnapshotDeltaServiceServer interface {
	ListVolumeSnapshotDeltas(context.Context, *VolumeSnapshotDeltaRequest) (*VolumeSnapshotDeltaResponse, error)
	mustEmbedUnimplementedVolumeSnapshotDeltaServiceServer()
}

// UnimplementedVolumeSnapshotDeltaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVolumeSnapshotDeltaServiceServer struct {
}

func (UnimplementedVolumeSnapshotDeltaServiceServer) ListVolumeSnapshotDeltas(context.Context, *VolumeSnapshotDeltaRequest) (*VolumeSnapshotDeltaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVolumeSnapshotDeltas not implemented")
}
func (UnimplementedVolumeSnapshotDeltaServiceServer) mustEmbedUnimplementedVolumeSnapshotDeltaServiceServer() {
}

// UnsafeVolumeSnapshotDeltaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VolumeSnapshotDeltaServiceServer will
// result in compilation errors.
type UnsafeVolumeSnapshotDeltaServiceServer interface {
	mustEmbedUnimplementedVolumeSnapshotDeltaServiceServer()
}

func RegisterVolumeSnapshotDeltaServiceServer(s grpc.ServiceRegistrar, srv VolumeSnapshotDeltaServiceServer) {
	s.RegisterService(&VolumeSnapshotDeltaService_ServiceDesc, srv)
}

func _VolumeSnapshotDeltaService_ListVolumeSnapshotDeltas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeSnapshotDeltaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeSnapshotDeltaServiceServer).ListVolumeSnapshotDeltas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbt.VolumeSnapshotDeltaService/ListVolumeSnapshotDeltas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeSnapshotDeltaServiceServer).ListVolumeSnapshotDeltas(ctx, req.(*VolumeSnapshotDeltaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VolumeSnapshotDeltaService_ServiceDesc is the grpc.ServiceDesc for VolumeSnapshotDeltaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VolumeSnapshotDeltaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cbt.VolumeSnapshotDeltaService",
	HandlerType: (*VolumeSnapshotDeltaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListVolumeSnapshotDeltas",
			Handler:    _VolumeSnapshotDeltaService_ListVolumeSnapshotDeltas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cbt.proto",
}
