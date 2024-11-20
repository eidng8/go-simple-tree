package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func listItem(t *testing.T) *ListItemResponse {
	hc := http.Client{}
	c, err := NewClientWithResponses(testServer, WithHTTPClient(&hc))
	assert.Nil(t, err)
	res, err := c.ListItemWithResponse(context.TODO(), &ListItemParams{})
	assert.Nil(t, err)
	return res
}

func Test_ListItemWithResponse_returns_1st_page(t *testing.T) {
	setupTest(t)
	res := listItem(t)
	assert.Equal(t, 50, res.JSON200.Total)
	assert.Equal(t, 1, res.JSON200.CurrentPage)
	assert.Equal(t, 5, res.JSON200.LastPage)
	assert.Equal(t, 1, res.JSON200.From)
	assert.Equal(t, 10, res.JSON200.To)
	assert.Equal(t, 10, res.JSON200.PerPage)
	assert.Equal(t, baseURL, res.JSON200.Path)
	assert.Equal(
		t, baseURL+"?page=1&per_page=10", res.JSON200.FirstPageUrl,
	)
	assert.Equal(
		t, baseURL+"?page=5&per_page=10", res.JSON200.LastPageUrl,
	)
	assert.Equal(
		t, baseURL+"?page=2&per_page=10", res.JSON200.NextPageUrl,
	)
	assert.Empty(t, res.JSON200.PrevPageUrl)
	assert.Equal(t, 10, len(res.JSON200.Data))
	assertJsonEquals(t, fixture[:10], res.JSON200.Data)
}

func Test_ListItemWithResponse_returns_2nd_page(t *testing.T) {
	setupTest(t)
	hc := http.Client{}
	c, err := NewClientWithResponses(testServer, WithHTTPClient(&hc))
	assert.Nil(t, err)
	u2 := 2
	res, err := c.ListItemWithResponse(
		context.TODO(), &ListItemParams{Page: &u2},
	)
	assert.Nil(t, err)
	assert.Equal(t, 50, res.JSON200.Total)
	assert.Equal(t, 2, res.JSON200.CurrentPage)
	assert.Equal(t, 5, res.JSON200.LastPage)
	assert.Equal(t, 11, res.JSON200.From)
	assert.Equal(t, 20, res.JSON200.To)
	assert.Equal(t, 10, res.JSON200.PerPage)
	assert.Equal(t, baseURL, res.JSON200.Path)
	assert.Equal(
		t, baseURL+"?page=1&per_page=10", res.JSON200.FirstPageUrl,
	)
	assert.Equal(
		t, baseURL+"?page=5&per_page=10", res.JSON200.LastPageUrl,
	)
	assert.Equal(
		t, baseURL+"?page=3&per_page=10", res.JSON200.NextPageUrl,
	)
	assert.Equal(
		t, baseURL+"?page=1&per_page=10", res.JSON200.PrevPageUrl,
	)
	assert.Equal(t, 10, len(res.JSON200.Data))
	assertJsonEquals(t, fixture[10:20], res.JSON200.Data)
}

func Test_ListItemWithResponse_returns_2nd_page_5_per_page(t *testing.T) {
	setupTest(t)
	hc := http.Client{}
	c, err := NewClientWithResponses(testServer, WithHTTPClient(&hc))
	assert.Nil(t, err)
	u2 := 2
	u5 := 5
	res, err := c.ListItemWithResponse(
		context.TODO(), &ListItemParams{Page: &u2, PerPage: &u5},
	)
	assert.Nil(t, err)
	assert.Equal(t, 50, res.JSON200.Total)
	assert.Equal(t, 2, res.JSON200.CurrentPage)
	assert.Equal(t, 10, res.JSON200.LastPage)
	assert.Equal(t, 6, res.JSON200.From)
	assert.Equal(t, 10, res.JSON200.To)
	assert.Equal(t, 5, res.JSON200.PerPage)
	assert.Equal(t, baseURL, res.JSON200.Path)
	assert.Equal(
		t, baseURL+"?page=1&per_page=5", res.JSON200.FirstPageUrl,
	)
	assert.Equal(
		t, baseURL+"?page=10&per_page=5", res.JSON200.LastPageUrl,
	)
	assert.Equal(
		t, baseURL+"?page=3&per_page=5", res.JSON200.NextPageUrl,
	)
	assert.Equal(
		t, baseURL+"?page=1&per_page=5", res.JSON200.PrevPageUrl,
	)
	assert.Equal(t, 5, len(res.JSON200.Data))
	assertJsonEquals(t, fixture[5:10], res.JSON200.Data)
}

