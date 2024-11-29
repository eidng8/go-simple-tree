package main

import (
	"context"
	"net/http"
	"unicode/utf8"

	"github.com/eidng8/go-ent/paginate"
	"github.com/eidng8/go-url"
	"github.com/gin-gonic/gin"

	"github.com/eidng8/go-simple-tree/ent"
	"github.com/eidng8/go-simple-tree/ent/item"
)

type ListItemChildrenPaginatedResponse struct {
	*paginate.PaginatedList[ent.Item]
}

func (response ListItemChildrenPaginatedResponse) VisitListItemChildrenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	return json.NewEncoder(w).Encode(response)
}

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
		return s.getDescendants(gc, ctx, query, id)
	}
	return s.getChildrenPage(gc, ctx, query, id)
}

func (s Server) getChildrenPage(
	gc *gin.Context, qc context.Context, query *ent.ItemQuery, id uint32,
) (ListItemChildrenResponseObject, error) {
	query.Where(item.HasParentWith(item.ID(id)))
	paginator := paginate.Paginator[ent.Item, ent.ItemQuery]{
		BaseUrl:  s.BaseURL,
		Query:    query,
		GinCtx:   gc,
		QueryCtx: qc,
	}
	areas, err := paginator.GetPage()
	if err != nil {
		return nil, err
	}
	return ListItemChildrenPaginatedResponse{PaginatedList: areas}, nil
}

func (s Server) getDescendants(
	gc *gin.Context, qc context.Context, query *ent.ItemQuery, id uint32,
) (ListItemChildrenResponseObject, error) {
	areas, err := query.QueryChildrenRecursive(id).All(qc)
	if err != nil {
		return nil, err
	}
	count := len(areas)
	req := gc.Request
	paginator := paginate.Paginator[ent.Item, ent.ItemQuery]{
		BaseUrl:  s.BaseURL,
		Query:    nil,
		GinCtx:   gc,
		QueryCtx: qc,
	}
	u := paginator.UrlWithoutPageParams()
	u = url.WithQueryParam(*u, "recurse", "1")
	return ListItemChildrenPaginatedResponse{
		PaginatedList: &paginate.PaginatedList[ent.Item]{
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
			Data:         areas,
		},
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
