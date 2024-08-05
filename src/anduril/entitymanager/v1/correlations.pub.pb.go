// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: anduril/entitymanager/v1/correlations.pub.proto

package entitymanagerv1

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

// The status of the correlation.
type CorrelationStatus int32

const (
	CorrelationStatus_CORRELATION_STATUS_INVALID CorrelationStatus = 0
	// potential correlation requested by manual inspection, not yet confirmed.
	CorrelationStatus_CORRELATION_STATUS_MANUAL_INSPECTION CorrelationStatus = 1
	// potential correlation suggested by system, not yet confirmed.
	CorrelationStatus_CORRELATION_STATUS_AUTO_SUGGESTED CorrelationStatus = 2
	// deprecated
	//
	// Deprecated: Marked as deprecated in anduril/entitymanager/v1/correlations.pub.proto.
	CorrelationStatus_CORRELATION_STATUS_START_CORRELATE CorrelationStatus = 3
	// correlation has been confirmed, treat non primary as hidden.
	CorrelationStatus_CORRELATION_STATUS_CONFIRMED CorrelationStatus = 4
	// correlation was explicitly rejected, treat as non correlated.
	CorrelationStatus_CORRELATION_STATUS_DENIED CorrelationStatus = 5
)

// Enum value maps for CorrelationStatus.
var (
	CorrelationStatus_name = map[int32]string{
		0: "CORRELATION_STATUS_INVALID",
		1: "CORRELATION_STATUS_MANUAL_INSPECTION",
		2: "CORRELATION_STATUS_AUTO_SUGGESTED",
		3: "CORRELATION_STATUS_START_CORRELATE",
		4: "CORRELATION_STATUS_CONFIRMED",
		5: "CORRELATION_STATUS_DENIED",
	}
	CorrelationStatus_value = map[string]int32{
		"CORRELATION_STATUS_INVALID":           0,
		"CORRELATION_STATUS_MANUAL_INSPECTION": 1,
		"CORRELATION_STATUS_AUTO_SUGGESTED":    2,
		"CORRELATION_STATUS_START_CORRELATE":   3,
		"CORRELATION_STATUS_CONFIRMED":         4,
		"CORRELATION_STATUS_DENIED":            5,
	}
)

func (x CorrelationStatus) Enum() *CorrelationStatus {
	p := new(CorrelationStatus)
	*p = x
	return p
}

func (x CorrelationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CorrelationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_correlations_pub_proto_enumTypes[0].Descriptor()
}

func (CorrelationStatus) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_correlations_pub_proto_enumTypes[0]
}

func (x CorrelationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CorrelationStatus.Descriptor instead.
func (CorrelationStatus) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_correlations_pub_proto_rawDescGZIP(), []int{0}
}

// The interpretation of the correlation score.
type ScoreInterpretation int32

const (
	ScoreInterpretation_SCORE_INTERPRETATION_INVALID ScoreInterpretation = 0
	// unlikely these are the same entity
	ScoreInterpretation_SCORE_INTERPRETATION_UNLIKELY ScoreInterpretation = 1
	// likely these are the same entity
	ScoreInterpretation_SCORE_INTERPRETATION_LIKELY ScoreInterpretation = 2
	// very likely these are the same entity
	ScoreInterpretation_SCORE_INTERPRETATION_VERY_LIKELY ScoreInterpretation = 3
)

// Enum value maps for ScoreInterpretation.
var (
	ScoreInterpretation_name = map[int32]string{
		0: "SCORE_INTERPRETATION_INVALID",
		1: "SCORE_INTERPRETATION_UNLIKELY",
		2: "SCORE_INTERPRETATION_LIKELY",
		3: "SCORE_INTERPRETATION_VERY_LIKELY",
	}
	ScoreInterpretation_value = map[string]int32{
		"SCORE_INTERPRETATION_INVALID":     0,
		"SCORE_INTERPRETATION_UNLIKELY":    1,
		"SCORE_INTERPRETATION_LIKELY":      2,
		"SCORE_INTERPRETATION_VERY_LIKELY": 3,
	}
)

func (x ScoreInterpretation) Enum() *ScoreInterpretation {
	p := new(ScoreInterpretation)
	*p = x
	return p
}

func (x ScoreInterpretation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScoreInterpretation) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_correlations_pub_proto_enumTypes[1].Descriptor()
}

func (ScoreInterpretation) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_correlations_pub_proto_enumTypes[1]
}

func (x ScoreInterpretation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScoreInterpretation.Descriptor instead.
func (ScoreInterpretation) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_correlations_pub_proto_rawDescGZIP(), []int{1}
}

