// Code generated from our API definition. DO NOT EDIT.

package Lattice

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
	time "time"
)

func TestSettersDeleteObjectRequest(t *testing.T) {
	t.Run("SetObjectPath", func(t *testing.T) {
		obj := &DeleteObjectRequest{}
		var fernTestValueObjectPath string
		obj.SetObjectPath(fernTestValueObjectPath)
		assert.Equal(t, fernTestValueObjectPath, obj.ObjectPath)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitDeleteObjectRequest(t *testing.T) {
	t.Run("SetObjectPath_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &DeleteObjectRequest{}
		var fernTestValueObjectPath string

		// Act
		obj.SetObjectPath(fernTestValueObjectPath)

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

func TestSettersGetObjectRequest(t *testing.T) {
	t.Run("SetAcceptEncoding", func(t *testing.T) {
		obj := &GetObjectRequest{}
		var fernTestValueAcceptEncoding *GetObjectRequestAcceptEncoding
		obj.SetAcceptEncoding(fernTestValueAcceptEncoding)
		assert.Equal(t, fernTestValueAcceptEncoding, obj.AcceptEncoding)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetPriority", func(t *testing.T) {
		obj := &GetObjectRequest{}
		var fernTestValuePriority *string
		obj.SetPriority(fernTestValuePriority)
		assert.Equal(t, fernTestValuePriority, obj.Priority)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetObjectPath", func(t *testing.T) {
		obj := &GetObjectRequest{}
		var fernTestValueObjectPath string
		obj.SetObjectPath(fernTestValueObjectPath)
		assert.Equal(t, fernTestValueObjectPath, obj.ObjectPath)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetObjectRequest(t *testing.T) {
	t.Run("SetAcceptEncoding_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetObjectRequest{}
		var fernTestValueAcceptEncoding *GetObjectRequestAcceptEncoding

		// Act
		obj.SetAcceptEncoding(fernTestValueAcceptEncoding)

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

	t.Run("SetPriority_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetObjectRequest{}
		var fernTestValuePriority *string

		// Act
		obj.SetPriority(fernTestValuePriority)

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

	t.Run("SetObjectPath_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetObjectRequest{}
		var fernTestValueObjectPath string

		// Act
		obj.SetObjectPath(fernTestValueObjectPath)

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

func TestSettersGetObjectMetadataRequest(t *testing.T) {
	t.Run("SetObjectPath", func(t *testing.T) {
		obj := &GetObjectMetadataRequest{}
		var fernTestValueObjectPath string
		obj.SetObjectPath(fernTestValueObjectPath)
		assert.Equal(t, fernTestValueObjectPath, obj.ObjectPath)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetObjectMetadataRequest(t *testing.T) {
	t.Run("SetObjectPath_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetObjectMetadataRequest{}
		var fernTestValueObjectPath string

		// Act
		obj.SetObjectPath(fernTestValueObjectPath)

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

func TestSettersListObjectsRequest(t *testing.T) {
	t.Run("SetPrefix", func(t *testing.T) {
		obj := &ListObjectsRequest{}
		var fernTestValuePrefix *string
		obj.SetPrefix(fernTestValuePrefix)
		assert.Equal(t, fernTestValuePrefix, obj.Prefix)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSinceTimestamp", func(t *testing.T) {
		obj := &ListObjectsRequest{}
		var fernTestValueSinceTimestamp *time.Time
		obj.SetSinceTimestamp(fernTestValueSinceTimestamp)
		assert.Equal(t, fernTestValueSinceTimestamp, obj.SinceTimestamp)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetPageToken", func(t *testing.T) {
		obj := &ListObjectsRequest{}
		var fernTestValuePageToken *string
		obj.SetPageToken(fernTestValuePageToken)
		assert.Equal(t, fernTestValuePageToken, obj.PageToken)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetAllObjectsInMesh", func(t *testing.T) {
		obj := &ListObjectsRequest{}
		var fernTestValueAllObjectsInMesh *bool
		obj.SetAllObjectsInMesh(fernTestValueAllObjectsInMesh)
		assert.Equal(t, fernTestValueAllObjectsInMesh, obj.AllObjectsInMesh)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMaxPageSize", func(t *testing.T) {
		obj := &ListObjectsRequest{}
		var fernTestValueMaxPageSize *int
		obj.SetMaxPageSize(fernTestValueMaxPageSize)
		assert.Equal(t, fernTestValueMaxPageSize, obj.MaxPageSize)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitListObjectsRequest(t *testing.T) {
	t.Run("SetPrefix_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListObjectsRequest{}
		var fernTestValuePrefix *string

		// Act
		obj.SetPrefix(fernTestValuePrefix)

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

	t.Run("SetSinceTimestamp_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListObjectsRequest{}
		var fernTestValueSinceTimestamp *time.Time

		// Act
		obj.SetSinceTimestamp(fernTestValueSinceTimestamp)

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

	t.Run("SetPageToken_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListObjectsRequest{}
		var fernTestValuePageToken *string

		// Act
		obj.SetPageToken(fernTestValuePageToken)

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

	t.Run("SetAllObjectsInMesh_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListObjectsRequest{}
		var fernTestValueAllObjectsInMesh *bool

		// Act
		obj.SetAllObjectsInMesh(fernTestValueAllObjectsInMesh)

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

	t.Run("SetMaxPageSize_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListObjectsRequest{}
		var fernTestValueMaxPageSize *int

		// Act
		obj.SetMaxPageSize(fernTestValueMaxPageSize)

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

func TestSettersContentIdentifier(t *testing.T) {
	t.Run("SetPath", func(t *testing.T) {
		obj := &ContentIdentifier{}
		var fernTestValuePath string
		obj.SetPath(fernTestValuePath)
		assert.Equal(t, fernTestValuePath, obj.Path)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetChecksum", func(t *testing.T) {
		obj := &ContentIdentifier{}
		var fernTestValueChecksum string
		obj.SetChecksum(fernTestValueChecksum)
		assert.Equal(t, fernTestValueChecksum, obj.Checksum)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersContentIdentifier(t *testing.T) {
	t.Run("GetPath", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ContentIdentifier{}
		var expected string
		obj.Path = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPath(), "getter should return the property value")
	})

	t.Run("GetPath_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ContentIdentifier
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPath() // Should return zero value
	})

	t.Run("GetChecksum", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ContentIdentifier{}
		var expected string
		obj.Checksum = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetChecksum(), "getter should return the property value")
	})

	t.Run("GetChecksum_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ContentIdentifier
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetChecksum() // Should return zero value
	})

}

func TestSettersMarkExplicitContentIdentifier(t *testing.T) {
	t.Run("SetPath_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ContentIdentifier{}
		var fernTestValuePath string

		// Act
		obj.SetPath(fernTestValuePath)

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

	t.Run("SetChecksum_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ContentIdentifier{}
		var fernTestValueChecksum string

		// Act
		obj.SetChecksum(fernTestValueChecksum)

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

func TestSettersListResponse(t *testing.T) {
	t.Run("SetPathMetadatas", func(t *testing.T) {
		obj := &ListResponse{}
		var fernTestValuePathMetadatas []*PathMetadata
		obj.SetPathMetadatas(fernTestValuePathMetadatas)
		assert.Equal(t, fernTestValuePathMetadatas, obj.PathMetadatas)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetNextPageToken", func(t *testing.T) {
		obj := &ListResponse{}
		var fernTestValueNextPageToken *string
		obj.SetNextPageToken(fernTestValueNextPageToken)
		assert.Equal(t, fernTestValueNextPageToken, obj.NextPageToken)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListResponse(t *testing.T) {
	t.Run("GetPathMetadatas", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListResponse{}
		var expected []*PathMetadata
		obj.PathMetadatas = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPathMetadatas(), "getter should return the property value")
	})

	t.Run("GetPathMetadatas_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListResponse{}
		obj.PathMetadatas = nil

		// Act & Assert
		assert.Nil(t, obj.GetPathMetadatas(), "getter should return nil when property is nil")
	})

	t.Run("GetPathMetadatas_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPathMetadatas() // Should return zero value
	})

	t.Run("GetNextPageToken", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListResponse{}
		var expected *string
		obj.NextPageToken = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetNextPageToken(), "getter should return the property value")
	})

	t.Run("GetNextPageToken_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListResponse{}
		obj.NextPageToken = nil

		// Act & Assert
		assert.Nil(t, obj.GetNextPageToken(), "getter should return nil when property is nil")
	})

	t.Run("GetNextPageToken_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetNextPageToken() // Should return zero value
	})

}

func TestSettersMarkExplicitListResponse(t *testing.T) {
	t.Run("SetPathMetadatas_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListResponse{}
		var fernTestValuePathMetadatas []*PathMetadata

		// Act
		obj.SetPathMetadatas(fernTestValuePathMetadatas)

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

	t.Run("SetNextPageToken_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListResponse{}
		var fernTestValueNextPageToken *string

		// Act
		obj.SetNextPageToken(fernTestValueNextPageToken)

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

func TestSettersPathMetadata(t *testing.T) {
	t.Run("SetContentIdentifier", func(t *testing.T) {
		obj := &PathMetadata{}
		var fernTestValueContentIdentifier *ContentIdentifier
		obj.SetContentIdentifier(fernTestValueContentIdentifier)
		assert.Equal(t, fernTestValueContentIdentifier, obj.ContentIdentifier)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSizeBytes", func(t *testing.T) {
		obj := &PathMetadata{}
		var fernTestValueSizeBytes int64
		obj.SetSizeBytes(fernTestValueSizeBytes)
		assert.Equal(t, fernTestValueSizeBytes, obj.SizeBytes)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetLastUpdatedAt", func(t *testing.T) {
		obj := &PathMetadata{}
		var fernTestValueLastUpdatedAt time.Time
		obj.SetLastUpdatedAt(fernTestValueLastUpdatedAt)
		assert.Equal(t, fernTestValueLastUpdatedAt, obj.LastUpdatedAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExpiryTime", func(t *testing.T) {
		obj := &PathMetadata{}
		var fernTestValueExpiryTime *time.Time
		obj.SetExpiryTime(fernTestValueExpiryTime)
		assert.Equal(t, fernTestValueExpiryTime, obj.ExpiryTime)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersPathMetadata(t *testing.T) {
	t.Run("GetContentIdentifier", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PathMetadata{}
		var expected *ContentIdentifier
		obj.ContentIdentifier = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetContentIdentifier(), "getter should return the property value")
	})

	t.Run("GetContentIdentifier_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PathMetadata{}
		obj.ContentIdentifier = nil

		// Act & Assert
		assert.Nil(t, obj.GetContentIdentifier(), "getter should return nil when property is nil")
	})

	t.Run("GetContentIdentifier_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PathMetadata
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetContentIdentifier() // Should return zero value
	})

	t.Run("GetSizeBytes", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PathMetadata{}
		var expected int64
		obj.SizeBytes = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSizeBytes(), "getter should return the property value")
	})

	t.Run("GetSizeBytes_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PathMetadata
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSizeBytes() // Should return zero value
	})

	t.Run("GetLastUpdatedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PathMetadata{}
		var expected time.Time
		obj.LastUpdatedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLastUpdatedAt(), "getter should return the property value")
	})

	t.Run("GetLastUpdatedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PathMetadata
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLastUpdatedAt() // Should return zero value
	})

	t.Run("GetExpiryTime", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PathMetadata{}
		var expected *time.Time
		obj.ExpiryTime = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetExpiryTime(), "getter should return the property value")
	})

	t.Run("GetExpiryTime_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PathMetadata{}
		obj.ExpiryTime = nil

		// Act & Assert
		assert.Nil(t, obj.GetExpiryTime(), "getter should return nil when property is nil")
	})

	t.Run("GetExpiryTime_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PathMetadata
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetExpiryTime() // Should return zero value
	})

}

func TestSettersMarkExplicitPathMetadata(t *testing.T) {
	t.Run("SetContentIdentifier_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PathMetadata{}
		var fernTestValueContentIdentifier *ContentIdentifier

		// Act
		obj.SetContentIdentifier(fernTestValueContentIdentifier)

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

	t.Run("SetSizeBytes_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PathMetadata{}
		var fernTestValueSizeBytes int64

		// Act
		obj.SetSizeBytes(fernTestValueSizeBytes)

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

	t.Run("SetLastUpdatedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PathMetadata{}
		var fernTestValueLastUpdatedAt time.Time

		// Act
		obj.SetLastUpdatedAt(fernTestValueLastUpdatedAt)

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

	t.Run("SetExpiryTime_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PathMetadata{}
		var fernTestValueExpiryTime *time.Time

		// Act
		obj.SetExpiryTime(fernTestValueExpiryTime)

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

func TestJSONMarshalingContentIdentifier(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ContentIdentifier{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ContentIdentifier
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ContentIdentifier
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ContentIdentifier
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingPathMetadata(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PathMetadata{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled PathMetadata
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj PathMetadata
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj PathMetadata
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestStringContentIdentifier(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ContentIdentifier{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ContentIdentifier
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringPathMetadata(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &PathMetadata{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PathMetadata
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestEnumGetObjectRequestAcceptEncoding(t *testing.T) {
	t.Run("NewFromString_identity", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetObjectRequestAcceptEncodingFromString("identity")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetObjectRequestAcceptEncoding("identity"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_zstd", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetObjectRequestAcceptEncodingFromString("zstd")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetObjectRequestAcceptEncoding("zstd"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewGetObjectRequestAcceptEncodingFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewGetObjectRequestAcceptEncodingFromString("identity")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumUploadObjectRequestDistributionMode(t *testing.T) {
	t.Run("NewFromString_force", func(t *testing.T) {
		t.Parallel()
		val, err := NewUploadObjectRequestDistributionModeFromString("force")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, UploadObjectRequestDistributionMode("force"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewUploadObjectRequestDistributionModeFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewUploadObjectRequestDistributionModeFromString("force")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestExtraPropertiesContentIdentifier(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ContentIdentifier{}
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
		var obj *ContentIdentifier
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListResponse{}
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
		var obj *ListResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesPathMetadata(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &PathMetadata{}
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
		var obj *PathMetadata
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}
