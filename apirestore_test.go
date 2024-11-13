package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eidng8/go-simple-tree/ent"
	"github.com/eidng8/go-simple-tree/ent/schema"
	"github.com/eidng8/go-simple-tree/ent/simpletree"
)

func Test_RestoreSimpleTree_should_restore_by_id(t *testing.T) {
	engine, entClient, response := setupGinTest(t)
	entClient.SimpleTree.Query().Where(simpletree.ID(1)).
		OnlyX(context.Background())
	entClient.SimpleTree.DeleteOneID(1).ExecX(context.Background())
	_, err := entClient.SimpleTree.Query().Where(simpletree.ID(1)).
		Only(context.Background())
	assert.True(t, ent.IsNotFound(err))
	req, _ := http.NewRequest(http.MethodPost, schema.BaseUri+"/1/restore", nil)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusNoContent, response.Code)
	rec := entClient.SimpleTree.Query().Where(simpletree.ID(1)).
		OnlyX(context.Background())
	assert.Nil(t, rec.DeletedAt)
}

func Test_RestoreSimpleTree_reports_404_if_not_found(t *testing.T) {
	engine, _, res := setupGinTest(t)
	req, _ := http.NewRequest(
		http.MethodPost, schema.BaseUri+"/987654321/restore", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}
