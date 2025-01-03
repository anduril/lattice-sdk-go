// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: anduril/entitymanager/v1/media.pub.proto

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

type MediaType int32

const (
	MediaType_MEDIA_TYPE_INVALID      MediaType = 0
	MediaType_MEDIA_TYPE_THUMBNAIL    MediaType = 1
	MediaType_MEDIA_TYPE_IMAGE        MediaType = 2
	MediaType_MEDIA_TYPE_VIDEO        MediaType = 3
	MediaType_MEDIA_TYPE_SLIPPY_TILES MediaType = 4
)

// Enum value maps for MediaType.
var (
	MediaType_name = map[int32]string{
		0: "MEDIA_TYPE_INVALID",
		1: "MEDIA_TYPE_THUMBNAIL",
		2: "MEDIA_TYPE_IMAGE",
		3: "MEDIA_TYPE_VIDEO",
		4: "MEDIA_TYPE_SLIPPY_TILES",
	}
	MediaType_value = map[string]int32{
		"MEDIA_TYPE_INVALID":      0,
		"MEDIA_TYPE_THUMBNAIL":    1,
		"MEDIA_TYPE_IMAGE":        2,
		"MEDIA_TYPE_VIDEO":        3,
		"MEDIA_TYPE_SLIPPY_TILES": 4,
	}
)

func (x MediaType) Enum() *MediaType {
	p := new(MediaType)
	*p = x
	return p
}

func (x MediaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MediaType) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_media_pub_proto_enumTypes[0].Descriptor()
}

func (MediaType) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_media_pub_proto_enumTypes[0]
}

func (x MediaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MediaType.Descriptor instead.
func (MediaType) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_media_pub_proto_rawDescGZIP(), []int{0}
}

// Media associated with an entity.
type Media struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Media         []*MediaItem           `protobuf:"bytes,1,rep,name=media,proto3" json:"media,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Media) Reset() {
	*x = Media{}
	mi := &file_anduril_entitymanager_v1_media_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_media_pub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Media.ProtoReflect.Descriptor instead.
func (*Media) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_media_pub_proto_rawDescGZIP(), []int{0}
}

func (x *Media) GetMedia() []*MediaItem {
	if x != nil {
		return x.Media
	}
	return nil
}

type MediaItem struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// To Be Deprecated, use relative_path.
	// The url where the media related to an entity can be accessed
	Url  string    `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Type MediaType `protobuf:"varint,2,opt,name=type,proto3,enum=anduril.entitymanager.v1.MediaType" json:"type,omitempty"`
	// The relative path where the media related to an entity can be accessed when used to query against a blobs service
	// node.
	RelativePath  string `protobuf:"bytes,3,opt,name=relative_path,json=relativePath,proto3" json:"relative_path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MediaItem) Reset() {
	*x = MediaItem{}
	mi := &file_anduril_entitymanager_v1_media_pub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MediaItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaItem) ProtoMessage() {}

func (x *MediaItem) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_media_pub_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaItem.ProtoReflect.Descriptor instead.
func (*MediaItem) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_media_pub_proto_rawDescGZIP(), []int{1}
}

func (x *MediaItem) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *MediaItem) GetType() MediaType {
	if x != nil {
		return x.Type
	}
	return MediaType_MEDIA_TYPE_INVALID
}

func (x *MediaItem) GetRelativePath() string {
	if x != nil {
		return x.RelativePath
	}
	return ""
}

var File_anduril_entitymanager_v1_media_pub_proto protoreflect.FileDescriptor

var file_anduril_entitymanager_v1_media_pub_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x47, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x3e, 0x0a, 0x05, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72,
	0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x03, 0xc8,
	0x3e, 0x01, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x22, 0x7b, 0x0a, 0x09, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x2a, 0x86, 0x01, 0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14,
	0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x48, 0x55, 0x4d, 0x42,
	0x4e, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10,
	0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f,
	0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x53, 0x4c, 0x49, 0x50, 0x50, 0x59, 0x5f, 0x54, 0x49, 0x4c, 0x45, 0x53, 0x10, 0x04, 0x42,
	0xff, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x42, 0x0d, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x50, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
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
	file_anduril_entitymanager_v1_media_pub_proto_rawDescOnce sync.Once
	file_anduril_entitymanager_v1_media_pub_proto_rawDescData = file_anduril_entitymanager_v1_media_pub_proto_rawDesc
)

func file_anduril_entitymanager_v1_media_pub_proto_rawDescGZIP() []byte {
	file_anduril_entitymanager_v1_media_pub_proto_rawDescOnce.Do(func() {
		file_anduril_entitymanager_v1_media_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_entitymanager_v1_media_pub_proto_rawDescData)
	})
	return file_anduril_entitymanager_v1_media_pub_proto_rawDescData
}

var file_anduril_entitymanager_v1_media_pub_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_anduril_entitymanager_v1_media_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_anduril_entitymanager_v1_media_pub_proto_goTypes = []any{
	(MediaType)(0),    // 0: anduril.entitymanager.v1.MediaType
	(*Media)(nil),     // 1: anduril.entitymanager.v1.Media
	(*MediaItem)(nil), // 2: anduril.entitymanager.v1.MediaItem
}
var file_anduril_entitymanager_v1_media_pub_proto_depIdxs = []int32{
	2, // 0: anduril.entitymanager.v1.Media.media:type_name -> anduril.entitymanager.v1.MediaItem
	0, // 1: anduril.entitymanager.v1.MediaItem.type:type_name -> anduril.entitymanager.v1.MediaType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_anduril_entitymanager_v1_media_pub_proto_init() }
func file_anduril_entitymanager_v1_media_pub_proto_init() {
	if File_anduril_entitymanager_v1_media_pub_proto != nil {
		return
	}
	file_anduril_entitymanager_v1_options_pub_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_anduril_entitymanager_v1_media_pub_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_entitymanager_v1_media_pub_proto_goTypes,
		DependencyIndexes: file_anduril_entitymanager_v1_media_pub_proto_depIdxs,
		EnumInfos:         file_anduril_entitymanager_v1_media_pub_proto_enumTypes,
		MessageInfos:      file_anduril_entitymanager_v1_media_pub_proto_msgTypes,
	}.Build()
	File_anduril_entitymanager_v1_media_pub_proto = out.File
	file_anduril_entitymanager_v1_media_pub_proto_rawDesc = nil
	file_anduril_entitymanager_v1_media_pub_proto_goTypes = nil
	file_anduril_entitymanager_v1_media_pub_proto_depIdxs = nil
}
