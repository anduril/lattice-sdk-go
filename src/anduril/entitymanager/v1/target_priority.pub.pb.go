// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: anduril/entitymanager/v1/target_priority.pub.proto

package entitymanagerv1

import (
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

// The target prioritization associated with an entity.
type TargetPriority struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Describes the target priority in relation to high value target lists.
	HighValueTarget *HighValueTarget `protobuf:"bytes,1,opt,name=high_value_target,json=highValueTarget,proto3" json:"high_value_target,omitempty"`
	// Describes whether the entity should be treated as a threat
	Threat *Threat `protobuf:"bytes,2,opt,name=threat,proto3" json:"threat,omitempty"`
}

func (x *TargetPriority) Reset() {
	*x = TargetPriority{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entitymanager_v1_target_priority_pub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetPriority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetPriority) ProtoMessage() {}

func (x *TargetPriority) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_target_priority_pub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetPriority.ProtoReflect.Descriptor instead.
func (*TargetPriority) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_target_priority_pub_proto_rawDescGZIP(), []int{0}
}

func (x *TargetPriority) GetHighValueTarget() *HighValueTarget {
	if x != nil {
		return x.HighValueTarget
	}
	return nil
}

func (x *TargetPriority) GetThreat() *Threat {
	if x != nil {
		return x.Threat
	}
	return nil
}

// Describes whether something is a high value target or not.
type HighValueTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates whether the target matches any description from a high value target list.
	IsHighValueTarget bool `protobuf:"varint,1,opt,name=is_high_value_target,json=isHighValueTarget,proto3" json:"is_high_value_target,omitempty"`
	// The priority associated with the target. If the target's description appears on multiple high value target lists,
	// the priority will be a reflection of the highest priority of all of those list's target description.
	//
	// A lower value indicates the target is of a higher priority, with 1 being the highest possible priority. A value of
	// 0 indicates there is no priority associated with this target.
	TargetPriority uint32 `protobuf:"varint,2,opt,name=target_priority,json=targetPriority,proto3" json:"target_priority,omitempty"`
	// All of the high value target descriptions that the target matches against.
	TargetMatches []*HighValueTargetMatch `protobuf:"bytes,3,rep,name=target_matches,json=targetMatches,proto3" json:"target_matches,omitempty"`
	// Indicates whether the target is a 'High Payoff Target'. Targets can be one or both of high value and high payoff.
	// Semantically a High Value Target characterizes the target's importance to Red, whereas a High Payoff Target
	// indicates prosecuting the target furthers Blue's specific objectives.
	IsHighPayoffTarget bool `protobuf:"varint,4,opt,name=is_high_payoff_target,json=isHighPayoffTarget,proto3" json:"is_high_payoff_target,omitempty"`
}

func (x *HighValueTarget) Reset() {
	*x = HighValueTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entitymanager_v1_target_priority_pub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HighValueTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HighValueTarget) ProtoMessage() {}

func (x *HighValueTarget) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_target_priority_pub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HighValueTarget.ProtoReflect.Descriptor instead.
func (*HighValueTarget) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_target_priority_pub_proto_rawDescGZIP(), []int{1}
}

func (x *HighValueTarget) GetIsHighValueTarget() bool {
	if x != nil {
		return x.IsHighValueTarget
	}
	return false
}

func (x *HighValueTarget) GetTargetPriority() uint32 {
	if x != nil {
		return x.TargetPriority
	}
	return 0
}

func (x *HighValueTarget) GetTargetMatches() []*HighValueTargetMatch {
	if x != nil {
		return x.TargetMatches
	}
	return nil
}

func (x *HighValueTarget) GetIsHighPayoffTarget() bool {
	if x != nil {
		return x.IsHighPayoffTarget
	}
	return false
}

type HighValueTargetMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the high value target list that matches the target description.
	HighValueTargetListId string `protobuf:"bytes,1,opt,name=high_value_target_list_id,json=highValueTargetListId,proto3" json:"high_value_target_list_id,omitempty"`
	// The ID of the specific high value target description within a high value target list that was matched against.
	// The ID is considered to be a globally unique identifier across all high value target IDs.
	HighValueTargetDescriptionId string `protobuf:"bytes,2,opt,name=high_value_target_description_id,json=highValueTargetDescriptionId,proto3" json:"high_value_target_description_id,omitempty"`
}

func (x *HighValueTargetMatch) Reset() {
	*x = HighValueTargetMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entitymanager_v1_target_priority_pub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HighValueTargetMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HighValueTargetMatch) ProtoMessage() {}

func (x *HighValueTargetMatch) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_target_priority_pub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HighValueTargetMatch.ProtoReflect.Descriptor instead.
func (*HighValueTargetMatch) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_target_priority_pub_proto_rawDescGZIP(), []int{2}
}

func (x *HighValueTargetMatch) GetHighValueTargetListId() string {
	if x != nil {
		return x.HighValueTargetListId
	}
	return ""
}

func (x *HighValueTargetMatch) GetHighValueTargetDescriptionId() string {
	if x != nil {
		return x.HighValueTargetDescriptionId
	}
	return ""
}

// Describes whether an entity is a threat or not.
type Threat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates that the entity has been determined to be a threat.
	IsThreat bool `protobuf:"varint,1,opt,name=is_threat,json=isThreat,proto3" json:"is_threat,omitempty"`
}

func (x *Threat) Reset() {
	*x = Threat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entitymanager_v1_target_priority_pub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Threat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Threat) ProtoMessage() {}

