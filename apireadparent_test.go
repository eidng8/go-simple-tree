package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/eidng8/go-simple-tree/ent/schema"
	"github.com/eidng8/go-simple-tree/ent/simpletree"
)

func Test_ReadSimpleTreeParent_should_return_one_record(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.SimpleTree.UpdateOneID(uint32(2)).SetParentID(1).
		SaveX(context.Background())
	rec := entClient.SimpleTree.Query().Where(simpletree.ID(1)).
		OnlyX(context.Background())
	eaa := ReadSimpleTreeParent200JSONResponse{
		Id:        1,
		Name:      "name 0",
		CreatedAt: rec.CreatedAt,
	}
	bytes, err := jsoniter.Marshal(eaa)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(http.MethodGet, schema.BaseUri+"/2/parent", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ReadSimpleTreeParent_does_not_returns_deleted_record(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.SimpleTree.DeleteOneID(2).ExecX(context.Background())
	req, _ := http.NewRequest(http.MethodGet, schema.BaseUri+"/2/parent", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}

func Test_ReadSimpleTreeParent_does_not_return_deleted_parent(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.SimpleTree.DeleteOneID(1).ExecX(context.Background())
	req, _ := http.NewRequest(http.MethodGet, schema.BaseUri+"/2/parent", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}

func Test_ReadSimpleTreeParent_should_return_404_if_not_found(t *testing.T) {
	engine, _, res := setupGinTest(t)
	req, _ := http.NewRequest(
		http.MethodGet, schema.BaseUri+"/987654321/parent", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}
