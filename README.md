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
8. Run `go test` to verify;
9. Do anything with the new service.

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