func Test_ListItemWithResponse_returns_all(t *testing.T) {
	setupTest(t)
	hc := http.Client{}
	c, err := NewClientWithResponses(testServer, WithHTTPClient(&hc))
	assert.Nil(t, err)
	u100 := 100
	res, err := c.ListItemWithResponse(
		context.TODO(), &ListItemParams{PerPage: &u100},
	)
	assert.Nil(t, err)
	assert.Equal(t, 50, res.JSON200.Total)
	assert.Equal(t, 1, res.JSON200.CurrentPage)
	assert.Equal(t, 1, res.JSON200.LastPage)
	assert.Equal(t, 1, res.JSON200.From)
	assert.Equal(t, 50, res.JSON200.To)
	assert.Equal(t, 100, res.JSON200.PerPage)
	assert.Equal(t, baseURL, res.JSON200.Path)
	assert.Equal(
		t, baseURL+"?page=1&per_page=100", res.JSON200.FirstPageUrl,
	)
	assert.Empty(t, res.JSON200.LastPageUrl)
	assert.Empty(t, res.JSON200.NextPageUrl)
	assert.Empty(t, res.JSON200.PrevPageUrl)
	assert.Equal(t, 50, len(res.JSON200.Data))
	assertJsonEquals(t, fixture, res.JSON200.Data)
}

func Test_ListItemWithResponse_returns_specified_name_prefix(t *testing.T) {
	setupTest(t)
	hc := http.Client{}
	c, err := NewClientWithResponses(testServer, WithHTTPClient(&hc))
	assert.Nil(t, err)
	name := "name 1"
	res, err := c.ListItemWithResponse(
		context.TODO(), &ListItemParams{Name: &name},
	)
	assert.Nil(t, err)
	assert.Equal(t, 11, res.JSON200.Total)
	assert.Equal(t, 1, res.JSON200.CurrentPage)
	assert.Equal(t, 2, res.JSON200.LastPage)
	assert.Equal(t, 1, res.JSON200.From)
	assert.Equal(t, 10, res.JSON200.To)
	assert.Equal(t, 10, res.JSON200.PerPage)
	assert.Equal(t, baseURL, res.JSON200.Path)
	assert.Equal(
		t, baseURL+"?name=name+1&page=1&per_page=10", res.JSON200.FirstPageUrl,
	)
	assert.Equal(
		t, baseURL+"?name=name+1&page=2&per_page=10", res.JSON200.LastPageUrl,
	)
	assert.Equal(
		t, baseURL+"?name=name+1&page=2&per_page=10", res.JSON200.NextPageUrl,
	)
	assert.Empty(t, res.JSON200.PrevPageUrl)
	assert.Equal(t, 10, len(res.JSON200.Data))
	assertJsonEquals(
		t, append(append([]ItemList{}, fixture[:1]...), fixture[9:18]...),
		res.JSON200.Data,
	)
}

func Test_ListItemWithResponse_returns_specified_no_record(t *testing.T) {
	setupTest(t)
	hc := http.Client{}
	c, err := NewClientWithResponses(testServer, WithHTTPClient(&hc))
	assert.Nil(t, err)
	name := "not exist"
	res, err := c.ListItemWithResponse(
		context.TODO(), &ListItemParams{Name: &name},
	)
	assert.Nil(t, err)
	assert.Equal(t, 0, res.JSON200.Total)
	assert.Equal(t, 1, res.JSON200.CurrentPage)
	assert.Equal(t, 1, res.JSON200.LastPage)
	assert.Equal(t, 0, res.JSON200.From)
	assert.Equal(t, 0, res.JSON200.To)
	assert.Equal(t, 10, res.JSON200.PerPage)
	assert.Equal(t, baseURL, res.JSON200.Path)
	assert.Equal(
		t, baseURL+"?name=not+exist&page=1&per_page=10",
		res.JSON200.FirstPageUrl,
	)
	assert.Empty(t, res.JSON200.LastPageUrl)
	assert.Empty(t, res.JSON200.NextPageUrl)
	assert.Empty(t, res.JSON200.PrevPageUrl)
	assert.Empty(t, res.JSON200.Data)
}
