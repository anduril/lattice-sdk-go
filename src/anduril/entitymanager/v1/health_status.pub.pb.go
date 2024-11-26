// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: anduril/entitymanager/v1/health_status.pub.proto

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

// Enumeration of possible connection states.
type ConnectionStatus int32

const (
	ConnectionStatus_CONNECTION_STATUS_INVALID ConnectionStatus = 0
	ConnectionStatus_CONNECTION_STATUS_ONLINE  ConnectionStatus = 1
	ConnectionStatus_CONNECTION_STATUS_OFFLINE ConnectionStatus = 2
)

// Enum value maps for ConnectionStatus.
var (
	ConnectionStatus_name = map[int32]string{
		0: "CONNECTION_STATUS_INVALID",
		1: "CONNECTION_STATUS_ONLINE",
		2: "CONNECTION_STATUS_OFFLINE",
	}
	ConnectionStatus_value = map[string]int32{
		"CONNECTION_STATUS_INVALID": 0,
		"CONNECTION_STATUS_ONLINE":  1,
		"CONNECTION_STATUS_OFFLINE": 2,
	}
)

func (x ConnectionStatus) Enum() *ConnectionStatus {
	p := new(ConnectionStatus)
	*p = x
	return p
}

func (x ConnectionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_health_status_pub_proto_enumTypes[0].Descriptor()
}

func (ConnectionStatus) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_health_status_pub_proto_enumTypes[0]
}

func (x ConnectionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionStatus.Descriptor instead.
func (ConnectionStatus) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_health_status_pub_proto_rawDescGZIP(), []int{0}
}

// Enumeration of possible health states.
type HealthStatus int32

const (
	HealthStatus_HEALTH_STATUS_INVALID HealthStatus = 0
	// Indicates that the component is operating as intended.
	HealthStatus_HEALTH_STATUS_HEALTHY HealthStatus = 1
	// Indicates that the component is at risk of transitioning into a HEALTH_STATUS_FAIL
	// state or that the component is operating in a degraded state.
	HealthStatus_HEALTH_STATUS_WARN HealthStatus = 2
	// Indicates that the component is not functioning as intended.
	HealthStatus_HEALTH_STATUS_FAIL HealthStatus = 3
	// Indicates that the component is offline.
	HealthStatus_HEALTH_STATUS_OFFLINE HealthStatus = 4
	// Indicates that the component is not yet functioning, but it is transitioning into a
	// HEALTH_STATUS_HEALTHY state. A component should only report this state temporarily.
	HealthStatus_HEALTH_STATUS_NOT_READY HealthStatus = 5
)

// Enum value maps for HealthStatus.
var (
	HealthStatus_name = map[int32]string{
		0: "HEALTH_STATUS_INVALID",
		1: "HEALTH_STATUS_HEALTHY",
		2: "HEALTH_STATUS_WARN",
		3: "HEALTH_STATUS_FAIL",
		4: "HEALTH_STATUS_OFFLINE",
		5: "HEALTH_STATUS_NOT_READY",
	}
	HealthStatus_value = map[string]int32{
		"HEALTH_STATUS_INVALID":   0,
		"HEALTH_STATUS_HEALTHY":   1,
		"HEALTH_STATUS_WARN":      2,
		"HEALTH_STATUS_FAIL":      3,
		"HEALTH_STATUS_OFFLINE":   4,
		"HEALTH_STATUS_NOT_READY": 5,
	}
)

func (x HealthStatus) Enum() *HealthStatus {
	p := new(HealthStatus)
	*p = x
	return p
}

func (x HealthStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_health_status_pub_proto_enumTypes[1].Descriptor()
}

func (HealthStatus) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_health_status_pub_proto_enumTypes[1]
}

func (x HealthStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthStatus.Descriptor instead.
func (HealthStatus) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_health_status_pub_proto_rawDescGZIP(), []int{1}
}

// Alerts are categorized into one of three levels - Warnings, Cautions, and Advisories (WCAs).
type AlertLevel int32

const (
	AlertLevel_ALERT_LEVEL_INVALID AlertLevel = 0
	// For conditions that require awareness and may require subsequent response.
	AlertLevel_ALERT_LEVEL_ADVISORY AlertLevel = 1
	// For conditions that require immediate awareness and subsequent response.
	AlertLevel_ALERT_LEVEL_CAUTION AlertLevel = 2
	// For conditions that require immediate awareness and response.
	AlertLevel_ALERT_LEVEL_WARNING AlertLevel = 3
)

// Enum value maps for AlertLevel.
var (
	AlertLevel_name = map[int32]string{
		0: "ALERT_LEVEL_INVALID",
		1: "ALERT_LEVEL_ADVISORY",
		2: "ALERT_LEVEL_CAUTION",
		3: "ALERT_LEVEL_WARNING",
	}
	AlertLevel_value = map[string]int32{
		"ALERT_LEVEL_INVALID":  0,
		"ALERT_LEVEL_ADVISORY": 1,
		"ALERT_LEVEL_CAUTION":  2,
		"ALERT_LEVEL_WARNING":  3,
	}
)

