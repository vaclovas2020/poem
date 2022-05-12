/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

package poemsserver

import (
	"context"

	"webimizer.dev/poem/poems"
)

/* gRPC GetCategories */
func (srv *poemsServer) GetCategories(ctx context.Context, req *poems.CategoriesRequest) (result *poems.CategoriesResponse, err error) {
	categoriesMap, err := srv.cmd.getCategories(req.Status)
	if err != nil {
		return nil, err
	}
	result = new(poems.CategoriesResponse)
	result.Success = true
	result.Categories = categoriesMap
	return result, nil
}

/* get categories list from mysql database */
func (p *poemsServerCmd) getCategories(status string) (result map[int32]*poems.Category, err error) {
	db, err := p.openDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query, err := db.Query("SELECT  poem_categories.category_id, poem_categories.name, poem_categories.slug FROM poem_categories WHERE poem_categories.status = ?;", status)
	if err != nil {
		return nil, err
	}
	result = make(map[int32]*poems.Category)
	for query.Next() {
		var id *int
		var category *poems.Category
		query.Scan(id, category.Name, category.Slug)
		result[int32(*id)] = category
	}
	return result, nil
}
