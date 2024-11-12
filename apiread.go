package main

import (
	"context"

	"github.com/eidng8/go-ent/softdelete"
	"github.com/oapi-codegen/nullable"

	"github.com/eidng8/go-simple-tree/ent"
	"github.com/eidng8/go-simple-tree/ent/simpletree"
)

// ReadSimpleTree Find a SimpleTree by ID
// (GET /simple-tree/{id})
func (s Server) ReadSimpleTree(
	ctx context.Context, request ReadSimpleTreeRequestObject,
) (ReadSimpleTreeResponseObject, error) {
	qc := softdelete.NewSoftDeleteQueryContext(request.Params.Trashed, ctx)
	area, err := s.EC.SimpleTree.Query().
		Where(simpletree.ID(request.Id)).Only(qc)
	if err != nil {
		if ent.IsNotFound(err) {
			return ReadSimpleTree404JSONResponse{}, nil
		}
		return nil, err
	}
	return newReadSimpleTree200JSONResponseFromEnt(area), nil
}

func newReadSimpleTree200JSONResponseFromEnt(eaa *ent.SimpleTree) ReadSimpleTree200JSONResponse {
	aa := ReadSimpleTree200JSONResponse{}
	aa.Id = eaa.ID
	aa.Name = eaa.Name
	if eaa.ParentID != nil {
		val := *eaa.ParentID
		aa.ParentId = &val
	}
	if eaa.DeletedAt != nil {
		aa.DeletedAt = nullable.NewNullableWithValue(*eaa.DeletedAt)
	}
	aa.CreatedAt = eaa.CreatedAt
	aa.UpdatedAt = eaa.UpdatedAt
	return aa
}