func (x AlertLevel) Enum() *AlertLevel {
	p := new(AlertLevel)
	*p = x
	return p
}

func (x AlertLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_health_status_pub_proto_enumTypes[2].Descriptor()
}

func (AlertLevel) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_health_status_pub_proto_enumTypes[2]
}

func (x AlertLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertLevel.Descriptor instead.
func (AlertLevel) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_health_status_pub_proto_rawDescGZIP(), []int{2}
}

// A message describing the component's health status.
type ComponentMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status associated with this message.
	Status HealthStatus `protobuf:"varint,1,opt,name=status,proto3,enum=anduril.entitymanager.v1.HealthStatus" json:"status,omitempty"`
	// The human-readable content of the message.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ComponentMessage) Reset() {
	*x = ComponentMessage{}
	mi := &file_anduril_entitymanager_v1_health_status_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComponentMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentMessage) ProtoMessage() {}

func (x *ComponentMessage) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_health_status_pub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentMessage.ProtoReflect.Descriptor instead.
func (*ComponentMessage) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_health_status_pub_proto_rawDescGZIP(), []int{0}
}

func (x *ComponentMessage) GetStatus() HealthStatus {
	if x != nil {
		return x.Status
	}
	return HealthStatus_HEALTH_STATUS_INVALID
}

func (x *ComponentMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Health of an individual component.
type ComponentHealth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Consistent internal ID for this component.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Display name for this component.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Health for this component.
	Health HealthStatus `protobuf:"varint,3,opt,name=health,proto3,enum=anduril.entitymanager.v1.HealthStatus" json:"health,omitempty"`
	// Human-readable describing the component state. These messages should be understandable by end users.
	Messages []*ComponentMessage `protobuf:"bytes,4,rep,name=messages,proto3" json:"messages,omitempty"`
	// The last update time for this specific component.
	// If this timestamp is unset, the data is assumed to be most recent
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *ComponentHealth) Reset() {
	*x = ComponentHealth{}
	mi := &file_anduril_entitymanager_v1_health_status_pub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComponentHealth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentHealth) ProtoMessage() {}

func (x *ComponentHealth) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_health_status_pub_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentHealth.ProtoReflect.Descriptor instead.
func (*ComponentHealth) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_health_status_pub_proto_rawDescGZIP(), []int{1}
}

func (x *ComponentHealth) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ComponentHealth) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ComponentHealth) GetHealth() HealthStatus {
	if x != nil {
		return x.Health
	}
	return HealthStatus_HEALTH_STATUS_INVALID
}

func (x *ComponentHealth) GetMessages() []*ComponentMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *ComponentHealth) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// General health of the entity as reported by the entity.
type Health struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status indicating whether the entity is able to communicate with Entity Manager.
	ConnectionStatus ConnectionStatus `protobuf:"varint,1,opt,name=connection_status,json=connectionStatus,proto3,enum=anduril.entitymanager.v1.ConnectionStatus" json:"connection_status,omitempty"`
	// Top-level health status; typically a roll-up of individual component healths.
	HealthStatus HealthStatus `protobuf:"varint,2,opt,name=health_status,json=healthStatus,proto3,enum=anduril.entitymanager.v1.HealthStatus" json:"health_status,omitempty"`
	// Health of individual components running on this Entity.
	Components []*ComponentHealth `protobuf:"bytes,3,rep,name=components,proto3" json:"components,omitempty"`
	// The update time for the top-level health information.
	// If this timestamp is unset, the data is assumed to be most recent
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Active alerts indicate a critical change in system state sent by the asset
	// that must be made known to an operator or consumer of the common operating picture.
	// Alerts are different from ComponentHealth messages--an active alert does not necessarily
	// indicate a component is in an unhealthy state. For example, an asset may trigger
	// an active alert based on fuel levels running low. Alerts should be removed from this list when their conditions
	// are cleared. In other words, only active alerts should be reported here.
	ActiveAlerts []*Alert `protobuf:"bytes,5,rep,name=active_alerts,json=activeAlerts,proto3" json:"active_alerts,omitempty"`
}

func (x *Health) Reset() {
	*x = Health{}
	mi := &file_anduril_entitymanager_v1_health_status_pub_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Health) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Health) ProtoMessage() {}

func (x *Health) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_health_status_pub_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Health.ProtoReflect.Descriptor instead.
func (*Health) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_health_status_pub_proto_rawDescGZIP(), []int{2}
}

func (x *Health) GetConnectionStatus() ConnectionStatus {
	if x != nil {
		return x.ConnectionStatus
	}
	return ConnectionStatus_CONNECTION_STATUS_INVALID
}

