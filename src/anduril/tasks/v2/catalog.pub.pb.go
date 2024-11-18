// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: anduril/tasks/v2/catalog.pub.proto

package tasksv2

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

// Catalog of supported tasks.
type TaskCatalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskDefinitions []*TaskDefinition `protobuf:"bytes,1,rep,name=task_definitions,json=taskDefinitions,proto3" json:"task_definitions,omitempty"`
	// Asset is inhibited by VCE.
	// Asset can still receive tasks but not be able to act on them until inhibition status is lifted.
	IsAssetInhibited bool `protobuf:"varint,2,opt,name=is_asset_inhibited,json=isAssetInhibited,proto3" json:"is_asset_inhibited,omitempty"`
}

func (x *TaskCatalog) Reset() {
	*x = TaskCatalog{}
	mi := &file_anduril_tasks_v2_catalog_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskCatalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCatalog) ProtoMessage() {}

func (x *TaskCatalog) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_v2_catalog_pub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCatalog.ProtoReflect.Descriptor instead.
func (*TaskCatalog) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_v2_catalog_pub_proto_rawDescGZIP(), []int{0}
}

func (x *TaskCatalog) GetTaskDefinitions() []*TaskDefinition {
	if x != nil {
		return x.TaskDefinitions
	}
	return nil
}

func (x *TaskCatalog) GetIsAssetInhibited() bool {
	if x != nil {
		return x.IsAssetInhibited
	}
	return false
}

// Defines a supported task by the task specification URL of its "Any" type.
type TaskDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskSpecificationUrl string `protobuf:"bytes,1,opt,name=task_specification_url,json=taskSpecificationUrl,proto3" json:"task_specification_url,omitempty"`
	DisplayName          string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *TaskDefinition) Reset() {
	*x = TaskDefinition{}
	mi := &file_anduril_tasks_v2_catalog_pub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDefinition) ProtoMessage() {}

func (x *TaskDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_v2_catalog_pub_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDefinition.ProtoReflect.Descriptor instead.
func (*TaskDefinition) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_v2_catalog_pub_proto_rawDescGZIP(), []int{1}
}

func (x *TaskDefinition) GetTaskSpecificationUrl() string {
	if x != nil {
		return x.TaskSpecificationUrl
	}
	return ""
}

func (x *TaskDefinition) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

var File_anduril_tasks_v2_catalog_pub_proto protoreflect.FileDescriptor

var file_anduril_tasks_v2_catalog_pub_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f,
	0x76, 0x32, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x22, 0x88, 0x01, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x4b, 0x0a, 0x10, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f,
	0x69, 0x6e, 0x68, 0x69, 0x62, 0x69, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x10, 0x69, 0x73, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x68, 0x69, 0x62, 0x69, 0x74, 0x65,
	0x64, 0x22, 0x69, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x16, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0xc9, 0x01, 0x0a,
	0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2e, 0x76, 0x32, 0x42, 0x0f, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x50, 0x75,
	0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x6c, 0x61, 0x74,
	0x74, 0x69, 0x63, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x32,
	0x3b, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x41, 0x54, 0x58, 0xaa, 0x02,
	0x10, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x56,
	0x32, 0xca, 0x02, 0x10, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x1c, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x3a, 0x3a, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_anduril_tasks_v2_catalog_pub_proto_rawDescOnce sync.Once
	file_anduril_tasks_v2_catalog_pub_proto_rawDescData = file_anduril_tasks_v2_catalog_pub_proto_rawDesc
)

func file_anduril_tasks_v2_catalog_pub_proto_rawDescGZIP() []byte {
	file_anduril_tasks_v2_catalog_pub_proto_rawDescOnce.Do(func() {
		file_anduril_tasks_v2_catalog_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_tasks_v2_catalog_pub_proto_rawDescData)
	})
	return file_anduril_tasks_v2_catalog_pub_proto_rawDescData
}

var file_anduril_tasks_v2_catalog_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_anduril_tasks_v2_catalog_pub_proto_goTypes = []any{
	(*TaskCatalog)(nil),    // 0: anduril.tasks.v2.TaskCatalog
	(*TaskDefinition)(nil), // 1: anduril.tasks.v2.TaskDefinition
}
var file_anduril_tasks_v2_catalog_pub_proto_depIdxs = []int32{
	1, // 0: anduril.tasks.v2.TaskCatalog.task_definitions:type_name -> anduril.tasks.v2.TaskDefinition
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_anduril_tasks_v2_catalog_pub_proto_init() }
func file_anduril_tasks_v2_catalog_pub_proto_init() {
	if File_anduril_tasks_v2_catalog_pub_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_anduril_tasks_v2_catalog_pub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_tasks_v2_catalog_pub_proto_goTypes,
		DependencyIndexes: file_anduril_tasks_v2_catalog_pub_proto_depIdxs,
		MessageInfos:      file_anduril_tasks_v2_catalog_pub_proto_msgTypes,
	}.Build()
	File_anduril_tasks_v2_catalog_pub_proto = out.File
	file_anduril_tasks_v2_catalog_pub_proto_rawDesc = nil
	file_anduril_tasks_v2_catalog_pub_proto_goTypes = nil
	file_anduril_tasks_v2_catalog_pub_proto_depIdxs = nil
}
