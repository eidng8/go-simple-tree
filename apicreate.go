package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/eidng8/go-simple-tree/ent"
)

type CreateItem201JSONResponse ItemCreate

func (response CreateItem201JSONResponse) VisitCreateItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	return json.NewEncoder(w).Encode(response)
}

// CreateItem Create a new Item
// (POST /simple-tree)
func (s Server) CreateItem(
	ctx context.Context, request CreateItemRequestObject,
) (CreateItemResponseObject, error) {
	tx, err := s.EC.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if nil != err {
			_ = tx.Rollback()
		}
	}()
	ac := tx.Item.Create()
	ac.SetName(request.Body.Name)
	if request.Body.ParentId != nil {
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
	return CreateItem201JSONResponse{
		Id:        aa.ID,
		ParentId:  pid,
		Name:      aa.Name,
		CreatedAt: aa.CreatedAt,
		UpdatedAt: aa.UpdatedAt,
	}, nil
}
