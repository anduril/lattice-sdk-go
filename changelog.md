# Changelog

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

