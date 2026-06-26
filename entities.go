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
	getEntityRequestFieldEntityID = big.NewInt(1 << 0)
)

type GetEntityRequest struct {
	// ID of the entity to return
	EntityID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetEntityRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetEntityID sets the EntityID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetEntityRequest) SetEntityID(entityID string) {
	g.EntityID = entityID
	g.require(getEntityRequestFieldEntityID)
}

var (
	entityEventRequestFieldSessionToken = big.NewInt(1 << 0)
	entityEventRequestFieldBatchSize    = big.NewInt(1 << 1)
)

type EntityEventRequest struct {
	// Long-poll session identifier. Leave empty to start a new polling session.
	SessionToken string `json:"sessionToken" url:"-"`
	// Maximum size of response batch. Defaults to 100. Must be between 1 and 2000 (inclusive).
	BatchSize *int `json:"batchSize,omitempty" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (e *EntityEventRequest) require(field *big.Int) {
	if e.explicitFields == nil {
		e.explicitFields = big.NewInt(0)
	}
	e.explicitFields.Or(e.explicitFields, field)
}

// SetSessionToken sets the SessionToken field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityEventRequest) SetSessionToken(sessionToken string) {
	e.SessionToken = sessionToken
	e.require(entityEventRequestFieldSessionToken)
}

// SetBatchSize sets the BatchSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityEventRequest) SetBatchSize(batchSize *int) {
	e.BatchSize = batchSize
	e.require(entityEventRequestFieldBatchSize)
}

func (e *EntityEventRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler EntityEventRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*e = EntityEventRequest(body)
	return nil
}

func (e *EntityEventRequest) MarshalJSON() ([]byte, error) {
	type embed EntityEventRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, e.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	entityOverrideFieldEntityID   = big.NewInt(1 << 0)
	entityOverrideFieldFieldPath  = big.NewInt(1 << 1)
	entityOverrideFieldEntity     = big.NewInt(1 << 2)
	entityOverrideFieldProvenance = big.NewInt(1 << 3)
)

type EntityOverride struct {
	// The unique ID of the entity to override
	EntityID string `json:"-" url:"-"`
	// fieldPath to override
	FieldPath string `json:"-" url:"-"`
	// The entity containing the overridden fields. The service will extract the overridable fields from
	// the object and ignore all other fields.
	Entity *Entity `json:"entity,omitempty" url:"-"`
	// Additional information about the source of the override.
	Provenance *Provenance `json:"provenance,omitempty" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (e *EntityOverride) require(field *big.Int) {
	if e.explicitFields == nil {
		e.explicitFields = big.NewInt(0)
	}
	e.explicitFields.Or(e.explicitFields, field)
}

// SetEntityID sets the EntityID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityOverride) SetEntityID(entityID string) {
	e.EntityID = entityID
	e.require(entityOverrideFieldEntityID)
}

// SetFieldPath sets the FieldPath field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityOverride) SetFieldPath(fieldPath string) {
	e.FieldPath = fieldPath
	e.require(entityOverrideFieldFieldPath)
}

// SetEntity sets the Entity field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityOverride) SetEntity(entity *Entity) {
	e.Entity = entity
	e.require(entityOverrideFieldEntity)
}

// SetProvenance sets the Provenance field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityOverride) SetProvenance(provenance *Provenance) {
	e.Provenance = provenance
	e.require(entityOverrideFieldProvenance)
}

func (e *EntityOverride) UnmarshalJSON(data []byte) error {
	type unmarshaler EntityOverride
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*e = EntityOverride(body)
	return nil
}

func (e *EntityOverride) MarshalJSON() ([]byte, error) {
	type embed EntityOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, e.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	removeEntityOverrideRequestFieldEntityID  = big.NewInt(1 << 0)
	removeEntityOverrideRequestFieldFieldPath = big.NewInt(1 << 1)
)

type RemoveEntityOverrideRequest struct {
	// The unique ID of the entity to undo an override from.
	EntityID string `json:"-" url:"-"`
	// The fieldPath to clear overrides from.
	FieldPath string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (r *RemoveEntityOverrideRequest) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetEntityID sets the EntityID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *RemoveEntityOverrideRequest) SetEntityID(entityID string) {
	r.EntityID = entityID
	r.require(removeEntityOverrideRequestFieldEntityID)
}

// SetFieldPath sets the FieldPath field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *RemoveEntityOverrideRequest) SetFieldPath(fieldPath string) {
	r.FieldPath = fieldPath
	r.require(removeEntityOverrideRequestFieldFieldPath)
}

var (
	entityStreamRequestFieldHeartbeatIntervalMs = big.NewInt(1 << 0)
	entityStreamRequestFieldPreExistingOnly     = big.NewInt(1 << 1)
	entityStreamRequestFieldComponentsToInclude = big.NewInt(1 << 2)
	entityStreamRequestFieldFilter              = big.NewInt(1 << 3)
)

type EntityStreamRequest struct {
	// at what interval to send heartbeat events, defaults to 30s.
	HeartbeatIntervalMs *int `json:"heartbeatIntervalMS,omitempty" url:"-"`
	// only stream pre-existing entities in the environment and then close the connection, defaults to false.
	PreExistingOnly *bool `json:"preExistingOnly,omitempty" url:"-"`
	// list of components to include, leave empty to include all components.
	ComponentsToInclude []string `json:"componentsToInclude,omitempty" url:"-"`
	// Optional root of a Statement filter expression tree. If provided, only entities matching
	// the filter are streamed. Applied dynamically: an entity that begins matching is delivered
	// as a CREATE, and one that stops matching is delivered as a DELETE. Mirrors the filter on
	// the gRPC StreamEntityComponents endpoint.
	Filter *Statement `json:"filter,omitempty" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (e *EntityStreamRequest) require(field *big.Int) {
	if e.explicitFields == nil {
		e.explicitFields = big.NewInt(0)
	}
	e.explicitFields.Or(e.explicitFields, field)
}

// SetHeartbeatIntervalMs sets the HeartbeatIntervalMs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityStreamRequest) SetHeartbeatIntervalMs(heartbeatIntervalMs *int) {
	e.HeartbeatIntervalMs = heartbeatIntervalMs
	e.require(entityStreamRequestFieldHeartbeatIntervalMs)
}

// SetPreExistingOnly sets the PreExistingOnly field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityStreamRequest) SetPreExistingOnly(preExistingOnly *bool) {
	e.PreExistingOnly = preExistingOnly
	e.require(entityStreamRequestFieldPreExistingOnly)
}

// SetComponentsToInclude sets the ComponentsToInclude field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityStreamRequest) SetComponentsToInclude(componentsToInclude []string) {
	e.ComponentsToInclude = componentsToInclude
	e.require(entityStreamRequestFieldComponentsToInclude)
}

// SetFilter sets the Filter field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityStreamRequest) SetFilter(filter *Statement) {
	e.Filter = filter
	e.require(entityStreamRequestFieldFilter)
}

func (e *EntityStreamRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler EntityStreamRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*e = EntityStreamRequest(body)
	return nil
}

func (e *EntityStreamRequest) MarshalJSON() ([]byte, error) {
	type embed EntityStreamRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, e.explicitFields)
	return json.Marshal(explicitMarshaler)
}

// The AndOperation represents the boolean AND operation, which is to be applied to the list of
//
//	children statement(s) or predicate(s).
var (
	andOperationFieldPredicateSet = big.NewInt(1 << 0)
	andOperationFieldStatementSet = big.NewInt(1 << 1)
)

type AndOperation struct {
	PredicateSet *PredicateSet `json:"predicateSet,omitempty" url:"predicateSet,omitempty"`
	StatementSet *StatementSet `json:"statementSet,omitempty" url:"statementSet,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AndOperation) GetPredicateSet() *PredicateSet {
	if a == nil {
		return nil
	}
	return a.PredicateSet
}

