package main

import (
	"context"

	"github.com/eidng8/go-simple-tree/ent"
	"github.com/eidng8/go-simple-tree/ent/item"
)

// ReadItemParent Find a Item by ID
// (GET /simple-tree/{id}/parent)
func (s Server) ReadItemParent(
	ctx context.Context, request ReadItemParentRequestObject,
) (ReadItemParentResponseObject, error) {
	area, err := s.EC.Item.Query().Where(item.ID(uint32(request.Id))).
		WithParent().Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ReadItemParent404JSONResponse{}, nil
		}
		return nil, err
	}
	if nil == area || nil == area.Edges.Parent {
		return ReadItemParent404JSONResponse{}, nil
	}
	return newReadItemParent200JSONResponseFromEnt(area.Edges.Parent), nil
}

func newReadItemParent200JSONResponseFromEnt(
	eaa *ent.Item,
) ReadItemParent200JSONResponse {
	aar := ReadItemParent200JSONResponse{}
	aar.Id = eaa.ID
	aar.Name = eaa.Name
	if eaa.ParentID != nil {
		val := *eaa.ParentID
		aar.ParentId = &val
	}
	aar.CreatedAt = eaa.CreatedAt
	aar.UpdatedAt = eaa.UpdatedAt
	return aar
}
