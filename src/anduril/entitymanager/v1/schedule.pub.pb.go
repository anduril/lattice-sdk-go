// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: anduril/entitymanager/v1/schedule.pub.proto

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

// The type of Schedule.
type ScheduleType int32

const (
	ScheduleType_SCHEDULE_TYPE_INVALID           ScheduleType = 0
	ScheduleType_SCHEDULE_TYPE_ZONE_ENABLED      ScheduleType = 1
	ScheduleType_SCHEDULE_TYPE_ZONE_TEMP_ENABLED ScheduleType = 2
)

// Enum value maps for ScheduleType.
var (
	ScheduleType_name = map[int32]string{
		0: "SCHEDULE_TYPE_INVALID",
		1: "SCHEDULE_TYPE_ZONE_ENABLED",
		2: "SCHEDULE_TYPE_ZONE_TEMP_ENABLED",
	}
	ScheduleType_value = map[string]int32{
		"SCHEDULE_TYPE_INVALID":           0,
		"SCHEDULE_TYPE_ZONE_ENABLED":      1,
		"SCHEDULE_TYPE_ZONE_TEMP_ENABLED": 2,
	}
)

func (x ScheduleType) Enum() *ScheduleType {
	p := new(ScheduleType)
	*p = x
	return p
}

func (x ScheduleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScheduleType) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_schedule_pub_proto_enumTypes[0].Descriptor()
}

func (ScheduleType) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_schedule_pub_proto_enumTypes[0]
}

func (x ScheduleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScheduleType.Descriptor instead.
func (ScheduleType) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_schedule_pub_proto_rawDescGZIP(), []int{0}
}

// Schedules associated with this entity
type Schedules struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Schedules     []*Schedule            `protobuf:"bytes,1,rep,name=schedules,proto3" json:"schedules,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Schedules) Reset() {
	*x = Schedules{}
	mi := &file_anduril_entitymanager_v1_schedule_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Schedules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedules) ProtoMessage() {}

func (x *Schedules) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_schedule_pub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedules.ProtoReflect.Descriptor instead.
func (*Schedules) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_schedule_pub_proto_rawDescGZIP(), []int{0}
}

func (x *Schedules) GetSchedules() []*Schedule {
	if x != nil {
		return x.Schedules
	}
	return nil
}

// A Schedule associated with this entity
type Schedule struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// expression that represents this schedule's "ON" state
	Windows []*CronWindow `protobuf:"bytes,1,rep,name=windows,proto3" json:"windows,omitempty"`
	// A unique identifier for this schedule.
	ScheduleId string `protobuf:"bytes,2,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	// The schedule type
	ScheduleType  ScheduleType `protobuf:"varint,3,opt,name=schedule_type,json=scheduleType,proto3,enum=anduril.entitymanager.v1.ScheduleType" json:"schedule_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	mi := &file_anduril_entitymanager_v1_schedule_pub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_schedule_pub_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_schedule_pub_proto_rawDescGZIP(), []int{1}
}

func (x *Schedule) GetWindows() []*CronWindow {
	if x != nil {
		return x.Windows
	}
	return nil
}

func (x *Schedule) GetScheduleId() string {
	if x != nil {
		return x.ScheduleId
	}
	return ""
}

func (x *Schedule) GetScheduleType() ScheduleType {
	if x != nil {
		return x.ScheduleType
	}
	return ScheduleType_SCHEDULE_TYPE_INVALID
}

type CronWindow struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// in UTC, describes when and at what cadence this window starts, in the quartz flavor of cron
	//
	// examples:
	//
	//	This schedule is begins at 7:00:00am UTC everyday between Monday and Friday
	//	    0 0 7 ? * MON-FRI *
	//	This schedule begins every 5 minutes starting at 12:00:00pm UTC until 8:00:00pm UTC everyday
	//	    0 0/5 12-20 * * ? *
	//	This schedule begins at 12:00:00pm UTC on March 2nd 2023
	//	    0 0 12 2 3 ? 2023
	CronExpression string `protobuf:"bytes,1,opt,name=cron_expression,json=cronExpression,proto3" json:"cron_expression,omitempty"`
	// describes the duration
	DurationMillis uint64 `protobuf:"varint,2,opt,name=duration_millis,json=durationMillis,proto3" json:"duration_millis,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CronWindow) Reset() {
	*x = CronWindow{}
	mi := &file_anduril_entitymanager_v1_schedule_pub_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CronWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronWindow) ProtoMessage() {}

