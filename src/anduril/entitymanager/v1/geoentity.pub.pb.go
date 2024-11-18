// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: anduril/entitymanager/v1/geoentity.pub.proto

package entitymanagerv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The type of geo entity.
type GeoType int32

const (
	GeoType_GEO_TYPE_INVALID   GeoType = 0
	GeoType_GEO_TYPE_GENERAL   GeoType = 1
	GeoType_GEO_TYPE_HAZARD    GeoType = 2
	GeoType_GEO_TYPE_EMERGENCY GeoType = 3
	// Fire support coordination measure
	GeoType_GEO_TYPE_FSCM GeoType = 4
	// Engagement zones allow for engaging an entity if it comes within the zone of another entity.
	GeoType_GEO_TYPE_ENGAGEMENT_ZONE GeoType = 5
	GeoType_GEO_TYPE_CONTROL_AREA    GeoType = 6
	GeoType_GEO_TYPE_BULLSEYE        GeoType = 7
	// Airspace Coordinating Measure
	GeoType_GEO_TYPE_ACM GeoType = 8
	// Maneuver Control Measure
	GeoType_GEO_TYPE_MCM GeoType = 9
)

// Enum value maps for GeoType.
var (
	GeoType_name = map[int32]string{
		0: "GEO_TYPE_INVALID",
		1: "GEO_TYPE_GENERAL",
		2: "GEO_TYPE_HAZARD",
		3: "GEO_TYPE_EMERGENCY",
		4: "GEO_TYPE_FSCM",
		5: "GEO_TYPE_ENGAGEMENT_ZONE",
		6: "GEO_TYPE_CONTROL_AREA",
		7: "GEO_TYPE_BULLSEYE",
		8: "GEO_TYPE_ACM",
		9: "GEO_TYPE_MCM",
	}
	GeoType_value = map[string]int32{
		"GEO_TYPE_INVALID":         0,
		"GEO_TYPE_GENERAL":         1,
		"GEO_TYPE_HAZARD":          2,
		"GEO_TYPE_EMERGENCY":       3,
		"GEO_TYPE_FSCM":            4,
		"GEO_TYPE_ENGAGEMENT_ZONE": 5,
		"GEO_TYPE_CONTROL_AREA":    6,
		"GEO_TYPE_BULLSEYE":        7,
		"GEO_TYPE_ACM":             8,
		"GEO_TYPE_MCM":             9,
	}
)

func (x GeoType) Enum() *GeoType {
	p := new(GeoType)
	*p = x
	return p
}

func (x GeoType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GeoType) Descriptor() protoreflect.EnumDescriptor {
	return file_anduril_entitymanager_v1_geoentity_pub_proto_enumTypes[0].Descriptor()
}

func (GeoType) Type() protoreflect.EnumType {
	return &file_anduril_entitymanager_v1_geoentity_pub_proto_enumTypes[0]
}

func (x GeoType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GeoType.Descriptor instead.
func (GeoType) EnumDescriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_geoentity_pub_proto_rawDescGZIP(), []int{0}
}

// A component that describes a geo-entity.
type GeoDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type GeoType `protobuf:"varint,1,opt,name=type,proto3,enum=anduril.entitymanager.v1.GeoType" json:"type,omitempty"`
}

func (x *GeoDetails) Reset() {
	*x = GeoDetails{}
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoDetails) ProtoMessage() {}

func (x *GeoDetails) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoDetails.ProtoReflect.Descriptor instead.
func (*GeoDetails) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_geoentity_pub_proto_rawDescGZIP(), []int{0}
}

func (x *GeoDetails) GetType() GeoType {
	if x != nil {
		return x.Type
	}
	return GeoType_GEO_TYPE_INVALID
}

// A component that describes the shape of a geo-entity.
type GeoShape struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// point, line_string, and polygon are convertible as a subset of GeoJSON
	//
	// Types that are assignable to Shape:
	//
	//	*GeoShape_Point
	//	*GeoShape_Line
	//	*GeoShape_Polygon
	//	*GeoShape_Ellipse
	//	*GeoShape_Ellipsoid
	Shape isGeoShape_Shape `protobuf_oneof:"shape"`
}

