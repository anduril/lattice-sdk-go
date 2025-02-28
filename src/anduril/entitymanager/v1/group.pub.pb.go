// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: anduril/entitymanager/v1/group.pub.proto

package entitymanagerv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Military units defined by the Army.
type ArmyEchelon int32

const (
	ArmyEchelon_ARMY_ECHELON_INVALID   ArmyEchelon = 0
	ArmyEchelon_ARMY_ECHELON_FIRE_TEAM ArmyEchelon = 1  // Smallest unit group, e.g., a few soldiers
	ArmyEchelon_ARMY_ECHELON_SQUAD     ArmyEchelon = 2  // E.g., a group of fire teams
	ArmyEchelon_ARMY_ECHELON_PLATOON   ArmyEchelon = 3  // E.g., several squads
	ArmyEchelon_ARMY_ECHELON_COMPANY   ArmyEchelon = 4  // E.g., several platoons
	ArmyEchelon_ARMY_ECHELON_BATTALION ArmyEchelon = 5  // E.g., several companies
	ArmyEchelon_ARMY_ECHELON_REGIMENT  ArmyEchelon = 6  // E.g., several battalions
	ArmyEchelon_ARMY_ECHELON_BRIGADE   ArmyEchelon = 7  // E.g., several regiments or battalions
	ArmyEchelon_ARMY_ECHELON_DIVISION  ArmyEchelon = 8  // E.g., several brigades
	ArmyEchelon_ARMY_ECHELON_CORPS     ArmyEchelon = 9  // E.g., several divisions
	ArmyEchelon_ARMY_ECHELON_ARMY      ArmyEchelon = 10 // E.g., several corps
)

// Enum value maps for ArmyEchelon.
var (
	ArmyEchelon_name = map[int32]string{
		0:  "ARMY_ECHELON_INVALID",
		1:  "ARMY_ECHELON_FIRE_TEAM",
		2:  "ARMY_ECHELON_SQUAD",
		3:  "ARMY_ECHELON_PLATOON",
		4:  "ARMY_ECHELON_COMPANY",
		5:  "ARMY_ECHELON_BATTALION",
		6:  "ARMY_ECHELON_REGIMENT",
		7:  "ARMY_ECHELON_BRIGADE",
		8:  "ARMY_ECHELON_DIVISION",
		9:  "ARMY_ECHELON_CORPS",
		10: "ARMY_ECHELON_ARMY",
	}
	ArmyEchelon_value = map[string]int32{
		"ARMY_ECHELON_INVALID":   0,
		"ARMY_ECHELON_FIRE_TEAM": 1,
		"ARMY_ECHELON_SQUAD":     2,
		"ARMY_ECHELON_PLATOON":   3,
		"ARMY_ECHELON_COMPANY":   4,
		"ARMY_ECHELON_BATTALION": 5,
		"ARMY_ECHELON_REGIMENT":  6,
		"ARMY_ECHELON_BRIGADE":   7,
		"ARMY_ECHELON_DIVISION":  8,
		"ARMY_ECHELON_CORPS":     9,
		"ARMY_ECHELON_ARMY":      10,
	}
)

func (x ArmyEchelon) Enum() *ArmyEchelon {
	p := new(ArmyEchelon)
	*p = x
	return p
}

func (x ArmyEchelon) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArmyEchelon) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_group_pub_proto_enumTypes[0].Descriptor()
}

func (ArmyEchelon) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_group_pub_proto_enumTypes[0]
}

func (x ArmyEchelon) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArmyEchelon.Descriptor instead.
func (ArmyEchelon) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_group_pub_proto_rawDescGZIP(), []int{0}
}

