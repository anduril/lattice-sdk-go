// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: anduril/tasks/ads/thirdparty/v1/formation.pub.proto

package thirdpartyv1

import (
	v2 "github.com/anduril/lattice-sdk-go/src/anduril/tasks/v2"
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

// Maps to a Line formation of assets with a speed. This is a simple line with two LLAs.
type LineFormation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Line start
	LineStart *v2.Objective `protobuf:"bytes,1,opt,name=line_start,json=lineStart,proto3" json:"line_start,omitempty"`
	// Line end
	LineEnd *v2.Objective `protobuf:"bytes,2,opt,name=line_end,json=lineEnd,proto3" json:"line_end,omitempty"`
	// Speed in Meters/Second to get in Line Formation
	SurfaceSpeedMs *wrapperspb.DoubleValue `protobuf:"bytes,3,opt,name=surface_speed_ms,json=surfaceSpeedMs,proto3" json:"surface_speed_ms,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *LineFormation) Reset() {
	*x = LineFormation{}
	mi := &file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LineFormation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineFormation) ProtoMessage() {}

func (x *LineFormation) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineFormation.ProtoReflect.Descriptor instead.
func (*LineFormation) Descriptor() ([]byte, []int) {
	return file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_rawDescGZIP(), []int{0}
}

func (x *LineFormation) GetLineStart() *v2.Objective {
	if x != nil {
		return x.LineStart
	}
	return nil
}

func (x *LineFormation) GetLineEnd() *v2.Objective {
	if x != nil {
		return x.LineEnd
	}
	return nil
}

func (x *LineFormation) GetSurfaceSpeedMs() *wrapperspb.DoubleValue {
	if x != nil {
		return x.SurfaceSpeedMs
	}
	return nil
}

var File_anduril_tasks_ads_thirdparty_v1_formation_pub_proto protoreflect.FileDescriptor

const file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_rawDesc = "" +
	"\n" +
	"3anduril/tasks/ads/thirdparty/v1/formation.pub.proto\x12\x1fanduril.tasks.ads.thirdparty.v1\x1a$anduril/tasks/v2/objective.pub.proto\x1a\x1egoogle/protobuf/wrappers.proto\"\xcb\x01\n" +
	"\rLineFormation\x12:\n" +
	"\n" +
	"line_start\x18\x01 \x01(\v2\x1b.anduril.tasks.v2.ObjectiveR\tlineStart\x126\n" +
	"\bline_end\x18\x02 \x01(\v2\x1b.anduril.tasks.v2.ObjectiveR\alineEnd\x12F\n" +
	"\x10surface_speed_ms\x18\x03 \x01(\v2\x1c.google.protobuf.DoubleValueR\x0esurfaceSpeedMsB\xad\x02\n" +
	"#com.anduril.tasks.ads.thirdparty.v1B\x11FormationPubProtoP\x01ZRgithub.com/anduril/lattice-sdk-go/src/anduril/tasks/ads/thirdparty/v1;thirdpartyv1\xa2\x02\x04ATAT\xaa\x02\x1fAnduril.Tasks.Ads.Thirdparty.V1\xca\x02\x1fAnduril\\Tasks\\Ads\\Thirdparty\\V1\xe2\x02+Anduril\\Tasks\\Ads\\Thirdparty\\V1\\GPBMetadata\xea\x02#Anduril::Tasks::Ads::Thirdparty::V1b\x06proto3"

var (
	file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_rawDescOnce sync.Once
	file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_rawDescData []byte
)

func file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_rawDescGZIP() []byte {
	file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_rawDescOnce.Do(func() {
		file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_rawDesc), len(file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_rawDesc)))
	})
	return file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_rawDescData
}

var file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_goTypes = []any{
	(*LineFormation)(nil),          // 0: anduril.tasks.ads.thirdparty.v1.LineFormation
	(*v2.Objective)(nil),           // 1: anduril.tasks.v2.Objective
	(*wrapperspb.DoubleValue)(nil), // 2: google.protobuf.DoubleValue
}
var file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_depIdxs = []int32{
	1, // 0: anduril.tasks.ads.thirdparty.v1.LineFormation.line_start:type_name -> anduril.tasks.v2.Objective
	1, // 1: anduril.tasks.ads.thirdparty.v1.LineFormation.line_end:type_name -> anduril.tasks.v2.Objective
	2, // 2: anduril.tasks.ads.thirdparty.v1.LineFormation.surface_speed_ms:type_name -> google.protobuf.DoubleValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_init() }
func file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_init() {
	if File_anduril_tasks_ads_thirdparty_v1_formation_pub_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_rawDesc), len(file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_goTypes,
		DependencyIndexes: file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_depIdxs,
		MessageInfos:      file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_msgTypes,
	}.Build()
	File_anduril_tasks_ads_thirdparty_v1_formation_pub_proto = out.File
	file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_goTypes = nil
	file_anduril_tasks_ads_thirdparty_v1_formation_pub_proto_depIdxs = nil
}
