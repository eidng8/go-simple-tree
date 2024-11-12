package main

import (
	"context"
	"log"
	"os"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/eidng8/go-db"
	"github.com/gin-gonic/gin"

	"github.com/eidng8/go-simple-tree/ent"
	"github.com/eidng8/go-simple-tree/ent/migrate"
	_ "github.com/eidng8/go-simple-tree/ent/runtime"
)

func main() {
	entClient := getEntClient()
	mode := getenvd("SERVER_MODE", gin.ReleaseMode)
	engine, err := newEngine(mode, entClient)
	if err != nil {
		log.Fatalf("Failed to create server: %s", err)
	}
	defer func(entClient *ent.Client) {
		err := entClient.Close()
		if err != nil {
			log.Fatalf("Failed to close ent client: %s", err)
		}
	}(entClient)
	err = setup(engine, entClient)
	if err != nil {
		return
	}
	if err = engine.Run(getenv("LISTEN_ADDR")); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

func setup(gc *gin.Engine, ec *ent.Client) error {
	// Just make sure we have a basic empty db to work with.
	// Import data to db to fully use the API.
	// Or remove this auto-migration and use your own.
	return ec.Schema.Create(
		context.Background(), migrate.WithDropIndex(true),
		migrate.WithDropColumn(true), migrate.WithForeignKeys(true),
	)
}

func getEntClient() *ent.Client {
	return ent.NewClient(ent.Driver(entsql.OpenDB(db.ConnectX())))
}

func getenv(name string) string {
	val := os.Getenv(name)
	if "" == val {
		log.Fatalf("%v environment variable is not set", name)
	}
	return val
}

func getenvd(name, defval string) string {
	val := os.Getenv(name)
	if "" == val {
		return defval
	}
	return val
}
