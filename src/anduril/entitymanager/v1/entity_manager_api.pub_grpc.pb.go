// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: anduril/entitymanager/v1/entity_manager_api.pub.proto

package entitymanagerv1

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
	EntityManagerAPI_GetEntity_FullMethodName              = "/anduril.entitymanager.v1.EntityManagerAPI/GetEntity"
	EntityManagerAPI_StreamEntityComponents_FullMethodName = "/anduril.entitymanager.v1.EntityManagerAPI/StreamEntityComponents"
	EntityManagerAPI_PutEntity_FullMethodName              = "/anduril.entitymanager.v1.EntityManagerAPI/PutEntity"
	EntityManagerAPI_PublishEntities_FullMethodName        = "/anduril.entitymanager.v1.EntityManagerAPI/PublishEntities"
	EntityManagerAPI_OverrideEntity_FullMethodName         = "/anduril.entitymanager.v1.EntityManagerAPI/OverrideEntity"
	EntityManagerAPI_RemoveEntityOverride_FullMethodName   = "/anduril.entitymanager.v1.EntityManagerAPI/RemoveEntityOverride"
	EntityManagerAPI_DeleteEntity_FullMethodName           = "/anduril.entitymanager.v1.EntityManagerAPI/DeleteEntity"
)

// EntityManagerAPIClient is the client API for EntityManagerAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The EntityManager provides a UI centric data model for understanding the entities in a battle space.
//
// Every object in a battle space is represented as an "Entity". Each Entity is essentially an ID, with a life cycle,
// and a collection of data components. Each data component is a separate protobuf message definition.
//
// EntityManager provides a way to query the currently live set of entities within a set of filter constraints,
// as well as a limited set of management APIs to change the grouping or relationships between entities.
type EntityManagerAPIClient interface {
	// Get a entity based on an entityId.
	GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*GetEntityResponse, error)
	// Returns a stream of entity with specified components populated.
	StreamEntityComponents(ctx context.Context, in *StreamEntityComponentsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamEntityComponentsResponse], error)
	// Create or Update an Entity. This should be used by low update rate situations where Entity Manager is
	// the source of truth, rather than an aggregator. The canonical example is a manually created entity.
	// Entities created in this fashion are stored as a Base entity, overrides on top are still possible.
	PutEntity(ctx context.Context, in *PutEntityRequest, opts ...grpc.CallOption) (*PutEntityResponse, error)
	// Create or Update one or more Entities. This should be used during high update rate situations where the originator
	// is both the aggregator and source of truth for the published entities, and the originator does not have
	// the ability to directly publish to Flux.
	PublishEntities(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[PublishEntitiesRequest, PublishEntitiesResponse], error)
	// Override an Entity Component. Only fields marked with overridable can be overwritten.
	OverrideEntity(ctx context.Context, in *OverrideEntityRequest, opts ...grpc.CallOption) (*OverrideEntityResponse, error)
	// Remove an override for an Entity component.
	RemoveEntityOverride(ctx context.Context, in *RemoveEntityOverrideRequest, opts ...grpc.CallOption) (*RemoveEntityOverrideResponse, error)
	// Delete an Entity - only works on entities created by PutEntity.
	DeleteEntity(ctx context.Context, in *DeleteEntityRequest, opts ...grpc.CallOption) (*DeleteEntityResponse, error)
}

type entityManagerAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewEntityManagerAPIClient(cc grpc.ClientConnInterface) EntityManagerAPIClient {
	return &entityManagerAPIClient{cc}
}