func (x *GeoShape) Reset() {
	*x = GeoShape{}
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoShape) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoShape) ProtoMessage() {}

func (x *GeoShape) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoShape.ProtoReflect.Descriptor instead.
func (*GeoShape) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_geoentity_pub_proto_rawDescGZIP(), []int{1}
}

func (m *GeoShape) GetShape() isGeoShape_Shape {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (x *GeoShape) GetPoint() *GeoPoint {
	if x, ok := x.GetShape().(*GeoShape_Point); ok {
		return x.Point
	}
	return nil
}

func (x *GeoShape) GetLine() *GeoLine {
	if x, ok := x.GetShape().(*GeoShape_Line); ok {
		return x.Line
	}
	return nil
}

func (x *GeoShape) GetPolygon() *GeoPolygon {
	if x, ok := x.GetShape().(*GeoShape_Polygon); ok {
		return x.Polygon
	}
	return nil
}

func (x *GeoShape) GetEllipse() *GeoEllipse {
	if x, ok := x.GetShape().(*GeoShape_Ellipse); ok {
		return x.Ellipse
	}
	return nil
}

func (x *GeoShape) GetEllipsoid() *GeoEllipsoid {
	if x, ok := x.GetShape().(*GeoShape_Ellipsoid); ok {
		return x.Ellipsoid
	}
	return nil
}

type isGeoShape_Shape interface {
	isGeoShape_Shape()
}

type GeoShape_Point struct {
	Point *GeoPoint `protobuf:"bytes,1,opt,name=point,proto3,oneof"`
}

type GeoShape_Line struct {
	Line *GeoLine `protobuf:"bytes,2,opt,name=line,proto3,oneof"`
}

type GeoShape_Polygon struct {
	Polygon *GeoPolygon `protobuf:"bytes,3,opt,name=polygon,proto3,oneof"`
}

type GeoShape_Ellipse struct {
	Ellipse *GeoEllipse `protobuf:"bytes,4,opt,name=ellipse,proto3,oneof"`
}

type GeoShape_Ellipsoid struct {
	Ellipsoid *GeoEllipsoid `protobuf:"bytes,5,opt,name=ellipsoid,proto3,oneof"`
}

func (*GeoShape_Point) isGeoShape_Shape() {}

func (*GeoShape_Line) isGeoShape_Shape() {}

func (*GeoShape_Polygon) isGeoShape_Shape() {}

func (*GeoShape_Ellipse) isGeoShape_Shape() {}

func (*GeoShape_Ellipsoid) isGeoShape_Shape() {}

// A point shaped geo-entity.
// See https://datatracker.ietf.org/doc/html/rfc7946#section-3.1.2
type GeoPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position *Position `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *GeoPoint) Reset() {
	*x = GeoPoint{}
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoPoint) ProtoMessage() {}

func (x *GeoPoint) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoPoint.ProtoReflect.Descriptor instead.
func (*GeoPoint) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_geoentity_pub_proto_rawDescGZIP(), []int{2}
}

func (x *GeoPoint) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

// A line shaped geo-entity.
// See https://datatracker.ietf.org/doc/html/rfc7946#section-3.1.4
type GeoLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Positions []*Position `protobuf:"bytes,1,rep,name=positions,proto3" json:"positions,omitempty"`
}

func (x *GeoLine) Reset() {
	*x = GeoLine{}
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoLine) ProtoMessage() {}

func (x *GeoLine) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoLine.ProtoReflect.Descriptor instead.
func (*GeoLine) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_geoentity_pub_proto_rawDescGZIP(), []int{3}
}

func (x *GeoLine) GetPositions() []*Position {
	if x != nil {
		return x.Positions
	}
	return nil
}

// A polygon shaped geo-entity.
// See https://datatracker.ietf.org/doc/html/rfc7946#section-3.1.6, only canonical representations accepted
type GeoPolygon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An array of LinearRings where the first item is the exterior ring and subsequent items are interior rings.
	// For a good introduction read (https://macwright.com/2015/03/23/geojson-second-bite.html#polygons)
	Rings []*LinearRing `protobuf:"bytes,1,rep,name=rings,proto3" json:"rings,omitempty"`
	// An extension hint that this polygon is a rectangle. When true this implies several things:
	// * exactly 1 linear ring with 5 points (starting corner, 3 other corners and start again)
	// * each point has the same altitude corresponding with the plane of the rectangle
	// * each point has the same height (either all present and equal, or all not present)
	IsRectangle bool `protobuf:"varint,2,opt,name=is_rectangle,json=isRectangle,proto3" json:"is_rectangle,omitempty"`
}

func (x *GeoPolygon) Reset() {
	*x = GeoPolygon{}
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoPolygon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoPolygon) ProtoMessage() {}

func (x *GeoPolygon) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoPolygon.ProtoReflect.Descriptor instead.
func (*GeoPolygon) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_geoentity_pub_proto_rawDescGZIP(), []int{4}
}

func (x *GeoPolygon) GetRings() []*LinearRing {
	if x != nil {
		return x.Rings
	}
	return nil
}

func (x *GeoPolygon) GetIsRectangle() bool {
	if x != nil {
		return x.IsRectangle
	}
	return false
}

// An ellipse shaped geo-entity.
// For a circle, the major and minor axis would be the same values.
// This shape is NOT Geo-JSON compatible.
type GeoEllipse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines the distance from the center point of the ellipse to the furthest distance on the perimeter in meters.
	SemiMajorAxisM *wrapperspb.DoubleValue `protobuf:"bytes,2,opt,name=semi_major_axis_m,json=semiMajorAxisM,proto3" json:"semi_major_axis_m,omitempty"`
	// Defines the distance from the center point of the ellipse to the shortest distance on the perimeter in meters.
	SemiMinorAxisM *wrapperspb.DoubleValue `protobuf:"bytes,3,opt,name=semi_minor_axis_m,json=semiMinorAxisM,proto3" json:"semi_minor_axis_m,omitempty"`
	// The orientation of the semi-major relative to true north in degrees from clockwise: 0-180 due to symmetry across the semi-minor axis.
	OrientationD *wrapperspb.DoubleValue `protobuf:"bytes,4,opt,name=orientation_d,json=orientationD,proto3" json:"orientation_d,omitempty"`
	// Optional height above entity position to extrude in meters. A non-zero value creates an elliptic cylinder
	HeightM *wrapperspb.DoubleValue `protobuf:"bytes,5,opt,name=height_m,json=heightM,proto3" json:"height_m,omitempty"`
}

func (x *GeoEllipse) Reset() {
	*x = GeoEllipse{}
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoEllipse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoEllipse) ProtoMessage() {}

func (x *GeoEllipse) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoEllipse.ProtoReflect.Descriptor instead.
func (*GeoEllipse) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_geoentity_pub_proto_rawDescGZIP(), []int{5}
}

func (x *GeoEllipse) GetSemiMajorAxisM() *wrapperspb.DoubleValue {
	if x != nil {
		return x.SemiMajorAxisM
	}
	return nil
}

func (x *GeoEllipse) GetSemiMinorAxisM() *wrapperspb.DoubleValue {
	if x != nil {
		return x.SemiMinorAxisM
	}
	return nil
}

func (x *GeoEllipse) GetOrientationD() *wrapperspb.DoubleValue {
	if x != nil {
		return x.OrientationD
	}
	return nil
}

func (x *GeoEllipse) GetHeightM() *wrapperspb.DoubleValue {
	if x != nil {
		return x.HeightM
	}
	return nil
}

// An ellipsoid shaped geo-entity.
// Principal axis lengths are defined in entity body space
// This shape is NOT Geo-JSON compatible.
type GeoEllipsoid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines the distance from the center point to the surface along the forward axis
	ForwardAxisM *wrapperspb.DoubleValue `protobuf:"bytes,1,opt,name=forward_axis_m,json=forwardAxisM,proto3" json:"forward_axis_m,omitempty"`
	// Defines the distance from the center point to the surface along the side axis
	SideAxisM *wrapperspb.DoubleValue `protobuf:"bytes,2,opt,name=side_axis_m,json=sideAxisM,proto3" json:"side_axis_m,omitempty"`
	// Defines the distance from the center point to the surface along the up axis
	UpAxisM *wrapperspb.DoubleValue `protobuf:"bytes,3,opt,name=up_axis_m,json=upAxisM,proto3" json:"up_axis_m,omitempty"`
}

func (x *GeoEllipsoid) Reset() {
	*x = GeoEllipsoid{}
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoEllipsoid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoEllipsoid) ProtoMessage() {}

func (x *GeoEllipsoid) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoEllipsoid.ProtoReflect.Descriptor instead.
func (*GeoEllipsoid) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_geoentity_pub_proto_rawDescGZIP(), []int{6}
}

func (x *GeoEllipsoid) GetForwardAxisM() *wrapperspb.DoubleValue {
	if x != nil {
		return x.ForwardAxisM
	}
	return nil
}

func (x *GeoEllipsoid) GetSideAxisM() *wrapperspb.DoubleValue {
	if x != nil {
		return x.SideAxisM
	}
	return nil
}

func (x *GeoEllipsoid) GetUpAxisM() *wrapperspb.DoubleValue {
	if x != nil {
		return x.UpAxisM
	}
	return nil
}

// A closed ring of points. The first and last point must be the same.
type LinearRing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: do not use, use positions instead
	//
	// Deprecated: Marked as deprecated in anduril/entitymanager/v1/geoentity.pub.proto.
	Points    []*Position           `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	Positions []*GeoPolygonPosition `protobuf:"bytes,2,rep,name=positions,proto3" json:"positions,omitempty"`
}

