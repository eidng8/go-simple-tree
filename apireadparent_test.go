package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eidng8/go-simple-tree/ent/item"
	"github.com/eidng8/go-simple-tree/ent/schema"
)

func Test_ReadItemParent_should_return_one_record(t *testing.T) {
	_, engine, entClient, res := setupGinTest(t)
	entClient.Item.UpdateOneID(uint32(2)).SetParentID(1).
		SaveX(context.Background())
	rec := entClient.Item.Query().Where(item.ID(1)).
		OnlyX(context.Background())
	eaa := ReadItemParent200JSONResponse{
		Id:        1,
		Name:      "name 0",
		CreatedAt: rec.CreatedAt,
	}
	bytes, err := json.Marshal(eaa)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(http.MethodGet, schema.BaseUri+"/2/parent", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	assert.JSONEq(t, expected, actual)
}

func Test_ReadItemParent_does_not_returns_deleted_record(t *testing.T) {
	_, engine, entClient, res := setupGinTest(t)
	entClient.Item.DeleteOneID(2).ExecX(context.Background())
	req, _ := http.NewRequest(http.MethodGet, schema.BaseUri+"/2/parent", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}

func Test_ReadItemParent_does_not_return_deleted_parent(t *testing.T) {
	_, engine, entClient, res := setupGinTest(t)
	entClient.Item.DeleteOneID(1).ExecX(context.Background())
	req, _ := http.NewRequest(http.MethodGet, schema.BaseUri+"/2/parent", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}

func Test_ReadItemParent_should_return_404_if_not_found(t *testing.T) {
	_, engine, _, res := setupGinTest(t)
	req, _ := http.NewRequest(
		http.MethodGet, schema.BaseUri+"/987654321/parent", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}
