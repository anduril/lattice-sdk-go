# Lattice SDK Go Library

![](https://www.anduril.com/lattice-sdk/)


The Lattice SDK Go library provides convenient access to the Lattice SDK APIs from Go.

## Table of Contents

- [Documentation](#documentation)
- [Requirements](#requirements)
- [Installation](#installation)
- [Support](#support)
- [Reference](#reference)
- [Usage](#usage)
- [Environments](#environments)
- [Oauth](#oauth)
- [Errors](#errors)
- [Request Options](#request-options)
- [Advanced](#advanced)
  - [Response Headers](#response-headers)
  - [Retries](#retries)
  - [Timeouts](#timeouts)
  - [Explicit Null](#explicit-null)

## Documentation

API reference documentation is available [here](https://developer.anduril.com/).

## Requirements

To use the SDK please ensure you have the following installed:

* [Go](https://go.dev/doc/install)

## Installation

Install the Lattice SDK package with

```
go get github.com/anduril/lattice-sdk-go/v4
```

## Support

For support with this library please reach out to your Anduril representative. 

## Reference

A full reference for this library is available [here](https://github.com/anduril/lattice-sdk-go/blob/HEAD/./reference.md).

## Usage

Instantiate and use the client with the following:

```go
package example

import (
    client "github.com/anduril/lattice-sdk-go/client"
    option "github.com/anduril/lattice-sdk-go/option"
    Lattice "github.com/anduril/lattice-sdk-go"
    context "context"
)

func do() {
    client := client.NewClient(
        option.WithClientCredentials(
            "<clientId>",
            "<clientSecret>",
        ),
    )
    request := &Lattice.EntityEventRequest{
        SessionToken: "sessionToken",
    }
    client.Entities.LongPollEntityEvents(
        context.TODO(),
        request,
    )
}
```

## Environments

You can choose between different environments by using the `option.WithBaseURL` option. You can configure any arbitrary base
URL, which is particularly useful in test environments.

```go
client := client.NewClient(
    option.WithBaseURL(Lattice.Environments.Default),
)
```

## Oauth

This SDK supports OAuth 2.0 authentication. You have two options for providing credentials:

**Option 1: Client Credentials** - Provide your client ID and secret, and the SDK will automatically handle
token fetching and refresh:

**Option 2: Direct Token** - If you already have an access token (e.g., obtained through your own OAuth flow),
you can provide it directly:

```go
// Option 1: Use client credentials (SDK will handle token fetching and refresh)
client := client.NewClient(
    option.WithClientCredentials(
        "<YOUR_CLIENT_ID>",
        "<YOUR_CLIENT_SECRET>",
    ),
)

// Option 2: Use a pre-fetched token directly
client := client.NewClient(
    option.WithToken("<YOUR_ACCESS_TOKEN>"),
)
```

## Errors

Structured error types are returned from API calls that return non-success status codes. These errors are compatible
with the `errors.Is` and `errors.As` APIs, so you can access the error like so:

```go
response, err := client.Entities.LongPollEntityEvents(...)
if err != nil {
    var apiError *core.APIError
    if errors.As(err, apiError) {
        // Do something with the API error ...
    }
    return err
}
```

## Request Options

A variety of request options are included to adapt the behavior of the library, which includes configuring
authorization tokens, or providing your own instrumented `*http.Client`.

These request options can either be
specified on the client so that they're applied on every request, or for an individual request, like so:

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

```go
// Specify default options applied on every request.
client := client.NewClient(
    option.WithToken("<YOUR_API_KEY>"),
    option.WithHTTPClient(
        &http.Client{
            Timeout: 5 * time.Second,
        },
    ),
)

// Specify options for an individual request.
response, err := client.Entities.LongPollEntityEvents(
    ...,
    option.WithToken("<YOUR_API_KEY>"),
)
```

## Advanced

### Response Headers

You can access the raw HTTP response data by using the `WithRawResponse` field on the client. This is useful
when you need to examine the response headers received from the API call. (When the endpoint is paginated,
the raw HTTP response data will be included automatically in the Page response object.)

```go
response, err := client.Entities.WithRawResponse.LongPollEntityEvents(...)
if err != nil {
    return err
}
fmt.Printf("Got response headers: %v", response.Header)
fmt.Printf("Got status code: %d", response.StatusCode)
```

### Retries

The SDK is instrumented with automatic retries with exponential backoff. A request will be retried as long
as the request is deemed retryable and the number of retry attempts has not grown larger than the configured
retry limit (default: 2).

A request is deemed retryable when any of the following HTTP status codes is returned:

- [408](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408) (Timeout)
- [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) (Too Many Requests)
- [5XX](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500) (Internal Server Errors)

If the `Retry-After` header is present in the response, the SDK will prioritize respecting its value exactly
over the default exponential backoff.

Use the `option.WithMaxAttempts` option to configure this behavior for the entire client or an individual request:

```go
client := client.NewClient(
    option.WithMaxAttempts(1),
)

response, err := client.Entities.LongPollEntityEvents(
    ...,
    option.WithMaxAttempts(1),
)
```

### Timeouts

Setting a timeout for each individual request is as simple as using the standard context library. Setting a one second timeout for an individual API call looks like the following:

```go
ctx, cancel := context.WithTimeout(ctx, time.Second)
defer cancel()

response, err := client.Entities.LongPollEntityEvents(ctx, ...)
```

### Explicit Null

If you want to send the explicit `null` JSON value through an optional parameter, you can use the setters\
that come with every object. Calling a setter method for a property will flip a bit in the `explicitFields`
bitfield for that setter's object; during serialization, any property with a flipped bit will have its
omittable status stripped, so zero or `nil` values will be sent explicitly rather than omitted altogether:

```go
type ExampleRequest struct {
    // An optional string parameter.
    Name *string `json:"name,omitempty" url:"-"`

    // Private bitmask of fields set to an explicit value and therefore not to be omitted
    explicitFields *big.Int `json:"-" url:"-"`
}

request := &ExampleRequest{}
request.SetName(nil)

response, err := client.Entities.LongPollEntityEvents(ctx, request, ...)
```