// Details related to grouping for this entity
type GroupDetails struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to GroupType:
	//
	//	*GroupDetails_Team
	//	*GroupDetails_Echelon
	GroupType     isGroupDetails_GroupType `protobuf_oneof:"group_type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GroupDetails) Reset() {
	*x = GroupDetails{}
	mi := &file_anduril_entitymanager_v1_group_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupDetails) ProtoMessage() {}

func (x *GroupDetails) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_group_pub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupDetails.ProtoReflect.Descriptor instead.
func (*GroupDetails) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_group_pub_proto_rawDescGZIP(), []int{0}
}

func (x *GroupDetails) GetGroupType() isGroupDetails_GroupType {
	if x != nil {
		return x.GroupType
	}
	return nil
}

func (x *GroupDetails) GetTeam() *Team {
	if x != nil {
		if x, ok := x.GroupType.(*GroupDetails_Team); ok {
			return x.Team
		}
	}
	return nil
}

func (x *GroupDetails) GetEchelon() *Echelon {
	if x != nil {
		if x, ok := x.GroupType.(*GroupDetails_Echelon); ok {
			return x.Echelon
		}
	}
	return nil
}

type isGroupDetails_GroupType interface {
	isGroupDetails_GroupType()
}

type GroupDetails_Team struct {
	Team *Team `protobuf:"bytes,1,opt,name=team,proto3,oneof"`
}

type GroupDetails_Echelon struct {
	Echelon *Echelon `protobuf:"bytes,3,opt,name=echelon,proto3,oneof"`
}

func (*GroupDetails_Team) isGroupDetails_GroupType() {}

func (*GroupDetails_Echelon) isGroupDetails_GroupType() {}

// Describes a Team group type. Comprised of autonomous entities where an entity
// in a Team can only be a part of a single Team at a time.
type Team struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Team) Reset() {
	*x = Team{}
	mi := &file_anduril_entitymanager_v1_group_pub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_group_pub_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_group_pub_proto_rawDescGZIP(), []int{1}
}

// Describes a Echelon group type.  Comprised of entities which are members of the
// same unit or echelon. Ex: A group of tanks within a armored company or that same company
// as a member of a battalion.
type Echelon struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to EchelonType:
	//
	//	*Echelon_ArmyEchelon
	EchelonType   isEchelon_EchelonType `protobuf_oneof:"echelon_type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Echelon) Reset() {
	*x = Echelon{}
	mi := &file_anduril_entitymanager_v1_group_pub_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Echelon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Echelon) ProtoMessage() {}

func (x *Echelon) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_group_pub_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Echelon.ProtoReflect.Descriptor instead.
func (*Echelon) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_group_pub_proto_rawDescGZIP(), []int{2}
}

func (x *Echelon) GetEchelonType() isEchelon_EchelonType {
	if x != nil {
		return x.EchelonType
	}
	return nil
}

func (x *Echelon) GetArmyEchelon() ArmyEchelon {
	if x != nil {
		if x, ok := x.EchelonType.(*Echelon_ArmyEchelon); ok {
			return x.ArmyEchelon
		}
	}
	return ArmyEchelon_ARMY_ECHELON_INVALID
}

type isEchelon_EchelonType interface {
	isEchelon_EchelonType()
}

type Echelon_ArmyEchelon struct {
	ArmyEchelon ArmyEchelon `protobuf:"varint,1,opt,name=army_echelon,json=armyEchelon,proto3,enum=anduril.entitymanager.v1.ArmyEchelon,oneof"`
}

func (*Echelon_ArmyEchelon) isEchelon_EchelonType() {}

var File_anduril_entitymanager_v1_group_pub_proto protoreflect.FileDescriptor

var file_anduril_entitymanager_v1_group_pub_proto_rawDesc = string([]byte{
	0x0a, 0x28, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x22, 0x91, 0x01, 0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x65, 0x61, 0x6d, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x3d, 0x0a, 0x07, 0x65,
	0x63, 0x68, 0x65, 0x6c, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x65, 0x6c, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x07, 0x65, 0x63, 0x68, 0x65, 0x6c, 0x6f, 0x6e, 0x42, 0x0c, 0x0a, 0x0a, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x06, 0x0a, 0x04, 0x54, 0x65, 0x61, 0x6d,
	0x22, 0x65, 0x0a, 0x07, 0x45, 0x63, 0x68, 0x65, 0x6c, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x0c, 0x61,
	0x72, 0x6d, 0x79, 0x5f, 0x65, 0x63, 0x68, 0x65, 0x6c, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x25, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x6d,
	0x79, 0x45, 0x63, 0x68, 0x65, 0x6c, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x72, 0x6d, 0x79,
	0x45, 0x63, 0x68, 0x65, 0x6c, 0x6f, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x65, 0x63, 0x68, 0x65, 0x6c,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2a, 0xaa, 0x02, 0x0a, 0x0b, 0x41, 0x72, 0x6d, 0x79,
	0x45, 0x63, 0x68, 0x65, 0x6c, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x52, 0x4d, 0x59, 0x5f,
	0x45, 0x43, 0x48, 0x45, 0x4c, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x52, 0x4d, 0x59, 0x5f, 0x45, 0x43, 0x48, 0x45, 0x4c, 0x4f,
	0x4e, 0x5f, 0x46, 0x49, 0x52, 0x45, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x10, 0x01, 0x12, 0x16, 0x0a,
	0x12, 0x41, 0x52, 0x4d, 0x59, 0x5f, 0x45, 0x43, 0x48, 0x45, 0x4c, 0x4f, 0x4e, 0x5f, 0x53, 0x51,
	0x55, 0x41, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x52, 0x4d, 0x59, 0x5f, 0x45, 0x43,
	0x48, 0x45, 0x4c, 0x4f, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x4f, 0x4f, 0x4e, 0x10, 0x03, 0x12,
	0x18, 0x0a, 0x14, 0x41, 0x52, 0x4d, 0x59, 0x5f, 0x45, 0x43, 0x48, 0x45, 0x4c, 0x4f, 0x4e, 0x5f,
	0x43, 0x4f, 0x4d, 0x50, 0x41, 0x4e, 0x59, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x52, 0x4d,
	0x59, 0x5f, 0x45, 0x43, 0x48, 0x45, 0x4c, 0x4f, 0x4e, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x41, 0x4c,
	0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x52, 0x4d, 0x59, 0x5f, 0x45, 0x43,
	0x48, 0x45, 0x4c, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x06,
	0x12, 0x18, 0x0a, 0x14, 0x41, 0x52, 0x4d, 0x59, 0x5f, 0x45, 0x43, 0x48, 0x45, 0x4c, 0x4f, 0x4e,
	0x5f, 0x42, 0x52, 0x49, 0x47, 0x41, 0x44, 0x45, 0x10, 0x07, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x52,
	0x4d, 0x59, 0x5f, 0x45, 0x43, 0x48, 0x45, 0x4c, 0x4f, 0x4e, 0x5f, 0x44, 0x49, 0x56, 0x49, 0x53,
	0x49, 0x4f, 0x4e, 0x10, 0x08, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x52, 0x4d, 0x59, 0x5f, 0x45, 0x43,
	0x48, 0x45, 0x4c, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x52, 0x50, 0x53, 0x10, 0x09, 0x12, 0x15, 0x0a,
	0x11, 0x41, 0x52, 0x4d, 0x59, 0x5f, 0x45, 0x43, 0x48, 0x45, 0x4c, 0x4f, 0x4e, 0x5f, 0x41, 0x52,
	0x4d, 0x59, 0x10, 0x0a, 0x42, 0xff, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x64,
	0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x75, 0x62, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x6c, 0x61, 0x74, 0x74, 0x69,
	0x63, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x6e,
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
})

