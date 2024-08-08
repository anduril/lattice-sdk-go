// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: anduril/taskmanager/v1/task_manager_api.pub.proto

package taskmanagerv1

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
	TaskManagerAPI_CreateTask_FullMethodName    = "/anduril.taskmanager.v1.TaskManagerAPI/CreateTask"
	TaskManagerAPI_GetTask_FullMethodName       = "/anduril.taskmanager.v1.TaskManagerAPI/GetTask"
	TaskManagerAPI_UpdateTask_FullMethodName    = "/anduril.taskmanager.v1.TaskManagerAPI/UpdateTask"
	TaskManagerAPI_UpdateStatus_FullMethodName  = "/anduril.taskmanager.v1.TaskManagerAPI/UpdateStatus"
	TaskManagerAPI_StreamTasks_FullMethodName   = "/anduril.taskmanager.v1.TaskManagerAPI/StreamTasks"
	TaskManagerAPI_ListenAsAgent_FullMethodName = "/anduril.taskmanager.v1.TaskManagerAPI/ListenAsAgent"
)

// TaskManagerAPIClient is the client API for TaskManagerAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Task Manager is a service that performs state management associated with Tasks, and also the execution of Tasks
//
//	on their designated agents.
type TaskManagerAPIClient interface {
	// Create a new Task.
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	// Get an existing Task.
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error)
	// Update definition of a Task, only works on Tasks that are not DONE or CANCEL_REQUESTED. Notes:
	//   - send the current task_version in Task, API will increment definition_version, and reset status_version to 1.
	//   - previous definition_version will have status set to REPLACED.
	//   - depending on assignee, replacing the definition will either update if capable on backend,
	//     or cancel previous and issue new.
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error)
	// Update the status of a Task.
	UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*UpdateStatusResponse, error)
	// Stream all existing live (not yet done) Tasks and any new updates.
	// Intended for clients to gain visibility into real time updates for live Tasks.
	StreamTasks(ctx context.Context, in *StreamTasksRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamTasksResponse], error)
	// Stream Tasks ready for RPC Agent execution that match agent selector (ex: entity_ids).
	// Intended for use by Taskable Agents that need to receive Tasks ready for execution by RPC (no Flux access)
	ListenAsAgent(ctx context.Context, in *ListenAsAgentRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ListenAsAgentResponse], error)
}

type taskManagerAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskManagerAPIClient(cc grpc.ClientConnInterface) TaskManagerAPIClient {
	return &taskManagerAPIClient{cc}
}

func (c *taskManagerAPIClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, TaskManagerAPI_CreateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerAPIClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTaskResponse)
	err := c.cc.Invoke(ctx, TaskManagerAPI_GetTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerAPIClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTaskResponse)
	err := c.cc.Invoke(ctx, TaskManagerAPI_UpdateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerAPIClient) UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*UpdateStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateStatusResponse)
	err := c.cc.Invoke(ctx, TaskManagerAPI_UpdateStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerAPIClient) StreamTasks(ctx context.Context, in *StreamTasksRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamTasksResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TaskManagerAPI_ServiceDesc.Streams[0], TaskManagerAPI_StreamTasks_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamTasksRequest, StreamTasksResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TaskManagerAPI_StreamTasksClient = grpc.ServerStreamingClient[StreamTasksResponse]

func (c *taskManagerAPIClient) ListenAsAgent(ctx context.Context, in *ListenAsAgentRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ListenAsAgentResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TaskManagerAPI_ServiceDesc.Streams[1], TaskManagerAPI_ListenAsAgent_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ListenAsAgentRequest, ListenAsAgentResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TaskManagerAPI_ListenAsAgentClient = grpc.ServerStreamingClient[ListenAsAgentResponse]

// TaskManagerAPIServer is the server API for TaskManagerAPI service.
// All implementations must embed UnimplementedTaskManagerAPIServer
// for forward compatibility.
//
// Task Manager is a service that performs state management associated with Tasks, and also the execution of Tasks
//
//	on their designated agents.
type TaskManagerAPIServer interface {
	// Create a new Task.
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	// Get an existing Task.
	GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error)
	// Update definition of a Task, only works on Tasks that are not DONE or CANCEL_REQUESTED. Notes:
	//   - send the current task_version in Task, API will increment definition_version, and reset status_version to 1.
	//   - previous definition_version will have status set to REPLACED.
	//   - depending on assignee, replacing the definition will either update if capable on backend,
	//     or cancel previous and issue new.
	UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error)
	// Update the status of a Task.
	UpdateStatus(context.Context, *UpdateStatusRequest) (*UpdateStatusResponse, error)
	// Stream all existing live (not yet done) Tasks and any new updates.
	// Intended for clients to gain visibility into real time updates for live Tasks.
	StreamTasks(*StreamTasksRequest, grpc.ServerStreamingServer[StreamTasksResponse]) error
	// Stream Tasks ready for RPC Agent execution that match agent selector (ex: entity_ids).
	// Intended for use by Taskable Agents that need to receive Tasks ready for execution by RPC (no Flux access)
	ListenAsAgent(*ListenAsAgentRequest, grpc.ServerStreamingServer[ListenAsAgentResponse]) error
	mustEmbedUnimplementedTaskManagerAPIServer()
}

// UnimplementedTaskManagerAPIServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTaskManagerAPIServer struct{}

func (UnimplementedTaskManagerAPIServer) CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTaskManagerAPIServer) GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedTaskManagerAPIServer) UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedTaskManagerAPIServer) UpdateStatus(context.Context, *UpdateStatusRequest) (*UpdateStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (UnimplementedTaskManagerAPIServer) StreamTasks(*StreamTasksRequest, grpc.ServerStreamingServer[StreamTasksResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamTasks not implemented")
}
func (UnimplementedTaskManagerAPIServer) ListenAsAgent(*ListenAsAgentRequest, grpc.ServerStreamingServer[ListenAsAgentResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ListenAsAgent not implemented")
}
func (UnimplementedTaskManagerAPIServer) mustEmbedUnimplementedTaskManagerAPIServer() {}
func (UnimplementedTaskManagerAPIServer) testEmbeddedByValue()                        {}

// UnsafeTaskManagerAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskManagerAPIServer will
// result in compilation errors.
type UnsafeTaskManagerAPIServer interface {
	mustEmbedUnimplementedTaskManagerAPIServer()
}

func RegisterTaskManagerAPIServer(s grpc.ServiceRegistrar, srv TaskManagerAPIServer) {
	// If the following call pancis, it indicates UnimplementedTaskManagerAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TaskManagerAPI_ServiceDesc, srv)
}

func _TaskManagerAPI_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerAPIServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskManagerAPI_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerAPIServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManagerAPI_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerAPIServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskManagerAPI_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerAPIServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManagerAPI_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerAPIServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskManagerAPI_UpdateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerAPIServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManagerAPI_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerAPIServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskManagerAPI_UpdateStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerAPIServer).UpdateStatus(ctx, req.(*UpdateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManagerAPI_StreamTasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamTasksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskManagerAPIServer).StreamTasks(m, &grpc.GenericServerStream[StreamTasksRequest, StreamTasksResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TaskManagerAPI_StreamTasksServer = grpc.ServerStreamingServer[StreamTasksResponse]

func _TaskManagerAPI_ListenAsAgent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenAsAgentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskManagerAPIServer).ListenAsAgent(m, &grpc.GenericServerStream[ListenAsAgentRequest, ListenAsAgentResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TaskManagerAPI_ListenAsAgentServer = grpc.ServerStreamingServer[ListenAsAgentResponse]

// TaskManagerAPI_ServiceDesc is the grpc.ServiceDesc for TaskManagerAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskManagerAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "anduril.taskmanager.v1.TaskManagerAPI",
	HandlerType: (*TaskManagerAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskManagerAPI_CreateTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _TaskManagerAPI_GetTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TaskManagerAPI_UpdateTask_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _TaskManagerAPI_UpdateStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTasks",
			Handler:       _TaskManagerAPI_StreamTasks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListenAsAgent",
			Handler:       _TaskManagerAPI_ListenAsAgent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "anduril/taskmanager/v1/task_manager_api.pub.proto",
}
