// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: anduril/entitymanager/v1/types.pub.proto

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

// The type of alternate id.
type AltIdType int32

const (
	AltIdType_ALT_ID_TYPE_INVALID AltIdType = 0
	// an Anduril trackId_2
	AltIdType_ALT_ID_TYPE_TRACK_ID_2 AltIdType = 1
	// an Anduril trackId_1
	AltIdType_ALT_ID_TYPE_TRACK_ID_1 AltIdType = 12
	// an Anduril Sensor Point of Interest ID
	AltIdType_ALT_ID_TYPE_SPI_ID AltIdType = 2
	// NITF file title
	AltIdType_ALT_ID_TYPE_NITF_FILE_TITLE AltIdType = 3
	// Track repo alert ID
	AltIdType_ALT_ID_TYPE_TRACK_REPO_ALERT_ID AltIdType = 4
	// an Anduril AssetId v2
	AltIdType_ALT_ID_TYPE_ASSET_ID AltIdType = 5
	// Use for Link 16 track identifiers for non-JTIDS Unit entities.
	AltIdType_ALT_ID_TYPE_LINK16_TRACK_NUMBER AltIdType = 6
	// Use for Link 16 JTIDS Unit identifiers.
	AltIdType_ALT_ID_TYPE_LINK16_JU AltIdType = 7
	// an NCCT message ID
	AltIdType_ALT_ID_TYPE_NCCT_MESSAGE_ID AltIdType = 8
	// callsign for the entity. e.g. a TAK callsign or an aircraft callsign
	AltIdType_ALT_ID_TYPE_CALLSIGN AltIdType = 9
	// the Maritime Mobile Service Identity for a maritime object (vessel, offshore installation, etc.)
	AltIdType_ALT_ID_TYPE_MMSI_ID AltIdType = 10
	// A VMF URN that uniquely identifies the URN on the VMF network.
	AltIdType_ALT_ID_TYPE_VMF_URN AltIdType = 11
	// the International Maritime Organization number for identifying maritime objects (vessel, offshore installation, etc.)
	AltIdType_ALT_ID_TYPE_IMO_ID AltIdType = 13
	// A VMF target number that uniquely identifies the target on the VMF network
	AltIdType_ALT_ID_TYPE_VMF_TARGET_NUMBER AltIdType = 14
	// A serial number that uniquely identifies the entity and is permanently associated with only one entity. This
	// identifier is assigned by some authority and only ever identifies a single thing. Examples include a
	// Vehicle Identification Number (VIN) or ship hull identification number (hull number). This is a generalized
	// component and should not be used if a more specific registration type is already defined (i.e., ALT_ID_TYPE_VMF_URN).
	AltIdType_ALT_ID_TYPE_SERIAL_NUMBER AltIdType = 15
	// A registration identifier assigned by a local or national authority. This identifier is not permanently fixed
	// to one specific entity and may be reassigned on change of ownership, destruction, or other conditions set
	// forth by the authority. Examples include a vehicle license plate or aircraft tail number. This is a generalized
	// component and should not be used if a more specific registration type is already defined (i.e., ALT_ID_TYPE_IMO_ID).
	AltIdType_ALT_ID_TYPE_REGISTRATION_ID AltIdType = 16
	// Integrated Broadcast Service Common Message Format Global Identifier
	AltIdType_ALT_ID_TYPE_IBS_GID AltIdType = 17
	// Department of Defense Activity Address Code.
	AltIdType_ALT_ID_TYPE_DODAAC AltIdType = 18
	// Unit Identification Code uniquely identifies each US Department of Defense entity
	AltIdType_ALT_ID_TYPE_UIC AltIdType = 19
	// A NORAD Satellite Catalog Number, a 9-digit number uniquely representing orbital objects around Earth.
	// of strictly numeric.
	AltIdType_ALT_ID_TYPE_NORAD_CAT_ID AltIdType = 20
	// Space object name. If populated, use names from the UN Office
	// of Outer Space Affairs designator index, otherwise set field to UNKNOWN.
	AltIdType_ALT_ID_TYPE_UNOOSA_NAME AltIdType = 23
	// Space object identifier. If populated, use the international spacecraft designator
	// as published in the UN Office of Outer Space Affairs designator index, otherwise set to UNKNOWN.
	// Recommended values have the format YYYYNNNP{PP}, where:
	//
	//	YYYY = Year of launch.
	//	NNN = Three-digit serial number of launch
	//	in year YYYY (with leading zeros).
	//	P{PP} = At least one capital letter for the
	//	identification of the part brought
	//	into space by the launch.
	AltIdType_ALT_ID_TYPE_UNOOSA_ID AltIdType = 24
)

