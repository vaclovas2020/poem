/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

package adminserver

import (
	"context"
	"database/sql"

	"webimizer.dev/poem/admin"
	"webimizer.dev/poem/runtime"
)

/* gRPC AddCategory */
func (srv *adminServer) AddCategory(_ context.Context, category *admin.AdminCategory) (response *admin.CategoryResponse, err error) {
	db, err := srv.cmd.openDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	_, err = runtime.ExecDb(db, category, func(db *sql.DB, category *admin.AdminCategory) (sql.Result, error) {
		return db.Exec("INSERT INTO `poem_categories`(name,slug,status,user_id) VALUES (?,?,?,?);", category.Name, category.Slug, category.Status.String(), category.UserId)
	})
	if err != nil {
		return nil, err
	}
	response = &admin.CategoryResponse{Success: true, Category: category}
	return response, nil
}
