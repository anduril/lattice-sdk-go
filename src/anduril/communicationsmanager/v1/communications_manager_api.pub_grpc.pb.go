// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: anduril/communicationsmanager/v1/communications_manager_api.pub.proto

package communicationsmanagerv1

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
	CommunicationsManagerAPI_PutRule_FullMethodName                 = "/anduril.communicationsmanager.v1.CommunicationsManagerAPI/PutRule"
	CommunicationsManagerAPI_StreamRules_FullMethodName             = "/anduril.communicationsmanager.v1.CommunicationsManagerAPI/StreamRules"
	CommunicationsManagerAPI_DeleteRule_FullMethodName              = "/anduril.communicationsmanager.v1.CommunicationsManagerAPI/DeleteRule"
	CommunicationsManagerAPI_StreamIntegrations_FullMethodName      = "/anduril.communicationsmanager.v1.CommunicationsManagerAPI/StreamIntegrations"
	CommunicationsManagerAPI_GetIntegrations_FullMethodName         = "/anduril.communicationsmanager.v1.CommunicationsManagerAPI/GetIntegrations"
	CommunicationsManagerAPI_IntegrationHealthUpdate_FullMethodName = "/anduril.communicationsmanager.v1.CommunicationsManagerAPI/IntegrationHealthUpdate"
)

// CommunicationsManagerAPIClient is the client API for CommunicationsManagerAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Manages the policies for distribution of data via application of filters & priorities.
type CommunicationsManagerAPIClient interface {
	// Puts a rule. Rules are distributed to all nodes in the mesh.
	PutRule(ctx context.Context, in *PutRuleRequest, opts ...grpc.CallOption) (*PutRuleResponse, error)
	// Streams all rules available to this distribution manager.
	StreamRules(ctx context.Context, in *StreamRulesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamRulesResponse], error)
	DeleteRule(ctx context.Context, in *DeleteRuleRequest, opts ...grpc.CallOption) (*DeleteRuleResponse, error)
	// Stream a list of integrations registered with the system. Supports various
	// filters to constrain to specific nodes.
	StreamIntegrations(ctx context.Context, in *StreamIntegrationsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamIntegrationsResponse], error)
	// Deprecated: Use ListIntegrations instead.
	GetIntegrations(ctx context.Context, in *GetIntegrationsRequest, opts ...grpc.CallOption) (*GetIntegrationsResponse, error)
	IntegrationHealthUpdate(ctx context.Context, in *IntegrationHealthUpdateRequest, opts ...grpc.CallOption) (*IntegrationHealthUpdateResponse, error)
}

type communicationsManagerAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunicationsManagerAPIClient(cc grpc.ClientConnInterface) CommunicationsManagerAPIClient {
	return &communicationsManagerAPIClient{cc}
}

