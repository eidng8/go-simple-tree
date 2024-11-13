// Package api provides primitives to interact with the openapi HTTP API.
// @formatter:off
package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
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
	// List SimpleTrees
	// (GET /simple-tree)
	ListSimpleTree(c *gin.Context, params ListSimpleTreeParams)
	// Create a new SimpleTree
	// (POST /simple-tree)
	CreateSimpleTree(c *gin.Context)
	// Deletes a SimpleTree by ID
	// (DELETE /simple-tree/{id})
	DeleteSimpleTree(c *gin.Context, id uint32, params DeleteSimpleTreeParams)
	// Find a SimpleTree by ID
	// (GET /simple-tree/{id})
	ReadSimpleTree(c *gin.Context, id uint32, params ReadSimpleTreeParams)
	// Updates a SimpleTree
	// (PATCH /simple-tree/{id})
	UpdateSimpleTree(c *gin.Context, id uint32)
	// List of subordinate items
	// (GET /simple-tree/{id}/children)
	ListSimpleTreeChildren(c *gin.Context, id uint32, params ListSimpleTreeChildrenParams)
	// Find the attached SimpleTree
	// (GET /simple-tree/{id}/parent)
	ReadSimpleTreeParent(c *gin.Context, id uint32)
	// Restore a trashed record
	// (POST /simple-tree/{id}/restore)
	RestoreSimpleTree(c *gin.Context, id uint32)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// ListSimpleTree operation middleware
func (siw *ServerInterfaceWrapper) ListSimpleTree(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListSimpleTreeParams

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

	siw.Handler.ListSimpleTree(c, params)
}

// CreateSimpleTree operation middleware
func (siw *ServerInterfaceWrapper) CreateSimpleTree(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateSimpleTree(c)
}

// DeleteSimpleTree operation middleware
func (siw *ServerInterfaceWrapper) DeleteSimpleTree(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uint32

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteSimpleTreeParams

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

	siw.Handler.DeleteSimpleTree(c, id, params)
}

// ReadSimpleTree operation middleware
func (siw *ServerInterfaceWrapper) ReadSimpleTree(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uint32

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params ReadSimpleTreeParams

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

	siw.Handler.ReadSimpleTree(c, id, params)
}

// UpdateSimpleTree operation middleware
func (siw *ServerInterfaceWrapper) UpdateSimpleTree(c *gin.Context) {

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

	siw.Handler.UpdateSimpleTree(c, id)
}

// ListSimpleTreeChildren operation middleware
func (siw *ServerInterfaceWrapper) ListSimpleTreeChildren(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uint32

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params ListSimpleTreeChildrenParams

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

	siw.Handler.ListSimpleTreeChildren(c, id, params)
}

// ReadSimpleTreeParent operation middleware
func (siw *ServerInterfaceWrapper) ReadSimpleTreeParent(c *gin.Context) {

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

	siw.Handler.ReadSimpleTreeParent(c, id)
}

