// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: anduril/entitymanager/v1/ontology.pub.proto

package entitymanagerv1

import (
	v1 "github.com/anduril/lattice-sdk-go/src/anduril/ontology/v1"
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

// Provides the disposition, environment, and nationality of an Entity.
type MilView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Disposition v1.Disposition `protobuf:"varint,1,opt,name=disposition,proto3,enum=anduril.ontology.v1.Disposition" json:"disposition,omitempty"`
	Environment v1.Environment `protobuf:"varint,2,opt,name=environment,proto3,enum=anduril.ontology.v1.Environment" json:"environment,omitempty"`
	Nationality v1.Nationality `protobuf:"varint,3,opt,name=nationality,proto3,enum=anduril.ontology.v1.Nationality" json:"nationality,omitempty"`
}

func (x *MilView) Reset() {
	*x = MilView{}
	mi := &file_anduril_entitymanager_v1_ontology_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MilView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MilView) ProtoMessage() {}

func (x *MilView) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_ontology_pub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MilView.ProtoReflect.Descriptor instead.
func (*MilView) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_ontology_pub_proto_rawDescGZIP(), []int{0}
}

func (x *MilView) GetDisposition() v1.Disposition {
	if x != nil {
		return x.Disposition
	}
	return v1.Disposition(0)
}

func (x *MilView) GetEnvironment() v1.Environment {
	if x != nil {
		return x.Environment
	}
	return v1.Environment(0)
}

func (x *MilView) GetNationality() v1.Nationality {
	if x != nil {
		return x.Nationality
	}
	return v1.Nationality(0)
}

// Ontology of the entity.
type Ontology struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A string that describes the entity's high-level type with natural language.
	PlatformType string `protobuf:"bytes,3,opt,name=platform_type,json=platformType,proto3" json:"platform_type,omitempty"`
	// A string that describes the entity's exact model or type.
	SpecificType string `protobuf:"bytes,4,opt,name=specific_type,json=specificType,proto3" json:"specific_type,omitempty"`
	// The template used when creating this entity. Specifies minimum required components.
	Template Template `protobuf:"varint,2,opt,name=template,proto3,enum=anduril.entitymanager.v1.Template" json:"template,omitempty"`
}

func (x *Ontology) Reset() {
	*x = Ontology{}
	mi := &file_anduril_entitymanager_v1_ontology_pub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ontology) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ontology) ProtoMessage() {}

func (x *Ontology) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_ontology_pub_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ontology.ProtoReflect.Descriptor instead.
func (*Ontology) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_ontology_pub_proto_rawDescGZIP(), []int{1}
}

func (x *Ontology) GetPlatformType() string {
	if x != nil {
		return x.PlatformType
	}
	return ""
}

func (x *Ontology) GetSpecificType() string {
	if x != nil {
		return x.SpecificType
	}
	return ""
}

func (x *Ontology) GetTemplate() Template {
	if x != nil {
		return x.Template
	}
	return Template_TEMPLATE_INVALID
}

var File_anduril_entitymanager_v1_ontology_pub_proto protoreflect.FileDescriptor

var file_anduril_entitymanager_v1_ontology_pub_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x6e, 0x74, 0x6f, 0x6c,
	0x6f, 0x67, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x61,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x6f, 0x6e, 0x74, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x07, 0x4d, 0x69, 0x6c, 0x56, 0x69, 0x65, 0x77, 0x12, 0x47, 0x0a,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x6f, 0x6e, 0x74,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xc8, 0x3e, 0x01, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x6f, 0x6e, 0x74, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x03, 0xc8,
	0x3e, 0x01, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x47, 0x0a, 0x0b, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x6f,
	0x6e, 0x74, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x03, 0xc8, 0x3e, 0x01, 0x52, 0x0b, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x22, 0xb1, 0x01, 0x0a, 0x08, 0x4f, 0x6e, 0x74,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x28, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc8, 0x3e,
	0x01, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x28, 0x0a, 0x0d, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc8, 0x3e, 0x01, 0x52, 0x0c, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x61, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x42, 0x82, 0x02, 0x0a,
	0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x4f,
	0x6e, 0x74, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x50, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x6c, 0x61, 0x74, 0x74, 0x69, 0x63, 0x65, 0x2d, 0x73, 0x64,
	0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x41, 0x45, 0x58, 0xaa, 0x02, 0x18, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x18, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24,
	0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x3a, 0x3a,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_anduril_entitymanager_v1_ontology_pub_proto_rawDescOnce sync.Once
	file_anduril_entitymanager_v1_ontology_pub_proto_rawDescData = file_anduril_entitymanager_v1_ontology_pub_proto_rawDesc
)

func file_anduril_entitymanager_v1_ontology_pub_proto_rawDescGZIP() []byte {
	file_anduril_entitymanager_v1_ontology_pub_proto_rawDescOnce.Do(func() {
		file_anduril_entitymanager_v1_ontology_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_entitymanager_v1_ontology_pub_proto_rawDescData)
	})
	return file_anduril_entitymanager_v1_ontology_pub_proto_rawDescData
}

var file_anduril_entitymanager_v1_ontology_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_anduril_entitymanager_v1_ontology_pub_proto_goTypes = []any{
	(*MilView)(nil),     // 0: anduril.entitymanager.v1.MilView
	(*Ontology)(nil),    // 1: anduril.entitymanager.v1.Ontology
	(v1.Disposition)(0), // 2: anduril.ontology.v1.Disposition
	(v1.Environment)(0), // 3: anduril.ontology.v1.Environment
	(v1.Nationality)(0), // 4: anduril.ontology.v1.Nationality
	(Template)(0),       // 5: anduril.entitymanager.v1.Template
}
var file_anduril_entitymanager_v1_ontology_pub_proto_depIdxs = []int32{
	2, // 0: anduril.entitymanager.v1.MilView.disposition:type_name -> anduril.ontology.v1.Disposition
	3, // 1: anduril.entitymanager.v1.MilView.environment:type_name -> anduril.ontology.v1.Environment
	4, // 2: anduril.entitymanager.v1.MilView.nationality:type_name -> anduril.ontology.v1.Nationality
	5, // 3: anduril.entitymanager.v1.Ontology.template:type_name -> anduril.entitymanager.v1.Template
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_anduril_entitymanager_v1_ontology_pub_proto_init() }
func file_anduril_entitymanager_v1_ontology_pub_proto_init() {
	if File_anduril_entitymanager_v1_ontology_pub_proto != nil {
		return
	}
	file_anduril_entitymanager_v1_options_pub_proto_init()
	file_anduril_entitymanager_v1_types_pub_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_anduril_entitymanager_v1_ontology_pub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_entitymanager_v1_ontology_pub_proto_goTypes,
		DependencyIndexes: file_anduril_entitymanager_v1_ontology_pub_proto_depIdxs,
		MessageInfos:      file_anduril_entitymanager_v1_ontology_pub_proto_msgTypes,
	}.Build()
	File_anduril_entitymanager_v1_ontology_pub_proto = out.File
	file_anduril_entitymanager_v1_ontology_pub_proto_rawDesc = nil
	file_anduril_entitymanager_v1_ontology_pub_proto_goTypes = nil
	file_anduril_entitymanager_v1_ontology_pub_proto_depIdxs = nil
}
