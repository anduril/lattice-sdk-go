// Code generated from our API definition. DO NOT EDIT.

package video

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/anduril/lattice-sdk-go/v4/internal"
	big "math/big"
	time "time"
)

var (
	createEgressStreamRequestFieldIngressID = big.NewInt(1 << 0)
	createEgressStreamRequestFieldRtsp      = big.NewInt(1 << 1)
	createEgressStreamRequestFieldSrt       = big.NewInt(1 << 2)
)

type CreateEgressStreamRequest struct {
	// Identifier of the live ingress stream to re-publish as an egress stream.
	IngressID *string       `json:"ingressId,omitempty" url:"-"`
	Rtsp      *RtspSettings `json:"rtsp,omitempty" url:"-"`
	Srt       *SrtSettings  `json:"srt,omitempty" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateEgressStreamRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetIngressID sets the IngressID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateEgressStreamRequest) SetIngressID(ingressID *string) {
	c.IngressID = ingressID
	c.require(createEgressStreamRequestFieldIngressID)
}

// SetRtsp sets the Rtsp field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateEgressStreamRequest) SetRtsp(rtsp *RtspSettings) {
	c.Rtsp = rtsp
	c.require(createEgressStreamRequestFieldRtsp)
}

// SetSrt sets the Srt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateEgressStreamRequest) SetSrt(srt *SrtSettings) {
	c.Srt = srt
	c.require(createEgressStreamRequestFieldSrt)
}

func (c *CreateEgressStreamRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateEgressStreamRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*c = CreateEgressStreamRequest(body)
	return nil
}

func (c *CreateEgressStreamRequest) MarshalJSON() ([]byte, error) {
	type embed CreateEgressStreamRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	createIngressStreamRequestFieldIngressID = big.NewInt(1 << 0)
	createIngressStreamRequestFieldTitle     = big.NewInt(1 << 1)
	createIngressStreamRequestFieldMpegTs    = big.NewInt(1 << 2)
	createIngressStreamRequestFieldRtsp      = big.NewInt(1 << 3)
	createIngressStreamRequestFieldSrt       = big.NewInt(1 << 4)
)

type CreateIngressStreamRequest struct {
	// Caller-supplied identifier for the new stream. If omitted, the service generates a GUID.
	//
	//	If supplied, a consistent and recognizable pattern is recommended. A common convention
	//	is a group prefix (organization, platform, or asset) followed by a specific identifier
	//	using underscore or dot as a separator, for example, `drone_1`, `vessel_2`, or
	//	`teamalpha.drone1`.
	//
	//	When supplied, an ingress_id must be between 4 and 36 characters long and use only
	//	printable ASCII characters with no spaces; the 36-character ceiling leaves room for a
	//	full GUID. A value outside that length range, or one containing spaces, control
	//	characters, or non-ASCII characters, is rejected, as is an ingress_id that another
	//	ingress stream is already using.
	IngressID *string `json:"ingressId,omitempty" url:"-"`
	// Human-readable title for the stream. A title is required: surrounding whitespace is
	//
	//	trimmed before it is stored, and what remains must be non-empty, valid UTF-8, and no
	//	longer than 64 characters. Otherwise the request is rejected.
	Title *string `json:"title,omitempty" url:"-"`
	// Receive an MPEG-TS push from the producer. The service allocates a UDP port and
	//
	//	returns the URL the producer must push to in CreateIngressStreamResponse.
	//
	//	MPEG-TS ingress may be disabled per deployment. When it is disabled, a request
	//	that selects mpeg_ts is rejected with a gRPC error rather than accepted, so
	//	callers should be prepared to fall back to another protocol.
	MpegTs *MpegTsSettings `json:"mpegTs,omitempty" url:"-"`
	// Pull from a caller-supplied RTSP URL.
	Rtsp *RtspSettings `json:"rtsp,omitempty" url:"-"`
	// Receive an SRT push from the producer. The service returns a URL and session_id
	//
	//	in CreateIngressStreamResponse.
	Srt *SrtSettings `json:"srt,omitempty" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateIngressStreamRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetIngressID sets the IngressID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateIngressStreamRequest) SetIngressID(ingressID *string) {
	c.IngressID = ingressID
	c.require(createIngressStreamRequestFieldIngressID)
}

// SetTitle sets the Title field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateIngressStreamRequest) SetTitle(title *string) {
	c.Title = title
	c.require(createIngressStreamRequestFieldTitle)
}

// SetMpegTs sets the MpegTs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateIngressStreamRequest) SetMpegTs(mpegTs *MpegTsSettings) {
	c.MpegTs = mpegTs
	c.require(createIngressStreamRequestFieldMpegTs)
}

// SetRtsp sets the Rtsp field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateIngressStreamRequest) SetRtsp(rtsp *RtspSettings) {
	c.Rtsp = rtsp
	c.require(createIngressStreamRequestFieldRtsp)
}

// SetSrt sets the Srt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateIngressStreamRequest) SetSrt(srt *SrtSettings) {
	c.Srt = srt
	c.require(createIngressStreamRequestFieldSrt)
}

func (c *CreateIngressStreamRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateIngressStreamRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*c = CreateIngressStreamRequest(body)
	return nil
}

