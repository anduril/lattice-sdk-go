// Code generated from our API definition. DO NOT EDIT.

package video

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
)

func TestSettersGoogleProtobufAny(t *testing.T) {
	t.Run("SetType", func(t *testing.T) {
		obj := &GoogleProtobufAny{}
		var fernTestValueType *string
		obj.SetType(fernTestValueType)
		assert.Equal(t, fernTestValueType, obj.Type)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersGoogleProtobufAny(t *testing.T) {
	t.Run("GetType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GoogleProtobufAny{}
		var expected *string
		obj.Type = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetType(), "getter should return the property value")
	})

	t.Run("GetType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GoogleProtobufAny{}
		obj.Type = nil

		// Act & Assert
		assert.Nil(t, obj.GetType(), "getter should return nil when property is nil")
	})

	t.Run("GetType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GoogleProtobufAny
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetType() // Should return zero value
	})

}

func TestSettersMarkExplicitGoogleProtobufAny(t *testing.T) {
	t.Run("SetType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GoogleProtobufAny{}
		var fernTestValueType *string

		// Act
		obj.SetType(fernTestValueType)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersGoogleRPCStatus(t *testing.T) {
	t.Run("SetCode", func(t *testing.T) {
		obj := &GoogleRPCStatus{}
		var fernTestValueCode *int
		obj.SetCode(fernTestValueCode)
		assert.Equal(t, fernTestValueCode, obj.Code)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMessage", func(t *testing.T) {
		obj := &GoogleRPCStatus{}
		var fernTestValueMessage *string
		obj.SetMessage(fernTestValueMessage)
		assert.Equal(t, fernTestValueMessage, obj.Message)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDetails", func(t *testing.T) {
		obj := &GoogleRPCStatus{}
		var fernTestValueDetails []*GoogleProtobufAny
		obj.SetDetails(fernTestValueDetails)
		assert.Equal(t, fernTestValueDetails, obj.Details)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersGoogleRPCStatus(t *testing.T) {
	t.Run("GetCode", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GoogleRPCStatus{}
		var expected *int
		obj.Code = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCode(), "getter should return the property value")
	})

	t.Run("GetCode_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GoogleRPCStatus{}
		obj.Code = nil

		// Act & Assert
		assert.Nil(t, obj.GetCode(), "getter should return nil when property is nil")
	})

	t.Run("GetCode_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GoogleRPCStatus
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCode() // Should return zero value
	})

	t.Run("GetMessage", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GoogleRPCStatus{}
		var expected *string
		obj.Message = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMessage(), "getter should return the property value")
	})

	t.Run("GetMessage_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GoogleRPCStatus{}
		obj.Message = nil

		// Act & Assert
		assert.Nil(t, obj.GetMessage(), "getter should return nil when property is nil")
	})

	t.Run("GetMessage_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GoogleRPCStatus
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMessage() // Should return zero value
	})

	t.Run("GetDetails", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GoogleRPCStatus{}
		var expected []*GoogleProtobufAny
		obj.Details = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDetails(), "getter should return the property value")
	})

	t.Run("GetDetails_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GoogleRPCStatus{}
		obj.Details = nil

		// Act & Assert
		assert.Nil(t, obj.GetDetails(), "getter should return nil when property is nil")
	})

	t.Run("GetDetails_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GoogleRPCStatus
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDetails() // Should return zero value
	})

}

func TestSettersMarkExplicitGoogleRPCStatus(t *testing.T) {
	t.Run("SetCode_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GoogleRPCStatus{}
		var fernTestValueCode *int

		// Act
		obj.SetCode(fernTestValueCode)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetMessage_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GoogleRPCStatus{}
		var fernTestValueMessage *string

		// Act
		obj.SetMessage(fernTestValueMessage)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetDetails_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GoogleRPCStatus{}
		var fernTestValueDetails []*GoogleProtobufAny

		// Act
		obj.SetDetails(fernTestValueDetails)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestJSONMarshalingGoogleProtobufAny(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GoogleProtobufAny{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled GoogleProtobufAny
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj GoogleProtobufAny
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj GoogleProtobufAny
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingGoogleRPCStatus(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GoogleRPCStatus{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled GoogleRPCStatus
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj GoogleRPCStatus
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj GoogleRPCStatus
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestStringGoogleProtobufAny(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &GoogleProtobufAny{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GoogleProtobufAny
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringGoogleRPCStatus(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &GoogleRPCStatus{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GoogleRPCStatus
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestExtraPropertiesGoogleProtobufAny(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &GoogleProtobufAny{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GoogleProtobufAny
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesGoogleRPCStatus(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &GoogleRPCStatus{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GoogleRPCStatus
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}
