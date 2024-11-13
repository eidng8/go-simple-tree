package main

import (
	"context"
	"unicode/utf8"

	"github.com/eidng8/go-ent/softdelete"
	"github.com/gin-gonic/gin"

	"github.com/eidng8/go-ent/paginate"

	"github.com/eidng8/go-simple-tree/ent"
	"github.com/eidng8/go-simple-tree/ent/item"
)

// ListItem List all Items
// (GET /simple-tree)
func (s Server) ListItem(
	ctx context.Context, request ListItemRequestObject,
) (ListItemResponseObject, error) {
	c := ctx.(*gin.Context)
	pageParams := paginate.GetPaginationParams(c)
	query := s.EC.Item.Query().Order(item.ByID())
	qc := softdelete.NewSoftDeleteQueryContext(request.Params.Trashed, ctx)
	applyNameFilter(request, query)
	areas, err := paginate.GetPage[ent.Item](c, qc, query, pageParams)
	if err != nil {
		return nil, err
	}
	return mapPage[ListItem200JSONResponse](areas), nil
}

func applyNameFilter(
	request ListItemRequestObject, query *ent.ItemQuery,
) {
	name := request.Params.Name
	if name != nil && utf8.RuneCountInString(*name) > 1 {
		query.Where(item.NameHasPrefix(*name))
	}
}
