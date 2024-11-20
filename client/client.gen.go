// Package client provides primitives to interact with the openapi HTTP API.
// @formatter:off
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/oapi-codegen/nullable"
	"github.com/oapi-codegen/runtime"
)

// Item defines model for Item.
type Item struct {
	Children  *[]Item    `json:"children,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Id        uint32     `json:"id"`

	// Name Item name
	Name   string `json:"name"`
	Parent *Item  `json:"parent,omitempty"`

	// ParentId Parent record ID
	ParentId  *uint32    `json:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// ItemCreate defines model for ItemCreate.
type ItemCreate struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Id        uint32     `json:"id"`

	// Name Item name
	Name string `json:"name"`

	// ParentId Parent record ID
	ParentId  *uint32    `json:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// ItemList defines model for ItemList.
type ItemList struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Id        uint32     `json:"id"`

	// Name Item name
	Name string `json:"name"`

	// ParentId Parent record ID
	ParentId  *uint32    `json:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// ItemRead defines model for ItemRead.
type ItemRead struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// DeletedAt Date and time when the record was deleted
	DeletedAt nullable.Nullable[time.Time] `json:"deleted_at,omitempty"`
	Id        uint32                       `json:"id"`

	// Name Item name
	Name string `json:"name"`

	// ParentId Parent record ID
	ParentId  *uint32    `json:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// ItemUpdate defines model for ItemUpdate.
type ItemUpdate struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Id        uint32     `json:"id"`

	// Name Item name
	Name string `json:"name"`

	// ParentId Parent record ID
	ParentId  *uint32    `json:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// ItemParentRead defines model for Item_ParentRead.
type ItemParentRead struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Id        uint32     `json:"id"`

	// Name Item name
	Name string `json:"name"`

	// ParentId Parent record ID
	ParentId  *uint32    `json:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// N400 defines model for 400.
type N400 struct {
	Code   int          `json:"code"`
	Errors *interface{} `json:"errors,omitempty"`
	Status string       `json:"status"`
}

// N404 defines model for 404.
type N404 struct {
	Code   int          `json:"code"`
	Errors *interface{} `json:"errors,omitempty"`
	Status string       `json:"status"`
}

// N409 defines model for 409.
type N409 struct {
	Code   int          `json:"code"`
	Errors *interface{} `json:"errors,omitempty"`
	Status string       `json:"status"`
}

// N500 defines model for 500.
type N500 struct {
	Code   int          `json:"code"`
	Errors *interface{} `json:"errors,omitempty"`
	Status string       `json:"status"`
}

// ListItemParams defines parameters for ListItem.
type ListItemParams struct {
	// Page what page to render
	Page *int `form:"page,omitempty" json:"page,omitempty"`

	// PerPage item count to render per page
	PerPage *int `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Name Name of the item
	Name *string `form:"name,omitempty" json:"name,omitempty"`

	// Trashed Whether to include trashed items
	Trashed *bool `form:"trashed,omitempty" json:"trashed,omitempty"`
}

// CreateItemJSONBody defines parameters for CreateItem.
type CreateItemJSONBody struct {
	// Name Item name
	Name string `json:"name"`

	// ParentId Parent record ID
	ParentId *uint32 `json:"parent_id,omitempty"`
}

// DeleteItemParams defines parameters for DeleteItem.
type DeleteItemParams struct {
	// Trashed Whether to include trashed items
	Trashed *bool `form:"trashed,omitempty" json:"trashed,omitempty"`
}

// ReadItemParams defines parameters for ReadItem.
type ReadItemParams struct {
	// Trashed Whether to include trashed items
	Trashed *bool `form:"trashed,omitempty" json:"trashed,omitempty"`
}

// UpdateItemJSONBody defines parameters for UpdateItem.
type UpdateItemJSONBody struct {
	// Name Item name
	Name *string `json:"name,omitempty"`

	// ParentId Parent record ID
	ParentId *uint32 `json:"parent_id,omitempty"`
}