func (x *CronWindow) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_schedule_pub_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronWindow.ProtoReflect.Descriptor instead.
func (*CronWindow) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_schedule_pub_proto_rawDescGZIP(), []int{2}
}

func (x *CronWindow) GetCronExpression() string {
	if x != nil {
		return x.CronExpression
	}
	return ""
}

func (x *CronWindow) GetDurationMillis() uint64 {
	if x != nil {
		return x.DurationMillis
	}
	return 0
}

var File_anduril_entitymanager_v1_schedule_pub_proto protoreflect.FileDescriptor

var file_anduril_entitymanager_v1_schedule_pub_proto_rawDesc = string([]byte{
	0x0a, 0x2b, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x4d, 0x0a, 0x09, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x09, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22, 0xb8, 0x01, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x6f, 0x6e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x52, 0x07, 0x77, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x61, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x5e, 0x0a, 0x0a, 0x43, 0x72, 0x6f, 0x6e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12,
	0x27, 0x0a, 0x0f, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x6f, 0x6e, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6c, 0x6c, 0x69,
	0x73, 0x2a, 0x6e, 0x0a, 0x0c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a,
	0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x5a, 0x4f,
	0x4e, 0x45, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f,
	0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x5a, 0x4f,
	0x4e, 0x45, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10,
	0x02, 0x42, 0x82, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x42, 0x10, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x75, 0x62, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x6c, 0x61, 0x74, 0x74, 0x69,
	0x63, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x45, 0x58, 0xaa, 0x02, 0x18, 0x41,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x24, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x41, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_anduril_entitymanager_v1_schedule_pub_proto_rawDescOnce sync.Once
	file_anduril_entitymanager_v1_schedule_pub_proto_rawDescData []byte
)

func file_anduril_entitymanager_v1_schedule_pub_proto_rawDescGZIP() []byte {
	file_anduril_entitymanager_v1_schedule_pub_proto_rawDescOnce.Do(func() {
		file_anduril_entitymanager_v1_schedule_pub_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_anduril_entitymanager_v1_schedule_pub_proto_rawDesc), len(file_anduril_entitymanager_v1_schedule_pub_proto_rawDesc)))
	})
	return file_anduril_entitymanager_v1_schedule_pub_proto_rawDescData
}

var file_anduril_entitymanager_v1_schedule_pub_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_anduril_entitymanager_v1_schedule_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_anduril_entitymanager_v1_schedule_pub_proto_goTypes = []any{
	(ScheduleType)(0),  // 0: anduril.entitymanager.v1.ScheduleType
	(*Schedules)(nil),  // 1: anduril.entitymanager.v1.Schedules
	(*Schedule)(nil),   // 2: anduril.entitymanager.v1.Schedule
	(*CronWindow)(nil), // 3: anduril.entitymanager.v1.CronWindow
}
var file_anduril_entitymanager_v1_schedule_pub_proto_depIdxs = []int32{
	2, // 0: anduril.entitymanager.v1.Schedules.schedules:type_name -> anduril.entitymanager.v1.Schedule
	3, // 1: anduril.entitymanager.v1.Schedule.windows:type_name -> anduril.entitymanager.v1.CronWindow
	0, // 2: anduril.entitymanager.v1.Schedule.schedule_type:type_name -> anduril.entitymanager.v1.ScheduleType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_anduril_entitymanager_v1_schedule_pub_proto_init() }
func file_anduril_entitymanager_v1_schedule_pub_proto_init() {
	if File_anduril_entitymanager_v1_schedule_pub_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_anduril_entitymanager_v1_schedule_pub_proto_rawDesc), len(file_anduril_entitymanager_v1_schedule_pub_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_entitymanager_v1_schedule_pub_proto_goTypes,
		DependencyIndexes: file_anduril_entitymanager_v1_schedule_pub_proto_depIdxs,
		EnumInfos:         file_anduril_entitymanager_v1_schedule_pub_proto_enumTypes,
		MessageInfos:      file_anduril_entitymanager_v1_schedule_pub_proto_msgTypes,
	}.Build()
	File_anduril_entitymanager_v1_schedule_pub_proto = out.File
	file_anduril_entitymanager_v1_schedule_pub_proto_goTypes = nil
	file_anduril_entitymanager_v1_schedule_pub_proto_depIdxs = nil
}