// Enum value maps for AltIdType.
var (
	AltIdType_name = map[int32]string{
		0:  "ALT_ID_TYPE_INVALID",
		1:  "ALT_ID_TYPE_TRACK_ID_2",
		12: "ALT_ID_TYPE_TRACK_ID_1",
		2:  "ALT_ID_TYPE_SPI_ID",
		3:  "ALT_ID_TYPE_NITF_FILE_TITLE",
		4:  "ALT_ID_TYPE_TRACK_REPO_ALERT_ID",
		5:  "ALT_ID_TYPE_ASSET_ID",
		6:  "ALT_ID_TYPE_LINK16_TRACK_NUMBER",
		7:  "ALT_ID_TYPE_LINK16_JU",
		8:  "ALT_ID_TYPE_NCCT_MESSAGE_ID",
		9:  "ALT_ID_TYPE_CALLSIGN",
		10: "ALT_ID_TYPE_MMSI_ID",
		11: "ALT_ID_TYPE_VMF_URN",
		13: "ALT_ID_TYPE_IMO_ID",
		14: "ALT_ID_TYPE_VMF_TARGET_NUMBER",
		15: "ALT_ID_TYPE_SERIAL_NUMBER",
		16: "ALT_ID_TYPE_REGISTRATION_ID",
		17: "ALT_ID_TYPE_IBS_GID",
		18: "ALT_ID_TYPE_DODAAC",
		19: "ALT_ID_TYPE_UIC",
		20: "ALT_ID_TYPE_NORAD_CAT_ID",
		23: "ALT_ID_TYPE_UNOOSA_NAME",
		24: "ALT_ID_TYPE_UNOOSA_ID",
	}
	AltIdType_value = map[string]int32{
		"ALT_ID_TYPE_INVALID":             0,
		"ALT_ID_TYPE_TRACK_ID_2":          1,
		"ALT_ID_TYPE_TRACK_ID_1":          12,
		"ALT_ID_TYPE_SPI_ID":              2,
		"ALT_ID_TYPE_NITF_FILE_TITLE":     3,
		"ALT_ID_TYPE_TRACK_REPO_ALERT_ID": 4,
		"ALT_ID_TYPE_ASSET_ID":            5,
		"ALT_ID_TYPE_LINK16_TRACK_NUMBER": 6,
		"ALT_ID_TYPE_LINK16_JU":           7,
		"ALT_ID_TYPE_NCCT_MESSAGE_ID":     8,
		"ALT_ID_TYPE_CALLSIGN":            9,
		"ALT_ID_TYPE_MMSI_ID":             10,
		"ALT_ID_TYPE_VMF_URN":             11,
		"ALT_ID_TYPE_IMO_ID":              13,
		"ALT_ID_TYPE_VMF_TARGET_NUMBER":   14,
		"ALT_ID_TYPE_SERIAL_NUMBER":       15,
		"ALT_ID_TYPE_REGISTRATION_ID":     16,
		"ALT_ID_TYPE_IBS_GID":             17,
		"ALT_ID_TYPE_DODAAC":              18,
		"ALT_ID_TYPE_UIC":                 19,
		"ALT_ID_TYPE_NORAD_CAT_ID":        20,
		"ALT_ID_TYPE_UNOOSA_NAME":         23,
		"ALT_ID_TYPE_UNOOSA_ID":           24,
	}
)

func (x AltIdType) Enum() *AltIdType {
	p := new(AltIdType)
	*p = x
	return p
}

