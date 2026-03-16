// Code generated from our API definition. DO NOT EDIT.

package Lattice

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
)

func TestSettersGetTokenRequest(t *testing.T) {
	t.Run("SetClientID", func(t *testing.T) {
		obj := &GetTokenRequest{}
		var fernTestValueClientID *string
		obj.SetClientID(fernTestValueClientID)
		assert.Equal(t, fernTestValueClientID, obj.ClientID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetClientSecret", func(t *testing.T) {
		obj := &GetTokenRequest{}
		var fernTestValueClientSecret *string
		obj.SetClientSecret(fernTestValueClientSecret)
		assert.Equal(t, fernTestValueClientSecret, obj.ClientSecret)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetTokenRequest(t *testing.T) {
	t.Run("SetClientID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenRequest{}
		var fernTestValueClientID *string

		// Act
		obj.SetClientID(fernTestValueClientID)

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

	t.Run("SetClientSecret_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenRequest{}
		var fernTestValueClientSecret *string

		// Act
		obj.SetClientSecret(fernTestValueClientSecret)

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

func TestSettersGetTokenResponse(t *testing.T) {
	t.Run("SetAccessToken", func(t *testing.T) {
		obj := &GetTokenResponse{}
		var fernTestValueAccessToken string
		obj.SetAccessToken(fernTestValueAccessToken)
		assert.Equal(t, fernTestValueAccessToken, obj.AccessToken)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTokenType", func(t *testing.T) {
		obj := &GetTokenResponse{}
		var fernTestValueTokenType string
		obj.SetTokenType(fernTestValueTokenType)
		assert.Equal(t, fernTestValueTokenType, obj.TokenType)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExpiresIn", func(t *testing.T) {
		obj := &GetTokenResponse{}
		var fernTestValueExpiresIn *int
		obj.SetExpiresIn(fernTestValueExpiresIn)
		assert.Equal(t, fernTestValueExpiresIn, obj.ExpiresIn)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRefreshExpiresIn", func(t *testing.T) {
		obj := &GetTokenResponse{}
		var fernTestValueRefreshExpiresIn *int
		obj.SetRefreshExpiresIn(fernTestValueRefreshExpiresIn)
		assert.Equal(t, fernTestValueRefreshExpiresIn, obj.RefreshExpiresIn)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetNotBeforePolicy", func(t *testing.T) {
		obj := &GetTokenResponse{}
		var fernTestValueNotBeforePolicy *int
		obj.SetNotBeforePolicy(fernTestValueNotBeforePolicy)
		assert.Equal(t, fernTestValueNotBeforePolicy, obj.NotBeforePolicy)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetScope", func(t *testing.T) {
		obj := &GetTokenResponse{}
		var fernTestValueScope *string
		obj.SetScope(fernTestValueScope)
		assert.Equal(t, fernTestValueScope, obj.Scope)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersGetTokenResponse(t *testing.T) {
	t.Run("GetAccessToken", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenResponse{}
		var expected string
		obj.AccessToken = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAccessToken(), "getter should return the property value")
	})

	t.Run("GetAccessToken_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetTokenResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAccessToken() // Should return zero value
	})

	t.Run("GetTokenType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenResponse{}
		var expected string
		obj.TokenType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTokenType(), "getter should return the property value")
	})

	t.Run("GetTokenType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetTokenResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTokenType() // Should return zero value
	})

	t.Run("GetExpiresIn", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenResponse{}
		var expected *int
		obj.ExpiresIn = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetExpiresIn(), "getter should return the property value")
	})

	t.Run("GetExpiresIn_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenResponse{}
		obj.ExpiresIn = nil

		// Act & Assert
		assert.Nil(t, obj.GetExpiresIn(), "getter should return nil when property is nil")
	})

	t.Run("GetExpiresIn_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetTokenResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetExpiresIn() // Should return zero value
	})

	t.Run("GetRefreshExpiresIn", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenResponse{}
		var expected *int
		obj.RefreshExpiresIn = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRefreshExpiresIn(), "getter should return the property value")
	})

	t.Run("GetRefreshExpiresIn_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenResponse{}
		obj.RefreshExpiresIn = nil

		// Act & Assert
		assert.Nil(t, obj.GetRefreshExpiresIn(), "getter should return nil when property is nil")
	})

	t.Run("GetRefreshExpiresIn_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetTokenResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRefreshExpiresIn() // Should return zero value
	})

	t.Run("GetNotBeforePolicy", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenResponse{}
		var expected *int
		obj.NotBeforePolicy = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetNotBeforePolicy(), "getter should return the property value")
	})

	t.Run("GetNotBeforePolicy_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenResponse{}
		obj.NotBeforePolicy = nil

		// Act & Assert
		assert.Nil(t, obj.GetNotBeforePolicy(), "getter should return nil when property is nil")
	})

	t.Run("GetNotBeforePolicy_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetTokenResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetNotBeforePolicy() // Should return zero value
	})

	t.Run("GetScope", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenResponse{}
		var expected *string
		obj.Scope = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetScope(), "getter should return the property value")
	})

	t.Run("GetScope_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenResponse{}
		obj.Scope = nil

		// Act & Assert
		assert.Nil(t, obj.GetScope(), "getter should return nil when property is nil")
	})

	t.Run("GetScope_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetTokenResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetScope() // Should return zero value
	})

}

func TestSettersMarkExplicitGetTokenResponse(t *testing.T) {
	t.Run("SetAccessToken_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenResponse{}
		var fernTestValueAccessToken string

		// Act
		obj.SetAccessToken(fernTestValueAccessToken)

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

	t.Run("SetTokenType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenResponse{}
		var fernTestValueTokenType string

		// Act
		obj.SetTokenType(fernTestValueTokenType)

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

	t.Run("SetExpiresIn_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenResponse{}
		var fernTestValueExpiresIn *int

		// Act
		obj.SetExpiresIn(fernTestValueExpiresIn)

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

	t.Run("SetRefreshExpiresIn_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenResponse{}
		var fernTestValueRefreshExpiresIn *int

		// Act
		obj.SetRefreshExpiresIn(fernTestValueRefreshExpiresIn)

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

	t.Run("SetNotBeforePolicy_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenResponse{}
		var fernTestValueNotBeforePolicy *int

		// Act
		obj.SetNotBeforePolicy(fernTestValueNotBeforePolicy)

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

	t.Run("SetScope_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenResponse{}
		var fernTestValueScope *string

		// Act
		obj.SetScope(fernTestValueScope)

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

func TestJSONMarshalingGetTokenResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTokenResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled GetTokenResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj GetTokenResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj GetTokenResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestStringGetTokenResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &GetTokenResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetTokenResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestExtraPropertiesGetTokenResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &GetTokenResponse{}
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
		var obj *GetTokenResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}
