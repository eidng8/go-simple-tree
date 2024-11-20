package main

import (
	"context"
	stdsql "database/sql"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/eidng8/go-simple-tree/ent/schema"
)

func Test_DeleteItem_should_delete_by_id(t *testing.T) {
	_, engine, entClient, response := setupGinTest(t)
	req, _ := http.NewRequest(http.MethodDelete, schema.BaseUri+"/1", nil)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusNoContent, response.Code)
	rs, err := entClient.QueryContext(
		context.Background(),
		fmt.Sprintf(
			"SELECT `id`, `deleted_at` FROM `%s` WHERE `id` = 1",
			schema.TableName,
		),
	)
	defer func(rs *stdsql.Rows) {
		err := rs.Close()
		assert.Nil(t, err)
	}(rs)
	assert.Nil(t, err)
	assert.NotNil(t, rs)
	rs.Next()
	var id int
	var deletedAt stdsql.NullTime
	err = rs.Scan(&id, &deletedAt)
	assert.Equal(t, 1, id)
	assert.True(t, deletedAt.Valid)
}

func Test_DeleteItem_should_physically_delete_if_requested(t *testing.T) {
	_, engine, entClient, response := setupGinTest(t)
	entClient.Item.UpdateOneID(1).SetDeletedAt(time.Now()).
		ExecX(context.Background())
	req, _ := http.NewRequest(
		http.MethodDelete, schema.BaseUri+"/1?trashed=1", nil,
	)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusNoContent, response.Code)
	rs, err := entClient.QueryContext(
		context.Background(),
		fmt.Sprintf(
			"SELECT `id`, `deleted_at` FROM `%s` WHERE `id` = 1",
			schema.TableName,
		),
	)
	defer func(rs *stdsql.Rows) {
		err := rs.Close()
		assert.Nil(t, err)
	}(rs)
	assert.Nil(t, err)
	assert.NotNil(t, rs)
	assert.False(t, rs.Next())
}

func Test_DeleteItem_should_return_404_if_not_found(t *testing.T) {
	_, engine, _, res := setupGinTest(t)
	req, _ := http.NewRequest(
		http.MethodDelete, schema.BaseUri+"/987654321", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}