func (c *CreateIngressStreamRequest) MarshalJSON() ([]byte, error) {
	type embed CreateIngressStreamRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	deleteEgressStreamRequestFieldEgressID = big.NewInt(1 << 0)
)

type DeleteEgressStreamRequest struct {
	// Identifier of the egress stream to delete.
	EgressID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteEgressStreamRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetEgressID sets the EgressID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteEgressStreamRequest) SetEgressID(egressID string) {
	d.EgressID = egressID
	d.require(deleteEgressStreamRequestFieldEgressID)
}

var (
	deleteIngressStreamRequestFieldIngressID = big.NewInt(1 << 0)
)

type DeleteIngressStreamRequest struct {
	// Identifier of the ingress stream to delete.
	IngressID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteIngressStreamRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetIngressID sets the IngressID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteIngressStreamRequest) SetIngressID(ingressID string) {
	d.IngressID = ingressID
	d.require(deleteIngressStreamRequestFieldIngressID)
}

var (
	getEgressStreamRequestFieldEgressID = big.NewInt(1 << 0)
)

type GetEgressStreamRequest struct {
	// Identifier of the egress stream to retrieve.
	EgressID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetEgressStreamRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetEgressID sets the EgressID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetEgressStreamRequest) SetEgressID(egressID string) {
	g.EgressID = egressID
	g.require(getEgressStreamRequestFieldEgressID)
}

var (
	getIngressStreamRequestFieldIngressID = big.NewInt(1 << 0)
)

type GetIngressStreamRequest struct {
	// Identifier of the ingress stream to retrieve.
	IngressID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetIngressStreamRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetIngressID sets the IngressID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetIngressStreamRequest) SetIngressID(ingressID string) {
	g.IngressID = ingressID
	g.require(getIngressStreamRequestFieldIngressID)
}

var (
	listEgressStreamsRequestFieldPageSize  = big.NewInt(1 << 0)
	listEgressStreamsRequestFieldPageToken = big.NewInt(1 << 1)
)

type ListEgressStreamsRequest struct {
	// Desired number of egress streams per page. Defaults to 50 if left blank,
	//
	//	and capped at 100. The response may contain fewer than max page size.
	PageSize *int `json:"-" url:"pageSize,omitempty"`
	// To retrieve the next page, pass the `next_page_token` from the previous
	//
	//	response. Leave empty for the first page.
	//
	//	Keep the rest of the request identical between pages, otherwise the
	//	server may reject it.
	PageToken *string `json:"-" url:"pageToken,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListEgressStreamsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListEgressStreamsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listEgressStreamsRequestFieldPageSize)
}

// SetPageToken sets the PageToken field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListEgressStreamsRequest) SetPageToken(pageToken *string) {
	l.PageToken = pageToken
	l.require(listEgressStreamsRequestFieldPageToken)
}

var (
	listIngressStreamsRequestFieldPageSize  = big.NewInt(1 << 0)
	listIngressStreamsRequestFieldPageToken = big.NewInt(1 << 1)
)

type ListIngressStreamsRequest struct {
	// Desired number of ingress streams per page. Defaults to 50 if left blank,
	//
	//	and capped at 100. The response may contain fewer than requested.
	PageSize *int `json:"-" url:"pageSize,omitempty"`
	// To retrieve the next page, pass the `next_page_token` from the previous
	//
	//	response. Leave empty for the first page.
	//
	//	Keep the rest of the request identical between pages, otherwise the
	//	server may reject it.
	PageToken *string `json:"-" url:"pageToken,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListIngressStreamsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListIngressStreamsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listIngressStreamsRequestFieldPageSize)
}

// SetPageToken sets the PageToken field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListIngressStreamsRequest) SetPageToken(pageToken *string) {
	l.PageToken = pageToken
	l.require(listIngressStreamsRequestFieldPageToken)
}

var (
	createEgressStreamResponseFieldEgressID = big.NewInt(1 << 0)
	createEgressStreamResponseFieldRtsp     = big.NewInt(1 << 1)
	createEgressStreamResponseFieldSrt      = big.NewInt(1 << 2)
)

