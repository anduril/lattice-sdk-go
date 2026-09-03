// Code generated from our API definition. DO NOT EDIT.

package video

import (
	json "encoding/json"
	core "github.com/anduril/lattice-sdk-go/v5/core"
)

// The request was malformed or violated a precondition. Possible causes include missing
// required fields, an unsupported value, an `oneof` with zero or more than one variant
// set, a malformed URL, or a state-based precondition such as the source ingress stream
// not being live.
type BadRequestError struct {
	*core.APIError
	Body *GoogleRPCStatus
}

func (b *BadRequestError) UnmarshalJSON(data []byte) error {
	var body *GoogleRPCStatus
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	b.StatusCode = 400
	b.Body = body
	return nil
}

func (b *BadRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Body)
}

func (b *BadRequestError) Unwrap() error {
	return b.APIError
}

// The supplied identifier is already in use. For ingress creation this means the
// requested `ingress_id` is taken by another stream; for egress creation an egress
// stream already exists for the requested source.
type ConflictError struct {
	*core.APIError
	Body *GoogleRPCStatus
}

func (c *ConflictError) UnmarshalJSON(data []byte) error {
	var body *GoogleRPCStatus
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	c.StatusCode = 409
	c.Body = body
	return nil
}

func (c *ConflictError) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Body)
}

func (c *ConflictError) Unwrap() error {
	return c.APIError
}

// The caller is authenticated but lacks permission for the requested resource.
type ForbiddenError struct {
	*core.APIError
	Body *GoogleRPCStatus
}

func (f *ForbiddenError) UnmarshalJSON(data []byte) error {
	var body *GoogleRPCStatus
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	f.StatusCode = 403
	f.Body = body
	return nil
}

func (f *ForbiddenError) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Body)
}

func (f *ForbiddenError) Unwrap() error {
	return f.APIError
}

// An unexpected error occurred while processing the request. The ingress or egress
// backend may have failed, or an underlying storage operation may have failed. Retry
// with backoff.
type InternalServerError struct {
	*core.APIError
	Body *GoogleRPCStatus
}

func (i *InternalServerError) UnmarshalJSON(data []byte) error {
	var body *GoogleRPCStatus
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	i.StatusCode = 500
	i.Body = body
	return nil
}

func (i *InternalServerError) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Body)
}

func (i *InternalServerError) Unwrap() error {
	return i.APIError
}

// The specified ingress or egress stream was not found.
type NotFoundError struct {
	*core.APIError
	Body *GoogleRPCStatus
}

func (n *NotFoundError) UnmarshalJSON(data []byte) error {
	var body *GoogleRPCStatus
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	n.StatusCode = 404
	n.Body = body
	return nil
}

func (n *NotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Body)
}

func (n *NotFoundError) Unwrap() error {
	return n.APIError
}

// The service is temporarily unable to handle the request, such as during startup or
// when a backend dependency is shutting down. The response carries a RetryInfo hint
// (typically around 2 seconds); clients should retry per that hint.
type ServiceUnavailableError struct {
	*core.APIError
	Body *GoogleRPCStatus
}

func (s *ServiceUnavailableError) UnmarshalJSON(data []byte) error {
	var body *GoogleRPCStatus
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	s.StatusCode = 503
	s.Body = body
	return nil
}

func (s *ServiceUnavailableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Body)
}

func (s *ServiceUnavailableError) Unwrap() error {
	return s.APIError
}

// A service-wide resource pool is exhausted. For egress this means a port could not be
// allocated or the egress instance is at capacity. There is no per-caller quota or rate
// limit, so this reflects deployment-wide capacity rather than caller behavior.
type TooManyRequestsError struct {
	*core.APIError
	Body *GoogleRPCStatus
}

func (t *TooManyRequestsError) UnmarshalJSON(data []byte) error {
	var body *GoogleRPCStatus
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	t.StatusCode = 429
	t.Body = body
	return nil
}

func (t *TooManyRequestsError) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Body)
}

func (t *TooManyRequestsError) Unwrap() error {
	return t.APIError
}

// The request was rejected because the bearer token was missing, malformed, or could
// not be resolved to a user identity.
type UnauthorizedError struct {
	*core.APIError
	Body *GoogleRPCStatus
}

func (u *UnauthorizedError) UnmarshalJSON(data []byte) error {
	var body *GoogleRPCStatus
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 401
	u.Body = body
	return nil
}

func (u *UnauthorizedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}

func (u *UnauthorizedError) Unwrap() error {
	return u.APIError
}
