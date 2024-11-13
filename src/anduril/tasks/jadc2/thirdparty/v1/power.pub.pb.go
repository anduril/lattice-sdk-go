// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: anduril/tasks/jadc2/thirdparty/v1/power.pub.proto

package thirdpartyv1

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

type PowerState int32

const (
	PowerState_POWER_STATE_INVALID PowerState = 0
	PowerState_POWER_STATE_ON      PowerState = 1
	PowerState_POWER_STATE_OFF     PowerState = 2
)

// Enum value maps for PowerState.
var (
	PowerState_name = map[int32]string{
		0: "POWER_STATE_INVALID",
		1: "POWER_STATE_ON",
		2: "POWER_STATE_OFF",
	}
	PowerState_value = map[string]int32{
		"POWER_STATE_INVALID": 0,
		"POWER_STATE_ON":      1,
		"POWER_STATE_OFF":     2,
	}
)

func (x PowerState) Enum() *PowerState {
	p := new(PowerState)
	*p = x
	return p
}

func (x PowerState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PowerState) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_enumTypes[0].Descriptor()
}

func (PowerState) Type() protoreflect.EnumType {
	return &file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_enumTypes[0]
}

func (x PowerState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PowerState.Descriptor instead.
func (PowerState) EnumDescriptor() ([]byte, []int) {
	return file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_rawDescGZIP(), []int{0}
}

// Set the power state of a robot. It is up to the robot to interpret the power state and act accordingly.
type SetPowerState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PowerState PowerState `protobuf:"varint,1,opt,name=power_state,json=powerState,proto3,enum=anduril.tasks.jadc2.thirdparty.v1.PowerState" json:"power_state,omitempty"`
}

func (x *SetPowerState) Reset() {
	*x = SetPowerState{}
	mi := &file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetPowerState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPowerState) ProtoMessage() {}

func (x *SetPowerState) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPowerState.ProtoReflect.Descriptor instead.
func (*SetPowerState) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_rawDescGZIP(), []int{0}
}

func (x *SetPowerState) GetPowerState() PowerState {
	if x != nil {
		return x.PowerState
	}
	return PowerState_POWER_STATE_INVALID
}

var File_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto protoreflect.FileDescriptor

var file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_rawDesc = []byte{
	0x0a, 0x31, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f,
	0x6a, 0x61, 0x64, 0x63, 0x32, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x21, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2e, 0x6a, 0x61, 0x64, 0x63, 0x32, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x5f, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x50, 0x6f, 0x77,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x70, 0x6f, 0x77, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x61,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x6a, 0x61, 0x64,
	0x63, 0x32, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x70, 0x6f, 0x77,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2a, 0x4e, 0x0a, 0x0a, 0x50, 0x6f, 0x77, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x12,
	0x0a, 0x0e, 0x50, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4e,
	0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x02, 0x42, 0xb1, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x6a, 0x61,
	0x64, 0x63, 0x32, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x42, 0x0d, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x50, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2d, 0x67,
	0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x2f, 0x6a, 0x61, 0x64, 0x63, 0x32, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x41, 0x54, 0x4a, 0x54, 0xaa, 0x02, 0x21, 0x41, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x4a, 0x61, 0x64, 0x63,
	0x32, 0x2e, 0x54, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x21, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x5c,
	0x4a, 0x61, 0x64, 0x63, 0x32, 0x5c, 0x54, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2d, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x5c, 0x4a, 0x61, 0x64, 0x63, 0x32, 0x5c, 0x54, 0x68, 0x69, 0x72, 0x64, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x25, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x3a, 0x3a, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x3a, 0x3a, 0x4a, 0x61, 0x64, 0x63, 0x32, 0x3a, 0x3a, 0x54, 0x68, 0x69,
	0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_rawDescOnce sync.Once
	file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_rawDescData = file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_rawDesc
)

func file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_rawDescGZIP() []byte {
	file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_rawDescOnce.Do(func() {
		file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_rawDescData)
	})
	return file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_rawDescData
}

var file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_goTypes = []any{
	(PowerState)(0),       // 0: anduril.tasks.jadc2.thirdparty.v1.PowerState
	(*SetPowerState)(nil), // 1: anduril.tasks.jadc2.thirdparty.v1.SetPowerState
}
var file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_depIdxs = []int32{
	0, // 0: anduril.tasks.jadc2.thirdparty.v1.SetPowerState.power_state:type_name -> anduril.tasks.jadc2.thirdparty.v1.PowerState
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_init() }
func file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_init() {
	if File_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_goTypes,
		DependencyIndexes: file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_depIdxs,
		EnumInfos:         file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_enumTypes,
		MessageInfos:      file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_msgTypes,
	}.Build()
	File_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto = out.File
	file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_rawDesc = nil
	file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_goTypes = nil
	file_anduril_tasks_jadc2_thirdparty_v1_power_pub_proto_depIdxs = nil
}
