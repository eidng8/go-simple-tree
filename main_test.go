package main

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	jitr "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"

	"github.com/eidng8/go-simple-tree/ent"
	"github.com/eidng8/go-simple-tree/ent/enttest"
)

var jsoniter = jitr.ConfigCompatibleWithStandardLibrary

func setupGinTest(tb testing.TB) (
	*gin.Engine, *ent.Client, *httptest.ResponseRecorder,
) {
	// assert.Nil(tb, os.Setenv("DB_DRIVER", "mysql"))
	// assert.Nil(tb, os.Setenv("DB_USER", "root"))
	// assert.Nil(tb, os.Setenv("DB_PASSWORD", "123456"))
	// assert.Nil(tb, os.Setenv("DB_HOST", "127.0.0.1:43306"))
	// assert.Nil(tb, os.Setenv("DB_NAME", "simple_tree"))
	entClient := enttest.Open(tb, "sqlite3", ":memory:?_fk=1")
	tb.Cleanup(
		func() {
			_ = entClient.Close()
		},
	)
	engine, err := newEngine(entClient)
	assert.Nil(tb, err)
	assert.Nil(tb, setup(engine, entClient))
	fixture(entClient)
	return engine, entClient, httptest.NewRecorder()
}

func fixture(client *ent.Client) {
	ctx := context.Background()
	items := make([]*ent.ItemCreate, 50)
	for i := range 50 {
		items[i] = client.Item.Create().SetName(fmt.Sprintf("name %d", i))
	}
	client.Item.CreateBulk(items...).SaveX(ctx)
}
