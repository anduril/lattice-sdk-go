// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: anduril/entitymanager/v1/payloads.pub.proto

package entitymanagerv1

import (
	v1 "github.com/anduril/lattice-sdk-go/src/anduril/ontology/v1"
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

// Describes the current operational state of a payload configuration.
type PayloadOperationalState int32

const (
	PayloadOperationalState_PAYLOAD_OPERATIONAL_STATE_INVALID         PayloadOperationalState = 0
	PayloadOperationalState_PAYLOAD_OPERATIONAL_STATE_OFF             PayloadOperationalState = 1
	PayloadOperationalState_PAYLOAD_OPERATIONAL_STATE_NON_OPERATIONAL PayloadOperationalState = 2
	PayloadOperationalState_PAYLOAD_OPERATIONAL_STATE_DEGRADED        PayloadOperationalState = 3
	PayloadOperationalState_PAYLOAD_OPERATIONAL_STATE_OPERATIONAL     PayloadOperationalState = 4
	PayloadOperationalState_PAYLOAD_OPERATIONAL_STATE_OUT_OF_SERVICE  PayloadOperationalState = 5
	PayloadOperationalState_PAYLOAD_OPERATIONAL_STATE_UNKNOWN         PayloadOperationalState = 6
)

// Enum value maps for PayloadOperationalState.
var (
	PayloadOperationalState_name = map[int32]string{
		0: "PAYLOAD_OPERATIONAL_STATE_INVALID",
		1: "PAYLOAD_OPERATIONAL_STATE_OFF",
		2: "PAYLOAD_OPERATIONAL_STATE_NON_OPERATIONAL",
		3: "PAYLOAD_OPERATIONAL_STATE_DEGRADED",
		4: "PAYLOAD_OPERATIONAL_STATE_OPERATIONAL",
		5: "PAYLOAD_OPERATIONAL_STATE_OUT_OF_SERVICE",
		6: "PAYLOAD_OPERATIONAL_STATE_UNKNOWN",
	}
	PayloadOperationalState_value = map[string]int32{
		"PAYLOAD_OPERATIONAL_STATE_INVALID":         0,
		"PAYLOAD_OPERATIONAL_STATE_OFF":             1,
		"PAYLOAD_OPERATIONAL_STATE_NON_OPERATIONAL": 2,
		"PAYLOAD_OPERATIONAL_STATE_DEGRADED":        3,
		"PAYLOAD_OPERATIONAL_STATE_OPERATIONAL":     4,
		"PAYLOAD_OPERATIONAL_STATE_OUT_OF_SERVICE":  5,
		"PAYLOAD_OPERATIONAL_STATE_UNKNOWN":         6,
	}
)

func (x PayloadOperationalState) Enum() *PayloadOperationalState {
	p := new(PayloadOperationalState)
	*p = x
	return p
}

func (x PayloadOperationalState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PayloadOperationalState) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_payloads_pub_proto_enumTypes[0].Descriptor()
}

func (PayloadOperationalState) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_payloads_pub_proto_enumTypes[0]
}

func (x PayloadOperationalState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PayloadOperationalState.Descriptor instead.
func (PayloadOperationalState) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_payloads_pub_proto_rawDescGZIP(), []int{0}
}

// List of payloads available for an entity.
type Payloads struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	PayloadConfigurations []*Payload             `protobuf:"bytes,1,rep,name=payload_configurations,json=payloadConfigurations,proto3" json:"payload_configurations,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *Payloads) Reset() {
	*x = Payloads{}
	mi := &file_anduril_entitymanager_v1_payloads_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Payloads) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payloads) ProtoMessage() {}

func (x *Payloads) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_payloads_pub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payloads.ProtoReflect.Descriptor instead.
func (*Payloads) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_payloads_pub_proto_rawDescGZIP(), []int{0}
}

func (x *Payloads) GetPayloadConfigurations() []*Payload {
	if x != nil {
		return x.PayloadConfigurations
	}
	return nil
}

// Individual payload configuration.
type Payload struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Config        *PayloadConfiguration  `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Payload) Reset() {
	*x = Payload{}
	mi := &file_anduril_entitymanager_v1_payloads_pub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_payloads_pub_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_payloads_pub_proto_rawDescGZIP(), []int{1}
}

