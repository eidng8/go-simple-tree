package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadItemWithResponse_returns_detail(t *testing.T) {
	setupTest(t)
	expected := fixture[9]
	c, err := NewClientWithResponses(testServer, WithHTTPClient(&http.Client{}))
	assert.Nil(t, err)
	res, err := c.ReadItemWithResponse(context.TODO(), expected.Id, nil)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode())
	assertJsonEquals(t, expected, res.JSON200)
}
