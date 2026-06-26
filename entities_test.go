// Code generated from our API definition. DO NOT EDIT.

package Lattice

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
	time "time"
)

func TestSettersGetEntityRequest(t *testing.T) {
	t.Run("SetEntityID", func(t *testing.T) {
		obj := &GetEntityRequest{}
		var fernTestValueEntityID string
		obj.SetEntityID(fernTestValueEntityID)
		assert.Equal(t, fernTestValueEntityID, obj.EntityID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetEntityRequest(t *testing.T) {
	t.Run("SetEntityID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetEntityRequest{}
		var fernTestValueEntityID string

		// Act
		obj.SetEntityID(fernTestValueEntityID)

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

func TestSettersEntityEventRequest(t *testing.T) {
	t.Run("SetSessionToken", func(t *testing.T) {
		obj := &EntityEventRequest{}
		var fernTestValueSessionToken string
		obj.SetSessionToken(fernTestValueSessionToken)
		assert.Equal(t, fernTestValueSessionToken, obj.SessionToken)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetBatchSize", func(t *testing.T) {
		obj := &EntityEventRequest{}
		var fernTestValueBatchSize *int
		obj.SetBatchSize(fernTestValueBatchSize)
		assert.Equal(t, fernTestValueBatchSize, obj.BatchSize)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitEntityEventRequest(t *testing.T) {
	t.Run("SetSessionToken_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEventRequest{}
		var fernTestValueSessionToken string

		// Act
		obj.SetSessionToken(fernTestValueSessionToken)

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

	t.Run("SetBatchSize_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEventRequest{}
		var fernTestValueBatchSize *int

		// Act
		obj.SetBatchSize(fernTestValueBatchSize)

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

func TestSettersEntityOverride(t *testing.T) {
	t.Run("SetEntityID", func(t *testing.T) {
		obj := &EntityOverride{}
		var fernTestValueEntityID string
		obj.SetEntityID(fernTestValueEntityID)
		assert.Equal(t, fernTestValueEntityID, obj.EntityID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetFieldPath", func(t *testing.T) {
		obj := &EntityOverride{}
		var fernTestValueFieldPath string
		obj.SetFieldPath(fernTestValueFieldPath)
		assert.Equal(t, fernTestValueFieldPath, obj.FieldPath)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEntity", func(t *testing.T) {
		obj := &EntityOverride{}
		var fernTestValueEntity *Entity
		obj.SetEntity(fernTestValueEntity)
		assert.Equal(t, fernTestValueEntity, obj.Entity)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetProvenance", func(t *testing.T) {
		obj := &EntityOverride{}
		var fernTestValueProvenance *Provenance
		obj.SetProvenance(fernTestValueProvenance)
		assert.Equal(t, fernTestValueProvenance, obj.Provenance)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitEntityOverride(t *testing.T) {
	t.Run("SetEntityID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityOverride{}
		var fernTestValueEntityID string

		// Act
		obj.SetEntityID(fernTestValueEntityID)

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

	t.Run("SetFieldPath_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityOverride{}
		var fernTestValueFieldPath string

		// Act
		obj.SetFieldPath(fernTestValueFieldPath)

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

	t.Run("SetEntity_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityOverride{}
		var fernTestValueEntity *Entity

		// Act
		obj.SetEntity(fernTestValueEntity)

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

	t.Run("SetProvenance_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityOverride{}
		var fernTestValueProvenance *Provenance

		// Act
		obj.SetProvenance(fernTestValueProvenance)

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

func TestSettersRemoveEntityOverrideRequest(t *testing.T) {
	t.Run("SetEntityID", func(t *testing.T) {
		obj := &RemoveEntityOverrideRequest{}
		var fernTestValueEntityID string
		obj.SetEntityID(fernTestValueEntityID)
		assert.Equal(t, fernTestValueEntityID, obj.EntityID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetFieldPath", func(t *testing.T) {
		obj := &RemoveEntityOverrideRequest{}
		var fernTestValueFieldPath string
		obj.SetFieldPath(fernTestValueFieldPath)
		assert.Equal(t, fernTestValueFieldPath, obj.FieldPath)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitRemoveEntityOverrideRequest(t *testing.T) {
	t.Run("SetEntityID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RemoveEntityOverrideRequest{}
		var fernTestValueEntityID string

		// Act
		obj.SetEntityID(fernTestValueEntityID)

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

	t.Run("SetFieldPath_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RemoveEntityOverrideRequest{}
		var fernTestValueFieldPath string

		// Act
		obj.SetFieldPath(fernTestValueFieldPath)

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

func TestSettersEntityStreamRequest(t *testing.T) {
	t.Run("SetHeartbeatIntervalMs", func(t *testing.T) {
		obj := &EntityStreamRequest{}
		var fernTestValueHeartbeatIntervalMs *int
		obj.SetHeartbeatIntervalMs(fernTestValueHeartbeatIntervalMs)
		assert.Equal(t, fernTestValueHeartbeatIntervalMs, obj.HeartbeatIntervalMs)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetPreExistingOnly", func(t *testing.T) {
		obj := &EntityStreamRequest{}
		var fernTestValuePreExistingOnly *bool
		obj.SetPreExistingOnly(fernTestValuePreExistingOnly)
		assert.Equal(t, fernTestValuePreExistingOnly, obj.PreExistingOnly)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetComponentsToInclude", func(t *testing.T) {
		obj := &EntityStreamRequest{}
		var fernTestValueComponentsToInclude []string
		obj.SetComponentsToInclude(fernTestValueComponentsToInclude)
		assert.Equal(t, fernTestValueComponentsToInclude, obj.ComponentsToInclude)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetFilter", func(t *testing.T) {
		obj := &EntityStreamRequest{}
		var fernTestValueFilter *Statement
		obj.SetFilter(fernTestValueFilter)
		assert.Equal(t, fernTestValueFilter, obj.Filter)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitEntityStreamRequest(t *testing.T) {
	t.Run("SetHeartbeatIntervalMs_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamRequest{}
		var fernTestValueHeartbeatIntervalMs *int

		// Act
		obj.SetHeartbeatIntervalMs(fernTestValueHeartbeatIntervalMs)

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

	t.Run("SetPreExistingOnly_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamRequest{}
		var fernTestValuePreExistingOnly *bool

		// Act
		obj.SetPreExistingOnly(fernTestValuePreExistingOnly)

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

	t.Run("SetComponentsToInclude_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamRequest{}
		var fernTestValueComponentsToInclude []string

		// Act
		obj.SetComponentsToInclude(fernTestValueComponentsToInclude)

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

	t.Run("SetFilter_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamRequest{}
		var fernTestValueFilter *Statement

		// Act
		obj.SetFilter(fernTestValueFilter)

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

func TestSettersAndOperation(t *testing.T) {
	t.Run("SetPredicateSet", func(t *testing.T) {
		obj := &AndOperation{}
		var fernTestValuePredicateSet *PredicateSet
		obj.SetPredicateSet(fernTestValuePredicateSet)
		assert.Equal(t, fernTestValuePredicateSet, obj.PredicateSet)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatementSet", func(t *testing.T) {
		obj := &AndOperation{}
		var fernTestValueStatementSet *StatementSet
		obj.SetStatementSet(fernTestValueStatementSet)
		assert.Equal(t, fernTestValueStatementSet, obj.StatementSet)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersAndOperation(t *testing.T) {
	t.Run("GetPredicateSet", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AndOperation{}
		var expected *PredicateSet
		obj.PredicateSet = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPredicateSet(), "getter should return the property value")
	})

	t.Run("GetPredicateSet_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AndOperation{}
		obj.PredicateSet = nil

		// Act & Assert
		assert.Nil(t, obj.GetPredicateSet(), "getter should return nil when property is nil")
	})

	t.Run("GetPredicateSet_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AndOperation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPredicateSet() // Should return zero value
	})

	t.Run("GetStatementSet", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AndOperation{}
		var expected *StatementSet
		obj.StatementSet = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatementSet(), "getter should return the property value")
	})

	t.Run("GetStatementSet_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AndOperation{}
		obj.StatementSet = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatementSet(), "getter should return nil when property is nil")
	})

	t.Run("GetStatementSet_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AndOperation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatementSet() // Should return zero value
	})

}

func TestSettersMarkExplicitAndOperation(t *testing.T) {
	t.Run("SetPredicateSet_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AndOperation{}
		var fernTestValuePredicateSet *PredicateSet

		// Act
		obj.SetPredicateSet(fernTestValuePredicateSet)

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

	t.Run("SetStatementSet_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AndOperation{}
		var fernTestValueStatementSet *StatementSet

		// Act
		obj.SetStatementSet(fernTestValueStatementSet)

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

func TestSettersBooleanType(t *testing.T) {
	t.Run("SetValue", func(t *testing.T) {
		obj := &BooleanType{}
		var fernTestValueValue *bool
		obj.SetValue(fernTestValueValue)
		assert.Equal(t, fernTestValueValue, obj.Value)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersBooleanType(t *testing.T) {
	t.Run("GetValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BooleanType{}
		var expected *bool
		obj.Value = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetValue(), "getter should return the property value")
	})

	t.Run("GetValue_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BooleanType{}
		obj.Value = nil

		// Act & Assert
		assert.Nil(t, obj.GetValue(), "getter should return nil when property is nil")
	})

	t.Run("GetValue_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BooleanType
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetValue() // Should return zero value
	})

}

func TestSettersMarkExplicitBooleanType(t *testing.T) {
	t.Run("SetValue_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BooleanType{}
		var fernTestValueValue *bool

		// Act
		obj.SetValue(fernTestValueValue)

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

func TestSettersBoundedShapeType(t *testing.T) {
	t.Run("SetPolygonValue", func(t *testing.T) {
		obj := &BoundedShapeType{}
		var fernTestValuePolygonValue *GeoPolygon
		obj.SetPolygonValue(fernTestValuePolygonValue)
		assert.Equal(t, fernTestValuePolygonValue, obj.PolygonValue)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersBoundedShapeType(t *testing.T) {
	t.Run("GetPolygonValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundedShapeType{}
		var expected *GeoPolygon
		obj.PolygonValue = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPolygonValue(), "getter should return the property value")
	})

	t.Run("GetPolygonValue_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundedShapeType{}
		obj.PolygonValue = nil

		// Act & Assert
		assert.Nil(t, obj.GetPolygonValue(), "getter should return nil when property is nil")
	})

	t.Run("GetPolygonValue_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BoundedShapeType
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPolygonValue() // Should return zero value
	})

}

func TestSettersMarkExplicitBoundedShapeType(t *testing.T) {
	t.Run("SetPolygonValue_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundedShapeType{}
		var fernTestValuePolygonValue *GeoPolygon

		// Act
		obj.SetPolygonValue(fernTestValuePolygonValue)

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

func TestSettersEntityEvent(t *testing.T) {
	t.Run("SetEventType", func(t *testing.T) {
		obj := &EntityEvent{}
		var fernTestValueEventType *EntityEventEventType
		obj.SetEventType(fernTestValueEventType)
		assert.Equal(t, fernTestValueEventType, obj.EventType)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTime", func(t *testing.T) {
		obj := &EntityEvent{}
		var fernTestValueTime *time.Time
		obj.SetTime(fernTestValueTime)
		assert.Equal(t, fernTestValueTime, obj.Time)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEntity", func(t *testing.T) {
		obj := &EntityEvent{}
		var fernTestValueEntity *Entity
		obj.SetEntity(fernTestValueEntity)
		assert.Equal(t, fernTestValueEntity, obj.Entity)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersEntityEvent(t *testing.T) {
	t.Run("GetEventType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEvent{}
		var expected *EntityEventEventType
		obj.EventType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEventType(), "getter should return the property value")
	})

	t.Run("GetEventType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEvent{}
		obj.EventType = nil

		// Act & Assert
		assert.Nil(t, obj.GetEventType(), "getter should return nil when property is nil")
	})

	t.Run("GetEventType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EntityEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEventType() // Should return zero value
	})

	t.Run("GetTime", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEvent{}
		var expected *time.Time
		obj.Time = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTime(), "getter should return the property value")
	})

	t.Run("GetTime_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEvent{}
		obj.Time = nil

		// Act & Assert
		assert.Nil(t, obj.GetTime(), "getter should return nil when property is nil")
	})

	t.Run("GetTime_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EntityEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTime() // Should return zero value
	})

	t.Run("GetEntity", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEvent{}
		var expected *Entity
		obj.Entity = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEntity(), "getter should return the property value")
	})

	t.Run("GetEntity_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEvent{}
		obj.Entity = nil

		// Act & Assert
		assert.Nil(t, obj.GetEntity(), "getter should return nil when property is nil")
	})

	t.Run("GetEntity_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EntityEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEntity() // Should return zero value
	})

}

func TestSettersMarkExplicitEntityEvent(t *testing.T) {
	t.Run("SetEventType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEvent{}
		var fernTestValueEventType *EntityEventEventType

		// Act
		obj.SetEventType(fernTestValueEventType)

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

	t.Run("SetTime_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEvent{}
		var fernTestValueTime *time.Time

		// Act
		obj.SetTime(fernTestValueTime)

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

	t.Run("SetEntity_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEvent{}
		var fernTestValueEntity *Entity

		// Act
		obj.SetEntity(fernTestValueEntity)

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

func TestSettersEntityEventResponse(t *testing.T) {
	t.Run("SetSessionToken", func(t *testing.T) {
		obj := &EntityEventResponse{}
		var fernTestValueSessionToken *string
		obj.SetSessionToken(fernTestValueSessionToken)
		assert.Equal(t, fernTestValueSessionToken, obj.SessionToken)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEntityEvents", func(t *testing.T) {
		obj := &EntityEventResponse{}
		var fernTestValueEntityEvents []*EntityEvent
		obj.SetEntityEvents(fernTestValueEntityEvents)
		assert.Equal(t, fernTestValueEntityEvents, obj.EntityEvents)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersEntityEventResponse(t *testing.T) {
	t.Run("GetSessionToken", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEventResponse{}
		var expected *string
		obj.SessionToken = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSessionToken(), "getter should return the property value")
	})

	t.Run("GetSessionToken_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEventResponse{}
		obj.SessionToken = nil

		// Act & Assert
		assert.Nil(t, obj.GetSessionToken(), "getter should return nil when property is nil")
	})

	t.Run("GetSessionToken_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EntityEventResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSessionToken() // Should return zero value
	})

	t.Run("GetEntityEvents", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEventResponse{}
		var expected []*EntityEvent
		obj.EntityEvents = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEntityEvents(), "getter should return the property value")
	})

	t.Run("GetEntityEvents_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEventResponse{}
		obj.EntityEvents = nil

		// Act & Assert
		assert.Nil(t, obj.GetEntityEvents(), "getter should return nil when property is nil")
	})

	t.Run("GetEntityEvents_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EntityEventResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEntityEvents() // Should return zero value
	})

}

func TestSettersMarkExplicitEntityEventResponse(t *testing.T) {
	t.Run("SetSessionToken_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEventResponse{}
		var fernTestValueSessionToken *string

		// Act
		obj.SetSessionToken(fernTestValueSessionToken)

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

	t.Run("SetEntityEvents_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEventResponse{}
		var fernTestValueEntityEvents []*EntityEvent

		// Act
		obj.SetEntityEvents(fernTestValueEntityEvents)

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

func TestSettersEntityStreamEvent(t *testing.T) {
	t.Run("SetEventType", func(t *testing.T) {
		obj := &EntityStreamEvent{}
		var fernTestValueEventType *EntityEventEventType
		obj.SetEventType(fernTestValueEventType)
		assert.Equal(t, fernTestValueEventType, obj.EventType)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTime", func(t *testing.T) {
		obj := &EntityStreamEvent{}
		var fernTestValueTime *time.Time
		obj.SetTime(fernTestValueTime)
		assert.Equal(t, fernTestValueTime, obj.Time)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEntity", func(t *testing.T) {
		obj := &EntityStreamEvent{}
		var fernTestValueEntity *Entity
		obj.SetEntity(fernTestValueEntity)
		assert.Equal(t, fernTestValueEntity, obj.Entity)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersEntityStreamEvent(t *testing.T) {
	t.Run("GetEventType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamEvent{}
		var expected *EntityEventEventType
		obj.EventType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEventType(), "getter should return the property value")
	})

	t.Run("GetEventType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamEvent{}
		obj.EventType = nil

		// Act & Assert
		assert.Nil(t, obj.GetEventType(), "getter should return nil when property is nil")
	})

	t.Run("GetEventType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EntityStreamEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEventType() // Should return zero value
	})

	t.Run("GetTime", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamEvent{}
		var expected *time.Time
		obj.Time = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTime(), "getter should return the property value")
	})

	t.Run("GetTime_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamEvent{}
		obj.Time = nil

		// Act & Assert
		assert.Nil(t, obj.GetTime(), "getter should return nil when property is nil")
	})

	t.Run("GetTime_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EntityStreamEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTime() // Should return zero value
	})

	t.Run("GetEntity", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamEvent{}
		var expected *Entity
		obj.Entity = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEntity(), "getter should return the property value")
	})

	t.Run("GetEntity_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamEvent{}
		obj.Entity = nil

		// Act & Assert
		assert.Nil(t, obj.GetEntity(), "getter should return nil when property is nil")
	})

	t.Run("GetEntity_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EntityStreamEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEntity() // Should return zero value
	})

}

func TestSettersMarkExplicitEntityStreamEvent(t *testing.T) {
	t.Run("SetEventType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamEvent{}
		var fernTestValueEventType *EntityEventEventType

		// Act
		obj.SetEventType(fernTestValueEventType)

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

	t.Run("SetTime_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamEvent{}
		var fernTestValueTime *time.Time

		// Act
		obj.SetTime(fernTestValueTime)

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

	t.Run("SetEntity_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamEvent{}
		var fernTestValueEntity *Entity

		// Act
		obj.SetEntity(fernTestValueEntity)

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

func TestSettersEntityStreamHeartbeat(t *testing.T) {
	t.Run("SetTimestamp", func(t *testing.T) {
		obj := &EntityStreamHeartbeat{}
		var fernTestValueTimestamp *string
		obj.SetTimestamp(fernTestValueTimestamp)
		assert.Equal(t, fernTestValueTimestamp, obj.Timestamp)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersEntityStreamHeartbeat(t *testing.T) {
	t.Run("GetTimestamp", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamHeartbeat{}
		var expected *string
		obj.Timestamp = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTimestamp(), "getter should return the property value")
	})

	t.Run("GetTimestamp_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamHeartbeat{}
		obj.Timestamp = nil

		// Act & Assert
		assert.Nil(t, obj.GetTimestamp(), "getter should return nil when property is nil")
	})

	t.Run("GetTimestamp_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EntityStreamHeartbeat
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTimestamp() // Should return zero value
	})

}

func TestSettersMarkExplicitEntityStreamHeartbeat(t *testing.T) {
	t.Run("SetTimestamp_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamHeartbeat{}
		var fernTestValueTimestamp *string

		// Act
		obj.SetTimestamp(fernTestValueTimestamp)

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

func TestSettersEnumType(t *testing.T) {
	t.Run("SetValue", func(t *testing.T) {
		obj := &EnumType{}
		var fernTestValueValue *int
		obj.SetValue(fernTestValueValue)
		assert.Equal(t, fernTestValueValue, obj.Value)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersEnumType(t *testing.T) {
	t.Run("GetValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EnumType{}
		var expected *int
		obj.Value = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetValue(), "getter should return the property value")
	})

	t.Run("GetValue_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EnumType{}
		obj.Value = nil

		// Act & Assert
		assert.Nil(t, obj.GetValue(), "getter should return nil when property is nil")
	})

	t.Run("GetValue_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EnumType
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetValue() // Should return zero value
	})

}

func TestSettersMarkExplicitEnumType(t *testing.T) {
	t.Run("SetValue_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EnumType{}
		var fernTestValueValue *int

		// Act
		obj.SetValue(fernTestValueValue)

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

func TestSettersHeadingType(t *testing.T) {
	t.Run("SetValue", func(t *testing.T) {
		obj := &HeadingType{}
		var fernTestValueValue *int
		obj.SetValue(fernTestValueValue)
		assert.Equal(t, fernTestValueValue, obj.Value)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersHeadingType(t *testing.T) {
	t.Run("GetValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &HeadingType{}
		var expected *int
		obj.Value = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetValue(), "getter should return the property value")
	})

	t.Run("GetValue_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &HeadingType{}
		obj.Value = nil

		// Act & Assert
		assert.Nil(t, obj.GetValue(), "getter should return nil when property is nil")
	})

	t.Run("GetValue_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *HeadingType
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetValue() // Should return zero value
	})

}

func TestSettersMarkExplicitHeadingType(t *testing.T) {
	t.Run("SetValue_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &HeadingType{}
		var fernTestValueValue *int

		// Act
		obj.SetValue(fernTestValueValue)

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

func TestSettersListOperation(t *testing.T) {
	t.Run("SetListPath", func(t *testing.T) {
		obj := &ListOperation{}
		var fernTestValueListPath *string
		obj.SetListPath(fernTestValueListPath)
		assert.Equal(t, fernTestValueListPath, obj.ListPath)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetListComparator", func(t *testing.T) {
		obj := &ListOperation{}
		var fernTestValueListComparator *ListOperationListComparator
		obj.SetListComparator(fernTestValueListComparator)
		assert.Equal(t, fernTestValueListComparator, obj.ListComparator)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatement", func(t *testing.T) {
		obj := &ListOperation{}
		var fernTestValueStatement *Statement
		obj.SetStatement(fernTestValueStatement)
		assert.Equal(t, fernTestValueStatement, obj.Statement)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListOperation(t *testing.T) {
	t.Run("GetListPath", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListOperation{}
		var expected *string
		obj.ListPath = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetListPath(), "getter should return the property value")
	})

	t.Run("GetListPath_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListOperation{}
		obj.ListPath = nil

		// Act & Assert
		assert.Nil(t, obj.GetListPath(), "getter should return nil when property is nil")
	})

	t.Run("GetListPath_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListOperation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetListPath() // Should return zero value
	})

	t.Run("GetListComparator", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListOperation{}
		var expected *ListOperationListComparator
		obj.ListComparator = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetListComparator(), "getter should return the property value")
	})

	t.Run("GetListComparator_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListOperation{}
		obj.ListComparator = nil

		// Act & Assert
		assert.Nil(t, obj.GetListComparator(), "getter should return nil when property is nil")
	})

	t.Run("GetListComparator_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListOperation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetListComparator() // Should return zero value
	})

	t.Run("GetStatement", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListOperation{}
		var expected *Statement
		obj.Statement = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatement(), "getter should return the property value")
	})

	t.Run("GetStatement_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListOperation{}
		obj.Statement = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatement(), "getter should return nil when property is nil")
	})

	t.Run("GetStatement_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListOperation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatement() // Should return zero value
	})

}

func TestSettersMarkExplicitListOperation(t *testing.T) {
	t.Run("SetListPath_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListOperation{}
		var fernTestValueListPath *string

		// Act
		obj.SetListPath(fernTestValueListPath)

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

	t.Run("SetListComparator_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListOperation{}
		var fernTestValueListComparator *ListOperationListComparator

		// Act
		obj.SetListComparator(fernTestValueListComparator)

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

	t.Run("SetStatement_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListOperation{}
		var fernTestValueStatement *Statement

		// Act
		obj.SetStatement(fernTestValueStatement)

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

func TestSettersListType(t *testing.T) {
	t.Run("SetValues", func(t *testing.T) {
		obj := &ListType{}
		var fernTestValueValues []*Value
		obj.SetValues(fernTestValueValues)
		assert.Equal(t, fernTestValueValues, obj.Values)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListType(t *testing.T) {
	t.Run("GetValues", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListType{}
		var expected []*Value
		obj.Values = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetValues(), "getter should return the property value")
	})

	t.Run("GetValues_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListType{}
		obj.Values = nil

		// Act & Assert
		assert.Nil(t, obj.GetValues(), "getter should return nil when property is nil")
	})

	t.Run("GetValues_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListType
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetValues() // Should return zero value
	})

}

func TestSettersMarkExplicitListType(t *testing.T) {
	t.Run("SetValues_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListType{}
		var fernTestValueValues []*Value

		// Act
		obj.SetValues(fernTestValueValues)

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

func TestSettersNotOperation(t *testing.T) {
	t.Run("SetPredicate", func(t *testing.T) {
		obj := &NotOperation{}
		var fernTestValuePredicate *Predicate
		obj.SetPredicate(fernTestValuePredicate)
		assert.Equal(t, fernTestValuePredicate, obj.Predicate)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatement", func(t *testing.T) {
		obj := &NotOperation{}
		var fernTestValueStatement *Statement
		obj.SetStatement(fernTestValueStatement)
		assert.Equal(t, fernTestValueStatement, obj.Statement)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersNotOperation(t *testing.T) {
	t.Run("GetPredicate", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotOperation{}
		var expected *Predicate
		obj.Predicate = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPredicate(), "getter should return the property value")
	})

	t.Run("GetPredicate_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotOperation{}
		obj.Predicate = nil

		// Act & Assert
		assert.Nil(t, obj.GetPredicate(), "getter should return nil when property is nil")
	})

	t.Run("GetPredicate_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotOperation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPredicate() // Should return zero value
	})

	t.Run("GetStatement", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotOperation{}
		var expected *Statement
		obj.Statement = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatement(), "getter should return the property value")
	})

	t.Run("GetStatement_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotOperation{}
		obj.Statement = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatement(), "getter should return nil when property is nil")
	})

	t.Run("GetStatement_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotOperation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatement() // Should return zero value
	})

}

func TestSettersMarkExplicitNotOperation(t *testing.T) {
	t.Run("SetPredicate_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotOperation{}
		var fernTestValuePredicate *Predicate

		// Act
		obj.SetPredicate(fernTestValuePredicate)

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

	t.Run("SetStatement_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotOperation{}
		var fernTestValueStatement *Statement

		// Act
		obj.SetStatement(fernTestValueStatement)

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

func TestSettersNumericType(t *testing.T) {
	t.Run("SetDoubleValue", func(t *testing.T) {
		obj := &NumericType{}
		var fernTestValueDoubleValue *float64
		obj.SetDoubleValue(fernTestValueDoubleValue)
		assert.Equal(t, fernTestValueDoubleValue, obj.DoubleValue)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetFloatValue", func(t *testing.T) {
		obj := &NumericType{}
		var fernTestValueFloatValue *float64
		obj.SetFloatValue(fernTestValueFloatValue)
		assert.Equal(t, fernTestValueFloatValue, obj.FloatValue)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetInt32Value", func(t *testing.T) {
		obj := &NumericType{}
		var fernTestValueInt32Value *int
		obj.SetInt32Value(fernTestValueInt32Value)
		assert.Equal(t, fernTestValueInt32Value, obj.Int32Value)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetInt64Value", func(t *testing.T) {
		obj := &NumericType{}
		var fernTestValueInt64Value *string
		obj.SetInt64Value(fernTestValueInt64Value)
		assert.Equal(t, fernTestValueInt64Value, obj.Int64Value)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetUint32Value", func(t *testing.T) {
		obj := &NumericType{}
		var fernTestValueUint32Value *int
		obj.SetUint32Value(fernTestValueUint32Value)
		assert.Equal(t, fernTestValueUint32Value, obj.Uint32Value)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetUint64Value", func(t *testing.T) {
		obj := &NumericType{}
		var fernTestValueUint64Value *string
		obj.SetUint64Value(fernTestValueUint64Value)
		assert.Equal(t, fernTestValueUint64Value, obj.Uint64Value)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersNumericType(t *testing.T) {
	t.Run("GetDoubleValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		var expected *float64
		obj.DoubleValue = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDoubleValue(), "getter should return the property value")
	})

	t.Run("GetDoubleValue_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		obj.DoubleValue = nil

		// Act & Assert
		assert.Nil(t, obj.GetDoubleValue(), "getter should return nil when property is nil")
	})

	t.Run("GetDoubleValue_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NumericType
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDoubleValue() // Should return zero value
	})

	t.Run("GetFloatValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		var expected *float64
		obj.FloatValue = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetFloatValue(), "getter should return the property value")
	})

	t.Run("GetFloatValue_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		obj.FloatValue = nil

		// Act & Assert
		assert.Nil(t, obj.GetFloatValue(), "getter should return nil when property is nil")
	})

	t.Run("GetFloatValue_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NumericType
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetFloatValue() // Should return zero value
	})

	t.Run("GetInt32Value", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		var expected *int
		obj.Int32Value = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetInt32Value(), "getter should return the property value")
	})

	t.Run("GetInt32Value_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		obj.Int32Value = nil

		// Act & Assert
		assert.Nil(t, obj.GetInt32Value(), "getter should return nil when property is nil")
	})

	t.Run("GetInt32Value_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NumericType
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetInt32Value() // Should return zero value
	})

	t.Run("GetInt64Value", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		var expected *string
		obj.Int64Value = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetInt64Value(), "getter should return the property value")
	})

	t.Run("GetInt64Value_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		obj.Int64Value = nil

		// Act & Assert
		assert.Nil(t, obj.GetInt64Value(), "getter should return nil when property is nil")
	})

	t.Run("GetInt64Value_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NumericType
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetInt64Value() // Should return zero value
	})

	t.Run("GetUint32Value", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		var expected *int
		obj.Uint32Value = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetUint32Value(), "getter should return the property value")
	})

	t.Run("GetUint32Value_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		obj.Uint32Value = nil

		// Act & Assert
		assert.Nil(t, obj.GetUint32Value(), "getter should return nil when property is nil")
	})

	t.Run("GetUint32Value_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NumericType
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetUint32Value() // Should return zero value
	})

	t.Run("GetUint64Value", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		var expected *string
		obj.Uint64Value = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetUint64Value(), "getter should return the property value")
	})

	t.Run("GetUint64Value_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		obj.Uint64Value = nil

		// Act & Assert
		assert.Nil(t, obj.GetUint64Value(), "getter should return nil when property is nil")
	})

	t.Run("GetUint64Value_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NumericType
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetUint64Value() // Should return zero value
	})

}

