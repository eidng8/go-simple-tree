package main

import (
	"context"

	"github.com/eidng8/go-ent/softdelete"

	"github.com/eidng8/go-simple-tree/ent"
)

// DeleteSimpleTree Deletes a SimpleTree by ID
// (DELETE /simple-tree/{id})
func (s Server) DeleteSimpleTree(
	ctx context.Context, request DeleteSimpleTreeRequestObject,
) (DeleteSimpleTreeResponseObject, error) {
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
	if err = tx.SimpleTree.DeleteOneID(request.Id).Exec(qc); err != nil {
		if ent.IsNotFound(err) {
			return DeleteSimpleTree404JSONResponse{}, nil
		}
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return DeleteSimpleTree204Response{}, nil
}
