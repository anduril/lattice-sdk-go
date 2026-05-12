// Code generated from our API definition. DO NOT EDIT.

package Lattice

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/anduril/lattice-sdk-go/v4/internal"
	big "math/big"
	time "time"
)

var (
	deleteObjectRequestFieldObjectPath = big.NewInt(1 << 0)
)

type DeleteObjectRequest struct {
	// The path of the object to delete.
	ObjectPath string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteObjectRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetObjectPath sets the ObjectPath field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteObjectRequest) SetObjectPath(objectPath string) {
	d.ObjectPath = objectPath
	d.require(deleteObjectRequestFieldObjectPath)
}

var (
	getObjectRequestFieldAcceptEncoding = big.NewInt(1 << 0)
	getObjectRequestFieldPriority       = big.NewInt(1 << 1)
	getObjectRequestFieldObjectPath     = big.NewInt(1 << 2)
)

type GetObjectRequest struct {
	// If set, Lattice will compress the response using the specified compression method. If the header is not defined, or the compression method is set to `identity`, no compression will be applied to the response.
	AcceptEncoding *GetObjectRequestAcceptEncoding `json:"-" url:"-"`
	// Indicates a client's preference for the priority of the response. The value is a structured header as defined in RFC 9218. If you do not set the header, Lattice uses the default priority set for the environment. Incremental delivery directives are not supported and will be ignored.
	Priority *string `json:"-" url:"-"`
	// The path of the object to fetch.
	ObjectPath string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetObjectRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetAcceptEncoding sets the AcceptEncoding field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetObjectRequest) SetAcceptEncoding(acceptEncoding *GetObjectRequestAcceptEncoding) {
	g.AcceptEncoding = acceptEncoding
	g.require(getObjectRequestFieldAcceptEncoding)
}

// SetPriority sets the Priority field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetObjectRequest) SetPriority(priority *string) {
	g.Priority = priority
	g.require(getObjectRequestFieldPriority)
}

// SetObjectPath sets the ObjectPath field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetObjectRequest) SetObjectPath(objectPath string) {
	g.ObjectPath = objectPath
	g.require(getObjectRequestFieldObjectPath)
}

var (
	getObjectMetadataRequestFieldObjectPath = big.NewInt(1 << 0)
)

type GetObjectMetadataRequest struct {
	// The path of the object to query.
	ObjectPath string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetObjectMetadataRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetObjectPath sets the ObjectPath field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetObjectMetadataRequest) SetObjectPath(objectPath string) {
	g.ObjectPath = objectPath
	g.require(getObjectMetadataRequestFieldObjectPath)
}

var (
	listDeletedObjectsRequestFieldPageToken   = big.NewInt(1 << 0)
	listDeletedObjectsRequestFieldMaxPageSize = big.NewInt(1 << 1)
)

type ListDeletedObjectsRequest struct {
	// Opaque cursor returned by a prior response to continue paging.
	PageToken *string `json:"-" url:"pageToken,omitempty"`
	// Maximum number of records to return in a single response. Server enforces an upper bound.
	MaxPageSize *int `json:"-" url:"maxPageSize,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListDeletedObjectsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPageToken sets the PageToken field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListDeletedObjectsRequest) SetPageToken(pageToken *string) {
	l.PageToken = pageToken
	l.require(listDeletedObjectsRequestFieldPageToken)
}

// SetMaxPageSize sets the MaxPageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListDeletedObjectsRequest) SetMaxPageSize(maxPageSize *int) {
	l.MaxPageSize = maxPageSize
	l.require(listDeletedObjectsRequestFieldMaxPageSize)
}

var (
	listObjectsRequestFieldPrefix           = big.NewInt(1 << 0)
	listObjectsRequestFieldSinceTimestamp   = big.NewInt(1 << 1)
	listObjectsRequestFieldPageToken        = big.NewInt(1 << 2)
	listObjectsRequestFieldAllObjectsInMesh = big.NewInt(1 << 3)
	listObjectsRequestFieldMaxPageSize      = big.NewInt(1 << 4)
)

type ListObjectsRequest struct {
	// Filters the objects based on the specified prefix path. If no path is specified, all objects are returned.
	Prefix *string `json:"-" url:"prefix,omitempty"`
	// Sets the age for the oldest objects to query across the environment.
	SinceTimestamp *time.Time `json:"-" url:"sinceTimestamp,omitempty"`
	// Base64 and URL-encoded cursor returned by the service to continue paging.
	PageToken *string `json:"-" url:"pageToken,omitempty"`
	// Lists objects across all environment nodes in a Lattice Mesh.
	AllObjectsInMesh *bool `json:"-" url:"allObjectsInMesh,omitempty"`
	// Sets the maximum number of items that should be returned on a single page.
	MaxPageSize *int `json:"-" url:"maxPageSize,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListObjectsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPrefix sets the Prefix field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListObjectsRequest) SetPrefix(prefix *string) {
	l.Prefix = prefix
	l.require(listObjectsRequestFieldPrefix)
}