func TestSettersMarkExplicitNumericType(t *testing.T) {
	t.Run("SetDoubleValue_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		var fernTestValueDoubleValue *float64

		// Act
		obj.SetDoubleValue(fernTestValueDoubleValue)

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

	t.Run("SetFloatValue_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		var fernTestValueFloatValue *float64

		// Act
		obj.SetFloatValue(fernTestValueFloatValue)

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

	t.Run("SetInt32Value_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		var fernTestValueInt32Value *int

		// Act
		obj.SetInt32Value(fernTestValueInt32Value)

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

	t.Run("SetInt64Value_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		var fernTestValueInt64Value *string

		// Act
		obj.SetInt64Value(fernTestValueInt64Value)

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

	t.Run("SetUint32Value_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		var fernTestValueUint32Value *int

		// Act
		obj.SetUint32Value(fernTestValueUint32Value)

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

	t.Run("SetUint64Value_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}
		var fernTestValueUint64Value *string

		// Act
		obj.SetUint64Value(fernTestValueUint64Value)

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

func TestSettersOrOperation(t *testing.T) {
	t.Run("SetPredicateSet", func(t *testing.T) {
		obj := &OrOperation{}
		var fernTestValuePredicateSet *PredicateSet
		obj.SetPredicateSet(fernTestValuePredicateSet)
		assert.Equal(t, fernTestValuePredicateSet, obj.PredicateSet)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatementSet", func(t *testing.T) {
		obj := &OrOperation{}
		var fernTestValueStatementSet *StatementSet
		obj.SetStatementSet(fernTestValueStatementSet)
		assert.Equal(t, fernTestValueStatementSet, obj.StatementSet)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersOrOperation(t *testing.T) {
	t.Run("GetPredicateSet", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OrOperation{}
		var expected *PredicateSet
		obj.PredicateSet = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPredicateSet(), "getter should return the property value")
	})

	t.Run("GetPredicateSet_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OrOperation{}
		obj.PredicateSet = nil

		// Act & Assert
		assert.Nil(t, obj.GetPredicateSet(), "getter should return nil when property is nil")
	})

	t.Run("GetPredicateSet_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *OrOperation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPredicateSet() // Should return zero value
	})

	t.Run("GetStatementSet", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OrOperation{}
		var expected *StatementSet
		obj.StatementSet = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatementSet(), "getter should return the property value")
	})

	t.Run("GetStatementSet_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OrOperation{}
		obj.StatementSet = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatementSet(), "getter should return nil when property is nil")
	})

	t.Run("GetStatementSet_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *OrOperation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatementSet() // Should return zero value
	})

}

func TestSettersMarkExplicitOrOperation(t *testing.T) {
	t.Run("SetPredicateSet_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OrOperation{}
		var fernTestValuePredicateSet *PredicateSet

		// Act
		obj.SetPredicateSet(fernTestValuePredicateSet)

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

	t.Run("SetStatementSet_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OrOperation{}
		var fernTestValueStatementSet *StatementSet

		// Act
		obj.SetStatementSet(fernTestValueStatementSet)

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

func TestSettersPositionType(t *testing.T) {
	t.Run("SetValue", func(t *testing.T) {
		obj := &PositionType{}
		var fernTestValueValue *Position
		obj.SetValue(fernTestValueValue)
		assert.Equal(t, fernTestValueValue, obj.Value)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersPositionType(t *testing.T) {
	t.Run("GetValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PositionType{}
		var expected *Position
		obj.Value = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetValue(), "getter should return the property value")
	})

	t.Run("GetValue_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PositionType{}
		obj.Value = nil

		// Act & Assert
		assert.Nil(t, obj.GetValue(), "getter should return nil when property is nil")
	})

	t.Run("GetValue_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PositionType
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetValue() // Should return zero value
	})

}

func TestSettersMarkExplicitPositionType(t *testing.T) {
	t.Run("SetValue_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PositionType{}
		var fernTestValueValue *Position

		// Act
		obj.SetValue(fernTestValueValue)

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

func TestSettersPredicate(t *testing.T) {
	t.Run("SetFieldPath", func(t *testing.T) {
		obj := &Predicate{}
		var fernTestValueFieldPath *string
		obj.SetFieldPath(fernTestValueFieldPath)
		assert.Equal(t, fernTestValueFieldPath, obj.FieldPath)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetValue", func(t *testing.T) {
		obj := &Predicate{}
		var fernTestValueValue *Value
		obj.SetValue(fernTestValueValue)
		assert.Equal(t, fernTestValueValue, obj.Value)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetComparator", func(t *testing.T) {
		obj := &Predicate{}
		var fernTestValueComparator *PredicateComparator
		obj.SetComparator(fernTestValueComparator)
		assert.Equal(t, fernTestValueComparator, obj.Comparator)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersPredicate(t *testing.T) {
	t.Run("GetFieldPath", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Predicate{}
		var expected *string
		obj.FieldPath = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetFieldPath(), "getter should return the property value")
	})

	t.Run("GetFieldPath_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Predicate{}
		obj.FieldPath = nil

		// Act & Assert
		assert.Nil(t, obj.GetFieldPath(), "getter should return nil when property is nil")
	})

	t.Run("GetFieldPath_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Predicate
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetFieldPath() // Should return zero value
	})

	t.Run("GetValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Predicate{}
		var expected *Value
		obj.Value = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetValue(), "getter should return the property value")
	})

	t.Run("GetValue_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Predicate{}
		obj.Value = nil

		// Act & Assert
		assert.Nil(t, obj.GetValue(), "getter should return nil when property is nil")
	})

	t.Run("GetValue_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Predicate
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetValue() // Should return zero value
	})

	t.Run("GetComparator", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Predicate{}
		var expected *PredicateComparator
		obj.Comparator = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetComparator(), "getter should return the property value")
	})

	t.Run("GetComparator_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Predicate{}
		obj.Comparator = nil

		// Act & Assert
		assert.Nil(t, obj.GetComparator(), "getter should return nil when property is nil")
	})

	t.Run("GetComparator_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Predicate
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetComparator() // Should return zero value
	})

}

