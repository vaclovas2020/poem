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
	poemsMap, err := srv.cmd.getPoems(db, req.Category)
	if err != nil {
		return nil, err
	}
	result = new(poems.PoemsResponse)
	result.Success = true
	result.Poems = poemsMap
	return result, nil
}

/* get poems list from mysql database */
func (p *poemsServerCmd) getPoems(db *sql.DB, category string) (result map[int32]*poems.Poem, err error) {
	query, err := runtime.QueryDb(db, category, func(db *sql.DB, category string) (*sql.Rows, error) {
		return db.Query("SELECT poem_poems.poem_id, poem_poems.title, poem_poems.text FROM poem_poems INNER JOIN poem_categories ON poem_poems.category_id = poem_categories.category_id WHERE poem_categories.slug = ?;", category)
	})
	if err != nil {
		return nil, err
	}
	result = make(map[int32]*poems.Poem)
	for query.Next() {
		var id *int
		var poem *poems.Poem
		query.Scan(id, poem.Title, poem.Text)
		result[int32(*id)] = poem
	}
	return result, nil
}