func (x *LinearRing) Reset() {
	*x = LinearRing{}
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LinearRing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinearRing) ProtoMessage() {}

func (x *LinearRing) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinearRing.ProtoReflect.Descriptor instead.
func (*LinearRing) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_geoentity_pub_proto_rawDescGZIP(), []int{7}
}

// Deprecated: Marked as deprecated in anduril/entitymanager/v1/geoentity.pub.proto.
func (x *LinearRing) GetPoints() []*Position {
	if x != nil {
		return x.Points
	}
	return nil
}

func (x *LinearRing) GetPositions() []*GeoPolygonPosition {
	if x != nil {
		return x.Positions
	}
	return nil
}

// A position in a GeoPolygon with an optional extruded height.
type GeoPolygonPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// base position. if no altitude set, its on the ground.
	Position *Position `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	// optional height above base position to extrude in meters.
	// for a given polygon, all points should have a height or none of them.
	// strictly GeoJSON compatible polygons will not have this set.
	HeightM *wrapperspb.FloatValue `protobuf:"bytes,2,opt,name=height_m,json=heightM,proto3" json:"height_m,omitempty"`
}

func (x *GeoPolygonPosition) Reset() {
	*x = GeoPolygonPosition{}
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoPolygonPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoPolygonPosition) ProtoMessage() {}

func (x *GeoPolygonPosition) ProtoReflect() protoreflect.Message {
	mi := &file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoPolygonPosition.ProtoReflect.Descriptor instead.
func (*GeoPolygonPosition) Descriptor() ([]byte, []int) {
	return file_anduril_entitymanager_v1_geoentity_pub_proto_rawDescGZIP(), []int{8}
}

func (x *GeoPolygonPosition) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *GeoPolygonPosition) GetHeightM() *wrapperspb.FloatValue {
	if x != nil {
		return x.HeightM
	}
	return nil
}

var File_anduril_entitymanager_v1_geoentity_pub_proto protoreflect.FileDescriptor

var file_anduril_entitymanager_v1_geoentity_pub_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x2b, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x75, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x6f, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x21, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x6f, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x4a, 0x04, 0x08, 0x02, 0x10,
	0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x4a, 0x04, 0x08,
	0x05, 0x10, 0x06, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x4a, 0x04, 0x08, 0x07, 0x10, 0x08, 0x52,
	0x09, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0e, 0x76, 0x69, 0x73, 0x75,
	0x61, 0x6c, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x04, 0x66, 0x73, 0x63, 0x6d,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x52, 0x03,
	0x61, 0x63, 0x6d, 0x52, 0x03, 0x6d, 0x63, 0x6d, 0x22, 0xd4, 0x02, 0x0a, 0x08, 0x47, 0x65, 0x6f,
	0x53, 0x68, 0x61, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x37, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f, 0x4c, 0x69,
	0x6e, 0x65, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x70, 0x6f,
	0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x07,
	0x65, 0x6c, 0x6c, 0x69, 0x70, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f, 0x45, 0x6c, 0x6c, 0x69,
	0x70, 0x73, 0x65, 0x48, 0x00, 0x52, 0x07, 0x65, 0x6c, 0x6c, 0x69, 0x70, 0x73, 0x65, 0x12, 0x46,
	0x0a, 0x09, 0x65, 0x6c, 0x6c, 0x69, 0x70, 0x73, 0x6f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f,
	0x45, 0x6c, 0x6c, 0x69, 0x70, 0x73, 0x6f, 0x69, 0x64, 0x48, 0x00, 0x52, 0x09, 0x65, 0x6c, 0x6c,
	0x69, 0x70, 0x73, 0x6f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x22,
	0x4a, 0x0a, 0x08, 0x47, 0x65, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x07, 0x47,
	0x65, 0x6f, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x6e, 0x64, 0x75,
	0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x6b, 0x0a, 0x0a, 0x47, 0x65, 0x6f, 0x50,
	0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x05, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x52, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x72, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x74, 0x61, 0x6e, 0x67,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x52, 0x65, 0x63, 0x74,
	0x61, 0x6e, 0x67, 0x6c, 0x65, 0x22, 0x9a, 0x02, 0x0a, 0x0a, 0x47, 0x65, 0x6f, 0x45, 0x6c, 0x6c,
	0x69, 0x70, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x11, 0x73, 0x65, 0x6d, 0x69, 0x5f, 0x6d, 0x61, 0x6a,
	0x6f, 0x72, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x5f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x73,
	0x65, 0x6d, 0x69, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x41, 0x78, 0x69, 0x73, 0x4d, 0x12, 0x47, 0x0a,
	0x11, 0x73, 0x65, 0x6d, 0x69, 0x5f, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x5f, 0x61, 0x78, 0x69, 0x73,
	0x5f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x73, 0x65, 0x6d, 0x69, 0x4d, 0x69, 0x6e, 0x6f,
	0x72, 0x41, 0x78, 0x69, 0x73, 0x4d, 0x12, 0x41, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x6f, 0x72, 0x69,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x12, 0x37, 0x0a, 0x08, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x5f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x4d, 0x22, 0xca, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x6f, 0x45, 0x6c, 0x6c, 0x69, 0x70, 0x73,
	0x6f, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x0e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x61,
	0x78, 0x69, 0x73, 0x5f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x41, 0x78, 0x69, 0x73, 0x4d, 0x12, 0x3c, 0x0a, 0x0b, 0x73, 0x69, 0x64, 0x65, 0x5f,
	0x61, 0x78, 0x69, 0x73, 0x5f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x73, 0x69, 0x64, 0x65,
	0x41, 0x78, 0x69, 0x73, 0x4d, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x5f, 0x61, 0x78, 0x69, 0x73,
	0x5f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x75, 0x70, 0x41, 0x78, 0x69, 0x73, 0x4d, 0x22,
	0x98, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x52, 0x69, 0x6e, 0x67, 0x12, 0x3e,
	0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x4a,
	0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f,
	0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x12, 0x47,
	0x65, 0x6f, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3e, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x36, 0x0a, 0x08, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x07, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x4d, 0x2a, 0xe9, 0x01, 0x0a, 0x07, 0x47, 0x65,
	0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x45, 0x4f, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x47,
	0x45, 0x4f, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x4c, 0x10,
	0x01, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x45, 0x4f, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x41,
	0x5a, 0x41, 0x52, 0x44, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x45, 0x4f, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x45, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x03, 0x12, 0x11,
	0x0a, 0x0d, 0x47, 0x45, 0x4f, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x53, 0x43, 0x4d, 0x10,
	0x04, 0x12, 0x1c, 0x0a, 0x18, 0x47, 0x45, 0x4f, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e,
	0x47, 0x41, 0x47, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x5a, 0x4f, 0x4e, 0x45, 0x10, 0x05, 0x12,
	0x19, 0x0a, 0x15, 0x47, 0x45, 0x4f, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54,
	0x52, 0x4f, 0x4c, 0x5f, 0x41, 0x52, 0x45, 0x41, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x45,
	0x4f, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x55, 0x4c, 0x4c, 0x53, 0x45, 0x59, 0x45, 0x10,
	0x07, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x45, 0x4f, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43,
	0x4d, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x45, 0x4f, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4d, 0x43, 0x4d, 0x10, 0x09, 0x42, 0x83, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e,
	0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x47, 0x65, 0x6f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x50, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f,
	0x6c, 0x61, 0x74, 0x74, 0x69, 0x63, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x73,
	0x72, 0x63, 0x2f, 0x61, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x45,
	0x58, 0xaa, 0x02, 0x18, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x2e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x41,
	0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69,
	0x6c, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x1a, 0x41, 0x6e, 0x64, 0x75, 0x72, 0x69, 0x6c, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_anduril_entitymanager_v1_geoentity_pub_proto_rawDescOnce sync.Once
	file_anduril_entitymanager_v1_geoentity_pub_proto_rawDescData = file_anduril_entitymanager_v1_geoentity_pub_proto_rawDesc
)

func file_anduril_entitymanager_v1_geoentity_pub_proto_rawDescGZIP() []byte {
	file_anduril_entitymanager_v1_geoentity_pub_proto_rawDescOnce.Do(func() {
		file_anduril_entitymanager_v1_geoentity_pub_proto_rawDescData = protoimpl.X.CompressGZIP(file_anduril_entitymanager_v1_geoentity_pub_proto_rawDescData)
	})
	return file_anduril_entitymanager_v1_geoentity_pub_proto_rawDescData
}

var file_anduril_entitymanager_v1_geoentity_pub_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_anduril_entitymanager_v1_geoentity_pub_proto_goTypes = []any{
	(GeoType)(0),                   // 0: anduril.entitymanager.v1.GeoType
	(*GeoDetails)(nil),             // 1: anduril.entitymanager.v1.GeoDetails
	(*GeoShape)(nil),               // 2: anduril.entitymanager.v1.GeoShape
	(*GeoPoint)(nil),               // 3: anduril.entitymanager.v1.GeoPoint
	(*GeoLine)(nil),                // 4: anduril.entitymanager.v1.GeoLine
	(*GeoPolygon)(nil),             // 5: anduril.entitymanager.v1.GeoPolygon
	(*GeoEllipse)(nil),             // 6: anduril.entitymanager.v1.GeoEllipse
	(*GeoEllipsoid)(nil),           // 7: anduril.entitymanager.v1.GeoEllipsoid
	(*LinearRing)(nil),             // 8: anduril.entitymanager.v1.LinearRing
	(*GeoPolygonPosition)(nil),     // 9: anduril.entitymanager.v1.GeoPolygonPosition
	(*Position)(nil),               // 10: anduril.entitymanager.v1.Position
	(*wrapperspb.DoubleValue)(nil), // 11: google.protobuf.DoubleValue
	(*wrapperspb.FloatValue)(nil),  // 12: google.protobuf.FloatValue
}
var file_anduril_entitymanager_v1_geoentity_pub_proto_depIdxs = []int32{
	0,  // 0: anduril.entitymanager.v1.GeoDetails.type:type_name -> anduril.entitymanager.v1.GeoType
	3,  // 1: anduril.entitymanager.v1.GeoShape.point:type_name -> anduril.entitymanager.v1.GeoPoint
	4,  // 2: anduril.entitymanager.v1.GeoShape.line:type_name -> anduril.entitymanager.v1.GeoLine
	5,  // 3: anduril.entitymanager.v1.GeoShape.polygon:type_name -> anduril.entitymanager.v1.GeoPolygon
	6,  // 4: anduril.entitymanager.v1.GeoShape.ellipse:type_name -> anduril.entitymanager.v1.GeoEllipse
	7,  // 5: anduril.entitymanager.v1.GeoShape.ellipsoid:type_name -> anduril.entitymanager.v1.GeoEllipsoid
	10, // 6: anduril.entitymanager.v1.GeoPoint.position:type_name -> anduril.entitymanager.v1.Position
	10, // 7: anduril.entitymanager.v1.GeoLine.positions:type_name -> anduril.entitymanager.v1.Position
	8,  // 8: anduril.entitymanager.v1.GeoPolygon.rings:type_name -> anduril.entitymanager.v1.LinearRing
	11, // 9: anduril.entitymanager.v1.GeoEllipse.semi_major_axis_m:type_name -> google.protobuf.DoubleValue
	11, // 10: anduril.entitymanager.v1.GeoEllipse.semi_minor_axis_m:type_name -> google.protobuf.DoubleValue
	11, // 11: anduril.entitymanager.v1.GeoEllipse.orientation_d:type_name -> google.protobuf.DoubleValue
	11, // 12: anduril.entitymanager.v1.GeoEllipse.height_m:type_name -> google.protobuf.DoubleValue
	11, // 13: anduril.entitymanager.v1.GeoEllipsoid.forward_axis_m:type_name -> google.protobuf.DoubleValue
	11, // 14: anduril.entitymanager.v1.GeoEllipsoid.side_axis_m:type_name -> google.protobuf.DoubleValue
	11, // 15: anduril.entitymanager.v1.GeoEllipsoid.up_axis_m:type_name -> google.protobuf.DoubleValue
	10, // 16: anduril.entitymanager.v1.LinearRing.points:type_name -> anduril.entitymanager.v1.Position
	9,  // 17: anduril.entitymanager.v1.LinearRing.positions:type_name -> anduril.entitymanager.v1.GeoPolygonPosition
	10, // 18: anduril.entitymanager.v1.GeoPolygonPosition.position:type_name -> anduril.entitymanager.v1.Position
	12, // 19: anduril.entitymanager.v1.GeoPolygonPosition.height_m:type_name -> google.protobuf.FloatValue
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_anduril_entitymanager_v1_geoentity_pub_proto_init() }
func file_anduril_entitymanager_v1_geoentity_pub_proto_init() {
	if File_anduril_entitymanager_v1_geoentity_pub_proto != nil {
		return
	}
	file_anduril_entitymanager_v1_location_pub_proto_init()
	file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes[1].OneofWrappers = []any{
		(*GeoShape_Point)(nil),
		(*GeoShape_Line)(nil),
		(*GeoShape_Polygon)(nil),
		(*GeoShape_Ellipse)(nil),
		(*GeoShape_Ellipsoid)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_anduril_entitymanager_v1_geoentity_pub_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anduril_entitymanager_v1_geoentity_pub_proto_goTypes,
		DependencyIndexes: file_anduril_entitymanager_v1_geoentity_pub_proto_depIdxs,
		EnumInfos:         file_anduril_entitymanager_v1_geoentity_pub_proto_enumTypes,
		MessageInfos:      file_anduril_entitymanager_v1_geoentity_pub_proto_msgTypes,
	}.Build()
	File_anduril_entitymanager_v1_geoentity_pub_proto = out.File
	file_anduril_entitymanager_v1_geoentity_pub_proto_rawDesc = nil
	file_anduril_entitymanager_v1_geoentity_pub_proto_goTypes = nil
	file_anduril_entitymanager_v1_geoentity_pub_proto_depIdxs = nil
}
