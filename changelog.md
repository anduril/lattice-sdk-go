# Changelog

## [5.0.0] - 2026-09-03

### Breaking Changes
- **`client.Video`** — video stream methods moved up one level from `client.Video.Video.X()` to `client.Video.X()`, so drop the extra `.Video` accessor (e.g. replace `client.Video.Video.ListEgressStreams(...)` with `client.Video.ListEgressStreams(...)`).

### Added
- **Video stream types** — ingress/egress management request and response types (`CreateIngressStreamRequest`/`Response`, `CreateEgressStreamRequest`/`Response`, plus matching `Get`/`List`/`Delete` types) along with `IngressStream` and `EgressStream` models.
- **Transport settings** — connection and transport types added, including `RtspSettings`, `SrtSettings`, `MpegTsSettings`, `RtspIngress`, `SrtIngress`, `MpegTsIngress`, `RtspEgress`, and `SrtEgress`.
- **`IngressStreamStatus`** — new lifecycle enum (`STREAM_STATUS_LIVE`, `STREAM_STATUS_INACTIVE`, etc.) with a `NewIngressStreamStatusFromString` helper.
- **`DeliveryConstraints.RequireAcknowledgement`** — new optional field with `GetRequireAcknowledgement()`/`SetRequireAcknowledgement()` accessors requiring agents to acknowledge task delivery.
- **`DeliveryErrorCode` value** — new `DELIVERY_ERROR_CODE_NOT_ACKNOWLEDGED` enum value, plus a new `PlatformSubcomponents` group detail type on `GroupDetails`.

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