type CreateEgressStreamResponse struct {
	// Service-generated identifier for the new egress stream. Use it for subsequent
	//
	//	`GetEgressStream` and `DeleteEgressStream` calls.
	EgressID *string     `json:"egressId,omitempty" url:"egressId,omitempty"`
	Rtsp     *RtspEgress `json:"rtsp,omitempty" url:"rtsp,omitempty"`
	Srt      *SrtEgress  `json:"srt,omitempty" url:"srt,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateEgressStreamResponse) GetEgressID() *string {
	if c == nil {
		return nil
	}
	return c.EgressID
}

func (c *CreateEgressStreamResponse) GetRtsp() *RtspEgress {
	if c == nil {
		return nil
	}
	return c.Rtsp
}

func (c *CreateEgressStreamResponse) GetSrt() *SrtEgress {
	if c == nil {
		return nil
	}
	return c.Srt
}

func (c *CreateEgressStreamResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateEgressStreamResponse) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetEgressID sets the EgressID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateEgressStreamResponse) SetEgressID(egressID *string) {
	c.EgressID = egressID
	c.require(createEgressStreamResponseFieldEgressID)
}

// SetRtsp sets the Rtsp field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateEgressStreamResponse) SetRtsp(rtsp *RtspEgress) {
	c.Rtsp = rtsp
	c.require(createEgressStreamResponseFieldRtsp)
}

// SetSrt sets the Srt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateEgressStreamResponse) SetSrt(srt *SrtEgress) {
	c.Srt = srt
	c.require(createEgressStreamResponseFieldSrt)
}

func (c *CreateEgressStreamResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateEgressStreamResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateEgressStreamResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateEgressStreamResponse) MarshalJSON() ([]byte, error) {
	type embed CreateEgressStreamResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateEgressStreamResponse) String() string {
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
	createIngressStreamResponseFieldIngressID = big.NewInt(1 << 0)
	createIngressStreamResponseFieldMpegTs    = big.NewInt(1 << 1)
	createIngressStreamResponseFieldSrt       = big.NewInt(1 << 2)
)

type CreateIngressStreamResponse struct {
	// Identifier of the newly created ingress stream. Echoes the caller-supplied
	//
	//	`ingress_id` if one was provided, otherwise a service-generated GUID.
	IngressID *string `json:"ingressId,omitempty" url:"ingressId,omitempty"`
	// Connection details for an MPEG-TS push. Only returned when MPEG-TS ingress is
	//
	//	enabled for the deployment and the request selected mpeg_ts.
	MpegTs *MpegTsIngress `json:"mpegTs,omitempty" url:"mpegTs,omitempty"`
	Srt    *SrtIngress    `json:"srt,omitempty" url:"srt,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateIngressStreamResponse) GetIngressID() *string {
	if c == nil {
		return nil
	}
	return c.IngressID
}

func (c *CreateIngressStreamResponse) GetMpegTs() *MpegTsIngress {
	if c == nil {
		return nil
	}
	return c.MpegTs
}

func (c *CreateIngressStreamResponse) GetSrt() *SrtIngress {
	if c == nil {
		return nil
	}
	return c.Srt
}

func (c *CreateIngressStreamResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateIngressStreamResponse) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetIngressID sets the IngressID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateIngressStreamResponse) SetIngressID(ingressID *string) {
	c.IngressID = ingressID
	c.require(createIngressStreamResponseFieldIngressID)
}

// SetMpegTs sets the MpegTs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateIngressStreamResponse) SetMpegTs(mpegTs *MpegTsIngress) {
	c.MpegTs = mpegTs
	c.require(createIngressStreamResponseFieldMpegTs)
}

// SetSrt sets the Srt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateIngressStreamResponse) SetSrt(srt *SrtIngress) {
	c.Srt = srt
	c.require(createIngressStreamResponseFieldSrt)
}

