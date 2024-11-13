package main

import (
	"context"

	"github.com/eidng8/go-ent/softdelete"

	"github.com/eidng8/go-simple-tree/ent"
)

// DeleteItem Deletes a Item by ID
// (DELETE /simple-tree/{id})
func (s Server) DeleteItem(
	ctx context.Context, request DeleteItemRequestObject,
) (DeleteItemResponseObject, error) {
	qc := softdelete.NewSoftDeleteQueryContext(request.Params.Trashed, ctx)
	tx, err := s.EC.Tx(qc)
	if err != nil {
		return nil, err
	}
	defer func() {
		if nil != err {
			_ = tx.Rollback()
		}
	}()
	if err = tx.Item.DeleteOneID(request.Id).Exec(qc); err != nil {
		if ent.IsNotFound(err) {
			return DeleteItem404JSONResponse{}, nil
		}
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return DeleteItem204Response{}, nil
}
