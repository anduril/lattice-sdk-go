// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: anduril/entitymanager/v1/relationship.pub.proto

package entitymanagerv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The relationships between this entity and other entities in the common operational picture.
type Relationships struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Relationships []*Relationship        `protobuf:"bytes,1,rep,name=relationships,proto3" json:"relationships,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Relationships) Reset() {
	*x = Relationships{}
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Relationships) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relationships) ProtoMessage() {}

func (x *Relationships) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relationships.ProtoReflect.Descriptor instead.
func (*Relationships) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_relationship_pub_proto_rawDescGZIP(), []int{0}
}

func (x *Relationships) GetRelationships() []*Relationship {
	if x != nil {
		return x.Relationships
	}
	return nil
}

// The relationship component indicates a relationship to another entity.
type Relationship struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The entity ID to which this entity is related.
	RelatedEntityId string `protobuf:"bytes,1,opt,name=related_entity_id,json=relatedEntityId,proto3" json:"related_entity_id,omitempty"`
	// A unique identifier for this relationship. Allows removing or updating relationships.
	RelationshipId string `protobuf:"bytes,2,opt,name=relationship_id,json=relationshipId,proto3" json:"relationship_id,omitempty"`
	// The relationship type
	RelationshipType *RelationshipType `protobuf:"bytes,3,opt,name=relationship_type,json=relationshipType,proto3" json:"relationship_type,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Relationship) Reset() {
	*x = Relationship{}
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Relationship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relationship) ProtoMessage() {}

func (x *Relationship) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relationship.ProtoReflect.Descriptor instead.
func (*Relationship) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_relationship_pub_proto_rawDescGZIP(), []int{1}
}

func (x *Relationship) GetRelatedEntityId() string {
	if x != nil {
		return x.RelatedEntityId
	}
	return ""
}

func (x *Relationship) GetRelationshipId() string {
	if x != nil {
		return x.RelationshipId
	}
	return ""
}

func (x *Relationship) GetRelationshipType() *RelationshipType {
	if x != nil {
		return x.RelationshipType
	}
	return nil
}

