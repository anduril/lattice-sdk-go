// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: anduril/communicationsmanager/v1/entity.pub.proto

package communicationsmanagerv1

import (
	v1 "github.com/dogun-anduril/anduril-sdk-go/src/anduril/entitymanager/v1"
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

type EntityIntegrationDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vendors []*EntityDataVendor `protobuf:"bytes,1,rep,name=vendors,proto3" json:"vendors,omitempty"`
}

func (x *EntityIntegrationDetails) Reset() {
	*x = EntityIntegrationDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_communicationsmanager_v1_entity_pub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityIntegrationDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityIntegrationDetails) ProtoMessage() {}

func (x *EntityIntegrationDetails) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_communicationsmanager_v1_entity_pub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityIntegrationDetails.ProtoReflect.Descriptor instead.
func (*EntityIntegrationDetails) Descriptor() ([]byte, []int) {
	return file_anduril_communicationsmanager_v1_entity_pub_proto_rawDescGZIP(), []int{0}
}

func (x *EntityIntegrationDetails) GetVendors() []*EntityDataVendor {
	if x != nil {
		return x.Vendors
	}
	return nil
}

type EntityDataVendor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VendorName string   `protobuf:"bytes,1,opt,name=vendor_name,json=vendorName,proto3" json:"vendor_name,omitempty"`
	DataType   []string `protobuf:"bytes,2,rep,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
}

func (x *EntityDataVendor) Reset() {
	*x = EntityDataVendor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_communicationsmanager_v1_entity_pub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityDataVendor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityDataVendor) ProtoMessage() {}

func (x *EntityDataVendor) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_communicationsmanager_v1_entity_pub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityDataVendor.ProtoReflect.Descriptor instead.
func (*EntityDataVendor) Descriptor() ([]byte, []int) {
	return file_anduril_communicationsmanager_v1_entity_pub_proto_rawDescGZIP(), []int{1}
}

func (x *EntityDataVendor) GetVendorName() string {
	if x != nil {
		return x.VendorName
	}
	return ""
}

func (x *EntityDataVendor) GetDataType() []string {
	if x != nil {
		return x.DataType
	}
	return nil
}

// Message to contain entity integration rules
type EntityIntegrationRuleDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The filter to be applied to the integration
	Filter *v1.Statement `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *EntityIntegrationRuleDetails) Reset() {
	*x = EntityIntegrationRuleDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_communicationsmanager_v1_entity_pub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityIntegrationRuleDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityIntegrationRuleDetails) ProtoMessage() {}

func (x *EntityIntegrationRuleDetails) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_communicationsmanager_v1_entity_pub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityIntegrationRuleDetails.ProtoReflect.Descriptor instead.
func (*EntityIntegrationRuleDetails) Descriptor() ([]byte, []int) {
	return file_anduril_communicationsmanager_v1_entity_pub_proto_rawDescGZIP(), []int{2}
}

func (x *EntityIntegrationRuleDetails) GetFilter() *v1.Statement {
	if x != nil {
		return x.Filter
	}
	return nil
}

var File_anduril_communicationsmanager_v1_entity_pub_proto protoreflect.FileDescriptor

var file_anduril_communicationsmanager_v1_entity_pub_proto_rawDesc = []byte{
	0x0a, 0x31, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x20, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x29, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x68, 0x0a, 0x18, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x4c, 0x0a, 0x07,
	0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x56, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x52, 0x07, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x22, 0x50, 0x0a, 0x10, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x22, 0x5b, 0x0a, 0x1c,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x3b, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0xbe, 0x02, 0x0a, 0x24, 0x63, 0x6f,
	0x6d, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x42, 0x0e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x75, 0x62, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x64, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x6f, 0x67, 0x75, 0x6e, 0x2d, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x61,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72,
	0x63, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58,
	0xaa, 0x02, 0x20, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x20, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x43, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2c, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c,
	0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x22, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x3a,
	0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_anduril_communicationsmanager_v1_entity_pub_proto_rawDescOnce sync.Once
	file_anduril_communicationsmanager_v1_entity_pub_proto_rawDescData = file_anduril_communicationsmanager_v1_entity_pub_proto_rawDesc
)

func file_anduril_communicationsmanager_v1_entity_pub_proto_rawDescGZIP() []byte {
	file_anduril_communicationsmanager_v1_entity_pub_proto_rawDescOnce.Do(func() {
		file_anduril_communicationsmanager_v1_entity_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_communicationsmanager_v1_entity_pub_proto_rawDescData)
	})
	return file_anduril_communicationsmanager_v1_entity_pub_proto_rawDescData
}

var file_anduril_communicationsmanager_v1_entity_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_anduril_communicationsmanager_v1_entity_pub_proto_goTypes = []any{
	(*EntityIntegrationDetails)(nil),     // 0: anduril.communicationsmanager.v1.EntityIntegrationDetails
	(*EntityDataVendor)(nil),             // 1: anduril.communicationsmanager.v1.EntityDataVendor
	(*EntityIntegrationRuleDetails)(nil), // 2: anduril.communicationsmanager.v1.EntityIntegrationRuleDetails
	(*v1.Statement)(nil),                 // 3: anduril.entitymanager.v1.Statement
}
var file_anduril_communicationsmanager_v1_entity_pub_proto_depIdxs = []int32{
	1, // 0: anduril.communicationsmanager.v1.EntityIntegrationDetails.vendors:type_name -> anduril.communicationsmanager.v1.EntityDataVendor
	3, // 1: anduril.communicationsmanager.v1.EntityIntegrationRuleDetails.filter:type_name -> anduril.entitymanager.v1.Statement
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_anduril_communicationsmanager_v1_entity_pub_proto_init() }
func file_anduril_communicationsmanager_v1_entity_pub_proto_init() {
	if File_anduril_communicationsmanager_v1_entity_pub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_anduril_communicationsmanager_v1_entity_pub_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EntityIntegrationDetails); i {
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
		file_anduril_communicationsmanager_v1_entity_pub_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EntityDataVendor); i {
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
		file_anduril_communicationsmanager_v1_entity_pub_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EntityIntegrationRuleDetails); i {
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
			RawDescriptor: file_anduril_communicationsmanager_v1_entity_pub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_communicationsmanager_v1_entity_pub_proto_goTypes,
		DependencyIndexes: file_anduril_communicationsmanager_v1_entity_pub_proto_depIdxs,
		MessageInfos:      file_anduril_communicationsmanager_v1_entity_pub_proto_msgTypes,
	}.Build()
	File_anduril_communicationsmanager_v1_entity_pub_proto = out.File
	file_anduril_communicationsmanager_v1_entity_pub_proto_rawDesc = nil
	file_anduril_communicationsmanager_v1_entity_pub_proto_goTypes = nil
	file_anduril_communicationsmanager_v1_entity_pub_proto_depIdxs = nil
}