func (c *CreateIngressStreamResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateIngressStreamResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateIngressStreamResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateIngressStreamResponse) MarshalJSON() ([]byte, error) {
	type embed CreateIngressStreamResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateIngressStreamResponse) String() string {
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

type DeleteEgressStreamResponse struct {

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DeleteEgressStreamResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d *DeleteEgressStreamResponse) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

func (d *DeleteEgressStreamResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteEgressStreamResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteEgressStreamResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeleteEgressStreamResponse) MarshalJSON() ([]byte, error) {
	type embed DeleteEgressStreamResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, d.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (d *DeleteEgressStreamResponse) String() string {
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

type DeleteIngressStreamResponse struct {

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DeleteIngressStreamResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d *DeleteIngressStreamResponse) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

func (d *DeleteIngressStreamResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteIngressStreamResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteIngressStreamResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeleteIngressStreamResponse) MarshalJSON() ([]byte, error) {
	type embed DeleteIngressStreamResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, d.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (d *DeleteIngressStreamResponse) String() string {
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

// An egress stream publishes a single stream to a downstream consumer over a chosen
//
//	transport.
var (
	egressStreamFieldEgressID  = big.NewInt(1 << 0)
	egressStreamFieldIngressID = big.NewInt(1 << 1)
	egressStreamFieldRtsp      = big.NewInt(1 << 2)
	egressStreamFieldSrt       = big.NewInt(1 << 3)
)

type EgressStream struct {
	// Service-generated identifier for the egress stream.
	EgressID *string `json:"egressId,omitempty" url:"egressId,omitempty"`
	// Identifier of the ingress stream this egress stream publishes.
	IngressID *string     `json:"ingressId,omitempty" url:"ingressId,omitempty"`
	Rtsp      *RtspEgress `json:"rtsp,omitempty" url:"rtsp,omitempty"`
	Srt       *SrtEgress  `json:"srt,omitempty" url:"srt,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *EgressStream) GetEgressID() *string {
	if e == nil {
		return nil
	}
	return e.EgressID
}

func (e *EgressStream) GetIngressID() *string {
	if e == nil {
		return nil
	}
	return e.IngressID
}

func (e *EgressStream) GetRtsp() *RtspEgress {
	if e == nil {
		return nil
	}
	return e.Rtsp
}

func (e *EgressStream) GetSrt() *SrtEgress {
	if e == nil {
		return nil
	}
	return e.Srt
}

func (e *EgressStream) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e *EgressStream) require(field *big.Int) {
	if e.explicitFields == nil {
		e.explicitFields = big.NewInt(0)
	}
	e.explicitFields.Or(e.explicitFields, field)
}

// SetEgressID sets the EgressID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EgressStream) SetEgressID(egressID *string) {
	e.EgressID = egressID
	e.require(egressStreamFieldEgressID)
}

// SetIngressID sets the IngressID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EgressStream) SetIngressID(ingressID *string) {
	e.IngressID = ingressID
	e.require(egressStreamFieldIngressID)
}

// SetRtsp sets the Rtsp field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EgressStream) SetRtsp(rtsp *RtspEgress) {
	e.Rtsp = rtsp
	e.require(egressStreamFieldRtsp)
}

// SetSrt sets the Srt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EgressStream) SetSrt(srt *SrtEgress) {
	e.Srt = srt
	e.require(egressStreamFieldSrt)
}

func (e *EgressStream) UnmarshalJSON(data []byte) error {
	type unmarshaler EgressStream
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EgressStream(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *EgressStream) MarshalJSON() ([]byte, error) {
	type embed EgressStream
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, e.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (e *EgressStream) String() string {
	if e == nil {
		return "<nil>"
	}
	if len(e.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

var (
	getEgressStreamResponseFieldEgressStream = big.NewInt(1 << 0)
)

type GetEgressStreamResponse struct {
	// The egress stream corresponding to the requested `egress_id`.
	EgressStream *EgressStream `json:"egressStream,omitempty" url:"egressStream,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetEgressStreamResponse) GetEgressStream() *EgressStream {
	if g == nil {
		return nil
	}
	return g.EgressStream
}

func (g *GetEgressStreamResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetEgressStreamResponse) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetEgressStream sets the EgressStream field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetEgressStreamResponse) SetEgressStream(egressStream *EgressStream) {
	g.EgressStream = egressStream
	g.require(getEgressStreamResponseFieldEgressStream)
}

func (g *GetEgressStreamResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetEgressStreamResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetEgressStreamResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetEgressStreamResponse) MarshalJSON() ([]byte, error) {
	type embed GetEgressStreamResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetEgressStreamResponse) String() string {
	if g == nil {
		return "<nil>"
	}
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

var (
	getIngressStreamResponseFieldIngressStream = big.NewInt(1 << 0)
)

type GetIngressStreamResponse struct {
	// The ingress stream corresponding to the requested `ingress_id`.
	IngressStream *IngressStream `json:"ingressStream,omitempty" url:"ingressStream,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetIngressStreamResponse) GetIngressStream() *IngressStream {
	if g == nil {
		return nil
	}
	return g.IngressStream
}

func (g *GetIngressStreamResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetIngressStreamResponse) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetIngressStream sets the IngressStream field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetIngressStreamResponse) SetIngressStream(ingressStream *IngressStream) {
	g.IngressStream = ingressStream
	g.require(getIngressStreamResponseFieldIngressStream)
}

func (g *GetIngressStreamResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetIngressStreamResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetIngressStreamResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetIngressStreamResponse) MarshalJSON() ([]byte, error) {
	type embed GetIngressStreamResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetIngressStreamResponse) String() string {
	if g == nil {
		return "<nil>"
	}
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

// Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
var (
	googleProtobufAnyFieldType = big.NewInt(1 << 0)
)

type GoogleProtobufAny struct {
	// The type of the serialized message.
	Type *string `json:"@type,omitempty" url:"@type,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	ExtraProperties map[string]interface{} `json:"-" url:"-"`

	rawJSON json.RawMessage
}

func (g *GoogleProtobufAny) GetType() *string {
	if g == nil {
		return nil
	}
	return g.Type
}

func (g *GoogleProtobufAny) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.ExtraProperties
}

func (g *GoogleProtobufAny) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetType sets the Type field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GoogleProtobufAny) SetType(type_ *string) {
	g.Type = type_
	g.require(googleProtobufAnyFieldType)
}

func (g *GoogleProtobufAny) UnmarshalJSON(data []byte) error {
	type embed GoogleProtobufAny
	var unmarshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*g = GoogleProtobufAny(unmarshaler.embed)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.ExtraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GoogleProtobufAny) MarshalJSON() ([]byte, error) {
	type embed GoogleProtobufAny
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return internal.MarshalJSONWithExtraProperties(explicitMarshaler, g.ExtraProperties)
}

func (g *GoogleProtobufAny) String() string {
	if g == nil {
		return "<nil>"
	}
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

// The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
var (
	googleRPCStatusFieldCode    = big.NewInt(1 << 0)
	googleRPCStatusFieldMessage = big.NewInt(1 << 1)
	googleRPCStatusFieldDetails = big.NewInt(1 << 2)
)

type GoogleRPCStatus struct {
	// The status code, which should be an enum value of [google.rpc.Code][google.rpc.Code].
	Code *int `json:"code,omitempty" url:"code,omitempty"`
	// A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
	Message *string `json:"message,omitempty" url:"message,omitempty"`
	// A list of messages that carry the error details.  There is a common set of message types for APIs to use.
	Details []*GoogleProtobufAny `json:"details,omitempty" url:"details,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GoogleRPCStatus) GetCode() *int {
	if g == nil {
		return nil
	}
	return g.Code
}

func (g *GoogleRPCStatus) GetMessage() *string {
	if g == nil {
		return nil
	}
	return g.Message
}

func (g *GoogleRPCStatus) GetDetails() []*GoogleProtobufAny {
	if g == nil {
		return nil
	}
	return g.Details
}

func (g *GoogleRPCStatus) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GoogleRPCStatus) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetCode sets the Code field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GoogleRPCStatus) SetCode(code *int) {
	g.Code = code
	g.require(googleRPCStatusFieldCode)
}

// SetMessage sets the Message field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GoogleRPCStatus) SetMessage(message *string) {
	g.Message = message
	g.require(googleRPCStatusFieldMessage)
}

// SetDetails sets the Details field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GoogleRPCStatus) SetDetails(details []*GoogleProtobufAny) {
	g.Details = details
	g.require(googleRPCStatusFieldDetails)
}

func (g *GoogleRPCStatus) UnmarshalJSON(data []byte) error {
	type unmarshaler GoogleRPCStatus
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GoogleRPCStatus(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GoogleRPCStatus) MarshalJSON() ([]byte, error) {
	type embed GoogleRPCStatus
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GoogleRPCStatus) String() string {
	if g == nil {
		return "<nil>"
	}
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

// An ingress stream represents a single source feeding frames into Lattice.
//
//	Ingress streams are replicated across Lattice and visible anywhere in the deployment.
var (
	ingressStreamFieldIngressID = big.NewInt(1 << 0)
	ingressStreamFieldTitle     = big.NewInt(1 << 1)
	ingressStreamFieldStatus    = big.NewInt(1 << 2)
	ingressStreamFieldMpegTs    = big.NewInt(1 << 3)
	ingressStreamFieldRtsp      = big.NewInt(1 << 4)
	ingressStreamFieldSrt       = big.NewInt(1 << 5)
	ingressStreamFieldCreatedAt = big.NewInt(1 << 6)
	ingressStreamFieldUpdatedAt = big.NewInt(1 << 7)
	ingressStreamFieldEgressIDs = big.NewInt(1 << 8)
)

type IngressStream struct {
	// Unique identifier for the ingress stream.
	IngressID *string `json:"ingressId,omitempty" url:"ingressId,omitempty"`
	// Human-readable title supplied at creation time.
	Title *string `json:"title,omitempty" url:"title,omitempty"`
	// Current lifecycle status of the stream. See StreamStatus for the full state machine.
	Status *IngressStreamStatus `json:"status,omitempty" url:"status,omitempty"`
	MpegTs *MpegTsIngress       `json:"mpegTs,omitempty" url:"mpegTs,omitempty"`
	Rtsp   *RtspIngress         `json:"rtsp,omitempty" url:"rtsp,omitempty"`
	Srt    *SrtIngress          `json:"srt,omitempty" url:"srt,omitempty"`
	// Wall-clock time the stream was created.
	CreatedAt *time.Time `json:"createdAt,omitempty" url:"createdAt,omitempty"`
	// Wall-clock time the stream's status (STREAM_STATUS) was changed. The status can change based on the activity or
	//
	//	the deletion of the stream.
	UpdatedAt *time.Time `json:"updatedAt,omitempty" url:"updatedAt,omitempty"`
	// Identifiers of the egress streams currently consuming this ingress stream.
	EgressIDs []string `json:"egressIds,omitempty" url:"egressIds,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (i *IngressStream) GetIngressID() *string {
	if i == nil {
		return nil
	}
	return i.IngressID
}

func (i *IngressStream) GetTitle() *string {
	if i == nil {
		return nil
	}
	return i.Title
}

func (i *IngressStream) GetStatus() *IngressStreamStatus {
	if i == nil {
		return nil
	}
	return i.Status
}

func (i *IngressStream) GetMpegTs() *MpegTsIngress {
	if i == nil {
		return nil
	}
	return i.MpegTs
}

func (i *IngressStream) GetRtsp() *RtspIngress {
	if i == nil {
		return nil
	}
	return i.Rtsp
}

func (i *IngressStream) GetSrt() *SrtIngress {
	if i == nil {
		return nil
	}
	return i.Srt
}

func (i *IngressStream) GetCreatedAt() *time.Time {
	if i == nil {
		return nil
	}
	return i.CreatedAt
}

func (i *IngressStream) GetUpdatedAt() *time.Time {
	if i == nil {
		return nil
	}
	return i.UpdatedAt
}

func (i *IngressStream) GetEgressIDs() []string {
	if i == nil {
		return nil
	}
	return i.EgressIDs
}

func (i *IngressStream) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i *IngressStream) require(field *big.Int) {
	if i.explicitFields == nil {
		i.explicitFields = big.NewInt(0)
	}
	i.explicitFields.Or(i.explicitFields, field)
}

// SetIngressID sets the IngressID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IngressStream) SetIngressID(ingressID *string) {
	i.IngressID = ingressID
	i.require(ingressStreamFieldIngressID)
}

// SetTitle sets the Title field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IngressStream) SetTitle(title *string) {
	i.Title = title
	i.require(ingressStreamFieldTitle)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IngressStream) SetStatus(status *IngressStreamStatus) {
	i.Status = status
	i.require(ingressStreamFieldStatus)
}

// SetMpegTs sets the MpegTs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IngressStream) SetMpegTs(mpegTs *MpegTsIngress) {
	i.MpegTs = mpegTs
	i.require(ingressStreamFieldMpegTs)
}

