// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: anduril/taskmanager/v1/generic_spec.pub.proto

package taskmanagerv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

// GenericSpec is a wrapper for arbitrary JSON payloads. Meant for wrapping a Task's specification field, if needed.
type GenericSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Can represent any JSON value.
	Payload       *structpb.Value `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenericSpec) Reset() {
	*x = GenericSpec{}
	mi := &file_anduril_taskmanager_v1_generic_spec_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenericSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericSpec) ProtoMessage() {}

func (x *GenericSpec) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_taskmanager_v1_generic_spec_pub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericSpec.ProtoReflect.Descriptor instead.
func (*GenericSpec) Descriptor() ([]byte, []int) {
	return file_anduril_taskmanager_v1_generic_spec_pub_proto_rawDescGZIP(), []int{0}
}

func (x *GenericSpec) GetPayload() *structpb.Value {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_anduril_taskmanager_v1_generic_spec_pub_proto protoreflect.FileDescriptor

var file_anduril_taskmanager_v1_generic_spec_pub_proto_rawDesc = string([]byte{
	0x0a, 0x2d, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x30, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0xf7, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x53, 0x70,
	0x65, 0x63, 0x50, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c,
	0x2f, 0x6c, 0x61, 0x74, 0x74, 0x69, 0x63, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f,
	0x73, 0x72, 0x63, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x61, 0x73, 0x6b, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x54, 0x58, 0xaa, 0x02,
	0x16, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x16, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x5c, 0x54, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x22, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x54, 0x61, 0x73, 0x6b, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x3a,
	0x3a, 0x54, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_anduril_taskmanager_v1_generic_spec_pub_proto_rawDescOnce sync.Once
	file_anduril_taskmanager_v1_generic_spec_pub_proto_rawDescData []byte
)

func file_anduril_taskmanager_v1_generic_spec_pub_proto_rawDescGZIP() []byte {
	file_anduril_taskmanager_v1_generic_spec_pub_proto_rawDescOnce.Do(func() {
		file_anduril_taskmanager_v1_generic_spec_pub_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_anduril_taskmanager_v1_generic_spec_pub_proto_rawDesc), len(file_anduril_taskmanager_v1_generic_spec_pub_proto_rawDesc)))
	})
	return file_anduril_taskmanager_v1_generic_spec_pub_proto_rawDescData
}

var file_anduril_taskmanager_v1_generic_spec_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_anduril_taskmanager_v1_generic_spec_pub_proto_goTypes = []any{
	(*GenericSpec)(nil),    // 0: anduril.taskmanager.v1.GenericSpec
	(*structpb.Value)(nil), // 1: google.protobuf.Value
}
var file_anduril_taskmanager_v1_generic_spec_pub_proto_depIdxs = []int32{
	1, // 0: anduril.taskmanager.v1.GenericSpec.payload:type_name -> google.protobuf.Value
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_anduril_taskmanager_v1_generic_spec_pub_proto_init() }
func file_anduril_taskmanager_v1_generic_spec_pub_proto_init() {
	if File_anduril_taskmanager_v1_generic_spec_pub_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_anduril_taskmanager_v1_generic_spec_pub_proto_rawDesc), len(file_anduril_taskmanager_v1_generic_spec_pub_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_taskmanager_v1_generic_spec_pub_proto_goTypes,
		DependencyIndexes: file_anduril_taskmanager_v1_generic_spec_pub_proto_depIdxs,
		MessageInfos:      file_anduril_taskmanager_v1_generic_spec_pub_proto_msgTypes,
	}.Build()
	File_anduril_taskmanager_v1_generic_spec_pub_proto = out.File
	file_anduril_taskmanager_v1_generic_spec_pub_proto_goTypes = nil
	file_anduril_taskmanager_v1_generic_spec_pub_proto_depIdxs = nil
}
