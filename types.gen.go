// Package api provides primitives to interact with the openapi HTTP API.
// @formatter:off
package main

import (
	"time"

	"github.com/oapi-codegen/nullable"
)

// Item defines model for Item.
type Item struct {
	Children  *[]Item    `json:"children,omitempty" yaml:"children,omitempty" xml:"children,omitempty" bson:"children,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at,omitempty"`
	Id        uint32     `json:"id" yaml:"id" xml:"id" bson:"id"`

	// Name Item name
	Name   string `json:"name" yaml:"name" xml:"name" bson:"name"`
	Parent *Item  `json:"parent,omitempty" yaml:"parent,omitempty" xml:"parent,omitempty" bson:"parent,omitempty"`

	// ParentId Parent record ID
	ParentId  *uint32    `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// ItemCreate defines model for ItemCreate.
type ItemCreate struct {
	CreatedAt *time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at,omitempty"`
	Id        uint32     `json:"id" yaml:"id" xml:"id" bson:"id"`

	// Name Item name
	Name string `json:"name" yaml:"name" xml:"name" bson:"name"`

	// ParentId Parent record ID
	ParentId  *uint32    `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// ItemList defines model for ItemList.
type ItemList struct {
	CreatedAt *time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at,omitempty"`
	Id        uint32     `json:"id" yaml:"id" xml:"id" bson:"id"`

	// Name Item name
	Name string `json:"name" yaml:"name" xml:"name" bson:"name"`

	// ParentId Parent record ID
	ParentId  *uint32    `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// ItemRead defines model for ItemRead.
type ItemRead struct {
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

// ItemUpdate defines model for ItemUpdate.
type ItemUpdate struct {
	CreatedAt *time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at,omitempty"`
	Id        uint32     `json:"id" yaml:"id" xml:"id" bson:"id"`

	// Name Item name
	Name string `json:"name" yaml:"name" xml:"name" bson:"name"`

	// ParentId Parent record ID
	ParentId  *uint32    `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// ItemParentRead defines model for Item_ParentRead.
type ItemParentRead struct {
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

// ListItemParams defines parameters for ListItem.
type ListItemParams struct {
	// Page what page to render
	Page *int `form:"page,omitempty" json:"page,omitempty" yaml:"page,omitempty" xml:"page,omitempty" bson:"page,omitempty"`

	// PerPage item count to render per page
	PerPage *int `form:"per_page,omitempty" json:"per_page,omitempty" yaml:"per_page,omitempty" xml:"per_page,omitempty" bson:"per_page,omitempty"`

	// Name Name of the item
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty" bson:"name,omitempty"`

	// Trashed Whether to include trashed items
	Trashed *bool `form:"trashed,omitempty" json:"trashed,omitempty" yaml:"trashed,omitempty" xml:"trashed,omitempty" bson:"trashed,omitempty"`
}

// CreateItemJSONBody defines parameters for CreateItem.
type CreateItemJSONBody struct {
	// Name Item name
	Name string `json:"name" yaml:"name" xml:"name" bson:"name"`

	// ParentId Parent record ID
	ParentId *uint32 `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
}

// DeleteItemParams defines parameters for DeleteItem.
type DeleteItemParams struct {
	// Trashed Whether to include trashed items
	Trashed *bool `form:"trashed,omitempty" json:"trashed,omitempty" yaml:"trashed,omitempty" xml:"trashed,omitempty" bson:"trashed,omitempty"`
}

// ReadItemParams defines parameters for ReadItem.
type ReadItemParams struct {
	// Trashed Whether to include trashed items
	Trashed *bool `form:"trashed,omitempty" json:"trashed,omitempty" yaml:"trashed,omitempty" xml:"trashed,omitempty" bson:"trashed,omitempty"`
}

// UpdateItemJSONBody defines parameters for UpdateItem.
type UpdateItemJSONBody struct {
	// Name Item name
	Name *string `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty" bson:"name,omitempty"`

	// ParentId Parent record ID
	ParentId *uint32 `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
}

// ListItemChildrenParams defines parameters for ListItemChildren.
type ListItemChildrenParams struct {
	// Page what page to render
	Page *int `form:"page,omitempty" json:"page,omitempty" yaml:"page,omitempty" xml:"page,omitempty" bson:"page,omitempty"`

	// PerPage item count to render per page
	PerPage *int `form:"per_page,omitempty" json:"per_page,omitempty" yaml:"per_page,omitempty" xml:"per_page,omitempty" bson:"per_page,omitempty"`

	// Name Name of the item
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty" bson:"name,omitempty"`

	// Recurse Whether to return all descendants (recurse to last leaf)
	Recurse *bool `form:"recurse,omitempty" json:"recurse,omitempty" yaml:"recurse,omitempty" xml:"recurse,omitempty" bson:"recurse,omitempty"`
}

// CreateItemJSONRequestBody defines body for CreateItem for application/json ContentType.
type CreateItemJSONRequestBody CreateItemJSONBody

// UpdateItemJSONRequestBody defines body for UpdateItem for application/json ContentType.
type UpdateItemJSONRequestBody UpdateItemJSONBody