func TestSettersMarkExplicitPredicate(t *testing.T) {
	t.Run("SetFieldPath_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Predicate{}
		var fernTestValueFieldPath *string

		// Act
		obj.SetFieldPath(fernTestValueFieldPath)

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

	t.Run("SetValue_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Predicate{}
		var fernTestValueValue *Value

		// Act
		obj.SetValue(fernTestValueValue)

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

	t.Run("SetComparator_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Predicate{}
		var fernTestValueComparator *PredicateComparator

		// Act
		obj.SetComparator(fernTestValueComparator)

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

func TestSettersPredicateSet(t *testing.T) {
	t.Run("SetPredicates", func(t *testing.T) {
		obj := &PredicateSet{}
		var fernTestValuePredicates []*Predicate
		obj.SetPredicates(fernTestValuePredicates)
		assert.Equal(t, fernTestValuePredicates, obj.Predicates)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersPredicateSet(t *testing.T) {
	t.Run("GetPredicates", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PredicateSet{}
		var expected []*Predicate
		obj.Predicates = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPredicates(), "getter should return the property value")
	})

	t.Run("GetPredicates_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PredicateSet{}
		obj.Predicates = nil

		// Act & Assert
		assert.Nil(t, obj.GetPredicates(), "getter should return nil when property is nil")
	})

	t.Run("GetPredicates_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PredicateSet
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPredicates() // Should return zero value
	})

}

func TestSettersMarkExplicitPredicateSet(t *testing.T) {
	t.Run("SetPredicates_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PredicateSet{}
		var fernTestValuePredicates []*Predicate

		// Act
		obj.SetPredicates(fernTestValuePredicates)

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

func TestSettersRangeType(t *testing.T) {
	t.Run("SetStart", func(t *testing.T) {
		obj := &RangeType{}
		var fernTestValueStart *NumericType
		obj.SetStart(fernTestValueStart)
		assert.Equal(t, fernTestValueStart, obj.Start)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEnd", func(t *testing.T) {
		obj := &RangeType{}
		var fernTestValueEnd *NumericType
		obj.SetEnd(fernTestValueEnd)
		assert.Equal(t, fernTestValueEnd, obj.End)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersRangeType(t *testing.T) {
	t.Run("GetStart", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RangeType{}
		var expected *NumericType
		obj.Start = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStart(), "getter should return the property value")
	})

	t.Run("GetStart_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RangeType{}
		obj.Start = nil

		// Act & Assert
		assert.Nil(t, obj.GetStart(), "getter should return nil when property is nil")
	})

	t.Run("GetStart_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RangeType
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStart() // Should return zero value
	})

	t.Run("GetEnd", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RangeType{}
		var expected *NumericType
		obj.End = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEnd(), "getter should return the property value")
	})

	t.Run("GetEnd_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RangeType{}
		obj.End = nil

		// Act & Assert
		assert.Nil(t, obj.GetEnd(), "getter should return nil when property is nil")
	})

	t.Run("GetEnd_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RangeType
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEnd() // Should return zero value
	})

}

