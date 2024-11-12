package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	ee "github.com/eidng8/go-ent"
	"github.com/eidng8/go-ent/simpletree"
	"github.com/eidng8/go-ent/softdelete"
	"github.com/ogen-go/ogen"

	gen "github.com/eidng8/go-simple-tree/ent"
	"github.com/eidng8/go-simple-tree/ent/intercept"
)

// BaseUri is the base URI for the OpenAPI endpoints.
const BaseUri = "/simple-tree"

// TableName is the name of the table in the database.
const TableName = "simple_tree"

type SimpleTree struct {
	ent.Schema
}

func (SimpleTree) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     TableName,
			Charset:   "utf8mb4",
			Collation: "utf8mb4_unicode_ci",
		},
		entsql.WithComments(true),
		entsql.OnDelete(entsql.Restrict),
		schema.Comment("Item table"),
	}
}

func (SimpleTree) Fields() []ent.Field {
	u2 := uint64(2)
	u255 := uint64(255)
	return append(
		[]ent.Field{
			field.Uint32("id").Unique().Immutable().Annotations(
				// adds constraints to the generated OpenAPI specification
				entoas.Schema(
					&ogen.Schema{
						Type:    "integer",
						Format:  "uint32",
						Minimum: ogen.Num("1"),
						Maximum: ogen.Num("4294967295"),
					},
				),
			),
			field.String("name").NotEmpty().MinLen(2).MaxLen(255).
				Comment("Item name").Annotations(
				// adds constraints to the generated OpenAPI specification
				entoas.Schema(
					&ogen.Schema{
						Type:        "string",
						MinLength:   &u2,
						MaxLength:   &u255,
						Description: "Item name",
					},
				),
			),
		},
		ee.Timestamps()...,
	)
}

func (SimpleTree) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Comment out these when running `go generate` for the first time
		softdelete.Mixin{},
		simpletree.ParentU32Mixin[SimpleTree]{},
	}
}

func (SimpleTree) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		// Comment out this when running `go generate` for the first time
		softdelete.Interceptor(intercept.NewQuery),
	}
}

func (SimpleTree) Hooks() []ent.Hook {
	return []ent.Hook{
		// Comment out this when running `go generate` for the first time
		softdelete.Mutator[*gen.Client](),
	}
}
