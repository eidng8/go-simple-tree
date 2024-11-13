package main

import (
	"context"

	"github.com/eidng8/go-ent/softdelete"

	"github.com/eidng8/go-simple-tree/ent"
)

func (s Server) RestoreItem(
	ctx context.Context, request RestoreItemRequestObject,
) (RestoreItemResponseObject, error) {
	qc := softdelete.IncludeTrashed(ctx)
	id := request.Id
	tx, err := s.EC.Tx(qc)
	if err != nil {
		return nil, err
	}
	defer func() {
		if nil != err {
			_ = tx.Rollback()
		}
	}()
	err = tx.Item.UpdateOneID(id).ClearDeletedAt().Exec(qc)
	if err != nil {
		if ent.IsNotFound(err) {
			return RestoreItem404JSONResponse{}, nil
		}
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return RestoreItem204Response{}, nil
}
