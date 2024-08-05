// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: anduril/tasks/v2/common.pub.proto

package tasksv2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ControlAreaType int32

const (
	ControlAreaType_CONTROL_AREA_TYPE_INVALID       ControlAreaType = 0
	ControlAreaType_CONTROL_AREA_TYPE_KEEP_IN_ZONE  ControlAreaType = 1
	ControlAreaType_CONTROL_AREA_TYPE_KEEP_OUT_ZONE ControlAreaType = 2
	// Zone for an autonomous asset to nose-dive into
	// when its assignment has been concluded
	ControlAreaType_CONTROL_AREA_TYPE_DITCH_ZONE ControlAreaType = 3
)

// Enum value maps for ControlAreaType.
var (
	ControlAreaType_name = map[int32]string{
		0: "CONTROL_AREA_TYPE_INVALID",
		1: "CONTROL_AREA_TYPE_KEEP_IN_ZONE",
		2: "CONTROL_AREA_TYPE_KEEP_OUT_ZONE",
		3: "CONTROL_AREA_TYPE_DITCH_ZONE",
	}
	ControlAreaType_value = map[string]int32{
		"CONTROL_AREA_TYPE_INVALID":       0,
		"CONTROL_AREA_TYPE_KEEP_IN_ZONE":  1,
		"CONTROL_AREA_TYPE_KEEP_OUT_ZONE": 2,
		"CONTROL_AREA_TYPE_DITCH_ZONE":    3,
	}
)

func (x ControlAreaType) Enum() *ControlAreaType {
	p := new(ControlAreaType)
	*p = x
	return p
}

func (x ControlAreaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ControlAreaType) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_tasks_v2_common_pub_proto_enumTypes[0].Descriptor()
}

func (ControlAreaType) Type() protoreflect.EnumType {
	return &file_anduril_tasks_v2_common_pub_proto_enumTypes[0]
}

func (x ControlAreaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ControlAreaType.Descriptor instead.
func (ControlAreaType) EnumDescriptor() ([]byte, []int) {
	return file_anduril_tasks_v2_common_pub_proto_rawDescGZIP(), []int{0}
}

// Maps to the UCI DurationRangeType.
type DurationRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min *durationpb.Duration `protobuf:"bytes,1,opt,name=min,proto3" json:"min,omitempty"`
	Max *durationpb.Duration `protobuf:"bytes,2,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *DurationRange) Reset() {
	*x = DurationRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_tasks_v2_common_pub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationRange) ProtoMessage() {}

func (x *DurationRange) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_v2_common_pub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationRange.ProtoReflect.Descriptor instead.
func (*DurationRange) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_v2_common_pub_proto_rawDescGZIP(), []int{0}
}

func (x *DurationRange) GetMin() *durationpb.Duration {
	if x != nil {
		return x.Min
	}
	return nil
}

func (x *DurationRange) GetMax() *durationpb.Duration {
	if x != nil {
		return x.Max
	}
	return nil
}

// Maps to the UCI AnglePair.
type AnglePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Angle lower bound in radians.
	Min float64 `protobuf:"fixed64,1,opt,name=min,proto3" json:"min,omitempty"`
	// Angle lower bound in radians.
	Max float64 `protobuf:"fixed64,2,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *AnglePair) Reset() {
	*x = AnglePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_tasks_v2_common_pub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnglePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnglePair) ProtoMessage() {}

func (x *AnglePair) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_v2_common_pub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnglePair.ProtoReflect.Descriptor instead.
func (*AnglePair) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_v2_common_pub_proto_rawDescGZIP(), []int{1}
}

func (x *AnglePair) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *AnglePair) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

// Maps to UCI AreaConstraints.
type AreaConstraints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AltitudeConstraint *AltitudeConstraint `protobuf:"bytes,1,opt,name=altitude_constraint,json=altitudeConstraint,proto3" json:"altitude_constraint,omitempty"`
}

func (x *AreaConstraints) Reset() {
	*x = AreaConstraints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_tasks_v2_common_pub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AreaConstraints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AreaConstraints) ProtoMessage() {}