func (x *Payload) GetConfig() *PayloadConfiguration {
	if x != nil {
		return x.Config
	}
	return nil
}

type PayloadConfiguration struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Identifying ID for the capability.
	// This ID may be used multiple times to represent payloads that are the same capability but have different operational states
	CapabilityId string `protobuf:"bytes,1,opt,name=capability_id,json=capabilityId,proto3" json:"capability_id,omitempty"`
	// The number of payloads currently available in the configuration.
	Quantity uint32 `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// The target environments the configuration is effective against.
	EffectiveEnvironment []v1.Environment `protobuf:"varint,5,rep,packed,name=effective_environment,json=effectiveEnvironment,proto3,enum=anduril.ontology.v1.Environment" json:"effective_environment,omitempty"`
	// The operational state of this payload.
	PayloadOperationalState PayloadOperationalState `protobuf:"varint,6,opt,name=payload_operational_state,json=payloadOperationalState,proto3,enum=anduril.entitymanager.v1.PayloadOperationalState" json:"payload_operational_state,omitempty"`
	// A human readable description of the payload
	PayloadDescription string `protobuf:"bytes,7,opt,name=payload_description,json=payloadDescription,proto3" json:"payload_description,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *PayloadConfiguration) Reset() {
	*x = PayloadConfiguration{}
	mi := &file_anduril_entitymanager_v1_payloads_pub_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PayloadConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadConfiguration) ProtoMessage() {}

func (x *PayloadConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_payloads_pub_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadConfiguration.ProtoReflect.Descriptor instead.
func (*PayloadConfiguration) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_payloads_pub_proto_rawDescGZIP(), []int{2}
}

func (x *PayloadConfiguration) GetCapabilityId() string {
	if x != nil {
		return x.CapabilityId
	}
	return ""
}

func (x *PayloadConfiguration) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *PayloadConfiguration) GetEffectiveEnvironment() []v1.Environment {
	if x != nil {
		return x.EffectiveEnvironment
	}
	return nil
}

func (x *PayloadConfiguration) GetPayloadOperationalState() PayloadOperationalState {
	if x != nil {
		return x.PayloadOperationalState
	}
	return PayloadOperationalState_PAYLOAD_OPERATIONAL_STATE_INVALID
}

func (x *PayloadConfiguration) GetPayloadDescription() string {
	if x != nil {
		return x.PayloadDescription
	}
	return ""
}

var File_anduril_entitymanager_v1_payloads_pub_proto protoreflect.FileDescriptor

