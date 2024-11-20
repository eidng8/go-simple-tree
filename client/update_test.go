package main

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_UpdateItemWithResponse_returns_detail(t *testing.T) {
	setupTest(t)
	justnow := time.Now().Add(-time.Second)
	fix := fixture[9]
	defer func() {
		c, err := NewClientWithResponses(
			testServer, WithHTTPClient(&http.Client{}),
		)
		body := UpdateItemJSONRequestBody{
			Name: &fix.Name, ParentId: fix.ParentId,
		}
		_, err = c.UpdateItemWithResponse(context.TODO(), fix.Id, body)
		assert.Nil(t, err)
	}()

	c, err := NewClientWithResponses(testServer, WithHTTPClient(&http.Client{}))
	assert.Nil(t, err)
	u1 := uint32(1)
	name := fmt.Sprintf("update %s", time.Now().Format(time.RFC3339))
	body := UpdateItemJSONRequestBody{Name: &name, ParentId: &u1}
	res, err := c.UpdateItemWithResponse(context.TODO(), fix.Id, body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode())
	assert.Equal(t, name, res.JSON200.Name)
	assert.Equal(t, uint32(1), *res.JSON200.ParentId)
	assert.True(t, res.JSON200.UpdatedAt.After(justnow))
}
