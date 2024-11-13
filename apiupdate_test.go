package main

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/eidng8/go-simple-tree/ent/item"
	"github.com/eidng8/go-simple-tree/ent/schema"
)

func Test_UpdateItem_updates_existing_record(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	body := `{"name":"test name","abbr":"test abbr","parent_id":1}`
	req, _ := http.NewRequest(
		http.MethodPatch, schema.BaseUri+"/2",
		io.NopCloser(strings.NewReader(body)),
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusCreated, res.Code)
	actual := res.Body.String()
	aa := entClient.Item.Query().Where(item.NameEQ("test name")).
		Where(item.IDEQ(2)).OnlyX(context.Background())
	pid := uint32(1)
	b, err := jsoniter.Marshal(
		CreateItem201JSONResponse{
			Id:        2,
			ParentId:  &pid,
			Name:      "test name",
			CreatedAt: aa.CreatedAt,
			UpdatedAt: aa.UpdatedAt,
		},
	)
	assert.Nil(t, err)
	expected := string(b)
	require.JSONEq(t, expected, actual)
}

func Test_UpdateItem_reports_404_if_update_deleted_record(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.Item.UpdateOneID(2).SetDeletedAt(time.Now()).ExecX(context.Background())
	body := `{"name":"test name","abbr":"test abbr","parent_id":1}`
	req, _ := http.NewRequest(
		http.MethodPatch, schema.BaseUri+"/2",
		io.NopCloser(strings.NewReader(body)),
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}

func Test_UpdateItem_reports_422_if_request_body_invalid(t *testing.T) {
	engine, _, res := setupGinTest(t)
	body := `{"name":"a","parent_id":1}`
	req, _ := http.NewRequest(
		http.MethodPatch, schema.BaseUri+"/2",
		io.NopCloser(strings.NewReader(body)),
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
}

func Test_UpdateItem_reports_422_if_parentId_equals_self(t *testing.T) {
	engine, _, res := setupGinTest(t)
	body := `{"parent_id":1}`
	req, _ := http.NewRequest(
		http.MethodPatch, schema.BaseUri+"/1",
		io.NopCloser(strings.NewReader(body)),
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
}