func TestSettersMarkExplicitRangeType(t *testing.T) {
	t.Run("SetStart_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RangeType{}
		var fernTestValueStart *NumericType

		// Act
		obj.SetStart(fernTestValueStart)

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

	t.Run("SetEnd_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RangeType{}
		var fernTestValueEnd *NumericType

		// Act
		obj.SetEnd(fernTestValueEnd)

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

func TestSettersStatement(t *testing.T) {
	t.Run("SetAnd", func(t *testing.T) {
		obj := &Statement{}
		var fernTestValueAnd *AndOperation
		obj.SetAnd(fernTestValueAnd)
		assert.Equal(t, fernTestValueAnd, obj.And)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetOr", func(t *testing.T) {
		obj := &Statement{}
		var fernTestValueOr *OrOperation
		obj.SetOr(fernTestValueOr)
		assert.Equal(t, fernTestValueOr, obj.Or)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetNot", func(t *testing.T) {
		obj := &Statement{}
		var fernTestValueNot *NotOperation
		obj.SetNot(fernTestValueNot)
		assert.Equal(t, fernTestValueNot, obj.Not)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetList", func(t *testing.T) {
		obj := &Statement{}
		var fernTestValueList *ListOperation
		obj.SetList(fernTestValueList)
		assert.Equal(t, fernTestValueList, obj.List)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetPredicate", func(t *testing.T) {
		obj := &Statement{}
		var fernTestValuePredicate *Predicate
		obj.SetPredicate(fernTestValuePredicate)
		assert.Equal(t, fernTestValuePredicate, obj.Predicate)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersStatement(t *testing.T) {
	t.Run("GetAnd", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Statement{}
		var expected *AndOperation
		obj.And = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAnd(), "getter should return the property value")
	})

	t.Run("GetAnd_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Statement{}
		obj.And = nil

		// Act & Assert
		assert.Nil(t, obj.GetAnd(), "getter should return nil when property is nil")
	})

	t.Run("GetAnd_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Statement
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAnd() // Should return zero value
	})

	t.Run("GetOr", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Statement{}
		var expected *OrOperation
		obj.Or = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetOr(), "getter should return the property value")
	})

	t.Run("GetOr_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Statement{}
		obj.Or = nil

		// Act & Assert
		assert.Nil(t, obj.GetOr(), "getter should return nil when property is nil")
	})

	t.Run("GetOr_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Statement
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetOr() // Should return zero value
	})

	t.Run("GetNot", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Statement{}
		var expected *NotOperation
		obj.Not = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetNot(), "getter should return the property value")
	})

	t.Run("GetNot_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Statement{}
		obj.Not = nil

		// Act & Assert
		assert.Nil(t, obj.GetNot(), "getter should return nil when property is nil")
	})

	t.Run("GetNot_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Statement
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetNot() // Should return zero value
	})

	t.Run("GetList", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Statement{}
		var expected *ListOperation
		obj.List = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetList(), "getter should return the property value")
	})

	t.Run("GetList_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Statement{}
		obj.List = nil

		// Act & Assert
		assert.Nil(t, obj.GetList(), "getter should return nil when property is nil")
	})

	t.Run("GetList_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Statement
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetList() // Should return zero value
	})

	t.Run("GetPredicate", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Statement{}
		var expected *Predicate
		obj.Predicate = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPredicate(), "getter should return the property value")
	})

	t.Run("GetPredicate_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Statement{}
		obj.Predicate = nil

		// Act & Assert
		assert.Nil(t, obj.GetPredicate(), "getter should return nil when property is nil")
	})

	t.Run("GetPredicate_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Statement
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPredicate() // Should return zero value
	})

}