func (a *AndOperation) GetStatementSet() *StatementSet {
	if a == nil {
		return nil
	}
	return a.StatementSet
}

func (a *AndOperation) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *AndOperation) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetPredicateSet sets the PredicateSet field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AndOperation) SetPredicateSet(predicateSet *PredicateSet) {
	a.PredicateSet = predicateSet
	a.require(andOperationFieldPredicateSet)
}

// SetStatementSet sets the StatementSet field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AndOperation) SetStatementSet(statementSet *StatementSet) {
	a.StatementSet = statementSet
	a.require(andOperationFieldStatementSet)
}

func (a *AndOperation) UnmarshalJSON(data []byte) error {
	type unmarshaler AndOperation
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AndOperation(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AndOperation) MarshalJSON() ([]byte, error) {
	type embed AndOperation
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *AndOperation) String() string {
	if a == nil {
		return "<nil>"
	}
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

// The BooleanType represents a static boolean value.
var (
	booleanTypeFieldValue = big.NewInt(1 << 0)
)

type BooleanType struct {
	Value *bool `json:"value,omitempty" url:"value,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BooleanType) GetValue() *bool {
	if b == nil {
		return nil
	}
	return b.Value
}

func (b *BooleanType) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b *BooleanType) require(field *big.Int) {
	if b.explicitFields == nil {
		b.explicitFields = big.NewInt(0)
	}
	b.explicitFields.Or(b.explicitFields, field)
}

// SetValue sets the Value field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BooleanType) SetValue(value *bool) {
	b.Value = value
	b.require(booleanTypeFieldValue)
}

func (b *BooleanType) UnmarshalJSON(data []byte) error {
	type unmarshaler BooleanType
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BooleanType(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BooleanType) MarshalJSON() ([]byte, error) {
	type embed BooleanType
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, b.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (b *BooleanType) String() string {
	if b == nil {
		return "<nil>"
	}
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// The BoundedShapeType represents any static fully-enclosed shape.
var (
	boundedShapeTypeFieldPolygonValue = big.NewInt(1 << 0)
)

type BoundedShapeType struct {
	PolygonValue *GeoPolygon `json:"polygonValue,omitempty" url:"polygonValue,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BoundedShapeType) GetPolygonValue() *GeoPolygon {
	if b == nil {
		return nil
	}
	return b.PolygonValue
}

func (b *BoundedShapeType) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b *BoundedShapeType) require(field *big.Int) {
	if b.explicitFields == nil {
		b.explicitFields = big.NewInt(0)
	}
	b.explicitFields.Or(b.explicitFields, field)
}

// SetPolygonValue sets the PolygonValue field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BoundedShapeType) SetPolygonValue(polygonValue *GeoPolygon) {
	b.PolygonValue = polygonValue
	b.require(boundedShapeTypeFieldPolygonValue)
}

func (b *BoundedShapeType) UnmarshalJSON(data []byte) error {
	type unmarshaler BoundedShapeType
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BoundedShapeType(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BoundedShapeType) MarshalJSON() ([]byte, error) {
	type embed BoundedShapeType
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, b.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (b *BoundedShapeType) String() string {
	if b == nil {
		return "<nil>"
	}
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Event representing some type of entity change.
var (
	entityEventFieldEventType = big.NewInt(1 << 0)
	entityEventFieldTime      = big.NewInt(1 << 1)
	entityEventFieldEntity    = big.NewInt(1 << 2)
)

type EntityEvent struct {
	EventType *EntityEventEventType `json:"eventType,omitempty" url:"eventType,omitempty"`
	Time      *time.Time            `json:"time,omitempty" url:"time,omitempty"`
	Entity    *Entity               `json:"entity,omitempty" url:"entity,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *EntityEvent) GetEventType() *EntityEventEventType {
	if e == nil {
		return nil
	}
	return e.EventType
}

func (e *EntityEvent) GetTime() *time.Time {
	if e == nil {
		return nil
	}
	return e.Time
}

func (e *EntityEvent) GetEntity() *Entity {
	if e == nil {
		return nil
	}
	return e.Entity
}

func (e *EntityEvent) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e *EntityEvent) require(field *big.Int) {
	if e.explicitFields == nil {
		e.explicitFields = big.NewInt(0)
	}
	e.explicitFields.Or(e.explicitFields, field)
}

// SetEventType sets the EventType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityEvent) SetEventType(eventType *EntityEventEventType) {
	e.EventType = eventType
	e.require(entityEventFieldEventType)
}

// SetTime sets the Time field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityEvent) SetTime(time *time.Time) {
	e.Time = time
	e.require(entityEventFieldTime)
}

// SetEntity sets the Entity field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityEvent) SetEntity(entity *Entity) {
	e.Entity = entity
	e.require(entityEventFieldEntity)
}

func (e *EntityEvent) UnmarshalJSON(data []byte) error {
	type embed EntityEvent
	var unmarshaler = struct {
		embed
		Time *internal.DateTime `json:"time,omitempty"`
	}{
		embed: embed(*e),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*e = EntityEvent(unmarshaler.embed)
	e.Time = unmarshaler.Time.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *EntityEvent) MarshalJSON() ([]byte, error) {
	type embed EntityEvent
	var marshaler = struct {
		embed
		Time *internal.DateTime `json:"time,omitempty"`
	}{
		embed: embed(*e),
		Time:  internal.NewOptionalDateTime(e.Time),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, e.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (e *EntityEvent) String() string {
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

type EntityEventEventType string

const (
	EntityEventEventTypeEventTypeInvalid            EntityEventEventType = "EVENT_TYPE_INVALID"
	EntityEventEventTypeEventTypeCreated            EntityEventEventType = "EVENT_TYPE_CREATED"
	EntityEventEventTypeEventTypeUpdate             EntityEventEventType = "EVENT_TYPE_UPDATE"
	EntityEventEventTypeEventTypeDeleted            EntityEventEventType = "EVENT_TYPE_DELETED"
	EntityEventEventTypeEventTypePreexisting        EntityEventEventType = "EVENT_TYPE_PREEXISTING"
	EntityEventEventTypeEventTypePostExpiryOverride EntityEventEventType = "EVENT_TYPE_POST_EXPIRY_OVERRIDE"
)

func NewEntityEventEventTypeFromString(s string) (EntityEventEventType, error) {
	switch s {
	case "EVENT_TYPE_INVALID":
		return EntityEventEventTypeEventTypeInvalid, nil
	case "EVENT_TYPE_CREATED":
		return EntityEventEventTypeEventTypeCreated, nil
	case "EVENT_TYPE_UPDATE":
		return EntityEventEventTypeEventTypeUpdate, nil
	case "EVENT_TYPE_DELETED":
		return EntityEventEventTypeEventTypeDeleted, nil
	case "EVENT_TYPE_PREEXISTING":
		return EntityEventEventTypeEventTypePreexisting, nil
	case "EVENT_TYPE_POST_EXPIRY_OVERRIDE":
		return EntityEventEventTypeEventTypePostExpiryOverride, nil
	}
	var t EntityEventEventType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e EntityEventEventType) Ptr() *EntityEventEventType {
	return &e
}

var (
	entityEventResponseFieldSessionToken = big.NewInt(1 << 0)
	entityEventResponseFieldEntityEvents = big.NewInt(1 << 1)
)

type EntityEventResponse struct {
	// Long-poll session identifier. Use this token to resume polling on subsequent requests.
	SessionToken *string        `json:"sessionToken,omitempty" url:"sessionToken,omitempty"`
	EntityEvents []*EntityEvent `json:"entityEvents,omitempty" url:"entityEvents,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *EntityEventResponse) GetSessionToken() *string {
	if e == nil {
		return nil
	}
	return e.SessionToken
}

func (e *EntityEventResponse) GetEntityEvents() []*EntityEvent {
	if e == nil {
		return nil
	}
	return e.EntityEvents
}

func (e *EntityEventResponse) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e *EntityEventResponse) require(field *big.Int) {
	if e.explicitFields == nil {
		e.explicitFields = big.NewInt(0)
	}
	e.explicitFields.Or(e.explicitFields, field)
}

// SetSessionToken sets the SessionToken field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityEventResponse) SetSessionToken(sessionToken *string) {
	e.SessionToken = sessionToken
	e.require(entityEventResponseFieldSessionToken)
}

// SetEntityEvents sets the EntityEvents field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityEventResponse) SetEntityEvents(entityEvents []*EntityEvent) {
	e.EntityEvents = entityEvents
	e.require(entityEventResponseFieldEntityEvents)
}

