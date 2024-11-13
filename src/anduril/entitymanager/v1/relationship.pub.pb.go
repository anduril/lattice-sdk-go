// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: anduril/entitymanager/v1/relationship.pub.proto

package entitymanagerv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The relationships between this entity and other entities in the battlespace.
type Relationships struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relationships []*Relationship `protobuf:"bytes,1,rep,name=relationships,proto3" json:"relationships,omitempty"`
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The entity ID to which this entity is related.
	RelatedEntityId string `protobuf:"bytes,1,opt,name=related_entity_id,json=relatedEntityId,proto3" json:"related_entity_id,omitempty"`
	// A unique identifier for this relationship. Allows removing or updating relationships.
	RelationshipId string `protobuf:"bytes,2,opt,name=relationship_id,json=relationshipId,proto3" json:"relationship_id,omitempty"`
	// The relationship type
	RelationshipType *RelationshipType `protobuf:"bytes,3,opt,name=relationship_type,json=relationshipType,proto3" json:"relationship_type,omitempty"`
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*RelationshipType_Tether
	//	*RelationshipType_TrackedBy
	//	*RelationshipType_Configure
	//	*RelationshipType_GroupChild
	//	*RelationshipType_GroupParent
	//	*RelationshipType_MergedFrom
	Type isRelationshipType_Type `protobuf_oneof:"type"`
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

func (m *RelationshipType) GetType() isRelationshipType_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *RelationshipType) GetTether() *Tether {
	if x, ok := x.GetType().(*RelationshipType_Tether); ok {
		return x.Tether
	}
	return nil
}

func (x *RelationshipType) GetTrackedBy() *TrackedBy {
	if x, ok := x.GetType().(*RelationshipType_TrackedBy); ok {
		return x.TrackedBy
	}
	return nil
}

func (x *RelationshipType) GetConfigure() *Configure {
	if x, ok := x.GetType().(*RelationshipType_Configure); ok {
		return x.Configure
	}
	return nil
}

func (x *RelationshipType) GetGroupChild() *GroupChild {
	if x, ok := x.GetType().(*RelationshipType_GroupChild); ok {
		return x.GroupChild
	}
	return nil
}

func (x *RelationshipType) GetGroupParent() *GroupParent {
	if x, ok := x.GetType().(*RelationshipType_GroupParent); ok {
		return x.GroupParent
	}
	return nil
}

func (x *RelationshipType) GetMergedFrom() *MergedFrom {
	if x, ok := x.GetType().(*RelationshipType_MergedFrom); ok {
		return x.MergedFrom
	}
	return nil
}

type isRelationshipType_Type interface {
	isRelationshipType_Type()
}

type RelationshipType_Tether struct {
	Tether *Tether `protobuf:"bytes,1,opt,name=tether,proto3,oneof"`
}

type RelationshipType_TrackedBy struct {
	TrackedBy *TrackedBy `protobuf:"bytes,2,opt,name=tracked_by,json=trackedBy,proto3,oneof"`
}

type RelationshipType_Configure struct {
	Configure *Configure `protobuf:"bytes,3,opt,name=configure,proto3,oneof"`
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

func (*RelationshipType_Tether) isRelationshipType_Type() {}

func (*RelationshipType_TrackedBy) isRelationshipType_Type() {}

func (*RelationshipType_Configure) isRelationshipType_Type() {}

func (*RelationshipType_GroupChild) isRelationshipType_Type() {}

func (*RelationshipType_GroupParent) isRelationshipType_Type() {}

func (*RelationshipType_MergedFrom) isRelationshipType_Type() {}

// A tether relationship indicates that this entity should take the position of the other entity.
type Tether struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Tether) Reset() {
	*x = Tether{}
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tether) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tether) ProtoMessage() {}

func (x *Tether) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Tether.ProtoReflect.Descriptor instead.
func (*Tether) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_relationship_pub_proto_rawDescGZIP(), []int{3}
}

