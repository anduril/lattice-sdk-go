// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: anduril/entitymanager/v1/classification.pub.proto

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

// An enumeration of security classification levels.
type ClassificationLevels int32

const (
	ClassificationLevels_CLASSIFICATION_LEVELS_INVALID                 ClassificationLevels = 0
	ClassificationLevels_CLASSIFICATION_LEVELS_UNCLASSIFIED            ClassificationLevels = 1
	ClassificationLevels_CLASSIFICATION_LEVELS_CONTROLLED_UNCLASSIFIED ClassificationLevels = 2
	ClassificationLevels_CLASSIFICATION_LEVELS_CONFIDENTIAL            ClassificationLevels = 3
	ClassificationLevels_CLASSIFICATION_LEVELS_SECRET                  ClassificationLevels = 4
	ClassificationLevels_CLASSIFICATION_LEVELS_TOP_SECRET              ClassificationLevels = 5
)

// Enum value maps for ClassificationLevels.
var (
	ClassificationLevels_name = map[int32]string{
		0: "CLASSIFICATION_LEVELS_INVALID",
		1: "CLASSIFICATION_LEVELS_UNCLASSIFIED",
		2: "CLASSIFICATION_LEVELS_CONTROLLED_UNCLASSIFIED",
		3: "CLASSIFICATION_LEVELS_CONFIDENTIAL",
		4: "CLASSIFICATION_LEVELS_SECRET",
		5: "CLASSIFICATION_LEVELS_TOP_SECRET",
	}
	ClassificationLevels_value = map[string]int32{
		"CLASSIFICATION_LEVELS_INVALID":                 0,
		"CLASSIFICATION_LEVELS_UNCLASSIFIED":            1,
		"CLASSIFICATION_LEVELS_CONTROLLED_UNCLASSIFIED": 2,
		"CLASSIFICATION_LEVELS_CONFIDENTIAL":            3,
		"CLASSIFICATION_LEVELS_SECRET":                  4,
		"CLASSIFICATION_LEVELS_TOP_SECRET":              5,
	}
)

func (x ClassificationLevels) Enum() *ClassificationLevels {
	p := new(ClassificationLevels)
	*p = x
	return p
}

func (x ClassificationLevels) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClassificationLevels) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_classification_pub_proto_enumTypes[0].Descriptor()
}

func (ClassificationLevels) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_classification_pub_proto_enumTypes[0]
}

func (x ClassificationLevels) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClassificationLevels.Descriptor instead.
func (ClassificationLevels) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_classification_pub_proto_rawDescGZIP(), []int{0}
}

// A component that describes an entity's security classification levels.
type Classification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Marked as deprecated in anduril/entitymanager/v1/classification.pub.proto.
	Level ClassificationLevels `protobuf:"varint,1,opt,name=level,proto3,enum=anduril.entitymanager.v1.ClassificationLevels" json:"level,omitempty"`
	// The default classification information which should be assumed to apply to everything in
	// the entity unless a specific field level classification is present.
	Default *ClassificationInformation `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
	// The set of individual field classification information which should always precedence
	// over the default classification information.
	Fields []*FieldClassificationInformation `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *Classification) Reset() {
	*x = Classification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entitymanager_v1_classification_pub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Classification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Classification) ProtoMessage() {}

func (x *Classification) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_classification_pub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Classification.ProtoReflect.Descriptor instead.
func (*Classification) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_classification_pub_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in anduril/entitymanager/v1/classification.pub.proto.
func (x *Classification) GetLevel() ClassificationLevels {
	if x != nil {
		return x.Level
	}
	return ClassificationLevels_CLASSIFICATION_LEVELS_INVALID
}

func (x *Classification) GetDefault() *ClassificationInformation {
	if x != nil {
		return x.Default
	}
	return nil
}

func (x *Classification) GetFields() []*FieldClassificationInformation {
	if x != nil {
		return x.Fields
	}
	return nil
}