// Determines the type of relationship between this entity and another.
type RelationshipType struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Type:
	//
	//	*RelationshipType_TrackedBy
	//	*RelationshipType_GroupChild
	//	*RelationshipType_GroupParent
	//	*RelationshipType_MergedFrom
	//	*RelationshipType_ActiveTarget
	Type          isRelationshipType_Type `protobuf_oneof:"type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RelationshipType) Reset() {
	*x = RelationshipType{}
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RelationshipType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationshipType) ProtoMessage() {}

func (x *RelationshipType) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationshipType.ProtoReflect.Descriptor instead.
func (*RelationshipType) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_relationship_pub_proto_rawDescGZIP(), []int{2}
}

func (x *RelationshipType) GetType() isRelationshipType_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *RelationshipType) GetTrackedBy() *TrackedBy {
	if x != nil {
		if x, ok := x.Type.(*RelationshipType_TrackedBy); ok {
			return x.TrackedBy
		}
	}
	return nil
}

func (x *RelationshipType) GetGroupChild() *GroupChild {
	if x != nil {
		if x, ok := x.Type.(*RelationshipType_GroupChild); ok {
			return x.GroupChild
		}
	}
	return nil
}

func (x *RelationshipType) GetGroupParent() *GroupParent {
	if x != nil {
		if x, ok := x.Type.(*RelationshipType_GroupParent); ok {
			return x.GroupParent
		}
	}
	return nil
}

func (x *RelationshipType) GetMergedFrom() *MergedFrom {
	if x != nil {
		if x, ok := x.Type.(*RelationshipType_MergedFrom); ok {
			return x.MergedFrom
		}
	}
	return nil
}

func (x *RelationshipType) GetActiveTarget() *ActiveTarget {
	if x != nil {
		if x, ok := x.Type.(*RelationshipType_ActiveTarget); ok {
			return x.ActiveTarget
		}
	}
	return nil
}

type isRelationshipType_Type interface {
	isRelationshipType_Type()
}

type RelationshipType_TrackedBy struct {
	TrackedBy *TrackedBy `protobuf:"bytes,2,opt,name=tracked_by,json=trackedBy,proto3,oneof"`
}

type RelationshipType_GroupChild struct {
	GroupChild *GroupChild `protobuf:"bytes,4,opt,name=group_child,json=groupChild,proto3,oneof"`
}

type RelationshipType_GroupParent struct {
	GroupParent *GroupParent `protobuf:"bytes,5,opt,name=group_parent,json=groupParent,proto3,oneof"`
}

type RelationshipType_MergedFrom struct {
	MergedFrom *MergedFrom `protobuf:"bytes,6,opt,name=merged_from,json=mergedFrom,proto3,oneof"`
}

type RelationshipType_ActiveTarget struct {
	ActiveTarget *ActiveTarget `protobuf:"bytes,7,opt,name=active_target,json=activeTarget,proto3,oneof"`
}

func (*RelationshipType_TrackedBy) isRelationshipType_Type() {}

func (*RelationshipType_GroupChild) isRelationshipType_Type() {}

func (*RelationshipType_GroupParent) isRelationshipType_Type() {}

func (*RelationshipType_MergedFrom) isRelationshipType_Type() {}

func (*RelationshipType_ActiveTarget) isRelationshipType_Type() {}

// Describes the relationship between the entity being tracked ("tracked entity") and the entity that is
// performing the tracking ("tracking entity").
type TrackedBy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Sensor details of the tracking entity's sensors that were active and tracking the tracked entity. This may be
	// a subset of the total sensors available on the tracking entity.
	ActivelyTrackingSensors *Sensors `protobuf:"bytes,1,opt,name=actively_tracking_sensors,json=activelyTrackingSensors,proto3" json:"actively_tracking_sensors,omitempty"`
	// Latest time that any sensor in actively_tracking_sensors detected the tracked entity.
	LastMeasurementTimestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_measurement_timestamp,json=lastMeasurementTimestamp,proto3" json:"last_measurement_timestamp,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *TrackedBy) Reset() {
	*x = TrackedBy{}
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrackedBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackedBy) ProtoMessage() {}

func (x *TrackedBy) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackedBy.ProtoReflect.Descriptor instead.
func (*TrackedBy) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_relationship_pub_proto_rawDescGZIP(), []int{3}
}

func (x *TrackedBy) GetActivelyTrackingSensors() *Sensors {
	if x != nil {
		return x.ActivelyTrackingSensors
	}
	return nil
}

func (x *TrackedBy) GetLastMeasurementTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.LastMeasurementTimestamp
	}
	return nil
}

// A GroupChild relationship is a uni-directional relationship indicating that (1) this entity
// represents an Entity Group and (2) the related entity is a child member of this group. The presence of this
// relationship alone determines that the type of group is an Entity Group.
type GroupChild struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GroupChild) Reset() {
	*x = GroupChild{}
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupChild) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupChild) ProtoMessage() {}

func (x *GroupChild) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupChild.ProtoReflect.Descriptor instead.
func (*GroupChild) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_relationship_pub_proto_rawDescGZIP(), []int{4}
}

// A GroupParent relationship is a uni-directional relationship indicating that this entity is a member of
// the Entity Group represented by the related entity. The presence of this relationship alone determines that
// the type of group that this entity is a member of is an Entity Group.
type GroupParent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GroupParent) Reset() {
	*x = GroupParent{}
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupParent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupParent) ProtoMessage() {}

func (x *GroupParent) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupParent.ProtoReflect.Descriptor instead.
func (*GroupParent) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_relationship_pub_proto_rawDescGZIP(), []int{5}
}

// A MergedFrom relationship is a uni-directional relationship indicating that this entity is a merged entity whose
// data has at least partially been merged from the related entity.
type MergedFrom struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MergedFrom) Reset() {
	*x = MergedFrom{}
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MergedFrom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergedFrom) ProtoMessage() {}

func (x *MergedFrom) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergedFrom.ProtoReflect.Descriptor instead.
func (*MergedFrom) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_relationship_pub_proto_rawDescGZIP(), []int{6}
}

// A target relationship is the inverse of TrackedBy; a one-way relation
// from sensor to target, indicating track(s) currently prioritized by a robot.
type ActiveTarget struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActiveTarget) Reset() {
	*x = ActiveTarget{}
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActiveTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveTarget) ProtoMessage() {}

func (x *ActiveTarget) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveTarget.ProtoReflect.Descriptor instead.
func (*ActiveTarget) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_relationship_pub_proto_rawDescGZIP(), []int{7}
}

var File_anduril_entitymanager_v1_relationship_pub_proto protoreflect.FileDescriptor

const file_anduril_entitymanager_v1_relationship_pub_proto_rawDesc = "" +
	"\n" +
	"/anduril/entitymanager/v1/relationship.pub.proto\x12\x18anduril.entitymanager.v1\x1a*anduril/entitymanager/v1/sensors.pub.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"]\n" +
	"\rRelationships\x12L\n" +
	"\rrelationships\x18\x01 \x03(\v2&.anduril.entitymanager.v1.RelationshipR\rrelationships\"\xbc\x01\n" +
	"\fRelationship\x12*\n" +
	"\x11related_entity_id\x18\x01 \x01(\tR\x0frelatedEntityId\x12'\n" +
	"\x0frelationship_id\x18\x02 \x01(\tR\x0erelationshipId\x12W\n" +
	"\x11relationship_type\x18\x03 \x01(\v2*.anduril.entitymanager.v1.RelationshipTypeR\x10relationshipType\"\x8d\x03\n" +
	"\x10RelationshipType\x12D\n" +
	"\n" +
	"tracked_by\x18\x02 \x01(\v2#.anduril.entitymanager.v1.TrackedByH\x00R\ttrackedBy\x12G\n" +
	"\vgroup_child\x18\x04 \x01(\v2$.anduril.entitymanager.v1.GroupChildH\x00R\n" +
	"groupChild\x12J\n" +
	"\fgroup_parent\x18\x05 \x01(\v2%.anduril.entitymanager.v1.GroupParentH\x00R\vgroupParent\x12G\n" +
	"\vmerged_from\x18\x06 \x01(\v2$.anduril.entitymanager.v1.MergedFromH\x00R\n" +
	"mergedFrom\x12M\n" +
	"\ractive_target\x18\a \x01(\v2&.anduril.entitymanager.v1.ActiveTargetH\x00R\factiveTargetB\x06\n" +
	"\x04type\"\xc4\x01\n" +
	"\tTrackedBy\x12]\n" +
	"\x19actively_tracking_sensors\x18\x01 \x01(\v2!.anduril.entitymanager.v1.SensorsR\x17activelyTrackingSensors\x12X\n" +
	"\x1alast_measurement_timestamp\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x18lastMeasurementTimestamp\"\f\n" +
	"\n" +
	"GroupChild\"\r\n" +
	"\vGroupParent\"\f\n" +
	"\n" +
	"MergedFrom\"\x0e\n" +
	"\fActiveTargetB\x86\x02\n" +
	"\x1ccom.anduril.entitymanager.v1B\x14RelationshipPubProtoP\x01ZNgithub.com/anduril/lattice-sdk-go/src/anduril/entitymanager/v1;entitymanagerv1\xa2\x02\x03AEX\xaa\x02\x18Anduril.Entitymanager.V1\xca\x02\x18Anduril\\Entitymanager\\V1\xe2\x02$Anduril\\Entitymanager\\V1\\GPBMetadata\xea\x02\x1aAnduril::Entitymanager::V1b\x06proto3"

var (
	file_anduril_entitymanager_v1_relationship_pub_proto_rawDescOnce sync.Once
	file_anduril_entitymanager_v1_relationship_pub_proto_rawDescData []byte
)

func file_anduril_entitymanager_v1_relationship_pub_proto_rawDescGZIP() []byte {
	file_anduril_entitymanager_v1_relationship_pub_proto_rawDescOnce.Do(func() {
		file_anduril_entitymanager_v1_relationship_pub_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_anduril_entitymanager_v1_relationship_pub_proto_rawDesc), len(file_anduril_entitymanager_v1_relationship_pub_proto_rawDesc)))
	})
	return file_anduril_entitymanager_v1_relationship_pub_proto_rawDescData
}

var file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_anduril_entitymanager_v1_relationship_pub_proto_goTypes = []any{
	(*Relationships)(nil),         // 0: anduril.entitymanager.v1.Relationships
	(*Relationship)(nil),          // 1: anduril.entitymanager.v1.Relationship
	(*RelationshipType)(nil),      // 2: anduril.entitymanager.v1.RelationshipType
	(*TrackedBy)(nil),             // 3: anduril.entitymanager.v1.TrackedBy
	(*GroupChild)(nil),            // 4: anduril.entitymanager.v1.GroupChild
	(*GroupParent)(nil),           // 5: anduril.entitymanager.v1.GroupParent
	(*MergedFrom)(nil),            // 6: anduril.entitymanager.v1.MergedFrom
	(*ActiveTarget)(nil),          // 7: anduril.entitymanager.v1.ActiveTarget
	(*Sensors)(nil),               // 8: anduril.entitymanager.v1.Sensors
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_anduril_entitymanager_v1_relationship_pub_proto_depIdxs = []int32{
	1, // 0: anduril.entitymanager.v1.Relationships.relationships:type_name -> anduril.entitymanager.v1.Relationship
	2, // 1: anduril.entitymanager.v1.Relationship.relationship_type:type_name -> anduril.entitymanager.v1.RelationshipType
	3, // 2: anduril.entitymanager.v1.RelationshipType.tracked_by:type_name -> anduril.entitymanager.v1.TrackedBy
	4, // 3: anduril.entitymanager.v1.RelationshipType.group_child:type_name -> anduril.entitymanager.v1.GroupChild
	5, // 4: anduril.entitymanager.v1.RelationshipType.group_parent:type_name -> anduril.entitymanager.v1.GroupParent
	6, // 5: anduril.entitymanager.v1.RelationshipType.merged_from:type_name -> anduril.entitymanager.v1.MergedFrom
	7, // 6: anduril.entitymanager.v1.RelationshipType.active_target:type_name -> anduril.entitymanager.v1.ActiveTarget
	8, // 7: anduril.entitymanager.v1.TrackedBy.actively_tracking_sensors:type_name -> anduril.entitymanager.v1.Sensors
	9, // 8: anduril.entitymanager.v1.TrackedBy.last_measurement_timestamp:type_name -> google.protobuf.Timestamp
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_anduril_entitymanager_v1_relationship_pub_proto_init() }
func file_anduril_entitymanager_v1_relationship_pub_proto_init() {
	if File_anduril_entitymanager_v1_relationship_pub_proto != nil {
		return
	}
	file_anduril_entitymanager_v1_sensors_pub_proto_init()
	file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[2].OneofWrappers = []any{
		(*RelationshipType_TrackedBy)(nil),
		(*RelationshipType_GroupChild)(nil),
		(*RelationshipType_GroupParent)(nil),
		(*RelationshipType_MergedFrom)(nil),
		(*RelationshipType_ActiveTarget)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_anduril_entitymanager_v1_relationship_pub_proto_rawDesc), len(file_anduril_entitymanager_v1_relationship_pub_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_entitymanager_v1_relationship_pub_proto_goTypes,
		DependencyIndexes: file_anduril_entitymanager_v1_relationship_pub_proto_depIdxs,
		MessageInfos:      file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes,
	}.Build()
	File_anduril_entitymanager_v1_relationship_pub_proto = out.File
	file_anduril_entitymanager_v1_relationship_pub_proto_goTypes = nil
	file_anduril_entitymanager_v1_relationship_pub_proto_depIdxs = nil
}