func (e *EntityEventResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler EntityEventResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EntityEventResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *EntityEventResponse) MarshalJSON() ([]byte, error) {
	type embed EntityEventResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, e.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (e *EntityEventResponse) String() string {
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
	entityStreamEventFieldEventType = big.NewInt(1 << 0)
	entityStreamEventFieldTime      = big.NewInt(1 << 1)
	entityStreamEventFieldEntity    = big.NewInt(1 << 2)
)

type EntityStreamEvent struct {
	EventType *EntityEventEventType `json:"eventType,omitempty" url:"eventType,omitempty"`
	Time      *time.Time            `json:"time,omitempty" url:"time,omitempty"`
	Entity    *Entity               `json:"entity,omitempty" url:"entity,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *EntityStreamEvent) GetEventType() *EntityEventEventType {
	if e == nil {
		return nil
	}
	return e.EventType
}

func (e *EntityStreamEvent) GetTime() *time.Time {
	if e == nil {
		return nil
	}
	return e.Time
}

func (e *EntityStreamEvent) GetEntity() *Entity {
	if e == nil {
		return nil
	}
	return e.Entity
}

func (e *EntityStreamEvent) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e *EntityStreamEvent) require(field *big.Int) {
	if e.explicitFields == nil {
		e.explicitFields = big.NewInt(0)
	}
	e.explicitFields.Or(e.explicitFields, field)
}

// SetEventType sets the EventType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityStreamEvent) SetEventType(eventType *EntityEventEventType) {
	e.EventType = eventType
	e.require(entityStreamEventFieldEventType)
}

// SetTime sets the Time field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityStreamEvent) SetTime(time *time.Time) {
	e.Time = time
	e.require(entityStreamEventFieldTime)
}

// SetEntity sets the Entity field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityStreamEvent) SetEntity(entity *Entity) {
	e.Entity = entity
	e.require(entityStreamEventFieldEntity)
}

func (e *EntityStreamEvent) UnmarshalJSON(data []byte) error {
	type embed EntityStreamEvent
	var unmarshaler = struct {
		embed
		Time *internal.DateTime `json:"time,omitempty"`
	}{
		embed: embed(*e),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*e = EntityStreamEvent(unmarshaler.embed)
	e.Time = unmarshaler.Time.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *EntityStreamEvent) MarshalJSON() ([]byte, error) {
	type embed EntityStreamEvent
	var marshaler = struct {
		embed
		Time *internal.DateTime `json:"time,omitempty"`
	}{
		embed: embed(*e),
		Time:  internal.NewOptionalDateTime(e.Time),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, e.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (e *EntityStreamEvent) String() string {
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
	entityStreamHeartbeatFieldTimestamp = big.NewInt(1 << 0)
)

type EntityStreamHeartbeat struct {
	// The timestamp at which the heartbeat message was sent.
	Timestamp *string `json:"timestamp,omitempty" url:"timestamp,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *EntityStreamHeartbeat) GetTimestamp() *string {
	if e == nil {
		return nil
	}
	return e.Timestamp
}

func (e *EntityStreamHeartbeat) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e *EntityStreamHeartbeat) require(field *big.Int) {
	if e.explicitFields == nil {
		e.explicitFields = big.NewInt(0)
	}
	e.explicitFields.Or(e.explicitFields, field)
}

// SetTimestamp sets the Timestamp field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EntityStreamHeartbeat) SetTimestamp(timestamp *string) {
	e.Timestamp = timestamp
	e.require(entityStreamHeartbeatFieldTimestamp)
}