// A field specific classification information definition.
type FieldClassificationInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Proto field path which is the string representation of a field.
	// > example: signal.bandwidth_hz would be bandwidth_hz in the signal component
	FieldPath string `protobuf:"bytes,1,opt,name=field_path,json=fieldPath,proto3" json:"field_path,omitempty"`
	// The information which makes up the field level classification marking.
	ClassificationInformation *ClassificationInformation `protobuf:"bytes,2,opt,name=classification_information,json=classificationInformation,proto3" json:"classification_information,omitempty"`
}

func (x *FieldClassificationInformation) Reset() {
	*x = FieldClassificationInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entitymanager_v1_classification_pub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldClassificationInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldClassificationInformation) ProtoMessage() {}

func (x *FieldClassificationInformation) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_classification_pub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldClassificationInformation.ProtoReflect.Descriptor instead.
func (*FieldClassificationInformation) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_classification_pub_proto_rawDescGZIP(), []int{1}
}

func (x *FieldClassificationInformation) GetFieldPath() string {
	if x != nil {
		return x.FieldPath
	}
	return ""
}

func (x *FieldClassificationInformation) GetClassificationInformation() *ClassificationInformation {
	if x != nil {
		return x.ClassificationInformation
	}
	return nil
}

// Represents all of the necessary information required to generate a summarized
// classification marking.
//
// > example: A summarized classification marking of "TOPSECRET//NOFORN//FISA"
//
//	would be defined as: { "level": 5, "caveats": [ "NOFORN, "FISA" ] }
type ClassificationInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Classification level to be applied to the information in question.
	Level ClassificationLevels `protobuf:"varint,1,opt,name=level,proto3,enum=anduril.entitymanager.v1.ClassificationLevels" json:"level,omitempty"`
	// Caveats that may further restrict how the information can be disseminated.
	Caveats []string `protobuf:"bytes,2,rep,name=caveats,proto3" json:"caveats,omitempty"`
}

func (x *ClassificationInformation) Reset() {
	*x = ClassificationInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entitymanager_v1_classification_pub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassificationInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassificationInformation) ProtoMessage() {}

func (x *ClassificationInformation) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_classification_pub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassificationInformation.ProtoReflect.Descriptor instead.
func (*ClassificationInformation) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_classification_pub_proto_rawDescGZIP(), []int{2}
}

func (x *ClassificationInformation) GetLevel() ClassificationLevels {
	if x != nil {
		return x.Level
	}
	return ClassificationLevels_CLASSIFICATION_LEVELS_INVALID
}

func (x *ClassificationInformation) GetCaveats() []string {
	if x != nil {
		return x.Caveats
	}
	return nil
}

var File_anduril_entitymanager_v1_classification_pub_proto protoreflect.FileDescriptor

var file_anduril_entitymanager_v1_classification_pub_proto_rawDesc = []byte{
	0x0a, 0x31, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0xfb, 0x01,
	0x0a, 0x0e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x48, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2e, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x4d, 0x0a, 0x07, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x61, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x50, 0x0a, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x61, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0xb3, 0x01, 0x0a, 0x1e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x68, 0x12, 0x72, 0x0a,
	0x1a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x33, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x19, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x7b, 0x0a, 0x19, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e,
	0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x76, 0x65, 0x61, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x76, 0x65, 0x61, 0x74, 0x73, 0x2a, 0x84,
	0x02, 0x0a, 0x14, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4c, 0x41, 0x53, 0x53,
	0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x53,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x22, 0x43, 0x4c,
	0x41, 0x53, 0x53, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x56,
	0x45, 0x4c, 0x53, 0x5f, 0x55, 0x4e, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x31, 0x0a, 0x2d, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49, 0x46, 0x49, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x53, 0x5f, 0x43, 0x4f, 0x4e, 0x54,
	0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x55, 0x4e, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x02, 0x12, 0x26, 0x0a, 0x22, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49, 0x46,
	0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x53, 0x5f, 0x43,
	0x4f, 0x4e, 0x46, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x20, 0x0a,
	0x1c, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x53, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x10, 0x04, 0x12,
	0x24, 0x0a, 0x20, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x53, 0x5f, 0x54, 0x4f, 0x50, 0x5f, 0x53, 0x45, 0x43,
	0x52, 0x45, 0x54, 0x10, 0x05, 0x42, 0x8e, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67,
	0x75, 0x6e, 0x2d, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72,
	0x69, 0x6c, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x6e,
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
	file_anduril_entitymanager_v1_classification_pub_proto_rawDescOnce sync.Once
	file_anduril_entitymanager_v1_classification_pub_proto_rawDescData = file_anduril_entitymanager_v1_classification_pub_proto_rawDesc
)

