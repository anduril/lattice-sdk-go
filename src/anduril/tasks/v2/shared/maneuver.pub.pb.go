// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: anduril/tasks/v2/shared/maneuver.pub.proto

package tasksv2

import (
	v2 "github.com/dogun-anduril/anduril-sdk-go/src/anduril/tasks/v2"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LaunchTrackingMode int32

const (
	LaunchTrackingMode_LAUNCH_TRACKING_MODE_INVALID           LaunchTrackingMode = 0
	LaunchTrackingMode_LAUNCH_TRACKING_MODE_GO_TO_WAYPOINT    LaunchTrackingMode = 1
	LaunchTrackingMode_LAUNCH_TRACKING_MODE_TRACK_TO_WAYPOINT LaunchTrackingMode = 2
)

// Enum value maps for LaunchTrackingMode.
var (
	LaunchTrackingMode_name = map[int32]string{
		0: "LAUNCH_TRACKING_MODE_INVALID",
		1: "LAUNCH_TRACKING_MODE_GO_TO_WAYPOINT",
		2: "LAUNCH_TRACKING_MODE_TRACK_TO_WAYPOINT",
	}
	LaunchTrackingMode_value = map[string]int32{
		"LAUNCH_TRACKING_MODE_INVALID":           0,
		"LAUNCH_TRACKING_MODE_GO_TO_WAYPOINT":    1,
		"LAUNCH_TRACKING_MODE_TRACK_TO_WAYPOINT": 2,
	}
)

func (x LaunchTrackingMode) Enum() *LaunchTrackingMode {
	p := new(LaunchTrackingMode)
	*p = x
	return p
}

func (x LaunchTrackingMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LaunchTrackingMode) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_tasks_v2_shared_maneuver_pub_proto_enumTypes[0].Descriptor()
}

func (LaunchTrackingMode) Type() protoreflect.EnumType {
	return &file_anduril_tasks_v2_shared_maneuver_pub_proto_enumTypes[0]
}

func (x LaunchTrackingMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LaunchTrackingMode.Descriptor instead.
func (LaunchTrackingMode) EnumDescriptor() ([]byte, []int) {
	return file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDescGZIP(), []int{0}
}

// Maps to BREVITY code Marshal.
// Establish(ed) at a specific point, typically used to posture forces in preparation for an offensive operation.
type Marshal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Objective to Marshal to.
	Objective *v2.Objective `protobuf:"bytes,1,opt,name=objective,proto3" json:"objective,omitempty"`
}

func (x *Marshal) Reset() {
	*x = Marshal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Marshal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Marshal) ProtoMessage() {}

func (x *Marshal) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Marshal.ProtoReflect.Descriptor instead.
func (*Marshal) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDescGZIP(), []int{0}
}

func (x *Marshal) GetObjective() *v2.Objective {
	if x != nil {
		return x.Objective
	}
	return nil
}

// Maps to UCI code RoutePlan.
// Used to command a platform between locations by requesting to make this RoutePlan the single primary active route.
type Transit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plan *RoutePlan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
}

func (x *Transit) Reset() {
	*x = Transit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transit) ProtoMessage() {}

func (x *Transit) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transit.ProtoReflect.Descriptor instead.
func (*Transit) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDescGZIP(), []int{1}
}

func (x *Transit) GetPlan() *RoutePlan {
	if x != nil {
		return x.Plan
	}
	return nil
}

type RoutePlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Route *Route `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *RoutePlan) Reset() {
	*x = RoutePlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutePlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutePlan) ProtoMessage() {}

func (x *RoutePlan) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutePlan.ProtoReflect.Descriptor instead.
func (*RoutePlan) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDescGZIP(), []int{2}
}

func (x *RoutePlan) GetRoute() *Route {
	if x != nil {
		return x.Route
	}
	return nil
}

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path []*PathSegment `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDescGZIP(), []int{3}
}

func (x *Route) GetPath() []*PathSegment {
	if x != nil {
		return x.Path
	}
	return nil
}

type PathSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to EndPoint:
	//
	//	*PathSegment_Waypoint
	//	*PathSegment_Loiter
	EndPoint isPathSegment_EndPoint `protobuf_oneof:"end_point"`
}

func (x *PathSegment) Reset() {
	*x = PathSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathSegment) ProtoMessage() {}

func (x *PathSegment) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathSegment.ProtoReflect.Descriptor instead.
func (*PathSegment) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDescGZIP(), []int{4}
}

func (m *PathSegment) GetEndPoint() isPathSegment_EndPoint {
	if m != nil {
		return m.EndPoint
	}
	return nil
}

func (x *PathSegment) GetWaypoint() *Waypoint {
	if x, ok := x.GetEndPoint().(*PathSegment_Waypoint); ok {
		return x.Waypoint
	}
	return nil
}