func (e *EntityStreamHeartbeat) UnmarshalJSON(data []byte) error {
	type unmarshaler EntityStreamHeartbeat
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EntityStreamHeartbeat(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *EntityStreamHeartbeat) MarshalJSON() ([]byte, error) {
	type embed EntityStreamHeartbeat
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, e.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (e *EntityStreamHeartbeat) String() string {
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

// The EnumType represents members of well-known anduril ontologies, such as "disposition." When
//
//	such a value is specified, the evaluation library expects the integer representation of the enum
//	value. For example, a disposition derived from ontology.v1 such as "DISPOSITION_HOSTILE" should be
//	represented with the integer value 2.
var (
	enumTypeFieldValue = big.NewInt(1 << 0)
)

type EnumType struct {
	Value *int `json:"value,omitempty" url:"value,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *EnumType) GetValue() *int {
	if e == nil {
		return nil
	}
	return e.Value
}

func (e *EnumType) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e *EnumType) require(field *big.Int) {
	if e.explicitFields == nil {
		e.explicitFields = big.NewInt(0)
	}
	e.explicitFields.Or(e.explicitFields, field)
}

// SetValue sets the Value field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *EnumType) SetValue(value *int) {
	e.Value = value
	e.require(enumTypeFieldValue)
}

func (e *EnumType) UnmarshalJSON(data []byte) error {
	type unmarshaler EnumType
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EnumType(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *EnumType) MarshalJSON() ([]byte, error) {
	type embed EnumType
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, e.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (e *EnumType) String() string {
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

// The HeadingType represents the heading in degrees for an entity's
//
//	attitudeEnu quaternion to be compared against. Defaults between a range of 0 to 360
var (
	headingTypeFieldValue = big.NewInt(1 << 0)
)

type HeadingType struct {
	Value *int `json:"value,omitempty" url:"value,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (h *HeadingType) GetValue() *int {
	if h == nil {
		return nil
	}
	return h.Value
}

func (h *HeadingType) GetExtraProperties() map[string]interface{} {
	if h == nil {
		return nil
	}
	return h.extraProperties
}

func (h *HeadingType) require(field *big.Int) {
	if h.explicitFields == nil {
		h.explicitFields = big.NewInt(0)
	}
	h.explicitFields.Or(h.explicitFields, field)
}

// SetValue sets the Value field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (h *HeadingType) SetValue(value *int) {
	h.Value = value
	h.require(headingTypeFieldValue)
}

func (h *HeadingType) UnmarshalJSON(data []byte) error {
	type unmarshaler HeadingType
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*h = HeadingType(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *h)
	if err != nil {
		return err
	}
	h.extraProperties = extraProperties
	h.rawJSON = json.RawMessage(data)
	return nil
}

func (h *HeadingType) MarshalJSON() ([]byte, error) {
	type embed HeadingType
	var marshaler = struct {
		embed
	}{
		embed: embed(*h),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, h.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (h *HeadingType) String() string {
	if h == nil {
		return "<nil>"
	}
	if len(h.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(h.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(h); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", h)
}

// The ListOperation represents an operation against a proto list. If the list is of primitive proto
//
//	type (e.g. int32), paths in all child predicates should be left empty. If the list is of message
//	proto type (e.g. Sensor), paths in all child predicates should be relative to the list path.
//
//	For example, the criteria "take an action if an entity has any sensor with sensor_id='sensor' and
//	OperationalState=STATE_OFF" would be modeled as:
//	Predicate1: { path: "sensor_id", comparator: EQUAL_TO, value: "sensor" }
//	Predicate2: { path: "operational_state", comparator: EQUAL_TO, value: STATE_OFF }
//
//	Statement2: { AndOperation: PredicateSet: { <Predicate1>, <Predicate2> } }
//	ListOperation: { list_path: "sensors.sensors", list_comparator: ANY, statement: <Statement2> }
//	Statement1: { ListOperation: <ListOperation> }
//
//	Note that in the above, the child predicates of the list operation have paths relative to the
//	list_path because the list is comprised of message not primitive types.
var (
	listOperationFieldListPath       = big.NewInt(1 << 0)
	listOperationFieldListComparator = big.NewInt(1 << 1)
	listOperationFieldStatement      = big.NewInt(1 << 2)
)

type ListOperation struct {
	// The list_path specifies the repeated field on an entity to which this operation applies.
	ListPath *string `json:"listPath,omitempty" url:"listPath,omitempty"`
	// The list_comparator specifies how to compose the boolean results from the child statement
	//
	//	for each member of the specified list.
	ListComparator *ListOperationListComparator `json:"listComparator,omitempty" url:"listComparator,omitempty"`
	// The statement is a new expression tree conceptually rooted at type of the list. It determines
	//
	//	how each member of the list is evaluated.
	Statement *Statement `json:"statement,omitempty" url:"statement,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListOperation) GetListPath() *string {
	if l == nil {
		return nil
	}
	return l.ListPath
}

func (l *ListOperation) GetListComparator() *ListOperationListComparator {
	if l == nil {
		return nil
	}
	return l.ListComparator
}

func (l *ListOperation) GetStatement() *Statement {
	if l == nil {
		return nil
	}
	return l.Statement
}

func (l *ListOperation) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListOperation) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetListPath sets the ListPath field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListOperation) SetListPath(listPath *string) {
	l.ListPath = listPath
	l.require(listOperationFieldListPath)
}

// SetListComparator sets the ListComparator field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListOperation) SetListComparator(listComparator *ListOperationListComparator) {
	l.ListComparator = listComparator
	l.require(listOperationFieldListComparator)
}

// SetStatement sets the Statement field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListOperation) SetStatement(statement *Statement) {
	l.Statement = statement
	l.require(listOperationFieldStatement)
}

func (l *ListOperation) UnmarshalJSON(data []byte) error {
	type unmarshaler ListOperation
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListOperation(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListOperation) MarshalJSON() ([]byte, error) {
	type embed ListOperation
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListOperation) String() string {
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

// The list_comparator specifies how to compose the boolean results from the child statement
//
//	for each member of the specified list.
type ListOperationListComparator string

const (
	ListOperationListComparatorListComparatorInvalid ListOperationListComparator = "LIST_COMPARATOR_INVALID"
	ListOperationListComparatorListComparatorAnyOf   ListOperationListComparator = "LIST_COMPARATOR_ANY_OF"
)

func NewListOperationListComparatorFromString(s string) (ListOperationListComparator, error) {
	switch s {
	case "LIST_COMPARATOR_INVALID":
		return ListOperationListComparatorListComparatorInvalid, nil
	case "LIST_COMPARATOR_ANY_OF":
		return ListOperationListComparatorListComparatorAnyOf, nil
	}
	var t ListOperationListComparator
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListOperationListComparator) Ptr() *ListOperationListComparator {
	return &l
}

// A List of Values for use with the IN comparator.
var (
	listTypeFieldValues = big.NewInt(1 << 0)
)

type ListType struct {
	Values []*Value `json:"values,omitempty" url:"values,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListType) GetValues() []*Value {
	if l == nil {
		return nil
	}
	return l.Values
}

func (l *ListType) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListType) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetValues sets the Values field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListType) SetValues(values []*Value) {
	l.Values = values
	l.require(listTypeFieldValues)
}

func (l *ListType) UnmarshalJSON(data []byte) error {
	type unmarshaler ListType
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListType(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListType) MarshalJSON() ([]byte, error) {
	type embed ListType
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListType) String() string {
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

// The NotOperation represents the boolean NOT operation, which can only be applied to a single
//
//	child predicate or statement.
var (
	notOperationFieldPredicate = big.NewInt(1 << 0)
	notOperationFieldStatement = big.NewInt(1 << 1)
)

type NotOperation struct {
	Predicate *Predicate `json:"predicate,omitempty" url:"predicate,omitempty"`
	Statement *Statement `json:"statement,omitempty" url:"statement,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (n *NotOperation) GetPredicate() *Predicate {
	if n == nil {
		return nil
	}
	return n.Predicate
}

func (n *NotOperation) GetStatement() *Statement {
	if n == nil {
		return nil
	}
	return n.Statement
}

func (n *NotOperation) GetExtraProperties() map[string]interface{} {
	if n == nil {
		return nil
	}
	return n.extraProperties
}

func (n *NotOperation) require(field *big.Int) {
	if n.explicitFields == nil {
		n.explicitFields = big.NewInt(0)
	}
	n.explicitFields.Or(n.explicitFields, field)
}

// SetPredicate sets the Predicate field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NotOperation) SetPredicate(predicate *Predicate) {
	n.Predicate = predicate
	n.require(notOperationFieldPredicate)
}

// SetStatement sets the Statement field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NotOperation) SetStatement(statement *Statement) {
	n.Statement = statement
	n.require(notOperationFieldStatement)
}

func (n *NotOperation) UnmarshalJSON(data []byte) error {
	type unmarshaler NotOperation
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*n = NotOperation(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *n)
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties
	n.rawJSON = json.RawMessage(data)
	return nil
}

func (n *NotOperation) MarshalJSON() ([]byte, error) {
	type embed NotOperation
	var marshaler = struct {
		embed
	}{
		embed: embed(*n),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, n.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (n *NotOperation) String() string {
	if n == nil {
		return "<nil>"
	}
	if len(n.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(n.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(n); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", n)
}

// The NumericType represents static numeric values. It supports all numeric primitives supported
//
//	by the proto3 language specification.
var (
	numericTypeFieldDoubleValue = big.NewInt(1 << 0)
	numericTypeFieldFloatValue  = big.NewInt(1 << 1)
	numericTypeFieldInt32Value  = big.NewInt(1 << 2)
	numericTypeFieldInt64Value  = big.NewInt(1 << 3)
	numericTypeFieldUint32Value = big.NewInt(1 << 4)
	numericTypeFieldUint64Value = big.NewInt(1 << 5)
)

type NumericType struct {
	DoubleValue *float64 `json:"doubleValue,omitempty" url:"doubleValue,omitempty"`
	FloatValue  *float64 `json:"floatValue,omitempty" url:"floatValue,omitempty"`
	Int32Value  *int     `json:"int32Value,omitempty" url:"int32Value,omitempty"`
	Int64Value  *string  `json:"int64Value,omitempty" url:"int64Value,omitempty"`
	Uint32Value *int     `json:"uint32Value,omitempty" url:"uint32Value,omitempty"`
	Uint64Value *string  `json:"uint64Value,omitempty" url:"uint64Value,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (n *NumericType) GetDoubleValue() *float64 {
	if n == nil {
		return nil
	}
	return n.DoubleValue
}

func (n *NumericType) GetFloatValue() *float64 {
	if n == nil {
		return nil
	}
	return n.FloatValue
}

func (n *NumericType) GetInt32Value() *int {
	if n == nil {
		return nil
	}
	return n.Int32Value
}

func (n *NumericType) GetInt64Value() *string {
	if n == nil {
		return nil
	}
	return n.Int64Value
}

func (n *NumericType) GetUint32Value() *int {
	if n == nil {
		return nil
	}
	return n.Uint32Value
}

func (n *NumericType) GetUint64Value() *string {
	if n == nil {
		return nil
	}
	return n.Uint64Value
}

func (n *NumericType) GetExtraProperties() map[string]interface{} {
	if n == nil {
		return nil
	}
	return n.extraProperties
}

func (n *NumericType) require(field *big.Int) {
	if n.explicitFields == nil {
		n.explicitFields = big.NewInt(0)
	}
	n.explicitFields.Or(n.explicitFields, field)
}

// SetDoubleValue sets the DoubleValue field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NumericType) SetDoubleValue(doubleValue *float64) {
	n.DoubleValue = doubleValue
	n.require(numericTypeFieldDoubleValue)
}

// SetFloatValue sets the FloatValue field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NumericType) SetFloatValue(floatValue *float64) {
	n.FloatValue = floatValue
	n.require(numericTypeFieldFloatValue)
}

// SetInt32Value sets the Int32Value field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NumericType) SetInt32Value(int32Value *int) {
	n.Int32Value = int32Value
	n.require(numericTypeFieldInt32Value)
}

// SetInt64Value sets the Int64Value field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NumericType) SetInt64Value(int64Value *string) {
	n.Int64Value = int64Value
	n.require(numericTypeFieldInt64Value)
}

// SetUint32Value sets the Uint32Value field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NumericType) SetUint32Value(uint32Value *int) {
	n.Uint32Value = uint32Value
	n.require(numericTypeFieldUint32Value)
}

// SetUint64Value sets the Uint64Value field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NumericType) SetUint64Value(uint64Value *string) {
	n.Uint64Value = uint64Value
	n.require(numericTypeFieldUint64Value)
}

