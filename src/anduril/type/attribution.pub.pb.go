// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: anduril/type/attribution.pub.proto

package _type

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

type Attribution struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The timestamp at which the event occurred, in UTC epoch microseconds.
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The user ID that initiated the event.
	UserId        string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Attribution) Reset() {
	*x = Attribution{}
	mi := &file_anduril_type_attribution_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Attribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribution) ProtoMessage() {}

func (x *Attribution) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_type_attribution_pub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribution.ProtoReflect.Descriptor instead.
func (*Attribution) Descriptor() ([]byte, []int) {
	return file_anduril_type_attribution_pub_proto_rawDescGZIP(), []int{0}
}

func (x *Attribution) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Attribution) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_anduril_type_attribution_pub_proto protoreflect.FileDescriptor

var file_anduril_type_attribution_pub_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x44, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0xac, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x42, 0x13, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x75, 0x62, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x6c, 0x61, 0x74, 0x74, 0x69, 0x63, 0x65,
	0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0xa2, 0x02, 0x03, 0x41, 0x54, 0x58, 0xaa, 0x02,
	0x0c, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x54, 0x79, 0x70, 0x65, 0xca, 0x02, 0x0c,
	0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x54, 0x79, 0x70, 0x65, 0xe2, 0x02, 0x18, 0x41,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_anduril_type_attribution_pub_proto_rawDescOnce sync.Once
	file_anduril_type_attribution_pub_proto_rawDescData = file_anduril_type_attribution_pub_proto_rawDesc
)

func file_anduril_type_attribution_pub_proto_rawDescGZIP() []byte {
	file_anduril_type_attribution_pub_proto_rawDescOnce.Do(func() {
		file_anduril_type_attribution_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_type_attribution_pub_proto_rawDescData)
	})
	return file_anduril_type_attribution_pub_proto_rawDescData
}

var file_anduril_type_attribution_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_anduril_type_attribution_pub_proto_goTypes = []any{
	(*Attribution)(nil), // 0: anduril.type.Attribution
}
var file_anduril_type_attribution_pub_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_anduril_type_attribution_pub_proto_init() }
func file_anduril_type_attribution_pub_proto_init() {
	if File_anduril_type_attribution_pub_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_anduril_type_attribution_pub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_type_attribution_pub_proto_goTypes,
		DependencyIndexes: file_anduril_type_attribution_pub_proto_depIdxs,
		MessageInfos:      file_anduril_type_attribution_pub_proto_msgTypes,
	}.Build()
	File_anduril_type_attribution_pub_proto = out.File
	file_anduril_type_attribution_pub_proto_rawDesc = nil
	file_anduril_type_attribution_pub_proto_goTypes = nil
	file_anduril_type_attribution_pub_proto_depIdxs = nil
}
