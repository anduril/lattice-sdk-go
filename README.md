# anduril-sdk-go

## Quick Start

```bash
mkdir cmd
touch cmd/main.go
go mod init anduril.com/demo/v2
go env -w GOPRIVATE=github.com/anduril/anduril-go
go get github.com/anduril/anduril-go
go mod vendor
```

main.go
```go
package main

import entitymanagerv1 "github.com/anduril/anduril-go/src/anduril/entitymanager/v1"

func main() {
	em := entitymanagerv1.EntityManagerAPI_PublishEntitiesClient()
}
```