func (x *PathSegment) GetLoiter() *Loiter {
	if x, ok := x.GetEndPoint().(*PathSegment_Loiter); ok {
		return x.Loiter
	}
	return nil
}

type isPathSegment_EndPoint interface {
	isPathSegment_EndPoint()
}

type PathSegment_Waypoint struct {
	Waypoint *Waypoint `protobuf:"bytes,1,opt,name=waypoint,proto3,oneof"`
}

type PathSegment_Loiter struct {
	Loiter *Loiter `protobuf:"bytes,2,opt,name=loiter,proto3,oneof"`
}

func (*PathSegment_Waypoint) isPathSegment_EndPoint() {}

func (*PathSegment_Loiter) isPathSegment_EndPoint() {}

type Waypoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Point:
	//
	//	*Waypoint_LlaPoint
	Point isWaypoint_Point `protobuf_oneof:"point"`
}

func (x *Waypoint) Reset() {
	*x = Waypoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Waypoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Waypoint) ProtoMessage() {}

func (x *Waypoint) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Waypoint.ProtoReflect.Descriptor instead.
func (*Waypoint) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDescGZIP(), []int{5}
}

func (m *Waypoint) GetPoint() isWaypoint_Point {
	if m != nil {
		return m.Point
	}
	return nil
}

func (x *Waypoint) GetLlaPoint() *v2.Point {
	if x, ok := x.GetPoint().(*Waypoint_LlaPoint); ok {
		return x.LlaPoint
	}
	return nil
}

type isWaypoint_Point interface {
	isWaypoint_Point()
}

type Waypoint_LlaPoint struct {
	LlaPoint *v2.Point `protobuf:"bytes,1,opt,name=lla_point,json=llaPoint,proto3,oneof"`
}

func (*Waypoint_LlaPoint) isWaypoint_Point() {}

type SetLaunchRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plan         *RoutePlan         `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	TrackingMode LaunchTrackingMode `protobuf:"varint,2,opt,name=tracking_mode,json=trackingMode,proto3,enum=anduril.tasks.v2.LaunchTrackingMode" json:"tracking_mode,omitempty"`
}

func (x *SetLaunchRoute) Reset() {
	*x = SetLaunchRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLaunchRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLaunchRoute) ProtoMessage() {}

func (x *SetLaunchRoute) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLaunchRoute.ProtoReflect.Descriptor instead.
func (*SetLaunchRoute) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDescGZIP(), []int{6}
}

func (x *SetLaunchRoute) GetPlan() *RoutePlan {
	if x != nil {
		return x.Plan
	}
	return nil
}

func (x *SetLaunchRoute) GetTrackingMode() LaunchTrackingMode {
	if x != nil {
		return x.TrackingMode
	}
	return LaunchTrackingMode_LAUNCH_TRACKING_MODE_INVALID
}

var File_anduril_tasks_v2_shared_maneuver_pub_proto protoreflect.FileDescriptor

var file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f,
	0x76, 0x32, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x6d, 0x61, 0x6e, 0x65, 0x75, 0x76,
	0x65, 0x72, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x24,
	0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x32,
	0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x69, 0x73,
	0x72, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x07, 0x4d,
	0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x12, 0x39, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x22, 0x3a, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x12, 0x2f, 0x0a, 0x04,
	0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x6e, 0x64,
	0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x22, 0x3a, 0x0a,
	0x09, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x2d, 0x0a, 0x05, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x3a, 0x0a, 0x05, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x88, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x74, 0x68, 0x53, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x57, 0x61, 0x79, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x08, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x32, 0x0a, 0x06, 0x6c, 0x6f, 0x69, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x69, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x6c, 0x6f, 0x69,
	0x74, 0x65, 0x72, 0x42, 0x0b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x22, 0x4b, 0x0a, 0x08, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x09,
	0x6c, 0x6c, 0x61, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x6c, 0x61, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x8c, 0x01,
	0x0a, 0x0e, 0x53, 0x65, 0x74, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x12, 0x2f, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61,
	0x6e, 0x12, 0x49, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72,
	0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x61, 0x75, 0x6e,
	0x63, 0x68, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0c,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x2a, 0x8b, 0x01, 0x0a,
	0x12, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x4c, 0x41, 0x55, 0x4e, 0x43, 0x48, 0x5f, 0x54, 0x52,
	0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x27, 0x0a, 0x23, 0x4c, 0x41, 0x55, 0x4e, 0x43, 0x48, 0x5f,
	0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x47, 0x4f,
	0x5f, 0x54, 0x4f, 0x5f, 0x57, 0x41, 0x59, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x2a,
	0x0a, 0x26, 0x4c, 0x41, 0x55, 0x4e, 0x43, 0x48, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e,
	0x47, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x5f, 0x54, 0x4f, 0x5f,
	0x57, 0x41, 0x59, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x02, 0x42, 0xd7, 0x01, 0x0a, 0x14, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x2e, 0x76, 0x32, 0x42, 0x10, 0x4d, 0x61, 0x6e, 0x65, 0x75, 0x76, 0x65, 0x72, 0x50, 0x75, 0x62,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x75, 0x6e, 0x2d, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f,
	0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x3b, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x41, 0x54, 0x58, 0xaa, 0x02, 0x10, 0x41, 0x6e, 0x64,
	0x75, 0x72, 0x69, 0x6c, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x10,
	0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x5c, 0x56, 0x32,
	0xe2, 0x02, 0x1c, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x12, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x3a, 0x3a, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDescOnce sync.Once
	file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDescData = file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDesc
)

func file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDescGZIP() []byte {
	file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDescOnce.Do(func() {
		file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDescData)
	})
	return file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDescData
}

var file_anduril_tasks_v2_shared_maneuver_pub_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_anduril_tasks_v2_shared_maneuver_pub_proto_goTypes = []any{
	(LaunchTrackingMode)(0), // 0: anduril.tasks.v2.LaunchTrackingMode
	(*Marshal)(nil),         // 1: anduril.tasks.v2.Marshal
	(*Transit)(nil),         // 2: anduril.tasks.v2.Transit
	(*RoutePlan)(nil),       // 3: anduril.tasks.v2.RoutePlan
	(*Route)(nil),           // 4: anduril.tasks.v2.Route
	(*PathSegment)(nil),     // 5: anduril.tasks.v2.PathSegment
	(*Waypoint)(nil),        // 6: anduril.tasks.v2.Waypoint
	(*SetLaunchRoute)(nil),  // 7: anduril.tasks.v2.SetLaunchRoute
	(*v2.Objective)(nil),    // 8: anduril.tasks.v2.Objective
	(*Loiter)(nil),          // 9: anduril.tasks.v2.Loiter
	(*v2.Point)(nil),        // 10: anduril.tasks.v2.Point
}
var file_anduril_tasks_v2_shared_maneuver_pub_proto_depIdxs = []int32{
	8,  // 0: anduril.tasks.v2.Marshal.objective:type_name -> anduril.tasks.v2.Objective
	3,  // 1: anduril.tasks.v2.Transit.plan:type_name -> anduril.tasks.v2.RoutePlan
	4,  // 2: anduril.tasks.v2.RoutePlan.route:type_name -> anduril.tasks.v2.Route
	5,  // 3: anduril.tasks.v2.Route.path:type_name -> anduril.tasks.v2.PathSegment
	6,  // 4: anduril.tasks.v2.PathSegment.waypoint:type_name -> anduril.tasks.v2.Waypoint
	9,  // 5: anduril.tasks.v2.PathSegment.loiter:type_name -> anduril.tasks.v2.Loiter
	10, // 6: anduril.tasks.v2.Waypoint.lla_point:type_name -> anduril.tasks.v2.Point
	3,  // 7: anduril.tasks.v2.SetLaunchRoute.plan:type_name -> anduril.tasks.v2.RoutePlan
	0,  // 8: anduril.tasks.v2.SetLaunchRoute.tracking_mode:type_name -> anduril.tasks.v2.LaunchTrackingMode
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_anduril_tasks_v2_shared_maneuver_pub_proto_init() }
func file_anduril_tasks_v2_shared_maneuver_pub_proto_init() {
	if File_anduril_tasks_v2_shared_maneuver_pub_proto != nil {
		return
	}
	file_anduril_tasks_v2_shared_isr_pub_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Marshal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Transit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RoutePlan); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Route); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PathSegment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Waypoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SetLaunchRoute); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[4].OneofWrappers = []any{
		(*PathSegment_Waypoint)(nil),
		(*PathSegment_Loiter)(nil),
	}
	file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes[5].OneofWrappers = []any{
		(*Waypoint_LlaPoint)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_tasks_v2_shared_maneuver_pub_proto_goTypes,
		DependencyIndexes: file_anduril_tasks_v2_shared_maneuver_pub_proto_depIdxs,
		EnumInfos:         file_anduril_tasks_v2_shared_maneuver_pub_proto_enumTypes,
		MessageInfos:      file_anduril_tasks_v2_shared_maneuver_pub_proto_msgTypes,
	}.Build()
	File_anduril_tasks_v2_shared_maneuver_pub_proto = out.File
	file_anduril_tasks_v2_shared_maneuver_pub_proto_rawDesc = nil
	file_anduril_tasks_v2_shared_maneuver_pub_proto_goTypes = nil
	file_anduril_tasks_v2_shared_maneuver_pub_proto_depIdxs = nil
}