func TestSettersMarkExplicitStatement(t *testing.T) {
	t.Run("SetAnd_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Statement{}
		var fernTestValueAnd *AndOperation

		// Act
		obj.SetAnd(fernTestValueAnd)

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

	t.Run("SetOr_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Statement{}
		var fernTestValueOr *OrOperation

		// Act
		obj.SetOr(fernTestValueOr)

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

	t.Run("SetNot_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Statement{}
		var fernTestValueNot *NotOperation

		// Act
		obj.SetNot(fernTestValueNot)

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

	t.Run("SetList_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Statement{}
		var fernTestValueList *ListOperation

		// Act
		obj.SetList(fernTestValueList)

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

	t.Run("SetPredicate_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Statement{}
		var fernTestValuePredicate *Predicate

		// Act
		obj.SetPredicate(fernTestValuePredicate)

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

func TestSettersStatementSet(t *testing.T) {
	t.Run("SetStatements", func(t *testing.T) {
		obj := &StatementSet{}
		var fernTestValueStatements []*Statement
		obj.SetStatements(fernTestValueStatements)
		assert.Equal(t, fernTestValueStatements, obj.Statements)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersStatementSet(t *testing.T) {
	t.Run("GetStatements", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StatementSet{}
		var expected []*Statement
		obj.Statements = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatements(), "getter should return the property value")
	})

	t.Run("GetStatements_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StatementSet{}
		obj.Statements = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatements(), "getter should return nil when property is nil")
	})

	t.Run("GetStatements_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *StatementSet
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatements() // Should return zero value
	})

}

func TestSettersMarkExplicitStatementSet(t *testing.T) {
	t.Run("SetStatements_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StatementSet{}
		var fernTestValueStatements []*Statement

		// Act
		obj.SetStatements(fernTestValueStatements)

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

func TestSettersStringType(t *testing.T) {
	t.Run("SetValue", func(t *testing.T) {
		obj := &StringType{}
		var fernTestValueValue *string
		obj.SetValue(fernTestValueValue)
		assert.Equal(t, fernTestValueValue, obj.Value)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersStringType(t *testing.T) {
	t.Run("GetValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StringType{}
		var expected *string
		obj.Value = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetValue(), "getter should return the property value")
	})

	t.Run("GetValue_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StringType{}
		obj.Value = nil

		// Act & Assert
		assert.Nil(t, obj.GetValue(), "getter should return nil when property is nil")
	})

	t.Run("GetValue_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *StringType
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetValue() // Should return zero value
	})

}

func TestSettersMarkExplicitStringType(t *testing.T) {
	t.Run("SetValue_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StringType{}
		var fernTestValueValue *string

		// Act
		obj.SetValue(fernTestValueValue)

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

func TestSettersTimestampType(t *testing.T) {
	t.Run("SetValue", func(t *testing.T) {
		obj := &TimestampType{}
		var fernTestValueValue *time.Time
		obj.SetValue(fernTestValueValue)
		assert.Equal(t, fernTestValueValue, obj.Value)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersTimestampType(t *testing.T) {
	t.Run("GetValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TimestampType{}
		var expected *time.Time
		obj.Value = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetValue(), "getter should return the property value")
	})

	t.Run("GetValue_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TimestampType{}
		obj.Value = nil

		// Act & Assert
		assert.Nil(t, obj.GetValue(), "getter should return nil when property is nil")
	})

	t.Run("GetValue_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TimestampType
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetValue() // Should return zero value
	})

}

func TestSettersMarkExplicitTimestampType(t *testing.T) {
	t.Run("SetValue_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TimestampType{}
		var fernTestValueValue *time.Time

		// Act
		obj.SetValue(fernTestValueValue)

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

func TestSettersValue(t *testing.T) {
	t.Run("SetBooleanType", func(t *testing.T) {
		obj := &Value{}
		var fernTestValueBooleanType *BooleanType
		obj.SetBooleanType(fernTestValueBooleanType)
		assert.Equal(t, fernTestValueBooleanType, obj.BooleanType)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetNumericType", func(t *testing.T) {
		obj := &Value{}
		var fernTestValueNumericType *NumericType
		obj.SetNumericType(fernTestValueNumericType)
		assert.Equal(t, fernTestValueNumericType, obj.NumericType)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStringType", func(t *testing.T) {
		obj := &Value{}
		var fernTestValueStringType *StringType
		obj.SetStringType(fernTestValueStringType)
		assert.Equal(t, fernTestValueStringType, obj.StringType)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEnumType", func(t *testing.T) {
		obj := &Value{}
		var fernTestValueEnumType *EnumType
		obj.SetEnumType(fernTestValueEnumType)
		assert.Equal(t, fernTestValueEnumType, obj.EnumType)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTimestampType", func(t *testing.T) {
		obj := &Value{}
		var fernTestValueTimestampType *TimestampType
		obj.SetTimestampType(fernTestValueTimestampType)
		assert.Equal(t, fernTestValueTimestampType, obj.TimestampType)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetBoundedShapeType", func(t *testing.T) {
		obj := &Value{}
		var fernTestValueBoundedShapeType *BoundedShapeType
		obj.SetBoundedShapeType(fernTestValueBoundedShapeType)
		assert.Equal(t, fernTestValueBoundedShapeType, obj.BoundedShapeType)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetPositionType", func(t *testing.T) {
		obj := &Value{}
		var fernTestValuePositionType *PositionType
		obj.SetPositionType(fernTestValuePositionType)
		assert.Equal(t, fernTestValuePositionType, obj.PositionType)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetHeadingType", func(t *testing.T) {
		obj := &Value{}
		var fernTestValueHeadingType *HeadingType
		obj.SetHeadingType(fernTestValueHeadingType)
		assert.Equal(t, fernTestValueHeadingType, obj.HeadingType)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetListType", func(t *testing.T) {
		obj := &Value{}
		var fernTestValueListType *ListType
		obj.SetListType(fernTestValueListType)
		assert.Equal(t, fernTestValueListType, obj.ListType)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRangeType", func(t *testing.T) {
		obj := &Value{}
		var fernTestValueRangeType *RangeType
		obj.SetRangeType(fernTestValueRangeType)
		assert.Equal(t, fernTestValueRangeType, obj.RangeType)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersValue(t *testing.T) {
	t.Run("GetBooleanType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var expected *BooleanType
		obj.BooleanType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetBooleanType(), "getter should return the property value")
	})

	t.Run("GetBooleanType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		obj.BooleanType = nil

		// Act & Assert
		assert.Nil(t, obj.GetBooleanType(), "getter should return nil when property is nil")
	})

	t.Run("GetBooleanType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Value
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetBooleanType() // Should return zero value
	})

	t.Run("GetNumericType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var expected *NumericType
		obj.NumericType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetNumericType(), "getter should return the property value")
	})

	t.Run("GetNumericType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		obj.NumericType = nil

		// Act & Assert
		assert.Nil(t, obj.GetNumericType(), "getter should return nil when property is nil")
	})

	t.Run("GetNumericType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Value
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetNumericType() // Should return zero value
	})

	t.Run("GetStringType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var expected *StringType
		obj.StringType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStringType(), "getter should return the property value")
	})

	t.Run("GetStringType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		obj.StringType = nil

		// Act & Assert
		assert.Nil(t, obj.GetStringType(), "getter should return nil when property is nil")
	})

	t.Run("GetStringType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Value
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStringType() // Should return zero value
	})

	t.Run("GetEnumType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var expected *EnumType
		obj.EnumType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEnumType(), "getter should return the property value")
	})

	t.Run("GetEnumType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		obj.EnumType = nil

		// Act & Assert
		assert.Nil(t, obj.GetEnumType(), "getter should return nil when property is nil")
	})

	t.Run("GetEnumType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Value
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEnumType() // Should return zero value
	})

	t.Run("GetTimestampType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var expected *TimestampType
		obj.TimestampType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTimestampType(), "getter should return the property value")
	})

	t.Run("GetTimestampType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		obj.TimestampType = nil

		// Act & Assert
		assert.Nil(t, obj.GetTimestampType(), "getter should return nil when property is nil")
	})

	t.Run("GetTimestampType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Value
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTimestampType() // Should return zero value
	})

	t.Run("GetBoundedShapeType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var expected *BoundedShapeType
		obj.BoundedShapeType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetBoundedShapeType(), "getter should return the property value")
	})

	t.Run("GetBoundedShapeType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		obj.BoundedShapeType = nil

		// Act & Assert
		assert.Nil(t, obj.GetBoundedShapeType(), "getter should return nil when property is nil")
	})

	t.Run("GetBoundedShapeType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Value
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetBoundedShapeType() // Should return zero value
	})

	t.Run("GetPositionType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var expected *PositionType
		obj.PositionType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPositionType(), "getter should return the property value")
	})

	t.Run("GetPositionType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		obj.PositionType = nil

		// Act & Assert
		assert.Nil(t, obj.GetPositionType(), "getter should return nil when property is nil")
	})

	t.Run("GetPositionType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Value
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPositionType() // Should return zero value
	})

	t.Run("GetHeadingType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var expected *HeadingType
		obj.HeadingType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHeadingType(), "getter should return the property value")
	})

	t.Run("GetHeadingType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		obj.HeadingType = nil

		// Act & Assert
		assert.Nil(t, obj.GetHeadingType(), "getter should return nil when property is nil")
	})

	t.Run("GetHeadingType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Value
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetHeadingType() // Should return zero value
	})

	t.Run("GetListType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var expected *ListType
		obj.ListType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetListType(), "getter should return the property value")
	})

	t.Run("GetListType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		obj.ListType = nil

		// Act & Assert
		assert.Nil(t, obj.GetListType(), "getter should return nil when property is nil")
	})

	t.Run("GetListType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Value
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetListType() // Should return zero value
	})

	t.Run("GetRangeType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var expected *RangeType
		obj.RangeType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRangeType(), "getter should return the property value")
	})

	t.Run("GetRangeType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		obj.RangeType = nil

		// Act & Assert
		assert.Nil(t, obj.GetRangeType(), "getter should return nil when property is nil")
	})

	t.Run("GetRangeType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Value
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRangeType() // Should return zero value
	})

}

