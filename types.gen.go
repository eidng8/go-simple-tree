// Package api provides primitives to interact with the openapi HTTP API.
// @formatter:off
package main

import (
	"time"

	"github.com/oapi-codegen/nullable"
)

// SimpleTree defines model for SimpleTree.
type SimpleTree struct {
	Children  *[]SimpleTree `json:"children,omitempty" yaml:"children,omitempty" xml:"children,omitempty" bson:"children,omitempty"`
	CreatedAt *time.Time    `json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at,omitempty"`
	Id        uint32        `json:"id" yaml:"id" xml:"id" bson:"id"`

	// Name Item name
	Name   string      `json:"name" yaml:"name" xml:"name" bson:"name"`
	Parent *SimpleTree `json:"parent,omitempty" yaml:"parent,omitempty" xml:"parent,omitempty" bson:"parent,omitempty"`

	// ParentId Parent record ID
	ParentId  *uint32    `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// SimpleTreeCreate defines model for SimpleTreeCreate.
type SimpleTreeCreate struct {
	CreatedAt *time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at,omitempty"`
	Id        uint32     `json:"id" yaml:"id" xml:"id" bson:"id"`

	// Name Item name
	Name string `json:"name" yaml:"name" xml:"name" bson:"name"`

	// ParentId Parent record ID
	ParentId  *uint32    `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// SimpleTreeList defines model for SimpleTreeList.
type SimpleTreeList struct {
	CreatedAt *time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at,omitempty"`
	Id        uint32     `json:"id" yaml:"id" xml:"id" bson:"id"`

	// Name Item name
	Name string `json:"name" yaml:"name" xml:"name" bson:"name"`

	// ParentId Parent record ID
	ParentId  *uint32    `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// SimpleTreeRead defines model for SimpleTreeRead.
type SimpleTreeRead struct {
	CreatedAt *time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at,omitempty"`

	// DeletedAt Date and time when the record was deleted
	DeletedAt nullable.Nullable[time.Time] `json:"deleted_at,omitempty" yaml:"deleted_at,omitempty" xml:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	Id        uint32                       `json:"id" yaml:"id" xml:"id" bson:"id"`

	// Name Item name
	Name string `json:"name" yaml:"name" xml:"name" bson:"name"`

	// ParentId Parent record ID
	ParentId  *uint32    `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// SimpleTreeUpdate defines model for SimpleTreeUpdate.
type SimpleTreeUpdate struct {
	CreatedAt *time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at,omitempty"`
	Id        uint32     `json:"id" yaml:"id" xml:"id" bson:"id"`

	// Name Item name
	Name string `json:"name" yaml:"name" xml:"name" bson:"name"`

	// ParentId Parent record ID
	ParentId  *uint32    `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// SimpleTreeParentRead defines model for SimpleTree_ParentRead.
type SimpleTreeParentRead struct {
	CreatedAt *time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at,omitempty"`
	Id        uint32     `json:"id" yaml:"id" xml:"id" bson:"id"`

	// Name Item name
	Name string `json:"name" yaml:"name" xml:"name" bson:"name"`

	// ParentId Parent record ID
	ParentId  *uint32    `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// N400 defines model for 400.
type N400 struct {
	Code   int          `json:"code" yaml:"code" xml:"code" bson:"code"`
	Errors *interface{} `json:"errors,omitempty" yaml:"errors,omitempty" xml:"errors,omitempty" bson:"errors,omitempty"`
	Status string       `json:"status" yaml:"status" xml:"status" bson:"status"`
}

// N404 defines model for 404.
type N404 struct {
	Code   int          `json:"code" yaml:"code" xml:"code" bson:"code"`
	Errors *interface{} `json:"errors,omitempty" yaml:"errors,omitempty" xml:"errors,omitempty" bson:"errors,omitempty"`
	Status string       `json:"status" yaml:"status" xml:"status" bson:"status"`
}

// N409 defines model for 409.
type N409 struct {
	Code   int          `json:"code" yaml:"code" xml:"code" bson:"code"`
	Errors *interface{} `json:"errors,omitempty" yaml:"errors,omitempty" xml:"errors,omitempty" bson:"errors,omitempty"`
	Status string       `json:"status" yaml:"status" xml:"status" bson:"status"`
}

// N500 defines model for 500.
type N500 struct {
	Code   int          `json:"code" yaml:"code" xml:"code" bson:"code"`
	Errors *interface{} `json:"errors,omitempty" yaml:"errors,omitempty" xml:"errors,omitempty" bson:"errors,omitempty"`
	Status string       `json:"status" yaml:"status" xml:"status" bson:"status"`
}

// ListSimpleTreeParams defines parameters for ListSimpleTree.
type ListSimpleTreeParams struct {
	// Page what page to render
	Page *int `form:"page,omitempty" json:"page,omitempty" yaml:"page,omitempty" xml:"page,omitempty" bson:"page,omitempty"`

	// PerPage item count to render per page
	PerPage *int `form:"per_page,omitempty" json:"per_page,omitempty" yaml:"per_page,omitempty" xml:"per_page,omitempty" bson:"per_page,omitempty"`

	// Name Name of the item
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty" bson:"name,omitempty"`

	// Trashed Whether to include trashed items
	Trashed *bool `form:"trashed,omitempty" json:"trashed,omitempty" yaml:"trashed,omitempty" xml:"trashed,omitempty" bson:"trashed,omitempty"`
}

// CreateSimpleTreeJSONBody defines parameters for CreateSimpleTree.
type CreateSimpleTreeJSONBody struct {
	// Name Item name
	Name string `json:"name" yaml:"name" xml:"name" bson:"name"`

	// ParentId Parent record ID
	ParentId *uint32 `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
}

// DeleteSimpleTreeParams defines parameters for DeleteSimpleTree.
type DeleteSimpleTreeParams struct {
	// Trashed Whether to include trashed items
	Trashed *bool `form:"trashed,omitempty" json:"trashed,omitempty" yaml:"trashed,omitempty" xml:"trashed,omitempty" bson:"trashed,omitempty"`
}

// ReadSimpleTreeParams defines parameters for ReadSimpleTree.
type ReadSimpleTreeParams struct {
	// Trashed Whether to include trashed items
	Trashed *bool `form:"trashed,omitempty" json:"trashed,omitempty" yaml:"trashed,omitempty" xml:"trashed,omitempty" bson:"trashed,omitempty"`
}

// UpdateSimpleTreeJSONBody defines parameters for UpdateSimpleTree.
type UpdateSimpleTreeJSONBody struct {
	// Name Item name
	Name *string `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty" bson:"name,omitempty"`

	// ParentId Parent record ID
	ParentId *uint32 `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
}

// ListSimpleTreeChildrenParams defines parameters for ListSimpleTreeChildren.
type ListSimpleTreeChildrenParams struct {
	// Page what page to render
	Page *int `form:"page,omitempty" json:"page,omitempty" yaml:"page,omitempty" xml:"page,omitempty" bson:"page,omitempty"`

	// PerPage item count to render per page
	PerPage *int `form:"per_page,omitempty" json:"per_page,omitempty" yaml:"per_page,omitempty" xml:"per_page,omitempty" bson:"per_page,omitempty"`

	// Name Name of the item
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty" bson:"name,omitempty"`

	// Recurse Whether to return all descendants (recurse to last leaf)
	Recurse *bool `form:"recurse,omitempty" json:"recurse,omitempty" yaml:"recurse,omitempty" xml:"recurse,omitempty" bson:"recurse,omitempty"`
}

// CreateSimpleTreeJSONRequestBody defines body for CreateSimpleTree for application/json ContentType.
type CreateSimpleTreeJSONRequestBody CreateSimpleTreeJSONBody

// UpdateSimpleTreeJSONRequestBody defines body for UpdateSimpleTree for application/json ContentType.
type UpdateSimpleTreeJSONRequestBody UpdateSimpleTreeJSONBody