func (c *communicationsManagerAPIClient) PutRule(ctx context.Context, in *PutRuleRequest, opts ...grpc.CallOption) (*PutRuleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PutRuleResponse)
	err := c.cc.Invoke(ctx, CommunicationsManagerAPI_PutRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationsManagerAPIClient) StreamRules(ctx context.Context, in *StreamRulesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamRulesResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CommunicationsManagerAPI_ServiceDesc.Streams[0], CommunicationsManagerAPI_StreamRules_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamRulesRequest, StreamRulesResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CommunicationsManagerAPI_StreamRulesClient = grpc.ServerStreamingClient[StreamRulesResponse]

func (c *communicationsManagerAPIClient) DeleteRule(ctx context.Context, in *DeleteRuleRequest, opts ...grpc.CallOption) (*DeleteRuleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRuleResponse)
	err := c.cc.Invoke(ctx, CommunicationsManagerAPI_DeleteRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationsManagerAPIClient) StreamIntegrations(ctx context.Context, in *StreamIntegrationsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamIntegrationsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CommunicationsManagerAPI_ServiceDesc.Streams[1], CommunicationsManagerAPI_StreamIntegrations_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamIntegrationsRequest, StreamIntegrationsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CommunicationsManagerAPI_StreamIntegrationsClient = grpc.ServerStreamingClient[StreamIntegrationsResponse]

func (c *communicationsManagerAPIClient) GetIntegrations(ctx context.Context, in *GetIntegrationsRequest, opts ...grpc.CallOption) (*GetIntegrationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIntegrationsResponse)
	err := c.cc.Invoke(ctx, CommunicationsManagerAPI_GetIntegrations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationsManagerAPIClient) IntegrationHealthUpdate(ctx context.Context, in *IntegrationHealthUpdateRequest, opts ...grpc.CallOption) (*IntegrationHealthUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IntegrationHealthUpdateResponse)
	err := c.cc.Invoke(ctx, CommunicationsManagerAPI_IntegrationHealthUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunicationsManagerAPIServer is the server API for CommunicationsManagerAPI service.
// All implementations must embed UnimplementedCommunicationsManagerAPIServer
// for forward compatibility.
//
// Manages the policies for distribution of data via application of filters & priorities.
type CommunicationsManagerAPIServer interface {
	// Puts a rule. Rules are distributed to all nodes in the mesh.
	PutRule(context.Context, *PutRuleRequest) (*PutRuleResponse, error)
	// Streams all rules available to this distribution manager.
	StreamRules(*StreamRulesRequest, grpc.ServerStreamingServer[StreamRulesResponse]) error
	DeleteRule(context.Context, *DeleteRuleRequest) (*DeleteRuleResponse, error)
	// Stream a list of integrations registered with the system. Supports various
	// filters to constrain to specific nodes.
	StreamIntegrations(*StreamIntegrationsRequest, grpc.ServerStreamingServer[StreamIntegrationsResponse]) error
	// Deprecated: Use ListIntegrations instead.
	GetIntegrations(context.Context, *GetIntegrationsRequest) (*GetIntegrationsResponse, error)
	IntegrationHealthUpdate(context.Context, *IntegrationHealthUpdateRequest) (*IntegrationHealthUpdateResponse, error)
	mustEmbedUnimplementedCommunicationsManagerAPIServer()
}

// UnimplementedCommunicationsManagerAPIServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommunicationsManagerAPIServer struct{}

func (UnimplementedCommunicationsManagerAPIServer) PutRule(context.Context, *PutRuleRequest) (*PutRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutRule not implemented")
}
func (UnimplementedCommunicationsManagerAPIServer) StreamRules(*StreamRulesRequest, grpc.ServerStreamingServer[StreamRulesResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamRules not implemented")
}
func (UnimplementedCommunicationsManagerAPIServer) DeleteRule(context.Context, *DeleteRuleRequest) (*DeleteRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRule not implemented")
}
func (UnimplementedCommunicationsManagerAPIServer) StreamIntegrations(*StreamIntegrationsRequest, grpc.ServerStreamingServer[StreamIntegrationsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamIntegrations not implemented")
}
func (UnimplementedCommunicationsManagerAPIServer) GetIntegrations(context.Context, *GetIntegrationsRequest) (*GetIntegrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIntegrations not implemented")
}
func (UnimplementedCommunicationsManagerAPIServer) IntegrationHealthUpdate(context.Context, *IntegrationHealthUpdateRequest) (*IntegrationHealthUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IntegrationHealthUpdate not implemented")
}
func (UnimplementedCommunicationsManagerAPIServer) mustEmbedUnimplementedCommunicationsManagerAPIServer() {
}
func (UnimplementedCommunicationsManagerAPIServer) testEmbeddedByValue() {}

// UnsafeCommunicationsManagerAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommunicationsManagerAPIServer will
// result in compilation errors.
type UnsafeCommunicationsManagerAPIServer interface {
	mustEmbedUnimplementedCommunicationsManagerAPIServer()
}

func RegisterCommunicationsManagerAPIServer(s grpc.ServiceRegistrar, srv CommunicationsManagerAPIServer) {
	// If the following call pancis, it indicates UnimplementedCommunicationsManagerAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CommunicationsManagerAPI_ServiceDesc, srv)
}

func _CommunicationsManagerAPI_PutRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationsManagerAPIServer).PutRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunicationsManagerAPI_PutRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationsManagerAPIServer).PutRule(ctx, req.(*PutRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationsManagerAPI_StreamRules_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRulesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommunicationsManagerAPIServer).StreamRules(m, &grpc.GenericServerStream[StreamRulesRequest, StreamRulesResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CommunicationsManagerAPI_StreamRulesServer = grpc.ServerStreamingServer[StreamRulesResponse]

func _CommunicationsManagerAPI_DeleteRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationsManagerAPIServer).DeleteRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunicationsManagerAPI_DeleteRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationsManagerAPIServer).DeleteRule(ctx, req.(*DeleteRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationsManagerAPI_StreamIntegrations_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamIntegrationsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommunicationsManagerAPIServer).StreamIntegrations(m, &grpc.GenericServerStream[StreamIntegrationsRequest, StreamIntegrationsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CommunicationsManagerAPI_StreamIntegrationsServer = grpc.ServerStreamingServer[StreamIntegrationsResponse]

func _CommunicationsManagerAPI_GetIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIntegrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationsManagerAPIServer).GetIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunicationsManagerAPI_GetIntegrations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationsManagerAPIServer).GetIntegrations(ctx, req.(*GetIntegrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationsManagerAPI_IntegrationHealthUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntegrationHealthUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationsManagerAPIServer).IntegrationHealthUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunicationsManagerAPI_IntegrationHealthUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationsManagerAPIServer).IntegrationHealthUpdate(ctx, req.(*IntegrationHealthUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommunicationsManagerAPI_ServiceDesc is the grpc.ServiceDesc for CommunicationsManagerAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommunicationsManagerAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "anduril.communicationsmanager.v1.CommunicationsManagerAPI",
	HandlerType: (*CommunicationsManagerAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutRule",
			Handler:    _CommunicationsManagerAPI_PutRule_Handler,
		},
		{
			MethodName: "DeleteRule",
			Handler:    _CommunicationsManagerAPI_DeleteRule_Handler,
		},
		{
			MethodName: "GetIntegrations",
			Handler:    _CommunicationsManagerAPI_GetIntegrations_Handler,
		},
		{
			MethodName: "IntegrationHealthUpdate",
			Handler:    _CommunicationsManagerAPI_IntegrationHealthUpdate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRules",
			Handler:       _CommunicationsManagerAPI_StreamRules_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamIntegrations",
			Handler:       _CommunicationsManagerAPI_StreamIntegrations_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "anduril/communicationsmanager/v1/communications_manager_api.pub.proto",
}