// Available for Entities that are a correlated (N to 1) set of entities. This will be present on each entity in the
// set.
type Correlated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// primary entity id
	PrimaryEntityId string `protobuf:"bytes,1,opt,name=primary_entity_id,json=primaryEntityId,proto3" json:"primary_entity_id,omitempty"`
	// status representing this correlation
	Status CorrelationStatus `protobuf:"varint,2,opt,name=status,proto3,enum=anduril.entitymanager.v1.CorrelationStatus" json:"status,omitempty"`
	// score pairings between this and other entity ids
	Scores []*CorrelationScore `protobuf:"bytes,3,rep,name=scores,proto3" json:"scores,omitempty"`
	// if not present, does not expire
	ExpiresTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expires_time,json=expiresTime,proto3" json:"expires_time,omitempty"`
}

func (x *Correlated) Reset() {
	*x = Correlated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entitymanager_v1_correlations_pub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Correlated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Correlated) ProtoMessage() {}

func (x *Correlated) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_correlations_pub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Correlated.ProtoReflect.Descriptor instead.
func (*Correlated) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_correlations_pub_proto_rawDescGZIP(), []int{0}
}

func (x *Correlated) GetPrimaryEntityId() string {
	if x != nil {
		return x.PrimaryEntityId
	}
	return ""
}

func (x *Correlated) GetStatus() CorrelationStatus {
	if x != nil {
		return x.Status
	}
	return CorrelationStatus_CORRELATION_STATUS_INVALID
}

func (x *Correlated) GetScores() []*CorrelationScore {
	if x != nil {
		return x.Scores
	}
	return nil
}

func (x *Correlated) GetExpiresTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresTime
	}
	return nil
}

