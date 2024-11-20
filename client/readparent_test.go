package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadParentWithResponse_returns_detail(t *testing.T) {
	setupTest(t)
	expected := fixture[0]
	c, err := NewClientWithResponses(testServer, WithHTTPClient(&http.Client{}))
	assert.Nil(t, err)
	res, err := c.ReadItemParentWithResponse(context.TODO(), 2)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode())
	assertJsonEquals(t, expected, res.JSON200)
}
