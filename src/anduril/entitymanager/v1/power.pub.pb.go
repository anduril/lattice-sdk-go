// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: anduril/entitymanager/v1/power.pub.proto

package entitymanagerv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

type PowerStatus int32

const (
	PowerStatus_POWER_STATUS_INVALID PowerStatus = 0
	// Indeterminate condition of whether the power system is present or absent.
	PowerStatus_POWER_STATUS_UNKNOWN PowerStatus = 1
	// Power system is not configured/present. This is considered a normal/expected condition, as opposed to the
	// system is expected to be present but is missing.
	PowerStatus_POWER_STATUS_NOT_PRESENT PowerStatus = 2
	// Power system is present and operating normally.
	PowerStatus_POWER_STATUS_OPERATING PowerStatus = 3
	// Power system is present and is in an expected disabled state. For example, if the generator was shut off for
	// operational reasons.
	PowerStatus_POWER_STATUS_DISABLED PowerStatus = 4
	// Power system is non-functional.
	PowerStatus_POWER_STATUS_ERROR PowerStatus = 5
)

// Enum value maps for PowerStatus.
var (
	PowerStatus_name = map[int32]string{
		0: "POWER_STATUS_INVALID",
		1: "POWER_STATUS_UNKNOWN",
		2: "POWER_STATUS_NOT_PRESENT",
		3: "POWER_STATUS_OPERATING",
		4: "POWER_STATUS_DISABLED",
		5: "POWER_STATUS_ERROR",
	}
	PowerStatus_value = map[string]int32{
		"POWER_STATUS_INVALID":     0,
		"POWER_STATUS_UNKNOWN":     1,
		"POWER_STATUS_NOT_PRESENT": 2,
		"POWER_STATUS_OPERATING":   3,
		"POWER_STATUS_DISABLED":    4,
		"POWER_STATUS_ERROR":       5,
	}
)

func (x PowerStatus) Enum() *PowerStatus {
	p := new(PowerStatus)
	*p = x
	return p
}

func (x PowerStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PowerStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_power_pub_proto_enumTypes[0].Descriptor()
}

func (PowerStatus) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_power_pub_proto_enumTypes[0]
}

func (x PowerStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PowerStatus.Descriptor instead.
func (PowerStatus) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_power_pub_proto_rawDescGZIP(), []int{0}
}

type PowerType int32

const (
	PowerType_POWER_TYPE_INVALID PowerType = 0
	PowerType_POWER_TYPE_UNKNOWN PowerType = 1
	PowerType_POWER_TYPE_GAS     PowerType = 2
	PowerType_POWER_TYPE_BATTERY PowerType = 3
)

// Enum value maps for PowerType.
var (
	PowerType_name = map[int32]string{
		0: "POWER_TYPE_INVALID",
		1: "POWER_TYPE_UNKNOWN",
		2: "POWER_TYPE_GAS",
		3: "POWER_TYPE_BATTERY",
	}
	PowerType_value = map[string]int32{
		"POWER_TYPE_INVALID": 0,
		"POWER_TYPE_UNKNOWN": 1,
		"POWER_TYPE_GAS":     2,
		"POWER_TYPE_BATTERY": 3,
	}
)

func (x PowerType) Enum() *PowerType {
	p := new(PowerType)
	*p = x
	return p
}

func (x PowerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PowerType) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_power_pub_proto_enumTypes[1].Descriptor()
}

func (PowerType) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_power_pub_proto_enumTypes[1]
}

func (x PowerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PowerType.Descriptor instead.
func (PowerType) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_power_pub_proto_rawDescGZIP(), []int{1}
}