func (x *AreaConstraints) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_v2_common_pub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AreaConstraints.ProtoReflect.Descriptor instead.
func (*AreaConstraints) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_v2_common_pub_proto_rawDescGZIP(), []int{2}
}

func (x *AreaConstraints) GetAltitudeConstraint() *AltitudeConstraint {
	if x != nil {
		return x.AltitudeConstraint
	}
	return nil
}

type AltitudeConstraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Minimum altitude (AGL) in meters.
	Min float64 `protobuf:"fixed64,1,opt,name=min,proto3" json:"min,omitempty"`
	// Maximum altitude (AGL) in meters.
	Max float64 `protobuf:"fixed64,2,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *AltitudeConstraint) Reset() {
	*x = AltitudeConstraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_tasks_v2_common_pub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AltitudeConstraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AltitudeConstraint) ProtoMessage() {}

func (x *AltitudeConstraint) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_v2_common_pub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AltitudeConstraint.ProtoReflect.Descriptor instead.
func (*AltitudeConstraint) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_v2_common_pub_proto_rawDescGZIP(), []int{3}
}

func (x *AltitudeConstraint) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *AltitudeConstraint) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

// Includes information about an Agent.
type Agent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetId  string `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	EntityId string `protobuf:"bytes,2,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
}

func (x *Agent) Reset() {
	*x = Agent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_tasks_v2_common_pub_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agent) ProtoMessage() {}

func (x *Agent) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_v2_common_pub_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agent.ProtoReflect.Descriptor instead.
func (*Agent) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_v2_common_pub_proto_rawDescGZIP(), []int{4}
}

func (x *Agent) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *Agent) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

// Models a Control Area within which Agents must operate.
type ControlArea struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reference to GeoPolygon GeoEntity representing the ControlArea.
	EntityId string `protobuf:"bytes,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	// Type of ControlArea.
	ControlAreaType ControlAreaType `protobuf:"varint,2,opt,name=control_area_type,json=controlAreaType,proto3,enum=anduril.tasks.v2.ControlAreaType" json:"control_area_type,omitempty"`
}

func (x *ControlArea) Reset() {
	*x = ControlArea{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_tasks_v2_common_pub_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlArea) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlArea) ProtoMessage() {}

func (x *ControlArea) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_v2_common_pub_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlArea.ProtoReflect.Descriptor instead.
func (*ControlArea) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_v2_common_pub_proto_rawDescGZIP(), []int{5}
}

func (x *ControlArea) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *ControlArea) GetControlAreaType() ControlAreaType {
	if x != nil {
		return x.ControlAreaType
	}
	return ControlAreaType_CONTROL_AREA_TYPE_INVALID
}

var File_anduril_tasks_v2_common_pub_proto protoreflect.FileDescriptor

var file_anduril_tasks_v2_common_pub_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f,
	0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0d, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03,
	0x6d, 0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x6d, 0x61, 0x78,
	0x22, 0x2f, 0x0a, 0x09, 0x41, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x61,
	0x78, 0x22, 0x68, 0x0a, 0x0f, 0x41, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x74, 0x73, 0x12, 0x55, 0x0a, 0x13, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x12, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x22, 0x38, 0x0a, 0x12, 0x41,
	0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x6d, 0x61, 0x78, 0x22, 0x3f, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x79, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x41, 0x72, 0x65, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x4d, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x61, 0x72,
	0x65, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e,
	0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x41, 0x72, 0x65, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x41, 0x72, 0x65, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x2a, 0x9b, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x41, 0x72, 0x65,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c,
	0x5f, 0x41, 0x52, 0x45, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f,
	0x41, 0x52, 0x45, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4b, 0x45, 0x45, 0x50, 0x5f, 0x49,
	0x4e, 0x5f, 0x5a, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x4f, 0x4e, 0x54,
	0x52, 0x4f, 0x4c, 0x5f, 0x41, 0x52, 0x45, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4b, 0x45,
	0x45, 0x50, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x5a, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x20, 0x0a,
	0x1c, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x41, 0x52, 0x45, 0x41, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x44, 0x49, 0x54, 0x43, 0x48, 0x5f, 0x5a, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x42,
	0xce, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x42, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x50, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x75, 0x6e, 0x2d, 0x61, 0x6e, 0x64,
	0x75, 0x72, 0x69, 0x6c, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2d, 0x73, 0x64, 0x6b,
	0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x32, 0x3b, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x76, 0x32,
	0xa2, 0x02, 0x03, 0x41, 0x54, 0x58, 0xaa, 0x02, 0x10, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x10, 0x41, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x5c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x1c, 0x41,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x5c, 0x56, 0x32, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x41, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x3a, 0x3a, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x3a, 0x3a, 0x56, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_anduril_tasks_v2_common_pub_proto_rawDescOnce sync.Once
	file_anduril_tasks_v2_common_pub_proto_rawDescData = file_anduril_tasks_v2_common_pub_proto_rawDesc
)