// SetRtsp sets the Rtsp field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IngressStream) SetRtsp(rtsp *RtspIngress) {
	i.Rtsp = rtsp
	i.require(ingressStreamFieldRtsp)
}

// SetSrt sets the Srt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IngressStream) SetSrt(srt *SrtIngress) {
	i.Srt = srt
	i.require(ingressStreamFieldSrt)
}

// SetCreatedAt sets the CreatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IngressStream) SetCreatedAt(createdAt *time.Time) {
	i.CreatedAt = createdAt
	i.require(ingressStreamFieldCreatedAt)
}

// SetUpdatedAt sets the UpdatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IngressStream) SetUpdatedAt(updatedAt *time.Time) {
	i.UpdatedAt = updatedAt
	i.require(ingressStreamFieldUpdatedAt)
}

// SetEgressIDs sets the EgressIDs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IngressStream) SetEgressIDs(egressIDs []string) {
	i.EgressIDs = egressIDs
	i.require(ingressStreamFieldEgressIDs)
}

func (i *IngressStream) UnmarshalJSON(data []byte) error {
	type embed IngressStream
	var unmarshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"createdAt,omitempty"`
		UpdatedAt *internal.DateTime `json:"updatedAt,omitempty"`
	}{
		embed: embed(*i),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*i = IngressStream(unmarshaler.embed)
	i.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	i.UpdatedAt = unmarshaler.UpdatedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	i.rawJSON = json.RawMessage(data)
	return nil
}

func (i *IngressStream) MarshalJSON() ([]byte, error) {
	type embed IngressStream
	var marshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"createdAt,omitempty"`
		UpdatedAt *internal.DateTime `json:"updatedAt,omitempty"`
	}{
		embed:     embed(*i),
		CreatedAt: internal.NewOptionalDateTime(i.CreatedAt),
		UpdatedAt: internal.NewOptionalDateTime(i.UpdatedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, i.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (i *IngressStream) String() string {
	if i == nil {
		return "<nil>"
	}
	if len(i.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(i.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}

// Current lifecycle status of the stream. See StreamStatus for the full state machine.
type IngressStreamStatus string

const (
	IngressStreamStatusStreamStatusUnspecified IngressStreamStatus = "STREAM_STATUS_UNSPECIFIED"
	IngressStreamStatusStreamStatusLive        IngressStreamStatus = "STREAM_STATUS_LIVE"
	IngressStreamStatusStreamStatusInactive    IngressStreamStatus = "STREAM_STATUS_INACTIVE"
	IngressStreamStatusStreamStatusUnavailable IngressStreamStatus = "STREAM_STATUS_UNAVAILABLE"
	IngressStreamStatusStreamStatusArchived    IngressStreamStatus = "STREAM_STATUS_ARCHIVED"
)

func NewIngressStreamStatusFromString(s string) (IngressStreamStatus, error) {
	switch s {
	case "STREAM_STATUS_UNSPECIFIED":
		return IngressStreamStatusStreamStatusUnspecified, nil
	case "STREAM_STATUS_LIVE":
		return IngressStreamStatusStreamStatusLive, nil
	case "STREAM_STATUS_INACTIVE":
		return IngressStreamStatusStreamStatusInactive, nil
	case "STREAM_STATUS_UNAVAILABLE":
		return IngressStreamStatusStreamStatusUnavailable, nil
	case "STREAM_STATUS_ARCHIVED":
		return IngressStreamStatusStreamStatusArchived, nil
	}
	var t IngressStreamStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (i IngressStreamStatus) Ptr() *IngressStreamStatus {
	return &i
}

var (
	listEgressStreamsResponseFieldEgressStreams = big.NewInt(1 << 0)
	listEgressStreamsResponseFieldNextPageToken = big.NewInt(1 << 1)
)

type ListEgressStreamsResponse struct {
	// The egress streams on this page. Up to `page_size` entries
	//
	//	(defaults to 50, capped at 100). Ordered by egress stream create time.
	EgressStreams []*EgressStream `json:"egressStreams,omitempty" url:"egressStreams,omitempty"`
	// Pass this back as `page_token` to retrieve the next page.
	//
	//	Empty when there are no more pages.
	NextPageToken *string `json:"nextPageToken,omitempty" url:"nextPageToken,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListEgressStreamsResponse) GetEgressStreams() []*EgressStream {
	if l == nil {
		return nil
	}
	return l.EgressStreams
}

func (l *ListEgressStreamsResponse) GetNextPageToken() *string {
	if l == nil {
		return nil
	}
	return l.NextPageToken
}

func (l *ListEgressStreamsResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListEgressStreamsResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetEgressStreams sets the EgressStreams field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListEgressStreamsResponse) SetEgressStreams(egressStreams []*EgressStream) {
	l.EgressStreams = egressStreams
	l.require(listEgressStreamsResponseFieldEgressStreams)
}

// SetNextPageToken sets the NextPageToken field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListEgressStreamsResponse) SetNextPageToken(nextPageToken *string) {
	l.NextPageToken = nextPageToken
	l.require(listEgressStreamsResponseFieldNextPageToken)
}

func (l *ListEgressStreamsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListEgressStreamsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListEgressStreamsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListEgressStreamsResponse) MarshalJSON() ([]byte, error) {
	type embed ListEgressStreamsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListEgressStreamsResponse) String() string {
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
	listIngressStreamsResponseFieldIngressStreams = big.NewInt(1 << 0)
	listIngressStreamsResponseFieldNextPageToken  = big.NewInt(1 << 1)
)

type ListIngressStreamsResponse struct {
	// The ingress streams on this page. Up to `page_size` entries
	//
	//	(defaults to 50, capped at 100). Ordered by ingress stream create time.
	IngressStreams []*IngressStream `json:"ingressStreams,omitempty" url:"ingressStreams,omitempty"`
	// Pass this back as `page_token` to retrieve the next page.
	//
	//	Empty when there are no more pages.
	NextPageToken *string `json:"nextPageToken,omitempty" url:"nextPageToken,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListIngressStreamsResponse) GetIngressStreams() []*IngressStream {
	if l == nil {
		return nil
	}
	return l.IngressStreams
}

func (l *ListIngressStreamsResponse) GetNextPageToken() *string {
	if l == nil {
		return nil
	}
	return l.NextPageToken
}

func (l *ListIngressStreamsResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListIngressStreamsResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetIngressStreams sets the IngressStreams field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListIngressStreamsResponse) SetIngressStreams(ingressStreams []*IngressStream) {
	l.IngressStreams = ingressStreams
	l.require(listIngressStreamsResponseFieldIngressStreams)
}

// SetNextPageToken sets the NextPageToken field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListIngressStreamsResponse) SetNextPageToken(nextPageToken *string) {
	l.NextPageToken = nextPageToken
	l.require(listIngressStreamsResponseFieldNextPageToken)
}

func (l *ListIngressStreamsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListIngressStreamsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListIngressStreamsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListIngressStreamsResponse) MarshalJSON() ([]byte, error) {
	type embed ListIngressStreamsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListIngressStreamsResponse) String() string {
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

// MPEG-TS ingress connection details.
var (
	mpegTsIngressFieldURL = big.NewInt(1 << 0)
)

type MpegTsIngress struct {
	// The URL that the producer should push the MPEG-TS stream to.
	URL *string `json:"url,omitempty" url:"url,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (m *MpegTsIngress) GetURL() *string {
	if m == nil {
		return nil
	}
	return m.URL
}

func (m *MpegTsIngress) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m *MpegTsIngress) require(field *big.Int) {
	if m.explicitFields == nil {
		m.explicitFields = big.NewInt(0)
	}
	m.explicitFields.Or(m.explicitFields, field)
}

// SetURL sets the URL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (m *MpegTsIngress) SetURL(url *string) {
	m.URL = url
	m.require(mpegTsIngressFieldURL)
}

func (m *MpegTsIngress) UnmarshalJSON(data []byte) error {
	type unmarshaler MpegTsIngress
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = MpegTsIngress(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	m.rawJSON = json.RawMessage(data)
	return nil
}

func (m *MpegTsIngress) MarshalJSON() ([]byte, error) {
	type embed MpegTsIngress
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, m.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (m *MpegTsIngress) String() string {
	if m == nil {
		return "<nil>"
	}
	if len(m.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(m.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

// Settings for MPEG-TS ingress. Empty by default, the service allocates a UDP port
//
//	from a service-wide pool and returns the push URL in CreateIngressStreamResponse.
//
//	MPEG-TS ingress may be disabled per deployment. When it is disabled, a
//	CreateIngressStream request that selects mpeg_ts is rejected with a gRPC error.
type MpegTsSettings = map[string]any

// RTSP egress connection details.
var (
	rtspEgressFieldURL = big.NewInt(1 << 0)
)

type RtspEgress struct {
	// The RTSP URL the downstream consumer should pull from.
	URL *string `json:"url,omitempty" url:"url,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *RtspEgress) GetURL() *string {
	if r == nil {
		return nil
	}
	return r.URL
}

func (r *RtspEgress) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *RtspEgress) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetURL sets the URL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *RtspEgress) SetURL(url *string) {
	r.URL = url
	r.require(rtspEgressFieldURL)
}

func (r *RtspEgress) UnmarshalJSON(data []byte) error {
	type unmarshaler RtspEgress
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RtspEgress(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RtspEgress) MarshalJSON() ([]byte, error) {
	type embed RtspEgress
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *RtspEgress) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

// RTSP ingress connection details.
var (
	rtspIngressFieldURL = big.NewInt(1 << 0)
)

type RtspIngress struct {
	// The upstream RTSP URL. Lattice will pull from the supplied URL.
	//
	//	The URL must be prefixed with `rtsp://`.
	URL *string `json:"url,omitempty" url:"url,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *RtspIngress) GetURL() *string {
	if r == nil {
		return nil
	}
	return r.URL
}

func (r *RtspIngress) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *RtspIngress) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetURL sets the URL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *RtspIngress) SetURL(url *string) {
	r.URL = url
	r.require(rtspIngressFieldURL)
}

func (r *RtspIngress) UnmarshalJSON(data []byte) error {
	type unmarshaler RtspIngress
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RtspIngress(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RtspIngress) MarshalJSON() ([]byte, error) {
	type embed RtspIngress
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *RtspIngress) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

// Settings for RTSP.
var (
	rtspSettingsFieldURL = big.NewInt(1 << 0)
)

type RtspSettings struct {
	// The upstream RTSP URL the service should pull frames from. Must use
	//
	//	the `rtsp://` scheme.
	URL *string `json:"url,omitempty" url:"url,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *RtspSettings) GetURL() *string {
	if r == nil {
		return nil
	}
	return r.URL
}

func (r *RtspSettings) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *RtspSettings) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetURL sets the URL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *RtspSettings) SetURL(url *string) {
	r.URL = url
	r.require(rtspSettingsFieldURL)
}

func (r *RtspSettings) UnmarshalJSON(data []byte) error {
	type unmarshaler RtspSettings
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RtspSettings(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RtspSettings) MarshalJSON() ([]byte, error) {
	type embed RtspSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *RtspSettings) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

// SRT egress connection details.
var (
	srtEgressFieldURL       = big.NewInt(1 << 0)
	srtEgressFieldSessionID = big.NewInt(1 << 1)
)

type SrtEgress struct {
	// The URL on which Lattice listens. The downstream consumer pulls from this URL.
	URL *string `json:"url,omitempty" url:"url,omitempty"`
	// Unique session identifier the consumer must supply on the SRT connection.
	SessionID *string `json:"sessionId,omitempty" url:"sessionId,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SrtEgress) GetURL() *string {
	if s == nil {
		return nil
	}
	return s.URL
}

func (s *SrtEgress) GetSessionID() *string {
	if s == nil {
		return nil
	}
	return s.SessionID
}

func (s *SrtEgress) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *SrtEgress) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetURL sets the URL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SrtEgress) SetURL(url *string) {
	s.URL = url
	s.require(srtEgressFieldURL)
}

// SetSessionID sets the SessionID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SrtEgress) SetSessionID(sessionID *string) {
	s.SessionID = sessionID
	s.require(srtEgressFieldSessionID)
}

func (s *SrtEgress) UnmarshalJSON(data []byte) error {
	type unmarshaler SrtEgress
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SrtEgress(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SrtEgress) MarshalJSON() ([]byte, error) {
	type embed SrtEgress
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *SrtEgress) String() string {
	if s == nil {
		return "<nil>"
	}
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// SRT ingress connection details. Returned to the producer so it knows where to
//
//	push the stream.
var (
	srtIngressFieldURL       = big.NewInt(1 << 0)
	srtIngressFieldSessionID = big.NewInt(1 << 1)
)

type SrtIngress struct {
	// The URL the producer should push the SRT stream to.
	URL *string `json:"url,omitempty" url:"url,omitempty"`
	// Unique session identifier the producer must include on the SRT connection. See
	//
	//	SrtSettings for context.
	SessionID *string `json:"sessionId,omitempty" url:"sessionId,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SrtIngress) GetURL() *string {
	if s == nil {
		return nil
	}
	return s.URL
}

func (s *SrtIngress) GetSessionID() *string {
	if s == nil {
		return nil
	}
	return s.SessionID
}

func (s *SrtIngress) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *SrtIngress) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetURL sets the URL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SrtIngress) SetURL(url *string) {
	s.URL = url
	s.require(srtIngressFieldURL)
}

// SetSessionID sets the SessionID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SrtIngress) SetSessionID(sessionID *string) {
	s.SessionID = sessionID
	s.require(srtIngressFieldSessionID)
}

func (s *SrtIngress) UnmarshalJSON(data []byte) error {
	type unmarshaler SrtIngress
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SrtIngress(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SrtIngress) MarshalJSON() ([]byte, error) {
	type embed SrtIngress
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *SrtIngress) String() string {
	if s == nil {
		return "<nil>"
	}
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// The Lattice video service supports SRT protocol for push operations (ingress)
//
//	and pull operations (egress).
//
//	When configuring SRT for ingress, CreateIngressStreamResponse will
//	return to the user a url to push to which contains a unique 'session_id' to use
//	on the connection. If supplied, passphrase will be applied on incoming
//	connections.
//
//	When configuring SRT for egress, CreateEgressStreamResponse will
//	return to the user a url from which to pull a stream. Use the supplied
//	session_id and passphrase in your StreamId if applicable.
//	See the SRT documentation on Access Control for more information.
var (
	srtSettingsFieldPassphrase = big.NewInt(1 << 0)
)

type SrtSettings struct {
	// Optional passphrase for the stream, set by the user, that applies AES encryption.
	Passphrase *string `json:"passphrase,omitempty" url:"passphrase,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SrtSettings) GetPassphrase() *string {
	if s == nil {
		return nil
	}
	return s.Passphrase
}

func (s *SrtSettings) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *SrtSettings) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetPassphrase sets the Passphrase field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SrtSettings) SetPassphrase(passphrase *string) {
	s.Passphrase = passphrase
	s.require(srtSettingsFieldPassphrase)
}

func (s *SrtSettings) UnmarshalJSON(data []byte) error {
	type unmarshaler SrtSettings
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SrtSettings(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SrtSettings) MarshalJSON() ([]byte, error) {
	type embed SrtSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *SrtSettings) String() string {
	if s == nil {
		return "<nil>"
	}
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}
