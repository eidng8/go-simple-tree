package main

import (
	"context"
	"unicode/utf8"

	"github.com/eidng8/go-ent/paginate"
	"github.com/eidng8/go-url"
	"github.com/gin-gonic/gin"

	"github.com/eidng8/go-simple-tree/ent"
	"github.com/eidng8/go-simple-tree/ent/simpletree"
)

// ListSimpleTreeChildren List attached Children
// (GET /simple-tree/{id}/children)
func (s Server) ListSimpleTreeChildren(
	ctx context.Context, request ListSimpleTreeChildrenRequestObject,
) (ListSimpleTreeChildrenResponseObject, error) {
	gc := ctx.(*gin.Context)
	query := s.EC.SimpleTree.Query().Order(simpletree.ByID())
	applyChildrenNameFilter(request, query)
	id := request.Id
	if nil != request.Params.Recurse && *request.Params.Recurse {
		return getDescendants(gc, ctx, query, id)
	}
	return getPage(gc, ctx, query, id)
}

func getPage(
	gc *gin.Context, qc context.Context, query *ent.SimpleTreeQuery, id uint32,
) (ListSimpleTreeChildrenResponseObject, error) {
	query.Where(simpletree.HasParentWith(simpletree.ID(id)))
	pageParams := paginate.GetPaginationParams(gc)
	areas, err := paginate.GetPage[ent.SimpleTree](gc, qc, query, pageParams)
	if err != nil {
		return nil, err
	}
	return mapPage[ListSimpleTreeChildren200JSONResponse](areas), nil
}

func getDescendants(
	gc *gin.Context, qc context.Context, query *ent.SimpleTreeQuery, id uint32,
) (ListSimpleTreeChildrenResponseObject, error) {
	areas, err := query.QueryChildrenRecursive(id).All(qc)
	if err != nil {
		return nil, err
	}
	count := len(areas)
	req := gc.Request
	u := paginate.UrlWithoutPageParams(req)
	u = url.WithQueryParam(*u, "recurse", "1")
	return ListSimpleTreeChildren200JSONResponse{
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
		Data:         mapSimpleTreeListFromEnt(areas),
	}, nil
}

func applyChildrenNameFilter(
	request ListSimpleTreeChildrenRequestObject, query *ent.SimpleTreeQuery,
) {
	name := request.Params.Name
	if name != nil && utf8.RuneCountInString(*name) > 1 {
		query.Where(simpletree.NameHasPrefix(*name))
	}
}
