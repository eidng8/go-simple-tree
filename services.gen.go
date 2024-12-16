// Package api provides primitives to interact with the openapi HTTP API.
// @formatter:off
package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List Items
	// (GET /simple-tree)
	ListItem(c *gin.Context, params ListItemParams)
	// Create a new Item
	// (POST /simple-tree)
	CreateItem(c *gin.Context)
	// Deletes a Item by ID
	// (DELETE /simple-tree/{id})
	DeleteItem(c *gin.Context, id uint32, params DeleteItemParams)
	// Find a Item by ID
	// (GET /simple-tree/{id})
	ReadItem(c *gin.Context, id uint32, params ReadItemParams)
	// Updates a Item
	// (PATCH /simple-tree/{id})
	UpdateItem(c *gin.Context, id uint32)
	// List of subordinate items
	// (GET /simple-tree/{id}/children)
	ListItemChildren(c *gin.Context, id uint32, params ListItemChildrenParams)
	// Find the attached Item
	// (GET /simple-tree/{id}/parent)
	ReadItemParent(c *gin.Context, id uint32)
	// Restore a trashed record
	// (POST /simple-tree/{id}/restore)
	RestoreItem(c *gin.Context, id uint32)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// ListItem operation middleware
func (siw *ServerInterfaceWrapper) ListItem(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListItemParams

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", c.Request.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "per_page" -------------

	err = runtime.BindQueryParameter("form", true, false, "per_page", c.Request.URL.Query(), &params.PerPage)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter per_page: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, false, "name", c.Request.URL.Query(), &params.Name)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter name: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "trashed" -------------

	err = runtime.BindQueryParameter("form", true, false, "trashed", c.Request.URL.Query(), &params.Trashed)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter trashed: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ListItem(c, params)
}

// CreateItem operation middleware
func (siw *ServerInterfaceWrapper) CreateItem(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateItem(c)
}

// DeleteItem operation middleware
func (siw *ServerInterfaceWrapper) DeleteItem(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uint32

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteItemParams

	// ------------- Optional query parameter "trashed" -------------

	err = runtime.BindQueryParameter("form", true, false, "trashed", c.Request.URL.Query(), &params.Trashed)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter trashed: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteItem(c, id, params)
}

// ReadItem operation middleware
func (siw *ServerInterfaceWrapper) ReadItem(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uint32

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params ReadItemParams

	// ------------- Optional query parameter "trashed" -------------

	err = runtime.BindQueryParameter("form", true, false, "trashed", c.Request.URL.Query(), &params.Trashed)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter trashed: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ReadItem(c, id, params)
}

// UpdateItem operation middleware
func (siw *ServerInterfaceWrapper) UpdateItem(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uint32

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateItem(c, id)
}

// ListItemChildren operation middleware
func (siw *ServerInterfaceWrapper) ListItemChildren(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uint32

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params ListItemChildrenParams

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", c.Request.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "per_page" -------------

	err = runtime.BindQueryParameter("form", true, false, "per_page", c.Request.URL.Query(), &params.PerPage)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter per_page: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, false, "name", c.Request.URL.Query(), &params.Name)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter name: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "recurse" -------------

	err = runtime.BindQueryParameter("form", true, false, "recurse", c.Request.URL.Query(), &params.Recurse)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter recurse: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ListItemChildren(c, id, params)
}

// ReadItemParent operation middleware
func (siw *ServerInterfaceWrapper) ReadItemParent(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uint32

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ReadItemParent(c, id)
}

