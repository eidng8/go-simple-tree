package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/eidng8/go-simple-tree/ent"
)

type UpdateItem201JSONResponse ItemCreate

func (response UpdateItem201JSONResponse) VisitUpdateItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	return json.NewEncoder(w).Encode(response)
}

// UpdateItem Updates a Item
// (PATCH /simple-tree/{id})
func (s Server) UpdateItem(
	ctx context.Context, request UpdateItemRequestObject,
) (UpdateItemResponseObject, error) {
	tx, err := s.EC.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if nil != err {
			_ = tx.Rollback()
		}
	}()
	ac := tx.Item.UpdateOneID(request.Id)
	if request.Body.Name != nil {
		ac.SetName(*request.Body.Name)
	}
	if request.Body.ParentId != nil {
		if *request.Body.ParentId == request.Id {
			return nil, ent.NewValidationError(
				"parent_id", fmt.Errorf("ParentId cannot be equal to self"),
			)
		}
		ac.SetParentID(*request.Body.ParentId)
	}
	var aa *ent.Item
	aa, err = ac.Save(ctx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	var pid *uint32
	if nil == aa.ParentID {
		pid = nil
	} else {
		val := *aa.ParentID
		pid = &val
	}
	return UpdateItem201JSONResponse{
		Id:        aa.ID,
		ParentId:  pid,
		Name:      aa.Name,
		CreatedAt: aa.CreatedAt,
		UpdatedAt: aa.UpdatedAt,
	}, nil
}
