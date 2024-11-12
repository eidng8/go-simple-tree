# go-simple-tree

A simple tree listing OpenAPI v3 microservice.


## Usage

1. Fork or download the package;
2. `go mod tidy` in `root` and `tools` dir;
3. Change constants in the following files to match your project (excerpts below):
4. Comment out stuffs at the bottom of `ent/schema/simpletree.go` following comments in the file;
5. Run `go generate` to generate the ent client, ignore errors during generation;
6. Bring back those lines commented in step 4;
7. Run `go generate` again, there's should be no error this time;
8. Manually `diff` and `merge` 3 files:
    - `ent/openapi.go` <=> `/openapi.go`
    - `tools/services.gen.go` <=> `/services.gen.go`
    - `tools/types.gen.go` <=> `/types.gen.go`
9. Update `/apirestore.go` to match the newly generate endpoint if necessary;
10. Run `go test` to verify;
11. Do anything with the new service.

```golang
// ent/schema/simpletree.go
package schema

import (
	// imports...
)

/* Change these 2 constants to match your project. */

// BaseUri is the base URI for the OpenAPI endpoints.
const BaseUri = "/your-endpoint-base-uri"

// TableName is the name of the table in the database.
const TableName = "your_database_table_name"

type SimpleTree struct {
	// ........
}
```

```golang
// tools/entc.go
package main

import (
	// imports...
)

/* Change these 2 constants to match your project. */

// BaseUri is the base URI for the OpenAPI endpoints.
const BaseUri = "/your-endpoint-base-uri"

// TableName is the name of the table in the database.
const TableName = "your_database_table_name"

func main() {
	// ........
}

// ...... snipped ......

func genSpec(s *ogen.Spec) {
	/* Change these information to match your project. */
	s.Info.SetTitle("Simple tree listing API").SetVersion("0.0.1").
		SetDescription("This is an API listing hierarchical tree data")
}
```

```golang
// main_test.go
package main

import (
	// imports...
)

/* Change these 2 constants to match your project. */

// BaseUri is the base URI for the OpenAPI endpoints.
const BaseUri = "/your-endpoint-base-uri"

// TableName is the name of the table in the database.
const TableName = "your_database_table_name"

func tests() {
	// ........
}
```
