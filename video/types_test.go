// Code generated from our API definition. DO NOT EDIT.

package video

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
	time "time"
)

func TestSettersCreateEgressStreamRequest(t *testing.T) {
	t.Run("SetIngressID", func(t *testing.T) {
		obj := &CreateEgressStreamRequest{}
		var fernTestValueIngressID *string
		obj.SetIngressID(fernTestValueIngressID)
		assert.Equal(t, fernTestValueIngressID, obj.IngressID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRtsp", func(t *testing.T) {
		obj := &CreateEgressStreamRequest{}
		var fernTestValueRtsp *RtspSettings
		obj.SetRtsp(fernTestValueRtsp)
		assert.Equal(t, fernTestValueRtsp, obj.Rtsp)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSrt", func(t *testing.T) {
		obj := &CreateEgressStreamRequest{}
		var fernTestValueSrt *SrtSettings
		obj.SetSrt(fernTestValueSrt)
		assert.Equal(t, fernTestValueSrt, obj.Srt)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateEgressStreamRequest(t *testing.T) {
	t.Run("SetIngressID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEgressStreamRequest{}
		var fernTestValueIngressID *string

		// Act
		obj.SetIngressID(fernTestValueIngressID)

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

	t.Run("SetRtsp_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEgressStreamRequest{}
		var fernTestValueRtsp *RtspSettings

		// Act
		obj.SetRtsp(fernTestValueRtsp)

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

	t.Run("SetSrt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEgressStreamRequest{}
		var fernTestValueSrt *SrtSettings

		// Act
		obj.SetSrt(fernTestValueSrt)

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

func TestSettersCreateIngressStreamRequest(t *testing.T) {
	t.Run("SetIngressID", func(t *testing.T) {
		obj := &CreateIngressStreamRequest{}
		var fernTestValueIngressID *string
		obj.SetIngressID(fernTestValueIngressID)
		assert.Equal(t, fernTestValueIngressID, obj.IngressID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTitle", func(t *testing.T) {
		obj := &CreateIngressStreamRequest{}
		var fernTestValueTitle *string
		obj.SetTitle(fernTestValueTitle)
		assert.Equal(t, fernTestValueTitle, obj.Title)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMpegTs", func(t *testing.T) {
		obj := &CreateIngressStreamRequest{}
		var fernTestValueMpegTs *MpegTsSettings
		obj.SetMpegTs(fernTestValueMpegTs)
		assert.Equal(t, fernTestValueMpegTs, obj.MpegTs)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRtsp", func(t *testing.T) {
		obj := &CreateIngressStreamRequest{}
		var fernTestValueRtsp *RtspSettings
		obj.SetRtsp(fernTestValueRtsp)
		assert.Equal(t, fernTestValueRtsp, obj.Rtsp)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSrt", func(t *testing.T) {
		obj := &CreateIngressStreamRequest{}
		var fernTestValueSrt *SrtSettings
		obj.SetSrt(fernTestValueSrt)
		assert.Equal(t, fernTestValueSrt, obj.Srt)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateIngressStreamRequest(t *testing.T) {
	t.Run("SetIngressID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateIngressStreamRequest{}
		var fernTestValueIngressID *string

		// Act
		obj.SetIngressID(fernTestValueIngressID)

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

	t.Run("SetTitle_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateIngressStreamRequest{}
		var fernTestValueTitle *string

		// Act
		obj.SetTitle(fernTestValueTitle)

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

	t.Run("SetMpegTs_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateIngressStreamRequest{}
		var fernTestValueMpegTs *MpegTsSettings

		// Act
		obj.SetMpegTs(fernTestValueMpegTs)

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

	t.Run("SetRtsp_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateIngressStreamRequest{}
		var fernTestValueRtsp *RtspSettings

		// Act
		obj.SetRtsp(fernTestValueRtsp)

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

	t.Run("SetSrt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateIngressStreamRequest{}
		var fernTestValueSrt *SrtSettings

		// Act
		obj.SetSrt(fernTestValueSrt)

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

func TestSettersDeleteEgressStreamRequest(t *testing.T) {
	t.Run("SetEgressID", func(t *testing.T) {
		obj := &DeleteEgressStreamRequest{}
		var fernTestValueEgressID string
		obj.SetEgressID(fernTestValueEgressID)
		assert.Equal(t, fernTestValueEgressID, obj.EgressID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitDeleteEgressStreamRequest(t *testing.T) {
	t.Run("SetEgressID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &DeleteEgressStreamRequest{}
		var fernTestValueEgressID string

		// Act
		obj.SetEgressID(fernTestValueEgressID)

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

func TestSettersDeleteIngressStreamRequest(t *testing.T) {
	t.Run("SetIngressID", func(t *testing.T) {
		obj := &DeleteIngressStreamRequest{}
		var fernTestValueIngressID string
		obj.SetIngressID(fernTestValueIngressID)
		assert.Equal(t, fernTestValueIngressID, obj.IngressID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitDeleteIngressStreamRequest(t *testing.T) {
	t.Run("SetIngressID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &DeleteIngressStreamRequest{}
		var fernTestValueIngressID string

		// Act
		obj.SetIngressID(fernTestValueIngressID)

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

func TestSettersGetEgressStreamRequest(t *testing.T) {
	t.Run("SetEgressID", func(t *testing.T) {
		obj := &GetEgressStreamRequest{}
		var fernTestValueEgressID string
		obj.SetEgressID(fernTestValueEgressID)
		assert.Equal(t, fernTestValueEgressID, obj.EgressID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetEgressStreamRequest(t *testing.T) {
	t.Run("SetEgressID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetEgressStreamRequest{}
		var fernTestValueEgressID string

		// Act
		obj.SetEgressID(fernTestValueEgressID)

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

func TestSettersGetIngressStreamRequest(t *testing.T) {
	t.Run("SetIngressID", func(t *testing.T) {
		obj := &GetIngressStreamRequest{}
		var fernTestValueIngressID string
		obj.SetIngressID(fernTestValueIngressID)
		assert.Equal(t, fernTestValueIngressID, obj.IngressID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetIngressStreamRequest(t *testing.T) {
	t.Run("SetIngressID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetIngressStreamRequest{}
		var fernTestValueIngressID string

		// Act
		obj.SetIngressID(fernTestValueIngressID)

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

func TestSettersListEgressStreamsRequest(t *testing.T) {
	t.Run("SetPageSize", func(t *testing.T) {
		obj := &ListEgressStreamsRequest{}
		var fernTestValuePageSize *int
		obj.SetPageSize(fernTestValuePageSize)
		assert.Equal(t, fernTestValuePageSize, obj.PageSize)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetPageToken", func(t *testing.T) {
		obj := &ListEgressStreamsRequest{}
		var fernTestValuePageToken *string
		obj.SetPageToken(fernTestValuePageToken)
		assert.Equal(t, fernTestValuePageToken, obj.PageToken)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitListEgressStreamsRequest(t *testing.T) {
	t.Run("SetPageSize_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListEgressStreamsRequest{}
		var fernTestValuePageSize *int

		// Act
		obj.SetPageSize(fernTestValuePageSize)

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
		obj := &ListEgressStreamsRequest{}
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

}

func TestSettersListIngressStreamsRequest(t *testing.T) {
	t.Run("SetPageSize", func(t *testing.T) {
		obj := &ListIngressStreamsRequest{}
		var fernTestValuePageSize *int
		obj.SetPageSize(fernTestValuePageSize)
		assert.Equal(t, fernTestValuePageSize, obj.PageSize)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetPageToken", func(t *testing.T) {
		obj := &ListIngressStreamsRequest{}
		var fernTestValuePageToken *string
		obj.SetPageToken(fernTestValuePageToken)
		assert.Equal(t, fernTestValuePageToken, obj.PageToken)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitListIngressStreamsRequest(t *testing.T) {
	t.Run("SetPageSize_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListIngressStreamsRequest{}
		var fernTestValuePageSize *int

		// Act
		obj.SetPageSize(fernTestValuePageSize)

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
		obj := &ListIngressStreamsRequest{}
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

}

func TestSettersCreateEgressStreamResponse(t *testing.T) {
	t.Run("SetEgressID", func(t *testing.T) {
		obj := &CreateEgressStreamResponse{}
		var fernTestValueEgressID *string
		obj.SetEgressID(fernTestValueEgressID)
		assert.Equal(t, fernTestValueEgressID, obj.EgressID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRtsp", func(t *testing.T) {
		obj := &CreateEgressStreamResponse{}
		var fernTestValueRtsp *RtspEgress
		obj.SetRtsp(fernTestValueRtsp)
		assert.Equal(t, fernTestValueRtsp, obj.Rtsp)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSrt", func(t *testing.T) {
		obj := &CreateEgressStreamResponse{}
		var fernTestValueSrt *SrtEgress
		obj.SetSrt(fernTestValueSrt)
		assert.Equal(t, fernTestValueSrt, obj.Srt)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateEgressStreamResponse(t *testing.T) {
	t.Run("GetEgressID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEgressStreamResponse{}
		var expected *string
		obj.EgressID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEgressID(), "getter should return the property value")
	})

	t.Run("GetEgressID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEgressStreamResponse{}
		obj.EgressID = nil

		// Act & Assert
		assert.Nil(t, obj.GetEgressID(), "getter should return nil when property is nil")
	})

	t.Run("GetEgressID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateEgressStreamResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEgressID() // Should return zero value
	})

	t.Run("GetRtsp", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEgressStreamResponse{}
		var expected *RtspEgress
		obj.Rtsp = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRtsp(), "getter should return the property value")
	})

	t.Run("GetRtsp_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEgressStreamResponse{}
		obj.Rtsp = nil

		// Act & Assert
		assert.Nil(t, obj.GetRtsp(), "getter should return nil when property is nil")
	})

	t.Run("GetRtsp_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateEgressStreamResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRtsp() // Should return zero value
	})

	t.Run("GetSrt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEgressStreamResponse{}
		var expected *SrtEgress
		obj.Srt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSrt(), "getter should return the property value")
	})

	t.Run("GetSrt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEgressStreamResponse{}
		obj.Srt = nil

		// Act & Assert
		assert.Nil(t, obj.GetSrt(), "getter should return nil when property is nil")
	})

	t.Run("GetSrt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateEgressStreamResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSrt() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateEgressStreamResponse(t *testing.T) {
	t.Run("SetEgressID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEgressStreamResponse{}
		var fernTestValueEgressID *string

		// Act
		obj.SetEgressID(fernTestValueEgressID)

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

	t.Run("SetRtsp_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEgressStreamResponse{}
		var fernTestValueRtsp *RtspEgress

		// Act
		obj.SetRtsp(fernTestValueRtsp)

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

	t.Run("SetSrt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEgressStreamResponse{}
		var fernTestValueSrt *SrtEgress

		// Act
		obj.SetSrt(fernTestValueSrt)

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

func TestSettersCreateIngressStreamResponse(t *testing.T) {
	t.Run("SetIngressID", func(t *testing.T) {
		obj := &CreateIngressStreamResponse{}
		var fernTestValueIngressID *string
		obj.SetIngressID(fernTestValueIngressID)
		assert.Equal(t, fernTestValueIngressID, obj.IngressID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMpegTs", func(t *testing.T) {
		obj := &CreateIngressStreamResponse{}
		var fernTestValueMpegTs *MpegTsIngress
		obj.SetMpegTs(fernTestValueMpegTs)
		assert.Equal(t, fernTestValueMpegTs, obj.MpegTs)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSrt", func(t *testing.T) {
		obj := &CreateIngressStreamResponse{}
		var fernTestValueSrt *SrtIngress
		obj.SetSrt(fernTestValueSrt)
		assert.Equal(t, fernTestValueSrt, obj.Srt)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateIngressStreamResponse(t *testing.T) {
	t.Run("GetIngressID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateIngressStreamResponse{}
		var expected *string
		obj.IngressID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIngressID(), "getter should return the property value")
	})

	t.Run("GetIngressID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateIngressStreamResponse{}
		obj.IngressID = nil

		// Act & Assert
		assert.Nil(t, obj.GetIngressID(), "getter should return nil when property is nil")
	})

	t.Run("GetIngressID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateIngressStreamResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIngressID() // Should return zero value
	})

	t.Run("GetMpegTs", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateIngressStreamResponse{}
		var expected *MpegTsIngress
		obj.MpegTs = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMpegTs(), "getter should return the property value")
	})

	t.Run("GetMpegTs_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateIngressStreamResponse{}
		obj.MpegTs = nil

		// Act & Assert
		assert.Nil(t, obj.GetMpegTs(), "getter should return nil when property is nil")
	})

	t.Run("GetMpegTs_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateIngressStreamResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMpegTs() // Should return zero value
	})

	t.Run("GetSrt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateIngressStreamResponse{}
		var expected *SrtIngress
		obj.Srt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSrt(), "getter should return the property value")
	})

	t.Run("GetSrt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateIngressStreamResponse{}
		obj.Srt = nil

		// Act & Assert
		assert.Nil(t, obj.GetSrt(), "getter should return nil when property is nil")
	})

	t.Run("GetSrt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateIngressStreamResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSrt() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateIngressStreamResponse(t *testing.T) {
	t.Run("SetIngressID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateIngressStreamResponse{}
		var fernTestValueIngressID *string

		// Act
		obj.SetIngressID(fernTestValueIngressID)

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

	t.Run("SetMpegTs_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateIngressStreamResponse{}
		var fernTestValueMpegTs *MpegTsIngress

		// Act
		obj.SetMpegTs(fernTestValueMpegTs)

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

	t.Run("SetSrt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateIngressStreamResponse{}
		var fernTestValueSrt *SrtIngress

		// Act
		obj.SetSrt(fernTestValueSrt)

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

func TestSettersEgressStream(t *testing.T) {
	t.Run("SetEgressID", func(t *testing.T) {
		obj := &EgressStream{}
		var fernTestValueEgressID *string
		obj.SetEgressID(fernTestValueEgressID)
		assert.Equal(t, fernTestValueEgressID, obj.EgressID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetIngressID", func(t *testing.T) {
		obj := &EgressStream{}
		var fernTestValueIngressID *string
		obj.SetIngressID(fernTestValueIngressID)
		assert.Equal(t, fernTestValueIngressID, obj.IngressID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRtsp", func(t *testing.T) {
		obj := &EgressStream{}
		var fernTestValueRtsp *RtspEgress
		obj.SetRtsp(fernTestValueRtsp)
		assert.Equal(t, fernTestValueRtsp, obj.Rtsp)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSrt", func(t *testing.T) {
		obj := &EgressStream{}
		var fernTestValueSrt *SrtEgress
		obj.SetSrt(fernTestValueSrt)
		assert.Equal(t, fernTestValueSrt, obj.Srt)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersEgressStream(t *testing.T) {
	t.Run("GetEgressID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EgressStream{}
		var expected *string
		obj.EgressID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEgressID(), "getter should return the property value")
	})

	t.Run("GetEgressID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EgressStream{}
		obj.EgressID = nil

		// Act & Assert
		assert.Nil(t, obj.GetEgressID(), "getter should return nil when property is nil")
	})

	t.Run("GetEgressID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EgressStream
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEgressID() // Should return zero value
	})

	t.Run("GetIngressID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EgressStream{}
		var expected *string
		obj.IngressID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIngressID(), "getter should return the property value")
	})

	t.Run("GetIngressID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EgressStream{}
		obj.IngressID = nil

		// Act & Assert
		assert.Nil(t, obj.GetIngressID(), "getter should return nil when property is nil")
	})

	t.Run("GetIngressID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EgressStream
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIngressID() // Should return zero value
	})

	t.Run("GetRtsp", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EgressStream{}
		var expected *RtspEgress
		obj.Rtsp = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRtsp(), "getter should return the property value")
	})

	t.Run("GetRtsp_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EgressStream{}
		obj.Rtsp = nil

		// Act & Assert
		assert.Nil(t, obj.GetRtsp(), "getter should return nil when property is nil")
	})

	t.Run("GetRtsp_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EgressStream
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRtsp() // Should return zero value
	})

	t.Run("GetSrt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EgressStream{}
		var expected *SrtEgress
		obj.Srt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSrt(), "getter should return the property value")
	})

	t.Run("GetSrt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EgressStream{}
		obj.Srt = nil

		// Act & Assert
		assert.Nil(t, obj.GetSrt(), "getter should return nil when property is nil")
	})

	t.Run("GetSrt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EgressStream
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSrt() // Should return zero value
	})

}

func TestSettersMarkExplicitEgressStream(t *testing.T) {
	t.Run("SetEgressID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EgressStream{}
		var fernTestValueEgressID *string

		// Act
		obj.SetEgressID(fernTestValueEgressID)

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

	t.Run("SetIngressID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EgressStream{}
		var fernTestValueIngressID *string

		// Act
		obj.SetIngressID(fernTestValueIngressID)

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

	t.Run("SetRtsp_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EgressStream{}
		var fernTestValueRtsp *RtspEgress

		// Act
		obj.SetRtsp(fernTestValueRtsp)

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

	t.Run("SetSrt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EgressStream{}
		var fernTestValueSrt *SrtEgress

		// Act
		obj.SetSrt(fernTestValueSrt)

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

func TestSettersGetEgressStreamResponse(t *testing.T) {
	t.Run("SetEgressStream", func(t *testing.T) {
		obj := &GetEgressStreamResponse{}
		var fernTestValueEgressStream *EgressStream
		obj.SetEgressStream(fernTestValueEgressStream)
		assert.Equal(t, fernTestValueEgressStream, obj.EgressStream)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersGetEgressStreamResponse(t *testing.T) {
	t.Run("GetEgressStream", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetEgressStreamResponse{}
		var expected *EgressStream
		obj.EgressStream = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEgressStream(), "getter should return the property value")
	})

	t.Run("GetEgressStream_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetEgressStreamResponse{}
		obj.EgressStream = nil

		// Act & Assert
		assert.Nil(t, obj.GetEgressStream(), "getter should return nil when property is nil")
	})

	t.Run("GetEgressStream_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetEgressStreamResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEgressStream() // Should return zero value
	})

}

func TestSettersMarkExplicitGetEgressStreamResponse(t *testing.T) {
	t.Run("SetEgressStream_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetEgressStreamResponse{}
		var fernTestValueEgressStream *EgressStream

		// Act
		obj.SetEgressStream(fernTestValueEgressStream)

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

func TestSettersGetIngressStreamResponse(t *testing.T) {
	t.Run("SetIngressStream", func(t *testing.T) {
		obj := &GetIngressStreamResponse{}
		var fernTestValueIngressStream *IngressStream
		obj.SetIngressStream(fernTestValueIngressStream)
		assert.Equal(t, fernTestValueIngressStream, obj.IngressStream)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersGetIngressStreamResponse(t *testing.T) {
	t.Run("GetIngressStream", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetIngressStreamResponse{}
		var expected *IngressStream
		obj.IngressStream = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIngressStream(), "getter should return the property value")
	})

	t.Run("GetIngressStream_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetIngressStreamResponse{}
		obj.IngressStream = nil

		// Act & Assert
		assert.Nil(t, obj.GetIngressStream(), "getter should return nil when property is nil")
	})

	t.Run("GetIngressStream_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetIngressStreamResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIngressStream() // Should return zero value
	})

}

func TestSettersMarkExplicitGetIngressStreamResponse(t *testing.T) {
	t.Run("SetIngressStream_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetIngressStreamResponse{}
		var fernTestValueIngressStream *IngressStream

		// Act
		obj.SetIngressStream(fernTestValueIngressStream)

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

func TestSettersIngressStream(t *testing.T) {
	t.Run("SetIngressID", func(t *testing.T) {
		obj := &IngressStream{}
		var fernTestValueIngressID *string
		obj.SetIngressID(fernTestValueIngressID)
		assert.Equal(t, fernTestValueIngressID, obj.IngressID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTitle", func(t *testing.T) {
		obj := &IngressStream{}
		var fernTestValueTitle *string
		obj.SetTitle(fernTestValueTitle)
		assert.Equal(t, fernTestValueTitle, obj.Title)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &IngressStream{}
		var fernTestValueStatus *IngressStreamStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMpegTs", func(t *testing.T) {
		obj := &IngressStream{}
		var fernTestValueMpegTs *MpegTsIngress
		obj.SetMpegTs(fernTestValueMpegTs)
		assert.Equal(t, fernTestValueMpegTs, obj.MpegTs)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRtsp", func(t *testing.T) {
		obj := &IngressStream{}
		var fernTestValueRtsp *RtspIngress
		obj.SetRtsp(fernTestValueRtsp)
		assert.Equal(t, fernTestValueRtsp, obj.Rtsp)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSrt", func(t *testing.T) {
		obj := &IngressStream{}
		var fernTestValueSrt *SrtIngress
		obj.SetSrt(fernTestValueSrt)
		assert.Equal(t, fernTestValueSrt, obj.Srt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCreatedAt", func(t *testing.T) {
		obj := &IngressStream{}
		var fernTestValueCreatedAt *time.Time
		obj.SetCreatedAt(fernTestValueCreatedAt)
		assert.Equal(t, fernTestValueCreatedAt, obj.CreatedAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetUpdatedAt", func(t *testing.T) {
		obj := &IngressStream{}
		var fernTestValueUpdatedAt *time.Time
		obj.SetUpdatedAt(fernTestValueUpdatedAt)
		assert.Equal(t, fernTestValueUpdatedAt, obj.UpdatedAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEgressIDs", func(t *testing.T) {
		obj := &IngressStream{}
		var fernTestValueEgressIDs []string
		obj.SetEgressIDs(fernTestValueEgressIDs)
		assert.Equal(t, fernTestValueEgressIDs, obj.EgressIDs)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersIngressStream(t *testing.T) {
	t.Run("GetIngressID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var expected *string
		obj.IngressID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIngressID(), "getter should return the property value")
	})

	t.Run("GetIngressID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		obj.IngressID = nil

		// Act & Assert
		assert.Nil(t, obj.GetIngressID(), "getter should return nil when property is nil")
	})

	t.Run("GetIngressID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IngressStream
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIngressID() // Should return zero value
	})

	t.Run("GetTitle", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var expected *string
		obj.Title = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTitle(), "getter should return the property value")
	})

	t.Run("GetTitle_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		obj.Title = nil

		// Act & Assert
		assert.Nil(t, obj.GetTitle(), "getter should return nil when property is nil")
	})

	t.Run("GetTitle_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IngressStream
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTitle() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var expected *IngressStreamStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		obj.Status = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IngressStream
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetMpegTs", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var expected *MpegTsIngress
		obj.MpegTs = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMpegTs(), "getter should return the property value")
	})

	t.Run("GetMpegTs_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		obj.MpegTs = nil

		// Act & Assert
		assert.Nil(t, obj.GetMpegTs(), "getter should return nil when property is nil")
	})

	t.Run("GetMpegTs_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IngressStream
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMpegTs() // Should return zero value
	})

	t.Run("GetRtsp", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var expected *RtspIngress
		obj.Rtsp = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRtsp(), "getter should return the property value")
	})

	t.Run("GetRtsp_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		obj.Rtsp = nil

		// Act & Assert
		assert.Nil(t, obj.GetRtsp(), "getter should return nil when property is nil")
	})

	t.Run("GetRtsp_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IngressStream
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRtsp() // Should return zero value
	})

	t.Run("GetSrt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var expected *SrtIngress
		obj.Srt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSrt(), "getter should return the property value")
	})

	t.Run("GetSrt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		obj.Srt = nil

		// Act & Assert
		assert.Nil(t, obj.GetSrt(), "getter should return nil when property is nil")
	})

	t.Run("GetSrt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IngressStream
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSrt() // Should return zero value
	})

	t.Run("GetCreatedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var expected *time.Time
		obj.CreatedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCreatedAt(), "getter should return the property value")
	})

	t.Run("GetCreatedAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		obj.CreatedAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetCreatedAt(), "getter should return nil when property is nil")
	})

	t.Run("GetCreatedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IngressStream
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCreatedAt() // Should return zero value
	})

	t.Run("GetUpdatedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var expected *time.Time
		obj.UpdatedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetUpdatedAt(), "getter should return the property value")
	})

	t.Run("GetUpdatedAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		obj.UpdatedAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetUpdatedAt(), "getter should return nil when property is nil")
	})

	t.Run("GetUpdatedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IngressStream
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetUpdatedAt() // Should return zero value
	})

	t.Run("GetEgressIDs", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var expected []string
		obj.EgressIDs = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEgressIDs(), "getter should return the property value")
	})

	t.Run("GetEgressIDs_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		obj.EgressIDs = nil

		// Act & Assert
		assert.Nil(t, obj.GetEgressIDs(), "getter should return nil when property is nil")
	})

	t.Run("GetEgressIDs_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IngressStream
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEgressIDs() // Should return zero value
	})

}

func TestSettersMarkExplicitIngressStream(t *testing.T) {
	t.Run("SetIngressID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var fernTestValueIngressID *string

		// Act
		obj.SetIngressID(fernTestValueIngressID)

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

	t.Run("SetTitle_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var fernTestValueTitle *string

		// Act
		obj.SetTitle(fernTestValueTitle)

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

	t.Run("SetStatus_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var fernTestValueStatus *IngressStreamStatus

		// Act
		obj.SetStatus(fernTestValueStatus)

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

	t.Run("SetMpegTs_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var fernTestValueMpegTs *MpegTsIngress

		// Act
		obj.SetMpegTs(fernTestValueMpegTs)

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

	t.Run("SetRtsp_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var fernTestValueRtsp *RtspIngress

		// Act
		obj.SetRtsp(fernTestValueRtsp)

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

	t.Run("SetSrt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var fernTestValueSrt *SrtIngress

		// Act
		obj.SetSrt(fernTestValueSrt)

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

	t.Run("SetCreatedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var fernTestValueCreatedAt *time.Time

		// Act
		obj.SetCreatedAt(fernTestValueCreatedAt)

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

	t.Run("SetUpdatedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var fernTestValueUpdatedAt *time.Time

		// Act
		obj.SetUpdatedAt(fernTestValueUpdatedAt)

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

	t.Run("SetEgressIDs_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}
		var fernTestValueEgressIDs []string

		// Act
		obj.SetEgressIDs(fernTestValueEgressIDs)

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

func TestSettersListEgressStreamsResponse(t *testing.T) {
	t.Run("SetEgressStreams", func(t *testing.T) {
		obj := &ListEgressStreamsResponse{}
		var fernTestValueEgressStreams []*EgressStream
		obj.SetEgressStreams(fernTestValueEgressStreams)
		assert.Equal(t, fernTestValueEgressStreams, obj.EgressStreams)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetNextPageToken", func(t *testing.T) {
		obj := &ListEgressStreamsResponse{}
		var fernTestValueNextPageToken *string
		obj.SetNextPageToken(fernTestValueNextPageToken)
		assert.Equal(t, fernTestValueNextPageToken, obj.NextPageToken)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListEgressStreamsResponse(t *testing.T) {
	t.Run("GetEgressStreams", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListEgressStreamsResponse{}
		var expected []*EgressStream
		obj.EgressStreams = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEgressStreams(), "getter should return the property value")
	})

	t.Run("GetEgressStreams_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListEgressStreamsResponse{}
		obj.EgressStreams = nil

		// Act & Assert
		assert.Nil(t, obj.GetEgressStreams(), "getter should return nil when property is nil")
	})

	t.Run("GetEgressStreams_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListEgressStreamsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEgressStreams() // Should return zero value
	})

	t.Run("GetNextPageToken", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListEgressStreamsResponse{}
		var expected *string
		obj.NextPageToken = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetNextPageToken(), "getter should return the property value")
	})

	t.Run("GetNextPageToken_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListEgressStreamsResponse{}
		obj.NextPageToken = nil

		// Act & Assert
		assert.Nil(t, obj.GetNextPageToken(), "getter should return nil when property is nil")
	})

	t.Run("GetNextPageToken_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListEgressStreamsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetNextPageToken() // Should return zero value
	})

}

func TestSettersMarkExplicitListEgressStreamsResponse(t *testing.T) {
	t.Run("SetEgressStreams_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListEgressStreamsResponse{}
		var fernTestValueEgressStreams []*EgressStream

		// Act
		obj.SetEgressStreams(fernTestValueEgressStreams)

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
		obj := &ListEgressStreamsResponse{}
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

func TestSettersListIngressStreamsResponse(t *testing.T) {
	t.Run("SetIngressStreams", func(t *testing.T) {
		obj := &ListIngressStreamsResponse{}
		var fernTestValueIngressStreams []*IngressStream
		obj.SetIngressStreams(fernTestValueIngressStreams)
		assert.Equal(t, fernTestValueIngressStreams, obj.IngressStreams)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetNextPageToken", func(t *testing.T) {
		obj := &ListIngressStreamsResponse{}
		var fernTestValueNextPageToken *string
		obj.SetNextPageToken(fernTestValueNextPageToken)
		assert.Equal(t, fernTestValueNextPageToken, obj.NextPageToken)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListIngressStreamsResponse(t *testing.T) {
	t.Run("GetIngressStreams", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListIngressStreamsResponse{}
		var expected []*IngressStream
		obj.IngressStreams = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIngressStreams(), "getter should return the property value")
	})

	t.Run("GetIngressStreams_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListIngressStreamsResponse{}
		obj.IngressStreams = nil

		// Act & Assert
		assert.Nil(t, obj.GetIngressStreams(), "getter should return nil when property is nil")
	})

	t.Run("GetIngressStreams_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListIngressStreamsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIngressStreams() // Should return zero value
	})

	t.Run("GetNextPageToken", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListIngressStreamsResponse{}
		var expected *string
		obj.NextPageToken = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetNextPageToken(), "getter should return the property value")
	})

	t.Run("GetNextPageToken_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListIngressStreamsResponse{}
		obj.NextPageToken = nil

		// Act & Assert
		assert.Nil(t, obj.GetNextPageToken(), "getter should return nil when property is nil")
	})

	t.Run("GetNextPageToken_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListIngressStreamsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetNextPageToken() // Should return zero value
	})

}

func TestSettersMarkExplicitListIngressStreamsResponse(t *testing.T) {
	t.Run("SetIngressStreams_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListIngressStreamsResponse{}
		var fernTestValueIngressStreams []*IngressStream

		// Act
		obj.SetIngressStreams(fernTestValueIngressStreams)

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
		obj := &ListIngressStreamsResponse{}
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

func TestSettersMpegTsIngress(t *testing.T) {
	t.Run("SetURL", func(t *testing.T) {
		obj := &MpegTsIngress{}
		var fernTestValueURL *string
		obj.SetURL(fernTestValueURL)
		assert.Equal(t, fernTestValueURL, obj.URL)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersMpegTsIngress(t *testing.T) {
	t.Run("GetURL", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MpegTsIngress{}
		var expected *string
		obj.URL = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetURL(), "getter should return the property value")
	})

	t.Run("GetURL_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MpegTsIngress{}
		obj.URL = nil

		// Act & Assert
		assert.Nil(t, obj.GetURL(), "getter should return nil when property is nil")
	})

	t.Run("GetURL_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *MpegTsIngress
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetURL() // Should return zero value
	})

}

func TestSettersMarkExplicitMpegTsIngress(t *testing.T) {
	t.Run("SetURL_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MpegTsIngress{}
		var fernTestValueURL *string

		// Act
		obj.SetURL(fernTestValueURL)

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

func TestSettersRtspEgress(t *testing.T) {
	t.Run("SetURL", func(t *testing.T) {
		obj := &RtspEgress{}
		var fernTestValueURL *string
		obj.SetURL(fernTestValueURL)
		assert.Equal(t, fernTestValueURL, obj.URL)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersRtspEgress(t *testing.T) {
	t.Run("GetURL", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RtspEgress{}
		var expected *string
		obj.URL = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetURL(), "getter should return the property value")
	})

	t.Run("GetURL_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RtspEgress{}
		obj.URL = nil

		// Act & Assert
		assert.Nil(t, obj.GetURL(), "getter should return nil when property is nil")
	})

	t.Run("GetURL_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RtspEgress
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetURL() // Should return zero value
	})

}

func TestSettersMarkExplicitRtspEgress(t *testing.T) {
	t.Run("SetURL_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RtspEgress{}
		var fernTestValueURL *string

		// Act
		obj.SetURL(fernTestValueURL)

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

func TestSettersRtspIngress(t *testing.T) {
	t.Run("SetURL", func(t *testing.T) {
		obj := &RtspIngress{}
		var fernTestValueURL *string
		obj.SetURL(fernTestValueURL)
		assert.Equal(t, fernTestValueURL, obj.URL)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersRtspIngress(t *testing.T) {
	t.Run("GetURL", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RtspIngress{}
		var expected *string
		obj.URL = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetURL(), "getter should return the property value")
	})

	t.Run("GetURL_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RtspIngress{}
		obj.URL = nil

		// Act & Assert
		assert.Nil(t, obj.GetURL(), "getter should return nil when property is nil")
	})

	t.Run("GetURL_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RtspIngress
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetURL() // Should return zero value
	})

}

func TestSettersMarkExplicitRtspIngress(t *testing.T) {
	t.Run("SetURL_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RtspIngress{}
		var fernTestValueURL *string

		// Act
		obj.SetURL(fernTestValueURL)

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

func TestSettersRtspSettings(t *testing.T) {
	t.Run("SetURL", func(t *testing.T) {
		obj := &RtspSettings{}
		var fernTestValueURL *string
		obj.SetURL(fernTestValueURL)
		assert.Equal(t, fernTestValueURL, obj.URL)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersRtspSettings(t *testing.T) {
	t.Run("GetURL", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RtspSettings{}
		var expected *string
		obj.URL = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetURL(), "getter should return the property value")
	})

	t.Run("GetURL_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RtspSettings{}
		obj.URL = nil

		// Act & Assert
		assert.Nil(t, obj.GetURL(), "getter should return nil when property is nil")
	})

	t.Run("GetURL_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RtspSettings
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetURL() // Should return zero value
	})

}

func TestSettersMarkExplicitRtspSettings(t *testing.T) {
	t.Run("SetURL_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RtspSettings{}
		var fernTestValueURL *string

		// Act
		obj.SetURL(fernTestValueURL)

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

func TestSettersSrtEgress(t *testing.T) {
	t.Run("SetURL", func(t *testing.T) {
		obj := &SrtEgress{}
		var fernTestValueURL *string
		obj.SetURL(fernTestValueURL)
		assert.Equal(t, fernTestValueURL, obj.URL)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSessionID", func(t *testing.T) {
		obj := &SrtEgress{}
		var fernTestValueSessionID *string
		obj.SetSessionID(fernTestValueSessionID)
		assert.Equal(t, fernTestValueSessionID, obj.SessionID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersSrtEgress(t *testing.T) {
	t.Run("GetURL", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtEgress{}
		var expected *string
		obj.URL = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetURL(), "getter should return the property value")
	})

	t.Run("GetURL_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtEgress{}
		obj.URL = nil

		// Act & Assert
		assert.Nil(t, obj.GetURL(), "getter should return nil when property is nil")
	})

	t.Run("GetURL_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *SrtEgress
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetURL() // Should return zero value
	})

	t.Run("GetSessionID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtEgress{}
		var expected *string
		obj.SessionID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSessionID(), "getter should return the property value")
	})

	t.Run("GetSessionID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtEgress{}
		obj.SessionID = nil

		// Act & Assert
		assert.Nil(t, obj.GetSessionID(), "getter should return nil when property is nil")
	})

	t.Run("GetSessionID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *SrtEgress
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSessionID() // Should return zero value
	})

}

func TestSettersMarkExplicitSrtEgress(t *testing.T) {
	t.Run("SetURL_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtEgress{}
		var fernTestValueURL *string

		// Act
		obj.SetURL(fernTestValueURL)

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

	t.Run("SetSessionID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtEgress{}
		var fernTestValueSessionID *string

		// Act
		obj.SetSessionID(fernTestValueSessionID)

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

func TestSettersSrtIngress(t *testing.T) {
	t.Run("SetURL", func(t *testing.T) {
		obj := &SrtIngress{}
		var fernTestValueURL *string
		obj.SetURL(fernTestValueURL)
		assert.Equal(t, fernTestValueURL, obj.URL)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSessionID", func(t *testing.T) {
		obj := &SrtIngress{}
		var fernTestValueSessionID *string
		obj.SetSessionID(fernTestValueSessionID)
		assert.Equal(t, fernTestValueSessionID, obj.SessionID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersSrtIngress(t *testing.T) {
	t.Run("GetURL", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtIngress{}
		var expected *string
		obj.URL = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetURL(), "getter should return the property value")
	})

	t.Run("GetURL_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtIngress{}
		obj.URL = nil

		// Act & Assert
		assert.Nil(t, obj.GetURL(), "getter should return nil when property is nil")
	})

	t.Run("GetURL_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *SrtIngress
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetURL() // Should return zero value
	})

	t.Run("GetSessionID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtIngress{}
		var expected *string
		obj.SessionID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSessionID(), "getter should return the property value")
	})

	t.Run("GetSessionID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtIngress{}
		obj.SessionID = nil

		// Act & Assert
		assert.Nil(t, obj.GetSessionID(), "getter should return nil when property is nil")
	})

	t.Run("GetSessionID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *SrtIngress
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSessionID() // Should return zero value
	})

}

func TestSettersMarkExplicitSrtIngress(t *testing.T) {
	t.Run("SetURL_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtIngress{}
		var fernTestValueURL *string

		// Act
		obj.SetURL(fernTestValueURL)

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

	t.Run("SetSessionID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtIngress{}
		var fernTestValueSessionID *string

		// Act
		obj.SetSessionID(fernTestValueSessionID)

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

func TestSettersSrtSettings(t *testing.T) {
	t.Run("SetPassphrase", func(t *testing.T) {
		obj := &SrtSettings{}
		var fernTestValuePassphrase *string
		obj.SetPassphrase(fernTestValuePassphrase)
		assert.Equal(t, fernTestValuePassphrase, obj.Passphrase)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersSrtSettings(t *testing.T) {
	t.Run("GetPassphrase", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtSettings{}
		var expected *string
		obj.Passphrase = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPassphrase(), "getter should return the property value")
	})

	t.Run("GetPassphrase_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtSettings{}
		obj.Passphrase = nil

		// Act & Assert
		assert.Nil(t, obj.GetPassphrase(), "getter should return nil when property is nil")
	})

	t.Run("GetPassphrase_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *SrtSettings
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPassphrase() // Should return zero value
	})

}

func TestSettersMarkExplicitSrtSettings(t *testing.T) {
	t.Run("SetPassphrase_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtSettings{}
		var fernTestValuePassphrase *string

		// Act
		obj.SetPassphrase(fernTestValuePassphrase)

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

func TestJSONMarshalingCreateEgressStreamResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEgressStreamResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateEgressStreamResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateEgressStreamResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateEgressStreamResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCreateIngressStreamResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateIngressStreamResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateIngressStreamResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateIngressStreamResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateIngressStreamResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingDeleteEgressStreamResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &DeleteEgressStreamResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled DeleteEgressStreamResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj DeleteEgressStreamResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj DeleteEgressStreamResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingDeleteIngressStreamResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &DeleteIngressStreamResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled DeleteIngressStreamResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj DeleteIngressStreamResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj DeleteIngressStreamResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingEgressStream(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &EgressStream{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled EgressStream
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj EgressStream
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj EgressStream
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingGetEgressStreamResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetEgressStreamResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled GetEgressStreamResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj GetEgressStreamResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj GetEgressStreamResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingGetIngressStreamResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetIngressStreamResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled GetIngressStreamResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj GetIngressStreamResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj GetIngressStreamResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
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

func TestJSONMarshalingIngressStream(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IngressStream{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled IngressStream
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj IngressStream
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj IngressStream
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListEgressStreamsResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListEgressStreamsResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListEgressStreamsResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListEgressStreamsResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListEgressStreamsResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListIngressStreamsResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListIngressStreamsResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListIngressStreamsResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListIngressStreamsResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListIngressStreamsResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingMpegTsIngress(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MpegTsIngress{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled MpegTsIngress
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj MpegTsIngress
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj MpegTsIngress
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingRtspEgress(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RtspEgress{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled RtspEgress
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj RtspEgress
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj RtspEgress
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingRtspIngress(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RtspIngress{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled RtspIngress
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj RtspIngress
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj RtspIngress
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingRtspSettings(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RtspSettings{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled RtspSettings
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj RtspSettings
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj RtspSettings
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingSrtEgress(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtEgress{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled SrtEgress
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj SrtEgress
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj SrtEgress
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingSrtIngress(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtIngress{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled SrtIngress
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj SrtIngress
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj SrtIngress
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingSrtSettings(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SrtSettings{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled SrtSettings
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj SrtSettings
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj SrtSettings
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestStringCreateEgressStreamResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateEgressStreamResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateEgressStreamResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCreateIngressStreamResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateIngressStreamResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateIngressStreamResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringDeleteEgressStreamResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &DeleteEgressStreamResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *DeleteEgressStreamResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringDeleteIngressStreamResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &DeleteIngressStreamResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *DeleteIngressStreamResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringEgressStream(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &EgressStream{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *EgressStream
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringGetEgressStreamResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &GetEgressStreamResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetEgressStreamResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringGetIngressStreamResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &GetIngressStreamResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetIngressStreamResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
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

func TestStringIngressStream(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &IngressStream{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IngressStream
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListEgressStreamsResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListEgressStreamsResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListEgressStreamsResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListIngressStreamsResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListIngressStreamsResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListIngressStreamsResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringMpegTsIngress(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &MpegTsIngress{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *MpegTsIngress
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringRtspEgress(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &RtspEgress{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RtspEgress
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringRtspIngress(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &RtspIngress{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RtspIngress
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringRtspSettings(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &RtspSettings{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RtspSettings
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringSrtEgress(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &SrtEgress{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *SrtEgress
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringSrtIngress(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &SrtIngress{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *SrtIngress
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringSrtSettings(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &SrtSettings{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *SrtSettings
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestEnumIngressStreamStatus(t *testing.T) {
	t.Run("NewFromString_STREAM_STATUS_UNSPECIFIED", func(t *testing.T) {
		t.Parallel()
		val, err := NewIngressStreamStatusFromString("STREAM_STATUS_UNSPECIFIED")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, IngressStreamStatus("STREAM_STATUS_UNSPECIFIED"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_STREAM_STATUS_LIVE", func(t *testing.T) {
		t.Parallel()
		val, err := NewIngressStreamStatusFromString("STREAM_STATUS_LIVE")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, IngressStreamStatus("STREAM_STATUS_LIVE"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_STREAM_STATUS_INACTIVE", func(t *testing.T) {
		t.Parallel()
		val, err := NewIngressStreamStatusFromString("STREAM_STATUS_INACTIVE")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, IngressStreamStatus("STREAM_STATUS_INACTIVE"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_STREAM_STATUS_UNAVAILABLE", func(t *testing.T) {
		t.Parallel()
		val, err := NewIngressStreamStatusFromString("STREAM_STATUS_UNAVAILABLE")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, IngressStreamStatus("STREAM_STATUS_UNAVAILABLE"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_STREAM_STATUS_ARCHIVED", func(t *testing.T) {
		t.Parallel()
		val, err := NewIngressStreamStatusFromString("STREAM_STATUS_ARCHIVED")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, IngressStreamStatus("STREAM_STATUS_ARCHIVED"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewIngressStreamStatusFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewIngressStreamStatusFromString("STREAM_STATUS_UNSPECIFIED")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestExtraPropertiesCreateEgressStreamResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateEgressStreamResponse{}
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
		var obj *CreateEgressStreamResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCreateIngressStreamResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateIngressStreamResponse{}
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
		var obj *CreateIngressStreamResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesDeleteEgressStreamResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &DeleteEgressStreamResponse{}
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
		var obj *DeleteEgressStreamResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesDeleteIngressStreamResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &DeleteIngressStreamResponse{}
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
		var obj *DeleteIngressStreamResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesEgressStream(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &EgressStream{}
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
		var obj *EgressStream
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesGetEgressStreamResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &GetEgressStreamResponse{}
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
		var obj *GetEgressStreamResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesGetIngressStreamResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &GetIngressStreamResponse{}
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
		var obj *GetIngressStreamResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
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

func TestExtraPropertiesIngressStream(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &IngressStream{}
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
		var obj *IngressStream
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListEgressStreamsResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListEgressStreamsResponse{}
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
		var obj *ListEgressStreamsResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListIngressStreamsResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListIngressStreamsResponse{}
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
		var obj *ListIngressStreamsResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesMpegTsIngress(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &MpegTsIngress{}
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
		var obj *MpegTsIngress
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesRtspEgress(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &RtspEgress{}
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
		var obj *RtspEgress
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesRtspIngress(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &RtspIngress{}
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
		var obj *RtspIngress
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesRtspSettings(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &RtspSettings{}
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
		var obj *RtspSettings
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesSrtEgress(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &SrtEgress{}
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
		var obj *SrtEgress
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesSrtIngress(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &SrtIngress{}
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
		var obj *SrtIngress
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesSrtSettings(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &SrtSettings{}
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
		var obj *SrtSettings
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}
