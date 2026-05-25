package core

import (
	"errors"
	"testing"
)

func TestRequestOptions_ToHeaderReturnsTokenGetterError(t *testing.T) {
	expectedErr := errors.New("token fetch failed")

	options := NewRequestOptions()
	options.SetTokenGetter(func() (string, error) {
		return "", expectedErr
	})

	_, err := options.ToHeader()
	if err == nil {
		t.Fatal("expected ToHeader to return token getter error")
	}

	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected error to wrap token getter error, got %v", err)
	}
}

func TestRequestOptions_ToHeaderSetsBearerTokenFromTokenGetter(t *testing.T) {
	options := NewRequestOptions()
	options.SetTokenGetter(func() (string, error) {
		return "test-token", nil
	})

	headers, err := options.ToHeader()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got := headers.Get("Authorization"); got != "Bearer test-token" {
		t.Fatalf("expected Authorization header %q, got %q", "Bearer test-token", got)
	}
}

func TestRequestOptions_ToHeaderExplicitTokenTakesPrecedence(t *testing.T) {
	options := NewRequestOptions(&TokenOption{Token: "explicit-token"})
	options.SetTokenGetter(func() (string, error) {
		return "oauth-token", nil
	})

	headers, err := options.ToHeader()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got := headers.Get("Authorization"); got != "Bearer explicit-token" {
		t.Fatalf("expected explicit Authorization header %q, got %q", "Bearer explicit-token", got)
	}
}
