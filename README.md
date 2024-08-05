# anduril-sdk-go

## Quick Start

```bash
mkdir cmd
touch cmd/main.go
go mod init anduril.com/demo/v2
go env -w GOPRIVATE=github.com/dogun-anduril/anduril-sdk-go
go get github.com/dogun-anduril/anduril-sdk-go
go mod vendor
```

main.go
```go
package main

import entitymanagerv1 "github.com/dogun-anduril/anduril-sdk-go/src/anduril/entitymanager/v1"

func main() {
	em := entitymanagerv1.EntityManagerAPI_PublishEntitiesClient()
}
```