// Describes the relationship between the entity being tracked ("tracked entity") and the entity that is
// performing the tracking ("tracking entity").
type TrackedBy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sensor details of the tracking entity's sensors that were active and tracking the tracked entity. This may be
	// a subset of the total sensors available on the tracking entity.
	ActivelyTrackingSensors *Sensors `protobuf:"bytes,1,opt,name=actively_tracking_sensors,json=activelyTrackingSensors,proto3" json:"actively_tracking_sensors,omitempty"`
	// Latest time that any sensor in actively_tracking_sensors detected the tracked entity.
	LastMeasurementTimestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_measurement_timestamp,json=lastMeasurementTimestamp,proto3" json:"last_measurement_timestamp,omitempty"`
}

func (x *TrackedBy) Reset() {
	*x = TrackedBy{}
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrackedBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackedBy) ProtoMessage() {}

func (x *TrackedBy) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TrackedBy.ProtoReflect.Descriptor instead.
func (*TrackedBy) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_relationship_pub_proto_rawDescGZIP(), []int{4}
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

// A configure relationship indicates that this entity is a configuration on other entity.
type Configure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Configure) Reset() {
	*x = Configure{}
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Configure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configure) ProtoMessage() {}

func (x *Configure) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Configure.ProtoReflect.Descriptor instead.
func (*Configure) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_relationship_pub_proto_rawDescGZIP(), []int{5}
}

// A GroupChild relationship is a uni-directional relationship indicating that (1) this entity
// represents an Entity Group and (2) the related entity is a child member of this group. The presence of this
// relationship alone determines that the type of group is an Entity Group.
type GroupChild struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GroupChild) Reset() {
	*x = GroupChild{}
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupChild) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupChild) ProtoMessage() {}

func (x *GroupChild) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GroupChild.ProtoReflect.Descriptor instead.
func (*GroupChild) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_relationship_pub_proto_rawDescGZIP(), []int{6}
}

// A GroupParent relationship is a uni-directional relationship indicating that this entity is a member of
// the Entity Group represented by the related entity. The presence of this relationship alone determines that
// the type of group that this entity is a member of is an Entity Group.
type GroupParent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GroupParent) Reset() {
	*x = GroupParent{}
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupParent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupParent) ProtoMessage() {}

func (x *GroupParent) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GroupParent.ProtoReflect.Descriptor instead.
func (*GroupParent) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_relationship_pub_proto_rawDescGZIP(), []int{7}
}

// A MergedFrom relationship is a uni-directional relationship indicating that this entity is a merged entity whose
// data has at least partially been merged from the related entity.
type MergedFrom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MergedFrom) Reset() {
	*x = MergedFrom{}
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MergedFrom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergedFrom) ProtoMessage() {}

func (x *MergedFrom) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[8]
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
	return file_anduril_entitymanager_v1_relationship_pub_proto_rawDescGZIP(), []int{8}
}

var File_anduril_entitymanager_v1_relationship_pub_proto protoreflect.FileDescriptor

var file_anduril_entitymanager_v1_relationship_pub_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x18, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x61, 0x6e, 0x64,
	0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x75,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x4c, 0x0a, 0x0d, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x49, 0x64, 0x12, 0x57, 0x0a,
	0x11, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72,
	0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68,
	0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22, 0xbf, 0x03, 0x0a, 0x10, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x74,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x74, 0x68, 0x65, 0x72, 0x48, 0x00, 0x52,
	0x06, 0x74, 0x65, 0x74, 0x68, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x79,
	0x48, 0x00, 0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x79, 0x12, 0x43, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x68, 0x69, 0x6c,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x48, 0x00, 0x52,
	0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x4a, 0x0a, 0x0c, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x67, 0x65,
	0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x64, 0x46, 0x72,
	0x6f, 0x6d, 0x48, 0x00, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d,
	0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x08, 0x0a, 0x06, 0x54, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x22, 0xc4, 0x01, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x79,
	0x12, 0x5d, 0x0a, 0x19, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x79, 0x5f, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x52, 0x17, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x79,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x12,
	0x58, 0x0a, 0x1a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x18, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x0b, 0x0a, 0x09, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43,
	0x68, 0x69, 0x6c, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x64, 0x46, 0x72, 0x6f,
	0x6d, 0x42, 0x82, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x42, 0x14, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x50, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x61,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x45, 0x58, 0xaa, 0x02, 0x18, 0x41,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x24, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x41, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_anduril_entitymanager_v1_relationship_pub_proto_rawDescOnce sync.Once
	file_anduril_entitymanager_v1_relationship_pub_proto_rawDescData = file_anduril_entitymanager_v1_relationship_pub_proto_rawDesc
)