func file_anduril_tasks_v2_common_pub_proto_rawDescGZIP() []byte {
	file_anduril_tasks_v2_common_pub_proto_rawDescOnce.Do(func() {
		file_anduril_tasks_v2_common_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_tasks_v2_common_pub_proto_rawDescData)
	})
	return file_anduril_tasks_v2_common_pub_proto_rawDescData
}

var file_anduril_tasks_v2_common_pub_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_anduril_tasks_v2_common_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_anduril_tasks_v2_common_pub_proto_goTypes = []any{
	(ControlAreaType)(0),        // 0: anduril.tasks.v2.ControlAreaType
	(*DurationRange)(nil),       // 1: anduril.tasks.v2.DurationRange
	(*AnglePair)(nil),           // 2: anduril.tasks.v2.AnglePair
	(*AreaConstraints)(nil),     // 3: anduril.tasks.v2.AreaConstraints
	(*AltitudeConstraint)(nil),  // 4: anduril.tasks.v2.AltitudeConstraint
	(*Agent)(nil),               // 5: anduril.tasks.v2.Agent
	(*ControlArea)(nil),         // 6: anduril.tasks.v2.ControlArea
	(*durationpb.Duration)(nil), // 7: google.protobuf.Duration
}
var file_anduril_tasks_v2_common_pub_proto_depIdxs = []int32{
	7, // 0: anduril.tasks.v2.DurationRange.min:type_name -> google.protobuf.Duration
	7, // 1: anduril.tasks.v2.DurationRange.max:type_name -> google.protobuf.Duration
	4, // 2: anduril.tasks.v2.AreaConstraints.altitude_constraint:type_name -> anduril.tasks.v2.AltitudeConstraint
	0, // 3: anduril.tasks.v2.ControlArea.control_area_type:type_name -> anduril.tasks.v2.ControlAreaType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_anduril_tasks_v2_common_pub_proto_init() }
func file_anduril_tasks_v2_common_pub_proto_init() {
	if File_anduril_tasks_v2_common_pub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_anduril_tasks_v2_common_pub_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DurationRange); i {
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
		file_anduril_tasks_v2_common_pub_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AnglePair); i {
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
		file_anduril_tasks_v2_common_pub_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AreaConstraints); i {
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
		file_anduril_tasks_v2_common_pub_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AltitudeConstraint); i {
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
		file_anduril_tasks_v2_common_pub_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Agent); i {
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
		file_anduril_tasks_v2_common_pub_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ControlArea); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_anduril_tasks_v2_common_pub_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_tasks_v2_common_pub_proto_goTypes,
		DependencyIndexes: file_anduril_tasks_v2_common_pub_proto_depIdxs,
		EnumInfos:         file_anduril_tasks_v2_common_pub_proto_enumTypes,
		MessageInfos:      file_anduril_tasks_v2_common_pub_proto_msgTypes,
	}.Build()
	File_anduril_tasks_v2_common_pub_proto = out.File
	file_anduril_tasks_v2_common_pub_proto_rawDesc = nil
	file_anduril_tasks_v2_common_pub_proto_goTypes = nil
	file_anduril_tasks_v2_common_pub_proto_depIdxs = nil
}