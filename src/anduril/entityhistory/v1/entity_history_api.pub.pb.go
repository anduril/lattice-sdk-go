// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: anduril/entityhistory/v1/entity_history_api.pub.proto

package entityhistoryv1

import (
	v1 "github.com/dogun-anduril/anduril-sdk-go/src/anduril/entitymanager/v1"
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

type ListHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the history returned to the client.
	HistoryQuery *HistoryQuery `protobuf:"bytes,1,opt,name=history_query,json=historyQuery,proto3" json:"history_query,omitempty"`
	// Page token in the case that the query results are paginated.
	// This will be deprecated in favor of history_page_token.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// history_page_token should be provided as part of a request to consume paginated history data and was received via
	// the most recent ListHistoryResponse. Clients should not create this on their own and the initial ListHistoryRequest
	// should leave this field empty.
	HistoryPageToken *HistoryPageToken `protobuf:"bytes,3,opt,name=history_page_token,json=historyPageToken,proto3" json:"history_page_token,omitempty"`
}

func (x *ListHistoryRequest) Reset() {
	*x = ListHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHistoryRequest) ProtoMessage() {}

func (x *ListHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHistoryRequest.ProtoReflect.Descriptor instead.
func (*ListHistoryRequest) Descriptor() ([]byte, []int) {
	return file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDescGZIP(), []int{0}
}

func (x *ListHistoryRequest) GetHistoryQuery() *HistoryQuery {
	if x != nil {
		return x.HistoryQuery
	}
	return nil
}

func (x *ListHistoryRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListHistoryRequest) GetHistoryPageToken() *HistoryPageToken {
	if x != nil {
		return x.HistoryPageToken
	}
	return nil
}

type ListHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HistoryPage *HistoryPage `protobuf:"bytes,1,opt,name=history_page,json=historyPage,proto3" json:"history_page,omitempty"`
	// If present this page token can be used to retrieve the next page of
	// history. If empty, there are no more results.
	// This will be deprecated in favor of next_history_page_token.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Next_history_page_token received should be provided as part of any followup requests to continue consuming the data
	// for the query requested until the HistoryPageToken reports is_complete as true.
	NextHistoryPageToken *HistoryPageToken `protobuf:"bytes,3,opt,name=next_history_page_token,json=nextHistoryPageToken,proto3" json:"next_history_page_token,omitempty"`
}

func (x *ListHistoryResponse) Reset() {
	*x = ListHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHistoryResponse) ProtoMessage() {}

func (x *ListHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHistoryResponse.ProtoReflect.Descriptor instead.
func (*ListHistoryResponse) Descriptor() ([]byte, []int) {
	return file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDescGZIP(), []int{1}
}

func (x *ListHistoryResponse) GetHistoryPage() *HistoryPage {
	if x != nil {
		return x.HistoryPage
	}
	return nil
}

func (x *ListHistoryResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListHistoryResponse) GetNextHistoryPageToken() *HistoryPageToken {
	if x != nil {
		return x.NextHistoryPageToken
	}
	return nil
}

// A query for entity history. Required fields are history type and time
// range. Downsampling and additional filter parameters are optional. If
// no additional filter is specified, all entity historical data of the
// specified type is returned for the provided time range.
type HistoryQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of history to query.
	HistoryType HistoryType `protobuf:"varint,1,opt,name=history_type,json=historyType,proto3,enum=anduril.entityhistory.v1.HistoryType" json:"history_type,omitempty"`
	// Time range for the below queries. If no query is specified, all data
	// for the history type during the time range is returned.
	TimeRange *TimeRange `protobuf:"bytes,2,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	// Downsampling to apply to a time ordered set of points. May not apply
	// to all history data types. Currently supported types are: trails.
	Downsample *Downsample `protobuf:"bytes,3,opt,name=downsample,proto3" json:"downsample,omitempty"`
	// Filter for historical data. The applied filter limits the data returned
	// by the query to only the data with matching current entity state. The
	// filter is applied to the latest observed state recorded in a snapshot.
	// Live state is not considered.
	Statement *v1.Statement `protobuf:"bytes,4,opt,name=statement,proto3" json:"statement,omitempty"`
	// Set of entityIds to request history data for. If empty, all entityIds will be found by querying TimeSeries.
	EntityIds []string `protobuf:"bytes,5,rep,name=entity_ids,json=entityIds,proto3" json:"entity_ids,omitempty"`
}

func (x *HistoryQuery) Reset() {
	*x = HistoryQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryQuery) ProtoMessage() {}

func (x *HistoryQuery) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryQuery.ProtoReflect.Descriptor instead.
func (*HistoryQuery) Descriptor() ([]byte, []int) {
	return file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDescGZIP(), []int{2}
}

func (x *HistoryQuery) GetHistoryType() HistoryType {
	if x != nil {
		return x.HistoryType
	}
	return HistoryType_HISTORY_TYPE_INVALID
}

func (x *HistoryQuery) GetTimeRange() *TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *HistoryQuery) GetDownsample() *Downsample {
	if x != nil {
		return x.Downsample
	}
	return nil
}

func (x *HistoryQuery) GetStatement() *v1.Statement {
	if x != nil {
		return x.Statement
	}
	return nil
}

func (x *HistoryQuery) GetEntityIds() []string {
	if x != nil {
		return x.EntityIds
	}
	return nil
}

// Range of wall-clock time. If upper bound timestamp is greater than the
// current time, the query will truncate the future time to now and return
// any history between lower bound and now. If both of the bounds are in
// the future, the query will report no available historical data.
type TimeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Earliest time in the TimeRange.
	LowerBoundInc *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=lower_bound_inc,json=lowerBoundInc,proto3" json:"lower_bound_inc,omitempty"`
	// Latest time in the TimeRange.
	UpperBoundExc *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=upper_bound_exc,json=upperBoundExc,proto3" json:"upper_bound_exc,omitempty"`
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDescGZIP(), []int{3}
}