func (n *NumericType) UnmarshalJSON(data []byte) error {
	type unmarshaler NumericType
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*n = NumericType(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *n)
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties
	n.rawJSON = json.RawMessage(data)
	return nil
}

func (n *NumericType) MarshalJSON() ([]byte, error) {
	type embed NumericType
	var marshaler = struct {
		embed
	}{
		embed: embed(*n),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, n.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (n *NumericType) String() string {
	if n == nil {
		return "<nil>"
	}
	if len(n.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(n.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(n); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", n)
}

// The OrOperation represents the boolean OR operation, which is to be applied to the list of
//
//	children statement(s) or predicate(s).
var (
	orOperationFieldPredicateSet = big.NewInt(1 << 0)
	orOperationFieldStatementSet = big.NewInt(1 << 1)
)

type OrOperation struct {
	PredicateSet *PredicateSet `json:"predicateSet,omitempty" url:"predicateSet,omitempty"`
	StatementSet *StatementSet `json:"statementSet,omitempty" url:"statementSet,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (o *OrOperation) GetPredicateSet() *PredicateSet {
	if o == nil {
		return nil
	}
	return o.PredicateSet
}

func (o *OrOperation) GetStatementSet() *StatementSet {
	if o == nil {
		return nil
	}
	return o.StatementSet
}

func (o *OrOperation) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o *OrOperation) require(field *big.Int) {
	if o.explicitFields == nil {
		o.explicitFields = big.NewInt(0)
	}
	o.explicitFields.Or(o.explicitFields, field)
}

// SetPredicateSet sets the PredicateSet field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (o *OrOperation) SetPredicateSet(predicateSet *PredicateSet) {
	o.PredicateSet = predicateSet
	o.require(orOperationFieldPredicateSet)
}

// SetStatementSet sets the StatementSet field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (o *OrOperation) SetStatementSet(statementSet *StatementSet) {
	o.StatementSet = statementSet
	o.require(orOperationFieldStatementSet)
}

func (o *OrOperation) UnmarshalJSON(data []byte) error {
	type unmarshaler OrOperation
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OrOperation(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	o.rawJSON = json.RawMessage(data)
	return nil
}

func (o *OrOperation) MarshalJSON() ([]byte, error) {
	type embed OrOperation
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, o.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (o *OrOperation) String() string {
	if o == nil {
		return "<nil>"
	}
	if len(o.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(o.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(o); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", o)
}

// The PositionType represents any fixed LLA point in space.
var (
	positionTypeFieldValue = big.NewInt(1 << 0)
)

type PositionType struct {
	Value *Position `json:"value,omitempty" url:"value,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PositionType) GetValue() *Position {
	if p == nil {
		return nil
	}
	return p.Value
}

func (p *PositionType) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p *PositionType) require(field *big.Int) {
	if p.explicitFields == nil {
		p.explicitFields = big.NewInt(0)
	}
	p.explicitFields.Or(p.explicitFields, field)
}

// SetValue sets the Value field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PositionType) SetValue(value *Position) {
	p.Value = value
	p.require(positionTypeFieldValue)
}

func (p *PositionType) UnmarshalJSON(data []byte) error {
	type unmarshaler PositionType
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PositionType(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PositionType) MarshalJSON() ([]byte, error) {
	type embed PositionType
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, p.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (p *PositionType) String() string {
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

// The Predicate fully encodes the information required to make an evaluation of an entity field
//
//	against a given static value, resulting in a boolean TRUE/FALSE result. The structure of a
//	predicate will always follow: "{entity-value} {comparator} {fixed-value}" where the entity value
//	is determined by the field path.
//
//	For example, a predicate would read as: "{entity.location.velocity_enu} {LESS_THAN} {500kph}"
var (
	predicateFieldFieldPath  = big.NewInt(1 << 0)
	predicateFieldValue      = big.NewInt(1 << 1)
	predicateFieldComparator = big.NewInt(1 << 2)
)

type Predicate struct {
	// The field_path determines which field on an entity is being referenced in this predicate. For
	//
	//	example: correlated.primary_entity_id would be primary_entity_id in correlated component.
	FieldPath *string `json:"fieldPath,omitempty" url:"fieldPath,omitempty"`
	// The value determines the fixed value against which the entity field is to be compared.
	//
	//	In the case of COMPARATOR_MATCH_ALL, the value contents do not matter as long as the Value is a supported
	//	type.
	Value *Value `json:"value,omitempty" url:"value,omitempty"`
	// The comparator determines the manner in which the entity field and static value are compared.
	//
	//	Comparators may only be applied to certain values. For example, the WITHIN comparator cannot
	//	be used for a boolean value comparison.
	Comparator *PredicateComparator `json:"comparator,omitempty" url:"comparator,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *Predicate) GetFieldPath() *string {
	if p == nil {
		return nil
	}
	return p.FieldPath
}

func (p *Predicate) GetValue() *Value {
	if p == nil {
		return nil
	}
	return p.Value
}

func (p *Predicate) GetComparator() *PredicateComparator {
	if p == nil {
		return nil
	}
	return p.Comparator
}

func (p *Predicate) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p *Predicate) require(field *big.Int) {
	if p.explicitFields == nil {
		p.explicitFields = big.NewInt(0)
	}
	p.explicitFields.Or(p.explicitFields, field)
}

// SetFieldPath sets the FieldPath field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *Predicate) SetFieldPath(fieldPath *string) {
	p.FieldPath = fieldPath
	p.require(predicateFieldFieldPath)
}

// SetValue sets the Value field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *Predicate) SetValue(value *Value) {
	p.Value = value
	p.require(predicateFieldValue)
}

// SetComparator sets the Comparator field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *Predicate) SetComparator(comparator *PredicateComparator) {
	p.Comparator = comparator
	p.require(predicateFieldComparator)
}

func (p *Predicate) UnmarshalJSON(data []byte) error {
	type unmarshaler Predicate
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = Predicate(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *Predicate) MarshalJSON() ([]byte, error) {
	type embed Predicate
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, p.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (p *Predicate) String() string {
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

// The comparator determines the manner in which the entity field and static value are compared.
//
//	Comparators may only be applied to certain values. For example, the WITHIN comparator cannot
//	be used for a boolean value comparison.
type PredicateComparator string

const (
	PredicateComparatorComparatorInvalid                   PredicateComparator = "COMPARATOR_INVALID"
	PredicateComparatorComparatorMatchAll                  PredicateComparator = "COMPARATOR_MATCH_ALL"
	PredicateComparatorComparatorEquality                  PredicateComparator = "COMPARATOR_EQUALITY"
	PredicateComparatorComparatorIn                        PredicateComparator = "COMPARATOR_IN"
	PredicateComparatorComparatorLessThan                  PredicateComparator = "COMPARATOR_LESS_THAN"
	PredicateComparatorComparatorGreaterThan               PredicateComparator = "COMPARATOR_GREATER_THAN"
	PredicateComparatorComparatorLessThanEqualTo           PredicateComparator = "COMPARATOR_LESS_THAN_EQUAL_TO"
	PredicateComparatorComparatorGreaterThanEqualTo        PredicateComparator = "COMPARATOR_GREATER_THAN_EQUAL_TO"
	PredicateComparatorComparatorWithin                    PredicateComparator = "COMPARATOR_WITHIN"
	PredicateComparatorComparatorExists                    PredicateComparator = "COMPARATOR_EXISTS"
	PredicateComparatorComparatorCaseInsensitiveEquality   PredicateComparator = "COMPARATOR_CASE_INSENSITIVE_EQUALITY"
	PredicateComparatorComparatorCaseInsensitiveEqualityIn PredicateComparator = "COMPARATOR_CASE_INSENSITIVE_EQUALITY_IN"
	PredicateComparatorComparatorRangeClosed               PredicateComparator = "COMPARATOR_RANGE_CLOSED"
)

func NewPredicateComparatorFromString(s string) (PredicateComparator, error) {
	switch s {
	case "COMPARATOR_INVALID":
		return PredicateComparatorComparatorInvalid, nil
	case "COMPARATOR_MATCH_ALL":
		return PredicateComparatorComparatorMatchAll, nil
	case "COMPARATOR_EQUALITY":
		return PredicateComparatorComparatorEquality, nil
	case "COMPARATOR_IN":
		return PredicateComparatorComparatorIn, nil
	case "COMPARATOR_LESS_THAN":
		return PredicateComparatorComparatorLessThan, nil
	case "COMPARATOR_GREATER_THAN":
		return PredicateComparatorComparatorGreaterThan, nil
	case "COMPARATOR_LESS_THAN_EQUAL_TO":
		return PredicateComparatorComparatorLessThanEqualTo, nil
	case "COMPARATOR_GREATER_THAN_EQUAL_TO":
		return PredicateComparatorComparatorGreaterThanEqualTo, nil
	case "COMPARATOR_WITHIN":
		return PredicateComparatorComparatorWithin, nil
	case "COMPARATOR_EXISTS":
		return PredicateComparatorComparatorExists, nil
	case "COMPARATOR_CASE_INSENSITIVE_EQUALITY":
		return PredicateComparatorComparatorCaseInsensitiveEquality, nil
	case "COMPARATOR_CASE_INSENSITIVE_EQUALITY_IN":
		return PredicateComparatorComparatorCaseInsensitiveEqualityIn, nil
	case "COMPARATOR_RANGE_CLOSED":
		return PredicateComparatorComparatorRangeClosed, nil
	}
	var t PredicateComparator
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p PredicateComparator) Ptr() *PredicateComparator {
	return &p
}

// The PredicateSet represents a list of predicates or "leaf nodes" in the expression tree, which
//
//	can be directly evaluated to a boolean TRUE/FALSE result.
var (
	predicateSetFieldPredicates = big.NewInt(1 << 0)
)

type PredicateSet struct {
	Predicates []*Predicate `json:"predicates,omitempty" url:"predicates,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PredicateSet) GetPredicates() []*Predicate {
	if p == nil {
		return nil
	}
	return p.Predicates
}

func (p *PredicateSet) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p *PredicateSet) require(field *big.Int) {
	if p.explicitFields == nil {
		p.explicitFields = big.NewInt(0)
	}
	p.explicitFields.Or(p.explicitFields, field)
}

// SetPredicates sets the Predicates field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PredicateSet) SetPredicates(predicates []*Predicate) {
	p.Predicates = predicates
	p.require(predicateSetFieldPredicates)
}

func (p *PredicateSet) UnmarshalJSON(data []byte) error {
	type unmarshaler PredicateSet
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PredicateSet(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PredicateSet) MarshalJSON() ([]byte, error) {
	type embed PredicateSet
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, p.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (p *PredicateSet) String() string {
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

// The RangeType represents a numeric range.
//
//	Whether endpoints are included are based on the comparator used.
//	Both endpoints must be of the same numeric type.
var (
	rangeTypeFieldStart = big.NewInt(1 << 0)
	rangeTypeFieldEnd   = big.NewInt(1 << 1)
)

type RangeType struct {
	Start *NumericType `json:"start,omitempty" url:"start,omitempty"`
	End   *NumericType `json:"end,omitempty" url:"end,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *RangeType) GetStart() *NumericType {
	if r == nil {
		return nil
	}
	return r.Start
}

func (r *RangeType) GetEnd() *NumericType {
	if r == nil {
		return nil
	}
	return r.End
}

func (r *RangeType) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *RangeType) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetStart sets the Start field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *RangeType) SetStart(start *NumericType) {
	r.Start = start
	r.require(rangeTypeFieldStart)
}

// SetEnd sets the End field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *RangeType) SetEnd(end *NumericType) {
	r.End = end
	r.require(rangeTypeFieldEnd)
}

func (r *RangeType) UnmarshalJSON(data []byte) error {
	type unmarshaler RangeType
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RangeType(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RangeType) MarshalJSON() ([]byte, error) {
	type embed RangeType
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *RangeType) String() string {
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

// A Statement is the building block of the entity filter. The outermost statement is conceptually
//
//	the root node of an "expression tree" which allows for the construction of complete boolean
//	logic statements. Statements are formed by grouping sets of children statement(s) or predicate(s)
//	according to the boolean operation which is to be applied.
//
//	For example, the criteria "take an action if an entity is hostile and an air vehicle" can be
//	represented as: Statement1: { AndOperation: { Predicate1, Predicate2 } }. Where Statement1
//	is the root of the expression tree, with an AND operation that is applied to children
//	predicates. The predicates themselves encode "entity is hostile" and "entity is air vehicle."
var (
	statementFieldAnd       = big.NewInt(1 << 0)
	statementFieldOr        = big.NewInt(1 << 1)
	statementFieldNot       = big.NewInt(1 << 2)
	statementFieldList      = big.NewInt(1 << 3)
	statementFieldPredicate = big.NewInt(1 << 4)
)

type Statement struct {
	And       *AndOperation  `json:"and,omitempty" url:"and,omitempty"`
	Or        *OrOperation   `json:"or,omitempty" url:"or,omitempty"`
	Not       *NotOperation  `json:"not,omitempty" url:"not,omitempty"`
	List      *ListOperation `json:"list,omitempty" url:"list,omitempty"`
	Predicate *Predicate     `json:"predicate,omitempty" url:"predicate,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *Statement) GetAnd() *AndOperation {
	if s == nil {
		return nil
	}
	return s.And
}

func (s *Statement) GetOr() *OrOperation {
	if s == nil {
		return nil
	}
	return s.Or
}

func (s *Statement) GetNot() *NotOperation {
	if s == nil {
		return nil
	}
	return s.Not
}

func (s *Statement) GetList() *ListOperation {
	if s == nil {
		return nil
	}
	return s.List
}

func (s *Statement) GetPredicate() *Predicate {
	if s == nil {
		return nil
	}
	return s.Predicate
}

func (s *Statement) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *Statement) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetAnd sets the And field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *Statement) SetAnd(and *AndOperation) {
	s.And = and
	s.require(statementFieldAnd)
}

// SetOr sets the Or field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *Statement) SetOr(or *OrOperation) {
	s.Or = or
	s.require(statementFieldOr)
}

// SetNot sets the Not field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *Statement) SetNot(not *NotOperation) {
	s.Not = not
	s.require(statementFieldNot)
}

// SetList sets the List field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *Statement) SetList(list *ListOperation) {
	s.List = list
	s.require(statementFieldList)
}

// SetPredicate sets the Predicate field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *Statement) SetPredicate(predicate *Predicate) {
	s.Predicate = predicate
	s.require(statementFieldPredicate)
}

