// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: binance/spot/margin/isolated/tradings/fishers/grids/grids.proto

package grids

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

// GridsClient is the client API for Grids service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GridsClient interface {
	Pagenate(ctx context.Context, in *PagenateRequest, opts ...grpc.CallOption) (*PagenateReply, error)
}

type gridsClient struct {
	cc grpc.ClientConnInterface
}

func NewGridsClient(cc grpc.ClientConnInterface) GridsClient {
	return &gridsClient{cc}
}

func (c *gridsClient) Pagenate(ctx context.Context, in *PagenateRequest, opts ...grpc.CallOption) (*PagenateReply, error) {
	out := new(PagenateReply)
	err := c.cc.Invoke(ctx, "/taoniu.local.cryptos.grpc.binance.spot.margin.isolated.tradings.fishers.grids.Grids/Pagenate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GridsServer is the server API for Grids service.
// All implementations must embed UnimplementedGridsServer
// for forward compatibility
type GridsServer interface {
	Pagenate(context.Context, *PagenateRequest) (*PagenateReply, error)
	mustEmbedUnimplementedGridsServer()
}

// UnimplementedGridsServer must be embedded to have forward compatible implementations.
type UnimplementedGridsServer struct {
}

func (UnimplementedGridsServer) Pagenate(context.Context, *PagenateRequest) (*PagenateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pagenate not implemented")
}
func (UnimplementedGridsServer) mustEmbedUnimplementedGridsServer() {}

// UnsafeGridsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GridsServer will
// result in compilation errors.
type UnsafeGridsServer interface {
	mustEmbedUnimplementedGridsServer()
}

func RegisterGridsServer(s grpc.ServiceRegistrar, srv GridsServer) {
	s.RegisterService(&Grids_ServiceDesc, srv)
}

func _Grids_Pagenate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PagenateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GridsServer).Pagenate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taoniu.local.cryptos.grpc.binance.spot.margin.isolated.tradings.fishers.grids.Grids/Pagenate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GridsServer).Pagenate(ctx, req.(*PagenateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Grids_ServiceDesc is the grpc.ServiceDesc for Grids service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Grids_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "taoniu.local.cryptos.grpc.binance.spot.margin.isolated.tradings.fishers.grids.Grids",
	HandlerType: (*GridsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pagenate",
			Handler:    _Grids_Pagenate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "binance/spot/margin/isolated/tradings/fishers/grids/grids.proto",
}