// ListItemChildrenParams defines parameters for ListItemChildren.
type ListItemChildrenParams struct {
	// Page what page to render
	Page *int `form:"page,omitempty" json:"page,omitempty"`

	// PerPage item count to render per page
	PerPage *int `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Name Name of the item
	Name *string `form:"name,omitempty" json:"name,omitempty"`

	// Recurse Whether to return all descendants (recurse to last leaf)
	Recurse *bool `form:"recurse,omitempty" json:"recurse,omitempty"`
}

// CreateItemJSONRequestBody defines body for CreateItem for application/json ContentType.
type CreateItemJSONRequestBody CreateItemJSONBody

// UpdateItemJSONRequestBody defines body for UpdateItem for application/json ContentType.
type UpdateItemJSONRequestBody UpdateItemJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// ListItem request
	ListItem(ctx context.Context, params *ListItemParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateItemWithBody request with any body
	CreateItemWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateItem(ctx context.Context, body CreateItemJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteItem request
	DeleteItem(ctx context.Context, id uint32, params *DeleteItemParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ReadItem request
	ReadItem(ctx context.Context, id uint32, params *ReadItemParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateItemWithBody request with any body
	UpdateItemWithBody(ctx context.Context, id uint32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateItem(ctx context.Context, id uint32, body UpdateItemJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListItemChildren request
	ListItemChildren(ctx context.Context, id uint32, params *ListItemChildrenParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ReadItemParent request
	ReadItemParent(ctx context.Context, id uint32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreItem request
	RestoreItem(ctx context.Context, id uint32, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ListItem(ctx context.Context, params *ListItemParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListItemRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateItemWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateItemRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateItem(ctx context.Context, body CreateItemJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateItemRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteItem(ctx context.Context, id uint32, params *DeleteItemParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteItemRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ReadItem(ctx context.Context, id uint32, params *ReadItemParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewReadItemRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateItemWithBody(ctx context.Context, id uint32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateItemRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateItem(ctx context.Context, id uint32, body UpdateItemJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateItemRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListItemChildren(ctx context.Context, id uint32, params *ListItemChildrenParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListItemChildrenRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ReadItemParent(ctx context.Context, id uint32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewReadItemParentRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RestoreItem(ctx context.Context, id uint32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRestoreItemRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewListItemRequest generates requests for ListItem
func NewListItemRequest(server string, params *ListItemParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/simple-tree")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Page != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.PerPage != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "per_page", runtime.ParamLocationQuery, *params.PerPage); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Name != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "name", runtime.ParamLocationQuery, *params.Name); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Trashed != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "trashed", runtime.ParamLocationQuery, *params.Trashed); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateItemRequest calls the generic CreateItem builder with application/json body
func NewCreateItemRequest(server string, body CreateItemJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateItemRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateItemRequestWithBody generates requests for CreateItem with any type of body
func NewCreateItemRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/simple-tree")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteItemRequest generates requests for DeleteItem
func NewDeleteItemRequest(server string, id uint32, params *DeleteItemParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/simple-tree/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Trashed != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "trashed", runtime.ParamLocationQuery, *params.Trashed); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewReadItemRequest generates requests for ReadItem
func NewReadItemRequest(server string, id uint32, params *ReadItemParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/simple-tree/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Trashed != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "trashed", runtime.ParamLocationQuery, *params.Trashed); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateItemRequest calls the generic UpdateItem builder with application/json body
func NewUpdateItemRequest(server string, id uint32, body UpdateItemJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateItemRequestWithBody(server, id, "application/json", bodyReader)
}

// NewUpdateItemRequestWithBody generates requests for UpdateItem with any type of body
func NewUpdateItemRequestWithBody(server string, id uint32, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/simple-tree/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewListItemChildrenRequest generates requests for ListItemChildren
func NewListItemChildrenRequest(server string, id uint32, params *ListItemChildrenParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/simple-tree/%s/children", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Page != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.PerPage != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "per_page", runtime.ParamLocationQuery, *params.PerPage); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Name != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "name", runtime.ParamLocationQuery, *params.Name); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Recurse != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "recurse", runtime.ParamLocationQuery, *params.Recurse); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewReadItemParentRequest generates requests for ReadItemParent
func NewReadItemParentRequest(server string, id uint32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/simple-tree/%s/parent", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewRestoreItemRequest generates requests for RestoreItem
func NewRestoreItemRequest(server string, id uint32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/simple-tree/%s/restore", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// ListItemWithResponse request
	ListItemWithResponse(ctx context.Context, params *ListItemParams, reqEditors ...RequestEditorFn) (*ListItemResponse, error)

	// CreateItemWithBodyWithResponse request with any body
	CreateItemWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateItemResponse, error)

	CreateItemWithResponse(ctx context.Context, body CreateItemJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateItemResponse, error)

	// DeleteItemWithResponse request
	DeleteItemWithResponse(ctx context.Context, id uint32, params *DeleteItemParams, reqEditors ...RequestEditorFn) (*DeleteItemResponse, error)

	// ReadItemWithResponse request
	ReadItemWithResponse(ctx context.Context, id uint32, params *ReadItemParams, reqEditors ...RequestEditorFn) (*ReadItemResponse, error)

	// UpdateItemWithBodyWithResponse request with any body
	UpdateItemWithBodyWithResponse(ctx context.Context, id uint32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateItemResponse, error)

	UpdateItemWithResponse(ctx context.Context, id uint32, body UpdateItemJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateItemResponse, error)

	// ListItemChildrenWithResponse request
	ListItemChildrenWithResponse(ctx context.Context, id uint32, params *ListItemChildrenParams, reqEditors ...RequestEditorFn) (*ListItemChildrenResponse, error)

	// ReadItemParentWithResponse request
	ReadItemParentWithResponse(ctx context.Context, id uint32, reqEditors ...RequestEditorFn) (*ReadItemParentResponse, error)

	// RestoreItemWithResponse request
	RestoreItemWithResponse(ctx context.Context, id uint32, reqEditors ...RequestEditorFn) (*RestoreItemResponse, error)
}

type ListItemResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// CurrentPage Page number (1-based)
		CurrentPage int `json:"current_page"`

		// Data List of items
		Data []ItemList `json:"data"`

		// FirstPageUrl URL to the first page
		FirstPageUrl string `json:"first_page_url"`

		// From Index (1-based) of the first item in the current page
		From int `json:"from"`

		// LastPage Last page number
		LastPage int `json:"last_page"`

		// LastPageUrl URL to the last page
		LastPageUrl string `json:"last_page_url"`

		// NextPageUrl URL to the next page
		NextPageUrl string `json:"next_page_url"`

		// Path Base path of the request
		Path string `json:"path"`

		// PerPage Number of items per page
		PerPage int `json:"per_page"`

		// PrevPageUrl URL to the previous page
		PrevPageUrl string `json:"prev_page_url"`

		// To Index (1-based) of the last item in the current page
		To int `json:"to"`

		// Total Total number of items
		Total int `json:"total"`
	}
	JSON400 *N400
	JSON404 *N404
	JSON409 *N409
	JSON500 *N500
}

// Status returns HTTPResponse.Status
func (r ListItemResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListItemResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateItemResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ItemCreate
	JSON400      *N400
	JSON409      *N409
	JSON500      *N500
}

// Status returns HTTPResponse.Status
func (r CreateItemResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateItemResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteItemResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *N400
	JSON404      *N404
	JSON409      *N409
	JSON500      *N500
}

// Status returns HTTPResponse.Status
func (r DeleteItemResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteItemResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ReadItemResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ItemRead
	JSON400      *N400
	JSON404      *N404
	JSON409      *N409
	JSON500      *N500
}

// Status returns HTTPResponse.Status
func (r ReadItemResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ReadItemResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateItemResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ItemUpdate
	JSON400      *N400
	JSON404      *N404
	JSON409      *N409
	JSON500      *N500
}

// Status returns HTTPResponse.Status
func (r UpdateItemResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateItemResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListItemChildrenResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// CurrentPage Page number (1-based)
		CurrentPage int `json:"current_page"`

		// Data List of items
		Data []ItemList `json:"data"`

		// FirstPageUrl URL to the first page
		FirstPageUrl string `json:"first_page_url"`

		// From Index (1-based) of the first item in the current page
		From int `json:"from"`

		// LastPage Last page number
		LastPage int `json:"last_page"`

		// LastPageUrl URL to the last page
		LastPageUrl string `json:"last_page_url"`

		// NextPageUrl URL to the next page
		NextPageUrl string `json:"next_page_url"`

		// Path Base path of the request
		Path string `json:"path"`

		// PerPage Number of items per page
		PerPage int `json:"per_page"`

		// PrevPageUrl URL to the previous page
		PrevPageUrl string `json:"prev_page_url"`

		// To Index (1-based) of the last item in the current page
		To int `json:"to"`

		// Total Total number of items
		Total int `json:"total"`
	}
	JSON400 *N400
	JSON404 *N404
	JSON409 *N409
	JSON500 *N500
}

// Status returns HTTPResponse.Status
func (r ListItemChildrenResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListItemChildrenResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ReadItemParentResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ItemParentRead
	JSON400      *N400
	JSON404      *N404
	JSON409      *N409
	JSON500      *N500
}

// Status returns HTTPResponse.Status
func (r ReadItemParentResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ReadItemParentResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RestoreItemResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *N400
	JSON404      *N404
	JSON409      *N409
	JSON500      *N500
}

// Status returns HTTPResponse.Status
func (r RestoreItemResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RestoreItemResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListItemWithResponse request returning *ListItemResponse
func (c *ClientWithResponses) ListItemWithResponse(ctx context.Context, params *ListItemParams, reqEditors ...RequestEditorFn) (*ListItemResponse, error) {
	rsp, err := c.ListItem(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListItemResponse(rsp)
}

// CreateItemWithBodyWithResponse request with arbitrary body returning *CreateItemResponse
func (c *ClientWithResponses) CreateItemWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateItemResponse, error) {
	rsp, err := c.CreateItemWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateItemResponse(rsp)
}

func (c *ClientWithResponses) CreateItemWithResponse(ctx context.Context, body CreateItemJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateItemResponse, error) {
	rsp, err := c.CreateItem(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateItemResponse(rsp)
}

// DeleteItemWithResponse request returning *DeleteItemResponse
func (c *ClientWithResponses) DeleteItemWithResponse(ctx context.Context, id uint32, params *DeleteItemParams, reqEditors ...RequestEditorFn) (*DeleteItemResponse, error) {
	rsp, err := c.DeleteItem(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteItemResponse(rsp)
}

// ReadItemWithResponse request returning *ReadItemResponse
func (c *ClientWithResponses) ReadItemWithResponse(ctx context.Context, id uint32, params *ReadItemParams, reqEditors ...RequestEditorFn) (*ReadItemResponse, error) {
	rsp, err := c.ReadItem(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseReadItemResponse(rsp)
}

// UpdateItemWithBodyWithResponse request with arbitrary body returning *UpdateItemResponse
func (c *ClientWithResponses) UpdateItemWithBodyWithResponse(ctx context.Context, id uint32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateItemResponse, error) {
	rsp, err := c.UpdateItemWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateItemResponse(rsp)
}

func (c *ClientWithResponses) UpdateItemWithResponse(ctx context.Context, id uint32, body UpdateItemJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateItemResponse, error) {
	rsp, err := c.UpdateItem(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateItemResponse(rsp)
}

// ListItemChildrenWithResponse request returning *ListItemChildrenResponse
func (c *ClientWithResponses) ListItemChildrenWithResponse(ctx context.Context, id uint32, params *ListItemChildrenParams, reqEditors ...RequestEditorFn) (*ListItemChildrenResponse, error) {
	rsp, err := c.ListItemChildren(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListItemChildrenResponse(rsp)
}

// ReadItemParentWithResponse request returning *ReadItemParentResponse
func (c *ClientWithResponses) ReadItemParentWithResponse(ctx context.Context, id uint32, reqEditors ...RequestEditorFn) (*ReadItemParentResponse, error) {
	rsp, err := c.ReadItemParent(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseReadItemParentResponse(rsp)
}

// RestoreItemWithResponse request returning *RestoreItemResponse
func (c *ClientWithResponses) RestoreItemWithResponse(ctx context.Context, id uint32, reqEditors ...RequestEditorFn) (*RestoreItemResponse, error) {
	rsp, err := c.RestoreItem(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRestoreItemResponse(rsp)
}

// ParseListItemResponse parses an HTTP response from a ListItemWithResponse call
func ParseListItemResponse(rsp *http.Response) (*ListItemResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListItemResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// CurrentPage Page number (1-based)
			CurrentPage int `json:"current_page"`

			// Data List of items
			Data []ItemList `json:"data"`

			// FirstPageUrl URL to the first page
			FirstPageUrl string `json:"first_page_url"`

			// From Index (1-based) of the first item in the current page
			From int `json:"from"`

			// LastPage Last page number
			LastPage int `json:"last_page"`

			// LastPageUrl URL to the last page
			LastPageUrl string `json:"last_page_url"`

			// NextPageUrl URL to the next page
			NextPageUrl string `json:"next_page_url"`

			// Path Base path of the request
			Path string `json:"path"`

			// PerPage Number of items per page
			PerPage int `json:"per_page"`

			// PrevPageUrl URL to the previous page
			PrevPageUrl string `json:"prev_page_url"`

			// To Index (1-based) of the last item in the current page
			To int `json:"to"`

			// Total Total number of items
			Total int `json:"total"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest N400
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest N404
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest N409
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest N500
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseCreateItemResponse parses an HTTP response from a CreateItemWithResponse call
func ParseCreateItemResponse(rsp *http.Response) (*CreateItemResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateItemResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest ItemCreate
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest N400
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest N409
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest N500
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseDeleteItemResponse parses an HTTP response from a DeleteItemWithResponse call
func ParseDeleteItemResponse(rsp *http.Response) (*DeleteItemResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteItemResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest N400
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest N404
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest N409
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest N500
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseReadItemResponse parses an HTTP response from a ReadItemWithResponse call
func ParseReadItemResponse(rsp *http.Response) (*ReadItemResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ReadItemResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ItemRead
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest N400
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest N404
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest N409
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest N500
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseUpdateItemResponse parses an HTTP response from a UpdateItemWithResponse call
func ParseUpdateItemResponse(rsp *http.Response) (*UpdateItemResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateItemResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ItemUpdate
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest N400
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest N404
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest N409
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest N500
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseListItemChildrenResponse parses an HTTP response from a ListItemChildrenWithResponse call
func ParseListItemChildrenResponse(rsp *http.Response) (*ListItemChildrenResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListItemChildrenResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// CurrentPage Page number (1-based)
			CurrentPage int `json:"current_page"`

			// Data List of items
			Data []ItemList `json:"data"`

			// FirstPageUrl URL to the first page
			FirstPageUrl string `json:"first_page_url"`

			// From Index (1-based) of the first item in the current page
			From int `json:"from"`

			// LastPage Last page number
			LastPage int `json:"last_page"`

			// LastPageUrl URL to the last page
			LastPageUrl string `json:"last_page_url"`

			// NextPageUrl URL to the next page
			NextPageUrl string `json:"next_page_url"`

			// Path Base path of the request
			Path string `json:"path"`

			// PerPage Number of items per page
			PerPage int `json:"per_page"`

			// PrevPageUrl URL to the previous page
			PrevPageUrl string `json:"prev_page_url"`

			// To Index (1-based) of the last item in the current page
			To int `json:"to"`

			// Total Total number of items
			Total int `json:"total"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest N400
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest N404
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest N409
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest N500
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseReadItemParentResponse parses an HTTP response from a ReadItemParentWithResponse call
func ParseReadItemParentResponse(rsp *http.Response) (*ReadItemParentResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ReadItemParentResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ItemParentRead
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest N400
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest N404
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest N409
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest N500
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseRestoreItemResponse parses an HTTP response from a RestoreItemWithResponse call
func ParseRestoreItemResponse(rsp *http.Response) (*RestoreItemResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RestoreItemResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest N400
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest N404
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest N409
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest N500
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