func (x *Health) GetHealthStatus() HealthStatus {
	if x != nil {
		return x.HealthStatus
	}
	return HealthStatus_HEALTH_STATUS_INVALID
}

func (x *Health) GetComponents() []*ComponentHealth {
	if x != nil {
		return x.Components
	}
	return nil
}

func (x *Health) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Health) GetActiveAlerts() []*Alert {
	if x != nil {
		return x.ActiveAlerts
	}
	return nil
}

// An alert informs operators of critical events related to system performance and mission
// execution. An alert is produced as a result of one or more alert conditions.
type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Short, machine-readable code that describes this alert. This code is intended to provide systems off-asset
	// with a lookup key to retrieve more detailed information about the alert.
	AlertCode string `protobuf:"bytes,1,opt,name=alert_code,json=alertCode,proto3" json:"alert_code,omitempty"`
	// Human-readable description of this alert. The description is intended for display in the UI for human
	// understanding and should not be used for machine processing. If the description is fixed and the vehicle controller
	// provides no dynamic substitutions, then prefer lookup based on alert_code.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Alert level (Warning, Caution, or Advisory).
	Level AlertLevel `protobuf:"varint,3,opt,name=level,proto3,enum=anduril.entitymanager.v1.AlertLevel" json:"level,omitempty"`
	// Time at which this alert was activated.
	ActivatedTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=activated_time,json=activatedTime,proto3" json:"activated_time,omitempty"`
	// Set of conditions which have activated this alert.
	ActiveConditions []*AlertCondition `protobuf:"bytes,5,rep,name=active_conditions,json=activeConditions,proto3" json:"active_conditions,omitempty"`
}

func (x *Alert) Reset() {
	*x = Alert{}
	mi := &file_anduril_entitymanager_v1_health_status_pub_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_health_status_pub_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_health_status_pub_proto_rawDescGZIP(), []int{3}
}

func (x *Alert) GetAlertCode() string {
	if x != nil {
		return x.AlertCode
	}
	return ""
}

func (x *Alert) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Alert) GetLevel() AlertLevel {
	if x != nil {
		return x.Level
	}
	return AlertLevel_ALERT_LEVEL_INVALID
}

func (x *Alert) GetActivatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ActivatedTime
	}
	return nil
}

func (x *Alert) GetActiveConditions() []*AlertCondition {
	if x != nil {
		return x.ActiveConditions
	}
	return nil
}

// A condition which may trigger an alert.
type AlertCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Short, machine-readable code that describes this condition. This code is intended to provide systems off-asset
	// with a lookup key to retrieve more detailed information about the condition.
	ConditionCode string `protobuf:"bytes,1,opt,name=condition_code,json=conditionCode,proto3" json:"condition_code,omitempty"`
	// Human-readable description of this condition. The description is intended for display in the UI for human
	// understanding and should not be used for machine processing. If the description is fixed and the vehicle controller
	// provides no dynamic substitutions, then prefer lookup based on condition_code.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AlertCondition) Reset() {
	*x = AlertCondition{}
	mi := &file_anduril_entitymanager_v1_health_status_pub_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertCondition) ProtoMessage() {}

func (x *AlertCondition) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_health_status_pub_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertCondition.ProtoReflect.Descriptor instead.
func (*AlertCondition) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_health_status_pub_proto_rawDescGZIP(), []int{4}
}

func (x *AlertCondition) GetConditionCode() string {
	if x != nil {
		return x.ConditionCode
	}
	return ""
}

func (x *AlertCondition) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_anduril_entitymanager_v1_health_status_pub_proto protoreflect.FileDescriptor

var file_anduril_entitymanager_v1_health_status_pub_proto_rawDesc = []byte{
	0x0a, 0x30, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x18, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a,
	0x10, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x3e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x26, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xfa, 0x01, 0x0a, 0x0f,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xfc, 0x02, 0x0a, 0x06, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x12, 0x57, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a,
	0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4b, 0x0a, 0x0d,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x49, 0x0a, 0x0a, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x44, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72,
	0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x22, 0x9e, 0x02, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x24, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x41,
	0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x55, 0x0a, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x59, 0x0a, 0x0e, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2a, 0x6e, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x4e, 0x4c, 0x49,
	0x4e, 0x45, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e,
	0x45, 0x10, 0x02, 0x2a, 0xac, 0x01, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x15, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12,
	0x19, 0x0a, 0x15, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x48, 0x45,
	0x41, 0x4c, 0x54, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x57, 0x41, 0x52, 0x4e,
	0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x48, 0x45,
	0x41, 0x4c, 0x54, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x46, 0x46, 0x4c,
	0x49, 0x4e, 0x45, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59,
	0x10, 0x05, 0x2a, 0x71, 0x0a, 0x0a, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x17, 0x0a, 0x13, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x4c, 0x45,
	0x52, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x41, 0x44, 0x56, 0x49, 0x53, 0x4f, 0x52,
	0x59, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x4c, 0x45, 0x56,
	0x45, 0x4c, 0x5f, 0x43, 0x41, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13,
	0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x57, 0x41, 0x52, 0x4e,
	0x49, 0x4e, 0x47, 0x10, 0x03, 0x42, 0x86, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x50, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e,
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
}

