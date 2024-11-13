//go:build ignore
// +build ignore

package main

import (
	"strings"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/eidng8/go-ent/paginate"
	"github.com/eidng8/go-ent/simpletree"
	"github.com/eidng8/go-ent/softdelete"
	"github.com/ogen-go/ogen"
)

// BaseUri is the base URI for the OpenAPI endpoints.
const BaseUri = "/simple-tree"

// TableName is the name of the table in the database.
const TableName = "simple_tree"

func main() {
	err := generate()
	if err != nil {
		panic(err)
	}
}

func generate() error {
	oas, err := newOasExtension()
	if err != nil {
		return err
	}
	ext := entc.Extensions(oas, &simpletree.Extension{})
	err = entc.Generate("./ent/schema", genConfig(), ext)
	if err != nil {
		return err
	}
	return nil
}

func newOasExtension() (*entoas.Extension, error) {
	return entoas.NewExtension(
		entoas.Mutations(
			func(_ *gen.Graph, s *ogen.Spec) error {
				// Comment out these when running `go generate` for the first time
				changeBaseUri(s)
				genSpec(s)
				constraintRequestBody(s.Paths)
				ep := s.Paths[BaseUri]
				op := ep.Get
				op.AddParameters(nameParam())
				simpletree.RemoveEdges(ep.Post)
				paginate.AttachTo(
					op, "Paginated list of items",
					"#/components/schemas/SimpleTreeList",
				)
				ep = s.Paths[BaseUri+"/{id}"]
				simpletree.RemoveEdges(ep.Patch)
				err := softdelete.AttachTo(
					s, BaseUri, s.Components.Schemas["SimpleTreeRead"],
					ep.Get.Parameters[0],
				)
				if err != nil {
					return err
				}
				op = s.Paths[BaseUri+"/{id}/children"].Get
				op.AddParameters(nameParam())
				op.SetSummary("List of subordinate items")
				simpletree.AttachTo(op)
				paginate.AttachTo(
					op,
					"Paginated list of subordinate items. Pagination is disabled when `recurse` is true.",
					"#/components/schemas/SimpleTreeList",
				)
				return nil
			},
		),
	)
}

func genConfig() *gen.Config {
	return &gen.Config{
		Features: []gen.Feature{
			gen.FeatureIntercept,
			gen.FeatureSnapshot,
			gen.FeatureExecQuery,
			gen.FeatureVersionedMigration,
		},
	}
}

func changeBaseUri(spec *ogen.Spec) {
	paths := make(ogen.Paths, len(spec.Paths))
	for key, path := range spec.Paths {
		nk := strings.Replace(key, "/simple-trees", BaseUri, 1)
		paths[nk] = path
	}
	spec.SetPaths(paths)
}

func genSpec(s *ogen.Spec) {
	s.Info.SetTitle("Simple tree listing API").SetVersion("0.0.1").
		SetDescription("This is an API listing hierarchical tree data")
}

func constraintRequestBody(paths ogen.Paths) {
	for _, path := range paths {
		for _, op := range []*ogen.Operation{path.Put, path.Post, path.Patch} {
			if nil == op || nil == op.RequestBody || nil == op.RequestBody.Content {
				continue
			}
			for _, param := range op.RequestBody.Content {
				if nil == param.Schema {
					continue
				}
				b := false
				param.Schema.AdditionalProperties = &ogen.AdditionalProperties{Bool: &b}
			}
		}
	}
}

func nameParam() *ogen.Parameter {
	u2 := uint64(2)
	u255 := uint64(255)
	return &ogen.Parameter{
		Name:        "name",
		In:          "query",
		Description: "Name of the item",
		Required:    false,
		Schema: &ogen.Schema{
			Type:      "string",
			MinLength: &u2,
			MaxLength: &u255,
		},
	}
}