func (s *Statement) UnmarshalJSON(data []byte) error {
	type unmarshaler Statement
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = Statement(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *Statement) MarshalJSON() ([]byte, error) {
	type embed Statement
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *Statement) String() string {
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

// The StatementSet represents a list of statements or "tree nodes," each of which follow the same
//
//	behavior as the Statement proto message.
var (
	statementSetFieldStatements = big.NewInt(1 << 0)
)

type StatementSet struct {
	Statements []*Statement `json:"statements,omitempty" url:"statements,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *StatementSet) GetStatements() []*Statement {
	if s == nil {
		return nil
	}
	return s.Statements
}

func (s *StatementSet) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *StatementSet) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetStatements sets the Statements field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *StatementSet) SetStatements(statements []*Statement) {
	s.Statements = statements
	s.require(statementSetFieldStatements)
}

func (s *StatementSet) UnmarshalJSON(data []byte) error {
	type unmarshaler StatementSet
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = StatementSet(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *StatementSet) MarshalJSON() ([]byte, error) {
	type embed StatementSet
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *StatementSet) String() string {
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

// The StringType represents static string values.
var (
	stringTypeFieldValue = big.NewInt(1 << 0)
)

type StringType struct {
	Value *string `json:"value,omitempty" url:"value,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *StringType) GetValue() *string {
	if s == nil {
		return nil
	}
	return s.Value
}

func (s *StringType) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *StringType) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetValue sets the Value field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *StringType) SetValue(value *string) {
	s.Value = value
	s.require(stringTypeFieldValue)
}