// SetSinceTimestamp sets the SinceTimestamp field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListObjectsRequest) SetSinceTimestamp(sinceTimestamp *time.Time) {
	l.SinceTimestamp = sinceTimestamp
	l.require(listObjectsRequestFieldSinceTimestamp)
}

// SetPageToken sets the PageToken field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListObjectsRequest) SetPageToken(pageToken *string) {
	l.PageToken = pageToken
	l.require(listObjectsRequestFieldPageToken)
}

// SetAllObjectsInMesh sets the AllObjectsInMesh field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListObjectsRequest) SetAllObjectsInMesh(allObjectsInMesh *bool) {
	l.AllObjectsInMesh = allObjectsInMesh
	l.require(listObjectsRequestFieldAllObjectsInMesh)
}

// SetMaxPageSize sets the MaxPageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListObjectsRequest) SetMaxPageSize(maxPageSize *int) {
	l.MaxPageSize = maxPageSize
	l.require(listObjectsRequestFieldMaxPageSize)
}

var (
	contentIdentifierFieldPath     = big.NewInt(1 << 0)
	contentIdentifierFieldChecksum = big.NewInt(1 << 1)
)

type ContentIdentifier struct {
	// A valid path must not contain the following:
	// - Spaces or Tabs
	// - Special characters other than underscore (_), dash (-), period (.) and slash (/)
	// - Non-ASCII characters such as accents or symbols
	// Paths must not start with a leading space.
	Path string `json:"path" url:"path"`
	// The SHA-256 checksum of this object.
	Checksum string `json:"checksum" url:"checksum"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *ContentIdentifier) GetPath() string {
	if c == nil {
		return ""
	}
	return c.Path
}

func (c *ContentIdentifier) GetChecksum() string {
	if c == nil {
		return ""
	}
	return c.Checksum
}

func (c *ContentIdentifier) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *ContentIdentifier) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetPath sets the Path field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ContentIdentifier) SetPath(path string) {
	c.Path = path
	c.require(contentIdentifierFieldPath)
}

// SetChecksum sets the Checksum field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ContentIdentifier) SetChecksum(checksum string) {
	c.Checksum = checksum
	c.require(contentIdentifierFieldChecksum)
}

func (c *ContentIdentifier) UnmarshalJSON(data []byte) error {
	type unmarshaler ContentIdentifier
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ContentIdentifier(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *ContentIdentifier) MarshalJSON() ([]byte, error) {
	type embed ContentIdentifier
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *ContentIdentifier) String() string {
	if c == nil {
		return "<nil>"
	}
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

var (
	deletedObjectEntryFieldPath      = big.NewInt(1 << 0)
	deletedObjectEntryFieldChecksum  = big.NewInt(1 << 1)
	deletedObjectEntryFieldDeletedAt = big.NewInt(1 << 2)
)

type DeletedObjectEntry struct {
	// Object path that was deleted on this node.
	// A valid path must not contain the following:
	// - Spaces or Tabs
	// - Special characters other than underscore (_), dash (-), period (.) and slash (/)
	// - Non-ASCII characters such as accents or symbols
	// Paths must not start with a leading space.
	Path string `json:"path" url:"path"`
	// The SHA-256 checksum of this object.
	Checksum string `json:"checksum" url:"checksum"`
	// Wall-clock time at which the deletion was initiated on this node.
	DeletedAt time.Time `json:"deleted_at" url:"deleted_at"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DeletedObjectEntry) GetPath() string {
	if d == nil {
		return ""
	}
	return d.Path
}

func (d *DeletedObjectEntry) GetChecksum() string {
	if d == nil {
		return ""
	}
	return d.Checksum
}

func (d *DeletedObjectEntry) GetDeletedAt() time.Time {
	if d == nil {
		return time.Time{}
	}
	return d.DeletedAt
}

func (d *DeletedObjectEntry) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d *DeletedObjectEntry) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetPath sets the Path field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeletedObjectEntry) SetPath(path string) {
	d.Path = path
	d.require(deletedObjectEntryFieldPath)
}

