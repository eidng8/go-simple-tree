package main

import (
	"context"
	"unicode/utf8"

	"github.com/eidng8/go-ent/softdelete"
	"github.com/gin-gonic/gin"

	"github.com/eidng8/go-ent/paginate"

	"github.com/eidng8/go-simple-tree/ent"
	"github.com/eidng8/go-simple-tree/ent/simpletree"
)

// ListSimpleTree List all SimpleTrees
// (GET /simple-tree)
func (s Server) ListSimpleTree(
	ctx context.Context, request ListSimpleTreeRequestObject,
) (ListSimpleTreeResponseObject, error) {
	c := ctx.(*gin.Context)
	pageParams := paginate.GetPaginationParams(c)
	query := s.EC.SimpleTree.Query().Order(simpletree.ByID())
	qc := softdelete.NewSoftDeleteQueryContext(request.Params.Trashed, ctx)
	applyNameFilter(request, query)
	areas, err := paginate.GetPage[ent.SimpleTree](c, qc, query, pageParams)
	if err != nil {
		return nil, err
	}
	return mapPage[ListSimpleTree200JSONResponse](areas), nil
}

func applyNameFilter(
	request ListSimpleTreeRequestObject, query *ent.SimpleTreeQuery,
) {
	name := request.Params.Name
	if name != nil && utf8.RuneCountInString(*name) > 1 {
		query.Where(simpletree.NameHasPrefix(*name))
	}
}