func (c *entityManagerAPIClient) GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*GetEntityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEntityResponse)
	err := c.cc.Invoke(ctx, EntityManagerAPI_GetEntity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityManagerAPIClient) StreamEntityComponents(ctx context.Context, in *StreamEntityComponentsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamEntityComponentsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &EntityManagerAPI_ServiceDesc.Streams[0], EntityManagerAPI_StreamEntityComponents_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamEntityComponentsRequest, StreamEntityComponentsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EntityManagerAPI_StreamEntityComponentsClient = grpc.ServerStreamingClient[StreamEntityComponentsResponse]

func (c *entityManagerAPIClient) PutEntity(ctx context.Context, in *PutEntityRequest, opts ...grpc.CallOption) (*PutEntityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PutEntityResponse)
	err := c.cc.Invoke(ctx, EntityManagerAPI_PutEntity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityManagerAPIClient) PublishEntities(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[PublishEntitiesRequest, PublishEntitiesResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &EntityManagerAPI_ServiceDesc.Streams[1], EntityManagerAPI_PublishEntities_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PublishEntitiesRequest, PublishEntitiesResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EntityManagerAPI_PublishEntitiesClient = grpc.ClientStreamingClient[PublishEntitiesRequest, PublishEntitiesResponse]

func (c *entityManagerAPIClient) OverrideEntity(ctx context.Context, in *OverrideEntityRequest, opts ...grpc.CallOption) (*OverrideEntityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OverrideEntityResponse)
	err := c.cc.Invoke(ctx, EntityManagerAPI_OverrideEntity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityManagerAPIClient) RemoveEntityOverride(ctx context.Context, in *RemoveEntityOverrideRequest, opts ...grpc.CallOption) (*RemoveEntityOverrideResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveEntityOverrideResponse)
	err := c.cc.Invoke(ctx, EntityManagerAPI_RemoveEntityOverride_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityManagerAPIClient) DeleteEntity(ctx context.Context, in *DeleteEntityRequest, opts ...grpc.CallOption) (*DeleteEntityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteEntityResponse)
	err := c.cc.Invoke(ctx, EntityManagerAPI_DeleteEntity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntityManagerAPIServer is the server API for EntityManagerAPI service.
// All implementations must embed UnimplementedEntityManagerAPIServer
// for forward compatibility.
//
// The EntityManager provides a UI centric data model for understanding the entities in a battle space.
//
// Every object in a battle space is represented as an "Entity". Each Entity is essentially an ID, with a life cycle,
// and a collection of data components. Each data component is a separate protobuf message definition.
//
// EntityManager provides a way to query the currently live set of entities within a set of filter constraints,
// as well as a limited set of management APIs to change the grouping or relationships between entities.
type EntityManagerAPIServer interface {
	// Get a entity based on an entityId.
	GetEntity(context.Context, *GetEntityRequest) (*GetEntityResponse, error)
	// Returns a stream of entity with specified components populated.
	StreamEntityComponents(*StreamEntityComponentsRequest, grpc.ServerStreamingServer[StreamEntityComponentsResponse]) error
	// Create or Update an Entity. This should be used by low update rate situations where Entity Manager is
	// the source of truth, rather than an aggregator. The canonical example is a manually created entity.
	// Entities created in this fashion are stored as a Base entity, overrides on top are still possible.
	PutEntity(context.Context, *PutEntityRequest) (*PutEntityResponse, error)
	// Create or Update one or more Entities. This should be used during high update rate situations where the originator
	// is both the aggregator and source of truth for the published entities, and the originator does not have
	// the ability to directly publish to Flux.
	PublishEntities(grpc.ClientStreamingServer[PublishEntitiesRequest, PublishEntitiesResponse]) error
	// Override an Entity Component. Only fields marked with overridable can be overwritten.
	OverrideEntity(context.Context, *OverrideEntityRequest) (*OverrideEntityResponse, error)
	// Remove an override for an Entity component.
	RemoveEntityOverride(context.Context, *RemoveEntityOverrideRequest) (*RemoveEntityOverrideResponse, error)
	// Delete an Entity - only works on entities created by PutEntity.
	DeleteEntity(context.Context, *DeleteEntityRequest) (*DeleteEntityResponse, error)
	mustEmbedUnimplementedEntityManagerAPIServer()
}

// UnimplementedEntityManagerAPIServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEntityManagerAPIServer struct{}

func (UnimplementedEntityManagerAPIServer) GetEntity(context.Context, *GetEntityRequest) (*GetEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntity not implemented")
}
func (UnimplementedEntityManagerAPIServer) StreamEntityComponents(*StreamEntityComponentsRequest, grpc.ServerStreamingServer[StreamEntityComponentsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamEntityComponents not implemented")
}
func (UnimplementedEntityManagerAPIServer) PutEntity(context.Context, *PutEntityRequest) (*PutEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutEntity not implemented")
}
func (UnimplementedEntityManagerAPIServer) PublishEntities(grpc.ClientStreamingServer[PublishEntitiesRequest, PublishEntitiesResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PublishEntities not implemented")
}
func (UnimplementedEntityManagerAPIServer) OverrideEntity(context.Context, *OverrideEntityRequest) (*OverrideEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OverrideEntity not implemented")
}
func (UnimplementedEntityManagerAPIServer) RemoveEntityOverride(context.Context, *RemoveEntityOverrideRequest) (*RemoveEntityOverrideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveEntityOverride not implemented")
}
func (UnimplementedEntityManagerAPIServer) DeleteEntity(context.Context, *DeleteEntityRequest) (*DeleteEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEntity not implemented")
}
func (UnimplementedEntityManagerAPIServer) mustEmbedUnimplementedEntityManagerAPIServer() {}
func (UnimplementedEntityManagerAPIServer) testEmbeddedByValue()                          {}

// UnsafeEntityManagerAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntityManagerAPIServer will
// result in compilation errors.
type UnsafeEntityManagerAPIServer interface {
	mustEmbedUnimplementedEntityManagerAPIServer()
}

func RegisterEntityManagerAPIServer(s grpc.ServiceRegistrar, srv EntityManagerAPIServer) {
	// If the following call pancis, it indicates UnimplementedEntityManagerAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EntityManagerAPI_ServiceDesc, srv)
}

func _EntityManagerAPI_GetEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityManagerAPIServer).GetEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntityManagerAPI_GetEntity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityManagerAPIServer).GetEntity(ctx, req.(*GetEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityManagerAPI_StreamEntityComponents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamEntityComponentsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EntityManagerAPIServer).StreamEntityComponents(m, &grpc.GenericServerStream[StreamEntityComponentsRequest, StreamEntityComponentsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EntityManagerAPI_StreamEntityComponentsServer = grpc.ServerStreamingServer[StreamEntityComponentsResponse]

func _EntityManagerAPI_PutEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityManagerAPIServer).PutEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntityManagerAPI_PutEntity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityManagerAPIServer).PutEntity(ctx, req.(*PutEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityManagerAPI_PublishEntities_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EntityManagerAPIServer).PublishEntities(&grpc.GenericServerStream[PublishEntitiesRequest, PublishEntitiesResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EntityManagerAPI_PublishEntitiesServer = grpc.ClientStreamingServer[PublishEntitiesRequest, PublishEntitiesResponse]

func _EntityManagerAPI_OverrideEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OverrideEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityManagerAPIServer).OverrideEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntityManagerAPI_OverrideEntity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityManagerAPIServer).OverrideEntity(ctx, req.(*OverrideEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityManagerAPI_RemoveEntityOverride_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveEntityOverrideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityManagerAPIServer).RemoveEntityOverride(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntityManagerAPI_RemoveEntityOverride_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityManagerAPIServer).RemoveEntityOverride(ctx, req.(*RemoveEntityOverrideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityManagerAPI_DeleteEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityManagerAPIServer).DeleteEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntityManagerAPI_DeleteEntity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityManagerAPIServer).DeleteEntity(ctx, req.(*DeleteEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EntityManagerAPI_ServiceDesc is the grpc.ServiceDesc for EntityManagerAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntityManagerAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "anduril.entitymanager.v1.EntityManagerAPI",
	HandlerType: (*EntityManagerAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntity",
			Handler:    _EntityManagerAPI_GetEntity_Handler,
		},
		{
			MethodName: "PutEntity",
			Handler:    _EntityManagerAPI_PutEntity_Handler,
		},
		{
			MethodName: "OverrideEntity",
			Handler:    _EntityManagerAPI_OverrideEntity_Handler,
		},
		{
			MethodName: "RemoveEntityOverride",
			Handler:    _EntityManagerAPI_RemoveEntityOverride_Handler,
		},
		{
			MethodName: "DeleteEntity",
			Handler:    _EntityManagerAPI_DeleteEntity_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEntityComponents",
			Handler:       _EntityManagerAPI_StreamEntityComponents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PublishEntities",
			Handler:       _EntityManagerAPI_PublishEntities_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "anduril/entitymanager/v1/entity_manager_api.pub.proto",
}