func (s *StringType) UnmarshalJSON(data []byte) error {
	type unmarshaler StringType
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = StringType(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *StringType) MarshalJSON() ([]byte, error) {
	type embed StringType
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *StringType) String() string {
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

// The TimestampType represents a static timestamp value.
var (
	timestampTypeFieldValue = big.NewInt(1 << 0)
)

type TimestampType struct {
	Value *time.Time `json:"value,omitempty" url:"value,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TimestampType) GetValue() *time.Time {
	if t == nil {
		return nil
	}
	return t.Value
}

func (t *TimestampType) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t *TimestampType) require(field *big.Int) {
	if t.explicitFields == nil {
		t.explicitFields = big.NewInt(0)
	}
	t.explicitFields.Or(t.explicitFields, field)
}

// SetValue sets the Value field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TimestampType) SetValue(value *time.Time) {
	t.Value = value
	t.require(timestampTypeFieldValue)
}

func (t *TimestampType) UnmarshalJSON(data []byte) error {
	type embed TimestampType
	var unmarshaler = struct {
		embed
		Value *internal.DateTime `json:"value,omitempty"`
	}{
		embed: embed(*t),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*t = TimestampType(unmarshaler.embed)
	t.Value = unmarshaler.Value.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TimestampType) MarshalJSON() ([]byte, error) {
	type embed TimestampType
	var marshaler = struct {
		embed
		Value *internal.DateTime `json:"value,omitempty"`
	}{
		embed: embed(*t),
		Value: internal.NewOptionalDateTime(t.Value),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, t.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (t *TimestampType) String() string {
	if t == nil {
		return "<nil>"
	}
	if len(t.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(t.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

// The Value represents the information against which an entity field is evaluated. It is one of
//
//	a fixed set of types, each of which correspond to specific comparators. See "ComparatorType"
//	for the full list of Value <-> Comparator mappings.
var (
	valueFieldBooleanType      = big.NewInt(1 << 0)
	valueFieldNumericType      = big.NewInt(1 << 1)
	valueFieldStringType       = big.NewInt(1 << 2)
	valueFieldEnumType         = big.NewInt(1 << 3)
	valueFieldTimestampType    = big.NewInt(1 << 4)
	valueFieldBoundedShapeType = big.NewInt(1 << 5)
	valueFieldPositionType     = big.NewInt(1 << 6)
	valueFieldHeadingType      = big.NewInt(1 << 7)
	valueFieldListType         = big.NewInt(1 << 8)
	valueFieldRangeType        = big.NewInt(1 << 9)
)

type Value struct {
	BooleanType      *BooleanType      `json:"booleanType,omitempty" url:"booleanType,omitempty"`
	NumericType      *NumericType      `json:"numericType,omitempty" url:"numericType,omitempty"`
	StringType       *StringType       `json:"stringType,omitempty" url:"stringType,omitempty"`
	EnumType         *EnumType         `json:"enumType,omitempty" url:"enumType,omitempty"`
	TimestampType    *TimestampType    `json:"timestampType,omitempty" url:"timestampType,omitempty"`
	BoundedShapeType *BoundedShapeType `json:"boundedShapeType,omitempty" url:"boundedShapeType,omitempty"`
	PositionType     *PositionType     `json:"positionType,omitempty" url:"positionType,omitempty"`
	HeadingType      *HeadingType      `json:"headingType,omitempty" url:"headingType,omitempty"`
	ListType         *ListType         `json:"listType,omitempty" url:"listType,omitempty"`
	RangeType        *RangeType        `json:"rangeType,omitempty" url:"rangeType,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (v *Value) GetBooleanType() *BooleanType {
	if v == nil {
		return nil
	}
	return v.BooleanType
}

func (v *Value) GetNumericType() *NumericType {
	if v == nil {
		return nil
	}
	return v.NumericType
}

func (v *Value) GetStringType() *StringType {
	if v == nil {
		return nil
	}
	return v.StringType
}

func (v *Value) GetEnumType() *EnumType {
	if v == nil {
		return nil
	}
	return v.EnumType
}

func (v *Value) GetTimestampType() *TimestampType {
	if v == nil {
		return nil
	}
	return v.TimestampType
}

func (v *Value) GetBoundedShapeType() *BoundedShapeType {
	if v == nil {
		return nil
	}
	return v.BoundedShapeType
}

func (v *Value) GetPositionType() *PositionType {
	if v == nil {
		return nil
	}
	return v.PositionType
}

func (v *Value) GetHeadingType() *HeadingType {
	if v == nil {
		return nil
	}
	return v.HeadingType
}

func (v *Value) GetListType() *ListType {
	if v == nil {
		return nil
	}
	return v.ListType
}

func (v *Value) GetRangeType() *RangeType {
	if v == nil {
		return nil
	}
	return v.RangeType
}

func (v *Value) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v *Value) require(field *big.Int) {
	if v.explicitFields == nil {
		v.explicitFields = big.NewInt(0)
	}
	v.explicitFields.Or(v.explicitFields, field)
}

// SetBooleanType sets the BooleanType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (v *Value) SetBooleanType(booleanType *BooleanType) {
	v.BooleanType = booleanType
	v.require(valueFieldBooleanType)
}

