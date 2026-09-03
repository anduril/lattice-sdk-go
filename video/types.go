// Code generated from our API definition. DO NOT EDIT.

package video

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/anduril/lattice-sdk-go/v4/internal"
	big "math/big"
)

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
