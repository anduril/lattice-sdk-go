# Changelog

## [v5.2.0] - 2026-10-10
### Added
- **`TaskStatus.LastUpdateTime`** — new optional field with `GetLastUpdateTime()`/`SetLastUpdateTime()` accessors that guard against out-of-order status updates.
- **`WithEventDiscriminator`** — now accepts a variadic `envelopeEvents` argument to wrap matching SSE events as `{"<field>":"<event>","data":<data>}` envelopes.
- **`EnvelopeEvents`** — new streaming support for event-discriminated envelope streams.
- **`GetBody()`** — nil-safe accessor added to all typed error types across the SDK (e.g. `BadRequestError`, `InternalServerError`, `NotFoundError`, `TooManyRequestsError`).
- **`ClientErrorWildcard` / `ServerErrorWildcard`** — error decoding now resolves unmapped 4XX/5XX status codes to the appropriate wildcard error constructor.

### Changed
- **`TaskStreamRequestTaskType`** — union deserialization now disambiguates variants using exact and partial object-key matching for more reliable unmarshaling.
- **`Date.UnmarshalJSON`** — now accepts RFC3339 and other date-time layouts, keeping the calendar date and discarding the time-of-day.
- **`DateTime.MarshalJSON`** — now serializes with `time.RFC3339Nano` to preserve sub-second precision.
- **Request body merging** — request body properties are now merged into JSON bodies with deterministic key ordering, overriding same-named properties.

## [5.1.0] - 2026-09-15

### Added
- **`ExecutionConstraints`** — new type with `StartAfter` and `CompleteBefore` fields describing when an agent may execute a task after delivery.
- **`ExecutionConstraints`** field with getters and setters (`GetExecutionConstraints()` / `SetExecutionConstraints()`) added to `Task` and `TaskCreation`.
- **`ExplicitFieldsFromJSON`** — internal helper that computes the explicit-fields bitmask from a raw JSON object.

### Changed
- **`HandleExplicitFields`** — now supports wrapper structs whose fields shadow an embedded struct, preserving overridden serialization (such as custom date formats) when removing `omitempty`.
- **`DeliveryConstraints`** — field documentation clarified to describe Lattice delivery scheduling to the agent.

## [5.0.0] - 2026-09-04

### Breaking Changes
- **`client.Video.Video.*`** — video client access flattened; replace `client.Video.Video.ListEgressStreams(...)` with `client.Video.ListEgressStreams(...)`.
- **Video request/response types** — moved from the `video` package to the root `Lattice` package; replace `video.ListEgressStreamsRequest` (and `RtspSettings`, `SrtSettings`, `MpegTsSettings`, etc.) with their `Lattice.*` equivalents.

### Added
- **`DeliveryConstraints.RequireAcknowledgement`** — new optional boolean field with getter/setter accessors requiring an agent to acknowledge a task before it is considered delivered.
- **`DeliveryErrorCode.DELIVERY_ERROR_CODE_NOT_ACKNOWLEDGED`** — new enum value for tasks that fail acknowledgement.
- **`GoogleRPCStatus`** — new type modeling gRPC-style status errors with `Code`, `Message`, and `Details` fields plus accessors.
- **`PlatformSubcomponents`** — new group type with accessors and JSON serialization, exposed via the new `GroupDetails.PlatformSubcomponents` field.

### Changed
- **Per-endpoint error handling** — each operation now maps its specific HTTP status codes (e.g. `404 NotFoundError`, `413 ContentTooLargeError`, `507 InsufficientStorageError`) to typed errors instead of a shared error set.
- **MPEG-TS ingress documentation** — clarifies that MPEG-TS is supported only at the edge in closed networks and may be disabled in cloud deployments, with fallback to RTSP or SRT.

## [4.25.0] - 2026-09-03

**Added**

* client.Video — new video client (with video.Client and video.RawClient) wired into the root Client for managing live video streams.
* Egress stream operations — ListEgressStreams, CreateEgressStream, GetEgressStream, and DeleteEgressStream for full egress stream lifecycle management.
* Ingress stream operations — ListIngressStreams, CreateIngressStream, GetIngressStream, and DeleteIngressStream for full ingress stream lifecycle management.
* Video streaming types — request/response models including IngressStream, EgressStream, their wrappers, transport settings (MpegTsSettings, RtspSettings, SrtSettings and related ingress/egress types), and the IngressStreamStatus enum with lifecycle states.
* Video error types — BadRequestError, UnauthorizedError, ForbiddenError, NotFoundError, ConflictError, TooManyRequestsError, ServiceUnavailableError, and InternalServerError mapped to HTTP status codes.

## [4.24.0] - 2026-08-26

## [4.23.1] - 2026-08-21

## [4.23.0] - 2026-08-20

## [4.22.0] - 2026-07-29

## [4.21.0] - 2026-07-22

## [v4.20.0] - 2026-07-21

## [4.19.0] - 2026-07-16

## [v4.18.1] - 2026-07-14