// RestoreSimpleTree operation middleware
func (siw *ServerInterfaceWrapper) RestoreSimpleTree(c *gin.Context) {

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

	siw.Handler.RestoreSimpleTree(c, id)
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

	router.GET(options.BaseURL+"/simple-tree", wrapper.ListSimpleTree)
	router.POST(options.BaseURL+"/simple-tree", wrapper.CreateSimpleTree)
	router.DELETE(options.BaseURL+"/simple-tree/:id", wrapper.DeleteSimpleTree)
	router.GET(options.BaseURL+"/simple-tree/:id", wrapper.ReadSimpleTree)
	router.PATCH(options.BaseURL+"/simple-tree/:id", wrapper.UpdateSimpleTree)
	router.GET(options.BaseURL+"/simple-tree/:id/children", wrapper.ListSimpleTreeChildren)
	router.GET(options.BaseURL+"/simple-tree/:id/parent", wrapper.ReadSimpleTreeParent)
	router.POST(options.BaseURL+"/simple-tree/:id/restore", wrapper.RestoreSimpleTree)
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

type ListSimpleTreeRequestObject struct {
	Params ListSimpleTreeParams
}

type ListSimpleTreeResponseObject interface {
	VisitListSimpleTreeResponse(w http.ResponseWriter) error
}

type ListSimpleTree200JSONResponse struct {
	// CurrentPage Page number (1-based)
	CurrentPage int `json:"current_page" yaml:"current_page" xml:"current_page" bson:"current_page"`

	// Data List of items
	Data []SimpleTreeList `json:"data" yaml:"data" xml:"data" bson:"data"`

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

func (response ListSimpleTree200JSONResponse) VisitListSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ListSimpleTree400JSONResponse struct{ N400JSONResponse }

func (response ListSimpleTree400JSONResponse) VisitListSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type ListSimpleTree404JSONResponse struct{ N404JSONResponse }

func (response ListSimpleTree404JSONResponse) VisitListSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type ListSimpleTree409JSONResponse struct{ N409JSONResponse }

func (response ListSimpleTree409JSONResponse) VisitListSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type ListSimpleTree500JSONResponse struct{ N500JSONResponse }

func (response ListSimpleTree500JSONResponse) VisitListSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type CreateSimpleTreeRequestObject struct {
	Body *CreateSimpleTreeJSONRequestBody
}

type CreateSimpleTreeResponseObject interface {
	VisitCreateSimpleTreeResponse(w http.ResponseWriter) error
}

type CreateSimpleTree200JSONResponse SimpleTreeCreate

func (response CreateSimpleTree200JSONResponse) VisitCreateSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type CreateSimpleTree400JSONResponse struct{ N400JSONResponse }

func (response CreateSimpleTree400JSONResponse) VisitCreateSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type CreateSimpleTree409JSONResponse struct{ N409JSONResponse }

func (response CreateSimpleTree409JSONResponse) VisitCreateSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type CreateSimpleTree500JSONResponse struct{ N500JSONResponse }

func (response CreateSimpleTree500JSONResponse) VisitCreateSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type DeleteSimpleTreeRequestObject struct {
	Id     uint32 `json:"id" yaml:"id" xml:"id" bson:"id"`
	Params DeleteSimpleTreeParams
}

type DeleteSimpleTreeResponseObject interface {
	VisitDeleteSimpleTreeResponse(w http.ResponseWriter) error
}

type DeleteSimpleTree204Response struct {
}

func (response DeleteSimpleTree204Response) VisitDeleteSimpleTreeResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteSimpleTree400JSONResponse struct{ N400JSONResponse }

func (response DeleteSimpleTree400JSONResponse) VisitDeleteSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type DeleteSimpleTree404JSONResponse struct{ N404JSONResponse }

func (response DeleteSimpleTree404JSONResponse) VisitDeleteSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type DeleteSimpleTree409JSONResponse struct{ N409JSONResponse }

func (response DeleteSimpleTree409JSONResponse) VisitDeleteSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type DeleteSimpleTree500JSONResponse struct{ N500JSONResponse }

func (response DeleteSimpleTree500JSONResponse) VisitDeleteSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type ReadSimpleTreeRequestObject struct {
	Id     uint32 `json:"id" yaml:"id" xml:"id" bson:"id"`
	Params ReadSimpleTreeParams
}

type ReadSimpleTreeResponseObject interface {
	VisitReadSimpleTreeResponse(w http.ResponseWriter) error
}

type ReadSimpleTree200JSONResponse SimpleTreeRead

func (response ReadSimpleTree200JSONResponse) VisitReadSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ReadSimpleTree400JSONResponse struct{ N400JSONResponse }

func (response ReadSimpleTree400JSONResponse) VisitReadSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type ReadSimpleTree404JSONResponse struct{ N404JSONResponse }

func (response ReadSimpleTree404JSONResponse) VisitReadSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type ReadSimpleTree409JSONResponse struct{ N409JSONResponse }

func (response ReadSimpleTree409JSONResponse) VisitReadSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type ReadSimpleTree500JSONResponse struct{ N500JSONResponse }

func (response ReadSimpleTree500JSONResponse) VisitReadSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type UpdateSimpleTreeRequestObject struct {
	Id   uint32 `json:"id" yaml:"id" xml:"id" bson:"id"`
	Body *UpdateSimpleTreeJSONRequestBody
}

type UpdateSimpleTreeResponseObject interface {
	VisitUpdateSimpleTreeResponse(w http.ResponseWriter) error
}

type UpdateSimpleTree200JSONResponse SimpleTreeUpdate

func (response UpdateSimpleTree200JSONResponse) VisitUpdateSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type UpdateSimpleTree400JSONResponse struct{ N400JSONResponse }

func (response UpdateSimpleTree400JSONResponse) VisitUpdateSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type UpdateSimpleTree404JSONResponse struct{ N404JSONResponse }

func (response UpdateSimpleTree404JSONResponse) VisitUpdateSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type UpdateSimpleTree409JSONResponse struct{ N409JSONResponse }

func (response UpdateSimpleTree409JSONResponse) VisitUpdateSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type UpdateSimpleTree500JSONResponse struct{ N500JSONResponse }

func (response UpdateSimpleTree500JSONResponse) VisitUpdateSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type ListSimpleTreeChildrenRequestObject struct {
	Id     uint32 `json:"id" yaml:"id" xml:"id" bson:"id"`
	Params ListSimpleTreeChildrenParams
}

type ListSimpleTreeChildrenResponseObject interface {
	VisitListSimpleTreeChildrenResponse(w http.ResponseWriter) error
}

type ListSimpleTreeChildren200JSONResponse struct {
	// CurrentPage Page number (1-based)
	CurrentPage int `json:"current_page" yaml:"current_page" xml:"current_page" bson:"current_page"`

	// Data List of items
	Data []SimpleTreeList `json:"data" yaml:"data" xml:"data" bson:"data"`

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

func (response ListSimpleTreeChildren200JSONResponse) VisitListSimpleTreeChildrenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ListSimpleTreeChildren400JSONResponse struct{ N400JSONResponse }

func (response ListSimpleTreeChildren400JSONResponse) VisitListSimpleTreeChildrenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type ListSimpleTreeChildren404JSONResponse struct{ N404JSONResponse }

func (response ListSimpleTreeChildren404JSONResponse) VisitListSimpleTreeChildrenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type ListSimpleTreeChildren409JSONResponse struct{ N409JSONResponse }

func (response ListSimpleTreeChildren409JSONResponse) VisitListSimpleTreeChildrenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type ListSimpleTreeChildren500JSONResponse struct{ N500JSONResponse }

func (response ListSimpleTreeChildren500JSONResponse) VisitListSimpleTreeChildrenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type ReadSimpleTreeParentRequestObject struct {
	Id uint32 `json:"id"`
}

type ReadSimpleTreeParentResponseObject interface {
	VisitReadSimpleTreeParentResponse(w http.ResponseWriter) error
}

type ReadSimpleTreeParent200JSONResponse SimpleTreeParentRead

func (response ReadSimpleTreeParent200JSONResponse) VisitReadSimpleTreeParentResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ReadSimpleTreeParent400JSONResponse struct{ N400JSONResponse }

func (response ReadSimpleTreeParent400JSONResponse) VisitReadSimpleTreeParentResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type ReadSimpleTreeParent404JSONResponse struct{ N404JSONResponse }

func (response ReadSimpleTreeParent404JSONResponse) VisitReadSimpleTreeParentResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type ReadSimpleTreeParent409JSONResponse struct{ N409JSONResponse }

func (response ReadSimpleTreeParent409JSONResponse) VisitReadSimpleTreeParentResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type ReadSimpleTreeParent500JSONResponse struct{ N500JSONResponse }

func (response ReadSimpleTreeParent500JSONResponse) VisitReadSimpleTreeParentResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type RestoreSimpleTreeRequestObject struct {
	Id uint32 `json:"id"`
}

type RestoreSimpleTreeResponseObject interface {
	VisitRestoreSimpleTreeResponse(w http.ResponseWriter) error
}

type RestoreSimpleTree204Response struct {
}

func (response RestoreSimpleTree204Response) VisitRestoreSimpleTreeResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type RestoreSimpleTree400JSONResponse struct{ N400JSONResponse }

func (response RestoreSimpleTree400JSONResponse) VisitRestoreSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type RestoreSimpleTree404JSONResponse struct{ N404JSONResponse }

func (response RestoreSimpleTree404JSONResponse) VisitRestoreSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type RestoreSimpleTree409JSONResponse struct{ N409JSONResponse }

func (response RestoreSimpleTree409JSONResponse) VisitRestoreSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type RestoreSimpleTree500JSONResponse struct{ N500JSONResponse }

func (response RestoreSimpleTree500JSONResponse) VisitRestoreSimpleTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// List SimpleTrees
	// (GET /simple-tree)
	ListSimpleTree(ctx context.Context, request ListSimpleTreeRequestObject) (ListSimpleTreeResponseObject, error)
	// Create a new SimpleTree
	// (POST /simple-tree)
	CreateSimpleTree(ctx context.Context, request CreateSimpleTreeRequestObject) (CreateSimpleTreeResponseObject, error)
	// Deletes a SimpleTree by ID
	// (DELETE /simple-tree/{id})
	DeleteSimpleTree(ctx context.Context, request DeleteSimpleTreeRequestObject) (DeleteSimpleTreeResponseObject, error)
	// Find a SimpleTree by ID
	// (GET /simple-tree/{id})
	ReadSimpleTree(ctx context.Context, request ReadSimpleTreeRequestObject) (ReadSimpleTreeResponseObject, error)
	// Updates a SimpleTree
	// (PATCH /simple-tree/{id})
	UpdateSimpleTree(ctx context.Context, request UpdateSimpleTreeRequestObject) (UpdateSimpleTreeResponseObject, error)
	// List of subordinate items
	// (GET /simple-tree/{id}/children)
	ListSimpleTreeChildren(ctx context.Context, request ListSimpleTreeChildrenRequestObject) (ListSimpleTreeChildrenResponseObject, error)
	// Find the attached SimpleTree
	// (GET /simple-tree/{id}/parent)
	ReadSimpleTreeParent(ctx context.Context, request ReadSimpleTreeParentRequestObject) (ReadSimpleTreeParentResponseObject, error)
	// Restore a trashed record
	// (POST /simple-tree/{id}/restore)
	RestoreSimpleTree(ctx context.Context, request RestoreSimpleTreeRequestObject) (RestoreSimpleTreeResponseObject, error)
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

// ListSimpleTree operation middleware
func (sh *strictHandler) ListSimpleTree(ctx *gin.Context, params ListSimpleTreeParams) {
	var request ListSimpleTreeRequestObject

	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ListSimpleTree(ctx, request.(ListSimpleTreeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListSimpleTree")
	}

	response, err := handler(ctx, request)

	if err != nil {
		handleErrorResponse(ctx, err)
	} else if validResponse, ok := response.(ListSimpleTreeResponseObject); ok {
		if err := validResponse.VisitListSimpleTreeResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// CreateSimpleTree operation middleware
func (sh *strictHandler) CreateSimpleTree(ctx *gin.Context) {
	var request CreateSimpleTreeRequestObject

	var body CreateSimpleTreeJSONRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreateSimpleTree(ctx, request.(CreateSimpleTreeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateSimpleTree")
	}

	response, err := handler(ctx, request)

	if err != nil {
		handleErrorResponse(ctx, err)
	} else if validResponse, ok := response.(CreateSimpleTreeResponseObject); ok {
		if err := validResponse.VisitCreateSimpleTreeResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteSimpleTree operation middleware
func (sh *strictHandler) DeleteSimpleTree(ctx *gin.Context, id uint32, params DeleteSimpleTreeParams) {
	var request DeleteSimpleTreeRequestObject

	request.Id = id
	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteSimpleTree(ctx, request.(DeleteSimpleTreeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteSimpleTree")
	}

	response, err := handler(ctx, request)

	if err != nil {
		handleErrorResponse(ctx, err)
	} else if validResponse, ok := response.(DeleteSimpleTreeResponseObject); ok {
		if err := validResponse.VisitDeleteSimpleTreeResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// ReadSimpleTree operation middleware
func (sh *strictHandler) ReadSimpleTree(ctx *gin.Context, id uint32, params ReadSimpleTreeParams) {
	var request ReadSimpleTreeRequestObject

	request.Id = id
	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ReadSimpleTree(ctx, request.(ReadSimpleTreeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ReadSimpleTree")
	}

	response, err := handler(ctx, request)

	if err != nil {
		handleErrorResponse(ctx, err)
	} else if validResponse, ok := response.(ReadSimpleTreeResponseObject); ok {
		if err := validResponse.VisitReadSimpleTreeResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// UpdateSimpleTree operation middleware
func (sh *strictHandler) UpdateSimpleTree(ctx *gin.Context, id uint32) {
	var request UpdateSimpleTreeRequestObject

	request.Id = id

	var body UpdateSimpleTreeJSONRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateSimpleTree(ctx, request.(UpdateSimpleTreeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateSimpleTree")
	}

	response, err := handler(ctx, request)

	if err != nil {
		handleErrorResponse(ctx, err)
	} else if validResponse, ok := response.(UpdateSimpleTreeResponseObject); ok {
		if err := validResponse.VisitUpdateSimpleTreeResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// ListSimpleTreeChildren operation middleware
func (sh *strictHandler) ListSimpleTreeChildren(ctx *gin.Context, id uint32, params ListSimpleTreeChildrenParams) {
	var request ListSimpleTreeChildrenRequestObject

	request.Id = id
	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ListSimpleTreeChildren(ctx, request.(ListSimpleTreeChildrenRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListSimpleTreeChildren")
	}

	response, err := handler(ctx, request)

	if err != nil {
		handleErrorResponse(ctx, err)
	} else if validResponse, ok := response.(ListSimpleTreeChildrenResponseObject); ok {
		if err := validResponse.VisitListSimpleTreeChildrenResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// ReadSimpleTreeParent operation middleware
func (sh *strictHandler) ReadSimpleTreeParent(ctx *gin.Context, id uint32) {
	var request ReadSimpleTreeParentRequestObject

	request.Id = id

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ReadSimpleTreeParent(ctx, request.(ReadSimpleTreeParentRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ReadSimpleTreeParent")
	}

	response, err := handler(ctx, request)

	if err != nil {
		handleErrorResponse(ctx, err)
	} else if validResponse, ok := response.(ReadSimpleTreeParentResponseObject); ok {
		if err := validResponse.VisitReadSimpleTreeParentResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// RestoreSimpleTree operation middleware
func (sh *strictHandler) RestoreSimpleTree(ctx *gin.Context, id uint32) {
	var request RestoreSimpleTreeRequestObject

	request.Id = id

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.RestoreSimpleTree(ctx, request.(RestoreSimpleTreeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "RestoreSimpleTree")
	}

	response, err := handler(ctx, request)

	if err != nil {
		handleErrorResponse(ctx, err)
	} else if validResponse, ok := response.(RestoreSimpleTreeResponseObject); ok {
		if err := validResponse.VisitRestoreSimpleTreeResponse(ctx.Writer); err != nil {
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
