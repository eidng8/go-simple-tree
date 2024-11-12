package main

import (
	"github.com/eidng8/go-ent/paginate"

	"github.com/eidng8/go-simple-tree/ent"
)

func newSimpleTreeFromEnt(eaa *ent.SimpleTree) *SimpleTree {
	aa := SimpleTree{}
	aa.Id = eaa.ID
	aa.Name = eaa.Name
	if eaa.ParentID != nil {
		val := *eaa.ParentID
		aa.ParentId = &val
	}
	aa.CreatedAt = eaa.CreatedAt
	aa.UpdatedAt = eaa.UpdatedAt
	if eaa.Edges.Parent != nil {
		aa.Parent = newSimpleTreeFromEnt(eaa.Edges.Parent)
	}
	if eaa.Edges.Children != nil {
		children := make([]SimpleTree, len(eaa.Edges.Children))
		for i, child := range eaa.Edges.Children {
			children[i] = *newSimpleTreeFromEnt(child)
		}
		aa.Children = &children
	}
	return &aa
}

func newSimpleTreeListFromEnt(eaa *ent.SimpleTree) SimpleTreeList {
	aa := SimpleTreeList{}
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

func mapSimpleTreeListFromEnt(array []*ent.SimpleTree) []SimpleTreeList {
	data := make([]SimpleTreeList, len(array))
	for i, row := range array {
		data[i] = newSimpleTreeListFromEnt(row)
	}
	return data
}

func mapPage[T ListSimpleTree200JSONResponse | ListSimpleTreeChildren200JSONResponse](
	page *paginate.PaginatedList[ent.SimpleTree],
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
		Data:         mapSimpleTreeListFromEnt(page.Data),
	}
}