func file_anduril_entitymanager_v1_relationship_pub_proto_rawDescGZIP() []byte {
	file_anduril_entitymanager_v1_relationship_pub_proto_rawDescOnce.Do(func() {
		file_anduril_entitymanager_v1_relationship_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_entitymanager_v1_relationship_pub_proto_rawDescData)
	})
	return file_anduril_entitymanager_v1_relationship_pub_proto_rawDescData
}

var file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_anduril_entitymanager_v1_relationship_pub_proto_goTypes = []any{
	(*Relationships)(nil),         // 0: anduril.entitymanager.v1.Relationships
	(*Relationship)(nil),          // 1: anduril.entitymanager.v1.Relationship
	(*RelationshipType)(nil),      // 2: anduril.entitymanager.v1.RelationshipType
	(*Tether)(nil),                // 3: anduril.entitymanager.v1.Tether
	(*TrackedBy)(nil),             // 4: anduril.entitymanager.v1.TrackedBy
	(*Configure)(nil),             // 5: anduril.entitymanager.v1.Configure
	(*GroupChild)(nil),            // 6: anduril.entitymanager.v1.GroupChild
	(*GroupParent)(nil),           // 7: anduril.entitymanager.v1.GroupParent
	(*MergedFrom)(nil),            // 8: anduril.entitymanager.v1.MergedFrom
	(*Sensors)(nil),               // 9: anduril.entitymanager.v1.Sensors
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_anduril_entitymanager_v1_relationship_pub_proto_depIdxs = []int32{
	1,  // 0: anduril.entitymanager.v1.Relationships.relationships:type_name -> anduril.entitymanager.v1.Relationship
	2,  // 1: anduril.entitymanager.v1.Relationship.relationship_type:type_name -> anduril.entitymanager.v1.RelationshipType
	3,  // 2: anduril.entitymanager.v1.RelationshipType.tether:type_name -> anduril.entitymanager.v1.Tether
	4,  // 3: anduril.entitymanager.v1.RelationshipType.tracked_by:type_name -> anduril.entitymanager.v1.TrackedBy
	5,  // 4: anduril.entitymanager.v1.RelationshipType.configure:type_name -> anduril.entitymanager.v1.Configure
	6,  // 5: anduril.entitymanager.v1.RelationshipType.group_child:type_name -> anduril.entitymanager.v1.GroupChild
	7,  // 6: anduril.entitymanager.v1.RelationshipType.group_parent:type_name -> anduril.entitymanager.v1.GroupParent
	8,  // 7: anduril.entitymanager.v1.RelationshipType.merged_from:type_name -> anduril.entitymanager.v1.MergedFrom
	9,  // 8: anduril.entitymanager.v1.TrackedBy.actively_tracking_sensors:type_name -> anduril.entitymanager.v1.Sensors
	10, // 9: anduril.entitymanager.v1.TrackedBy.last_measurement_timestamp:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_anduril_entitymanager_v1_relationship_pub_proto_init() }
func file_anduril_entitymanager_v1_relationship_pub_proto_init() {
	if File_anduril_entitymanager_v1_relationship_pub_proto != nil {
		return
	}
	file_anduril_entitymanager_v1_sensors_pub_proto_init()
	file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes[2].OneofWrappers = []any{
		(*RelationshipType_Tether)(nil),
		(*RelationshipType_TrackedBy)(nil),
		(*RelationshipType_Configure)(nil),
		(*RelationshipType_GroupChild)(nil),
		(*RelationshipType_GroupParent)(nil),
		(*RelationshipType_MergedFrom)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_anduril_entitymanager_v1_relationship_pub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_entitymanager_v1_relationship_pub_proto_goTypes,
		DependencyIndexes: file_anduril_entitymanager_v1_relationship_pub_proto_depIdxs,
		MessageInfos:      file_anduril_entitymanager_v1_relationship_pub_proto_msgTypes,
	}.Build()
	File_anduril_entitymanager_v1_relationship_pub_proto = out.File
	file_anduril_entitymanager_v1_relationship_pub_proto_rawDesc = nil
	file_anduril_entitymanager_v1_relationship_pub_proto_goTypes = nil
	file_anduril_entitymanager_v1_relationship_pub_proto_depIdxs = nil
}
