// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: anduril/blobs/v1/filter.pub.proto

package blobsv1

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

// If match and not_match are both populated, both must evaluate to true for Statement to evaluate to true.
// If match is not populated, only not_match is used. If not_match is empty, only match is used.
// If match is not populated and not_match is empty, Statement evaluates to true.
type Statement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Positive match fields.
	Match *FieldMatcher `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	// Negative match fields. All not_match FieldMatchers must evaluate to false for this to evaluate to true.
	NotMatch []*FieldMatcher `protobuf:"bytes,2,rep,name=not_match,json=notMatch,proto3" json:"not_match,omitempty"`
}

func (x *Statement) Reset() {
	*x = Statement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_blobs_v1_filter_pub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statement) ProtoMessage() {}

func (x *Statement) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_blobs_v1_filter_pub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statement.ProtoReflect.Descriptor instead.
func (*Statement) Descriptor() ([]byte, []int) {
	return file_anduril_blobs_v1_filter_pub_proto_rawDescGZIP(), []int{0}
}

func (x *Statement) GetMatch() *FieldMatcher {
	if x != nil {
		return x.Match
	}
	return nil
}

func (x *Statement) GetNotMatch() []*FieldMatcher {
	if x != nil {
		return x.NotMatch
	}
	return nil
}

// All populated fields must be an exact match (logical AND).
type FieldMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntegrationName string `protobuf:"bytes,1,opt,name=integration_name,json=integrationName,proto3" json:"integration_name,omitempty"`
	MimeType        string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	DataType        string `protobuf:"bytes,3,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
}

func (x *FieldMatcher) Reset() {
	*x = FieldMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_blobs_v1_filter_pub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldMatcher) ProtoMessage() {}

func (x *FieldMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_blobs_v1_filter_pub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldMatcher.ProtoReflect.Descriptor instead.
func (*FieldMatcher) Descriptor() ([]byte, []int) {
	return file_anduril_blobs_v1_filter_pub_proto_rawDescGZIP(), []int{1}
}

func (x *FieldMatcher) GetIntegrationName() string {
	if x != nil {
		return x.IntegrationName
	}
	return ""
}

func (x *FieldMatcher) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *FieldMatcher) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

var File_anduril_blobs_v1_filter_pub_proto protoreflect.FileDescriptor

var file_anduril_blobs_v1_filter_pub_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x62, 0x6c, 0x6f,
	0x62, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x7e, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x62, 0x6c, 0x6f, 0x62,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x3b, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x5f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x08, 0x6e, 0x6f, 0x74,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x22, 0x73, 0x0a, 0x0c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x42, 0xc4, 0x01, 0x0a, 0x14, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x73,
	0x2e, 0x76, 0x31, 0x42, 0x0e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x75, 0x62, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c,
	0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x41, 0x42, 0x58, 0xaa, 0x02, 0x10, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x41, 0x6e, 0x64,
	0x75, 0x72, 0x69, 0x6c, 0x5c, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c,
	0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x41,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x3a, 0x3a, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_anduril_blobs_v1_filter_pub_proto_rawDescOnce sync.Once
	file_anduril_blobs_v1_filter_pub_proto_rawDescData = file_anduril_blobs_v1_filter_pub_proto_rawDesc
)

func file_anduril_blobs_v1_filter_pub_proto_rawDescGZIP() []byte {
	file_anduril_blobs_v1_filter_pub_proto_rawDescOnce.Do(func() {
		file_anduril_blobs_v1_filter_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_blobs_v1_filter_pub_proto_rawDescData)
	})
	return file_anduril_blobs_v1_filter_pub_proto_rawDescData
}

var file_anduril_blobs_v1_filter_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_anduril_blobs_v1_filter_pub_proto_goTypes = []any{
	(*Statement)(nil),    // 0: anduril.blobs.v1.Statement
	(*FieldMatcher)(nil), // 1: anduril.blobs.v1.FieldMatcher
}
var file_anduril_blobs_v1_filter_pub_proto_depIdxs = []int32{
	1, // 0: anduril.blobs.v1.Statement.match:type_name -> anduril.blobs.v1.FieldMatcher
	1, // 1: anduril.blobs.v1.Statement.not_match:type_name -> anduril.blobs.v1.FieldMatcher
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_anduril_blobs_v1_filter_pub_proto_init() }
func file_anduril_blobs_v1_filter_pub_proto_init() {
	if File_anduril_blobs_v1_filter_pub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_anduril_blobs_v1_filter_pub_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Statement); i {
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
		file_anduril_blobs_v1_filter_pub_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FieldMatcher); i {
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
			RawDescriptor: file_anduril_blobs_v1_filter_pub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_blobs_v1_filter_pub_proto_goTypes,
		DependencyIndexes: file_anduril_blobs_v1_filter_pub_proto_depIdxs,
		MessageInfos:      file_anduril_blobs_v1_filter_pub_proto_msgTypes,
	}.Build()
	File_anduril_blobs_v1_filter_pub_proto = out.File
	file_anduril_blobs_v1_filter_pub_proto_rawDesc = nil
	file_anduril_blobs_v1_filter_pub_proto_goTypes = nil
	file_anduril_blobs_v1_filter_pub_proto_depIdxs = nil
}