// SetChecksum sets the Checksum field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeletedObjectEntry) SetChecksum(checksum string) {
	d.Checksum = checksum
	d.require(deletedObjectEntryFieldChecksum)
}

// SetDeletedAt sets the DeletedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeletedObjectEntry) SetDeletedAt(deletedAt time.Time) {
	d.DeletedAt = deletedAt
	d.require(deletedObjectEntryFieldDeletedAt)
}

func (d *DeletedObjectEntry) UnmarshalJSON(data []byte) error {
	type embed DeletedObjectEntry
	var unmarshaler = struct {
		embed
		DeletedAt *internal.DateTime `json:"deleted_at"`
	}{
		embed: embed(*d),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*d = DeletedObjectEntry(unmarshaler.embed)
	d.DeletedAt = unmarshaler.DeletedAt.Time()
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeletedObjectEntry) MarshalJSON() ([]byte, error) {
	type embed DeletedObjectEntry
	var marshaler = struct {
		embed
		DeletedAt *internal.DateTime `json:"deleted_at"`
	}{
		embed:     embed(*d),
		DeletedAt: internal.NewDateTime(d.DeletedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, d.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (d *DeletedObjectEntry) String() string {
	if d == nil {
		return "<nil>"
	}
	if len(d.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

var (
	listDeletedObjectsResponseFieldDeletedObjects = big.NewInt(1 << 0)
	listDeletedObjectsResponseFieldNextPageToken  = big.NewInt(1 << 1)
)

type ListDeletedObjectsResponse struct {
	DeletedObjects []*DeletedObjectEntry `json:"deleted_objects" url:"deleted_objects"`
	// Present when more pages are available. Pass back as `pageToken`.
	NextPageToken *string `json:"next_page_token,omitempty" url:"next_page_token,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListDeletedObjectsResponse) GetDeletedObjects() []*DeletedObjectEntry {
	if l == nil {
		return nil
	}
	return l.DeletedObjects
}

func (l *ListDeletedObjectsResponse) GetNextPageToken() *string {
	if l == nil {
		return nil
	}
	return l.NextPageToken
}

func (l *ListDeletedObjectsResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListDeletedObjectsResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetDeletedObjects sets the DeletedObjects field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListDeletedObjectsResponse) SetDeletedObjects(deletedObjects []*DeletedObjectEntry) {
	l.DeletedObjects = deletedObjects
	l.require(listDeletedObjectsResponseFieldDeletedObjects)
}

// SetNextPageToken sets the NextPageToken field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListDeletedObjectsResponse) SetNextPageToken(nextPageToken *string) {
	l.NextPageToken = nextPageToken
	l.require(listDeletedObjectsResponseFieldNextPageToken)
}

func (l *ListDeletedObjectsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListDeletedObjectsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListDeletedObjectsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListDeletedObjectsResponse) MarshalJSON() ([]byte, error) {
	type embed ListDeletedObjectsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListDeletedObjectsResponse) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

var (
	listResponseFieldPathMetadatas = big.NewInt(1 << 0)
	listResponseFieldNextPageToken = big.NewInt(1 << 1)
)

type ListResponse struct {
	PathMetadatas []*PathMetadata `json:"path_metadatas" url:"path_metadatas"`
	NextPageToken *string         `json:"next_page_token,omitempty" url:"next_page_token,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListResponse) GetPathMetadatas() []*PathMetadata {
	if l == nil {
		return nil
	}
	return l.PathMetadatas
}

func (l *ListResponse) GetNextPageToken() *string {
	if l == nil {
		return nil
	}
	return l.NextPageToken
}

func (l *ListResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPathMetadatas sets the PathMetadatas field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListResponse) SetPathMetadatas(pathMetadatas []*PathMetadata) {
	l.PathMetadatas = pathMetadatas
	l.require(listResponseFieldPathMetadatas)
}

// SetNextPageToken sets the NextPageToken field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListResponse) SetNextPageToken(nextPageToken *string) {
	l.NextPageToken = nextPageToken
	l.require(listResponseFieldNextPageToken)
}

func (l *ListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListResponse) MarshalJSON() ([]byte, error) {
	type embed ListResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListResponse) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

var (
	pathMetadataFieldContentIdentifier = big.NewInt(1 << 0)
	pathMetadataFieldSizeBytes         = big.NewInt(1 << 1)
	pathMetadataFieldLastUpdatedAt     = big.NewInt(1 << 2)
	pathMetadataFieldExpiryTime        = big.NewInt(1 << 3)
)

type PathMetadata struct {
	ContentIdentifier *ContentIdentifier `json:"content_identifier" url:"content_identifier"`
	SizeBytes         int64              `json:"size_bytes" url:"size_bytes"`
	LastUpdatedAt     time.Time          `json:"last_updated_at" url:"last_updated_at"`
	ExpiryTime        *time.Time         `json:"expiry_time,omitempty" url:"expiry_time,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PathMetadata) GetContentIdentifier() *ContentIdentifier {
	if p == nil {
		return nil
	}
	return p.ContentIdentifier
}

func (p *PathMetadata) GetSizeBytes() int64 {
	if p == nil {
		return 0
	}
	return p.SizeBytes
}

func (p *PathMetadata) GetLastUpdatedAt() time.Time {
	if p == nil {
		return time.Time{}
	}
	return p.LastUpdatedAt
}

func (p *PathMetadata) GetExpiryTime() *time.Time {
	if p == nil {
		return nil
	}
	return p.ExpiryTime
}

func (p *PathMetadata) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p *PathMetadata) require(field *big.Int) {
	if p.explicitFields == nil {
		p.explicitFields = big.NewInt(0)
	}
	p.explicitFields.Or(p.explicitFields, field)
}

// SetContentIdentifier sets the ContentIdentifier field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PathMetadata) SetContentIdentifier(contentIdentifier *ContentIdentifier) {
	p.ContentIdentifier = contentIdentifier
	p.require(pathMetadataFieldContentIdentifier)
}

// SetSizeBytes sets the SizeBytes field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PathMetadata) SetSizeBytes(sizeBytes int64) {
	p.SizeBytes = sizeBytes
	p.require(pathMetadataFieldSizeBytes)
}

// SetLastUpdatedAt sets the LastUpdatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PathMetadata) SetLastUpdatedAt(lastUpdatedAt time.Time) {
	p.LastUpdatedAt = lastUpdatedAt
	p.require(pathMetadataFieldLastUpdatedAt)
}

// SetExpiryTime sets the ExpiryTime field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PathMetadata) SetExpiryTime(expiryTime *time.Time) {
	p.ExpiryTime = expiryTime
	p.require(pathMetadataFieldExpiryTime)
}

func (p *PathMetadata) UnmarshalJSON(data []byte) error {
	type embed PathMetadata
	var unmarshaler = struct {
		embed
		LastUpdatedAt *internal.DateTime `json:"last_updated_at"`
		ExpiryTime    *internal.DateTime `json:"expiry_time,omitempty"`
	}{
		embed: embed(*p),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*p = PathMetadata(unmarshaler.embed)
	p.LastUpdatedAt = unmarshaler.LastUpdatedAt.Time()
	p.ExpiryTime = unmarshaler.ExpiryTime.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PathMetadata) MarshalJSON() ([]byte, error) {
	type embed PathMetadata
	var marshaler = struct {
		embed
		LastUpdatedAt *internal.DateTime `json:"last_updated_at"`
		ExpiryTime    *internal.DateTime `json:"expiry_time,omitempty"`
	}{
		embed:         embed(*p),
		LastUpdatedAt: internal.NewDateTime(p.LastUpdatedAt),
		ExpiryTime:    internal.NewOptionalDateTime(p.ExpiryTime),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, p.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (p *PathMetadata) String() string {
	if p == nil {
		return "<nil>"
	}
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type GetObjectRequestAcceptEncoding string

const (
	GetObjectRequestAcceptEncodingIdentity GetObjectRequestAcceptEncoding = "identity"
	GetObjectRequestAcceptEncodingZstd     GetObjectRequestAcceptEncoding = "zstd"
)

func NewGetObjectRequestAcceptEncodingFromString(s string) (GetObjectRequestAcceptEncoding, error) {
	switch s {
	case "identity":
		return GetObjectRequestAcceptEncodingIdentity, nil
	case "zstd":
		return GetObjectRequestAcceptEncodingZstd, nil
	}
	var t GetObjectRequestAcceptEncoding
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (g GetObjectRequestAcceptEncoding) Ptr() *GetObjectRequestAcceptEncoding {
	return &g
}

type UploadObjectRequestDistributionMode string

const (
	UploadObjectRequestDistributionModeForce UploadObjectRequestDistributionMode = "force"
)

func NewUploadObjectRequestDistributionModeFromString(s string) (UploadObjectRequestDistributionMode, error) {
	switch s {
	case "force":
		return UploadObjectRequestDistributionModeForce, nil
	}
	var t UploadObjectRequestDistributionMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (u UploadObjectRequestDistributionMode) Ptr() *UploadObjectRequestDistributionMode {
	return &u
}