// RestoreItem operation middleware
func (siw *ServerInterfaceWrapper) RestoreItem(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uint32

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.RestoreItem(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/simple-tree", wrapper.ListItem)
	router.POST(options.BaseURL+"/simple-tree", wrapper.CreateItem)
	router.DELETE(options.BaseURL+"/simple-tree/:id", wrapper.DeleteItem)
	router.GET(options.BaseURL+"/simple-tree/:id", wrapper.ReadItem)
	router.PATCH(options.BaseURL+"/simple-tree/:id", wrapper.UpdateItem)
	router.GET(options.BaseURL+"/simple-tree/:id/children", wrapper.ListItemChildren)
	router.GET(options.BaseURL+"/simple-tree/:id/parent", wrapper.ReadItemParent)
	router.POST(options.BaseURL+"/simple-tree/:id/restore", wrapper.RestoreItem)
}

type N400JSONResponse struct {
	Code   int          `json:"code" yaml:"code" xml:"code" bson:"code"`
	Errors *interface{} `json:"errors,omitempty" yaml:"errors,omitempty" xml:"errors,omitempty" bson:"errors,omitempty"`
	Status string       `json:"status" yaml:"status" xml:"status" bson:"status"`
}

type N404JSONResponse struct {
	Code   int          `json:"code" yaml:"code" xml:"code" bson:"code"`
	Errors *interface{} `json:"errors,omitempty" yaml:"errors,omitempty" xml:"errors,omitempty" bson:"errors,omitempty"`
	Status string       `json:"status" yaml:"status" xml:"status" bson:"status"`
}

type N409JSONResponse struct {
	Code   int          `json:"code" yaml:"code" xml:"code" bson:"code"`
	Errors *interface{} `json:"errors,omitempty" yaml:"errors,omitempty" xml:"errors,omitempty" bson:"errors,omitempty"`
	Status string       `json:"status" yaml:"status" xml:"status" bson:"status"`
}

type N500JSONResponse struct {
	Code   int          `json:"code" yaml:"code" xml:"code" bson:"code"`
	Errors *interface{} `json:"errors,omitempty" yaml:"errors,omitempty" xml:"errors,omitempty" bson:"errors,omitempty"`
	Status string       `json:"status" yaml:"status" xml:"status" bson:"status"`
}

type ListItemRequestObject struct {
	Params ListItemParams
}

type ListItemResponseObject interface {
	VisitListItemResponse(w http.ResponseWriter) error
}

type ListItem200JSONResponse struct {
	// CurrentPage Page number (1-based)
	CurrentPage int `json:"current_page" yaml:"current_page" xml:"current_page" bson:"current_page"`

	// Data List of items
	Data []ItemList `json:"data" yaml:"data" xml:"data" bson:"data"`

	// FirstPageUrl URL to the first page
	FirstPageUrl string `json:"first_page_url" yaml:"first_page_url" xml:"first_page_url" bson:"first_page_url"`

	// From Index (1-based) of the first item in the current page
	From int `json:"from" yaml:"from" xml:"from" bson:"from"`

	// LastPage Last page number
	LastPage int `json:"last_page" yaml:"last_page" xml:"last_page" bson:"last_page"`

	// LastPageUrl URL to the last page
	LastPageUrl string `json:"last_page_url" yaml:"last_page_url" xml:"last_page_url" bson:"last_page_url"`

	// NextPageUrl URL to the next page
	NextPageUrl string `json:"next_page_url" yaml:"next_page_url" xml:"next_page_url" bson:"next_page_url"`

	// Path Base path of the request
	Path string `json:"path" yaml:"path" xml:"path" bson:"path"`

	// PerPage Number of items per page
	PerPage int `json:"per_page" yaml:"per_page" xml:"per_page" bson:"per_page"`

	// PrevPageUrl URL to the previous page
	PrevPageUrl string `json:"prev_page_url" yaml:"prev_page_url" xml:"prev_page_url" bson:"prev_page_url"`

	// To Index (1-based) of the last item in the current page
	To int `json:"to" yaml:"to" xml:"to" bson:"to"`

	// Total Total number of items
	Total int `json:"total" yaml:"total" xml:"total" bson:"total"`
}

func (response ListItem200JSONResponse) VisitListItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ListItem400JSONResponse struct{ N400JSONResponse }

func (response ListItem400JSONResponse) VisitListItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type ListItem404JSONResponse struct{ N404JSONResponse }

func (response ListItem404JSONResponse) VisitListItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type ListItem409JSONResponse struct{ N409JSONResponse }

func (response ListItem409JSONResponse) VisitListItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type ListItem500JSONResponse struct{ N500JSONResponse }

func (response ListItem500JSONResponse) VisitListItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type CreateItemRequestObject struct {
	Body *CreateItemJSONRequestBody
}

type CreateItemResponseObject interface {
	VisitCreateItemResponse(w http.ResponseWriter) error
}

type CreateItem200JSONResponse ItemCreate

func (response CreateItem200JSONResponse) VisitCreateItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type CreateItem400JSONResponse struct{ N400JSONResponse }

func (response CreateItem400JSONResponse) VisitCreateItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type CreateItem409JSONResponse struct{ N409JSONResponse }

func (response CreateItem409JSONResponse) VisitCreateItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type CreateItem500JSONResponse struct{ N500JSONResponse }

func (response CreateItem500JSONResponse) VisitCreateItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type DeleteItemRequestObject struct {
	Id     uint32 `json:"id" yaml:"id" xml:"id" bson:"id"`
	Params DeleteItemParams
}

type DeleteItemResponseObject interface {
	VisitDeleteItemResponse(w http.ResponseWriter) error
}

type DeleteItem204Response struct {
}

func (response DeleteItem204Response) VisitDeleteItemResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteItem400JSONResponse struct{ N400JSONResponse }

func (response DeleteItem400JSONResponse) VisitDeleteItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type DeleteItem404JSONResponse struct{ N404JSONResponse }

func (response DeleteItem404JSONResponse) VisitDeleteItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type DeleteItem409JSONResponse struct{ N409JSONResponse }

func (response DeleteItem409JSONResponse) VisitDeleteItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type DeleteItem500JSONResponse struct{ N500JSONResponse }

func (response DeleteItem500JSONResponse) VisitDeleteItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type ReadItemRequestObject struct {
	Id     uint32 `json:"id" yaml:"id" xml:"id" bson:"id"`
	Params ReadItemParams
}

type ReadItemResponseObject interface {
	VisitReadItemResponse(w http.ResponseWriter) error
}

type ReadItem200JSONResponse ItemRead

func (response ReadItem200JSONResponse) VisitReadItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ReadItem400JSONResponse struct{ N400JSONResponse }

func (response ReadItem400JSONResponse) VisitReadItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type ReadItem404JSONResponse struct{ N404JSONResponse }

func (response ReadItem404JSONResponse) VisitReadItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type ReadItem409JSONResponse struct{ N409JSONResponse }

func (response ReadItem409JSONResponse) VisitReadItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type ReadItem500JSONResponse struct{ N500JSONResponse }

func (response ReadItem500JSONResponse) VisitReadItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type UpdateItemRequestObject struct {
	Id   uint32 `json:"id" yaml:"id" xml:"id" bson:"id"`
	Body *UpdateItemJSONRequestBody
}

type UpdateItemResponseObject interface {
	VisitUpdateItemResponse(w http.ResponseWriter) error
}

type UpdateItem200JSONResponse ItemUpdate

func (response UpdateItem200JSONResponse) VisitUpdateItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type UpdateItem400JSONResponse struct{ N400JSONResponse }

func (response UpdateItem400JSONResponse) VisitUpdateItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type UpdateItem404JSONResponse struct{ N404JSONResponse }

func (response UpdateItem404JSONResponse) VisitUpdateItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type UpdateItem409JSONResponse struct{ N409JSONResponse }

func (response UpdateItem409JSONResponse) VisitUpdateItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type UpdateItem500JSONResponse struct{ N500JSONResponse }

func (response UpdateItem500JSONResponse) VisitUpdateItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type ListItemChildrenRequestObject struct {
	Id     uint32 `json:"id" yaml:"id" xml:"id" bson:"id"`
	Params ListItemChildrenParams
}

type ListItemChildrenResponseObject interface {
	VisitListItemChildrenResponse(w http.ResponseWriter) error
}

type ListItemChildren200JSONResponse struct {
	// CurrentPage Page number (1-based)
	CurrentPage int `json:"current_page" yaml:"current_page" xml:"current_page" bson:"current_page"`

	// Data List of items
	Data []ItemList `json:"data" yaml:"data" xml:"data" bson:"data"`

	// FirstPageUrl URL to the first page
	FirstPageUrl string `json:"first_page_url" yaml:"first_page_url" xml:"first_page_url" bson:"first_page_url"`

	// From Index (1-based) of the first item in the current page
	From int `json:"from" yaml:"from" xml:"from" bson:"from"`

	// LastPage Last page number
	LastPage int `json:"last_page" yaml:"last_page" xml:"last_page" bson:"last_page"`

	// LastPageUrl URL to the last page
	LastPageUrl string `json:"last_page_url" yaml:"last_page_url" xml:"last_page_url" bson:"last_page_url"`

	// NextPageUrl URL to the next page
	NextPageUrl string `json:"next_page_url" yaml:"next_page_url" xml:"next_page_url" bson:"next_page_url"`

	// Path Base path of the request
	Path string `json:"path" yaml:"path" xml:"path" bson:"path"`

	// PerPage Number of items per page
	PerPage int `json:"per_page" yaml:"per_page" xml:"per_page" bson:"per_page"`

	// PrevPageUrl URL to the previous page
	PrevPageUrl string `json:"prev_page_url" yaml:"prev_page_url" xml:"prev_page_url" bson:"prev_page_url"`

	// To Index (1-based) of the last item in the current page
	To int `json:"to" yaml:"to" xml:"to" bson:"to"`

	// Total Total number of items
	Total int `json:"total" yaml:"total" xml:"total" bson:"total"`
}

func (response ListItemChildren200JSONResponse) VisitListItemChildrenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ListItemChildren400JSONResponse struct{ N400JSONResponse }

func (response ListItemChildren400JSONResponse) VisitListItemChildrenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type ListItemChildren404JSONResponse struct{ N404JSONResponse }

func (response ListItemChildren404JSONResponse) VisitListItemChildrenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type ListItemChildren409JSONResponse struct{ N409JSONResponse }

func (response ListItemChildren409JSONResponse) VisitListItemChildrenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type ListItemChildren500JSONResponse struct{ N500JSONResponse }

func (response ListItemChildren500JSONResponse) VisitListItemChildrenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type ReadItemParentRequestObject struct {
	Id uint32 `json:"id"`
}

type ReadItemParentResponseObject interface {
	VisitReadItemParentResponse(w http.ResponseWriter) error
}

type ReadItemParent200JSONResponse ItemParentRead

func (response ReadItemParent200JSONResponse) VisitReadItemParentResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ReadItemParent400JSONResponse struct{ N400JSONResponse }

func (response ReadItemParent400JSONResponse) VisitReadItemParentResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type ReadItemParent404JSONResponse struct{ N404JSONResponse }

func (response ReadItemParent404JSONResponse) VisitReadItemParentResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type ReadItemParent409JSONResponse struct{ N409JSONResponse }

func (response ReadItemParent409JSONResponse) VisitReadItemParentResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type ReadItemParent500JSONResponse struct{ N500JSONResponse }

func (response ReadItemParent500JSONResponse) VisitReadItemParentResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type RestoreItemRequestObject struct {
	Id uint32 `json:"id"`
}

type RestoreItemResponseObject interface {
	VisitRestoreItemResponse(w http.ResponseWriter) error
}

type RestoreItem204Response struct {
}

func (response RestoreItem204Response) VisitRestoreItemResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type RestoreItem400JSONResponse struct{ N400JSONResponse }

func (response RestoreItem400JSONResponse) VisitRestoreItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type RestoreItem404JSONResponse struct{ N404JSONResponse }

func (response RestoreItem404JSONResponse) VisitRestoreItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type RestoreItem409JSONResponse struct{ N409JSONResponse }

func (response RestoreItem409JSONResponse) VisitRestoreItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type RestoreItem500JSONResponse struct{ N500JSONResponse }

func (response RestoreItem500JSONResponse) VisitRestoreItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// List Items
	// (GET /simple-tree)
	ListItem(ctx context.Context, request ListItemRequestObject) (ListItemResponseObject, error)
	// Create a new Item
	// (POST /simple-tree)
	CreateItem(ctx context.Context, request CreateItemRequestObject) (CreateItemResponseObject, error)
	// Deletes a Item by ID
	// (DELETE /simple-tree/{id})
	DeleteItem(ctx context.Context, request DeleteItemRequestObject) (DeleteItemResponseObject, error)
	// Find a Item by ID
	// (GET /simple-tree/{id})
	ReadItem(ctx context.Context, request ReadItemRequestObject) (ReadItemResponseObject, error)
	// Updates a Item
	// (PATCH /simple-tree/{id})
	UpdateItem(ctx context.Context, request UpdateItemRequestObject) (UpdateItemResponseObject, error)
	// List of subordinate items
	// (GET /simple-tree/{id}/children)
	ListItemChildren(ctx context.Context, request ListItemChildrenRequestObject) (ListItemChildrenResponseObject, error)
	// Find the attached Item
	// (GET /simple-tree/{id}/parent)
	ReadItemParent(ctx context.Context, request ReadItemParentRequestObject) (ReadItemParentResponseObject, error)
	// Restore a trashed record
	// (POST /simple-tree/{id}/restore)
	RestoreItem(ctx context.Context, request RestoreItemRequestObject) (RestoreItemResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// ListItem operation middleware
func (sh *strictHandler) ListItem(ctx *gin.Context, params ListItemParams) {
	var request ListItemRequestObject

	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ListItem(ctx, request.(ListItemRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListItem")
	}

	response, err := handler(ctx, request)

	if err != nil {
		handleErrorResponse(ctx, err)
	} else if validResponse, ok := response.(ListItemResponseObject); ok {
		if err := validResponse.VisitListItemResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// CreateItem operation middleware
func (sh *strictHandler) CreateItem(ctx *gin.Context) {
	var request CreateItemRequestObject

	var body CreateItemJSONRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreateItem(ctx, request.(CreateItemRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateItem")
	}

	response, err := handler(ctx, request)

	if err != nil {
		handleErrorResponse(ctx, err)
	} else if validResponse, ok := response.(CreateItemResponseObject); ok {
		if err := validResponse.VisitCreateItemResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteItem operation middleware
func (sh *strictHandler) DeleteItem(ctx *gin.Context, id uint32, params DeleteItemParams) {
	var request DeleteItemRequestObject

	request.Id = id
	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteItem(ctx, request.(DeleteItemRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteItem")
	}

	response, err := handler(ctx, request)

	if err != nil {
		handleErrorResponse(ctx, err)
	} else if validResponse, ok := response.(DeleteItemResponseObject); ok {
		if err := validResponse.VisitDeleteItemResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// ReadItem operation middleware
func (sh *strictHandler) ReadItem(ctx *gin.Context, id uint32, params ReadItemParams) {
	var request ReadItemRequestObject

	request.Id = id
	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ReadItem(ctx, request.(ReadItemRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ReadItem")
	}

	response, err := handler(ctx, request)

	if err != nil {
		handleErrorResponse(ctx, err)
	} else if validResponse, ok := response.(ReadItemResponseObject); ok {
		if err := validResponse.VisitReadItemResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// UpdateItem operation middleware
func (sh *strictHandler) UpdateItem(ctx *gin.Context, id uint32) {
	var request UpdateItemRequestObject

	request.Id = id

	var body UpdateItemJSONRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateItem(ctx, request.(UpdateItemRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateItem")
	}

	response, err := handler(ctx, request)

	if err != nil {
		handleErrorResponse(ctx, err)
	} else if validResponse, ok := response.(UpdateItemResponseObject); ok {
		if err := validResponse.VisitUpdateItemResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// ListItemChildren operation middleware
func (sh *strictHandler) ListItemChildren(ctx *gin.Context, id uint32, params ListItemChildrenParams) {
	var request ListItemChildrenRequestObject

	request.Id = id
	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ListItemChildren(ctx, request.(ListItemChildrenRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListItemChildren")
	}

	response, err := handler(ctx, request)

	if err != nil {
		handleErrorResponse(ctx, err)
	} else if validResponse, ok := response.(ListItemChildrenResponseObject); ok {
		if err := validResponse.VisitListItemChildrenResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// ReadItemParent operation middleware
func (sh *strictHandler) ReadItemParent(ctx *gin.Context, id uint32) {
	var request ReadItemParentRequestObject

	request.Id = id

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ReadItemParent(ctx, request.(ReadItemParentRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ReadItemParent")
	}

	response, err := handler(ctx, request)

	if err != nil {
		handleErrorResponse(ctx, err)
	} else if validResponse, ok := response.(ReadItemParentResponseObject); ok {
		if err := validResponse.VisitReadItemParentResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// RestoreItem operation middleware
func (sh *strictHandler) RestoreItem(ctx *gin.Context, id uint32) {
	var request RestoreItemRequestObject

	request.Id = id

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.RestoreItem(ctx, request.(RestoreItemRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "RestoreItem")
	}

	response, err := handler(ctx, request)

	if err != nil {
		handleErrorResponse(ctx, err)
	} else if validResponse, ok := response.(RestoreItemResponseObject); ok {
		if err := validResponse.VisitRestoreItemResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