func TestSettersMarkExplicitValue(t *testing.T) {
	t.Run("SetBooleanType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var fernTestValueBooleanType *BooleanType

		// Act
		obj.SetBooleanType(fernTestValueBooleanType)

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

	t.Run("SetNumericType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var fernTestValueNumericType *NumericType

		// Act
		obj.SetNumericType(fernTestValueNumericType)

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

	t.Run("SetStringType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var fernTestValueStringType *StringType

		// Act
		obj.SetStringType(fernTestValueStringType)

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

	t.Run("SetEnumType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var fernTestValueEnumType *EnumType

		// Act
		obj.SetEnumType(fernTestValueEnumType)

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

	t.Run("SetTimestampType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var fernTestValueTimestampType *TimestampType

		// Act
		obj.SetTimestampType(fernTestValueTimestampType)

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

	t.Run("SetBoundedShapeType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var fernTestValueBoundedShapeType *BoundedShapeType

		// Act
		obj.SetBoundedShapeType(fernTestValueBoundedShapeType)

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

	t.Run("SetPositionType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var fernTestValuePositionType *PositionType

		// Act
		obj.SetPositionType(fernTestValuePositionType)

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

	t.Run("SetHeadingType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var fernTestValueHeadingType *HeadingType

		// Act
		obj.SetHeadingType(fernTestValueHeadingType)

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

	t.Run("SetListType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var fernTestValueListType *ListType

		// Act
		obj.SetListType(fernTestValueListType)

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

	t.Run("SetRangeType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}
		var fernTestValueRangeType *RangeType

		// Act
		obj.SetRangeType(fernTestValueRangeType)

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

func TestGettersStreamEntitiesResponse(t *testing.T) {
	t.Run("GetEvent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StreamEntitiesResponse{}
		var expected string
		obj.Event = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEvent(), "getter should return the property value")
	})

	t.Run("GetEvent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *StreamEntitiesResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEvent() // Should return zero value
	})

	t.Run("GetHeartbeat", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StreamEntitiesResponse{}
		var expected *EntityStreamHeartbeat
		obj.Heartbeat = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHeartbeat(), "getter should return the property value")
	})

	t.Run("GetHeartbeat_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StreamEntitiesResponse{}
		obj.Heartbeat = nil

		// Act & Assert
		assert.Nil(t, obj.GetHeartbeat(), "getter should return nil when property is nil")
	})

	t.Run("GetHeartbeat_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *StreamEntitiesResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetHeartbeat() // Should return zero value
	})

	t.Run("GetEntity", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StreamEntitiesResponse{}
		var expected *EntityStreamEvent
		obj.Entity = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEntity(), "getter should return the property value")
	})

	t.Run("GetEntity_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StreamEntitiesResponse{}
		obj.Entity = nil

		// Act & Assert
		assert.Nil(t, obj.GetEntity(), "getter should return nil when property is nil")
	})

	t.Run("GetEntity_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *StreamEntitiesResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEntity() // Should return zero value
	})

}

