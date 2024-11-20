package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/eidng8/go-ent/paginate"
	"github.com/stretchr/testify/assert"

	"github.com/eidng8/go-simple-tree/ent/item"
	"github.com/eidng8/go-simple-tree/ent/schema"
)

func Test_ListItemChildren_should_return_1st_page(t *testing.T) {
	server, engine, entClient, res := setupGinTest(t)
	entClient.Item.Update().Where(item.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	rows := entClient.Item.Query().Order(item.ByID()).Limit(10).
		Where(item.ParentID(2)).AllX(context.Background())
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	page := paginate.PaginatedList[Item]{
		Total:        48,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     5,
		FirstPageUrl: server.BaseUrl() + "/2/children?page=1&per_page=10",
		LastPageUrl:  server.BaseUrl() + "/2/children?page=5&per_page=10",
		NextPageUrl:  server.BaseUrl() + "/2/children?page=2&per_page=10",
		PrevPageUrl:  "",
		Path:         server.BaseUrl() + "/2/children",
		From:         1,
		To:           10,
		Data:         list,
	}
	bytes, err := json.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, schema.BaseUri+"/2/children", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	assert.JSONEq(t, expected, actual)
}

func Test_ListItemChildren_should_return_4th_page(t *testing.T) {
	server, engine, entClient, res := setupGinTest(t)
	entClient.Item.Update().Where(item.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	rows := entClient.Item.Query().Order(item.ByID()).Limit(10).
		Offset(30).Where(item.IDGT(2)).AllX(context.Background())
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	page := paginate.PaginatedList[Item]{
		Total:        48,
		PerPage:      10,
		CurrentPage:  4,
		LastPage:     5,
		FirstPageUrl: server.BaseUrl() + "/2/children?page=1&per_page=10",
		LastPageUrl:  server.BaseUrl() + "/2/children?page=5&per_page=10",
		NextPageUrl:  server.BaseUrl() + "/2/children?page=5&per_page=10",
		PrevPageUrl:  server.BaseUrl() + "/2/children?page=3&per_page=10",
		Path:         server.BaseUrl() + "/2/children",
		From:         31,
		To:           40,
		Data:         list,
	}
	bytes, err := json.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, schema.BaseUri+"/2/children?page=4",
		nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	assert.JSONEq(t, expected, actual)
}

func Test_ListItemChildren_should_return_all_records(t *testing.T) {
	server, engine, entClient, res := setupGinTest(t)
	entClient.Item.Update().Where(item.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	rows := entClient.Item.Query().Order(item.ByID()).
		Where(item.IDGT(2)).AllX(context.Background())
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	page := paginate.PaginatedList[Item]{
		Total:        48,
		PerPage:      12345,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: server.BaseUrl() + "/2/children?page=1&per_page=12345",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         server.BaseUrl() + "/2/children",
		From:         1,
		To:           48,
		Data:         list,
	}
	bytes, err := json.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet,
		server.BaseUrl()+"/2/children?per_page=12345", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	assert.JSONEq(t, expected, actual)
}

func Test_ListItemChildren_should_return_2nd_page_exclude_deleted(t *testing.T) {
	server, engine, entClient, res := setupGinTest(t)
	entClient.Item.Update().Where(item.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	entClient.Item.Delete().
		Where(item.Or(item.IDIn(5, 3, 21))).
		ExecX(context.Background())
	rows := entClient.Item.Query().Order(item.ByID()).
		Where(item.And(item.IDNotIn(5, 3, 21))).
		Where(item.IDGT(2)).Offset(10).Limit(10).
		AllX(context.Background())
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	page := paginate.PaginatedList[Item]{
		Total:        45,
		PerPage:      10,
		CurrentPage:  2,
		LastPage:     5,
		FirstPageUrl: server.BaseUrl() + "/2/children?page=1&per_page=10",
		LastPageUrl:  server.BaseUrl() + "/2/children?page=5&per_page=10",
		NextPageUrl:  server.BaseUrl() + "/2/children?page=3&per_page=10",
		PrevPageUrl:  server.BaseUrl() + "/2/children?page=1&per_page=10",
		Path:         server.BaseUrl() + "/2/children",
		From:         11,
		To:           20,
		Data:         list,
	}
	bytes, err := json.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, schema.BaseUri+"/2/children?page=2",
		nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	assert.JSONEq(t, expected, actual)
}

func Test_ListItemChildren_should_return_all_records_exclude_deleted(t *testing.T) {
	server, engine, entClient, res := setupGinTest(t)
	entClient.Item.Update().Where(item.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	entClient.Item.Delete().Where(item.IDIn(5, 3, 21)).
		ExecX(context.Background())
	rows := entClient.Item.Query().Order(item.ByID()).
		Where(item.IDNotIn(5, 3, 21)).Where(item.ParentIDEQ(2)).
		AllX(context.Background())
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	page := paginate.PaginatedList[Item]{
		Total:        45,
		PerPage:      12345,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: server.BaseUrl() + "/2/children?page=1&per_page=12345",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         server.BaseUrl() + "/2/children",
		From:         1,
		To:           45,
		Data:         list,
	}
	bytes, err := json.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet,
		server.BaseUrl()+"/2/children?per_page=12345", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	assert.JSONEq(t, expected, actual)
}

func Test_ListItemChildren_should_return_4th_page_5_per_page(t *testing.T) {
	server, engine, entClient, res := setupGinTest(t)
	entClient.Item.Update().Where(item.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	rows := entClient.Item.Query().Order(item.ByID()).
		Where(item.IDGT(2)).Limit(5).Offset(15).
		AllX(context.Background())
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	page := paginate.PaginatedList[Item]{
		Total:        48,
		PerPage:      5,
		CurrentPage:  4,
		LastPage:     10,
		FirstPageUrl: server.BaseUrl() + "/2/children?page=1&per_page=5",
		LastPageUrl:  server.BaseUrl() + "/2/children?page=10&per_page=5",
		NextPageUrl:  server.BaseUrl() + "/2/children?page=5&per_page=5",
		PrevPageUrl:  server.BaseUrl() + "/2/children?page=3&per_page=5",
		Path:         server.BaseUrl() + "/2/children",
		From:         16,
		To:           20,
		Data:         list,
	}
	bytes, err := json.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet,
		server.BaseUrl()+"/2/children?page=4&per_page=5", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	assert.JSONEq(t, expected, actual)
}

func Test_ListItemChildren_should_return_specified_name_prefix(t *testing.T) {
	server, engine, entClient, res := setupGinTest(t)
	entClient.Item.Update().Where(item.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	rows := entClient.Item.Query().Order(item.ByID()).
		Where(item.ParentIDEQ(2)).Where(item.NameHasPrefix("name 1")).
		Limit(10).AllX(context.Background())
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	page := paginate.PaginatedList[Item]{
		Total:        10,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: server.BaseUrl() + "/2/children?name=name+1&page=1&per_page=10",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         server.BaseUrl() + "/2/children",
		From:         1,
		To:           10,
		Data:         list,
	}
	bytes, err := json.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet,
		server.BaseUrl()+"/2/children?name=name%201",
		nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	assert.JSONEq(t, expected, actual)
}

func Test_ListItemChildren_should_apply_all_filter(t *testing.T) {
	server, engine, entClient, res := setupGinTest(t)
	entClient.Item.Update().Where(item.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	entClient.Item.DeleteOneID(11).ExecX(context.Background())
	rows := entClient.Item.Query().Order(item.ByID()).
		Where(item.NameHasPrefix("name 1")).
		Where(item.ParentIDEQ(2)).
		Limit(10).AllX(context.Background())
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	page := paginate.PaginatedList[Item]{
		Total:        9,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: server.BaseUrl() + "/2/children?abbr=abbr+1&name=name+1&page=1&per_page=10",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         server.BaseUrl() + "/2/children",
		From:         1,
		To:           9,
		Data:         list,
	}
	bytes, err := json.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet,
		server.BaseUrl()+"/2/children?name=name+1&abbr=abbr%201",
		nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	assert.JSONEq(t, expected, actual)
}

func Test_ListItemChildren_should_return_no_record(t *testing.T) {
	server, engine, entClient, res := setupGinTest(t)
	entClient.Item.Update().Where(item.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	page := paginate.PaginatedList[Item]{
		Total:        0,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: server.BaseUrl() + "/2/children?name=not+exist&page=1&per_page=10",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         server.BaseUrl() + "/2/children",
		From:         0,
		To:           0,
		Data:         []*Item{},
	}
	bytes, err := json.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet,
		server.BaseUrl()+"/2/children?name=not+exist", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	assert.JSONEq(t, expected, actual)
}

func Test_ListItemChildren_should_report_400_for_invalid_page(t *testing.T) {
	_, engine, _, res := setupGinTest(t)
	req, _ := http.NewRequest(
		http.MethodGet, schema.BaseUri+"/2/children?page=a", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func Test_ListItemChildren_should_report_400_for_invalid_perPage(t *testing.T) {
	_, engine, _, res := setupGinTest(t)
	req, _ := http.NewRequest(
		http.MethodGet, schema.BaseUri+"?per_page=a", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func Test_ListItemChildren_should_return_all_descendants(t *testing.T) {
	server, engine, entClient, res := setupGinTest(t)
	entClient.Item.Update().SetParentID(1).
		Where(item.IDIn(2, 3)).ExecX(context.Background())
	entClient.Item.Update().SetParentID(2).
		Where(item.IDIn(4, 5, 6)).ExecX(context.Background())
	entClient.Item.Update().SetParentID(3).
		Where(item.IDIn(7, 8)).ExecX(context.Background())
	entClient.Item.Update().SetParentID(4).
		Where(item.IDIn(9, 10, 11, 12)).
		ExecX(context.Background())
	rows := entClient.Item.Query().Order(item.ByID()).
		Where(item.IDGT(1)).Where(item.IDLTE(12)).
		AllX(context.Background())
	list := make([]*Item, len(rows))
	for i, row := range rows {
		list[i] = newItemFromEnt(row)
	}
	page := paginate.PaginatedList[Item]{
		Total:        11,
		PerPage:      11,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: server.BaseUrl() + "/1/children?recurse=1",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         server.BaseUrl() + "/1/children",
		From:         1,
		To:           11,
		Data:         list,
	}
	bytes, err := json.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet,
		server.BaseUrl()+"/1/children?recurse=1",
		nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	assert.JSONEq(t, expected, actual)
}
