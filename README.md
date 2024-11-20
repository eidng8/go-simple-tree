# go-simple-tree

A simple hierarchical tree listing OpenAPI v3 microservice.

Every item has one parent at maximum, and can have multiple children.
The service is generated using `ent`, `entoas` and `oapi-codegen` tools.


## Usage

1. Fork or download the package;
2. `go mod tidy` in `root` and `tools` dir;
3. Change constants in the following files to match your project (excerpts below):
4. Comment out stuffs at the bottom of `ent/schema/item.go` following comments in the file;
5. Run `go generate` to generate the ent client, ignore errors during generation;
6. Bring back those lines commented in step 4;
7. Run `go generate` again, there's should be no error this time;
8. Manually `diff` and `merge` 3 files:
    - `ent/openapi.go` <=> `/openapi.go`
    - `tools/services.gen.go` <=> `/services.gen.go`
    - `tools/types.gen.go` <=> `/types.gen.go`
9. Update any file to match the newly generate endpoint if necessary;
10. Run `go test` to verify;
11. Do anything with the new service.

> Manually merging the files would be a better way to avoid messing up the existing code.

```golang
// ent/schema/item.go
package schema

import (
    // imports...
)

/* Change these 2 constants to match your project. */

// BaseUri is the base URI for the OpenAPI endpoints.
const BaseUri = "/your-endpoint-base-uri"

// TableName is the name of the table in the database.
const TableName = "your_database_table_name"

type Item struct {
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

## Environment Variables

The following environment variables are needed to stat the service.

#### BASE_URL

REQUIRED and cannot be empty. Determines the base URL for all URL generation. It should be the fully qualified URL without endpoint path.

#### LISTEN

REQUIRED and defaults to `:80`. Determines the gin server listening TCP network address. e.g. `localhost:8080`.

#### GIN_MODE

OPTIONAL and defaults to `release`. Can be one of `debug`, `test`, or `release`.

#### DB_DRIVER

REQUIRED and cannot be empty. Determines what kind of database to connect. Can be any driver supported by `database/sql`, such as `mysql`, `sqlite3`, `pgx`, etc. Remember to import proper driver module to your package.

#### DB_DSN

REQUIRED for connections other than MySQL. A complete DSN string to be used to establish connection to database. When set, this variable is passed directly to `sql.Open()`. For MySQL, if this variable is not set, variables below will be used to configure the connection.

#### Variables specific to MySQL

These variables are used to configure the connection to MySQL, if `DB_DSN` is not set.

##### DB_USER

REQUIRED and cannot be empty. Determines the username to connect to database.

##### DB_PASSWORD

REQUIRED and cannot be empty. Determines the password to connect to database.

##### DB_HOST

REQUIRED and cannot be empty. Determines the host to connect to.

##### DB_NAME

REQUIRED and cannot be empty. Determines the database name to connect to.

##### DB_PROTOCOL

OPTIONAL and defaults to `tcp`.

##### DB_COLLATION

OPTIONAL and defaults to `utf8mb4_unicode_ci`.

##### DB_TIMEZONE

OPTIONAL and defaults to `UTC`.