func (x AltIdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AltIdType) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_types_pub_proto_enumTypes[0].Descriptor()
}

func (AltIdType) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_types_pub_proto_enumTypes[0]
}

func (x AltIdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AltIdType.Descriptor instead.
func (AltIdType) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_types_pub_proto_rawDescGZIP(), []int{0}
}

// Set of possible templates used when creating an entity.
// This impacts minimum required component sets and can be used by edge systems that need to distinguish.
type Template int32

const (
	Template_TEMPLATE_INVALID Template = 0
	// Refers to any detected object.
	// Requires setting the location, and mil_view components.
	Template_TEMPLATE_TRACK Template = 1
	// Refers to any sensors detected at a specific location.
	// Requires setting location, and mil_view.
	Template_TEMPLATE_SENSOR_POINT_OF_INTEREST Template = 2
	// Refers to a taskable entity under the control of friendly forces.
	// Requires setting location, and mil_view, and ontology.
	Template_TEMPLATE_ASSET Template = 3
	// Refers to shapes or points of interest drawn on the map.
	// Requires setting geo_shape and geo_details.
	Template_TEMPLATE_GEO Template = 4
	// Refers to signal detection with characteristics such as emitter notation, frequency, or lines of bearing.
	// Requires setting signal, and mil_view, and ontology. Requies setting location, if the signal.fixed component is populated.
	Template_TEMPLATE_SIGNAL_OF_INTEREST Template = 5
)

// Enum value maps for Template.
var (
	Template_name = map[int32]string{
		0: "TEMPLATE_INVALID",
		1: "TEMPLATE_TRACK",
		2: "TEMPLATE_SENSOR_POINT_OF_INTEREST",
		3: "TEMPLATE_ASSET",
		4: "TEMPLATE_GEO",
		5: "TEMPLATE_SIGNAL_OF_INTEREST",
	}
	Template_value = map[string]int32{
		"TEMPLATE_INVALID":                  0,
		"TEMPLATE_TRACK":                    1,
		"TEMPLATE_SENSOR_POINT_OF_INTEREST": 2,
		"TEMPLATE_ASSET":                    3,
		"TEMPLATE_GEO":                      4,
		"TEMPLATE_SIGNAL_OF_INTEREST":       5,
	}
)

func (x Template) Enum() *Template {
	p := new(Template)
	*p = x
	return p
}

func (x Template) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Template) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_types_pub_proto_enumTypes[1].Descriptor()
}

func (Template) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_types_pub_proto_enumTypes[1]
}

func (x Template) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Template.Descriptor instead.
func (Template) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_types_pub_proto_rawDescGZIP(), []int{1}
}

// The state of an override.
type OverrideStatus int32

const (
	OverrideStatus_OVERRIDE_STATUS_INVALID OverrideStatus = 0
	// the override was applied to the entity.
	OverrideStatus_OVERRIDE_STATUS_APPLIED OverrideStatus = 1
	// the override is pending action.
	OverrideStatus_OVERRIDE_STATUS_PENDING OverrideStatus = 2
	// the override has been timed out.
	OverrideStatus_OVERRIDE_STATUS_TIMEOUT OverrideStatus = 3
	// the override has been rejected
	OverrideStatus_OVERRIDE_STATUS_REJECTED OverrideStatus = 4
	// The override is pending deletion.
	OverrideStatus_OVERRIDE_STATUS_DELETION_PENDING OverrideStatus = 5
)

// Enum value maps for OverrideStatus.
var (
	OverrideStatus_name = map[int32]string{
		0: "OVERRIDE_STATUS_INVALID",
		1: "OVERRIDE_STATUS_APPLIED",
		2: "OVERRIDE_STATUS_PENDING",
		3: "OVERRIDE_STATUS_TIMEOUT",
		4: "OVERRIDE_STATUS_REJECTED",
		5: "OVERRIDE_STATUS_DELETION_PENDING",
	}
	OverrideStatus_value = map[string]int32{
		"OVERRIDE_STATUS_INVALID":          0,
		"OVERRIDE_STATUS_APPLIED":          1,
		"OVERRIDE_STATUS_PENDING":          2,
		"OVERRIDE_STATUS_TIMEOUT":          3,
		"OVERRIDE_STATUS_REJECTED":         4,
		"OVERRIDE_STATUS_DELETION_PENDING": 5,
	}
)

