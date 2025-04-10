// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: anduril/entitymanager/v1/transponder_codes.pub.proto

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

// Indicates the interrogation status of a target.
type InterrogationResponse int32

const (
	// Note that INTERROGATION_INVALID indicates that the target has not been interrogated.
	InterrogationResponse_INTERROGATION_RESPONSE_INVALID     InterrogationResponse = 0
	InterrogationResponse_INTERROGATION_RESPONSE_CORRECT     InterrogationResponse = 1
	InterrogationResponse_INTERROGATION_RESPONSE_INCORRECT   InterrogationResponse = 2
	InterrogationResponse_INTERROGATION_RESPONSE_NO_RESPONSE InterrogationResponse = 3
)

// Enum value maps for InterrogationResponse.
var (
	InterrogationResponse_name = map[int32]string{
		0: "INTERROGATION_RESPONSE_INVALID",
		1: "INTERROGATION_RESPONSE_CORRECT",
		2: "INTERROGATION_RESPONSE_INCORRECT",
		3: "INTERROGATION_RESPONSE_NO_RESPONSE",
	}
	InterrogationResponse_value = map[string]int32{
		"INTERROGATION_RESPONSE_INVALID":     0,
		"INTERROGATION_RESPONSE_CORRECT":     1,
		"INTERROGATION_RESPONSE_INCORRECT":   2,
		"INTERROGATION_RESPONSE_NO_RESPONSE": 3,
	}
)

func (x InterrogationResponse) Enum() *InterrogationResponse {
	p := new(InterrogationResponse)
	*p = x
	return p
}

func (x InterrogationResponse) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InterrogationResponse) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_transponder_codes_pub_proto_enumTypes[0].Descriptor()
}

func (InterrogationResponse) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_transponder_codes_pub_proto_enumTypes[0]
}

func (x InterrogationResponse) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InterrogationResponse.Descriptor instead.
func (InterrogationResponse) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_transponder_codes_pub_proto_rawDescGZIP(), []int{0}
}

// A message describing any transponder codes associated with Mode 1, 2, 3, 4, 5, S interrogations.
type TransponderCodes struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The mode 1 code assigned to military assets.
	Mode1 uint32 `protobuf:"varint,1,opt,name=mode1,proto3" json:"mode1,omitempty"`
	// The Mode 2 code assigned to military assets.
	Mode2 uint32 `protobuf:"varint,2,opt,name=mode2,proto3" json:"mode2,omitempty"`
	// The Mode 3 code assigned by ATC to the asset.
	Mode3 uint32 `protobuf:"varint,3,opt,name=mode3,proto3" json:"mode3,omitempty"`
	// The validity of the response from the Mode 4 interrogation.
	Mode4InterrogationResponse InterrogationResponse `protobuf:"varint,4,opt,name=mode4_interrogation_response,json=mode4InterrogationResponse,proto3,enum=anduril.entitymanager.v1.InterrogationResponse" json:"mode4_interrogation_response,omitempty"`
	// The Mode 5 transponder codes.
	Mode5 *Mode5 `protobuf:"bytes,5,opt,name=mode5,proto3" json:"mode5,omitempty"`
	// The Mode S transponder codes.
	ModeS         *ModeS `protobuf:"bytes,6,opt,name=mode_s,json=modeS,proto3" json:"mode_s,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransponderCodes) Reset() {
	*x = TransponderCodes{}
	mi := &file_anduril_entitymanager_v1_transponder_codes_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransponderCodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransponderCodes) ProtoMessage() {}

func (x *TransponderCodes) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_transponder_codes_pub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransponderCodes.ProtoReflect.Descriptor instead.
func (*TransponderCodes) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_transponder_codes_pub_proto_rawDescGZIP(), []int{0}
}

func (x *TransponderCodes) GetMode1() uint32 {
	if x != nil {
		return x.Mode1
	}
	return 0
}

func (x *TransponderCodes) GetMode2() uint32 {
	if x != nil {
		return x.Mode2
	}
	return 0
}

func (x *TransponderCodes) GetMode3() uint32 {
	if x != nil {
		return x.Mode3
	}
	return 0
}

func (x *TransponderCodes) GetMode4InterrogationResponse() InterrogationResponse {
	if x != nil {
		return x.Mode4InterrogationResponse
	}
	return InterrogationResponse_INTERROGATION_RESPONSE_INVALID
}

func (x *TransponderCodes) GetMode5() *Mode5 {
	if x != nil {
		return x.Mode5
	}
	return nil
}

func (x *TransponderCodes) GetModeS() *ModeS {
	if x != nil {
		return x.ModeS
	}
	return nil
}

// Describes the Mode 5 transponder interrogation status and codes.
type Mode5 struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The validity of the response from the Mode 5 interrogation.
	Mode5InterrogationResponse InterrogationResponse `protobuf:"varint,1,opt,name=mode5_interrogation_response,json=mode5InterrogationResponse,proto3,enum=anduril.entitymanager.v1.InterrogationResponse" json:"mode5_interrogation_response,omitempty"`
	// The Mode 5 code assigned to military assets.
	Mode5 uint32 `protobuf:"varint,2,opt,name=mode5,proto3" json:"mode5,omitempty"`
	// The Mode 5 platform identification code.
	Mode5PlatformId uint32 `protobuf:"varint,3,opt,name=mode5_platform_id,json=mode5PlatformId,proto3" json:"mode5_platform_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Mode5) Reset() {
	*x = Mode5{}
	mi := &file_anduril_entitymanager_v1_transponder_codes_pub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Mode5) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mode5) ProtoMessage() {}

