// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: anduril/entitymanager/v1/contact_details.pub.proto

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

// Contains details on how to make contact with an entity.
type ContactDetails struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The primary phone number for this entity.
	PhoneNumber   string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ContactDetails) Reset() {
	*x = ContactDetails{}
	mi := &file_anduril_entitymanager_v1_contact_details_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContactDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactDetails) ProtoMessage() {}

func (x *ContactDetails) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_contact_details_pub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactDetails.ProtoReflect.Descriptor instead.
func (*ContactDetails) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_contact_details_pub_proto_rawDescGZIP(), []int{0}
}

func (x *ContactDetails) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

var File_anduril_entitymanager_v1_contact_details_pub_proto protoreflect.FileDescriptor

var file_anduril_entitymanager_v1_contact_details_pub_proto_rawDesc = string([]byte{
	0x0a, 0x32, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x33,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x42, 0x88, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x50, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72,
	0x69, 0x6c, 0x2f, 0x6c, 0x61, 0x74, 0x74, 0x69, 0x63, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67,
	0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x41, 0x45, 0x58, 0xaa, 0x02, 0x18, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x18, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x41, 0x6e, 0x64,
	0x75, 0x72, 0x69, 0x6c, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x1a, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x3a, 0x3a, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_anduril_entitymanager_v1_contact_details_pub_proto_rawDescOnce sync.Once
	file_anduril_entitymanager_v1_contact_details_pub_proto_rawDescData []byte
)

func file_anduril_entitymanager_v1_contact_details_pub_proto_rawDescGZIP() []byte {
	file_anduril_entitymanager_v1_contact_details_pub_proto_rawDescOnce.Do(func() {
		file_anduril_entitymanager_v1_contact_details_pub_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_anduril_entitymanager_v1_contact_details_pub_proto_rawDesc), len(file_anduril_entitymanager_v1_contact_details_pub_proto_rawDesc)))
	})
	return file_anduril_entitymanager_v1_contact_details_pub_proto_rawDescData
}

var file_anduril_entitymanager_v1_contact_details_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_anduril_entitymanager_v1_contact_details_pub_proto_goTypes = []any{
	(*ContactDetails)(nil), // 0: anduril.entitymanager.v1.ContactDetails
}
var file_anduril_entitymanager_v1_contact_details_pub_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_anduril_entitymanager_v1_contact_details_pub_proto_init() }
func file_anduril_entitymanager_v1_contact_details_pub_proto_init() {
	if File_anduril_entitymanager_v1_contact_details_pub_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_anduril_entitymanager_v1_contact_details_pub_proto_rawDesc), len(file_anduril_entitymanager_v1_contact_details_pub_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_entitymanager_v1_contact_details_pub_proto_goTypes,
		DependencyIndexes: file_anduril_entitymanager_v1_contact_details_pub_proto_depIdxs,
		MessageInfos:      file_anduril_entitymanager_v1_contact_details_pub_proto_msgTypes,
	}.Build()
	File_anduril_entitymanager_v1_contact_details_pub_proto = out.File
	file_anduril_entitymanager_v1_contact_details_pub_proto_goTypes = nil
	file_anduril_entitymanager_v1_contact_details_pub_proto_depIdxs = nil
}
