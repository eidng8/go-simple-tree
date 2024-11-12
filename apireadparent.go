package main

import (
	"context"

	"github.com/eidng8/go-simple-tree/ent"
	"github.com/eidng8/go-simple-tree/ent/simpletree"
)

// ReadSimpleTreeParent Find a SimpleTree by ID
// (GET /simple-tree/{id}/parent)
func (s Server) ReadSimpleTreeParent(
	ctx context.Context, request ReadSimpleTreeParentRequestObject,
) (ReadSimpleTreeParentResponseObject, error) {
	area, err := s.EC.SimpleTree.Query().Where(simpletree.ID(uint32(request.Id))).
		WithParent().Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ReadSimpleTreeParent404JSONResponse{}, nil
		}
		return nil, err
	}
	if nil == area || nil == area.Edges.Parent {
		return ReadSimpleTreeParent404JSONResponse{}, nil
	}
	return newReadSimpleTreeParent200JSONResponseFromEnt(area.Edges.Parent), nil
}

func newReadSimpleTreeParent200JSONResponseFromEnt(
	eaa *ent.SimpleTree,
) ReadSimpleTreeParent200JSONResponse {
	aar := ReadSimpleTreeParent200JSONResponse{}
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