func (x *TimeRange) GetLowerBoundInc() *timestamppb.Timestamp {
	if x != nil {
		return x.LowerBoundInc
	}
	return nil
}

func (x *TimeRange) GetUpperBoundExc() *timestamppb.Timestamp {
	if x != nil {
		return x.UpperBoundExc
	}
	return nil
}

// Downsampling is applied to points in an individual historical data model.
// Enforces each point is separated by at least downsample_duration millis.
//
// For example: downsample_duration = 10,000.
// stored points = [ 0, 5,000, 10,000, 13,000, 22,000, 30,000, 35,000 ].
// returned points = [ 0, 10,000, 22,000, 35,000 ].
type Downsample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*Downsample_DownsampleDuration
	Type isDownsample_Type `protobuf_oneof:"type"`
}

func (x *Downsample) Reset() {
	*x = Downsample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Downsample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Downsample) ProtoMessage() {}

func (x *Downsample) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Downsample.ProtoReflect.Descriptor instead.
func (*Downsample) Descriptor() ([]byte, []int) {
	return file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDescGZIP(), []int{4}
}

func (m *Downsample) GetType() isDownsample_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Downsample) GetDownsampleDuration() *DownsampleDuration {
	if x, ok := x.GetType().(*Downsample_DownsampleDuration); ok {
		return x.DownsampleDuration
	}
	return nil
}

type isDownsample_Type interface {
	isDownsample_Type()
}

type Downsample_DownsampleDuration struct {
	DownsampleDuration *DownsampleDuration `protobuf:"bytes,1,opt,name=downsample_duration,json=downsampleDuration,proto3,oneof"`
}

func (*Downsample_DownsampleDuration) isDownsample_Type() {}

type DownsampleDuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Downsample duration specified in milliseconds.
	DurationMs uint32 `protobuf:"varint,1,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
}

func (x *DownsampleDuration) Reset() {
	*x = DownsampleDuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownsampleDuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownsampleDuration) ProtoMessage() {}

func (x *DownsampleDuration) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownsampleDuration.ProtoReflect.Descriptor instead.
func (*DownsampleDuration) Descriptor() ([]byte, []int) {
	return file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDescGZIP(), []int{5}
}

func (x *DownsampleDuration) GetDurationMs() uint32 {
	if x != nil {
		return x.DurationMs
	}
	return 0
}

// Open a server stream to receive updates to entity time windows after Entity History has processed
// a backfill event or healing job. The stream is long-lived and only delivers new updates as they arrive.
type StreamBackfillUpdatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamBackfillUpdatesRequest) Reset() {
	*x = StreamBackfillUpdatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamBackfillUpdatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamBackfillUpdatesRequest) ProtoMessage() {}

func (x *StreamBackfillUpdatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamBackfillUpdatesRequest.ProtoReflect.Descriptor instead.
func (*StreamBackfillUpdatesRequest) Descriptor() ([]byte, []int) {
	return file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDescGZIP(), []int{6}
}

type StreamBackfillUpdatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BackfillUpdates []*BackfillUpdate `protobuf:"bytes,1,rep,name=backfill_updates,json=backfillUpdates,proto3" json:"backfill_updates,omitempty"`
}

func (x *StreamBackfillUpdatesResponse) Reset() {
	*x = StreamBackfillUpdatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamBackfillUpdatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamBackfillUpdatesResponse) ProtoMessage() {}

func (x *StreamBackfillUpdatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamBackfillUpdatesResponse.ProtoReflect.Descriptor instead.
func (*StreamBackfillUpdatesResponse) Descriptor() ([]byte, []int) {
	return file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDescGZIP(), []int{7}
}

func (x *StreamBackfillUpdatesResponse) GetBackfillUpdates() []*BackfillUpdate {
	if x != nil {
		return x.BackfillUpdates
	}
	return nil
}

// Indicates the entity identified by entity_id has new data available in EntityHistory for the time period described
// by time_range.
type BackfillUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId  string     `protobuf:"bytes,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	TimeRange *TimeRange `protobuf:"bytes,2,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
}

