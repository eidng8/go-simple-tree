//go:build ignore

package main

import (
	"context"
	"log"
	"os"

	"github.com/eidng8/go-simple-tree/ent/migrate"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create a local migration directory able to understand Atlas migration file format for replay.
	if _, err := os.Stat("ent/migrate/migrations"); os.IsNotExist(err) {
		err = os.Mkdir("ent/migrate/migrations", 0o755)
		if err != nil {
			log.Fatalf("failed creating migration directory: %v", err)
		}
	}
	dir, err := atlas.NewLocalDir("ent/migrate/migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.MySQL),           // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
	}
	if len(os.Args) != 2 {
		log.Fatalln("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
	}
	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	err = migrate.NamedDiff(
		context.Background(),
		"mysql://root:pass@localhost:54322/dev-simple-tree", os.Args[1],
		opts...,
	)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
