package main

import (
	"github.com/eidng8/go-ent/paginate"

	"github.com/eidng8/go-simple-tree/ent"
)

func newItemFromEnt(eaa *ent.Item) *Item {
	aa := Item{}
	aa.Id = eaa.ID
	aa.Name = eaa.Name
	if eaa.ParentID != nil {
		val := *eaa.ParentID
		aa.ParentId = &val
	}
	aa.CreatedAt = eaa.CreatedAt
	aa.UpdatedAt = eaa.UpdatedAt
	if eaa.Edges.Parent != nil {
		aa.Parent = newItemFromEnt(eaa.Edges.Parent)
	}
	if eaa.Edges.Children != nil {
		children := make([]Item, len(eaa.Edges.Children))
		for i, child := range eaa.Edges.Children {
			children[i] = *newItemFromEnt(child)
		}
		aa.Children = &children
	}
	return &aa
}

func newItemListFromEnt(eaa *ent.Item) ItemList {
	aa := ItemList{}
	aa.Id = eaa.ID
	aa.Name = eaa.Name
	if eaa.ParentID != nil {
		val := *eaa.ParentID
		aa.ParentId = &val
	}
	aa.CreatedAt = eaa.CreatedAt
	aa.UpdatedAt = eaa.UpdatedAt
	return aa
}

func mapItemListFromEnt(array []*ent.Item) []ItemList {
	data := make([]ItemList, len(array))
	for i, row := range array {
		data[i] = newItemListFromEnt(row)
	}
	return data
}

func mapPage[T ListItem200JSONResponse | ListItemChildren200JSONResponse](
	page *paginate.PaginatedList[ent.Item],
) T {
	return T{
		CurrentPage:  page.CurrentPage,
		FirstPageUrl: page.FirstPageUrl,
		From:         page.From,
		LastPage:     page.LastPage,
		LastPageUrl:  page.LastPageUrl,
		NextPageUrl:  page.NextPageUrl,
		Path:         page.Path,
		PerPage:      page.PerPage,
		PrevPageUrl:  page.PrevPageUrl,
		To:           page.To,
		Total:        page.Total,
		Data:         mapItemListFromEnt(page.Data),
	}
}
