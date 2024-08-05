// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: anduril/auth/v2/idp.pub.proto

package authv2

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
	Idps_GetSSOURL_FullMethodName = "/anduril.auth.v2.Idps/GetSSOURL"
)

// IdpsClient is the client API for Idps service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdpsClient interface {
	// GetSSOURL returns an identity provider's single sign on (SSO) URL which a user should be redirected to
	// authenticate. If no IDP can be located for the email's domain then a NotFound error is returned.
	GetSSOURL(ctx context.Context, in *GetSSOURLRequest, opts ...grpc.CallOption) (*GetSSOURLResponse, error)
}

type idpsClient struct {
	cc grpc.ClientConnInterface
}

func NewIdpsClient(cc grpc.ClientConnInterface) IdpsClient {
	return &idpsClient{cc}
}

func (c *idpsClient) GetSSOURL(ctx context.Context, in *GetSSOURLRequest, opts ...grpc.CallOption) (*GetSSOURLResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSSOURLResponse)
	err := c.cc.Invoke(ctx, Idps_GetSSOURL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdpsServer is the server API for Idps service.
// All implementations must embed UnimplementedIdpsServer
// for forward compatibility.
type IdpsServer interface {
	// GetSSOURL returns an identity provider's single sign on (SSO) URL which a user should be redirected to
	// authenticate. If no IDP can be located for the email's domain then a NotFound error is returned.
	GetSSOURL(context.Context, *GetSSOURLRequest) (*GetSSOURLResponse, error)
	mustEmbedUnimplementedIdpsServer()
}

// UnimplementedIdpsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedIdpsServer struct{}

func (UnimplementedIdpsServer) GetSSOURL(context.Context, *GetSSOURLRequest) (*GetSSOURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSSOURL not implemented")
}
func (UnimplementedIdpsServer) mustEmbedUnimplementedIdpsServer() {}
func (UnimplementedIdpsServer) testEmbeddedByValue()              {}

// UnsafeIdpsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdpsServer will
// result in compilation errors.
type UnsafeIdpsServer interface {
	mustEmbedUnimplementedIdpsServer()
}

func RegisterIdpsServer(s grpc.ServiceRegistrar, srv IdpsServer) {
	// If the following call pancis, it indicates UnimplementedIdpsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Idps_ServiceDesc, srv)
}

func _Idps_GetSSOURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSSOURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdpsServer).GetSSOURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Idps_GetSSOURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdpsServer).GetSSOURL(ctx, req.(*GetSSOURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Idps_ServiceDesc is the grpc.ServiceDesc for Idps service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Idps_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "anduril.auth.v2.Idps",
	HandlerType: (*IdpsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSSOURL",
			Handler:    _Idps_GetSSOURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "anduril/auth/v2/idp.pub.proto",
}