var file_anduril_entitymanager_v1_payloads_pub_proto_rawDesc = string([]byte{
	0x0a, 0x2b, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x6f, 0x6e, 0x74,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x75,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x08, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x12, 0x5d, 0x0a, 0x16, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x03, 0xc8, 0x3e, 0x01, 0x52, 0x15, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x51, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x46, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xce, 0x02, 0x0a, 0x14, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x55, 0x0a, 0x15, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x20,
	0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x6f, 0x6e, 0x74, 0x6f, 0x6c, 0x6f, 0x67,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x14, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x6d, 0x0a, 0x19, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x61, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x17, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0xba, 0x02, 0x0a, 0x17, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x25, 0x0a, 0x21, 0x50, 0x41, 0x59, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x4f, 0x50,
	0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x50, 0x41, 0x59,
	0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x01, 0x12, 0x2d, 0x0a, 0x29,
	0x50, 0x41, 0x59, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f, 0x4f, 0x50,
	0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x26, 0x0a, 0x22, 0x50,
	0x41, 0x59, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x45, 0x47, 0x52, 0x41, 0x44, 0x45,
	0x44, 0x10, 0x03, 0x12, 0x29, 0x0a, 0x25, 0x50, 0x41, 0x59, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x04, 0x12, 0x2c,
	0x0a, 0x28, 0x50, 0x41, 0x59, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x5f,
	0x4f, 0x46, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x05, 0x12, 0x25, 0x0a, 0x21,
	0x50, 0x41, 0x59, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x06, 0x42, 0x82, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x50, 0x75,
	0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x6c, 0x61, 0x74,
	0x74, 0x69, 0x63, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x45, 0x58, 0xaa, 0x02,
	0x18, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x41, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x41, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_anduril_entitymanager_v1_payloads_pub_proto_rawDescOnce sync.Once
	file_anduril_entitymanager_v1_payloads_pub_proto_rawDescData []byte
)

func file_anduril_entitymanager_v1_payloads_pub_proto_rawDescGZIP() []byte {
	file_anduril_entitymanager_v1_payloads_pub_proto_rawDescOnce.Do(func() {
		file_anduril_entitymanager_v1_payloads_pub_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_anduril_entitymanager_v1_payloads_pub_proto_rawDesc), len(file_anduril_entitymanager_v1_payloads_pub_proto_rawDesc)))
	})
	return file_anduril_entitymanager_v1_payloads_pub_proto_rawDescData
}

var file_anduril_entitymanager_v1_payloads_pub_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_anduril_entitymanager_v1_payloads_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_anduril_entitymanager_v1_payloads_pub_proto_goTypes = []any{
	(PayloadOperationalState)(0), // 0: anduril.entitymanager.v1.PayloadOperationalState
	(*Payloads)(nil),             // 1: anduril.entitymanager.v1.Payloads
	(*Payload)(nil),              // 2: anduril.entitymanager.v1.Payload
	(*PayloadConfiguration)(nil), // 3: anduril.entitymanager.v1.PayloadConfiguration
	(v1.Environment)(0),          // 4: anduril.ontology.v1.Environment
}
var file_anduril_entitymanager_v1_payloads_pub_proto_depIdxs = []int32{
	2, // 0: anduril.entitymanager.v1.Payloads.payload_configurations:type_name -> anduril.entitymanager.v1.Payload
	3, // 1: anduril.entitymanager.v1.Payload.config:type_name -> anduril.entitymanager.v1.PayloadConfiguration
	4, // 2: anduril.entitymanager.v1.PayloadConfiguration.effective_environment:type_name -> anduril.ontology.v1.Environment
	0, // 3: anduril.entitymanager.v1.PayloadConfiguration.payload_operational_state:type_name -> anduril.entitymanager.v1.PayloadOperationalState
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_anduril_entitymanager_v1_payloads_pub_proto_init() }
func file_anduril_entitymanager_v1_payloads_pub_proto_init() {
	if File_anduril_entitymanager_v1_payloads_pub_proto != nil {
		return
	}
	file_anduril_entitymanager_v1_options_pub_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_anduril_entitymanager_v1_payloads_pub_proto_rawDesc), len(file_anduril_entitymanager_v1_payloads_pub_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_entitymanager_v1_payloads_pub_proto_goTypes,
		DependencyIndexes: file_anduril_entitymanager_v1_payloads_pub_proto_depIdxs,
		EnumInfos:         file_anduril_entitymanager_v1_payloads_pub_proto_enumTypes,
		MessageInfos:      file_anduril_entitymanager_v1_payloads_pub_proto_msgTypes,
	}.Build()
	File_anduril_entitymanager_v1_payloads_pub_proto = out.File
	file_anduril_entitymanager_v1_payloads_pub_proto_goTypes = nil
	file_anduril_entitymanager_v1_payloads_pub_proto_depIdxs = nil
}
