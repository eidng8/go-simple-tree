package main

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CreateItemWithResponse_returns_detail(t *testing.T) {
	setupTest(t)
	c, err := NewClientWithResponses(testServer, WithHTTPClient(&http.Client{}))
	assert.Nil(t, err)
	u1 := uint32(1)
	name := fmt.Sprintf("new %s", time.Now().Format(time.RFC3339))
	body := CreateItemJSONRequestBody{Name: name, ParentId: &u1}
	res, err := c.CreateItemWithResponse(context.TODO(), body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode())
	assert.Equal(t, name, res.JSON200.Name)
	assert.Equal(t, uint32(1), *res.JSON200.ParentId)
	assert.Nil(t, res.JSON200.UpdatedAt)
}
