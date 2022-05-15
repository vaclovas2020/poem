/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

package poemsserver

import (
	"context"
	"database/sql"

	"webimizer.dev/poem/poems"
	"webimizer.dev/poem/runtime"
)

/* gRPC GetCategories */
func (srv *poemsServer) GetCategories(ctx context.Context, req *poems.CategoriesRequest) (result *poems.CategoriesResponse, err error) {
	db, err := srv.cmd.openDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	categoriesMap, err := srv.cmd.getCategories(db, req)
	if err != nil {
		return nil, err
	}
	result = new(poems.CategoriesResponse)
	result.Success = true
	result.Categories = categoriesMap
	return result, nil
}

/* get categories list from mysql database */
func (p *poemsServerCmd) getCategories(db *sql.DB, req *poems.CategoriesRequest) (result map[int32]*poems.Category, err error) {
	query, err := runtime.QueryDb(db, req, func(db *sql.DB, req *poems.CategoriesRequest) (*sql.Rows, error) {
		return db.Query("SELECT poem_categories.category_id, poem_categories.name, poem_categories.slug FROM poem_categories WHERE poem_categories.status = ? AND poem_categories.user_id = ?;", req.Status.String(), req.UserId)
	})
	if err != nil {
		return nil, err
	}
	result = make(map[int32]*poems.Category)
	for query.Next() {
		var id int
		var name string
		var slug string
		query.Scan(&id, &name, &slug)
		result[int32(id)] = &poems.Category{Name: name, Slug: slug}
	}
	return result, nil
}