// Represents the state of power sources connected to this entity.
type PowerState struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This is a map where the key is a unique id of the power source and the value is additional information about the
	// power source.
	SourceIdToState map[string]*PowerSource `protobuf:"bytes,5,rep,name=source_id_to_state,json=sourceIdToState,proto3" json:"source_id_to_state,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *PowerState) Reset() {
	*x = PowerState{}
	mi := &file_anduril_entitymanager_v1_power_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PowerState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowerState) ProtoMessage() {}

func (x *PowerState) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_power_pub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowerState.ProtoReflect.Descriptor instead.
func (*PowerState) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_power_pub_proto_rawDescGZIP(), []int{0}
}

func (x *PowerState) GetSourceIdToState() map[string]*PowerSource {
	if x != nil {
		return x.SourceIdToState
	}
	return nil
}

// Represents the state of a single power source that is connected to this entity.
type PowerSource struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Status of the power source.
	PowerStatus PowerStatus `protobuf:"varint,1,opt,name=power_status,json=powerStatus,proto3,enum=anduril.entitymanager.v1.PowerStatus" json:"power_status,omitempty"`
	// Used to determine the type of power source.
	PowerType PowerType `protobuf:"varint,2,opt,name=power_type,json=powerType,proto3,enum=anduril.entitymanager.v1.PowerType" json:"power_type,omitempty"`
	// Power level of the system. If absent, the power level is assumed to be unknown.
	PowerLevel *PowerLevel `protobuf:"bytes,3,opt,name=power_level,json=powerLevel,proto3" json:"power_level,omitempty"`
	// Set of human-readable messages with status of the power system. Typically this would be used in an error state
	// to provide additional error information. This can also be used for informational messages.
	Messages []string `protobuf:"bytes,4,rep,name=messages,proto3" json:"messages,omitempty"`
	// Whether the power source is offloadable. If the value is missing (as opposed to false) then the entity does not
	// report whether the power source is offloadable.
	Offloadable   *wrapperspb.BoolValue `protobuf:"bytes,5,opt,name=offloadable,proto3" json:"offloadable,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PowerSource) Reset() {
	*x = PowerSource{}
	mi := &file_anduril_entitymanager_v1_power_pub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PowerSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowerSource) ProtoMessage() {}

func (x *PowerSource) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_power_pub_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowerSource.ProtoReflect.Descriptor instead.
func (*PowerSource) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_power_pub_proto_rawDescGZIP(), []int{1}
}

func (x *PowerSource) GetPowerStatus() PowerStatus {
	if x != nil {
		return x.PowerStatus
	}
	return PowerStatus_POWER_STATUS_INVALID
}

func (x *PowerSource) GetPowerType() PowerType {
	if x != nil {
		return x.PowerType
	}
	return PowerType_POWER_TYPE_INVALID
}

func (x *PowerSource) GetPowerLevel() *PowerLevel {
	if x != nil {
		return x.PowerLevel
	}
	return nil
}

func (x *PowerSource) GetMessages() []string {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *PowerSource) GetOffloadable() *wrapperspb.BoolValue {
	if x != nil {
		return x.Offloadable
	}
	return nil
}

// Represents the power level of a system.
type PowerLevel struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Total power capacity of the system.
	Capacity float32 `protobuf:"fixed32,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	// Remaining power capacity of the system.
	Remaining float32 `protobuf:"fixed32,2,opt,name=remaining,proto3" json:"remaining,omitempty"`
	// Percent of power remaining.
	PercentRemaining float32 `protobuf:"fixed32,3,opt,name=percent_remaining,json=percentRemaining,proto3" json:"percent_remaining,omitempty"`
	// Voltage of the power source subsystem, as reported by the power source. If the source does not report this value
	// this field will be null.
	Voltage *wrapperspb.DoubleValue `protobuf:"bytes,4,opt,name=voltage,proto3" json:"voltage,omitempty"`
	// Current in amps of the power source subsystem, as reported by the power source. If the source does not
	// report this value this field will be null.
	CurrentAmps *wrapperspb.DoubleValue `protobuf:"bytes,5,opt,name=current_amps,json=currentAmps,proto3" json:"current_amps,omitempty"`
	// Estimated minutes until empty. Calculated with consumption at the moment, as reported by the power source. If the source does not
	// report this value this field will be null.
	RunTimeToEmptyMins *wrapperspb.DoubleValue `protobuf:"bytes,6,opt,name=run_time_to_empty_mins,json=runTimeToEmptyMins,proto3" json:"run_time_to_empty_mins,omitempty"`
	// Fuel consumption rate in liters per second.
	ConsumptionRateLPerS *wrapperspb.DoubleValue `protobuf:"bytes,7,opt,name=consumption_rate_l_per_s,json=consumptionRateLPerS,proto3" json:"consumption_rate_l_per_s,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *PowerLevel) Reset() {
	*x = PowerLevel{}
	mi := &file_anduril_entitymanager_v1_power_pub_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PowerLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowerLevel) ProtoMessage() {}

func (x *PowerLevel) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_power_pub_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowerLevel.ProtoReflect.Descriptor instead.
func (*PowerLevel) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_power_pub_proto_rawDescGZIP(), []int{2}
}

func (x *PowerLevel) GetCapacity() float32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *PowerLevel) GetRemaining() float32 {
	if x != nil {
		return x.Remaining
	}
	return 0
}

func (x *PowerLevel) GetPercentRemaining() float32 {
	if x != nil {
		return x.PercentRemaining
	}
	return 0
}

