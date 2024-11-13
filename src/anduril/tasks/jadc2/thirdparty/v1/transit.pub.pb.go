// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: anduril/tasks/jadc2/thirdparty/v1/transit.pub.proto

package thirdpartyv1

import (
	_type "github.com/anduril/anduril-go/src/anduril/type"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Transit represents moving a vehicle on a path through one or more points.
type Transit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path consisting of all segments to be taken for this transit task.
	Path []*PathSegment `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
	// Speed in which the vehicle will move through each of the path segments.
	SurfaceSpeedMs *wrapperspb.DoubleValue `protobuf:"bytes,2,opt,name=surface_speed_ms,json=surfaceSpeedMs,proto3" json:"surface_speed_ms,omitempty"`
}

func (x *Transit) Reset() {
	*x = Transit{}
	mi := &file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Transit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transit) ProtoMessage() {}

func (x *Transit) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_msgTypes[0]
	if x != nil {
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
	return file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_rawDescGZIP(), []int{0}
}

func (x *Transit) GetPath() []*PathSegment {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *Transit) GetSurfaceSpeedMs() *wrapperspb.DoubleValue {
	if x != nil {
		return x.SurfaceSpeedMs
	}
	return nil
}

type PathSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Describes the end of the path segment, the starting point is the end of the previous segment or the
	// current position if first. Note that the Altitude reference for a given waypoint dictates the height
	//
	//	mode used when traversing TO that waypoint.
	Endpoint *_type.LLA `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *PathSegment) Reset() {
	*x = PathSegment{}
	mi := &file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PathSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathSegment) ProtoMessage() {}

func (x *PathSegment) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_msgTypes[1]
	if x != nil {
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
	return file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_rawDescGZIP(), []int{1}
}

func (x *PathSegment) GetEndpoint() *_type.LLA {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

// TeamTransit represents moving a team of vehicles into a zone.
// The specifics of how each vehicle in the team behaves is determined by the specific autonomy logic.
type TeamTransit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reference to GeoPolygon GeoEntity representing the transit zone area.
	TransitZoneEntityId string `protobuf:"bytes,1,opt,name=transit_zone_entity_id,json=transitZoneEntityId,proto3" json:"transit_zone_entity_id,omitempty"`
	// Speed in which the vehicles will move to the zone.
	SurfaceSpeedMs *wrapperspb.DoubleValue `protobuf:"bytes,2,opt,name=surface_speed_ms,json=surfaceSpeedMs,proto3" json:"surface_speed_ms,omitempty"`
}

func (x *TeamTransit) Reset() {
	*x = TeamTransit{}
	mi := &file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TeamTransit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamTransit) ProtoMessage() {}

func (x *TeamTransit) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamTransit.ProtoReflect.Descriptor instead.
func (*TeamTransit) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_rawDescGZIP(), []int{2}
}

func (x *TeamTransit) GetTransitZoneEntityId() string {
	if x != nil {
		return x.TransitZoneEntityId
	}
	return ""
}

func (x *TeamTransit) GetSurfaceSpeedMs() *wrapperspb.DoubleValue {
	if x != nil {
		return x.SurfaceSpeedMs
	}
	return nil
}

var File_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto protoreflect.FileDescriptor

var file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_rawDesc = []byte{
	0x0a, 0x33, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f,
	0x6a, 0x61, 0x64, 0x63, 0x32, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x2e, 0x70, 0x75, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2e, 0x6a, 0x61, 0x64, 0x63, 0x32, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x75,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x69, 0x74, 0x12, 0x42, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x2e, 0x6a, 0x61, 0x64, 0x63, 0x32, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x46, 0x0a, 0x10, 0x73, 0x75, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0e, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x70, 0x65, 0x65, 0x64, 0x4d, 0x73, 0x22,
	0x3c, 0x0a, 0x0b, 0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2d,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x4c, 0x4c, 0x41, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x8a, 0x01,
	0x0a, 0x0b, 0x54, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x12, 0x33, 0x0a,
	0x16, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x46, 0x0a, 0x10, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x70,
	0x65, 0x65, 0x64, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x73, 0x75, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x53, 0x70, 0x65, 0x65, 0x64, 0x4d, 0x73, 0x42, 0xb3, 0x02, 0x0a, 0x25, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x2e, 0x6a, 0x61, 0x64, 0x63, 0x32, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x50, 0x75, 0x62,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x61, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72,
	0x69, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x6a, 0x61, 0x64, 0x63, 0x32, 0x2f, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x41, 0x54, 0x4a, 0x54,
	0xaa, 0x02, 0x21, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x2e, 0x4a, 0x61, 0x64, 0x63, 0x32, 0x2e, 0x54, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x21, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x5c, 0x4a, 0x61, 0x64, 0x63, 0x32, 0x5c, 0x54, 0x68, 0x69, 0x72, 0x64,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2d, 0x41, 0x6e, 0x64, 0x75, 0x72,
	0x69, 0x6c, 0x5c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x5c, 0x4a, 0x61, 0x64, 0x63, 0x32, 0x5c, 0x54,
	0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25, 0x41, 0x6e, 0x64, 0x75, 0x72,
	0x69, 0x6c, 0x3a, 0x3a, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x3a, 0x3a, 0x4a, 0x61, 0x64, 0x63, 0x32,
	0x3a, 0x3a, 0x54, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_rawDescOnce sync.Once
	file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_rawDescData = file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_rawDesc
)

func file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_rawDescGZIP() []byte {
	file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_rawDescOnce.Do(func() {
		file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_rawDescData)
	})
	return file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_rawDescData
}

var file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_goTypes = []any{
	(*Transit)(nil),                // 0: anduril.tasks.jadc2.thirdparty.v1.Transit
	(*PathSegment)(nil),            // 1: anduril.tasks.jadc2.thirdparty.v1.PathSegment
	(*TeamTransit)(nil),            // 2: anduril.tasks.jadc2.thirdparty.v1.TeamTransit
	(*wrapperspb.DoubleValue)(nil), // 3: google.protobuf.DoubleValue
	(*_type.LLA)(nil),              // 4: anduril.type.LLA
}
var file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_depIdxs = []int32{
	1, // 0: anduril.tasks.jadc2.thirdparty.v1.Transit.path:type_name -> anduril.tasks.jadc2.thirdparty.v1.PathSegment
	3, // 1: anduril.tasks.jadc2.thirdparty.v1.Transit.surface_speed_ms:type_name -> google.protobuf.DoubleValue
	4, // 2: anduril.tasks.jadc2.thirdparty.v1.PathSegment.endpoint:type_name -> anduril.type.LLA
	3, // 3: anduril.tasks.jadc2.thirdparty.v1.TeamTransit.surface_speed_ms:type_name -> google.protobuf.DoubleValue
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_init() }
func file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_init() {
	if File_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_goTypes,
		DependencyIndexes: file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_depIdxs,
		MessageInfos:      file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_msgTypes,
	}.Build()
	File_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto = out.File
	file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_rawDesc = nil
	file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_goTypes = nil
	file_anduril_tasks_jadc2_thirdparty_v1_transit_pub_proto_depIdxs = nil
}
