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

/* gRPC GetPoems */
func (srv *poemsServer) GetPoems(ctx context.Context, req *poems.PoemsRequest) (result *poems.PoemsResponse, err error) {
	db, err := srv.cmd.openDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	poemsMap, err := srv.cmd.getPoems(db, req)
	if err != nil {
		return nil, err
	}
	result = new(poems.PoemsResponse)
	result.Success = true
	result.Poems = poemsMap
	return result, nil
}

/* get poems list from mysql database */
func (p *poemsServerCmd) getPoems(db *sql.DB, req *poems.PoemsRequest) (result map[int32]*poems.Poem, err error) {
	query, err := runtime.QueryDb(db, req, func(db *sql.DB, req *poems.PoemsRequest) (*sql.Rows, error) {
		return db.Query("SELECT poem_poems.poem_id, poem_poems.category_id, poem_categories.name, poem_poems.title, poem_poems.text FROM poem_poems INNER JOIN poem_categories ON poem_poems.category_id = poem_categories.category_id WHERE poem_poems.user_id = ?;", req.UserId)
	})
	if err != nil {
		return nil, err
	}
	result = make(map[int32]*poems.Poem)
	for query.Next() {
		var id int
		var title string
		var text string
		var categoryId int
		var categoryName string
		query.Scan(&id, &categoryId, &categoryName, &title, &text)
		result[int32(id)] = &poems.Poem{Title: title, Text: text, CategoryId: int32(categoryId), CategoryName: categoryName}
	}
	return result, nil
}