func (x *PowerLevel) GetVoltage() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Voltage
	}
	return nil
}

func (x *PowerLevel) GetCurrentAmps() *wrapperspb.DoubleValue {
	if x != nil {
		return x.CurrentAmps
	}
	return nil
}

func (x *PowerLevel) GetRunTimeToEmptyMins() *wrapperspb.DoubleValue {
	if x != nil {
		return x.RunTimeToEmptyMins
	}
	return nil
}

func (x *PowerLevel) GetConsumptionRateLPerS() *wrapperspb.DoubleValue {
	if x != nil {
		return x.ConsumptionRateLPerS
	}
	return nil
}

var File_anduril_entitymanager_v1_power_pub_proto protoreflect.FileDescriptor

const file_anduril_entitymanager_v1_power_pub_proto_rawDesc = "" +
	"\n" +
	"(anduril/entitymanager/v1/power.pub.proto\x12\x18anduril.entitymanager.v1\x1a\x1egoogle/protobuf/wrappers.proto\"\xf7\x01\n" +
	"\n" +
	"PowerState\x12f\n" +
	"\x12source_id_to_state\x18\x05 \x03(\v29.anduril.entitymanager.v1.PowerState.SourceIdToStateEntryR\x0fsourceIdToState\x1ai\n" +
	"\x14SourceIdToStateEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12;\n" +
	"\x05value\x18\x02 \x01(\v2%.anduril.entitymanager.v1.PowerSourceR\x05value:\x028\x01J\x04\b\x01\x10\x02J\x04\b\x02\x10\x03J\x04\b\x03\x10\x04J\x04\b\x04\x10\x05\"\xbc\x02\n" +
	"\vPowerSource\x12H\n" +
	"\fpower_status\x18\x01 \x01(\x0e2%.anduril.entitymanager.v1.PowerStatusR\vpowerStatus\x12B\n" +
	"\n" +
	"power_type\x18\x02 \x01(\x0e2#.anduril.entitymanager.v1.PowerTypeR\tpowerType\x12E\n" +
	"\vpower_level\x18\x03 \x01(\v2$.anduril.entitymanager.v1.PowerLevelR\n" +
	"powerLevel\x12\x1a\n" +
	"\bmessages\x18\x04 \x03(\tR\bmessages\x12<\n" +
	"\voffloadable\x18\x05 \x01(\v2\x1a.google.protobuf.BoolValueR\voffloadable\"\x94\x03\n" +
	"\n" +
	"PowerLevel\x12\x1a\n" +
	"\bcapacity\x18\x01 \x01(\x02R\bcapacity\x12\x1c\n" +
	"\tremaining\x18\x02 \x01(\x02R\tremaining\x12+\n" +
	"\x11percent_remaining\x18\x03 \x01(\x02R\x10percentRemaining\x126\n" +
	"\avoltage\x18\x04 \x01(\v2\x1c.google.protobuf.DoubleValueR\avoltage\x12?\n" +
	"\fcurrent_amps\x18\x05 \x01(\v2\x1c.google.protobuf.DoubleValueR\vcurrentAmps\x12P\n" +
	"\x16run_time_to_empty_mins\x18\x06 \x01(\v2\x1c.google.protobuf.DoubleValueR\x12runTimeToEmptyMins\x12T\n" +
	"\x18consumption_rate_l_per_s\x18\a \x01(\v2\x1c.google.protobuf.DoubleValueR\x14consumptionRateLPerS*\xae\x01\n" +
	"\vPowerStatus\x12\x18\n" +
	"\x14POWER_STATUS_INVALID\x10\x00\x12\x18\n" +
	"\x14POWER_STATUS_UNKNOWN\x10\x01\x12\x1c\n" +
	"\x18POWER_STATUS_NOT_PRESENT\x10\x02\x12\x1a\n" +
	"\x16POWER_STATUS_OPERATING\x10\x03\x12\x19\n" +
	"\x15POWER_STATUS_DISABLED\x10\x04\x12\x16\n" +
	"\x12POWER_STATUS_ERROR\x10\x05*g\n" +
	"\tPowerType\x12\x16\n" +
	"\x12POWER_TYPE_INVALID\x10\x00\x12\x16\n" +
	"\x12POWER_TYPE_UNKNOWN\x10\x01\x12\x12\n" +
	"\x0ePOWER_TYPE_GAS\x10\x02\x12\x16\n" +
	"\x12POWER_TYPE_BATTERY\x10\x03B\xff\x01\n" +
	"\x1ccom.anduril.entitymanager.v1B\rPowerPubProtoP\x01ZNgithub.com/anduril/lattice-sdk-go/src/anduril/entitymanager/v1;entitymanagerv1\xa2\x02\x03AEX\xaa\x02\x18Anduril.Entitymanager.V1\xca\x02\x18Anduril\\Entitymanager\\V1\xe2\x02$Anduril\\Entitymanager\\V1\\GPBMetadata\xea\x02\x1aAnduril::Entitymanager::V1b\x06proto3"