func (x *Threat) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_target_priority_pub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Threat.ProtoReflect.Descriptor instead.
func (*Threat) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_target_priority_pub_proto_rawDescGZIP(), []int{3}
}

func (x *Threat) GetIsThreat() bool {
	if x != nil {
		return x.IsThreat
	}
	return false
}

var File_anduril_entitymanager_v1_target_priority_pub_proto protoreflect.FileDescriptor

var file_anduril_entitymanager_v1_target_priority_pub_proto_rawDesc = []byte{
	0x0a, 0x32, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x2a,
	0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x0e, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x5a, 0x0a,
	0x11, 0x68, 0x69, 0x67, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72,
	0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x69, 0x67, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x42, 0x03, 0xc8, 0x3e, 0x01, 0x52, 0x0f, 0x68, 0x69, 0x67, 0x68, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x74, 0x68, 0x72,
	0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x52, 0x06, 0x74, 0x68, 0x72,
	0x65, 0x61, 0x74, 0x22, 0xf5, 0x01, 0x0a, 0x0f, 0x48, 0x69, 0x67, 0x68, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x2f, 0x0a, 0x14, 0x69, 0x73, 0x5f, 0x68, 0x69,
	0x67, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x73, 0x48, 0x69, 0x67, 0x68, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x55, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x69, 0x67, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x15, 0x69, 0x73, 0x5f, 0x68,
	0x69, 0x67, 0x68, 0x5f, 0x70, 0x61, 0x79, 0x6f, 0x66, 0x66, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x73, 0x48, 0x69, 0x67, 0x68, 0x50,
	0x61, 0x79, 0x6f, 0x66, 0x66, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x98, 0x01, 0x0a, 0x14,
	0x48, 0x69, 0x67, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x38, 0x0a, 0x19, 0x68, 0x69, 0x67, 0x68, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x68, 0x69, 0x67, 0x68, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x46,
	0x0a, 0x20, 0x68, 0x69, 0x67, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1c, 0x68, 0x69, 0x67, 0x68, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x06, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x42, 0x84, 0x02,
	0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x16,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x50, 0x75,
	0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x61, 0x6e, 0x64,
	0x75, 0x72, 0x69, 0x6c, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x45, 0x58, 0xaa, 0x02, 0x18, 0x41, 0x6e, 0x64,
	0x75, 0x72, 0x69, 0x6c, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x24, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_anduril_entitymanager_v1_target_priority_pub_proto_rawDescOnce sync.Once
	file_anduril_entitymanager_v1_target_priority_pub_proto_rawDescData = file_anduril_entitymanager_v1_target_priority_pub_proto_rawDesc
)

func file_anduril_entitymanager_v1_target_priority_pub_proto_rawDescGZIP() []byte {
	file_anduril_entitymanager_v1_target_priority_pub_proto_rawDescOnce.Do(func() {
		file_anduril_entitymanager_v1_target_priority_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_entitymanager_v1_target_priority_pub_proto_rawDescData)
	})
	return file_anduril_entitymanager_v1_target_priority_pub_proto_rawDescData
}

var file_anduril_entitymanager_v1_target_priority_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_anduril_entitymanager_v1_target_priority_pub_proto_goTypes = []any{
	(*TargetPriority)(nil),       // 0: anduril.entitymanager.v1.TargetPriority
	(*HighValueTarget)(nil),      // 1: anduril.entitymanager.v1.HighValueTarget
	(*HighValueTargetMatch)(nil), // 2: anduril.entitymanager.v1.HighValueTargetMatch
	(*Threat)(nil),               // 3: anduril.entitymanager.v1.Threat
}
var file_anduril_entitymanager_v1_target_priority_pub_proto_depIdxs = []int32{
	1, // 0: anduril.entitymanager.v1.TargetPriority.high_value_target:type_name -> anduril.entitymanager.v1.HighValueTarget
	3, // 1: anduril.entitymanager.v1.TargetPriority.threat:type_name -> anduril.entitymanager.v1.Threat
	2, // 2: anduril.entitymanager.v1.HighValueTarget.target_matches:type_name -> anduril.entitymanager.v1.HighValueTargetMatch
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_anduril_entitymanager_v1_target_priority_pub_proto_init() }
func file_anduril_entitymanager_v1_target_priority_pub_proto_init() {
	if File_anduril_entitymanager_v1_target_priority_pub_proto != nil {
		return
	}
	file_anduril_entitymanager_v1_options_pub_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_anduril_entitymanager_v1_target_priority_pub_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TargetPriority); i {
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
		file_anduril_entitymanager_v1_target_priority_pub_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*HighValueTarget); i {
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
		file_anduril_entitymanager_v1_target_priority_pub_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*HighValueTargetMatch); i {
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
		file_anduril_entitymanager_v1_target_priority_pub_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Threat); i {
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
			RawDescriptor: file_anduril_entitymanager_v1_target_priority_pub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_entitymanager_v1_target_priority_pub_proto_goTypes,
		DependencyIndexes: file_anduril_entitymanager_v1_target_priority_pub_proto_depIdxs,
		MessageInfos:      file_anduril_entitymanager_v1_target_priority_pub_proto_msgTypes,
	}.Build()
	File_anduril_entitymanager_v1_target_priority_pub_proto = out.File
	file_anduril_entitymanager_v1_target_priority_pub_proto_rawDesc = nil
	file_anduril_entitymanager_v1_target_priority_pub_proto_goTypes = nil
	file_anduril_entitymanager_v1_target_priority_pub_proto_depIdxs = nil
}
