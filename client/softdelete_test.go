package main

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_DeleteItemWithResponse_RestoreItemWithResponse(t *testing.T) {
	justnow := time.Now().Add(-time.Second)
	setupTest(t)
	c, err := NewClientWithResponses(testServer, WithHTTPClient(&http.Client{}))
	assert.Nil(t, err)
	res, err := c.DeleteItemWithResponse(context.TODO(), 6, nil)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, res.StatusCode())
	expected := append([]ItemList{}, fixture[:5]...)
	expected = append(expected, fixture[6:11]...)
	assertJsonEquals(t, expected, listItem(t).JSON200.Data)
	c, err = NewClientWithResponses(testServer, WithHTTPClient(&http.Client{}))
	assert.Nil(t, err)
	rest, err := c.RestoreItemWithResponse(context.TODO(), 6)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, rest.StatusCode())
	expected = append([]ItemList{}, fixture[:10]...)
	list := listItem(t)
	assert.NotNil(t, list.JSON200.Data[5].UpdatedAt)
	assert.True(t, list.JSON200.Data[5].UpdatedAt.After(justnow))
	expected[5].UpdatedAt = list.JSON200.Data[5].UpdatedAt
	assertJsonEquals(t, expected, list.JSON200.Data)
}