func (x OverrideStatus) Enum() *OverrideStatus {
	p := new(OverrideStatus)
	*p = x
	return p
}

func (x OverrideStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OverrideStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_types_pub_proto_enumTypes[2].Descriptor()
}

func (OverrideStatus) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_types_pub_proto_enumTypes[2]
}

func (x OverrideStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OverrideStatus.Descriptor instead.
func (OverrideStatus) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_types_pub_proto_rawDescGZIP(), []int{2}
}

type OverrideType int32

const (
	// The override type value was not set. This value is interpreted as OVERRIDE_TYPE_LIVE for backward compatibility.
	OverrideType_OVERRIDE_TYPE_INVALID OverrideType = 0
	// Override was requested when the entity was live according to the Entity Manager instance that handled the request.
	OverrideType_OVERRIDE_TYPE_LIVE OverrideType = 1
	// Override was requested after the entity expired according to the Entity Manager instance that handled the request.
	OverrideType_OVERRIDE_TYPE_POST_EXPIRY OverrideType = 2
)

// Enum value maps for OverrideType.
var (
	OverrideType_name = map[int32]string{
		0: "OVERRIDE_TYPE_INVALID",
		1: "OVERRIDE_TYPE_LIVE",
		2: "OVERRIDE_TYPE_POST_EXPIRY",
	}
	OverrideType_value = map[string]int32{
		"OVERRIDE_TYPE_INVALID":     0,
		"OVERRIDE_TYPE_LIVE":        1,
		"OVERRIDE_TYPE_POST_EXPIRY": 2,
	}
)

func (x OverrideType) Enum() *OverrideType {
	p := new(OverrideType)
	*p = x
	return p
}

func (x OverrideType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OverrideType) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_types_pub_proto_enumTypes[3].Descriptor()
}

func (OverrideType) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_types_pub_proto_enumTypes[3]
}

func (x OverrideType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OverrideType.Descriptor instead.
func (OverrideType) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_types_pub_proto_rawDescGZIP(), []int{3}
}

type UInt32Range struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LowerBound    uint32                 `protobuf:"varint,1,opt,name=lower_bound,json=lowerBound,proto3" json:"lower_bound,omitempty"`
	UpperBound    uint32                 `protobuf:"varint,2,opt,name=upper_bound,json=upperBound,proto3" json:"upper_bound,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UInt32Range) Reset() {
	*x = UInt32Range{}
	mi := &file_anduril_entitymanager_v1_types_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UInt32Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UInt32Range) ProtoMessage() {}

func (x *UInt32Range) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_types_pub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UInt32Range.ProtoReflect.Descriptor instead.
func (*UInt32Range) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_types_pub_proto_rawDescGZIP(), []int{0}
}

func (x *UInt32Range) GetLowerBound() uint32 {
	if x != nil {
		return x.LowerBound
	}
	return 0
}

func (x *UInt32Range) GetUpperBound() uint32 {
	if x != nil {
		return x.UpperBound
	}
	return 0
}

type FloatRange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LowerBound    float32                `protobuf:"fixed32,1,opt,name=lower_bound,json=lowerBound,proto3" json:"lower_bound,omitempty"`
	UpperBound    float32                `protobuf:"fixed32,2,opt,name=upper_bound,json=upperBound,proto3" json:"upper_bound,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FloatRange) Reset() {
	*x = FloatRange{}
	mi := &file_anduril_entitymanager_v1_types_pub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FloatRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FloatRange) ProtoMessage() {}

func (x *FloatRange) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_types_pub_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FloatRange.ProtoReflect.Descriptor instead.
func (*FloatRange) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_types_pub_proto_rawDescGZIP(), []int{1}
}

func (x *FloatRange) GetLowerBound() float32 {
	if x != nil {
		return x.LowerBound
	}
	return 0
}