func (x *BackfillUpdate) Reset() {
	*x = BackfillUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackfillUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackfillUpdate) ProtoMessage() {}

func (x *BackfillUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackfillUpdate.ProtoReflect.Descriptor instead.
func (*BackfillUpdate) Descriptor() ([]byte, []int) {
	return file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDescGZIP(), []int{8}
}

func (x *BackfillUpdate) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *BackfillUpdate) GetTimeRange() *TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

var File_anduril_entityhistory_v1_entity_history_api_pub_proto protoreflect.FileDescriptor

var file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDesc = []byte{
	0x0a, 0x35, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x1a, 0x2a, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x61,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x4b, 0x0a, 0x0d, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x0c, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x58, 0x0a, 0x12,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72,
	0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x10, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xea, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48,
	0x0a, 0x0c, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x61, 0x0a, 0x17, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x14, 0x6e,
	0x65, 0x78, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0xc4, 0x02, 0x0a, 0x0c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x48, 0x0a, 0x0c, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x61, 0x6e, 0x64,
	0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x42,
	0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x44, 0x0a, 0x0a, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x0a, 0x64, 0x6f,
	0x77, 0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x41, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x09, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x63, 0x12, 0x42, 0x0a, 0x0f,
	0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x65, 0x78, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0d, 0x75, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x78, 0x63,
	0x22, 0x75, 0x0a, 0x0a, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x5f,
	0x0a, 0x13, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x12, 0x64, 0x6f, 0x77,
	0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x35, 0x0a, 0x12, 0x44, 0x6f, 0x77, 0x6e, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x22, 0x1e,
	0x0a, 0x1c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x74,
	0x0a, 0x1d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x53, 0x0a, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x73, 0x22, 0x71, 0x0a, 0x0e, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x32, 0x7e, 0x0a, 0x10, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x41, 0x50, 0x49, 0x12, 0x6a, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2c, 0x2e, 0x61, 0x6e, 0x64,
	0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72,
	0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x90, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x18, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x41, 0x70, 0x69, 0x50, 0x75, 0x62, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x6f, 0x67, 0x75, 0x6e, 0x2d, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x61,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72,
	0x63, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x45, 0x58,
	0xaa, 0x02, 0x18, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x41, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c,
	0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a,
	0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDescOnce sync.Once
	file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDescData = file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDesc
)

func file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDescGZIP() []byte {
	file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDescOnce.Do(func() {
		file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDescData)
	})
	return file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDescData
}