func file_anduril_entitymanager_v1_classification_pub_proto_rawDescGZIP() []byte {
	file_anduril_entitymanager_v1_classification_pub_proto_rawDescOnce.Do(func() {
		file_anduril_entitymanager_v1_classification_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_entitymanager_v1_classification_pub_proto_rawDescData)
	})
	return file_anduril_entitymanager_v1_classification_pub_proto_rawDescData
}

var file_anduril_entitymanager_v1_classification_pub_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_anduril_entitymanager_v1_classification_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_anduril_entitymanager_v1_classification_pub_proto_goTypes = []any{
	(ClassificationLevels)(0),              // 0: anduril.entitymanager.v1.ClassificationLevels
	(*Classification)(nil),                 // 1: anduril.entitymanager.v1.Classification
	(*FieldClassificationInformation)(nil), // 2: anduril.entitymanager.v1.FieldClassificationInformation
	(*ClassificationInformation)(nil),      // 3: anduril.entitymanager.v1.ClassificationInformation
}
var file_anduril_entitymanager_v1_classification_pub_proto_depIdxs = []int32{
	0, // 0: anduril.entitymanager.v1.Classification.level:type_name -> anduril.entitymanager.v1.ClassificationLevels
	3, // 1: anduril.entitymanager.v1.Classification.default:type_name -> anduril.entitymanager.v1.ClassificationInformation
	2, // 2: anduril.entitymanager.v1.Classification.fields:type_name -> anduril.entitymanager.v1.FieldClassificationInformation
	3, // 3: anduril.entitymanager.v1.FieldClassificationInformation.classification_information:type_name -> anduril.entitymanager.v1.ClassificationInformation
	0, // 4: anduril.entitymanager.v1.ClassificationInformation.level:type_name -> anduril.entitymanager.v1.ClassificationLevels
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_anduril_entitymanager_v1_classification_pub_proto_init() }
func file_anduril_entitymanager_v1_classification_pub_proto_init() {
	if File_anduril_entitymanager_v1_classification_pub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_anduril_entitymanager_v1_classification_pub_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Classification); i {
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
		file_anduril_entitymanager_v1_classification_pub_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FieldClassificationInformation); i {
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
		file_anduril_entitymanager_v1_classification_pub_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ClassificationInformation); i {
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
			RawDescriptor: file_anduril_entitymanager_v1_classification_pub_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_entitymanager_v1_classification_pub_proto_goTypes,
		DependencyIndexes: file_anduril_entitymanager_v1_classification_pub_proto_depIdxs,
		EnumInfos:         file_anduril_entitymanager_v1_classification_pub_proto_enumTypes,
		MessageInfos:      file_anduril_entitymanager_v1_classification_pub_proto_msgTypes,
	}.Build()
	File_anduril_entitymanager_v1_classification_pub_proto = out.File
	file_anduril_entitymanager_v1_classification_pub_proto_rawDesc = nil
	file_anduril_entitymanager_v1_classification_pub_proto_goTypes = nil
	file_anduril_entitymanager_v1_classification_pub_proto_depIdxs = nil
}
