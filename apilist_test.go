package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/eidng8/go-ent/paginate"
	"github.com/eidng8/go-ent/softdelete"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/eidng8/go-simple-tree/ent/item"
	"github.com/eidng8/go-simple-tree/ent/schema"
)

func Test_ListItem_should_return_1st_page(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	rows := entClient.Item.Query().Order(item.ByID()).Limit(10).
		AllX(context.Background())
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	page := paginate.PaginatedList[Item]{
		Total:        50,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     5,
		FirstPageUrl: "http://127.0.0.1" + schema.BaseUri + "?page=1&per_page=10",
		LastPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?page=5&per_page=10",
		NextPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?page=2&per_page=10",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1" + schema.BaseUri,
		From:         1,
		To:           10,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1"+schema.BaseUri, nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListItem_should_return_4th_page(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	rows := entClient.Item.Query().Order(item.ByID()).Limit(10).
		Offset(30).AllX(context.Background())
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	page := paginate.PaginatedList[Item]{
		Total:        50,
		PerPage:      10,
		CurrentPage:  4,
		LastPage:     5,
		FirstPageUrl: "http://127.0.0.1" + schema.BaseUri + "?page=1&per_page=10",
		LastPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?page=5&per_page=10",
		NextPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?page=5&per_page=10",
		PrevPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?page=3&per_page=10",
		Path:         "http://127.0.0.1" + schema.BaseUri,
		From:         31,
		To:           40,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1"+schema.BaseUri+"?page=4", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListItem_should_return_all_records(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	rows := entClient.Item.Query().Order(item.ByID()).
		AllX(context.Background())
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	page := paginate.PaginatedList[Item]{
		Total:        50,
		PerPage:      12345,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: "http://127.0.0.1" + schema.BaseUri + "?page=1&per_page=12345",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1" + schema.BaseUri,
		From:         1,
		To:           50,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1"+schema.BaseUri+"?per_page=12345",
		nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListItem_should_return_2nd_page_exclude_deleted(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.Item.Delete().
		Where(item.Or(item.IDIn(5, 3, 21))).
		ExecX(context.Background())
	rows := entClient.Item.Query().Order(item.ByID()).
		Where(item.And(item.IDNotIn(5, 3, 21))).
		Offset(10).Limit(10).AllX(context.Background())
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	page := paginate.PaginatedList[Item]{
		Total:        47,
		PerPage:      10,
		CurrentPage:  2,
		LastPage:     5,
		FirstPageUrl: "http://127.0.0.1" + schema.BaseUri + "?page=1&per_page=10",
		LastPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?page=5&per_page=10",
		NextPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?page=3&per_page=10",
		PrevPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?page=1&per_page=10",
		Path:         "http://127.0.0.1" + schema.BaseUri,
		From:         11,
		To:           20,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1"+schema.BaseUri+"?page=2", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListItem_should_return_2nd_page_include_deleted(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.Item.Delete().
		Where(item.IDIn(5, 3, 11)).
		ExecX(context.Background())
	rows := entClient.Item.Query().Order(item.ByID()).
		Where(item.IDLTE(20)).
		Offset(10).Limit(10).
		AllX(softdelete.IncludeTrashed(context.Background()))
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	page := paginate.PaginatedList[Item]{
		Total:        50,
		PerPage:      10,
		CurrentPage:  2,
		LastPage:     5,
		FirstPageUrl: "http://127.0.0.1" + schema.BaseUri + "?page=1&per_page=10&trashed=1",
		LastPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?page=5&per_page=10&trashed=1",
		NextPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?page=3&per_page=10&trashed=1",
		PrevPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?page=1&per_page=10&trashed=1",
		Path:         "http://127.0.0.1" + schema.BaseUri,
		From:         11,
		To:           20,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1"+schema.BaseUri+"?page=2&trashed=1",
		nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListItem_should_return_all_records_exclude_deleted(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.Item.Delete().
		Where(item.IDIn(5, 3, 21)).
		ExecX(context.Background())
	rows := entClient.Item.Query().Order(item.ByID()).
		Where(item.IDNotIn(5, 3, 21)).
		AllX(context.Background())
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	page := paginate.PaginatedList[Item]{
		Total:        47,
		PerPage:      12345,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: "http://127.0.0.1" + schema.BaseUri + "?page=1&per_page=12345",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1" + schema.BaseUri,
		From:         1,
		To:           47,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1"+schema.BaseUri+"?per_page=12345",
		nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListItem_should_return_4th_page_5_per_page(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	rows := entClient.Item.Query().Order(item.ByID()).Limit(5).
		Offset(15).AllX(context.Background())
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	last := 10
	page := paginate.PaginatedList[Item]{
		Total:        50,
		PerPage:      5,
		CurrentPage:  4,
		LastPage:     last,
		FirstPageUrl: "http://127.0.0.1" + schema.BaseUri + "?page=1&per_page=5",
		LastPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?page=10&per_page=5",
		NextPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?page=5&per_page=5",
		PrevPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?page=3&per_page=5",
		Path:         "http://127.0.0.1" + schema.BaseUri,
		From:         16,
		To:           20,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1"+schema.BaseUri+"?page=4&per_page=5",
		nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListItem_should_return_specified_name_prefix(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	rows := entClient.Item.Query().Order(item.ByID()).
		Where(item.NameHasPrefix("name 1")).Limit(10).
		AllX(context.Background())
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	page := paginate.PaginatedList[Item]{
		Total:        11,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     2,
		FirstPageUrl: "http://127.0.0.1" + schema.BaseUri + "?name=name+1&page=1&per_page=10",
		LastPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?name=name+1&page=2&per_page=10",
		NextPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?name=name+1&page=2&per_page=10",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1" + schema.BaseUri,
		From:         1,
		To:           10,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1"+schema.BaseUri+"?name=name%201", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListItem_should_apply_all_filter(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.Item.DeleteOneID(1).ExecX(context.Background())
	rows := entClient.Item.Query().Order(item.ByID()).Limit(10).
		Where(item.NameHasPrefix("name 1")).
		AllX(softdelete.IncludeTrashed(context.Background()))
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	page := paginate.PaginatedList[Item]{
		Total:        11,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     2,
		FirstPageUrl: "http://127.0.0.1" + schema.BaseUri + "?name=name+1&page=1&per_page=10",
		LastPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?name=name+1&page=2&per_page=10",
		NextPageUrl:  "http://127.0.0.1" + schema.BaseUri + "?name=name+1&page=2&per_page=10",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1" + schema.BaseUri,
		From:         1,
		To:           10,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet,
		"http://127.0.0.1"+schema.BaseUri+"?name=name+1", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListItem_should_return_no_record(t *testing.T) {
	engine, _, res := setupGinTest(t)
	page := paginate.PaginatedList[Item]{
		Total:        0,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: "http://127.0.0.1" + schema.BaseUri + "?name=not+exist&page=1&per_page=10",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1" + schema.BaseUri,
		From:         0,
		To:           0,
		Data:         []*Item{},
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet,
		"http://127.0.0.1"+schema.BaseUri+"?name=not+exist", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListItem_should_report_400_for_invalid_page(t *testing.T) {
	engine, _, res := setupGinTest(t)
	req, _ := http.NewRequest(http.MethodGet, schema.BaseUri+"?page=a", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func Test_ListItem_should_report_400_for_invalid_perPage(t *testing.T) {
	engine, _, res := setupGinTest(t)
	req, _ := http.NewRequest(http.MethodGet, schema.BaseUri+"?per_page=a", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code)
}