func TestJSONMarshalingAndOperation(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AndOperation{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled AndOperation
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj AndOperation
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj AndOperation
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingBooleanType(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BooleanType{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled BooleanType
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj BooleanType
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj BooleanType
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingBoundedShapeType(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundedShapeType{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled BoundedShapeType
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj BoundedShapeType
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj BoundedShapeType
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingEntityEvent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEvent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled EntityEvent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj EntityEvent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj EntityEvent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingEntityEventResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityEventResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled EntityEventResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj EntityEventResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj EntityEventResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingEntityStreamEvent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamEvent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled EntityStreamEvent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj EntityStreamEvent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj EntityStreamEvent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingEntityStreamHeartbeat(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EntityStreamHeartbeat{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled EntityStreamHeartbeat
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj EntityStreamHeartbeat
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj EntityStreamHeartbeat
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingEnumType(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EnumType{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled EnumType
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj EnumType
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj EnumType
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingHeadingType(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &HeadingType{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled HeadingType
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj HeadingType
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj HeadingType
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListOperation(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListOperation{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListOperation
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListOperation
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListOperation
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListType(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListType{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListType
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListType
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListType
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingNotOperation(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotOperation{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled NotOperation
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj NotOperation
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj NotOperation
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingNumericType(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NumericType{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled NumericType
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj NumericType
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj NumericType
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingOrOperation(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OrOperation{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled OrOperation
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj OrOperation
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj OrOperation
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingPositionType(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PositionType{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled PositionType
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj PositionType
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj PositionType
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingPredicate(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Predicate{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled Predicate
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj Predicate
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj Predicate
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingPredicateSet(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PredicateSet{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled PredicateSet
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj PredicateSet
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj PredicateSet
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingRangeType(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RangeType{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled RangeType
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj RangeType
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj RangeType
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingStatement(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Statement{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled Statement
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj Statement
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj Statement
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingStatementSet(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StatementSet{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled StatementSet
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj StatementSet
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj StatementSet
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingStringType(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StringType{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled StringType
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj StringType
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj StringType
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingTimestampType(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TimestampType{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled TimestampType
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj TimestampType
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj TimestampType
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingValue(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Value{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled Value
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj Value
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj Value
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestStringAndOperation(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &AndOperation{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AndOperation
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringBooleanType(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &BooleanType{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BooleanType
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringBoundedShapeType(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &BoundedShapeType{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BoundedShapeType
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringEntityEvent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &EntityEvent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EntityEvent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringEntityEventResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &EntityEventResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EntityEventResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringEntityStreamEvent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &EntityStreamEvent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EntityStreamEvent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringEntityStreamHeartbeat(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &EntityStreamHeartbeat{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EntityStreamHeartbeat
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringEnumType(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &EnumType{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EnumType
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringHeadingType(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &HeadingType{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *HeadingType
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListOperation(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListOperation{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListOperation
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListType(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListType{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListType
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringNotOperation(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &NotOperation{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotOperation
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringNumericType(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &NumericType{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NumericType
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringOrOperation(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &OrOperation{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *OrOperation
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringPositionType(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &PositionType{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PositionType
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringPredicate(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &Predicate{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Predicate
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringPredicateSet(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &PredicateSet{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PredicateSet
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringRangeType(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &RangeType{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RangeType
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringStatement(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &Statement{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Statement
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringStatementSet(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &StatementSet{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *StatementSet
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringStringType(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &StringType{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *StringType
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringTimestampType(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &TimestampType{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TimestampType
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringValue(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &Value{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Value
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestEnumEntityEventEventType(t *testing.T) {
	t.Run("NewFromString_EVENT_TYPE_INVALID", func(t *testing.T) {
		t.Parallel()
		val, err := NewEntityEventEventTypeFromString("EVENT_TYPE_INVALID")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, EntityEventEventType("EVENT_TYPE_INVALID"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_EVENT_TYPE_CREATED", func(t *testing.T) {
		t.Parallel()
		val, err := NewEntityEventEventTypeFromString("EVENT_TYPE_CREATED")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, EntityEventEventType("EVENT_TYPE_CREATED"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_EVENT_TYPE_UPDATE", func(t *testing.T) {
		t.Parallel()
		val, err := NewEntityEventEventTypeFromString("EVENT_TYPE_UPDATE")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, EntityEventEventType("EVENT_TYPE_UPDATE"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_EVENT_TYPE_DELETED", func(t *testing.T) {
		t.Parallel()
		val, err := NewEntityEventEventTypeFromString("EVENT_TYPE_DELETED")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, EntityEventEventType("EVENT_TYPE_DELETED"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_EVENT_TYPE_PREEXISTING", func(t *testing.T) {
		t.Parallel()
		val, err := NewEntityEventEventTypeFromString("EVENT_TYPE_PREEXISTING")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, EntityEventEventType("EVENT_TYPE_PREEXISTING"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_EVENT_TYPE_POST_EXPIRY_OVERRIDE", func(t *testing.T) {
		t.Parallel()
		val, err := NewEntityEventEventTypeFromString("EVENT_TYPE_POST_EXPIRY_OVERRIDE")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, EntityEventEventType("EVENT_TYPE_POST_EXPIRY_OVERRIDE"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewEntityEventEventTypeFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewEntityEventEventTypeFromString("EVENT_TYPE_INVALID")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumListOperationListComparator(t *testing.T) {
	t.Run("NewFromString_LIST_COMPARATOR_INVALID", func(t *testing.T) {
		t.Parallel()
		val, err := NewListOperationListComparatorFromString("LIST_COMPARATOR_INVALID")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListOperationListComparator("LIST_COMPARATOR_INVALID"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_LIST_COMPARATOR_ANY_OF", func(t *testing.T) {
		t.Parallel()
		val, err := NewListOperationListComparatorFromString("LIST_COMPARATOR_ANY_OF")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListOperationListComparator("LIST_COMPARATOR_ANY_OF"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewListOperationListComparatorFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewListOperationListComparatorFromString("LIST_COMPARATOR_INVALID")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumPredicateComparator(t *testing.T) {
	t.Run("NewFromString_COMPARATOR_INVALID", func(t *testing.T) {
		t.Parallel()
		val, err := NewPredicateComparatorFromString("COMPARATOR_INVALID")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PredicateComparator("COMPARATOR_INVALID"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_COMPARATOR_MATCH_ALL", func(t *testing.T) {
		t.Parallel()
		val, err := NewPredicateComparatorFromString("COMPARATOR_MATCH_ALL")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PredicateComparator("COMPARATOR_MATCH_ALL"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_COMPARATOR_EQUALITY", func(t *testing.T) {
		t.Parallel()
		val, err := NewPredicateComparatorFromString("COMPARATOR_EQUALITY")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PredicateComparator("COMPARATOR_EQUALITY"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_COMPARATOR_IN", func(t *testing.T) {
		t.Parallel()
		val, err := NewPredicateComparatorFromString("COMPARATOR_IN")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PredicateComparator("COMPARATOR_IN"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_COMPARATOR_LESS_THAN", func(t *testing.T) {
		t.Parallel()
		val, err := NewPredicateComparatorFromString("COMPARATOR_LESS_THAN")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PredicateComparator("COMPARATOR_LESS_THAN"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_COMPARATOR_GREATER_THAN", func(t *testing.T) {
		t.Parallel()
		val, err := NewPredicateComparatorFromString("COMPARATOR_GREATER_THAN")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PredicateComparator("COMPARATOR_GREATER_THAN"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_COMPARATOR_LESS_THAN_EQUAL_TO", func(t *testing.T) {
		t.Parallel()
		val, err := NewPredicateComparatorFromString("COMPARATOR_LESS_THAN_EQUAL_TO")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PredicateComparator("COMPARATOR_LESS_THAN_EQUAL_TO"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_COMPARATOR_GREATER_THAN_EQUAL_TO", func(t *testing.T) {
		t.Parallel()
		val, err := NewPredicateComparatorFromString("COMPARATOR_GREATER_THAN_EQUAL_TO")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PredicateComparator("COMPARATOR_GREATER_THAN_EQUAL_TO"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_COMPARATOR_WITHIN", func(t *testing.T) {
		t.Parallel()
		val, err := NewPredicateComparatorFromString("COMPARATOR_WITHIN")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PredicateComparator("COMPARATOR_WITHIN"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_COMPARATOR_EXISTS", func(t *testing.T) {
		t.Parallel()
		val, err := NewPredicateComparatorFromString("COMPARATOR_EXISTS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PredicateComparator("COMPARATOR_EXISTS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_COMPARATOR_CASE_INSENSITIVE_EQUALITY", func(t *testing.T) {
		t.Parallel()
		val, err := NewPredicateComparatorFromString("COMPARATOR_CASE_INSENSITIVE_EQUALITY")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PredicateComparator("COMPARATOR_CASE_INSENSITIVE_EQUALITY"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_COMPARATOR_CASE_INSENSITIVE_EQUALITY_IN", func(t *testing.T) {
		t.Parallel()
		val, err := NewPredicateComparatorFromString("COMPARATOR_CASE_INSENSITIVE_EQUALITY_IN")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PredicateComparator("COMPARATOR_CASE_INSENSITIVE_EQUALITY_IN"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_COMPARATOR_RANGE_CLOSED", func(t *testing.T) {
		t.Parallel()
		val, err := NewPredicateComparatorFromString("COMPARATOR_RANGE_CLOSED")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PredicateComparator("COMPARATOR_RANGE_CLOSED"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewPredicateComparatorFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewPredicateComparatorFromString("COMPARATOR_INVALID")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestExtraPropertiesAndOperation(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &AndOperation{}
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
		var obj *AndOperation
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesBooleanType(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &BooleanType{}
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
		var obj *BooleanType
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesBoundedShapeType(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &BoundedShapeType{}
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
		var obj *BoundedShapeType
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesEntityEvent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &EntityEvent{}
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
		var obj *EntityEvent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesEntityEventResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &EntityEventResponse{}
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
		var obj *EntityEventResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesEntityStreamEvent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &EntityStreamEvent{}
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
		var obj *EntityStreamEvent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesEntityStreamHeartbeat(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &EntityStreamHeartbeat{}
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
		var obj *EntityStreamHeartbeat
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesEnumType(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &EnumType{}
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
		var obj *EnumType
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesHeadingType(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &HeadingType{}
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
		var obj *HeadingType
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListOperation(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListOperation{}
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
		var obj *ListOperation
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListType(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListType{}
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
		var obj *ListType
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesNotOperation(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &NotOperation{}
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
		var obj *NotOperation
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesNumericType(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &NumericType{}
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
		var obj *NumericType
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesOrOperation(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &OrOperation{}
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
		var obj *OrOperation
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesPositionType(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &PositionType{}
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
		var obj *PositionType
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesPredicate(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &Predicate{}
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
		var obj *Predicate
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesPredicateSet(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &PredicateSet{}
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
		var obj *PredicateSet
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesRangeType(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &RangeType{}
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
		var obj *RangeType
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesStatement(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &Statement{}
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
		var obj *Statement
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesStatementSet(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &StatementSet{}
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
		var obj *StatementSet
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesStringType(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &StringType{}
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
		var obj *StringType
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesTimestampType(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &TimestampType{}
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
		var obj *TimestampType
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesValue(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &Value{}
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
		var obj *Value
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}