var (
	file_anduril_entitymanager_v1_health_status_pub_proto_rawDescOnce sync.Once
	file_anduril_entitymanager_v1_health_status_pub_proto_rawDescData = file_anduril_entitymanager_v1_health_status_pub_proto_rawDesc
)

func file_anduril_entitymanager_v1_health_status_pub_proto_rawDescGZIP() []byte {
	file_anduril_entitymanager_v1_health_status_pub_proto_rawDescOnce.Do(func() {
		file_anduril_entitymanager_v1_health_status_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_entitymanager_v1_health_status_pub_proto_rawDescData)
	})
	return file_anduril_entitymanager_v1_health_status_pub_proto_rawDescData
}

var file_anduril_entitymanager_v1_health_status_pub_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_anduril_entitymanager_v1_health_status_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_anduril_entitymanager_v1_health_status_pub_proto_goTypes = []any{
	(ConnectionStatus)(0),         // 0: anduril.entitymanager.v1.ConnectionStatus
	(HealthStatus)(0),             // 1: anduril.entitymanager.v1.HealthStatus
	(AlertLevel)(0),               // 2: anduril.entitymanager.v1.AlertLevel
	(*ComponentMessage)(nil),      // 3: anduril.entitymanager.v1.ComponentMessage
	(*ComponentHealth)(nil),       // 4: anduril.entitymanager.v1.ComponentHealth
	(*Health)(nil),                // 5: anduril.entitymanager.v1.Health
	(*Alert)(nil),                 // 6: anduril.entitymanager.v1.Alert
	(*AlertCondition)(nil),        // 7: anduril.entitymanager.v1.AlertCondition
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_anduril_entitymanager_v1_health_status_pub_proto_depIdxs = []int32{
	1,  // 0: anduril.entitymanager.v1.ComponentMessage.status:type_name -> anduril.entitymanager.v1.HealthStatus
	1,  // 1: anduril.entitymanager.v1.ComponentHealth.health:type_name -> anduril.entitymanager.v1.HealthStatus
	3,  // 2: anduril.entitymanager.v1.ComponentHealth.messages:type_name -> anduril.entitymanager.v1.ComponentMessage
	8,  // 3: anduril.entitymanager.v1.ComponentHealth.update_time:type_name -> google.protobuf.Timestamp
	0,  // 4: anduril.entitymanager.v1.Health.connection_status:type_name -> anduril.entitymanager.v1.ConnectionStatus
	1,  // 5: anduril.entitymanager.v1.Health.health_status:type_name -> anduril.entitymanager.v1.HealthStatus
	4,  // 6: anduril.entitymanager.v1.Health.components:type_name -> anduril.entitymanager.v1.ComponentHealth
	8,  // 7: anduril.entitymanager.v1.Health.update_time:type_name -> google.protobuf.Timestamp
	6,  // 8: anduril.entitymanager.v1.Health.active_alerts:type_name -> anduril.entitymanager.v1.Alert
	2,  // 9: anduril.entitymanager.v1.Alert.level:type_name -> anduril.entitymanager.v1.AlertLevel
	8,  // 10: anduril.entitymanager.v1.Alert.activated_time:type_name -> google.protobuf.Timestamp
	7,  // 11: anduril.entitymanager.v1.Alert.active_conditions:type_name -> anduril.entitymanager.v1.AlertCondition
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_anduril_entitymanager_v1_health_status_pub_proto_init() }
func file_anduril_entitymanager_v1_health_status_pub_proto_init() {
	if File_anduril_entitymanager_v1_health_status_pub_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_anduril_entitymanager_v1_health_status_pub_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_entitymanager_v1_health_status_pub_proto_goTypes,
		DependencyIndexes: file_anduril_entitymanager_v1_health_status_pub_proto_depIdxs,
		EnumInfos:         file_anduril_entitymanager_v1_health_status_pub_proto_enumTypes,
		MessageInfos:      file_anduril_entitymanager_v1_health_status_pub_proto_msgTypes,
	}.Build()
	File_anduril_entitymanager_v1_health_status_pub_proto = out.File
	file_anduril_entitymanager_v1_health_status_pub_proto_rawDesc = nil
	file_anduril_entitymanager_v1_health_status_pub_proto_goTypes = nil
	file_anduril_entitymanager_v1_health_status_pub_proto_depIdxs = nil
}