var (
	file_anduril_entitymanager_v1_group_pub_proto_rawDescOnce sync.Once
	file_anduril_entitymanager_v1_group_pub_proto_rawDescData []byte
)

func file_anduril_entitymanager_v1_group_pub_proto_rawDescGZIP() []byte {
	file_anduril_entitymanager_v1_group_pub_proto_rawDescOnce.Do(func() {
		file_anduril_entitymanager_v1_group_pub_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_anduril_entitymanager_v1_group_pub_proto_rawDesc), len(file_anduril_entitymanager_v1_group_pub_proto_rawDesc)))
	})
	return file_anduril_entitymanager_v1_group_pub_proto_rawDescData
}

var file_anduril_entitymanager_v1_group_pub_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_anduril_entitymanager_v1_group_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_anduril_entitymanager_v1_group_pub_proto_goTypes = []any{
	(ArmyEchelon)(0),     // 0: anduril.entitymanager.v1.ArmyEchelon
	(*GroupDetails)(nil), // 1: anduril.entitymanager.v1.GroupDetails
	(*Team)(nil),         // 2: anduril.entitymanager.v1.Team
	(*Echelon)(nil),      // 3: anduril.entitymanager.v1.Echelon
}
var file_anduril_entitymanager_v1_group_pub_proto_depIdxs = []int32{
	2, // 0: anduril.entitymanager.v1.GroupDetails.team:type_name -> anduril.entitymanager.v1.Team
	3, // 1: anduril.entitymanager.v1.GroupDetails.echelon:type_name -> anduril.entitymanager.v1.Echelon
	0, // 2: anduril.entitymanager.v1.Echelon.army_echelon:type_name -> anduril.entitymanager.v1.ArmyEchelon
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_anduril_entitymanager_v1_group_pub_proto_init() }
func file_anduril_entitymanager_v1_group_pub_proto_init() {
	if File_anduril_entitymanager_v1_group_pub_proto != nil {
		return
	}
	file_anduril_entitymanager_v1_group_pub_proto_msgTypes[0].OneofWrappers = []any{
		(*GroupDetails_Team)(nil),
		(*GroupDetails_Echelon)(nil),
	}
	file_anduril_entitymanager_v1_group_pub_proto_msgTypes[2].OneofWrappers = []any{
		(*Echelon_ArmyEchelon)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_anduril_entitymanager_v1_group_pub_proto_rawDesc), len(file_anduril_entitymanager_v1_group_pub_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_entitymanager_v1_group_pub_proto_goTypes,
		DependencyIndexes: file_anduril_entitymanager_v1_group_pub_proto_depIdxs,
		EnumInfos:         file_anduril_entitymanager_v1_group_pub_proto_enumTypes,
		MessageInfos:      file_anduril_entitymanager_v1_group_pub_proto_msgTypes,
	}.Build()
	File_anduril_entitymanager_v1_group_pub_proto = out.File
	file_anduril_entitymanager_v1_group_pub_proto_goTypes = nil
	file_anduril_entitymanager_v1_group_pub_proto_depIdxs = nil
}