func (x *FloatRange) GetUpperBound() float32 {
	if x != nil {
		return x.UpperBound
	}
	return 0
}

var File_anduril_entitymanager_v1_types_pub_proto protoreflect.FileDescriptor

const file_anduril_entitymanager_v1_types_pub_proto_rawDesc = "" +
	"\n" +
	"(anduril/entitymanager/v1/types.pub.proto\x12\x18anduril.entitymanager.v1\"O\n" +
	"\vUInt32Range\x12\x1f\n" +
	"\vlower_bound\x18\x01 \x01(\rR\n" +
	"lowerBound\x12\x1f\n" +
	"\vupper_bound\x18\x02 \x01(\rR\n" +
	"upperBound\"N\n" +
	"\n" +
	"FloatRange\x12\x1f\n" +
	"\vlower_bound\x18\x01 \x01(\x02R\n" +
	"lowerBound\x12\x1f\n" +
	"\vupper_bound\x18\x02 \x01(\x02R\n" +
	"upperBound*\x98\x05\n" +
	"\tAltIdType\x12\x17\n" +
	"\x13ALT_ID_TYPE_INVALID\x10\x00\x12\x1a\n" +
	"\x16ALT_ID_TYPE_TRACK_ID_2\x10\x01\x12\x1a\n" +
	"\x16ALT_ID_TYPE_TRACK_ID_1\x10\f\x12\x16\n" +
	"\x12ALT_ID_TYPE_SPI_ID\x10\x02\x12\x1f\n" +
	"\x1bALT_ID_TYPE_NITF_FILE_TITLE\x10\x03\x12#\n" +
	"\x1fALT_ID_TYPE_TRACK_REPO_ALERT_ID\x10\x04\x12\x18\n" +
	"\x14ALT_ID_TYPE_ASSET_ID\x10\x05\x12#\n" +
	"\x1fALT_ID_TYPE_LINK16_TRACK_NUMBER\x10\x06\x12\x19\n" +
	"\x15ALT_ID_TYPE_LINK16_JU\x10\a\x12\x1f\n" +
	"\x1bALT_ID_TYPE_NCCT_MESSAGE_ID\x10\b\x12\x18\n" +
	"\x14ALT_ID_TYPE_CALLSIGN\x10\t\x12\x17\n" +
	"\x13ALT_ID_TYPE_MMSI_ID\x10\n" +
	"\x12\x17\n" +
	"\x13ALT_ID_TYPE_VMF_URN\x10\v\x12\x16\n" +
	"\x12ALT_ID_TYPE_IMO_ID\x10\r\x12!\n" +
	"\x1dALT_ID_TYPE_VMF_TARGET_NUMBER\x10\x0e\x12\x1d\n" +
	"\x19ALT_ID_TYPE_SERIAL_NUMBER\x10\x0f\x12\x1f\n" +
	"\x1bALT_ID_TYPE_REGISTRATION_ID\x10\x10\x12\x17\n" +
	"\x13ALT_ID_TYPE_IBS_GID\x10\x11\x12\x16\n" +
	"\x12ALT_ID_TYPE_DODAAC\x10\x12\x12\x13\n" +
	"\x0fALT_ID_TYPE_UIC\x10\x13\x12\x1c\n" +
	"\x18ALT_ID_TYPE_NORAD_CAT_ID\x10\x14\x12\x1b\n" +
	"\x17ALT_ID_TYPE_UNOOSA_NAME\x10\x17\x12\x19\n" +
	"\x15ALT_ID_TYPE_UNOOSA_ID\x10\x18*\xa2\x01\n" +
	"\bTemplate\x12\x14\n" +
	"\x10TEMPLATE_INVALID\x10\x00\x12\x12\n" +
	"\x0eTEMPLATE_TRACK\x10\x01\x12%\n" +
	"!TEMPLATE_SENSOR_POINT_OF_INTEREST\x10\x02\x12\x12\n" +
	"\x0eTEMPLATE_ASSET\x10\x03\x12\x10\n" +
	"\fTEMPLATE_GEO\x10\x04\x12\x1f\n" +
	"\x1bTEMPLATE_SIGNAL_OF_INTEREST\x10\x05*\xc8\x01\n" +
	"\x0eOverrideStatus\x12\x1b\n" +
	"\x17OVERRIDE_STATUS_INVALID\x10\x00\x12\x1b\n" +
	"\x17OVERRIDE_STATUS_APPLIED\x10\x01\x12\x1b\n" +
	"\x17OVERRIDE_STATUS_PENDING\x10\x02\x12\x1b\n" +
	"\x17OVERRIDE_STATUS_TIMEOUT\x10\x03\x12\x1c\n" +
	"\x18OVERRIDE_STATUS_REJECTED\x10\x04\x12$\n" +
	" OVERRIDE_STATUS_DELETION_PENDING\x10\x05*`\n" +
	"\fOverrideType\x12\x19\n" +
	"\x15OVERRIDE_TYPE_INVALID\x10\x00\x12\x16\n" +
	"\x12OVERRIDE_TYPE_LIVE\x10\x01\x12\x1d\n" +
	"\x19OVERRIDE_TYPE_POST_EXPIRY\x10\x02B\xff\x01\n" +
	"\x1ccom.anduril.entitymanager.v1B\rTypesPubProtoP\x01ZNgithub.com/anduril/lattice-sdk-go/src/anduril/entitymanager/v1;entitymanagerv1\xa2\x02\x03AEX\xaa\x02\x18Anduril.Entitymanager.V1\xca\x02\x18Anduril\\Entitymanager\\V1\xe2\x02$Anduril\\Entitymanager\\V1\\GPBMetadata\xea\x02\x1aAnduril::Entitymanager::V1b\x06proto3"

