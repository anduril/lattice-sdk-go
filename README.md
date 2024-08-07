# Anduril SDK GoLang

The official [Anduril](https://www.anduril.com/) client library.

## Requirements

[Go installed](https://go.dev/doc/install). Code samples were written in 1.22.5

## Installation

### Authentication

To authenticate with the Github package repository, you will need to generate a [personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-personal-access-token-classic). This should have at least `read:packages` scope. Please keep the token safe for the next stage of the setup procedure.

### Setup

```
mkdir cmd
touch cmd/main.go
go mod init anduril.com/demo/v2
go env -w GOPRIVATE=github.com/anduril/anduril-go
go get github.com/anduril/anduril-go
go mod vendor
```

## Usage
main.go

```go
package main

import (
	"context"
	"log"

	entitymanagerv1 "github.com/anduril/anduril-go/src/anduril/entitymanager/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// BearerTokenAuth supplies PerRPCCredentials from a given token.
type BearerTokenAuth struct {
	Token string
}

// GetRequestMetadata gets the current request metadata, adding the bearer token.
func (b *BearerTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + b.Token,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security.
func (b *BearerTokenAuth) RequireTransportSecurity() bool {
	return true // or false if you are developing/testing without TLS
}

func main() {
	bearerToken := "<YOUR BEARER TOKEN>"
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
		grpc.WithPerRPCCredentials(&BearerTokenAuth{Token: bearerToken}),
	}

	conn, err := grpc.NewClient("desert-guardian.anduril.com:443", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	em := entitymanagerv1.NewEntityManagerAPIClient(conn)

	ctx := context.Background()
	request := &entitymanagerv1.GetEntityRequest{
		EntityId: "<ENTITY ID>",
	}
	response, err := em.GetEntity(ctx, request)

	if err != nil {
		log.Fatalf("Error fetching entity: %v", err)
	}
	log.Printf("Response: %v", response)
}
```

## Support

For support with this library please [file an issue](https://github.com/anduril/anduril-go/issues/new) or reach out to your Anduril representative. 



