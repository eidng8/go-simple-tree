package main

import (
	"context"

	"github.com/eidng8/go-ent/softdelete"
	"github.com/oapi-codegen/nullable"

	"github.com/eidng8/go-simple-tree/ent"
	"github.com/eidng8/go-simple-tree/ent/item"
)

// ReadItem Find a Item by ID
// (GET /simple-tree/{id})
func (s Server) ReadItem(
	ctx context.Context, request ReadItemRequestObject,
) (ReadItemResponseObject, error) {
	qc := softdelete.NewSoftDeleteQueryContext(request.Params.Trashed, ctx)
	area, err := s.EC.Item.Query().
		Where(item.ID(request.Id)).Only(qc)
	if err != nil {
		if ent.IsNotFound(err) {
			return ReadItem404JSONResponse{}, nil
		}
		return nil, err
	}
	return newReadItem200JSONResponseFromEnt(area), nil
}

func newReadItem200JSONResponseFromEnt(eaa *ent.Item) ReadItem200JSONResponse {
	aa := ReadItem200JSONResponse{}
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
