package main

import (
	"context"
	"unicode/utf8"

	"github.com/eidng8/go-ent/paginate"
	"github.com/eidng8/go-url"
	"github.com/gin-gonic/gin"

	"github.com/eidng8/go-simple-tree/ent"
	"github.com/eidng8/go-simple-tree/ent/item"
)

// ListItemChildren List attached Children
// (GET /simple-tree/{id}/children)
func (s Server) ListItemChildren(
	ctx context.Context, request ListItemChildrenRequestObject,
) (ListItemChildrenResponseObject, error) {
	gc := ctx.(*gin.Context)
	query := s.EC.Item.Query().Order(item.ByID())
	applyChildrenNameFilter(request, query)
	id := request.Id
	if nil != request.Params.Recurse && *request.Params.Recurse {
		return getDescendants(gc, ctx, query, id)
	}
	return getPage(gc, ctx, query, id)
}

func getPage(
	gc *gin.Context, qc context.Context, query *ent.ItemQuery, id uint32,
) (ListItemChildrenResponseObject, error) {
	query.Where(item.HasParentWith(item.ID(id)))
	pageParams := paginate.GetPaginationParams(gc)
	areas, err := paginate.GetPage[ent.Item](gc, qc, query, pageParams)
	if err != nil {
		return nil, err
	}
	return mapPage[ListItemChildren200JSONResponse](areas), nil
}

func getDescendants(
	gc *gin.Context, qc context.Context, query *ent.ItemQuery, id uint32,
) (ListItemChildrenResponseObject, error) {
	areas, err := query.QueryChildrenRecursive(id).All(qc)
	if err != nil {
		return nil, err
	}
	count := len(areas)
	req := gc.Request
	u := paginate.UrlWithoutPageParams(req)
	u = url.WithQueryParam(*u, "recurse", "1")
	return ListItemChildren200JSONResponse{
		CurrentPage:  1,
		FirstPageUrl: u.String(),
		From:         1,
		LastPage:     1,
		LastPageUrl:  "",
		NextPageUrl:  "",
		Path:         url.RequestBaseUrl(req).String(),
		PerPage:      count,
		PrevPageUrl:  "",
		To:           count,
		Total:        count,
		Data:         mapItemListFromEnt(areas),
	}, nil
}

func applyChildrenNameFilter(
	request ListItemChildrenRequestObject, query *ent.ItemQuery,
) {
	name := request.Params.Name
	if name != nil && utf8.RuneCountInString(*name) > 1 {
		query.Where(item.NameHasPrefix(*name))
	}
}