var (
	file_anduril_entitymanager_v1_types_pub_proto_rawDescOnce sync.Once
	file_anduril_entitymanager_v1_types_pub_proto_rawDescData []byte
)

func file_anduril_entitymanager_v1_types_pub_proto_rawDescGZIP() []byte {
	file_anduril_entitymanager_v1_types_pub_proto_rawDescOnce.Do(func() {
		file_anduril_entitymanager_v1_types_pub_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_anduril_entitymanager_v1_types_pub_proto_rawDesc), len(file_anduril_entitymanager_v1_types_pub_proto_rawDesc)))
	})
	return file_anduril_entitymanager_v1_types_pub_proto_rawDescData
}

var file_anduril_entitymanager_v1_types_pub_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_anduril_entitymanager_v1_types_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_anduril_entitymanager_v1_types_pub_proto_goTypes = []any{
	(AltIdType)(0),      // 0: anduril.entitymanager.v1.AltIdType
	(Template)(0),       // 1: anduril.entitymanager.v1.Template
	(OverrideStatus)(0), // 2: anduril.entitymanager.v1.OverrideStatus
	(OverrideType)(0),   // 3: anduril.entitymanager.v1.OverrideType
	(*UInt32Range)(nil), // 4: anduril.entitymanager.v1.UInt32Range
	(*FloatRange)(nil),  // 5: anduril.entitymanager.v1.FloatRange
}
var file_anduril_entitymanager_v1_types_pub_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_anduril_entitymanager_v1_types_pub_proto_init() }
func file_anduril_entitymanager_v1_types_pub_proto_init() {
	if File_anduril_entitymanager_v1_types_pub_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_anduril_entitymanager_v1_types_pub_proto_rawDesc), len(file_anduril_entitymanager_v1_types_pub_proto_rawDesc)),
			NumEnums:      4,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_entitymanager_v1_types_pub_proto_goTypes,
		DependencyIndexes: file_anduril_entitymanager_v1_types_pub_proto_depIdxs,
		EnumInfos:         file_anduril_entitymanager_v1_types_pub_proto_enumTypes,
		MessageInfos:      file_anduril_entitymanager_v1_types_pub_proto_msgTypes,
	}.Build()
	File_anduril_entitymanager_v1_types_pub_proto = out.File
	file_anduril_entitymanager_v1_types_pub_proto_goTypes = nil
	file_anduril_entitymanager_v1_types_pub_proto_depIdxs = nil
}
