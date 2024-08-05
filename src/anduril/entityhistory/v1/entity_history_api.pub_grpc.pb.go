// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: anduril/entityhistory/v1/entity_history_api.pub.proto

package entityhistoryv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	EntityHistoryAPI_ListHistory_FullMethodName = "/anduril.entityhistory.v1.EntityHistoryAPI/ListHistory"
)

// EntityHistoryAPIClient is the client API for EntityHistoryAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The EntityHistoryAPI provides UI-centric data models for understanding
// the historic view of a battle space. The API draws heavily on the Entity
// data model and Entity filter to form the core of its client contract.
type EntityHistoryAPIClient interface {
	// Returns a page of history given a history query.
	ListHistory(ctx context.Context, in *ListHistoryRequest, opts ...grpc.CallOption) (*ListHistoryResponse, error)
}

type entityHistoryAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewEntityHistoryAPIClient(cc grpc.ClientConnInterface) EntityHistoryAPIClient {
	return &entityHistoryAPIClient{cc}
}

func (c *entityHistoryAPIClient) ListHistory(ctx context.Context, in *ListHistoryRequest, opts ...grpc.CallOption) (*ListHistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListHistoryResponse)
	err := c.cc.Invoke(ctx, EntityHistoryAPI_ListHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntityHistoryAPIServer is the server API for EntityHistoryAPI service.
// All implementations must embed UnimplementedEntityHistoryAPIServer
// for forward compatibility.
//
// The EntityHistoryAPI provides UI-centric data models for understanding
// the historic view of a battle space. The API draws heavily on the Entity
// data model and Entity filter to form the core of its client contract.
type EntityHistoryAPIServer interface {
	// Returns a page of history given a history query.
	ListHistory(context.Context, *ListHistoryRequest) (*ListHistoryResponse, error)
	mustEmbedUnimplementedEntityHistoryAPIServer()
}

// UnimplementedEntityHistoryAPIServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEntityHistoryAPIServer struct{}

func (UnimplementedEntityHistoryAPIServer) ListHistory(context.Context, *ListHistoryRequest) (*ListHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHistory not implemented")
}
func (UnimplementedEntityHistoryAPIServer) mustEmbedUnimplementedEntityHistoryAPIServer() {}
func (UnimplementedEntityHistoryAPIServer) testEmbeddedByValue()                          {}

// UnsafeEntityHistoryAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntityHistoryAPIServer will
// result in compilation errors.
type UnsafeEntityHistoryAPIServer interface {
	mustEmbedUnimplementedEntityHistoryAPIServer()
}

func RegisterEntityHistoryAPIServer(s grpc.ServiceRegistrar, srv EntityHistoryAPIServer) {
	// If the following call pancis, it indicates UnimplementedEntityHistoryAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EntityHistoryAPI_ServiceDesc, srv)
}

func _EntityHistoryAPI_ListHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityHistoryAPIServer).ListHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntityHistoryAPI_ListHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityHistoryAPIServer).ListHistory(ctx, req.(*ListHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EntityHistoryAPI_ServiceDesc is the grpc.ServiceDesc for EntityHistoryAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntityHistoryAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "anduril.entityhistory.v1.EntityHistoryAPI",
	HandlerType: (*EntityHistoryAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListHistory",
			Handler:    _EntityHistoryAPI_ListHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "anduril/entityhistory/v1/entity_history_api.pub.proto",
}