func (x *Mode5) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_transponder_codes_pub_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mode5.ProtoReflect.Descriptor instead.
func (*Mode5) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_transponder_codes_pub_proto_rawDescGZIP(), []int{1}
}

func (x *Mode5) GetMode5InterrogationResponse() InterrogationResponse {
	if x != nil {
		return x.Mode5InterrogationResponse
	}
	return InterrogationResponse_INTERROGATION_RESPONSE_INVALID
}

func (x *Mode5) GetMode5() uint32 {
	if x != nil {
		return x.Mode5
	}
	return 0
}

func (x *Mode5) GetMode5PlatformId() uint32 {
	if x != nil {
		return x.Mode5PlatformId
	}
	return 0
}

// Describes the Mode S codes.
type ModeS struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Mode S identifier which comprises of 8 alphanumeric characters.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The Mode S ICAO aircraft address. Expected values are between 1 and 16777214 decimal. The Mode S address is
	// considered unique.
	Address       uint32 `protobuf:"varint,2,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ModeS) Reset() {
	*x = ModeS{}
	mi := &file_anduril_entitymanager_v1_transponder_codes_pub_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ModeS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModeS) ProtoMessage() {}

func (x *ModeS) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_transponder_codes_pub_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModeS.ProtoReflect.Descriptor instead.
func (*ModeS) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_transponder_codes_pub_proto_rawDescGZIP(), []int{2}
}

func (x *ModeS) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ModeS) GetAddress() uint32 {
	if x != nil {
		return x.Address
	}
	return 0
}

var File_anduril_entitymanager_v1_transponder_codes_pub_proto protoreflect.FileDescriptor

const file_anduril_entitymanager_v1_transponder_codes_pub_proto_rawDesc = "" +
	"\n" +
	"4anduril/entitymanager/v1/transponder_codes.pub.proto\x12\x18anduril.entitymanager.v1\x1a*anduril/entitymanager/v1/options.pub.proto\"\xd4\x02\n" +
	"\x10TransponderCodes\x12\x19\n" +
	"\x05mode1\x18\x01 \x01(\rB\x03\xc8>\x01R\x05mode1\x12\x19\n" +
	"\x05mode2\x18\x02 \x01(\rB\x03\xc8>\x01R\x05mode2\x12\x19\n" +
	"\x05mode3\x18\x03 \x01(\rB\x03\xc8>\x01R\x05mode3\x12v\n" +
	"\x1cmode4_interrogation_response\x18\x04 \x01(\x0e2/.anduril.entitymanager.v1.InterrogationResponseB\x03\xc8>\x01R\x1amode4InterrogationResponse\x12:\n" +
	"\x05mode5\x18\x05 \x01(\v2\x1f.anduril.entitymanager.v1.Mode5B\x03\xc8>\x01R\x05mode5\x12;\n" +
	"\x06mode_s\x18\x06 \x01(\v2\x1f.anduril.entitymanager.v1.ModeSB\x03\xc8>\x01R\x05modeS\"\xbc\x01\n" +
	"\x05Mode5\x12q\n" +
	"\x1cmode5_interrogation_response\x18\x01 \x01(\x0e2/.anduril.entitymanager.v1.InterrogationResponseR\x1amode5InterrogationResponse\x12\x14\n" +
	"\x05mode5\x18\x02 \x01(\rR\x05mode5\x12*\n" +
	"\x11mode5_platform_id\x18\x03 \x01(\rR\x0fmode5PlatformId\"1\n" +
	"\x05ModeS\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x18\n" +
	"\aaddress\x18\x02 \x01(\rR\aaddress*\xad\x01\n" +
	"\x15InterrogationResponse\x12\"\n" +
	"\x1eINTERROGATION_RESPONSE_INVALID\x10\x00\x12\"\n" +
	"\x1eINTERROGATION_RESPONSE_CORRECT\x10\x01\x12$\n" +
	" INTERROGATION_RESPONSE_INCORRECT\x10\x02\x12&\n" +
	"\"INTERROGATION_RESPONSE_NO_RESPONSE\x10\x03B\x8a\x02\n" +
	"\x1ccom.anduril.entitymanager.v1B\x18TransponderCodesPubProtoP\x01ZNgithub.com/anduril/lattice-sdk-go/src/anduril/entitymanager/v1;entitymanagerv1\xa2\x02\x03AEX\xaa\x02\x18Anduril.Entitymanager.V1\xca\x02\x18Anduril\\Entitymanager\\V1\xe2\x02$Anduril\\Entitymanager\\V1\\GPBMetadata\xea\x02\x1aAnduril::Entitymanager::V1b\x06proto3"

var (
	file_anduril_entitymanager_v1_transponder_codes_pub_proto_rawDescOnce sync.Once
	file_anduril_entitymanager_v1_transponder_codes_pub_proto_rawDescData []byte
)

func file_anduril_entitymanager_v1_transponder_codes_pub_proto_rawDescGZIP() []byte {
	file_anduril_entitymanager_v1_transponder_codes_pub_proto_rawDescOnce.Do(func() {
		file_anduril_entitymanager_v1_transponder_codes_pub_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_anduril_entitymanager_v1_transponder_codes_pub_proto_rawDesc), len(file_anduril_entitymanager_v1_transponder_codes_pub_proto_rawDesc)))
	})
	return file_anduril_entitymanager_v1_transponder_codes_pub_proto_rawDescData
}

var file_anduril_entitymanager_v1_transponder_codes_pub_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_anduril_entitymanager_v1_transponder_codes_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_anduril_entitymanager_v1_transponder_codes_pub_proto_goTypes = []any{
	(InterrogationResponse)(0), // 0: anduril.entitymanager.v1.InterrogationResponse
	(*TransponderCodes)(nil),   // 1: anduril.entitymanager.v1.TransponderCodes
	(*Mode5)(nil),              // 2: anduril.entitymanager.v1.Mode5
	(*ModeS)(nil),              // 3: anduril.entitymanager.v1.ModeS
}
var file_anduril_entitymanager_v1_transponder_codes_pub_proto_depIdxs = []int32{
	0, // 0: anduril.entitymanager.v1.TransponderCodes.mode4_interrogation_response:type_name -> anduril.entitymanager.v1.InterrogationResponse
	2, // 1: anduril.entitymanager.v1.TransponderCodes.mode5:type_name -> anduril.entitymanager.v1.Mode5
	3, // 2: anduril.entitymanager.v1.TransponderCodes.mode_s:type_name -> anduril.entitymanager.v1.ModeS
	0, // 3: anduril.entitymanager.v1.Mode5.mode5_interrogation_response:type_name -> anduril.entitymanager.v1.InterrogationResponse
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_anduril_entitymanager_v1_transponder_codes_pub_proto_init() }
func file_anduril_entitymanager_v1_transponder_codes_pub_proto_init() {
	if File_anduril_entitymanager_v1_transponder_codes_pub_proto != nil {
		return
	}
	file_anduril_entitymanager_v1_options_pub_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_anduril_entitymanager_v1_transponder_codes_pub_proto_rawDesc), len(file_anduril_entitymanager_v1_transponder_codes_pub_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_entitymanager_v1_transponder_codes_pub_proto_goTypes,
		DependencyIndexes: file_anduril_entitymanager_v1_transponder_codes_pub_proto_depIdxs,
		EnumInfos:         file_anduril_entitymanager_v1_transponder_codes_pub_proto_enumTypes,
		MessageInfos:      file_anduril_entitymanager_v1_transponder_codes_pub_proto_msgTypes,
	}.Build()
	File_anduril_entitymanager_v1_transponder_codes_pub_proto = out.File
	file_anduril_entitymanager_v1_transponder_codes_pub_proto_goTypes = nil
	file_anduril_entitymanager_v1_transponder_codes_pub_proto_depIdxs = nil
}