// SetNumericType sets the NumericType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (v *Value) SetNumericType(numericType *NumericType) {
	v.NumericType = numericType
	v.require(valueFieldNumericType)
}

// SetStringType sets the StringType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (v *Value) SetStringType(stringType *StringType) {
	v.StringType = stringType
	v.require(valueFieldStringType)
}

// SetEnumType sets the EnumType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (v *Value) SetEnumType(enumType *EnumType) {
	v.EnumType = enumType
	v.require(valueFieldEnumType)
}

// SetTimestampType sets the TimestampType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (v *Value) SetTimestampType(timestampType *TimestampType) {
	v.TimestampType = timestampType
	v.require(valueFieldTimestampType)
}

// SetBoundedShapeType sets the BoundedShapeType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (v *Value) SetBoundedShapeType(boundedShapeType *BoundedShapeType) {
	v.BoundedShapeType = boundedShapeType
	v.require(valueFieldBoundedShapeType)
}

// SetPositionType sets the PositionType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (v *Value) SetPositionType(positionType *PositionType) {
	v.PositionType = positionType
	v.require(valueFieldPositionType)
}

// SetHeadingType sets the HeadingType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (v *Value) SetHeadingType(headingType *HeadingType) {
	v.HeadingType = headingType
	v.require(valueFieldHeadingType)
}

// SetListType sets the ListType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (v *Value) SetListType(listType *ListType) {
	v.ListType = listType
	v.require(valueFieldListType)
}

// SetRangeType sets the RangeType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (v *Value) SetRangeType(rangeType *RangeType) {
	v.RangeType = rangeType
	v.require(valueFieldRangeType)
}

func (v *Value) UnmarshalJSON(data []byte) error {
	type unmarshaler Value
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = Value(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	v.rawJSON = json.RawMessage(data)
	return nil
}

func (v *Value) MarshalJSON() ([]byte, error) {
	type embed Value
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, v.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (v *Value) String() string {
	if v == nil {
		return "<nil>"
	}
	if len(v.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(v.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

// The stream event response.
type StreamEntitiesResponse struct {
	Event     string
	Heartbeat *EntityStreamHeartbeat
	Entity    *EntityStreamEvent

	rawJSON json.RawMessage
}

func (s *StreamEntitiesResponse) GetEvent() string {
	if s == nil {
		return ""
	}
	return s.Event
}

func (s *StreamEntitiesResponse) GetHeartbeat() *EntityStreamHeartbeat {
	if s == nil {
		return nil
	}
	return s.Heartbeat
}

func (s *StreamEntitiesResponse) GetEntity() *EntityStreamEvent {
	if s == nil {
		return nil
	}
	return s.Entity
}

func (s *StreamEntitiesResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Event string `json:"event"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	s.Event = unmarshaler.Event
	if unmarshaler.Event == "" {
		return fmt.Errorf("%T did not include discriminant event", s)
	}
	switch unmarshaler.Event {
	case "heartbeat":
		value := new(EntityStreamHeartbeat)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Heartbeat = value
	case "entity":
		value := new(EntityStreamEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Entity = value
	}
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s StreamEntitiesResponse) MarshalJSON() ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	if s.Heartbeat != nil {
		return internal.MarshalJSONWithExtraProperty(s.Heartbeat, "event", "heartbeat")
	}
	if s.Entity != nil {
		return internal.MarshalJSONWithExtraProperty(s.Entity, "event", "entity")
	}
	if len(s.rawJSON) > 0 {
		return s.rawJSON, nil
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", s)
}

type StreamEntitiesResponseVisitor interface {
	VisitHeartbeat(*EntityStreamHeartbeat) error
	VisitEntity(*EntityStreamEvent) error
}

func (s *StreamEntitiesResponse) Accept(visitor StreamEntitiesResponseVisitor) error {
	if s.Heartbeat != nil {
		return visitor.VisitHeartbeat(s.Heartbeat)
	}
	if s.Entity != nil {
		return visitor.VisitEntity(s.Entity)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", s)
}

func (s *StreamEntitiesResponse) validate() error {
	if s == nil {
		return fmt.Errorf("type %T is nil", s)
	}
	var fields []string
	if s.Heartbeat != nil {
		fields = append(fields, "heartbeat")
	}
	if s.Entity != nil {
		fields = append(fields, "entity")
	}
	if len(fields) == 0 {
		if s.Event != "" {
			if len(s.rawJSON) > 0 {
				return nil
			}
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", s, s.Event)
		}
		return fmt.Errorf("type %T is empty", s)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", s, fields)
	}
	if s.Event != "" {
		field := fields[0]
		if s.Event != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				s,
				s.Event,
				s,
			)
		}
	}
	return nil
}
