package main

import (
	"context"
	"net/http"
	"unicode/utf8"

	"github.com/eidng8/go-ent/softdelete"
	"github.com/gin-gonic/gin"

	"github.com/eidng8/go-ent/paginate"

	"github.com/eidng8/go-simple-tree/ent"
	"github.com/eidng8/go-simple-tree/ent/item"
)

type ListItemPaginatedResponse struct {
	*paginate.PaginatedList[ent.Item]
}

func (response ListItemPaginatedResponse) VisitListItemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	return json.NewEncoder(w).Encode(response)
}

// ListItem List all Items
// (GET /simple-tree)
func (s Server) ListItem(
	ctx context.Context, request ListItemRequestObject,
) (ListItemResponseObject, error) {
	gc := ctx.(*gin.Context)
	query := s.EC.Item.Query().Order(item.ByID())
	qc := softdelete.NewSoftDeleteQueryContext(request.Params.Trashed, ctx)
	applyNameFilter(request, query)
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
	return ListItemPaginatedResponse{PaginatedList: areas}, nil
}

func applyNameFilter(
	request ListItemRequestObject, query *ent.ItemQuery,
) {
	name := request.Params.Name
	if name != nil && utf8.RuneCountInString(*name) > 1 {
		query.Where(item.NameHasPrefix(*name))
	}
}
