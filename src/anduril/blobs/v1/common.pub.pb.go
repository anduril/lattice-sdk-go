// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: anduril/blobs/v1/common.pub.proto

package blobsv1

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

// Metadata about the blob.
type BlobMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// timestamp of when blob was created
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// timestamp of when this blob will be retained until.
	RetentionTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=retention_time,json=retentionTime,proto3" json:"retention_time,omitempty"`
	// size of the blob contents in bytes
	SizeBytes uint64 `protobuf:"varint,3,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	// md5 sum/hash of blob contents
	Md5 []byte `protobuf:"bytes,4,opt,name=md5,proto3" json:"md5,omitempty"`
	// MIME type of blob contents
	ContentType string `protobuf:"bytes,5,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// details regarding the blob's origin - aka source and type
	Provenance *BlobProvenance `protobuf:"bytes,6,opt,name=provenance,proto3" json:"provenance,omitempty"`
}

func (x *BlobMetadata) Reset() {
	*x = BlobMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_blobs_v1_common_pub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlobMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlobMetadata) ProtoMessage() {}

func (x *BlobMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_blobs_v1_common_pub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlobMetadata.ProtoReflect.Descriptor instead.
func (*BlobMetadata) Descriptor() ([]byte, []int) {
	return file_anduril_blobs_v1_common_pub_proto_rawDescGZIP(), []int{0}
}

func (x *BlobMetadata) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *BlobMetadata) GetRetentionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RetentionTime
	}
	return nil
}

func (x *BlobMetadata) GetSizeBytes() uint64 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

func (x *BlobMetadata) GetMd5() []byte {
	if x != nil {
		return x.Md5
	}
	return nil
}

func (x *BlobMetadata) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *BlobMetadata) GetProvenance() *BlobProvenance {
	if x != nil {
		return x.Provenance
	}
	return nil
}

type BlobProvenance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// registered integration_name producing the Blob.
	IntegrationName string `protobuf:"bytes,1,opt,name=integration_name,json=integrationName,proto3" json:"integration_name,omitempty"`
	// data type defines what is inside of the Blob. Must be registered.
	DataType string `protobuf:"bytes,2,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
}

func (x *BlobProvenance) Reset() {
	*x = BlobProvenance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_blobs_v1_common_pub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlobProvenance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlobProvenance) ProtoMessage() {}

func (x *BlobProvenance) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_blobs_v1_common_pub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlobProvenance.ProtoReflect.Descriptor instead.
func (*BlobProvenance) Descriptor() ([]byte, []int) {
	return file_anduril_blobs_v1_common_pub_proto_rawDescGZIP(), []int{1}
}

func (x *BlobProvenance) GetIntegrationName() string {
	if x != nil {
		return x.IntegrationName
	}
	return ""
}

func (x *BlobProvenance) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

var File_anduril_blobs_v1_common_pub_proto protoreflect.FileDescriptor

var file_anduril_blobs_v1_common_pub_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x62, 0x6c, 0x6f,
	0x62, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x02, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x62, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x72, 0x65, 0x74, 0x65,
	0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x7a,
	0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73,
	0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d, 0x64, 0x35, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x62, 0x6c, 0x6f, 0x62,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x22,
	0x58, 0x0a, 0x0e, 0x42, 0x6c, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x42, 0xce, 0x01, 0x0a, 0x14, 0x63, 0x6f,
	0x6d, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x2e,
	0x76, 0x31, 0x42, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x75, 0x62, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x6f, 0x67, 0x75, 0x6e, 0x2d, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x61,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72,
	0x63, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x42, 0x58,
	0xaa, 0x02, 0x10, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x73,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x42, 0x6c,
	0x6f, 0x62, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c,
	0x5c, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x3a,
	0x3a, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_anduril_blobs_v1_common_pub_proto_rawDescOnce sync.Once
	file_anduril_blobs_v1_common_pub_proto_rawDescData = file_anduril_blobs_v1_common_pub_proto_rawDesc
)

func file_anduril_blobs_v1_common_pub_proto_rawDescGZIP() []byte {
	file_anduril_blobs_v1_common_pub_proto_rawDescOnce.Do(func() {
		file_anduril_blobs_v1_common_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_blobs_v1_common_pub_proto_rawDescData)
	})
	return file_anduril_blobs_v1_common_pub_proto_rawDescData
}

var file_anduril_blobs_v1_common_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_anduril_blobs_v1_common_pub_proto_goTypes = []any{
	(*BlobMetadata)(nil),          // 0: anduril.blobs.v1.BlobMetadata
	(*BlobProvenance)(nil),        // 1: anduril.blobs.v1.BlobProvenance
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_anduril_blobs_v1_common_pub_proto_depIdxs = []int32{
	2, // 0: anduril.blobs.v1.BlobMetadata.created_time:type_name -> google.protobuf.Timestamp
	2, // 1: anduril.blobs.v1.BlobMetadata.retention_time:type_name -> google.protobuf.Timestamp
	1, // 2: anduril.blobs.v1.BlobMetadata.provenance:type_name -> anduril.blobs.v1.BlobProvenance
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_anduril_blobs_v1_common_pub_proto_init() }
func file_anduril_blobs_v1_common_pub_proto_init() {
	if File_anduril_blobs_v1_common_pub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_anduril_blobs_v1_common_pub_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BlobMetadata); i {
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
		file_anduril_blobs_v1_common_pub_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BlobProvenance); i {
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
			RawDescriptor: file_anduril_blobs_v1_common_pub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_blobs_v1_common_pub_proto_goTypes,
		DependencyIndexes: file_anduril_blobs_v1_common_pub_proto_depIdxs,
		MessageInfos:      file_anduril_blobs_v1_common_pub_proto_msgTypes,
	}.Build()
	File_anduril_blobs_v1_common_pub_proto = out.File
	file_anduril_blobs_v1_common_pub_proto_rawDesc = nil
	file_anduril_blobs_v1_common_pub_proto_goTypes = nil
	file_anduril_blobs_v1_common_pub_proto_depIdxs = nil
}