var (
	file_anduril_entitymanager_v1_power_pub_proto_rawDescOnce sync.Once
	file_anduril_entitymanager_v1_power_pub_proto_rawDescData []byte
)

func file_anduril_entitymanager_v1_power_pub_proto_rawDescGZIP() []byte {
	file_anduril_entitymanager_v1_power_pub_proto_rawDescOnce.Do(func() {
		file_anduril_entitymanager_v1_power_pub_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_anduril_entitymanager_v1_power_pub_proto_rawDesc), len(file_anduril_entitymanager_v1_power_pub_proto_rawDesc)))
	})
	return file_anduril_entitymanager_v1_power_pub_proto_rawDescData
}

var file_anduril_entitymanager_v1_power_pub_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_anduril_entitymanager_v1_power_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_anduril_entitymanager_v1_power_pub_proto_goTypes = []any{
	(PowerStatus)(0),               // 0: anduril.entitymanager.v1.PowerStatus
	(PowerType)(0),                 // 1: anduril.entitymanager.v1.PowerType
	(*PowerState)(nil),             // 2: anduril.entitymanager.v1.PowerState
	(*PowerSource)(nil),            // 3: anduril.entitymanager.v1.PowerSource
	(*PowerLevel)(nil),             // 4: anduril.entitymanager.v1.PowerLevel
	nil,                            // 5: anduril.entitymanager.v1.PowerState.SourceIdToStateEntry
	(*wrapperspb.BoolValue)(nil),   // 6: google.protobuf.BoolValue
	(*wrapperspb.DoubleValue)(nil), // 7: google.protobuf.DoubleValue
}
var file_anduril_entitymanager_v1_power_pub_proto_depIdxs = []int32{
	5,  // 0: anduril.entitymanager.v1.PowerState.source_id_to_state:type_name -> anduril.entitymanager.v1.PowerState.SourceIdToStateEntry
	0,  // 1: anduril.entitymanager.v1.PowerSource.power_status:type_name -> anduril.entitymanager.v1.PowerStatus
	1,  // 2: anduril.entitymanager.v1.PowerSource.power_type:type_name -> anduril.entitymanager.v1.PowerType
	4,  // 3: anduril.entitymanager.v1.PowerSource.power_level:type_name -> anduril.entitymanager.v1.PowerLevel
	6,  // 4: anduril.entitymanager.v1.PowerSource.offloadable:type_name -> google.protobuf.BoolValue
	7,  // 5: anduril.entitymanager.v1.PowerLevel.voltage:type_name -> google.protobuf.DoubleValue
	7,  // 6: anduril.entitymanager.v1.PowerLevel.current_amps:type_name -> google.protobuf.DoubleValue
	7,  // 7: anduril.entitymanager.v1.PowerLevel.run_time_to_empty_mins:type_name -> google.protobuf.DoubleValue
	7,  // 8: anduril.entitymanager.v1.PowerLevel.consumption_rate_l_per_s:type_name -> google.protobuf.DoubleValue
	3,  // 9: anduril.entitymanager.v1.PowerState.SourceIdToStateEntry.value:type_name -> anduril.entitymanager.v1.PowerSource
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_anduril_entitymanager_v1_power_pub_proto_init() }
func file_anduril_entitymanager_v1_power_pub_proto_init() {
	if File_anduril_entitymanager_v1_power_pub_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_anduril_entitymanager_v1_power_pub_proto_rawDesc), len(file_anduril_entitymanager_v1_power_pub_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_entitymanager_v1_power_pub_proto_goTypes,
		DependencyIndexes: file_anduril_entitymanager_v1_power_pub_proto_depIdxs,
		EnumInfos:         file_anduril_entitymanager_v1_power_pub_proto_enumTypes,
		MessageInfos:      file_anduril_entitymanager_v1_power_pub_proto_msgTypes,
	}.Build()
	File_anduril_entitymanager_v1_power_pub_proto = out.File
	file_anduril_entitymanager_v1_power_pub_proto_goTypes = nil
	file_anduril_entitymanager_v1_power_pub_proto_depIdxs = nil
}