// A correlation scoring between two entities.
type CorrelationScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtherEntityId  string              `protobuf:"bytes,1,opt,name=other_entity_id,json=otherEntityId,proto3" json:"other_entity_id,omitempty"`
	Score          float32             `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	Interpretation ScoreInterpretation `protobuf:"varint,3,opt,name=interpretation,proto3,enum=anduril.entitymanager.v1.ScoreInterpretation" json:"interpretation,omitempty"`
	// Deprecated: do not use
	//
	// Deprecated: Marked as deprecated in anduril/entitymanager/v1/correlations.pub.proto.
	Link16Compliant bool `protobuf:"varint,4,opt,name=link16_compliant,json=link16Compliant,proto3" json:"link16_compliant,omitempty"`
	// status of other_entity_id correlation, expresses relationship of other to correlation set this entity is part of.
	OtherStatus CorrelationStatus `protobuf:"varint,5,opt,name=other_status,json=otherStatus,proto3,enum=anduril.entitymanager.v1.CorrelationStatus" json:"other_status,omitempty"`
}

func (x *CorrelationScore) Reset() {
	*x = CorrelationScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entitymanager_v1_correlations_pub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CorrelationScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CorrelationScore) ProtoMessage() {}

func (x *CorrelationScore) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_correlations_pub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CorrelationScore.ProtoReflect.Descriptor instead.
func (*CorrelationScore) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_correlations_pub_proto_rawDescGZIP(), []int{1}
}

func (x *CorrelationScore) GetOtherEntityId() string {
	if x != nil {
		return x.OtherEntityId
	}
	return ""
}

func (x *CorrelationScore) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *CorrelationScore) GetInterpretation() ScoreInterpretation {
	if x != nil {
		return x.Interpretation
	}
	return ScoreInterpretation_SCORE_INTERPRETATION_INVALID
}

// Deprecated: Marked as deprecated in anduril/entitymanager/v1/correlations.pub.proto.
func (x *CorrelationScore) GetLink16Compliant() bool {
	if x != nil {
		return x.Link16Compliant
	}
	return false
}

func (x *CorrelationScore) GetOtherStatus() CorrelationStatus {
	if x != nil {
		return x.OtherStatus
	}
	return CorrelationStatus_CORRELATION_STATUS_INVALID
}

var File_anduril_entitymanager_v1_correlations_pub_proto protoreflect.FileDescriptor

var file_anduril_entitymanager_v1_correlations_pub_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x18, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x02, 0x0a,
	0x0a, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x42, 0x0a, 0x06,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x12, 0x3d, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0xa6, 0x02, 0x0a, 0x10, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x61, 0x6e, 0x64,
	0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x70, 0x72, 0x65, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x70, 0x72, 0x65, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x10, 0x6c, 0x69, 0x6e,
	0x6b, 0x31, 0x36, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0f, 0x6c, 0x69, 0x6e, 0x6b, 0x31, 0x36, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x0c, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b,
	0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0xf1, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x72,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e,
	0x0a, 0x1a, 0x43, 0x4f, 0x52, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x28,
	0x0a, 0x24, 0x43, 0x4f, 0x52, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x5f, 0x49, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x43, 0x4f, 0x52, 0x52,
	0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41,
	0x55, 0x54, 0x4f, 0x5f, 0x53, 0x55, 0x47, 0x47, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x2a, 0x0a, 0x22, 0x43, 0x4f, 0x52, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x43, 0x4f, 0x52, 0x52,
	0x45, 0x4c, 0x41, 0x54, 0x45, 0x10, 0x03, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x43,
	0x4f, 0x52, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1d, 0x0a,
	0x19, 0x43, 0x4f, 0x52, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x05, 0x2a, 0xa1, 0x01, 0x0a,
	0x13, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x5f, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x50, 0x52, 0x45, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x5f,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x50, 0x52, 0x45, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x4e, 0x4c, 0x49, 0x4b, 0x45, 0x4c, 0x59, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x43, 0x4f,
	0x52, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x50, 0x52, 0x45, 0x54, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x4c, 0x59, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x53, 0x43,
	0x4f, 0x52, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x50, 0x52, 0x45, 0x54, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x4c, 0x59, 0x10, 0x03,
	0x42, 0x8c, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x14, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50,
	0x75, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x75, 0x6e, 0x2d, 0x61, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2d, 0x73, 0x64, 0x6b, 0x2d,
	0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x41, 0x45, 0x58, 0xaa, 0x02, 0x18, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x18, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x41, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x1a, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x3a, 0x3a, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_anduril_entitymanager_v1_correlations_pub_proto_rawDescOnce sync.Once
	file_anduril_entitymanager_v1_correlations_pub_proto_rawDescData = file_anduril_entitymanager_v1_correlations_pub_proto_rawDesc
)

func file_anduril_entitymanager_v1_correlations_pub_proto_rawDescGZIP() []byte {
	file_anduril_entitymanager_v1_correlations_pub_proto_rawDescOnce.Do(func() {
		file_anduril_entitymanager_v1_correlations_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_entitymanager_v1_correlations_pub_proto_rawDescData)
	})
	return file_anduril_entitymanager_v1_correlations_pub_proto_rawDescData
}

var file_anduril_entitymanager_v1_correlations_pub_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_anduril_entitymanager_v1_correlations_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_anduril_entitymanager_v1_correlations_pub_proto_goTypes = []any{
	(CorrelationStatus)(0),        // 0: anduril.entitymanager.v1.CorrelationStatus
	(ScoreInterpretation)(0),      // 1: anduril.entitymanager.v1.ScoreInterpretation
	(*Correlated)(nil),            // 2: anduril.entitymanager.v1.Correlated
	(*CorrelationScore)(nil),      // 3: anduril.entitymanager.v1.CorrelationScore
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_anduril_entitymanager_v1_correlations_pub_proto_depIdxs = []int32{
	0, // 0: anduril.entitymanager.v1.Correlated.status:type_name -> anduril.entitymanager.v1.CorrelationStatus
	3, // 1: anduril.entitymanager.v1.Correlated.scores:type_name -> anduril.entitymanager.v1.CorrelationScore
	4, // 2: anduril.entitymanager.v1.Correlated.expires_time:type_name -> google.protobuf.Timestamp
	1, // 3: anduril.entitymanager.v1.CorrelationScore.interpretation:type_name -> anduril.entitymanager.v1.ScoreInterpretation
	0, // 4: anduril.entitymanager.v1.CorrelationScore.other_status:type_name -> anduril.entitymanager.v1.CorrelationStatus
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_anduril_entitymanager_v1_correlations_pub_proto_init() }
func file_anduril_entitymanager_v1_correlations_pub_proto_init() {
	if File_anduril_entitymanager_v1_correlations_pub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_anduril_entitymanager_v1_correlations_pub_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Correlated); i {
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
		file_anduril_entitymanager_v1_correlations_pub_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CorrelationScore); i {
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
			RawDescriptor: file_anduril_entitymanager_v1_correlations_pub_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_entitymanager_v1_correlations_pub_proto_goTypes,
		DependencyIndexes: file_anduril_entitymanager_v1_correlations_pub_proto_depIdxs,
		EnumInfos:         file_anduril_entitymanager_v1_correlations_pub_proto_enumTypes,
		MessageInfos:      file_anduril_entitymanager_v1_correlations_pub_proto_msgTypes,
	}.Build()
	File_anduril_entitymanager_v1_correlations_pub_proto = out.File
	file_anduril_entitymanager_v1_correlations_pub_proto_rawDesc = nil
	file_anduril_entitymanager_v1_correlations_pub_proto_goTypes = nil
	file_anduril_entitymanager_v1_correlations_pub_proto_depIdxs = nil
}