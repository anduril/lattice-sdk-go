package internal

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	core "github.com/anduril/lattice-sdk-go/v4/core"
)

// HTTPClient is an interface for a subset of the *http.Client.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// ResolveBaseURL resolves the base URL from the given arguments,
// preferring the first non-empty value.
func ResolveBaseURL(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

// EncodeURL encodes the given arguments into the URL, escaping
// values as needed. Pointer arguments are dereferenced before processing.
func EncodeURL(urlFormat string, args ...interface{}) string {
	escapedArgs := make([]interface{}, 0, len(args))
	for _, arg := range args {
		value := dereferenceArg(arg)
		escapedArgs = append(escapedArgs, url.PathEscape(fmt.Sprintf("%v", value)))
	}
	return fmt.Sprintf(urlFormat, escapedArgs...)
}

// dereferenceArg dereferences a pointer argument if necessary, returning the underlying value.
// If the argument is not a pointer or is nil, it returns the argument as-is.
func dereferenceArg(arg interface{}) interface{} {
	if arg == nil {
		return arg
	}

	v := reflect.ValueOf(arg)
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	return v.Interface()
}

// MergeHeaders merges the given headers together, where the right
// takes precedence over the left.
func MergeHeaders(left, right http.Header) http.Header {
	for key, values := range right {
		if len(values) > 1 {
			left[key] = values
			continue
		}

		if value := right.Get(key); value != "" {
			left.Set(key, value)
		}
	}

	return left
}

// MergeRequestHeaders converts request options into headers and merges them,
// where the right options take precedence over the left.
func MergeRequestHeaders(leftOptions, rightOptions *core.RequestOptions) (http.Header, error) {
	left, err := leftOptions.ToHeader()
	if err != nil {
		return nil, err
	}

	right, err := rightOptions.ToHeader()
	if err != nil {
		return nil, err
	}

	return MergeHeaders(left, right), nil
}