var file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_anduril_entityhistory_v1_entity_history_api_pub_proto_goTypes = []any{
	(*ListHistoryRequest)(nil),            // 0: anduril.entityhistory.v1.ListHistoryRequest
	(*ListHistoryResponse)(nil),           // 1: anduril.entityhistory.v1.ListHistoryResponse
	(*HistoryQuery)(nil),                  // 2: anduril.entityhistory.v1.HistoryQuery
	(*TimeRange)(nil),                     // 3: anduril.entityhistory.v1.TimeRange
	(*Downsample)(nil),                    // 4: anduril.entityhistory.v1.Downsample
	(*DownsampleDuration)(nil),            // 5: anduril.entityhistory.v1.DownsampleDuration
	(*StreamBackfillUpdatesRequest)(nil),  // 6: anduril.entityhistory.v1.StreamBackfillUpdatesRequest
	(*StreamBackfillUpdatesResponse)(nil), // 7: anduril.entityhistory.v1.StreamBackfillUpdatesResponse
	(*BackfillUpdate)(nil),                // 8: anduril.entityhistory.v1.BackfillUpdate
	(*HistoryPageToken)(nil),              // 9: anduril.entityhistory.v1.HistoryPageToken
	(*HistoryPage)(nil),                   // 10: anduril.entityhistory.v1.HistoryPage
	(HistoryType)(0),                      // 11: anduril.entityhistory.v1.HistoryType
	(*v1.Statement)(nil),                  // 12: anduril.entitymanager.v1.Statement
	(*timestamppb.Timestamp)(nil),         // 13: google.protobuf.Timestamp
}
var file_anduril_entityhistory_v1_entity_history_api_pub_proto_depIdxs = []int32{
	2,  // 0: anduril.entityhistory.v1.ListHistoryRequest.history_query:type_name -> anduril.entityhistory.v1.HistoryQuery
	9,  // 1: anduril.entityhistory.v1.ListHistoryRequest.history_page_token:type_name -> anduril.entityhistory.v1.HistoryPageToken
	10, // 2: anduril.entityhistory.v1.ListHistoryResponse.history_page:type_name -> anduril.entityhistory.v1.HistoryPage
	9,  // 3: anduril.entityhistory.v1.ListHistoryResponse.next_history_page_token:type_name -> anduril.entityhistory.v1.HistoryPageToken
	11, // 4: anduril.entityhistory.v1.HistoryQuery.history_type:type_name -> anduril.entityhistory.v1.HistoryType
	3,  // 5: anduril.entityhistory.v1.HistoryQuery.time_range:type_name -> anduril.entityhistory.v1.TimeRange
	4,  // 6: anduril.entityhistory.v1.HistoryQuery.downsample:type_name -> anduril.entityhistory.v1.Downsample
	12, // 7: anduril.entityhistory.v1.HistoryQuery.statement:type_name -> anduril.entitymanager.v1.Statement
	13, // 8: anduril.entityhistory.v1.TimeRange.lower_bound_inc:type_name -> google.protobuf.Timestamp
	13, // 9: anduril.entityhistory.v1.TimeRange.upper_bound_exc:type_name -> google.protobuf.Timestamp
	5,  // 10: anduril.entityhistory.v1.Downsample.downsample_duration:type_name -> anduril.entityhistory.v1.DownsampleDuration
	8,  // 11: anduril.entityhistory.v1.StreamBackfillUpdatesResponse.backfill_updates:type_name -> anduril.entityhistory.v1.BackfillUpdate
	3,  // 12: anduril.entityhistory.v1.BackfillUpdate.time_range:type_name -> anduril.entityhistory.v1.TimeRange
	0,  // 13: anduril.entityhistory.v1.EntityHistoryAPI.ListHistory:input_type -> anduril.entityhistory.v1.ListHistoryRequest
	1,  // 14: anduril.entityhistory.v1.EntityHistoryAPI.ListHistory:output_type -> anduril.entityhistory.v1.ListHistoryResponse
	14, // [14:15] is the sub-list for method output_type
	13, // [13:14] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_anduril_entityhistory_v1_entity_history_api_pub_proto_init() }
func file_anduril_entityhistory_v1_entity_history_api_pub_proto_init() {
	if File_anduril_entityhistory_v1_entity_history_api_pub_proto != nil {
		return
	}
	file_anduril_entityhistory_v1_history_pub_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListHistoryRequest); i {
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
		file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListHistoryResponse); i {
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
		file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*HistoryQuery); i {
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
		file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TimeRange); i {
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
		file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Downsample); i {
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
		file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DownsampleDuration); i {
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
		file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*StreamBackfillUpdatesRequest); i {
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
		file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*StreamBackfillUpdatesResponse); i {
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
		file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*BackfillUpdate); i {
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
	file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes[4].OneofWrappers = []any{
		(*Downsample_DownsampleDuration)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_anduril_entityhistory_v1_entity_history_api_pub_proto_goTypes,
		DependencyIndexes: file_anduril_entityhistory_v1_entity_history_api_pub_proto_depIdxs,
		MessageInfos:      file_anduril_entityhistory_v1_entity_history_api_pub_proto_msgTypes,
	}.Build()
	File_anduril_entityhistory_v1_entity_history_api_pub_proto = out.File
	file_anduril_entityhistory_v1_entity_history_api_pub_proto_rawDesc = nil
	file_anduril_entityhistory_v1_entity_history_api_pub_proto_goTypes = nil
	file_anduril_entityhistory_v1_entity_history_api_pub_proto_depIdxs = nil
}
