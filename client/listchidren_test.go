package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func listChildren(t *testing.T, id uint32) *ListItemChildrenResponse {
	hc := http.Client{}
	c, err := NewClientWithResponses(testServer, WithHTTPClient(&hc))
	assert.Nil(t, err)
	res, err := c.ListItemChildrenWithResponse(
		context.TODO(), id, &ListItemChildrenParams{},
	)
	assert.Nil(t, err)
	return res
}

func childrenUrl(id uint32) string {
	return baseURL + "/8/children"
}

func Test_ListChildrenWithResponse_returns_1st_page(t *testing.T) {
	setupTest(t)
	res := listChildren(t, 8)
	assert.Equal(t, 11, res.JSON200.Total)
	assert.Equal(t, 1, res.JSON200.CurrentPage)
	assert.Equal(t, 2, res.JSON200.LastPage)
	assert.Equal(t, 1, res.JSON200.From)
	assert.Equal(t, 10, res.JSON200.To)
	assert.Equal(t, 10, res.JSON200.PerPage)
	assert.Equal(t, baseURL+"/8/children", res.JSON200.Path)
	assert.Equal(
		t, childrenUrl(8)+"?page=1&per_page=10", res.JSON200.FirstPageUrl,
	)
	assert.Equal(
		t, childrenUrl(8)+"?page=2&per_page=10", res.JSON200.LastPageUrl,
	)
	assert.Equal(
		t, childrenUrl(8)+"?page=2&per_page=10", res.JSON200.NextPageUrl,
	)
	assert.Empty(t, res.JSON200.PrevPageUrl)
	assert.Equal(t, 10, len(res.JSON200.Data))
	assertJsonEquals(t, fixture[39:49], res.JSON200.Data)
}
